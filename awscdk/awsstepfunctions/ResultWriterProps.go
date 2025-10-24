package awsstepfunctions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Interface for Result Writer configuration props.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket Bucket
//
//   resultWriterProps := &ResultWriterProps{
//   	Bucket: bucket,
//
//   	// the properties below are optional
//   	Prefix: jsii.String("prefix"),
//   }
//
// Deprecated: use {@link ResultWriterV2Props } instead.
type ResultWriterProps struct {
	// S3 Bucket in which to save Map Run results.
	// Deprecated: use {@link ResultWriterV2Props } instead.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// S3 prefix in which to save Map Run results.
	// Default: - No prefix.
	//
	// Deprecated: use {@link ResultWriterV2Props } instead.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

