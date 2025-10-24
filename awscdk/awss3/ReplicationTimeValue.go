package awss3

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The replication time value used for S3 Replication Time Control (S3 RTC).
//
// Example:
//   var destinationBucket1 IBucket
//   var destinationBucket2 IBucket
//   var replicationRole IRole
//   var encryptionKey IKey
//   var destinationEncryptionKey IKey
//
//
//   sourceBucket := s3.NewBucket(this, jsii.String("SourceBucket"), &BucketProps{
//   	// Versioning must be enabled on both the source and destination bucket
//   	Versioned: jsii.Boolean(true),
//   	// Optional. Specify the KMS key to use for encrypts objects in the source bucket.
//   	EncryptionKey: EncryptionKey,
//   	// Optional. If not specified, a new role will be created.
//   	ReplicationRole: ReplicationRole,
//   	ReplicationRules: []ReplicationRule{
//   		&ReplicationRule{
//   			// The destination bucket for the replication rule.
//   			Destination: destinationBucket1,
//   			// The priority of the rule.
//   			// Amazon S3 will attempt to replicate objects according to all replication rules.
//   			// However, if there are two or more rules with the same destination bucket, then objects will be replicated according to the rule with the highest priority.
//   			// The higher the number, the higher the priority.
//   			// It is essential to specify priority explicitly when the replication configuration has multiple rules.
//   			Priority: jsii.Number(1),
//   		},
//   		&ReplicationRule{
//   			Destination: destinationBucket2,
//   			Priority: jsii.Number(2),
//   			// Whether to specify S3 Replication Time Control (S3 RTC).
//   			// S3 RTC replicates most objects that you upload to Amazon S3 in seconds,
//   			// and 99.99 percent of those objects within specified time.
//   			ReplicationTimeControl: s3.ReplicationTimeValue_FIFTEEN_MINUTES(),
//   			// Whether to enable replication metrics about S3 RTC.
//   			// If set, metrics will be output to indicate whether replication by S3 RTC took longer than the configured time.
//   			Metrics: s3.ReplicationTimeValue_FIFTEEN_MINUTES(),
//   			// The kms key to use for the destination bucket.
//   			KmsKey: destinationEncryptionKey,
//   			// The storage class to use for the destination bucket.
//   			StorageClass: s3.StorageClass_INFREQUENT_ACCESS(),
//   			// Whether to replicate objects with SSE-KMS encryption.
//   			SseKmsEncryptedObjects: jsii.Boolean(false),
//   			// Whether to replicate modifications on replicas.
//   			ReplicaModifications: jsii.Boolean(true),
//   			// Whether to replicate delete markers.
//   			// This property cannot be enabled if the replication rule has a tag filter.
//   			DeleteMarkerReplication: jsii.Boolean(false),
//   			// The ID of the rule.
//   			Id: jsii.String("full-settings-rule"),
//   			// The object filter for the rule.
//   			Filter: &Filter{
//   				// The prefix filter for the rule.
//   				Prefix: jsii.String("prefix"),
//   				// The tag filter for the rule.
//   				Tags: []Tag{
//   					&Tag{
//   						Key: jsii.String("tagKey"),
//   						Value: jsii.String("tagValue"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   })
//
//   // Grant permissions to the replication role.
//   // This method is not required if you choose to use an auto-generated replication role or manually grant permissions.
//   sourceBucket.GrantReplicationPermission(replicationRole, &GrantReplicationPermissionProps{
//   	// Optional. Specify the KMS key to use for decrypting objects in the source bucket.
//   	SourceDecryptionKey: encryptionKey,
//   	Destinations: []GrantReplicationPermissionDestinationProps{
//   		&GrantReplicationPermissionDestinationProps{
//   			Bucket: destinationBucket1,
//   		},
//   		&GrantReplicationPermissionDestinationProps{
//   			Bucket: destinationBucket2,
//   			EncryptionKey: destinationEncryptionKey,
//   		},
//   	},
//   })
//
type ReplicationTimeValue interface {
	// the time in minutes.
	Minutes() *float64
}

// The jsii proxy struct for ReplicationTimeValue
type jsiiProxy_ReplicationTimeValue struct {
	_ byte // padding
}

func (j *jsiiProxy_ReplicationTimeValue) Minutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minutes",
		&returns,
	)
	return returns
}


func ReplicationTimeValue_FIFTEEN_MINUTES() ReplicationTimeValue {
	_init_.Initialize()
	var returns ReplicationTimeValue
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_s3.ReplicationTimeValue",
		"FIFTEEN_MINUTES",
		&returns,
	)
	return returns
}

