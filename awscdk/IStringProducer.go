// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for lazy string producers.
type IStringProducer interface {
	// Produce the string value.
	Produce(context IResolveContext) *string
}

// The jsii proxy for IStringProducer
type jsiiProxy_IStringProducer struct {
	_ byte // padding
}

func (i *jsiiProxy_IStringProducer) Produce(context IResolveContext) *string {
	if err := i.validateProduceParameters(context); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"produce",
		[]interface{}{context},
		&returns,
	)

	return returns
}

