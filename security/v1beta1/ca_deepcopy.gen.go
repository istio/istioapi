// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: security/v1alpha1/ca.proto

// Keep this package for backward compatibility.

package v1beta1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using IstioCertificateRequest within kubernetes types, where deepcopy-gen is used.
func (in *IstioCertificateRequest) DeepCopyInto(out *IstioCertificateRequest) {
	p := proto.Clone(in).(*IstioCertificateRequest)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioCertificateRequest. Required by controller-gen.
func (in *IstioCertificateRequest) DeepCopy() *IstioCertificateRequest {
	if in == nil {
		return nil
	}
	out := new(IstioCertificateRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using IstioCertificateResponse within kubernetes types, where deepcopy-gen is used.
func (in *IstioCertificateResponse) DeepCopyInto(out *IstioCertificateResponse) {
	p := proto.Clone(in).(*IstioCertificateResponse)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioCertificateResponse. Required by controller-gen.
func (in *IstioCertificateResponse) DeepCopy() *IstioCertificateResponse {
	if in == nil {
		return nil
	}
	out := new(IstioCertificateResponse)
	in.DeepCopyInto(out)
	return out
}
