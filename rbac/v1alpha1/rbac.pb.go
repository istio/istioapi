// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rbac/v1alpha1/rbac.proto

/*
Package istio_rbac_v1alpha1 is a generated protocol buffer package.

It is generated from these files:
	rbac/v1alpha1/rbac.proto

It has these top-level messages:
	ServiceRole
	AccessRule
	ServiceRoleBinding
	Subject
	RoleRef
*/
package istio_rbac_v1alpha1

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

// ServiceRole specification contains a list of access rules (permissions).
// This represent the "Spec" part of the ServiceRole object. The name and namespace
// of the ServiceRole is specified in "metadata" section of the ServiceRole object.
type ServiceRole struct {
	// Required. The set of access rules (permissions) that the role has.
	Rules []*AccessRule `protobuf:"bytes,1,rep,name=rules" json:"rules,omitempty"`
}

func (m *ServiceRole) Reset()                    { *m = ServiceRole{} }
func (m *ServiceRole) String() string            { return proto.CompactTextString(m) }
func (*ServiceRole) ProtoMessage()               {}
func (*ServiceRole) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ServiceRole) GetRules() []*AccessRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// AccessRule defines a permission to access a list of services.
type AccessRule struct {
	// Required. A list of service names.
	// Exact match, prefix match, and suffix match are supported for service names.
	// For example, the service name "bookstore.mtv.cluster.local" matches
	// "bookstore.mtv.cluster.local" (exact match), or "bookstore*" (prefix match),
	// or "*.mtv.cluster.local" (suffix match).
	// If set to ["*"], it refers to all services in the namespace.
	Services []string `protobuf:"bytes,1,rep,name=services" json:"services,omitempty"`
	// Optional. A list of HTTP paths.
	// Exact match, prefix match, and suffix match are supported for paths.
	// For example, the path "/books/review" matches
	// "/books/review" (exact match), or "/books/*" (prefix match),
	// or "*/review" (suffix match).
	// If not specified, it applies to any path.
	Paths []string `protobuf:"bytes,2,rep,name=paths" json:"paths,omitempty"`
	// Required. A list of HTTP methods (e.g., "GET", "POST") or gRPC methods.
	// gRPC methods must be presented as fully-qualified name in the form of
	// packageName.serviceName/methodName.
	// If set to ["*"], it applies to any method.
	Methods []string `protobuf:"bytes,3,rep,name=methods" json:"methods,omitempty"`
	// Optional. Extra constraints in the ServiceRole specification.
	// The above ServiceRole examples shows an example of constraint "version".
	Constraints []*AccessRule_Constraint `protobuf:"bytes,4,rep,name=constraints" json:"constraints,omitempty"`
}

func (m *AccessRule) Reset()                    { *m = AccessRule{} }
func (m *AccessRule) String() string            { return proto.CompactTextString(m) }
func (*AccessRule) ProtoMessage()               {}
func (*AccessRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AccessRule) GetServices() []string {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *AccessRule) GetPaths() []string {
	if m != nil {
		return m.Paths
	}
	return nil
}

func (m *AccessRule) GetMethods() []string {
	if m != nil {
		return m.Methods
	}
	return nil
}

func (m *AccessRule) GetConstraints() []*AccessRule_Constraint {
	if m != nil {
		return m.Constraints
	}
	return nil
}

// Definition of a custom constraint.
type AccessRule_Constraint struct {
	// Key of the constraint.
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// List of valid values for the constraint.
	// Exact match, prefix match, and suffix match are supported for constraint values.
	// For example, the value "v1alpha2" matches
	// "v1alpha2" (exact match), or "v1*" (prefix match),
	// or "*alpha2" (suffix match).
	Values []string `protobuf:"bytes,2,rep,name=values" json:"values,omitempty"`
}

func (m *AccessRule_Constraint) Reset()                    { *m = AccessRule_Constraint{} }
func (m *AccessRule_Constraint) String() string            { return proto.CompactTextString(m) }
func (*AccessRule_Constraint) ProtoMessage()               {}
func (*AccessRule_Constraint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *AccessRule_Constraint) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *AccessRule_Constraint) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

