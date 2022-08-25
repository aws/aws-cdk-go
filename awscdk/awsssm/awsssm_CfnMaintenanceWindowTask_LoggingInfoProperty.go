package awsssm


// The `LoggingInfo` property type specifies information about the Amazon S3 bucket to write instance-level logs to.
//
// `LoggingInfo` is a property of the [AWS::SSM::MaintenanceWindowTask](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html) resource.
//
// > `LoggingInfo` has been deprecated. To specify an Amazon S3 bucket to contain logs, instead use the `OutputS3BucketName` and `OutputS3KeyPrefix` options in the `TaskInvocationParameters` structure. For information about how Systems Manager handles these options for the supported maintenance window task types, see [AWS Systems Manager MaintenanceWindowTask TaskInvocationParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingInfoProperty := &loggingInfoProperty{
//   	region: jsii.String("region"),
//   	s3Bucket: jsii.String("s3Bucket"),
//
//   	// the properties below are optional
//   	s3Prefix: jsii.String("s3Prefix"),
//   }
//
type CfnMaintenanceWindowTask_LoggingInfoProperty struct {
	// The AWS Region where the S3 bucket is located.
	Region *string `field:"required" json:"region" yaml:"region"`
	// The name of an S3 bucket where execution logs are stored .
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// The Amazon S3 bucket subfolder.
	S3Prefix *string `field:"optional" json:"s3Prefix" yaml:"s3Prefix"`
}

