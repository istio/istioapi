// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mesh/v1alpha1/config.proto

// Configuration affecting the service mesh as a whole.

package v1alpha1

import (
	bytes "bytes"
	fmt "fmt"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "google/protobuf/duration.proto"
	_ "google/protobuf/wrappers.proto"
	_ "istio.io/api/networking/v1alpha3"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for MeshConfig
func (this *MeshConfig) MarshalJSON() ([]byte, error) {
	str, err := ConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshConfig
func (this *MeshConfig) UnmarshalJSON(b []byte) error {
	return ConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MeshConfig_OutboundTrafficPolicy
func (this *MeshConfig_OutboundTrafficPolicy) MarshalJSON() ([]byte, error) {
	str, err := ConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshConfig_OutboundTrafficPolicy
func (this *MeshConfig_OutboundTrafficPolicy) UnmarshalJSON(b []byte) error {
	return ConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MeshConfig_ThriftConfig
func (this *MeshConfig_ThriftConfig) MarshalJSON() ([]byte, error) {
	str, err := ConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshConfig_ThriftConfig
func (this *MeshConfig_ThriftConfig) UnmarshalJSON(b []byte) error {
	return ConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MeshConfig_ServiceSettings
func (this *MeshConfig_ServiceSettings) MarshalJSON() ([]byte, error) {
	str, err := ConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshConfig_ServiceSettings
func (this *MeshConfig_ServiceSettings) UnmarshalJSON(b []byte) error {
	return ConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MeshConfig_ServiceSettings_Settings
func (this *MeshConfig_ServiceSettings_Settings) MarshalJSON() ([]byte, error) {
	str, err := ConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshConfig_ServiceSettings_Settings
func (this *MeshConfig_ServiceSettings_Settings) UnmarshalJSON(b []byte) error {
	return ConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MeshConfig_CA
func (this *MeshConfig_CA) MarshalJSON() ([]byte, error) {
	str, err := ConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshConfig_CA
func (this *MeshConfig_CA) UnmarshalJSON(b []byte) error {
	return ConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ConfigSource
func (this *ConfigSource) MarshalJSON() ([]byte, error) {
	str, err := ConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ConfigSource
func (this *ConfigSource) UnmarshalJSON(b []byte) error {
	return ConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Certificate
func (this *Certificate) MarshalJSON() ([]byte, error) {
	str, err := ConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Certificate
func (this *Certificate) UnmarshalJSON(b []byte) error {
	return ConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	ConfigMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	ConfigUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)