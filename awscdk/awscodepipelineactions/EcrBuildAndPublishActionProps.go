package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Construction properties of the `EcrBuildAndPublishAction`.
//
// Example:
//   import ecr "github.com/aws/aws-cdk-go/awscdk"
//
//   var pipeline pipeline
//   var repository iRepository
//
//
//   sourceOutput := codepipeline.NewArtifact()
//   // your source repository
//   sourceAction := codepipeline_actions.NewCodeStarConnectionsSourceAction(&CodeStarConnectionsSourceActionProps{
//   	ActionName: jsii.String("CodeStarConnectionsSourceAction"),
//   	Output: sourceOutput,
//   	ConnectionArn: jsii.String("your-connection-arn"),
//   	Owner: jsii.String("your-owner"),
//   	Repo: jsii.String("your-repo"),
//   })
//
//   buildAction := codepipeline_actions.NewEcrBuildAndPublishAction(&EcrBuildAndPublishActionProps{
//   	ActionName: jsii.String("EcrBuildAndPublishAction"),
//   	RepositoryName: repository.RepositoryName,
//   	RegistryType: codepipeline_actions.RegistryType_PRIVATE,
//   	DockerfileDirectoryPath: jsii.String("./my-dir"),
//   	 // The path indicates ./my-dir/Dockerfile in the source repository
//   	ImageTags: []*string{
//   		jsii.String("my-tag-1"),
//   		jsii.String("my-tag-2"),
//   	},
//   	Input: sourceOutput,
//   })
//
//   pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("Source"),
//   	Actions: []iAction{
//   		sourceAction,
//   	},
//   })
//   pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("Build"),
//   	Actions: []*iAction{
//   		buildAction,
//   	},
//   })
//
type EcrBuildAndPublishActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Default: 1.
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Default: - a name will be generated, based on the stage and action names,
	// if any of the action's variables were referenced - otherwise,
	// no namespace will be set.
	//
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your `IAction.bind`
	// method in the `ActionBindOptions.role` property.
	// Default: a new Role will be generated.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The artifact produced by the source action that contains the Dockerfile needed to build the image.
	Input awscodepipeline.Artifact `field:"required" json:"input" yaml:"input"`
	// The name of the ECR repository where the image is pushed.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// The directory path of Dockerfile used to build the image.
	//
	// Optionally, you can provide an alternate directory path if Dockerfile is not at the root level.
	// Default: - the source repository root level.
	//
	DockerfileDirectoryPath *string `field:"optional" json:"dockerfileDirectoryPath" yaml:"dockerfileDirectoryPath"`
	// The tags used for the image.
	// Default: - latest.
	//
	ImageTags *[]*string `field:"optional" json:"imageTags" yaml:"imageTags"`
	// Specifies whether the repository is public or private.
	// Default: - RegistryType.PRIVATE
	//
	RegistryType RegistryType `field:"optional" json:"registryType" yaml:"registryType"`
}

