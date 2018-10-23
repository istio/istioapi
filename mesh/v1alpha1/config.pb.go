// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mesh/v1alpha1/config.proto

/*
	Package v1alpha1 is a generated protocol buffer package.

	It is generated from these files:
		mesh/v1alpha1/config.proto
		mesh/v1alpha1/network.proto
		mesh/v1alpha1/proxy.proto

	It has these top-level messages:
		MeshConfig
		Network
		MeshNetworks
		Tracing
		ProxyConfig
*/
package v1alpha1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/types"

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

type MeshConfig_IngressControllerMode int32

const (
	// Disables Istio ingress controller.
	MeshConfig_OFF MeshConfig_IngressControllerMode = 0
	// Istio ingress controller will act on ingress resources that do not
	// contain any annotation or whose annotations match the value
	// specified in the ingress_class parameter described earlier. Use this
	// mode if Istio ingress controller will be the default ingress
	// controller for the entire kubernetes cluster.
	MeshConfig_DEFAULT MeshConfig_IngressControllerMode = 1
	// Istio ingress controller will only act on ingress resources whose
	// annotations match the value specified in the ingress_class parameter
	// described earlier. Use this mode if Istio ingress controller will be
	// a secondary ingress controller (e.g., in addition to a
	// cloud-provided ingress controller).
	MeshConfig_STRICT MeshConfig_IngressControllerMode = 2
)

var MeshConfig_IngressControllerMode_name = map[int32]string{
	0: "OFF",
	1: "DEFAULT",
	2: "STRICT",
}
var MeshConfig_IngressControllerMode_value = map[string]int32{
	"OFF":     0,
	"DEFAULT": 1,
	"STRICT":  2,
}

func (x MeshConfig_IngressControllerMode) String() string {
	return proto.EnumName(MeshConfig_IngressControllerMode_name, int32(x))
}
func (MeshConfig_IngressControllerMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorConfig, []int{0, 0}
}

// $hide_from_docs
type MeshConfig_AuthPolicy int32

const (
	MeshConfig_NONE       MeshConfig_AuthPolicy = 0
	MeshConfig_MUTUAL_TLS MeshConfig_AuthPolicy = 1
)

var MeshConfig_AuthPolicy_name = map[int32]string{
	0: "NONE",
	1: "MUTUAL_TLS",
}
var MeshConfig_AuthPolicy_value = map[string]int32{
	"NONE":       0,
	"MUTUAL_TLS": 1,
}

func (x MeshConfig_AuthPolicy) String() string {
	return proto.EnumName(MeshConfig_AuthPolicy_name, int32(x))
}
func (MeshConfig_AuthPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorConfig, []int{0, 1}
}

