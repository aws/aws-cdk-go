package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Extract a field from the State Machine Context data.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#wait-token-contextobject
//
// Deprecated: replaced by `JsonPath`.
type Context interface {
}

// The jsii proxy struct for Context
type jsiiProxy_Context struct {
	_ byte // padding
}

// Instead of using a literal number, get the value from a JSON path.
// Deprecated: replaced by `JsonPath`.
func Context_NumberAt(path *string) *float64 {
	_init_.Initialize()

	if err := validateContext_NumberAtParameters(path); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Context",
		"numberAt",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Instead of using a literal string, get the value from a JSON path.
// Deprecated: replaced by `JsonPath`.
func Context_StringAt(path *string) *string {
	_init_.Initialize()

	if err := validateContext_StringAtParameters(path); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Context",
		"stringAt",
		[]interface{}{path},
		&returns,
	)

	return returns
}

func Context_EntireContext() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_stepfunctions.Context",
		"entireContext",
		&returns,
	)
	return returns
}

func Context_TaskToken() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_stepfunctions.Context",
		"taskToken",
		&returns,
	)
	return returns
}

