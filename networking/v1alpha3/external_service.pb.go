// Code generated by protoc-gen-go. DO NOT EDIT.
// source: networking/v1alpha3/external_service.proto

package v1alpha3

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
func (ExternalService_Discovery) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

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
//       destination:
//         name: example.com
//       tls:
//         mode: simple # initiates HTTPS when talking to example.com
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
//       destination:
//         name: 192.192.192.192/24
//       tls:
//         mode: mutual
//         clientCertificate: /etc/certs/myclientcert.pem
//         privateKey: /etc/certs/client_private_key.pem
//         caCertificates: /etc/certs/rootcacerts.pem
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
//       destination:
//         name: foo.bar.com
//       tls:
//         mode: simple # initiates HTTPS
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
	Ports []*Port `protobuf:"bytes,2,rep,name=ports" json:"ports,omitempty"`
	// Service discovery mode for the hosts. If not set, Istio will attempt
	// to infer the discovery mode based on the value of hosts and endpoints.
	Discovery ExternalService_Discovery `protobuf:"varint,3,opt,name=discovery,enum=istio.networking.v1alpha3.ExternalService_Discovery" json:"discovery,omitempty"`
	// One or more endpoints associated with the service. Endpoints must be
	// accessible over the set of outPorts defined at the service level.
	Endpoints []*ExternalService_Endpoint `protobuf:"bytes,4,rep,name=endpoints" json:"endpoints,omitempty"`
}

func (m *ExternalService) Reset()                    { *m = ExternalService{} }
func (m *ExternalService) String() string            { return proto.CompactTextString(m) }
func (*ExternalService) ProtoMessage()               {}
func (*ExternalService) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ExternalService) GetHosts() []string {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func (m *ExternalService) GetPorts() []*Port {
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

func (m *ExternalService) GetEndpoints() []*ExternalService_Endpoint {
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
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	// Set of ports associated with the endpoint. The ports must be
	// associated with a port name that was declared as part of the
	// service.
	Ports map[string]uint32 `protobuf:"bytes,2,rep,name=ports" json:"ports,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	// One or more labels associated with the endpoint.
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ExternalService_Endpoint) Reset()                    { *m = ExternalService_Endpoint{} }
func (m *ExternalService_Endpoint) String() string            { return proto.CompactTextString(m) }
func (*ExternalService_Endpoint) ProtoMessage()               {}
func (*ExternalService_Endpoint) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

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

func init() { proto.RegisterFile("networking/v1alpha3/external_service.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0x86, 0x2b, 0xc9, 0x5f, 0x1a, 0xd3, 0x56, 0x2c, 0x3d, 0xa8, 0xba, 0x54, 0xf5, 0x49, 0xf8,
	0x20, 0x51, 0xbb, 0x05, 0xb7, 0x87, 0x96, 0x26, 0xd6, 0x21, 0x10, 0x9c, 0x44, 0x36, 0x04, 0x72,
	0x09, 0x6b, 0x6b, 0xb1, 0x17, 0x0b, 0xad, 0xd8, 0xdd, 0xc8, 0xd1, 0x3f, 0xcc, 0x8f, 0xca, 0x21,
	0x58, 0x1f, 0x91, 0x13, 0x9c, 0x84, 0xf8, 0xa6, 0x11, 0xfb, 0x3c, 0xb3, 0xfb, 0xce, 0x40, 0x3f,
	0x26, 0x72, 0xc3, 0xf8, 0x9a, 0xc6, 0x4b, 0x2f, 0xfd, 0x81, 0xa3, 0x64, 0x85, 0x87, 0x1e, 0xb9,
	0x95, 0x84, 0xc7, 0x38, 0xba, 0x16, 0x84, 0xa7, 0x74, 0x41, 0xdc, 0x84, 0x33, 0xc9, 0xd0, 0x57,
	0x2a, 0x24, 0x65, 0x6e, 0x4d, 0xb8, 0x15, 0x61, 0x7d, 0xdf, 0xa7, 0x59, 0x62, 0x49, 0x36, 0x38,
	0x2b, 0xe8, 0xde, 0x7d, 0x03, 0x3e, 0xfb, 0xa5, 0x78, 0x5a, 0x78, 0xd1, 0x17, 0x68, 0xae, 0x98,
	0x90, 0xc2, 0x54, 0x6c, 0xcd, 0xd1, 0x83, 0xa2, 0x40, 0xbf, 0xa0, 0x99, 0x30, 0x2e, 0x85, 0xa9,
	0xda, 0x9a, 0xd3, 0x1d, 0x7c, 0x73, 0x5f, 0xec, 0xeb, 0x9e, 0x33, 0x2e, 0x83, 0xe2, 0x34, 0x0a,
	0x40, 0x0f, 0xa9, 0x58, 0xb0, 0x94, 0xf0, 0xcc, 0xd4, 0x6c, 0xc5, 0xf9, 0x34, 0xf8, 0xf9, 0x0a,
	0xfa, 0xec, 0x2e, 0xee, 0xb8, 0x62, 0x83, 0x5a, 0x83, 0x2e, 0x40, 0x27, 0x71, 0x98, 0x30, 0x1a,
	0x4b, 0x61, 0x36, 0xf2, 0xeb, 0x0c, 0xdf, 0xe1, 0xf4, 0x4b, 0x36, 0xa8, 0x2d, 0xd6, 0x9d, 0x0a,
	0x9d, 0xea, 0x3f, 0x32, 0xa1, 0x8d, 0xc3, 0x90, 0x13, 0xb1, 0x8d, 0x40, 0x71, 0xf4, 0xa0, 0x2a,
	0xd1, 0xec, 0x69, 0x08, 0x7f, 0x0f, 0xe8, 0x9a, 0xa7, 0x23, 0xfc, 0x58, 0xf2, 0xac, 0xca, 0xe8,
	0x12, 0x5a, 0x11, 0x9e, 0x93, 0x48, 0x98, 0x5a, 0xae, 0xfd, 0x77, 0x88, 0xf6, 0x34, 0x37, 0x14,
	0xde, 0x52, 0x67, 0x8d, 0x00, 0xea, 0x6e, 0xc8, 0x00, 0x6d, 0x4d, 0xb2, 0xf2, 0x49, 0xdb, 0xcf,
	0xed, 0xa4, 0x53, 0x1c, 0xdd, 0x10, 0x53, 0xb5, 0x15, 0xe7, 0x63, 0x50, 0x14, 0x7f, 0xd4, 0x91,
	0x62, 0xfd, 0x86, 0xee, 0x8e, 0xf0, 0x2d, 0x54, 0xdf, 0x41, 0x7b, 0x7d, 0xd0, 0x1f, 0xa7, 0x86,
	0x3a, 0xd0, 0x98, 0x9c, 0x4d, 0x7c, 0xe3, 0x03, 0x02, 0x68, 0x4d, 0x67, 0xff, 0x67, 0x27, 0xc7,
	0x86, 0x82, 0xda, 0xa0, 0x8d, 0x27, 0x53, 0x43, 0x3d, 0xea, 0x5d, 0xd9, 0xc5, 0x53, 0x29, 0xf3,
	0x70, 0x42, 0xbd, 0x3d, 0x0b, 0x3b, 0x6f, 0xe5, 0x9b, 0x3a, 0x7c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x83, 0x97, 0x1b, 0xdc, 0x15, 0x03, 0x00, 0x00,
}
