// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for examining a construct and exposing metadata.
type IInspectable interface {
	// Examines construct.
	Inspect(inspector TreeInspector)
}

// The jsii proxy for IInspectable
type jsiiProxy_IInspectable struct {
	_ byte // padding
}

func (i *jsiiProxy_IInspectable) Inspect(inspector TreeInspector) {
	if err := i.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"inspect",
		[]interface{}{inspector},
	)
}

