// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: operator/v1alpha1/operator.proto

// Configuration affecting Istio control plane installation version and shape.

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "google/protobuf/any.proto"
	_ "k8s.io/apimachinery/pkg/apis/meta/v1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using IstioOperatorSpec within kubernetes types, where deepcopy-gen is used.
func (in *IstioOperatorSpec) DeepCopyInto(out *IstioOperatorSpec) {
	p := proto.Clone(in).(*IstioOperatorSpec)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioOperatorSpec. Required by controller-gen.
func (in *IstioOperatorSpec) DeepCopy() *IstioOperatorSpec {
	if in == nil {
		return nil
	}
	out := new(IstioOperatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using InstallStatus within kubernetes types, where deepcopy-gen is used.
func (in *InstallStatus) DeepCopyInto(out *InstallStatus) {
	p := proto.Clone(in).(*InstallStatus)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallStatus. Required by controller-gen.
func (in *InstallStatus) DeepCopy() *InstallStatus {
	if in == nil {
		return nil
	}
	out := new(InstallStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using InstallStatus_VersionStatus within kubernetes types, where deepcopy-gen is used.
func (in *InstallStatus_VersionStatus) DeepCopyInto(out *InstallStatus_VersionStatus) {
	p := proto.Clone(in).(*InstallStatus_VersionStatus)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallStatus_VersionStatus. Required by controller-gen.
func (in *InstallStatus_VersionStatus) DeepCopy() *InstallStatus_VersionStatus {
	if in == nil {
		return nil
	}
	out := new(InstallStatus_VersionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using IstioComponentSetSpec within kubernetes types, where deepcopy-gen is used.
func (in *IstioComponentSetSpec) DeepCopyInto(out *IstioComponentSetSpec) {
	p := proto.Clone(in).(*IstioComponentSetSpec)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioComponentSetSpec. Required by controller-gen.
func (in *IstioComponentSetSpec) DeepCopy() *IstioComponentSetSpec {
	if in == nil {
		return nil
	}
	out := new(IstioComponentSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using BaseComponentSpec within kubernetes types, where deepcopy-gen is used.
func (in *BaseComponentSpec) DeepCopyInto(out *BaseComponentSpec) {
	p := proto.Clone(in).(*BaseComponentSpec)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaseComponentSpec. Required by controller-gen.
func (in *BaseComponentSpec) DeepCopy() *BaseComponentSpec {
	if in == nil {
		return nil
	}
	out := new(BaseComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using ComponentSpec within kubernetes types, where deepcopy-gen is used.
func (in *ComponentSpec) DeepCopyInto(out *ComponentSpec) {
	p := proto.Clone(in).(*ComponentSpec)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentSpec. Required by controller-gen.
func (in *ComponentSpec) DeepCopy() *ComponentSpec {
	if in == nil {
		return nil
	}
	out := new(ComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using ExternalComponentSpec within kubernetes types, where deepcopy-gen is used.
func (in *ExternalComponentSpec) DeepCopyInto(out *ExternalComponentSpec) {
	p := proto.Clone(in).(*ExternalComponentSpec)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalComponentSpec. Required by controller-gen.
func (in *ExternalComponentSpec) DeepCopy() *ExternalComponentSpec {
	if in == nil {
		return nil
	}
	out := new(ExternalComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using GatewaySpec within kubernetes types, where deepcopy-gen is used.
func (in *GatewaySpec) DeepCopyInto(out *GatewaySpec) {
	p := proto.Clone(in).(*GatewaySpec)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewaySpec. Required by controller-gen.
func (in *GatewaySpec) DeepCopy() *GatewaySpec {
	if in == nil {
		return nil
	}
	out := new(GatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using KubernetesResourcesSpec within kubernetes types, where deepcopy-gen is used.
func (in *KubernetesResourcesSpec) DeepCopyInto(out *KubernetesResourcesSpec) {
	p := proto.Clone(in).(*KubernetesResourcesSpec)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesResourcesSpec. Required by controller-gen.
func (in *KubernetesResourcesSpec) DeepCopy() *KubernetesResourcesSpec {
	if in == nil {
		return nil
	}
	out := new(KubernetesResourcesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using K8SObjectOverlay within kubernetes types, where deepcopy-gen is used.
func (in *K8SObjectOverlay) DeepCopyInto(out *K8SObjectOverlay) {
	p := proto.Clone(in).(*K8SObjectOverlay)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8SObjectOverlay. Required by controller-gen.
func (in *K8SObjectOverlay) DeepCopy() *K8SObjectOverlay {
	if in == nil {
		return nil
	}
	out := new(K8SObjectOverlay)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using K8SObjectOverlay_PathValue within kubernetes types, where deepcopy-gen is used.
func (in *K8SObjectOverlay_PathValue) DeepCopyInto(out *K8SObjectOverlay_PathValue) {
	p := proto.Clone(in).(*K8SObjectOverlay_PathValue)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8SObjectOverlay_PathValue. Required by controller-gen.
func (in *K8SObjectOverlay_PathValue) DeepCopy() *K8SObjectOverlay_PathValue {
	if in == nil {
		return nil
	}
	out := new(K8SObjectOverlay_PathValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using Affinity within kubernetes types, where deepcopy-gen is used.
func (in *Affinity) DeepCopyInto(out *Affinity) {
	p := proto.Clone(in).(*Affinity)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Affinity. Required by controller-gen.
func (in *Affinity) DeepCopy() *Affinity {
	if in == nil {
		return nil
	}
	out := new(Affinity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using ConfigMapKeySelector within kubernetes types, where deepcopy-gen is used.
func (in *ConfigMapKeySelector) DeepCopyInto(out *ConfigMapKeySelector) {
	p := proto.Clone(in).(*ConfigMapKeySelector)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapKeySelector. Required by controller-gen.
func (in *ConfigMapKeySelector) DeepCopy() *ConfigMapKeySelector {
	if in == nil {
		return nil
	}
	out := new(ConfigMapKeySelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using ClientIPConfig within kubernetes types, where deepcopy-gen is used.
func (in *ClientIPConfig) DeepCopyInto(out *ClientIPConfig) {
	p := proto.Clone(in).(*ClientIPConfig)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientIPConfig. Required by controller-gen.
func (in *ClientIPConfig) DeepCopy() *ClientIPConfig {
	if in == nil {
		return nil
	}
	out := new(ClientIPConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using CrossVersionObjectReference within kubernetes types, where deepcopy-gen is used.
func (in *CrossVersionObjectReference) DeepCopyInto(out *CrossVersionObjectReference) {
	p := proto.Clone(in).(*CrossVersionObjectReference)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrossVersionObjectReference. Required by controller-gen.
func (in *CrossVersionObjectReference) DeepCopy() *CrossVersionObjectReference {
	if in == nil {
		return nil
	}
	out := new(CrossVersionObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using DeploymentStrategy within kubernetes types, where deepcopy-gen is used.
func (in *DeploymentStrategy) DeepCopyInto(out *DeploymentStrategy) {
	p := proto.Clone(in).(*DeploymentStrategy)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentStrategy. Required by controller-gen.
func (in *DeploymentStrategy) DeepCopy() *DeploymentStrategy {
	if in == nil {
		return nil
	}
	out := new(DeploymentStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using EnvVar within kubernetes types, where deepcopy-gen is used.
func (in *EnvVar) DeepCopyInto(out *EnvVar) {
	p := proto.Clone(in).(*EnvVar)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvVar. Required by controller-gen.
func (in *EnvVar) DeepCopy() *EnvVar {
	if in == nil {
		return nil
	}
	out := new(EnvVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using EnvVarSource within kubernetes types, where deepcopy-gen is used.
func (in *EnvVarSource) DeepCopyInto(out *EnvVarSource) {
	p := proto.Clone(in).(*EnvVarSource)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvVarSource. Required by controller-gen.
func (in *EnvVarSource) DeepCopy() *EnvVarSource {
	if in == nil {
		return nil
	}
	out := new(EnvVarSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using ExecAction within kubernetes types, where deepcopy-gen is used.
func (in *ExecAction) DeepCopyInto(out *ExecAction) {
	p := proto.Clone(in).(*ExecAction)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecAction. Required by controller-gen.
func (in *ExecAction) DeepCopy() *ExecAction {
	if in == nil {
		return nil
	}
	out := new(ExecAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using ExternalMetricSource within kubernetes types, where deepcopy-gen is used.
func (in *ExternalMetricSource) DeepCopyInto(out *ExternalMetricSource) {
	p := proto.Clone(in).(*ExternalMetricSource)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalMetricSource. Required by controller-gen.
func (in *ExternalMetricSource) DeepCopy() *ExternalMetricSource {
	if in == nil {
		return nil
	}
	out := new(ExternalMetricSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using HTTPGetAction within kubernetes types, where deepcopy-gen is used.
func (in *HTTPGetAction) DeepCopyInto(out *HTTPGetAction) {
	p := proto.Clone(in).(*HTTPGetAction)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPGetAction. Required by controller-gen.
func (in *HTTPGetAction) DeepCopy() *HTTPGetAction {
	if in == nil {
		return nil
	}
	out := new(HTTPGetAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using HTTPHeader within kubernetes types, where deepcopy-gen is used.
func (in *HTTPHeader) DeepCopyInto(out *HTTPHeader) {
	p := proto.Clone(in).(*HTTPHeader)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPHeader. Required by controller-gen.
func (in *HTTPHeader) DeepCopy() *HTTPHeader {
	if in == nil {
		return nil
	}
	out := new(HTTPHeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using HorizontalPodAutoscalerSpec within kubernetes types, where deepcopy-gen is used.
func (in *HorizontalPodAutoscalerSpec) DeepCopyInto(out *HorizontalPodAutoscalerSpec) {
	p := proto.Clone(in).(*HorizontalPodAutoscalerSpec)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HorizontalPodAutoscalerSpec. Required by controller-gen.
func (in *HorizontalPodAutoscalerSpec) DeepCopy() *HorizontalPodAutoscalerSpec {
	if in == nil {
		return nil
	}
	out := new(HorizontalPodAutoscalerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using LocalObjectReference within kubernetes types, where deepcopy-gen is used.
func (in *LocalObjectReference) DeepCopyInto(out *LocalObjectReference) {
	p := proto.Clone(in).(*LocalObjectReference)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalObjectReference. Required by controller-gen.
func (in *LocalObjectReference) DeepCopy() *LocalObjectReference {
	if in == nil {
		return nil
	}
	out := new(LocalObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using MetricSpec within kubernetes types, where deepcopy-gen is used.
func (in *MetricSpec) DeepCopyInto(out *MetricSpec) {
	p := proto.Clone(in).(*MetricSpec)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricSpec. Required by controller-gen.
func (in *MetricSpec) DeepCopy() *MetricSpec {
	if in == nil {
		return nil
	}
	out := new(MetricSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using NodeAffinity within kubernetes types, where deepcopy-gen is used.
func (in *NodeAffinity) DeepCopyInto(out *NodeAffinity) {
	p := proto.Clone(in).(*NodeAffinity)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeAffinity. Required by controller-gen.
func (in *NodeAffinity) DeepCopy() *NodeAffinity {
	if in == nil {
		return nil
	}
	out := new(NodeAffinity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using NodeSelector within kubernetes types, where deepcopy-gen is used.
func (in *NodeSelector) DeepCopyInto(out *NodeSelector) {
	p := proto.Clone(in).(*NodeSelector)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeSelector. Required by controller-gen.
func (in *NodeSelector) DeepCopy() *NodeSelector {
	if in == nil {
		return nil
	}
	out := new(NodeSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using NodeSelectorTerm within kubernetes types, where deepcopy-gen is used.
func (in *NodeSelectorTerm) DeepCopyInto(out *NodeSelectorTerm) {
	p := proto.Clone(in).(*NodeSelectorTerm)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeSelectorTerm. Required by controller-gen.
func (in *NodeSelectorTerm) DeepCopy() *NodeSelectorTerm {
	if in == nil {
		return nil
	}
	out := new(NodeSelectorTerm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using NodeSelectorRequirement within kubernetes types, where deepcopy-gen is used.
func (in *NodeSelectorRequirement) DeepCopyInto(out *NodeSelectorRequirement) {
	p := proto.Clone(in).(*NodeSelectorRequirement)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeSelectorRequirement. Required by controller-gen.
func (in *NodeSelectorRequirement) DeepCopy() *NodeSelectorRequirement {
	if in == nil {
		return nil
	}
	out := new(NodeSelectorRequirement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using ObjectFieldSelector within kubernetes types, where deepcopy-gen is used.
func (in *ObjectFieldSelector) DeepCopyInto(out *ObjectFieldSelector) {
	p := proto.Clone(in).(*ObjectFieldSelector)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectFieldSelector. Required by controller-gen.
func (in *ObjectFieldSelector) DeepCopy() *ObjectFieldSelector {
	if in == nil {
		return nil
	}
	out := new(ObjectFieldSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using ObjectMeta within kubernetes types, where deepcopy-gen is used.
func (in *ObjectMeta) DeepCopyInto(out *ObjectMeta) {
	p := proto.Clone(in).(*ObjectMeta)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectMeta. Required by controller-gen.
func (in *ObjectMeta) DeepCopy() *ObjectMeta {
	if in == nil {
		return nil
	}
	out := new(ObjectMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using ObjectMetricSource within kubernetes types, where deepcopy-gen is used.
func (in *ObjectMetricSource) DeepCopyInto(out *ObjectMetricSource) {
	p := proto.Clone(in).(*ObjectMetricSource)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectMetricSource. Required by controller-gen.
func (in *ObjectMetricSource) DeepCopy() *ObjectMetricSource {
	if in == nil {
		return nil
	}
	out := new(ObjectMetricSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using PodAffinity within kubernetes types, where deepcopy-gen is used.
func (in *PodAffinity) DeepCopyInto(out *PodAffinity) {
	p := proto.Clone(in).(*PodAffinity)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodAffinity. Required by controller-gen.
func (in *PodAffinity) DeepCopy() *PodAffinity {
	if in == nil {
		return nil
	}
	out := new(PodAffinity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using PodAntiAffinity within kubernetes types, where deepcopy-gen is used.
func (in *PodAntiAffinity) DeepCopyInto(out *PodAntiAffinity) {
	p := proto.Clone(in).(*PodAntiAffinity)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodAntiAffinity. Required by controller-gen.
func (in *PodAntiAffinity) DeepCopy() *PodAntiAffinity {
	if in == nil {
		return nil
	}
	out := new(PodAntiAffinity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using PodAffinityTerm within kubernetes types, where deepcopy-gen is used.
func (in *PodAffinityTerm) DeepCopyInto(out *PodAffinityTerm) {
	p := proto.Clone(in).(*PodAffinityTerm)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodAffinityTerm. Required by controller-gen.
func (in *PodAffinityTerm) DeepCopy() *PodAffinityTerm {
	if in == nil {
		return nil
	}
	out := new(PodAffinityTerm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using PodDisruptionBudgetSpec within kubernetes types, where deepcopy-gen is used.
func (in *PodDisruptionBudgetSpec) DeepCopyInto(out *PodDisruptionBudgetSpec) {
	p := proto.Clone(in).(*PodDisruptionBudgetSpec)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodDisruptionBudgetSpec. Required by controller-gen.
func (in *PodDisruptionBudgetSpec) DeepCopy() *PodDisruptionBudgetSpec {
	if in == nil {
		return nil
	}
	out := new(PodDisruptionBudgetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using PodsMetricSource within kubernetes types, where deepcopy-gen is used.
func (in *PodsMetricSource) DeepCopyInto(out *PodsMetricSource) {
	p := proto.Clone(in).(*PodsMetricSource)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodsMetricSource. Required by controller-gen.
func (in *PodsMetricSource) DeepCopy() *PodsMetricSource {
	if in == nil {
		return nil
	}
	out := new(PodsMetricSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using PreferredSchedulingTerm within kubernetes types, where deepcopy-gen is used.
func (in *PreferredSchedulingTerm) DeepCopyInto(out *PreferredSchedulingTerm) {
	p := proto.Clone(in).(*PreferredSchedulingTerm)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreferredSchedulingTerm. Required by controller-gen.
func (in *PreferredSchedulingTerm) DeepCopy() *PreferredSchedulingTerm {
	if in == nil {
		return nil
	}
	out := new(PreferredSchedulingTerm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using ReadinessProbe within kubernetes types, where deepcopy-gen is used.
func (in *ReadinessProbe) DeepCopyInto(out *ReadinessProbe) {
	p := proto.Clone(in).(*ReadinessProbe)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadinessProbe. Required by controller-gen.
func (in *ReadinessProbe) DeepCopy() *ReadinessProbe {
	if in == nil {
		return nil
	}
	out := new(ReadinessProbe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using ResourceFieldSelector within kubernetes types, where deepcopy-gen is used.
func (in *ResourceFieldSelector) DeepCopyInto(out *ResourceFieldSelector) {
	p := proto.Clone(in).(*ResourceFieldSelector)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceFieldSelector. Required by controller-gen.
func (in *ResourceFieldSelector) DeepCopy() *ResourceFieldSelector {
	if in == nil {
		return nil
	}
	out := new(ResourceFieldSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using ResourceMetricSource within kubernetes types, where deepcopy-gen is used.
func (in *ResourceMetricSource) DeepCopyInto(out *ResourceMetricSource) {
	p := proto.Clone(in).(*ResourceMetricSource)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMetricSource. Required by controller-gen.
func (in *ResourceMetricSource) DeepCopy() *ResourceMetricSource {
	if in == nil {
		return nil
	}
	out := new(ResourceMetricSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using Resources within kubernetes types, where deepcopy-gen is used.
func (in *Resources) DeepCopyInto(out *Resources) {
	p := proto.Clone(in).(*Resources)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resources. Required by controller-gen.
func (in *Resources) DeepCopy() *Resources {
	if in == nil {
		return nil
	}
	out := new(Resources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using RollingUpdateDeployment within kubernetes types, where deepcopy-gen is used.
func (in *RollingUpdateDeployment) DeepCopyInto(out *RollingUpdateDeployment) {
	p := proto.Clone(in).(*RollingUpdateDeployment)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RollingUpdateDeployment. Required by controller-gen.
func (in *RollingUpdateDeployment) DeepCopy() *RollingUpdateDeployment {
	if in == nil {
		return nil
	}
	out := new(RollingUpdateDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using SecretKeySelector within kubernetes types, where deepcopy-gen is used.
func (in *SecretKeySelector) DeepCopyInto(out *SecretKeySelector) {
	p := proto.Clone(in).(*SecretKeySelector)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeySelector. Required by controller-gen.
func (in *SecretKeySelector) DeepCopy() *SecretKeySelector {
	if in == nil {
		return nil
	}
	out := new(SecretKeySelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using ServiceSpec within kubernetes types, where deepcopy-gen is used.
func (in *ServiceSpec) DeepCopyInto(out *ServiceSpec) {
	p := proto.Clone(in).(*ServiceSpec)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpec. Required by controller-gen.
func (in *ServiceSpec) DeepCopy() *ServiceSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using ServicePort within kubernetes types, where deepcopy-gen is used.
func (in *ServicePort) DeepCopyInto(out *ServicePort) {
	p := proto.Clone(in).(*ServicePort)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicePort. Required by controller-gen.
func (in *ServicePort) DeepCopy() *ServicePort {
	if in == nil {
		return nil
	}
	out := new(ServicePort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using SessionAffinityConfig within kubernetes types, where deepcopy-gen is used.
func (in *SessionAffinityConfig) DeepCopyInto(out *SessionAffinityConfig) {
	p := proto.Clone(in).(*SessionAffinityConfig)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SessionAffinityConfig. Required by controller-gen.
func (in *SessionAffinityConfig) DeepCopy() *SessionAffinityConfig {
	if in == nil {
		return nil
	}
	out := new(SessionAffinityConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using TCPSocketAction within kubernetes types, where deepcopy-gen is used.
func (in *TCPSocketAction) DeepCopyInto(out *TCPSocketAction) {
	p := proto.Clone(in).(*TCPSocketAction)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPSocketAction. Required by controller-gen.
func (in *TCPSocketAction) DeepCopy() *TCPSocketAction {
	if in == nil {
		return nil
	}
	out := new(TCPSocketAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using Toleration within kubernetes types, where deepcopy-gen is used.
func (in *Toleration) DeepCopyInto(out *Toleration) {
	p := proto.Clone(in).(*Toleration)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Toleration. Required by controller-gen.
func (in *Toleration) DeepCopy() *Toleration {
	if in == nil {
		return nil
	}
	out := new(Toleration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using WeightedPodAffinityTerm within kubernetes types, where deepcopy-gen is used.
func (in *WeightedPodAffinityTerm) DeepCopyInto(out *WeightedPodAffinityTerm) {
	p := proto.Clone(in).(*WeightedPodAffinityTerm)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeightedPodAffinityTerm. Required by controller-gen.
func (in *WeightedPodAffinityTerm) DeepCopy() *WeightedPodAffinityTerm {
	if in == nil {
		return nil
	}
	out := new(WeightedPodAffinityTerm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using PodSecurityContext within kubernetes types, where deepcopy-gen is used.
func (in *PodSecurityContext) DeepCopyInto(out *PodSecurityContext) {
	p := proto.Clone(in).(*PodSecurityContext)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSecurityContext. Required by controller-gen.
func (in *PodSecurityContext) DeepCopy() *PodSecurityContext {
	if in == nil {
		return nil
	}
	out := new(PodSecurityContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using SELinuxOptions within kubernetes types, where deepcopy-gen is used.
func (in *SELinuxOptions) DeepCopyInto(out *SELinuxOptions) {
	p := proto.Clone(in).(*SELinuxOptions)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SELinuxOptions. Required by controller-gen.
func (in *SELinuxOptions) DeepCopy() *SELinuxOptions {
	if in == nil {
		return nil
	}
	out := new(SELinuxOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using Sysctl within kubernetes types, where deepcopy-gen is used.
func (in *Sysctl) DeepCopyInto(out *Sysctl) {
	p := proto.Clone(in).(*Sysctl)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sysctl. Required by controller-gen.
func (in *Sysctl) DeepCopy() *Sysctl {
	if in == nil {
		return nil
	}
	out := new(Sysctl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using WindowsSecurityContextOptions within kubernetes types, where deepcopy-gen is used.
func (in *WindowsSecurityContextOptions) DeepCopyInto(out *WindowsSecurityContextOptions) {
	p := proto.Clone(in).(*WindowsSecurityContextOptions)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WindowsSecurityContextOptions. Required by controller-gen.
func (in *WindowsSecurityContextOptions) DeepCopy() *WindowsSecurityContextOptions {
	if in == nil {
		return nil
	}
	out := new(WindowsSecurityContextOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using SeccompProfile within kubernetes types, where deepcopy-gen is used.
func (in *SeccompProfile) DeepCopyInto(out *SeccompProfile) {
	p := proto.Clone(in).(*SeccompProfile)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeccompProfile. Required by controller-gen.
func (in *SeccompProfile) DeepCopy() *SeccompProfile {
	if in == nil {
		return nil
	}
	out := new(SeccompProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using TypeInterface within kubernetes types, where deepcopy-gen is used.
func (in *TypeInterface) DeepCopyInto(out *TypeInterface) {
	p := proto.Clone(in).(*TypeInterface)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TypeInterface. Required by controller-gen.
func (in *TypeInterface) DeepCopy() *TypeInterface {
	if in == nil {
		return nil
	}
	out := new(TypeInterface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using TypeMapStringInterface within kubernetes types, where deepcopy-gen is used.
func (in *TypeMapStringInterface) DeepCopyInto(out *TypeMapStringInterface) {
	p := proto.Clone(in).(*TypeMapStringInterface)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TypeMapStringInterface. Required by controller-gen.
func (in *TypeMapStringInterface) DeepCopy() *TypeMapStringInterface {
	if in == nil {
		return nil
	}
	out := new(TypeMapStringInterface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using TypeIntOrStringForPB within kubernetes types, where deepcopy-gen is used.
func (in *TypeIntOrStringForPB) DeepCopyInto(out *TypeIntOrStringForPB) {
	p := proto.Clone(in).(*TypeIntOrStringForPB)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TypeIntOrStringForPB. Required by controller-gen.
func (in *TypeIntOrStringForPB) DeepCopy() *TypeIntOrStringForPB {
	if in == nil {
		return nil
	}
	out := new(TypeIntOrStringForPB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using TypeBoolValueForPB within kubernetes types, where deepcopy-gen is used.
func (in *TypeBoolValueForPB) DeepCopyInto(out *TypeBoolValueForPB) {
	p := proto.Clone(in).(*TypeBoolValueForPB)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TypeBoolValueForPB. Required by controller-gen.
func (in *TypeBoolValueForPB) DeepCopy() *TypeBoolValueForPB {
	if in == nil {
		return nil
	}
	out := new(TypeBoolValueForPB)
	in.DeepCopyInto(out)
	return out
}
