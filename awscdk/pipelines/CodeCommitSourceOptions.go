package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipelineactions"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Configuration options for a CodeCommit source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role Role
//
//   codeCommitSourceOptions := &CodeCommitSourceOptions{
//   	ActionName: jsii.String("actionName"),
//   	CodeBuildCloneOutput: jsii.Boolean(false),
//   	EventRole: role,
//   	Trigger: awscdk.Aws_codepipeline_actions.CodeCommitTrigger_NONE,
//   }
//
type CodeCommitSourceOptions struct {
	// The action name used for this source in the CodePipeline.
	// Default: - The repository name.
	//
	ActionName *string `field:"optional" json:"actionName" yaml:"actionName"`
	// If this is set, the next CodeBuild job clones the repository (instead of CodePipeline downloading the files).
	//
	// This provides access to repository history, and retains symlinks (symlinks would otherwise be
	// removed by CodePipeline).
	//
	// **Note**: if this option is true, only CodeBuild jobs can use the output artifact.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodeCommit.html
	//
	// Default: false.
	//
	CodeBuildCloneOutput *bool `field:"optional" json:"codeBuildCloneOutput" yaml:"codeBuildCloneOutput"`
	// Role to be used by on commit event rule.
	//
	// Used only when trigger value is CodeCommitTrigger.EVENTS.
	// Default: a new role will be created.
	//
	EventRole awsiam.IRole `field:"optional" json:"eventRole" yaml:"eventRole"`
	// How should CodePipeline detect source changes for this Action.
	// Default: CodeCommitTrigger.EVENTS
	//
	Trigger awscodepipelineactions.CodeCommitTrigger `field:"optional" json:"trigger" yaml:"trigger"`
}

