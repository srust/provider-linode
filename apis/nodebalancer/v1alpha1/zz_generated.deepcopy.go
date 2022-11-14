//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Nodebalancer) DeepCopyInto(out *Nodebalancer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Nodebalancer.
func (in *Nodebalancer) DeepCopy() *Nodebalancer {
	if in == nil {
		return nil
	}
	out := new(Nodebalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Nodebalancer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodebalancerList) DeepCopyInto(out *NodebalancerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Nodebalancer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodebalancerList.
func (in *NodebalancerList) DeepCopy() *NodebalancerList {
	if in == nil {
		return nil
	}
	out := new(NodebalancerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodebalancerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodebalancerObservation) DeepCopyInto(out *NodebalancerObservation) {
	*out = *in
	if in.Created != nil {
		in, out := &in.Created, &out.Created
		*out = new(string)
		**out = **in
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IPv4 != nil {
		in, out := &in.IPv4, &out.IPv4
		*out = new(string)
		**out = **in
	}
	if in.IPv6 != nil {
		in, out := &in.IPv6, &out.IPv6
		*out = new(string)
		**out = **in
	}
	if in.Transfer != nil {
		in, out := &in.Transfer, &out.Transfer
		*out = make([]TransferObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Updated != nil {
		in, out := &in.Updated, &out.Updated
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodebalancerObservation.
func (in *NodebalancerObservation) DeepCopy() *NodebalancerObservation {
	if in == nil {
		return nil
	}
	out := new(NodebalancerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodebalancerParameters) DeepCopyInto(out *NodebalancerParameters) {
	*out = *in
	if in.ClientConnThrottle != nil {
		in, out := &in.ClientConnThrottle, &out.ClientConnThrottle
		*out = new(float64)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodebalancerParameters.
func (in *NodebalancerParameters) DeepCopy() *NodebalancerParameters {
	if in == nil {
		return nil
	}
	out := new(NodebalancerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodebalancerSpec) DeepCopyInto(out *NodebalancerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodebalancerSpec.
func (in *NodebalancerSpec) DeepCopy() *NodebalancerSpec {
	if in == nil {
		return nil
	}
	out := new(NodebalancerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodebalancerStatus) DeepCopyInto(out *NodebalancerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodebalancerStatus.
func (in *NodebalancerStatus) DeepCopy() *NodebalancerStatus {
	if in == nil {
		return nil
	}
	out := new(NodebalancerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferObservation) DeepCopyInto(out *TransferObservation) {
	*out = *in
	if in.In != nil {
		in, out := &in.In, &out.In
		*out = new(float64)
		**out = **in
	}
	if in.Out != nil {
		in, out := &in.Out, &out.Out
		*out = new(float64)
		**out = **in
	}
	if in.Total != nil {
		in, out := &in.Total, &out.Total
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferObservation.
func (in *TransferObservation) DeepCopy() *TransferObservation {
	if in == nil {
		return nil
	}
	out := new(TransferObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferParameters) DeepCopyInto(out *TransferParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferParameters.
func (in *TransferParameters) DeepCopy() *TransferParameters {
	if in == nil {
		return nil
	}
	out := new(TransferParameters)
	in.DeepCopyInto(out)
	return out
}