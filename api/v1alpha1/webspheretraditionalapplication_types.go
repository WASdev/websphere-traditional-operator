/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"time"

	"github.com/application-stacks/runtime-component-operator/common"
	//	prometheusv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	routev1 "github.com/openshift/api/route/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// WebsphereTraditionalApplicationSpec defines the desired state of WebsphereTraditionalApplication
type WebsphereTraditionalApplicationSpec struct {
	// Application image to deploy.
	// +operator-sdk:csv:customresourcedefinitions:order=1,type=spec,displayName="Application Image",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	ApplicationImage string `json:"applicationImage"`

	// Name of the application. Defaults to the name of this custom resource.
	// +operator-sdk:csv:customresourcedefinitions:order=2,type=spec,displayName="Application Name",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	ApplicationName string `json:"applicationName,omitempty"`

	// Version of the application.
	// +operator-sdk:csv:customresourcedefinitions:order=3,type=spec,displayName="Application Version",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	ApplicationVersion string `json:"applicationVersion,omitempty"`

	// Policy for pulling container images. Defaults to IfNotPresent.
	// +operator-sdk:csv:customresourcedefinitions:order=4,type=spec,displayName="Pull Policy",xDescriptors="urn:alm:descriptor:com.tectonic.ui:imagePullPolicy"
	PullPolicy *corev1.PullPolicy `json:"pullPolicy,omitempty"`

	// Expose the application externally via a Route, a Knative Route or an Ingress resource.
	// +operator-sdk:csv:customresourcedefinitions:order=8,type=spec,displayName="Expose",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Expose *bool `json:"expose,omitempty"`

	// Resource requests and limits for the application container.
	// +operator-sdk:csv:customresourcedefinitions:order=11,type=spec,displayName="Resource Requirements",xDescriptors="urn:alm:descriptor:com.tectonic.ui:resourceRequirements"
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`

	// +operator-sdk:csv:customresourcedefinitions:order=13,type=spec,displayName="Deployment"
	Deployment *WebsphereTraditionalApplicationDeployment `json:"deployment,omitempty"`

	// +operator-sdk:csv:customresourcedefinitions:order=15,type=spec,displayName="Service"
	Service *WebsphereTraditionalApplicationService `json:"service,omitempty"`

	// +operator-sdk:csv:customresourcedefinitions:order=16,type=spec,displayName="Route"
	Route *WebsphereTraditionalApplicationRoute `json:"route,omitempty"`

	// +operator-sdk:csv:customresourcedefinitions:order=26,type=spec,displayName="Affinity"
	Affinity *WebsphereTraditionalApplicationAffinity `json:"affinity,omitempty"`
}

type StatusCondition struct {
	LastTransitionTime *metav1.Time           `json:"lastTransitionTime,omitempty"`
	Reason             string                 `json:"reason,omitempty"`
	Message            string                 `json:"message,omitempty"`
	Status             corev1.ConditionStatus `json:"status,omitempty"`
	Type               StatusConditionType    `json:"type,omitempty"`
}

// Defines the type of status condition.
type StatusConditionType string

const (
	// StatusConditionTypeReconciled ...
	StatusConditionTypeReconciled StatusConditionType = "Reconciled"
)

// Configure pods to run on particular Nodes.
type WebsphereTraditionalApplicationAffinity struct {
	// Controls which nodes the pod are scheduled to run on, based on labels on the node.
	// +operator-sdk:csv:customresourcedefinitions:order=37,type=spec,displayName="Node Affinity",xDescriptors="urn:alm:descriptor:com.tectonic.ui:nodeAffinity"
	NodeAffinity *corev1.NodeAffinity `json:"nodeAffinity,omitempty"`

	// Controls the nodes the pod are scheduled to run on, based on labels on the pods that are already running on the node.
	// +operator-sdk:csv:customresourcedefinitions:order=38,type=spec,displayName="Pod Affinity",xDescriptors="urn:alm:descriptor:com.tectonic.ui:podAffinity"
	PodAffinity *corev1.PodAffinity `json:"podAffinity,omitempty"`

	// Enables the ability to prevent running a pod on the same node as another pod.
	// +operator-sdk:csv:customresourcedefinitions:order=39,type=spec,displayName="Pod Anti Affinity",xDescriptors="urn:alm:descriptor:com.tectonic.ui:podAntiAffinity"
	PodAntiAffinity *corev1.PodAntiAffinity `json:"podAntiAffinity,omitempty"`

	// A YAML object that contains a set of required labels and their values.
	// +operator-sdk:csv:customresourcedefinitions:order=40,type=spec,displayName="Node Affinity Labels",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	NodeAffinityLabels map[string]string `json:"nodeAffinityLabels,omitempty"`

	// An array of architectures to be considered for deployment. Their position in the array indicates preference.
	// +listType=set
	Architecture []string `json:"architecture,omitempty"`
}

// Configures parameters for the network service of pods.

type WebsphereTraditionalApplicationService struct {
	// The port exposed by the container.
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:validation:Minimum=1
	// +operator-sdk:csv:customresourcedefinitions:order=9,type=spec,displayName="Service Port",xDescriptors="urn:alm:descriptor:com.tectonic.ui:number"
	Port int32 `json:"port,omitempty"`

	// +operator-sdk:csv:customresourcedefinitions:order=10,type=spec,displayName="Service Type",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	Type *corev1.ServiceType `json:"type,omitempty"`

	// Node proxies this port into your service.
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:validation:Minimum=0
	// +operator-sdk:csv:customresourcedefinitions:order=11,type=spec,displayName="Node Port",xDescriptors="urn:alm:descriptor:com.tectonic.ui:number"
	NodePort *int32 `json:"nodePort,omitempty"`

	// The name for the port exposed by the container.
	// +operator-sdk:csv:customresourcedefinitions:order=12,type=spec,displayName="Port Name",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	PortName string `json:"portName,omitempty"`

	// Annotations to be added to the service.
	// +operator-sdk:csv:customresourcedefinitions:order=13,type=spec,displayName="Service Annotations",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	Annotations map[string]string `json:"annotations,omitempty"`

	// The port that the operator assigns to containers inside pods. Defaults to the value of .spec.service.port.
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:validation:Minimum=1
	// +operator-sdk:csv:customresourcedefinitions:order=14,type=spec,displayName="Target Port",xDescriptors="urn:alm:descriptor:com.tectonic.ui:number"
	TargetPort *int32 `json:"targetPort,omitempty"`

	// A name of a secret that already contains TLS key, certificate and CA to be mounted in the pod. The following keys are valid in the secret: ca.crt, tls.crt, and tls.key.
	// +operator-sdk:csv:customresourcedefinitions:order=15,type=spec,displayName="Certificate Secret Reference",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	CertificateSecretRef *string `json:"certificateSecretRef,omitempty"`

	// An array consisting of service ports.
	// +operator-sdk:csv:customresourcedefinitions:order=16,type=spec
	Ports []corev1.ServicePort `json:"ports,omitempty"`

	// Expose the application as a bindable service. Defaults to false.
	// +operator-sdk:csv:customresourcedefinitions:order=17,type=spec,displayName="Bindable",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Bindable *bool `json:"bindable,omitempty"`
}

// Defines the desired state and cycle of applications.
type WebsphereTraditionalApplicationDeployment struct {
	// Specifies the strategy to replace old deployment pods with new pods.
	// +operator-sdk:csv:customresourcedefinitions:order=21,type=spec,displayName="Deployment Update Strategy",xDescriptors="urn:alm:descriptor:com.tectonic.ui:updateStrategy"
	UpdateStrategy *appsv1.DeploymentStrategy `json:"updateStrategy,omitempty"`

	// Annotations to be added only to the Deployment and resources owned by the Deployment
	Annotations map[string]string `json:"annotations,omitempty"`
}

// Configures the ingress resource.
type WebsphereTraditionalApplicationRoute struct {

	// Annotations to be added to the Route.
	// +operator-sdk:csv:customresourcedefinitions:order=42,type=spec,displayName="Route Annotations",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	Annotations map[string]string `json:"annotations,omitempty"`

	// Hostname to be used for the Route.
	// +operator-sdk:csv:customresourcedefinitions:order=43,type=spec,displayName="Route Host",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	Host string `json:"host,omitempty"`

	// Path to be used for Route.
	// +operator-sdk:csv:customresourcedefinitions:order=44,type=spec,displayName="Route Path",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	Path string `json:"path,omitempty"`

	// Path type to be used for Ingress.
	PathType networkingv1.PathType `json:"pathType,omitempty"`

	// A name of a secret that already contains TLS key, certificate and CA to be used in the route. It can also contain destination CA certificate. The following keys are valid in the secret: ca.crt, destCA.crt, tls.crt, and tls.key.
	// +operator-sdk:csv:customresourcedefinitions:order=45,type=spec,displayName="Certificate Secret Reference",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	CertificateSecretRef *string `json:"certificateSecretRef,omitempty"`

	// TLS termination policy. Can be one of edge, reencrypt and passthrough.
	// +operator-sdk:csv:customresourcedefinitions:order=46,type=spec,displayName="Termination",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	Termination *routev1.TLSTerminationType `json:"termination,omitempty"`

	// HTTP traffic policy with TLS enabled. Can be one of Allow, Redirect and None.
	// +operator-sdk:csv:customresourcedefinitions:order=47,type=spec,displayName="Insecure Edge Termination Policy",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	InsecureEdgeTerminationPolicy *routev1.InsecureEdgeTerminationPolicyType `json:"insecureEdgeTerminationPolicy,omitempty"`
}

// WebsphereTraditionalApplicationStatus defines the observed state of WebsphereTraditionalApplication
type WebsphereTraditionalApplicationStatus struct {
	// +listType=atomic
	// +operator-sdk:csv:customresourcedefinitions:type=status,displayName="Status Conditions",xDescriptors="urn:alm:descriptor:io.kubernetes.conditions"
	Conditions     []StatusCondition `json:"conditions,omitempty"`
	RouteAvailable *bool             `json:"routeAvailable,omitempty"`
	ImageReference string            `json:"imageReference,omitempty"`

	// +operator-sdk:csv:customresourcedefinitions:type=status,displayName="Service Binding"
	Binding *corev1.LocalObjectReference `json:"binding,omitempty"`
}

// +kubebuilder:resource:path=webspheretradtitionalapplications,scope=Namespaced,shortName=wtapp;wtapps
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Image",type="string",JSONPath=".spec.applicationImage",priority=0,description="Absolute name of the deployed image containing registry and tag"
// +kubebuilder:printcolumn:name="Exposed",type="boolean",JSONPath=".spec.expose",priority=0,description="Specifies whether deployment is exposed externally via default Route"
// +kubebuilder:printcolumn:name="Reconciled",type="string",JSONPath=".status.conditions[?(@.type=='Reconciled')].status",priority=0,description="Status of the reconcile condition"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Reconciled')].reason",priority=1,description="Reason for the failure of reconcile condition"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Reconciled')].message",priority=1,description="Failure message from reconcile condition"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",priority=0,description="Age of the resource"
//+operator-sdk:csv:customresourcedefinitions:displayName="WebsphereTradtitionalApplication",resources={{Deployment,v1},{Service,v1},{StatefulSet,v1},{Route,v1},{HorizontalPodAutoscaler,v1},{ServiceAccount,v1},{Secret,v1}}

// WebsphereTraditionalApplication is the Schema for the webspheretraditionalapplications API
type WebsphereTraditionalApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WebsphereTraditionalApplicationSpec   `json:"spec,omitempty"`
	Status WebsphereTraditionalApplicationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// WebsphereTraditionalApplicationList contains a list of WebsphereTraditionalApplication
type WebsphereTraditionalApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebsphereTraditionalApplication `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WebsphereTraditionalApplication{}, &WebsphereTraditionalApplicationList{})
}

