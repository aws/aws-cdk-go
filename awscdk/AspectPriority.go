package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Default Priority values for Aspects.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   aspectPriority := cdk.NewAspectPriority()
//
type AspectPriority interface {
}

// The jsii proxy struct for AspectPriority
type jsiiProxy_AspectPriority struct {
	_ byte // padding
}

func NewAspectPriority() AspectPriority {
	_init_.Initialize()

	j := jsiiProxy_AspectPriority{}

	_jsii_.Create(
		"aws-cdk-lib.AspectPriority",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewAspectPriority_Override(a AspectPriority) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.AspectPriority",
		nil, // no parameters
		a,
	)
}

func AspectPriority_DEFAULT() *float64 {
	_init_.Initialize()
	var returns *float64
	_jsii_.StaticGet(
		"aws-cdk-lib.AspectPriority",
		"DEFAULT",
		&returns,
	)
	return returns
}

func AspectPriority_MUTATING() *float64 {
	_init_.Initialize()
	var returns *float64
	_jsii_.StaticGet(
		"aws-cdk-lib.AspectPriority",
		"MUTATING",
		&returns,
	)
	return returns
}

func AspectPriority_READONLY() *float64 {
	_init_.Initialize()
	var returns *float64
	_jsii_.StaticGet(
		"aws-cdk-lib.AspectPriority",
		"READONLY",
		&returns,
	)
	return returns
}

