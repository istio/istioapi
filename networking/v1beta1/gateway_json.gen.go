// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: networking/v1beta1/gateway.proto

// `Gateway` describes a load balancer operating at the edge of the mesh
// receiving incoming or outgoing HTTP/TCP connections. The specification
// describes a set of ports that should be exposed, the type of protocol to
// use, SNI configuration for the load balancer, etc.
//
// For example, the following Gateway configuration sets up a proxy to act
// as a load balancer exposing port 80 and 9080 (http), 443 (https),
// 9443(https) and port 2379 (TCP) for ingress.  The gateway will be
// applied to the proxy running on a pod with labels `app:
// my-gateway-controller`. While Istio will configure the proxy to listen
// on these ports, it is the responsibility of the user to ensure that
// external traffic to these ports are allowed into the mesh.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: Gateway
// metadata:
//   name: my-gateway
//   namespace: some-config-namespace
// spec:
//   selector:
//     app: my-gateway-controller
//   servers:
//   - port:
//       number: 80
//       name: http
//       protocol: HTTP
//     hosts:
//     - uk.bookinfo.com
//     - eu.bookinfo.com
//     tls:
//       httpsRedirect: true # sends 301 redirect for http requests
//   - port:
//       number: 443
//       name: https-443
//       protocol: HTTPS
//     hosts:
//     - uk.bookinfo.com
//     - eu.bookinfo.com
//     tls:
//       mode: SIMPLE # enables HTTPS on this port
//       serverCertificate: /etc/certs/servercert.pem
//       privateKey: /etc/certs/privatekey.pem
//   - port:
//       number: 9443
//       name: https-9443
//       protocol: HTTPS
//     hosts:
//     - "bookinfo-namespace/*.bookinfo.com"
//     tls:
//       mode: SIMPLE # enables HTTPS on this port
//       credentialName: bookinfo-secret # fetches certs from Kubernetes secret
//   - port:
//       number: 9080
//       name: http-wildcard
//       protocol: HTTP
//     hosts:
//     - "*"
//   - port:
//       number: 2379 # to expose internal service via external port 2379
//       name: mongo
//       protocol: MONGO
//     hosts:
//     - "*"
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: Gateway
// metadata:
//   name: my-gateway
//   namespace: some-config-namespace
// spec:
//   selector:
//     app: my-gateway-controller
//   servers:
//   - port:
//       number: 80
//       name: http
//       protocol: HTTP
//     hosts:
//     - uk.bookinfo.com
//     - eu.bookinfo.com
//     tls:
//       httpsRedirect: true # sends 301 redirect for http requests
//   - port:
//       number: 443
//       name: https-443
//       protocol: HTTPS
//     hosts:
//     - uk.bookinfo.com
//     - eu.bookinfo.com
//     tls:
//       mode: SIMPLE # enables HTTPS on this port
//       serverCertificate: /etc/certs/servercert.pem
//       privateKey: /etc/certs/privatekey.pem
//   - port:
//       number: 9443
//       name: https-9443
//       protocol: HTTPS
//     hosts:
//     - "bookinfo-namespace/*.bookinfo.com"
//     tls:
//       mode: SIMPLE # enables HTTPS on this port
//       credentialName: bookinfo-secret # fetches certs from Kubernetes secret
//   - port:
//       number: 9080
//       name: http-wildcard
//       protocol: HTTP
//     hosts:
//     - "*"
//   - port:
//       number: 2379 # to expose internal service via external port 2379
//       name: mongo
//       protocol: MONGO
//     hosts:
//     - "*"
// ```
// {{</tab>}}
// {{</tabset>}}
//
// The Gateway specification above describes the L4-L6 properties of a load
// balancer. A `VirtualService` can then be bound to a gateway to control
// the forwarding of traffic arriving at a particular host or gateway port.
//
// For example, the following VirtualService splits traffic for
// `https://uk.bookinfo.com/reviews`, `https://eu.bookinfo.com/reviews`,
// `http://uk.bookinfo.com:9080/reviews`,
// `http://eu.bookinfo.com:9080/reviews` into two versions (prod and qa) of
// an internal reviews service on port 9080. In addition, requests
// containing the cookie "user: dev-123" will be sent to special port 7777
// in the qa version. The same rule is also applicable inside the mesh for
// requests to the "reviews.prod.svc.cluster.local" service. This rule is
// applicable across ports 443, 9080. Note that `http://uk.bookinfo.com`
// gets redirected to `https://uk.bookinfo.com` (i.e. 80 redirects to 443).
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: VirtualService
// metadata:
//   name: bookinfo-rule
//   namespace: bookinfo-namespace
// spec:
//   hosts:
//   - reviews.prod.svc.cluster.local
//   - uk.bookinfo.com
//   - eu.bookinfo.com
//   gateways:
//   - some-config-namespace/my-gateway
//   - mesh # applies to all the sidecars in the mesh
//   http:
//   - match:
//     - headers:
//         cookie:
//           exact: "user=dev-123"
//     route:
//     - destination:
//         port:
//           number: 7777
//         host: reviews.qa.svc.cluster.local
//   - match:
//     - uri:
//         prefix: /reviews/
//     route:
//     - destination:
//         port:
//           number: 9080 # can be omitted if it's the only port for reviews
//         host: reviews.prod.svc.cluster.local
//       weight: 80
//     - destination:
//         host: reviews.qa.svc.cluster.local
//       weight: 20
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: VirtualService
// metadata:
//   name: bookinfo-rule
//   namespace: bookinfo-namespace
// spec:
//   hosts:
//   - reviews.prod.svc.cluster.local
//   - uk.bookinfo.com
//   - eu.bookinfo.com
//   gateways:
//   - some-config-namespace/my-gateway
//   - mesh # applies to all the sidecars in the mesh
//   http:
//   - match:
//     - headers:
//         cookie:
//           exact: "user=dev-123"
//     route:
//     - destination:
//         port:
//           number: 7777
//         host: reviews.qa.svc.cluster.local
//   - match:
//     - uri:
//         prefix: /reviews/
//     route:
//     - destination:
//         port:
//           number: 9080 # can be omitted if it's the only port for reviews
//         host: reviews.prod.svc.cluster.local
//       weight: 80
//     - destination:
//         host: reviews.qa.svc.cluster.local
//       weight: 20
// ```
// {{</tab>}}
// {{</tabset>}}
//
// The following VirtualService forwards traffic arriving at (external)
// port 27017 to internal Mongo server on port 5555. This rule is not
// applicable internally in the mesh as the gateway list omits the
// reserved name `mesh`.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: VirtualService
// metadata:
//   name: bookinfo-mongo
//   namespace: bookinfo-namespace
// spec:
//   hosts:
//   - mongosvr.prod.svc.cluster.local # name of internal Mongo service
//   gateways:
//   - some-config-namespace/my-gateway # can omit the namespace if gateway is
//   in same
//                                        namespace as virtual service.
//   tcp:
//   - match:
//     - port: 27017
//     route:
//     - destination:
//         host: mongo.prod.svc.cluster.local
//         port:
//           number: 5555
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: VirtualService
// metadata:
//   name: bookinfo-mongo
//   namespace: bookinfo-namespace
// spec:
//   hosts:
//   - mongosvr.prod.svc.cluster.local # name of internal Mongo service
//   gateways:
//   - some-config-namespace/my-gateway # can omit the namespace if gateway is
//   in same
//                                        namespace as virtual service.
//   tcp:
//   - match:
//     - port: 27017
//     route:
//     - destination:
//         host: mongo.prod.svc.cluster.local
//         port:
//           number: 5555
// ```
// {{</tab>}}
// {{</tabset>}}
//
// It is possible to restrict the set of virtual services that can bind to
// a gateway server using the namespace/hostname syntax in the hosts field.
// For example, the following Gateway allows any virtual service in the ns1
// namespace to bind to it, while restricting only the virtual service with
// foo.bar.com host in the ns2 namespace to bind to it.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: Gateway
// metadata:
//   name: my-gateway
//   namespace: some-config-namespace
// spec:
//   selector:
//     app: my-gateway-controller
//   servers:
//   - port:
//       number: 80
//       name: http
//       protocol: HTTP
//     hosts:
//     - "ns1/*"
//     - "ns2/foo.bar.com"
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: Gateway
// metadata:
//   name: my-gateway
//   namespace: some-config-namespace
// spec:
//   selector:
//     app: my-gateway-controller
//   servers:
//   - port:
//       number: 80
//       name: http
//       protocol: HTTP
//     hosts:
//     - "ns1/*"
//     - "ns2/foo.bar.com"
// ```
// {{</tab>}}
// {{</tabset>}}
//

package v1beta1

import (
	bytes "bytes"
	fmt "fmt"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for Gateway
func (this *Gateway) MarshalJSON() ([]byte, error) {
	str, err := GatewayMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Gateway
func (this *Gateway) UnmarshalJSON(b []byte) error {
	return GatewayUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Server
func (this *Server) MarshalJSON() ([]byte, error) {
	str, err := GatewayMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Server
func (this *Server) UnmarshalJSON(b []byte) error {
	return GatewayUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Port
func (this *Port) MarshalJSON() ([]byte, error) {
	str, err := GatewayMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Port
func (this *Port) UnmarshalJSON(b []byte) error {
	return GatewayUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ServerTLSSettings
func (this *ServerTLSSettings) MarshalJSON() ([]byte, error) {
	str, err := GatewayMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ServerTLSSettings
func (this *ServerTLSSettings) UnmarshalJSON(b []byte) error {
	return GatewayUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	GatewayMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	GatewayUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
