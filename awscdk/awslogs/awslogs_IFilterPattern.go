package awslogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for objects that can render themselves to log patterns.
// Experimental.
type IFilterPattern interface {
	// Experimental.
	LogPatternString() *string
}

// The jsii proxy for IFilterPattern
type jsiiProxy_IFilterPattern struct {
	_ byte // padding
}

func (j *jsiiProxy_IFilterPattern) LogPatternString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logPatternString",
		&returns,
	)
	return returns
}

