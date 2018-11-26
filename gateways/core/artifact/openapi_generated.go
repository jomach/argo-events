// +build !ignore_autogenerated

/*
Copyright 2018 BlackRock, Inc.

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
// Code generated by main. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package artifact

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/argoproj/argo-events/gateways/core/artifact.S3Artifact":    schema_argo_events_gateways_core_artifact_S3Artifact(ref),
		"github.com/argoproj/argo-events/gateways/core/artifact.S3EventConfig": schema_argo_events_gateways_core_artifact_S3EventConfig(ref),
		"github.com/argoproj/argo-events/gateways/core/artifact.S3Filter":      schema_argo_events_gateways_core_artifact_S3Filter(ref),
	}
}

func schema_argo_events_gateways_core_artifact_S3Artifact(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "S3Artifact contains information about an artifact in S3",
				Properties: map[string]spec.Schema{
					"s3EventConfig": {
						SchemaProps: spec.SchemaProps{
							Description: "S3EventConfig contains configuration for bucket notification",
							Ref:         ref("github.com/argoproj/argo-events/gateways/core/artifact.S3EventConfig"),
						},
					},
					"insecure": {
						SchemaProps: spec.SchemaProps{
							Description: "Mode of operation for s3 client",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
				Required: []string{"s3EventConfig"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/gateways/core/artifact.S3EventConfig"},
	}
}

func schema_argo_events_gateways_core_artifact_S3EventConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "S3EventConfig contains configuration for bucket notification",
				Properties: map[string]spec.Schema{
					"endpoint": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"bucket": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"region": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"event": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"filter": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj/argo-events/gateways/core/artifact.S3Filter"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/gateways/core/artifact.S3Filter"},
	}
}

func schema_argo_events_gateways_core_artifact_S3Filter(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "S3Filter represents filters to apply to bucket nofifications for specifying constraints on objects",
				Properties: map[string]spec.Schema{
					"prefix": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"suffix": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"prefix", "suffix"},
			},
		},
		Dependencies: []string{},
	}
}