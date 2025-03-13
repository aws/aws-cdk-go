package awss3


// A container for information about the replication destination and its configurations including enabling the S3 Replication Time Control (S3 RTC).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationDestinationProperty := &ReplicationDestinationProperty{
//   	Bucket: jsii.String("bucket"),
//
//   	// the properties below are optional
//   	AccessControlTranslation: &AccessControlTranslationProperty{
//   		Owner: jsii.String("owner"),
//   	},
//   	Account: jsii.String("account"),
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		ReplicaKmsKeyId: jsii.String("replicaKmsKeyId"),
//   	},
//   	Metrics: &MetricsProperty{
//   		Status: jsii.String("status"),
//
//   		// the properties below are optional
//   		EventThreshold: &ReplicationTimeValueProperty{
//   			Minutes: jsii.Number(123),
//   		},
//   	},
//   	ReplicationTime: &ReplicationTimeProperty{
//   		Status: jsii.String("status"),
//   		Time: &ReplicationTimeValueProperty{
//   			Minutes: jsii.Number(123),
//   		},
//   	},
//   	StorageClass: jsii.String("storageClass"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationdestination.html
//
type CfnBucket_ReplicationDestinationProperty struct {
	// The Amazon Resource Name (ARN) of the bucket where you want Amazon S3 to store the results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationdestination.html#cfn-s3-bucket-replicationdestination-bucket
	//
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Specify this only in a cross-account scenario (where source and destination bucket owners are not the same), and you want to change replica ownership to the AWS account that owns the destination bucket.
	//
	// If this is not specified in the replication configuration, the replicas are owned by same AWS account that owns the source object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationdestination.html#cfn-s3-bucket-replicationdestination-accesscontroltranslation
	//
	AccessControlTranslation interface{} `field:"optional" json:"accessControlTranslation" yaml:"accessControlTranslation"`
	// Destination bucket owner account ID.
	//
	// In a cross-account scenario, if you direct Amazon S3 to change replica ownership to the AWS account that owns the destination bucket by specifying the `AccessControlTranslation` property, this is the account ID of the destination bucket owner. For more information, see [Cross-Region Replication Additional Configuration: Change Replica Owner](https://docs.aws.amazon.com/AmazonS3/latest/dev/crr-change-owner.html) in the *Amazon S3 User Guide* .
	//
	// If you specify the `AccessControlTranslation` property, the `Account` property is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationdestination.html#cfn-s3-bucket-replicationdestination-account
	//
	Account *string `field:"optional" json:"account" yaml:"account"`
	// Specifies encryption-related information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationdestination.html#cfn-s3-bucket-replicationdestination-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// A container specifying replication metrics-related settings enabling replication metrics and events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationdestination.html#cfn-s3-bucket-replicationdestination-metrics
	//
	Metrics interface{} `field:"optional" json:"metrics" yaml:"metrics"`
	// A container specifying S3 Replication Time Control (S3 RTC), including whether S3 RTC is enabled and the time when all objects and operations on objects must be replicated.
	//
	// Must be specified together with a `Metrics` block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationdestination.html#cfn-s3-bucket-replicationdestination-replicationtime
	//
	ReplicationTime interface{} `field:"optional" json:"replicationTime" yaml:"replicationTime"`
	// The storage class to use when replicating objects, such as S3 Standard or reduced redundancy.
	//
	// By default, Amazon S3 uses the storage class of the source object to create the object replica.
	//
	// For valid values, see the `StorageClass` element of the [PUT Bucket replication](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTreplication.html) action in the *Amazon S3 API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationdestination.html#cfn-s3-bucket-replicationdestination-storageclass
	//
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
}

