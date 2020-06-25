// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: security/v1beta1/request_authentication.proto

package v1beta1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	v1beta1 "istio.io/api/type/v1beta1"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// RequestAuthentication defines what request authentication methods are supported by a workload.
// It will reject a request if the request contains invalid authentication information, based on the
// configured authentication rules. A request that does not contain any authentication credentials
// will be accepted but will not have any authenticated identity. To restrict access to authenticated
// requests only, this should be accompanied by an authorization rule.
// Examples:
//
// - Require JWT for all request for workloads that have label `app:httpbin`
//
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: RequestAuthentication
// metadata:
//   name: httpbin
//   namespace: foo
// spec:
//   selector:
//     matchLabels:
//       app: httpbin
//   jwtRules:
//   - issuer: "issuer-foo"
//     jwksUri: https://example.com/.well-known/jwks.json
// ---
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//   name: httpbin
//   namespace: foo
// spec:
//   selector:
//     matchLabels:
//       app: httpbin
//   rules:
//   - from:
//     - source:
//         requestPrincipals: ["*"]
// ```
//
// - The next example shows how to set a different JWT requirement for a different `host`. The `RequestAuthentication`
// declares it can accept JWTs issued by either `issuer-foo` or `issuer-bar` (the public key set is implicitly
// set from the OpenID Connect spec).
//
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: RequestAuthentication
// metadata:
//   name: httpbin
//   namespace: foo
// spec:
//   selector:
//     matchLabels:
//       app: httpbin
//   jwtRules:
//   - issuer: "issuer-foo"
//   - issuer: "issuer-bar"
// ---
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//   name: httpbin
//   namespace: foo
// spec:
//   selector:
//     matchLabels:
//       app: httpbin
//  rules:
//  - from:
//    - source:
//        requestPrincipals: ["issuer-foo/*"]
//    to:
//      hosts: ["example.com"]
//  - from:
//    - source:
//        requestPrincipals: ["issuer-bar/*"]
//    to:
//      hosts: ["another-host.com"]
// ```
//
// - You can fine tune the authorization policy to set different requirement per path. For example,
// to require JWT on all paths, except /healthz, the same `RequestAuthentication` can be used, but the
// authorization policy could be:
//
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//   name: httpbin
//   namespace: foo
// spec:
//   selector:
//     matchLabels:
//       app: httpbin
//  rules:
//  - from:
//    - source:
//        requestPrincipals: ["*"]
//  - to:
//    - operation:
//        paths: ["/healthz]
// ```
//
// <!-- crd generation tags
// +cue-gen:RequestAuthentication:groupName:security.istio.io
// +cue-gen:RequestAuthentication:version:v1beta1
// +cue-gen:RequestAuthentication:storageVersion
// +cue-gen:RequestAuthentication:annotations:helm.sh/resource-policy=keep
// +cue-gen:RequestAuthentication:labels:app=istio-pilot,chart=istio,istio=security,heritage=Tiller,release=istio
// +cue-gen:RequestAuthentication:subresource:status
// +cue-gen:RequestAuthentication:scope:Namespaced
// +cue-gen:RequestAuthentication:resource:categories=istio-io,security-istio-io,shortNames=ra
// +cue-gen:RequestAuthentication:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=security.istio.io/v1beta1
// +genclient
// +k8s:deepcopy-gen=true
// -->
type RequestAuthentication struct {
	// The selector determines the workloads to apply the RequestAuthentication on.
	// If not set, the policy will be applied to all workloads in the same namespace as the policy.
	Selector *v1beta1.WorkloadSelector `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// Define the list of JWTs that can be validated at the selected workloads' proxy. A valid token
	// will be used to extract the authenticated identity.
	// Each rule will be activated only when a token is presented at the location recorgnized by the
	// rule. The token will be validated based on the JWT rule config. If validation fails, the request will
	// be rejected.
	// Note: if more than one token is presented (at different locations), the output principal is nondeterministic.
	JwtRules             []*JWTRule `protobuf:"bytes,2,rep,name=jwt_rules,json=jwtRules,proto3" json:"jwt_rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RequestAuthentication) Reset()         { *m = RequestAuthentication{} }
func (m *RequestAuthentication) String() string { return proto.CompactTextString(m) }
func (*RequestAuthentication) ProtoMessage()    {}
func (*RequestAuthentication) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a3374592471772b, []int{0}
}
func (m *RequestAuthentication) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestAuthentication) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestAuthentication.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestAuthentication) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestAuthentication.Merge(m, src)
}
func (m *RequestAuthentication) XXX_Size() int {
	return m.Size()
}
func (m *RequestAuthentication) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestAuthentication.DiscardUnknown(m)
}

var xxx_messageInfo_RequestAuthentication proto.InternalMessageInfo

