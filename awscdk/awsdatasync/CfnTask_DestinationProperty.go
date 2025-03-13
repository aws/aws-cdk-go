package awsdatasync


// Specifies where DataSync uploads your [task report](https://docs.aws.amazon.com/datasync/latest/userguide/task-reports.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationProperty := &DestinationProperty{
//   	S3: &S3Property{
//   		BucketAccessRoleArn: jsii.String("bucketAccessRoleArn"),
//   		S3BucketArn: jsii.String("s3BucketArn"),
//   		Subdirectory: jsii.String("subdirectory"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-destination.html
//
type CfnTask_DestinationProperty struct {
	// Specifies the Amazon S3 bucket where DataSync uploads your task report.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-destination.html#cfn-datasync-task-destination-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

