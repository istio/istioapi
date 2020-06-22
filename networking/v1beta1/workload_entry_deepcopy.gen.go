// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: networking/v1beta1/workload_entry.proto

// `WorkloadEntry` enables operators to describe the properties of a
// single non-Kubernetes workload such as a VM or a bare metal server
// as it is onboarded into the mesh. A `WorkloadEntry` must be
// accompanied by an Istio `ServiceEntry` that selects the workload
// through the appropriate labels and provides the service definition
// for a `MESH_INTERNAL` service (hostnames, port properties, etc.). A
// `ServiceEntry` object can select multiple workload entries as well
// as Kubernetes pods based on the label selector specified in the
// service entry.
//
// When a workload connects to `istiod`, the status field in the
// custom resource will be updated to indicate the health of the
// workload along with other details, similar to how Kubernetes
// updates the status of a pod.
//
// The following example declares a workload entry representing a VM
// for the `details.bookinfo.com` service. This VM has sidecar
// installed and bootstrapped using the `details-legacy` service
// account. The service is exposed by the sidecar on port 80 to
// applications in the mesh. The HTTP traffic is then wrapped in Istio
// mutual TLS and sent to sidecars on VMs on target port 8080, that in
// turn forward it to the application on localhost on the same port.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: WorkloadEntry
// metadata:
//   name: details-svc
// spec:
//   # use of the service account indicates that the workload has a
//   # sidecar proxy bootstrapped with this service account. Pods with
//   # sidecars will automatically communicate with the workload using
//   # istio mutual TLS.
//   serviceAccount: details-legacy
//   address: 2.2.2.2
//   labels:
//     app: details-legacy
//     instance-id: vm1
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: WorkloadEntry
// metadata:
//   name: details-svc
// spec:
//   # use of the service account indicates that the workload has a
//   # sidecar proxy bootstrapped with this service account. Pods with
//   # sidecars will automatically communicate with the workload using
//   # istio mutual TLS.
//   serviceAccount: details-legacy
//   address: 2.2.2.2
//   labels:
//     app: details-legacy
//     instance-id: vm1
// ```
// {{</tab>}}
// {{</tabset>}}
//
// and the associated service entry
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: details-svc
// spec:
//   hosts:
//   - details.bookinfo.com
//   location: MESH_INTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//     targetPort: 8080
//   resolution: STATIC
//   workloadSelector:
//     labels:
//       app: details-legacy
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: details-svc
// spec:
//   hosts:
//   - details.bookinfo.com
//   location: MESH_INTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//     targetPort: 8080
//   resolution: STATIC
//   workloadSelector:
//     labels:
//       app: details-legacy
// ```
// {{</tab>}}
// {{</tabset>}}
//
//
// The following example declares the same VM workload using
// its fully qualified DNS name. The service entry's resolution
// mode should be changed to DNS to indicate that the client-side
// sidecars should dynamically resolve the DNS name at runtime before
// forwarding the request.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: WorkloadEntry
// metadata:
//   name: details-svc
// spec:
//   # use of the service account indicates that the workload has a
//   # sidecar proxy bootstrapped with this service account. Pods with
//   # sidecars will automatically communicate with the workload using
//   # istio mutual TLS.
//   serviceAccount: details-legacy
//   address: vm1.vpc01.corp.net
//   labels:
//     app: details-legacy
//     instance-id: vm1
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: WorkloadEntry
// metadata:
//   name: details-svc
// spec:
//   # use of the service account indicates that the workload has a
//   # sidecar proxy bootstrapped with this service account. Pods with
//   # sidecars will automatically communicate with the workload using
//   # istio mutual TLS.
//   serviceAccount: details-legacy
//   address: vm1.vpc01.corp.net
//   labels:
//     app: details-legacy
//     instance-id: vm1
// ```
// {{</tab>}}
// {{</tabset>}}
//
// and the associated service entry
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: details-svc
// spec:
//   hosts:
//   - details.bookinfo.com
//   location: MESH_INTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//     targetPort: 8080
//   resolution: DNS
//   workloadSelector:
//     labels:
//       app: details-legacy
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: details-svc
// spec:
//   hosts:
//   - details.bookinfo.com
//   location: MESH_INTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//     targetPort: 8080
//   resolution: DNS
//   workloadSelector:
//     labels:
//       app: details-legacy
// ```
// {{</tab>}}
// {{</tabset>}}
//

package v1beta1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using WorkloadEntry within kubernetes types, where deepcopy-gen is used.
func (in *WorkloadEntry) DeepCopyInto(out *WorkloadEntry) {
	p := proto.Clone(in).(*WorkloadEntry)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadEntry. Required by controller-gen.
func (in *WorkloadEntry) DeepCopy() *WorkloadEntry {
	if in == nil {
		return nil
	}
	out := new(WorkloadEntry)
	in.DeepCopyInto(out)
	return out
}
