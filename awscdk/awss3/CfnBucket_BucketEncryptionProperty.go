package awss3


// Specifies default encryption for a bucket using server-side encryption with Amazon S3-managed keys (SSE-S3), AWS KMS-managed keys (SSE-KMS), or dual-layer server-side encryption with KMS-managed keys (DSSE-KMS).
//
// For information about the Amazon S3 default encryption feature, see [Amazon S3 Default Encryption for S3 Buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bucketEncryptionProperty := &BucketEncryptionProperty{
//   	ServerSideEncryptionConfiguration: []interface{}{
//   		&ServerSideEncryptionRuleProperty{
//   			BucketKeyEnabled: jsii.Boolean(false),
//   			ServerSideEncryptionByDefault: &ServerSideEncryptionByDefaultProperty{
//   				SseAlgorithm: jsii.String("sseAlgorithm"),
//
//   				// the properties below are optional
//   				KmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-bucketencryption.html
//
type CfnBucket_BucketEncryptionProperty struct {
	// Specifies the default server-side-encryption configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-bucketencryption.html#cfn-s3-bucket-bucketencryption-serversideencryptionconfiguration
	//
	ServerSideEncryptionConfiguration interface{} `field:"required" json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
}

