package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
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
//   actionRole := iam.NewRole(otherAccountStack, jsii.String("ActionRole"), &RoleProps{
//   	AssumedBy: iam.NewAccountPrincipal(jsii.String("123456789012")),
//   	// the role has to have a physical name set
//   	RoleName: awscdk.PhysicalName_GENERATE_IF_NEEDED(),
//   })
//
//   // in the pipeline stack...
//   sourceOutput := codepipeline.NewArtifact()
//   codepipeline_actions.NewCloudFormationCreateUpdateStackAction(&CloudFormationCreateUpdateStackActionProps{
//   	ActionName: jsii.String("CloudFormationCreateUpdate"),
//   	StackName: jsii.String("MyStackName"),
//   	AdminPermissions: jsii.Boolean(true),
//   	TemplatePath: sourceOutput.AtPath(jsii.String("template.yaml")),
//   	Role: actionRole,
//   })
//
type CloudFormationCreateUpdateStackAction interface {
	Action
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the `bind` callback.
	ActionProperties() *awscodepipeline.ActionProperties
	DeploymentRole() awsiam.IRole
	// This is a renamed version of the `IAction.actionProperties` property.
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	// Add statement to the service role assumed by CloudFormation while executing this action.
	AddToDeploymentRolePolicy(statement awsiam.PolicyStatement) *bool
	// The callback invoked when this Action is added to a Pipeline.
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// This is a renamed version of the `IAction.bind` method.
	Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
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


func NewCloudFormationCreateUpdateStackAction(props *CloudFormationCreateUpdateStackActionProps) CloudFormationCreateUpdateStackAction {
	_init_.Initialize()

	if err := validateNewCloudFormationCreateUpdateStackActionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudFormationCreateUpdateStackAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CloudFormationCreateUpdateStackAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewCloudFormationCreateUpdateStackAction_Override(c CloudFormationCreateUpdateStackAction, props *CloudFormationCreateUpdateStackActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CloudFormationCreateUpdateStackAction",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_CloudFormationCreateUpdateStackAction) AddToDeploymentRolePolicy(statement awsiam.PolicyStatement) *bool {
	if err := c.validateAddToDeploymentRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"addToDeploymentRolePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationCreateUpdateStackAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := c.validateBindParameters(scope, stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationCreateUpdateStackAction) Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := c.validateBoundParameters(scope, stage, options); err != nil {
		panic(err)
	}
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
	if err := c.validateOnStateChangeParameters(name, options); err != nil {
		panic(err)
	}
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
	if err := c.validateVariableExpressionParameters(variableName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