// GetApplicationImage returns application image
func (cr *WebsphereTraditionalApplication) GetApplicationImage() string {
	return cr.Spec.ApplicationImage
}

// GetExpose returns expose flag
func (cr *WebsphereTraditionalApplication) GetExpose() *bool {
	return cr.Spec.Expose
}

// GetService returns service settings
func (cr *WebsphereTraditionalApplication) GetService() common.BaseComponentService {
	if cr.Spec.Service == nil {
		return nil
	}
	return cr.Spec.Service
}

// GetApplicationName returns Application name
func (cr *WebsphereTraditionalApplication) GetApplicationName() string {
	return cr.Spec.ApplicationName
}

// GetStatus returns WebsphereTraditionalApplication status
func (cr *WebsphereTraditionalApplication) GetStatus() common.BaseComponentStatus {
	return &cr.Status
}

// GetGroupName returns group name to be used in labels and annotation
func (cr *WebsphereTraditionalApplication) GetGroupName() string {
	return "apps.webspheretraditional.io"
}

// GetAnnotations returns a set of annotations to be added to the service
func (s *WebsphereTraditionalApplicationService) GetAnnotations() map[string]string {
	return s.Annotations
}

// GetBinding returns BindingStatus representing binding status
func (s *WebsphereTraditionalApplicationStatus) GetBinding() *corev1.LocalObjectReference {
	return s.Binding
}

