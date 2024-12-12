package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Helper to check if an error is a SynthesisErrors.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   errors := cdk.NewErrors()
//
type Errors interface {
}

// The jsii proxy struct for Errors
type jsiiProxy_Errors struct {
	_ byte // padding
}

func NewErrors() Errors {
	_init_.Initialize()

	j := jsiiProxy_Errors{}

	_jsii_.Create(
		"aws-cdk-lib.Errors",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewErrors_Override(e Errors) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.Errors",
		nil, // no parameters
		e,
	)
}

// Test whether the given errors is a ConstructionError.
func Errors_IsConstructError(x interface{}) *bool {
	_init_.Initialize()

	if err := validateErrors_IsConstructErrorParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Errors",
		"isConstructError",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Test whether the given error is a ValidationError.
func Errors_IsValidationError(x interface{}) *bool {
	_init_.Initialize()

	if err := validateErrors_IsValidationErrorParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Errors",
		"isValidationError",
		[]interface{}{x},
		&returns,
	)

	return returns
}

