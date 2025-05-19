package awscodebuild


// `LogsConfig` is a property of the [AWS CodeBuild Project](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html) resource that specifies information about logs for a build project. These can be logs in Amazon CloudWatch Logs, built in a specified S3 bucket, or both.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logsConfigProperty := &LogsConfigProperty{
//   	CloudWatchLogs: &CloudWatchLogsConfigProperty{
//   		Status: jsii.String("status"),
//
//   		// the properties below are optional
//   		GroupName: jsii.String("groupName"),
//   		StreamName: jsii.String("streamName"),
//   	},
//   	S3Logs: &S3LogsConfigProperty{
//   		Status: jsii.String("status"),
//
//   		// the properties below are optional
//   		EncryptionDisabled: jsii.Boolean(false),
//   		Location: jsii.String("location"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-logsconfig.html
//
type CfnProject_LogsConfigProperty struct {
	// Information about CloudWatch Logs for a build project.
	//
	// CloudWatch Logs are enabled by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-logsconfig.html#cfn-codebuild-project-logsconfig-cloudwatchlogs
	//
	CloudWatchLogs interface{} `field:"optional" json:"cloudWatchLogs" yaml:"cloudWatchLogs"`
	// Information about logs built to an S3 bucket for a build project.
	//
	// S3 logs are not enabled by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-logsconfig.html#cfn-codebuild-project-logsconfig-s3logs
	//
	S3Logs interface{} `field:"optional" json:"s3Logs" yaml:"s3Logs"`
}

