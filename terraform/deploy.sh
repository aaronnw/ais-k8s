#!/bin/bash

set -eo pipefail

source utils.sh
source volumes.sh

deploy_ais() {
  node_cnt=$(kubectl get nodes --no-headers | wc -l | xargs)
  if [[ ${node_cnt} -le 0 ]]; then
    print_error "kubectl does not have any nodes assigned, try 'deploy.sh all'"
  fi

  echo "🔥 Deploying AIStore on the cluster"

  cloud_provider=$(get_state_var "CLOUD_PROVIDER")

  echo "💾 Initializing persistent storage for AIS"
  pushd k8s/ 1>/dev/null
  terraform init -input=false "${cloud_provider}" 1>/dev/null
  terraform apply -input=false -auto-approve "${cloud_provider}"
  popd 1>/dev/null

  restore_persisted_volumes
  # Remove already used yamls.
  clear_persisted_volumes

  set_state_var "VOLUMES_DEPLOYED" "true"

  # Remove labels (from all nodes) if exist.
  remove_nodes_labels

  # Label nodes so they match a required selector.
  kubectl label nodes --all \
    nvidia.com/ais-target="${release_name}-ais" \
    nvidia.com/ais-proxy="${release_name}-ais-electable" \
    1>/dev/null

  # Label one of the nodes to mark it as a primary.
  primary_node=$(kubectl get nodes -o jsonpath='{.items[0].metadata.name}')
  kubectl label nodes "${primary_node}" \
    nvidia.com/ais-admin="${release_name}-ais" \
    nvidia.com/ais-initial-primary-proxy="${release_name}-ais" \
    1>/dev/null

  external_ip=$(terraform output external_ip)
  pushd ../helm/ais 1>/dev/null

  helm_args="--set tags.builtin_monitoring=false,tags.prometheus=false,aiscluster.expected_target_nodes=${node_cnt},aiscluster.skipHostIP=true,admin.enabled=true"
  if [[ -n ${wait_timeout} ]]; then
    helm_args="${helm_args} --timeout ${wait_timeout} --wait"
    echo "⏳ Waiting for the AIStore to fully start, it may take couple minutes..."
  fi

  AIS_NAME="${release_name}" \
    AIS_GATEWAY_EXTERNAL_IP="${external_ip}" \
    AIS_K8S_CLUSTER_CIDR="10.64.0.0/14" \
    AISNODE_IMAGE="aistore/aisnode:3.3-k8s" \
    KUBECTL_IMAGE="gmaltby/ais-kubectl:1" \
    EXTERNAL_VOLUMES_COUNT="$(get_state_var "DISK_CNT")" \
    STATS_NODENAME="${primary_node}" \
    HELM_ARGS="${helm_args}" \
    ./run_ais_sample.sh 1>/dev/null

  popd 1>/dev/null

  sleep 3 # Just to make sure that pods have started.

  admin_container_id=$(kubectl get pods --namespace default -l "release=${release_name},component=admin" -o jsonpath="{.items[0].metadata.name}")
  echo "✨ AIStore cluster deployed"
  echo -e "\nLIST PODS:\n  $ kubectl get pods"
  echo -e "\nACCESS ADMIN CONTAINER:\n  $ kubectl exec -it ${admin_container_id} -- /bin/bash"
  echo -e "\nSHOW CLUSTER STATUS:\n  $ kubectl exec ${admin_container_id} -- ais show cluster"

  set_state_var "AIS_DEPLOYED" "true"
}