// GetAnnotations returns route annotations
func (r *WebsphereTraditionalApplicationRoute) GetAnnotations() map[string]string {
	return r.Annotations
}

// GetArchitecture returns list of architecture names
func (a *WebsphereTraditionalApplicationAffinity) GetArchitecture() []string {
	return a.Architecture
}

// GetAnnotations returns annotations to be added only to the Deployment and its child resources
func (rcd *WebsphereTraditionalApplicationDeployment) GetAnnotations() map[string]string {
	return rcd.Annotations
}

// GetBindable returns whether the application should be exposable as a service
func (s *WebsphereTraditionalApplicationService) GetBindable() *bool {
	return s.Bindable
}

// GetConditions returns slice of conditions
func (s *WebsphereTraditionalApplicationStatus) GetConditions() []common.StatusCondition {
	var conditions = make([]common.StatusCondition, len(s.Conditions))
	for i := range s.Conditions {
		conditions[i] = &s.Conditions[i]
	}
	return conditions
}

// GetCondition ...
func (s *WebsphereTraditionalApplicationStatus) GetCondition(t common.StatusConditionType) common.StatusCondition {
	for i := range s.Conditions {
		if s.Conditions[i].GetType() == t {
			return &s.Conditions[i]
		}
	}
	return nil
}

