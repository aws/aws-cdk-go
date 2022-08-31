package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// CodePipeline action to delete a stack.
//
// Deletes a stack. If you specify a stack that doesn't exist, the action completes successfully
// without deleting a stack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//   	capabilities: []cloudFormationCapabilities{
//   		awscdk.Aws_cloudformation.*cloudFormationCapabilities_NONE,
//   	},
//   	cfnCapabilities: []cfnCapabilities{
//   		monocdk.*cfnCapabilities_NONE,
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
// Experimental.
type CloudFormationDeleteStackAction interface {
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


// Experimental.
func NewCloudFormationDeleteStackAction(props *CloudFormationDeleteStackActionProps) CloudFormationDeleteStackAction {
	_init_.Initialize()

	j := jsiiProxy_CloudFormationDeleteStackAction{}

	_jsii_.Create(
		"monocdk.aws_codepipeline_actions.CloudFormationDeleteStackAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewCloudFormationDeleteStackAction_Override(c CloudFormationDeleteStackAction, props *CloudFormationDeleteStackActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codepipeline_actions.CloudFormationDeleteStackAction",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_CloudFormationDeleteStackAction) AddToDeploymentRolePolicy(statement awsiam.PolicyStatement) *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"addToDeploymentRolePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationDeleteStackAction) Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationDeleteStackAction) Bound(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
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
	var returns *string

	_jsii_.Invoke(
		c,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