// ServiceRoleBinding assigns a ServiceRole to a list of subjects.
// This represents the "Spec" part of the ServiceRoleBinding object. The name and namespace
// of the ServiceRoleBinding is specified in "metadata" section of the ServiceRoleBinding
// object.
type ServiceRoleBinding struct {
	// Required. List of subjects that are assigned the ServiceRole object.
	Subjects []*Subject `protobuf:"bytes,1,rep,name=subjects" json:"subjects,omitempty"`
	// Required. Reference to the ServiceRole object.
	RoleRef *RoleRef `protobuf:"bytes,2,opt,name=roleRef" json:"roleRef,omitempty"`
}

func (m *ServiceRoleBinding) Reset()                    { *m = ServiceRoleBinding{} }
func (m *ServiceRoleBinding) String() string            { return proto.CompactTextString(m) }
func (*ServiceRoleBinding) ProtoMessage()               {}
func (*ServiceRoleBinding) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ServiceRoleBinding) GetSubjects() []*Subject {
	if m != nil {
		return m.Subjects
	}
	return nil
}

func (m *ServiceRoleBinding) GetRoleRef() *RoleRef {
	if m != nil {
		return m.RoleRef
	}
	return nil
}

// Subject defines an identity or a group of identities.
type Subject struct {
	// Optional. The user name/ID that the subject represents.
	User string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	// Optional. The group that the subject belongs to.
	Group string `protobuf:"bytes,2,opt,name=group" json:"group,omitempty"`
	// Optional. The set of properties that identify the subject.
	// In the above ServiceRoleBinding example, the second subject has two properties:
	//     service: "reviews"
	//     namespace: "abc"
	Properties map[string]string `protobuf:"bytes,3,rep,name=properties" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Subject) Reset()                    { *m = Subject{} }
func (m *Subject) String() string            { return proto.CompactTextString(m) }
func (*Subject) ProtoMessage()               {}
func (*Subject) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Subject) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Subject) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *Subject) GetProperties() map[string]string {
	if m != nil {
		return m.Properties
	}
	return nil
}

// RoleRef refers to a role object.
type RoleRef struct {
	// Required. The type of the role being referenced.
	// Currently, "ServiceRole" is the only supported value for "kind".
	Kind string `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	// Required. The name of the ServiceRole object being referenced.
	// The ServiceRole object must be in the same namespace as the ServiceRoleBinding
	// object.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *RoleRef) Reset()                    { *m = RoleRef{} }
func (m *RoleRef) String() string            { return proto.CompactTextString(m) }
func (*RoleRef) ProtoMessage()               {}
func (*RoleRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RoleRef) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *RoleRef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*ServiceRole)(nil), "istio.rbac.v1alpha1.ServiceRole")
	proto.RegisterType((*AccessRule)(nil), "istio.rbac.v1alpha1.AccessRule")
	proto.RegisterType((*AccessRule_Constraint)(nil), "istio.rbac.v1alpha1.AccessRule.Constraint")
	proto.RegisterType((*ServiceRoleBinding)(nil), "istio.rbac.v1alpha1.ServiceRoleBinding")
	proto.RegisterType((*Subject)(nil), "istio.rbac.v1alpha1.Subject")
	proto.RegisterType((*RoleRef)(nil), "istio.rbac.v1alpha1.RoleRef")
}

