package v0alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Scope struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ScopeSpec `json:"spec,omitempty"`
}

type ScopeSpec struct {
	// The name displayed in the UI
	Title string `json:"title"`
	// Scope type?? should this be an enumeration
	Type string `json:"type"`
	// Optional scope description
	Description string `json:"description,omitempty"`
	// User configured category
	Category string `json:"category,omitempty"`
	// Filters that can get added ???
	// +listType=atomic
	Filters []ScopeFilter `json:"filters"`
}

type ScopeFilter struct {
	Key   string `json:"key"`
	Value string `json:"value"`

	// TODO? should this be an enumeration? or is it really anything
	Operator string `json:"operator"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ScopeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Scope `json:"items,omitempty"`
}

// ScopeDashboards is a list of dashboards that belong to a given scope
// The kubernetes name for this object matches the scope name
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ScopeDashboardBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ScopeDashboardBindingSpec `json:"spec"`
}

type ScopeDashboardBindingSpec struct {
	// List of dashboard names (eg, grafana uids)
	// +listType=set
	Dashboards []string `json:"dashboards"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ScopeDashboardBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ScopeDashboardBinding `json:"items,omitempty"`
}
