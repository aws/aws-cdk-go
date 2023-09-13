package awsdatasync


// Specifies the Amazon S3 bucket where DataSync uploads your [task report](https://docs.aws.amazon.com/datasync/latest/userguide/creating-task-reports.html) .
//
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
	// Specifies the Amazon Resource Name (ARN) of the IAM policy that allows DataSync to upload a task report to your S3 bucket.
	//
	// For more information, see [Allowing DataSync to upload a task report to an Amazon S3 bucket](https://docs.aws.amazon.com/datasync/latest/userguide/creating-task-reports.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-s3.html#cfn-datasync-task-s3-bucketaccessrolearn
	//
	BucketAccessRoleArn *string `field:"optional" json:"bucketAccessRoleArn" yaml:"bucketAccessRoleArn"`
	// Specifies the ARN of the S3 bucket where DataSync uploads your report.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-s3.html#cfn-datasync-task-s3-s3bucketarn
	//
	S3BucketArn *string `field:"optional" json:"s3BucketArn" yaml:"s3BucketArn"`
	// Specifies a bucket prefix for your report.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-s3.html#cfn-datasync-task-s3-subdirectory
	//
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
}

