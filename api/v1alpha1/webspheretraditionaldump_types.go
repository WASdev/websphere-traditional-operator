package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// WebsphereTraditionalDumpSpec defines the desired state of WebsphereTraditionalDump
type WebsphereTraditionalDumpSpec struct {
	// The name of the Pod, which must be in the same namespace as the WebsphereTraditionalDump CR.
	PodName string `json:"podName"`
	// Optional. List of memory dump types to request: thread, heap, system.
	// +listType=set
	Include []WebsphereTraditionalDumpInclude `json:"include,omitempty"`
}

// Defines the possible values for dump types
// +kubebuilder:validation:Enum=dumpThreads;generateHeapDump;generateSystemDump;generateJavaCore
type WebsphereTraditionalDumpInclude string

const (
	WebsphereTraditionalDumpIncludeHeap     WebsphereTraditionalDumpInclude = "generateHeapDump"
	WebsphereTraditionalDumpIncludeThread   WebsphereTraditionalDumpInclude = "dumpThreads"
	WebsphereTraditionalDumpIncludeSystem   WebsphereTraditionalDumpInclude = "generateSystemDump"
	WebsphereTraditionalDumpIncludeJavaCore WebsphereTraditionalDumpInclude = "generateJavaCore"
)

// Defines the observed state of WebsphereTraditionalDump
type WebsphereTraditionalDumpStatus struct {
	// +listType=atomic
	Conditions []OperationStatusCondition `json:"conditions,omitempty"`
	// Location of the generated dump file
	// +operator-sdk:csv:customresourcedefinitions:type=status,displayName="Dump File Path",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	DumpFile string `json:"dumpFile,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=webspheretraditionaldumps,scope=Namespaced,shortName=oldump;oldumps
// +kubebuilder:printcolumn:name="Started",type="string",JSONPath=".status.conditions[?(@.type=='Started')].status",priority=0,description="Indicates if dump operation has started"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Started')].reason",priority=1,description="Reason for dump operation failing to start"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Started')].message",priority=1,description="Message for dump operation failing to start"
// +kubebuilder:printcolumn:name="Completed",type="string",JSONPath=".status.conditions[?(@.type=='Completed')].status",priority=0,description="Indicates if dump operation has completed"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Completed')].reason",priority=1,description="Reason for dump operation failing to complete"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Completed')].message",priority=1,description="Message for dump operation failing to complete"
// +kubebuilder:printcolumn:name="Dump file",type="string",JSONPath=".status.dumpFile",priority=0,description="Indicates filename of the server dump"
//+operator-sdk:csv:customresourcedefinitions:displayName="WebsphereTraditionalDump"
// Day-2 operation for generating server dumps
type WebsphereTraditionalDump struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WebsphereTraditionalDumpSpec   `json:"spec,omitempty"`
	Status WebsphereTraditionalDumpStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// WebsphereTraditionalDumpList contains a list of WebsphereTraditionalDump
type WebsphereTraditionalDumpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebsphereTraditionalDump `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WebsphereTraditionalDump{}, &WebsphereTraditionalDumpList{})
}