// GetCertificateSecretRef returns a secret reference with a certificate
func (r *WebsphereTraditionalApplicationRoute) GetCertificateSecretRef() *string {
	return r.CertificateSecretRef
}

// GetCertificateSecretRef returns a secret reference with a certificate
func (s *WebsphereTraditionalApplicationService) GetCertificateSecretRef() *string {
	return s.CertificateSecretRef
}

// GetImageReference returns Docker image reference to be deployed by the CR
func (s *WebsphereTraditionalApplicationStatus) GetImageReference() string {
	return s.ImageReference
}

// GetLastTransitionTime return time of last status change
func (c *StatusCondition) GetLastTransitionTime() *metav1.Time {
	return c.LastTransitionTime
}

// GetType returns status condition type
func (c *StatusCondition) GetType() common.StatusConditionType {
	return convertToCommonStatusConditionType(c.Type)
}

// GetHost returns hostname to be used by the route
func (r *WebsphereTraditionalApplicationRoute) GetHost() string {
	return r.Host
}

// GetNodeAffinityLabels returns list of architecture names
func (a *WebsphereTraditionalApplicationAffinity) GetNodeAffinityLabels() map[string]string {
	return a.NodeAffinityLabels
}

// GetNodeAffinity returns node affinity
func (a *WebsphereTraditionalApplicationAffinity) GetNodeAffinity() *corev1.NodeAffinity {
	return a.NodeAffinity
}

