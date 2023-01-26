package awscodepipeline

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

// Low-level class for generic CodePipeline Actions implementing the `IAction` interface.
//
// Contains some common logic that can be re-used by all `IAction` implementations.
// If you're writing your own Action class,
// feel free to extend this class.
type Action interface {
	IAction
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the `bind` callback.
	ActionProperties() *ActionProperties
	// This is a renamed version of the `IAction.actionProperties` property.
	ProvidedActionProperties() *ActionProperties
	// The callback invoked when this Action is added to a Pipeline.
	Bind(scope constructs.Construct, stage IStage, options *ActionBindOptions) *ActionConfig
	// This is a renamed version of the `IAction.bind` method.
	Bound(scope constructs.Construct, stage IStage, options *ActionBindOptions) *ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for Action
type jsiiProxy_Action struct {
	jsiiProxy_IAction
}

func (j *jsiiProxy_Action) ActionProperties() *ActionProperties {
	var returns *ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Action) ProvidedActionProperties() *ActionProperties {
	var returns *ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewAction_Override(a Action) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.Action",
		nil, // no parameters
		a,
	)
}

func (a *jsiiProxy_Action) Bind(scope constructs.Construct, stage IStage, options *ActionBindOptions) *ActionConfig {
	if err := a.validateBindParameters(scope, stage, options); err != nil {
		panic(err)
	}
	var returns *ActionConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Action) Bound(scope constructs.Construct, stage IStage, options *ActionBindOptions) *ActionConfig {
	if err := a.validateBoundParameters(scope, stage, options); err != nil {
		panic(err)
	}
	var returns *ActionConfig

	_jsii_.Invoke(
		a,
		"bound",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Action) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	if err := a.validateOnStateChangeParameters(name, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		a,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Action) VariableExpression(variableName *string) *string {
	if err := a.validateVariableExpressionParameters(variableName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

