//go:build !ignore_autogenerated

/*
 * Copyright (c) 2021-2024, NVIDIA CORPORATION. All rights reserved.
 */

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/NVIDIA/aistore/cmn/cos"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AIStore) DeepCopyInto(out *AIStore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AIStore.
func (in *AIStore) DeepCopy() *AIStore {
	if in == nil {
		return nil
	}
	out := new(AIStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AIStore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AIStoreList) DeepCopyInto(out *AIStoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AIStore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AIStoreList.
func (in *AIStoreList) DeepCopy() *AIStoreList {
	if in == nil {
		return nil
	}
	out := new(AIStoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AIStoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AIStoreSpec) DeepCopyInto(out *AIStoreSpec) {
	*out = *in
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(int32)
		**out = **in
	}
	if in.HostpathPrefix != nil {
		in, out := &in.HostpathPrefix, &out.HostpathPrefix
		*out = new(string)
		**out = **in
	}
	if in.StateStorageClass != nil {
		in, out := &in.StateStorageClass, &out.StateStorageClass
		*out = new(string)
		**out = **in
	}
	if in.ConfigToUpdate != nil {
		in, out := &in.ConfigToUpdate, &out.ConfigToUpdate
		*out = new(ConfigToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.HostnameMap != nil {
		in, out := &in.HostnameMap, &out.HostnameMap
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NetAttachment != nil {
		in, out := &in.NetAttachment, &out.NetAttachment
		*out = new(string)
		**out = **in
	}
	in.ProxySpec.DeepCopyInto(&out.ProxySpec)
	in.TargetSpec.DeepCopyInto(&out.TargetSpec)
	if in.ShutdownCluster != nil {
		in, out := &in.ShutdownCluster, &out.ShutdownCluster
		*out = new(bool)
		**out = **in
	}
	if in.CleanupMetadata != nil {
		in, out := &in.CleanupMetadata, &out.CleanupMetadata
		*out = new(bool)
		**out = **in
	}
	if in.CleanupData != nil {
		in, out := &in.CleanupData, &out.CleanupData
		*out = new(bool)
		**out = **in
	}
	if in.EnablePromExporter != nil {
		in, out := &in.EnablePromExporter, &out.EnablePromExporter
		*out = new(bool)
		**out = **in
	}
	if in.ClusterDomain != nil {
		in, out := &in.ClusterDomain, &out.ClusterDomain
		*out = new(string)
		**out = **in
	}
	if in.GCPSecretName != nil {
		in, out := &in.GCPSecretName, &out.GCPSecretName
		*out = new(string)
		**out = **in
	}
	if in.AWSSecretName != nil {
		in, out := &in.AWSSecretName, &out.AWSSecretName
		*out = new(string)
		**out = **in
	}
	if in.TLSSecretName != nil {
		in, out := &in.TLSSecretName, &out.TLSSecretName
		*out = new(string)
		**out = **in
	}
	if in.TLSCertManagerIssuerName != nil {
		in, out := &in.TLSCertManagerIssuerName, &out.TLSCertManagerIssuerName
		*out = new(string)
		**out = **in
	}
	if in.AuthNSecretName != nil {
		in, out := &in.AuthNSecretName, &out.AuthNSecretName
		*out = new(string)
		**out = **in
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.DisablePodAntiAffinity != nil {
		in, out := &in.DisablePodAntiAffinity, &out.DisablePodAntiAffinity
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AIStoreSpec.
func (in *AIStoreSpec) DeepCopy() *AIStoreSpec {
	if in == nil {
		return nil
	}
	out := new(AIStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AIStoreStatus) DeepCopyInto(out *AIStoreStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AIStoreStatus.
func (in *AIStoreStatus) DeepCopy() *AIStoreStatus {
	if in == nil {
		return nil
	}
	out := new(AIStoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthConfToUpdate) DeepCopyInto(out *AuthConfToUpdate) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthConfToUpdate.
func (in *AuthConfToUpdate) DeepCopy() *AuthConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(AuthConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CksumConfToUpdate) DeepCopyInto(out *CksumConfToUpdate) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.ValidateColdGet != nil {
		in, out := &in.ValidateColdGet, &out.ValidateColdGet
		*out = new(bool)
		**out = **in
	}
	if in.ValidateWarmGet != nil {
		in, out := &in.ValidateWarmGet, &out.ValidateWarmGet
		*out = new(bool)
		**out = **in
	}
	if in.ValidateObjMove != nil {
		in, out := &in.ValidateObjMove, &out.ValidateObjMove
		*out = new(bool)
		**out = **in
	}
	if in.EnableReadRange != nil {
		in, out := &in.EnableReadRange, &out.EnableReadRange
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CksumConfToUpdate.
func (in *CksumConfToUpdate) DeepCopy() *CksumConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(CksumConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientConfToUpdate) DeepCopyInto(out *ClientConfToUpdate) {
	*out = *in
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(Duration)
		**out = **in
	}
	if in.TimeoutLong != nil {
		in, out := &in.TimeoutLong, &out.TimeoutLong
		*out = new(Duration)
		**out = **in
	}
	if in.ListObjects != nil {
		in, out := &in.ListObjects, &out.ListObjects
		*out = new(Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientConfToUpdate.
func (in *ClientConfToUpdate) DeepCopy() *ClientConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(ClientConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigToUpdate) DeepCopyInto(out *ConfigToUpdate) {
	*out = *in
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(map[string]Empty)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]Empty, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.Mirror != nil {
		in, out := &in.Mirror, &out.Mirror
		*out = new(MirrorConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.EC != nil {
		in, out := &in.EC, &out.EC
		*out = new(ECConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Log != nil {
		in, out := &in.Log, &out.Log
		*out = new(LogConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Periodic != nil {
		in, out := &in.Periodic, &out.Periodic
		*out = new(PeriodConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(TimeoutConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Client != nil {
		in, out := &in.Client, &out.Client
		*out = new(ClientConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Space != nil {
		in, out := &in.Space, &out.Space
		*out = new(SpaceConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.LRU != nil {
		in, out := &in.LRU, &out.LRU
		*out = new(LRUConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Disk != nil {
		in, out := &in.Disk, &out.Disk
		*out = new(DiskConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Rebalance != nil {
		in, out := &in.Rebalance, &out.Rebalance
		*out = new(RebalanceConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Resilver != nil {
		in, out := &in.Resilver, &out.Resilver
		*out = new(ResilverConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Cksum != nil {
		in, out := &in.Cksum, &out.Cksum
		*out = new(CksumConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Versioning != nil {
		in, out := &in.Versioning, &out.Versioning
		*out = new(VersionConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Net != nil {
		in, out := &in.Net, &out.Net
		*out = new(NetConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.FSHC != nil {
		in, out := &in.FSHC, &out.FSHC
		*out = new(FSHCConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Auth != nil {
		in, out := &in.Auth, &out.Auth
		*out = new(AuthConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Keepalive != nil {
		in, out := &in.Keepalive, &out.Keepalive
		*out = new(KeepaliveConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Downloader != nil {
		in, out := &in.Downloader, &out.Downloader
		*out = new(DownloaderConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.DSort != nil {
		in, out := &in.DSort, &out.DSort
		*out = new(DSortConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Transport != nil {
		in, out := &in.Transport, &out.Transport
		*out = new(TransportConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Memsys != nil {
		in, out := &in.Memsys, &out.Memsys
		*out = new(MemsysConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.TCB != nil {
		in, out := &in.TCB, &out.TCB
		*out = new(TCBConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.WritePolicy != nil {
		in, out := &in.WritePolicy, &out.WritePolicy
		*out = new(WritePolicyConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(ProxyConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigToUpdate.
func (in *ConfigToUpdate) DeepCopy() *ConfigToUpdate {
	if in == nil {
		return nil
	}
	out := new(ConfigToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DSortConfToUpdate) DeepCopyInto(out *DSortConfToUpdate) {
	*out = *in
	if in.DuplicatedRecords != nil {
		in, out := &in.DuplicatedRecords, &out.DuplicatedRecords
		*out = new(string)
		**out = **in
	}
	if in.MissingShards != nil {
		in, out := &in.MissingShards, &out.MissingShards
		*out = new(string)
		**out = **in
	}
	if in.EKMMalformedLine != nil {
		in, out := &in.EKMMalformedLine, &out.EKMMalformedLine
		*out = new(string)
		**out = **in
	}
	if in.EKMMissingKey != nil {
		in, out := &in.EKMMissingKey, &out.EKMMissingKey
		*out = new(string)
		**out = **in
	}
	if in.DefaultMaxMemUsage != nil {
		in, out := &in.DefaultMaxMemUsage, &out.DefaultMaxMemUsage
		*out = new(string)
		**out = **in
	}
	if in.CallTimeout != nil {
		in, out := &in.CallTimeout, &out.CallTimeout
		*out = new(Duration)
		**out = **in
	}
	if in.DSorterMemThreshold != nil {
		in, out := &in.DSorterMemThreshold, &out.DSorterMemThreshold
		*out = new(string)
		**out = **in
	}
	if in.Compression != nil {
		in, out := &in.Compression, &out.Compression
		*out = new(string)
		**out = **in
	}
	if in.SbundleMult != nil {
		in, out := &in.SbundleMult, &out.SbundleMult
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DSortConfToUpdate.
func (in *DSortConfToUpdate) DeepCopy() *DSortConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(DSortConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaemonSpec) DeepCopyInto(out *DaemonSpec) {
	*out = *in
	out.ServiceSpec = in.ServiceSpec
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(int32)
		**out = **in
	}
	if in.ContainerSecurity != nil {
		in, out := &in.ContainerSecurity, &out.ContainerSecurity
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HostPort != nil {
		in, out := &in.HostPort, &out.HostPort
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaemonSpec.
func (in *DaemonSpec) DeepCopy() *DaemonSpec {
	if in == nil {
		return nil
	}
	out := new(DaemonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskConfToUpdate) DeepCopyInto(out *DiskConfToUpdate) {
	*out = *in
	if in.DiskUtilLowWM != nil {
		in, out := &in.DiskUtilLowWM, &out.DiskUtilLowWM
		*out = new(int64)
		**out = **in
	}
	if in.DiskUtilHighWM != nil {
		in, out := &in.DiskUtilHighWM, &out.DiskUtilHighWM
		*out = new(int64)
		**out = **in
	}
	if in.DiskUtilMaxWM != nil {
		in, out := &in.DiskUtilMaxWM, &out.DiskUtilMaxWM
		*out = new(int64)
		**out = **in
	}
	if in.IostatTimeLong != nil {
		in, out := &in.IostatTimeLong, &out.IostatTimeLong
		*out = new(Duration)
		**out = **in
	}
	if in.IostatTimeShort != nil {
		in, out := &in.IostatTimeShort, &out.IostatTimeShort
		*out = new(Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskConfToUpdate.
func (in *DiskConfToUpdate) DeepCopy() *DiskConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(DiskConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DownloaderConfToUpdate) DeepCopyInto(out *DownloaderConfToUpdate) {
	*out = *in
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DownloaderConfToUpdate.
func (in *DownloaderConfToUpdate) DeepCopy() *DownloaderConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(DownloaderConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ECConfToUpdate) DeepCopyInto(out *ECConfToUpdate) {
	*out = *in
	if in.ObjSizeLimit != nil {
		in, out := &in.ObjSizeLimit, &out.ObjSizeLimit
		*out = new(int64)
		**out = **in
	}
	if in.Compression != nil {
		in, out := &in.Compression, &out.Compression
		*out = new(string)
		**out = **in
	}
	if in.SbundleMult != nil {
		in, out := &in.SbundleMult, &out.SbundleMult
		*out = new(int)
		**out = **in
	}
	if in.DataSlices != nil {
		in, out := &in.DataSlices, &out.DataSlices
		*out = new(int)
		**out = **in
	}
	if in.ParitySlices != nil {
		in, out := &in.ParitySlices, &out.ParitySlices
		*out = new(int)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.DiskOnly != nil {
		in, out := &in.DiskOnly, &out.DiskOnly
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ECConfToUpdate.
func (in *ECConfToUpdate) DeepCopy() *ECConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(ECConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Empty) DeepCopyInto(out *Empty) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Empty.
func (in *Empty) DeepCopy() *Empty {
	if in == nil {
		return nil
	}
	out := new(Empty)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FSHCConfToUpdate) DeepCopyInto(out *FSHCConfToUpdate) {
	*out = *in
	if in.TestFileCount != nil {
		in, out := &in.TestFileCount, &out.TestFileCount
		*out = new(int)
		**out = **in
	}
	if in.HardErrs != nil {
		in, out := &in.HardErrs, &out.HardErrs
		*out = new(int)
		**out = **in
	}
	if in.IOErrs != nil {
		in, out := &in.IOErrs, &out.IOErrs
		*out = new(int)
		**out = **in
	}
	if in.IOErrTime != nil {
		in, out := &in.IOErrTime, &out.IOErrTime
		*out = new(cos.Duration)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FSHCConfToUpdate.
func (in *FSHCConfToUpdate) DeepCopy() *FSHCConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(FSHCConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPConfToUpdate) DeepCopyInto(out *HTTPConfToUpdate) {
	*out = *in
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(string)
		**out = **in
	}
	if in.CertKey != nil {
		in, out := &in.CertKey, &out.CertKey
		*out = new(string)
		**out = **in
	}
	if in.ServerNameTLS != nil {
		in, out := &in.ServerNameTLS, &out.ServerNameTLS
		*out = new(string)
		**out = **in
	}
	if in.ClientCA != nil {
		in, out := &in.ClientCA, &out.ClientCA
		*out = new(string)
		**out = **in
	}
	if in.WriteBufferSize != nil {
		in, out := &in.WriteBufferSize, &out.WriteBufferSize
		*out = new(int)
		**out = **in
	}
	if in.ReadBufferSize != nil {
		in, out := &in.ReadBufferSize, &out.ReadBufferSize
		*out = new(int)
		**out = **in
	}
	if in.ClientAuthTLS != nil {
		in, out := &in.ClientAuthTLS, &out.ClientAuthTLS
		*out = new(int)
		**out = **in
	}
	if in.UseHTTPS != nil {
		in, out := &in.UseHTTPS, &out.UseHTTPS
		*out = new(bool)
		**out = **in
	}
	if in.SkipVerifyCrt != nil {
		in, out := &in.SkipVerifyCrt, &out.SkipVerifyCrt
		*out = new(bool)
		**out = **in
	}
	if in.Chunked != nil {
		in, out := &in.Chunked, &out.Chunked
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPConfToUpdate.
func (in *HTTPConfToUpdate) DeepCopy() *HTTPConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(HTTPConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeepaliveConfToUpdate) DeepCopyInto(out *KeepaliveConfToUpdate) {
	*out = *in
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(KeepaliveTrackerConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(KeepaliveTrackerConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.RetryFactor != nil {
		in, out := &in.RetryFactor, &out.RetryFactor
		*out = new(uint8)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeepaliveConfToUpdate.
func (in *KeepaliveConfToUpdate) DeepCopy() *KeepaliveConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(KeepaliveConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeepaliveTrackerConfToUpdate) DeepCopyInto(out *KeepaliveTrackerConfToUpdate) {
	*out = *in
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(Duration)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Factor != nil {
		in, out := &in.Factor, &out.Factor
		*out = new(uint8)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeepaliveTrackerConfToUpdate.
func (in *KeepaliveTrackerConfToUpdate) DeepCopy() *KeepaliveTrackerConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(KeepaliveTrackerConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LRUConfToUpdate) DeepCopyInto(out *LRUConfToUpdate) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.DontEvictTime != nil {
		in, out := &in.DontEvictTime, &out.DontEvictTime
		*out = new(Duration)
		**out = **in
	}
	if in.CapacityUpdTime != nil {
		in, out := &in.CapacityUpdTime, &out.CapacityUpdTime
		*out = new(Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LRUConfToUpdate.
func (in *LRUConfToUpdate) DeepCopy() *LRUConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(LRUConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogConfToUpdate) DeepCopyInto(out *LogConfToUpdate) {
	*out = *in
	if in.Level != nil {
		in, out := &in.Level, &out.Level
		*out = new(string)
		**out = **in
	}
	if in.ToStderr != nil {
		in, out := &in.ToStderr, &out.ToStderr
		*out = new(bool)
		**out = **in
	}
	if in.MaxSize != nil {
		in, out := &in.MaxSize, &out.MaxSize
		*out = new(cos.SizeIEC)
		**out = **in
	}
	if in.MaxTotal != nil {
		in, out := &in.MaxTotal, &out.MaxTotal
		*out = new(cos.SizeIEC)
		**out = **in
	}
	if in.FlushTime != nil {
		in, out := &in.FlushTime, &out.FlushTime
		*out = new(Duration)
		**out = **in
	}
	if in.StatsTime != nil {
		in, out := &in.StatsTime, &out.StatsTime
		*out = new(Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogConfToUpdate.
func (in *LogConfToUpdate) DeepCopy() *LogConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(LogConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemsysConfToUpdate) DeepCopyInto(out *MemsysConfToUpdate) {
	*out = *in
	if in.MinFree != nil {
		in, out := &in.MinFree, &out.MinFree
		*out = new(cos.SizeIEC)
		**out = **in
	}
	if in.DefaultBufSize != nil {
		in, out := &in.DefaultBufSize, &out.DefaultBufSize
		*out = new(cos.SizeIEC)
		**out = **in
	}
	if in.SizeToGC != nil {
		in, out := &in.SizeToGC, &out.SizeToGC
		*out = new(cos.SizeIEC)
		**out = **in
	}
	if in.HousekeepTime != nil {
		in, out := &in.HousekeepTime, &out.HousekeepTime
		*out = new(Duration)
		**out = **in
	}
	if in.MinPctTotal != nil {
		in, out := &in.MinPctTotal, &out.MinPctTotal
		*out = new(int)
		**out = **in
	}
	if in.MinPctFree != nil {
		in, out := &in.MinPctFree, &out.MinPctFree
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemsysConfToUpdate.
func (in *MemsysConfToUpdate) DeepCopy() *MemsysConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(MemsysConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MirrorConfToUpdate) DeepCopyInto(out *MirrorConfToUpdate) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Copies != nil {
		in, out := &in.Copies, &out.Copies
		*out = new(int64)
		**out = **in
	}
	if in.Burst != nil {
		in, out := &in.Burst, &out.Burst
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MirrorConfToUpdate.
func (in *MirrorConfToUpdate) DeepCopy() *MirrorConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(MirrorConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Mount) DeepCopyInto(out *Mount) {
	*out = *in
	out.Size = in.Size.DeepCopy()
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = new(string)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Mount.
func (in *Mount) DeepCopy() *Mount {
	if in == nil {
		return nil
	}
	out := new(Mount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetConfToUpdate) DeepCopyInto(out *NetConfToUpdate) {
	*out = *in
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		*out = new(HTTPConfToUpdate)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetConfToUpdate.
func (in *NetConfToUpdate) DeepCopy() *NetConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(NetConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeriodConfToUpdate) DeepCopyInto(out *PeriodConfToUpdate) {
	*out = *in
	if in.StatsTime != nil {
		in, out := &in.StatsTime, &out.StatsTime
		*out = new(Duration)
		**out = **in
	}
	if in.RetrySyncTime != nil {
		in, out := &in.RetrySyncTime, &out.RetrySyncTime
		*out = new(Duration)
		**out = **in
	}
	if in.NotifTime != nil {
		in, out := &in.NotifTime, &out.NotifTime
		*out = new(Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeriodConfToUpdate.
func (in *PeriodConfToUpdate) DeepCopy() *PeriodConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(PeriodConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyConfToUpdate) DeepCopyInto(out *ProxyConfToUpdate) {
	*out = *in
	if in.PrimaryURL != nil {
		in, out := &in.PrimaryURL, &out.PrimaryURL
		*out = new(string)
		**out = **in
	}
	if in.OriginalURL != nil {
		in, out := &in.OriginalURL, &out.OriginalURL
		*out = new(string)
		**out = **in
	}
	if in.DiscoveryURL != nil {
		in, out := &in.DiscoveryURL, &out.DiscoveryURL
		*out = new(string)
		**out = **in
	}
	if in.NonElectable != nil {
		in, out := &in.NonElectable, &out.NonElectable
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyConfToUpdate.
func (in *ProxyConfToUpdate) DeepCopy() *ProxyConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(ProxyConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RebalanceConfToUpdate) DeepCopyInto(out *RebalanceConfToUpdate) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.DestRetryTime != nil {
		in, out := &in.DestRetryTime, &out.DestRetryTime
		*out = new(Duration)
		**out = **in
	}
	if in.Compression != nil {
		in, out := &in.Compression, &out.Compression
		*out = new(string)
		**out = **in
	}
	if in.SbundleMult != nil {
		in, out := &in.SbundleMult, &out.SbundleMult
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RebalanceConfToUpdate.
func (in *RebalanceConfToUpdate) DeepCopy() *RebalanceConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(RebalanceConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResilverConfToUpdate) DeepCopyInto(out *ResilverConfToUpdate) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResilverConfToUpdate.
func (in *ResilverConfToUpdate) DeepCopy() *ResilverConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(ResilverConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpec) DeepCopyInto(out *ServiceSpec) {
	*out = *in
	out.ServicePort = in.ServicePort
	out.PublicPort = in.PublicPort
	out.IntraControlPort = in.IntraControlPort
	out.IntraDataPort = in.IntraDataPort
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpec.
func (in *ServiceSpec) DeepCopy() *ServiceSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpaceConfToUpdate) DeepCopyInto(out *SpaceConfToUpdate) {
	*out = *in
	if in.CleanupWM != nil {
		in, out := &in.CleanupWM, &out.CleanupWM
		*out = new(int64)
		**out = **in
	}
	if in.LowWM != nil {
		in, out := &in.LowWM, &out.LowWM
		*out = new(int64)
		**out = **in
	}
	if in.HighWM != nil {
		in, out := &in.HighWM, &out.HighWM
		*out = new(int64)
		**out = **in
	}
	if in.OOS != nil {
		in, out := &in.OOS, &out.OOS
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpaceConfToUpdate.
func (in *SpaceConfToUpdate) DeepCopy() *SpaceConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(SpaceConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCBConfToUpdate) DeepCopyInto(out *TCBConfToUpdate) {
	*out = *in
	if in.Compression != nil {
		in, out := &in.Compression, &out.Compression
		*out = new(string)
		**out = **in
	}
	if in.SbundleMult != nil {
		in, out := &in.SbundleMult, &out.SbundleMult
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCBConfToUpdate.
func (in *TCBConfToUpdate) DeepCopy() *TCBConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(TCBConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetSpec) DeepCopyInto(out *TargetSpec) {
	*out = *in
	in.DaemonSpec.DeepCopyInto(&out.DaemonSpec)
	if in.Mounts != nil {
		in, out := &in.Mounts, &out.Mounts
		*out = make([]Mount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AllowSharedOrNoDisks != nil {
		in, out := &in.AllowSharedOrNoDisks, &out.AllowSharedOrNoDisks
		*out = new(bool)
		**out = **in
	}
	if in.DisablePodAntiAffinity != nil {
		in, out := &in.DisablePodAntiAffinity, &out.DisablePodAntiAffinity
		*out = new(bool)
		**out = **in
	}
	if in.HostNetwork != nil {
		in, out := &in.HostNetwork, &out.HostNetwork
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetSpec.
func (in *TargetSpec) DeepCopy() *TargetSpec {
	if in == nil {
		return nil
	}
	out := new(TargetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeoutConfToUpdate) DeepCopyInto(out *TimeoutConfToUpdate) {
	*out = *in
	if in.CplaneOperation != nil {
		in, out := &in.CplaneOperation, &out.CplaneOperation
		*out = new(Duration)
		**out = **in
	}
	if in.MaxKeepalive != nil {
		in, out := &in.MaxKeepalive, &out.MaxKeepalive
		*out = new(Duration)
		**out = **in
	}
	if in.MaxHostBusy != nil {
		in, out := &in.MaxHostBusy, &out.MaxHostBusy
		*out = new(Duration)
		**out = **in
	}
	if in.Startup != nil {
		in, out := &in.Startup, &out.Startup
		*out = new(Duration)
		**out = **in
	}
	if in.JoinAtStartup != nil {
		in, out := &in.JoinAtStartup, &out.JoinAtStartup
		*out = new(Duration)
		**out = **in
	}
	if in.SendFile != nil {
		in, out := &in.SendFile, &out.SendFile
		*out = new(Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeoutConfToUpdate.
func (in *TimeoutConfToUpdate) DeepCopy() *TimeoutConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(TimeoutConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransportConfToUpdate) DeepCopyInto(out *TransportConfToUpdate) {
	*out = *in
	if in.MaxHeaderSize != nil {
		in, out := &in.MaxHeaderSize, &out.MaxHeaderSize
		*out = new(int)
		**out = **in
	}
	if in.Burst != nil {
		in, out := &in.Burst, &out.Burst
		*out = new(int)
		**out = **in
	}
	if in.IdleTeardown != nil {
		in, out := &in.IdleTeardown, &out.IdleTeardown
		*out = new(Duration)
		**out = **in
	}
	if in.QuiesceTime != nil {
		in, out := &in.QuiesceTime, &out.QuiesceTime
		*out = new(Duration)
		**out = **in
	}
	if in.LZ4BlockMaxSize != nil {
		in, out := &in.LZ4BlockMaxSize, &out.LZ4BlockMaxSize
		*out = new(int)
		**out = **in
	}
	if in.LZ4FrameChecksum != nil {
		in, out := &in.LZ4FrameChecksum, &out.LZ4FrameChecksum
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransportConfToUpdate.
func (in *TransportConfToUpdate) DeepCopy() *TransportConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(TransportConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionConfToUpdate) DeepCopyInto(out *VersionConfToUpdate) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ValidateWarmGet != nil {
		in, out := &in.ValidateWarmGet, &out.ValidateWarmGet
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionConfToUpdate.
func (in *VersionConfToUpdate) DeepCopy() *VersionConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(VersionConfToUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WritePolicyConfToUpdate) DeepCopyInto(out *WritePolicyConfToUpdate) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(string)
		**out = **in
	}
	if in.MD != nil {
		in, out := &in.MD, &out.MD
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WritePolicyConfToUpdate.
func (in *WritePolicyConfToUpdate) DeepCopy() *WritePolicyConfToUpdate {
	if in == nil {
		return nil
	}
	out := new(WritePolicyConfToUpdate)
	in.DeepCopyInto(out)
	return out
}
