package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Construction properties for `CodeStarConnectionsSourceAction`.
//
// Example:
//   sourceOutput := codepipeline.NewArtifact()
//   sourceAction := codepipeline_actions.NewCodeStarConnectionsSourceAction(&CodeStarConnectionsSourceActionProps{
//   	ActionName: jsii.String("BitBucket_Source"),
//   	Owner: jsii.String("aws"),
//   	Repo: jsii.String("aws-cdk"),
//   	Output: sourceOutput,
//   	ConnectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:123456789012:connection/12345678-abcd-12ab-34cdef5678gh"),
//   })
//
type CodeStarConnectionsSourceActionProps struct {
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
	// The ARN of the CodeStar Connection created in the AWS console that has permissions to access this GitHub or BitBucket repository.
	//
	// Example:
	//   "arn:aws:codestar-connections:us-east-1:123456789012:connection/12345678-abcd-12ab-34cdef5678gh"
	//
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/connections-create.html
	//
	ConnectionArn *string `field:"required" json:"connectionArn" yaml:"connectionArn"`
	// The output artifact that this action produces.
	//
	// Can be used as input for further pipeline actions.
	Output awscodepipeline.Artifact `field:"required" json:"output" yaml:"output"`
	// The owning user or organization of the repository.
	//
	// Example:
	//   "aws"
	//
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// The name of the repository.
	//
	// Example:
	//   "aws-cdk"
	//
	Repo *string `field:"required" json:"repo" yaml:"repo"`
	// The branch to build.
	// Default: 'master'.
	//
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// Whether the output should be the contents of the repository (which is the default), or a link that allows CodeBuild to clone the repository before building.
	//
	// **Note**: if this option is true,
	// then only CodeBuild actions can use the resulting `output`.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodestarConnectionSource.html#action-reference-CodestarConnectionSource-config
	//
	// Default: false.
	//
	CodeBuildCloneOutput *bool `field:"optional" json:"codeBuildCloneOutput" yaml:"codeBuildCloneOutput"`
	// Controls automatically starting your pipeline when a new commit is made on the configured repository and branch.
	//
	// If unspecified,
	// the default value is true, and the field does not display by default.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodestarConnectionSource.html
	//
	// Default: true.
	//
	TriggerOnPush *bool `field:"optional" json:"triggerOnPush" yaml:"triggerOnPush"`
}

