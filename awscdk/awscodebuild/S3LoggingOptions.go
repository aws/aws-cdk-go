package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Information about logs built to an S3 bucket for a build project.
//
// Example:
//   codebuild.NewProject(this, jsii.String("Project"), &ProjectProps{
//   	Logging: &LoggingOptions{
//   		S3: &S3LoggingOptions{
//   			Bucket: s3.NewBucket(this, jsii.String("LogBucket")),
//   		},
//   	},
//   })
//
type S3LoggingOptions struct {
	// The S3 Bucket to send logs to.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The current status of the logs in Amazon CloudWatch Logs for a build project.
	// Default: true.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Encrypt the S3 build log output.
	// Default: true.
	//
	Encrypted *bool `field:"optional" json:"encrypted" yaml:"encrypted"`
	// The path prefix for S3 logs.
	// Default: - no prefix.
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

