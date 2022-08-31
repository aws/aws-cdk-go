package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Extract a field from the State Machine data that gets passed around between states.
// Deprecated: replaced by `JsonPath`.
type Data interface {
}

// The jsii proxy struct for Data
type jsiiProxy_Data struct {
	_ byte // padding
}

// Determines if the indicated string is an encoded JSON path.
// Deprecated: replaced by `JsonPath`.
func Data_IsJsonPathString(value *string) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Data",
		"isJsonPathString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Instead of using a literal string list, get the value from a JSON path.
// Deprecated: replaced by `JsonPath`.
func Data_ListAt(path *string) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Data",
		"listAt",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Instead of using a literal number, get the value from a JSON path.
// Deprecated: replaced by `JsonPath`.
func Data_NumberAt(path *string) *float64 {
	_init_.Initialize()

	var returns *float64

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Data",
		"numberAt",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Instead of using a literal string, get the value from a JSON path.
// Deprecated: replaced by `JsonPath`.
func Data_StringAt(path *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Data",
		"stringAt",
		[]interface{}{path},
		&returns,
	)

	return returns
}

func Data_EntirePayload() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_stepfunctions.Data",
		"entirePayload",
		&returns,
	)
	return returns
}

