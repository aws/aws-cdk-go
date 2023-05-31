package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
)

// Customization options for a DeployCdkStackAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var artifact artifact
//
//   deployCdkStackActionOptions := &DeployCdkStackActionOptions{
//   	CloudAssemblyInput: artifact,
//
//   	// the properties below are optional
//   	BaseActionName: jsii.String("baseActionName"),
//   	ChangeSetName: jsii.String("changeSetName"),
//   	ExecuteRunOrder: jsii.Number(123),
//   	Output: artifact,
//   	OutputFileName: jsii.String("outputFileName"),
//   	PrepareRunOrder: jsii.Number(123),
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type DeployCdkStackActionOptions struct {
	// The CodePipeline artifact that holds the Cloud Assembly.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CloudAssemblyInput awscodepipeline.Artifact `field:"required" json:"cloudAssemblyInput" yaml:"cloudAssemblyInput"`
	// Base name of the action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	BaseActionName *string `field:"optional" json:"baseActionName" yaml:"baseActionName"`
	// Name of the change set to create and deploy.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ChangeSetName *string `field:"optional" json:"changeSetName" yaml:"changeSetName"`
	// Run order for the Execute action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ExecuteRunOrder *float64 `field:"optional" json:"executeRunOrder" yaml:"executeRunOrder"`
	// Artifact to write Stack Outputs to.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Output awscodepipeline.Artifact `field:"optional" json:"output" yaml:"output"`
	// Filename in output to write Stack outputs to.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OutputFileName *string `field:"optional" json:"outputFileName" yaml:"outputFileName"`
	// Run order for the Prepare action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	PrepareRunOrder *float64 `field:"optional" json:"prepareRunOrder" yaml:"prepareRunOrder"`
}

