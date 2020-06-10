// Code generated by protoc-gen-gogo. DO NOT EDIT.
// mixer/v1/attributes.proto is a deprecated file.

package v1

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for Attributes
func (this *Attributes) MarshalJSON() ([]byte, error) {
	str, err := AttributesMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Attributes
func (this *Attributes) UnmarshalJSON(b []byte) error {
	return AttributesUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Attributes_AttributeValue
func (this *Attributes_AttributeValue) MarshalJSON() ([]byte, error) {
	str, err := AttributesMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Attributes_AttributeValue
func (this *Attributes_AttributeValue) UnmarshalJSON(b []byte) error {
	return AttributesUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Attributes_StringMap
func (this *Attributes_StringMap) MarshalJSON() ([]byte, error) {
	str, err := AttributesMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Attributes_StringMap
func (this *Attributes_StringMap) UnmarshalJSON(b []byte) error {
	return AttributesUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for CompressedAttributes
func (this *CompressedAttributes) MarshalJSON() ([]byte, error) {
	str, err := AttributesMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for CompressedAttributes
func (this *CompressedAttributes) UnmarshalJSON(b []byte) error {
	return AttributesUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for StringMap
func (this *StringMap) MarshalJSON() ([]byte, error) {
	str, err := AttributesMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for StringMap
func (this *StringMap) UnmarshalJSON(b []byte) error {
	return AttributesUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	AttributesMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	AttributesUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
