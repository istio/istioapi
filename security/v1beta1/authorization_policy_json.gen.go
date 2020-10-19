// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: security/v1beta1/authorization_policy.proto

// Istio Authorization Policy enables access control on workloads in the mesh.
//
// Authorization policy supports both allow and deny policies. When allow and
// deny policies are used for a workload at the same time, the deny policies are
// evaluated first. The evaluation is determined by the following rules:
//
// 1. If there are any DENY policies that match the request, deny the request.
// 2. If there are no ALLOW policies for the workload, allow the request.
// 3. If any of the ALLOW policies match the request, allow the request.
// 4. Deny the request.
//
// Istio Authorization Policy also supports the AUDIT action to decide whether to log requests.
// AUDIT policies do not affect whether requests are allowed or denied to the workload.
// Requests will be allowed or denied based solely on ALLOW and DENY policies.
//
// A request will be internally marked that it should be audited if there is an AUDIT policy on the workload that matches the request.
// A separate plugin must be configured and enabled to actually fulfill the audit decision and complete the audit behavior.
// The request will not be audited if there are no such supporting plugins enabled.
// Currently, the only supported plugin is the [Stackdriver](https://istio.io/latest/docs/reference/config/proxy_extensions/stackdriver/) plugin.
//
// Here is an example of Istio Authorization Policy:
//
// It sets the `action` to "ALLOW" to create an allow policy. The default action is "ALLOW"
// but it is useful to be explicit in the policy.
//
// It allows requests from:
//
// - service account "cluster.local/ns/default/sa/sleep" or
// - namespace "test"
//
// to access the workload with:
//
// - "GET" method at paths of prefix "/info" or,
// - "POST" method at path "/data".
//
// when the request has a valid JWT token issued by "https://accounts.google.com".
//
// Any other requests will be denied.
//
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//  name: httpbin
//  namespace: foo
// spec:
//  action: ALLOW
//  rules:
//  - from:
//    - source:
//        principals: ["cluster.local/ns/default/sa/sleep"]
//    - source:
//        namespaces: ["test"]
//    to:
//    - operation:
//        methods: ["GET"]
//        paths: ["/info*"]
//    - operation:
//        methods: ["POST"]
//        paths: ["/data"]
//    when:
//    - key: request.auth.claims[iss]
//      values: ["https://accounts.google.com"]
// ```
//
// The following is another example that sets `action` to "DENY" to create a deny policy.
// It denies requests from the "dev" namespace to the "POST" method on all workloads
// in the "foo" namespace.
//
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//  name: httpbin
//  namespace: foo
// spec:
//  action: DENY
//  rules:
//  - from:
//    - source:
//        namespaces: ["dev"]
//    to:
//    - operation:
//        methods: ["POST"]
// ```
//
// The following authorization policy sets the `action` to "AUDIT". It will audit any GET requests to the path with the
// prefix "/user/profile".
//
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//   namespace: ns1
//   name: anyname
// spec:
//   selector:
//     matchLabels:
//       app: myapi
//   action: audit
//   rules:
//   - to:
//     - operation:
//         methods: ["GET"]
//         paths: ["/user/profile/*"]
// ```
//
// Authorization Policy scope (target) is determined by "metadata/namespace" and
// an optional "selector".
//
// - "metadata/namespace" tells which namespace the policy applies. If set to root
// namespace, the policy applies to all namespaces in a mesh.
// - workload "selector" can be used to further restrict where a policy applies.
//
// For example,
//
// The following authorization policy applies to workloads containing label
// "app: httpbin" in namespace bar.
//
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//  name: policy
//  namespace: bar
// spec:
//  selector:
//    matchLabels:
//      app: httpbin
// ```
//
// The following authorization policy applies to all workloads in namespace foo.
//
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//  name: policy
//  namespace: foo
// spec:
//   {}
// ```
//
// The following authorization policy applies to workloads containing label
// "version: v1" in all namespaces in the mesh. (Assuming the root namespace is
// configured to "istio-config").
//
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//  name: policy
//  namespace: istio-config
// spec:
//  selector:
//    matchLabels:
//      version: v1
// ```

package v1beta1

import (
	bytes "bytes"
	fmt "fmt"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "istio.io/api/type/v1beta1"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for AuthorizationPolicy
func (this *AuthorizationPolicy) MarshalJSON() ([]byte, error) {
	str, err := AuthorizationPolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for AuthorizationPolicy
func (this *AuthorizationPolicy) UnmarshalJSON(b []byte) error {
	return AuthorizationPolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for AuthorizationPolicy_External
func (this *AuthorizationPolicy_External) MarshalJSON() ([]byte, error) {
	str, err := AuthorizationPolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for AuthorizationPolicy_External
func (this *AuthorizationPolicy_External) UnmarshalJSON(b []byte) error {
	return AuthorizationPolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Rule
func (this *Rule) MarshalJSON() ([]byte, error) {
	str, err := AuthorizationPolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Rule
func (this *Rule) UnmarshalJSON(b []byte) error {
	return AuthorizationPolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Rule_From
func (this *Rule_From) MarshalJSON() ([]byte, error) {
	str, err := AuthorizationPolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Rule_From
func (this *Rule_From) UnmarshalJSON(b []byte) error {
	return AuthorizationPolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Rule_To
func (this *Rule_To) MarshalJSON() ([]byte, error) {
	str, err := AuthorizationPolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Rule_To
func (this *Rule_To) UnmarshalJSON(b []byte) error {
	return AuthorizationPolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Source
func (this *Source) MarshalJSON() ([]byte, error) {
	str, err := AuthorizationPolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Source
func (this *Source) UnmarshalJSON(b []byte) error {
	return AuthorizationPolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Operation
func (this *Operation) MarshalJSON() ([]byte, error) {
	str, err := AuthorizationPolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Operation
func (this *Operation) UnmarshalJSON(b []byte) error {
	return AuthorizationPolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Condition
func (this *Condition) MarshalJSON() ([]byte, error) {
	str, err := AuthorizationPolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Condition
func (this *Condition) UnmarshalJSON(b []byte) error {
	return AuthorizationPolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	AuthorizationPolicyMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	AuthorizationPolicyUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{AllowUnknownFields: true}
)