// GetNodePort returns service nodePort
func (s *WebsphereTraditionalApplicationService) GetNodePort() *int32 {
	if s.NodePort == nil {
		return nil
	}
	return s.NodePort
}

// NewCondition returns new condition
func (s *WebsphereTraditionalApplicationStatus) NewCondition() common.StatusCondition {
	return &StatusCondition{}
}

// GetMessage return condition's message
func (c *StatusCondition) GetMessage() string {
	return c.Message
}

// GetInsecureEdgeTerminationPolicy returns terminatation of the route's TLS
func (r *WebsphereTraditionalApplicationRoute) GetInsecureEdgeTerminationPolicy() *routev1.InsecureEdgeTerminationPolicyType {
	return r.InsecureEdgeTerminationPolicy
}

// GetPodAffinity returns pod affinity
func (a *WebsphereTraditionalApplicationAffinity) GetPodAffinity() *corev1.PodAffinity {
	return a.PodAffinity
}

// GetPortName returns name of service port
func (s *WebsphereTraditionalApplicationService) GetPortName() string {
	return s.PortName
}

// SetBinding sets BindingStatus representing binding status
func (s *WebsphereTraditionalApplicationStatus) SetBinding(r *corev1.LocalObjectReference) {
	s.Binding = r
}

// GetReason return condition's message
func (c *StatusCondition) GetReason() string {
	return c.Reason
}

// GetPath returns path to use for the route
func (r *WebsphereTraditionalApplicationRoute) GetPath() string {
	return r.Path
}

// GetPathType returns pathType to use for the route
func (r *WebsphereTraditionalApplicationRoute) GetPathType() networkingv1.PathType {
	return r.PathType
}

// GetPorts returns a list of service ports
func (s *WebsphereTraditionalApplicationService) GetPorts() []corev1.ServicePort {
	return s.Ports
}

// SetCondition ...
func (s *WebsphereTraditionalApplicationStatus) SetCondition(c common.StatusCondition) {
	condition := &StatusCondition{}
	found := false
	for i := range s.Conditions {
		if s.Conditions[i].GetType() == c.GetType() {
			condition = &s.Conditions[i]
			found = true
		}
	}

	if condition.GetStatus() != c.GetStatus() {
		condition.SetLastTransitionTime(&metav1.Time{Time: time.Now()})
	}

	condition.SetReason(c.GetReason())
	condition.SetMessage(c.GetMessage())
	condition.SetStatus(c.GetStatus())
	condition.SetType(c.GetType())
	if !found {
		s.Conditions = append(s.Conditions, *condition)
	}
}

// GetStatus return condition's status
func (c *StatusCondition) GetStatus() corev1.ConditionStatus {
	return c.Status
}

// GetTermination returns terminatation of the route's TLS
func (r *WebsphereTraditionalApplicationRoute) GetTermination() *routev1.TLSTerminationType {
	return r.Termination
}

// GetAutoscaling returns autoscaling settings
func (cr *WebsphereTraditionalApplication) GetAutoscaling() common.BaseComponentAutoscaling {
	if cr.Spec.Autoscaling == nil {
		return nil
	}
	return cr.Spec.Autoscaling
}

// SetImageReference sets Docker image reference on the status portion of the CR
func (s *WebsphereTraditionalApplicationStatus) SetImageReference(imageReference string) {
	s.ImageReference = imageReference
}

// GetTargetPort returns the internal target port for containers
func (s *WebsphereTraditionalApplicationService) GetTargetPort() *int32 {
	return s.TargetPort
}

// GetType returns service type
func (s *WebsphereTraditionalApplicationService) GetType() *corev1.ServiceType {
	return s.Type
}

// GetApplicationVersion returns application version
func (cr *WebsphereTraditionalApplication) GetApplicationVersion() string {
	return cr.Spec.ApplicationVersion
}

