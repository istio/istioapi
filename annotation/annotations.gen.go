// Copyright 2019 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// GENERATED FILE -- DO NOT EDIT
//
// nolint: lll
//go:generate go run $GOPATH/src/istio.io/api/annotation/cmd/gen/main.go --input=$GOPATH/src/istio.io/api/annotation/annotations.yaml --output=$GOPATH/src/istio.io/api/annotation/annotations.gen.go
//go:generate goimports -w annotations.gen.go

package annotation

var (
	IngressClass = Instance{
		Name: "kubernetes.io/ingress.class",
		Description: "Annotation on ingress resources for the class of " +
			"controllers responsible for it.",
		Hidden:     false,
		Deprecated: false,
	}

	KubernetesServiceaccounts = Instance{
		Name: "alpha.istio.io/kubernetes-serviceaccounts",
		Description: "Specifies the Kubernetes service accounts that are " +
			"allowed to run this service on the VMs.",
		Hidden:     false,
		Deprecated: false,
	}

	CanonicalServiceaccounts = Instance{
		Name: "alpha.istio.io/canonical-serviceaccounts",
		Description: "Specifies the non-Kubernetes service accounts that are " +
			"allowed to run this service.",
		Hidden:     false,
		Deprecated: false,
	}

	ExportTo = Instance{
		Name: "networking.istio.io/exportTo",
		Description: "Specifies the namespaces to which this service should be " +
			"exported to. A value of '*' indicates it is reachable " +
			"within the mesh '.' indicates it is reachable within its " +
			"namespace.",
		Hidden:     false,
		Deprecated: false,
	}

	Identity = Instance{
		Name:        "alpha.istio.io/identity",
		Description: "",
		Hidden:      false,
		Deprecated:  false,
	}

	Inject = Instance{
		Name: "sidecar.istio.io/inject",
		Description: "Specifies whether or not an istio-proxy sidecar should be " +
			"automatically injected into the workload.",
		Hidden:     false,
		Deprecated: false,
	}

	Status = Instance{
		Name: "sidecar.istio.io/status",
		Description: "Generated by istio-proxy sidecar injection that indicates " +
			"the status of the operation. Includes a version hash of " +
			"the executed template, as well as names of injected " +
			"resources.",
		Hidden:     false,
		Deprecated: false,
	}

	RewriteAppHTTPProbers = Instance{
		Name: "sidecar.istio.io/rewriteAppHTTPProbers",
		Description: "Rewrite HTTP readiness and liveness probes to be " +
			"redirected to istio-proxy sidecar.",
		Hidden:     false,
		Deprecated: false,
	}

	ControlPlaneAuthPolicy = Instance{
		Name: "sidecar.istio.io/controlPlaneAuthPolicy",
		Description: "Specifies the auth policy used by the Istio control " +
			"plane. If NONE, traffic will not be encrypted. If " +
			"MUTUAL_TLS, traffic between istio-proxy sidecars will be " +
			"wrapped into mutual TLS connections.",
		Hidden:     false,
		Deprecated: false,
	}

	DiscoveryAddress = Instance{
		Name: "sidecar.istio.io/discoveryAddress",
		Description: "Specifies the XDS discovery address to be used by the " +
			"istio-proxy sidecar.",
		Hidden:     false,
		Deprecated: false,
	}

	ProxyImage = Instance{
		Name: "sidecar.istio.io/proxyImage",
		Description: "Specifies the Docker image to be used by the istio-proxy " +
			"sidecar.",
		Hidden:     false,
		Deprecated: false,
	}

	ProxyCPU = Instance{
		Name: "sidecar.istio.io/proxyCPU",
		Description: "Specifies the requested CPU setting for the istio-proxy " +
			"sidecar.",
		Hidden:     false,
		Deprecated: false,
	}

	ProxyMemory = Instance{
		Name: "sidecar.istio.io/proxyMemory",
		Description: "Specifies the requested memory setting for the " +
			"istio-proxy sidecar.",
		Hidden:     false,
		Deprecated: false,
	}

	InterceptionMode = Instance{
		Name: "sidecar.istio.io/interceptionMode",
		Description: "Specifies the mode used to redirect inbound connections " +
			"to Envoy (REDIRECT or TPROXY).",
		Hidden:     false,
		Deprecated: false,
	}

	BootstrapOverride = Instance{
		Name: "sidecar.istio.io/bootstrapOverride",
		Description: "Specifies an alternative Envoy bootstrap configuration " +
			"file.",
		Hidden:     false,
		Deprecated: false,
	}

	StatsInclusionPrefixes = Instance{
		Name: "sidecar.istio.io/statsInclusionPrefixes",
		Description: "Specifies the comma separated list of prefixes of the " +
			"stats to be emitted by Envoy.",
		Hidden:     false,
		Deprecated: false,
	}

	StatsInclusionSuffixes = Instance{
		Name: "sidecar.istio.io/statsInclusionSuffixes",
		Description: "Specifies the comma separated list of suffixes of the " +
			"stats to be emitted by Envoy.",
		Hidden:     false,
		Deprecated: false,
	}

	StatsInclusionRegexps = Instance{
		Name: "sidecar.istio.io/statsInclusionRegexps",
		Description: "Specifies the comma separated list of regexes the stats " +
			"should match to be emitted by Envoy.",
		Hidden:     false,
		Deprecated: false,
	}

	UserVolume = Instance{
		Name: "sidecar.istio.io/userVolume",
		Description: "Specifies one or more user volumes (as a JSON array) to " +
			"be added to the istio-proxy sidecar.",
		Hidden:     false,
		Deprecated: false,
	}

	UserVolumeMount = Instance{
		Name: "sidecar.istio.io/userVolumeMount",
		Description: "Specifies one or more user volume mounts (as a JSON " +
			"array) to be added to the istio-proxy sidecar.",
		Hidden:     false,
		Deprecated: false,
	}

	Port = Instance{
		Name: "status.sidecar.istio.io/port",
		Description: "Specifies the HTTP status Port for the istio-proxy " +
			"sidecar. If zero, the istio-proxy will not provide " +
			"status.",
		Hidden:     false,
		Deprecated: false,
	}

	InitialDelaySeconds = Instance{
		Name: "readiness.status.sidecar.istio.io/initialDelaySeconds",
		Description: "Specifies the initial delay (in seconds) for the " +
			"istio-proxy readiness probe.",
		Hidden:     false,
		Deprecated: false,
	}

	PeriodSeconds = Instance{
		Name: "readiness.status.sidecar.istio.io/periodSeconds",
		Description: "Specifies the period (in seconds) for the istio-proxy " +
			"readiness probe.",
		Hidden:     false,
		Deprecated: false,
	}

	FailureThreshold = Instance{
		Name: "readiness.status.sidecar.istio.io/failureThreshold",
		Description: "Specifies the failure threshold for the istio-proxy " +
			"readiness probe.",
		Hidden:     false,
		Deprecated: false,
	}

	ApplicationPorts = Instance{
		Name: "readiness.status.sidecar.istio.io/applicationPorts",
		Description: "Specifies the list of ports exposed by the application " +
			"container. Used by the istio-proxy readiness probe to " +
			"determine that Envoy is configured and ready to receive " +
			"traffic.",
		Hidden:     false,
		Deprecated: false,
	}

	IncludeOutboundIPRanges = Instance{
		Name: "traffic.sidecar.istio.io/includeOutboundIPRanges",
		Description: "A comma separated list of IP ranges in CIDR form to " +
			"redirect to envoy (optional). The wildcard character '*' " +
			"can be used to redirect all outbound traffic. An empty " +
			"list will disable all outbound redirection.",
		Hidden:     false,
		Deprecated: false,
	}

	ExcludeOutboundIPRanges = Instance{
		Name: "traffic.sidecar.istio.io/excludeOutboundIPRanges",
		Description: "A comma separated list of IP ranges in CIDR form to be " +
			"excluded from redirection. Only applies when all outbound " +
			"traffic (i.e. '*') is being redirected.",
		Hidden:     false,
		Deprecated: false,
	}

	IncludeInboundPorts = Instance{
		Name: "traffic.sidecar.istio.io/includeInboundPorts",
		Description: "A comma separated list of inbound ports for which traffic " +
			"is to be redirected to Envoy. The wildcard character '*' " +
			"can be used to configure redirection for all ports. An " +
			"empty list will disable all inbound redirection.",
		Hidden:     false,
		Deprecated: false,
	}

	ExcludeInboundPorts = Instance{
		Name: "traffic.sidecar.istio.io/excludeInboundPorts",
		Description: "A comma separated list of inbound ports to be excluded " +
			"from redirection to Envoy. Only applies when all inbound " +
			"traffic (i.e. '*') is being redirected.",
		Hidden:     false,
		Deprecated: false,
	}

	ExcludeOutboundPorts = Instance{
		Name: "traffic.sidecar.istio.io/excludeOutboundPorts",
		Description: "A comma separated list of outbound ports to be excluded " +
			"from redirection to Envoy.",
		Hidden:     false,
		Deprecated: false,
	}

	KubevirtInterfaces = Instance{
		Name: "traffic.sidecar.istio.io/kubevirtInterfaces",
		Description: "A comma separated list of virtual interfaces whose " +
			"inbound traffic (from VM) will be treated as outbound.",
		Hidden:     false,
		Deprecated: false,
	}
)