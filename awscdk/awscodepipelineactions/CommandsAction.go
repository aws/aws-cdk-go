package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

// CodePipeline compute action that uses AWS Commands.
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
type CommandsAction interface {
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
	Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	// Reference a CodePipeline variable exported in the Commands action.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-env-vars.html
	//
	Variable(variableName *string) *string
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for CommandsAction
type jsiiProxy_CommandsAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_CommandsAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CommandsAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewCommandsAction(props *CommandsActionProps) CommandsAction {
	_init_.Initialize()

	if err := validateNewCommandsActionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CommandsAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CommandsAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewCommandsAction_Override(c CommandsAction, props *CommandsActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CommandsAction",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_CommandsAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
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

func (c *jsiiProxy_CommandsAction) Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
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

func (c *jsiiProxy_CommandsAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
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

func (c *jsiiProxy_CommandsAction) Variable(variableName *string) *string {
	if err := c.validateVariableParameters(variableName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"variable",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CommandsAction) VariableExpression(variableName *string) *string {
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

