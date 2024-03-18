// Copyright 2023 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"maps"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

//+kubebuilder:object:root=true

// NodeConfig is the Schema for the workerconfigs API.
type NodeConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	//+optional
	Spec NodeConfigSpec `json:"spec,omitempty"`
}

// NodeConfigSpec defines the desired state of NodeConfig.
// Place any configuration that can be applied to individual Nodes here.
// Otherwise, it should go into the ClusterConfigSpec.
type NodeConfigSpec struct {
	// +optional
	AWS *AWSNodeSpec `json:"aws,omitempty"`
	// +optional
	Docker *DockerNodeSpec `json:"docker,omitempty"`

	GenericNodeSpec `json:",inline"`
}

func (s NodeConfigSpec) VariableSchema() clusterv1.VariableSchema {
	nodeConfigProps := GenericNodeSpec{}.VariableSchema()

	switch {
	case s.AWS != nil:
		maps.Copy(
			nodeConfigProps.OpenAPIV3Schema.Properties,
			map[string]clusterv1.JSONSchemaProps{
				AWSVariableName: AWSNodeSpec{}.VariableSchema().OpenAPIV3Schema,
			},
		)
	case s.Docker != nil:
		maps.Copy(
			nodeConfigProps.OpenAPIV3Schema.Properties,
			map[string]clusterv1.JSONSchemaProps{
				"docker": DockerNodeSpec{}.VariableSchema().OpenAPIV3Schema,
			},
		)
	}

	return nodeConfigProps
}

type GenericNodeSpec struct {
	// +optional
	Users Users `json:"users,omitempty"`
}

func (GenericNodeSpec) VariableSchema() clusterv1.VariableSchema {
	return clusterv1.VariableSchema{
		OpenAPIV3Schema: clusterv1.JSONSchemaProps{
			Description: "Node configuration",
			Type:        "object",
			Properties: map[string]clusterv1.JSONSchemaProps{
				"users": Users{}.VariableSchema().OpenAPIV3Schema,
			},
		},
	}
}

type Users []User

func (Users) VariableSchema() clusterv1.VariableSchema {
	return clusterv1.VariableSchema{
		OpenAPIV3Schema: clusterv1.JSONSchemaProps{
			Description: "TBD",
			Type:        "array",
			Items:       ptr.To(User{}.VariableSchema().OpenAPIV3Schema),
		},
	}
}

// User defines the input for a generated user in cloud-init.
type User struct {
	// Name specifies the user name
	Name string `json:"name"`

	// LockPassword specifies if password login should be disabled
	// +optional
	LockPassword *bool `json:"lockPassword,omitempty"`

	// Passwd specifies a hashed password for the user
	// +optional
	Passwd *string `json:"passwd,omitempty"`

	// SSHAuthorizedKeys specifies a list of ssh authorized keys for the user
	// +optional
	SSHAuthorizedKeys []string `json:"sshAuthorizedKeys,omitempty"`

	// Sudo specifies a sudo role for the user
	// +optional
	Sudo *string `json:"sudo,omitempty"`
}

func (User) VariableSchema() clusterv1.VariableSchema {
	return clusterv1.VariableSchema{
		OpenAPIV3Schema: clusterv1.JSONSchemaProps{
			Description: "User",
			Type:        "TBD",
			Properties: map[string]clusterv1.JSONSchemaProps{
				"name": {
					Description: "TBD",
					Type:        "string",
				},
				"password": {
					Description: "TBD",
					Type:        "string",
				},
				"sshAuthorizedKeys": {
					Description: "TBD",
					Type:        "array",
					Items: &clusterv1.JSONSchemaProps{
						// No description, because the one for the parent array is enough.
						Type: "string",
					},
				},
				"sudo": {
					Description: "TBD",
					Type:        "string",
				},
			},
		},
	}
}

func init() {
	SchemeBuilder.Register(&NodeConfig{})
}