// MeshConfig defines mesh-wide variables shared by all Envoy instances in the
// Istio service mesh.
//
// NOTE: This configuration type should be used for the low-level global
// configuration, such as component addresses and port numbers. It should not
// be used for the features of the mesh that can be scoped by service or by
// namespace. Some of the fields in the mesh config are going to be deprecated
// and replaced with several individual configuration types (for example,
// tracing configuration).
type MeshConfig struct {
	// Address of the server that will be used by the proxies for policy
	// check calls. By using different names for mixerCheckServer and
	// mixerReportServer, it is possible to have one set of mixer servers handle
	// policy check calls while another set of mixer servers handle telemetry
	// calls.
	//
	// NOTE: Omitting mixerCheckServer while specifying mixerReportServer is
	// equivalent to setting disablePolicyChecks to true.
	MixerCheckServer string `protobuf:"bytes,1,opt,name=mixer_check_server,json=mixerCheckServer,proto3" json:"mixer_check_server,omitempty"`
	// Address of the server that will be used by the proxies for policy report
	// calls.
	MixerReportServer string `protobuf:"bytes,2,opt,name=mixer_report_server,json=mixerReportServer,proto3" json:"mixer_report_server,omitempty"`
	// Disable policy checks by the mixer service. Default
	// is false, i.e. mixer policy check is enabled by default.
	DisablePolicyChecks bool `protobuf:"varint,3,opt,name=disable_policy_checks,json=disablePolicyChecks,proto3" json:"disable_policy_checks,omitempty"`
	// Allow all traffic in cases when the mixer policy service cannot be reached.
	// Default is false which means the traffic is denied when the client is unable
	// to connect to Mixer.
	PolicyCheckFailOpen bool `protobuf:"varint,25,opt,name=policy_check_fail_open,json=policyCheckFailOpen,proto3" json:"policy_check_fail_open,omitempty"`
	// Port on which Envoy should listen for incoming connections from
	// other services.
	ProxyListenPort int32 `protobuf:"varint,4,opt,name=proxy_listen_port,json=proxyListenPort,proto3" json:"proxy_listen_port,omitempty"`
	// Port on which Envoy should listen for HTTP PROXY requests if set.
	ProxyHttpPort int32 `protobuf:"varint,5,opt,name=proxy_http_port,json=proxyHttpPort,proto3" json:"proxy_http_port,omitempty"`
	// Connection timeout used by Envoy. (MUST BE >=1ms)
	ConnectTimeout *google_protobuf.Duration `protobuf:"bytes,6,opt,name=connect_timeout,json=connectTimeout" json:"connect_timeout,omitempty"`
	// Class of ingress resources to be processed by Istio ingress
	// controller.  This corresponds to the value of
	// "kubernetes.io/ingress.class" annotation.
	IngressClass string `protobuf:"bytes,7,opt,name=ingress_class,json=ingressClass,proto3" json:"ingress_class,omitempty"`
	// Name of the kubernetes service used for the istio ingress controller.
	IngressService string `protobuf:"bytes,8,opt,name=ingress_service,json=ingressService,proto3" json:"ingress_service,omitempty"`
	// Defines whether to use Istio ingress controller for annotated or all ingress resources.
	IngressControllerMode MeshConfig_IngressControllerMode `protobuf:"varint,9,opt,name=ingress_controller_mode,json=ingressControllerMode,proto3,enum=istio.mesh.v1alpha1.MeshConfig_IngressControllerMode" json:"ingress_controller_mode,omitempty"`
	// $hide_from_docs
	AuthPolicy MeshConfig_AuthPolicy `protobuf:"varint,10,opt,name=auth_policy,json=authPolicy,proto3,enum=istio.mesh.v1alpha1.MeshConfig_AuthPolicy" json:"auth_policy,omitempty"`
	// Flag to control generation of trace spans and request IDs.
	// Requires a trace span collector defined in the proxy configuration.
	EnableTracing bool `protobuf:"varint,12,opt,name=enable_tracing,json=enableTracing,proto3" json:"enable_tracing,omitempty"`
	// File address for the proxy access log (e.g. /dev/stdout).
	// Empty value disables access logging.
	AccessLogFile string `protobuf:"bytes,13,opt,name=access_log_file,json=accessLogFile,proto3" json:"access_log_file,omitempty"`
	// Format for the proxy access log (text or json).
	// Default value is text.
	AccessLogFormat string `protobuf:"bytes,24,opt,name=access_log_format,json=accessLogFormat,proto3" json:"access_log_format,omitempty"`
	// Default proxy config used by the proxy injection mechanism operating in the mesh
	// (e.g. Kubernetes admission controller)
	// In case of Kubernetes, the proxy config is applied once during the injection process,
	// and remain constant for the duration of the pod. The rest of the mesh config can be changed
	// at runtime and config gets distributed dynamically.
	DefaultConfig *ProxyConfig `protobuf:"bytes,14,opt,name=default_config,json=defaultConfig" json:"default_config,omitempty"`
	// Enables clide side policy checks.
	EnableClientSidePolicyCheck bool `protobuf:"varint,19,opt,name=enable_client_side_policy_check,json=enableClientSidePolicyCheck,proto3" json:"enable_client_side_policy_check,omitempty"`
	// Unix Domain Socket through which envoy communicates with NodeAgent SDS to get key/cert for mTLS.
	// Use secret-mount files instead of SDS if set to empty.
	SdsUdsPath string `protobuf:"bytes,20,opt,name=sds_uds_path,json=sdsUdsPath,proto3" json:"sds_uds_path,omitempty"`
	// Address of the galley service exposing the Mesh Control Protocol (MCP).
	GalleyAddress string `protobuf:"bytes,22,opt,name=galley_address,json=galleyAddress,proto3" json:"galley_address,omitempty"`
	// $hide_from_docs
	// This flag is used by secret discovery service(SDS).
	// If set to true(prerequisite: https://kubernetes.io/docs/concepts/storage/volumes/#projected), Istio will inject volumes mount
	// for k8s service account JWT, so that K8s API server mounts k8s service account JWT to envoy container, which
	// will be used to generate key/cert eventually. This isn't supported for non-k8s case.
	EnableSdsTokenMount bool `protobuf:"varint,23,opt,name=enable_sds_token_mount,json=enableSdsTokenMount,proto3" json:"enable_sds_token_mount,omitempty"`
}

