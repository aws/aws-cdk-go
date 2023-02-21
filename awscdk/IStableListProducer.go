// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for (stable) lazy list producers.
type IStableListProducer interface {
	// Produce the list value.
	Produce() *[]*string
}

// The jsii proxy for IStableListProducer
type jsiiProxy_IStableListProducer struct {
	_ byte // padding
}

func (i *jsiiProxy_IStableListProducer) Produce() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"produce",
		nil, // no parameters
		&returns,
	)

	return returns
}

