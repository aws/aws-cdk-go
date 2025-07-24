package awscloudwatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A single dashboard variable.
type IVariable interface {
	// Return the variable JSON for use in the dashboard.
	ToJson() interface{}
}

// The jsii proxy for IVariable
type jsiiProxy_IVariable struct {
	_ byte // padding
}

func (i *jsiiProxy_IVariable) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