func init() { proto.RegisterFile("rbac/v1alpha1/rbac.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0x6a, 0x2a, 0x31,
	0x18, 0x25, 0xfe, 0x8d, 0x7e, 0xb3, 0xb8, 0x97, 0x5c, 0xb9, 0x04, 0x29, 0x54, 0x66, 0x25, 0xa5,
	0x8c, 0x68, 0xa9, 0x48, 0xa1, 0x8b, 0xfe, 0xed, 0x5c, 0x94, 0xf8, 0x04, 0xe3, 0x98, 0x6a, 0xea,
	0x38, 0x19, 0x92, 0x8c, 0xe0, 0x0b, 0xf4, 0xb9, 0xfa, 0x24, 0x7d, 0x96, 0x32, 0xc9, 0x64, 0x94,
	0x62, 0xdb, 0xdd, 0x77, 0xbe, 0xef, 0x9c, 0xc3, 0xc9, 0x21, 0x40, 0xe4, 0x22, 0x8a, 0x87, 0xbb,
	0x51, 0x94, 0x64, 0xeb, 0x68, 0x34, 0x2c, 0x50, 0x98, 0x49, 0xa1, 0x05, 0xfe, 0xc7, 0x95, 0xe6,
	0x22, 0x34, 0x1b, 0x77, 0x0f, 0x1e, 0xc1, 0x9f, 0x33, 0xb9, 0xe3, 0x31, 0xa3, 0x22, 0x61, 0xf8,
	0x1a, 0x9a, 0x32, 0x4f, 0x98, 0x22, 0xa8, 0x5f, 0x1f, 0xf8, 0xe3, 0xf3, 0xf0, 0x84, 0x26, 0xbc,
	0x8b, 0x63, 0xa6, 0x14, 0xcd, 0x13, 0x46, 0x2d, 0x3b, 0xf8, 0x40, 0x00, 0x87, 0x2d, 0xee, 0x41,
	0x5b, 0x59, 0x53, 0x6b, 0xd4, 0xa1, 0x15, 0xc6, 0x5d, 0x68, 0x66, 0x91, 0x5e, 0x2b, 0x52, 0x33,
	0x07, 0x0b, 0x30, 0x01, 0x6f, 0xcb, 0xf4, 0x5a, 0x2c, 0x15, 0xa9, 0x9b, 0xbd, 0x83, 0x78, 0x06,
	0x7e, 0x2c, 0x52, 0xa5, 0x65, 0xc4, 0x53, 0xad, 0x48, 0xc3, 0xe4, 0xba, 0xf8, 0x25, 0x57, 0xf8,
	0x50, 0x49, 0xe8, 0xb1, 0xbc, 0x37, 0x01, 0x38, 0x9c, 0xf0, 0x5f, 0xa8, 0x6f, 0xd8, 0x9e, 0xa0,
	0x3e, 0x1a, 0x74, 0x68, 0x31, 0xe2, 0xff, 0xd0, 0xda, 0x45, 0x49, 0xce, 0x5c, 0xbc, 0x12, 0x05,
	0x6f, 0x08, 0xf0, 0x51, 0x4f, 0xf7, 0x3c, 0x5d, 0xf2, 0x74, 0x85, 0xa7, 0xd0, 0x56, 0xf9, 0xe2,
	0x95, 0xc5, 0xda, 0x35, 0x76, 0x76, 0x32, 0xd9, 0xdc, 0x92, 0x68, 0xc5, 0xc6, 0x13, 0xf0, 0xa4,
	0x48, 0x18, 0x65, 0x2f, 0xa4, 0xd6, 0x47, 0xdf, 0x0a, 0xa9, 0xe5, 0x50, 0x47, 0x0e, 0xde, 0x11,
	0x78, 0xa5, 0x1b, 0xc6, 0xd0, 0xc8, 0x15, 0x93, 0x65, 0x7e, 0x33, 0x17, 0xf5, 0xae, 0xa4, 0xc8,
	0x33, 0xe3, 0xda, 0xa1, 0x16, 0xe0, 0x19, 0x40, 0x26, 0x45, 0xc6, 0xa4, 0xe6, 0xcc, 0x36, 0xec,
	0x8f, 0x2f, 0x7f, 0x4a, 0x1a, 0x3e, 0x57, 0xf4, 0xa7, 0x54, 0xcb, 0x3d, 0x3d, 0xd2, 0xf7, 0x6e,
	0xe1, 0xcf, 0x97, 0xf3, 0x89, 0x26, 0xbb, 0xd0, 0x34, 0xdd, 0xb9, 0x20, 0x06, 0xdc, 0xd4, 0xa6,
	0x28, 0x18, 0x81, 0x57, 0x3e, 0xab, 0x78, 0xc1, 0x86, 0xa7, 0x4b, 0xf7, 0x82, 0x62, 0x2e, 0x76,
	0x69, 0xb4, 0x75, 0x3a, 0x33, 0x2f, 0x5a, 0xe6, 0x07, 0x5f, 0x7d, 0x06, 0x00, 0x00, 0xff, 0xff,
	0xf4, 0x3d, 0x4a, 0xf1, 0xdd, 0x02, 0x00, 0x00,
}
