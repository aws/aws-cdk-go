package awscodepipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// The abstract interface of a Pipeline Stage that is used by Actions.
type IStage interface {
	AddAction(action IAction)
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	// The actions belonging to this stage.
	Actions() *[]IAction
	Pipeline() IPipeline
	// The physical, human-readable name of this Pipeline Stage.
	StageName() *string
}

// The jsii proxy for IStage
type jsiiProxy_IStage struct {
	_ byte // padding
}

func (i *jsiiProxy_IStage) AddAction(action IAction) {
	if err := i.validateAddActionParameters(action); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addAction",
		[]interface{}{action},
	)
}

func (i *jsiiProxy_IStage) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
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

func (j *jsiiProxy_IStage) Actions() *[]IAction {
	var returns *[]IAction
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStage) Pipeline() IPipeline {
	var returns IPipeline
	_jsii_.Get(
		j,
		"pipeline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStage) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}

