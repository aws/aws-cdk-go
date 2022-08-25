package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// CodePipeline action to deploy a stack.
//
// Creates the stack if the specified stack doesn't exist. If the stack exists,
// AWS CloudFormation updates the stack. Use this action to update existing
// stacks.
//
// AWS CodePipeline won't replace the stack, and will fail deployment if the
// stack is in a failed state. Use `ReplaceOnFailure` for an action that
// will delete and recreate the stack to try and recover from failed states.
//
// Use this action to automatically replace failed stacks without recovering or
// troubleshooting them. You would typically choose this mode for testing.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   // in stack for account 123456789012...
//   var otherAccountStack stack
//
//   actionRole := iam.NewRole(otherAccountStack, jsii.String("ActionRole"), &roleProps{
//   	assumedBy: iam.NewAccountPrincipal(jsii.String("123456789012")),
//   	// the role has to have a physical name set
//   	roleName: awscdk.PhysicalName_GENERATE_IF_NEEDED(),
//   })
//
//   // in the pipeline stack...
//   sourceOutput := codepipeline.NewArtifact()
//   codepipeline_actions.NewCloudFormationCreateUpdateStackAction(&cloudFormationCreateUpdateStackActionProps{
//   	actionName: jsii.String("CloudFormationCreateUpdate"),
//   	stackName: jsii.String("MyStackName"),
//   	adminPermissions: jsii.Boolean(true),
//   	templatePath: sourceOutput.atPath(jsii.String("template.yaml")),
//   	role: actionRole,
//   })
//
// Experimental.
type CloudFormationCreateUpdateStackAction interface {
	Action
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the {@link bind} callback.
	// Experimental.
	ActionProperties() *awscodepipeline.ActionProperties
	// Experimental.
	DeploymentRole() awsiam.IRole
	// This is a renamed version of the {@link IAction.actionProperties} property.
	// Experimental.
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	// Add statement to the service role assumed by CloudFormation while executing this action.
	// Experimental.
	AddToDeploymentRolePolicy(statement awsiam.PolicyStatement) *bool
	// The callback invoked when this Action is added to a Pipeline.
	// Experimental.
	Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// This is a renamed version of the {@link IAction.bind} method.
	// Experimental.
	Bound(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	// Experimental.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	// Experimental.
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for CloudFormationCreateUpdateStackAction
type jsiiProxy_CloudFormationCreateUpdateStackAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_CloudFormationCreateUpdateStackAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationCreateUpdateStackAction) DeploymentRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"deploymentRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationCreateUpdateStackAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


// Experimental.
func NewCloudFormationCreateUpdateStackAction(props *CloudFormationCreateUpdateStackActionProps) CloudFormationCreateUpdateStackAction {
	_init_.Initialize()

	j := jsiiProxy_CloudFormationCreateUpdateStackAction{}

	_jsii_.Create(
		"monocdk.aws_codepipeline_actions.CloudFormationCreateUpdateStackAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewCloudFormationCreateUpdateStackAction_Override(c CloudFormationCreateUpdateStackAction, props *CloudFormationCreateUpdateStackActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codepipeline_actions.CloudFormationCreateUpdateStackAction",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_CloudFormationCreateUpdateStackAction) AddToDeploymentRolePolicy(statement awsiam.PolicyStatement) *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"addToDeploymentRolePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationCreateUpdateStackAction) Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationCreateUpdateStackAction) Bound(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bound",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationCreateUpdateStackAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		c,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationCreateUpdateStackAction) VariableExpression(variableName *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

