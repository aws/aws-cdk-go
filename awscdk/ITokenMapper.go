// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface to apply operation to tokens in a string.
//
// Interface so it can be exported via jsii.
type ITokenMapper interface {
	// Replace a single token.
	MapToken(t IResolvable) interface{}
}

// The jsii proxy for ITokenMapper
type jsiiProxy_ITokenMapper struct {
	_ byte // padding
}

func (i *jsiiProxy_ITokenMapper) MapToken(t IResolvable) interface{} {
	if err := i.validateMapTokenParameters(t); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"mapToken",
		[]interface{}{t},
		&returns,
	)

	return returns
}

