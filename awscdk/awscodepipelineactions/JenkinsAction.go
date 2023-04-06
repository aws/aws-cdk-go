package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

// Jenkins build CodePipeline Action.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var jenkinsProvider jenkinsProvider
//
//   buildAction := codepipeline_actions.NewJenkinsAction(&JenkinsActionProps{
//   	ActionName: jsii.String("JenkinsBuild"),
//   	JenkinsProvider: jenkinsProvider,
//   	ProjectName: jsii.String("MyProject"),
//   	Type: codepipeline_actions.JenkinsActionType_BUILD,
//   })
//
// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/tutorials-four-stage-pipeline.html
//
type JenkinsAction interface {
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
	Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, _options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for JenkinsAction
type jsiiProxy_JenkinsAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_JenkinsAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JenkinsAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewJenkinsAction(props *JenkinsActionProps) JenkinsAction {
	_init_.Initialize()

	if err := validateNewJenkinsActionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_JenkinsAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.JenkinsAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewJenkinsAction_Override(j JenkinsAction, props *JenkinsActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.JenkinsAction",
		[]interface{}{props},
		j,
	)
}

func (j *jsiiProxy_JenkinsAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := j.validateBindParameters(scope, stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		j,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JenkinsAction) Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, _options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := j.validateBoundParameters(_scope, _stage, _options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		j,
		"bound",
		[]interface{}{_scope, _stage, _options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JenkinsAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	if err := j.validateOnStateChangeParameters(name, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		j,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JenkinsAction) VariableExpression(variableName *string) *string {
	if err := j.validateVariableExpressionParameters(variableName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

