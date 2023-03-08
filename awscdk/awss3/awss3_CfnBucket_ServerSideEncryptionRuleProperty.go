package awss3


// Specifies the default server-side encryption configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverSideEncryptionRuleProperty := &serverSideEncryptionRuleProperty{
//   	bucketKeyEnabled: jsii.Boolean(false),
//   	serverSideEncryptionByDefault: &serverSideEncryptionByDefaultProperty{
//   		sseAlgorithm: jsii.String("sseAlgorithm"),
//
//   		// the properties below are optional
//   		kmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   	},
//   }
//
type CfnBucket_ServerSideEncryptionRuleProperty struct {
	// Specifies whether Amazon S3 should use an S3 Bucket Key with server-side encryption using KMS (SSE-KMS) for new objects in the bucket.
	//
	// Existing objects are not affected. Setting the `BucketKeyEnabled` element to `true` causes Amazon S3 to use an S3 Bucket Key. By default, S3 Bucket Key is not enabled.
	//
	// For more information, see [Amazon S3 Bucket Keys](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-key.html) in the *Amazon S3 User Guide* .
	BucketKeyEnabled interface{} `field:"optional" json:"bucketKeyEnabled" yaml:"bucketKeyEnabled"`
	// Specifies the default server-side encryption to apply to new objects in the bucket.
	//
	// If a PUT Object request doesn't specify any server-side encryption, this default encryption will be applied.
	ServerSideEncryptionByDefault interface{} `field:"optional" json:"serverSideEncryptionByDefault" yaml:"serverSideEncryptionByDefault"`
}

