// CDK Integration Testing Constructs
package awscdkintegtestsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkintegtestsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents the "expected" results to compare.
//
// Example:
//   var app app
//   var integ integTest
//
//   integ.assertions.awsApiCall(jsii.String("SQS"), jsii.String("sendMessage"), map[string]*string{
//   	"QueueUrl": jsii.String("url"),
//   	"MessageBody": jsii.String("hello"),
//   })
//   message := integ.assertions.awsApiCall(jsii.String("SQS"), jsii.String("receiveMessage"), map[string]*string{
//   	"QueueUrl": jsii.String("url"),
//   })
//   message.expect(awscdkintegtestsalpha.ExpectedResult.objectLike(map[string]interface{}{
//   	"Messages": []interface{}{
//   		map[string]*string{
//   			"Body": jsii.String("hello"),
//   		},
//   	},
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
//   awscdkintegtestsalpha.ExpectedResult.arrayWith([]interface{}{
//   	map[string]*string{
//   		"stringParam": jsii.String("hello"),
//   	},
//   })
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
//   awscdkintegtestsalpha.ExpectedResult.exact(map[string]interface{}{
//   	"stringParam": jsii.String("hello"),
//   	"numberParam": jsii.Number(3),
//   	"booleanParam": jsii.Boolean(true),
//   })
//
//   // fail
//   awscdkintegtestsalpha.ExpectedResult.exact(map[string]*string{
//   	"stringParam": jsii.String("hello"),
//   })
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
//   }
//   // pass
//   awscdkintegtestsalpha.ExpectedResult.objectLike(map[string]interface{}{
//   	"stringParam": jsii.String("hello"),
//   })
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
//   awscdkintegtestsalpha.ExpectedResult.stringLikeRegexp(jsii.String("value"))
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

