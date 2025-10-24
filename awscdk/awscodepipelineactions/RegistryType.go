package awscodepipelineactions


// The type of registry to use for the EcrBuildAndPublish action.
//
// Example:
//   import ecr "github.com/aws/aws-cdk-go/awscdk"
//
//   var pipeline Pipeline
//   var repository IRepository
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
//   	Actions: []IAction{
//   		sourceAction,
//   	},
//   })
//   pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("Build"),
//   	Actions: []IAction{
//   		buildAction,
//   	},
//   })
//
type RegistryType string

const (
	// Private registry.
	RegistryType_PRIVATE RegistryType = "PRIVATE"
	// Public registry.
	RegistryType_PUBLIC RegistryType = "PUBLIC"
)

