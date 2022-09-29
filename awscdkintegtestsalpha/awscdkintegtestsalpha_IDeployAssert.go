// CDK Integration Testing Constructs
package awscdkintegtestsalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface that allows for registering a list of assertions that should be performed on a construct.
//
// This is only necessary
// when writing integration tests.
// Experimental.
type IDeployAssert interface {
	// Query AWS using JavaScript SDK V2 API calls.
	//
	// This can be used to either
	// trigger an action or to return a result that can then be asserted against
	// an expected value.
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
	AwsApiCall(service *string, api *string, parameters interface{}) IApiCall
	// Assert that the ExpectedResult is equal to the ActualResult.
	//
	// Example:
	//   var integ integTest
	//   var apiCall awsApiCall
	//
	//   integ.assertions.expect(jsii.String("invoke"), awscdkintegtestsalpha.ExpectedResult.objectLike(map[string]interface{}{
	//   	"Payload": jsii.String("OK"),
	//   }), awscdkintegtestsalpha.ActualResult.fromAwsApiCall(apiCall, jsii.String("Body")))
	//
	// Experimental.
	Expect(id *string, expected ExpectedResult, actual ActualResult)
	// Invoke a lambda function and return the response which can be asserted.
	//
	// Example:
	//   var app app
	//   var integ integTest
	//
	//   invoke := integ.assertions.invokeFunction(&lambdaInvokeFunctionProps{
	//   	functionName: jsii.String("my-function"),
	//   })
	//   invoke.expect(awscdkintegtestsalpha.ExpectedResult.objectLike(map[string]interface{}{
	//   	"Payload": jsii.String("200"),
	//   }))
	//
	// Experimental.
	InvokeFunction(props *LambdaInvokeFunctionProps) IApiCall
}

// The jsii proxy for IDeployAssert
type jsiiProxy_IDeployAssert struct {
	_ byte // padding
}

func (i *jsiiProxy_IDeployAssert) AwsApiCall(service *string, api *string, parameters interface{}) IApiCall {
	if err := i.validateAwsApiCallParameters(service, api); err != nil {
		panic(err)
	}
	var returns IApiCall

	_jsii_.Invoke(
		i,
		"awsApiCall",
		[]interface{}{service, api, parameters},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDeployAssert) Expect(id *string, expected ExpectedResult, actual ActualResult) {
	if err := i.validateExpectParameters(id, expected, actual); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"expect",
		[]interface{}{id, expected, actual},
	)
}

func (i *jsiiProxy_IDeployAssert) InvokeFunction(props *LambdaInvokeFunctionProps) IApiCall {
	if err := i.validateInvokeFunctionParameters(props); err != nil {
		panic(err)
	}
	var returns IApiCall

	_jsii_.Invoke(
		i,
		"invokeFunction",
		[]interface{}{props},
		&returns,
	)

	return returns
}

