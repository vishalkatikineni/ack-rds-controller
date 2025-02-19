// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DBClusterSnapshotSpec defines the desired state of DBClusterSnapshot.
//
// # Contains the details for an Amazon RDS DB cluster snapshot
//
// This data type is used as a response element in the DescribeDBClusterSnapshots
// action.
type DBClusterSnapshotSpec struct {

	// The identifier of the DB cluster to create a snapshot for. This parameter
	// isn't case-sensitive.
	//
	// Constraints:
	//
	//    * Must match the identifier of an existing DBCluster.
	//
	// Example: my-cluster1

	DBClusterIdentifier *string `json:"dbClusterIdentifier,omitempty"`

	DBClusterIdentifierRef *ackv1alpha1.AWSResourceReferenceWrapper `json:"dbClusterIdentifierRef,omitempty"`
	// The identifier of the DB cluster snapshot. This parameter is stored as a
	// lowercase string.
	//
	// Constraints:
	//
	//    * Must contain from 1 to 63 letters, numbers, or hyphens.
	//
	//    * First character must be a letter.
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens.
	//
	// Example: my-cluster1-snapshot1

	// +kubebuilder:validation:Required

	DBClusterSnapshotIdentifier *string `json:"dbClusterSnapshotIdentifier"`
	// The tags to be assigned to the DB cluster snapshot.

	Tags []*Tag `json:"tags,omitempty"`
}

// DBClusterSnapshotStatus defines the observed state of DBClusterSnapshot
type DBClusterSnapshotStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRs managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The allocated storage size of the DB cluster snapshot in gibibytes (GiB).
	// +kubebuilder:validation:Optional
	AllocatedStorage *int64 `json:"allocatedStorage,omitempty"`
	// The list of Availability Zones (AZs) where instances in the DB cluster snapshot
	// can be restored.
	// +kubebuilder:validation:Optional
	AvailabilityZones []*string `json:"availabilityZones,omitempty"`
	// The time when the DB cluster was created, in Universal Coordinated Time (UTC).
	// +kubebuilder:validation:Optional
	ClusterCreateTime *metav1.Time `json:"clusterCreateTime,omitempty"`
	// Reserved for future use.
	// +kubebuilder:validation:Optional
	DBSystemID *string `json:"dbSystemID,omitempty"`
	// The name of the database engine for this DB cluster snapshot.
	// +kubebuilder:validation:Optional
	Engine *string `json:"engine,omitempty"`
	// The engine mode of the database engine for this DB cluster snapshot.
	// +kubebuilder:validation:Optional
	EngineMode *string `json:"engineMode,omitempty"`
	// The version of the database engine for this DB cluster snapshot.
	// +kubebuilder:validation:Optional
	EngineVersion *string `json:"engineVersion,omitempty"`
	// Indicates whether mapping of Amazon Web Services Identity and Access Management
	// (IAM) accounts to database accounts is enabled.
	// +kubebuilder:validation:Optional
	IAMDatabaseAuthenticationEnabled *bool `json:"iamDatabaseAuthenticationEnabled,omitempty"`
	// If StorageEncrypted is true, the Amazon Web Services KMS key identifier for
	// the encrypted DB cluster snapshot.
	//
	// The Amazon Web Services KMS key identifier is the key ARN, key ID, alias
	// ARN, or alias name for the KMS key.
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyID,omitempty"`
	// The license model information for this DB cluster snapshot.
	// +kubebuilder:validation:Optional
	LicenseModel *string `json:"licenseModel,omitempty"`
	// The master username for this DB cluster snapshot.
	// +kubebuilder:validation:Optional
	MasterUsername *string `json:"masterUsername,omitempty"`
	// The percentage of the estimated data that has been transferred.
	// +kubebuilder:validation:Optional
	PercentProgress *int64 `json:"percentProgress,omitempty"`
	// The port that the DB cluster was listening on at the time of the snapshot.
	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty"`
	// The time when the snapshot was taken, in Universal Coordinated Time (UTC).
	// +kubebuilder:validation:Optional
	SnapshotCreateTime *metav1.Time `json:"snapshotCreateTime,omitempty"`
	// The type of the DB cluster snapshot.
	// +kubebuilder:validation:Optional
	SnapshotType *string `json:"snapshotType,omitempty"`
	// If the DB cluster snapshot was copied from a source DB cluster snapshot,
	// the Amazon Resource Name (ARN) for the source DB cluster snapshot, otherwise,
	// a null value.
	// +kubebuilder:validation:Optional
	SourceDBClusterSnapshotARN *string `json:"sourceDBClusterSnapshotARN,omitempty"`
	// The status of this DB cluster snapshot. Valid statuses are the following:
	//
	//    * available
	//
	//    * copying
	//
	//    * creating
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty"`
	// Indicates whether the DB cluster snapshot is encrypted.
	// +kubebuilder:validation:Optional
	StorageEncrypted *bool `json:"storageEncrypted,omitempty"`
	// +kubebuilder:validation:Optional
	TagList []*Tag `json:"tagList,omitempty"`
	// The VPC ID associated with the DB cluster snapshot.
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcID,omitempty"`
}

// DBClusterSnapshot is the Schema for the DBClusterSnapshots API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type DBClusterSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DBClusterSnapshotSpec   `json:"spec,omitempty"`
	Status            DBClusterSnapshotStatus `json:"status,omitempty"`
}

// DBClusterSnapshotList contains a list of DBClusterSnapshot
// +kubebuilder:object:root=true
type DBClusterSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DBClusterSnapshot `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DBClusterSnapshot{}, &DBClusterSnapshotList{})
}
