package awss3express


// Specifies default encryption for a bucket using server-side encryption with Amazon S3 managed keys (SSE-S3) or AWS KMS keys (SSE-KMS).
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
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-bucketencryption.html
//
type CfnDirectoryBucket_BucketEncryptionProperty struct {
	// Specifies the default server-side-encryption configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-bucketencryption.html#cfn-s3express-directorybucket-bucketencryption-serversideencryptionconfiguration
	//
	ServerSideEncryptionConfiguration interface{} `field:"required" json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
}

