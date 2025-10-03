package awscloudtrail

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Selecting an S3 bucket and an optional prefix to be logged for data events.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   s3EventSelector := &S3EventSelector{
//   	Bucket: bucket,
//
//   	// the properties below are optional
//   	ObjectPrefix: jsii.String("objectPrefix"),
//   }
//
type S3EventSelector struct {
	// S3 bucket.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// Data events for objects whose key matches this prefix will be logged.
	// Default: - all objects.
	//
	ObjectPrefix *string `field:"optional" json:"objectPrefix" yaml:"objectPrefix"`
}

