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
//   var bucket Bucket
//
//   itemReaderProps := &ItemReaderProps{
//   	Bucket: bucket,
//   	BucketNamePath: jsii.String("bucketNamePath"),
//   	MaxItems: jsii.Number(123),
//   }
//
type ItemReaderProps struct {
	// S3 Bucket containing objects to iterate over or a file with a list to iterate over.
	// See: bucketNamePath.
	//
	// Default: - S3 bucket will be determined from.
	//
	Bucket awss3.IBucket `field:"optional" json:"bucket" yaml:"bucket"`
	// S3 bucket name containing objects to iterate over or a file with a list to iterate over, as JsonPath.
	// See: bucket.
	//
	// Default: - S3 bucket will be determined from.
	//
	BucketNamePath *string `field:"optional" json:"bucketNamePath" yaml:"bucketNamePath"`
	// Limits the number of items passed to the Distributed Map state.
	// Default: - Distributed Map state will iterate over all items provided by the ItemReader.
	//
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
}

