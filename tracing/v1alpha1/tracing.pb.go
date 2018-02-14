// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tracing/v1alpha1/tracing.proto

/*
Package v1alpha1 is a generated protocol buffer package.

Provides configuration parameters for Istio tracing.

It is generated from these files:
	tracing/v1alpha1/tracing.proto

It has these top-level messages:
	TracingRule
	TraceCollector
	ZipkinCollector
	LightStepCollector
	Percent
	TraceSampling
*/
package v1alpha1

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

// Controls the mode of sampling.
type TraceSampling_Mode int32

const (
	// Do NOT generate trace spans.
	TraceSampling_NONE TraceSampling_Mode = 0
	// Use the pre-configured minimal trace span generation settings.
	// This will generate trace spans for 1% of all requests (and 100% for
	// client-directed forced traces).
	TraceSampling_MINIMAL TraceSampling_Mode = 1
	// Use a custom sampling configuration.
	TraceSampling_CUSTOM TraceSampling_Mode = 2
	// Generate trace spans for 100% of requests, client-requested or
	// otherwise.
	TraceSampling_ALL TraceSampling_Mode = 3
)

var TraceSampling_Mode_name = map[int32]string{
	0: "NONE",
	1: "MINIMAL",
	2: "CUSTOM",
	3: "ALL",
}
var TraceSampling_Mode_value = map[string]int32{
	"NONE":    0,
	"MINIMAL": 1,
	"CUSTOM":  2,
	"ALL":     3,
}

func (x TraceSampling_Mode) String() string {
	return proto.EnumName(TraceSampling_Mode_name, int32(x))
}
func (TraceSampling_Mode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

// TracingRule controls Istio mesh trace collection configuration
// parameters. It includes the configuration for the tracing backends as well
// as the rates at which distributed tracing data should be generated.
//
// Example configuration:
// ```yaml
// apiVersion: tracing.istio.io/v1alpha1
// kind: TracingConfiguration
// metadata:
//   name: meshwide
// spec:
//   traceCollector:
//     zipkinCollector:
//       address: zipkin:9411
// ```
//
// WARNING: This is not currently supported by Istio.
type TracingRule struct {
	// Provides a way to tailor the configuration for a specific service.
	// If not specifed, this applies to all services in the namespace of the
	// configured resource. When the namespace is `istio-system`, the settings
	// will be applied to the entire mesh.
	//
	// Example: ratings
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	// Provides configuration for a tracing backend
	TraceCollector *TraceCollector `protobuf:"bytes,3,opt,name=trace_collector,json=traceCollector" json:"trace_collector,omitempty"`
	// $hide_from_docs
	//
	// Controls the rate of trace generation.
	//
	// NOTE: This will only be effective when support for Envoy V2 APIs is added
	// to Istio Pilot.
	TraceSampling *TraceSampling `protobuf:"bytes,4,opt,name=trace_sampling,json=traceSampling" json:"trace_sampling,omitempty"`
}

func (m *TracingRule) Reset()                    { *m = TracingRule{} }
func (m *TracingRule) String() string            { return proto.CompactTextString(m) }
func (*TracingRule) ProtoMessage()               {}
func (*TracingRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TracingRule) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *TracingRule) GetTraceCollector() *TraceCollector {
	if m != nil {
		return m.TraceCollector
	}
	return nil
}

func (m *TracingRule) GetTraceSampling() *TraceSampling {
	if m != nil {
		return m.TraceSampling
	}
	return nil
}

// TraceCollector holds the configuration for a specific tracing backend.
type TraceCollector struct {
	// Types that are valid to be assigned to Collector:
	//	*TraceCollector_ZipkinCollector
	//	*TraceCollector_LightstepCollector
	Collector isTraceCollector_Collector `protobuf_oneof:"collector"`
}

func (m *TraceCollector) Reset()                    { *m = TraceCollector{} }
func (m *TraceCollector) String() string            { return proto.CompactTextString(m) }
func (*TraceCollector) ProtoMessage()               {}
func (*TraceCollector) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isTraceCollector_Collector interface{ isTraceCollector_Collector() }

type TraceCollector_ZipkinCollector struct {
	ZipkinCollector *ZipkinCollector `protobuf:"bytes,1,opt,name=zipkin_collector,json=zipkinCollector,oneof"`
}
type TraceCollector_LightstepCollector struct {
	LightstepCollector *LightStepCollector `protobuf:"bytes,2,opt,name=lightstep_collector,json=lightstepCollector,oneof"`
}

