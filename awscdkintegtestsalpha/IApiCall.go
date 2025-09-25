package awscdkintegtestsalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkintegtestsalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an ApiCall.
// Experimental.
type IApiCall interface {
	constructs.IConstruct
	// Assert that the ExpectedResult is equal to the result of the AwsApiCall at the given path.
	//
	// Providing a path will filter the output of the initial API call.
	//
	// For example the SQS.receiveMessage api response would look
	// like:
	//
	// If you wanted to assert the value of `Body` you could do.
	//
	// Example:
	//   var integ integTest
	//   actual := map[string][]map[string]interface{}{
	//   	"Messages": []map[string]interface{}{
	//   		map[string]interface{}{
	//   			"MessageId": jsii.String(""),
	//   			"ReceiptHandle": jsii.String(""),
	//   			"MD5OfBody": jsii.String(""),
	//   			"Body": jsii.String("hello"),
	//   			"Attributes": map[string]interface{}{
	//   			},
	//   			"MD5OfMessageAttributes": map[string]interface{}{
	//   			},
	//   			"MessageAttributes": map[string]interface{}{
	//   			},
	//   		},
	//   	},
	//   }
	//   message := integ.Assertions.AwsApiCall(jsii.String("SQS"), jsii.String("receiveMessage"))
	//
	//   message.AssertAtPath(jsii.String("Messages.0.Body"), awscdkintegtestsalpha.ExpectedResult_StringLikeRegexp(jsii.String("hello")))
	//
	// Experimental.
	AssertAtPath(path *string, expected ExpectedResult) IApiCall
	// Assert that the ExpectedResult is equal to the result of the AwsApiCall.
	//
	// Example:
	//   var integ integTest
	//
	//   invoke := integ.Assertions.InvokeFunction(&LambdaInvokeFunctionProps{
	//   	FunctionName: jsii.String("my-func"),
	//   })
	//   invoke.Expect(awscdkintegtestsalpha.ExpectedResult_ObjectLike(map[string]interface{}{
	//   	"Payload": jsii.String("OK"),
	//   }))
	//
	// Experimental.
	Expect(expected ExpectedResult) IApiCall
	// Returns the value of an attribute of the custom resource of an arbitrary type.
	//
	// Attributes are returned from the custom resource provider through the
	// `Data` map where the key is the attribute name.
	//
	// Returns: a token for `Fn::GetAtt`. Use `Token.asXxx` to encode the returned `Reference` as a specific type or
	// use the convenience `getAttString` for string attributes.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Returns the value of an attribute of the custom resource of type string.
	//
	// Attributes are returned from the custom resource provider through the
	// `Data` map where the key is the attribute name.
	//
	// Returns: a token for `Fn::GetAtt` encoded as a string.
	// Experimental.
	GetAttString(attributeName *string) *string
	// Allows you to chain IApiCalls. This adds an explicit dependency betweent the two resources.
	//
	// Returns the IApiCall provided as `next`.
	//
	// Example:
	//   var first iApiCall
	//   var second iApiCall
	//
	//
	//   first.Next(second)
	//
	// Experimental.
	Next(next IApiCall) IApiCall
	// Wait for the IApiCall to return the expected response.
	//
	// If no expected response is specified then it will wait for
	// the IApiCall to return a success.
	//
	// Example:
	//   var integ integTest
	//   var executionArn string
	//
	//   integ.Assertions.AwsApiCall(jsii.String("StepFunctions"), jsii.String("describeExecution"), map[string]*string{
	//   	"executionArn": jsii.String(executionArn),
	//   }).WaitForAssertions()
	//
	// Experimental.
	WaitForAssertions(options *WaiterStateMachineOptions) IApiCall
	// access the AssertionsProvider.
	//
	// This can be used to add additional IAM policies
	// the the provider role policy.
	//
	// Example:
	//   var apiCall awsApiCall
	//
	//   apiCall.Provider.AddToRolePolicy(map[string]interface{}{
	//   	"Effect": jsii.String("Allow"),
	//   	"Action": []*string{
	//   		jsii.String("s3:GetObject"),
	//   	},
	//   	"Resource": []*string{
	//   		jsii.String("*"),
	//   	},
	//   })
	//
	// Experimental.
	Provider() AssertionsProvider
}

// The jsii proxy for IApiCall
type jsiiProxy_IApiCall struct {
	internal.Type__constructsIConstruct
}

func (i *jsiiProxy_IApiCall) AssertAtPath(path *string, expected ExpectedResult) IApiCall {
	if err := i.validateAssertAtPathParameters(path, expected); err != nil {
		panic(err)
	}
	var returns IApiCall

	_jsii_.Invoke(
		i,
		"assertAtPath",
		[]interface{}{path, expected},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApiCall) Expect(expected ExpectedResult) IApiCall {
	if err := i.validateExpectParameters(expected); err != nil {
		panic(err)
	}
	var returns IApiCall

	_jsii_.Invoke(
		i,
		"expect",
		[]interface{}{expected},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApiCall) GetAtt(attributeName *string) awscdk.Reference {
	if err := i.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		i,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApiCall) GetAttString(attributeName *string) *string {
	if err := i.validateGetAttStringParameters(attributeName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getAttString",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApiCall) Next(next IApiCall) IApiCall {
	if err := i.validateNextParameters(next); err != nil {
		panic(err)
	}
	var returns IApiCall

	_jsii_.Invoke(
		i,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApiCall) WaitForAssertions(options *WaiterStateMachineOptions) IApiCall {
	if err := i.validateWaitForAssertionsParameters(options); err != nil {
		panic(err)
	}
	var returns IApiCall

	_jsii_.Invoke(
		i,
		"waitForAssertions",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IApiCall) Provider() AssertionsProvider {
	var returns AssertionsProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

