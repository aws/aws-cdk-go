package integtests

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// Represents the "actual" results to compare.
//
// Example:
//   var myCustomResource customResource
//   var stack stack
//   var app app
//
//
//   integ := awscdk.NewIntegTest(app, jsii.String("Integ"), &IntegTestProps{
//   	TestCases: []*stack{
//   		stack,
//   	},
//   })
//   integ.Assertions.Expect(jsii.String("CustomAssertion"), awscdk.ExpectedResult_ObjectLike(map[string]interface{}{
//   	"foo": jsii.String("bar"),
//   }), awscdk.ActualResult_FromCustomResource(myCustomResource, jsii.String("data")))
//
// Experimental.
type ActualResult interface {
	// The actual results as a string.
	// Experimental.
	Result() *string
	// Experimental.
	SetResult(val *string)
}

// The jsii proxy struct for ActualResult
type jsiiProxy_ActualResult struct {
	_ byte // padding
}

func (j *jsiiProxy_ActualResult) Result() *string {
	var returns *string
	_jsii_.Get(
		j,
		"result",
		&returns,
	)
	return returns
}


// Experimental.
func NewActualResult_Override(a ActualResult) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.integ_tests.ActualResult",
		nil, // no parameters
		a,
	)
}

func (j *jsiiProxy_ActualResult)SetResult(val *string) {
	if err := j.validateSetResultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"result",
		val,
	)
}

// Get the actual results from a AwsApiCall.
// Experimental.
func ActualResult_FromAwsApiCall(query IAwsApiCall, attribute *string) ActualResult {
	_init_.Initialize()

	if err := validateActualResult_FromAwsApiCallParameters(query, attribute); err != nil {
		panic(err)
	}
	var returns ActualResult

	_jsii_.StaticInvoke(
		"monocdk.integ_tests.ActualResult",
		"fromAwsApiCall",
		[]interface{}{query, attribute},
		&returns,
	)

	return returns
}

// Get the actual results from a CustomResource.
// Experimental.
func ActualResult_FromCustomResource(customResource awscdk.CustomResource, attribute *string) ActualResult {
	_init_.Initialize()

	if err := validateActualResult_FromCustomResourceParameters(customResource, attribute); err != nil {
		panic(err)
	}
	var returns ActualResult

	_jsii_.StaticInvoke(
		"monocdk.integ_tests.ActualResult",
		"fromCustomResource",
		[]interface{}{customResource, attribute},
		&returns,
	)

	return returns
}

