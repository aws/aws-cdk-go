package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
)

// Options for CdkDeployAction.fromStackArtifact.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var artifact artifact
//
//   fromStackArtifactOptions := &fromStackArtifactOptions{
//   	cloudAssemblyInput: artifact,
//
//   	// the properties below are optional
//   	executeRunOrder: jsii.Number(123),
//   	output: artifact,
//   	outputFileName: jsii.String("outputFileName"),
//   	prepareRunOrder: jsii.Number(123),
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type FromStackArtifactOptions struct {
	// The CodePipeline artifact that holds the Cloud Assembly.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CloudAssemblyInput awscodepipeline.Artifact `field:"required" json:"cloudAssemblyInput" yaml:"cloudAssemblyInput"`
	// Run order for the Execute action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ExecuteRunOrder *float64 `field:"optional" json:"executeRunOrder" yaml:"executeRunOrder"`
	// Artifact to write Stack Outputs to.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Output awscodepipeline.Artifact `field:"optional" json:"output" yaml:"output"`
	// Filename in output to write Stack outputs to.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OutputFileName *string `field:"optional" json:"outputFileName" yaml:"outputFileName"`
	// Run order for the 2 actions that will be created.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	PrepareRunOrder *float64 `field:"optional" json:"prepareRunOrder" yaml:"prepareRunOrder"`
}

