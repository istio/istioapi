// Code generated by protoc-gen-gogo.
// source: mixer/v1/config/descriptor/quota_descriptor.proto
// DO NOT EDIT!

package istio_mixer_v1_config_descriptor

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strings "strings"
import reflect "reflect"
import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Configuration state for a particular quota.
//
// Quotas are similar to metrics, except that they are mutated through method
// calls and there are limits on the allowed values.
// The descriptor below lets you define a quota and indicate the maximum
// amount values of this quota are allowed to hold.
//
// A given quota is described by a set of attributes. These attributes represent
// the different dimensions to associate with the quota. A given quota holds a
// unique value for potentially any combination of these attributes.
type QuotaDescriptor struct {
	// Required. The name of this descriptor.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. A concise name for the quota which can be displayed in user
	// interfaces.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Optional. A description of the quota which can be used in documentation.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// The set of labels that are necessary to describe a specific value cell
	// for a quota of this type.
	Labels map[string]ValueType `protobuf:"bytes,4,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=istio.mixer.v1.config.descriptor.ValueType"`
	// Indicates whether the quota represents a rate limit or represents a
	// resource quota.
	RateLimit bool `protobuf:"varint,5,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
}

func (m *QuotaDescriptor) Reset()                    { *m = QuotaDescriptor{} }
func (*QuotaDescriptor) ProtoMessage()               {}
func (*QuotaDescriptor) Descriptor() ([]byte, []int) { return fileDescriptorQuotaDescriptor, []int{0} }

func (m *QuotaDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *QuotaDescriptor) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *QuotaDescriptor) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *QuotaDescriptor) GetLabels() map[string]ValueType {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *QuotaDescriptor) GetRateLimit() bool {
	if m != nil {
		return m.RateLimit
	}
	return false
}

func init() {
	proto.RegisterType((*QuotaDescriptor)(nil), "istio.mixer.v1.config.descriptor.QuotaDescriptor")
}
func (this *QuotaDescriptor) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*QuotaDescriptor)
	if !ok {
		that2, ok := that.(QuotaDescriptor)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.DisplayName != that1.DisplayName {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if len(this.Labels) != len(that1.Labels) {
		return false
	}
	for i := range this.Labels {
		if this.Labels[i] != that1.Labels[i] {
			return false
		}
	}
	if this.RateLimit != that1.RateLimit {
		return false
	}
	return true
}
func (this *QuotaDescriptor) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&istio_mixer_v1_config_descriptor.QuotaDescriptor{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "DisplayName: "+fmt.Sprintf("%#v", this.DisplayName)+",\n")
	s = append(s, "Description: "+fmt.Sprintf("%#v", this.Description)+",\n")
	keysForLabels := make([]string, 0, len(this.Labels))
	for k, _ := range this.Labels {
		keysForLabels = append(keysForLabels, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForLabels)
	mapStringForLabels := "map[string]ValueType{"
	for _, k := range keysForLabels {
		mapStringForLabels += fmt.Sprintf("%#v: %#v,", k, this.Labels[k])
	}
	mapStringForLabels += "}"
	if this.Labels != nil {
		s = append(s, "Labels: "+mapStringForLabels+",\n")
	}
	s = append(s, "RateLimit: "+fmt.Sprintf("%#v", this.RateLimit)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringQuotaDescriptor(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *QuotaDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuotaDescriptor) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintQuotaDescriptor(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.DisplayName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintQuotaDescriptor(dAtA, i, uint64(len(m.DisplayName)))
		i += copy(dAtA[i:], m.DisplayName)
	}
	if len(m.Description) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintQuotaDescriptor(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	}
	if len(m.Labels) > 0 {
		for k, _ := range m.Labels {
			dAtA[i] = 0x22
			i++
			v := m.Labels[k]
			mapSize := 1 + len(k) + sovQuotaDescriptor(uint64(len(k))) + 1 + sovQuotaDescriptor(uint64(v))
			i = encodeVarintQuotaDescriptor(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintQuotaDescriptor(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x10
			i++
			i = encodeVarintQuotaDescriptor(dAtA, i, uint64(v))
		}
	}
	if m.RateLimit {
		dAtA[i] = 0x28
		i++
		if m.RateLimit {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeFixed64QuotaDescriptor(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32QuotaDescriptor(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintQuotaDescriptor(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *QuotaDescriptor) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovQuotaDescriptor(uint64(l))
	}
	l = len(m.DisplayName)
	if l > 0 {
		n += 1 + l + sovQuotaDescriptor(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovQuotaDescriptor(uint64(l))
	}
	if len(m.Labels) > 0 {
		for k, v := range m.Labels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovQuotaDescriptor(uint64(len(k))) + 1 + sovQuotaDescriptor(uint64(v))
			n += mapEntrySize + 1 + sovQuotaDescriptor(uint64(mapEntrySize))
		}
	}
	if m.RateLimit {
		n += 2
	}
	return n
}

func sovQuotaDescriptor(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozQuotaDescriptor(x uint64) (n int) {
	return sovQuotaDescriptor(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *QuotaDescriptor) String() string {
	if this == nil {
		return "nil"
	}
	keysForLabels := make([]string, 0, len(this.Labels))
	for k, _ := range this.Labels {
		keysForLabels = append(keysForLabels, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForLabels)
	mapStringForLabels := "map[string]ValueType{"
	for _, k := range keysForLabels {
		mapStringForLabels += fmt.Sprintf("%v: %v,", k, this.Labels[k])
	}
	mapStringForLabels += "}"
	s := strings.Join([]string{`&QuotaDescriptor{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`DisplayName:` + fmt.Sprintf("%v", this.DisplayName) + `,`,
		`Description:` + fmt.Sprintf("%v", this.Description) + `,`,
		`Labels:` + mapStringForLabels + `,`,
		`RateLimit:` + fmt.Sprintf("%v", this.RateLimit) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringQuotaDescriptor(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *QuotaDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuotaDescriptor
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
			return fmt.Errorf("proto: QuotaDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuotaDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuotaDescriptor
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
				return ErrInvalidLengthQuotaDescriptor
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisplayName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuotaDescriptor
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
				return ErrInvalidLengthQuotaDescriptor
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DisplayName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuotaDescriptor
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
				return ErrInvalidLengthQuotaDescriptor
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuotaDescriptor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuotaDescriptor
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuotaDescriptor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuotaDescriptor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthQuotaDescriptor
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.Labels == nil {
				m.Labels = make(map[string]ValueType)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQuotaDescriptor
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var mapvalue ValueType
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQuotaDescriptor
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					mapvalue |= (ValueType(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Labels[mapkey] = mapvalue
			} else {
				var mapvalue ValueType
				m.Labels[mapkey] = mapvalue
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RateLimit", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuotaDescriptor
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
			m.RateLimit = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipQuotaDescriptor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuotaDescriptor
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
func skipQuotaDescriptor(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuotaDescriptor
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
					return 0, ErrIntOverflowQuotaDescriptor
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
					return 0, ErrIntOverflowQuotaDescriptor
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
				return 0, ErrInvalidLengthQuotaDescriptor
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowQuotaDescriptor
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
				next, err := skipQuotaDescriptor(dAtA[start:])
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
	ErrInvalidLengthQuotaDescriptor = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuotaDescriptor   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("mixer/v1/config/descriptor/quota_descriptor.proto", fileDescriptorQuotaDescriptor)
}

var fileDescriptorQuotaDescriptor = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xcc, 0xcd, 0xac, 0x48,
	0x2d, 0xd2, 0x2f, 0x33, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0x49, 0x2d, 0x4e,
	0x2e, 0xca, 0x2c, 0x28, 0xc9, 0x2f, 0xd2, 0x2f, 0x2c, 0xcd, 0x2f, 0x49, 0x8c, 0x47, 0x08, 0xe8,
	0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x29, 0x64, 0x16, 0x97, 0x64, 0xe6, 0xeb, 0x81, 0x35, 0xea,
	0x95, 0x19, 0xea, 0x41, 0x34, 0xea, 0x21, 0xd4, 0x49, 0x69, 0xe3, 0x31, 0xb4, 0x2c, 0x31, 0xa7,
	0x34, 0x35, 0xbe, 0xa4, 0xb2, 0x20, 0x15, 0x62, 0x9c, 0xd2, 0x09, 0x26, 0x2e, 0xfe, 0x40, 0x90,
	0x4d, 0x2e, 0x70, 0x45, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x22, 0x17, 0x4f, 0x4a, 0x66, 0x71, 0x41, 0x4e, 0x62, 0x65,
	0x3c, 0x58, 0x8e, 0x09, 0x2c, 0xc7, 0x0d, 0x15, 0xf3, 0x03, 0x29, 0x51, 0xe0, 0xe2, 0x86, 0xd9,
	0x94, 0x99, 0x9f, 0x27, 0xc1, 0x0c, 0x55, 0x81, 0x10, 0x12, 0x0a, 0xe5, 0x62, 0xcb, 0x49, 0x4c,
	0x4a, 0xcd, 0x29, 0x96, 0x60, 0x51, 0x60, 0xd6, 0xe0, 0x36, 0xb2, 0xd5, 0x23, 0xe4, 0x19, 0x3d,
	0x34, 0xb7, 0xe9, 0xf9, 0x80, 0xf5, 0xbb, 0xe6, 0x95, 0x14, 0x55, 0x06, 0x41, 0x0d, 0x13, 0x92,
	0xe5, 0xe2, 0x2a, 0x4a, 0x2c, 0x49, 0x8d, 0xcf, 0xc9, 0xcc, 0xcd, 0x2c, 0x91, 0x60, 0x55, 0x60,
	0xd4, 0xe0, 0x08, 0xe2, 0x04, 0x89, 0xf8, 0x80, 0x04, 0xa4, 0xd2, 0xb8, 0xb8, 0x91, 0x74, 0x09,
	0x09, 0x70, 0x31, 0x67, 0xa7, 0x56, 0x42, 0x3d, 0x07, 0x62, 0x0a, 0x39, 0x72, 0xb1, 0x82, 0xc3,
	0x05, 0xec, 0x29, 0x3e, 0x23, 0x6d, 0xc2, 0xae, 0x0a, 0x03, 0x29, 0x0f, 0xa9, 0x2c, 0x48, 0x0d,
	0x82, 0xe8, 0xb4, 0x62, 0xb2, 0x60, 0x74, 0xd2, 0xb9, 0xf0, 0x50, 0x8e, 0xe1, 0xc6, 0x43, 0x39,
	0x86, 0x0f, 0x0f, 0xe5, 0x18, 0x1b, 0x1e, 0xc9, 0x31, 0xae, 0x78, 0x24, 0xc7, 0x78, 0xe2, 0x91,
	0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0xbe, 0x78, 0x24, 0xc7, 0xf0, 0xe1,
	0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x49, 0x6c, 0xe0, 0xf0, 0x37, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xd3, 0x07, 0x23, 0xa3, 0x03, 0x02, 0x00, 0x00,
}
