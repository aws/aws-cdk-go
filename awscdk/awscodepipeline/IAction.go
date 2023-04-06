package awscodepipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Pipeline Action.
//
// If you want to implement this interface,
// consider extending the `Action` class,
// which contains some common logic.
type IAction interface {
	// The callback invoked when this Action is added to a Pipeline.
	Bind(scope constructs.Construct, stage IStage, options *ActionBindOptions) *ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the `bind` callback.
	ActionProperties() *ActionProperties
}

// The jsii proxy for IAction
type jsiiProxy_IAction struct {
	_ byte // padding
}

func (i *jsiiProxy_IAction) Bind(scope constructs.Construct, stage IStage, options *ActionBindOptions) *ActionConfig {
	if err := i.validateBindParameters(scope, stage, options); err != nil {
		panic(err)
	}
	var returns *ActionConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	if err := i.validateOnStateChangeParameters(name, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IAction) ActionProperties() *ActionProperties {
	var returns *ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

