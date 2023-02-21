package awsssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A StringList SSM Parameter.
type IStringListParameter interface {
	IParameter
	// The parameter value.
	//
	// Value must not nest another parameter. Do not use {{}} in the value. Values in the array
	// cannot contain commas (``,``).
	StringListValue() *[]*string
}

// The jsii proxy for IStringListParameter
type jsiiProxy_IStringListParameter struct {
	jsiiProxy_IParameter
}

func (j *jsiiProxy_IStringListParameter) StringListValue() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stringListValue",
		&returns,
	)
	return returns
}

