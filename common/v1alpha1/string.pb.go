// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: common/v1alpha1/string.proto

/*
	Package v1alpha1 is a generated protocol buffer package.

	This package defines common string protos.

	It is generated from these files:
		common/v1alpha1/string.proto

	It has these top-level messages:
		StringMatch
*/
package v1alpha1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Describes how to match a given string. Match is case-sensitive.
type StringMatch struct {
	// Types that are valid to be assigned to MatchType:
	//	*StringMatch_Exact
	//	*StringMatch_Prefix
	//	*StringMatch_Suffix
	//	*StringMatch_Regex
	MatchType isStringMatch_MatchType `protobuf_oneof:"match_type"`
	// If true, the match result will be inverted. Defaults to false.
	InvertMatch bool `protobuf:"varint,5,opt,name=invert_match,json=invertMatch,proto3" json:"invert_match,omitempty"`
}

func (m *StringMatch) Reset()                    { *m = StringMatch{} }
func (m *StringMatch) String() string            { return proto.CompactTextString(m) }
func (*StringMatch) ProtoMessage()               {}
func (*StringMatch) Descriptor() ([]byte, []int) { return fileDescriptorString, []int{0} }

type isStringMatch_MatchType interface {
	isStringMatch_MatchType()
	MarshalTo([]byte) (int, error)
	Size() int
}

type StringMatch_Exact struct {
	Exact string `protobuf:"bytes,1,opt,name=exact,proto3,oneof"`
}
type StringMatch_Prefix struct {
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3,oneof"`
}
type StringMatch_Suffix struct {
	Suffix string `protobuf:"bytes,3,opt,name=suffix,proto3,oneof"`
}
type StringMatch_Regex struct {
	Regex string `protobuf:"bytes,4,opt,name=regex,proto3,oneof"`
}

func (*StringMatch_Exact) isStringMatch_MatchType()  {}
func (*StringMatch_Prefix) isStringMatch_MatchType() {}
func (*StringMatch_Suffix) isStringMatch_MatchType() {}
func (*StringMatch_Regex) isStringMatch_MatchType()  {}

func (m *StringMatch) GetMatchType() isStringMatch_MatchType {
	if m != nil {
		return m.MatchType
	}
	return nil
}

func (m *StringMatch) GetExact() string {
	if x, ok := m.GetMatchType().(*StringMatch_Exact); ok {
		return x.Exact
	}
	return ""
}

func (m *StringMatch) GetPrefix() string {
	if x, ok := m.GetMatchType().(*StringMatch_Prefix); ok {
		return x.Prefix
	}
	return ""
}

func (m *StringMatch) GetSuffix() string {
	if x, ok := m.GetMatchType().(*StringMatch_Suffix); ok {
		return x.Suffix
	}
	return ""
}

func (m *StringMatch) GetRegex() string {
	if x, ok := m.GetMatchType().(*StringMatch_Regex); ok {
		return x.Regex
	}
	return ""
}

func (m *StringMatch) GetInvertMatch() bool {
	if m != nil {
		return m.InvertMatch
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StringMatch) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StringMatch_OneofMarshaler, _StringMatch_OneofUnmarshaler, _StringMatch_OneofSizer, []interface{}{
		(*StringMatch_Exact)(nil),
		(*StringMatch_Prefix)(nil),
		(*StringMatch_Suffix)(nil),
		(*StringMatch_Regex)(nil),
	}
}

func _StringMatch_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StringMatch)
	// match_type
	switch x := m.MatchType.(type) {
	case *StringMatch_Exact:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Exact)
	case *StringMatch_Prefix:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Prefix)
	case *StringMatch_Suffix:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Suffix)
	case *StringMatch_Regex:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Regex)
	case nil:
	default:
		return fmt.Errorf("StringMatch.MatchType has unexpected type %T", x)
	}
	return nil
}

func _StringMatch_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StringMatch)
	switch tag {
	case 1: // match_type.exact
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.MatchType = &StringMatch_Exact{x}
		return true, err
	case 2: // match_type.prefix
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.MatchType = &StringMatch_Prefix{x}
		return true, err
	case 3: // match_type.suffix
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.MatchType = &StringMatch_Suffix{x}
		return true, err
	case 4: // match_type.regex
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.MatchType = &StringMatch_Regex{x}
		return true, err
	default:
		return false, nil
	}
}

