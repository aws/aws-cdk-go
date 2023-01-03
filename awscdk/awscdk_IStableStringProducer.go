// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for (stable) lazy string producers.
// Experimental.
type IStableStringProducer interface {
	// Produce the string value.
	// Experimental.
	Produce() *string
}

// The jsii proxy for IStableStringProducer
type jsiiProxy_IStableStringProducer struct {
	_ byte // padding
}

func (i *jsiiProxy_IStableStringProducer) Produce() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"produce",
		nil, // no parameters
		&returns,
	)

	return returns
}