deploy_k8s() {
  if [[ ${cloud_provider} == "aws" ]]; then
    print_error "'aws' provider is not yet supported"

    check_command aws

    # TODO: Check if `aws` is initialized with project id and region.
    aws configure
    terraform_args=(-var "node_count=${node_cnt}" -var "ais_release_name=${release_name}" -var "cluster_name=${cluster_name}")
  elif [[ ${cloud_provider} == "azure" ]]; then
    print_error "'azure' provider is not yet supported"

    terraform_args=(-var "node_count=${node_cnt}" -var "ais_release_name=${release_name}" -var "cluster_name=${cluster_name}")
  elif [[ ${cloud_provider} == "gcp" ]]; then
    check_command gcloud

    # Check if user is logged into `gcloud`.
    if [[ -z $(gcloud config list account --format "value(core.account)") ]]; then
      if [[ -n ${GOOGLE_APPLICATION_CREDENTIALS} ]]; then
        check_command jq

        gcloud auth activate-service-account "$(jq -r '.client_email' "${GOOGLE_APPLICATION_CREDENTIALS}")" --key-file="${GOOGLE_APPLICATION_CREDENTIALS}" --project="$(jq -r '.project_id' "${GOOGLE_APPLICATION_CREDENTIALS}")"
      else
        gcloud init --console-only
      fi
    fi

    # Check if project ID is set. If it is then use it as input for the terraform.
    project_id=$(gcloud config get-value core/project)
    if [[ -z ${project_id} ]]; then
      print_error "project id is not set in 'gcloud'"
    fi

    username=$(gcloud config get-value account)
    if [[ -z ${username} ]]; then
      print_error "username is not set in 'gcloud'"
    fi

    set_state_var "GKE_PROJECT_ID" "${project_id}"
    set_state_var "GKE_USERNAME" "${username}"

    terraform_args=(-var "project_id=${project_id}" -var "user=${username}" -var "node_count=${node_cnt}" -var "ais_release_name=${release_name}" -var "cluster_name=${cluster_name}")
  fi

  # Initialize terraform and download necessary plugins.
  echo "Initializing terraform cluster environment"
  terraform init -input=false "${cloud_provider}" 1>/dev/null

  # Execute terraform plan. The approved automatically as we assume that everything is correct.
  echo "🔥 Starting Kubernetes cluster (${username}/${project_id})..."
  terraform apply -input=false -auto-approve "${terraform_args[@]}" "${cloud_provider}"

  echo "🔄 Updating kubectl config..."
  if [[ ${cloud_provider} == "aws" ]]; then
    aws eks --region us-east-2 update-kubeconfig --name training-eks-sR8eLIil
  elif [[ ${cloud_provider} == "azure" ]]; then
    :
  elif [[ ${cloud_provider} == "gcp" ]]; then
    gcloud container clusters get-credentials "$(terraform output kubernetes_cluster_name)" --zone "$(terraform output zone)"
    echo "✅ kubectl configured to use '$(kubectl config current-context)' context"
  fi
}

deploy_dashboard() {
  echo "🔄 Setting up k8s dashboard..."
  kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.0-beta8/aio/deploy/recommended.yaml
  kubectl proxy &
  echo "✅ kubectl proxy started"
  echo "🌐 Visit: http://127.0.0.1:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/"
  echo "🔑 Use this token to authenticate: $(kubectl -n kube-system get secret/"$(kubectl -n kube-system get secret | grep service-controller-token | awk '{print $1}')" --template="{{.data.token}}" | base64 -D)"
}

print_help() {
  printf "%-15s\tStops K8s pods, and destroys started nodes.\n" "all"
  printf "%-15s\tOnly stops AIStore Pods so the cluster can be redeployed.\n" "ais"
  printf "%-15s\t\t\tShows this help message.\n" "--help"
}

init_state_dir
deploy_type=$1; shift

while (( "$#" )); do
  case "$1" in
    --cloud)   cloud_provider=$2; shift; shift;;
    --cloud=*) cloud_provider="${1#*=}"; shift;;

    --node-cnt)   node_cnt=$2; shift; shift;;
    --node-cnt=*) node_cnt="${1#*=}"; shift;;

    --disk-cnt)   disk_cnt=$2; shift; shift;;
    --disk-cnt=*) disk_cnt="${1#*=}"; shift;;

    --cluster-name)   cluster_name=$2; shift; shift;;
    --cluster-name=*) cluster_name="${1#*=}"; shift;;

    --wait)   wait_timeout=$2; shift; shift;;
    --wait=*) wait_timeout="${1#*=}"; shift;;

    --help) print_help; exit 0;;

    *) echo "fatal: unknown argument '$1'"; exit 1;;
  esac
done

case ${deploy_type} in
all)
  check_command terraform
  check_command kubectl
  check_command helm

  if [[ -f ${state_file} ]]; then
    print_error "state file exists, please run 'destroy.sh all' or remove it manually: 'rm -f ${state_file}'"
  fi

  select_provider
  select_node_count
  select_disk_count
  validate_cluster_name

  set_state_var "CLOUD_PROVIDER" "${cloud_provider}"
  set_state_var "NODE_CNT" "${node_cnt}"
  set_state_var "DISK_CNT" "${disk_cnt}"

  deploy_k8s
  sleep 10
  deploy_ais
  ;;
ais)
  check_command kubectl
  check_command helm

  deploy_ais
  ;;
dashboard)
  deploy_dashboard
  wait # Wait indefinitely for `kubectl proxy`.
  ;;
--help)
  print_help
  ;;
*)
  print_error "invalid deployment type: '${deploy_type}' (expected 'all', 'ais' or 'dashboard')"
  ;;
esac
