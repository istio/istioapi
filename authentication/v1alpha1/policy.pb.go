// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authentication/v1alpha1/policy.proto

/*
Package v1alpha1 is a generated protocol buffer package.

It is generated from these files:
	authentication/v1alpha1/policy.proto

It has these top-level messages:
	None
	MutualTls
	Jwt
	AuthenticationMechanism
	AuthenticationPolicy
*/
package v1alpha1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import istio_routing_v1alpha2 "istio.io/api/routing/v1alpha2"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Place holder for None authentication params
type None struct {
}

func (m *None) Reset()                    { *m = None{} }
func (m *None) String() string            { return proto.CompactTextString(m) }
func (*None) ProtoMessage()               {}
func (*None) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Place holder for mTLS authentication params.
type MutualTls struct {
}

func (m *MutualTls) Reset()                    { *m = MutualTls{} }
func (m *MutualTls) String() string            { return proto.CompactTextString(m) }
func (*MutualTls) ProtoMessage()               {}
func (*MutualTls) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// JSON Web Token (JWT) token format for authentication as defined by
// https://tools.ietf.org/html/rfc7519. See [OAuth
// 2.0](https://tools.ietf.org/html/rfc6749) and [OIDC
// 1.0](http://openid.net/connect) for how this is used in the whole
// authentication flow.
//
// Example,
//
//     issuer: https://example.com
//     audiences:
//     - bookstore_android.apps.googleusercontent.com
//       bookstore_web.apps.googleusercontent.com
//     jwks_uri: https://example.com/.well-known/jwks.json
//
type Jwt struct {
	// Identifies the issuer that issued the JWT. See
	// https://tools.ietf.org/html/rfc7519#section-4.1.1
	// Usually a URL or an email address.
	//
	// Example: https://securetoken.google.com
	// Example: 1234567-compute@developer.gserviceaccount.com
	//
	Issuer string `protobuf:"bytes,1,opt,name=issuer" json:"issuer,omitempty"`
	// The list of JWT
	// [audiences](https://tools.ietf.org/html/rfc7519#section-4.1.3).
	// that are allowed to access. A JWT containing any of these
	// audiences will be accepted.
	//
	// The service name will be accepted if audiences is empty.
	//
	// Example:
	//
	//     audiences:
	//     - bookstore_android.apps.googleusercontent.com
	//       bookstore_web.apps.googleusercontent.com
	//
	Audiences []string `protobuf:"bytes,2,rep,name=audiences" json:"audiences,omitempty"`
	// URL of the provider's public key set to validate signature of the
	// JWT. See [OpenID
	// Discovery](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).
	//
	// Optional if the key set document can either (a) be retrieved from
	// [OpenID
	// Discovery](https://openid.net/specs/openid-connect-discovery-1_0.html) of
	// the issuer or (b) inferred from the email domain of the issuer (e.g. a
	// Google service account).
	//
	// Example: https://www.googleapis.com/oauth2/v1/certs
	JwksUri string `protobuf:"bytes,3,opt,name=jwks_uri,json=jwksUri" json:"jwks_uri,omitempty"`
	// JWT is sent in a request header. `header` represents the
	// header name.
	//
	// For example, if `header=x-goog-iap-jwt-assertion`, the header
	// format will be x-goog-iap-jwt-assertion: <JWT>.
	JwtHeaders []string `protobuf:"bytes,6,rep,name=jwt_headers,json=jwtHeaders" json:"jwt_headers,omitempty"`
	// JWT is sent in a query parameter. `query` represents the
	// query parameter name.
	//
	// For example, `query=jwt_token`.
	JwtParams []string `protobuf:"bytes,7,rep,name=jwt_params,json=jwtParams" json:"jwt_params,omitempty"`
}

func (m *Jwt) Reset()                    { *m = Jwt{} }
func (m *Jwt) String() string            { return proto.CompactTextString(m) }
func (*Jwt) ProtoMessage()               {}
func (*Jwt) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Jwt) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *Jwt) GetAudiences() []string {
	if m != nil {
		return m.Audiences
	}
	return nil
}

func (m *Jwt) GetJwksUri() string {
	if m != nil {
		return m.JwksUri
	}
	return ""
}

func (m *Jwt) GetJwtHeaders() []string {
	if m != nil {
		return m.JwtHeaders
	}
	return nil
}

