package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var deploymentGroup serverDeploymentGroup
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"), &pipelineProps{
//   	pipelineName: jsii.String("MyPipeline"),
//   })
//
//   // add the source and build Stages to the Pipeline...
//   buildOutput := codepipeline.NewArtifact()
//   deployAction := codepipeline_actions.NewCodeDeployServerDeployAction(&codeDeployServerDeployActionProps{
//   	actionName: jsii.String("CodeDeploy"),
//   	input: buildOutput,
//   	deploymentGroup: deploymentGroup,
//   })
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("Deploy"),
//   	actions: []iAction{
//   		deployAction,
//   	},
//   })
//
type CodeDeployServerDeployAction interface {
	Action
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the `bind` callback.
	ActionProperties() *awscodepipeline.ActionProperties
	// This is a renamed version of the `IAction.actionProperties` property.
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	// The callback invoked when this Action is added to a Pipeline.
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// This is a renamed version of the `IAction.bind` method.
	Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for CodeDeployServerDeployAction
type jsiiProxy_CodeDeployServerDeployAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_CodeDeployServerDeployAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeDeployServerDeployAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewCodeDeployServerDeployAction(props *CodeDeployServerDeployActionProps) CodeDeployServerDeployAction {
	_init_.Initialize()

	if err := validateNewCodeDeployServerDeployActionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodeDeployServerDeployAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CodeDeployServerDeployAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewCodeDeployServerDeployAction_Override(c CodeDeployServerDeployAction, props *CodeDeployServerDeployActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CodeDeployServerDeployAction",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_CodeDeployServerDeployAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
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

func (c *jsiiProxy_CodeDeployServerDeployAction) Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := c.validateBoundParameters(_scope, _stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bound",
		[]interface{}{_scope, _stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeDeployServerDeployAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
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

func (c *jsiiProxy_CodeDeployServerDeployAction) VariableExpression(variableName *string) *string {
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

