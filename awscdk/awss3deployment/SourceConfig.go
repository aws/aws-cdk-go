package awss3deployment

import (
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

// Source information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//   var markers interface{}
//
//   sourceConfig := &SourceConfig{
//   	Bucket: bucket,
//   	ZipObjectKey: jsii.String("zipObjectKey"),
//
//   	// the properties below are optional
//   	Markers: map[string]interface{}{
//   		"markersKey": markers,
//   	},
//   }
//
// Experimental.
type SourceConfig struct {
	// The source bucket to deploy from.
	// Experimental.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// An S3 object key in the source bucket that points to a zip file.
	// Experimental.
	ZipObjectKey *string `field:"required" json:"zipObjectKey" yaml:"zipObjectKey"`
	// A set of markers to substitute in the source content.
	// Experimental.
	Markers *map[string]interface{} `field:"optional" json:"markers" yaml:"markers"`
}

