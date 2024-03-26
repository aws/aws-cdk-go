package awsdatasync


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3Property := &S3Property{
//   	BucketAccessRoleArn: jsii.String("bucketAccessRoleArn"),
//   	S3BucketArn: jsii.String("s3BucketArn"),
//   	Subdirectory: jsii.String("subdirectory"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-s3.html
//
type CfnTask_S3Property struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-s3.html#cfn-datasync-task-s3-bucketaccessrolearn
	//
	BucketAccessRoleArn *string `field:"optional" json:"bucketAccessRoleArn" yaml:"bucketAccessRoleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-s3.html#cfn-datasync-task-s3-s3bucketarn
	//
	S3BucketArn *string `field:"optional" json:"s3BucketArn" yaml:"s3BucketArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-s3.html#cfn-datasync-task-s3-subdirectory
	//
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
}

