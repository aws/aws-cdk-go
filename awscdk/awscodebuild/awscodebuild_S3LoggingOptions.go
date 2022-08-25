package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

// Information about logs built to an S3 bucket for a build project.
//
// Example:
//   codebuild.NewProject(this, jsii.String("Project"), &projectProps{
//   	logging: &loggingOptions{
//   		s3: &s3LoggingOptions{
//   			bucket: s3.NewBucket(this, jsii.String("LogBucket")),
//   		},
//   	},
//   })
//
// Experimental.
type S3LoggingOptions struct {
	// The S3 Bucket to send logs to.
	// Experimental.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The current status of the logs in Amazon CloudWatch Logs for a build project.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Encrypt the S3 build log output.
	// Experimental.
	Encrypted *bool `field:"optional" json:"encrypted" yaml:"encrypted"`
	// The path prefix for S3 logs.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

