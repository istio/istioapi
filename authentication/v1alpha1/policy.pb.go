// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authentication/v1alpha1/policy.proto

/*
Package v1alpha1 is a generated protocol buffer package.

It is generated from these files:
	authentication/v1alpha1/policy.proto

It has these top-level messages:
	MutualTLS
	Mechanism
	Destination
	Policy
*/
package v1alpha1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "gogoproto"
import istio_mixer_v1_config_client1 "istio.io/api/mixer/v1/config/client"
import _ "istio.io/api/routing/v1alpha2"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Placer holder for mTLS authentication params.
type MutualTLS struct {
}

func (m *MutualTLS) Reset()                    { *m = MutualTLS{} }
func (m *MutualTLS) String() string            { return proto.CompactTextString(m) }
func (*MutualTLS) ProtoMessage()               {}
func (*MutualTLS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// AuthenticationMechanism defines one particular type of authentication (i.e
// mutual TLS, JWT etc). The type can be progammatically determine by checking
// the type of the "params" field.
type Mechanism struct {
	// Types that are valid to be assigned to Params:
	//	*Mechanism_None
	//	*Mechanism_Mtls
	//	*Mechanism_Jwt
	Params isMechanism_Params `protobuf_oneof:"params"`
}

func (m *Mechanism) Reset()                    { *m = Mechanism{} }
func (m *Mechanism) String() string            { return proto.CompactTextString(m) }
func (*Mechanism) ProtoMessage()               {}
func (*Mechanism) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isMechanism_Params interface{ isMechanism_Params() }

type Mechanism_None struct {
	None bool `protobuf:"varint,1,opt,name=none,oneof"`
}
type Mechanism_Mtls struct {
	Mtls *MutualTLS `protobuf:"bytes,2,opt,name=mtls,oneof"`
}
type Mechanism_Jwt struct {
	Jwt *istio_mixer_v1_config_client1.JWT `protobuf:"bytes,3,opt,name=jwt,oneof"`
}

func (*Mechanism_None) isMechanism_Params() {}
func (*Mechanism_Mtls) isMechanism_Params() {}
func (*Mechanism_Jwt) isMechanism_Params()  {}

func (m *Mechanism) GetParams() isMechanism_Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *Mechanism) GetNone() bool {
	if x, ok := m.GetParams().(*Mechanism_None); ok {
		return x.None
	}
	return false
}

func (m *Mechanism) GetMtls() *MutualTLS {
	if x, ok := m.GetParams().(*Mechanism_Mtls); ok {
		return x.Mtls
	}
	return nil
}

func (m *Mechanism) GetJwt() *istio_mixer_v1_config_client1.JWT {
	if x, ok := m.GetParams().(*Mechanism_Jwt); ok {
		return x.Jwt
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Mechanism) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Mechanism_OneofMarshaler, _Mechanism_OneofUnmarshaler, _Mechanism_OneofSizer, []interface{}{
		(*Mechanism_None)(nil),
		(*Mechanism_Mtls)(nil),
		(*Mechanism_Jwt)(nil),
	}
}

func _Mechanism_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Mechanism)
	// params
	switch x := m.Params.(type) {
	case *Mechanism_None:
		t := uint64(0)
		if x.None {
			t = 1
		}
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *Mechanism_Mtls:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Mtls); err != nil {
			return err
		}
	case *Mechanism_Jwt:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Jwt); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Mechanism.Params has unexpected type %T", x)
	}
	return nil
}

func _Mechanism_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Mechanism)
	switch tag {
	case 1: // params.none
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Params = &Mechanism_None{x != 0}
		return true, err
	case 2: // params.mtls
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MutualTLS)
		err := b.DecodeMessage(msg)
		m.Params = &Mechanism_Mtls{msg}
		return true, err
	case 3: // params.jwt
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(istio_mixer_v1_config_client1.JWT)
		err := b.DecodeMessage(msg)
		m.Params = &Mechanism_Jwt{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Mechanism_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Mechanism)
	// params
	switch x := m.Params.(type) {
	case *Mechanism_None:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += 1
	case *Mechanism_Mtls:
		s := proto.Size(x.Mtls)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Mechanism_Jwt:
		s := proto.Size(x.Jwt)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Destination struct {
	// REQUIRED. The name can be a short name or a fully qualified domain
	// name from the service registry, a resolvable DNS name, or an IP
	// address.
	// If short names are used, the FQDN of the service will be resolved in a
	// platform specific manner.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Optional. List of ports that service are exposed on. Leave blank to applied
	// to all ports.
	Ports []uint32 `protobuf:"varint,3,rep,packed,name=ports" json:"ports,omitempty"`
}

func (m *Destination) Reset()                    { *m = Destination{} }
func (m *Destination) String() string            { return proto.CompactTextString(m) }
func (*Destination) ProtoMessage()               {}
func (*Destination) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Destination) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Destination) GetPorts() []uint32 {
	if m != nil {
		return m.Ports
	}
	return nil
}

