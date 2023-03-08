package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
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
// Experimental.
type CloudWatchLoggingOptions struct {
	// The current status of the logs in Amazon CloudWatch Logs for a build project.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The Log Group to send logs to.
	// Experimental.
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// The prefix of the stream name of the Amazon CloudWatch Logs.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

