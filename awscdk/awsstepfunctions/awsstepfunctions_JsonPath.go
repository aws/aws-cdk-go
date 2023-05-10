package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// Extract a field from the State Machine data or context that gets passed around between states.
//
// Example:
//   var fn function
//
//   tasks.NewLambdaInvoke(this, jsii.String("Invoke Handler"), &LambdaInvokeProps{
//   	LambdaFunction: fn,
//   	ResultSelector: map[string]interface{}{
//   		"lambdaOutput": sfn.JsonPath_stringAt(jsii.String("$.Payload")),
//   		"invokeRequestId": sfn.JsonPath_stringAt(jsii.String("$.SdkResponseMetadata.RequestId")),
//   		"staticValue": map[string]*string{
//   			"foo": jsii.String("bar"),
//   		},
//   		"stateName": sfn.JsonPath_stringAt(jsii.String("$.State.Name")),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-paths.html
//
// Experimental.
type JsonPath interface {
}

// The jsii proxy struct for JsonPath
type jsiiProxy_JsonPath struct {
	_ byte // padding
}

// Make an intrinsic States.Array expression.
//
// Combine any number of string literals or JsonPath expressions into an array.
//
// Use this function if the value of an array element directly has to come
// from a JSON Path expression (either the State object or the Context object).
//
// If the array contains object literals whose values come from a JSON path
// expression, you do not need to use this function.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-intrinsic-functions.html
//
// Experimental.
func JsonPath_Array(values ...*string) *string {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range values {
		args = append(args, a)
	}

	var returns *string

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.JsonPath",
		"array",
		args,
		&returns,
	)

	return returns
}

// Make an intrinsic States.Format expression.
//
// This can be used to embed JSON Path variables inside a format string.
//
// For example:
//
// ```ts
// sfn.JsonPath.format('Hello, my name is {}.', sfn.JsonPath.stringAt('$.name'))
// ```.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-intrinsic-functions.html
//
// Experimental.
func JsonPath_Format(formatString *string, values ...*string) *string {
	_init_.Initialize()

	if err := validateJsonPath_FormatParameters(formatString); err != nil {
		panic(err)
	}
	args := []interface{}{formatString}
	for _, a := range values {
		args = append(args, a)
	}

	var returns *string

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.JsonPath",
		"format",
		args,
		&returns,
	)

	return returns
}

// Determines if the indicated string is an encoded JSON path.
// Experimental.
func JsonPath_IsEncodedJsonPath(value *string) *bool {
	_init_.Initialize()

	if err := validateJsonPath_IsEncodedJsonPathParameters(value); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.JsonPath",
		"isEncodedJsonPath",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Make an intrinsic States.JsonToString expression.
//
// During the execution of the Step Functions state machine, encode the
// given object into a JSON string.
//
// For example:
//
// ```ts
// sfn.JsonPath.jsonToString(sfn.JsonPath.objectAt('$.someObject'))
// ```.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-intrinsic-functions.html
//
// Experimental.
func JsonPath_JsonToString(value interface{}) *string {
	_init_.Initialize()

	if err := validateJsonPath_JsonToStringParameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.JsonPath",
		"jsonToString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Instead of using a literal string list, get the value from a JSON path.
// Experimental.
func JsonPath_ListAt(path *string) *[]*string {
	_init_.Initialize()

	if err := validateJsonPath_ListAtParameters(path); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.JsonPath",
		"listAt",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Instead of using a literal number, get the value from a JSON path.
// Experimental.
func JsonPath_NumberAt(path *string) *float64 {
	_init_.Initialize()

	if err := validateJsonPath_NumberAtParameters(path); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.JsonPath",
		"numberAt",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Reference a complete (complex) object in a JSON path location.
// Experimental.
func JsonPath_ObjectAt(path *string) awscdk.IResolvable {
	_init_.Initialize()

	if err := validateJsonPath_ObjectAtParameters(path); err != nil {
		panic(err)
	}
	var returns awscdk.IResolvable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.JsonPath",
		"objectAt",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Instead of using a literal string, get the value from a JSON path.
// Experimental.
func JsonPath_StringAt(path *string) *string {
	_init_.Initialize()

	if err := validateJsonPath_StringAtParameters(path); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.JsonPath",
		"stringAt",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Make an intrinsic States.StringToJson expression.
//
// During the execution of the Step Functions state machine, parse the given
// argument as JSON into its object form.
//
// For example:
//
// ```ts
// sfn.JsonPath.stringToJson(sfn.JsonPath.stringAt('$.someJsonBody'))
// ```.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-intrinsic-functions.html
//
// Experimental.
func JsonPath_StringToJson(jsonString *string) awscdk.IResolvable {
	_init_.Initialize()

	if err := validateJsonPath_StringToJsonParameters(jsonString); err != nil {
		panic(err)
	}
	var returns awscdk.IResolvable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.JsonPath",
		"stringToJson",
		[]interface{}{jsonString},
		&returns,
	)

	return returns
}

func JsonPath_DISCARD() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_stepfunctions.JsonPath",
		"DISCARD",
		&returns,
	)
	return returns
}

func JsonPath_EntireContext() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_stepfunctions.JsonPath",
		"entireContext",
		&returns,
	)
	return returns
}

func JsonPath_EntirePayload() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_stepfunctions.JsonPath",
		"entirePayload",
		&returns,
	)
	return returns
}

func JsonPath_TaskToken() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_stepfunctions.JsonPath",
		"taskToken",
		&returns,
	)
	return returns
}

