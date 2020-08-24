// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mcp/v1alpha1/metadata.proto

package v1alpha1

import (
	bytes "bytes"
	fmt "fmt"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "google/protobuf/timestamp.proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for Metadata
func (this *Metadata) MarshalJSON() ([]byte, error) {
	str, err := MetadataMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Metadata
func (this *Metadata) UnmarshalJSON(b []byte) error {
	return MetadataUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	MetadataMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	MetadataUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
