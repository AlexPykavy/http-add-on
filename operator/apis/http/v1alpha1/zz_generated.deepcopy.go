//go:build !ignore_autogenerated

/*
Copyright 2023 The KEDA Authors.

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
func (in *ConcurrencyMetricSpec) DeepCopyInto(out *ConcurrencyMetricSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConcurrencyMetricSpec.
func (in *ConcurrencyMetricSpec) DeepCopy() *ConcurrencyMetricSpec {
	if in == nil {
		return nil
	}
	out := new(ConcurrencyMetricSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Conditions) DeepCopyInto(out *Conditions) {
	{
		in := &in
		*out = make(Conditions, len(*in))
		copy(*out, *in)
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
func (in *HTTPScaledObject) DeepCopyInto(out *HTTPScaledObject) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPScaledObject.
func (in *HTTPScaledObject) DeepCopy() *HTTPScaledObject {
	if in == nil {
		return nil
	}
	out := new(HTTPScaledObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HTTPScaledObject) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPScaledObjectCondition) DeepCopyInto(out *HTTPScaledObjectCondition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPScaledObjectCondition.
func (in *HTTPScaledObjectCondition) DeepCopy() *HTTPScaledObjectCondition {
	if in == nil {
		return nil
	}
	out := new(HTTPScaledObjectCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPScaledObjectList) DeepCopyInto(out *HTTPScaledObjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HTTPScaledObject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPScaledObjectList.
func (in *HTTPScaledObjectList) DeepCopy() *HTTPScaledObjectList {
	if in == nil {
		return nil
	}
	out := new(HTTPScaledObjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HTTPScaledObjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPScaledObjectSpec) DeepCopyInto(out *HTTPScaledObjectSpec) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PathPrefixes != nil {
		in, out := &in.PathPrefixes, &out.PathPrefixes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.ScaleTargetRef = in.ScaleTargetRef
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(ReplicaStruct)
		(*in).DeepCopyInto(*out)
	}
	if in.TargetPendingRequests != nil {
		in, out := &in.TargetPendingRequests, &out.TargetPendingRequests
		*out = new(int32)
		**out = **in
	}
	if in.CooldownPeriod != nil {
		in, out := &in.CooldownPeriod, &out.CooldownPeriod
		*out = new(int32)
		**out = **in
	}
	if in.InitialCooldownPeriod != nil {
		in, out := &in.InitialCooldownPeriod, &out.InitialCooldownPeriod
		*out = new(int32)
		**out = **in
	}
	if in.ScalingMetric != nil {
		in, out := &in.ScalingMetric, &out.ScalingMetric
		*out = new(ScalingMetricSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(HTTPScaledObjectTimeoutsSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPScaledObjectSpec.
func (in *HTTPScaledObjectSpec) DeepCopy() *HTTPScaledObjectSpec {
	if in == nil {
		return nil
	}
	out := new(HTTPScaledObjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPScaledObjectStatus) DeepCopyInto(out *HTTPScaledObjectStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(Conditions, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPScaledObjectStatus.
func (in *HTTPScaledObjectStatus) DeepCopy() *HTTPScaledObjectStatus {
	if in == nil {
		return nil
	}
	out := new(HTTPScaledObjectStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPScaledObjectTimeoutsSpec) DeepCopyInto(out *HTTPScaledObjectTimeoutsSpec) {
	*out = *in
	out.ConditionWait = in.ConditionWait
	out.ResponseHeader = in.ResponseHeader
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPScaledObjectTimeoutsSpec.
func (in *HTTPScaledObjectTimeoutsSpec) DeepCopy() *HTTPScaledObjectTimeoutsSpec {
	if in == nil {
		return nil
	}
	out := new(HTTPScaledObjectTimeoutsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateMetricSpec) DeepCopyInto(out *RateMetricSpec) {
	*out = *in
	out.Window = in.Window
	out.Granularity = in.Granularity
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateMetricSpec.
func (in *RateMetricSpec) DeepCopy() *RateMetricSpec {
	if in == nil {
		return nil
	}
	out := new(RateMetricSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaStruct) DeepCopyInto(out *ReplicaStruct) {
	*out = *in
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		*out = new(int32)
		**out = **in
	}
	if in.Max != nil {
		in, out := &in.Max, &out.Max
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaStruct.
func (in *ReplicaStruct) DeepCopy() *ReplicaStruct {
	if in == nil {
		return nil
	}
	out := new(ReplicaStruct)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaleTargetRef) DeepCopyInto(out *ScaleTargetRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaleTargetRef.
func (in *ScaleTargetRef) DeepCopy() *ScaleTargetRef {
	if in == nil {
		return nil
	}
	out := new(ScaleTargetRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingMetricSpec) DeepCopyInto(out *ScalingMetricSpec) {
	*out = *in
	if in.Concurrency != nil {
		in, out := &in.Concurrency, &out.Concurrency
		*out = new(ConcurrencyMetricSpec)
		**out = **in
	}
	if in.Rate != nil {
		in, out := &in.Rate, &out.Rate
		*out = new(RateMetricSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingMetricSpec.
func (in *ScalingMetricSpec) DeepCopy() *ScalingMetricSpec {
	if in == nil {
		return nil
	}
	out := new(ScalingMetricSpec)
	in.DeepCopyInto(out)
	return out
}
