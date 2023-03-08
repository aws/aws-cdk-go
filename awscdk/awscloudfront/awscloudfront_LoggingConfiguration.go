package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

// Logging configuration for incoming requests.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   loggingConfiguration := &loggingConfiguration{
//   	bucket: bucket,
//   	includeCookies: jsii.Boolean(false),
//   	prefix: jsii.String("prefix"),
//   }
//
// Experimental.
type LoggingConfiguration struct {
	// Bucket to log requests to.
	// Experimental.
	Bucket awss3.IBucket `field:"optional" json:"bucket" yaml:"bucket"`
	// Whether to include the cookies in the logs.
	// Experimental.
	IncludeCookies *bool `field:"optional" json:"includeCookies" yaml:"includeCookies"`
	// Where in the bucket to store logs.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

