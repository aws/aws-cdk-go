package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
)

// Props for the UpdatePipelineAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var artifact artifact
//   var buildSpec buildSpec
//   var dockerCredential dockerCredential
//
//   updatePipelineActionProps := &updatePipelineActionProps{
//   	cloudAssemblyInput: artifact,
//   	pipelineStackHierarchicalId: jsii.String("pipelineStackHierarchicalId"),
//
//   	// the properties below are optional
//   	buildSpec: buildSpec,
//   	cdkCliVersion: jsii.String("cdkCliVersion"),
//   	dockerCredentials: []*dockerCredential{
//   		dockerCredential,
//   	},
//   	pipelineStackName: jsii.String("pipelineStackName"),
//   	privileged: jsii.Boolean(false),
//   	projectName: jsii.String("projectName"),
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type UpdatePipelineActionProps struct {
	// The CodePipeline artifact that holds the Cloud Assembly.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CloudAssemblyInput awscodepipeline.Artifact `field:"required" json:"cloudAssemblyInput" yaml:"cloudAssemblyInput"`
	// Hierarchical id of the pipeline stack.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	PipelineStackHierarchicalId *string `field:"required" json:"pipelineStackHierarchicalId" yaml:"pipelineStackHierarchicalId"`
	// Custom BuildSpec that is merged with generated one.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	BuildSpec awscodebuild.BuildSpec `field:"optional" json:"buildSpec" yaml:"buildSpec"`
	// Version of CDK CLI to 'npm install'.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CdkCliVersion *string `field:"optional" json:"cdkCliVersion" yaml:"cdkCliVersion"`
	// Docker registries and associated credentials necessary during the pipeline self-update stage.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	DockerCredentials *[]DockerCredential `field:"optional" json:"dockerCredentials" yaml:"dockerCredentials"`
	// Name of the pipeline stack.
	// Deprecated: - Use `pipelineStackHierarchicalId` instead.
	PipelineStackName *string `field:"optional" json:"pipelineStackName" yaml:"pipelineStackName"`
	// Whether the build step should run in privileged mode.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Privileged *bool `field:"optional" json:"privileged" yaml:"privileged"`
	// Name of the CodeBuild project.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ProjectName *string `field:"optional" json:"projectName" yaml:"projectName"`
}

