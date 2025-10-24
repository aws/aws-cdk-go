package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// S3 bucket configuration for the validation data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket Bucket
//
//   validationBucketConfiguration := &ValidationBucketConfiguration{
//   	Bucket: bucket,
//
//   	// the properties below are optional
//   	Path: jsii.String("path"),
//   }
//
type ValidationBucketConfiguration struct {
	// The S3 bucket.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// Path to file or directory within the bucket.
	// Default: - root of the bucket.
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
}

