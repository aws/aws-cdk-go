package awsapigatewayv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Calculations and operations for HTTP APIs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var httpApiRef IHttpApiRef
//
//   httpApiHelper := awscdk.Aws_apigatewayv2.HttpApiHelper_FromHttpApi(httpApiRef)
//
type HttpApiHelper interface {
	// Get the "execute-api" ARN.
	//
	// When 'ANY' is passed to the method, an ARN with the method set to '*' is obtained.
	// Default: - The default behavior applies when no specific method, path, or stage is provided.
	// In this case, the ARN will cover all methods, all resources, and all stages of this API.
	// Specifically, if 'method' is not specified, it defaults to '*', representing all methods.
	// If 'path' is not specified, it defaults to '/*', representing all paths.
	// If 'stage' is not specified, it also defaults to '*', representing all stages.
	//
	ArnForExecuteApi(method *string, path *string, stage *string) *string
}

// The jsii proxy struct for HttpApiHelper
type jsiiProxy_HttpApiHelper struct {
	_ byte // padding
}

// Return an `HttpApiHelper` for the given HTTP API.
func HttpApiHelper_FromHttpApi(httpApi IHttpApiRef) HttpApiHelper {
	_init_.Initialize()

	if err := validateHttpApiHelper_FromHttpApiParameters(httpApi); err != nil {
		panic(err)
	}
	var returns HttpApiHelper

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.HttpApiHelper",
		"fromHttpApi",
		[]interface{}{httpApi},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpApiHelper) ArnForExecuteApi(method *string, path *string, stage *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"arnForExecuteApi",
		[]interface{}{method, path, stage},
		&returns,
	)

	return returns
}

