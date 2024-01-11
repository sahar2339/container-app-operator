//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationLinks) DeepCopyInto(out *ApplicationLinks) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationLinks.
func (in *ApplicationLinks) DeepCopy() *ApplicationLinks {
	if in == nil {
		return nil
	}
	out := new(ApplicationLinks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Capp) DeepCopyInto(out *Capp) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Capp.
func (in *Capp) DeepCopy() *Capp {
	if in == nil {
		return nil
	}
	out := new(Capp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Capp) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CappList) DeepCopyInto(out *CappList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Capp, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CappList.
func (in *CappList) DeepCopy() *CappList {
	if in == nil {
		return nil
	}
	out := new(CappList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CappList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CappSpec) DeepCopyInto(out *CappSpec) {
	*out = *in
	in.ConfigurationSpec.DeepCopyInto(&out.ConfigurationSpec)
	in.RouteSpec.DeepCopyInto(&out.RouteSpec)
	out.LogSpec = in.LogSpec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CappSpec.
func (in *CappSpec) DeepCopy() *CappSpec {
	if in == nil {
		return nil
	}
	out := new(CappSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CappStatus) DeepCopyInto(out *CappStatus) {
	*out = *in
	out.ApplicationLinks = in.ApplicationLinks
	in.KnativeObjectStatus.DeepCopyInto(&out.KnativeObjectStatus)
	if in.RevisionInfo != nil {
		in, out := &in.RevisionInfo, &out.RevisionInfo
		*out = make([]RevisionInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.StateStatus.DeepCopyInto(&out.StateStatus)
	in.LoggingStatus.DeepCopyInto(&out.LoggingStatus)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CappStatus.
func (in *CappStatus) DeepCopy() *CappStatus {
	if in == nil {
		return nil
	}
	out := new(CappStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogSpec) DeepCopyInto(out *LogSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogSpec.
func (in *LogSpec) DeepCopy() *LogSpec {
	if in == nil {
		return nil
	}
	out := new(LogSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingStatus) DeepCopyInto(out *LoggingStatus) {
	*out = *in
	in.Flow.DeepCopyInto(&out.Flow)
	in.Output.DeepCopyInto(&out.Output)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingStatus.
func (in *LoggingStatus) DeepCopy() *LoggingStatus {
	if in == nil {
		return nil
	}
	out := new(LoggingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevisionInfo) DeepCopyInto(out *RevisionInfo) {
	*out = *in
	in.RevisionStatus.DeepCopyInto(&out.RevisionStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevisionInfo.
func (in *RevisionInfo) DeepCopy() *RevisionInfo {
	if in == nil {
		return nil
	}
	out := new(RevisionInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteSpec) DeepCopyInto(out *RouteSpec) {
	*out = *in
	in.TrafficTarget.DeepCopyInto(&out.TrafficTarget)
	if in.RouteTimeoutSeconds != nil {
		in, out := &in.RouteTimeoutSeconds, &out.RouteTimeoutSeconds
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteSpec.
func (in *RouteSpec) DeepCopy() *RouteSpec {
	if in == nil {
		return nil
	}
	out := new(RouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StateStatus) DeepCopyInto(out *StateStatus) {
	*out = *in
	in.LastChange.DeepCopyInto(&out.LastChange)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StateStatus.
func (in *StateStatus) DeepCopy() *StateStatus {
	if in == nil {
		return nil
	}
	out := new(StateStatus)
	in.DeepCopyInto(out)
	return out
}