func (m *Jwt) GetJwtParams() []string {
	if m != nil {
		return m.JwtParams
	}
	return nil
}

// AuthenticaitonMechanism defines one particular type of authentication, e.g
// mutual TLS, JWT etc, (no authentication is one type by itsefl).
// The type can be progammatically determine by checking the type of the
// "params" field.
type AuthenticationMechanism struct {
	// Types that are valid to be assigned to Params:
	//	*AuthenticationMechanism_None
	//	*AuthenticationMechanism_Mtls
	//	*AuthenticationMechanism_Jwt
	Params isAuthenticationMechanism_Params `protobuf_oneof:"params"`
}

func (m *AuthenticationMechanism) Reset()                    { *m = AuthenticationMechanism{} }
func (m *AuthenticationMechanism) String() string            { return proto.CompactTextString(m) }
func (*AuthenticationMechanism) ProtoMessage()               {}
func (*AuthenticationMechanism) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isAuthenticationMechanism_Params interface {
	isAuthenticationMechanism_Params()
}

type AuthenticationMechanism_None struct {
	None *None `protobuf:"bytes,1,opt,name=none,oneof"`
}
type AuthenticationMechanism_Mtls struct {
	Mtls *MutualTls `protobuf:"bytes,2,opt,name=mtls,oneof"`
}
type AuthenticationMechanism_Jwt struct {
	Jwt *Jwt `protobuf:"bytes,3,opt,name=jwt,oneof"`
}

func (*AuthenticationMechanism_None) isAuthenticationMechanism_Params() {}
func (*AuthenticationMechanism_Mtls) isAuthenticationMechanism_Params() {}
func (*AuthenticationMechanism_Jwt) isAuthenticationMechanism_Params()  {}

func (m *AuthenticationMechanism) GetParams() isAuthenticationMechanism_Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *AuthenticationMechanism) GetNone() *None {
	if x, ok := m.GetParams().(*AuthenticationMechanism_None); ok {
		return x.None
	}
	return nil
}

func (m *AuthenticationMechanism) GetMtls() *MutualTls {
	if x, ok := m.GetParams().(*AuthenticationMechanism_Mtls); ok {
		return x.Mtls
	}
	return nil
}

func (m *AuthenticationMechanism) GetJwt() *Jwt {
	if x, ok := m.GetParams().(*AuthenticationMechanism_Jwt); ok {
		return x.Jwt
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AuthenticationMechanism) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AuthenticationMechanism_OneofMarshaler, _AuthenticationMechanism_OneofUnmarshaler, _AuthenticationMechanism_OneofSizer, []interface{}{
		(*AuthenticationMechanism_None)(nil),
		(*AuthenticationMechanism_Mtls)(nil),
		(*AuthenticationMechanism_Jwt)(nil),
	}
}

func _AuthenticationMechanism_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AuthenticationMechanism)
	// params
	switch x := m.Params.(type) {
	case *AuthenticationMechanism_None:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.None); err != nil {
			return err
		}
	case *AuthenticationMechanism_Mtls:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Mtls); err != nil {
			return err
		}
	case *AuthenticationMechanism_Jwt:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Jwt); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AuthenticationMechanism.Params has unexpected type %T", x)
	}
	return nil
}

func _AuthenticationMechanism_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AuthenticationMechanism)
	switch tag {
	case 1: // params.none
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(None)
		err := b.DecodeMessage(msg)
		m.Params = &AuthenticationMechanism_None{msg}
		return true, err
	case 2: // params.mtls
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MutualTls)
		err := b.DecodeMessage(msg)
		m.Params = &AuthenticationMechanism_Mtls{msg}
		return true, err
	case 3: // params.jwt
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Jwt)
		err := b.DecodeMessage(msg)
		m.Params = &AuthenticationMechanism_Jwt{msg}
		return true, err
	default:
		return false, nil
	}
}

