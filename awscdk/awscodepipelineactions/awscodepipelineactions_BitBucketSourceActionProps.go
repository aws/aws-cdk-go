package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Construction properties for {@link BitBucketSourceAction}.
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
//   bitBucketSourceActionProps := &BitBucketSourceActionProps{
//   	ActionName: jsii.String("actionName"),
//   	ConnectionArn: jsii.String("connectionArn"),
//   	Output: artifact,
//   	Owner: jsii.String("owner"),
//   	Repo: jsii.String("repo"),
//
//   	// the properties below are optional
//   	Branch: jsii.String("branch"),
//   	CodeBuildCloneOutput: jsii.Boolean(false),
//   	Role: role,
//   	RunOrder: jsii.Number(123),
//   	TriggerOnPush: jsii.Boolean(false),
//   	VariablesNamespace: jsii.String("variablesNamespace"),
//   }
//
// Deprecated: use CodeStarConnectionsSourceActionProps instead.
type BitBucketSourceActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Deprecated: use CodeStarConnectionsSourceActionProps instead.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Deprecated: use CodeStarConnectionsSourceActionProps instead.
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Deprecated: use CodeStarConnectionsSourceActionProps instead.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	// Deprecated: use CodeStarConnectionsSourceActionProps instead.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The ARN of the CodeStar Connection created in the AWS console that has permissions to access this GitHub or BitBucket repository.
	//
	// Example:
	//   'arn:aws:codestar-connections:us-east-1:123456789012:connection/12345678-abcd-12ab-34cdef5678gh'
	//
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/connections-create.html
	//
	// Deprecated: use CodeStarConnectionsSourceActionProps instead.
	ConnectionArn *string `field:"required" json:"connectionArn" yaml:"connectionArn"`
	// The output artifact that this action produces.
	//
	// Can be used as input for further pipeline actions.
	// Deprecated: use CodeStarConnectionsSourceActionProps instead.
	Output awscodepipeline.Artifact `field:"required" json:"output" yaml:"output"`
	// The owning user or organization of the repository.
	//
	// Example:
	//   'aws'
	//
	// Deprecated: use CodeStarConnectionsSourceActionProps instead.
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// The name of the repository.
	//
	// Example:
	//   'aws-cdk'
	//
	// Deprecated: use CodeStarConnectionsSourceActionProps instead.
	Repo *string `field:"required" json:"repo" yaml:"repo"`
	// The branch to build.
	// Deprecated: use CodeStarConnectionsSourceActionProps instead.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// Whether the output should be the contents of the repository (which is the default), or a link that allows CodeBuild to clone the repository before building.
	//
	// **Note**: if this option is true,
	// then only CodeBuild actions can use the resulting {@link output}.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodestarConnectionSource.html#action-reference-CodestarConnectionSource-config
	//
	// Deprecated: use CodeStarConnectionsSourceActionProps instead.
	CodeBuildCloneOutput *bool `field:"optional" json:"codeBuildCloneOutput" yaml:"codeBuildCloneOutput"`
	// Controls automatically starting your pipeline when a new commit is made on the configured repository and branch.
	//
	// If unspecified,
	// the default value is true, and the field does not display by default.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodestarConnectionSource.html
	//
	// Deprecated: use CodeStarConnectionsSourceActionProps instead.
	TriggerOnPush *bool `field:"optional" json:"triggerOnPush" yaml:"triggerOnPush"`
}

