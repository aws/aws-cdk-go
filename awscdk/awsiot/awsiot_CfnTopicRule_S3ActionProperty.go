package awsiot


// Describes an action to write data to an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ActionProperty := &s3ActionProperty{
//   	bucketName: jsii.String("bucketName"),
//   	key: jsii.String("key"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	cannedAcl: jsii.String("cannedAcl"),
//   }
//
type CfnTopicRule_S3ActionProperty struct {
	// The Amazon S3 bucket.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The object key.
	//
	// For more information, see [Actions, resources, and condition keys for Amazon S3](https://docs.aws.amazon.com/AmazonS3/latest/dev/list_amazons3.html) .
	Key *string `field:"required" json:"key" yaml:"key"`
	// The ARN of the IAM role that grants access.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The Amazon S3 canned ACL that controls access to the object identified by the object key.
	//
	// For more information, see [S3 canned ACLs](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) .
	CannedAcl *string `field:"optional" json:"cannedAcl" yaml:"cannedAcl"`
}

