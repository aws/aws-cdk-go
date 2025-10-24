package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Construction properties of the `CommandsAction`.
//
// Example:
//   var sourceArtifact Artifact
//   var outputArtifact Artifact
//
//
//   commandsAction := codepipeline_actions.NewCommandsAction(&CommandsActionProps{
//   	ActionName: jsii.String("Commands"),
//   	Commands: []*string{
//   		jsii.String("export MY_OUTPUT=my-key"),
//   	},
//   	Input: sourceArtifact,
//   	Output: outputArtifact,
//   	OutputVariables: []*string{
//   		jsii.String("MY_OUTPUT"),
//   		jsii.String("CODEBUILD_BUILD_ID"),
//   	},
//   })
//
//   // Deploy action
//   deployAction := codepipeline_actions.NewS3DeployAction(&S3DeployActionProps{
//   	ActionName: jsii.String("DeployAction"),
//   	Extract: jsii.Boolean(true),
//   	Input: outputArtifact,
//   	Bucket: s3.NewBucket(this, jsii.String("DeployBucket")),
//   	ObjectKey: commandsAction.Variable(jsii.String("MY_OUTPUT")),
//   })
//
type CommandsActionProps struct {
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
	// Shell commands for the Commands action to run.
	//
	// All formats are supported except multi-line formats.
	//
	// The length of the commands array must be between 1 and 50.
	Commands *[]*string `field:"required" json:"commands" yaml:"commands"`
	// The source to use as input for this action.
	Input awscodepipeline.Artifact `field:"required" json:"input" yaml:"input"`
	// The list of additional input artifacts for this action.
	// Default: - no extra inputs.
	//
	ExtraInputs *[]awscodepipeline.Artifact `field:"optional" json:"extraInputs" yaml:"extraInputs"`
	// The output artifact for this action.
	//
	// You can filter files that you want to export as the output artifact for the action.
	//
	// Example:
	//   codepipeline.NewArtifact(jsii.String("CommandsArtifact"), []*string{
	//   	jsii.String("my-dir/**"),
	//   })
	//
	// Default: - no output artifact.
	//
	Output awscodepipeline.Artifact `field:"optional" json:"output" yaml:"output"`
	// The names of the variables in your environment that you want to export.
	//
	// These variables can be referenced in other actions by using the `variable` method
	// of this class.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-env-vars.html
	//
	// Default: - No output variables are exported.
	//
	OutputVariables *[]*string `field:"optional" json:"outputVariables" yaml:"outputVariables"`
}

