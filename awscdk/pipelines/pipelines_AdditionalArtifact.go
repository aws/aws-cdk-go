package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
)

// Specification of an additional artifact to generate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var artifact artifact
//
//   additionalArtifact := &additionalArtifact{
//   	artifact: artifact,
//   	directory: jsii.String("directory"),
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type AdditionalArtifact struct {
	// Artifact to represent the build directory in the pipeline.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Artifact awscodepipeline.Artifact `field:"required" json:"artifact" yaml:"artifact"`
	// Directory to be packaged.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Directory *string `field:"required" json:"directory" yaml:"directory"`
}

