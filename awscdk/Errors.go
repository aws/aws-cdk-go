package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Helper to check if an error is of a certain type.
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

// Test whether the given error is a AssertionError.
//
// An AssertionError is thrown when an assertion fails.
func Errors_IsAssertionError(x interface{}) *bool {
	_init_.Initialize()

	if err := validateErrors_IsAssertionErrorParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Errors",
		"isAssertionError",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Test whether the given error is an AssumptionError.
//
// An AssumptionError is thrown when a construct made an assumption somewhere that doesn't hold true.
// This error always indicates a bug in the construct.
func Errors_IsAssumptionError(x interface{}) *bool {
	_init_.Initialize()

	if err := validateErrors_IsAssumptionErrorParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Errors",
		"isAssumptionError",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Test whether the given error is a CloudAssemblyError.
//
// A CloudAssemblyError is thrown for unexpected problems with the synthesized assembly.
func Errors_IsCloudAssemblyError(x interface{}) *bool {
	_init_.Initialize()

	if err := validateErrors_IsCloudAssemblyErrorParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Errors",
		"isCloudAssemblyError",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Test whether the given errors is a ConstructionError.
//
// A ConstructionError is a generic error that will be thrown during the App construction or synthesis.
// To check for more specific errors, use the respective methods.
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

// Test whether the given error is an ExecutionError.
//
// An ExecutionError is thrown if an externally executed script or code failed.
func Errors_IsExecutionError(x interface{}) *bool {
	_init_.Initialize()

	if err := validateErrors_IsExecutionErrorParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Errors",
		"isExecutionError",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Test whether the given error is a ValidationError.
//
// A ValidationError is thrown when input props are failing to pass the rules of the construct.
// It usually means the underlying CloudFormation resource(s) would not deploy with a given configuration.
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

