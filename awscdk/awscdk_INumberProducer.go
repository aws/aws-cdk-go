// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for lazy number producers.
type INumberProducer interface {
	// Produce the number value.
	Produce(context IResolveContext) *float64
}

// The jsii proxy for INumberProducer
type jsiiProxy_INumberProducer struct {
	_ byte // padding
}

func (i *jsiiProxy_INumberProducer) Produce(context IResolveContext) *float64 {
	if err := i.validateProduceParameters(context); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"produce",
		[]interface{}{context},
		&returns,
	)

	return returns
}

