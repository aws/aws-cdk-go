package appstagingsynthesizeralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Information returned by the Staging Stack for each file asset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import app_staging_synthesizer_alpha "github.com/aws/aws-cdk-go/appstagingsynthesizeralpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var stack stack
//
//   fileStagingLocation := &FileStagingLocation{
//   	BucketName: jsii.String("bucketName"),
//
//   	// the properties below are optional
//   	AssumeRoleArn: jsii.String("assumeRoleArn"),
//   	DependencyStack: stack,
//   	Prefix: jsii.String("prefix"),
//   }
//
// Experimental.
type FileStagingLocation struct {
	// The name of the staging bucket.
	// Experimental.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The ARN to assume to write files to this bucket.
	// Default: - Don't assume a role.
	//
	// Experimental.
	AssumeRoleArn *string `field:"optional" json:"assumeRoleArn" yaml:"assumeRoleArn"`
	// The stack that creates this bucket (leads to dependencies on it).
	// Default: - Don't add dependencies.
	//
	// Experimental.
	DependencyStack awscdk.Stack `field:"optional" json:"dependencyStack" yaml:"dependencyStack"`
	// A prefix to add to the keys.
	// Default: ''.
	//
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

