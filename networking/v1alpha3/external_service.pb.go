// Code generated by protoc-gen-gogo.
// source: networking/v1alpha3/external_service.proto
// DO NOT EDIT!

package v1alpha3

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Different ways of discovering the IP addresses associated with the
// service.
type ExternalService_Discovery int32

const (
	// If set to "none", the proxy will assume that incoming connections
	// have already been resolved (to a specific destination IP
	// address). Such connections are typically routed via the proxy using
	// mechanisms such as IP table REDIRECT/ eBPF. After performing any
	// routing related transformations, the proxy will forward the
	// connection to the IP address to which the connection was bound.
	ExternalService_NONE ExternalService_Discovery = 0
	// If set to "static", the proxy will use the IP addresses specified in
	// endpoints (See below) as the backing nodes associated with the
	// external service.
	ExternalService_STATIC ExternalService_Discovery = 1
	// If set to "dns", the proxy will attempt to resolve the DNS address
	// during request processing. If no endpoints are specified, the proxy
	// will resolve the DNS address specified in the hosts field, if
	// wildcards are not used. If endpoints are specified, the DNS
	// addresses specified in the endpoints will be resolved to determine
	// the destination IP address.
	ExternalService_DNS ExternalService_Discovery = 2
)

var ExternalService_Discovery_name = map[int32]string{
	0: "NONE",
	1: "STATIC",
	2: "DNS",
}
var ExternalService_Discovery_value = map[string]int32{
	"NONE":   0,
	"STATIC": 1,
	"DNS":    2,
}

func (x ExternalService_Discovery) String() string {
	return proto.EnumName(ExternalService_Discovery_name, int32(x))
}
func (ExternalService_Discovery) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorExternalService, []int{0, 0}
}

