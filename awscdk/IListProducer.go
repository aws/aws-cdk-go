// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for lazy list producers.
type IListProducer interface {
	// Produce the list value.
	Produce(context IResolveContext) *[]*string
}

// The jsii proxy for IListProducer
type jsiiProxy_IListProducer struct {
	_ byte // padding
}

func (i *jsiiProxy_IListProducer) Produce(context IResolveContext) *[]*string {
	if err := i.validateProduceParameters(context); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"produce",
		[]interface{}{context},
		&returns,
	)

	return returns
}

