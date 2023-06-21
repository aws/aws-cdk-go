package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Extract a field from the State Machine data or context that gets passed around between states.
//
// Example:
//   submitJobActivity := sfn.NewActivity(this, jsii.String("SubmitJob"))
//
//   tasks.NewStepFunctionsInvokeActivity(this, jsii.String("Submit Job"), &StepFunctionsInvokeActivityProps{
//   	Activity: submitJobActivity,
//   	Parameters: map[string]interface{}{
//   		"comment": jsii.String("Selecting what I care about."),
//   		"MyDetails": map[string]interface{}{
//   			"size": sfn.JsonPath_stringAt(jsii.String("$.product.details.size")),
//   			"exists": sfn.JsonPath_stringAt(jsii.String("$.product.availability")),
//   			"StaticValue": jsii.String("foo"),
//   		},
//   	},
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-paths.html
//
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
func JsonPath_Array(values ...*string) *string {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range values {
		args = append(args, a)
	}

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"array",
		args,
		&returns,
	)

	return returns
}

// Make an intrinsic States.ArrayContains expression.
//
// Use this function to determine if a specific value is present in an array. For example, you can use this function to detect if there was an error in a Map state iteration.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-intrinsic-functions.html
//
func JsonPath_ArrayContains(array interface{}, value interface{}) *string {
	_init_.Initialize()

	if err := validateJsonPath_ArrayContainsParameters(array, value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"arrayContains",
		[]interface{}{array, value},
		&returns,
	)

	return returns
}

// Make an intrinsic States.ArrayGetItem expression.
//
// Use this function to get a specified index's value in an array.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-intrinsic-functions.html
//
func JsonPath_ArrayGetItem(array interface{}, index *float64) *string {
	_init_.Initialize()

	if err := validateJsonPath_ArrayGetItemParameters(array, index); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"arrayGetItem",
		[]interface{}{array, index},
		&returns,
	)

	return returns
}

// Make an intrinsic States.ArrayLength expression.
//
// Use this function to get the length of an array.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-intrinsic-functions.html
//
func JsonPath_ArrayLength(array interface{}) *string {
	_init_.Initialize()

	if err := validateJsonPath_ArrayLengthParameters(array); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"arrayLength",
		[]interface{}{array},
		&returns,
	)

	return returns
}

// Make an intrinsic States.ArrayPartition expression.
//
// Use this function to partition a large array. You can also use this intrinsic to slice the data and then send the payload in smaller chunks.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-intrinsic-functions.html
//
func JsonPath_ArrayPartition(array interface{}, chunkSize *float64) *string {
	_init_.Initialize()

	if err := validateJsonPath_ArrayPartitionParameters(array, chunkSize); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"arrayPartition",
		[]interface{}{array, chunkSize},
		&returns,
	)

	return returns
}

// Make an intrinsic States.ArrayRange expression.
//
// Use this function to create a new array containing a specific range of elements. The new array can contain up to 1000 elements.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-intrinsic-functions.html
//
func JsonPath_ArrayRange(start *float64, end *float64, step *float64) *string {
	_init_.Initialize()

	if err := validateJsonPath_ArrayRangeParameters(start, end, step); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"arrayRange",
		[]interface{}{start, end, step},
		&returns,
	)

	return returns
}

// Make an intrinsic States.ArrayUnique expression.
//
// Use this function to get the length of an array.
// Use this function to remove duplicate values from an array and returns an array containing only unique elements. This function takes an array, which can be unsorted, as its sole argument.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-intrinsic-functions.html
//
func JsonPath_ArrayUnique(array interface{}) *string {
	_init_.Initialize()

	if err := validateJsonPath_ArrayUniqueParameters(array); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"arrayUnique",
		[]interface{}{array},
		&returns,
	)

	return returns
}

// Make an intrinsic States.Base64Decode expression.
//
// Use this function to decode data based on MIME Base64 decoding scheme. You can use this function to pass data to other AWS services without using a Lambda function.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-intrinsic-functions.html
//
func JsonPath_Base64Decode(base64 *string) *string {
	_init_.Initialize()

	if err := validateJsonPath_Base64DecodeParameters(base64); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"base64Decode",
		[]interface{}{base64},
		&returns,
	)

	return returns
}

// Make an intrinsic States.Base64Encode expression.
//
// Use this function to encode data based on MIME Base64 encoding scheme. You can use this function to pass data to other AWS services without using an AWS Lambda function.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-intrinsic-functions.html
//
func JsonPath_Base64Encode(input *string) *string {
	_init_.Initialize()

	if err := validateJsonPath_Base64EncodeParameters(input); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"base64Encode",
		[]interface{}{input},
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
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"format",
		args,
		&returns,
	)

	return returns
}

// Make an intrinsic States.Hash expression.
//
// Use this function to calculate the hash value of a given input. You can use this function to pass data to other AWS services without using a Lambda function.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-intrinsic-functions.html
//
func JsonPath_Hash(data interface{}, algorithm *string) *string {
	_init_.Initialize()

	if err := validateJsonPath_HashParameters(data, algorithm); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"hash",
		[]interface{}{data, algorithm},
		&returns,
	)

	return returns
}