// External service describes the endpoints, ports and protocols of a
// white-listed set of mesh-external domains and IP blocks that services in
// the mesh are allowed to access.
//
// For example, the following external service configuration describes the
// set of services at https://example.com to be accessed internally over
// plaintext http (i.e. http://example.com:443), with the sidecar originating
// TLS.
//
//     apiVersion: networking.istio.io/v1alpha3
//     kind: ExternalService
//     metadata:
//       name: external-svc-example
//     spec:
//       hosts:
//       - example.com
//       ports:
//       - number: 443
//         name: example-http
//         protocol: http # not HTTPS.
//       discovery: DNS
//
// and a destination rule to initiate TLS connections to the external service.
//
//     apiVersion: networking.istio.io/v1alpha3
//     kind: DestinationRule
//     metadata:
//       name: tls-example
//     spec:
//       name: example.com
//       trafficPolicy:
//         tls:
//           mode: SIMPLE # initiates HTTPS when talking to example.com
//
// The following specification specifies a static set of backend nodes for
// a MongoDB cluster behind a set of virtual IPs, and sets up a destination
// rule to initiate mTLS connections upstream.
//
//     apiVersion: networking.istio.io/v1alpha3
//     kind: ExternalService
//     metadata:
//       name: external-svc-mongocluster
//     spec:
//       hosts:
//       - 192.192.192.192/24
//       ports:
//       - number: 27018
//         name: mongodb
//         protocol: mongo
//       discovery: STATIC
//       endpoints:
//       - address: 2.2.2.2
//       - address: 3.3.3.3
//
// and the associated destination rule
//
//     apiVersion: networking.istio.io/v1alpha3
//     kind: DestinationRule
//     metadata:
//       name: mtls-mongocluster
//     spec:
//       name: 192.192.192.192/24
//       trafficPolicy:
//         tls:
//           mode: MUTUAL
//           clientCertificate: /etc/certs/myclientcert.pem
//           privateKey: /etc/certs/client_private_key.pem
//           caCertificates: /etc/certs/rootcacerts.pem
//
// The following example demonstrates the use of wildcards in the hosts. If
// the connection has to be routed to the IP address requested by the
// application (i.e. application resolves DNS and attempts to connect to a
// specific IP), the discovery mode must be set to "none".
//
//     apiVersion: networking.istio.io/v1alpha3
//     kind: ExternalService
//     metadata:
//       name: external-svc-wildcard-example
//     spec:
//       hosts:
//       - *.bar.com
//       ports:
//       - number: 80
//         name: http
//         protocol: http
//       discovery: NONE
//
// For HTTP based services, it is possible to create a virtual service
// backed by multiple DNS addressible endpoints. In such a scenario, the
// application can use the HTTP_PROXY environment variable to transparently
// reroute API calls for the virtual service to a chosen backend. For
// example, the following configuration creates a non-existent service
// called foo.bar.com backed by three domains: us.foo.bar.com:8443,
// uk.foo.bar.com:9443, and in.foo.bar.com:7443
//
//     apiVersion: networking.istio.io/v1alpha3
//     kind: ExternalService
//     metadata:
//       name: external-svc-dns
//     spec:
//       hosts:
//       - foo.bar.com
//       ports:
//       - number: 443
//         name: https
//         protocol: http
//       discovery: DNS
//       endpoints:
//       - address: us.foo.bar.com
//         ports:
//         - https: 8443
//       - address: uk.foo.bar.com
//         ports:
//         - https: 9443
//       - address: in.foo.bar.com
//         ports:
//         - https: 7443
//
// and a destination rule to initiate TLS connections to the external service.
//
//     apiVersion: networking.istio.io/v1alpha3
//     kind: DestinationRule
//     metadata:
//       name: tls-foobar
//     spec:
//       name: foo.bar.com
//       trafficPolicy:
//         tls:
//           mode: SIMPLE # initiates HTTPS
//
// With HTTP_PROXY=http://localhost:443, calls from the application to
// http://foo.bar.com will be upgraded to HTTPS and load balanced across
// the three domains specified above. In other words, a call to
// http://foo.bar.com/baz would be translated to
// https://uk.foo.bar.com/baz.
//
// NOTE: In the scenario above, the value of the HTTP Authority/host header
// associated with the outbound HTTP requests will be based on the
// endpoint's DNS name, i.e. ":authority: uk.foo.bar.com". Refer to Envoy's
// auto_host_rewrite for further details. The automatic rewrite can be
// overridden using a host rewrite route rule.
//
type ExternalService struct {
	// REQUIRED. The hosts associated with the external service. Could be a
	// DNS name with wildcard prefix or a CIDR prefix. Note that the hosts
	// field applies to all protocols. DNS names in hosts will be ignored if
	// the application accesses the service over non-HTTP protocols such as
	// mongo/opaque TCP/even HTTPS. In such scenarios, the port on which the
	// external service is being accessed must not be shared by any other
	// service in the mesh. In other words, the sidecar will behave as a
	// simple TCP proxy, forwarding incoming traffic on a specified port to
	// the specified destination endpoint IP/host.
	Hosts []string `protobuf:"bytes,1,rep,name=hosts" json:"hosts,omitempty"`
	// REQUIRED. The ports associated with the external service.
	Ports []Port `protobuf:"bytes,2,rep,name=ports" json:"ports"`
	// Service discovery mode for the hosts. If not set, Istio will attempt
	// to infer the discovery mode based on the value of hosts and endpoints.
	Discovery ExternalService_Discovery `protobuf:"varint,3,opt,name=discovery,proto3,enum=istio.networking.v1alpha3.ExternalService_Discovery" json:"discovery,omitempty"`
	// One or more endpoints associated with the service. Endpoints must be
	// accessible over the set of outPorts defined at the service level.
	Endpoints []ExternalService_Endpoint `protobuf:"bytes,4,rep,name=endpoints" json:"endpoints"`
}

func (m *ExternalService) Reset()                    { *m = ExternalService{} }
func (m *ExternalService) String() string            { return proto.CompactTextString(m) }
func (*ExternalService) ProtoMessage()               {}
func (*ExternalService) Descriptor() ([]byte, []int) { return fileDescriptorExternalService, []int{0} }

