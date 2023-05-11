package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The result of a Pass operation.
//
// Example:
//   // Makes the current JSON state { ..., "subObject": { "hello": "world" } }
//   pass := sfn.NewPass(this, jsii.String("Add Hello World"), &PassProps{
//   	Result: sfn.Result_FromObject(map[string]interface{}{
//   		"hello": jsii.String("world"),
//   	}),
//   	ResultPath: jsii.String("$.subObject"),
//   })
//
//   // Set the next state
//   nextState := sfn.NewPass(this, jsii.String("NextState"))
//   pass.Next(nextState)
//
type Result interface {
	// result of the Pass operation.
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


func NewResult(value interface{}) Result {
	_init_.Initialize()

	if err := validateNewResultParameters(value); err != nil {
		panic(err)
	}
	j := jsiiProxy_Result{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.Result",
		[]interface{}{value},
		&j,
	)

	return &j
}

func NewResult_Override(r Result, value interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.Result",
		[]interface{}{value},
		r,
	)
}

// The result of the operation is an array.
func Result_FromArray(value *[]interface{}) Result {
	_init_.Initialize()

	if err := validateResult_FromArrayParameters(value); err != nil {
		panic(err)
	}
	var returns Result

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Result",
		"fromArray",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// The result of the operation is a boolean.
func Result_FromBoolean(value *bool) Result {
	_init_.Initialize()

	if err := validateResult_FromBooleanParameters(value); err != nil {
		panic(err)
	}
	var returns Result

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Result",
		"fromBoolean",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// The result of the operation is a number.
func Result_FromNumber(value *float64) Result {
	_init_.Initialize()

	if err := validateResult_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns Result

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Result",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// The result of the operation is an object.
func Result_FromObject(value *map[string]interface{}) Result {
	_init_.Initialize()

	if err := validateResult_FromObjectParameters(value); err != nil {
		panic(err)
	}
	var returns Result

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Result",
		"fromObject",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// The result of the operation is a string.
func Result_FromString(value *string) Result {
	_init_.Initialize()

	if err := validateResult_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns Result

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Result",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