func (*TraceCollector_ZipkinCollector) isTraceCollector_Collector()    {}
func (*TraceCollector_LightstepCollector) isTraceCollector_Collector() {}

func (m *TraceCollector) GetCollector() isTraceCollector_Collector {
	if m != nil {
		return m.Collector
	}
	return nil
}

func (m *TraceCollector) GetZipkinCollector() *ZipkinCollector {
	if x, ok := m.GetCollector().(*TraceCollector_ZipkinCollector); ok {
		return x.ZipkinCollector
	}
	return nil
}

func (m *TraceCollector) GetLightstepCollector() *LightStepCollector {
	if x, ok := m.GetCollector().(*TraceCollector_LightstepCollector); ok {
		return x.LightstepCollector
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TraceCollector) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TraceCollector_OneofMarshaler, _TraceCollector_OneofUnmarshaler, _TraceCollector_OneofSizer, []interface{}{
		(*TraceCollector_ZipkinCollector)(nil),
		(*TraceCollector_LightstepCollector)(nil),
	}
}

func _TraceCollector_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TraceCollector)
	// collector
	switch x := m.Collector.(type) {
	case *TraceCollector_ZipkinCollector:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ZipkinCollector); err != nil {
			return err
		}
	case *TraceCollector_LightstepCollector:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LightstepCollector); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TraceCollector.Collector has unexpected type %T", x)
	}
	return nil
}

func _TraceCollector_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TraceCollector)
	switch tag {
	case 1: // collector.zipkin_collector
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ZipkinCollector)
		err := b.DecodeMessage(msg)
		m.Collector = &TraceCollector_ZipkinCollector{msg}
		return true, err
	case 2: // collector.lightstep_collector
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LightStepCollector)
		err := b.DecodeMessage(msg)
		m.Collector = &TraceCollector_LightstepCollector{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TraceCollector_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TraceCollector)
	// collector
	switch x := m.Collector.(type) {
	case *TraceCollector_ZipkinCollector:
		s := proto.Size(x.ZipkinCollector)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TraceCollector_LightstepCollector:
		s := proto.Size(x.LightstepCollector)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Configuration for a Zipkin collector.
type ZipkinCollector struct {
	// Required. Address of the Zipkin service (e.g. _zipkin:9411_).
	Host string `protobuf:"bytes,1,opt,name=host" json:"host,omitempty"`
	// Optional. API endpoint of the Zipkin service.
	//
	// Default: "/api/v1/spans"
	ApiEndpoint string `protobuf:"bytes,2,opt,name=api_endpoint,json=apiEndpoint" json:"api_endpoint,omitempty"`
}

func (m *ZipkinCollector) Reset()                    { *m = ZipkinCollector{} }
func (m *ZipkinCollector) String() string            { return proto.CompactTextString(m) }
func (*ZipkinCollector) ProtoMessage()               {}
func (*ZipkinCollector) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ZipkinCollector) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *ZipkinCollector) GetApiEndpoint() string {
	if m != nil {
		return m.ApiEndpoint
	}
	return ""
}

// $hide_from_docs
//
// Configuration for a LightStep collector.
type LightStepCollector struct {
	// Required. Address of the LightStep service (e.g. _collector-grpc.lightstep.com:443_).
	Host string `protobuf:"bytes,1,opt,name=host" json:"host,omitempty"`
	// Required. File containing the access token for the LightStep API.
	//
	// Default: "/app/access_token"
	AccessTokenFile string `protobuf:"bytes,2,opt,name=access_token_file,json=accessTokenFile" json:"access_token_file,omitempty"`
}

func (m *LightStepCollector) Reset()                    { *m = LightStepCollector{} }
func (m *LightStepCollector) String() string            { return proto.CompactTextString(m) }
func (*LightStepCollector) ProtoMessage()               {}
func (*LightStepCollector) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LightStepCollector) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *LightStepCollector) GetAccessTokenFile() string {
	if m != nil {
		return m.AccessTokenFile
	}
	return ""
}

// Identifies a percentage.
type Percent struct {
	// Valid range: `[0.0, 100.0]`
	Value float64 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
}