func (m *ExternalService) GetHosts() []string {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func (m *ExternalService) GetPorts() []Port {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *ExternalService) GetDiscovery() ExternalService_Discovery {
	if m != nil {
		return m.Discovery
	}
	return ExternalService_NONE
}

func (m *ExternalService) GetEndpoints() []ExternalService_Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

// Endpoint defines a network address (IP or hostname) associated with
// the external service.
type ExternalService_Endpoint struct {
	// REQUIRED: Address associated with the network endpoint without the
	// port ( IP or fully qualified domain name without wildcards).
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Set of ports associated with the endpoint. The ports must be
	// associated with a port name that was declared as part of the
	// service.
	Ports map[string]uint32 `protobuf:"bytes,2,rep,name=ports" json:"ports,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// One or more labels associated with the endpoint.
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *ExternalService_Endpoint) Reset()         { *m = ExternalService_Endpoint{} }
func (m *ExternalService_Endpoint) String() string { return proto.CompactTextString(m) }
func (*ExternalService_Endpoint) ProtoMessage()    {}
func (*ExternalService_Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptorExternalService, []int{0, 0}
}

func (m *ExternalService_Endpoint) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ExternalService_Endpoint) GetPorts() map[string]uint32 {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *ExternalService_Endpoint) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*ExternalService)(nil), "istio.networking.v1alpha3.ExternalService")
	proto.RegisterType((*ExternalService_Endpoint)(nil), "istio.networking.v1alpha3.ExternalService.Endpoint")
	proto.RegisterEnum("istio.networking.v1alpha3.ExternalService_Discovery", ExternalService_Discovery_name, ExternalService_Discovery_value)
}
func (m *ExternalService) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExternalService) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Hosts) > 0 {
		for _, s := range m.Hosts {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Ports) > 0 {
		for _, msg := range m.Ports {
			dAtA[i] = 0x12
			i++
			i = encodeVarintExternalService(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Discovery != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintExternalService(dAtA, i, uint64(m.Discovery))
	}
	if len(m.Endpoints) > 0 {
		for _, msg := range m.Endpoints {
			dAtA[i] = 0x22
			i++
			i = encodeVarintExternalService(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ExternalService_Endpoint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExternalService_Endpoint) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintExternalService(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
	}
	if len(m.Ports) > 0 {
		for k, _ := range m.Ports {
			dAtA[i] = 0x12
			i++
			v := m.Ports[k]
			mapSize := 1 + len(k) + sovExternalService(uint64(len(k))) + 1 + sovExternalService(uint64(v))
			i = encodeVarintExternalService(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintExternalService(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x10
			i++
			i = encodeVarintExternalService(dAtA, i, uint64(v))
		}
	}
	if len(m.Labels) > 0 {
		for k, _ := range m.Labels {
			dAtA[i] = 0x1a
			i++
			v := m.Labels[k]
			mapSize := 1 + len(k) + sovExternalService(uint64(len(k))) + 1 + len(v) + sovExternalService(uint64(len(v)))
			i = encodeVarintExternalService(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintExternalService(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintExternalService(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	return i, nil
}

func encodeFixed64ExternalService(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32ExternalService(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintExternalService(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ExternalService) Size() (n int) {
	var l int
	_ = l
	if len(m.Hosts) > 0 {
		for _, s := range m.Hosts {
			l = len(s)
			n += 1 + l + sovExternalService(uint64(l))
		}
	}
	if len(m.Ports) > 0 {
		for _, e := range m.Ports {
			l = e.Size()
			n += 1 + l + sovExternalService(uint64(l))
		}
	}
	if m.Discovery != 0 {
		n += 1 + sovExternalService(uint64(m.Discovery))
	}
	if len(m.Endpoints) > 0 {
		for _, e := range m.Endpoints {
			l = e.Size()
			n += 1 + l + sovExternalService(uint64(l))
		}
	}
	return n
}

func (m *ExternalService_Endpoint) Size() (n int) {
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovExternalService(uint64(l))
	}
	if len(m.Ports) > 0 {
		for k, v := range m.Ports {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovExternalService(uint64(len(k))) + 1 + sovExternalService(uint64(v))
			n += mapEntrySize + 1 + sovExternalService(uint64(mapEntrySize))
		}
	}
	if len(m.Labels) > 0 {
		for k, v := range m.Labels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovExternalService(uint64(len(k))) + 1 + len(v) + sovExternalService(uint64(len(v)))
			n += mapEntrySize + 1 + sovExternalService(uint64(mapEntrySize))
		}
	}
	return n
}

func sovExternalService(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozExternalService(x uint64) (n int) {
	return sovExternalService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExternalService) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExternalService
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
			return fmt.Errorf("proto: ExternalService: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExternalService: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hosts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalService
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
				return ErrInvalidLengthExternalService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hosts = append(m.Hosts, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ports", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalService
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
				return ErrInvalidLengthExternalService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ports = append(m.Ports, Port{})
			if err := m.Ports[len(m.Ports)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Discovery", wireType)
			}
			m.Discovery = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Discovery |= (ExternalService_Discovery(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Endpoints", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalService
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
				return ErrInvalidLengthExternalService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Endpoints = append(m.Endpoints, ExternalService_Endpoint{})
			if err := m.Endpoints[len(m.Endpoints)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExternalService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExternalService
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
func (m *ExternalService_Endpoint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExternalService
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
			return fmt.Errorf("proto: Endpoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Endpoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalService
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
				return ErrInvalidLengthExternalService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ports", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalService
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
				return ErrInvalidLengthExternalService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalService
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
					return ErrIntOverflowExternalService
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
				return ErrInvalidLengthExternalService
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.Ports == nil {
				m.Ports = make(map[string]uint32)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowExternalService
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
				var mapvalue uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowExternalService
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					mapvalue |= (uint32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Ports[mapkey] = mapvalue
			} else {
				var mapvalue uint32
				m.Ports[mapkey] = mapvalue
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalService
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
				return ErrInvalidLengthExternalService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalService
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
					return ErrIntOverflowExternalService
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
				return ErrInvalidLengthExternalService
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.Labels == nil {
				m.Labels = make(map[string]string)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowExternalService
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
				var stringLenmapvalue uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowExternalService
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLenmapvalue |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLenmapvalue := int(stringLenmapvalue)
				if intStringLenmapvalue < 0 {
					return ErrInvalidLengthExternalService
				}
				postStringIndexmapvalue := iNdEx + intStringLenmapvalue
				if postStringIndexmapvalue > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := string(dAtA[iNdEx:postStringIndexmapvalue])
				iNdEx = postStringIndexmapvalue
				m.Labels[mapkey] = mapvalue
			} else {
				var mapvalue string
				m.Labels[mapkey] = mapvalue
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExternalService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExternalService
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
func skipExternalService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExternalService
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
					return 0, ErrIntOverflowExternalService
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
					return 0, ErrIntOverflowExternalService
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
				return 0, ErrInvalidLengthExternalService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowExternalService
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
				next, err := skipExternalService(dAtA[start:])
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
	ErrInvalidLengthExternalService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExternalService   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("networking/v1alpha3/external_service.proto", fileDescriptorExternalService)
}

var fileDescriptorExternalService = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcd, 0xae, 0xd2, 0x40,
	0x14, 0xc7, 0xef, 0xb4, 0xc0, 0xbd, 0x3d, 0xd7, 0x8f, 0x66, 0xc2, 0xa2, 0x76, 0x01, 0x95, 0x55,
	0xc3, 0xa2, 0x44, 0x70, 0x81, 0x9a, 0x68, 0x44, 0xba, 0x30, 0x31, 0x68, 0x0a, 0x09, 0x4b, 0x33,
	0xd0, 0x49, 0x69, 0x68, 0x3a, 0xcd, 0xcc, 0x58, 0xec, 0x1b, 0xf8, 0x68, 0x2c, 0x7d, 0x02, 0x63,
	0x78, 0x12, 0xd3, 0x2f, 0x8b, 0xc6, 0x8f, 0x5c, 0x76, 0xe7, 0xb4, 0xf9, 0xfd, 0xce, 0xe9, 0xff,
	0x14, 0x86, 0x31, 0x95, 0x07, 0xc6, 0xf7, 0x61, 0x1c, 0x8c, 0xd2, 0x27, 0x24, 0x4a, 0x76, 0x64,
	0x32, 0xa2, 0x9f, 0x25, 0xe5, 0x31, 0x89, 0x3e, 0x0a, 0xca, 0xd3, 0x70, 0x4b, 0x9d, 0x84, 0x33,
	0xc9, 0xf0, 0xa3, 0x50, 0xc8, 0x90, 0x39, 0x0d, 0xe1, 0xd4, 0x84, 0xf9, 0xf8, 0x4f, 0x9a, 0x80,
	0x48, 0x7a, 0x20, 0x59, 0x49, 0x9b, 0xdd, 0x80, 0x05, 0xac, 0x28, 0x47, 0x79, 0x55, 0x3e, 0x1d,
	0x7c, 0x69, 0xc3, 0x43, 0xb7, 0x1a, 0xb7, 0x2c, 0xa7, 0xe1, 0x2e, 0xb4, 0x77, 0x4c, 0x48, 0x61,
	0x20, 0x4b, 0xb5, 0x35, 0xaf, 0x6c, 0xf0, 0x0b, 0x68, 0x27, 0x8c, 0x4b, 0x61, 0x28, 0x96, 0x6a,
	0xdf, 0x8e, 0xfb, 0xce, 0x5f, 0xb7, 0x71, 0x3e, 0x30, 0x2e, 0x67, 0xad, 0xe3, 0xb7, 0xfe, 0x95,
	0x57, 0x32, 0xd8, 0x03, 0xcd, 0x0f, 0xc5, 0x96, 0xa5, 0x94, 0x67, 0x86, 0x6a, 0x21, 0xfb, 0xc1,
	0xf8, 0xe9, 0x3f, 0x04, 0xbf, 0x6d, 0xe4, 0xcc, 0x6b, 0xd6, 0x6b, 0x34, 0x78, 0x0d, 0x1a, 0x8d,
	0xfd, 0x84, 0x85, 0xb1, 0x14, 0x46, 0xab, 0x58, 0x6a, 0x72, 0x07, 0xa7, 0x5b, 0xb1, 0xd5, 0xa2,
	0x8d, 0xcb, 0x3c, 0x2a, 0x70, 0x53, 0xbf, 0xc5, 0x06, 0x5c, 0x13, 0xdf, 0xe7, 0x54, 0xe4, 0x71,
	0x20, 0x5b, 0xf3, 0xea, 0x16, 0xaf, 0x7e, 0x0d, 0xe4, 0xe5, 0x05, 0xb3, 0x8b, 0xa4, 0x84, 0x1b,
	0x4b, 0x9e, 0xd5, 0x49, 0xad, 0xa1, 0x13, 0x91, 0x0d, 0x8d, 0x84, 0xa1, 0x16, 0xda, 0x57, 0x97,
	0x68, 0xdf, 0x15, 0x86, 0xd2, 0x5b, 0xe9, 0xcc, 0x29, 0x40, 0x33, 0x0d, 0xeb, 0xa0, 0xee, 0x69,
	0x56, 0x7d, 0x52, 0x5e, 0xe6, 0x57, 0x4f, 0x49, 0xf4, 0x89, 0x1a, 0x8a, 0x85, 0xec, 0xfb, 0x5e,
	0xd9, 0x3c, 0x57, 0xa6, 0xc8, 0x7c, 0x06, 0xb7, 0x67, 0xc2, 0xff, 0xa1, 0xda, 0x19, 0x3a, 0x18,
	0x82, 0xf6, 0xf3, 0x76, 0xf8, 0x06, 0x5a, 0x8b, 0xf7, 0x0b, 0x57, 0xbf, 0xc2, 0x00, 0x9d, 0xe5,
	0xea, 0xf5, 0xea, 0xed, 0x1b, 0x1d, 0xe1, 0x6b, 0x50, 0xe7, 0x8b, 0xa5, 0xae, 0xcc, 0xee, 0x1d,
	0x4f, 0x3d, 0xf4, 0xf5, 0xd4, 0x43, 0xdf, 0x4f, 0x3d, 0xb4, 0xe9, 0x14, 0xff, 0xe7, 0xe4, 0x47,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x02, 0x41, 0xd3, 0xff, 0x21, 0x03, 0x00, 0x00,
}
