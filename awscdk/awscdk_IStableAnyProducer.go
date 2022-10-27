// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for (stable) lazy untyped value producers.
type IStableAnyProducer interface {
	// Produce the value.
	Produce() interface{}
}

// The jsii proxy for IStableAnyProducer
type jsiiProxy_IStableAnyProducer struct {
	_ byte // padding
}

func (i *jsiiProxy_IStableAnyProducer) Produce() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"produce",
		nil, // no parameters
		&returns,
	)

	return returns
}