// AuthenticationPolicy binds credentials to workload(s).
// Authentication policy is composed of 2-part authentication:
// - peer: verify caller service credentials.
// - end_user: verify end-user credentials.
// For each part, if it's not empty, at least one of those listed credential
// must be provided and  (successfully) verified for the authentication to pass.
//
// Examples:
// Policy to enable mTLS for all services in namespace frod
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: RouteRule
//     metadata:
//       name: mTLS-enable
//       namespace: frod
//     spec:
//       match:
//       peer:
//       - mtls: {}
//
// Policy to enable mTLS, and use JWT for productpage:9000
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: RouteRule
//     metadata:
//       name: mTLS-enable
//       namespace: frod
//     spec:
//       match:
//       - name: productpage
//         ports:
//         - 9000
//       peer:
//       - mtls:
//       end_user:
//       - jwt:
//           issuer: "https://securetoken.google.com"
//           audiences:
//           - "productpage"
//           jwks_uri: "https://www.googleapis.com/oauth2/v1/certs"
//           locations:
//           - header: x-goog-iap-jwt-assertion
type Policy struct {
	// List of destinations (workloads) that the policy should be applied on.
	// If empty, policy will be used on all destinations in the same namespace.
	Match []*Destination `protobuf:"bytes,1,rep,name=match" json:"match,omitempty"`
	// List of credential that should be checked by peer authentication. They
	// will be validated in sequence, until the first one satisfied. If none of
	// the specified mechanism valid, the whole authentication should fail.
	// On the other hand, the first valid credential will be used to extract
	// peer identity (i.e the source.user attribute in the request to mixer).
	Peers []*Mechanism `protobuf:"bytes,2,rep,name=peers" json:"peers,omitempty"`
	// Similar to above, but for end_user authentication, which will extract
	// request.auth.principal/audiences/presenter if authentication succeed.
	EndUsers []*Mechanism `protobuf:"bytes,3,rep,name=end_users,json=endUsers" json:"end_users,omitempty"`
}