func (m *RequestAuthentication) GetSelector() *v1beta1.WorkloadSelector {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (m *RequestAuthentication) GetJwtRules() []*JWTRule {
	if m != nil {
		return m.JwtRules
	}
	return nil
}

func init() {
	proto.RegisterType((*RequestAuthentication)(nil), "istio.security.v1beta1.RequestAuthentication")
}

func init() {
	proto.RegisterFile("security/v1beta1/request_authentication.proto", fileDescriptor_4a3374592471772b)
}

var fileDescriptor_4a3374592471772b = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2d, 0x4e, 0x4d, 0x2e,
	0x2d, 0xca, 0x2c, 0xa9, 0xd4, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x2f, 0x4a, 0x2d,
	0x2c, 0x4d, 0x2d, 0x2e, 0x89, 0x4f, 0x2c, 0x2d, 0xc9, 0x48, 0xcd, 0x2b, 0xc9, 0x4c, 0x4e, 0x2c,
	0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xcb, 0x2c, 0x2e, 0xc9, 0xcc,
	0xd7, 0x83, 0x69, 0xd2, 0x83, 0x6a, 0x92, 0x92, 0x2e, 0xa9, 0x2c, 0x48, 0x85, 0x1b, 0x51, 0x9c,
	0x9a, 0x93, 0x9a, 0x5c, 0x92, 0x5f, 0x04, 0xd1, 0x24, 0x25, 0x85, 0x61, 0x47, 0x56, 0x79, 0x09,
	0x44, 0x4e, 0x69, 0x3a, 0x23, 0x97, 0x68, 0x10, 0xc4, 0x46, 0x47, 0x14, 0x0b, 0x85, 0x1c, 0xb8,
	0x38, 0x60, 0xe6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0xa9, 0xe8, 0x41, 0x6c, 0x07, 0xd9,
	0x05, 0xb3, 0x59, 0x2f, 0x3c, 0xbf, 0x28, 0x3b, 0x27, 0x3f, 0x31, 0x25, 0x18, 0xaa, 0x36, 0x08,
	0xae, 0x4b, 0xc8, 0x86, 0x8b, 0x33, 0xab, 0xbc, 0x24, 0xbe, 0xa8, 0x34, 0x27, 0xb5, 0x58, 0x82,
	0x49, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x5e, 0x0f, 0xbb, 0x07, 0xf4, 0xbc, 0xc2, 0x43, 0x82, 0x4a,
	0x73, 0x52, 0x83, 0x38, 0xb2, 0xca, 0x4b, 0x40, 0x8c, 0x62, 0x27, 0xed, 0x13, 0x8f, 0xe4, 0x18,
	0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x31, 0x4a, 0x16, 0xa2, 0x2f, 0x33, 0x5f, 0x3f,
	0xb1, 0x20, 0x53, 0x1f, 0xdd, 0x43, 0x49, 0x6c, 0x60, 0xdf, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xe1, 0xe0, 0xd4, 0xfd, 0x4f, 0x01, 0x00, 0x00,
}

func (m *RequestAuthentication) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestAuthentication) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestAuthentication) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.JwtRules) > 0 {
		for iNdEx := len(m.JwtRules) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.JwtRules[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRequestAuthentication(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Selector != nil {
		{
			size, err := m.Selector.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRequestAuthentication(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRequestAuthentication(dAtA []byte, offset int, v uint64) int {
	offset -= sovRequestAuthentication(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RequestAuthentication) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Selector != nil {
		l = m.Selector.Size()
		n += 1 + l + sovRequestAuthentication(uint64(l))
	}
	if len(m.JwtRules) > 0 {
		for _, e := range m.JwtRules {
			l = e.Size()
			n += 1 + l + sovRequestAuthentication(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRequestAuthentication(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRequestAuthentication(x uint64) (n int) {
	return sovRequestAuthentication(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RequestAuthentication) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRequestAuthentication
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RequestAuthentication: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestAuthentication: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Selector", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestAuthentication
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRequestAuthentication
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRequestAuthentication
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Selector == nil {
				m.Selector = &v1beta1.WorkloadSelector{}
			}
			if err := m.Selector.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JwtRules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestAuthentication
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRequestAuthentication
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRequestAuthentication
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JwtRules = append(m.JwtRules, &JWTRule{})
			if err := m.JwtRules[len(m.JwtRules)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRequestAuthentication(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRequestAuthentication
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRequestAuthentication
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRequestAuthentication(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRequestAuthentication
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRequestAuthentication
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRequestAuthentication
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthRequestAuthentication
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthRequestAuthentication
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRequestAuthentication
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRequestAuthentication(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthRequestAuthentication
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRequestAuthentication = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRequestAuthentication   = fmt.Errorf("proto: integer overflow")
)
