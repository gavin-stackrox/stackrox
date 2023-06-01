//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 Red Hat.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdditionalCA) DeepCopyInto(out *AdditionalCA) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdditionalCA.
func (in *AdditionalCA) DeepCopy() *AdditionalCA {
	if in == nil {
		return nil
	}
	out := new(AdditionalCA)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminPasswordStatus) DeepCopyInto(out *AdminPasswordStatus) {
	*out = *in
	if in.SecretReference != nil {
		in, out := &in.SecretReference, &out.SecretReference
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminPasswordStatus.
func (in *AdminPasswordStatus) DeepCopy() *AdminPasswordStatus {
	if in == nil {
		return nil
	}
	out := new(AdminPasswordStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdmissionControlComponentSpec) DeepCopyInto(out *AdmissionControlComponentSpec) {
	*out = *in
	if in.ListenOnCreates != nil {
		in, out := &in.ListenOnCreates, &out.ListenOnCreates
		*out = new(bool)
		**out = **in
	}
	if in.ListenOnUpdates != nil {
		in, out := &in.ListenOnUpdates, &out.ListenOnUpdates
		*out = new(bool)
		**out = **in
	}
	if in.ListenOnEvents != nil {
		in, out := &in.ListenOnEvents, &out.ListenOnEvents
		*out = new(bool)
		**out = **in
	}
	if in.ContactImageScanners != nil {
		in, out := &in.ContactImageScanners, &out.ContactImageScanners
		*out = new(ImageScanPolicy)
		**out = **in
	}
	if in.TimeoutSeconds != nil {
		in, out := &in.TimeoutSeconds, &out.TimeoutSeconds
		*out = new(int32)
		**out = **in
	}
	if in.Bypass != nil {
		in, out := &in.Bypass, &out.Bypass
		*out = new(BypassPolicy)
		**out = **in
	}
	in.DeploymentSpec.DeepCopyInto(&out.DeploymentSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdmissionControlComponentSpec.
func (in *AdmissionControlComponentSpec) DeepCopy() *AdmissionControlComponentSpec {
	if in == nil {
		return nil
	}
	out := new(AdmissionControlComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuditLogsSpec) DeepCopyInto(out *AuditLogsSpec) {
	*out = *in
	if in.Collection != nil {
		in, out := &in.Collection, &out.Collection
		*out = new(AuditLogsCollectionSetting)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuditLogsSpec.
func (in *AuditLogsSpec) DeepCopy() *AuditLogsSpec {
	if in == nil {
		return nil
	}
	out := new(AuditLogsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Central) DeepCopyInto(out *Central) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Central.
func (in *Central) DeepCopy() *Central {
	if in == nil {
		return nil
	}
	out := new(Central)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Central) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CentralComponentSpec) DeepCopyInto(out *CentralComponentSpec) {
	*out = *in
	if in.AdminPasswordSecret != nil {
		in, out := &in.AdminPasswordSecret, &out.AdminPasswordSecret
		*out = new(LocalSecretReference)
		**out = **in
	}
	if in.AdminPasswordGenerationDisabled != nil {
		in, out := &in.AdminPasswordGenerationDisabled, &out.AdminPasswordGenerationDisabled
		*out = new(bool)
		**out = **in
	}
	if in.Exposure != nil {
		in, out := &in.Exposure, &out.Exposure
		*out = new(Exposure)
		(*in).DeepCopyInto(*out)
	}
	if in.DefaultTLSSecret != nil {
		in, out := &in.DefaultTLSSecret, &out.DefaultTLSSecret
		*out = new(LocalSecretReference)
		**out = **in
	}
	if in.Monitoring != nil {
		in, out := &in.Monitoring, &out.Monitoring
		*out = new(Monitoring)
		(*in).DeepCopyInto(*out)
	}
	if in.Persistence != nil {
		in, out := &in.Persistence, &out.Persistence
		*out = new(Persistence)
		(*in).DeepCopyInto(*out)
	}
	if in.DB != nil {
		in, out := &in.DB, &out.DB
		*out = new(CentralDBSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Telemetry != nil {
		in, out := &in.Telemetry, &out.Telemetry
		*out = new(Telemetry)
		(*in).DeepCopyInto(*out)
	}
	if in.DeclarativeConfiguration != nil {
		in, out := &in.DeclarativeConfiguration, &out.DeclarativeConfiguration
		*out = new(DeclarativeConfiguration)
		(*in).DeepCopyInto(*out)
	}
	in.DeploymentSpec.DeepCopyInto(&out.DeploymentSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CentralComponentSpec.
func (in *CentralComponentSpec) DeepCopy() *CentralComponentSpec {
	if in == nil {
		return nil
	}
	out := new(CentralComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CentralComponentStatus) DeepCopyInto(out *CentralComponentStatus) {
	*out = *in
	if in.AdminPassword != nil {
		in, out := &in.AdminPassword, &out.AdminPassword
		*out = new(AdminPasswordStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CentralComponentStatus.
func (in *CentralComponentStatus) DeepCopy() *CentralComponentStatus {
	if in == nil {
		return nil
	}
	out := new(CentralComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CentralDBSpec) DeepCopyInto(out *CentralDBSpec) {
	*out = *in
	if in.IsEnabled != nil {
		in, out := &in.IsEnabled, &out.IsEnabled
		*out = new(CentralDBEnabled)
		**out = **in
	}
	if in.PasswordSecret != nil {
		in, out := &in.PasswordSecret, &out.PasswordSecret
		*out = new(LocalSecretReference)
		**out = **in
	}
	if in.ConnectionStringOverride != nil {
		in, out := &in.ConnectionStringOverride, &out.ConnectionStringOverride
		*out = new(string)
		**out = **in
	}
	if in.Persistence != nil {
		in, out := &in.Persistence, &out.Persistence
		*out = new(DBPersistence)
		(*in).DeepCopyInto(*out)
	}
	out.ConfigOverride = in.ConfigOverride
	in.DeploymentSpec.DeepCopyInto(&out.DeploymentSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CentralDBSpec.
func (in *CentralDBSpec) DeepCopy() *CentralDBSpec {
	if in == nil {
		return nil
	}
	out := new(CentralDBSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CentralList) DeepCopyInto(out *CentralList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Central, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CentralList.
func (in *CentralList) DeepCopy() *CentralList {
	if in == nil {
		return nil
	}
	out := new(CentralList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CentralList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CentralSpec) DeepCopyInto(out *CentralSpec) {
	*out = *in
	if in.Central != nil {
		in, out := &in.Central, &out.Central
		*out = new(CentralComponentSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Scanner != nil {
		in, out := &in.Scanner, &out.Scanner
		*out = new(ScannerComponentSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Egress != nil {
		in, out := &in.Egress, &out.Egress
		*out = new(Egress)
		(*in).DeepCopyInto(*out)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]LocalSecretReference, len(*in))
		copy(*out, *in)
	}
	if in.Customize != nil {
		in, out := &in.Customize, &out.Customize
		*out = new(CustomizeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Misc != nil {
		in, out := &in.Misc, &out.Misc
		*out = new(MiscSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CentralSpec.
func (in *CentralSpec) DeepCopy() *CentralSpec {
	if in == nil {
		return nil
	}
	out := new(CentralSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CentralStatus) DeepCopyInto(out *CentralStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]StackRoxCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DeployedRelease != nil {
		in, out := &in.DeployedRelease, &out.DeployedRelease
		*out = new(StackRoxRelease)
		**out = **in
	}
	if in.Central != nil {
		in, out := &in.Central, &out.Central
		*out = new(CentralComponentStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CentralStatus.
func (in *CentralStatus) DeepCopy() *CentralStatus {
	if in == nil {
		return nil
	}
	out := new(CentralStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectorContainerSpec) DeepCopyInto(out *CollectorContainerSpec) {
	*out = *in
	if in.Collection != nil {
		in, out := &in.Collection, &out.Collection
		*out = new(CollectionMethod)
		**out = **in
	}
	if in.ImageFlavor != nil {
		in, out := &in.ImageFlavor, &out.ImageFlavor
		*out = new(CollectorImageFlavor)
		**out = **in
	}
	in.ContainerSpec.DeepCopyInto(&out.ContainerSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectorContainerSpec.
func (in *CollectorContainerSpec) DeepCopy() *CollectorContainerSpec {
	if in == nil {
		return nil
	}
	out := new(CollectorContainerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerSpec) DeepCopyInto(out *ContainerSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerSpec.
func (in *ContainerSpec) DeepCopy() *ContainerSpec {
	if in == nil {
		return nil
	}
	out := new(ContainerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizeSpec) DeepCopyInto(out *CustomizeSpec) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizeSpec.
func (in *CustomizeSpec) DeepCopy() *CustomizeSpec {
	if in == nil {
		return nil
	}
	out := new(CustomizeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBPersistence) DeepCopyInto(out *DBPersistence) {
	*out = *in
	if in.PersistentVolumeClaim != nil {
		in, out := &in.PersistentVolumeClaim, &out.PersistentVolumeClaim
		*out = new(DBPersistentVolumeClaim)
		(*in).DeepCopyInto(*out)
	}
	if in.HostPath != nil {
		in, out := &in.HostPath, &out.HostPath
		*out = new(HostPathSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBPersistence.
func (in *DBPersistence) DeepCopy() *DBPersistence {
	if in == nil {
		return nil
	}
	out := new(DBPersistence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBPersistentVolumeClaim) DeepCopyInto(out *DBPersistentVolumeClaim) {
	*out = *in
	if in.ClaimName != nil {
		in, out := &in.ClaimName, &out.ClaimName
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(string)
		**out = **in
	}
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBPersistentVolumeClaim.
func (in *DBPersistentVolumeClaim) DeepCopy() *DBPersistentVolumeClaim {
	if in == nil {
		return nil
	}
	out := new(DBPersistentVolumeClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeclarativeConfiguration) DeepCopyInto(out *DeclarativeConfiguration) {
	*out = *in
	if in.ConfigMaps != nil {
		in, out := &in.ConfigMaps, &out.ConfigMaps
		*out = make([]LocalConfigMapReference, len(*in))
		copy(*out, *in)
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]LocalSecretReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeclarativeConfiguration.
func (in *DeclarativeConfiguration) DeepCopy() *DeclarativeConfiguration {
	if in == nil {
		return nil
	}
	out := new(DeclarativeConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentSpec) DeepCopyInto(out *DeploymentSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
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
		*out = make([]*v1.Toleration, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(v1.Toleration)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentSpec.
func (in *DeploymentSpec) DeepCopy() *DeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(DeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Egress) DeepCopyInto(out *Egress) {
	*out = *in
	if in.ConnectivityPolicy != nil {
		in, out := &in.ConnectivityPolicy, &out.ConnectivityPolicy
		*out = new(ConnectivityPolicy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Egress.
func (in *Egress) DeepCopy() *Egress {
	if in == nil {
		return nil
	}
	out := new(Egress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Exposure) DeepCopyInto(out *Exposure) {
	*out = *in
	if in.Route != nil {
		in, out := &in.Route, &out.Route
		*out = new(ExposureRoute)
		(*in).DeepCopyInto(*out)
	}
	if in.LoadBalancer != nil {
		in, out := &in.LoadBalancer, &out.LoadBalancer
		*out = new(ExposureLoadBalancer)
		(*in).DeepCopyInto(*out)
	}
	if in.NodePort != nil {
		in, out := &in.NodePort, &out.NodePort
		*out = new(ExposureNodePort)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Exposure.
func (in *Exposure) DeepCopy() *Exposure {
	if in == nil {
		return nil
	}
	out := new(Exposure)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExposureLoadBalancer) DeepCopyInto(out *ExposureLoadBalancer) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExposureLoadBalancer.
func (in *ExposureLoadBalancer) DeepCopy() *ExposureLoadBalancer {
	if in == nil {
		return nil
	}
	out := new(ExposureLoadBalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExposureNodePort) DeepCopyInto(out *ExposureNodePort) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExposureNodePort.
func (in *ExposureNodePort) DeepCopy() *ExposureNodePort {
	if in == nil {
		return nil
	}
	out := new(ExposureNodePort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExposureRoute) DeepCopyInto(out *ExposureRoute) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExposureRoute.
func (in *ExposureRoute) DeepCopy() *ExposureRoute {
	if in == nil {
		return nil
	}
	out := new(ExposureRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostPathSpec) DeepCopyInto(out *HostPathSpec) {
	*out = *in
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostPathSpec.
func (in *HostPathSpec) DeepCopy() *HostPathSpec {
	if in == nil {
		return nil
	}
	out := new(HostPathSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalConfigMapReference) DeepCopyInto(out *LocalConfigMapReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalConfigMapReference.
func (in *LocalConfigMapReference) DeepCopy() *LocalConfigMapReference {
	if in == nil {
		return nil
	}
	out := new(LocalConfigMapReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalScannerComponentSpec) DeepCopyInto(out *LocalScannerComponentSpec) {
	*out = *in
	if in.ScannerComponent != nil {
		in, out := &in.ScannerComponent, &out.ScannerComponent
		*out = new(LocalScannerComponentPolicy)
		**out = **in
	}
	if in.Analyzer != nil {
		in, out := &in.Analyzer, &out.Analyzer
		*out = new(ScannerAnalyzerComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.DB != nil {
		in, out := &in.DB, &out.DB
		*out = new(DeploymentSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalScannerComponentSpec.
func (in *LocalScannerComponentSpec) DeepCopy() *LocalScannerComponentSpec {
	if in == nil {
		return nil
	}
	out := new(LocalScannerComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalSecretReference) DeepCopyInto(out *LocalSecretReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalSecretReference.
func (in *LocalSecretReference) DeepCopy() *LocalSecretReference {
	if in == nil {
		return nil
	}
	out := new(LocalSecretReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MiscSpec) DeepCopyInto(out *MiscSpec) {
	*out = *in
	if in.CreateSCCs != nil {
		in, out := &in.CreateSCCs, &out.CreateSCCs
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MiscSpec.
func (in *MiscSpec) DeepCopy() *MiscSpec {
	if in == nil {
		return nil
	}
	out := new(MiscSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Monitoring) DeepCopyInto(out *Monitoring) {
	*out = *in
	if in.ExposeEndpoint != nil {
		in, out := &in.ExposeEndpoint, &out.ExposeEndpoint
		*out = new(ExposeEndpoint)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Monitoring.
func (in *Monitoring) DeepCopy() *Monitoring {
	if in == nil {
		return nil
	}
	out := new(Monitoring)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerNodeSpec) DeepCopyInto(out *PerNodeSpec) {
	*out = *in
	if in.Collector != nil {
		in, out := &in.Collector, &out.Collector
		*out = new(CollectorContainerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Compliance != nil {
		in, out := &in.Compliance, &out.Compliance
		*out = new(ContainerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeInventory != nil {
		in, out := &in.NodeInventory, &out.NodeInventory
		*out = new(ContainerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.TaintToleration != nil {
		in, out := &in.TaintToleration, &out.TaintToleration
		*out = new(TaintTolerationPolicy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerNodeSpec.
func (in *PerNodeSpec) DeepCopy() *PerNodeSpec {
	if in == nil {
		return nil
	}
	out := new(PerNodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Persistence) DeepCopyInto(out *Persistence) {
	*out = *in
	if in.PersistentVolumeClaim != nil {
		in, out := &in.PersistentVolumeClaim, &out.PersistentVolumeClaim
		*out = new(PersistentVolumeClaim)
		(*in).DeepCopyInto(*out)
	}
	if in.HostPath != nil {
		in, out := &in.HostPath, &out.HostPath
		*out = new(HostPathSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Persistence.
func (in *Persistence) DeepCopy() *Persistence {
	if in == nil {
		return nil
	}
	out := new(Persistence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolumeClaim) DeepCopyInto(out *PersistentVolumeClaim) {
	*out = *in
	if in.ClaimName != nil {
		in, out := &in.ClaimName, &out.ClaimName
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(string)
		**out = **in
	}
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolumeClaim.
func (in *PersistentVolumeClaim) DeepCopy() *PersistentVolumeClaim {
	if in == nil {
		return nil
	}
	out := new(PersistentVolumeClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScannerAnalyzerComponent) DeepCopyInto(out *ScannerAnalyzerComponent) {
	*out = *in
	if in.Scaling != nil {
		in, out := &in.Scaling, &out.Scaling
		*out = new(ScannerAnalyzerScaling)
		(*in).DeepCopyInto(*out)
	}
	in.DeploymentSpec.DeepCopyInto(&out.DeploymentSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScannerAnalyzerComponent.
func (in *ScannerAnalyzerComponent) DeepCopy() *ScannerAnalyzerComponent {
	if in == nil {
		return nil
	}
	out := new(ScannerAnalyzerComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScannerAnalyzerScaling) DeepCopyInto(out *ScannerAnalyzerScaling) {
	*out = *in
	if in.AutoScaling != nil {
		in, out := &in.AutoScaling, &out.AutoScaling
		*out = new(AutoScalingPolicy)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	if in.MaxReplicas != nil {
		in, out := &in.MaxReplicas, &out.MaxReplicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScannerAnalyzerScaling.
func (in *ScannerAnalyzerScaling) DeepCopy() *ScannerAnalyzerScaling {
	if in == nil {
		return nil
	}
	out := new(ScannerAnalyzerScaling)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScannerComponentSpec) DeepCopyInto(out *ScannerComponentSpec) {
	*out = *in
	if in.ScannerComponent != nil {
		in, out := &in.ScannerComponent, &out.ScannerComponent
		*out = new(ScannerComponentPolicy)
		**out = **in
	}
	if in.Analyzer != nil {
		in, out := &in.Analyzer, &out.Analyzer
		*out = new(ScannerAnalyzerComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.DB != nil {
		in, out := &in.DB, &out.DB
		*out = new(DeploymentSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Monitoring != nil {
		in, out := &in.Monitoring, &out.Monitoring
		*out = new(Monitoring)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScannerComponentSpec.
func (in *ScannerComponentSpec) DeepCopy() *ScannerComponentSpec {
	if in == nil {
		return nil
	}
	out := new(ScannerComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecuredCluster) DeepCopyInto(out *SecuredCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecuredCluster.
func (in *SecuredCluster) DeepCopy() *SecuredCluster {
	if in == nil {
		return nil
	}
	out := new(SecuredCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecuredCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecuredClusterList) DeepCopyInto(out *SecuredClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecuredCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecuredClusterList.
func (in *SecuredClusterList) DeepCopy() *SecuredClusterList {
	if in == nil {
		return nil
	}
	out := new(SecuredClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecuredClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecuredClusterSpec) DeepCopyInto(out *SecuredClusterSpec) {
	*out = *in
	if in.ClusterLabels != nil {
		in, out := &in.ClusterLabels, &out.ClusterLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Sensor != nil {
		in, out := &in.Sensor, &out.Sensor
		*out = new(SensorComponentSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.AdmissionControl != nil {
		in, out := &in.AdmissionControl, &out.AdmissionControl
		*out = new(AdmissionControlComponentSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.PerNode != nil {
		in, out := &in.PerNode, &out.PerNode
		*out = new(PerNodeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.AuditLogs != nil {
		in, out := &in.AuditLogs, &out.AuditLogs
		*out = new(AuditLogsSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Scanner != nil {
		in, out := &in.Scanner, &out.Scanner
		*out = new(LocalScannerComponentSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]LocalSecretReference, len(*in))
		copy(*out, *in)
	}
	if in.Customize != nil {
		in, out := &in.Customize, &out.Customize
		*out = new(CustomizeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Misc != nil {
		in, out := &in.Misc, &out.Misc
		*out = new(MiscSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecuredClusterSpec.
func (in *SecuredClusterSpec) DeepCopy() *SecuredClusterSpec {
	if in == nil {
		return nil
	}
	out := new(SecuredClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecuredClusterStatus) DeepCopyInto(out *SecuredClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]StackRoxCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DeployedRelease != nil {
		in, out := &in.DeployedRelease, &out.DeployedRelease
		*out = new(StackRoxRelease)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecuredClusterStatus.
func (in *SecuredClusterStatus) DeepCopy() *SecuredClusterStatus {
	if in == nil {
		return nil
	}
	out := new(SecuredClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SensorComponentSpec) DeepCopyInto(out *SensorComponentSpec) {
	*out = *in
	in.DeploymentSpec.DeepCopyInto(&out.DeploymentSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SensorComponentSpec.
func (in *SensorComponentSpec) DeepCopy() *SensorComponentSpec {
	if in == nil {
		return nil
	}
	out := new(SensorComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackRoxCondition) DeepCopyInto(out *StackRoxCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackRoxCondition.
func (in *StackRoxCondition) DeepCopy() *StackRoxCondition {
	if in == nil {
		return nil
	}
	out := new(StackRoxCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackRoxRelease) DeepCopyInto(out *StackRoxRelease) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackRoxRelease.
func (in *StackRoxRelease) DeepCopy() *StackRoxRelease {
	if in == nil {
		return nil
	}
	out := new(StackRoxRelease)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSConfig) DeepCopyInto(out *TLSConfig) {
	*out = *in
	if in.AdditionalCAs != nil {
		in, out := &in.AdditionalCAs, &out.AdditionalCAs
		*out = make([]AdditionalCA, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSConfig.
func (in *TLSConfig) DeepCopy() *TLSConfig {
	if in == nil {
		return nil
	}
	out := new(TLSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Telemetry) DeepCopyInto(out *Telemetry) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(TelemetryStorage)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Telemetry.
func (in *Telemetry) DeepCopy() *Telemetry {
	if in == nil {
		return nil
	}
	out := new(Telemetry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TelemetryStorage) DeepCopyInto(out *TelemetryStorage) {
	*out = *in
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TelemetryStorage.
func (in *TelemetryStorage) DeepCopy() *TelemetryStorage {
	if in == nil {
		return nil
	}
	out := new(TelemetryStorage)
	in.DeepCopyInto(out)
	return out
}
