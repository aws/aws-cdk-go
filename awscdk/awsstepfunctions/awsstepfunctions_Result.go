package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The result of a Pass operation.
//
// Example:
//   // Makes the current JSON state { ..., "subObject": { "hello": "world" } }
//   pass := sfn.NewPass(this, jsii.String("Add Hello World"), &passProps{
//   	result: sfn.result.fromObject(map[string]interface{}{
//   		"hello": jsii.String("world"),
//   	}),
//   	resultPath: jsii.String("$.subObject"),
//   })
//
//   // Set the next state
//   nextState := sfn.NewPass(this, jsii.String("NextState"))
//   pass.next(nextState)
//
// Experimental.
type Result interface {
	// result of the Pass operation.
	// Experimental.
	Value() interface{}
}

// The jsii proxy struct for Result
type jsiiProxy_Result struct {
	_ byte // padding
}

func (j *jsiiProxy_Result) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func NewResult(value interface{}) Result {
	_init_.Initialize()

	j := jsiiProxy_Result{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions.Result",
		[]interface{}{value},
		&j,
	)

	return &j
}

// Experimental.
func NewResult_Override(r Result, value interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions.Result",
		[]interface{}{value},
		r,
	)
}

// The result of the operation is an array.
// Experimental.
func Result_FromArray(value *[]interface{}) Result {
	_init_.Initialize()

	var returns Result

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Result",
		"fromArray",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// The result of the operation is a boolean.
// Experimental.
func Result_FromBoolean(value *bool) Result {
	_init_.Initialize()

	var returns Result

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Result",
		"fromBoolean",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// The result of the operation is a number.
// Experimental.
func Result_FromNumber(value *float64) Result {
	_init_.Initialize()

	var returns Result

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Result",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// The result of the operation is an object.
// Experimental.
func Result_FromObject(value *map[string]interface{}) Result {
	_init_.Initialize()

	var returns Result

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Result",
		"fromObject",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// The result of the operation is a string.
// Experimental.
func Result_FromString(value *string) Result {
	_init_.Initialize()

	var returns Result

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Result",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

