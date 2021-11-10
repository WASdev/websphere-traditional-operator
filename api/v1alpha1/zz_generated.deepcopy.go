//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebsphereTraditionalApplication) DeepCopyInto(out *WebsphereTraditionalApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebsphereTraditionalApplication.
func (in *WebsphereTraditionalApplication) DeepCopy() *WebsphereTraditionalApplication {
	if in == nil {
		return nil
	}
	out := new(WebsphereTraditionalApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WebsphereTraditionalApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebsphereTraditionalApplicationList) DeepCopyInto(out *WebsphereTraditionalApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WebsphereTraditionalApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebsphereTraditionalApplicationList.
func (in *WebsphereTraditionalApplicationList) DeepCopy() *WebsphereTraditionalApplicationList {
	if in == nil {
		return nil
	}
	out := new(WebsphereTraditionalApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WebsphereTraditionalApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebsphereTraditionalApplicationSpec) DeepCopyInto(out *WebsphereTraditionalApplicationSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebsphereTraditionalApplicationSpec.
func (in *WebsphereTraditionalApplicationSpec) DeepCopy() *WebsphereTraditionalApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(WebsphereTraditionalApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebsphereTraditionalApplicationStatus) DeepCopyInto(out *WebsphereTraditionalApplicationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebsphereTraditionalApplicationStatus.
func (in *WebsphereTraditionalApplicationStatus) DeepCopy() *WebsphereTraditionalApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(WebsphereTraditionalApplicationStatus)
	in.DeepCopyInto(out)
	return out
}
