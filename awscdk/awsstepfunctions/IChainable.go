package awsstepfunctions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for objects that can be used in a Chain.
type IChainable interface {
	// The chainable end state(s) of this chainable.
	EndStates() *[]INextable
	// Descriptive identifier for this chainable.
	Id() *string
	// The start state of this chainable.
	StartState() State
}

// The jsii proxy for IChainable
type jsiiProxy_IChainable struct {
	_ byte // padding
}

func (j *jsiiProxy_IChainable) EndStates() *[]INextable {
	var returns *[]INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChainable) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChainable) StartState() State {
	var returns State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

