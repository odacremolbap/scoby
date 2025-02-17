//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2023 TriggerMesh Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	"knative.dev/pkg/apis"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddConfiguration) DeepCopyInto(out *AddConfiguration) {
	*out = *in
	if in.ToEnv != nil {
		in, out := &in.ToEnv, &out.ToEnv
		*out = make([]AddToEnvConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ToVolume != nil {
		in, out := &in.ToVolume, &out.ToVolume
		*out = make([]AddToVolumeConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddConfiguration.
func (in *AddConfiguration) DeepCopy() *AddConfiguration {
	if in == nil {
		return nil
	}
	out := new(AddConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddToEnvConfiguration) DeepCopyInto(out *AddToEnvConfiguration) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(AddToEnvValueFrom)
		(*in).DeepCopyInto(*out)
	}
	if in.ValueFromControllerConfigMap != nil {
		in, out := &in.ValueFromControllerConfigMap, &out.ValueFromControllerConfigMap
		*out = new(v1.ConfigMapKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddToEnvConfiguration.
func (in *AddToEnvConfiguration) DeepCopy() *AddToEnvConfiguration {
	if in == nil {
		return nil
	}
	out := new(AddToEnvConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddToEnvValueFrom) DeepCopyInto(out *AddToEnvValueFrom) {
	*out = *in
	if in.ConfigMap != nil {
		in, out := &in.ConfigMap, &out.ConfigMap
		*out = new(v1.ConfigMapKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.FieldRef != nil {
		in, out := &in.FieldRef, &out.FieldRef
		*out = new(v1.ObjectFieldSelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddToEnvValueFrom.
func (in *AddToEnvValueFrom) DeepCopy() *AddToEnvValueFrom {
	if in == nil {
		return nil
	}
	out := new(AddToEnvValueFrom)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddToVolumeConfiguration) DeepCopyInto(out *AddToVolumeConfiguration) {
	*out = *in
	in.MountFrom.DeepCopyInto(&out.MountFrom)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddToVolumeConfiguration.
func (in *AddToVolumeConfiguration) DeepCopy() *AddToVolumeConfiguration {
	if in == nil {
		return nil
	}
	out := new(AddToVolumeConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Address) DeepCopyInto(out *Address) {
	*out = *in
	if in.Ref != nil {
		in, out := &in.Ref, &out.Ref
		*out = new(Reference)
		**out = **in
	}
	if in.URI != nil {
		in, out := &in.URI, &out.URI
		*out = new(apis.URL)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Address.
func (in *Address) DeepCopy() *Address {
	if in == nil {
		return nil
	}
	out := new(Address)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuiltInfunction) DeepCopyInto(out *BuiltInfunction) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuiltInfunction.
func (in *BuiltInfunction) DeepCopy() *BuiltInfunction {
	if in == nil {
		return nil
	}
	out := new(BuiltInfunction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Conditions) DeepCopyInto(out *Conditions) {
	{
		in := &in
		*out = make(Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Conditions.
func (in Conditions) DeepCopy() Conditions {
	if in == nil {
		return nil
	}
	out := new(Conditions)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionsFromHook) DeepCopyInto(out *ConditionsFromHook) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionsFromHook.
func (in *ConditionsFromHook) DeepCopy() *ConditionsFromHook {
	if in == nil {
		return nil
	}
	out := new(ConditionsFromHook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentFormFactor) DeepCopyInto(out *DeploymentFormFactor) {
	*out = *in
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(DeploymentService)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentFormFactor.
func (in *DeploymentFormFactor) DeepCopy() *DeploymentFormFactor {
	if in == nil {
		return nil
	}
	out := new(DeploymentFormFactor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentService) DeepCopyInto(out *DeploymentService) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentService.
func (in *DeploymentService) DeepCopy() *DeploymentService {
	if in == nil {
		return nil
	}
	out := new(DeploymentService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FormFactor) DeepCopyInto(out *FormFactor) {
	*out = *in
	if in.Deployment != nil {
		in, out := &in.Deployment, &out.Deployment
		*out = new(DeploymentFormFactor)
		(*in).DeepCopyInto(*out)
	}
	if in.KnativeService != nil {
		in, out := &in.KnativeService, &out.KnativeService
		*out = new(KnativeServiceFormFactor)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FormFactor.
func (in *FormFactor) DeepCopy() *FormFactor {
	if in == nil {
		return nil
	}
	out := new(FormFactor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FromSpecConfiguration) DeepCopyInto(out *FromSpecConfiguration) {
	*out = *in
	if in.Skip != nil {
		in, out := &in.Skip, &out.Skip
		*out = make([]FromSpecSkip, len(*in))
		copy(*out, *in)
	}
	if in.ToEnv != nil {
		in, out := &in.ToEnv, &out.ToEnv
		*out = make([]FromSpecToEnv, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ToVolume != nil {
		in, out := &in.ToVolume, &out.ToVolume
		*out = make([]FromSpecToVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FromSpecConfiguration.
func (in *FromSpecConfiguration) DeepCopy() *FromSpecConfiguration {
	if in == nil {
		return nil
	}
	out := new(FromSpecConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FromSpecSkip) DeepCopyInto(out *FromSpecSkip) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FromSpecSkip.
func (in *FromSpecSkip) DeepCopy() *FromSpecSkip {
	if in == nil {
		return nil
	}
	out := new(FromSpecSkip)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FromSpecToEnv) DeepCopyInto(out *FromSpecToEnv) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(SpecToEnvDefaultValue)
		(*in).DeepCopyInto(*out)
	}
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(SpecToEnvValueFrom)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FromSpecToEnv.
func (in *FromSpecToEnv) DeepCopy() *FromSpecToEnv {
	if in == nil {
		return nil
	}
	out := new(FromSpecToEnv)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FromSpecToVolume) DeepCopyInto(out *FromSpecToVolume) {
	*out = *in
	in.MountFrom.DeepCopyInto(&out.MountFrom)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FromSpecToVolume.
func (in *FromSpecToVolume) DeepCopy() *FromSpecToVolume {
	if in == nil {
		return nil
	}
	out := new(FromSpecToVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalParameterConfiguration) DeepCopyInto(out *GlobalParameterConfiguration) {
	*out = *in
	if in.DefaultPrefix != nil {
		in, out := &in.DefaultPrefix, &out.DefaultPrefix
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalParameterConfiguration.
func (in *GlobalParameterConfiguration) DeepCopy() *GlobalParameterConfiguration {
	if in == nil {
		return nil
	}
	out := new(GlobalParameterConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hook) DeepCopyInto(out *Hook) {
	*out = *in
	in.Address.DeepCopyInto(&out.Address)
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(string)
		**out = **in
	}
	if in.Capabilities != nil {
		in, out := &in.Capabilities, &out.Capabilities
		*out = make(HookCapabilities, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hook.
func (in *Hook) DeepCopy() *Hook {
	if in == nil {
		return nil
	}
	out := new(Hook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in HookCapabilities) DeepCopyInto(out *HookCapabilities) {
	{
		in := &in
		*out = make(HookCapabilities, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HookCapabilities.
func (in HookCapabilities) DeepCopy() HookCapabilities {
	if in == nil {
		return nil
	}
	out := new(HookCapabilities)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KnativeServiceFormFactor) DeepCopyInto(out *KnativeServiceFormFactor) {
	*out = *in
	if in.MinScale != nil {
		in, out := &in.MinScale, &out.MinScale
		*out = new(int)
		**out = **in
	}
	if in.MaxScale != nil {
		in, out := &in.MaxScale, &out.MaxScale
		*out = new(int)
		**out = **in
	}
	if in.Visibility != nil {
		in, out := &in.Visibility, &out.Visibility
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KnativeServiceFormFactor.
func (in *KnativeServiceFormFactor) DeepCopy() *KnativeServiceFormFactor {
	if in == nil {
		return nil
	}
	out := new(KnativeServiceFormFactor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MountFrom) DeepCopyInto(out *MountFrom) {
	*out = *in
	if in.ConfigMap != nil {
		in, out := &in.ConfigMap, &out.ConfigMap
		*out = new(v1.ConfigMapKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MountFrom.
func (in *MountFrom) DeepCopy() *MountFrom {
	if in == nil {
		return nil
	}
	out := new(MountFrom)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterConfiguration) DeepCopyInto(out *ParameterConfiguration) {
	*out = *in
	if in.Global != nil {
		in, out := &in.Global, &out.Global
		*out = new(GlobalParameterConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Add != nil {
		in, out := &in.Add, &out.Add
		*out = new(AddConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.FromSpec != nil {
		in, out := &in.FromSpec, &out.FromSpec
		*out = new(FromSpecConfiguration)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterConfiguration.
func (in *ParameterConfiguration) DeepCopy() *ParameterConfiguration {
	if in == nil {
		return nil
	}
	out := new(ParameterConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Reference) DeepCopyInto(out *Reference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Reference.
func (in *Reference) DeepCopy() *Reference {
	if in == nil {
		return nil
	}
	out := new(Reference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistrationFromImage) DeepCopyInto(out *RegistrationFromImage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistrationFromImage.
func (in *RegistrationFromImage) DeepCopy() *RegistrationFromImage {
	if in == nil {
		return nil
	}
	out := new(RegistrationFromImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpecToEnvDefaultValue) DeepCopyInto(out *SpecToEnvDefaultValue) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	if in.ConfigMap != nil {
		in, out := &in.ConfigMap, &out.ConfigMap
		*out = new(v1.ConfigMapKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpecToEnvDefaultValue.
func (in *SpecToEnvDefaultValue) DeepCopy() *SpecToEnvDefaultValue {
	if in == nil {
		return nil
	}
	out := new(SpecToEnvDefaultValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpecToEnvValueFrom) DeepCopyInto(out *SpecToEnvValueFrom) {
	*out = *in
	if in.ConfigMap != nil {
		in, out := &in.ConfigMap, &out.ConfigMap
		*out = new(v1.ConfigMapKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.BuiltInFunc != nil {
		in, out := &in.BuiltInFunc, &out.BuiltInFunc
		*out = new(BuiltInfunction)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpecToEnvValueFrom.
func (in *SpecToEnvValueFrom) DeepCopy() *SpecToEnvValueFrom {
	if in == nil {
		return nil
	}
	out := new(SpecToEnvValueFrom)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Status) DeepCopyInto(out *Status) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Status.
func (in *Status) DeepCopy() *Status {
	if in == nil {
		return nil
	}
	out := new(Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusAddElement) DeepCopyInto(out *StatusAddElement) {
	*out = *in
	out.ValueFrom = in.ValueFrom
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusAddElement.
func (in *StatusAddElement) DeepCopy() *StatusAddElement {
	if in == nil {
		return nil
	}
	out := new(StatusAddElement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusConfiguration) DeepCopyInto(out *StatusConfiguration) {
	*out = *in
	if in.AddElements != nil {
		in, out := &in.AddElements, &out.AddElements
		*out = make([]StatusAddElement, len(*in))
		copy(*out, *in)
	}
	if in.ConditionsFromHook != nil {
		in, out := &in.ConditionsFromHook, &out.ConditionsFromHook
		*out = make([]ConditionsFromHook, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusConfiguration.
func (in *StatusConfiguration) DeepCopy() *StatusConfiguration {
	if in == nil {
		return nil
	}
	out := new(StatusConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusValueFrom) DeepCopyInto(out *StatusValueFrom) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusValueFrom.
func (in *StatusValueFrom) DeepCopy() *StatusValueFrom {
	if in == nil {
		return nil
	}
	out := new(StatusValueFrom)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workload) DeepCopyInto(out *Workload) {
	*out = *in
	if in.FormFactor != nil {
		in, out := &in.FormFactor, &out.FormFactor
		*out = new(FormFactor)
		(*in).DeepCopyInto(*out)
	}
	out.FromImage = in.FromImage
	if in.ParameterConfiguration != nil {
		in, out := &in.ParameterConfiguration, &out.ParameterConfiguration
		*out = new(ParameterConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.StatusConfiguration != nil {
		in, out := &in.StatusConfiguration, &out.StatusConfiguration
		*out = new(StatusConfiguration)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workload.
func (in *Workload) DeepCopy() *Workload {
	if in == nil {
		return nil
	}
	out := new(Workload)
	in.DeepCopyInto(out)
	return out
}
