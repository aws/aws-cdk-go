package awscdkappconfigalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IValidator interface {
	// The content of the validator.
	// Experimental.
	Content() *string
	// The type of validator.
	// Experimental.
	Type() ValidatorType
}

// The jsii proxy for IValidator
type jsiiProxy_IValidator struct {
	_ byte // padding
}

func (j *jsiiProxy_IValidator) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IValidator) Type() ValidatorType {
	var returns ValidatorType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

