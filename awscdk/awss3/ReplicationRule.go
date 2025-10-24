package awss3

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Specifies which Amazon S3 objects to replicate and where to store the replicas.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket Bucket
//   var key Key
//   var replicationTimeValue ReplicationTimeValue
//   var storageClass StorageClass
//
//   replicationRule := &ReplicationRule{
//   	Destination: bucket,
//
//   	// the properties below are optional
//   	AccessControlTransition: jsii.Boolean(false),
//   	DeleteMarkerReplication: jsii.Boolean(false),
//   	Filter: &Filter{
//   		Prefix: jsii.String("prefix"),
//   		Tags: []Tag{
//   			&Tag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	Id: jsii.String("id"),
//   	KmsKey: key,
//   	Metrics: replicationTimeValue,
//   	Priority: jsii.Number(123),
//   	ReplicaModifications: jsii.Boolean(false),
//   	ReplicationTimeControl: replicationTimeValue,
//   	SseKmsEncryptedObjects: jsii.Boolean(false),
//   	StorageClass: storageClass,
//   }
//
type ReplicationRule struct {
	// The destination bucket for the replicated objects.
	//
	// The destination can be either in the same AWS account or a cross account.
	//
	// If you want to configure cross-account replication,
	// the destination bucket must have a policy that allows the source bucket to replicate objects to it.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/userguide/replication-walkthrough-2.html
	//
	Destination IBucket `field:"required" json:"destination" yaml:"destination"`
	// Whether to want to change replica ownership to the AWS account that owns the destination bucket.
	//
	// This can only be specified if the source bucket and the destination bucket are not in the same AWS account.
	// Default: - The replicas are owned by same AWS account that owns the source object.
	//
	AccessControlTransition *bool `field:"optional" json:"accessControlTransition" yaml:"accessControlTransition"`
	// Specifies whether Amazon S3 replicates delete markers.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/userguide/delete-marker-replication.html
	//
	// Default: - delete markers in source bucket is not replicated to destination bucket.
	//
	DeleteMarkerReplication *bool `field:"optional" json:"deleteMarkerReplication" yaml:"deleteMarkerReplication"`
	// A filter that identifies the subset of objects to which the replication rule applies.
	// Default: - applies to all objects.
	//
	Filter *Filter `field:"optional" json:"filter" yaml:"filter"`
	// A unique identifier for the rule.
	//
	// The maximum value is 255 characters.
	// Default: - auto generated random ID.
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The customer managed AWS KMS key stored in AWS Key Management Service (KMS) for the destination bucket.
	//
	// Amazon S3 uses this key to encrypt replica objects.
	//
	// Amazon S3 only supports symmetric encryption KMS keys.
	// See: https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html
	//
	// Default: - Amazon S3 uses the AWS managed KMS key for encryption.
	//
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// A container specifying replication metrics-related settings enabling replication metrics and events.
	//
	// When a value is set, metrics will be output to indicate whether the replication took longer than the specified time.
	// Default: - Replication metrics are not enabled.
	//
	Metrics ReplicationTimeValue `field:"optional" json:"metrics" yaml:"metrics"`
	// The priority indicates which rule has precedence whenever two or more replication rules conflict.
	//
	// Amazon S3 will attempt to replicate objects according to all replication rules.
	// However, if there are two or more rules with the same destination bucket,
	// then objects will be replicated according to the rule with the highest priority.
	//
	// The higher the number, the higher the priority.
	//
	// It is essential to specify priority explicitly when the replication configuration has multiple rules.
	// Default: 0.
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Specifies whether Amazon S3 replicates modifications on replicas.
	// Default: false.
	//
	ReplicaModifications *bool `field:"optional" json:"replicaModifications" yaml:"replicaModifications"`
	// Specifying S3 Replication Time Control (S3 RTC), including whether S3 RTC is enabled and the time when all objects and operations on objects must be replicated.
	// Default: - S3 Replication Time Control is not enabled.
	//
	ReplicationTimeControl ReplicationTimeValue `field:"optional" json:"replicationTimeControl" yaml:"replicationTimeControl"`
	// Specifies whether Amazon S3 replicates objects created with server-side encryption using an AWS KMS key stored in AWS Key Management Service.
	// Default: false.
	//
	SseKmsEncryptedObjects *bool `field:"optional" json:"sseKmsEncryptedObjects" yaml:"sseKmsEncryptedObjects"`
	// The storage class to use when replicating objects, such as S3 Standard or reduced redundancy.
	// Default: - The storage class of the source object.
	//
	StorageClass StorageClass `field:"optional" json:"storageClass" yaml:"storageClass"`
}

