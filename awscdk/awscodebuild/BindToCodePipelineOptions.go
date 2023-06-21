package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// The extra options passed to the `IProject.bindToCodePipeline` method.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   bindToCodePipelineOptions := &BindToCodePipelineOptions{
//   	ArtifactBucket: bucket,
//   }
//
type BindToCodePipelineOptions struct {
	// The artifact bucket that will be used by the action that invokes this project.
	ArtifactBucket awss3.IBucket `field:"required" json:"artifactBucket" yaml:"artifactBucket"`
}

