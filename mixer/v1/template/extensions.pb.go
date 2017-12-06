// Code generated by protoc-gen-gogo.
// source: mixer/v1/template/extensions.proto
// DO NOT EDIT!

/*
Package istio_mixer_v1_template is a generated protocol buffer package.

It is generated from these files:
	mixer/v1/template/extensions.proto
	mixer/v1/template/standard_types.proto

It has these top-level messages:
	IPAddress
	Duration
	TimeStamp
	DNSName
	EmailAddress
	Uri
*/
package istio_mixer_v1_template

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

import strconv "strconv"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Specifies the varity of the the Template.
type TemplateVariety int32

const (
	TEMPLATE_VARIETY_CHECK               TemplateVariety = 0
	TEMPLATE_VARIETY_REPORT              TemplateVariety = 1
	TEMPLATE_VARIETY_QUOTA               TemplateVariety = 2
	TEMPLATE_VARIETY_ATTRIBUTE_GENERATOR TemplateVariety = 3
)

var TemplateVariety_name = map[int32]string{
	0: "TEMPLATE_VARIETY_CHECK",
	1: "TEMPLATE_VARIETY_REPORT",
	2: "TEMPLATE_VARIETY_QUOTA",
	3: "TEMPLATE_VARIETY_ATTRIBUTE_GENERATOR",
}
var TemplateVariety_value = map[string]int32{
	"TEMPLATE_VARIETY_CHECK":               0,
	"TEMPLATE_VARIETY_REPORT":              1,
	"TEMPLATE_VARIETY_QUOTA":               2,
	"TEMPLATE_VARIETY_ATTRIBUTE_GENERATOR": 3,
}

func (TemplateVariety) EnumDescriptor() ([]byte, []int) { return fileDescriptorExtensions, []int{0} }

var E_TemplateVariety = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*TemplateVariety)(nil),
	Field:         72295727,
	Name:          "istio.mixer.v1.template.template_variety",
	Tag:           "varint,72295727,opt,name=template_variety,json=templateVariety,enum=istio.mixer.v1.template.TemplateVariety",
	Filename:      "mixer/v1/template/extensions.proto",
}

func init() {
	proto.RegisterEnum("istio.mixer.v1.template.TemplateVariety", TemplateVariety_name, TemplateVariety_value)
	proto.RegisterExtension(E_TemplateVariety)
}
func (x TemplateVariety) String() string {
	s, ok := TemplateVariety_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

func init() { proto.RegisterFile("mixer/v1/template/extensions.proto", fileDescriptorExtensions) }

var fileDescriptorExtensions = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xca, 0xcd, 0xac, 0x48,
	0x2d, 0xd2, 0x2f, 0x33, 0xd4, 0x2f, 0x49, 0xcd, 0x2d, 0xc8, 0x49, 0x2c, 0x49, 0xd5, 0x4f, 0xad,
	0x28, 0x49, 0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0x2b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12,
	0xcf, 0x2c, 0x2e, 0xc9, 0xcc, 0xd7, 0x03, 0xab, 0xd4, 0x2b, 0x33, 0xd4, 0x83, 0xa9, 0x94, 0x52,
	0x48, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x2b, 0x4b, 0x2a, 0x4d, 0xd3, 0x4f, 0x49, 0x2d,
	0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xc9, 0x2f, 0x82, 0x68, 0xd5, 0x9a, 0xc0, 0xc8, 0xc5, 0x1f, 0x02,
	0x55, 0x1e, 0x96, 0x58, 0x94, 0x99, 0x5a, 0x52, 0x29, 0x24, 0xc5, 0x25, 0x16, 0xe2, 0xea, 0x1b,
	0xe0, 0xe3, 0x18, 0xe2, 0x1a, 0x1f, 0xe6, 0x18, 0xe4, 0xe9, 0x1a, 0x12, 0x19, 0xef, 0xec, 0xe1,
	0xea, 0xec, 0x2d, 0xc0, 0x20, 0x24, 0xcd, 0x25, 0x8e, 0x21, 0x17, 0xe4, 0x1a, 0xe0, 0x1f, 0x14,
	0x22, 0xc0, 0x88, 0x55, 0x63, 0x60, 0xa8, 0x7f, 0x88, 0xa3, 0x00, 0x93, 0x90, 0x06, 0x97, 0x0a,
	0x86, 0x9c, 0x63, 0x48, 0x48, 0x90, 0xa7, 0x53, 0x68, 0x88, 0x6b, 0xbc, 0xbb, 0xab, 0x9f, 0x6b,
	0x90, 0x63, 0x88, 0x7f, 0x90, 0x00, 0xb3, 0x55, 0x09, 0x97, 0x00, 0xcc, 0x03, 0xf1, 0x65, 0x50,
	0x27, 0xc9, 0xe8, 0x41, 0x7c, 0xa2, 0x07, 0xf3, 0x89, 0x9e, 0x5b, 0x66, 0x4e, 0xaa, 0x7f, 0x41,
	0x09, 0x28, 0x14, 0x24, 0xd6, 0x9f, 0xda, 0xa3, 0xa4, 0xc0, 0xa8, 0xc1, 0x67, 0xa4, 0xa1, 0x87,
	0x23, 0x24, 0xf4, 0xd0, 0xfc, 0x18, 0xc4, 0x5f, 0x82, 0x2a, 0xe0, 0xa4, 0x73, 0xe1, 0xa1, 0x1c,
	0xc3, 0x8d, 0x87, 0x72, 0x0c, 0x1f, 0x1e, 0xca, 0x31, 0x36, 0x3c, 0x92, 0x63, 0x5c, 0xf1, 0x48,
	0x8e, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x7c, 0xf1,
	0x48, 0x8e, 0xe1, 0xc3, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x92, 0xd8, 0xc0, 0xee, 0x30,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xdb, 0xe2, 0x9b, 0xa0, 0x9e, 0x01, 0x00, 0x00,
}