func (m *Policy) Reset()                    { *m = Policy{} }
func (m *Policy) String() string            { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()               {}
func (*Policy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Policy) GetMatch() []*Destination {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *Policy) GetPeers() []*Mechanism {
	if m != nil {
		return m.Peers
	}
	return nil
}

func (m *Policy) GetEndUsers() []*Mechanism {
	if m != nil {
		return m.EndUsers
	}
	return nil
}

func init() {
	proto.RegisterType((*MutualTLS)(nil), "istio.authentication.v1alpha1.MutualTLS")
	proto.RegisterType((*Mechanism)(nil), "istio.authentication.v1alpha1.Mechanism")
	proto.RegisterType((*Destination)(nil), "istio.authentication.v1alpha1.Destination")
	proto.RegisterType((*Policy)(nil), "istio.authentication.v1alpha1.Policy")
}

func init() { proto.RegisterFile("authentication/v1alpha1/policy.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x8a, 0xdb, 0x30,
	0x18, 0xc5, 0xed, 0x3a, 0x09, 0xb1, 0x4c, 0x37, 0x22, 0x0b, 0x13, 0x68, 0x70, 0x4c, 0x16, 0xa6,
	0x0b, 0x09, 0xbb, 0x94, 0xee, 0x42, 0x09, 0x2d, 0x84, 0xd2, 0x40, 0x71, 0x53, 0x0a, 0xdd, 0x04,
	0xd5, 0x51, 0x6d, 0x15, 0x5b, 0x32, 0xb6, 0x9c, 0x76, 0xae, 0x33, 0xab, 0x39, 0xca, 0x6c, 0xe7,
	0x08, 0x99, 0x13, 0xcc, 0x11, 0x06, 0x49, 0xe3, 0x61, 0x66, 0x31, 0xff, 0x76, 0x92, 0x79, 0xbf,
	0xf7, 0x7d, 0xef, 0x59, 0x60, 0x41, 0x3a, 0x59, 0x50, 0x2e, 0x59, 0x46, 0x24, 0x13, 0x1c, 0x1f,
	0x62, 0x52, 0xd6, 0x05, 0x89, 0x71, 0x2d, 0x4a, 0x96, 0x9d, 0xa0, 0xba, 0x11, 0x52, 0xc0, 0x37,
	0xac, 0x95, 0x4c, 0xa0, 0xfb, 0x5a, 0xd4, 0x6b, 0xa7, 0x93, 0x5c, 0xe4, 0x42, 0x2b, 0xb1, 0x3a,
	0x19, 0x68, 0x3a, 0xaf, 0xd8, 0x7f, 0xda, 0xe0, 0x43, 0x8c, 0x33, 0xc1, 0xff, 0xb0, 0x1c, 0x67,
	0x25, 0xa3, 0x5c, 0x62, 0xe5, 0xd2, 0x4b, 0x1a, 0xd1, 0x49, 0xc6, 0xf3, 0x7e, 0x6c, 0x82, 0xd5,
	0x07, 0xba, 0x6b, 0xba, 0x92, 0x1a, 0x49, 0xe8, 0x01, 0x77, 0xd3, 0xc9, 0x8e, 0x94, 0xdb, 0xaf,
	0xdf, 0xc3, 0x53, 0x1b, 0xb8, 0x1b, 0x9a, 0x15, 0x84, 0xb3, 0xb6, 0x82, 0x13, 0x30, 0xe0, 0x82,
	0x53, 0xdf, 0x0e, 0xec, 0x68, 0xbc, 0xb6, 0x52, 0x7d, 0x83, 0x4b, 0x30, 0xa8, 0x64, 0xd9, 0xfa,
	0xaf, 0x02, 0x3b, 0xf2, 0x92, 0x08, 0x3d, 0xba, 0x3a, 0xba, 0xf5, 0x56, 0xbc, 0xe2, 0xe0, 0x7b,
	0xe0, 0xfc, 0xfd, 0x27, 0x7d, 0x47, 0xe3, 0xf3, 0x1b, 0x5c, 0x47, 0x41, 0x87, 0x18, 0x99, 0x28,
	0xc8, 0x44, 0x41, 0x5f, 0x7e, 0x6e, 0xd7, 0x56, 0xaa, 0xf4, 0xab, 0x31, 0x18, 0xd5, 0xa4, 0x21,
	0x55, 0x1b, 0x7e, 0x00, 0xde, 0x27, 0xda, 0x4a, 0xc6, 0xf5, 0x28, 0x08, 0xc1, 0x80, 0x93, 0xca,
	0x6c, 0xe9, 0xa6, 0xfa, 0x0c, 0x27, 0x60, 0x58, 0x8b, 0x46, 0xb6, 0xbe, 0x13, 0x38, 0xd1, 0xeb,
	0xd4, 0x5c, 0xc2, 0x0b, 0x1b, 0x8c, 0xbe, 0xe9, 0xda, 0xe1, 0x47, 0x30, 0xac, 0x88, 0xcc, 0x0a,
	0xdf, 0x0e, 0x9c, 0xc8, 0x4b, 0xde, 0x3e, 0x91, 0xe2, 0xce, 0xbc, 0xd4, 0x80, 0x70, 0x09, 0x86,
	0x35, 0xa5, 0x8d, 0xea, 0xc1, 0x79, 0x4e, 0x0f, 0x7d, 0xab, 0xa9, 0xc1, 0xe0, 0x67, 0xe0, 0x52,
	0xbe, 0xdf, 0x75, 0xad, 0xf2, 0x70, 0x5e, 0xe8, 0x31, 0xa6, 0x7c, 0xff, 0x43, 0x91, 0xab, 0xe4,
	0xfc, 0x38, 0xb3, 0xae, 0x8e, 0x33, 0xeb, 0xec, 0x72, 0x66, 0xfd, 0x5a, 0x18, 0x03, 0x26, 0x30,
	0xa9, 0x19, 0x7e, 0xe0, 0xe9, 0xfd, 0x1e, 0xe9, 0x3f, 0xff, 0xee, 0x3a, 0x00, 0x00, 0xff, 0xff,
	0x0a, 0x19, 0x4c, 0x39, 0x9c, 0x02, 0x00, 0x00,
}
