package awscodebuild


// `LogsConfig` is a property of the [AWS CodeBuild Project](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html) resource that specifies information about logs for a build project. These can be logs in Amazon CloudWatch Logs, built in a specified S3 bucket, or both.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logsConfigProperty := &logsConfigProperty{
//   	cloudWatchLogs: &cloudWatchLogsConfigProperty{
//   		status: jsii.String("status"),
//
//   		// the properties below are optional
//   		groupName: jsii.String("groupName"),
//   		streamName: jsii.String("streamName"),
//   	},
//   	s3Logs: &s3LogsConfigProperty{
//   		status: jsii.String("status"),
//
//   		// the properties below are optional
//   		encryptionDisabled: jsii.Boolean(false),
//   		location: jsii.String("location"),
//   	},
//   }
//
type CfnProject_LogsConfigProperty struct {
	// Information about CloudWatch Logs for a build project.
	//
	// CloudWatch Logs are enabled by default.
	CloudWatchLogs interface{} `field:"optional" json:"cloudWatchLogs" yaml:"cloudWatchLogs"`
	// Information about logs built to an S3 bucket for a build project.
	//
	// S3 logs are not enabled by default.
	S3Logs interface{} `field:"optional" json:"s3Logs" yaml:"s3Logs"`
}

