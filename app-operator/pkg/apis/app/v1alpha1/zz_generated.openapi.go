// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/example-inc/app-operator/pkg/apis/app/v1alpha1.AppService":       schema_pkg_apis_app_v1alpha1_AppService(ref),
		"github.com/example-inc/app-operator/pkg/apis/app/v1alpha1.AppServiceSpec":   schema_pkg_apis_app_v1alpha1_AppServiceSpec(ref),
		"github.com/example-inc/app-operator/pkg/apis/app/v1alpha1.AppServiceStatus": schema_pkg_apis_app_v1alpha1_AppServiceStatus(ref),
	}
}

func schema_pkg_apis_app_v1alpha1_AppService(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AppService is the Schema for the appservices API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/example-inc/app-operator/pkg/apis/app/v1alpha1.AppServiceSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/example-inc/app-operator/pkg/apis/app/v1alpha1.AppServiceStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/example-inc/app-operator/pkg/apis/app/v1alpha1.AppServiceSpec", "github.com/example-inc/app-operator/pkg/apis/app/v1alpha1.AppServiceStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_app_v1alpha1_AppServiceSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AppServiceSpec defines the desired state of AppService",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_app_v1alpha1_AppServiceStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AppServiceStatus defines the observed state of AppService",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}
