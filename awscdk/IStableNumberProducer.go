// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for (stable) lazy number producers.
type IStableNumberProducer interface {
	// Produce the number value.
	Produce() *float64
}

// The jsii proxy for IStableNumberProducer
type jsiiProxy_IStableNumberProducer struct {
	_ byte // padding
}

func (i *jsiiProxy_IStableNumberProducer) Produce() *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"produce",
		nil, // no parameters
		&returns,
	)

	return returns
}

