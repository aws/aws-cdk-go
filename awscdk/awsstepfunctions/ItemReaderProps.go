package awsstepfunctions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Base interface for Item Reader configuration properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   itemReaderProps := &ItemReaderProps{
//   	Bucket: bucket,
//
//   	// the properties below are optional
//   	MaxItems: jsii.Number(123),
//   }
//
type ItemReaderProps struct {
	// S3 Bucket containing objects to iterate over or a file with a list to iterate over.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// Limits the number of items passed to the Distributed Map state.
	// Default: - Distributed Map state will iterate over all items provided by the ItemReader.
	//
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
}

