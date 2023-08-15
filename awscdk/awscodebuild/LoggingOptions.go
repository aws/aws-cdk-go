package awscodebuild


// Information about logs for the build project.
//
// A project can create logs in Amazon CloudWatch Logs, an S3 bucket, or both.
//
// Example:
//   codebuild.NewProject(this, jsii.String("Project"), &ProjectProps{
//   	Logging: &LoggingOptions{
//   		CloudWatch: &CloudWatchLoggingOptions{
//   			LogGroup: logs.NewLogGroup(this, jsii.String("MyLogGroup")),
//   		},
//   	},
//   })
//
type LoggingOptions struct {
	// Information about Amazon CloudWatch Logs for a build project.
	// Default: - enabled.
	//
	CloudWatch *CloudWatchLoggingOptions `field:"optional" json:"cloudWatch" yaml:"cloudWatch"`
	// Information about logs built to an S3 bucket for a build project.
	// Default: - disabled.
	//
	S3 *S3LoggingOptions `field:"optional" json:"s3" yaml:"s3"`
}

