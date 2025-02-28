package awscdkintegtestsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkintegtestsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents the "expected" results to compare.
//
// Example:
//   var app app
//   var stack stack
//   var queue queue
//   var fn iFunction
//
//
//   integ := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("Integ"), &IntegTestProps{
//   	TestCases: []*stack{
//   		stack,
//   	},
//   })
//
//   integ.Assertions.InvokeFunction(&LambdaInvokeFunctionProps{
//   	FunctionName: fn.FunctionName,
//   	InvocationType: awscdkintegtestsalpha.InvocationType_EVENT,
//   	Payload: jSON.stringify(map[string]*string{
//   		"status": jsii.String("OK"),
//   	}),
//   })
//
//   message := integ.Assertions.AwsApiCall(jsii.String("SQS"), jsii.String("receiveMessage"), map[string]interface{}{
//   	"QueueUrl": queue.queueUrl,
//   	"WaitTimeSeconds": jsii.Number(20),
//   })
//
//   message.AssertAtPath(jsii.String("Messages.0.Body"), awscdkintegtestsalpha.ExpectedResult_ObjectLike(map[string]interface{}{
//   	"requestContext": map[string]*string{
//   		"condition": jsii.String("Success"),
//   	},
//   	"requestPayload": map[string]*string{
//   		"status": jsii.String("OK"),
//   	},
//   	"responseContext": map[string]*f64{
//   		"statusCode": jsii.Number(200),
//   	},
//   	"responsePayload": jsii.String("success"),
//   }))
//
// Experimental.
type ExpectedResult interface {
	// The expected results encoded as a string.
	// Experimental.
	Result() *string
	// Experimental.
	SetResult(val *string)
}

// The jsii proxy struct for ExpectedResult
type jsiiProxy_ExpectedResult struct {
	_ byte // padding
}

func (j *jsiiProxy_ExpectedResult) Result() *string {
	var returns *string
	_jsii_.Get(
		j,
		"result",
		&returns,
	)
	return returns
}


// Experimental.
func NewExpectedResult_Override(e ExpectedResult) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/integ-tests-alpha.ExpectedResult",
		nil, // no parameters
		e,
	)
}

func (j *jsiiProxy_ExpectedResult)SetResult(val *string) {
	if err := j.validateSetResultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"result",
		val,
	)
}

// The actual results must be a list and must contain an item with the expected results.
//
// Example:
//   // actual results
//   actual := []map[string]*string{
//   	map[string]*string{
//   		"stringParam": jsii.String("hello"),
//   	},
//   	map[string]*string{
//   		"stringParam": jsii.String("world"),
//   	},
//   }
//   // pass
//   awscdkintegtestsalpha.ExpectedResult_ArrayWith([]interface{}{
//   	map[string]*string{
//   		"stringParam": jsii.String("hello"),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.assertions.Match.html#static-arraywbrwithpattern
//
// Experimental.
func ExpectedResult_ArrayWith(expected *[]interface{}) ExpectedResult {
	_init_.Initialize()

	if err := validateExpectedResult_ArrayWithParameters(expected); err != nil {
		panic(err)
	}
	var returns ExpectedResult

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.ExpectedResult",
		"arrayWith",
		[]interface{}{expected},
		&returns,
	)

	return returns
}

// The actual results must match exactly.
//
// Missing data
// will result in a failure.
//
// Example:
//   // actual results
//   actual := map[string]interface{}{
//   	"stringParam": jsii.String("hello"),
//   	"numberParam": jsii.Number(3),
//   	"booleanParam": jsii.Boolean(true),
//   }
//   // pass
//   awscdkintegtestsalpha.ExpectedResult_Exact(map[string]interface{}{
//   	"stringParam": jsii.String("hello"),
//   	"numberParam": jsii.Number(3),
//   	"booleanParam": jsii.Boolean(true),
//   })
//
//   // fail
//   awscdkintegtestsalpha.ExpectedResult_Exact(map[string]*string{
//   	"stringParam": jsii.String("hello"),
//   })
//
// See: https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.assertions.Match.html#static-exactpattern
//
// Experimental.
func ExpectedResult_Exact(expected interface{}) ExpectedResult {
	_init_.Initialize()

	if err := validateExpectedResult_ExactParameters(expected); err != nil {
		panic(err)
	}
	var returns ExpectedResult

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.ExpectedResult",
		"exact",
		[]interface{}{expected},
		&returns,
	)

	return returns
}

// The expected results must be a subset of the actual results.
//
// Example:
//   // actual results
//   actual := map[string]interface{}{
//   	"stringParam": jsii.String("hello"),
//   	"numberParam": jsii.Number(3),
//   	"booleanParam": jsii.Boolean(true),
//   	"objectParam": map[string]*string{
//   		"prop1": jsii.String("value"),
//   		"prop2": jsii.String("value"),
//   	},
//   }
//   // pass
//   awscdkintegtestsalpha.ExpectedResult_ObjectLike(map[string]interface{}{
//   	"stringParam": jsii.String("hello"),
//   	"objectParam": map[string]*string{
//   		"prop1": jsii.String("value"),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.assertions.Match.html#static-objectwbrlikepattern
//
// Experimental.
func ExpectedResult_ObjectLike(expected *map[string]interface{}) ExpectedResult {
	_init_.Initialize()

	if err := validateExpectedResult_ObjectLikeParameters(expected); err != nil {
		panic(err)
	}
	var returns ExpectedResult

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.ExpectedResult",
		"objectLike",
		[]interface{}{expected},
		&returns,
	)

	return returns
}

// Actual results is a string that matches the Expected result regex.
//
// Example:
//   // actual results
//   actual := "some string value"
//
//   // pass
//   awscdkintegtestsalpha.ExpectedResult_StringLikeRegexp(jsii.String("value"))
//
// See: https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.assertions.Match.html#static-stringwbrlikewbrregexppattern
//
// Experimental.
func ExpectedResult_StringLikeRegexp(expected *string) ExpectedResult {
	_init_.Initialize()

	if err := validateExpectedResult_StringLikeRegexpParameters(expected); err != nil {
		panic(err)
	}
	var returns ExpectedResult

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.ExpectedResult",
		"stringLikeRegexp",
		[]interface{}{expected},
		&returns,
	)

	return returns
}

