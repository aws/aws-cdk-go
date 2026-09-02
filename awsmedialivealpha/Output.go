package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents an output within an output group.
//
// Each output group type has its own Output subclass that knows how to render its output settings.
// Experimental.
type Output interface {
	// Experimental.
	OutputName() *string
}

// The jsii proxy struct for Output
type jsiiProxy_Output struct {
	_ byte // padding
}

func (j *jsiiProxy_Output) OutputName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputName",
		&returns,
	)
	return returns
}


// Experimental.
func NewOutput_Override(o Output, outputName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-medialive-alpha.Output",
		[]interface{}{outputName},
		o,
	)
}

