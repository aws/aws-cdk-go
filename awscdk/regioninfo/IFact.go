package regioninfo

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A fact that can be registered about a particular region.
type IFact interface {
	// The name of this fact.
	//
	// Standardized values are provided by the `Facts` class.
	Name() *string
	// The region for which this fact applies.
	Region() *string
	// The value of this fact.
	Value() *string
}

// The jsii proxy for IFact
type jsiiProxy_IFact struct {
	_ byte // padding
}

func (j *jsiiProxy_IFact) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFact) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFact) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

