/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/linode/provider-linode/apis/database/v1alpha1"
	v1alpha1domain "github.com/linode/provider-linode/apis/domain/v1alpha1"
	v1alpha1firewall "github.com/linode/provider-linode/apis/firewall/v1alpha1"
	v1alpha1instance "github.com/linode/provider-linode/apis/instance/v1alpha1"
	v1alpha1linode "github.com/linode/provider-linode/apis/linode/v1alpha1"
	v1alpha1lke "github.com/linode/provider-linode/apis/lke/v1alpha1"
	v1alpha1nodebalancer "github.com/linode/provider-linode/apis/nodebalancer/v1alpha1"
	v1alpha1object "github.com/linode/provider-linode/apis/object/v1alpha1"
	v1alpha1apis "github.com/linode/provider-linode/apis/v1alpha1"
	v1beta1 "github.com/linode/provider-linode/apis/v1beta1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1domain.SchemeBuilder.AddToScheme,
		v1alpha1firewall.SchemeBuilder.AddToScheme,
		v1alpha1instance.SchemeBuilder.AddToScheme,
		v1alpha1linode.SchemeBuilder.AddToScheme,
		v1alpha1lke.SchemeBuilder.AddToScheme,
		v1alpha1nodebalancer.SchemeBuilder.AddToScheme,
		v1alpha1object.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
