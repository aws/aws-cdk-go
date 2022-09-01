package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// CodePipeline action to delete a stack.
//
// Deletes a stack. If you specify a stack that doesn't exist, the action completes successfully
// without deleting a stack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var artifact artifact
//   var artifactPath artifactPath
//   var parameterOverrides interface{}
//   var role role
//
//   cloudFormationDeleteStackAction := awscdk.Aws_codepipeline_actions.NewCloudFormationDeleteStackAction(&cloudFormationDeleteStackActionProps{
//   	actionName: jsii.String("actionName"),
//   	adminPermissions: jsii.Boolean(false),
//   	stackName: jsii.String("stackName"),
//
//   	// the properties below are optional
//   	account: jsii.String("account"),
//   	cfnCapabilities: []cfnCapabilities{
//   		cdk.*cfnCapabilities_NONE,
//   	},
//   	deploymentRole: role,
//   	extraInputs: []*artifact{
//   		artifact,
//   	},
//   	output: artifact,
//   	outputFileName: jsii.String("outputFileName"),
//   	parameterOverrides: map[string]interface{}{
//   		"parameterOverridesKey": parameterOverrides,
//   	},
//   	region: jsii.String("region"),
//   	role: role,
//   	runOrder: jsii.Number(123),
//   	templateConfiguration: artifactPath,
//   	variablesNamespace: jsii.String("variablesNamespace"),
//   })
//
type CloudFormationDeleteStackAction interface {
	Action
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the {@link bind} callback.
	ActionProperties() *awscodepipeline.ActionProperties
	DeploymentRole() awsiam.IRole
	// This is a renamed version of the {@link IAction.actionProperties} property.
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	// Add statement to the service role assumed by CloudFormation while executing this action.
	AddToDeploymentRolePolicy(statement awsiam.PolicyStatement) *bool
	// The callback invoked when this Action is added to a Pipeline.
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// This is a renamed version of the {@link IAction.bind} method.
	Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for CloudFormationDeleteStackAction
type jsiiProxy_CloudFormationDeleteStackAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_CloudFormationDeleteStackAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationDeleteStackAction) DeploymentRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"deploymentRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationDeleteStackAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewCloudFormationDeleteStackAction(props *CloudFormationDeleteStackActionProps) CloudFormationDeleteStackAction {
	_init_.Initialize()

	if err := validateNewCloudFormationDeleteStackActionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudFormationDeleteStackAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CloudFormationDeleteStackAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewCloudFormationDeleteStackAction_Override(c CloudFormationDeleteStackAction, props *CloudFormationDeleteStackActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CloudFormationDeleteStackAction",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_CloudFormationDeleteStackAction) AddToDeploymentRolePolicy(statement awsiam.PolicyStatement) *bool {
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

func (c *jsiiProxy_CloudFormationDeleteStackAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
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

func (c *jsiiProxy_CloudFormationDeleteStackAction) Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
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

func (c *jsiiProxy_CloudFormationDeleteStackAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
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

func (c *jsiiProxy_CloudFormationDeleteStackAction) VariableExpression(variableName *string) *string {
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

