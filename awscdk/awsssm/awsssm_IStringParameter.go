package awsssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A String SSM Parameter.
type IStringParameter interface {
	IParameter
	// The parameter value.
	//
	// Value must not nest another parameter. Do not use {{}} in the value.
	StringValue() *string
}

// The jsii proxy for IStringParameter
type jsiiProxy_IStringParameter struct {
	jsiiProxy_IParameter
}

func (j *jsiiProxy_IStringParameter) StringValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValue",
		&returns,
	)
	return returns
}

