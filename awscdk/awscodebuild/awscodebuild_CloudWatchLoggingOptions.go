package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Information about logs built to a CloudWatch Log Group for a build project.
//
// Example:
//   codebuild.NewProject(this, jsii.String("Project"), &projectProps{
//   	logging: &loggingOptions{
//   		cloudWatch: &cloudWatchLoggingOptions{
//   			logGroup: logs.NewLogGroup(this, jsii.String("MyLogGroup")),
//   		},
//   	},
//   })
//
type CloudWatchLoggingOptions struct {
	// The current status of the logs in Amazon CloudWatch Logs for a build project.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The Log Group to send logs to.
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// The prefix of the stream name of the Amazon CloudWatch Logs.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