func (m *Percent) Reset()                    { *m = Percent{} }
func (m *Percent) String() string            { return proto.CompactTextString(m) }
func (*Percent) ProtoMessage()               {}
func (*Percent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Percent) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// $hide_from_docs
//
// Controls the rate at which Istio will generate trace spans for requests
// within a mesh.
type TraceSampling struct {
	// Required. Trace generation mode.
	//
	// Default: NONE
	Mode TraceSampling_Mode `protobuf:"varint,1,opt,name=mode,enum=istio.tracing.v1alpha1.TraceSampling_Mode" json:"mode,omitempty"`
	// Target percentage of requests that will be force traced if requested
	// by a client.
	//
	// Optional. Only valid when mode is set to CUSTOM.
	//
	// Default: 0.0%
	ClientSampling *Percent `protobuf:"bytes,2,opt,name=client_sampling,json=clientSampling" json:"client_sampling,omitempty"`
	// Target percentage of requests managed that will be randomly selected for
	// trace generation, if not requested by the client.
	//
	// Optional. Only valid when mode is set to CUSTOM.
	//
	// Default: 0.0%
	RandomSampling *Percent `protobuf:"bytes,3,opt,name=random_sampling,json=randomSampling" json:"random_sampling,omitempty"`
}

func (m *TraceSampling) Reset()                    { *m = TraceSampling{} }
func (m *TraceSampling) String() string            { return proto.CompactTextString(m) }
func (*TraceSampling) ProtoMessage()               {}
func (*TraceSampling) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TraceSampling) GetMode() TraceSampling_Mode {
	if m != nil {
		return m.Mode
	}
	return TraceSampling_NONE
}

func (m *TraceSampling) GetClientSampling() *Percent {
	if m != nil {
		return m.ClientSampling
	}
	return nil
}

func (m *TraceSampling) GetRandomSampling() *Percent {
	if m != nil {
		return m.RandomSampling
	}
	return nil
}

func init() {
	proto.RegisterType((*TracingRule)(nil), "istio.tracing.v1alpha1.TracingRule")
	proto.RegisterType((*TraceCollector)(nil), "istio.tracing.v1alpha1.TraceCollector")
	proto.RegisterType((*ZipkinCollector)(nil), "istio.tracing.v1alpha1.ZipkinCollector")
	proto.RegisterType((*LightStepCollector)(nil), "istio.tracing.v1alpha1.LightStepCollector")
	proto.RegisterType((*Percent)(nil), "istio.tracing.v1alpha1.Percent")
	proto.RegisterType((*TraceSampling)(nil), "istio.tracing.v1alpha1.TraceSampling")
	proto.RegisterEnum("istio.tracing.v1alpha1.TraceSampling_Mode", TraceSampling_Mode_name, TraceSampling_Mode_value)
}

