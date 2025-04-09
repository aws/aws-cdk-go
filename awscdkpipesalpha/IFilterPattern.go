package awscdkpipesalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Filter events using an event pattern.
// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-event-filtering.html
//
// Experimental.
type IFilterPattern interface {
	// Stringified version of the filter pattern.
	// Experimental.
	Pattern() *string
	// Experimental.
	SetPattern(p *string)
}

// The jsii proxy for IFilterPattern
type jsiiProxy_IFilterPattern struct {
	_ byte // padding
}

func (j *jsiiProxy_IFilterPattern) Pattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFilterPattern)SetPattern(val *string) {
	if err := j.validateSetPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pattern",
		val,
	)
}

