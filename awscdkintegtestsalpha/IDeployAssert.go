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
	// Query AWS using JavaScript SDK API calls.
	//
	// This can be used to either
	// trigger an action or to return a result that can then be asserted against
	// an expected value
	//
	// The `service` is the name of an AWS service, in one of the following forms:
	// - An AWS SDK for JavaScript v3 package name (`@aws-sdk/client-api-gateway`)
	// - An AWS SDK for JavaScript v3 client name (`api-gateway`)
	// - An AWS SDK for JavaScript v2 constructor name (`APIGateway`)
	// - A lowercase AWS SDK for JavaScript v2 constructor name (`apigateway`)
	//
	// The `api` is the name of an AWS API call, in one of the following forms:
	// - An API call name as found in the API Reference documentation (`GetObject`)
	// - The API call name starting with a lowercase letter (`getObject`)
	// - The AWS SDK for JavaScript v3 command class name (`GetObjectCommand`).
	//
	// Example:
	//   var app App
	//   var integ IntegTest
	//
	//   integ.Assertions.AwsApiCall(jsii.String("SQS"), jsii.String("sendMessage"), map[string]*string{
	//   	"QueueUrl": jsii.String("url"),
	//   	"MessageBody": jsii.String("hello"),
	//   })
	//   message := integ.Assertions.AwsApiCall(jsii.String("SQS"), jsii.String("receiveMessage"), map[string]*string{
	//   	"QueueUrl": jsii.String("url"),
	//   })
	//   message.Expect(awscdkintegtestsalpha.ExpectedResult_ObjectLike(map[string]interface{}{
	//   	"Messages": []interface{}{
	//   		map[string]*string{
	//   			"Body": jsii.String("hello"),
	//   		},
	//   	},
	//   }))
	//
	// Experimental.
	AwsApiCall(service *string, api *string, parameters interface{}, outputPaths *[]*string) IApiCall
	// Assert that the ExpectedResult is equal to the ActualResult.
	//
	// Example:
	//   var integ IntegTest
	//   var apiCall AwsApiCall
	//
	//   integ.Assertions.Expect(jsii.String("invoke"), awscdkintegtestsalpha.ExpectedResult_ObjectLike(map[string]interface{}{
	//   	"Payload": jsii.String("OK"),
	//   }), awscdkintegtestsalpha.ActualResult_FromAwsApiCall(apiCall, jsii.String("Body")))
	//
	// Experimental.
	Expect(id *string, expected ExpectedResult, actual ActualResult)
	// Make an HTTP call to the provided endpoint.
	//
	// Example:
	//   var app App
	//   var integ IntegTest
	//
	//   call := integ.Assertions.HttpApiCall(jsii.String("https://example.com/test"))
	//   call.Expect(awscdkintegtestsalpha.ExpectedResult_ObjectLike(map[string]interface{}{
	//   	"Message": jsii.String("Hello World!"),
	//   }))
	//
	// Experimental.
	HttpApiCall(url *string, options *FetchOptions) IApiCall
	// Invoke a lambda function and return the response which can be asserted.
	//
	// Example:
	//   var app App
	//   var integ IntegTest
	//
	//   invoke := integ.Assertions.InvokeFunction(&LambdaInvokeFunctionProps{
	//   	FunctionName: jsii.String("my-function"),
	//   })
	//   invoke.Expect(awscdkintegtestsalpha.ExpectedResult_ObjectLike(map[string]interface{}{
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

func (i *jsiiProxy_IDeployAssert) AwsApiCall(service *string, api *string, parameters interface{}, outputPaths *[]*string) IApiCall {
	if err := i.validateAwsApiCallParameters(service, api); err != nil {
		panic(err)
	}
	var returns IApiCall

	_jsii_.Invoke(
		i,
		"awsApiCall",
		[]interface{}{service, api, parameters, outputPaths},
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

func (i *jsiiProxy_IDeployAssert) HttpApiCall(url *string, options *FetchOptions) IApiCall {
	if err := i.validateHttpApiCallParameters(url, options); err != nil {
		panic(err)
	}
	var returns IApiCall

	_jsii_.Invoke(
		i,
		"httpApiCall",
		[]interface{}{url, options},
		&returns,
	)

	return returns
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