func _StringMatch_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StringMatch)
	// match_type
	switch x := m.MatchType.(type) {
	case *StringMatch_Exact:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Exact)))
		n += len(x.Exact)
	case *StringMatch_Prefix:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Prefix)))
		n += len(x.Prefix)
	case *StringMatch_Suffix:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Suffix)))
		n += len(x.Suffix)
	case *StringMatch_Regex:
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Regex)))
		n += len(x.Regex)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*StringMatch)(nil), "istio.common.v1alpha1.StringMatch")
}
func (m *StringMatch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StringMatch) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.MatchType != nil {
		nn1, err := m.MatchType.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	if m.InvertMatch {
		dAtA[i] = 0x28
		i++
		if m.InvertMatch {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *StringMatch_Exact) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0xa
	i++
	i = encodeVarintString(dAtA, i, uint64(len(m.Exact)))
	i += copy(dAtA[i:], m.Exact)
	return i, nil
}
func (m *StringMatch_Prefix) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x12
	i++
	i = encodeVarintString(dAtA, i, uint64(len(m.Prefix)))
	i += copy(dAtA[i:], m.Prefix)
	return i, nil
}
func (m *StringMatch_Suffix) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x1a
	i++
	i = encodeVarintString(dAtA, i, uint64(len(m.Suffix)))
	i += copy(dAtA[i:], m.Suffix)
	return i, nil
}
func (m *StringMatch_Regex) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x22
	i++
	i = encodeVarintString(dAtA, i, uint64(len(m.Regex)))
	i += copy(dAtA[i:], m.Regex)
	return i, nil
}
func encodeVarintString(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *StringMatch) Size() (n int) {
	var l int
	_ = l
	if m.MatchType != nil {
		n += m.MatchType.Size()
	}
	if m.InvertMatch {
		n += 2
	}
	return n
}

func (m *StringMatch_Exact) Size() (n int) {
	var l int
	_ = l
	l = len(m.Exact)
	n += 1 + l + sovString(uint64(l))
	return n
}
func (m *StringMatch_Prefix) Size() (n int) {
	var l int
	_ = l
	l = len(m.Prefix)
	n += 1 + l + sovString(uint64(l))
	return n
}
func (m *StringMatch_Suffix) Size() (n int) {
	var l int
	_ = l
	l = len(m.Suffix)
	n += 1 + l + sovString(uint64(l))
	return n
}
func (m *StringMatch_Regex) Size() (n int) {
	var l int
	_ = l
	l = len(m.Regex)
	n += 1 + l + sovString(uint64(l))
	return n
}

func sovString(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozString(x uint64) (n int) {
	return sovString(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StringMatch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowString
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StringMatch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StringMatch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exact", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowString
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthString
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MatchType = &StringMatch_Exact{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowString
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthString
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MatchType = &StringMatch_Prefix{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Suffix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowString
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthString
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MatchType = &StringMatch_Suffix{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Regex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowString
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthString
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MatchType = &StringMatch_Regex{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InvertMatch", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowString
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.InvertMatch = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipString(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthString
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipString(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowString
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
					return 0, ErrIntOverflowString
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
					return 0, ErrIntOverflowString
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthString
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowString
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
				next, err := skipString(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthString = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowString   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("common/v1alpha1/string.proto", fileDescriptorString) }

var fileDescriptorString = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2f, 0x33, 0x4c, 0xcc, 0x29, 0xc8, 0x48, 0x34, 0xd4, 0x2f, 0x2e, 0x29, 0xca,
	0xcc, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xcd, 0x2c, 0x2e, 0xc9, 0xcc, 0xd7,
	0x83, 0xa8, 0xd1, 0x83, 0xa9, 0x51, 0x5a, 0xc4, 0xc8, 0xc5, 0x1d, 0x0c, 0x56, 0xe7, 0x9b, 0x58,
	0x92, 0x9c, 0x21, 0x24, 0xc6, 0xc5, 0x9a, 0x5a, 0x91, 0x98, 0x5c, 0x22, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0xe9, 0xc1, 0x10, 0x04, 0xe1, 0x0a, 0x49, 0x70, 0xb1, 0x15, 0x14, 0xa5, 0xa6, 0x65, 0x56,
	0x48, 0x30, 0x41, 0x25, 0xa0, 0x7c, 0x90, 0x4c, 0x71, 0x69, 0x1a, 0x48, 0x86, 0x19, 0x26, 0x03,
	0xe1, 0x83, 0xcc, 0x2a, 0x4a, 0x4d, 0x4f, 0xad, 0x90, 0x60, 0x81, 0x99, 0x05, 0xe6, 0x0a, 0x29,
	0x72, 0xf1, 0x64, 0xe6, 0x95, 0xa5, 0x16, 0x95, 0xc4, 0xe7, 0x82, 0xec, 0x94, 0x60, 0x55, 0x60,
	0xd4, 0xe0, 0x08, 0xe2, 0x86, 0x88, 0x81, 0x9d, 0xe1, 0xc4, 0xc3, 0xc5, 0x05, 0x96, 0x8b, 0x2f,
	0xa9, 0x2c, 0x48, 0x75, 0xd2, 0x3a, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f,
	0xe4, 0x18, 0xa3, 0x64, 0x20, 0x3e, 0xc9, 0xcc, 0xd7, 0x4f, 0x2c, 0xc8, 0xd4, 0x47, 0xf3, 0x74,
	0x12, 0x1b, 0xd8, 0xbb, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfc, 0xac, 0x8e, 0x13, 0x0e,
	0x01, 0x00, 0x00,
}
