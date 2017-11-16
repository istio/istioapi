// Code generated by protoc-gen-go. DO NOT EDIT.
// source: broker/v1/config/service_class.proto

/*
Package istio_broker_v1_config is a generated protocol buffer package.

It is generated from these files:
	broker/v1/config/service_class.proto
	broker/v1/config/service_plan.proto

It has these top-level messages:
	ServiceClass
	Deployment
	CatalogEntry
	ServicePlan
	CatalogPlan
*/
package istio_broker_v1_config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ServiceClass defines a service that are exposed to Istio service consumers.
// The service is linked into one or more ServicePlan.
type ServiceClass struct {
	// Required. Istio deployment spec for the service class.
	Deployment *Deployment `protobuf:"bytes,1,opt,name=deployment" json:"deployment,omitempty"`
	// Required. Listing information for the public catalog.
	Entry *CatalogEntry `protobuf:"bytes,2,opt,name=entry" json:"entry,omitempty"`
}

func (m *ServiceClass) Reset()                    { *m = ServiceClass{} }
func (m *ServiceClass) String() string            { return proto.CompactTextString(m) }
func (*ServiceClass) ProtoMessage()               {}
func (*ServiceClass) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ServiceClass) GetDeployment() *Deployment {
	if m != nil {
		return m.Deployment
	}
	return nil
}

func (m *ServiceClass) GetEntry() *CatalogEntry {
	if m != nil {
		return m.Entry
	}
	return nil
}

// Deployment defines how the service instances are deployed.
type Deployment struct {
	// For truely multi-tenant service, the deployed service instance name.
	Instance string `protobuf:"bytes,1,opt,name=instance" json:"instance,omitempty"`
}

func (m *Deployment) Reset()                    { *m = Deployment{} }
func (m *Deployment) String() string            { return proto.CompactTextString(m) }
func (*Deployment) ProtoMessage()               {}
func (*Deployment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Deployment) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

// CatalogEntry defines listing information for this service within the exposed
// catalog.  The message is a subset of OSBI service fields defined in
// https://github.com/openservicebrokerapi
type CatalogEntry struct {
	// Required. Public service name.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Required. Public unique service guid.
	Id string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	// Required. Public short service description.
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
}

func (m *CatalogEntry) Reset()                    { *m = CatalogEntry{} }
func (m *CatalogEntry) String() string            { return proto.CompactTextString(m) }
func (*CatalogEntry) ProtoMessage()               {}
func (*CatalogEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CatalogEntry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CatalogEntry) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CatalogEntry) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*ServiceClass)(nil), "istio.broker.v1.config.ServiceClass")
	proto.RegisterType((*Deployment)(nil), "istio.broker.v1.config.Deployment")
	proto.RegisterType((*CatalogEntry)(nil), "istio.broker.v1.config.CatalogEntry")
}

func init() { proto.RegisterFile("broker/v1/config/service_class.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x95, 0xf0, 0x47, 0xf4, 0x5a, 0x31, 0xdc, 0x80, 0x22, 0xa6, 0x2a, 0xea, 0xd0, 0xc9,
	0x51, 0x61, 0x63, 0xa4, 0xf0, 0x02, 0x81, 0x1d, 0xb9, 0xf6, 0x51, 0x9d, 0x48, 0x7d, 0x91, 0x6d,
	0x45, 0xea, 0x4b, 0xf0, 0xcc, 0xa8, 0xb6, 0x08, 0x19, 0xe8, 0x66, 0xdf, 0xfd, 0x7e, 0x9f, 0x3e,
	0x1d, 0xac, 0x76, 0x5e, 0xbe, 0xc8, 0x37, 0xc3, 0xa6, 0x31, 0xe2, 0x3e, 0x79, 0xdf, 0x04, 0xf2,
	0x03, 0x1b, 0xfa, 0x30, 0x9d, 0x0e, 0x41, 0xf5, 0x5e, 0xa2, 0xe0, 0x1d, 0x87, 0xc8, 0xa2, 0x32,
	0xab, 0x86, 0x8d, 0xca, 0x6c, 0xfd, 0x5d, 0xc0, 0xe2, 0x2d, 0xf3, 0xdb, 0x13, 0x8e, 0xcf, 0x00,
	0x96, 0xfa, 0x4e, 0x8e, 0x07, 0x72, 0xb1, 0x2a, 0x96, 0xc5, 0x7a, 0xfe, 0x50, 0xab, 0xff, 0x6d,
	0xf5, 0x32, 0x92, 0xed, 0xc4, 0xc2, 0x27, 0xb8, 0x22, 0x17, 0xfd, 0xb1, 0x2a, 0x93, 0xbe, 0x3a,
	0xa7, 0x6f, 0x75, 0xd4, 0x9d, 0xec, 0x5f, 0x4f, 0x6c, 0x9b, 0x95, 0x7a, 0x0d, 0xf0, 0x97, 0x8a,
	0xf7, 0x70, 0xc3, 0x2e, 0x44, 0xed, 0x0c, 0xa5, 0x2e, 0xb3, 0x76, 0xfc, 0xd7, 0xef, 0xb0, 0x98,
	0x06, 0x20, 0xc2, 0xa5, 0xd3, 0x87, 0x5f, 0x2e, 0xbd, 0xf1, 0x16, 0x4a, 0xb6, 0xa9, 0xc6, 0xac,
	0x2d, 0xd9, 0xe2, 0x12, 0xe6, 0x96, 0x82, 0xf1, 0xdc, 0x47, 0x16, 0x57, 0x5d, 0xa4, 0xc5, 0x74,
	0xb4, 0xbb, 0x4e, 0xf7, 0x7a, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x99, 0x33, 0x90, 0x8d, 0x57,
	0x01, 0x00, 0x00,
}
