package awsstepfunctions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Properties for configuring an Item Reader that iterates over objects in an S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   s3ObjectsItemReaderProps := &S3ObjectsItemReaderProps{
//   	Bucket: bucket,
//
//   	// the properties below are optional
//   	MaxItems: jsii.Number(123),
//   	Prefix: jsii.String("prefix"),
//   }
//
type S3ObjectsItemReaderProps struct {
	// S3 Bucket containing objects to iterate over or a file with a list to iterate over.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// Limits the number of items passed to the Distributed Map state.
	// Default: - Distributed Map state will iterate over all items provided by the ItemReader.
	//
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// S3 prefix used to limit objects to iterate over.
	// Default: - No prefix.
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