// SetType returns status condition type
func (c *StatusCondition) SetType(ct common.StatusConditionType) {
	c.Type = convertFromCommonStatusConditionType(ct)
}

// SetLastTransitionTime sets time of last status change
func (c *StatusCondition) SetLastTransitionTime(t *metav1.Time) {
	c.LastTransitionTime = t
}

// GetPodAntiAffinity returns pod anti-affinity
func (a *WebsphereTraditionalApplicationAffinity) GetPodAntiAffinity() *corev1.PodAntiAffinity {
	return a.PodAntiAffinity
}

// GetRoute returns route
func (cr *WebsphereTraditionalApplication) GetRoute() common.BaseComponentRoute {
	if cr.Spec.Route == nil {
		return nil
	}
	return cr.Spec.Route
}

// SetReason sets condition's reason
func (c *StatusCondition) SetReason(r string) {
	c.Reason = r
}

// SetMessage sets condition's message
func (c *StatusCondition) SetMessage(m string) {
	c.Message = m
}

// SetStatus sets condition's status
func (c *StatusCondition) SetStatus(s corev1.ConditionStatus) {
	c.Status = s
}

// GetAffinity returns deployment's node and pod affinity settings
func (cr *WebsphereTraditionalApplication) GetAffinity() common.BaseComponentAffinity {
	if cr.Spec.Affinity == nil {
		return nil
	}
	return cr.Spec.Affinity
}

// GetDeployment returns deployment settings
func (cr *WebsphereTraditionalApplication) GetDeployment() common.BaseComponentDeployment {
	if cr.Spec.Deployment == nil {
		return nil
	}
	return cr.Spec.Deployment
}

// GetDeploymentStrategy returns deployment strategy struct
func (cr *WebsphereTraditionalApplicationDeployment) GetDeploymentUpdateStrategy() *appsv1.DeploymentStrategy {
	return cr.UpdateStrategy
}

// GetPort returns service port
func (s *WebsphereTraditionalApplicationService) GetPort() int32 {
	if s != nil && s.Port != 0 {
		return s.Port
	}
	return 9080
}

// Initialize sets default values
func (cr *WebsphereTraditionalApplication) Initialize() {
	if cr.Spec.PullPolicy == nil {
		pp := corev1.PullIfNotPresent
		cr.Spec.PullPolicy = &pp
	}

	if cr.Spec.Resources == nil {
		cr.Spec.Resources = &corev1.ResourceRequirements{}
	}

	// Default applicationName to cr.Name, if a user sets createAppDefinition to true but doesn't set applicationName
	if cr.Spec.ApplicationName == "" {
		if cr.Labels != nil && cr.Labels["app.kubernetes.io/part-of"] != "" {
			cr.Spec.ApplicationName = cr.Labels["app.kubernetes.io/part-of"]
		} else {
			cr.Spec.ApplicationName = cr.Name
		}
	}

	if cr.Labels != nil {
		cr.Labels["app.kubernetes.io/part-of"] = cr.Spec.ApplicationName
	}

	// This is to handle when there is no service in the CR
	if cr.Spec.Service == nil {
		cr.Spec.Service = &WebsphereTraditionalApplicationService{}
	}

	if cr.Spec.Service.Type == nil {
		st := corev1.ServiceTypeClusterIP
		cr.Spec.Service.Type = &st
	}

	if cr.Spec.Service.Port == 0 {
		cr.Spec.Service.Port = 9080
	}

}

func convertToCommonStatusConditionType(c StatusConditionType) common.StatusConditionType {
	switch c {
	case StatusConditionTypeReconciled:
		return common.StatusConditionTypeReconciled
	default:
		panic(c)
	}
}

func convertFromCommonStatusConditionType(c common.StatusConditionType) StatusConditionType {
	switch c {
	case common.StatusConditionTypeReconciled:
		return StatusConditionTypeReconciled
	default:
		panic(c)
	}
}
