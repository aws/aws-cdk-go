package awsstepfunctions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for states that can have 'next' states.
type INextable interface {
	// Go to the indicated state after this state.
	//
	// Returns: The chain of states built up.
	Next(state IChainable) Chain
}

// The jsii proxy for INextable
type jsiiProxy_INextable struct {
	_ byte // padding
}

func (i *jsiiProxy_INextable) Next(state IChainable) Chain {
	if err := i.validateNextParameters(state); err != nil {
		panic(err)
	}
	var returns Chain

	_jsii_.Invoke(
		i,
		"next",
		[]interface{}{state},
		&returns,
	)

	return returns
}

