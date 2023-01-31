package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a Mapping Value.
// Experimental.
type IMappingValue interface {
	// Represents a Mapping Value.
	// Experimental.
	Value() *string
}

// The jsii proxy for IMappingValue
type jsiiProxy_IMappingValue struct {
	_ byte // padding
}

func (j *jsiiProxy_IMappingValue) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

