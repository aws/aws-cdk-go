package awssynthetics

import (
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

// Options for specifying the s3 location that stores the data of each canary run.
//
// The artifacts bucket location **cannot**
// be updated once the canary is created.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   artifactsBucketLocation := &artifactsBucketLocation{
//   	bucket: bucket,
//
//   	// the properties below are optional
//   	prefix: jsii.String("prefix"),
//   }
//
// Experimental.
type ArtifactsBucketLocation struct {
	// The s3 location that stores the data of each run.
	// Experimental.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The S3 bucket prefix.
	//
	// Specify this if you want a more specific path within the artifacts bucket.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