// Determines if the indicated string is an encoded JSON path.
func JsonPath_IsEncodedJsonPath(value *string) *bool {
	_init_.Initialize()

	if err := validateJsonPath_IsEncodedJsonPathParameters(value); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"isEncodedJsonPath",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Make an intrinsic States.JsonMerge expression.
//
// Use this function to merge two JSON objects into a single object.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-intrinsic-functions.html
//
func JsonPath_JsonMerge(value1 interface{}, value2 interface{}) *string {
	_init_.Initialize()

	if err := validateJsonPath_JsonMergeParameters(value1, value2); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"jsonMerge",
		[]interface{}{value1, value2},
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
func JsonPath_JsonToString(value interface{}) *string {
	_init_.Initialize()

	if err := validateJsonPath_JsonToStringParameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"jsonToString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Instead of using a literal string list, get the value from a JSON path.
func JsonPath_ListAt(path *string) *[]*string {
	_init_.Initialize()

	if err := validateJsonPath_ListAtParameters(path); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"listAt",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Make an intrinsic States.MathAdd expression.
//
// Use this function to return the sum of two numbers. For example, you can use this function to increment values inside a loop without invoking a Lambda function.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-intrinsic-functions.html
//
func JsonPath_MathAdd(num1 *float64, num2 *float64) *string {
	_init_.Initialize()

	if err := validateJsonPath_MathAddParameters(num1, num2); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"mathAdd",
		[]interface{}{num1, num2},
		&returns,
	)

	return returns
}

// Make an intrinsic States.MathRandom expression.
//
// Use this function to return a random number between the specified start and end number. For example, you can use this function to distribute a specific task between two or more resources.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-intrinsic-functions.html
//
func JsonPath_MathRandom(start *float64, end *float64) *string {
	_init_.Initialize()

	if err := validateJsonPath_MathRandomParameters(start, end); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"mathRandom",
		[]interface{}{start, end},
		&returns,
	)

	return returns
}

// Instead of using a literal number, get the value from a JSON path.
func JsonPath_NumberAt(path *string) *float64 {
	_init_.Initialize()

	if err := validateJsonPath_NumberAtParameters(path); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"numberAt",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Reference a complete (complex) object in a JSON path location.
func JsonPath_ObjectAt(path *string) awscdk.IResolvable {
	_init_.Initialize()

	if err := validateJsonPath_ObjectAtParameters(path); err != nil {
		panic(err)
	}
	var returns awscdk.IResolvable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"objectAt",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Instead of using a literal string, get the value from a JSON path.
func JsonPath_StringAt(path *string) *string {
	_init_.Initialize()

	if err := validateJsonPath_StringAtParameters(path); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"stringAt",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Make an intrinsic States.StringSplit expression.
//
// Use this function to split a string into an array of values. This function takes two arguments.The first argument is a string and the second argument is the delimiting character that the function will use to divide the string.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-intrinsic-functions.html
//
func JsonPath_StringSplit(inputString *string, splitter *string) *string {
	_init_.Initialize()

	if err := validateJsonPath_StringSplitParameters(inputString, splitter); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"stringSplit",
		[]interface{}{inputString, splitter},
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
func JsonPath_StringToJson(jsonString *string) awscdk.IResolvable {
	_init_.Initialize()

	if err := validateJsonPath_StringToJsonParameters(jsonString); err != nil {
		panic(err)
	}
	var returns awscdk.IResolvable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"stringToJson",
		[]interface{}{jsonString},
		&returns,
	)

	return returns
}

// Make an intrinsic States.UUID expression.
//
// Use this function to return a version 4 universally unique identifier (v4 UUID) generated using random numbers. For example, you can use this function to call other AWS services or resources that need a UUID parameter or insert items in a DynamoDB table.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-intrinsic-functions.html
//
func JsonPath_Uuid() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"uuid",
		nil, // no parameters
		&returns,
	)

	return returns
}

func JsonPath_DISCARD() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"DISCARD",
		&returns,
	)
	return returns
}

func JsonPath_EntireContext() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"entireContext",
		&returns,
	)
	return returns
}

func JsonPath_EntirePayload() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"entirePayload",
		&returns,
	)
	return returns
}

func JsonPath_ExecutionId() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"executionId",
		&returns,
	)
	return returns
}

func JsonPath_ExecutionInput() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"executionInput",
		&returns,
	)
	return returns
}

func JsonPath_ExecutionName() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"executionName",
		&returns,
	)
	return returns
}

func JsonPath_ExecutionRoleArn() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"executionRoleArn",
		&returns,
	)
	return returns
}

func JsonPath_ExecutionStartTime() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"executionStartTime",
		&returns,
	)
	return returns
}

func JsonPath_StateEnteredTime() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"stateEnteredTime",
		&returns,
	)
	return returns
}

func JsonPath_StateMachineId() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"stateMachineId",
		&returns,
	)
	return returns
}

func JsonPath_StateMachineName() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"stateMachineName",
		&returns,
	)
	return returns
}

func JsonPath_StateName() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"stateName",
		&returns,
	)
	return returns
}

func JsonPath_StateRetryCount() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"stateRetryCount",
		&returns,
	)
	return returns
}

func JsonPath_TaskToken() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.JsonPath",
		"taskToken",
		&returns,
	)
	return returns
}