func (m *MeshConfig) Reset()                    { *m = MeshConfig{} }
func (m *MeshConfig) String() string            { return proto.CompactTextString(m) }
func (*MeshConfig) ProtoMessage()               {}
func (*MeshConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

func (m *MeshConfig) GetMixerCheckServer() string {
	if m != nil {
		return m.MixerCheckServer
	}
	return ""
}

func (m *MeshConfig) GetMixerReportServer() string {
	if m != nil {
		return m.MixerReportServer
	}
	return ""
}

func (m *MeshConfig) GetDisablePolicyChecks() bool {
	if m != nil {
		return m.DisablePolicyChecks
	}
	return false
}

func (m *MeshConfig) GetPolicyCheckFailOpen() bool {
	if m != nil {
		return m.PolicyCheckFailOpen
	}
	return false
}

func (m *MeshConfig) GetProxyListenPort() int32 {
	if m != nil {
		return m.ProxyListenPort
	}
	return 0
}

func (m *MeshConfig) GetProxyHttpPort() int32 {
	if m != nil {
		return m.ProxyHttpPort
	}
	return 0
}

func (m *MeshConfig) GetConnectTimeout() *google_protobuf.Duration {
	if m != nil {
		return m.ConnectTimeout
	}
	return nil
}

func (m *MeshConfig) GetIngressClass() string {
	if m != nil {
		return m.IngressClass
	}
	return ""
}

func (m *MeshConfig) GetIngressService() string {
	if m != nil {
		return m.IngressService
	}
	return ""
}

func (m *MeshConfig) GetIngressControllerMode() MeshConfig_IngressControllerMode {
	if m != nil {
		return m.IngressControllerMode
	}
	return MeshConfig_OFF
}

func (m *MeshConfig) GetAuthPolicy() MeshConfig_AuthPolicy {
	if m != nil {
		return m.AuthPolicy
	}
	return MeshConfig_NONE
}

func (m *MeshConfig) GetEnableTracing() bool {
	if m != nil {
		return m.EnableTracing
	}
	return false
}

func (m *MeshConfig) GetAccessLogFile() string {
	if m != nil {
		return m.AccessLogFile
	}
	return ""
}

func (m *MeshConfig) GetAccessLogFormat() string {
	if m != nil {
		return m.AccessLogFormat
	}
	return ""
}

func (m *MeshConfig) GetDefaultConfig() *ProxyConfig {
	if m != nil {
		return m.DefaultConfig
	}
	return nil
}

func (m *MeshConfig) GetEnableClientSidePolicyCheck() bool {
	if m != nil {
		return m.EnableClientSidePolicyCheck
	}
	return false
}

func (m *MeshConfig) GetSdsUdsPath() string {
	if m != nil {
		return m.SdsUdsPath
	}
	return ""
}

func (m *MeshConfig) GetGalleyAddress() string {
	if m != nil {
		return m.GalleyAddress
	}
	return ""
}

func (m *MeshConfig) GetEnableSdsTokenMount() bool {
	if m != nil {
		return m.EnableSdsTokenMount
	}
	return false
}

func init() {
	proto.RegisterType((*MeshConfig)(nil), "istio.mesh.v1alpha1.MeshConfig")
	proto.RegisterEnum("istio.mesh.v1alpha1.MeshConfig_IngressControllerMode", MeshConfig_IngressControllerMode_name, MeshConfig_IngressControllerMode_value)
	proto.RegisterEnum("istio.mesh.v1alpha1.MeshConfig_AuthPolicy", MeshConfig_AuthPolicy_name, MeshConfig_AuthPolicy_value)
}
func (m *MeshConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MeshConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.MixerCheckServer) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.MixerCheckServer)))
		i += copy(dAtA[i:], m.MixerCheckServer)
	}
	if len(m.MixerReportServer) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.MixerReportServer)))
		i += copy(dAtA[i:], m.MixerReportServer)
	}
	if m.DisablePolicyChecks {
		dAtA[i] = 0x18
		i++
		if m.DisablePolicyChecks {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.ProxyListenPort != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.ProxyListenPort))
	}
	if m.ProxyHttpPort != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.ProxyHttpPort))
	}
	if m.ConnectTimeout != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.ConnectTimeout.Size()))
		n1, err := m.ConnectTimeout.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.IngressClass) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.IngressClass)))
		i += copy(dAtA[i:], m.IngressClass)
	}
	if len(m.IngressService) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.IngressService)))
		i += copy(dAtA[i:], m.IngressService)
	}
	if m.IngressControllerMode != 0 {
		dAtA[i] = 0x48
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.IngressControllerMode))
	}
	if m.AuthPolicy != 0 {
		dAtA[i] = 0x50
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.AuthPolicy))
	}
	if m.EnableTracing {
		dAtA[i] = 0x60
		i++
		if m.EnableTracing {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.AccessLogFile) > 0 {
		dAtA[i] = 0x6a
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.AccessLogFile)))
		i += copy(dAtA[i:], m.AccessLogFile)
	}
	if m.DefaultConfig != nil {
		dAtA[i] = 0x72
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.DefaultConfig.Size()))
		n2, err := m.DefaultConfig.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.EnableClientSidePolicyCheck {
		dAtA[i] = 0x98
		i++
		dAtA[i] = 0x1
		i++
		if m.EnableClientSidePolicyCheck {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.SdsUdsPath) > 0 {
		dAtA[i] = 0xa2
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.SdsUdsPath)))
		i += copy(dAtA[i:], m.SdsUdsPath)
	}
	if len(m.GalleyAddress) > 0 {
		dAtA[i] = 0xb2
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.GalleyAddress)))
		i += copy(dAtA[i:], m.GalleyAddress)
	}
	if m.EnableSdsTokenMount {
		dAtA[i] = 0xb8
		i++
		dAtA[i] = 0x1
		i++
		if m.EnableSdsTokenMount {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.AccessLogFormat) > 0 {
		dAtA[i] = 0xc2
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.AccessLogFormat)))
		i += copy(dAtA[i:], m.AccessLogFormat)
	}
	if m.PolicyCheckFailOpen {
		dAtA[i] = 0xc8
		i++
		dAtA[i] = 0x1
		i++
		if m.PolicyCheckFailOpen {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *MeshConfig) Size() (n int) {
	var l int
	_ = l
	l = len(m.MixerCheckServer)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.MixerReportServer)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.DisablePolicyChecks {
		n += 2
	}
	if m.ProxyListenPort != 0 {
		n += 1 + sovConfig(uint64(m.ProxyListenPort))
	}
	if m.ProxyHttpPort != 0 {
		n += 1 + sovConfig(uint64(m.ProxyHttpPort))
	}
	if m.ConnectTimeout != nil {
		l = m.ConnectTimeout.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.IngressClass)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.IngressService)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.IngressControllerMode != 0 {
		n += 1 + sovConfig(uint64(m.IngressControllerMode))
	}
	if m.AuthPolicy != 0 {
		n += 1 + sovConfig(uint64(m.AuthPolicy))
	}
	if m.EnableTracing {
		n += 2
	}
	l = len(m.AccessLogFile)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.DefaultConfig != nil {
		l = m.DefaultConfig.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.EnableClientSidePolicyCheck {
		n += 3
	}
	l = len(m.SdsUdsPath)
	if l > 0 {
		n += 2 + l + sovConfig(uint64(l))
	}
	l = len(m.GalleyAddress)
	if l > 0 {
		n += 2 + l + sovConfig(uint64(l))
	}
	if m.EnableSdsTokenMount {
		n += 3
	}
	l = len(m.AccessLogFormat)
	if l > 0 {
		n += 2 + l + sovConfig(uint64(l))
	}
	if m.PolicyCheckFailOpen {
		n += 3
	}
	return n
}

func sovConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MeshConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: MeshConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MeshConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MixerCheckServer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MixerCheckServer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MixerReportServer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MixerReportServer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisablePolicyChecks", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
			m.DisablePolicyChecks = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProxyListenPort", wireType)
			}
			m.ProxyListenPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProxyListenPort |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProxyHttpPort", wireType)
			}
			m.ProxyHttpPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProxyHttpPort |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectTimeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ConnectTimeout == nil {
				m.ConnectTimeout = &google_protobuf.Duration{}
			}
			if err := m.ConnectTimeout.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IngressClass", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IngressClass = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IngressService", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IngressService = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IngressControllerMode", wireType)
			}
			m.IngressControllerMode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IngressControllerMode |= (MeshConfig_IngressControllerMode(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthPolicy", wireType)
			}
			m.AuthPolicy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuthPolicy |= (MeshConfig_AuthPolicy(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableTracing", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
			m.EnableTracing = bool(v != 0)
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessLogFile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccessLogFile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DefaultConfig == nil {
				m.DefaultConfig = &ProxyConfig{}
			}
			if err := m.DefaultConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 19:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableClientSidePolicyCheck", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
			m.EnableClientSidePolicyCheck = bool(v != 0)
		case 20:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SdsUdsPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SdsUdsPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 22:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GalleyAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GalleyAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 23:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableSdsTokenMount", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
			m.EnableSdsTokenMount = bool(v != 0)
		case 24:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessLogFormat", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccessLogFormat = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 25:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PolicyCheckFailOpen", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
			m.PolicyCheckFailOpen = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
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
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfig
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
				next, err := skipConfig(dAtA[start:])
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
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("mesh/v1alpha1/config.proto", fileDescriptorConfig) }

var fileDescriptorConfig = []byte{
	// 750 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xdd, 0x6e, 0xdb, 0x36,
	0x18, 0xad, 0x52, 0x35, 0x51, 0xe8, 0x3f, 0x99, 0x9e, 0x53, 0xc6, 0x03, 0x3c, 0x23, 0xc3, 0x3a,
	0xa3, 0x18, 0x64, 0x34, 0xc5, 0x2e, 0x76, 0xe9, 0x38, 0xf5, 0x56, 0xc3, 0xae, 0x0d, 0x59, 0xbe,
	0xd9, 0x0d, 0xc1, 0x48, 0xb4, 0x44, 0x94, 0x16, 0x05, 0x91, 0x2a, 0x92, 0x37, 0xd8, 0xa3, 0xed,
	0x72, 0x8f, 0x30, 0xe4, 0x49, 0x06, 0x91, 0x72, 0x9c, 0x00, 0x06, 0x76, 0xa9, 0x73, 0xce, 0xf7,
	0x7b, 0x3e, 0x0a, 0xf4, 0x76, 0x54, 0x26, 0xa3, 0x6f, 0x1f, 0x08, 0xcf, 0x12, 0xf2, 0x61, 0x14,
	0x8a, 0x74, 0xcb, 0x62, 0x2f, 0xcb, 0x85, 0x12, 0xb0, 0xc3, 0xa4, 0x62, 0xc2, 0x2b, 0x15, 0xde,
	0x5e, 0xd1, 0xeb, 0xc7, 0x42, 0xc4, 0x9c, 0x8e, 0xb4, 0xe4, 0xae, 0xd8, 0x8e, 0xa2, 0x22, 0x27,
	0x8a, 0x89, 0xd4, 0x04, 0xf5, 0x2e, 0x5f, 0x26, 0xcc, 0x72, 0x71, 0xff, 0x60, 0xa8, 0xab, 0xbf,
	0xce, 0x01, 0x58, 0x50, 0x99, 0x4c, 0x74, 0x11, 0xf8, 0x0b, 0x80, 0x3b, 0x76, 0x4f, 0x73, 0x1c,
	0x26, 0x34, 0xfc, 0x8a, 0x25, 0xcd, 0xbf, 0xd1, 0x1c, 0x59, 0x03, 0x6b, 0x78, 0xee, 0xbb, 0x9a,
	0x99, 0x94, 0xc4, 0x5a, 0xe3, 0xd0, 0x03, 0x1d, 0xa3, 0xce, 0x69, 0x26, 0x72, 0xb5, 0x97, 0x9f,
	0x68, 0x79, 0x5b, 0x53, 0xbe, 0x66, 0x2a, 0xfd, 0x35, 0xe8, 0x46, 0x4c, 0x92, 0x3b, 0x4e, 0x71,
	0x26, 0x38, 0x0b, 0x1f, 0x4c, 0x19, 0x89, 0x5e, 0x0f, 0xac, 0xa1, 0xe3, 0x77, 0x2a, 0x72, 0xa5,
	0x39, 0x5d, 0x48, 0xc2, 0xf7, 0xa0, 0xad, 0xfb, 0xc5, 0x9c, 0x49, 0x45, 0x53, 0x5c, 0xa6, 0x43,
	0xf6, 0xc0, 0x1a, 0xbe, 0xf1, 0x5b, 0x9a, 0x98, 0x6b, 0x7c, 0x25, 0x72, 0x05, 0xdf, 0x01, 0x03,
	0xe1, 0x44, 0xa9, 0xcc, 0x28, 0xdf, 0x68, 0x65, 0x43, 0xc3, 0x7f, 0x28, 0x95, 0x69, 0xdd, 0x0d,
	0x68, 0x85, 0x22, 0x4d, 0x69, 0xa8, 0xb0, 0x62, 0x3b, 0x2a, 0x0a, 0x85, 0x4e, 0x07, 0xd6, 0xb0,
	0x76, 0x7d, 0xe9, 0x99, 0x4d, 0x7a, 0xfb, 0x4d, 0x7a, 0xb7, 0xd5, 0x26, 0xfd, 0x66, 0x15, 0x11,
	0x98, 0x00, 0xf8, 0x23, 0x68, 0xb0, 0x34, 0xce, 0xa9, 0x94, 0x38, 0xe4, 0x44, 0x4a, 0x74, 0xa6,
	0xa7, 0xae, 0x57, 0xe0, 0xa4, 0xc4, 0xe0, 0xcf, 0xa0, 0xb5, 0x17, 0x95, 0xbb, 0x61, 0x21, 0x45,
	0x8e, 0x96, 0x35, 0x2b, 0x78, 0x6d, 0x50, 0xb8, 0x03, 0x6f, 0x9f, 0xb2, 0x89, 0x54, 0xe5, 0x82,
	0x73, 0x9a, 0xe3, 0x9d, 0x88, 0x28, 0x3a, 0x1f, 0x58, 0xc3, 0xe6, 0xf5, 0xaf, 0xde, 0x11, 0xe3,
	0xbd, 0x83, 0x73, 0xde, 0xe7, 0xaa, 0xee, 0x53, 0xf4, 0x42, 0x44, 0xd4, 0xef, 0xb2, 0x63, 0x30,
	0x5c, 0x82, 0x1a, 0x29, 0x54, 0x52, 0xb9, 0x80, 0x80, 0x2e, 0xf1, 0xfe, 0xff, 0x4a, 0x8c, 0x0b,
	0x95, 0x18, 0x6f, 0x6e, 0x4e, 0x90, 0xe5, 0x03, 0xf2, 0xf4, 0x0d, 0x7f, 0x02, 0x4d, 0x9a, 0x6a,
	0x63, 0x55, 0x4e, 0x42, 0x96, 0xc6, 0xa8, 0xae, 0x2d, 0x6d, 0x18, 0x34, 0x30, 0x60, 0x69, 0x10,
	0x09, 0xc3, 0x72, 0x4a, 0x2e, 0x62, 0xbc, 0x65, 0x9c, 0xa2, 0x86, 0xde, 0x47, 0xc3, 0xc0, 0x73,
	0x11, 0x4f, 0x19, 0xa7, 0xf0, 0x77, 0xd0, 0x8c, 0xe8, 0x96, 0x14, 0x5c, 0x61, 0x73, 0xfd, 0xa8,
	0xa9, 0xfd, 0x19, 0x1c, 0x6d, 0x71, 0x55, 0x9a, 0x6b, 0x7a, 0xf4, 0x1b, 0x55, 0x5c, 0x75, 0xcf,
	0xb7, 0xe0, 0x87, 0xaa, 0xaf, 0x90, 0x33, 0x9a, 0x2a, 0x2c, 0x59, 0xf4, 0xf2, 0xf8, 0x50, 0x47,
	0x37, 0xfa, 0xbd, 0x91, 0x4d, 0xb4, 0x6a, 0xcd, 0xa2, 0xe7, 0x47, 0x08, 0x07, 0xa0, 0x2e, 0x23,
	0x89, 0x8b, 0x48, 0xe2, 0x8c, 0xa8, 0x04, 0x7d, 0xa7, 0x7b, 0x06, 0x32, 0x92, 0x9b, 0x48, 0xae,
	0x88, 0x4a, 0xca, 0xf9, 0x63, 0xc2, 0x39, 0x7d, 0xc0, 0x24, 0x8a, 0xca, 0x85, 0xa3, 0x0b, 0x33,
	0x97, 0x41, 0xc7, 0x06, 0x84, 0x1f, 0xc1, 0x45, 0xd5, 0x4e, 0x99, 0x4f, 0x89, 0xaf, 0x34, 0xc5,
	0x3b, 0x51, 0xa4, 0x0a, 0xbd, 0x35, 0x2f, 0xc0, 0xb0, 0xeb, 0x48, 0x06, 0x25, 0xb7, 0x28, 0xa9,
	0xf2, 0x05, 0x3c, 0x5f, 0x9a, 0xc8, 0x77, 0x44, 0x21, 0xa4, 0xd3, 0xb7, 0x0e, 0x6b, 0xd3, 0x70,
	0x59, 0xe0, 0xf9, 0x70, 0x78, 0x4b, 0x18, 0xc7, 0x22, 0xa3, 0x29, 0xba, 0x34, 0x05, 0xb2, 0xc3,
	0x58, 0x53, 0xc2, 0xf8, 0x32, 0xa3, 0xe9, 0xd5, 0x6f, 0xa0, 0x7b, 0xf4, 0x7a, 0xe0, 0x19, 0x78,
	0xbd, 0x9c, 0x4e, 0xdd, 0x57, 0xb0, 0x06, 0xce, 0x6e, 0x3f, 0x4d, 0xc7, 0x9b, 0x79, 0xe0, 0x5a,
	0x10, 0x80, 0xd3, 0x75, 0xe0, 0x7f, 0x9e, 0x04, 0xee, 0xc9, 0xd5, 0x3b, 0x00, 0x0e, 0x57, 0x01,
	0x1d, 0x60, 0x7f, 0x59, 0x7e, 0xf9, 0xe4, 0xbe, 0x82, 0x4d, 0x00, 0x16, 0x9b, 0x60, 0x33, 0x9e,
	0xe3, 0x60, 0xbe, 0x76, 0xad, 0x99, 0xed, 0xd4, 0xdc, 0xfa, 0xcc, 0x76, 0x5a, 0xae, 0x3b, 0xb3,
	0x1d, 0xd7, 0x6d, 0xcf, 0x6c, 0xa7, 0xed, 0xc2, 0x99, 0xed, 0x40, 0xb7, 0x33, 0xb3, 0x9d, 0xae,
	0x7b, 0x71, 0x33, 0xfc, 0xfb, 0xb1, 0x6f, 0xfd, 0xf3, 0xd8, 0xb7, 0xfe, 0x7d, 0xec, 0x5b, 0x7f,
	0xf6, 0x8c, 0xd3, 0x4c, 0x8c, 0x48, 0xc6, 0x46, 0x2f, 0x7e, 0x60, 0x77, 0xa7, 0xfa, 0x79, 0x7e,
	0xfc, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x06, 0x2a, 0xee, 0xaa, 0x29, 0x05, 0x00, 0x00,
}
