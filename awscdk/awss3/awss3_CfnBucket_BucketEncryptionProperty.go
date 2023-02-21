package awss3


// Specifies default encryption for a bucket using server-side encryption with Amazon S3-managed keys (SSE-S3) or AWS KMS-managed keys (SSE-KMS) bucket.
//
// For information about the Amazon S3 default encryption feature, see [Amazon S3 Default Encryption for S3 Buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bucketEncryptionProperty := &bucketEncryptionProperty{
//   	serverSideEncryptionConfiguration: []interface{}{
//   		&serverSideEncryptionRuleProperty{
//   			bucketKeyEnabled: jsii.Boolean(false),
//   			serverSideEncryptionByDefault: &serverSideEncryptionByDefaultProperty{
//   				sseAlgorithm: jsii.String("sseAlgorithm"),
//
//   				// the properties below are optional
//   				kmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   			},
//   		},
//   	},
//   }
//
type CfnBucket_BucketEncryptionProperty struct {
	// Specifies the default server-side-encryption configuration.
	ServerSideEncryptionConfiguration interface{} `field:"required" json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
}

