package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Construction properties for `S3Source`.
//
// Example:
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//
//   codebuild.NewProject(this, jsii.String("MyProject"), &ProjectProps{
//   	Source: codebuild.Source_S3(&S3SourceProps{
//   		Bucket: bucket,
//   		Path: jsii.String("path/to/file.zip"),
//   	}),
//   })
//
type S3SourceProps struct {
	// The source identifier.
	//
	// This property is required on secondary sources.
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	Path *string `field:"required" json:"path" yaml:"path"`
	// The version ID of the object that represents the build input ZIP file to use.
	// Default: latest.
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

