package appstagingsynthesizeralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Information returned by the Staging Stack for each image asset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import app_staging_synthesizer_alpha "github.com/aws/aws-cdk-go/appstagingsynthesizeralpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var stack Stack
//
//   imageStagingLocation := &ImageStagingLocation{
//   	RepoName: jsii.String("repoName"),
//
//   	// the properties below are optional
//   	AssumeRoleArn: jsii.String("assumeRoleArn"),
//   	DependencyStack: stack,
//   }
//
// Experimental.
type ImageStagingLocation struct {
	// The name of the staging repository.
	// Experimental.
	RepoName *string `field:"required" json:"repoName" yaml:"repoName"`
	// The arn to assume to write files to this repository.
	// Default: - Don't assume a role.
	//
	// Experimental.
	AssumeRoleArn *string `field:"optional" json:"assumeRoleArn" yaml:"assumeRoleArn"`
	// The stack that creates this repository (leads to dependencies on it).
	// Default: - Don't add dependencies.
	//
	// Experimental.
	DependencyStack awscdk.Stack `field:"optional" json:"dependencyStack" yaml:"dependencyStack"`
}

