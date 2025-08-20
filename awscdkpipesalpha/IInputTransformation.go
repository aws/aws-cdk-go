package awscdkpipesalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Transform or replace the input event payload.
// Experimental.
type IInputTransformation interface {
	// Bind the input transformation to the pipe and returns the inputTemplate string.
	// Experimental.
	Bind(pipe IPipe) *InputTransformationConfig
}

// The jsii proxy for IInputTransformation
type jsiiProxy_IInputTransformation struct {
	_ byte // padding
}

func (i *jsiiProxy_IInputTransformation) Bind(pipe IPipe) *InputTransformationConfig {
	if err := i.validateBindParameters(pipe); err != nil {
		panic(err)
	}
	var returns *InputTransformationConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{pipe},
		&returns,
	)

	return returns
}

