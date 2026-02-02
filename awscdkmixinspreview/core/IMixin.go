package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// A mixin is a reusable piece of functionality that can be applied to constructs to add behavior, properties, or modify existing functionality without inheritance.
// Experimental.
type IMixin interface {
	// Applies the mixin functionality to the target construct.
	// Experimental.
	ApplyTo(construct constructs.IConstruct)
	// Determines whether this mixin can be applied to the given construct.
	// Experimental.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy for IMixin
type jsiiProxy_IMixin struct {
	_ byte // padding
}

func (i *jsiiProxy_IMixin) ApplyTo(construct constructs.IConstruct) {
	if err := i.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyTo",
		[]interface{}{construct},
	)
}

func (i *jsiiProxy_IMixin) Supports(construct constructs.IConstruct) *bool {
	if err := i.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		i,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