func _AuthenticationMechanism_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AuthenticationMechanism)
	// params
	switch x := m.Params.(type) {
	case *AuthenticationMechanism_None:
		s := proto.Size(x.None)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AuthenticationMechanism_Mtls:
		s := proto.Size(x.Mtls)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AuthenticationMechanism_Jwt:
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
//     apiVersion: config.istio.io/v1alpha1
//     kind: AuthenticationPolicy
//     metadata:
//       name: mTLS-enable
//       namespace: frod
//     spec:
//       match:
//       peers:
//       - mtls: {}
//
// Policy to enable mTLS, and use JWT for productpage:9000
//
//     apiVersion: config.istio.io/v1alpha1
//     kind: AuthenticationPolicy
//     metadata:
//       name: mTLS-enable
//       namespace: frod
//     spec:
//       match:
//       - name: productpage
//         port:
//           number: 9000
//       peers:
//       - mtls:
//       endUsers:
//       - jwt:
//           issuer: "https://securetoken.google.com"
//           audiences:
//           - "productpage"
//           jwksUri: "https://www.googleapis.com/oauth2/v1/certs"
//           locations:
//           - header: x-goog-iap-jwt-assertion
type AuthenticationPolicy struct {
	// List of destinations (workloads) that the policy should be applied on.
	// If empty, policy will be used on all destinations in the same namespace.
	Match []*istio_routing_v1alpha2.Destination `protobuf:"bytes,1,rep,name=match" json:"match,omitempty"`
	// List of credential that should be checked by peer authentication. They
	// will be validated in sequence, until the first one satisfied. If none of
	// the specified mechanism valid, the whole authentication should fail.
	// On the other hand, the first valid credential will be used to extract
	// peer identity (i.e the source.user attribute in the request to mixer).
	Peers []*AuthenticationMechanism `protobuf:"bytes,2,rep,name=peers" json:"peers,omitempty"`
	// Similar to above, but for end_user authentication, which will extract
	// request.auth.principal/audiences/presenter if authentication succeed.
	EndUsers []*AuthenticationMechanism `protobuf:"bytes,3,rep,name=end_users,json=endUsers" json:"end_users,omitempty"`
}

func (m *AuthenticationPolicy) Reset()                    { *m = AuthenticationPolicy{} }
func (m *AuthenticationPolicy) String() string            { return proto.CompactTextString(m) }
func (*AuthenticationPolicy) ProtoMessage()               {}
func (*AuthenticationPolicy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AuthenticationPolicy) GetMatch() []*istio_routing_v1alpha2.Destination {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *AuthenticationPolicy) GetPeers() []*AuthenticationMechanism {
	if m != nil {
		return m.Peers
	}
	return nil
}

func (m *AuthenticationPolicy) GetEndUsers() []*AuthenticationMechanism {
	if m != nil {
		return m.EndUsers
	}
	return nil
}

func init() {
	proto.RegisterType((*None)(nil), "istio.authentication.v1alpha1.None")
	proto.RegisterType((*MutualTls)(nil), "istio.authentication.v1alpha1.MutualTls")
	proto.RegisterType((*Jwt)(nil), "istio.authentication.v1alpha1.Jwt")
	proto.RegisterType((*AuthenticationMechanism)(nil), "istio.authentication.v1alpha1.AuthenticationMechanism")
	proto.RegisterType((*AuthenticationPolicy)(nil), "istio.authentication.v1alpha1.AuthenticationPolicy")
}

func init() { proto.RegisterFile("authentication/v1alpha1/policy.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x41, 0x6e, 0xd4, 0x30,
	0x14, 0x86, 0x9b, 0x66, 0x9a, 0x4e, 0x5e, 0x76, 0x16, 0x82, 0x80, 0xa8, 0x18, 0x42, 0x85, 0x66,
	0x95, 0xa8, 0x41, 0xaa, 0xd4, 0x0d, 0x12, 0x15, 0x8b, 0x51, 0x45, 0x51, 0x15, 0x98, 0x0d, 0x9b,
	0xc8, 0x24, 0x16, 0x71, 0x48, 0xec, 0xc8, 0x7e, 0x26, 0xe2, 0x22, 0x1c, 0x8d, 0x93, 0x70, 0x00,
	0x64, 0x67, 0x06, 0x98, 0xc5, 0x30, 0x8b, 0x2e, 0xdf, 0x93, 0xbf, 0x4f, 0xfe, 0x7f, 0x1b, 0xce,
	0xa9, 0xc1, 0x86, 0x09, 0xe4, 0x15, 0x45, 0x2e, 0x45, 0xf6, 0xed, 0x82, 0x76, 0x43, 0x43, 0x2f,
	0xb2, 0x41, 0x76, 0xbc, 0xfa, 0x9e, 0x0e, 0x4a, 0xa2, 0x24, 0x67, 0x5c, 0x23, 0x97, 0xe9, 0xee,
	0xd9, 0x74, 0x7b, 0xf6, 0xc9, 0x73, 0x25, 0x0d, 0x72, 0xf1, 0x65, 0x4b, 0xe7, 0x99, 0x5d, 0xb0,
	0x52, 0x99, 0x8e, 0x4d, 0x86, 0x24, 0x80, 0xd9, 0x7b, 0x29, 0x58, 0x12, 0x41, 0x78, 0x6b, 0xd0,
	0xd0, 0xee, 0x63, 0xa7, 0x93, 0x1f, 0x1e, 0xf8, 0x37, 0x23, 0x92, 0x87, 0x10, 0x70, 0xad, 0x0d,
	0x53, 0xb1, 0xb7, 0xf0, 0x96, 0x61, 0xb1, 0x99, 0xc8, 0x53, 0x08, 0xa9, 0xa9, 0x39, 0x13, 0x15,
	0xd3, 0xf1, 0xf1, 0xc2, 0x5f, 0x86, 0xc5, 0xdf, 0x05, 0x79, 0x0c, 0xf3, 0x76, 0xfc, 0xaa, 0x4b,
	0xa3, 0x78, 0xec, 0x3b, 0xee, 0xd4, 0xce, 0x6b, 0xc5, 0xc9, 0x33, 0x88, 0xda, 0x11, 0xcb, 0x86,
	0xd1, 0x9a, 0x29, 0x1d, 0x07, 0x0e, 0x85, 0x76, 0xc4, 0xd5, 0xb4, 0x21, 0x67, 0x60, 0xa7, 0x72,
	0xa0, 0x8a, 0xf6, 0x3a, 0x3e, 0x9d, 0xd4, 0xed, 0x88, 0x77, 0x6e, 0x91, 0xfc, 0xf4, 0xe0, 0xd1,
	0x9b, 0x9d, 0xb0, 0xb7, 0xac, 0x6a, 0xa8, 0xe0, 0xba, 0x27, 0x57, 0x30, 0x13, 0x52, 0x30, 0x77,
	0xd5, 0x28, 0x7f, 0x91, 0xfe, 0xb7, 0x9a, 0xd4, 0x86, 0x5e, 0x1d, 0x15, 0x0e, 0x21, 0xaf, 0x61,
	0xd6, 0x63, 0x67, 0xa3, 0x58, 0x74, 0x79, 0x00, 0xfd, 0xd3, 0x93, 0xe5, 0x2d, 0x47, 0x2e, 0xc1,
	0x6f, 0x47, 0x74, 0x61, 0xa3, 0x3c, 0x39, 0x80, 0xdf, 0x8c, 0xb8, 0x3a, 0x2a, 0x2c, 0x70, 0x3d,
	0x87, 0x60, 0x4a, 0x9a, 0xfc, 0xf2, 0xe0, 0xc1, 0x6e, 0xb0, 0x3b, 0xf7, 0xce, 0xe4, 0x0a, 0x4e,
	0x7a, 0x8a, 0x55, 0x13, 0x7b, 0x0b, 0xff, 0x9f, 0x58, 0x9b, 0x87, 0xdd, 0x5a, 0xf3, 0xf4, 0x2d,
	0xd3, 0xc8, 0x85, 0x23, 0x8b, 0x89, 0x20, 0xef, 0xe0, 0x64, 0x60, 0xb6, 0xe6, 0x63, 0x87, 0x5e,
	0x1e, 0xb8, 0xd7, 0x9e, 0x5e, 0x8b, 0x49, 0x42, 0x3e, 0x40, 0xc8, 0x44, 0x5d, 0x1a, 0x6d, 0x8d,
	0xfe, 0xbd, 0x8c, 0x73, 0x26, 0xea, 0xb5, 0xf5, 0x5c, 0xbf, 0xfc, 0x74, 0x3e, 0x29, 0xb8, 0xcc,
	0xe8, 0xc0, 0xb3, 0x3d, 0x9f, 0xfe, 0x73, 0xe0, 0x3e, 0xeb, 0xab, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xb4, 0x96, 0xae, 0xb4, 0x16, 0x03, 0x00, 0x00,
}
