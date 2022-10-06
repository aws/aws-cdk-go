package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Properties for a DeployCdkStackAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var artifact artifact
//   var role role
//
//   deployCdkStackActionProps := &deployCdkStackActionProps{
//   	actionRole: role,
//   	cloudAssemblyInput: artifact,
//   	stackName: jsii.String("stackName"),
//   	templatePath: jsii.String("templatePath"),
//
//   	// the properties below are optional
//   	baseActionName: jsii.String("baseActionName"),
//   	changeSetName: jsii.String("changeSetName"),
//   	cloudFormationExecutionRole: role,
//   	dependencyStackArtifactIds: []*string{
//   		jsii.String("dependencyStackArtifactIds"),
//   	},
//   	executeRunOrder: jsii.Number(123),
//   	output: artifact,
//   	outputFileName: jsii.String("outputFileName"),
//   	prepareRunOrder: jsii.Number(123),
//   	region: jsii.String("region"),
//   	stackArtifactId: jsii.String("stackArtifactId"),
//   	templateConfigurationPath: jsii.String("templateConfigurationPath"),
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type DeployCdkStackActionProps struct {
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
	// Role for the action to assume.
	//
	// This controls the account to deploy into.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ActionRole awsiam.IRole `field:"required" json:"actionRole" yaml:"actionRole"`
	// The name of the stack that should be created/updated.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
	// Relative path of template in the input artifact.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	TemplatePath *string `field:"required" json:"templatePath" yaml:"templatePath"`
	// Role to execute CloudFormation under.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CloudFormationExecutionRole awsiam.IRole `field:"optional" json:"cloudFormationExecutionRole" yaml:"cloudFormationExecutionRole"`
	// Artifact ID for the stacks this stack depends on.
	//
	// Used for pipeline order checking.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	DependencyStackArtifactIds *[]*string `field:"optional" json:"dependencyStackArtifactIds" yaml:"dependencyStackArtifactIds"`
	// Region to deploy into.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Artifact ID for the stack deployed here.
	//
	// Used for pipeline order checking.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	StackArtifactId *string `field:"optional" json:"stackArtifactId" yaml:"stackArtifactId"`
	// Template configuration path relative to the input artifact.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	TemplateConfigurationPath *string `field:"optional" json:"templateConfigurationPath" yaml:"templateConfigurationPath"`
}

