package awsiot


// Describes an action to write data to an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ActionProperty := &S3ActionProperty{
//   	BucketName: jsii.String("bucketName"),
//   	Key: jsii.String("key"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	CannedAcl: jsii.String("cannedAcl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-s3action.html
//
type CfnTopicRule_S3ActionProperty struct {
	// The Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-s3action.html#cfn-iot-topicrule-s3action-bucketname
	//
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The object key.
	//
	// For more information, see [Actions, resources, and condition keys for Amazon S3](https://docs.aws.amazon.com/AmazonS3/latest/dev/list_amazons3.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-s3action.html#cfn-iot-topicrule-s3action-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The ARN of the IAM role that grants access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-s3action.html#cfn-iot-topicrule-s3action-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The Amazon S3 canned ACL that controls access to the object identified by the object key.
	//
	// For more information, see [S3 canned ACLs](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-s3action.html#cfn-iot-topicrule-s3action-cannedacl
	//
	CannedAcl *string `field:"optional" json:"cannedAcl" yaml:"cannedAcl"`
}

