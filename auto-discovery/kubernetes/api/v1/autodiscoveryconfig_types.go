// SPDX-FileCopyrightText: 2021 iteratec GmbH
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	cfg "sigs.k8s.io/controller-runtime/pkg/config/v1alpha1"
)

//+kubebuilder:object:root=true
type AutoDiscoveryConfig struct {
	metav1.TypeMeta `json:",inline"`

	// ControllerManagerConfigurationSpec returns the contfigurations for controllers
	cfg.ControllerManagerConfigurationSpec `json:",inline"`

	Cluster                    ClusterConfig              `json:"cluster"`
	ResourceInclusion          ResourceInclusionConfig    `json:"resourceInclusion"`
	ServiceAutoDiscoveryConfig ServiceAutoDiscoveryConfig `json:"serviceAutoDiscovery"`
}

type ServiceAutoDiscoveryConfig struct {
	PassiveReconcileInterval metav1.Duration `json:"passiveReconcileInterval"`
	ScanConfig               ScanConfig      `json:"scanConfig"`
}

type ClusterConfig struct {
	Name string `json:"name"`
}

// ResourceInclusionMode Modes which can be used to trigger "ResourceInclusion" behavior which dictaes which resources are considered by the auto-discovery.
type ResourceInclusionMode string

const (
	EnabledPerNamespace ResourceInclusionMode = "enabled-per-namespace"
	EnabledPerResource  ResourceInclusionMode = "enabled-per-resource"
	All                 ResourceInclusionMode = "all"
)

type ResourceInclusionConfig struct {
	Mode ResourceInclusionMode `json:"mode"`
}

type ScanConfig struct {
	RepeatInterval metav1.Duration   `json:"repeatInterval"`
	Annotations    map[string]string `json:"annotations"`
	Labels         map[string]string `json:"labels"`
	ScanType       string            `json:"scanType"`
	Parameters     []string          `json:"parameters"`
}

func init() {
	SchemeBuilder.Register(&AutoDiscoveryConfig{})
}
