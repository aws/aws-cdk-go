package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An abstract class which represents an AWS Lambda event source.
type IEventSource interface {
	// Called by `lambda.addEventSource` to allow the event source to bind to this function.
	Bind(target IFunction)
}

// The jsii proxy for IEventSource
type jsiiProxy_IEventSource struct {
	_ byte // padding
}

func (i *jsiiProxy_IEventSource) Bind(target IFunction) {
	if err := i.validateBindParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"bind",
		[]interface{}{target},
	)
}

