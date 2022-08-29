package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CodePipeline source action for the CodeStar Connections source, which allows connecting to GitHub and BitBucket.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   sourceOutput := codepipeline.NewArtifact()
//   sourceAction := codepipeline_actions.NewCodeStarConnectionsSourceAction(&codeStarConnectionsSourceActionProps{
//   	actionName: jsii.String("BitBucket_Source"),
//   	owner: jsii.String("aws"),
//   	repo: jsii.String("aws-cdk"),
//   	output: sourceOutput,
//   	connectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:123456789012:connection/12345678-abcd-12ab-34cdef5678gh"),
//   })
//
type CodeStarConnectionsSourceAction interface {
	Action
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the {@link bind} callback.
	ActionProperties() *awscodepipeline.ActionProperties
	// This is a renamed version of the {@link IAction.actionProperties} property.
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	// The variables emitted by this action.
	Variables() *CodeStarSourceVariables
	// The callback invoked when this Action is added to a Pipeline.
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// This is a renamed version of the {@link IAction.bind} method.
	Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for CodeStarConnectionsSourceAction
type jsiiProxy_CodeStarConnectionsSourceAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_CodeStarConnectionsSourceAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeStarConnectionsSourceAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeStarConnectionsSourceAction) Variables() *CodeStarSourceVariables {
	var returns *CodeStarSourceVariables
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}


func NewCodeStarConnectionsSourceAction(props *CodeStarConnectionsSourceActionProps) CodeStarConnectionsSourceAction {
	_init_.Initialize()

	j := jsiiProxy_CodeStarConnectionsSourceAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CodeStarConnectionsSourceAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewCodeStarConnectionsSourceAction_Override(c CodeStarConnectionsSourceAction, props *CodeStarConnectionsSourceActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CodeStarConnectionsSourceAction",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_CodeStarConnectionsSourceAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeStarConnectionsSourceAction) Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bound",
		[]interface{}{_scope, _stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeStarConnectionsSourceAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		c,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeStarConnectionsSourceAction) VariableExpression(variableName *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