func init() { proto.RegisterFile("tracing/v1alpha1/tracing.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x97, 0xb6, 0xac, 0xf4, 0x84, 0x35, 0xc1, 0x20, 0xd4, 0x1b, 0xe8, 0x88, 0x04, 0x4c,
	0xbb, 0x48, 0xb5, 0x71, 0x8f, 0xb4, 0x4d, 0x43, 0x9d, 0x94, 0xb6, 0x28, 0x0d, 0x37, 0x93, 0x50,
	0x64, 0xd2, 0x43, 0x6b, 0xcd, 0xb1, 0xad, 0xc4, 0xeb, 0xc5, 0x5e, 0x84, 0xc7, 0x42, 0xe2, 0x89,
	0x50, 0xe2, 0x34, 0x6b, 0xc6, 0x8a, 0xb8, 0x8b, 0xff, 0x9c, 0xf3, 0xfd, 0xbf, 0x7c, 0x8e, 0xe1,
	0x8d, 0xce, 0x68, 0xc2, 0xc4, 0x72, 0xb4, 0x3e, 0xa1, 0x5c, 0xad, 0xe8, 0xc9, 0xa8, 0x12, 0x7c,
	0x95, 0x49, 0x2d, 0xc9, 0x2b, 0x96, 0x6b, 0x26, 0xfd, 0x8d, 0xb8, 0xa9, 0xf2, 0x7e, 0x59, 0x60,
	0x47, 0x46, 0x0c, 0x6f, 0x39, 0x92, 0xb7, 0xf0, 0x2c, 0xc7, 0x6c, 0xcd, 0x12, 0x8c, 0x05, 0x4d,
	0x71, 0x60, 0x1d, 0x5a, 0x47, 0xbd, 0xd0, 0xae, 0xb4, 0x29, 0x4d, 0x91, 0xcc, 0xc0, 0x29, 0x30,
	0x18, 0x27, 0x92, 0x73, 0x4c, 0xb4, 0xcc, 0x06, 0xed, 0x43, 0xeb, 0xc8, 0x3e, 0x7d, 0xef, 0x3f,
	0x6e, 0xe2, 0x17, 0x06, 0x78, 0xb1, 0xa9, 0x0e, 0xfb, 0xba, 0x71, 0x26, 0x01, 0x18, 0x25, 0xce,
	0x69, 0xaa, 0x38, 0x13, 0xcb, 0x41, 0xa7, 0xe4, 0xbd, 0xfb, 0x27, 0x6f, 0x5e, 0x15, 0x87, 0x07,
	0x7a, 0xfb, 0xe8, 0xfd, 0xb6, 0xa0, 0xdf, 0x34, 0x24, 0x11, 0xb8, 0x77, 0x4c, 0xdd, 0x30, 0xb1,
	0x15, 0xd9, 0x2a, 0x2d, 0x3e, 0xec, 0xb2, 0xb8, 0x2e, 0xeb, 0x6b, 0xc4, 0x78, 0x2f, 0x74, 0xee,
	0x9a, 0x12, 0xf9, 0x06, 0x2f, 0x38, 0x5b, 0xae, 0x74, 0xae, 0x51, 0x6d, 0x81, 0x5b, 0x25, 0xf8,
	0x78, 0x17, 0x38, 0x28, 0x5a, 0xe6, 0x1a, 0xd5, 0x36, 0x9b, 0xd4, 0xa0, 0x5a, 0x3d, 0xb7, 0xa1,
	0x57, 0x43, 0xbd, 0x31, 0x38, 0x0f, 0x12, 0x11, 0x02, 0x9d, 0x95, 0xcc, 0x75, 0x35, 0xa1, 0xf2,
	0xbb, 0x98, 0x1e, 0x55, 0x2c, 0x46, 0xb1, 0x50, 0x92, 0x09, 0x5d, 0x66, 0xe9, 0x85, 0x36, 0x55,
	0xec, 0xb2, 0x92, 0xbc, 0x08, 0xc8, 0xdf, 0x11, 0x1e, 0x85, 0x1d, 0xc3, 0x73, 0x9a, 0x24, 0x98,
	0xe7, 0xb1, 0x96, 0x37, 0x28, 0xe2, 0x1f, 0x8c, 0x63, 0x45, 0x74, 0xcc, 0x8f, 0xa8, 0xd0, 0x3f,
	0x33, 0x8e, 0xde, 0x10, 0xba, 0x5f, 0x30, 0x4b, 0x50, 0x68, 0xf2, 0x12, 0x9e, 0xac, 0x29, 0xbf,
	0x35, 0xab, 0x63, 0x85, 0xe6, 0xe0, 0xfd, 0x6c, 0xc1, 0x41, 0x63, 0x6c, 0xe4, 0x13, 0x74, 0x52,
	0xb9, 0x30, 0x65, 0xfd, 0xdd, 0xf7, 0xd5, 0x68, 0xf2, 0x27, 0x72, 0x81, 0x61, 0xd9, 0x47, 0xc6,
	0xe0, 0x24, 0x9c, 0xa1, 0xd0, 0xf7, 0x6b, 0x63, 0xae, 0x7e, 0xb8, 0x0b, 0x55, 0x25, 0x0c, 0xfb,
	0xa6, 0xaf, 0x4e, 0x32, 0x06, 0x27, 0xa3, 0x62, 0x21, 0xd3, 0x7b, 0x52, 0xfb, 0x3f, 0x49, 0xa6,
	0xaf, 0xde, 0xbd, 0x53, 0xe8, 0x14, 0x09, 0xc9, 0x53, 0xe8, 0x4c, 0x67, 0xd3, 0x4b, 0x77, 0x8f,
	0xd8, 0xd0, 0x9d, 0x5c, 0x4d, 0xaf, 0x26, 0x67, 0x81, 0x6b, 0x11, 0x80, 0xfd, 0x8b, 0xaf, 0xf3,
	0x68, 0x36, 0x71, 0x5b, 0xa4, 0x0b, 0xed, 0xb3, 0x20, 0x70, 0xdb, 0xe7, 0xc3, 0xeb, 0xd7, 0xc6,
	0x85, 0xc9, 0x11, 0x55, 0x6c, 0xf4, 0xf0, 0x21, 0x7f, 0xdf, 0x2f, 0x5f, 0xf0, 0xc7, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x8d, 0xd6, 0x7a, 0x4b, 0xe3, 0x03, 0x00, 0x00,
}
