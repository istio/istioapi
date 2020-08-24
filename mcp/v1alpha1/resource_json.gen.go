// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mcp/v1alpha1/resource.proto

// This package defines the common, core types used by the Mesh Configuration Protocol.

package v1alpha1

import (
	bytes "bytes"
	fmt "fmt"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "google/protobuf/any.proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for Resource
func (this *Resource) MarshalJSON() ([]byte, error) {
	str, err := ResourceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Resource
func (this *Resource) UnmarshalJSON(b []byte) error {
	return ResourceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	ResourceMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	ResourceUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)