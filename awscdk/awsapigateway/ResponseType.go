package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Supported types of gateway responses.
//
// Example:
//   api := apigateway.NewRestApi(this, jsii.String("books-api"))
//   api.AddGatewayResponse(jsii.String("test-response"), &GatewayResponseOptions{
//   	Type: apigateway.ResponseType_ACCESS_DENIED(),
//   	StatusCode: jsii.String("500"),
//   	ResponseHeaders: map[string]*string{
//   		// Note that values must be enclosed within a pair of single quotes
//   		"Access-Control-Allow-Origin": jsii.String("'test.com'"),
//   		"test-key": jsii.String("'test-value'"),
//   	},
//   	Templates: map[string]*string{
//   		"application/json": jsii.String("{ \"message\": $context.error.messageString, \"statusCode\": \"488\", \"type\": \"$context.error.responseType\" }"),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/supported-gateway-response-types.html
//
type ResponseType interface {
	// Valid value of response type.
	ResponseType() *string
}

// The jsii proxy struct for ResponseType
type jsiiProxy_ResponseType struct {
	_ byte // padding
}

func (j *jsiiProxy_ResponseType) ResponseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseType",
		&returns,
	)
	return returns
}


// A custom response type to support future cases.
func ResponseType_Of(type_ *string) ResponseType {
	_init_.Initialize()

	if err := validateResponseType_OfParameters(type_); err != nil {
		panic(err)
	}
	var returns ResponseType

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.ResponseType",
		"of",
		[]interface{}{type_},
		&returns,
	)

	return returns
}

func ResponseType_ACCESS_DENIED() ResponseType {
	_init_.Initialize()
	var returns ResponseType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.ResponseType",
		"ACCESS_DENIED",
		&returns,
	)
	return returns
}

func ResponseType_API_CONFIGURATION_ERROR() ResponseType {
	_init_.Initialize()
	var returns ResponseType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.ResponseType",
		"API_CONFIGURATION_ERROR",
		&returns,
	)
	return returns
}

func ResponseType_AUTHORIZER_CONFIGURATION_ERROR() ResponseType {
	_init_.Initialize()
	var returns ResponseType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.ResponseType",
		"AUTHORIZER_CONFIGURATION_ERROR",
		&returns,
	)
	return returns
}

func ResponseType_AUTHORIZER_FAILURE() ResponseType {
	_init_.Initialize()
	var returns ResponseType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.ResponseType",
		"AUTHORIZER_FAILURE",
		&returns,
	)
	return returns
}

func ResponseType_BAD_REQUEST_BODY() ResponseType {
	_init_.Initialize()
	var returns ResponseType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.ResponseType",
		"BAD_REQUEST_BODY",
		&returns,
	)
	return returns
}

func ResponseType_BAD_REQUEST_PARAMETERS() ResponseType {
	_init_.Initialize()
	var returns ResponseType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.ResponseType",
		"BAD_REQUEST_PARAMETERS",
		&returns,
	)
	return returns
}

func ResponseType_DEFAULT_4XX() ResponseType {
	_init_.Initialize()
	var returns ResponseType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.ResponseType",
		"DEFAULT_4XX",
		&returns,
	)
	return returns
}

func ResponseType_DEFAULT_5XX() ResponseType {
	_init_.Initialize()
	var returns ResponseType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.ResponseType",
		"DEFAULT_5XX",
		&returns,
	)
	return returns
}

func ResponseType_EXPIRED_TOKEN() ResponseType {
	_init_.Initialize()
	var returns ResponseType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.ResponseType",
		"EXPIRED_TOKEN",
		&returns,
	)
	return returns
}

func ResponseType_INTEGRATION_FAILURE() ResponseType {
	_init_.Initialize()
	var returns ResponseType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.ResponseType",
		"INTEGRATION_FAILURE",
		&returns,
	)
	return returns
}

func ResponseType_INTEGRATION_TIMEOUT() ResponseType {
	_init_.Initialize()
	var returns ResponseType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.ResponseType",
		"INTEGRATION_TIMEOUT",
		&returns,
	)
	return returns
}

func ResponseType_INVALID_API_KEY() ResponseType {
	_init_.Initialize()
	var returns ResponseType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.ResponseType",
		"INVALID_API_KEY",
		&returns,
	)
	return returns
}

func ResponseType_INVALID_SIGNATURE() ResponseType {
	_init_.Initialize()
	var returns ResponseType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.ResponseType",
		"INVALID_SIGNATURE",
		&returns,
	)
	return returns
}

func ResponseType_MISSING_AUTHENTICATION_TOKEN() ResponseType {
	_init_.Initialize()
	var returns ResponseType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.ResponseType",
		"MISSING_AUTHENTICATION_TOKEN",
		&returns,
	)
	return returns
}

func ResponseType_QUOTA_EXCEEDED() ResponseType {
	_init_.Initialize()
	var returns ResponseType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.ResponseType",
		"QUOTA_EXCEEDED",
		&returns,
	)
	return returns
}

func ResponseType_REQUEST_TOO_LARGE() ResponseType {
	_init_.Initialize()
	var returns ResponseType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.ResponseType",
		"REQUEST_TOO_LARGE",
		&returns,
	)
	return returns
}

func ResponseType_RESOURCE_NOT_FOUND() ResponseType {
	_init_.Initialize()
	var returns ResponseType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.ResponseType",
		"RESOURCE_NOT_FOUND",
		&returns,
	)
	return returns
}

func ResponseType_THROTTLED() ResponseType {
	_init_.Initialize()
	var returns ResponseType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.ResponseType",
		"THROTTLED",
		&returns,
	)
	return returns
}

func ResponseType_UNAUTHORIZED() ResponseType {
	_init_.Initialize()
	var returns ResponseType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.ResponseType",
		"UNAUTHORIZED",
		&returns,
	)
	return returns
}

func ResponseType_UNSUPPORTED_MEDIA_TYPE() ResponseType {
	_init_.Initialize()
	var returns ResponseType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.ResponseType",
		"UNSUPPORTED_MEDIA_TYPE",
		&returns,
	)
	return returns
}

func ResponseType_WAF_FILTERED() ResponseType {
	_init_.Initialize()
	var returns ResponseType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.ResponseType",
		"WAF_FILTERED",
		&returns,
	)
	return returns
}

