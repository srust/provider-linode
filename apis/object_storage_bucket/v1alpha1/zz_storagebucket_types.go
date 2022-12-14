/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CertObservation struct {
}

type CertParameters struct {

	// The Base64 encoded and PEM formatted SSL certificate.
	// The Base64 encoded and PEM formatted SSL certificate.
	// +kubebuilder:validation:Required
	CertificateSecretRef v1.SecretKeySelector `json:"certificateSecretRef" tf:"-"`

	// The private key associated with the TLS/SSL certificate.
	// The private key associated with the TLS/SSL certificate.
	// +kubebuilder:validation:Required
	PrivateKeySecretRef v1.SecretKeySelector `json:"privateKeySecretRef" tf:"-"`
}

type ExpirationObservation struct {
}

type ExpirationParameters struct {

	// Specifies the date after which you want the corresponding action to take effect.
	// Specifies the date after which you want the corresponding action to take effect.
	// +kubebuilder:validation:Optional
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// Specifies the number of days after object creation when the specific rule action takes effect.
	// Specifies the number of days after object creation when the specific rule action takes effect.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// On a versioned bucket (versioning-enabled or versioning-suspended bucket), you can add this element in the lifecycle configuration to direct Linode Object Storage to delete expired object delete markers. This cannot be specified with Days or Date in a Lifecycle Expiration Policy.
	// Directs Linode Object Storage to remove expired deleted markers.
	// +kubebuilder:validation:Optional
	ExpiredObjectDeleteMarker *bool `json:"expiredObjectDeleteMarker,omitempty" tf:"expired_object_delete_marker,omitempty"`
}

type LifecycleRuleObservation struct {
}

type LifecycleRuleParameters struct {

	// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.
	// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.
	// +kubebuilder:validation:Optional
	AbortIncompleteMultipartUploadDays *float64 `json:"abortIncompleteMultipartUploadDays,omitempty" tf:"abort_incomplete_multipart_upload_days,omitempty"`

	// Specifies whether the lifecycle rule is active.
	// Specifies whether the lifecycle rule is active.
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// Specifies a period in the object's expire.
	// +kubebuilder:validation:Optional
	Expiration []ExpirationParameters `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// The unique identifier for the rule.
	// The unique identifier for the rule.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies when non-current object versions expire.
	// +kubebuilder:validation:Optional
	NoncurrentVersionExpiration []NoncurrentVersionExpirationParameters `json:"noncurrentVersionExpiration,omitempty" tf:"noncurrent_version_expiration,omitempty"`

	// The object key prefix identifying one or more objects to which the rule applies.
	// The object key prefix identifying one or more objects to which the rule applies.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type NoncurrentVersionExpirationObservation struct {
}

type NoncurrentVersionExpirationParameters struct {

	// Specifies the number of days after object creation when the specific rule action takes effect.
	// Specifies the number of days non-current object versions expire.
	// +kubebuilder:validation:Required
	Days *float64 `json:"days" tf:"days,omitempty"`
}

type StorageBucketObservation struct {

	// The hostname where this bucket can be accessed. This hostname can be accessed through a browser if the bucket is made public.
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// The unique identifier for the rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type StorageBucketParameters struct {

	// The Access Control Level of the bucket using a canned ACL string. See all ACL strings in the Linode API v4 documentation.
	// The Access Control Level of the bucket using a canned ACL string.
	// +kubebuilder:validation:Optional
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// The access key to authenticate with.
	// The S3 access key to use for this resource. (Required for lifecycle_rule and versioning)
	// +kubebuilder:validation:Optional
	AccessKey *string `json:"accessKey,omitempty" tf:"access_key,omitempty"`

	// The cert used by this Object Storage Bucket.
	// +kubebuilder:validation:Optional
	Cert []CertParameters `json:"cert,omitempty" tf:"cert,omitempty"`

	// The cluster of the Linode Object Storage Bucket.
	// The cluster of the Linode Object Storage Bucket.
	// +kubebuilder:validation:Required
	Cluster *string `json:"cluster" tf:"cluster,omitempty"`

	// If true, the bucket will have CORS enabled for all origins.
	// If true, the bucket will be created with CORS enabled for all origins.
	// +kubebuilder:validation:Optional
	CorsEnabled *bool `json:"corsEnabled,omitempty" tf:"cors_enabled,omitempty"`

	// The label of the Linode Object Storage Bucket.
	// The label of the Linode Object Storage Bucket.
	// +kubebuilder:validation:Required
	Label *string `json:"label" tf:"label,omitempty"`

	// Lifecycle rules to be applied to the bucket.
	// +kubebuilder:validation:Optional
	LifecycleRule []LifecycleRuleParameters `json:"lifecycleRule,omitempty" tf:"lifecycle_rule,omitempty"`

	// The secret key to authenticate with.
	// The S3 secret key to use for this resource. (Required for lifecycle_rule and versioning)
	// +kubebuilder:validation:Optional
	SecretKey *string `json:"secretKey,omitempty" tf:"secret_key,omitempty"`

	// Whether to enable versioning. Once you version-enable a bucket, it can never return to an unversioned state. You can, however, suspend versioning on that bucket. (Requires access_key and secret_key)
	// Whether to enable versioning.
	// +kubebuilder:validation:Optional
	Versioning *bool `json:"versioning,omitempty" tf:"versioning,omitempty"`
}

// StorageBucketSpec defines the desired state of StorageBucket
type StorageBucketSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StorageBucketParameters `json:"forProvider"`
}

// StorageBucketStatus defines the observed state of StorageBucket.
type StorageBucketStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StorageBucketObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StorageBucket is the Schema for the StorageBuckets API. Manages a Linode Object Storage Bucket.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,linode}
type StorageBucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageBucketSpec   `json:"spec"`
	Status            StorageBucketStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StorageBucketList contains a list of StorageBuckets
type StorageBucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageBucket `json:"items"`
}

// Repository type metadata.
var (
	StorageBucket_Kind             = "StorageBucket"
	StorageBucket_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StorageBucket_Kind}.String()
	StorageBucket_KindAPIVersion   = StorageBucket_Kind + "." + CRDGroupVersion.String()
	StorageBucket_GroupVersionKind = CRDGroupVersion.WithKind(StorageBucket_Kind)
)

func init() {
	SchemeBuilder.Register(&StorageBucket{}, &StorageBucketList{})
}
