package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
)

// Options when binding a log destination to a RestApi Stage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessLogDestinationConfig := &accessLogDestinationConfig{
//   	destinationArn: jsii.String("destinationArn"),
//   }
//
type AccessLogDestinationConfig struct {
	// The Amazon Resource Name (ARN) of the destination resource.
	DestinationArn *string `field:"required" json:"destinationArn" yaml:"destinationArn"`
}

// $context variables that can be used to customize access log pattern.
//
// Example:
//   apigateway.accessLogFormat.custom(jSON.stringify(map[string]interface{}{
//   	"requestId": apigateway.AccessLogField.contextRequestId(),
//   	"sourceIp": apigateway.AccessLogField.contextIdentitySourceIp(),
//   	"method": apigateway.AccessLogField.contextHttpMethod(),
//   	"userContext": map[string]*string{
//   		"sub": apigateway.AccessLogField.contextAuthorizerClaims(jsii.String("sub")),
//   		"email": apigateway.AccessLogField.contextAuthorizerClaims(jsii.String("email")),
//   	},
//   }))
//
type AccessLogField interface {
}

// The jsii proxy struct for AccessLogField
type jsiiProxy_AccessLogField struct {
	_ byte // padding
}

func NewAccessLogField() AccessLogField {
	_init_.Initialize()

	j := jsiiProxy_AccessLogField{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewAccessLogField_Override(a AccessLogField) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		nil, // no parameters
		a,
	)
}

// The API owner's AWS account ID.
func AccessLogField_ContextAccountId() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextAccountId",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The identifier API Gateway assigns to your API.
func AccessLogField_ContextApiId() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextApiId",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The stringified value of the specified key-value pair of the `context` map returned from an API Gateway Lambda authorizer function.
// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-use-lambda-authorizer.html
//
func AccessLogField_ContextAuthorizer(property *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextAuthorizer",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// A property of the claims returned from the Amazon Cognito user pool after the method caller is successfully authenticated.
// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-integrate-with-cognito.html
//
func AccessLogField_ContextAuthorizerClaims(property *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextAuthorizerClaims",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// The authorizer latency in ms.
func AccessLogField_ContextAuthorizerIntegrationLatency() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextAuthorizerIntegrationLatency",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The principal user identification associated with the token sent by the client and returned from an API Gateway Lambda authorizer (formerly known as a custom authorizer).
// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-use-lambda-authorizer.html
//
func AccessLogField_ContextAuthorizerPrincipalId() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextAuthorizerPrincipalId",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The AWS endpoint's request ID.
func AccessLogField_ContextAwsEndpointRequestId() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextAwsEndpointRequestId",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The full domain name used to invoke the API.
//
// This should be the same as the incoming `Host` header.
func AccessLogField_ContextDomainName() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextDomainName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The first label of the `$context.domainName`. This is often used as a caller/customer identifier.
func AccessLogField_ContextDomainPrefix() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextDomainPrefix",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A string containing an API Gateway error message.
func AccessLogField_ContextErrorMessage() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextErrorMessage",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The quoted value of $context.error.message, namely "$context.error.message".
func AccessLogField_ContextErrorMessageString() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextErrorMessageString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A type of GatewayResponse.
//
// This variable can only be used for simple variable substitution in a GatewayResponse body-mapping template,
// which is not processed by the Velocity Template Language engine, and in access logging.
// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/customize-gateway-responses.html
//
func AccessLogField_ContextErrorResponseType() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextErrorResponseType",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A string containing a detailed validation error message.
func AccessLogField_ContextErrorValidationErrorString() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextErrorValidationErrorString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The extended ID that API Gateway assigns to the API request, which contains more useful information for debugging/troubleshooting.
func AccessLogField_ContextExtendedRequestId() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextExtendedRequestId",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The HTTP method used.
//
// Valid values include: `DELETE`, `GET`, `HEAD`, `OPTIONS`, `PATCH`, `POST`, and `PUT`.
func AccessLogField_ContextHttpMethod() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextHttpMethod",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The AWS account ID associated with the request.
func AccessLogField_ContextIdentityAccountId() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextIdentityAccountId",
		nil, // no parameters
		&returns,
	)

	return returns
}

// For API methods that require an API key, this variable is the API key associated with the method request.
//
// For methods that don't require an API key, this variable is.
// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-api-usage-plans.html
//
func AccessLogField_ContextIdentityApiKey() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextIdentityApiKey",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The API key ID associated with an API request that requires an API key.
func AccessLogField_ContextIdentityApiKeyId() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextIdentityApiKeyId",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The principal identifier of the caller making the request.
func AccessLogField_ContextIdentityCaller() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextIdentityCaller",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The Amazon Cognito authentication provider used by the caller making the request.
//
// Available only if the request was signed with Amazon Cognito credentials.
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-identity.html
//
func AccessLogField_ContextIdentityCognitoAuthenticationProvider() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextIdentityCognitoAuthenticationProvider",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The Amazon Cognito authentication type of the caller making the request.
//
// Available only if the request was signed with Amazon Cognito credentials.
func AccessLogField_ContextIdentityCognitoAuthenticationType() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextIdentityCognitoAuthenticationType",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The Amazon Cognito identity ID of the caller making the request.
//
// Available only if the request was signed with Amazon Cognito credentials.
func AccessLogField_ContextIdentityCognitoIdentityId() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextIdentityCognitoIdentityId",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The Amazon Cognito identity pool ID of the caller making the request.
//
// Available only if the request was signed with Amazon Cognito credentials.
func AccessLogField_ContextIdentityCognitoIdentityPoolId() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextIdentityCognitoIdentityPoolId",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The AWS organization ID.
func AccessLogField_ContextIdentityPrincipalOrgId() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextIdentityPrincipalOrgId",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The source IP address of the TCP connection making the request to API Gateway.
//
// Warning: You should not trust this value if there is any chance that the `X-Forwarded-For` header could be forged.
func AccessLogField_ContextIdentitySourceIp() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextIdentitySourceIp",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The principal identifier of the user making the request.
//
// Used in Lambda authorizers.
// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-lambda-authorizer-output.html
//
func AccessLogField_ContextIdentityUser() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextIdentityUser",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The User-Agent header of the API caller.
func AccessLogField_ContextIdentityUserAgent() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextIdentityUserAgent",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The Amazon Resource Name (ARN) of the effective user identified after authentication.
// See: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users.html
//
func AccessLogField_ContextIdentityUserArn() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextIdentityUserArn",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The integration latency in ms.
func AccessLogField_ContextIntegrationLatency() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextIntegrationLatency",
		nil, // no parameters
		&returns,
	)

	return returns
}

// For Lambda proxy integration, this parameter represents the status code returned from AWS Lambda, not from the backend Lambda function.
func AccessLogField_ContextIntegrationStatus() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextIntegrationStatus",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The request path.
//
// For example, for a non-proxy request URL of https://{rest-api-id.execute-api.{region}.amazonaws.com/{stage}/root/child,
// this value is /{stage}/root/child.
func AccessLogField_ContextPath() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextPath",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The request protocol, for example, HTTP/1.1.
func AccessLogField_ContextProtocol() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextProtocol",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The ID that API Gateway assigns to the API request.
func AccessLogField_ContextRequestId() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextRequestId",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The request header override.
//
// If this parameter is defined, it contains the headers to be used instead of the HTTP Headers that are defined in the Integration Request pane.
// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-override-request-response-parameters.html
//
func AccessLogField_ContextRequestOverrideHeader(headerName *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextRequestOverrideHeader",
		[]interface{}{headerName},
		&returns,
	)

	return returns
}

// The request path override.
//
// If this parameter is defined,
// it contains the request path to be used instead of the URL Path Parameters that are defined in the Integration Request pane.
// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-override-request-response-parameters.html
//
func AccessLogField_ContextRequestOverridePath(pathName *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextRequestOverridePath",
		[]interface{}{pathName},
		&returns,
	)

	return returns
}

// The request query string override.
//
// If this parameter is defined, it contains the request query strings to be used instead
// of the URL Query String Parameters that are defined in the Integration Request pane.
func AccessLogField_ContextRequestOverrideQuerystring(querystringName *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextRequestOverrideQuerystring",
		[]interface{}{querystringName},
		&returns,
	)

	return returns
}

// The CLF-formatted request time (dd/MMM/yyyy:HH:mm:ss +-hhmm).
func AccessLogField_ContextRequestTime() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextRequestTime",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The Epoch-formatted request time.
func AccessLogField_ContextRequestTimeEpoch() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextRequestTimeEpoch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The identifier that API Gateway assigns to your resource.
func AccessLogField_ContextResourceId() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextResourceId",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The path to your resource.
//
// For example, for the non-proxy request URI of `https://{rest-api-id.execute-api.{region}.amazonaws.com/{stage}/root/child`,
// The $context.resourcePath value is `/root/child`.
// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-create-api-step-by-step.html
//
func AccessLogField_ContextResourcePath() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextResourcePath",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The response latency in ms.
func AccessLogField_ContextResponseLatency() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextResponseLatency",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The response payload length.
func AccessLogField_ContextResponseLength() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextResponseLength",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The response header override.
//
// If this parameter is defined, it contains the header to be returned instead of the Response header
// that is defined as the Default mapping in the Integration Response pane.
// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-override-request-response-parameters.html
//
func AccessLogField_ContextResponseOverrideHeader(headerName *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextResponseOverrideHeader",
		[]interface{}{headerName},
		&returns,
	)

	return returns
}

// The response status code override.
//
// If this parameter is defined, it contains the status code to be returned instead of the Method response status
// that is defined as the Default mapping in the Integration Response pane.
// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-override-request-response-parameters.html
//
func AccessLogField_ContextResponseOverrideStatus() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextResponseOverrideStatus",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The deployment stage of the API request (for example, `Beta` or `Prod`).
func AccessLogField_ContextStage() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextStage",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The method response status.
func AccessLogField_ContextStatus() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextStatus",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The response received from AWS WAF: `WAF_ALLOW` or `WAF_BLOCK`.
//
// Will not be set if the stage is not associated with a web ACL.
// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-control-access-aws-waf.html
//
func AccessLogField_ContextWafResponseCode() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextWafResponseCode",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The complete ARN of the web ACL that is used to decide whether to allow or block the request.
//
// Will not be set if the stage is not associated with a web ACL.
// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-control-access-aws-waf.html
//
func AccessLogField_ContextWebaclArn() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextWebaclArn",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The trace ID for the X-Ray trace.
// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-enabling-xray.html
//
func AccessLogField_ContextXrayTraceId() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextXrayTraceId",
		nil, // no parameters
		&returns,
	)

	return returns
}

// factory methods for access log format.
//
// Example:
//   logGroup := logs.NewLogGroup(this, jsii.String("ApiGatewayAccessLogs"))
//   apigateway.NewRestApi(this, jsii.String("books"), &restApiProps{
//   	deployOptions: &stageOptions{
//   		accessLogDestination: apigateway.NewLogGroupLogDestination(logGroup),
//   		accessLogFormat: apigateway.accessLogFormat.custom(
//   		fmt.Sprintf("%v %v %v", apigateway.accessLogField.contextRequestId(), apigateway.*accessLogField.contextErrorMessage(), apigateway.*accessLogField.contextErrorMessageString())),
//   	},
//   })
//
type AccessLogFormat interface {
	// Output a format string to be used with CloudFormation.
	ToString() *string
}

// The jsii proxy struct for AccessLogFormat
type jsiiProxy_AccessLogFormat struct {
	_ byte // padding
}

// Generate Common Log Format.
func AccessLogFormat_Clf() AccessLogFormat {
	_init_.Initialize()

	var returns AccessLogFormat

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogFormat",
		"clf",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Custom log format.
//
// You can create any log format string. You can easily get the $ context variable by using the methods of AccessLogField.
//
// Example:
//   apigateway.accessLogFormat.custom(jSON.stringify(map[string]interface{}{
//   	"requestId": apigateway.AccessLogField.contextRequestId(),
//   	"sourceIp": apigateway.AccessLogField.contextIdentitySourceIp(),
//   	"method": apigateway.AccessLogField.contextHttpMethod(),
//   	"userContext": map[string]*string{
//   		"sub": apigateway.AccessLogField.contextAuthorizerClaims(jsii.String("sub")),
//   		"email": apigateway.AccessLogField.contextAuthorizerClaims(jsii.String("email")),
//   	},
//   }))
//
func AccessLogFormat_Custom(format *string) AccessLogFormat {
	_init_.Initialize()

	var returns AccessLogFormat

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogFormat",
		"custom",
		[]interface{}{format},
		&returns,
	)

	return returns
}

// Access log will be produced in the JSON format with a set of fields most useful in the access log.
//
// All fields are turned on by default with the
// option to turn off specific fields.
func AccessLogFormat_JsonWithStandardFields(fields *JsonWithStandardFieldProps) AccessLogFormat {
	_init_.Initialize()

	var returns AccessLogFormat

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogFormat",
		"jsonWithStandardFields",
		[]interface{}{fields},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessLogFormat) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options to the UsagePlan.addApiKey() method.
//
// Example:
//   var usageplan usagePlan
//   var apiKey apiKey
//
//
//   usageplan.addApiKey(apiKey, &addApiKeyOptions{
//   	overrideLogicalId: jsii.String("..."),
//   })
//
type AddApiKeyOptions struct {
	// Override the CloudFormation logical id of the AWS::ApiGateway::UsagePlanKey resource.
	OverrideLogicalId *string `field:"optional" json:"overrideLogicalId" yaml:"overrideLogicalId"`
}

// Represents an OpenAPI definition asset.
//
// Example:
//   var integration integration
//
//
//   api := apigateway.NewSpecRestApi(this, jsii.String("books-api"), &specRestApiProps{
//   	apiDefinition: apigateway.apiDefinition.fromAsset(jsii.String("path-to-file.json")),
//   })
//
//   booksResource := api.root.addResource(jsii.String("books"))
//   booksResource.addMethod(jsii.String("GET"), integration)
//
type ApiDefinition interface {
	// Called when the specification is initialized to allow this object to bind to the stack, add resources and have fun.
	Bind(scope constructs.Construct) *ApiDefinitionConfig
	// Called after the CFN RestApi resource has been created to allow the Api Definition to bind to it.
	//
	// Specifically it's required to allow assets to add
	// metadata for tooling like SAM CLI to be able to find their origins.
	BindAfterCreate(_scope constructs.Construct, _restApi IRestApi)
}

// The jsii proxy struct for ApiDefinition
type jsiiProxy_ApiDefinition struct {
	_ byte // padding
}

func NewApiDefinition_Override(a ApiDefinition) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.ApiDefinition",
		nil, // no parameters
		a,
	)
}

// Loads the API specification from a local disk asset.
func ApiDefinition_FromAsset(file *string, options *awss3assets.AssetOptions) AssetApiDefinition {
	_init_.Initialize()

	var returns AssetApiDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.ApiDefinition",
		"fromAsset",
		[]interface{}{file, options},
		&returns,
	)

	return returns
}

// Creates an API definition from a specification file in an S3 bucket.
func ApiDefinition_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3ApiDefinition {
	_init_.Initialize()

	var returns S3ApiDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.ApiDefinition",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Create an API definition from an inline object.
//
// The inline object must follow the
// schema of OpenAPI 2.0 or OpenAPI 3.0
//
// Example:
//   apigateway.apiDefinition.fromInline(map[string]interface{}{
//   	"openapi": jsii.String("3.0.2"),
//   	"paths": map[string]map[string]map[string]map[string]map[string]map[string]map[string]map[string]*string{
//   		"/pets": map[string]map[string]map[string]map[string]map[string]map[string]map[string]*string{
//   			"get": map[string]map[string]map[string]map[string]map[string]map[string]*string{
//   				"responses": map[string]map[string]map[string]map[string]map[string]*string{
//   					jsii.Number(200): map[string]map[string]map[string]map[string]*string{
//   						"content": map[string]map[string]map[string]*string{
//   							"application/json": map[string]map[string]*string{
//   								"schema": map[string]*string{
//   									"$ref": jsii.String("#/components/schemas/Empty"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   				"x-amazon-apigateway-integration": map[string]interface{}{
//   					"responses": map[string]map[string]*string{
//   						"default": map[string]*string{
//   							"statusCode": jsii.String("200"),
//   						},
//   					},
//   					"requestTemplates": map[string]*string{
//   						"application/json": jsii.String("{\"statusCode\": 200}"),
//   					},
//   					"passthroughBehavior": jsii.String("when_no_match"),
//   					"type": jsii.String("mock"),
//   				},
//   			},
//   		},
//   	},
//   	"components": map[string]map[string]map[string]*string{
//   		"schemas": map[string]map[string]*string{
//   			"Empty": map[string]*string{
//   				"title": jsii.String("Empty Schema"),
//   				"type": jsii.String("object"),
//   			},
//   		},
//   	},
//   })
//
func ApiDefinition_FromInline(definition interface{}) InlineApiDefinition {
	_init_.Initialize()

	var returns InlineApiDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.ApiDefinition",
		"fromInline",
		[]interface{}{definition},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiDefinition) Bind(scope constructs.Construct) *ApiDefinitionConfig {
	var returns *ApiDefinitionConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiDefinition) BindAfterCreate(_scope constructs.Construct, _restApi IRestApi) {
	_jsii_.InvokeVoid(
		a,
		"bindAfterCreate",
		[]interface{}{_scope, _restApi},
	)
}

// Post-Binding Configuration for a CDK construct.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var inlineDefinition interface{}
//
//   apiDefinitionConfig := &apiDefinitionConfig{
//   	inlineDefinition: inlineDefinition,
//   	s3Location: &apiDefinitionS3Location{
//   		bucket: jsii.String("bucket"),
//   		key: jsii.String("key"),
//
//   		// the properties below are optional
//   		version: jsii.String("version"),
//   	},
//   }
//
type ApiDefinitionConfig struct {
	// Inline specification (mutually exclusive with `s3Location`).
	InlineDefinition interface{} `field:"optional" json:"inlineDefinition" yaml:"inlineDefinition"`
	// The location of the specification in S3 (mutually exclusive with `inlineDefinition`).
	S3Location *ApiDefinitionS3Location `field:"optional" json:"s3Location" yaml:"s3Location"`
}

// S3 location of the API definition file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiDefinitionS3Location := &apiDefinitionS3Location{
//   	bucket: jsii.String("bucket"),
//   	key: jsii.String("key"),
//
//   	// the properties below are optional
//   	version: jsii.String("version"),
//   }
//
type ApiDefinitionS3Location struct {
	// The S3 bucket.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The S3 key.
	Key *string `field:"required" json:"key" yaml:"key"`
	// An optional version.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

// An API Gateway ApiKey.
//
// An ApiKey can be distributed to API clients that are executing requests
// for Method resources that require an Api Key.
//
// Example:
//   importedKey := apigateway.apiKey.fromApiKeyId(this, jsii.String("imported-key"), jsii.String("<api-key-id>"))
//
type ApiKey interface {
	awscdk.Resource
	IApiKey
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The API key ARN.
	KeyArn() *string
	// The API key ID.
	KeyId() *string
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Permits the IAM principal all read operations through this key.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Permits the IAM principal all read and write operations through this key.
	GrantReadWrite(grantee awsiam.IGrantable) awsiam.Grant
	// Permits the IAM principal all write operations through this key.
	GrantWrite(grantee awsiam.IGrantable) awsiam.Grant
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for ApiKey
type jsiiProxy_ApiKey struct {
	internal.Type__awscdkResource
	jsiiProxy_IApiKey
}

func (j *jsiiProxy_ApiKey) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiKey) KeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiKey) KeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiKey) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiKey) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiKey) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewApiKey(scope constructs.Construct, id *string, props *ApiKeyProps) ApiKey {
	_init_.Initialize()

	j := jsiiProxy_ApiKey{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.ApiKey",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewApiKey_Override(a ApiKey, scope constructs.Construct, id *string, props *ApiKeyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.ApiKey",
		[]interface{}{scope, id, props},
		a,
	)
}

// Import an ApiKey by its Id.
func ApiKey_FromApiKeyId(scope constructs.Construct, id *string, apiKeyId *string) IApiKey {
	_init_.Initialize()

	var returns IApiKey

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.ApiKey",
		"fromApiKeyId",
		[]interface{}{scope, id, apiKeyId},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func ApiKey_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.ApiKey",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func ApiKey_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.ApiKey",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ApiKey_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.ApiKey",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiKey) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (a *jsiiProxy_ApiKey) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiKey) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiKey) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiKey) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		a,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiKey) GrantReadWrite(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		a,
		"grantReadWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiKey) GrantWrite(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		a,
		"grantWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The options for creating an API Key.
//
// Example:
//   var api restApi
//
//   key := api.addApiKey(jsii.String("ApiKey"), &apiKeyOptions{
//   	apiKeyName: jsii.String("myApiKey1"),
//   	value: jsii.String("MyApiKeyThatIsAtLeast20Characters"),
//   })
//
type ApiKeyOptions struct {
	// Adds a CORS preflight OPTIONS method to this resource and all child resources.
	//
	// You can add CORS at the resource-level using `addCorsPreflight`.
	DefaultCorsPreflightOptions *CorsOptions `field:"optional" json:"defaultCorsPreflightOptions" yaml:"defaultCorsPreflightOptions"`
	// An integration to use as a default for all methods created within this API unless an integration is specified.
	DefaultIntegration Integration `field:"optional" json:"defaultIntegration" yaml:"defaultIntegration"`
	// Method options to use as a default for all methods created within this API unless custom options are specified.
	DefaultMethodOptions *MethodOptions `field:"optional" json:"defaultMethodOptions" yaml:"defaultMethodOptions"`
	// A name for the API key.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the API key name.
	ApiKeyName *string `field:"optional" json:"apiKeyName" yaml:"apiKeyName"`
	// A description of the purpose of the API key.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The value of the API key.
	//
	// Must be at least 20 characters long.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

// ApiKey Properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var authorizer authorizer
//   var integration integration
//   var model model
//   var requestValidator requestValidator
//   var restApi restApi
//
//   apiKeyProps := &apiKeyProps{
//   	apiKeyName: jsii.String("apiKeyName"),
//   	customerId: jsii.String("customerId"),
//   	defaultCorsPreflightOptions: &corsOptions{
//   		allowOrigins: []*string{
//   			jsii.String("allowOrigins"),
//   		},
//
//   		// the properties below are optional
//   		allowCredentials: jsii.Boolean(false),
//   		allowHeaders: []*string{
//   			jsii.String("allowHeaders"),
//   		},
//   		allowMethods: []*string{
//   			jsii.String("allowMethods"),
//   		},
//   		disableCache: jsii.Boolean(false),
//   		exposeHeaders: []*string{
//   			jsii.String("exposeHeaders"),
//   		},
//   		maxAge: cdk.duration.minutes(jsii.Number(30)),
//   		statusCode: jsii.Number(123),
//   	},
//   	defaultIntegration: integration,
//   	defaultMethodOptions: &methodOptions{
//   		apiKeyRequired: jsii.Boolean(false),
//   		authorizationScopes: []*string{
//   			jsii.String("authorizationScopes"),
//   		},
//   		authorizationType: awscdk.Aws_apigateway.authorizationType_NONE,
//   		authorizer: authorizer,
//   		methodResponses: []methodResponse{
//   			&methodResponse{
//   				statusCode: jsii.String("statusCode"),
//
//   				// the properties below are optional
//   				responseModels: map[string]iModel{
//   					"responseModelsKey": model,
//   				},
//   				responseParameters: map[string]*bool{
//   					"responseParametersKey": jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		operationName: jsii.String("operationName"),
//   		requestModels: map[string]*iModel{
//   			"requestModelsKey": model,
//   		},
//   		requestParameters: map[string]*bool{
//   			"requestParametersKey": jsii.Boolean(false),
//   		},
//   		requestValidator: requestValidator,
//   		requestValidatorOptions: &requestValidatorOptions{
//   			requestValidatorName: jsii.String("requestValidatorName"),
//   			validateRequestBody: jsii.Boolean(false),
//   			validateRequestParameters: jsii.Boolean(false),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	enabled: jsii.Boolean(false),
//   	generateDistinctId: jsii.Boolean(false),
//   	resources: []iRestApi{
//   		restApi,
//   	},
//   	value: jsii.String("value"),
//   }
//
type ApiKeyProps struct {
	// Adds a CORS preflight OPTIONS method to this resource and all child resources.
	//
	// You can add CORS at the resource-level using `addCorsPreflight`.
	DefaultCorsPreflightOptions *CorsOptions `field:"optional" json:"defaultCorsPreflightOptions" yaml:"defaultCorsPreflightOptions"`
	// An integration to use as a default for all methods created within this API unless an integration is specified.
	DefaultIntegration Integration `field:"optional" json:"defaultIntegration" yaml:"defaultIntegration"`
	// Method options to use as a default for all methods created within this API unless custom options are specified.
	DefaultMethodOptions *MethodOptions `field:"optional" json:"defaultMethodOptions" yaml:"defaultMethodOptions"`
	// A name for the API key.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the API key name.
	ApiKeyName *string `field:"optional" json:"apiKeyName" yaml:"apiKeyName"`
	// A description of the purpose of the API key.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The value of the API key.
	//
	// Must be at least 20 characters long.
	Value *string `field:"optional" json:"value" yaml:"value"`
	// An AWS Marketplace customer identifier to use when integrating with the AWS SaaS Marketplace.
	CustomerId *string `field:"optional" json:"customerId" yaml:"customerId"`
	// Indicates whether the API key can be used by clients.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Specifies whether the key identifier is distinct from the created API key value.
	GenerateDistinctId *bool `field:"optional" json:"generateDistinctId" yaml:"generateDistinctId"`
	// A list of resources this api key is associated with.
	Resources *[]IRestApi `field:"optional" json:"resources" yaml:"resources"`
}

type ApiKeySourceType string

const (
	// To read the API key from the `X-API-Key` header of a request.
	ApiKeySourceType_HEADER ApiKeySourceType = "HEADER"
	// To read the API key from the `UsageIdentifierKey` from a custom authorizer.
	ApiKeySourceType_AUTHORIZER ApiKeySourceType = "AUTHORIZER"
)

// OpenAPI specification from a local file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dockerImage dockerImage
//   var grantable iGrantable
//   var localBundling iLocalBundling
//
//   assetApiDefinition := awscdk.Aws_apigateway.NewAssetApiDefinition(jsii.String("path"), &assetOptions{
//   	assetHash: jsii.String("assetHash"),
//   	assetHashType: cdk.assetHashType_SOURCE,
//   	bundling: &bundlingOptions{
//   		image: dockerImage,
//
//   		// the properties below are optional
//   		command: []*string{
//   			jsii.String("command"),
//   		},
//   		entrypoint: []*string{
//   			jsii.String("entrypoint"),
//   		},
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		local: localBundling,
//   		network: jsii.String("network"),
//   		outputType: cdk.bundlingOutput_ARCHIVED,
//   		securityOpt: jsii.String("securityOpt"),
//   		user: jsii.String("user"),
//   		volumes: []dockerVolume{
//   			&dockerVolume{
//   				containerPath: jsii.String("containerPath"),
//   				hostPath: jsii.String("hostPath"),
//
//   				// the properties below are optional
//   				consistency: cdk.dockerVolumeConsistency_CONSISTENT,
//   			},
//   		},
//   		workingDirectory: jsii.String("workingDirectory"),
//   	},
//   	exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	followSymlinks: cdk.symlinkFollowMode_NEVER,
//   	ignoreMode: cdk.ignoreMode_GLOB,
//   	readers: []*iGrantable{
//   		grantable,
//   	},
//   })
//
type AssetApiDefinition interface {
	ApiDefinition
	// Called when the specification is initialized to allow this object to bind to the stack, add resources and have fun.
	Bind(scope constructs.Construct) *ApiDefinitionConfig
	// Called after the CFN RestApi resource has been created to allow the Api Definition to bind to it.
	//
	// Specifically it's required to allow assets to add
	// metadata for tooling like SAM CLI to be able to find their origins.
	BindAfterCreate(scope constructs.Construct, restApi IRestApi)
}

// The jsii proxy struct for AssetApiDefinition
type jsiiProxy_AssetApiDefinition struct {
	jsiiProxy_ApiDefinition
}

func NewAssetApiDefinition(path *string, options *awss3assets.AssetOptions) AssetApiDefinition {
	_init_.Initialize()

	j := jsiiProxy_AssetApiDefinition{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.AssetApiDefinition",
		[]interface{}{path, options},
		&j,
	)

	return &j
}

func NewAssetApiDefinition_Override(a AssetApiDefinition, path *string, options *awss3assets.AssetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.AssetApiDefinition",
		[]interface{}{path, options},
		a,
	)
}

// Loads the API specification from a local disk asset.
func AssetApiDefinition_FromAsset(file *string, options *awss3assets.AssetOptions) AssetApiDefinition {
	_init_.Initialize()

	var returns AssetApiDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AssetApiDefinition",
		"fromAsset",
		[]interface{}{file, options},
		&returns,
	)

	return returns
}

// Creates an API definition from a specification file in an S3 bucket.
func AssetApiDefinition_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3ApiDefinition {
	_init_.Initialize()

	var returns S3ApiDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AssetApiDefinition",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Create an API definition from an inline object.
//
// The inline object must follow the
// schema of OpenAPI 2.0 or OpenAPI 3.0
//
// Example:
//     apigateway.ApiDefinition.fromInline({
//       openapi: '3.0.2',
//       paths: {
//         '/pets': {
//           get: {
//             'responses': {
//               200: {
//                 content: {
//                   'application/json': {
//                     schema: {
//                       $ref: '#/components/schemas/Empty',
//                     },
//                   },
//                 },
//               },
//             },
//             'x-amazon-apigateway-integration': {
//               responses: {
//                 default: {
//                   statusCode: '200',
//                 },
//               },
//               requestTemplates: {
//                 'application/json': '{"statusCode": 200}',
//               },
//               passthroughBehavior: 'when_no_match',
//               type: 'mock',
//             },
//           },
//         },
//       },
//       components: {
//         schemas: {
//           Empty: {
//             title: 'Empty Schema',
//             type: 'object',
//           },
//         },
//       },
//     });
//
func AssetApiDefinition_FromInline(definition interface{}) InlineApiDefinition {
	_init_.Initialize()

	var returns InlineApiDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AssetApiDefinition",
		"fromInline",
		[]interface{}{definition},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetApiDefinition) Bind(scope constructs.Construct) *ApiDefinitionConfig {
	var returns *ApiDefinitionConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetApiDefinition) BindAfterCreate(scope constructs.Construct, restApi IRestApi) {
	_jsii_.InvokeVoid(
		a,
		"bindAfterCreate",
		[]interface{}{scope, restApi},
	)
}

// Example:
//   var books resource
//   userPool := cognito.NewUserPool(this, jsii.String("UserPool"))
//
//   auth := apigateway.NewCognitoUserPoolsAuthorizer(this, jsii.String("booksAuthorizer"), &cognitoUserPoolsAuthorizerProps{
//   	cognitoUserPools: []iUserPool{
//   		userPool,
//   	},
//   })
//   books.addMethod(jsii.String("GET"), apigateway.NewHttpIntegration(jsii.String("http://amazon.com")), &methodOptions{
//   	authorizer: auth,
//   	authorizationType: apigateway.authorizationType_COGNITO,
//   })
//
type AuthorizationType string

const (
	// Open access.
	AuthorizationType_NONE AuthorizationType = "NONE"
	// Use AWS IAM permissions.
	AuthorizationType_IAM AuthorizationType = "IAM"
	// Use a custom authorizer.
	AuthorizationType_CUSTOM AuthorizationType = "CUSTOM"
	// Use an AWS Cognito user pool.
	AuthorizationType_COGNITO AuthorizationType = "COGNITO"
)

// Base class for all custom authorizers.
type Authorizer interface {
	awscdk.Resource
	IAuthorizer
	// The authorization type of this authorizer.
	AuthorizationType() AuthorizationType
	// The authorizer ID.
	AuthorizerId() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Authorizer
type jsiiProxy_Authorizer struct {
	internal.Type__awscdkResource
	jsiiProxy_IAuthorizer
}

func (j *jsiiProxy_Authorizer) AuthorizationType() AuthorizationType {
	var returns AuthorizationType
	_jsii_.Get(
		j,
		"authorizationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authorizer) AuthorizerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authorizer) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authorizer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authorizer) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authorizer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewAuthorizer_Override(a Authorizer, scope constructs.Construct, id *string, props *awscdk.ResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.Authorizer",
		[]interface{}{scope, id, props},
		a,
	)
}

// Return whether the given object is an Authorizer.
func Authorizer_IsAuthorizer(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.Authorizer",
		"isAuthorizer",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Authorizer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.Authorizer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func Authorizer_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.Authorizer",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Authorizer_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.Authorizer",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Authorizer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (a *jsiiProxy_Authorizer) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Authorizer) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Authorizer) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Authorizer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// This type of integration lets an API expose AWS service actions.
//
// It is
// intended for calling all AWS service actions, but is not recommended for
// calling a Lambda function, because the Lambda custom integration is a legacy
// technology.
//
// Example:
//   getMessageIntegration := apigateway.NewAwsIntegration(&awsIntegrationProps{
//   	service: jsii.String("sqs"),
//   	path: jsii.String("queueName"),
//   	region: jsii.String("eu-west-1"),
//   })
//
type AwsIntegration interface {
	Integration
	// Can be overridden by subclasses to allow the integration to interact with the method being integrated, access the REST API object, method ARNs, etc.
	Bind(method Method) *IntegrationConfig
}

// The jsii proxy struct for AwsIntegration
type jsiiProxy_AwsIntegration struct {
	jsiiProxy_Integration
}

func NewAwsIntegration(props *AwsIntegrationProps) AwsIntegration {
	_init_.Initialize()

	j := jsiiProxy_AwsIntegration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.AwsIntegration",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewAwsIntegration_Override(a AwsIntegration, props *AwsIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.AwsIntegration",
		[]interface{}{props},
		a,
	)
}

func (a *jsiiProxy_AwsIntegration) Bind(method Method) *IntegrationConfig {
	var returns *IntegrationConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{method},
		&returns,
	)

	return returns
}

// Example:
//   getMessageIntegration := apigateway.NewAwsIntegration(&awsIntegrationProps{
//   	service: jsii.String("sqs"),
//   	path: jsii.String("queueName"),
//   	region: jsii.String("eu-west-1"),
//   })
//
type AwsIntegrationProps struct {
	// The name of the integrated AWS service (e.g. `s3`).
	Service *string `field:"required" json:"service" yaml:"service"`
	// The AWS action to perform in the integration.
	//
	// Use `actionParams` to specify key-value params for the action.
	//
	// Mutually exclusive with `path`.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Parameters for the action.
	//
	// `action` must be set, and `path` must be undefined.
	// The action params will be URL encoded.
	ActionParameters *map[string]*string `field:"optional" json:"actionParameters" yaml:"actionParameters"`
	// The integration's HTTP method type.
	IntegrationHttpMethod *string `field:"optional" json:"integrationHttpMethod" yaml:"integrationHttpMethod"`
	// Integration options, such as content handling, request/response mapping, etc.
	Options *IntegrationOptions `field:"optional" json:"options" yaml:"options"`
	// The path to use for path-base APIs.
	//
	// For example, for S3 GET, you can set path to `bucket/key`.
	// For lambda, you can set path to `2015-03-31/functions/${function-arn}/invocations`
	//
	// Mutually exclusive with the `action` options.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Use AWS_PROXY integration.
	Proxy *bool `field:"optional" json:"proxy" yaml:"proxy"`
	// The region of the integrated AWS service.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// A designated subdomain supported by certain AWS service for fast host-name lookup.
	Subdomain *string `field:"optional" json:"subdomain" yaml:"subdomain"`
}

// This resource creates a base path that clients who call your API must use in the invocation URL.
//
// Unless you're importing a domain with `DomainName.fromDomainNameAttributes()`,
// you can use `DomainName.addBasePathMapping()` to define mappings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var domainName domainName
//   var restApi restApi
//   var stage stage
//
//   basePathMapping := awscdk.Aws_apigateway.NewBasePathMapping(this, jsii.String("MyBasePathMapping"), &basePathMappingProps{
//   	domainName: domainName,
//   	restApi: restApi,
//
//   	// the properties below are optional
//   	basePath: jsii.String("basePath"),
//   	stage: stage,
//   })
//
type BasePathMapping interface {
	awscdk.Resource
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for BasePathMapping
type jsiiProxy_BasePathMapping struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_BasePathMapping) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BasePathMapping) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BasePathMapping) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BasePathMapping) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewBasePathMapping(scope constructs.Construct, id *string, props *BasePathMappingProps) BasePathMapping {
	_init_.Initialize()

	j := jsiiProxy_BasePathMapping{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.BasePathMapping",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewBasePathMapping_Override(b BasePathMapping, scope constructs.Construct, id *string, props *BasePathMappingProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.BasePathMapping",
		[]interface{}{scope, id, props},
		b,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func BasePathMapping_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.BasePathMapping",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func BasePathMapping_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.BasePathMapping",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func BasePathMapping_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.BasePathMapping",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BasePathMapping) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		b,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (b *jsiiProxy_BasePathMapping) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BasePathMapping) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BasePathMapping) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BasePathMapping) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Example:
//   var domain domainName
//   var api1 restApi
//   var api2 restApi
//
//
//   domain.addBasePathMapping(api1, &basePathMappingOptions{
//   	basePath: jsii.String("go-to-api1"),
//   })
//   domain.addBasePathMapping(api2, &basePathMappingOptions{
//   	basePath: jsii.String("boom"),
//   })
//
type BasePathMappingOptions struct {
	// The base path name that callers of the API must provide in the URL after the domain name (e.g. `example.com/base-path`). If you specify this property, it can't be an empty string.
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
	// The Deployment stage of API [disable-awslint:ref-via-interface].
	Stage Stage `field:"optional" json:"stage" yaml:"stage"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var domainName domainName
//   var restApi restApi
//   var stage stage
//
//   basePathMappingProps := &basePathMappingProps{
//   	domainName: domainName,
//   	restApi: restApi,
//
//   	// the properties below are optional
//   	basePath: jsii.String("basePath"),
//   	stage: stage,
//   }
//
type BasePathMappingProps struct {
	// The base path name that callers of the API must provide in the URL after the domain name (e.g. `example.com/base-path`). If you specify this property, it can't be an empty string.
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
	// The Deployment stage of API [disable-awslint:ref-via-interface].
	Stage Stage `field:"optional" json:"stage" yaml:"stage"`
	// The DomainName to associate with this base path mapping.
	DomainName IDomainName `field:"required" json:"domainName" yaml:"domainName"`
	// The RestApi resource to target.
	RestApi IRestApi `field:"required" json:"restApi" yaml:"restApi"`
}

// A CloudFormation `AWS::ApiGateway::Account`.
//
// The `AWS::ApiGateway::Account` resource specifies the IAM role that Amazon API Gateway uses to write API logs to Amazon CloudWatch Logs.
//
// > If an API Gateway resource has never been created in your AWS account , you must add a dependency on another API Gateway resource, such as an [AWS::ApiGateway::RestApi](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html) or [AWS::ApiGateway::ApiKey](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html) resource.
// >
// > If an API Gateway resource has been created in your AWS account , no dependency is required (even if the resource was deleted).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccount := awscdk.Aws_apigateway.NewCfnAccount(this, jsii.String("MyCfnAccount"), &cfnAccountProps{
//   	cloudWatchRoleArn: jsii.String("cloudWatchRoleArn"),
//   })
//
type CfnAccount interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ID for the account.
	//
	// For example: `abc123` .
	AttrId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The Amazon Resource Name (ARN) of an IAM role that has write access to CloudWatch Logs in your account.
	CloudWatchRoleArn() *string
	SetCloudWatchRoleArn(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnAccount
type jsiiProxy_CfnAccount struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAccount) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccount) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccount) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccount) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccount) CloudWatchRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudWatchRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccount) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccount) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccount) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccount) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccount) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccount) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGateway::Account`.
func NewCfnAccount(scope constructs.Construct, id *string, props *CfnAccountProps) CfnAccount {
	_init_.Initialize()

	j := jsiiProxy_CfnAccount{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnAccount",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGateway::Account`.
func NewCfnAccount_Override(c CfnAccount, scope constructs.Construct, id *string, props *CfnAccountProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnAccount",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAccount) SetCloudWatchRoleArn(val *string) {
	_jsii_.Set(
		j,
		"cloudWatchRoleArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnAccount_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnAccount",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnAccount_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnAccount",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnAccount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnAccount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAccount_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.CfnAccount",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAccount) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAccount) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAccount) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAccount) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAccount) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAccount) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAccount) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnAccount) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAccount) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAccount) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAccount) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAccount) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAccount) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAccount) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAccount) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnAccount`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccountProps := &cfnAccountProps{
//   	cloudWatchRoleArn: jsii.String("cloudWatchRoleArn"),
//   }
//
type CfnAccountProps struct {
	// The Amazon Resource Name (ARN) of an IAM role that has write access to CloudWatch Logs in your account.
	CloudWatchRoleArn *string `field:"optional" json:"cloudWatchRoleArn" yaml:"cloudWatchRoleArn"`
}

// A CloudFormation `AWS::ApiGateway::ApiKey`.
//
// The `AWS::ApiGateway::ApiKey` resource creates a unique key that you can distribute to clients who are executing API Gateway `Method` resources that require an API key. To specify which API key clients must use, map the API key with the `RestApi` and `Stage` resources that include the methods that require a key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApiKey := awscdk.Aws_apigateway.NewCfnApiKey(this, jsii.String("MyCfnApiKey"), &cfnApiKeyProps{
//   	customerId: jsii.String("customerId"),
//   	description: jsii.String("description"),
//   	enabled: jsii.Boolean(false),
//   	generateDistinctId: jsii.Boolean(false),
//   	name: jsii.String("name"),
//   	stageKeys: []interface{}{
//   		&stageKeyProperty{
//   			restApiId: jsii.String("restApiId"),
//   			stageName: jsii.String("stageName"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	value: jsii.String("value"),
//   })
//
type CfnApiKey interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ID for the API key.
	//
	// For example: `abc123` .
	AttrApiKeyId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// An AWS Marketplace customer identifier to use when integrating with the AWS SaaS Marketplace.
	CustomerId() *string
	SetCustomerId(val *string)
	// A description of the purpose of the API key.
	Description() *string
	SetDescription(val *string)
	// Indicates whether the API key can be used by clients.
	Enabled() interface{}
	SetEnabled(val interface{})
	// Specifies whether the key identifier is distinct from the created API key value.
	//
	// This parameter is deprecated and should not be used.
	GenerateDistinctId() interface{}
	SetGenerateDistinctId(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// A name for the API key.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the API key name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// A list of stages to associate with this API key.
	StageKeys() interface{}
	SetStageKeys(val interface{})
	// An array of arbitrary tags (key-value pairs) to associate with the API key.
	Tags() awscdk.TagManager
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// The value of the API key.
	//
	// Must be at least 20 characters long.
	Value() *string
	SetValue(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnApiKey
type jsiiProxy_CfnApiKey struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnApiKey) AttrApiKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrApiKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) CustomerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) GenerateDistinctId() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generateDistinctId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) StageKeys() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stageKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGateway::ApiKey`.
func NewCfnApiKey(scope constructs.Construct, id *string, props *CfnApiKeyProps) CfnApiKey {
	_init_.Initialize()

	j := jsiiProxy_CfnApiKey{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnApiKey",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGateway::ApiKey`.
func NewCfnApiKey_Override(c CfnApiKey, scope constructs.Construct, id *string, props *CfnApiKeyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnApiKey",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApiKey) SetCustomerId(val *string) {
	_jsii_.Set(
		j,
		"customerId",
		val,
	)
}

func (j *jsiiProxy_CfnApiKey) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnApiKey) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CfnApiKey) SetGenerateDistinctId(val interface{}) {
	_jsii_.Set(
		j,
		"generateDistinctId",
		val,
	)
}

func (j *jsiiProxy_CfnApiKey) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnApiKey) SetStageKeys(val interface{}) {
	_jsii_.Set(
		j,
		"stageKeys",
		val,
	)
}

func (j *jsiiProxy_CfnApiKey) SetValue(val *string) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnApiKey_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnApiKey",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnApiKey_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnApiKey",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnApiKey_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnApiKey",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApiKey_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.CfnApiKey",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApiKey) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnApiKey) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnApiKey) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnApiKey) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnApiKey) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnApiKey) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnApiKey) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnApiKey) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApiKey) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApiKey) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnApiKey) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnApiKey) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApiKey) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApiKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApiKey) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// `StageKey` is a property of the [AWS::ApiGateway::ApiKey](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html) resource that specifies the stage to associate with the API key. This association allows only clients with the key to make requests to methods in that stage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stageKeyProperty := &stageKeyProperty{
//   	restApiId: jsii.String("restApiId"),
//   	stageName: jsii.String("stageName"),
//   }
//
type CfnApiKey_StageKeyProperty struct {
	// The ID of a `RestApi` resource that includes the stage with which you want to associate the API key.
	RestApiId *string `field:"optional" json:"restApiId" yaml:"restApiId"`
	// The name of the stage with which to associate the API key.
	//
	// The stage must be included in the `RestApi` resource that you specified in the `RestApiId` property.
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
}

// Properties for defining a `CfnApiKey`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApiKeyProps := &cfnApiKeyProps{
//   	customerId: jsii.String("customerId"),
//   	description: jsii.String("description"),
//   	enabled: jsii.Boolean(false),
//   	generateDistinctId: jsii.Boolean(false),
//   	name: jsii.String("name"),
//   	stageKeys: []interface{}{
//   		&stageKeyProperty{
//   			restApiId: jsii.String("restApiId"),
//   			stageName: jsii.String("stageName"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	value: jsii.String("value"),
//   }
//
type CfnApiKeyProps struct {
	// An AWS Marketplace customer identifier to use when integrating with the AWS SaaS Marketplace.
	CustomerId *string `field:"optional" json:"customerId" yaml:"customerId"`
	// A description of the purpose of the API key.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates whether the API key can be used by clients.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Specifies whether the key identifier is distinct from the created API key value.
	//
	// This parameter is deprecated and should not be used.
	GenerateDistinctId interface{} `field:"optional" json:"generateDistinctId" yaml:"generateDistinctId"`
	// A name for the API key.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the API key name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of stages to associate with this API key.
	StageKeys interface{} `field:"optional" json:"stageKeys" yaml:"stageKeys"`
	// An array of arbitrary tags (key-value pairs) to associate with the API key.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The value of the API key.
	//
	// Must be at least 20 characters long.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

// A CloudFormation `AWS::ApiGateway::Authorizer`.
//
// The `AWS::ApiGateway::Authorizer` resource creates an authorization layer that API Gateway activates for methods that have authorization enabled. API Gateway activates the authorizer when a client calls those methods.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAuthorizer := awscdk.Aws_apigateway.NewCfnAuthorizer(this, jsii.String("MyCfnAuthorizer"), &cfnAuthorizerProps{
//   	name: jsii.String("name"),
//   	restApiId: jsii.String("restApiId"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	authorizerCredentials: jsii.String("authorizerCredentials"),
//   	authorizerResultTtlInSeconds: jsii.Number(123),
//   	authorizerUri: jsii.String("authorizerUri"),
//   	authType: jsii.String("authType"),
//   	identitySource: jsii.String("identitySource"),
//   	identityValidationExpression: jsii.String("identityValidationExpression"),
//   	providerArns: []*string{
//   		jsii.String("providerArns"),
//   	},
//   })
//
type CfnAuthorizer interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ID for the authorizer.
	//
	// For example: `abc123` .
	AttrAuthorizerId() *string
	// The credentials that are required for the authorizer.
	//
	// To specify an IAM role that API Gateway assumes, specify the role's Amazon Resource Name (ARN). To use resource-based permissions on the Lambda function, specify null.
	AuthorizerCredentials() *string
	SetAuthorizerCredentials(val *string)
	// The time-to-live (TTL) period, in seconds, that specifies how long API Gateway caches authorizer results.
	//
	// If you specify a value greater than 0, API Gateway caches the authorizer responses. By default, API Gateway sets this property to 300. The maximum value is 3600, or 1 hour.
	AuthorizerResultTtlInSeconds() *float64
	SetAuthorizerResultTtlInSeconds(val *float64)
	// The authorizer's Uniform Resource Identifier (URI).
	//
	// If you specify `TOKEN` for the authorizer's `Type` property, specify a Lambda function URI that has the form `arn:aws:apigateway: *region* :lambda:path/ *path*` . The path usually has the form /2015-03-31/functions/ *LambdaFunctionARN* /invocations.
	AuthorizerUri() *string
	SetAuthorizerUri(val *string)
	// An optional customer-defined field that's used in OpenApi imports and exports without functional impact.
	AuthType() *string
	SetAuthType(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The source of the identity in an incoming request.
	//
	// If you specify `TOKEN` or `COGNITO_USER_POOLS` for the `Type` property, this property is required. Specify a header mapping expression using the form `method.request.header. *name*` , where *name* is the name of a custom authorization header that clients submit as part of their requests.
	//
	// If you specify `REQUEST` for the `Type` property, this property is required when authorization caching is enabled. Specify a comma-separated string of one or more mapping expressions of the specified request parameter using the form `method.request.parameter. *name*` . For supported parameter types, see [Configure Lambda Authorizer Using the API Gateway Console](https://docs.aws.amazon.com/apigateway/latest/developerguide/configure-api-gateway-lambda-authorization-with-console.html) in the *API Gateway Developer Guide* .
	IdentitySource() *string
	SetIdentitySource(val *string)
	// A validation expression for the incoming identity.
	//
	// If you specify `TOKEN` for the authorizer's `Type` property, specify a regular expression. API Gateway uses the expression to attempt to match the incoming client token, and proceeds if the token matches. If the token doesn't match, API Gateway responds with a 401 (unauthorized request) error code.
	IdentityValidationExpression() *string
	SetIdentityValidationExpression(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The name of the authorizer.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// A list of the Amazon Cognito user pool Amazon Resource Names (ARNs) to associate with this authorizer.
	//
	// Required if you specify `COGNITO_USER_POOLS` as the authorizer `Type` . For more information, see [Use Amazon Cognito User Pools](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-integrate-with-cognito.html#apigateway-enable-cognito-user-pool) in the *API Gateway Developer Guide* .
	ProviderArns() *[]*string
	SetProviderArns(val *[]*string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The ID of the `RestApi` resource that API Gateway creates the authorizer in.
	RestApiId() *string
	SetRestApiId(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The type of authorizer. Valid values include:.
	//
	// - `TOKEN` : A custom authorizer that uses a Lambda function.
	// - `COGNITO_USER_POOLS` : An authorizer that uses Amazon Cognito user pools.
	// - `REQUEST` : An authorizer that uses a Lambda function using incoming request parameters.
	Type() *string
	SetType(val *string)
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnAuthorizer
type jsiiProxy_CfnAuthorizer struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAuthorizer) AttrAuthorizerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAuthorizerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) AuthorizerCredentials() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) AuthorizerResultTtlInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authorizerResultTtlInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) AuthorizerUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) AuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) IdentitySource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identitySource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) IdentityValidationExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityValidationExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) ProviderArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"providerArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGateway::Authorizer`.
func NewCfnAuthorizer(scope constructs.Construct, id *string, props *CfnAuthorizerProps) CfnAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_CfnAuthorizer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnAuthorizer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGateway::Authorizer`.
func NewCfnAuthorizer_Override(c CfnAuthorizer, scope constructs.Construct, id *string, props *CfnAuthorizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnAuthorizer",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAuthorizer) SetAuthorizerCredentials(val *string) {
	_jsii_.Set(
		j,
		"authorizerCredentials",
		val,
	)
}

func (j *jsiiProxy_CfnAuthorizer) SetAuthorizerResultTtlInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"authorizerResultTtlInSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnAuthorizer) SetAuthorizerUri(val *string) {
	_jsii_.Set(
		j,
		"authorizerUri",
		val,
	)
}

func (j *jsiiProxy_CfnAuthorizer) SetAuthType(val *string) {
	_jsii_.Set(
		j,
		"authType",
		val,
	)
}

func (j *jsiiProxy_CfnAuthorizer) SetIdentitySource(val *string) {
	_jsii_.Set(
		j,
		"identitySource",
		val,
	)
}

func (j *jsiiProxy_CfnAuthorizer) SetIdentityValidationExpression(val *string) {
	_jsii_.Set(
		j,
		"identityValidationExpression",
		val,
	)
}

func (j *jsiiProxy_CfnAuthorizer) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnAuthorizer) SetProviderArns(val *[]*string) {
	_jsii_.Set(
		j,
		"providerArns",
		val,
	)
}

func (j *jsiiProxy_CfnAuthorizer) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

func (j *jsiiProxy_CfnAuthorizer) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnAuthorizer_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnAuthorizer",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnAuthorizer_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnAuthorizer",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnAuthorizer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnAuthorizer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAuthorizer_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.CfnAuthorizer",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAuthorizer) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAuthorizer) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAuthorizer) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAuthorizer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAuthorizer) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAuthorizer) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAuthorizer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnAuthorizer) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAuthorizer) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAuthorizer) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAuthorizer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAuthorizer) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAuthorizer) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAuthorizer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAuthorizer) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnAuthorizer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAuthorizerProps := &cfnAuthorizerProps{
//   	name: jsii.String("name"),
//   	restApiId: jsii.String("restApiId"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	authorizerCredentials: jsii.String("authorizerCredentials"),
//   	authorizerResultTtlInSeconds: jsii.Number(123),
//   	authorizerUri: jsii.String("authorizerUri"),
//   	authType: jsii.String("authType"),
//   	identitySource: jsii.String("identitySource"),
//   	identityValidationExpression: jsii.String("identityValidationExpression"),
//   	providerArns: []*string{
//   		jsii.String("providerArns"),
//   	},
//   }
//
type CfnAuthorizerProps struct {
	// The name of the authorizer.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the `RestApi` resource that API Gateway creates the authorizer in.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// The type of authorizer. Valid values include:.
	//
	// - `TOKEN` : A custom authorizer that uses a Lambda function.
	// - `COGNITO_USER_POOLS` : An authorizer that uses Amazon Cognito user pools.
	// - `REQUEST` : An authorizer that uses a Lambda function using incoming request parameters.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The credentials that are required for the authorizer.
	//
	// To specify an IAM role that API Gateway assumes, specify the role's Amazon Resource Name (ARN). To use resource-based permissions on the Lambda function, specify null.
	AuthorizerCredentials *string `field:"optional" json:"authorizerCredentials" yaml:"authorizerCredentials"`
	// The time-to-live (TTL) period, in seconds, that specifies how long API Gateway caches authorizer results.
	//
	// If you specify a value greater than 0, API Gateway caches the authorizer responses. By default, API Gateway sets this property to 300. The maximum value is 3600, or 1 hour.
	AuthorizerResultTtlInSeconds *float64 `field:"optional" json:"authorizerResultTtlInSeconds" yaml:"authorizerResultTtlInSeconds"`
	// The authorizer's Uniform Resource Identifier (URI).
	//
	// If you specify `TOKEN` for the authorizer's `Type` property, specify a Lambda function URI that has the form `arn:aws:apigateway: *region* :lambda:path/ *path*` . The path usually has the form /2015-03-31/functions/ *LambdaFunctionARN* /invocations.
	AuthorizerUri *string `field:"optional" json:"authorizerUri" yaml:"authorizerUri"`
	// An optional customer-defined field that's used in OpenApi imports and exports without functional impact.
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// The source of the identity in an incoming request.
	//
	// If you specify `TOKEN` or `COGNITO_USER_POOLS` for the `Type` property, this property is required. Specify a header mapping expression using the form `method.request.header. *name*` , where *name* is the name of a custom authorization header that clients submit as part of their requests.
	//
	// If you specify `REQUEST` for the `Type` property, this property is required when authorization caching is enabled. Specify a comma-separated string of one or more mapping expressions of the specified request parameter using the form `method.request.parameter. *name*` . For supported parameter types, see [Configure Lambda Authorizer Using the API Gateway Console](https://docs.aws.amazon.com/apigateway/latest/developerguide/configure-api-gateway-lambda-authorization-with-console.html) in the *API Gateway Developer Guide* .
	IdentitySource *string `field:"optional" json:"identitySource" yaml:"identitySource"`
	// A validation expression for the incoming identity.
	//
	// If you specify `TOKEN` for the authorizer's `Type` property, specify a regular expression. API Gateway uses the expression to attempt to match the incoming client token, and proceeds if the token matches. If the token doesn't match, API Gateway responds with a 401 (unauthorized request) error code.
	IdentityValidationExpression *string `field:"optional" json:"identityValidationExpression" yaml:"identityValidationExpression"`
	// A list of the Amazon Cognito user pool Amazon Resource Names (ARNs) to associate with this authorizer.
	//
	// Required if you specify `COGNITO_USER_POOLS` as the authorizer `Type` . For more information, see [Use Amazon Cognito User Pools](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-integrate-with-cognito.html#apigateway-enable-cognito-user-pool) in the *API Gateway Developer Guide* .
	ProviderArns *[]*string `field:"optional" json:"providerArns" yaml:"providerArns"`
}

// A CloudFormation `AWS::ApiGateway::BasePathMapping`.
//
// The `AWS::ApiGateway::BasePathMapping` resource creates a base path that clients who call your API must use in the invocation URL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBasePathMapping := awscdk.Aws_apigateway.NewCfnBasePathMapping(this, jsii.String("MyCfnBasePathMapping"), &cfnBasePathMappingProps{
//   	domainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	basePath: jsii.String("basePath"),
//   	id: jsii.String("id"),
//   	restApiId: jsii.String("restApiId"),
//   	stage: jsii.String("stage"),
//   })
//
type CfnBasePathMapping interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The base path name that callers of the API must provide in the URL after the domain name.
	BasePath() *string
	SetBasePath(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The `DomainName` of an [AWS::ApiGateway::DomainName](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html) resource.
	DomainName() *string
	SetDomainName(val *string)
	// `AWS::ApiGateway::BasePathMapping.Id`.
	Id() *string
	SetId(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The ID of the API.
	RestApiId() *string
	SetRestApiId(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The name of the API's stage.
	Stage() *string
	SetStage(val *string)
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnBasePathMapping
type jsiiProxy_CfnBasePathMapping struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnBasePathMapping) BasePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBasePathMapping) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBasePathMapping) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBasePathMapping) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBasePathMapping) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBasePathMapping) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBasePathMapping) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBasePathMapping) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBasePathMapping) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBasePathMapping) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBasePathMapping) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBasePathMapping) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBasePathMapping) Stage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBasePathMapping) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBasePathMapping) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGateway::BasePathMapping`.
func NewCfnBasePathMapping(scope constructs.Construct, id *string, props *CfnBasePathMappingProps) CfnBasePathMapping {
	_init_.Initialize()

	j := jsiiProxy_CfnBasePathMapping{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnBasePathMapping",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGateway::BasePathMapping`.
func NewCfnBasePathMapping_Override(c CfnBasePathMapping, scope constructs.Construct, id *string, props *CfnBasePathMappingProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnBasePathMapping",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnBasePathMapping) SetBasePath(val *string) {
	_jsii_.Set(
		j,
		"basePath",
		val,
	)
}

func (j *jsiiProxy_CfnBasePathMapping) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_CfnBasePathMapping) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CfnBasePathMapping) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

func (j *jsiiProxy_CfnBasePathMapping) SetStage(val *string) {
	_jsii_.Set(
		j,
		"stage",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnBasePathMapping_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnBasePathMapping",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnBasePathMapping_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnBasePathMapping",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnBasePathMapping_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnBasePathMapping",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBasePathMapping_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.CfnBasePathMapping",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBasePathMapping) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnBasePathMapping) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnBasePathMapping) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnBasePathMapping) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnBasePathMapping) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnBasePathMapping) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnBasePathMapping) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnBasePathMapping) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBasePathMapping) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBasePathMapping) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnBasePathMapping) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnBasePathMapping) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBasePathMapping) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBasePathMapping) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBasePathMapping) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnBasePathMapping`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBasePathMappingProps := &cfnBasePathMappingProps{
//   	domainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	basePath: jsii.String("basePath"),
//   	id: jsii.String("id"),
//   	restApiId: jsii.String("restApiId"),
//   	stage: jsii.String("stage"),
//   }
//
type CfnBasePathMappingProps struct {
	// The `DomainName` of an [AWS::ApiGateway::DomainName](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html) resource.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The base path name that callers of the API must provide in the URL after the domain name.
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
	// `AWS::ApiGateway::BasePathMapping.Id`.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ID of the API.
	RestApiId *string `field:"optional" json:"restApiId" yaml:"restApiId"`
	// The name of the API's stage.
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
}

// A CloudFormation `AWS::ApiGateway::ClientCertificate`.
//
// The `AWS::ApiGateway::ClientCertificate` resource creates a client certificate that API Gateway uses to configure client-side SSL authentication for sending requests to the integration endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClientCertificate := awscdk.Aws_apigateway.NewCfnClientCertificate(this, jsii.String("MyCfnClientCertificate"), &cfnClientCertificateProps{
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnClientCertificate interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ID for the client certificate.
	//
	// For example: `abc123` .
	AttrClientCertificateId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A description of the client certificate.
	Description() *string
	SetDescription(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of arbitrary tags (key-value pairs) to associate with the client certificate.
	Tags() awscdk.TagManager
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnClientCertificate
type jsiiProxy_CfnClientCertificate struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnClientCertificate) AttrClientCertificateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrClientCertificateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientCertificate) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientCertificate) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientCertificate) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientCertificate) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientCertificate) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientCertificate) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientCertificate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientCertificate) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientCertificate) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientCertificate) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientCertificate) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientCertificate) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGateway::ClientCertificate`.
func NewCfnClientCertificate(scope constructs.Construct, id *string, props *CfnClientCertificateProps) CfnClientCertificate {
	_init_.Initialize()

	j := jsiiProxy_CfnClientCertificate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnClientCertificate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGateway::ClientCertificate`.
func NewCfnClientCertificate_Override(c CfnClientCertificate, scope constructs.Construct, id *string, props *CfnClientCertificateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnClientCertificate",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnClientCertificate) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnClientCertificate_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnClientCertificate",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnClientCertificate_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnClientCertificate",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnClientCertificate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnClientCertificate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnClientCertificate_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.CfnClientCertificate",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnClientCertificate) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnClientCertificate) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnClientCertificate) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnClientCertificate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnClientCertificate) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnClientCertificate) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnClientCertificate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnClientCertificate) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClientCertificate) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClientCertificate) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnClientCertificate) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnClientCertificate) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClientCertificate) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClientCertificate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClientCertificate) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnClientCertificate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClientCertificateProps := &cfnClientCertificateProps{
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnClientCertificateProps struct {
	// A description of the client certificate.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of arbitrary tags (key-value pairs) to associate with the client certificate.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::ApiGateway::Deployment`.
//
// The `AWS::ApiGateway::Deployment` resource deploys an API Gateway `RestApi` resource to a stage so that clients can call the API over the internet. The stage acts as an environment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeployment := awscdk.Aws_apigateway.NewCfnDeployment(this, jsii.String("MyCfnDeployment"), &cfnDeploymentProps{
//   	restApiId: jsii.String("restApiId"),
//
//   	// the properties below are optional
//   	deploymentCanarySettings: &deploymentCanarySettingsProperty{
//   		percentTraffic: jsii.Number(123),
//   		stageVariableOverrides: map[string]*string{
//   			"stageVariableOverridesKey": jsii.String("stageVariableOverrides"),
//   		},
//   		useStageCache: jsii.Boolean(false),
//   	},
//   	description: jsii.String("description"),
//   	stageDescription: &stageDescriptionProperty{
//   		accessLogSetting: &accessLogSettingProperty{
//   			destinationArn: jsii.String("destinationArn"),
//   			format: jsii.String("format"),
//   		},
//   		cacheClusterEnabled: jsii.Boolean(false),
//   		cacheClusterSize: jsii.String("cacheClusterSize"),
//   		cacheDataEncrypted: jsii.Boolean(false),
//   		cacheTtlInSeconds: jsii.Number(123),
//   		cachingEnabled: jsii.Boolean(false),
//   		canarySetting: &canarySettingProperty{
//   			percentTraffic: jsii.Number(123),
//   			stageVariableOverrides: map[string]*string{
//   				"stageVariableOverridesKey": jsii.String("stageVariableOverrides"),
//   			},
//   			useStageCache: jsii.Boolean(false),
//   		},
//   		clientCertificateId: jsii.String("clientCertificateId"),
//   		dataTraceEnabled: jsii.Boolean(false),
//   		description: jsii.String("description"),
//   		documentationVersion: jsii.String("documentationVersion"),
//   		loggingLevel: jsii.String("loggingLevel"),
//   		methodSettings: []interface{}{
//   			&methodSettingProperty{
//   				cacheDataEncrypted: jsii.Boolean(false),
//   				cacheTtlInSeconds: jsii.Number(123),
//   				cachingEnabled: jsii.Boolean(false),
//   				dataTraceEnabled: jsii.Boolean(false),
//   				httpMethod: jsii.String("httpMethod"),
//   				loggingLevel: jsii.String("loggingLevel"),
//   				metricsEnabled: jsii.Boolean(false),
//   				resourcePath: jsii.String("resourcePath"),
//   				throttlingBurstLimit: jsii.Number(123),
//   				throttlingRateLimit: jsii.Number(123),
//   			},
//   		},
//   		metricsEnabled: jsii.Boolean(false),
//   		tags: []cfnTag{
//   			&cfnTag{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		throttlingBurstLimit: jsii.Number(123),
//   		throttlingRateLimit: jsii.Number(123),
//   		tracingEnabled: jsii.Boolean(false),
//   		variables: map[string]*string{
//   			"variablesKey": jsii.String("variables"),
//   		},
//   	},
//   	stageName: jsii.String("stageName"),
//   })
//
type CfnDeployment interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ID for the deployment.
	//
	// For example: `abc123` .
	AttrDeploymentId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Specifies settings for the canary deployment.
	DeploymentCanarySettings() interface{}
	SetDeploymentCanarySettings(val interface{})
	// A description of the purpose of the API Gateway deployment.
	Description() *string
	SetDescription(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The ID of the `RestApi` resource to deploy.
	RestApiId() *string
	SetRestApiId(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Configures the stage that API Gateway creates with this deployment.
	StageDescription() interface{}
	SetStageDescription(val interface{})
	// A name for the stage that API Gateway creates with this deployment.
	//
	// Use only alphanumeric characters.
	StageName() *string
	SetStageName(val *string)
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDeployment
type jsiiProxy_CfnDeployment struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDeployment) AttrDeploymentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDeploymentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeployment) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeployment) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeployment) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeployment) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeployment) DeploymentCanarySettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentCanarySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeployment) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeployment) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeployment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeployment) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeployment) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeployment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeployment) StageDescription() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stageDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeployment) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeployment) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeployment) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGateway::Deployment`.
func NewCfnDeployment(scope constructs.Construct, id *string, props *CfnDeploymentProps) CfnDeployment {
	_init_.Initialize()

	j := jsiiProxy_CfnDeployment{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnDeployment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGateway::Deployment`.
func NewCfnDeployment_Override(c CfnDeployment, scope constructs.Construct, id *string, props *CfnDeploymentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnDeployment",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDeployment) SetDeploymentCanarySettings(val interface{}) {
	_jsii_.Set(
		j,
		"deploymentCanarySettings",
		val,
	)
}

func (j *jsiiProxy_CfnDeployment) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnDeployment) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

func (j *jsiiProxy_CfnDeployment) SetStageDescription(val interface{}) {
	_jsii_.Set(
		j,
		"stageDescription",
		val,
	)
}

func (j *jsiiProxy_CfnDeployment) SetStageName(val *string) {
	_jsii_.Set(
		j,
		"stageName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDeployment_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnDeployment",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDeployment_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnDeployment",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnDeployment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnDeployment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDeployment_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.CfnDeployment",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDeployment) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDeployment) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDeployment) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDeployment) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDeployment) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDeployment) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDeployment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDeployment) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeployment) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeployment) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDeployment) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDeployment) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeployment) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeployment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeployment) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The `AccessLogSetting` property type specifies settings for logging access in this stage.
//
// `AccessLogSetting` is a property of the [StageDescription](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessLogSettingProperty := &accessLogSettingProperty{
//   	destinationArn: jsii.String("destinationArn"),
//   	format: jsii.String("format"),
//   }
//
type CfnDeployment_AccessLogSettingProperty struct {
	// The Amazon Resource Name (ARN) of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs.
	//
	// If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with `amazon-apigateway-` .
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
	// A single line format of the access logs of data, as specified by selected [$context variables](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html#context-variable-reference) . The format must include at least `$context.requestId` .
	Format *string `field:"optional" json:"format" yaml:"format"`
}

// The `CanarySetting` property type specifies settings for the canary deployment in this stage.
//
// `CanarySetting` is a property of the [StageDescription](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   canarySettingProperty := &canarySettingProperty{
//   	percentTraffic: jsii.Number(123),
//   	stageVariableOverrides: map[string]*string{
//   		"stageVariableOverridesKey": jsii.String("stageVariableOverrides"),
//   	},
//   	useStageCache: jsii.Boolean(false),
//   }
//
type CfnDeployment_CanarySettingProperty struct {
	// The percent (0-100) of traffic diverted to a canary deployment.
	PercentTraffic *float64 `field:"optional" json:"percentTraffic" yaml:"percentTraffic"`
	// Stage variables overridden for a canary release deployment, including new stage variables introduced in the canary.
	//
	// These stage variables are represented as a string-to-string map between stage variable names and their values.
	StageVariableOverrides interface{} `field:"optional" json:"stageVariableOverrides" yaml:"stageVariableOverrides"`
	// Whether the canary deployment uses the stage cache or not.
	UseStageCache interface{} `field:"optional" json:"useStageCache" yaml:"useStageCache"`
}

// The `DeploymentCanarySettings` property type specifies settings for the canary deployment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentCanarySettingsProperty := &deploymentCanarySettingsProperty{
//   	percentTraffic: jsii.Number(123),
//   	stageVariableOverrides: map[string]*string{
//   		"stageVariableOverridesKey": jsii.String("stageVariableOverrides"),
//   	},
//   	useStageCache: jsii.Boolean(false),
//   }
//
type CfnDeployment_DeploymentCanarySettingsProperty struct {
	// The percentage (0-100) of traffic diverted to a canary deployment.
	PercentTraffic *float64 `field:"optional" json:"percentTraffic" yaml:"percentTraffic"`
	// Stage variables overridden for a canary release deployment, including new stage variables introduced in the canary.
	//
	// These stage variables are represented as a string-to-string map between stage variable names and their values.
	//
	// Duplicates are not allowed.
	StageVariableOverrides interface{} `field:"optional" json:"stageVariableOverrides" yaml:"stageVariableOverrides"`
	// Whether the canary deployment uses the stage cache.
	UseStageCache interface{} `field:"optional" json:"useStageCache" yaml:"useStageCache"`
}

// The `MethodSetting` property type configures settings for all methods in a stage.
//
// The `MethodSettings` property of the [Amazon API Gateway Deployment StageDescription](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription.html) property type contains a list of `MethodSetting` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   methodSettingProperty := &methodSettingProperty{
//   	cacheDataEncrypted: jsii.Boolean(false),
//   	cacheTtlInSeconds: jsii.Number(123),
//   	cachingEnabled: jsii.Boolean(false),
//   	dataTraceEnabled: jsii.Boolean(false),
//   	httpMethod: jsii.String("httpMethod"),
//   	loggingLevel: jsii.String("loggingLevel"),
//   	metricsEnabled: jsii.Boolean(false),
//   	resourcePath: jsii.String("resourcePath"),
//   	throttlingBurstLimit: jsii.Number(123),
//   	throttlingRateLimit: jsii.Number(123),
//   }
//
type CfnDeployment_MethodSettingProperty struct {
	// Indicates whether the cached responses are encrypted.
	CacheDataEncrypted interface{} `field:"optional" json:"cacheDataEncrypted" yaml:"cacheDataEncrypted"`
	// The time-to-live (TTL) period, in seconds, that specifies how long API Gateway caches responses.
	CacheTtlInSeconds *float64 `field:"optional" json:"cacheTtlInSeconds" yaml:"cacheTtlInSeconds"`
	// Indicates whether responses are cached and returned for requests.
	//
	// You must enable a cache cluster on the stage to cache responses. For more information, see [Enable API Gateway Caching in a Stage to Enhance API Performance](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-caching.html) in the *API Gateway Developer Guide* .
	CachingEnabled interface{} `field:"optional" json:"cachingEnabled" yaml:"cachingEnabled"`
	// Indicates whether data trace logging is enabled for methods in the stage.
	//
	// API Gateway pushes these logs to Amazon CloudWatch Logs.
	DataTraceEnabled interface{} `field:"optional" json:"dataTraceEnabled" yaml:"dataTraceEnabled"`
	// The HTTP method.
	HttpMethod *string `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// The logging level for this method.
	//
	// For valid values, see the `loggingLevel` property of the [Stage](https://docs.aws.amazon.com/apigateway/api-reference/resource/stage/#loggingLevel) resource in the *Amazon API Gateway API Reference* .
	LoggingLevel *string `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
	// Indicates whether Amazon CloudWatch metrics are enabled for methods in the stage.
	MetricsEnabled interface{} `field:"optional" json:"metricsEnabled" yaml:"metricsEnabled"`
	// The resource path for this method.
	//
	// Forward slashes ( `/` ) are encoded as `~1` and the initial slash must include a forward slash. For example, the path value `/resource/subresource` must be encoded as `/~1resource~1subresource` . To specify the root path, use only a slash ( `/` ).
	ResourcePath *string `field:"optional" json:"resourcePath" yaml:"resourcePath"`
	// The number of burst requests per second that API Gateway permits across all APIs, stages, and methods in your AWS account .
	//
	// For more information, see [Manage API Request Throttling](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html) in the *API Gateway Developer Guide* .
	ThrottlingBurstLimit *float64 `field:"optional" json:"throttlingBurstLimit" yaml:"throttlingBurstLimit"`
	// The number of steady-state requests per second that API Gateway permits across all APIs, stages, and methods in your AWS account .
	//
	// For more information, see [Manage API Request Throttling](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html) in the *API Gateway Developer Guide* .
	ThrottlingRateLimit *float64 `field:"optional" json:"throttlingRateLimit" yaml:"throttlingRateLimit"`
}

// `StageDescription` is a property of the [AWS::ApiGateway::Deployment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html) resource that configures a deployment stage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stageDescriptionProperty := &stageDescriptionProperty{
//   	accessLogSetting: &accessLogSettingProperty{
//   		destinationArn: jsii.String("destinationArn"),
//   		format: jsii.String("format"),
//   	},
//   	cacheClusterEnabled: jsii.Boolean(false),
//   	cacheClusterSize: jsii.String("cacheClusterSize"),
//   	cacheDataEncrypted: jsii.Boolean(false),
//   	cacheTtlInSeconds: jsii.Number(123),
//   	cachingEnabled: jsii.Boolean(false),
//   	canarySetting: &canarySettingProperty{
//   		percentTraffic: jsii.Number(123),
//   		stageVariableOverrides: map[string]*string{
//   			"stageVariableOverridesKey": jsii.String("stageVariableOverrides"),
//   		},
//   		useStageCache: jsii.Boolean(false),
//   	},
//   	clientCertificateId: jsii.String("clientCertificateId"),
//   	dataTraceEnabled: jsii.Boolean(false),
//   	description: jsii.String("description"),
//   	documentationVersion: jsii.String("documentationVersion"),
//   	loggingLevel: jsii.String("loggingLevel"),
//   	methodSettings: []interface{}{
//   		&methodSettingProperty{
//   			cacheDataEncrypted: jsii.Boolean(false),
//   			cacheTtlInSeconds: jsii.Number(123),
//   			cachingEnabled: jsii.Boolean(false),
//   			dataTraceEnabled: jsii.Boolean(false),
//   			httpMethod: jsii.String("httpMethod"),
//   			loggingLevel: jsii.String("loggingLevel"),
//   			metricsEnabled: jsii.Boolean(false),
//   			resourcePath: jsii.String("resourcePath"),
//   			throttlingBurstLimit: jsii.Number(123),
//   			throttlingRateLimit: jsii.Number(123),
//   		},
//   	},
//   	metricsEnabled: jsii.Boolean(false),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	throttlingBurstLimit: jsii.Number(123),
//   	throttlingRateLimit: jsii.Number(123),
//   	tracingEnabled: jsii.Boolean(false),
//   	variables: map[string]*string{
//   		"variablesKey": jsii.String("variables"),
//   	},
//   }
//
type CfnDeployment_StageDescriptionProperty struct {
	// Specifies settings for logging access in this stage.
	AccessLogSetting interface{} `field:"optional" json:"accessLogSetting" yaml:"accessLogSetting"`
	// Indicates whether cache clustering is enabled for the stage.
	CacheClusterEnabled interface{} `field:"optional" json:"cacheClusterEnabled" yaml:"cacheClusterEnabled"`
	// The size of the stage's cache cluster.
	CacheClusterSize *string `field:"optional" json:"cacheClusterSize" yaml:"cacheClusterSize"`
	// Indicates whether the cached responses are encrypted.
	CacheDataEncrypted interface{} `field:"optional" json:"cacheDataEncrypted" yaml:"cacheDataEncrypted"`
	// The time-to-live (TTL) period, in seconds, that specifies how long API Gateway caches responses.
	CacheTtlInSeconds *float64 `field:"optional" json:"cacheTtlInSeconds" yaml:"cacheTtlInSeconds"`
	// Indicates whether responses are cached and returned for requests.
	//
	// You must enable a cache cluster on the stage to cache responses. For more information, see [Enable API Gateway Caching in a Stage to Enhance API Performance](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-caching.html) in the *API Gateway Developer Guide* .
	CachingEnabled interface{} `field:"optional" json:"cachingEnabled" yaml:"cachingEnabled"`
	// Specifies settings for the canary deployment in this stage.
	CanarySetting interface{} `field:"optional" json:"canarySetting" yaml:"canarySetting"`
	// The identifier of the client certificate that API Gateway uses to call your integration endpoints in the stage.
	ClientCertificateId *string `field:"optional" json:"clientCertificateId" yaml:"clientCertificateId"`
	// Indicates whether data trace logging is enabled for methods in the stage.
	//
	// API Gateway pushes these logs to Amazon CloudWatch Logs.
	DataTraceEnabled interface{} `field:"optional" json:"dataTraceEnabled" yaml:"dataTraceEnabled"`
	// A description of the purpose of the stage.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The version identifier of the API documentation snapshot.
	DocumentationVersion *string `field:"optional" json:"documentationVersion" yaml:"documentationVersion"`
	// The logging level for this method.
	//
	// For valid values, see the `loggingLevel` property of the [Stage](https://docs.aws.amazon.com/apigateway/api-reference/resource/stage/#loggingLevel) resource in the *Amazon API Gateway API Reference* .
	LoggingLevel *string `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
	// Configures settings for all of the stage's methods.
	MethodSettings interface{} `field:"optional" json:"methodSettings" yaml:"methodSettings"`
	// Indicates whether Amazon CloudWatch metrics are enabled for methods in the stage.
	MetricsEnabled interface{} `field:"optional" json:"metricsEnabled" yaml:"metricsEnabled"`
	// An array of arbitrary tags (key-value pairs) to associate with the stage.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The target request burst rate limit.
	//
	// This allows more requests through for a period of time than the target rate limit. For more information, see [Manage API Request Throttling](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html) in the *API Gateway Developer Guide* .
	ThrottlingBurstLimit *float64 `field:"optional" json:"throttlingBurstLimit" yaml:"throttlingBurstLimit"`
	// The target request steady-state rate limit.
	//
	// For more information, see [Manage API Request Throttling](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html) in the *API Gateway Developer Guide* .
	ThrottlingRateLimit *float64 `field:"optional" json:"throttlingRateLimit" yaml:"throttlingRateLimit"`
	// Specifies whether active tracing with X-ray is enabled for this stage.
	//
	// For more information, see [Trace API Gateway API Execution with AWS X-Ray](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-xray.html) in the *API Gateway Developer Guide* .
	TracingEnabled interface{} `field:"optional" json:"tracingEnabled" yaml:"tracingEnabled"`
	// A map that defines the stage variables.
	//
	// Variable names must consist of alphanumeric characters, and the values must match the following regular expression: `[A-Za-z0-9-._~:/?#&=,]+` .
	Variables interface{} `field:"optional" json:"variables" yaml:"variables"`
}

// Properties for defining a `CfnDeployment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeploymentProps := &cfnDeploymentProps{
//   	restApiId: jsii.String("restApiId"),
//
//   	// the properties below are optional
//   	deploymentCanarySettings: &deploymentCanarySettingsProperty{
//   		percentTraffic: jsii.Number(123),
//   		stageVariableOverrides: map[string]*string{
//   			"stageVariableOverridesKey": jsii.String("stageVariableOverrides"),
//   		},
//   		useStageCache: jsii.Boolean(false),
//   	},
//   	description: jsii.String("description"),
//   	stageDescription: &stageDescriptionProperty{
//   		accessLogSetting: &accessLogSettingProperty{
//   			destinationArn: jsii.String("destinationArn"),
//   			format: jsii.String("format"),
//   		},
//   		cacheClusterEnabled: jsii.Boolean(false),
//   		cacheClusterSize: jsii.String("cacheClusterSize"),
//   		cacheDataEncrypted: jsii.Boolean(false),
//   		cacheTtlInSeconds: jsii.Number(123),
//   		cachingEnabled: jsii.Boolean(false),
//   		canarySetting: &canarySettingProperty{
//   			percentTraffic: jsii.Number(123),
//   			stageVariableOverrides: map[string]*string{
//   				"stageVariableOverridesKey": jsii.String("stageVariableOverrides"),
//   			},
//   			useStageCache: jsii.Boolean(false),
//   		},
//   		clientCertificateId: jsii.String("clientCertificateId"),
//   		dataTraceEnabled: jsii.Boolean(false),
//   		description: jsii.String("description"),
//   		documentationVersion: jsii.String("documentationVersion"),
//   		loggingLevel: jsii.String("loggingLevel"),
//   		methodSettings: []interface{}{
//   			&methodSettingProperty{
//   				cacheDataEncrypted: jsii.Boolean(false),
//   				cacheTtlInSeconds: jsii.Number(123),
//   				cachingEnabled: jsii.Boolean(false),
//   				dataTraceEnabled: jsii.Boolean(false),
//   				httpMethod: jsii.String("httpMethod"),
//   				loggingLevel: jsii.String("loggingLevel"),
//   				metricsEnabled: jsii.Boolean(false),
//   				resourcePath: jsii.String("resourcePath"),
//   				throttlingBurstLimit: jsii.Number(123),
//   				throttlingRateLimit: jsii.Number(123),
//   			},
//   		},
//   		metricsEnabled: jsii.Boolean(false),
//   		tags: []cfnTag{
//   			&cfnTag{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		throttlingBurstLimit: jsii.Number(123),
//   		throttlingRateLimit: jsii.Number(123),
//   		tracingEnabled: jsii.Boolean(false),
//   		variables: map[string]*string{
//   			"variablesKey": jsii.String("variables"),
//   		},
//   	},
//   	stageName: jsii.String("stageName"),
//   }
//
type CfnDeploymentProps struct {
	// The ID of the `RestApi` resource to deploy.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// Specifies settings for the canary deployment.
	DeploymentCanarySettings interface{} `field:"optional" json:"deploymentCanarySettings" yaml:"deploymentCanarySettings"`
	// A description of the purpose of the API Gateway deployment.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Configures the stage that API Gateway creates with this deployment.
	StageDescription interface{} `field:"optional" json:"stageDescription" yaml:"stageDescription"`
	// A name for the stage that API Gateway creates with this deployment.
	//
	// Use only alphanumeric characters.
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
}

// A CloudFormation `AWS::ApiGateway::DocumentationPart`.
//
// The `AWS::ApiGateway::DocumentationPart` resource creates a documentation part for an API. For more information, see [Representation of API Documentation in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-documenting-api-content-representation.html) in the *API Gateway Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDocumentationPart := awscdk.Aws_apigateway.NewCfnDocumentationPart(this, jsii.String("MyCfnDocumentationPart"), &cfnDocumentationPartProps{
//   	location: &locationProperty{
//   		method: jsii.String("method"),
//   		name: jsii.String("name"),
//   		path: jsii.String("path"),
//   		statusCode: jsii.String("statusCode"),
//   		type: jsii.String("type"),
//   	},
//   	properties: jsii.String("properties"),
//   	restApiId: jsii.String("restApiId"),
//   })
//
type CfnDocumentationPart interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrDocumentationPartId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The location of the API entity that the documentation applies to.
	Location() interface{}
	SetLocation(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// The documentation content map of the targeted API entity.
	Properties() *string
	SetProperties(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The identifier of the targeted API entity.
	RestApiId() *string
	SetRestApiId(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDocumentationPart
type jsiiProxy_CfnDocumentationPart struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDocumentationPart) AttrDocumentationPartId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDocumentationPartId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocumentationPart) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocumentationPart) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocumentationPart) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocumentationPart) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocumentationPart) Location() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocumentationPart) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocumentationPart) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocumentationPart) Properties() *string {
	var returns *string
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocumentationPart) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocumentationPart) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocumentationPart) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocumentationPart) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocumentationPart) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGateway::DocumentationPart`.
func NewCfnDocumentationPart(scope constructs.Construct, id *string, props *CfnDocumentationPartProps) CfnDocumentationPart {
	_init_.Initialize()

	j := jsiiProxy_CfnDocumentationPart{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnDocumentationPart",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGateway::DocumentationPart`.
func NewCfnDocumentationPart_Override(c CfnDocumentationPart, scope constructs.Construct, id *string, props *CfnDocumentationPartProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnDocumentationPart",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDocumentationPart) SetLocation(val interface{}) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_CfnDocumentationPart) SetProperties(val *string) {
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_CfnDocumentationPart) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDocumentationPart_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnDocumentationPart",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDocumentationPart_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnDocumentationPart",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnDocumentationPart_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnDocumentationPart",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDocumentationPart_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.CfnDocumentationPart",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDocumentationPart) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDocumentationPart) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDocumentationPart) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDocumentationPart) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDocumentationPart) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDocumentationPart) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDocumentationPart) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDocumentationPart) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDocumentationPart) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDocumentationPart) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDocumentationPart) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDocumentationPart) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDocumentationPart) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDocumentationPart) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDocumentationPart) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The `Location` property specifies the location of the Amazon API Gateway API entity that the documentation applies to.
//
// `Location` is a property of the [AWS::ApiGateway::DocumentationPart](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-documentationpart.html) resource.
//
// > For more information about each property, including constraints and valid values, see [DocumentationPart](https://docs.aws.amazon.com/apigateway/api-reference/resource/documentation-part/#location) in the *Amazon API Gateway REST API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationProperty := &locationProperty{
//   	method: jsii.String("method"),
//   	name: jsii.String("name"),
//   	path: jsii.String("path"),
//   	statusCode: jsii.String("statusCode"),
//   	type: jsii.String("type"),
//   }
//
type CfnDocumentationPart_LocationProperty struct {
	// The HTTP verb of a method.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// The name of the targeted API entity.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The URL path of the target.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The HTTP status code of a response.
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
	// The type of API entity that the documentation content applies to.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

// Properties for defining a `CfnDocumentationPart`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDocumentationPartProps := &cfnDocumentationPartProps{
//   	location: &locationProperty{
//   		method: jsii.String("method"),
//   		name: jsii.String("name"),
//   		path: jsii.String("path"),
//   		statusCode: jsii.String("statusCode"),
//   		type: jsii.String("type"),
//   	},
//   	properties: jsii.String("properties"),
//   	restApiId: jsii.String("restApiId"),
//   }
//
type CfnDocumentationPartProps struct {
	// The location of the API entity that the documentation applies to.
	Location interface{} `field:"required" json:"location" yaml:"location"`
	// The documentation content map of the targeted API entity.
	Properties *string `field:"required" json:"properties" yaml:"properties"`
	// The identifier of the targeted API entity.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
}

// A CloudFormation `AWS::ApiGateway::DocumentationVersion`.
//
// The `AWS::ApiGateway::DocumentationVersion` resource creates a snapshot of the documentation for an API. For more information, see [Representation of API Documentation in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-documenting-api-content-representation.html) in the *API Gateway Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDocumentationVersion := awscdk.Aws_apigateway.NewCfnDocumentationVersion(this, jsii.String("MyCfnDocumentationVersion"), &cfnDocumentationVersionProps{
//   	documentationVersion: jsii.String("documentationVersion"),
//   	restApiId: jsii.String("restApiId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   })
//
type CfnDocumentationVersion interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of the API documentation snapshot.
	Description() *string
	SetDescription(val *string)
	// The version identifier of the API documentation snapshot.
	DocumentationVersion() *string
	SetDocumentationVersion(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The identifier of the API.
	RestApiId() *string
	SetRestApiId(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDocumentationVersion
type jsiiProxy_CfnDocumentationVersion struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDocumentationVersion) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocumentationVersion) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocumentationVersion) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocumentationVersion) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocumentationVersion) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocumentationVersion) DocumentationVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentationVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocumentationVersion) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocumentationVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocumentationVersion) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocumentationVersion) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocumentationVersion) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocumentationVersion) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDocumentationVersion) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGateway::DocumentationVersion`.
func NewCfnDocumentationVersion(scope constructs.Construct, id *string, props *CfnDocumentationVersionProps) CfnDocumentationVersion {
	_init_.Initialize()

	j := jsiiProxy_CfnDocumentationVersion{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnDocumentationVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGateway::DocumentationVersion`.
func NewCfnDocumentationVersion_Override(c CfnDocumentationVersion, scope constructs.Construct, id *string, props *CfnDocumentationVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnDocumentationVersion",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDocumentationVersion) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnDocumentationVersion) SetDocumentationVersion(val *string) {
	_jsii_.Set(
		j,
		"documentationVersion",
		val,
	)
}

func (j *jsiiProxy_CfnDocumentationVersion) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDocumentationVersion_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnDocumentationVersion",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDocumentationVersion_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnDocumentationVersion",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnDocumentationVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnDocumentationVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDocumentationVersion_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.CfnDocumentationVersion",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDocumentationVersion) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDocumentationVersion) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDocumentationVersion) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDocumentationVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDocumentationVersion) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDocumentationVersion) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDocumentationVersion) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDocumentationVersion) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDocumentationVersion) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDocumentationVersion) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDocumentationVersion) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDocumentationVersion) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDocumentationVersion) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDocumentationVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDocumentationVersion) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnDocumentationVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDocumentationVersionProps := &cfnDocumentationVersionProps{
//   	documentationVersion: jsii.String("documentationVersion"),
//   	restApiId: jsii.String("restApiId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   }
//
type CfnDocumentationVersionProps struct {
	// The version identifier of the API documentation snapshot.
	DocumentationVersion *string `field:"required" json:"documentationVersion" yaml:"documentationVersion"`
	// The identifier of the API.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// The description of the API documentation snapshot.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

// A CloudFormation `AWS::ApiGateway::DomainName`.
//
// The `AWS::ApiGateway::DomainName` resource specifies a custom domain name for your API in API Gateway.
//
// You can use a custom domain name to provide a URL that's more intuitive and easier to recall. For more information about using custom domain names, see [Set up Custom Domain Name for an API in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html) in the *API Gateway Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDomainName := awscdk.Aws_apigateway.NewCfnDomainName(this, jsii.String("MyCfnDomainName"), &cfnDomainNameProps{
//   	certificateArn: jsii.String("certificateArn"),
//   	domainName: jsii.String("domainName"),
//   	endpointConfiguration: &endpointConfigurationProperty{
//   		types: []*string{
//   			jsii.String("types"),
//   		},
//   	},
//   	mutualTlsAuthentication: &mutualTlsAuthenticationProperty{
//   		truststoreUri: jsii.String("truststoreUri"),
//   		truststoreVersion: jsii.String("truststoreVersion"),
//   	},
//   	ownershipVerificationCertificateArn: jsii.String("ownershipVerificationCertificateArn"),
//   	regionalCertificateArn: jsii.String("regionalCertificateArn"),
//   	securityPolicy: jsii.String("securityPolicy"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnDomainName interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon CloudFront distribution domain name that's mapped to the custom domain name.
	//
	// This is only applicable for endpoints whose type is `EDGE` .
	//
	// Example: `d111111abcdef8.cloudfront.net`
	AttrDistributionDomainName() *string
	// The region-agnostic Amazon Route 53 Hosted Zone ID of the edge-optimized endpoint.
	//
	// The only valid value is `Z2FDTNDATAQYW2` for all regions.
	AttrDistributionHostedZoneId() *string
	// The domain name associated with the regional endpoint for this custom domain name.
	//
	// You set up this association by adding a DNS record that points the custom domain name to this regional domain name.
	AttrRegionalDomainName() *string
	// The region-specific Amazon Route 53 Hosted Zone ID of the regional endpoint.
	AttrRegionalHostedZoneId() *string
	// The reference to an AWS -managed certificate for use by the edge-optimized endpoint for this domain name.
	//
	// AWS Certificate Manager is the only supported source. For requirements and additional information about setting up certificates, see [Get Certificates Ready in AWS Certificate Manager](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html#how-to-custom-domains-prerequisites) in the *API Gateway Developer Guide* .
	CertificateArn() *string
	SetCertificateArn(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The custom domain name for your API.
	//
	// Uppercase letters are not supported.
	DomainName() *string
	SetDomainName(val *string)
	// A list of the endpoint types of the domain name.
	EndpointConfiguration() interface{}
	SetEndpointConfiguration(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The mutual TLS authentication configuration for a custom domain name.
	MutualTlsAuthentication() interface{}
	SetMutualTlsAuthentication(val interface{})
	// The tree node.
	Node() constructs.Node
	// The ARN of the public certificate issued by ACM to validate ownership of your custom domain.
	//
	// Only required when configuring mutual TLS and using an ACM imported or private CA certificate ARN as the RegionalCertificateArn.
	OwnershipVerificationCertificateArn() *string
	SetOwnershipVerificationCertificateArn(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The reference to an AWS -managed certificate for use by the regional endpoint for the domain name.
	//
	// AWS Certificate Manager is the only supported source.
	RegionalCertificateArn() *string
	SetRegionalCertificateArn(val *string)
	// The Transport Layer Security (TLS) version + cipher suite for this domain name.
	//
	// Valid values include `TLS_1_0` and `TLS_1_2` .
	SecurityPolicy() *string
	SetSecurityPolicy(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of arbitrary tags (key-value pairs) to associate with the domain name.
	Tags() awscdk.TagManager
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDomainName
type jsiiProxy_CfnDomainName struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDomainName) AttrDistributionDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDistributionDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) AttrDistributionHostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDistributionHostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) AttrRegionalDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrRegionalDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) AttrRegionalHostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrRegionalHostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) EndpointConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) MutualTlsAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mutualTlsAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) OwnershipVerificationCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownershipVerificationCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) RegionalCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) SecurityPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGateway::DomainName`.
func NewCfnDomainName(scope constructs.Construct, id *string, props *CfnDomainNameProps) CfnDomainName {
	_init_.Initialize()

	j := jsiiProxy_CfnDomainName{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnDomainName",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGateway::DomainName`.
func NewCfnDomainName_Override(c CfnDomainName, scope constructs.Construct, id *string, props *CfnDomainNameProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnDomainName",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDomainName) SetCertificateArn(val *string) {
	_jsii_.Set(
		j,
		"certificateArn",
		val,
	)
}

func (j *jsiiProxy_CfnDomainName) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_CfnDomainName) SetEndpointConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"endpointConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDomainName) SetMutualTlsAuthentication(val interface{}) {
	_jsii_.Set(
		j,
		"mutualTlsAuthentication",
		val,
	)
}

func (j *jsiiProxy_CfnDomainName) SetOwnershipVerificationCertificateArn(val *string) {
	_jsii_.Set(
		j,
		"ownershipVerificationCertificateArn",
		val,
	)
}

func (j *jsiiProxy_CfnDomainName) SetRegionalCertificateArn(val *string) {
	_jsii_.Set(
		j,
		"regionalCertificateArn",
		val,
	)
}

func (j *jsiiProxy_CfnDomainName) SetSecurityPolicy(val *string) {
	_jsii_.Set(
		j,
		"securityPolicy",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDomainName_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnDomainName",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDomainName_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnDomainName",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnDomainName_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnDomainName",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDomainName_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.CfnDomainName",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDomainName) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDomainName) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDomainName) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDomainName) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDomainName) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDomainName) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDomainName) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDomainName) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomainName) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomainName) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDomainName) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDomainName) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomainName) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomainName) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomainName) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The `EndpointConfiguration` property type specifies the endpoint types of an Amazon API Gateway domain name.
//
// `EndpointConfiguration` is a property of the [AWS::ApiGateway::DomainName](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointConfigurationProperty := &endpointConfigurationProperty{
//   	types: []*string{
//   		jsii.String("types"),
//   	},
//   }
//
type CfnDomainName_EndpointConfigurationProperty struct {
	// A list of endpoint types of an API or its custom domain name.
	//
	// For an edge-optimized API and its custom domain name, the endpoint type is `EDGE` . For a regional API and its custom domain name, the endpoint type is `REGIONAL` .
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

// If specified, API Gateway performs two-way authentication between the client and the server.
//
// Clients must present a trusted certificate to access your API.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mutualTlsAuthenticationProperty := &mutualTlsAuthenticationProperty{
//   	truststoreUri: jsii.String("truststoreUri"),
//   	truststoreVersion: jsii.String("truststoreVersion"),
//   }
//
type CfnDomainName_MutualTlsAuthenticationProperty struct {
	// An Amazon S3 URL that specifies the truststore for mutual TLS authentication, for example, `s3:// bucket-name / key-name` .
	//
	// The truststore can contain certificates from public or private certificate authorities. To update the truststore, upload a new version to S3, and then update your custom domain name to use the new version. To update the truststore, you must have permissions to access the S3 object.
	TruststoreUri *string `field:"optional" json:"truststoreUri" yaml:"truststoreUri"`
	// The version of the S3 object that contains your truststore.
	//
	// To specify a version, you must have versioning enabled for the S3 bucket.
	TruststoreVersion *string `field:"optional" json:"truststoreVersion" yaml:"truststoreVersion"`
}

// Properties for defining a `CfnDomainName`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDomainNameProps := &cfnDomainNameProps{
//   	certificateArn: jsii.String("certificateArn"),
//   	domainName: jsii.String("domainName"),
//   	endpointConfiguration: &endpointConfigurationProperty{
//   		types: []*string{
//   			jsii.String("types"),
//   		},
//   	},
//   	mutualTlsAuthentication: &mutualTlsAuthenticationProperty{
//   		truststoreUri: jsii.String("truststoreUri"),
//   		truststoreVersion: jsii.String("truststoreVersion"),
//   	},
//   	ownershipVerificationCertificateArn: jsii.String("ownershipVerificationCertificateArn"),
//   	regionalCertificateArn: jsii.String("regionalCertificateArn"),
//   	securityPolicy: jsii.String("securityPolicy"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDomainNameProps struct {
	// The reference to an AWS -managed certificate for use by the edge-optimized endpoint for this domain name.
	//
	// AWS Certificate Manager is the only supported source. For requirements and additional information about setting up certificates, see [Get Certificates Ready in AWS Certificate Manager](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html#how-to-custom-domains-prerequisites) in the *API Gateway Developer Guide* .
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// The custom domain name for your API.
	//
	// Uppercase letters are not supported.
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// A list of the endpoint types of the domain name.
	EndpointConfiguration interface{} `field:"optional" json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// The mutual TLS authentication configuration for a custom domain name.
	MutualTlsAuthentication interface{} `field:"optional" json:"mutualTlsAuthentication" yaml:"mutualTlsAuthentication"`
	// The ARN of the public certificate issued by ACM to validate ownership of your custom domain.
	//
	// Only required when configuring mutual TLS and using an ACM imported or private CA certificate ARN as the RegionalCertificateArn.
	OwnershipVerificationCertificateArn *string `field:"optional" json:"ownershipVerificationCertificateArn" yaml:"ownershipVerificationCertificateArn"`
	// The reference to an AWS -managed certificate for use by the regional endpoint for the domain name.
	//
	// AWS Certificate Manager is the only supported source.
	RegionalCertificateArn *string `field:"optional" json:"regionalCertificateArn" yaml:"regionalCertificateArn"`
	// The Transport Layer Security (TLS) version + cipher suite for this domain name.
	//
	// Valid values include `TLS_1_0` and `TLS_1_2` .
	SecurityPolicy *string `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
	// An array of arbitrary tags (key-value pairs) to associate with the domain name.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::ApiGateway::GatewayResponse`.
//
// The `AWS::ApiGateway::GatewayResponse` resource creates a gateway response for your API. For more information, see [API Gateway Responses](https://docs.aws.amazon.com/apigateway/latest/developerguide/customize-gateway-responses.html#api-gateway-gatewayResponse-definition) in the *API Gateway Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGatewayResponse := awscdk.Aws_apigateway.NewCfnGatewayResponse(this, jsii.String("MyCfnGatewayResponse"), &cfnGatewayResponseProps{
//   	responseType: jsii.String("responseType"),
//   	restApiId: jsii.String("restApiId"),
//
//   	// the properties below are optional
//   	responseParameters: map[string]*string{
//   		"responseParametersKey": jsii.String("responseParameters"),
//   	},
//   	responseTemplates: map[string]*string{
//   		"responseTemplatesKey": jsii.String("responseTemplates"),
//   	},
//   	statusCode: jsii.String("statusCode"),
//   })
//
type CfnGatewayResponse interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ID for the gateway response.
	//
	// For example: `abc123` .
	AttrId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The response parameters (paths, query strings, and headers) for the response.
	//
	// Duplicates not allowed.
	ResponseParameters() interface{}
	SetResponseParameters(val interface{})
	// The response templates for the response.
	//
	// Duplicates not allowed.
	ResponseTemplates() interface{}
	SetResponseTemplates(val interface{})
	// The response type.
	//
	// For valid values, see [GatewayResponse](https://docs.aws.amazon.com/apigateway/api-reference/resource/gateway-response/) in the *API Gateway API Reference* .
	ResponseType() *string
	SetResponseType(val *string)
	// The identifier of the API.
	RestApiId() *string
	SetRestApiId(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The HTTP status code for the response.
	StatusCode() *string
	SetStatusCode(val *string)
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnGatewayResponse
type jsiiProxy_CfnGatewayResponse struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnGatewayResponse) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayResponse) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayResponse) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayResponse) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayResponse) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayResponse) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayResponse) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayResponse) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayResponse) ResponseParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responseParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayResponse) ResponseTemplates() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responseTemplates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayResponse) ResponseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayResponse) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayResponse) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayResponse) StatusCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayResponse) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayResponse) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGateway::GatewayResponse`.
func NewCfnGatewayResponse(scope constructs.Construct, id *string, props *CfnGatewayResponseProps) CfnGatewayResponse {
	_init_.Initialize()

	j := jsiiProxy_CfnGatewayResponse{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnGatewayResponse",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGateway::GatewayResponse`.
func NewCfnGatewayResponse_Override(c CfnGatewayResponse, scope constructs.Construct, id *string, props *CfnGatewayResponseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnGatewayResponse",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnGatewayResponse) SetResponseParameters(val interface{}) {
	_jsii_.Set(
		j,
		"responseParameters",
		val,
	)
}

func (j *jsiiProxy_CfnGatewayResponse) SetResponseTemplates(val interface{}) {
	_jsii_.Set(
		j,
		"responseTemplates",
		val,
	)
}

func (j *jsiiProxy_CfnGatewayResponse) SetResponseType(val *string) {
	_jsii_.Set(
		j,
		"responseType",
		val,
	)
}

func (j *jsiiProxy_CfnGatewayResponse) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

func (j *jsiiProxy_CfnGatewayResponse) SetStatusCode(val *string) {
	_jsii_.Set(
		j,
		"statusCode",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnGatewayResponse_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnGatewayResponse",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnGatewayResponse_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnGatewayResponse",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnGatewayResponse_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnGatewayResponse",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGatewayResponse_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.CfnGatewayResponse",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGatewayResponse) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnGatewayResponse) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnGatewayResponse) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnGatewayResponse) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnGatewayResponse) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnGatewayResponse) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnGatewayResponse) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnGatewayResponse) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGatewayResponse) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGatewayResponse) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnGatewayResponse) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnGatewayResponse) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGatewayResponse) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGatewayResponse) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGatewayResponse) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnGatewayResponse`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGatewayResponseProps := &cfnGatewayResponseProps{
//   	responseType: jsii.String("responseType"),
//   	restApiId: jsii.String("restApiId"),
//
//   	// the properties below are optional
//   	responseParameters: map[string]*string{
//   		"responseParametersKey": jsii.String("responseParameters"),
//   	},
//   	responseTemplates: map[string]*string{
//   		"responseTemplatesKey": jsii.String("responseTemplates"),
//   	},
//   	statusCode: jsii.String("statusCode"),
//   }
//
type CfnGatewayResponseProps struct {
	// The response type.
	//
	// For valid values, see [GatewayResponse](https://docs.aws.amazon.com/apigateway/api-reference/resource/gateway-response/) in the *API Gateway API Reference* .
	ResponseType *string `field:"required" json:"responseType" yaml:"responseType"`
	// The identifier of the API.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// The response parameters (paths, query strings, and headers) for the response.
	//
	// Duplicates not allowed.
	ResponseParameters interface{} `field:"optional" json:"responseParameters" yaml:"responseParameters"`
	// The response templates for the response.
	//
	// Duplicates not allowed.
	ResponseTemplates interface{} `field:"optional" json:"responseTemplates" yaml:"responseTemplates"`
	// The HTTP status code for the response.
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
}

// A CloudFormation `AWS::ApiGateway::Method`.
//
// The `AWS::ApiGateway::Method` resource creates API Gateway methods that define the parameters and body that clients must send in their requests.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMethod := awscdk.Aws_apigateway.NewCfnMethod(this, jsii.String("MyCfnMethod"), &cfnMethodProps{
//   	httpMethod: jsii.String("httpMethod"),
//   	resourceId: jsii.String("resourceId"),
//   	restApiId: jsii.String("restApiId"),
//
//   	// the properties below are optional
//   	apiKeyRequired: jsii.Boolean(false),
//   	authorizationScopes: []*string{
//   		jsii.String("authorizationScopes"),
//   	},
//   	authorizationType: jsii.String("authorizationType"),
//   	authorizerId: jsii.String("authorizerId"),
//   	integration: &integrationProperty{
//   		cacheKeyParameters: []*string{
//   			jsii.String("cacheKeyParameters"),
//   		},
//   		cacheNamespace: jsii.String("cacheNamespace"),
//   		connectionId: jsii.String("connectionId"),
//   		connectionType: jsii.String("connectionType"),
//   		contentHandling: jsii.String("contentHandling"),
//   		credentials: jsii.String("credentials"),
//   		integrationHttpMethod: jsii.String("integrationHttpMethod"),
//   		integrationResponses: []interface{}{
//   			&integrationResponseProperty{
//   				statusCode: jsii.String("statusCode"),
//
//   				// the properties below are optional
//   				contentHandling: jsii.String("contentHandling"),
//   				responseParameters: map[string]*string{
//   					"responseParametersKey": jsii.String("responseParameters"),
//   				},
//   				responseTemplates: map[string]*string{
//   					"responseTemplatesKey": jsii.String("responseTemplates"),
//   				},
//   				selectionPattern: jsii.String("selectionPattern"),
//   			},
//   		},
//   		passthroughBehavior: jsii.String("passthroughBehavior"),
//   		requestParameters: map[string]*string{
//   			"requestParametersKey": jsii.String("requestParameters"),
//   		},
//   		requestTemplates: map[string]*string{
//   			"requestTemplatesKey": jsii.String("requestTemplates"),
//   		},
//   		timeoutInMillis: jsii.Number(123),
//   		type: jsii.String("type"),
//   		uri: jsii.String("uri"),
//   	},
//   	methodResponses: []interface{}{
//   		&methodResponseProperty{
//   			statusCode: jsii.String("statusCode"),
//
//   			// the properties below are optional
//   			responseModels: map[string]*string{
//   				"responseModelsKey": jsii.String("responseModels"),
//   			},
//   			responseParameters: map[string]interface{}{
//   				"responseParametersKey": jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	operationName: jsii.String("operationName"),
//   	requestModels: map[string]*string{
//   		"requestModelsKey": jsii.String("requestModels"),
//   	},
//   	requestParameters: map[string]interface{}{
//   		"requestParametersKey": jsii.Boolean(false),
//   	},
//   	requestValidatorId: jsii.String("requestValidatorId"),
//   })
//
type CfnMethod interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Indicates whether the method requires clients to submit a valid API key.
	ApiKeyRequired() interface{}
	SetApiKeyRequired(val interface{})
	// A list of authorization scopes configured on the method.
	//
	// The scopes are used with a `COGNITO_USER_POOLS` authorizer to authorize the method invocation. The authorization works by matching the method scopes against the scopes parsed from the access token in the incoming request. The method invocation is authorized if any method scopes match a claimed scope in the access token. Otherwise, the invocation is not authorized. When the method scope is configured, the client must provide an access token instead of an identity token for authorization purposes.
	AuthorizationScopes() *[]*string
	SetAuthorizationScopes(val *[]*string)
	// The method's authorization type.
	//
	// This parameter is required. For valid values, see [Method](https://docs.aws.amazon.com/apigateway/api-reference/resource/method/) in the *API Gateway API Reference* .
	//
	// > If you specify the `AuthorizerId` property, specify `CUSTOM` or `COGNITO_USER_POOLS` for this property.
	AuthorizationType() *string
	SetAuthorizationType(val *string)
	// The identifier of the [authorizer](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html) to use on this method. If you specify this property, specify `CUSTOM` or `COGNITO_USER_POOLS` for the `AuthorizationType` property.
	AuthorizerId() *string
	SetAuthorizerId(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The HTTP method that clients use to call this method.
	HttpMethod() *string
	SetHttpMethod(val *string)
	// The backend system that the method calls when it receives a request.
	Integration() interface{}
	SetIntegration(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The responses that can be sent to the client who calls the method.
	MethodResponses() interface{}
	SetMethodResponses(val interface{})
	// The tree node.
	Node() constructs.Node
	// A friendly operation name for the method.
	//
	// For example, you can assign the `OperationName` of `ListPets` for the `GET /pets` method.
	OperationName() *string
	SetOperationName(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The resources that are used for the request's content type.
	//
	// Specify request models as key-value pairs (string-to-string mapping), with a content type as the key and a `Model` resource name as the value. To use the same model regardless of the content type, specify `$default` as the key.
	RequestModels() interface{}
	SetRequestModels(val interface{})
	// The request parameters that API Gateway accepts.
	//
	// Specify request parameters as key-value pairs (string-to-Boolean mapping), with a source as the key and a Boolean as the value. The Boolean specifies whether a parameter is required. A source must match the format `method.request. *location* . *name*` , where the location is querystring, path, or header, and *name* is a valid, unique parameter name.
	RequestParameters() interface{}
	SetRequestParameters(val interface{})
	// The ID of the associated request validator.
	RequestValidatorId() *string
	SetRequestValidatorId(val *string)
	// The ID of an API Gateway [resource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html) . For root resource methods, specify the `RestApi` root resource ID, such as `{ "Fn::GetAtt": ["MyRestApi", "RootResourceId"] }` .
	ResourceId() *string
	SetResourceId(val *string)
	// The ID of the [RestApi](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html) resource in which API Gateway creates the method.
	RestApiId() *string
	SetRestApiId(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnMethod
type jsiiProxy_CfnMethod struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnMethod) ApiKeyRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiKeyRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMethod) AuthorizationScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorizationScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMethod) AuthorizationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMethod) AuthorizerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMethod) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMethod) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMethod) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMethod) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMethod) HttpMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMethod) Integration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"integration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMethod) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMethod) MethodResponses() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"methodResponses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMethod) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMethod) OperationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMethod) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMethod) RequestModels() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestModels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMethod) RequestParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMethod) RequestValidatorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestValidatorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMethod) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMethod) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMethod) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMethod) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMethod) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGateway::Method`.
func NewCfnMethod(scope constructs.Construct, id *string, props *CfnMethodProps) CfnMethod {
	_init_.Initialize()

	j := jsiiProxy_CfnMethod{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnMethod",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGateway::Method`.
func NewCfnMethod_Override(c CfnMethod, scope constructs.Construct, id *string, props *CfnMethodProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnMethod",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnMethod) SetApiKeyRequired(val interface{}) {
	_jsii_.Set(
		j,
		"apiKeyRequired",
		val,
	)
}

func (j *jsiiProxy_CfnMethod) SetAuthorizationScopes(val *[]*string) {
	_jsii_.Set(
		j,
		"authorizationScopes",
		val,
	)
}

func (j *jsiiProxy_CfnMethod) SetAuthorizationType(val *string) {
	_jsii_.Set(
		j,
		"authorizationType",
		val,
	)
}

func (j *jsiiProxy_CfnMethod) SetAuthorizerId(val *string) {
	_jsii_.Set(
		j,
		"authorizerId",
		val,
	)
}

func (j *jsiiProxy_CfnMethod) SetHttpMethod(val *string) {
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_CfnMethod) SetIntegration(val interface{}) {
	_jsii_.Set(
		j,
		"integration",
		val,
	)
}

func (j *jsiiProxy_CfnMethod) SetMethodResponses(val interface{}) {
	_jsii_.Set(
		j,
		"methodResponses",
		val,
	)
}

func (j *jsiiProxy_CfnMethod) SetOperationName(val *string) {
	_jsii_.Set(
		j,
		"operationName",
		val,
	)
}

func (j *jsiiProxy_CfnMethod) SetRequestModels(val interface{}) {
	_jsii_.Set(
		j,
		"requestModels",
		val,
	)
}

func (j *jsiiProxy_CfnMethod) SetRequestParameters(val interface{}) {
	_jsii_.Set(
		j,
		"requestParameters",
		val,
	)
}

func (j *jsiiProxy_CfnMethod) SetRequestValidatorId(val *string) {
	_jsii_.Set(
		j,
		"requestValidatorId",
		val,
	)
}

func (j *jsiiProxy_CfnMethod) SetResourceId(val *string) {
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_CfnMethod) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnMethod_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnMethod",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnMethod_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnMethod",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnMethod_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnMethod",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMethod_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.CfnMethod",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMethod) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnMethod) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnMethod) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnMethod) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnMethod) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnMethod) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnMethod) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnMethod) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMethod) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMethod) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnMethod) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnMethod) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMethod) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMethod) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMethod) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// `Integration` is a property of the [AWS::ApiGateway::Method](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html) resource that specifies information about the target backend that a method calls.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integrationProperty := &integrationProperty{
//   	cacheKeyParameters: []*string{
//   		jsii.String("cacheKeyParameters"),
//   	},
//   	cacheNamespace: jsii.String("cacheNamespace"),
//   	connectionId: jsii.String("connectionId"),
//   	connectionType: jsii.String("connectionType"),
//   	contentHandling: jsii.String("contentHandling"),
//   	credentials: jsii.String("credentials"),
//   	integrationHttpMethod: jsii.String("integrationHttpMethod"),
//   	integrationResponses: []interface{}{
//   		&integrationResponseProperty{
//   			statusCode: jsii.String("statusCode"),
//
//   			// the properties below are optional
//   			contentHandling: jsii.String("contentHandling"),
//   			responseParameters: map[string]*string{
//   				"responseParametersKey": jsii.String("responseParameters"),
//   			},
//   			responseTemplates: map[string]*string{
//   				"responseTemplatesKey": jsii.String("responseTemplates"),
//   			},
//   			selectionPattern: jsii.String("selectionPattern"),
//   		},
//   	},
//   	passthroughBehavior: jsii.String("passthroughBehavior"),
//   	requestParameters: map[string]*string{
//   		"requestParametersKey": jsii.String("requestParameters"),
//   	},
//   	requestTemplates: map[string]*string{
//   		"requestTemplatesKey": jsii.String("requestTemplates"),
//   	},
//   	timeoutInMillis: jsii.Number(123),
//   	type: jsii.String("type"),
//   	uri: jsii.String("uri"),
//   }
//
type CfnMethod_IntegrationProperty struct {
	// A list of request parameters whose values API Gateway caches.
	//
	// For cases where the integration type allows for RequestParameters to be set, these parameters must also be specified in [RequestParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-requestparameters) to be supported in `CacheKeyParameters` .
	CacheKeyParameters *[]*string `field:"optional" json:"cacheKeyParameters" yaml:"cacheKeyParameters"`
	// An API-specific tag group of related cached parameters.
	CacheNamespace *string `field:"optional" json:"cacheNamespace" yaml:"cacheNamespace"`
	// The ID of the `VpcLink` used for the integration when `connectionType=VPC_LINK` , otherwise undefined.
	ConnectionId *string `field:"optional" json:"connectionId" yaml:"connectionId"`
	// The type of the network connection to the integration endpoint.
	//
	// The valid value is `INTERNET` for connections through the public routable internet or `VPC_LINK` for private connections between API Gateway and a network load balancer in a VPC. The default value is `INTERNET` .
	ConnectionType *string `field:"optional" json:"connectionType" yaml:"connectionType"`
	// Specifies how to handle request payload content type conversions. Valid values are:.
	//
	// - `CONVERT_TO_BINARY` : Converts a request payload from a base64-encoded string to a binary blob.
	// - `CONVERT_TO_TEXT` : Converts a request payload from a binary blob to a base64-encoded string.
	//
	// If this property isn't defined, the request payload is passed through from the method request to the integration request without modification, provided that the `PassthroughBehaviors` property is configured to support payload pass-through.
	ContentHandling *string `field:"optional" json:"contentHandling" yaml:"contentHandling"`
	// The credentials that are required for the integration.
	//
	// To specify an AWS Identity and Access Management (IAM) role that API Gateway assumes, specify the role's Amazon Resource Name (ARN). To require that the caller's identity be passed through from the request, specify arn:aws:iam::*:user/*.
	//
	// To use resource-based permissions on the AWS Lambda (Lambda) function, don't specify this property. Use the [AWS::Lambda::Permission](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html) resource to permit API Gateway to call the function. For more information, see [Allow Amazon API Gateway to Invoke a Lambda Function](https://docs.aws.amazon.com/lambda/latest/dg/access-control-resource-based.html#access-control-resource-based-example-apigateway-invoke-function) in the *AWS Lambda Developer Guide* .
	Credentials *string `field:"optional" json:"credentials" yaml:"credentials"`
	// The integration's HTTP method type.
	//
	// For the `Type` property, if you specify `MOCK` , this property is optional. For all other types, you must specify this property.
	IntegrationHttpMethod *string `field:"optional" json:"integrationHttpMethod" yaml:"integrationHttpMethod"`
	// The response that API Gateway provides after a method's backend completes processing a request.
	//
	// API Gateway intercepts the response from the backend so that you can control how API Gateway surfaces backend responses. For example, you can map the backend status codes to codes that you define.
	IntegrationResponses interface{} `field:"optional" json:"integrationResponses" yaml:"integrationResponses"`
	// Indicates when API Gateway passes requests to the targeted backend.
	//
	// This behavior depends on the request's `Content-Type` header and whether you defined a mapping template for it.
	//
	// For more information and valid values, see the [passthroughBehavior](https://docs.aws.amazon.com/apigateway/api-reference/link-relation/integration-put/#passthroughBehavior) field in the *API Gateway API Reference* .
	PassthroughBehavior *string `field:"optional" json:"passthroughBehavior" yaml:"passthroughBehavior"`
	// The request parameters that API Gateway sends with the backend request.
	//
	// Specify request parameters as key-value pairs (string-to-string mappings), with a destination as the key and a source as the value.
	//
	// Specify the destination by using the following pattern `integration.request. *location* . *name*` , where *location* is query string, path, or header, and *name* is a valid, unique parameter name.
	//
	// The source must be an existing method request parameter or a static value. You must enclose static values in single quotation marks and pre-encode these values based on their destination in the request.
	RequestParameters interface{} `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// A map of Apache Velocity templates that are applied on the request payload.
	//
	// The template that API Gateway uses is based on the value of the `Content-Type` header that's sent by the client. The content type value is the key, and the template is the value (specified as a string), such as the following snippet:
	//
	// `"application/json": "{\n \"statusCode\": 200\n}"`
	//
	// For more information about templates, see [API Gateway Mapping Template and Access Logging Variable Reference](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html) in the *API Gateway Developer Guide* .
	RequestTemplates interface{} `field:"optional" json:"requestTemplates" yaml:"requestTemplates"`
	// Custom timeout between 50 and 29,000 milliseconds.
	//
	// The default value is 29,000 milliseconds or 29 seconds.
	TimeoutInMillis *float64 `field:"optional" json:"timeoutInMillis" yaml:"timeoutInMillis"`
	// The type of backend that your method is running, such as `HTTP` or `MOCK` .
	//
	// For all of the valid values, see the [type](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/#type) property for the `Integration` resource in the *Amazon API Gateway REST API Reference* .
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The Uniform Resource Identifier (URI) for the integration.
	//
	// If you specify `HTTP` for the `Type` property, specify the API endpoint URL.
	//
	// If you specify `MOCK` for the `Type` property, don't specify this property.
	//
	// If you specify `AWS` for the `Type` property, specify an AWS service that follows this form: arn:aws:apigateway: *region* : *subdomain* . *service|service* : *path|action* / *service_api* . For example, a Lambda function URI follows this form: arn:aws:apigateway: *region* :lambda:path/ *path* . The path is usually in the form /2015-03-31/functions/ *LambdaFunctionARN* /invocations. For more information, see the `uri` property of the [Integration](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/) resource in the Amazon API Gateway REST API Reference.
	//
	// If you specified `HTTP` or `AWS` for the `Type` property, you must specify this property.
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

// `IntegrationResponse` is a property of the [Amazon API Gateway Method Integration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html) property type that specifies the response that API Gateway sends after a method's backend finishes processing a request.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integrationResponseProperty := &integrationResponseProperty{
//   	statusCode: jsii.String("statusCode"),
//
//   	// the properties below are optional
//   	contentHandling: jsii.String("contentHandling"),
//   	responseParameters: map[string]*string{
//   		"responseParametersKey": jsii.String("responseParameters"),
//   	},
//   	responseTemplates: map[string]*string{
//   		"responseTemplatesKey": jsii.String("responseTemplates"),
//   	},
//   	selectionPattern: jsii.String("selectionPattern"),
//   }
//
type CfnMethod_IntegrationResponseProperty struct {
	// The status code that API Gateway uses to map the integration response to a [MethodResponse](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-methodresponse.html) status code.
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
	// Specifies how to handle request payload content type conversions. Valid values are:.
	//
	// - `CONVERT_TO_BINARY` : Converts a request payload from a base64-encoded string to a binary blob.
	// - `CONVERT_TO_TEXT` : Converts a request payload from a binary blob to a base64-encoded string.
	//
	// If this property isn't defined, the request payload is passed through from the method request to the integration request without modification.
	ContentHandling *string `field:"optional" json:"contentHandling" yaml:"contentHandling"`
	// The response parameters from the backend response that API Gateway sends to the method response.
	//
	// Specify response parameters as key-value pairs ( [string-to-string mappings](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/mappings-section-structure.html) ).
	//
	// Use the destination as the key and the source as the value:
	//
	// - The destination must be an existing response parameter in the [MethodResponse](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-methodresponse.html) property.
	// - The source must be an existing method request parameter or a static value. You must enclose static values in single quotation marks and pre-encode these values based on the destination specified in the request.
	//
	// For more information about templates, see [API Gateway Mapping Template and Access Logging Variable Reference](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html) in the *API Gateway Developer Guide* .
	ResponseParameters interface{} `field:"optional" json:"responseParameters" yaml:"responseParameters"`
	// The templates that are used to transform the integration response body.
	//
	// Specify templates as key-value pairs (string-to-string mappings), with a content type as the key and a template as the value. For more information, see [API Gateway Mapping Template and Access Logging Variable Reference](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html) in the *API Gateway Developer Guide* .
	ResponseTemplates interface{} `field:"optional" json:"responseTemplates" yaml:"responseTemplates"`
	// A [regular expression](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-regexes.html) that specifies which error strings or status codes from the backend map to the integration response.
	SelectionPattern *string `field:"optional" json:"selectionPattern" yaml:"selectionPattern"`
}

// `MethodResponse` is a property of the [AWS::ApiGateway::Method](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html) resource that defines the responses that can be sent to the client that calls a method.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   methodResponseProperty := &methodResponseProperty{
//   	statusCode: jsii.String("statusCode"),
//
//   	// the properties below are optional
//   	responseModels: map[string]*string{
//   		"responseModelsKey": jsii.String("responseModels"),
//   	},
//   	responseParameters: map[string]interface{}{
//   		"responseParametersKey": jsii.Boolean(false),
//   	},
//   }
//
type CfnMethod_MethodResponseProperty struct {
	// The method response's status code, which you map to an [IntegrationResponse](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration-integrationresponse.html) .
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
	// The resources used for the response's content type.
	//
	// Specify response models as key-value pairs (string-to-string maps), with a content type as the key and a [Model](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-model.html) resource name as the value.
	ResponseModels interface{} `field:"optional" json:"responseModels" yaml:"responseModels"`
	// Response parameters that API Gateway sends to the client that called a method.
	//
	// Specify response parameters as key-value pairs (string-to-Boolean maps), with a destination as the key and a Boolean as the value. Specify the destination using the following pattern: `method.response.header. *name*` , where *name* is a valid, unique header name. The Boolean specifies whether a parameter is required.
	ResponseParameters interface{} `field:"optional" json:"responseParameters" yaml:"responseParameters"`
}

// Properties for defining a `CfnMethod`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMethodProps := &cfnMethodProps{
//   	httpMethod: jsii.String("httpMethod"),
//   	resourceId: jsii.String("resourceId"),
//   	restApiId: jsii.String("restApiId"),
//
//   	// the properties below are optional
//   	apiKeyRequired: jsii.Boolean(false),
//   	authorizationScopes: []*string{
//   		jsii.String("authorizationScopes"),
//   	},
//   	authorizationType: jsii.String("authorizationType"),
//   	authorizerId: jsii.String("authorizerId"),
//   	integration: &integrationProperty{
//   		cacheKeyParameters: []*string{
//   			jsii.String("cacheKeyParameters"),
//   		},
//   		cacheNamespace: jsii.String("cacheNamespace"),
//   		connectionId: jsii.String("connectionId"),
//   		connectionType: jsii.String("connectionType"),
//   		contentHandling: jsii.String("contentHandling"),
//   		credentials: jsii.String("credentials"),
//   		integrationHttpMethod: jsii.String("integrationHttpMethod"),
//   		integrationResponses: []interface{}{
//   			&integrationResponseProperty{
//   				statusCode: jsii.String("statusCode"),
//
//   				// the properties below are optional
//   				contentHandling: jsii.String("contentHandling"),
//   				responseParameters: map[string]*string{
//   					"responseParametersKey": jsii.String("responseParameters"),
//   				},
//   				responseTemplates: map[string]*string{
//   					"responseTemplatesKey": jsii.String("responseTemplates"),
//   				},
//   				selectionPattern: jsii.String("selectionPattern"),
//   			},
//   		},
//   		passthroughBehavior: jsii.String("passthroughBehavior"),
//   		requestParameters: map[string]*string{
//   			"requestParametersKey": jsii.String("requestParameters"),
//   		},
//   		requestTemplates: map[string]*string{
//   			"requestTemplatesKey": jsii.String("requestTemplates"),
//   		},
//   		timeoutInMillis: jsii.Number(123),
//   		type: jsii.String("type"),
//   		uri: jsii.String("uri"),
//   	},
//   	methodResponses: []interface{}{
//   		&methodResponseProperty{
//   			statusCode: jsii.String("statusCode"),
//
//   			// the properties below are optional
//   			responseModels: map[string]*string{
//   				"responseModelsKey": jsii.String("responseModels"),
//   			},
//   			responseParameters: map[string]interface{}{
//   				"responseParametersKey": jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	operationName: jsii.String("operationName"),
//   	requestModels: map[string]*string{
//   		"requestModelsKey": jsii.String("requestModels"),
//   	},
//   	requestParameters: map[string]interface{}{
//   		"requestParametersKey": jsii.Boolean(false),
//   	},
//   	requestValidatorId: jsii.String("requestValidatorId"),
//   }
//
type CfnMethodProps struct {
	// The HTTP method that clients use to call this method.
	HttpMethod *string `field:"required" json:"httpMethod" yaml:"httpMethod"`
	// The ID of an API Gateway [resource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html) . For root resource methods, specify the `RestApi` root resource ID, such as `{ "Fn::GetAtt": ["MyRestApi", "RootResourceId"] }` .
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// The ID of the [RestApi](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html) resource in which API Gateway creates the method.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// Indicates whether the method requires clients to submit a valid API key.
	ApiKeyRequired interface{} `field:"optional" json:"apiKeyRequired" yaml:"apiKeyRequired"`
	// A list of authorization scopes configured on the method.
	//
	// The scopes are used with a `COGNITO_USER_POOLS` authorizer to authorize the method invocation. The authorization works by matching the method scopes against the scopes parsed from the access token in the incoming request. The method invocation is authorized if any method scopes match a claimed scope in the access token. Otherwise, the invocation is not authorized. When the method scope is configured, the client must provide an access token instead of an identity token for authorization purposes.
	AuthorizationScopes *[]*string `field:"optional" json:"authorizationScopes" yaml:"authorizationScopes"`
	// The method's authorization type.
	//
	// This parameter is required. For valid values, see [Method](https://docs.aws.amazon.com/apigateway/api-reference/resource/method/) in the *API Gateway API Reference* .
	//
	// > If you specify the `AuthorizerId` property, specify `CUSTOM` or `COGNITO_USER_POOLS` for this property.
	AuthorizationType *string `field:"optional" json:"authorizationType" yaml:"authorizationType"`
	// The identifier of the [authorizer](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html) to use on this method. If you specify this property, specify `CUSTOM` or `COGNITO_USER_POOLS` for the `AuthorizationType` property.
	AuthorizerId *string `field:"optional" json:"authorizerId" yaml:"authorizerId"`
	// The backend system that the method calls when it receives a request.
	Integration interface{} `field:"optional" json:"integration" yaml:"integration"`
	// The responses that can be sent to the client who calls the method.
	MethodResponses interface{} `field:"optional" json:"methodResponses" yaml:"methodResponses"`
	// A friendly operation name for the method.
	//
	// For example, you can assign the `OperationName` of `ListPets` for the `GET /pets` method.
	OperationName *string `field:"optional" json:"operationName" yaml:"operationName"`
	// The resources that are used for the request's content type.
	//
	// Specify request models as key-value pairs (string-to-string mapping), with a content type as the key and a `Model` resource name as the value. To use the same model regardless of the content type, specify `$default` as the key.
	RequestModels interface{} `field:"optional" json:"requestModels" yaml:"requestModels"`
	// The request parameters that API Gateway accepts.
	//
	// Specify request parameters as key-value pairs (string-to-Boolean mapping), with a source as the key and a Boolean as the value. The Boolean specifies whether a parameter is required. A source must match the format `method.request. *location* . *name*` , where the location is querystring, path, or header, and *name* is a valid, unique parameter name.
	RequestParameters interface{} `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// The ID of the associated request validator.
	RequestValidatorId *string `field:"optional" json:"requestValidatorId" yaml:"requestValidatorId"`
}

// A CloudFormation `AWS::ApiGateway::Model`.
//
// The `AWS::ApiGateway::Model` resource defines the structure of a request or response payload for an API method.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var schema interface{}
//
//   cfnModel := awscdk.Aws_apigateway.NewCfnModel(this, jsii.String("MyCfnModel"), &cfnModelProps{
//   	restApiId: jsii.String("restApiId"),
//
//   	// the properties below are optional
//   	contentType: jsii.String("contentType"),
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	schema: schema,
//   })
//
type CfnModel interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The content type for the model.
	ContentType() *string
	SetContentType(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A description that identifies this model.
	Description() *string
	SetDescription(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// A name for the model.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the model name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The ID of a REST API with which to associate this model.
	RestApiId() *string
	SetRestApiId(val *string)
	// The schema to use to transform data to one or more output formats.
	//
	// Specify null ( `{}` ) if you don't want to specify a schema.
	Schema() interface{}
	SetSchema(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnModel
type jsiiProxy_CfnModel struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnModel) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) Schema() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGateway::Model`.
func NewCfnModel(scope constructs.Construct, id *string, props *CfnModelProps) CfnModel {
	_init_.Initialize()

	j := jsiiProxy_CfnModel{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnModel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGateway::Model`.
func NewCfnModel_Override(c CfnModel, scope constructs.Construct, id *string, props *CfnModelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnModel",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnModel) SetContentType(val *string) {
	_jsii_.Set(
		j,
		"contentType",
		val,
	)
}

func (j *jsiiProxy_CfnModel) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnModel) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnModel) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

func (j *jsiiProxy_CfnModel) SetSchema(val interface{}) {
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnModel_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnModel",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnModel_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnModel",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnModel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnModel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnModel_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.CfnModel",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnModel) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnModel) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnModel) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnModel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnModel) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnModel) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnModel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnModel) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModel) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModel) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnModel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnModel) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModel) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModel) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnModel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var schema interface{}
//
//   cfnModelProps := &cfnModelProps{
//   	restApiId: jsii.String("restApiId"),
//
//   	// the properties below are optional
//   	contentType: jsii.String("contentType"),
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	schema: schema,
//   }
//
type CfnModelProps struct {
	// The ID of a REST API with which to associate this model.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// The content type for the model.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// A description that identifies this model.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name for the model.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the model name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The schema to use to transform data to one or more output formats.
	//
	// Specify null ( `{}` ) if you don't want to specify a schema.
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
}

// A CloudFormation `AWS::ApiGateway::RequestValidator`.
//
// The `AWS::ApiGateway::RequestValidator` resource sets up basic validation rules for incoming requests to your API. For more information, see [Enable Basic Request Validation for an API in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-method-request-validation.html) in the *API Gateway Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRequestValidator := awscdk.Aws_apigateway.NewCfnRequestValidator(this, jsii.String("MyCfnRequestValidator"), &cfnRequestValidatorProps{
//   	restApiId: jsii.String("restApiId"),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	validateRequestBody: jsii.Boolean(false),
//   	validateRequestParameters: jsii.Boolean(false),
//   })
//
type CfnRequestValidator interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ID for the request validator.
	//
	// For example: `abc123` .
	AttrRequestValidatorId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The name of this request validator.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The identifier of the targeted API entity.
	RestApiId() *string
	SetRestApiId(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Indicates whether to validate the request body according to the configured schema for the targeted API and method.
	ValidateRequestBody() interface{}
	SetValidateRequestBody(val interface{})
	// Indicates whether to validate request parameters.
	ValidateRequestParameters() interface{}
	SetValidateRequestParameters(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnRequestValidator
type jsiiProxy_CfnRequestValidator struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnRequestValidator) AttrRequestValidatorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrRequestValidatorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRequestValidator) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRequestValidator) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRequestValidator) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRequestValidator) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRequestValidator) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRequestValidator) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRequestValidator) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRequestValidator) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRequestValidator) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRequestValidator) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRequestValidator) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRequestValidator) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRequestValidator) ValidateRequestBody() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateRequestBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRequestValidator) ValidateRequestParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateRequestParameters",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGateway::RequestValidator`.
func NewCfnRequestValidator(scope constructs.Construct, id *string, props *CfnRequestValidatorProps) CfnRequestValidator {
	_init_.Initialize()

	j := jsiiProxy_CfnRequestValidator{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnRequestValidator",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGateway::RequestValidator`.
func NewCfnRequestValidator_Override(c CfnRequestValidator, scope constructs.Construct, id *string, props *CfnRequestValidatorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnRequestValidator",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnRequestValidator) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnRequestValidator) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

func (j *jsiiProxy_CfnRequestValidator) SetValidateRequestBody(val interface{}) {
	_jsii_.Set(
		j,
		"validateRequestBody",
		val,
	)
}

func (j *jsiiProxy_CfnRequestValidator) SetValidateRequestParameters(val interface{}) {
	_jsii_.Set(
		j,
		"validateRequestParameters",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnRequestValidator_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnRequestValidator",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnRequestValidator_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnRequestValidator",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnRequestValidator_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnRequestValidator",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRequestValidator_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.CfnRequestValidator",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRequestValidator) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnRequestValidator) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnRequestValidator) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnRequestValidator) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnRequestValidator) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnRequestValidator) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnRequestValidator) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnRequestValidator) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRequestValidator) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRequestValidator) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnRequestValidator) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRequestValidator) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRequestValidator) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRequestValidator) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRequestValidator) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnRequestValidator`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRequestValidatorProps := &cfnRequestValidatorProps{
//   	restApiId: jsii.String("restApiId"),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	validateRequestBody: jsii.Boolean(false),
//   	validateRequestParameters: jsii.Boolean(false),
//   }
//
type CfnRequestValidatorProps struct {
	// The identifier of the targeted API entity.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// The name of this request validator.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Indicates whether to validate the request body according to the configured schema for the targeted API and method.
	ValidateRequestBody interface{} `field:"optional" json:"validateRequestBody" yaml:"validateRequestBody"`
	// Indicates whether to validate request parameters.
	ValidateRequestParameters interface{} `field:"optional" json:"validateRequestParameters" yaml:"validateRequestParameters"`
}

// A CloudFormation `AWS::ApiGateway::Resource`.
//
// The `AWS::ApiGateway::Resource` resource creates a resource in an API.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResource := awscdk.Aws_apigateway.NewCfnResource(this, jsii.String("MyCfnResource"), &cfnResourceProps{
//   	parentId: jsii.String("parentId"),
//   	pathPart: jsii.String("pathPart"),
//   	restApiId: jsii.String("restApiId"),
//   })
//
type CfnResource interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ID for the resource.
	//
	// For example: `abc123` .
	AttrResourceId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// If you want to create a child resource, the ID of the parent resource.
	//
	// For resources without a parent, specify the `RestApi` root resource ID, such as `{ "Fn::GetAtt": ["MyRestApi", "RootResourceId"] }` .
	ParentId() *string
	SetParentId(val *string)
	// A path name for the resource.
	PathPart() *string
	SetPathPart(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The ID of the [RestApi](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html) resource in which you want to create this resource.
	RestApiId() *string
	SetRestApiId(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnResource
type jsiiProxy_CfnResource struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnResource) AttrResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) ParentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) PathPart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathPart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGateway::Resource`.
func NewCfnResource(scope constructs.Construct, id *string, props *CfnResourceProps) CfnResource {
	_init_.Initialize()

	j := jsiiProxy_CfnResource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnResource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGateway::Resource`.
func NewCfnResource_Override(c CfnResource, scope constructs.Construct, id *string, props *CfnResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnResource",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnResource) SetParentId(val *string) {
	_jsii_.Set(
		j,
		"parentId",
		val,
	)
}

func (j *jsiiProxy_CfnResource) SetPathPart(val *string) {
	_jsii_.Set(
		j,
		"pathPart",
		val,
	)
}

func (j *jsiiProxy_CfnResource) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnResource_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnResource",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnResource_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnResource",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnResource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnResource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResource_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.CfnResource",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResource) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnResource) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnResource) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnResource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnResource) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnResource) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnResource) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnResource) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResource) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResource) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnResource) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResource) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResource) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResource) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnResource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceProps := &cfnResourceProps{
//   	parentId: jsii.String("parentId"),
//   	pathPart: jsii.String("pathPart"),
//   	restApiId: jsii.String("restApiId"),
//   }
//
type CfnResourceProps struct {
	// If you want to create a child resource, the ID of the parent resource.
	//
	// For resources without a parent, specify the `RestApi` root resource ID, such as `{ "Fn::GetAtt": ["MyRestApi", "RootResourceId"] }` .
	ParentId *string `field:"required" json:"parentId" yaml:"parentId"`
	// A path name for the resource.
	PathPart *string `field:"required" json:"pathPart" yaml:"pathPart"`
	// The ID of the [RestApi](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html) resource in which you want to create this resource.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
}

// A CloudFormation `AWS::ApiGateway::RestApi`.
//
// The `AWS::ApiGateway::RestApi` resource creates a REST API. For more information, see [restapi:create](https://docs.aws.amazon.com/apigateway/api-reference/link-relation/restapi-create/) in the *Amazon API Gateway REST API Reference* .
//
// > On January 1, 2016, the Swagger Specification was donated to the [OpenAPI initiative](https://docs.aws.amazon.com/https://www.openapis.org/) , becoming the foundation of the OpenAPI Specification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var body interface{}
//   var policy interface{}
//
//   cfnRestApi := awscdk.Aws_apigateway.NewCfnRestApi(this, jsii.String("MyCfnRestApi"), &cfnRestApiProps{
//   	apiKeySourceType: jsii.String("apiKeySourceType"),
//   	binaryMediaTypes: []*string{
//   		jsii.String("binaryMediaTypes"),
//   	},
//   	body: body,
//   	bodyS3Location: &s3LocationProperty{
//   		bucket: jsii.String("bucket"),
//   		eTag: jsii.String("eTag"),
//   		key: jsii.String("key"),
//   		version: jsii.String("version"),
//   	},
//   	cloneFrom: jsii.String("cloneFrom"),
//   	description: jsii.String("description"),
//   	disableExecuteApiEndpoint: jsii.Boolean(false),
//   	endpointConfiguration: &endpointConfigurationProperty{
//   		types: []*string{
//   			jsii.String("types"),
//   		},
//   		vpcEndpointIds: []*string{
//   			jsii.String("vpcEndpointIds"),
//   		},
//   	},
//   	failOnWarnings: jsii.Boolean(false),
//   	minimumCompressionSize: jsii.Number(123),
//   	mode: jsii.String("mode"),
//   	name: jsii.String("name"),
//   	parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	policy: policy,
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnRestApi interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The source of the API key for metering requests according to a usage plan. Valid values are:.
	//
	// - `HEADER` to read the API key from the `X-API-Key` header of a request.
	// - `AUTHORIZER` to read the API key from the `UsageIdentifierKey` from a Lambda authorizer.
	ApiKeySourceType() *string
	SetApiKeySourceType(val *string)
	// The root resource ID for a `RestApi` resource, such as `a0bc123d4e` .
	AttrRootResourceId() *string
	// The list of binary media types that are supported by the `RestApi` resource.
	//
	// Use `~1` instead of `/` in the media types, for example `image~1png` or `application~1octet-stream` . By default, `RestApi` supports only UTF-8-encoded text payloads. Duplicates are not allowed. For more information, see [Enable Support for Binary Payloads in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-payload-encodings.html) in the *API Gateway Developer Guide* .
	BinaryMediaTypes() *[]*string
	SetBinaryMediaTypes(val *[]*string)
	// An OpenAPI specification that defines a set of RESTful APIs in JSON format.
	//
	// For YAML templates, you can also provide the specification in YAML format.
	Body() interface{}
	SetBody(val interface{})
	// The Amazon Simple Storage Service (Amazon S3) location that points to an OpenAPI file, which defines a set of RESTful APIs in JSON or YAML format.
	BodyS3Location() interface{}
	SetBodyS3Location(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The ID of the `RestApi` resource that you want to clone.
	CloneFrom() *string
	SetCloneFrom(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A description of the `RestApi` resource.
	Description() *string
	SetDescription(val *string)
	// Specifies whether clients can invoke your API by using the default `execute-api` endpoint.
	//
	// By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint.
	DisableExecuteApiEndpoint() interface{}
	SetDisableExecuteApiEndpoint(val interface{})
	// A list of the endpoint types of the API.
	//
	// Use this property when creating an API. When importing an existing API, specify the endpoint configuration types using the `Parameters` property.
	EndpointConfiguration() interface{}
	SetEndpointConfiguration(val interface{})
	// Indicates whether to roll back the resource if a warning occurs while API Gateway is creating the `RestApi` resource.
	FailOnWarnings() interface{}
	SetFailOnWarnings(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// A nullable integer that is used to enable compression (with non-negative between 0 and 10485760 (10M) bytes, inclusive) or disable compression (with a null value) on an API.
	//
	// When compression is enabled, compression or decompression is not applied on the payload if the payload size is smaller than this value. Setting it to zero allows compression for any payload size.
	MinimumCompressionSize() *float64
	SetMinimumCompressionSize(val *float64)
	// This property applies only when you use OpenAPI to define your REST API.
	//
	// The `Mode` determines how API Gateway handles resource updates.
	//
	// Valid values are `overwrite` or `merge` .
	//
	// For `overwrite` , the new API definition replaces the existing one. The existing API identifier remains unchanged.
	//
	// For `merge` , the new API definition takes precedence, but any container types such as endpoint configurations and binary media types are merged with the existing API. Use `merge` to define top-level `RestApi` properties in addition to using OpenAPI. Generally, it's preferred to use API Gateway's OpenAPI extensions to model these properties.
	//
	// If you don't specify this property, a default value is chosen. For REST APIs created before March 29, 2021, the default is `overwrite` . Otherwise, the default value is `merge` .
	Mode() *string
	SetMode(val *string)
	// A name for the `RestApi` resource.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Custom header parameters for the request.
	Parameters() interface{}
	SetParameters(val interface{})
	// A policy document that contains the permissions for the `RestApi` resource.
	//
	// To set the ARN for the policy, use the `!Join` intrinsic function with `""` as delimiter and values of `"execute-api:/"` and `"*"` .
	Policy() interface{}
	SetPolicy(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of arbitrary tags (key-value pairs) to associate with the API.
	Tags() awscdk.TagManager
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnRestApi
type jsiiProxy_CfnRestApi struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnRestApi) ApiKeySourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeySourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRestApi) AttrRootResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrRootResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRestApi) BinaryMediaTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"binaryMediaTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRestApi) Body() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRestApi) BodyS3Location() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bodyS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRestApi) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRestApi) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRestApi) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRestApi) CloneFrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloneFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRestApi) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRestApi) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRestApi) DisableExecuteApiEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableExecuteApiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRestApi) EndpointConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRestApi) FailOnWarnings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failOnWarnings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRestApi) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRestApi) MinimumCompressionSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumCompressionSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRestApi) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRestApi) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRestApi) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRestApi) Parameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRestApi) Policy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRestApi) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRestApi) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRestApi) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRestApi) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRestApi) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGateway::RestApi`.
func NewCfnRestApi(scope constructs.Construct, id *string, props *CfnRestApiProps) CfnRestApi {
	_init_.Initialize()

	j := jsiiProxy_CfnRestApi{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnRestApi",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGateway::RestApi`.
func NewCfnRestApi_Override(c CfnRestApi, scope constructs.Construct, id *string, props *CfnRestApiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnRestApi",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnRestApi) SetApiKeySourceType(val *string) {
	_jsii_.Set(
		j,
		"apiKeySourceType",
		val,
	)
}

func (j *jsiiProxy_CfnRestApi) SetBinaryMediaTypes(val *[]*string) {
	_jsii_.Set(
		j,
		"binaryMediaTypes",
		val,
	)
}

func (j *jsiiProxy_CfnRestApi) SetBody(val interface{}) {
	_jsii_.Set(
		j,
		"body",
		val,
	)
}

func (j *jsiiProxy_CfnRestApi) SetBodyS3Location(val interface{}) {
	_jsii_.Set(
		j,
		"bodyS3Location",
		val,
	)
}

func (j *jsiiProxy_CfnRestApi) SetCloneFrom(val *string) {
	_jsii_.Set(
		j,
		"cloneFrom",
		val,
	)
}

func (j *jsiiProxy_CfnRestApi) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnRestApi) SetDisableExecuteApiEndpoint(val interface{}) {
	_jsii_.Set(
		j,
		"disableExecuteApiEndpoint",
		val,
	)
}

func (j *jsiiProxy_CfnRestApi) SetEndpointConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"endpointConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnRestApi) SetFailOnWarnings(val interface{}) {
	_jsii_.Set(
		j,
		"failOnWarnings",
		val,
	)
}

func (j *jsiiProxy_CfnRestApi) SetMinimumCompressionSize(val *float64) {
	_jsii_.Set(
		j,
		"minimumCompressionSize",
		val,
	)
}

func (j *jsiiProxy_CfnRestApi) SetMode(val *string) {
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_CfnRestApi) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnRestApi) SetParameters(val interface{}) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_CfnRestApi) SetPolicy(val interface{}) {
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnRestApi_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnRestApi",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnRestApi_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnRestApi",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnRestApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnRestApi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRestApi_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.CfnRestApi",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRestApi) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnRestApi) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnRestApi) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnRestApi) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnRestApi) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnRestApi) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnRestApi) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnRestApi) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRestApi) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRestApi) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnRestApi) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRestApi) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRestApi) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRestApi) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRestApi) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The `EndpointConfiguration` property type specifies the endpoint types of a REST API.
//
// `EndpointConfiguration` is a property of the [AWS::ApiGateway::RestApi](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointConfigurationProperty := &endpointConfigurationProperty{
//   	types: []*string{
//   		jsii.String("types"),
//   	},
//   	vpcEndpointIds: []*string{
//   		jsii.String("vpcEndpointIds"),
//   	},
//   }
//
type CfnRestApi_EndpointConfigurationProperty struct {
	// A list of endpoint types of an API or its custom domain name. Valid values include:.
	//
	// - `EDGE` : For an edge-optimized API and its custom domain name.
	// - `REGIONAL` : For a regional API and its custom domain name.
	// - `PRIVATE` : For a private API.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
	// A list of VPC endpoint IDs of an API ( [AWS::ApiGateway::RestApi](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html) ) against which to create Route53 ALIASes. It is only supported for `PRIVATE` endpoint type.
	VpcEndpointIds *[]*string `field:"optional" json:"vpcEndpointIds" yaml:"vpcEndpointIds"`
}

// `S3Location` is a property of the [AWS::ApiGateway::RestApi](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html) resource that specifies the Amazon S3 location of a OpenAPI (formerly Swagger) file that defines a set of RESTful APIs in JSON or YAML.
//
// > On January 1, 2016, the Swagger Specification was donated to the [OpenAPI initiative](https://docs.aws.amazon.com/https://www.openapis.org/) , becoming the foundation of the OpenAPI Specification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LocationProperty := &s3LocationProperty{
//   	bucket: jsii.String("bucket"),
//   	eTag: jsii.String("eTag"),
//   	key: jsii.String("key"),
//   	version: jsii.String("version"),
//   }
//
type CfnRestApi_S3LocationProperty struct {
	// The name of the S3 bucket where the OpenAPI file is stored.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The Amazon S3 ETag (a file checksum) of the OpenAPI file.
	//
	// If you don't specify a value, API Gateway skips ETag validation of your OpenAPI file.
	ETag *string `field:"optional" json:"eTag" yaml:"eTag"`
	// The file name of the OpenAPI file (Amazon S3 object name).
	Key *string `field:"optional" json:"key" yaml:"key"`
	// For versioning-enabled buckets, a specific version of the OpenAPI file.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

// Properties for defining a `CfnRestApi`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var body interface{}
//   var policy interface{}
//
//   cfnRestApiProps := &cfnRestApiProps{
//   	apiKeySourceType: jsii.String("apiKeySourceType"),
//   	binaryMediaTypes: []*string{
//   		jsii.String("binaryMediaTypes"),
//   	},
//   	body: body,
//   	bodyS3Location: &s3LocationProperty{
//   		bucket: jsii.String("bucket"),
//   		eTag: jsii.String("eTag"),
//   		key: jsii.String("key"),
//   		version: jsii.String("version"),
//   	},
//   	cloneFrom: jsii.String("cloneFrom"),
//   	description: jsii.String("description"),
//   	disableExecuteApiEndpoint: jsii.Boolean(false),
//   	endpointConfiguration: &endpointConfigurationProperty{
//   		types: []*string{
//   			jsii.String("types"),
//   		},
//   		vpcEndpointIds: []*string{
//   			jsii.String("vpcEndpointIds"),
//   		},
//   	},
//   	failOnWarnings: jsii.Boolean(false),
//   	minimumCompressionSize: jsii.Number(123),
//   	mode: jsii.String("mode"),
//   	name: jsii.String("name"),
//   	parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	policy: policy,
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnRestApiProps struct {
	// The source of the API key for metering requests according to a usage plan. Valid values are:.
	//
	// - `HEADER` to read the API key from the `X-API-Key` header of a request.
	// - `AUTHORIZER` to read the API key from the `UsageIdentifierKey` from a Lambda authorizer.
	ApiKeySourceType *string `field:"optional" json:"apiKeySourceType" yaml:"apiKeySourceType"`
	// The list of binary media types that are supported by the `RestApi` resource.
	//
	// Use `~1` instead of `/` in the media types, for example `image~1png` or `application~1octet-stream` . By default, `RestApi` supports only UTF-8-encoded text payloads. Duplicates are not allowed. For more information, see [Enable Support for Binary Payloads in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-payload-encodings.html) in the *API Gateway Developer Guide* .
	BinaryMediaTypes *[]*string `field:"optional" json:"binaryMediaTypes" yaml:"binaryMediaTypes"`
	// An OpenAPI specification that defines a set of RESTful APIs in JSON format.
	//
	// For YAML templates, you can also provide the specification in YAML format.
	Body interface{} `field:"optional" json:"body" yaml:"body"`
	// The Amazon Simple Storage Service (Amazon S3) location that points to an OpenAPI file, which defines a set of RESTful APIs in JSON or YAML format.
	BodyS3Location interface{} `field:"optional" json:"bodyS3Location" yaml:"bodyS3Location"`
	// The ID of the `RestApi` resource that you want to clone.
	CloneFrom *string `field:"optional" json:"cloneFrom" yaml:"cloneFrom"`
	// A description of the `RestApi` resource.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether clients can invoke your API by using the default `execute-api` endpoint.
	//
	// By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint.
	DisableExecuteApiEndpoint interface{} `field:"optional" json:"disableExecuteApiEndpoint" yaml:"disableExecuteApiEndpoint"`
	// A list of the endpoint types of the API.
	//
	// Use this property when creating an API. When importing an existing API, specify the endpoint configuration types using the `Parameters` property.
	EndpointConfiguration interface{} `field:"optional" json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// Indicates whether to roll back the resource if a warning occurs while API Gateway is creating the `RestApi` resource.
	FailOnWarnings interface{} `field:"optional" json:"failOnWarnings" yaml:"failOnWarnings"`
	// A nullable integer that is used to enable compression (with non-negative between 0 and 10485760 (10M) bytes, inclusive) or disable compression (with a null value) on an API.
	//
	// When compression is enabled, compression or decompression is not applied on the payload if the payload size is smaller than this value. Setting it to zero allows compression for any payload size.
	MinimumCompressionSize *float64 `field:"optional" json:"minimumCompressionSize" yaml:"minimumCompressionSize"`
	// This property applies only when you use OpenAPI to define your REST API.
	//
	// The `Mode` determines how API Gateway handles resource updates.
	//
	// Valid values are `overwrite` or `merge` .
	//
	// For `overwrite` , the new API definition replaces the existing one. The existing API identifier remains unchanged.
	//
	// For `merge` , the new API definition takes precedence, but any container types such as endpoint configurations and binary media types are merged with the existing API. Use `merge` to define top-level `RestApi` properties in addition to using OpenAPI. Generally, it's preferred to use API Gateway's OpenAPI extensions to model these properties.
	//
	// If you don't specify this property, a default value is chosen. For REST APIs created before March 29, 2021, the default is `overwrite` . Otherwise, the default value is `merge` .
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// A name for the `RestApi` resource.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Custom header parameters for the request.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// A policy document that contains the permissions for the `RestApi` resource.
	//
	// To set the ARN for the policy, use the `!Join` intrinsic function with `""` as delimiter and values of `"execute-api:/"` and `"*"` .
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
	// An array of arbitrary tags (key-value pairs) to associate with the API.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::ApiGateway::Stage`.
//
// The `AWS::ApiGateway::Stage` resource creates a stage for a deployment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStage := awscdk.Aws_apigateway.NewCfnStage(this, jsii.String("MyCfnStage"), &cfnStageProps{
//   	restApiId: jsii.String("restApiId"),
//
//   	// the properties below are optional
//   	accessLogSetting: &accessLogSettingProperty{
//   		destinationArn: jsii.String("destinationArn"),
//   		format: jsii.String("format"),
//   	},
//   	cacheClusterEnabled: jsii.Boolean(false),
//   	cacheClusterSize: jsii.String("cacheClusterSize"),
//   	canarySetting: &canarySettingProperty{
//   		deploymentId: jsii.String("deploymentId"),
//   		percentTraffic: jsii.Number(123),
//   		stageVariableOverrides: map[string]*string{
//   			"stageVariableOverridesKey": jsii.String("stageVariableOverrides"),
//   		},
//   		useStageCache: jsii.Boolean(false),
//   	},
//   	clientCertificateId: jsii.String("clientCertificateId"),
//   	deploymentId: jsii.String("deploymentId"),
//   	description: jsii.String("description"),
//   	documentationVersion: jsii.String("documentationVersion"),
//   	methodSettings: []interface{}{
//   		&methodSettingProperty{
//   			cacheDataEncrypted: jsii.Boolean(false),
//   			cacheTtlInSeconds: jsii.Number(123),
//   			cachingEnabled: jsii.Boolean(false),
//   			dataTraceEnabled: jsii.Boolean(false),
//   			httpMethod: jsii.String("httpMethod"),
//   			loggingLevel: jsii.String("loggingLevel"),
//   			metricsEnabled: jsii.Boolean(false),
//   			resourcePath: jsii.String("resourcePath"),
//   			throttlingBurstLimit: jsii.Number(123),
//   			throttlingRateLimit: jsii.Number(123),
//   		},
//   	},
//   	stageName: jsii.String("stageName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	tracingEnabled: jsii.Boolean(false),
//   	variables: map[string]*string{
//   		"variablesKey": jsii.String("variables"),
//   	},
//   })
//
type CfnStage interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Specifies settings for logging access in this stage.
	AccessLogSetting() interface{}
	SetAccessLogSetting(val interface{})
	// Indicates whether cache clustering is enabled for the stage.
	CacheClusterEnabled() interface{}
	SetCacheClusterEnabled(val interface{})
	// The stage's cache cluster size.
	CacheClusterSize() *string
	SetCacheClusterSize(val *string)
	// Specifies settings for the canary deployment in this stage.
	CanarySetting() interface{}
	SetCanarySetting(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The ID of the client certificate that API Gateway uses to call your integration endpoints in the stage.
	ClientCertificateId() *string
	SetClientCertificateId(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The ID of the deployment that the stage is associated with.
	//
	// This parameter is required to create a stage.
	DeploymentId() *string
	SetDeploymentId(val *string)
	// A description of the stage.
	Description() *string
	SetDescription(val *string)
	// The version ID of the API documentation snapshot.
	DocumentationVersion() *string
	SetDocumentationVersion(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// Settings for all methods in the stage.
	MethodSettings() interface{}
	SetMethodSettings(val interface{})
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The ID of the `RestApi` resource that you're deploying with this stage.
	RestApiId() *string
	SetRestApiId(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The name of the stage, which API Gateway uses as the first path segment in the invoked Uniform Resource Identifier (URI).
	StageName() *string
	SetStageName(val *string)
	// An array of arbitrary tags (key-value pairs) to associate with the stage.
	Tags() awscdk.TagManager
	// Specifies whether active X-Ray tracing is enabled for this stage.
	//
	// For more information, see [Trace API Gateway API Execution with AWS X-Ray](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-xray.html) in the *API Gateway Developer Guide* .
	TracingEnabled() interface{}
	SetTracingEnabled(val interface{})
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// A map (string-to-string map) that defines the stage variables, where the variable name is the key and the variable value is the value.
	//
	// Variable names are limited to alphanumeric characters. Values must match the following regular expression: `[A-Za-z0-9-._~:/?#&=,]+` .
	Variables() interface{}
	SetVariables(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnStage
type jsiiProxy_CfnStage struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnStage) AccessLogSetting() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessLogSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) CacheClusterEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheClusterEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) CacheClusterSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheClusterSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) CanarySetting() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canarySetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) ClientCertificateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) DeploymentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) DocumentationVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentationVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) MethodSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"methodSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) TracingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tracingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) Variables() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGateway::Stage`.
func NewCfnStage(scope constructs.Construct, id *string, props *CfnStageProps) CfnStage {
	_init_.Initialize()

	j := jsiiProxy_CfnStage{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnStage",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGateway::Stage`.
func NewCfnStage_Override(c CfnStage, scope constructs.Construct, id *string, props *CfnStageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnStage",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnStage) SetAccessLogSetting(val interface{}) {
	_jsii_.Set(
		j,
		"accessLogSetting",
		val,
	)
}

func (j *jsiiProxy_CfnStage) SetCacheClusterEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"cacheClusterEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnStage) SetCacheClusterSize(val *string) {
	_jsii_.Set(
		j,
		"cacheClusterSize",
		val,
	)
}

func (j *jsiiProxy_CfnStage) SetCanarySetting(val interface{}) {
	_jsii_.Set(
		j,
		"canarySetting",
		val,
	)
}

func (j *jsiiProxy_CfnStage) SetClientCertificateId(val *string) {
	_jsii_.Set(
		j,
		"clientCertificateId",
		val,
	)
}

func (j *jsiiProxy_CfnStage) SetDeploymentId(val *string) {
	_jsii_.Set(
		j,
		"deploymentId",
		val,
	)
}

func (j *jsiiProxy_CfnStage) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnStage) SetDocumentationVersion(val *string) {
	_jsii_.Set(
		j,
		"documentationVersion",
		val,
	)
}

func (j *jsiiProxy_CfnStage) SetMethodSettings(val interface{}) {
	_jsii_.Set(
		j,
		"methodSettings",
		val,
	)
}

func (j *jsiiProxy_CfnStage) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

func (j *jsiiProxy_CfnStage) SetStageName(val *string) {
	_jsii_.Set(
		j,
		"stageName",
		val,
	)
}

func (j *jsiiProxy_CfnStage) SetTracingEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"tracingEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnStage) SetVariables(val interface{}) {
	_jsii_.Set(
		j,
		"variables",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnStage_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnStage",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnStage_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnStage",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnStage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnStage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStage_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.CfnStage",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStage) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnStage) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnStage) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnStage) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnStage) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnStage) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnStage) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnStage) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStage) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStage) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnStage) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnStage) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStage) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStage) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStage) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The `AccessLogSetting` property type specifies settings for logging access in this stage.
//
// `AccessLogSetting` is a property of the [AWS::ApiGateway::Stage](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessLogSettingProperty := &accessLogSettingProperty{
//   	destinationArn: jsii.String("destinationArn"),
//   	format: jsii.String("format"),
//   }
//
type CfnStage_AccessLogSettingProperty struct {
	// The Amazon Resource Name (ARN) of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs.
	//
	// If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with `amazon-apigateway-` . This parameter is required to enable access logging.
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
	// A single line format of the access logs of data, as specified by selected [$context variables](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html#context-variable-reference) . The format must include at least `$context.requestId` . This parameter is required to enable access logging.
	Format *string `field:"optional" json:"format" yaml:"format"`
}

// The `CanarySetting` property type specifies settings for the canary deployment in this stage.
//
// `CanarySetting` is a property of the [AWS::ApiGateway::Stage](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   canarySettingProperty := &canarySettingProperty{
//   	deploymentId: jsii.String("deploymentId"),
//   	percentTraffic: jsii.Number(123),
//   	stageVariableOverrides: map[string]*string{
//   		"stageVariableOverridesKey": jsii.String("stageVariableOverrides"),
//   	},
//   	useStageCache: jsii.Boolean(false),
//   }
//
type CfnStage_CanarySettingProperty struct {
	// The identifier of the deployment that the stage points to.
	DeploymentId *string `field:"optional" json:"deploymentId" yaml:"deploymentId"`
	// The percentage (0-100) of traffic diverted to a canary deployment.
	PercentTraffic *float64 `field:"optional" json:"percentTraffic" yaml:"percentTraffic"`
	// Stage variables overridden for a canary release deployment, including new stage variables introduced in the canary.
	//
	// These stage variables are represented as a string-to-string map between stage variable names and their values.
	//
	// Duplicates are not allowed.
	StageVariableOverrides interface{} `field:"optional" json:"stageVariableOverrides" yaml:"stageVariableOverrides"`
	// Whether the canary deployment uses the stage cache or not.
	UseStageCache interface{} `field:"optional" json:"useStageCache" yaml:"useStageCache"`
}

// The `MethodSetting` property type configures settings for all methods in a stage.
//
// The `MethodSettings` property of the `AWS::ApiGateway::Stage` resource contains a list of `MethodSetting` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   methodSettingProperty := &methodSettingProperty{
//   	cacheDataEncrypted: jsii.Boolean(false),
//   	cacheTtlInSeconds: jsii.Number(123),
//   	cachingEnabled: jsii.Boolean(false),
//   	dataTraceEnabled: jsii.Boolean(false),
//   	httpMethod: jsii.String("httpMethod"),
//   	loggingLevel: jsii.String("loggingLevel"),
//   	metricsEnabled: jsii.Boolean(false),
//   	resourcePath: jsii.String("resourcePath"),
//   	throttlingBurstLimit: jsii.Number(123),
//   	throttlingRateLimit: jsii.Number(123),
//   }
//
type CfnStage_MethodSettingProperty struct {
	// Indicates whether the cached responses are encrypted.
	CacheDataEncrypted interface{} `field:"optional" json:"cacheDataEncrypted" yaml:"cacheDataEncrypted"`
	// The time-to-live (TTL) period, in seconds, that specifies how long API Gateway caches responses.
	CacheTtlInSeconds *float64 `field:"optional" json:"cacheTtlInSeconds" yaml:"cacheTtlInSeconds"`
	// Indicates whether responses are cached and returned for requests.
	//
	// You must enable a cache cluster on the stage to cache responses.
	CachingEnabled interface{} `field:"optional" json:"cachingEnabled" yaml:"cachingEnabled"`
	// Indicates whether data trace logging is enabled for methods in the stage.
	//
	// API Gateway pushes these logs to Amazon CloudWatch Logs.
	DataTraceEnabled interface{} `field:"optional" json:"dataTraceEnabled" yaml:"dataTraceEnabled"`
	// The HTTP method.
	//
	// To apply settings to multiple resources and methods, specify an asterisk ( `*` ) for the `HttpMethod` and `/*` for the `ResourcePath` . This parameter is required when you specify a `MethodSetting` .
	HttpMethod *string `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// The logging level for this method.
	//
	// For valid values, see the `loggingLevel` property of the [Stage](https://docs.aws.amazon.com/apigateway/api-reference/resource/stage/#loggingLevel) resource in the *Amazon API Gateway API Reference* .
	LoggingLevel *string `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
	// Indicates whether Amazon CloudWatch metrics are enabled for methods in the stage.
	MetricsEnabled interface{} `field:"optional" json:"metricsEnabled" yaml:"metricsEnabled"`
	// The resource path for this method.
	//
	// Forward slashes ( `/` ) are encoded as `~1` and the initial slash must include a forward slash. For example, the path value `/resource/subresource` must be encoded as `/~1resource~1subresource` . To specify the root path, use only a slash ( `/` ). To apply settings to multiple resources and methods, specify an asterisk ( `*` ) for the `HttpMethod` and `/*` for the `ResourcePath` . This parameter is required when you specify a `MethodSetting` .
	ResourcePath *string `field:"optional" json:"resourcePath" yaml:"resourcePath"`
	// The number of burst requests per second that API Gateway permits across all APIs, stages, and methods in your AWS account .
	//
	// For more information, see [Manage API Request Throttling](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html) in the *API Gateway Developer Guide* .
	ThrottlingBurstLimit *float64 `field:"optional" json:"throttlingBurstLimit" yaml:"throttlingBurstLimit"`
	// The number of steady-state requests per second that API Gateway permits across all APIs, stages, and methods in your AWS account .
	//
	// For more information, see [Manage API Request Throttling](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html) in the *API Gateway Developer Guide* .
	ThrottlingRateLimit *float64 `field:"optional" json:"throttlingRateLimit" yaml:"throttlingRateLimit"`
}

// Properties for defining a `CfnStage`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStageProps := &cfnStageProps{
//   	restApiId: jsii.String("restApiId"),
//
//   	// the properties below are optional
//   	accessLogSetting: &accessLogSettingProperty{
//   		destinationArn: jsii.String("destinationArn"),
//   		format: jsii.String("format"),
//   	},
//   	cacheClusterEnabled: jsii.Boolean(false),
//   	cacheClusterSize: jsii.String("cacheClusterSize"),
//   	canarySetting: &canarySettingProperty{
//   		deploymentId: jsii.String("deploymentId"),
//   		percentTraffic: jsii.Number(123),
//   		stageVariableOverrides: map[string]*string{
//   			"stageVariableOverridesKey": jsii.String("stageVariableOverrides"),
//   		},
//   		useStageCache: jsii.Boolean(false),
//   	},
//   	clientCertificateId: jsii.String("clientCertificateId"),
//   	deploymentId: jsii.String("deploymentId"),
//   	description: jsii.String("description"),
//   	documentationVersion: jsii.String("documentationVersion"),
//   	methodSettings: []interface{}{
//   		&methodSettingProperty{
//   			cacheDataEncrypted: jsii.Boolean(false),
//   			cacheTtlInSeconds: jsii.Number(123),
//   			cachingEnabled: jsii.Boolean(false),
//   			dataTraceEnabled: jsii.Boolean(false),
//   			httpMethod: jsii.String("httpMethod"),
//   			loggingLevel: jsii.String("loggingLevel"),
//   			metricsEnabled: jsii.Boolean(false),
//   			resourcePath: jsii.String("resourcePath"),
//   			throttlingBurstLimit: jsii.Number(123),
//   			throttlingRateLimit: jsii.Number(123),
//   		},
//   	},
//   	stageName: jsii.String("stageName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	tracingEnabled: jsii.Boolean(false),
//   	variables: map[string]*string{
//   		"variablesKey": jsii.String("variables"),
//   	},
//   }
//
type CfnStageProps struct {
	// The ID of the `RestApi` resource that you're deploying with this stage.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// Specifies settings for logging access in this stage.
	AccessLogSetting interface{} `field:"optional" json:"accessLogSetting" yaml:"accessLogSetting"`
	// Indicates whether cache clustering is enabled for the stage.
	CacheClusterEnabled interface{} `field:"optional" json:"cacheClusterEnabled" yaml:"cacheClusterEnabled"`
	// The stage's cache cluster size.
	CacheClusterSize *string `field:"optional" json:"cacheClusterSize" yaml:"cacheClusterSize"`
	// Specifies settings for the canary deployment in this stage.
	CanarySetting interface{} `field:"optional" json:"canarySetting" yaml:"canarySetting"`
	// The ID of the client certificate that API Gateway uses to call your integration endpoints in the stage.
	ClientCertificateId *string `field:"optional" json:"clientCertificateId" yaml:"clientCertificateId"`
	// The ID of the deployment that the stage is associated with.
	//
	// This parameter is required to create a stage.
	DeploymentId *string `field:"optional" json:"deploymentId" yaml:"deploymentId"`
	// A description of the stage.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The version ID of the API documentation snapshot.
	DocumentationVersion *string `field:"optional" json:"documentationVersion" yaml:"documentationVersion"`
	// Settings for all methods in the stage.
	MethodSettings interface{} `field:"optional" json:"methodSettings" yaml:"methodSettings"`
	// The name of the stage, which API Gateway uses as the first path segment in the invoked Uniform Resource Identifier (URI).
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
	// An array of arbitrary tags (key-value pairs) to associate with the stage.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies whether active X-Ray tracing is enabled for this stage.
	//
	// For more information, see [Trace API Gateway API Execution with AWS X-Ray](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-xray.html) in the *API Gateway Developer Guide* .
	TracingEnabled interface{} `field:"optional" json:"tracingEnabled" yaml:"tracingEnabled"`
	// A map (string-to-string map) that defines the stage variables, where the variable name is the key and the variable value is the value.
	//
	// Variable names are limited to alphanumeric characters. Values must match the following regular expression: `[A-Za-z0-9-._~:/?#&=,]+` .
	Variables interface{} `field:"optional" json:"variables" yaml:"variables"`
}

// A CloudFormation `AWS::ApiGateway::UsagePlan`.
//
// The `AWS::ApiGateway::UsagePlan` resource creates a usage plan for deployed APIs. A usage plan sets a target for the throttling and quota limits on individual client API keys. For more information, see [Creating and Using API Usage Plans in Amazon API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-api-usage-plans.html) in the *API Gateway Developer Guide* .
//
// In some cases clients can exceed the targets that you set. Don’t rely on usage plans to control costs. Consider using [AWS Budgets](https://docs.aws.amazon.com/cost-management/latest/userguide/budgets-managing-costs.html) to monitor costs and [AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) to manage API requests.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUsagePlan := awscdk.Aws_apigateway.NewCfnUsagePlan(this, jsii.String("MyCfnUsagePlan"), &cfnUsagePlanProps{
//   	apiStages: []interface{}{
//   		&apiStageProperty{
//   			apiId: jsii.String("apiId"),
//   			stage: jsii.String("stage"),
//   			throttle: map[string]interface{}{
//   				"throttleKey": &ThrottleSettingsProperty{
//   					"burstLimit": jsii.Number(123),
//   					"rateLimit": jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	description: jsii.String("description"),
//   	quota: &quotaSettingsProperty{
//   		limit: jsii.Number(123),
//   		offset: jsii.Number(123),
//   		period: jsii.String("period"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	throttle: &throttleSettingsProperty{
//   		burstLimit: jsii.Number(123),
//   		rateLimit: jsii.Number(123),
//   	},
//   	usagePlanName: jsii.String("usagePlanName"),
//   })
//
type CfnUsagePlan interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The API stages to associate with this usage plan.
	ApiStages() interface{}
	SetApiStages(val interface{})
	// The ID for the usage plan.
	//
	// For example: `abc123` .
	AttrId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A description of the usage plan.
	Description() *string
	SetDescription(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Configures the number of requests that users can make within a given interval.
	Quota() interface{}
	SetQuota(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of arbitrary tags (key-value pairs) to associate with the usage plan.
	Tags() awscdk.TagManager
	// Configures the overall request rate (average requests per second) and burst capacity.
	Throttle() interface{}
	SetThrottle(val interface{})
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// A name for the usage plan.
	UsagePlanName() *string
	SetUsagePlanName(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnUsagePlan
type jsiiProxy_CfnUsagePlan struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnUsagePlan) ApiStages() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiStages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUsagePlan) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUsagePlan) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUsagePlan) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUsagePlan) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUsagePlan) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUsagePlan) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUsagePlan) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUsagePlan) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUsagePlan) Quota() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"quota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUsagePlan) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUsagePlan) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUsagePlan) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUsagePlan) Throttle() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"throttle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUsagePlan) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUsagePlan) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUsagePlan) UsagePlanName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usagePlanName",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGateway::UsagePlan`.
func NewCfnUsagePlan(scope constructs.Construct, id *string, props *CfnUsagePlanProps) CfnUsagePlan {
	_init_.Initialize()

	j := jsiiProxy_CfnUsagePlan{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnUsagePlan",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGateway::UsagePlan`.
func NewCfnUsagePlan_Override(c CfnUsagePlan, scope constructs.Construct, id *string, props *CfnUsagePlanProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnUsagePlan",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnUsagePlan) SetApiStages(val interface{}) {
	_jsii_.Set(
		j,
		"apiStages",
		val,
	)
}

func (j *jsiiProxy_CfnUsagePlan) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnUsagePlan) SetQuota(val interface{}) {
	_jsii_.Set(
		j,
		"quota",
		val,
	)
}

func (j *jsiiProxy_CfnUsagePlan) SetThrottle(val interface{}) {
	_jsii_.Set(
		j,
		"throttle",
		val,
	)
}

func (j *jsiiProxy_CfnUsagePlan) SetUsagePlanName(val *string) {
	_jsii_.Set(
		j,
		"usagePlanName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnUsagePlan_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnUsagePlan",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnUsagePlan_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnUsagePlan",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnUsagePlan_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnUsagePlan",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUsagePlan_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.CfnUsagePlan",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUsagePlan) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnUsagePlan) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnUsagePlan) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnUsagePlan) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnUsagePlan) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnUsagePlan) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnUsagePlan) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnUsagePlan) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUsagePlan) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUsagePlan) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnUsagePlan) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnUsagePlan) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUsagePlan) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUsagePlan) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUsagePlan) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// `ApiStage` is a property of the [AWS::ApiGateway::UsagePlan](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html) resource that specifies which stages and APIs to associate with a usage plan.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiStageProperty := &apiStageProperty{
//   	apiId: jsii.String("apiId"),
//   	stage: jsii.String("stage"),
//   	throttle: map[string]interface{}{
//   		"throttleKey": &ThrottleSettingsProperty{
//   			"burstLimit": jsii.Number(123),
//   			"rateLimit": jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnUsagePlan_ApiStageProperty struct {
	// The ID of an API that is in the specified `Stage` property that you want to associate with the usage plan.
	ApiId *string `field:"optional" json:"apiId" yaml:"apiId"`
	// The name of the stage to associate with the usage plan.
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
	// Map containing method-level throttling information for an API stage in a usage plan.
	//
	// The key for the map is the path and method for which to configure custom throttling, for example, "/pets/GET".
	//
	// Duplicates are not allowed.
	Throttle interface{} `field:"optional" json:"throttle" yaml:"throttle"`
}

// `QuotaSettings` is a property of the [AWS::ApiGateway::UsagePlan](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html) resource that specifies a target for the maximum number of requests users can make to your REST APIs.
//
// In some cases clients can exceed the targets that you set. Don’t rely on usage plans to control costs. Consider using [AWS Budgets](https://docs.aws.amazon.com/cost-management/latest/userguide/budgets-managing-costs.html) to monitor costs and [AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) to manage API requests.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   quotaSettingsProperty := &quotaSettingsProperty{
//   	limit: jsii.Number(123),
//   	offset: jsii.Number(123),
//   	period: jsii.String("period"),
//   }
//
type CfnUsagePlan_QuotaSettingsProperty struct {
	// The target maximum number of requests that can be made in a given time period.
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// The day that a time period starts.
	//
	// For example, with a time period of `WEEK` , an offset of `0` starts on Sunday, and an offset of `1` starts on Monday.
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
	// The time period for which the target maximum limit of requests applies, such as `DAY` or `WEEK` .
	//
	// For valid values, see the period property for the [UsagePlan](https://docs.aws.amazon.com/apigateway/api-reference/resource/usage-plan) resource in the *Amazon API Gateway REST API Reference* .
	Period *string `field:"optional" json:"period" yaml:"period"`
}

// `ThrottleSettings` is a property of the [AWS::ApiGateway::UsagePlan](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html) resource that specifies the overall request rate (average requests per second) and burst capacity when users call your REST APIs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   throttleSettingsProperty := &throttleSettingsProperty{
//   	burstLimit: jsii.Number(123),
//   	rateLimit: jsii.Number(123),
//   }
//
type CfnUsagePlan_ThrottleSettingsProperty struct {
	// The API target request burst rate limit.
	//
	// This allows more requests through for a period of time than the target rate limit. For more information about request throttling, see [Manage API Request Throttling](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html) in the *API Gateway Developer Guide* .
	BurstLimit *float64 `field:"optional" json:"burstLimit" yaml:"burstLimit"`
	// The API target request steady-state rate limit.
	//
	// For more information about request throttling, see [Manage API Request Throttling](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html) in the *API Gateway Developer Guide* .
	RateLimit *float64 `field:"optional" json:"rateLimit" yaml:"rateLimit"`
}

// A CloudFormation `AWS::ApiGateway::UsagePlanKey`.
//
// The `AWS::ApiGateway::UsagePlanKey` resource associates an API key with a usage plan. This association determines which users the usage plan is applied to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUsagePlanKey := awscdk.Aws_apigateway.NewCfnUsagePlanKey(this, jsii.String("MyCfnUsagePlanKey"), &cfnUsagePlanKeyProps{
//   	keyId: jsii.String("keyId"),
//   	keyType: jsii.String("keyType"),
//   	usagePlanId: jsii.String("usagePlanId"),
//   })
//
type CfnUsagePlanKey interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The ID of the usage plan key.
	KeyId() *string
	SetKeyId(val *string)
	// The type of usage plan key.
	//
	// Currently, the only valid key type is `API_KEY` .
	KeyType() *string
	SetKeyType(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// The ID of the usage plan.
	UsagePlanId() *string
	SetUsagePlanId(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnUsagePlanKey
type jsiiProxy_CfnUsagePlanKey struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnUsagePlanKey) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUsagePlanKey) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUsagePlanKey) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUsagePlanKey) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUsagePlanKey) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUsagePlanKey) KeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUsagePlanKey) KeyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUsagePlanKey) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUsagePlanKey) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUsagePlanKey) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUsagePlanKey) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUsagePlanKey) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUsagePlanKey) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUsagePlanKey) UsagePlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usagePlanId",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGateway::UsagePlanKey`.
func NewCfnUsagePlanKey(scope constructs.Construct, id *string, props *CfnUsagePlanKeyProps) CfnUsagePlanKey {
	_init_.Initialize()

	j := jsiiProxy_CfnUsagePlanKey{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnUsagePlanKey",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGateway::UsagePlanKey`.
func NewCfnUsagePlanKey_Override(c CfnUsagePlanKey, scope constructs.Construct, id *string, props *CfnUsagePlanKeyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnUsagePlanKey",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnUsagePlanKey) SetKeyId(val *string) {
	_jsii_.Set(
		j,
		"keyId",
		val,
	)
}

func (j *jsiiProxy_CfnUsagePlanKey) SetKeyType(val *string) {
	_jsii_.Set(
		j,
		"keyType",
		val,
	)
}

func (j *jsiiProxy_CfnUsagePlanKey) SetUsagePlanId(val *string) {
	_jsii_.Set(
		j,
		"usagePlanId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnUsagePlanKey_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnUsagePlanKey",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnUsagePlanKey_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnUsagePlanKey",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnUsagePlanKey_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnUsagePlanKey",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUsagePlanKey_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.CfnUsagePlanKey",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUsagePlanKey) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnUsagePlanKey) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnUsagePlanKey) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnUsagePlanKey) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnUsagePlanKey) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnUsagePlanKey) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnUsagePlanKey) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnUsagePlanKey) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUsagePlanKey) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUsagePlanKey) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnUsagePlanKey) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnUsagePlanKey) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUsagePlanKey) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUsagePlanKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUsagePlanKey) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnUsagePlanKey`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUsagePlanKeyProps := &cfnUsagePlanKeyProps{
//   	keyId: jsii.String("keyId"),
//   	keyType: jsii.String("keyType"),
//   	usagePlanId: jsii.String("usagePlanId"),
//   }
//
type CfnUsagePlanKeyProps struct {
	// The ID of the usage plan key.
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
	// The type of usage plan key.
	//
	// Currently, the only valid key type is `API_KEY` .
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
	// The ID of the usage plan.
	UsagePlanId *string `field:"required" json:"usagePlanId" yaml:"usagePlanId"`
}

// Properties for defining a `CfnUsagePlan`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUsagePlanProps := &cfnUsagePlanProps{
//   	apiStages: []interface{}{
//   		&apiStageProperty{
//   			apiId: jsii.String("apiId"),
//   			stage: jsii.String("stage"),
//   			throttle: map[string]interface{}{
//   				"throttleKey": &ThrottleSettingsProperty{
//   					"burstLimit": jsii.Number(123),
//   					"rateLimit": jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	description: jsii.String("description"),
//   	quota: &quotaSettingsProperty{
//   		limit: jsii.Number(123),
//   		offset: jsii.Number(123),
//   		period: jsii.String("period"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	throttle: &throttleSettingsProperty{
//   		burstLimit: jsii.Number(123),
//   		rateLimit: jsii.Number(123),
//   	},
//   	usagePlanName: jsii.String("usagePlanName"),
//   }
//
type CfnUsagePlanProps struct {
	// The API stages to associate with this usage plan.
	ApiStages interface{} `field:"optional" json:"apiStages" yaml:"apiStages"`
	// A description of the usage plan.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Configures the number of requests that users can make within a given interval.
	Quota interface{} `field:"optional" json:"quota" yaml:"quota"`
	// An array of arbitrary tags (key-value pairs) to associate with the usage plan.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Configures the overall request rate (average requests per second) and burst capacity.
	Throttle interface{} `field:"optional" json:"throttle" yaml:"throttle"`
	// A name for the usage plan.
	UsagePlanName *string `field:"optional" json:"usagePlanName" yaml:"usagePlanName"`
}

// A CloudFormation `AWS::ApiGateway::VpcLink`.
//
// The `AWS::ApiGateway::VpcLink` resource creates an API Gateway VPC link for a REST API to access resources in an Amazon Virtual Private Cloud (VPC). For more information, see [vpclink:create](https://docs.aws.amazon.com/apigateway/api-reference/link-relation/vpclink-create/) in the `Amazon API Gateway REST API Reference` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVpcLink := awscdk.Aws_apigateway.NewCfnVpcLink(this, jsii.String("MyCfnVpcLink"), &cfnVpcLinkProps{
//   	name: jsii.String("name"),
//   	targetArns: []*string{
//   		jsii.String("targetArns"),
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnVpcLink interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A description of the VPC link.
	Description() *string
	SetDescription(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// A name for the VPC link.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of arbitrary tags (key-value pairs) to associate with the VPC link.
	Tags() awscdk.TagManager
	// The ARN of network load balancer of the VPC targeted by the VPC link.
	//
	// The network load balancer must be owned by the same AWS account of the API owner.
	TargetArns() *[]*string
	SetTargetArns(val *[]*string)
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnVpcLink
type jsiiProxy_CfnVpcLink struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnVpcLink) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcLink) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcLink) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcLink) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcLink) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcLink) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcLink) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcLink) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcLink) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcLink) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcLink) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcLink) TargetArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcLink) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcLink) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGateway::VpcLink`.
func NewCfnVpcLink(scope constructs.Construct, id *string, props *CfnVpcLinkProps) CfnVpcLink {
	_init_.Initialize()

	j := jsiiProxy_CfnVpcLink{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnVpcLink",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGateway::VpcLink`.
func NewCfnVpcLink_Override(c CfnVpcLink, scope constructs.Construct, id *string, props *CfnVpcLinkProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CfnVpcLink",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnVpcLink) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnVpcLink) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnVpcLink) SetTargetArns(val *[]*string) {
	_jsii_.Set(
		j,
		"targetArns",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnVpcLink_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnVpcLink",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnVpcLink_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnVpcLink",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnVpcLink_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CfnVpcLink",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVpcLink_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.CfnVpcLink",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVpcLink) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnVpcLink) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnVpcLink) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnVpcLink) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnVpcLink) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnVpcLink) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnVpcLink) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnVpcLink) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVpcLink) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVpcLink) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnVpcLink) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnVpcLink) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVpcLink) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVpcLink) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVpcLink) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnVpcLink`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVpcLinkProps := &cfnVpcLinkProps{
//   	name: jsii.String("name"),
//   	targetArns: []*string{
//   		jsii.String("targetArns"),
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnVpcLinkProps struct {
	// A name for the VPC link.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN of network load balancer of the VPC targeted by the VPC link.
	//
	// The network load balancer must be owned by the same AWS account of the API owner.
	TargetArns *[]*string `field:"required" json:"targetArns" yaml:"targetArns"`
	// A description of the VPC link.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of arbitrary tags (key-value pairs) to associate with the VPC link.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// Cognito user pools based custom authorizer.
//
// Example:
//   var books resource
//   userPool := cognito.NewUserPool(this, jsii.String("UserPool"))
//
//   auth := apigateway.NewCognitoUserPoolsAuthorizer(this, jsii.String("booksAuthorizer"), &cognitoUserPoolsAuthorizerProps{
//   	cognitoUserPools: []iUserPool{
//   		userPool,
//   	},
//   })
//   books.addMethod(jsii.String("GET"), apigateway.NewHttpIntegration(jsii.String("http://amazon.com")), &methodOptions{
//   	authorizer: auth,
//   	authorizationType: apigateway.authorizationType_COGNITO,
//   })
//
type CognitoUserPoolsAuthorizer interface {
	Authorizer
	IAuthorizer
	// The authorization type of this authorizer.
	AuthorizationType() AuthorizationType
	// The ARN of the authorizer to be used in permission policies, such as IAM and resource-based grants.
	AuthorizerArn() *string
	// The id of the authorizer.
	AuthorizerId() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CognitoUserPoolsAuthorizer
type jsiiProxy_CognitoUserPoolsAuthorizer struct {
	jsiiProxy_Authorizer
	jsiiProxy_IAuthorizer
}

func (j *jsiiProxy_CognitoUserPoolsAuthorizer) AuthorizationType() AuthorizationType {
	var returns AuthorizationType
	_jsii_.Get(
		j,
		"authorizationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolsAuthorizer) AuthorizerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolsAuthorizer) AuthorizerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolsAuthorizer) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolsAuthorizer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolsAuthorizer) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolsAuthorizer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewCognitoUserPoolsAuthorizer(scope constructs.Construct, id *string, props *CognitoUserPoolsAuthorizerProps) CognitoUserPoolsAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_CognitoUserPoolsAuthorizer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CognitoUserPoolsAuthorizer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCognitoUserPoolsAuthorizer_Override(c CognitoUserPoolsAuthorizer, scope constructs.Construct, id *string, props *CognitoUserPoolsAuthorizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CognitoUserPoolsAuthorizer",
		[]interface{}{scope, id, props},
		c,
	)
}

// Return whether the given object is an Authorizer.
func CognitoUserPoolsAuthorizer_IsAuthorizer(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CognitoUserPoolsAuthorizer",
		"isAuthorizer",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CognitoUserPoolsAuthorizer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CognitoUserPoolsAuthorizer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func CognitoUserPoolsAuthorizer_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CognitoUserPoolsAuthorizer",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func CognitoUserPoolsAuthorizer_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CognitoUserPoolsAuthorizer",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolsAuthorizer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (c *jsiiProxy_CognitoUserPoolsAuthorizer) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolsAuthorizer) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolsAuthorizer) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolsAuthorizer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for CognitoUserPoolsAuthorizer.
//
// Example:
//   var books resource
//   userPool := cognito.NewUserPool(this, jsii.String("UserPool"))
//
//   auth := apigateway.NewCognitoUserPoolsAuthorizer(this, jsii.String("booksAuthorizer"), &cognitoUserPoolsAuthorizerProps{
//   	cognitoUserPools: []iUserPool{
//   		userPool,
//   	},
//   })
//   books.addMethod(jsii.String("GET"), apigateway.NewHttpIntegration(jsii.String("http://amazon.com")), &methodOptions{
//   	authorizer: auth,
//   	authorizationType: apigateway.authorizationType_COGNITO,
//   })
//
type CognitoUserPoolsAuthorizerProps struct {
	// The user pools to associate with this authorizer.
	CognitoUserPools *[]awscognito.IUserPool `field:"required" json:"cognitoUserPools" yaml:"cognitoUserPools"`
	// An optional human friendly name for the authorizer.
	//
	// Note that, this is not the primary identifier of the authorizer.
	AuthorizerName *string `field:"optional" json:"authorizerName" yaml:"authorizerName"`
	// The request header mapping expression for the bearer token.
	//
	// This is typically passed as part of the header, in which case
	// this should be `method.request.header.Authorizer` where Authorizer is the header containing the bearer token.
	// See: https://docs.aws.amazon.com/apigateway/api-reference/link-relation/authorizer-create/#identitySource
	//
	IdentitySource *string `field:"optional" json:"identitySource" yaml:"identitySource"`
	// How long APIGateway should cache the results.
	//
	// Max 1 hour.
	// Disable caching by setting this to 0.
	ResultsCacheTtl awscdk.Duration `field:"optional" json:"resultsCacheTtl" yaml:"resultsCacheTtl"`
}

// Example:
//   import elbv2 "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("VPC"))
//   nlb := elbv2.NewNetworkLoadBalancer(this, jsii.String("NLB"), &networkLoadBalancerProps{
//   	vpc: vpc,
//   })
//   link := apigateway.NewVpcLink(this, jsii.String("link"), &vpcLinkProps{
//   	targets: []iNetworkLoadBalancer{
//   		nlb,
//   	},
//   })
//
//   integration := apigateway.NewIntegration(&integrationProps{
//   	type: apigateway.integrationType_HTTP_PROXY,
//   	options: &integrationOptions{
//   		connectionType: apigateway.connectionType_VPC_LINK,
//   		vpcLink: link,
//   	},
//   })
//
type ConnectionType string

const (
	// For connections through the public routable internet.
	ConnectionType_INTERNET ConnectionType = "INTERNET"
	// For private connections between API Gateway and a network load balancer in a VPC.
	ConnectionType_VPC_LINK ConnectionType = "VPC_LINK"
)

// Example:
//   var getBookHandler function
//   var getBookIntegration lambdaIntegration
//
//
//   getBookIntegration := apigateway.NewLambdaIntegration(getBookHandler, &lambdaIntegrationOptions{
//   	contentHandling: apigateway.contentHandling_CONVERT_TO_TEXT,
//   	 // convert to base64
//   	credentialsPassthrough: jsii.Boolean(true),
//   })
//
type ContentHandling string

const (
	// Converts a request payload from a base64-encoded string to a binary blob.
	ContentHandling_CONVERT_TO_BINARY ContentHandling = "CONVERT_TO_BINARY"
	// Converts a request payload from a binary blob to a base64-encoded string.
	ContentHandling_CONVERT_TO_TEXT ContentHandling = "CONVERT_TO_TEXT"
)

// Example:
//   apigateway.NewRestApi(this, jsii.String("api"), &restApiProps{
//   	defaultCorsPreflightOptions: &corsOptions{
//   		allowOrigins: apigateway.cors_ALL_ORIGINS(),
//   		allowMethods: apigateway.*cors_ALL_METHODS(),
//   	},
//   })
//
type Cors interface {
}

// The jsii proxy struct for Cors
type jsiiProxy_Cors struct {
	_ byte // padding
}

func Cors_ALL_METHODS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.Cors",
		"ALL_METHODS",
		&returns,
	)
	return returns
}

func Cors_ALL_ORIGINS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.Cors",
		"ALL_ORIGINS",
		&returns,
	)
	return returns
}

func Cors_DEFAULT_HEADERS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.Cors",
		"DEFAULT_HEADERS",
		&returns,
	)
	return returns
}

// Example:
//   var myResource resource
//
//
//   myResource.addCorsPreflight(&corsOptions{
//   	allowOrigins: []*string{
//   		jsii.String("https://amazon.com"),
//   	},
//   	allowMethods: []*string{
//   		jsii.String("GET"),
//   		jsii.String("PUT"),
//   	},
//   })
//
type CorsOptions struct {
	// Specifies the list of origins that are allowed to make requests to this resource.
	//
	// If you wish to allow all origins, specify `Cors.ALL_ORIGINS` or
	// `[ * ]`.
	//
	// Responses will include the `Access-Control-Allow-Origin` response header.
	// If `Cors.ALL_ORIGINS` is specified, the `Vary: Origin` response header will
	// also be included.
	// See: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin
	//
	AllowOrigins *[]*string `field:"required" json:"allowOrigins" yaml:"allowOrigins"`
	// The Access-Control-Allow-Credentials response header tells browsers whether to expose the response to frontend JavaScript code when the request's credentials mode (Request.credentials) is "include".
	//
	// When a request's credentials mode (Request.credentials) is "include",
	// browsers will only expose the response to frontend JavaScript code if the
	// Access-Control-Allow-Credentials value is true.
	//
	// Credentials are cookies, authorization headers or TLS client certificates.
	// See: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Credentials
	//
	AllowCredentials *bool `field:"optional" json:"allowCredentials" yaml:"allowCredentials"`
	// The Access-Control-Allow-Headers response header is used in response to a preflight request which includes the Access-Control-Request-Headers to indicate which HTTP headers can be used during the actual request.
	// See: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Headers
	//
	AllowHeaders *[]*string `field:"optional" json:"allowHeaders" yaml:"allowHeaders"`
	// The Access-Control-Allow-Methods response header specifies the method or methods allowed when accessing the resource in response to a preflight request.
	//
	// If `ANY` is specified, it will be expanded to `Cors.ALL_METHODS`.
	// See: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Methods
	//
	AllowMethods *[]*string `field:"optional" json:"allowMethods" yaml:"allowMethods"`
	// Sets Access-Control-Max-Age to -1, which means that caching is disabled.
	//
	// This option cannot be used with `maxAge`.
	DisableCache *bool `field:"optional" json:"disableCache" yaml:"disableCache"`
	// The Access-Control-Expose-Headers response header indicates which headers can be exposed as part of the response by listing their names.
	//
	// If you want clients to be able to access other headers, you have to list
	// them using the Access-Control-Expose-Headers header.
	// See: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Expose-Headers
	//
	ExposeHeaders *[]*string `field:"optional" json:"exposeHeaders" yaml:"exposeHeaders"`
	// The Access-Control-Max-Age response header indicates how long the results of a preflight request (that is the information contained in the Access-Control-Allow-Methods and Access-Control-Allow-Headers headers) can be cached.
	//
	// To disable caching altogether use `disableCache: true`.
	// See: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Max-Age
	//
	MaxAge awscdk.Duration `field:"optional" json:"maxAge" yaml:"maxAge"`
	// Specifies the response status code returned from the OPTIONS method.
	StatusCode *float64 `field:"optional" json:"statusCode" yaml:"statusCode"`
}

// A Deployment of a REST API.
//
// An immutable representation of a RestApi resource that can be called by users
// using Stages. A deployment must be associated with a Stage for it to be
// callable over the Internet.
//
// Normally, you don't need to define deployments manually. The RestApi
// construct manages a Deployment resource that represents the latest model. It
// can be accessed through `restApi.latestDeployment` (unless `deploy: false` is
// set when defining the `RestApi`).
//
// If you manually define this resource, you will need to know that since
// deployments are immutable, as long as the resource's logical ID doesn't
// change, the deployment will represent the snapshot in time in which the
// resource was created. This means that if you modify the RestApi model (i.e.
// add methods or resources), these changes will not be reflected unless a new
// deployment resource is created.
//
// To achieve this behavior, the method `addToLogicalId(data)` can be used to
// augment the logical ID generated for the deployment resource such that it
// will include arbitrary data. This is done automatically for the
// `restApi.latestDeployment` deployment.
//
// Furthermore, since a deployment does not reference any of the REST API
// resources and methods, CloudFormation will likely provision it before these
// resources are created, which means that it will represent a "half-baked"
// model. Use the `node.addDependency(dep)` method to circumvent that. This is done
// automatically for the `restApi.latestDeployment` deployment.
//
// Example:
//   // production stage
//   prdLogGroup := logs.NewLogGroup(this, jsii.String("PrdLogs"))
//   api := apigateway.NewRestApi(this, jsii.String("books"), &restApiProps{
//   	deployOptions: &stageOptions{
//   		accessLogDestination: apigateway.NewLogGroupLogDestination(prdLogGroup),
//   		accessLogFormat: apigateway.accessLogFormat.jsonWithStandardFields(),
//   	},
//   })
//   deployment := apigateway.NewDeployment(this, jsii.String("Deployment"), &deploymentProps{
//   	api: api,
//   })
//
//   // development stage
//   devLogGroup := logs.NewLogGroup(this, jsii.String("DevLogs"))
//   apigateway.NewStage(this, jsii.String("dev"), &stageProps{
//   	deployment: deployment,
//   	accessLogDestination: apigateway.NewLogGroupLogDestination(devLogGroup),
//   	accessLogFormat: apigateway.*accessLogFormat.jsonWithStandardFields(&jsonWithStandardFieldProps{
//   		caller: jsii.Boolean(false),
//   		httpMethod: jsii.Boolean(true),
//   		ip: jsii.Boolean(true),
//   		protocol: jsii.Boolean(true),
//   		requestTime: jsii.Boolean(true),
//   		resourcePath: jsii.Boolean(true),
//   		responseLength: jsii.Boolean(true),
//   		status: jsii.Boolean(true),
//   		user: jsii.Boolean(true),
//   	}),
//   })
//
type Deployment interface {
	awscdk.Resource
	Api() IRestApi
	DeploymentId() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Adds a component to the hash that determines this Deployment resource's logical ID.
	//
	// This should be called by constructs of the API Gateway model that want to
	// invalidate the deployment when their settings change. The component will
	// be resolve()ed during synthesis so tokens are welcome.
	AddToLogicalId(data interface{})
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Deployment
type jsiiProxy_Deployment struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_Deployment) Api() IRestApi {
	var returns IRestApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) DeploymentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deployment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewDeployment(scope constructs.Construct, id *string, props *DeploymentProps) Deployment {
	_init_.Initialize()

	j := jsiiProxy_Deployment{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.Deployment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDeployment_Override(d Deployment, scope constructs.Construct, id *string, props *DeploymentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.Deployment",
		[]interface{}{scope, id, props},
		d,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Deployment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.Deployment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func Deployment_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.Deployment",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Deployment_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.Deployment",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Deployment) AddToLogicalId(data interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addToLogicalId",
		[]interface{}{data},
	)
}

func (d *jsiiProxy_Deployment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (d *jsiiProxy_Deployment) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Deployment) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Deployment) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Deployment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Example:
//   // production stage
//   prdLogGroup := logs.NewLogGroup(this, jsii.String("PrdLogs"))
//   api := apigateway.NewRestApi(this, jsii.String("books"), &restApiProps{
//   	deployOptions: &stageOptions{
//   		accessLogDestination: apigateway.NewLogGroupLogDestination(prdLogGroup),
//   		accessLogFormat: apigateway.accessLogFormat.jsonWithStandardFields(),
//   	},
//   })
//   deployment := apigateway.NewDeployment(this, jsii.String("Deployment"), &deploymentProps{
//   	api: api,
//   })
//
//   // development stage
//   devLogGroup := logs.NewLogGroup(this, jsii.String("DevLogs"))
//   apigateway.NewStage(this, jsii.String("dev"), &stageProps{
//   	deployment: deployment,
//   	accessLogDestination: apigateway.NewLogGroupLogDestination(devLogGroup),
//   	accessLogFormat: apigateway.*accessLogFormat.jsonWithStandardFields(&jsonWithStandardFieldProps{
//   		caller: jsii.Boolean(false),
//   		httpMethod: jsii.Boolean(true),
//   		ip: jsii.Boolean(true),
//   		protocol: jsii.Boolean(true),
//   		requestTime: jsii.Boolean(true),
//   		resourcePath: jsii.Boolean(true),
//   		responseLength: jsii.Boolean(true),
//   		status: jsii.Boolean(true),
//   		user: jsii.Boolean(true),
//   	}),
//   })
//
type DeploymentProps struct {
	// The Rest API to deploy.
	Api IRestApi `field:"required" json:"api" yaml:"api"`
	// A description of the purpose of the API Gateway deployment.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// When an API Gateway model is updated, a new deployment will automatically be created.
	//
	// If this is true, the old API Gateway Deployment resource will not be deleted.
	// This will allow manually reverting back to a previous deployment in case for example.
	RetainDeployments *bool `field:"optional" json:"retainDeployments" yaml:"retainDeployments"`
}

// Example:
//   var acm interface{}
//
//
//   apigateway.NewDomainName(this, jsii.String("domain-name"), &domainNameProps{
//   	domainName: jsii.String("example.com"),
//   	certificate: acm.certificate_FromCertificateArn(this, jsii.String("cert"), jsii.String("arn:aws:acm:us-east-1:1111111:certificate/11-3336f1-44483d-adc7-9cd375c5169d")),
//   	mtls: &mTLSConfig{
//   		bucket: s3.NewBucket(this, jsii.String("bucket")),
//   		key: jsii.String("truststore.pem"),
//   		version: jsii.String("version"),
//   	},
//   })
//
type DomainName interface {
	awscdk.Resource
	IDomainName
	// The domain name (e.g. `example.com`).
	DomainName() *string
	// The Route53 alias target to use in order to connect a record set to this domain through an alias.
	DomainNameAliasDomainName() *string
	// The Route53 hosted zone ID to use in order to connect a record set to this domain through an alias.
	DomainNameAliasHostedZoneId() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Maps this domain to an API endpoint.
	AddBasePathMapping(targetApi IRestApi, options *BasePathMappingOptions) BasePathMapping
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for DomainName
type jsiiProxy_DomainName struct {
	internal.Type__awscdkResource
	jsiiProxy_IDomainName
}

func (j *jsiiProxy_DomainName) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DomainName) DomainNameAliasDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameAliasDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DomainName) DomainNameAliasHostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameAliasHostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DomainName) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DomainName) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DomainName) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DomainName) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewDomainName(scope constructs.Construct, id *string, props *DomainNameProps) DomainName {
	_init_.Initialize()

	j := jsiiProxy_DomainName{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.DomainName",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDomainName_Override(d DomainName, scope constructs.Construct, id *string, props *DomainNameProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.DomainName",
		[]interface{}{scope, id, props},
		d,
	)
}

// Imports an existing domain name.
func DomainName_FromDomainNameAttributes(scope constructs.Construct, id *string, attrs *DomainNameAttributes) IDomainName {
	_init_.Initialize()

	var returns IDomainName

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.DomainName",
		"fromDomainNameAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func DomainName_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.DomainName",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func DomainName_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.DomainName",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func DomainName_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.DomainName",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DomainName) AddBasePathMapping(targetApi IRestApi, options *BasePathMappingOptions) BasePathMapping {
	var returns BasePathMapping

	_jsii_.Invoke(
		d,
		"addBasePathMapping",
		[]interface{}{targetApi, options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DomainName) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (d *jsiiProxy_DomainName) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DomainName) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DomainName) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DomainName) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainNameAttributes := &domainNameAttributes{
//   	domainName: jsii.String("domainName"),
//   	domainNameAliasHostedZoneId: jsii.String("domainNameAliasHostedZoneId"),
//   	domainNameAliasTarget: jsii.String("domainNameAliasTarget"),
//   }
//
type DomainNameAttributes struct {
	// The domain name (e.g. `example.com`).
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The Route53 hosted zone ID to use in order to connect a record set to this domain through an alias.
	DomainNameAliasHostedZoneId *string `field:"required" json:"domainNameAliasHostedZoneId" yaml:"domainNameAliasHostedZoneId"`
	// The Route53 alias target to use in order to connect a record set to this domain through an alias.
	DomainNameAliasTarget *string `field:"required" json:"domainNameAliasTarget" yaml:"domainNameAliasTarget"`
}

// Example:
//   var acmCertificateForExampleCom interface{}
//
//
//   api := apigateway.NewRestApi(this, jsii.String("MyDomain"), &restApiProps{
//   	domainName: &domainNameOptions{
//   		domainName: jsii.String("example.com"),
//   		certificate: acmCertificateForExampleCom,
//   	},
//   })
//
type DomainNameOptions struct {
	// The reference to an AWS-managed certificate for use by the edge-optimized endpoint for the domain name.
	//
	// For "EDGE" domain names, the certificate
	// needs to be in the US East (N. Virginia) region.
	Certificate awscertificatemanager.ICertificate `field:"required" json:"certificate" yaml:"certificate"`
	// The custom domain name for your API.
	//
	// Uppercase letters are not supported.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The base path name that callers of the API must provide in the URL after the domain name (e.g. `example.com/base-path`). If you specify this property, it can't be an empty string.
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
	// The type of endpoint for this DomainName.
	EndpointType EndpointType `field:"optional" json:"endpointType" yaml:"endpointType"`
	// The mutual TLS authentication configuration for a custom domain name.
	Mtls *MTLSConfig `field:"optional" json:"mtls" yaml:"mtls"`
	// The Transport Layer Security (TLS) version + cipher suite for this domain name.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html
	//
	SecurityPolicy SecurityPolicy `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
}

// Example:
//   var acm interface{}
//
//
//   apigateway.NewDomainName(this, jsii.String("domain-name"), &domainNameProps{
//   	domainName: jsii.String("example.com"),
//   	certificate: acm.certificate_FromCertificateArn(this, jsii.String("cert"), jsii.String("arn:aws:acm:us-east-1:1111111:certificate/11-3336f1-44483d-adc7-9cd375c5169d")),
//   	mtls: &mTLSConfig{
//   		bucket: s3.NewBucket(this, jsii.String("bucket")),
//   		key: jsii.String("truststore.pem"),
//   		version: jsii.String("version"),
//   	},
//   })
//
type DomainNameProps struct {
	// The reference to an AWS-managed certificate for use by the edge-optimized endpoint for the domain name.
	//
	// For "EDGE" domain names, the certificate
	// needs to be in the US East (N. Virginia) region.
	Certificate awscertificatemanager.ICertificate `field:"required" json:"certificate" yaml:"certificate"`
	// The custom domain name for your API.
	//
	// Uppercase letters are not supported.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The base path name that callers of the API must provide in the URL after the domain name (e.g. `example.com/base-path`). If you specify this property, it can't be an empty string.
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
	// The type of endpoint for this DomainName.
	EndpointType EndpointType `field:"optional" json:"endpointType" yaml:"endpointType"`
	// The mutual TLS authentication configuration for a custom domain name.
	Mtls *MTLSConfig `field:"optional" json:"mtls" yaml:"mtls"`
	// The Transport Layer Security (TLS) version + cipher suite for this domain name.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html
	//
	SecurityPolicy SecurityPolicy `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
	// If specified, all requests to this domain will be mapped to the production deployment of this API.
	//
	// If you wish to map this domain to multiple APIs
	// with different base paths, don't specify this option and use
	// `addBasePathMapping`.
	Mapping IRestApi `field:"optional" json:"mapping" yaml:"mapping"`
}

// The endpoint configuration of a REST API, including VPCs and endpoint types.
//
// EndpointConfiguration is a property of the AWS::ApiGateway::RestApi resource.
//
// Example:
//   api := apigateway.NewRestApi(this, jsii.String("api"), &restApiProps{
//   	endpointConfiguration: &endpointConfiguration{
//   		types: []endpointType{
//   			apigateway.*endpointType_EDGE,
//   		},
//   	},
//   })
//
type EndpointConfiguration struct {
	// A list of endpoint types of an API or its custom domain name.
	Types *[]EndpointType `field:"required" json:"types" yaml:"types"`
	// A list of VPC Endpoints against which to create Route53 ALIASes.
	VpcEndpoints *[]awsec2.IVpcEndpoint `field:"optional" json:"vpcEndpoints" yaml:"vpcEndpoints"`
}

// Example:
//   var apiDefinition apiDefinition
//
//
//   api := apigateway.NewSpecRestApi(this, jsii.String("ExampleRestApi"), &specRestApiProps{
//   	apiDefinition: apiDefinition,
//   	endpointTypes: []endpointType{
//   		apigateway.*endpointType_PRIVATE,
//   	},
//   })
//
type EndpointType string

const (
	// For an edge-optimized API and its custom domain name.
	EndpointType_EDGE EndpointType = "EDGE"
	// For a regional API and its custom domain name.
	EndpointType_REGIONAL EndpointType = "REGIONAL"
	// For a private API and its custom domain name.
	EndpointType_PRIVATE EndpointType = "PRIVATE"
)

// Configure the response received by clients, produced from the API Gateway backend.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var responseType responseType
//   var restApi restApi
//
//   gatewayResponse := awscdk.Aws_apigateway.NewGatewayResponse(this, jsii.String("MyGatewayResponse"), &gatewayResponseProps{
//   	restApi: restApi,
//   	type: responseType,
//
//   	// the properties below are optional
//   	responseHeaders: map[string]*string{
//   		"responseHeadersKey": jsii.String("responseHeaders"),
//   	},
//   	statusCode: jsii.String("statusCode"),
//   	templates: map[string]*string{
//   		"templatesKey": jsii.String("templates"),
//   	},
//   })
//
type GatewayResponse interface {
	awscdk.Resource
	IGatewayResponse
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for GatewayResponse
type jsiiProxy_GatewayResponse struct {
	internal.Type__awscdkResource
	jsiiProxy_IGatewayResponse
}

func (j *jsiiProxy_GatewayResponse) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayResponse) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayResponse) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayResponse) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewGatewayResponse(scope constructs.Construct, id *string, props *GatewayResponseProps) GatewayResponse {
	_init_.Initialize()

	j := jsiiProxy_GatewayResponse{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.GatewayResponse",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewGatewayResponse_Override(g GatewayResponse, scope constructs.Construct, id *string, props *GatewayResponseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.GatewayResponse",
		[]interface{}{scope, id, props},
		g,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func GatewayResponse_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.GatewayResponse",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func GatewayResponse_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.GatewayResponse",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func GatewayResponse_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.GatewayResponse",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GatewayResponse) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		g,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (g *jsiiProxy_GatewayResponse) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GatewayResponse) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GatewayResponse) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GatewayResponse) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options to add gateway response.
//
// Example:
//   api := apigateway.NewRestApi(this, jsii.String("books-api"))
//   api.addGatewayResponse(jsii.String("test-response"), &gatewayResponseOptions{
//   	type: apigateway.responseType_ACCESS_DENIED(),
//   	statusCode: jsii.String("500"),
//   	responseHeaders: map[string]*string{
//   		"Access-Control-Allow-Origin": jsii.String("test.com"),
//   		"test-key": jsii.String("test-value"),
//   	},
//   	templates: map[string]*string{
//   		"application/json": jsii.String("{ \"message\": $context.error.messageString, \"statusCode\": \"488\", \"type\": \"$context.error.responseType\" }"),
//   	},
//   })
//
type GatewayResponseOptions struct {
	// Response type to associate with gateway response.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/supported-gateway-response-types.html
	//
	Type ResponseType `field:"required" json:"type" yaml:"type"`
	// Custom headers parameters for response.
	ResponseHeaders *map[string]*string `field:"optional" json:"responseHeaders" yaml:"responseHeaders"`
	// Http status code for response.
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
	// Custom templates to get mapped as response.
	Templates *map[string]*string `field:"optional" json:"templates" yaml:"templates"`
}

// Properties for a new gateway response.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var responseType responseType
//   var restApi restApi
//
//   gatewayResponseProps := &gatewayResponseProps{
//   	restApi: restApi,
//   	type: responseType,
//
//   	// the properties below are optional
//   	responseHeaders: map[string]*string{
//   		"responseHeadersKey": jsii.String("responseHeaders"),
//   	},
//   	statusCode: jsii.String("statusCode"),
//   	templates: map[string]*string{
//   		"templatesKey": jsii.String("templates"),
//   	},
//   }
//
type GatewayResponseProps struct {
	// Response type to associate with gateway response.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/supported-gateway-response-types.html
	//
	Type ResponseType `field:"required" json:"type" yaml:"type"`
	// Custom headers parameters for response.
	ResponseHeaders *map[string]*string `field:"optional" json:"responseHeaders" yaml:"responseHeaders"`
	// Http status code for response.
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
	// Custom templates to get mapped as response.
	Templates *map[string]*string `field:"optional" json:"templates" yaml:"templates"`
	// Rest api resource to target.
	RestApi IRestApi `field:"required" json:"restApi" yaml:"restApi"`
}

// You can integrate an API method with an HTTP endpoint using the HTTP proxy integration or the HTTP custom integration,.
//
// With the proxy integration, the setup is simple. You only need to set the
// HTTP method and the HTTP endpoint URI, according to the backend requirements,
// if you are not concerned with content encoding or caching.
//
// With the custom integration, the setup is more involved. In addition to the
// proxy integration setup steps, you need to specify how the incoming request
// data is mapped to the integration request and how the resulting integration
// response data is mapped to the method response.
//
// Example:
//   var authFn function
//   var books resource
//
//
//   auth := apigateway.NewRequestAuthorizer(this, jsii.String("booksAuthorizer"), &requestAuthorizerProps{
//   	handler: authFn,
//   	identitySources: []*string{
//   		apigateway.identitySource.header(jsii.String("Authorization")),
//   	},
//   })
//
//   books.addMethod(jsii.String("GET"), apigateway.NewHttpIntegration(jsii.String("http://amazon.com")), &methodOptions{
//   	authorizer: auth,
//   })
//
type HttpIntegration interface {
	Integration
	// Can be overridden by subclasses to allow the integration to interact with the method being integrated, access the REST API object, method ARNs, etc.
	Bind(_method Method) *IntegrationConfig
}

// The jsii proxy struct for HttpIntegration
type jsiiProxy_HttpIntegration struct {
	jsiiProxy_Integration
}

func NewHttpIntegration(url *string, props *HttpIntegrationProps) HttpIntegration {
	_init_.Initialize()

	j := jsiiProxy_HttpIntegration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.HttpIntegration",
		[]interface{}{url, props},
		&j,
	)

	return &j
}

func NewHttpIntegration_Override(h HttpIntegration, url *string, props *HttpIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.HttpIntegration",
		[]interface{}{url, props},
		h,
	)
}

func (h *jsiiProxy_HttpIntegration) Bind(_method Method) *IntegrationConfig {
	var returns *IntegrationConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{_method},
		&returns,
	)

	return returns
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//   var vpcLink vpcLink
//
//   httpIntegrationProps := &httpIntegrationProps{
//   	httpMethod: jsii.String("httpMethod"),
//   	options: &integrationOptions{
//   		cacheKeyParameters: []*string{
//   			jsii.String("cacheKeyParameters"),
//   		},
//   		cacheNamespace: jsii.String("cacheNamespace"),
//   		connectionType: awscdk.Aws_apigateway.connectionType_INTERNET,
//   		contentHandling: awscdk.*Aws_apigateway.contentHandling_CONVERT_TO_BINARY,
//   		credentialsPassthrough: jsii.Boolean(false),
//   		credentialsRole: role,
//   		integrationResponses: []integrationResponse{
//   			&integrationResponse{
//   				statusCode: jsii.String("statusCode"),
//
//   				// the properties below are optional
//   				contentHandling: awscdk.*Aws_apigateway.*contentHandling_CONVERT_TO_BINARY,
//   				responseParameters: map[string]*string{
//   					"responseParametersKey": jsii.String("responseParameters"),
//   				},
//   				responseTemplates: map[string]*string{
//   					"responseTemplatesKey": jsii.String("responseTemplates"),
//   				},
//   				selectionPattern: jsii.String("selectionPattern"),
//   			},
//   		},
//   		passthroughBehavior: awscdk.*Aws_apigateway.passthroughBehavior_WHEN_NO_MATCH,
//   		requestParameters: map[string]*string{
//   			"requestParametersKey": jsii.String("requestParameters"),
//   		},
//   		requestTemplates: map[string]*string{
//   			"requestTemplatesKey": jsii.String("requestTemplates"),
//   		},
//   		timeout: cdk.duration.minutes(jsii.Number(30)),
//   		vpcLink: vpcLink,
//   	},
//   	proxy: jsii.Boolean(false),
//   }
//
type HttpIntegrationProps struct {
	// HTTP method to use when invoking the backend URL.
	HttpMethod *string `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// Integration options, such as request/resopnse mapping, content handling, etc.
	Options *IntegrationOptions `field:"optional" json:"options" yaml:"options"`
	// Determines whether to use proxy integration or custom integration.
	Proxy *bool `field:"optional" json:"proxy" yaml:"proxy"`
}

// Access log destination for a RestApi Stage.
type IAccessLogDestination interface {
	// Binds this destination to the RestApi Stage.
	Bind(stage IStage) *AccessLogDestinationConfig
}

// The jsii proxy for IAccessLogDestination
type jsiiProxy_IAccessLogDestination struct {
	_ byte // padding
}

func (i *jsiiProxy_IAccessLogDestination) Bind(stage IStage) *AccessLogDestinationConfig {
	var returns *AccessLogDestinationConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{stage},
		&returns,
	)

	return returns
}

// API keys are alphanumeric string values that you distribute to app developer customers to grant access to your API.
type IApiKey interface {
	awscdk.IResource
	// The API key ARN.
	KeyArn() *string
	// The API key ID.
	KeyId() *string
}

// The jsii proxy for IApiKey
type jsiiProxy_IApiKey struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IApiKey) KeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiKey) KeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyId",
		&returns,
	)
	return returns
}

// Represents an API Gateway authorizer.
type IAuthorizer interface {
	// The authorization type of this authorizer.
	AuthorizationType() AuthorizationType
	// The authorizer ID.
	AuthorizerId() *string
}

// The jsii proxy for IAuthorizer
type jsiiProxy_IAuthorizer struct {
	_ byte // padding
}

func (j *jsiiProxy_IAuthorizer) AuthorizationType() AuthorizationType {
	var returns AuthorizationType
	_jsii_.Get(
		j,
		"authorizationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAuthorizer) AuthorizerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerId",
		&returns,
	)
	return returns
}

type IDomainName interface {
	awscdk.IResource
	// The domain name (e.g. `example.com`).
	DomainName() *string
	// The Route53 alias target to use in order to connect a record set to this domain through an alias.
	DomainNameAliasDomainName() *string
	// The Route53 hosted zone ID to use in order to connect a record set to this domain through an alias.
	DomainNameAliasHostedZoneId() *string
}

// The jsii proxy for IDomainName
type jsiiProxy_IDomainName struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IDomainName) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomainName) DomainNameAliasDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameAliasDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomainName) DomainNameAliasHostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameAliasHostedZoneId",
		&returns,
	)
	return returns
}

// Represents gateway response resource.
type IGatewayResponse interface {
	awscdk.IResource
}

// The jsii proxy for IGatewayResponse
type jsiiProxy_IGatewayResponse struct {
	internal.Type__awscdkIResource
}

type IModel interface {
	// Returns the model name, such as 'myModel'.
	ModelId() *string
}

// The jsii proxy for IModel
type jsiiProxy_IModel struct {
	_ byte // padding
}

func (j *jsiiProxy_IModel) ModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelId",
		&returns,
	)
	return returns
}

type IRequestValidator interface {
	awscdk.IResource
	// ID of the request validator, such as abc123.
	RequestValidatorId() *string
}

// The jsii proxy for IRequestValidator
type jsiiProxy_IRequestValidator struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IRequestValidator) RequestValidatorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestValidatorId",
		&returns,
	)
	return returns
}

type IResource interface {
	awscdk.IResource
	// Adds an OPTIONS method to this resource which responds to Cross-Origin Resource Sharing (CORS) preflight requests.
	//
	// Cross-Origin Resource Sharing (CORS) is a mechanism that uses additional
	// HTTP headers to tell browsers to give a web application running at one
	// origin, access to selected resources from a different origin. A web
	// application executes a cross-origin HTTP request when it requests a
	// resource that has a different origin (domain, protocol, or port) from its
	// own.
	//
	// Returns: a `Method` object.
	// See: https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS
	//
	AddCorsPreflight(options *CorsOptions) Method
	// Defines a new method for this resource.
	//
	// Returns: The newly created `Method` object.
	AddMethod(httpMethod *string, target Integration, options *MethodOptions) Method
	// Adds a greedy proxy resource ("{proxy+}") and an ANY method to this route.
	AddProxy(options *ProxyResourceOptions) ProxyResource
	// Defines a new child resource where this resource is the parent.
	//
	// Returns: A Resource object.
	AddResource(pathPart *string, options *ResourceOptions) Resource
	// Retrieves a child resource by path part.
	//
	// Returns: the child resource or undefined if not found.
	GetResource(pathPart *string) IResource
	// Gets or create all resources leading up to the specified path.
	//
	// - Path may only start with "/" if this method is called on the root resource.
	// - All resources are created using default options.
	//
	// Returns: a new or existing resource.
	ResourceForPath(path *string) Resource
	// The rest API that this resource is part of.
	//
	// The reason we need the RestApi object itself and not just the ID is because the model
	// is being tracked by the top-level RestApi object for the purpose of calculating it's
	// hash to determine the ID of the deployment. This allows us to automatically update
	// the deployment when the model of the REST API changes.
	Api() IRestApi
	// Default options for CORS preflight OPTIONS method.
	DefaultCorsPreflightOptions() *CorsOptions
	// An integration to use as a default for all methods created within this API unless an integration is specified.
	DefaultIntegration() Integration
	// Method options to use as a default for all methods created within this API unless custom options are specified.
	DefaultMethodOptions() *MethodOptions
	// The parent of this resource or undefined for the root resource.
	ParentResource() IResource
	// The full path of this resource.
	Path() *string
	// The ID of the resource.
	ResourceId() *string
}

// The jsii proxy for IResource
type jsiiProxy_IResource struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IResource) AddCorsPreflight(options *CorsOptions) Method {
	var returns Method

	_jsii_.Invoke(
		i,
		"addCorsPreflight",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IResource) AddMethod(httpMethod *string, target Integration, options *MethodOptions) Method {
	var returns Method

	_jsii_.Invoke(
		i,
		"addMethod",
		[]interface{}{httpMethod, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IResource) AddProxy(options *ProxyResourceOptions) ProxyResource {
	var returns ProxyResource

	_jsii_.Invoke(
		i,
		"addProxy",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IResource) AddResource(pathPart *string, options *ResourceOptions) Resource {
	var returns Resource

	_jsii_.Invoke(
		i,
		"addResource",
		[]interface{}{pathPart, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IResource) GetResource(pathPart *string) IResource {
	var returns IResource

	_jsii_.Invoke(
		i,
		"getResource",
		[]interface{}{pathPart},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IResource) ResourceForPath(path *string) Resource {
	var returns Resource

	_jsii_.Invoke(
		i,
		"resourceForPath",
		[]interface{}{path},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IResource) Api() IRestApi {
	var returns IRestApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResource) DefaultCorsPreflightOptions() *CorsOptions {
	var returns *CorsOptions
	_jsii_.Get(
		j,
		"defaultCorsPreflightOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResource) DefaultIntegration() Integration {
	var returns Integration
	_jsii_.Get(
		j,
		"defaultIntegration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResource) DefaultMethodOptions() *MethodOptions {
	var returns *MethodOptions
	_jsii_.Get(
		j,
		"defaultMethodOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResource) ParentResource() IResource {
	var returns IResource
	_jsii_.Get(
		j,
		"parentResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResource) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResource) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

type IRestApi interface {
	awscdk.IResource
	// Gets the "execute-api" ARN.
	//
	// Returns: The "execute-api" ARN.
	ArnForExecuteApi(method *string, path *string, stage *string) *string
	// API Gateway stage that points to the latest deployment (if defined).
	DeploymentStage() Stage
	SetDeploymentStage(d Stage)
	// API Gateway deployment that represents the latest changes of the API.
	//
	// This resource will be automatically updated every time the REST API model changes.
	// `undefined` when no deployment is configured.
	LatestDeployment() Deployment
	// The ID of this API Gateway RestApi.
	RestApiId() *string
	// The name of this API Gateway RestApi.
	RestApiName() *string
	// The resource ID of the root resource.
	RestApiRootResourceId() *string
	// Represents the root resource ("/") of this API. Use it to define the API model:.
	//
	// api.root.addMethod('ANY', redirectToHomePage); // "ANY /"
	//     api.root.addResource('friends').addMethod('GET', getFriendsHandler); // "GET /friends"
	Root() IResource
}

// The jsii proxy for IRestApi
type jsiiProxy_IRestApi struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IRestApi) ArnForExecuteApi(method *string, path *string, stage *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"arnForExecuteApi",
		[]interface{}{method, path, stage},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IRestApi) DeploymentStage() Stage {
	var returns Stage
	_jsii_.Get(
		j,
		"deploymentStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRestApi) SetDeploymentStage(val Stage) {
	_jsii_.Set(
		j,
		"deploymentStage",
		val,
	)
}

func (j *jsiiProxy_IRestApi) LatestDeployment() Deployment {
	var returns Deployment
	_jsii_.Get(
		j,
		"latestDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRestApi) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRestApi) RestApiName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRestApi) RestApiRootResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiRootResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRestApi) Root() IResource {
	var returns IResource
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

// Represents an APIGateway Stage.
type IStage interface {
	awscdk.IResource
	// RestApi to which this stage is associated.
	RestApi() IRestApi
	// Name of this stage.
	StageName() *string
}

// The jsii proxy for IStage
type jsiiProxy_IStage struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IStage) RestApi() IRestApi {
	var returns IRestApi
	_jsii_.Get(
		j,
		"restApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStage) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}

// A UsagePlan, either managed by this CDK app, or imported.
type IUsagePlan interface {
	awscdk.IResource
	// Adds an ApiKey.
	AddApiKey(apiKey IApiKey, options *AddApiKeyOptions)
	// Id of the usage plan.
	UsagePlanId() *string
}

// The jsii proxy for IUsagePlan
type jsiiProxy_IUsagePlan struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IUsagePlan) AddApiKey(apiKey IApiKey, options *AddApiKeyOptions) {
	_jsii_.InvokeVoid(
		i,
		"addApiKey",
		[]interface{}{apiKey, options},
	)
}

func (j *jsiiProxy_IUsagePlan) UsagePlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usagePlanId",
		&returns,
	)
	return returns
}

// Represents an API Gateway VpcLink.
type IVpcLink interface {
	awscdk.IResource
	// Physical ID of the VpcLink resource.
	VpcLinkId() *string
}

// The jsii proxy for IVpcLink
type jsiiProxy_IVpcLink struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IVpcLink) VpcLinkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcLinkId",
		&returns,
	)
	return returns
}

// Represents an identity source.
//
// The source can be specified either as a literal value (e.g: `Auth`) which
// cannot be blank, or as an unresolved string token.
//
// Example:
//   var authFn function
//   var books resource
//
//
//   auth := apigateway.NewRequestAuthorizer(this, jsii.String("booksAuthorizer"), &requestAuthorizerProps{
//   	handler: authFn,
//   	identitySources: []*string{
//   		apigateway.identitySource.header(jsii.String("Authorization")),
//   	},
//   })
//
//   books.addMethod(jsii.String("GET"), apigateway.NewHttpIntegration(jsii.String("http://amazon.com")), &methodOptions{
//   	authorizer: auth,
//   })
//
type IdentitySource interface {
}

// The jsii proxy struct for IdentitySource
type jsiiProxy_IdentitySource struct {
	_ byte // padding
}

func NewIdentitySource() IdentitySource {
	_init_.Initialize()

	j := jsiiProxy_IdentitySource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.IdentitySource",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewIdentitySource_Override(i IdentitySource) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.IdentitySource",
		nil, // no parameters
		i,
	)
}

// Provides a properly formatted request context identity source.
//
// Returns: a request context identity source.
func IdentitySource_Context(context *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.IdentitySource",
		"context",
		[]interface{}{context},
		&returns,
	)

	return returns
}

// Provides a properly formatted header identity source.
//
// Returns: a header identity source.
func IdentitySource_Header(headerName *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.IdentitySource",
		"header",
		[]interface{}{headerName},
		&returns,
	)

	return returns
}

// Provides a properly formatted query string identity source.
//
// Returns: a query string identity source.
func IdentitySource_QueryString(queryString *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.IdentitySource",
		"queryString",
		[]interface{}{queryString},
		&returns,
	)

	return returns
}

// Provides a properly formatted API Gateway stage variable identity source.
//
// Returns: an API Gateway stage variable identity source.
func IdentitySource_StageVariable(stageVariable *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.IdentitySource",
		"stageVariable",
		[]interface{}{stageVariable},
		&returns,
	)

	return returns
}

// OpenAPI specification from an inline JSON object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var definition interface{}
//
//   inlineApiDefinition := awscdk.Aws_apigateway.NewInlineApiDefinition(definition)
//
type InlineApiDefinition interface {
	ApiDefinition
	// Called when the specification is initialized to allow this object to bind to the stack, add resources and have fun.
	Bind(_scope constructs.Construct) *ApiDefinitionConfig
	// Called after the CFN RestApi resource has been created to allow the Api Definition to bind to it.
	//
	// Specifically it's required to allow assets to add
	// metadata for tooling like SAM CLI to be able to find their origins.
	BindAfterCreate(_scope constructs.Construct, _restApi IRestApi)
}

// The jsii proxy struct for InlineApiDefinition
type jsiiProxy_InlineApiDefinition struct {
	jsiiProxy_ApiDefinition
}

func NewInlineApiDefinition(definition interface{}) InlineApiDefinition {
	_init_.Initialize()

	j := jsiiProxy_InlineApiDefinition{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.InlineApiDefinition",
		[]interface{}{definition},
		&j,
	)

	return &j
}

func NewInlineApiDefinition_Override(i InlineApiDefinition, definition interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.InlineApiDefinition",
		[]interface{}{definition},
		i,
	)
}

// Loads the API specification from a local disk asset.
func InlineApiDefinition_FromAsset(file *string, options *awss3assets.AssetOptions) AssetApiDefinition {
	_init_.Initialize()

	var returns AssetApiDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.InlineApiDefinition",
		"fromAsset",
		[]interface{}{file, options},
		&returns,
	)

	return returns
}

// Creates an API definition from a specification file in an S3 bucket.
func InlineApiDefinition_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3ApiDefinition {
	_init_.Initialize()

	var returns S3ApiDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.InlineApiDefinition",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Create an API definition from an inline object.
//
// The inline object must follow the
// schema of OpenAPI 2.0 or OpenAPI 3.0
//
// Example:
//     apigateway.ApiDefinition.fromInline({
//       openapi: '3.0.2',
//       paths: {
//         '/pets': {
//           get: {
//             'responses': {
//               200: {
//                 content: {
//                   'application/json': {
//                     schema: {
//                       $ref: '#/components/schemas/Empty',
//                     },
//                   },
//                 },
//               },
//             },
//             'x-amazon-apigateway-integration': {
//               responses: {
//                 default: {
//                   statusCode: '200',
//                 },
//               },
//               requestTemplates: {
//                 'application/json': '{"statusCode": 200}',
//               },
//               passthroughBehavior: 'when_no_match',
//               type: 'mock',
//             },
//           },
//         },
//       },
//       components: {
//         schemas: {
//           Empty: {
//             title: 'Empty Schema',
//             type: 'object',
//           },
//         },
//       },
//     });
//
func InlineApiDefinition_FromInline(definition interface{}) InlineApiDefinition {
	_init_.Initialize()

	var returns InlineApiDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.InlineApiDefinition",
		"fromInline",
		[]interface{}{definition},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InlineApiDefinition) Bind(_scope constructs.Construct) *ApiDefinitionConfig {
	var returns *ApiDefinitionConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InlineApiDefinition) BindAfterCreate(_scope constructs.Construct, _restApi IRestApi) {
	_jsii_.InvokeVoid(
		i,
		"bindAfterCreate",
		[]interface{}{_scope, _restApi},
	)
}

// Base class for backend integrations for an API Gateway method.
//
// Use one of the concrete classes such as `MockIntegration`, `AwsIntegration`, `LambdaIntegration`
// or implement on your own by specifying the set of props.
//
// Example:
//   var books resource
//   var iamUser user
//
//
//   getBooks := books.addMethod(jsii.String("GET"), apigateway.NewHttpIntegration(jsii.String("http://amazon.com")), &methodOptions{
//   	authorizationType: apigateway.authorizationType_IAM,
//   })
//
//   iamUser.attachInlinePolicy(iam.NewPolicy(this, jsii.String("AllowBooks"), &policyProps{
//   	statements: []policyStatement{
//   		iam.NewPolicyStatement(&policyStatementProps{
//   			actions: []*string{
//   				jsii.String("execute-api:Invoke"),
//   			},
//   			effect: iam.effect_ALLOW,
//   			resources: []*string{
//   				getBooks.methodArn,
//   			},
//   		}),
//   	},
//   }))
//
type Integration interface {
	// Can be overridden by subclasses to allow the integration to interact with the method being integrated, access the REST API object, method ARNs, etc.
	Bind(_method Method) *IntegrationConfig
}

// The jsii proxy struct for Integration
type jsiiProxy_Integration struct {
	_ byte // padding
}

func NewIntegration(props *IntegrationProps) Integration {
	_init_.Initialize()

	j := jsiiProxy_Integration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.Integration",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewIntegration_Override(i Integration, props *IntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.Integration",
		[]interface{}{props},
		i,
	)
}

func (i *jsiiProxy_Integration) Bind(_method Method) *IntegrationConfig {
	var returns *IntegrationConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{_method},
		&returns,
	)

	return returns
}

// Result of binding an Integration to a Method.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//   var vpcLink vpcLink
//
//   integrationConfig := &integrationConfig{
//   	type: awscdk.Aws_apigateway.integrationType_AWS,
//
//   	// the properties below are optional
//   	deploymentToken: jsii.String("deploymentToken"),
//   	integrationHttpMethod: jsii.String("integrationHttpMethod"),
//   	options: &integrationOptions{
//   		cacheKeyParameters: []*string{
//   			jsii.String("cacheKeyParameters"),
//   		},
//   		cacheNamespace: jsii.String("cacheNamespace"),
//   		connectionType: awscdk.*Aws_apigateway.connectionType_INTERNET,
//   		contentHandling: awscdk.*Aws_apigateway.contentHandling_CONVERT_TO_BINARY,
//   		credentialsPassthrough: jsii.Boolean(false),
//   		credentialsRole: role,
//   		integrationResponses: []integrationResponse{
//   			&integrationResponse{
//   				statusCode: jsii.String("statusCode"),
//
//   				// the properties below are optional
//   				contentHandling: awscdk.*Aws_apigateway.*contentHandling_CONVERT_TO_BINARY,
//   				responseParameters: map[string]*string{
//   					"responseParametersKey": jsii.String("responseParameters"),
//   				},
//   				responseTemplates: map[string]*string{
//   					"responseTemplatesKey": jsii.String("responseTemplates"),
//   				},
//   				selectionPattern: jsii.String("selectionPattern"),
//   			},
//   		},
//   		passthroughBehavior: awscdk.*Aws_apigateway.passthroughBehavior_WHEN_NO_MATCH,
//   		requestParameters: map[string]*string{
//   			"requestParametersKey": jsii.String("requestParameters"),
//   		},
//   		requestTemplates: map[string]*string{
//   			"requestTemplatesKey": jsii.String("requestTemplates"),
//   		},
//   		timeout: cdk.duration.minutes(jsii.Number(30)),
//   		vpcLink: vpcLink,
//   	},
//   	uri: jsii.String("uri"),
//   }
//
type IntegrationConfig struct {
	// Specifies an API method integration type.
	Type IntegrationType `field:"required" json:"type" yaml:"type"`
	// This value is included in computing the Deployment's fingerprint.
	//
	// When the fingerprint
	// changes, a new deployment is triggered.
	// This property should contain values associated with the Integration that upon changing
	// should trigger a fresh the Deployment needs to be refreshed.
	DeploymentToken *string `field:"optional" json:"deploymentToken" yaml:"deploymentToken"`
	// The integration's HTTP method type.
	IntegrationHttpMethod *string `field:"optional" json:"integrationHttpMethod" yaml:"integrationHttpMethod"`
	// Integration options.
	Options *IntegrationOptions `field:"optional" json:"options" yaml:"options"`
	// The Uniform Resource Identifier (URI) for the integration.
	// See: https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/#uri
	//
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

// Example:
//   import path "github.com/aws-samples/dummy/path"
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   // Against the RestApi endpoint from the stack output, run
//   // `curl -s -o /dev/null -w "%{http_code}" <url>` should return 401
//   // `curl -s -o /dev/null -w "%{http_code}" -H 'Authorization: deny' <url>?allow=yes` should return 403
//   // `curl -s -o /dev/null -w "%{http_code}" -H 'Authorization: allow' <url>?allow=yes` should return 200
//
//   app := awscdk.NewApp()
//   stack := awscdk.NewStack(app, jsii.String("RequestAuthorizerInteg"))
//
//   authorizerFn := lambda.NewFunction(stack, jsii.String("MyAuthorizerFunction"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.assetCode.fromAsset(path.join(__dirname, jsii.String("integ.request-authorizer.handler"))),
//   })
//
//   restapi := awscdk.NewRestApi(stack, jsii.String("MyRestApi"))
//
//   authorizer := awscdk.NewRequestAuthorizer(stack, jsii.String("MyAuthorizer"), &requestAuthorizerProps{
//   	handler: authorizerFn,
//   	identitySources: []*string{
//   		awscdk.IdentitySource.header(jsii.String("Authorization")),
//   		awscdk.IdentitySource.queryString(jsii.String("allow")),
//   	},
//   })
//
//   restapi.root.addMethod(jsii.String("ANY"), awscdk.NewMockIntegration(&integrationOptions{
//   	integrationResponses: []integrationResponse{
//   		&integrationResponse{
//   			statusCode: jsii.String("200"),
//   		},
//   	},
//   	passthroughBehavior: awscdk.PassthroughBehavior_NEVER,
//   	requestTemplates: map[string]*string{
//   		"application/json": jsii.String("{ \"statusCode\": 200 }"),
//   	},
//   }), &methodOptions{
//   	methodResponses: []methodResponse{
//   		&methodResponse{
//   			statusCode: jsii.String("200"),
//   		},
//   	},
//   	authorizer: authorizer,
//   })
//
type IntegrationOptions struct {
	// A list of request parameters whose values are to be cached.
	//
	// It determines
	// request parameters that will make it into the cache key.
	CacheKeyParameters *[]*string `field:"optional" json:"cacheKeyParameters" yaml:"cacheKeyParameters"`
	// An API-specific tag group of related cached parameters.
	CacheNamespace *string `field:"optional" json:"cacheNamespace" yaml:"cacheNamespace"`
	// The type of network connection to the integration endpoint.
	ConnectionType ConnectionType `field:"optional" json:"connectionType" yaml:"connectionType"`
	// Specifies how to handle request payload content type conversions.
	ContentHandling ContentHandling `field:"optional" json:"contentHandling" yaml:"contentHandling"`
	// Requires that the caller's identity be passed through from the request.
	CredentialsPassthrough *bool `field:"optional" json:"credentialsPassthrough" yaml:"credentialsPassthrough"`
	// An IAM role that API Gateway assumes.
	//
	// Mutually exclusive with `credentialsPassThrough`.
	CredentialsRole awsiam.IRole `field:"optional" json:"credentialsRole" yaml:"credentialsRole"`
	// The response that API Gateway provides after a method's backend completes processing a request.
	//
	// API Gateway intercepts the response from the
	// backend so that you can control how API Gateway surfaces backend
	// responses. For example, you can map the backend status codes to codes
	// that you define.
	IntegrationResponses *[]*IntegrationResponse `field:"optional" json:"integrationResponses" yaml:"integrationResponses"`
	// Specifies the pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the requestTemplates property on the Integration resource.
	//
	// There are three valid values: WHEN_NO_MATCH, WHEN_NO_TEMPLATES, and
	// NEVER.
	PassthroughBehavior PassthroughBehavior `field:"optional" json:"passthroughBehavior" yaml:"passthroughBehavior"`
	// The request parameters that API Gateway sends with the backend request.
	//
	// Specify request parameters as key-value pairs (string-to-string
	// mappings), with a destination as the key and a source as the value.
	//
	// Specify the destination by using the following pattern
	// integration.request.location.name, where location is querystring, path,
	// or header, and name is a valid, unique parameter name.
	//
	// The source must be an existing method request parameter or a static
	// value. You must enclose static values in single quotation marks and
	// pre-encode these values based on their destination in the request.
	RequestParameters *map[string]*string `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// A map of Apache Velocity templates that are applied on the request payload.
	//
	// The template that API Gateway uses is based on the value of the
	// Content-Type header that's sent by the client. The content type value is
	// the key, and the template is the value (specified as a string), such as
	// the following snippet:
	//
	// ```
	//    { "application/json": "{ \"statusCode\": 200 }" }
	// ```.
	// See: http://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html
	//
	RequestTemplates *map[string]*string `field:"optional" json:"requestTemplates" yaml:"requestTemplates"`
	// The maximum amount of time an integration will run before it returns without a response.
	//
	// Must be between 50 milliseconds and 29 seconds.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The VpcLink used for the integration.
	//
	// Required if connectionType is VPC_LINK.
	VpcLink IVpcLink `field:"optional" json:"vpcLink" yaml:"vpcLink"`
}

// Example:
//   import elbv2 "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("VPC"))
//   nlb := elbv2.NewNetworkLoadBalancer(this, jsii.String("NLB"), &networkLoadBalancerProps{
//   	vpc: vpc,
//   })
//   link := apigateway.NewVpcLink(this, jsii.String("link"), &vpcLinkProps{
//   	targets: []iNetworkLoadBalancer{
//   		nlb,
//   	},
//   })
//
//   integration := apigateway.NewIntegration(&integrationProps{
//   	type: apigateway.integrationType_HTTP_PROXY,
//   	options: &integrationOptions{
//   		connectionType: apigateway.connectionType_VPC_LINK,
//   		vpcLink: link,
//   	},
//   })
//
type IntegrationProps struct {
	// Specifies an API method integration type.
	Type IntegrationType `field:"required" json:"type" yaml:"type"`
	// The integration's HTTP method type.
	//
	// Required unless you use a MOCK integration.
	IntegrationHttpMethod *string `field:"optional" json:"integrationHttpMethod" yaml:"integrationHttpMethod"`
	// Integration options.
	Options *IntegrationOptions `field:"optional" json:"options" yaml:"options"`
	// The Uniform Resource Identifier (URI) for the integration.
	//
	// - If you specify HTTP for the `type` property, specify the API endpoint URL.
	// - If you specify MOCK for the `type` property, don't specify this property.
	// - If you specify AWS for the `type` property, specify an AWS service that
	//    follows this form: `arn:partition:apigateway:region:subdomain.service|service:path|action/service_api.`
	//    For example, a Lambda function URI follows this form:
	//    arn:partition:apigateway:region:lambda:path/path. The path is usually in the
	//    form /2015-03-31/functions/LambdaFunctionARN/invocations.
	// See: https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/#uri
	//
	Uri interface{} `field:"optional" json:"uri" yaml:"uri"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integrationResponse := &integrationResponse{
//   	statusCode: jsii.String("statusCode"),
//
//   	// the properties below are optional
//   	contentHandling: awscdk.Aws_apigateway.contentHandling_CONVERT_TO_BINARY,
//   	responseParameters: map[string]*string{
//   		"responseParametersKey": jsii.String("responseParameters"),
//   	},
//   	responseTemplates: map[string]*string{
//   		"responseTemplatesKey": jsii.String("responseTemplates"),
//   	},
//   	selectionPattern: jsii.String("selectionPattern"),
//   }
//
type IntegrationResponse struct {
	// The status code that API Gateway uses to map the integration response to a MethodResponse status code.
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
	// Specifies how to handle request payload content type conversions.
	ContentHandling ContentHandling `field:"optional" json:"contentHandling" yaml:"contentHandling"`
	// The response parameters from the backend response that API Gateway sends to the method response.
	//
	// Use the destination as the key and the source as the value:
	//
	// - The destination must be an existing response parameter in the
	//    MethodResponse property.
	// - The source must be an existing method request parameter or a static
	//    value. You must enclose static values in single quotation marks and
	//    pre-encode these values based on the destination specified in the
	//    request.
	// See: http://docs.aws.amazon.com/apigateway/latest/developerguide/request-response-data-mappings.html
	//
	ResponseParameters *map[string]*string `field:"optional" json:"responseParameters" yaml:"responseParameters"`
	// The templates that are used to transform the integration response body.
	//
	// Specify templates as key-value pairs, with a content type as the key and
	// a template as the value.
	// See: http://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html
	//
	ResponseTemplates *map[string]*string `field:"optional" json:"responseTemplates" yaml:"responseTemplates"`
	// Specifies the regular expression (regex) pattern used to choose an integration response based on the response from the back end.
	//
	// For example, if the success response returns nothing and the error response returns some string, you
	// could use the ``.+`` regex to match error response. However, make sure that the error response does not contain any
	// newline (``\n``) character in such cases. If the back end is an AWS Lambda function, the AWS Lambda function error
	// header is matched. For all other HTTP and AWS back ends, the HTTP status code is matched.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-integration-settings-integration-response.html
	//
	SelectionPattern *string `field:"optional" json:"selectionPattern" yaml:"selectionPattern"`
}

// Example:
//   import elbv2 "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("VPC"))
//   nlb := elbv2.NewNetworkLoadBalancer(this, jsii.String("NLB"), &networkLoadBalancerProps{
//   	vpc: vpc,
//   })
//   link := apigateway.NewVpcLink(this, jsii.String("link"), &vpcLinkProps{
//   	targets: []iNetworkLoadBalancer{
//   		nlb,
//   	},
//   })
//
//   integration := apigateway.NewIntegration(&integrationProps{
//   	type: apigateway.integrationType_HTTP_PROXY,
//   	options: &integrationOptions{
//   		connectionType: apigateway.connectionType_VPC_LINK,
//   		vpcLink: link,
//   	},
//   })
//
type IntegrationType string

const (
	// For integrating the API method request with an AWS service action, including the Lambda function-invoking action.
	//
	// With the Lambda
	// function-invoking action, this is referred to as the Lambda custom
	// integration. With any other AWS service action, this is known as AWS
	// integration.
	IntegrationType_AWS IntegrationType = "AWS"
	// For integrating the API method request with the Lambda function-invoking action with the client request passed through as-is.
	//
	// This integration is
	// also referred to as the Lambda proxy integration.
	IntegrationType_AWS_PROXY IntegrationType = "AWS_PROXY"
	// For integrating the API method request with an HTTP endpoint, including a private HTTP endpoint within a VPC.
	//
	// This integration is also referred to
	// as the HTTP custom integration.
	IntegrationType_HTTP IntegrationType = "HTTP"
	// For integrating the API method request with an HTTP endpoint, including a private HTTP endpoint within a VPC, with the client request passed through as-is.
	//
	// This is also referred to as the HTTP proxy integration.
	IntegrationType_HTTP_PROXY IntegrationType = "HTTP_PROXY"
	// For integrating the API method request with API Gateway as a "loop-back" endpoint without invoking any backend.
	IntegrationType_MOCK IntegrationType = "MOCK"
)

// Represents a JSON schema definition of the structure of a REST API model.
//
// Copied from npm module jsonschema.
//
// Example:
//   var api restApi
//
//
//   // We define the JSON Schema for the transformed valid response
//   responseModel := api.addModel(jsii.String("ResponseModel"), &modelOptions{
//   	contentType: jsii.String("application/json"),
//   	modelName: jsii.String("ResponseModel"),
//   	schema: &jsonSchema{
//   		schema: apigateway.jsonSchemaVersion_DRAFT4,
//   		title: jsii.String("pollResponse"),
//   		type: apigateway.jsonSchemaType_OBJECT,
//   		properties: map[string]*jsonSchema{
//   			"state": &jsonSchema{
//   				"type": apigateway.*jsonSchemaType_STRING,
//   			},
//   			"greeting": &jsonSchema{
//   				"type": apigateway.*jsonSchemaType_STRING,
//   			},
//   		},
//   	},
//   })
//
//   // We define the JSON Schema for the transformed error response
//   errorResponseModel := api.addModel(jsii.String("ErrorResponseModel"), &modelOptions{
//   	contentType: jsii.String("application/json"),
//   	modelName: jsii.String("ErrorResponseModel"),
//   	schema: &jsonSchema{
//   		schema: apigateway.*jsonSchemaVersion_DRAFT4,
//   		title: jsii.String("errorResponse"),
//   		type: apigateway.*jsonSchemaType_OBJECT,
//   		properties: map[string]*jsonSchema{
//   			"state": &jsonSchema{
//   				"type": apigateway.*jsonSchemaType_STRING,
//   			},
//   			"message": &jsonSchema{
//   				"type": apigateway.*jsonSchemaType_STRING,
//   			},
//   		},
//   	},
//   })
//
// See: https://github.com/tdegrunt/jsonschema
//
type JsonSchema struct {
	AdditionalItems *[]*JsonSchema `field:"optional" json:"additionalItems" yaml:"additionalItems"`
	AdditionalProperties interface{} `field:"optional" json:"additionalProperties" yaml:"additionalProperties"`
	AllOf *[]*JsonSchema `field:"optional" json:"allOf" yaml:"allOf"`
	AnyOf *[]*JsonSchema `field:"optional" json:"anyOf" yaml:"anyOf"`
	Contains interface{} `field:"optional" json:"contains" yaml:"contains"`
	// The default value if you use an enum.
	Default interface{} `field:"optional" json:"default" yaml:"default"`
	Definitions *map[string]*JsonSchema `field:"optional" json:"definitions" yaml:"definitions"`
	Dependencies *map[string]interface{} `field:"optional" json:"dependencies" yaml:"dependencies"`
	Description *string `field:"optional" json:"description" yaml:"description"`
	Enum *[]interface{} `field:"optional" json:"enum" yaml:"enum"`
	ExclusiveMaximum *bool `field:"optional" json:"exclusiveMaximum" yaml:"exclusiveMaximum"`
	ExclusiveMinimum *bool `field:"optional" json:"exclusiveMinimum" yaml:"exclusiveMinimum"`
	Format *string `field:"optional" json:"format" yaml:"format"`
	Id *string `field:"optional" json:"id" yaml:"id"`
	Items interface{} `field:"optional" json:"items" yaml:"items"`
	Maximum *float64 `field:"optional" json:"maximum" yaml:"maximum"`
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	MaxLength *float64 `field:"optional" json:"maxLength" yaml:"maxLength"`
	MaxProperties *float64 `field:"optional" json:"maxProperties" yaml:"maxProperties"`
	Minimum *float64 `field:"optional" json:"minimum" yaml:"minimum"`
	MinItems *float64 `field:"optional" json:"minItems" yaml:"minItems"`
	MinLength *float64 `field:"optional" json:"minLength" yaml:"minLength"`
	MinProperties *float64 `field:"optional" json:"minProperties" yaml:"minProperties"`
	MultipleOf *float64 `field:"optional" json:"multipleOf" yaml:"multipleOf"`
	Not **JsonSchema `field:"optional" json:"not" yaml:"not"`
	OneOf *[]*JsonSchema `field:"optional" json:"oneOf" yaml:"oneOf"`
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	PatternProperties *map[string]*JsonSchema `field:"optional" json:"patternProperties" yaml:"patternProperties"`
	Properties *map[string]*JsonSchema `field:"optional" json:"properties" yaml:"properties"`
	PropertyNames **JsonSchema `field:"optional" json:"propertyNames" yaml:"propertyNames"`
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
	Required *[]*string `field:"optional" json:"required" yaml:"required"`
	Schema JsonSchemaVersion `field:"optional" json:"schema" yaml:"schema"`
	Title *string `field:"optional" json:"title" yaml:"title"`
	Type interface{} `field:"optional" json:"type" yaml:"type"`
	UniqueItems *bool `field:"optional" json:"uniqueItems" yaml:"uniqueItems"`
}

// Example:
//   var api restApi
//
//
//   // We define the JSON Schema for the transformed valid response
//   responseModel := api.addModel(jsii.String("ResponseModel"), &modelOptions{
//   	contentType: jsii.String("application/json"),
//   	modelName: jsii.String("ResponseModel"),
//   	schema: &jsonSchema{
//   		schema: apigateway.jsonSchemaVersion_DRAFT4,
//   		title: jsii.String("pollResponse"),
//   		type: apigateway.jsonSchemaType_OBJECT,
//   		properties: map[string]*jsonSchema{
//   			"state": &jsonSchema{
//   				"type": apigateway.*jsonSchemaType_STRING,
//   			},
//   			"greeting": &jsonSchema{
//   				"type": apigateway.*jsonSchemaType_STRING,
//   			},
//   		},
//   	},
//   })
//
//   // We define the JSON Schema for the transformed error response
//   errorResponseModel := api.addModel(jsii.String("ErrorResponseModel"), &modelOptions{
//   	contentType: jsii.String("application/json"),
//   	modelName: jsii.String("ErrorResponseModel"),
//   	schema: &jsonSchema{
//   		schema: apigateway.*jsonSchemaVersion_DRAFT4,
//   		title: jsii.String("errorResponse"),
//   		type: apigateway.*jsonSchemaType_OBJECT,
//   		properties: map[string]*jsonSchema{
//   			"state": &jsonSchema{
//   				"type": apigateway.*jsonSchemaType_STRING,
//   			},
//   			"message": &jsonSchema{
//   				"type": apigateway.*jsonSchemaType_STRING,
//   			},
//   		},
//   	},
//   })
//
type JsonSchemaType string

const (
	JsonSchemaType_NULL JsonSchemaType = "NULL"
	JsonSchemaType_BOOLEAN JsonSchemaType = "BOOLEAN"
	JsonSchemaType_OBJECT JsonSchemaType = "OBJECT"
	JsonSchemaType_ARRAY JsonSchemaType = "ARRAY"
	JsonSchemaType_NUMBER JsonSchemaType = "NUMBER"
	JsonSchemaType_INTEGER JsonSchemaType = "INTEGER"
	JsonSchemaType_STRING JsonSchemaType = "STRING"
)

// Example:
//   var api restApi
//
//
//   // We define the JSON Schema for the transformed valid response
//   responseModel := api.addModel(jsii.String("ResponseModel"), &modelOptions{
//   	contentType: jsii.String("application/json"),
//   	modelName: jsii.String("ResponseModel"),
//   	schema: &jsonSchema{
//   		schema: apigateway.jsonSchemaVersion_DRAFT4,
//   		title: jsii.String("pollResponse"),
//   		type: apigateway.jsonSchemaType_OBJECT,
//   		properties: map[string]*jsonSchema{
//   			"state": &jsonSchema{
//   				"type": apigateway.*jsonSchemaType_STRING,
//   			},
//   			"greeting": &jsonSchema{
//   				"type": apigateway.*jsonSchemaType_STRING,
//   			},
//   		},
//   	},
//   })
//
//   // We define the JSON Schema for the transformed error response
//   errorResponseModel := api.addModel(jsii.String("ErrorResponseModel"), &modelOptions{
//   	contentType: jsii.String("application/json"),
//   	modelName: jsii.String("ErrorResponseModel"),
//   	schema: &jsonSchema{
//   		schema: apigateway.*jsonSchemaVersion_DRAFT4,
//   		title: jsii.String("errorResponse"),
//   		type: apigateway.*jsonSchemaType_OBJECT,
//   		properties: map[string]*jsonSchema{
//   			"state": &jsonSchema{
//   				"type": apigateway.*jsonSchemaType_STRING,
//   			},
//   			"message": &jsonSchema{
//   				"type": apigateway.*jsonSchemaType_STRING,
//   			},
//   		},
//   	},
//   })
//
type JsonSchemaVersion string

const (
	// In API Gateway models are defined using the JSON schema draft 4.
	// See: https://tools.ietf.org/html/draft-zyp-json-schema-04
	//
	JsonSchemaVersion_DRAFT4 JsonSchemaVersion = "DRAFT4"
	JsonSchemaVersion_DRAFT7 JsonSchemaVersion = "DRAFT7"
)

// Properties for controlling items output in JSON standard format.
//
// Example:
//   // production stage
//   prdLogGroup := logs.NewLogGroup(this, jsii.String("PrdLogs"))
//   api := apigateway.NewRestApi(this, jsii.String("books"), &restApiProps{
//   	deployOptions: &stageOptions{
//   		accessLogDestination: apigateway.NewLogGroupLogDestination(prdLogGroup),
//   		accessLogFormat: apigateway.accessLogFormat.jsonWithStandardFields(),
//   	},
//   })
//   deployment := apigateway.NewDeployment(this, jsii.String("Deployment"), &deploymentProps{
//   	api: api,
//   })
//
//   // development stage
//   devLogGroup := logs.NewLogGroup(this, jsii.String("DevLogs"))
//   apigateway.NewStage(this, jsii.String("dev"), &stageProps{
//   	deployment: deployment,
//   	accessLogDestination: apigateway.NewLogGroupLogDestination(devLogGroup),
//   	accessLogFormat: apigateway.*accessLogFormat.jsonWithStandardFields(&jsonWithStandardFieldProps{
//   		caller: jsii.Boolean(false),
//   		httpMethod: jsii.Boolean(true),
//   		ip: jsii.Boolean(true),
//   		protocol: jsii.Boolean(true),
//   		requestTime: jsii.Boolean(true),
//   		resourcePath: jsii.Boolean(true),
//   		responseLength: jsii.Boolean(true),
//   		status: jsii.Boolean(true),
//   		user: jsii.Boolean(true),
//   	}),
//   })
//
type JsonWithStandardFieldProps struct {
	// If this flag is enabled, the principal identifier of the caller will be output to the log.
	Caller *bool `field:"required" json:"caller" yaml:"caller"`
	// If this flag is enabled, the http method will be output to the log.
	HttpMethod *bool `field:"required" json:"httpMethod" yaml:"httpMethod"`
	// If this flag is enabled, the source IP of request will be output to the log.
	Ip *bool `field:"required" json:"ip" yaml:"ip"`
	// If this flag is enabled, the request protocol will be output to the log.
	Protocol *bool `field:"required" json:"protocol" yaml:"protocol"`
	// If this flag is enabled, the CLF-formatted request time((dd/MMM/yyyy:HH:mm:ss +-hhmm) will be output to the log.
	RequestTime *bool `field:"required" json:"requestTime" yaml:"requestTime"`
	// If this flag is enabled, the path to your resource will be output to the log.
	ResourcePath *bool `field:"required" json:"resourcePath" yaml:"resourcePath"`
	// If this flag is enabled, the response payload length will be output to the log.
	ResponseLength *bool `field:"required" json:"responseLength" yaml:"responseLength"`
	// If this flag is enabled, the method response status will be output to the log.
	Status *bool `field:"required" json:"status" yaml:"status"`
	// If this flag is enabled, the principal identifier of the user will be output to the log.
	User *bool `field:"required" json:"user" yaml:"user"`
}

// Base properties for all lambda authorizers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var function_ function
//   var role role
//
//   lambdaAuthorizerProps := &lambdaAuthorizerProps{
//   	handler: function_,
//
//   	// the properties below are optional
//   	assumeRole: role,
//   	authorizerName: jsii.String("authorizerName"),
//   	resultsCacheTtl: cdk.duration.minutes(jsii.Number(30)),
//   }
//
type LambdaAuthorizerProps struct {
	// The handler for the authorizer lambda function.
	//
	// The handler must follow a very specific protocol on the input it receives and the output it needs to produce.
	// API Gateway has documented the handler's input specification
	// {@link https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-lambda-authorizer-input.html | here} and output specification
	// {@link https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-lambda-authorizer-output.html | here}.
	Handler awslambda.IFunction `field:"required" json:"handler" yaml:"handler"`
	// An optional IAM role for APIGateway to assume before calling the Lambda-based authorizer.
	//
	// The IAM role must be
	// assumable by 'apigateway.amazonaws.com'.
	AssumeRole awsiam.IRole `field:"optional" json:"assumeRole" yaml:"assumeRole"`
	// An optional human friendly name for the authorizer.
	//
	// Note that, this is not the primary identifier of the authorizer.
	AuthorizerName *string `field:"optional" json:"authorizerName" yaml:"authorizerName"`
	// How long APIGateway should cache the results.
	//
	// Max 1 hour.
	// Disable caching by setting this to 0.
	ResultsCacheTtl awscdk.Duration `field:"optional" json:"resultsCacheTtl" yaml:"resultsCacheTtl"`
}

// Integrates an AWS Lambda function to an API Gateway method.
//
// Example:
//   var resource resource
//   var handler function
//
//   resource.addMethod(jsii.String("GET"), apigateway.NewLambdaIntegration(handler))
//
type LambdaIntegration interface {
	AwsIntegration
	// Can be overridden by subclasses to allow the integration to interact with the method being integrated, access the REST API object, method ARNs, etc.
	Bind(method Method) *IntegrationConfig
}

// The jsii proxy struct for LambdaIntegration
type jsiiProxy_LambdaIntegration struct {
	jsiiProxy_AwsIntegration
}

func NewLambdaIntegration(handler awslambda.IFunction, options *LambdaIntegrationOptions) LambdaIntegration {
	_init_.Initialize()

	j := jsiiProxy_LambdaIntegration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.LambdaIntegration",
		[]interface{}{handler, options},
		&j,
	)

	return &j
}

func NewLambdaIntegration_Override(l LambdaIntegration, handler awslambda.IFunction, options *LambdaIntegrationOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.LambdaIntegration",
		[]interface{}{handler, options},
		l,
	)
}

func (l *jsiiProxy_LambdaIntegration) Bind(method Method) *IntegrationConfig {
	var returns *IntegrationConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{method},
		&returns,
	)

	return returns
}

// Example:
//   var backend function
//
//
//   api := apigateway.NewLambdaRestApi(this, jsii.String("myapi"), &lambdaRestApiProps{
//   	handler: backend,
//   	integrationOptions: &lambdaIntegrationOptions{
//   		allowTestInvoke: jsii.Boolean(false),
//   		timeout: awscdk.Duration.seconds(jsii.Number(1)),
//   	},
//   })
//
type LambdaIntegrationOptions struct {
	// A list of request parameters whose values are to be cached.
	//
	// It determines
	// request parameters that will make it into the cache key.
	CacheKeyParameters *[]*string `field:"optional" json:"cacheKeyParameters" yaml:"cacheKeyParameters"`
	// An API-specific tag group of related cached parameters.
	CacheNamespace *string `field:"optional" json:"cacheNamespace" yaml:"cacheNamespace"`
	// The type of network connection to the integration endpoint.
	ConnectionType ConnectionType `field:"optional" json:"connectionType" yaml:"connectionType"`
	// Specifies how to handle request payload content type conversions.
	ContentHandling ContentHandling `field:"optional" json:"contentHandling" yaml:"contentHandling"`
	// Requires that the caller's identity be passed through from the request.
	CredentialsPassthrough *bool `field:"optional" json:"credentialsPassthrough" yaml:"credentialsPassthrough"`
	// An IAM role that API Gateway assumes.
	//
	// Mutually exclusive with `credentialsPassThrough`.
	CredentialsRole awsiam.IRole `field:"optional" json:"credentialsRole" yaml:"credentialsRole"`
	// The response that API Gateway provides after a method's backend completes processing a request.
	//
	// API Gateway intercepts the response from the
	// backend so that you can control how API Gateway surfaces backend
	// responses. For example, you can map the backend status codes to codes
	// that you define.
	IntegrationResponses *[]*IntegrationResponse `field:"optional" json:"integrationResponses" yaml:"integrationResponses"`
	// Specifies the pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the requestTemplates property on the Integration resource.
	//
	// There are three valid values: WHEN_NO_MATCH, WHEN_NO_TEMPLATES, and
	// NEVER.
	PassthroughBehavior PassthroughBehavior `field:"optional" json:"passthroughBehavior" yaml:"passthroughBehavior"`
	// The request parameters that API Gateway sends with the backend request.
	//
	// Specify request parameters as key-value pairs (string-to-string
	// mappings), with a destination as the key and a source as the value.
	//
	// Specify the destination by using the following pattern
	// integration.request.location.name, where location is querystring, path,
	// or header, and name is a valid, unique parameter name.
	//
	// The source must be an existing method request parameter or a static
	// value. You must enclose static values in single quotation marks and
	// pre-encode these values based on their destination in the request.
	RequestParameters *map[string]*string `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// A map of Apache Velocity templates that are applied on the request payload.
	//
	// The template that API Gateway uses is based on the value of the
	// Content-Type header that's sent by the client. The content type value is
	// the key, and the template is the value (specified as a string), such as
	// the following snippet:
	//
	// ```
	//    { "application/json": "{ \"statusCode\": 200 }" }
	// ```.
	// See: http://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html
	//
	RequestTemplates *map[string]*string `field:"optional" json:"requestTemplates" yaml:"requestTemplates"`
	// The maximum amount of time an integration will run before it returns without a response.
	//
	// Must be between 50 milliseconds and 29 seconds.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The VpcLink used for the integration.
	//
	// Required if connectionType is VPC_LINK.
	VpcLink IVpcLink `field:"optional" json:"vpcLink" yaml:"vpcLink"`
	// Allow invoking method from AWS Console UI (for testing purposes).
	//
	// This will add another permission to the AWS Lambda resource policy which
	// will allow the `test-invoke-stage` stage to invoke this handler. If this
	// is set to `false`, the function will only be usable from the deployment
	// endpoint.
	AllowTestInvoke *bool `field:"optional" json:"allowTestInvoke" yaml:"allowTestInvoke"`
	// Use proxy integration or normal (request/response mapping) integration.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-lambda-proxy-integrations.html#api-gateway-simple-proxy-for-lambda-output-format
	//
	Proxy *bool `field:"optional" json:"proxy" yaml:"proxy"`
}

// Defines an API Gateway REST API with AWS Lambda proxy integration.
//
// Use the `proxy` property to define a greedy proxy ("{proxy+}") and "ANY"
// method from the specified path. If not defined, you will need to explicity
// add resources and methods to the API.
//
// Example:
//   var backend function
//
//   api := apigateway.NewLambdaRestApi(this, jsii.String("myapi"), &lambdaRestApiProps{
//   	handler: backend,
//   	proxy: jsii.Boolean(false),
//   })
//
//   items := api.root.addResource(jsii.String("items"))
//   items.addMethod(jsii.String("GET")) // GET /items
//   items.addMethod(jsii.String("POST")) // POST /items
//
//   item := items.addResource(jsii.String("{item}"))
//   item.addMethod(jsii.String("GET")) // GET /items/{item}
//
//   // the default integration for methods is "handler", but one can
//   // customize this behavior per method or even a sub path.
//   item.addMethod(jsii.String("DELETE"), apigateway.NewHttpIntegration(jsii.String("http://amazon.com")))
//
type LambdaRestApi interface {
	RestApi
	CloudWatchAccount() CfnAccount
	SetCloudWatchAccount(val CfnAccount)
	// API Gateway stage that points to the latest deployment (if defined).
	//
	// If `deploy` is disabled, you will need to explicitly assign this value in order to
	// set up integrations.
	DeploymentStage() Stage
	SetDeploymentStage(val Stage)
	// The first domain name mapped to this API, if defined through the `domainName` configuration prop, or added via `addDomainName`.
	DomainName() DomainName
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// API Gateway deployment that represents the latest changes of the API.
	//
	// This resource will be automatically updated every time the REST API model changes.
	// This will be undefined if `deploy` is false.
	LatestDeployment() Deployment
	// The list of methods bound to this RestApi.
	Methods() *[]Method
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The ID of this API Gateway RestApi.
	RestApiId() *string
	// A human friendly name for this Rest API.
	//
	// Note that this is different from `restApiId`.
	RestApiName() *string
	// The resource ID of the root resource.
	RestApiRootResourceId() *string
	// Represents the root resource of this API endpoint ('/').
	//
	// Resources and Methods are added to this resource.
	Root() IResource
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The deployed root URL of this REST API.
	Url() *string
	// Add an ApiKey.
	AddApiKey(id *string, options *ApiKeyOptions) IApiKey
	// Defines an API Gateway domain name and maps it to this API.
	AddDomainName(id *string, options *DomainNameOptions) DomainName
	// Adds a new gateway response.
	AddGatewayResponse(id *string, options *GatewayResponseOptions) GatewayResponse
	// Adds a new model.
	AddModel(id *string, props *ModelOptions) Model
	// Adds a new request validator.
	AddRequestValidator(id *string, props *RequestValidatorOptions) RequestValidator
	// Adds a usage plan.
	AddUsagePlan(id *string, props *UsagePlanProps) UsagePlan
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Gets the "execute-api" ARN.
	ArnForExecuteApi(method *string, path *string, stage *string) *string
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns the given named metric for this API.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of requests served from the API cache in a given period.
	//
	// Default: sum over 5 minutes.
	MetricCacheHitCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of requests served from the backend in a given period, when API caching is enabled.
	//
	// Default: sum over 5 minutes.
	MetricCacheMissCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of client-side errors captured in a given period.
	//
	// Default: sum over 5 minutes.
	MetricClientError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the total number API requests in a given period.
	//
	// Default: sample count over 5 minutes.
	MetricCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the time between when API Gateway relays a request to the backend and when it receives a response from the backend.
	//
	// Default: average over 5 minutes.
	MetricIntegrationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time between when API Gateway receives a request from a client and when it returns a response to the client.
	//
	// The latency includes the integration latency and other API Gateway overhead.
	//
	// Default: average over 5 minutes.
	MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of server-side errors captured in a given period.
	//
	// Default: sum over 5 minutes.
	MetricServerError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns a string representation of this construct.
	ToString() *string
	// Returns the URL for an HTTP path.
	//
	// Fails if `deploymentStage` is not set either by `deploy` or explicitly.
	UrlForPath(path *string) *string
}

// The jsii proxy struct for LambdaRestApi
type jsiiProxy_LambdaRestApi struct {
	jsiiProxy_RestApi
}

func (j *jsiiProxy_LambdaRestApi) CloudWatchAccount() CfnAccount {
	var returns CfnAccount
	_jsii_.Get(
		j,
		"cloudWatchAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRestApi) DeploymentStage() Stage {
	var returns Stage
	_jsii_.Get(
		j,
		"deploymentStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRestApi) DomainName() DomainName {
	var returns DomainName
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRestApi) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRestApi) LatestDeployment() Deployment {
	var returns Deployment
	_jsii_.Get(
		j,
		"latestDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRestApi) Methods() *[]Method {
	var returns *[]Method
	_jsii_.Get(
		j,
		"methods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRestApi) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRestApi) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRestApi) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRestApi) RestApiName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRestApi) RestApiRootResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiRootResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRestApi) Root() IResource {
	var returns IResource
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRestApi) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRestApi) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}


func NewLambdaRestApi(scope constructs.Construct, id *string, props *LambdaRestApiProps) LambdaRestApi {
	_init_.Initialize()

	j := jsiiProxy_LambdaRestApi{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.LambdaRestApi",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewLambdaRestApi_Override(l LambdaRestApi, scope constructs.Construct, id *string, props *LambdaRestApiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.LambdaRestApi",
		[]interface{}{scope, id, props},
		l,
	)
}

func (j *jsiiProxy_LambdaRestApi) SetCloudWatchAccount(val CfnAccount) {
	_jsii_.Set(
		j,
		"cloudWatchAccount",
		val,
	)
}

func (j *jsiiProxy_LambdaRestApi) SetDeploymentStage(val Stage) {
	_jsii_.Set(
		j,
		"deploymentStage",
		val,
	)
}

// Import an existing RestApi that can be configured with additional Methods and Resources.
func LambdaRestApi_FromRestApiAttributes(scope constructs.Construct, id *string, attrs *RestApiAttributes) IRestApi {
	_init_.Initialize()

	var returns IRestApi

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.LambdaRestApi",
		"fromRestApiAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Import an existing RestApi.
func LambdaRestApi_FromRestApiId(scope constructs.Construct, id *string, restApiId *string) IRestApi {
	_init_.Initialize()

	var returns IRestApi

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.LambdaRestApi",
		"fromRestApiId",
		[]interface{}{scope, id, restApiId},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func LambdaRestApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.LambdaRestApi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func LambdaRestApi_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.LambdaRestApi",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func LambdaRestApi_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.LambdaRestApi",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaRestApi) AddApiKey(id *string, options *ApiKeyOptions) IApiKey {
	var returns IApiKey

	_jsii_.Invoke(
		l,
		"addApiKey",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaRestApi) AddDomainName(id *string, options *DomainNameOptions) DomainName {
	var returns DomainName

	_jsii_.Invoke(
		l,
		"addDomainName",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaRestApi) AddGatewayResponse(id *string, options *GatewayResponseOptions) GatewayResponse {
	var returns GatewayResponse

	_jsii_.Invoke(
		l,
		"addGatewayResponse",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaRestApi) AddModel(id *string, props *ModelOptions) Model {
	var returns Model

	_jsii_.Invoke(
		l,
		"addModel",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaRestApi) AddRequestValidator(id *string, props *RequestValidatorOptions) RequestValidator {
	var returns RequestValidator

	_jsii_.Invoke(
		l,
		"addRequestValidator",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaRestApi) AddUsagePlan(id *string, props *UsagePlanProps) UsagePlan {
	var returns UsagePlan

	_jsii_.Invoke(
		l,
		"addUsagePlan",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaRestApi) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		l,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (l *jsiiProxy_LambdaRestApi) ArnForExecuteApi(method *string, path *string, stage *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"arnForExecuteApi",
		[]interface{}{method, path, stage},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaRestApi) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaRestApi) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaRestApi) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaRestApi) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		l,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaRestApi) MetricCacheHitCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		l,
		"metricCacheHitCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaRestApi) MetricCacheMissCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		l,
		"metricCacheMissCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaRestApi) MetricClientError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		l,
		"metricClientError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaRestApi) MetricCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		l,
		"metricCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaRestApi) MetricIntegrationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		l,
		"metricIntegrationLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaRestApi) MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		l,
		"metricLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaRestApi) MetricServerError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		l,
		"metricServerError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaRestApi) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaRestApi) UrlForPath(path *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"urlForPath",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Example:
//   var backend function
//
//   api := apigateway.NewLambdaRestApi(this, jsii.String("myapi"), &lambdaRestApiProps{
//   	handler: backend,
//   	proxy: jsii.Boolean(false),
//   })
//
//   items := api.root.addResource(jsii.String("items"))
//   items.addMethod(jsii.String("GET")) // GET /items
//   items.addMethod(jsii.String("POST")) // POST /items
//
//   item := items.addResource(jsii.String("{item}"))
//   item.addMethod(jsii.String("GET")) // GET /items/{item}
//
//   // the default integration for methods is "handler", but one can
//   // customize this behavior per method or even a sub path.
//   item.addMethod(jsii.String("DELETE"), apigateway.NewHttpIntegration(jsii.String("http://amazon.com")))
//
type LambdaRestApiProps struct {
	// Adds a CORS preflight OPTIONS method to this resource and all child resources.
	//
	// You can add CORS at the resource-level using `addCorsPreflight`.
	DefaultCorsPreflightOptions *CorsOptions `field:"optional" json:"defaultCorsPreflightOptions" yaml:"defaultCorsPreflightOptions"`
	// An integration to use as a default for all methods created within this API unless an integration is specified.
	DefaultIntegration Integration `field:"optional" json:"defaultIntegration" yaml:"defaultIntegration"`
	// Method options to use as a default for all methods created within this API unless custom options are specified.
	DefaultMethodOptions *MethodOptions `field:"optional" json:"defaultMethodOptions" yaml:"defaultMethodOptions"`
	// Automatically configure an AWS CloudWatch role for API Gateway.
	CloudWatchRole *bool `field:"optional" json:"cloudWatchRole" yaml:"cloudWatchRole"`
	// Indicates if a Deployment should be automatically created for this API, and recreated when the API model (resources, methods) changes.
	//
	// Since API Gateway deployments are immutable, When this option is enabled
	// (by default), an AWS::ApiGateway::Deployment resource will automatically
	// created with a logical ID that hashes the API model (methods, resources
	// and options). This means that when the model changes, the logical ID of
	// this CloudFormation resource will change, and a new deployment will be
	// created.
	//
	// If this is set, `latestDeployment` will refer to the `Deployment` object
	// and `deploymentStage` will refer to a `Stage` that points to this
	// deployment. To customize the stage options, use the `deployOptions`
	// property.
	//
	// A CloudFormation Output will also be defined with the root URL endpoint
	// of this REST API.
	Deploy *bool `field:"optional" json:"deploy" yaml:"deploy"`
	// Options for the API Gateway stage that will always point to the latest deployment when `deploy` is enabled.
	//
	// If `deploy` is disabled,
	// this value cannot be set.
	DeployOptions *StageOptions `field:"optional" json:"deployOptions" yaml:"deployOptions"`
	// A description of the RestApi construct.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether clients can invoke the API using the default execute-api endpoint.
	//
	// To require that clients use a custom domain name to invoke the
	// API, disable the default endpoint.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html
	//
	DisableExecuteApiEndpoint *bool `field:"optional" json:"disableExecuteApiEndpoint" yaml:"disableExecuteApiEndpoint"`
	// Configure a custom domain name and map it to this API.
	DomainName *DomainNameOptions `field:"optional" json:"domainName" yaml:"domainName"`
	// Export name for the CfnOutput containing the API endpoint.
	EndpointExportName *string `field:"optional" json:"endpointExportName" yaml:"endpointExportName"`
	// A list of the endpoint types of the API.
	//
	// Use this property when creating
	// an API.
	EndpointTypes *[]EndpointType `field:"optional" json:"endpointTypes" yaml:"endpointTypes"`
	// Indicates whether to roll back the resource if a warning occurs while API Gateway is creating the RestApi resource.
	FailOnWarnings *bool `field:"optional" json:"failOnWarnings" yaml:"failOnWarnings"`
	// Custom header parameters for the request.
	// See: https://docs.aws.amazon.com/cli/latest/reference/apigateway/import-rest-api.html
	//
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// A policy document that contains the permissions for this RestApi.
	Policy awsiam.PolicyDocument `field:"optional" json:"policy" yaml:"policy"`
	// A name for the API Gateway RestApi resource.
	RestApiName *string `field:"optional" json:"restApiName" yaml:"restApiName"`
	// Retains old deployment resources when the API changes.
	//
	// This allows
	// manually reverting stages to point to old deployments via the AWS
	// Console.
	RetainDeployments *bool `field:"optional" json:"retainDeployments" yaml:"retainDeployments"`
	// The source of the API key for metering requests according to a usage plan.
	ApiKeySourceType ApiKeySourceType `field:"optional" json:"apiKeySourceType" yaml:"apiKeySourceType"`
	// The list of binary media mime-types that are supported by the RestApi resource, such as "image/png" or "application/octet-stream".
	BinaryMediaTypes *[]*string `field:"optional" json:"binaryMediaTypes" yaml:"binaryMediaTypes"`
	// The ID of the API Gateway RestApi resource that you want to clone.
	CloneFrom IRestApi `field:"optional" json:"cloneFrom" yaml:"cloneFrom"`
	// The EndpointConfiguration property type specifies the endpoint types of a REST API.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-restapi-endpointconfiguration.html
	//
	EndpointConfiguration *EndpointConfiguration `field:"optional" json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// A nullable integer that is used to enable compression (with non-negative between 0 and 10485760 (10M) bytes, inclusive) or disable compression (when undefined) on an API.
	//
	// When compression is enabled, compression or
	// decompression is not applied on the payload if the payload size is
	// smaller than this value. Setting it to zero allows compression for any
	// payload size.
	MinimumCompressionSize *float64 `field:"optional" json:"minimumCompressionSize" yaml:"minimumCompressionSize"`
	// The default Lambda function that handles all requests from this API.
	//
	// This handler will be used as a the default integration for all methods in
	// this API, unless specified otherwise in `addMethod`.
	Handler awslambda.IFunction `field:"required" json:"handler" yaml:"handler"`
	// Specific Lambda integration options.
	IntegrationOptions *LambdaIntegrationOptions `field:"optional" json:"integrationOptions" yaml:"integrationOptions"`
	// If true, route all requests to the Lambda Function.
	//
	// If set to false, you will need to explicitly define the API model using
	// `addResource` and `addMethod` (or `addProxy`).
	Proxy *bool `field:"optional" json:"proxy" yaml:"proxy"`
}

// Use CloudWatch Logs as a custom access log destination for API Gateway.
//
// Example:
//   logGroup := logs.NewLogGroup(this, jsii.String("ApiGatewayAccessLogs"))
//   apigateway.NewRestApi(this, jsii.String("books"), &restApiProps{
//   	deployOptions: &stageOptions{
//   		accessLogDestination: apigateway.NewLogGroupLogDestination(logGroup),
//   		accessLogFormat: apigateway.accessLogFormat.custom(
//   		fmt.Sprintf("%v %v %v", apigateway.accessLogField.contextRequestId(), apigateway.*accessLogField.contextErrorMessage(), apigateway.*accessLogField.contextErrorMessageString())),
//   	},
//   })
//
type LogGroupLogDestination interface {
	IAccessLogDestination
	// Binds this destination to the CloudWatch Logs.
	Bind(_stage IStage) *AccessLogDestinationConfig
}

// The jsii proxy struct for LogGroupLogDestination
type jsiiProxy_LogGroupLogDestination struct {
	jsiiProxy_IAccessLogDestination
}

func NewLogGroupLogDestination(logGroup awslogs.ILogGroup) LogGroupLogDestination {
	_init_.Initialize()

	j := jsiiProxy_LogGroupLogDestination{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.LogGroupLogDestination",
		[]interface{}{logGroup},
		&j,
	)

	return &j
}

func NewLogGroupLogDestination_Override(l LogGroupLogDestination, logGroup awslogs.ILogGroup) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.LogGroupLogDestination",
		[]interface{}{logGroup},
		l,
	)
}

func (l *jsiiProxy_LogGroupLogDestination) Bind(_stage IStage) *AccessLogDestinationConfig {
	var returns *AccessLogDestinationConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{_stage},
		&returns,
	)

	return returns
}

// The mTLS authentication configuration for a custom domain name.
//
// Example:
//   var acm interface{}
//
//
//   apigateway.NewDomainName(this, jsii.String("domain-name"), &domainNameProps{
//   	domainName: jsii.String("example.com"),
//   	certificate: acm.certificate_FromCertificateArn(this, jsii.String("cert"), jsii.String("arn:aws:acm:us-east-1:1111111:certificate/11-3336f1-44483d-adc7-9cd375c5169d")),
//   	mtls: &mTLSConfig{
//   		bucket: s3.NewBucket(this, jsii.String("bucket")),
//   		key: jsii.String("truststore.pem"),
//   		version: jsii.String("version"),
//   	},
//   })
//
type MTLSConfig struct {
	// The bucket that the trust store is hosted in.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The key in S3 to look at for the trust store.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The version of the S3 object that contains your truststore.
	//
	// To specify a version, you must have versioning enabled for the S3 bucket.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

// Example:
//   var integration lambdaIntegration
//
//
//   api := apigateway.NewRestApi(this, jsii.String("hello-api"))
//
//   v1 := api.root.addResource(jsii.String("v1"))
//   echo := v1.addResource(jsii.String("echo"))
//   echoMethod := echo.addMethod(jsii.String("GET"), integration, &methodOptions{
//   	apiKeyRequired: jsii.Boolean(true),
//   })
//
//   plan := api.addUsagePlan(jsii.String("UsagePlan"), &usagePlanProps{
//   	name: jsii.String("Easy"),
//   	throttle: &throttleSettings{
//   		rateLimit: jsii.Number(10),
//   		burstLimit: jsii.Number(2),
//   	},
//   })
//
//   key := api.addApiKey(jsii.String("ApiKey"))
//   plan.addApiKey(key)
//
type Method interface {
	awscdk.Resource
	// The API Gateway RestApi associated with this method.
	Api() IRestApi
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	HttpMethod() *string
	// Returns an execute-api ARN for this method:.
	//
	// arn:aws:execute-api:{region}:{account}:{restApiId}/{stage}/{method}/{path}
	//
	// NOTE: {stage} will refer to the `restApi.deploymentStage`, which will
	// automatically set if auto-deploy is enabled, or can be explicitly assigned.
	// When not configured, {stage} will be set to '*', as a shorthand for 'all stages'.
	MethodArn() *string
	MethodId() *string
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	Resource() IResource
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Returns an execute-api ARN for this method's "test-invoke-stage" stage.
	//
	// This stage is used by the AWS Console UI when testing the method.
	TestMethodArn() *string
	// Add a method response to this method.
	AddMethodResponse(methodResponse *MethodResponse)
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns the given named metric for this API method.
	Metric(metricName *string, stage IStage, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of requests served from the API cache in a given period.
	MetricCacheHitCount(stage IStage, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of requests served from the backend in a given period, when API caching is enabled.
	MetricCacheMissCount(stage IStage, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of client-side errors captured in a given period.
	MetricClientError(stage IStage, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the total number API requests in a given period.
	MetricCount(stage IStage, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the time between when API Gateway relays a request to the backend and when it receives a response from the backend.
	MetricIntegrationLatency(stage IStage, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time between when API Gateway receives a request from a client and when it returns a response to the client.
	//
	// The latency includes the integration latency and other API Gateway overhead.
	MetricLatency(stage IStage, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of server-side errors captured in a given period.
	MetricServerError(stage IStage, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Method
type jsiiProxy_Method struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_Method) Api() IRestApi {
	var returns IRestApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Method) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Method) HttpMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Method) MethodArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"methodArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Method) MethodId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"methodId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Method) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Method) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Method) Resource() IResource {
	var returns IResource
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Method) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Method) TestMethodArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"testMethodArn",
		&returns,
	)
	return returns
}


func NewMethod(scope constructs.Construct, id *string, props *MethodProps) Method {
	_init_.Initialize()

	j := jsiiProxy_Method{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.Method",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewMethod_Override(m Method, scope constructs.Construct, id *string, props *MethodProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.Method",
		[]interface{}{scope, id, props},
		m,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Method_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.Method",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func Method_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.Method",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Method_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.Method",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Method) AddMethodResponse(methodResponse *MethodResponse) {
	_jsii_.InvokeVoid(
		m,
		"addMethodResponse",
		[]interface{}{methodResponse},
	)
}

func (m *jsiiProxy_Method) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		m,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (m *jsiiProxy_Method) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Method) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Method) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Method) Metric(metricName *string, stage IStage, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		m,
		"metric",
		[]interface{}{metricName, stage, props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Method) MetricCacheHitCount(stage IStage, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		m,
		"metricCacheHitCount",
		[]interface{}{stage, props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Method) MetricCacheMissCount(stage IStage, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		m,
		"metricCacheMissCount",
		[]interface{}{stage, props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Method) MetricClientError(stage IStage, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		m,
		"metricClientError",
		[]interface{}{stage, props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Method) MetricCount(stage IStage, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		m,
		"metricCount",
		[]interface{}{stage, props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Method) MetricIntegrationLatency(stage IStage, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		m,
		"metricIntegrationLatency",
		[]interface{}{stage, props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Method) MetricLatency(stage IStage, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		m,
		"metricLatency",
		[]interface{}{stage, props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Method) MetricServerError(stage IStage, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		m,
		"metricServerError",
		[]interface{}{stage, props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Method) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Example:
//   api := apigateway.NewRestApi(this, jsii.String("books"))
//   deployment := apigateway.NewDeployment(this, jsii.String("my-deployment"), &deploymentProps{
//   	api: api,
//   })
//   stage := apigateway.NewStage(this, jsii.String("my-stage"), &stageProps{
//   	deployment: deployment,
//   	methodOptions: map[string]methodDeploymentOptions{
//   		"/*/*": &methodDeploymentOptions{
//   			 // This special path applies to all resource paths and all HTTP methods
//   			"throttlingRateLimit": jsii.Number(100),
//   			"throttlingBurstLimit": jsii.Number(200),
//   		},
//   	},
//   })
//
type MethodDeploymentOptions struct {
	// Indicates whether the cached responses are encrypted.
	CacheDataEncrypted *bool `field:"optional" json:"cacheDataEncrypted" yaml:"cacheDataEncrypted"`
	// Specifies the time to live (TTL), in seconds, for cached responses.
	//
	// The
	// higher the TTL, the longer the response will be cached.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-caching.html
	//
	CacheTtl awscdk.Duration `field:"optional" json:"cacheTtl" yaml:"cacheTtl"`
	// Specifies whether responses should be cached and returned for requests.
	//
	// A
	// cache cluster must be enabled on the stage for responses to be cached.
	CachingEnabled *bool `field:"optional" json:"cachingEnabled" yaml:"cachingEnabled"`
	// Specifies whether data trace logging is enabled for this method.
	//
	// When enabled, API gateway will log the full API requests and responses.
	// This can be useful to troubleshoot APIs, but can result in logging sensitive data.
	// We recommend that you don't enable this feature for production APIs.
	DataTraceEnabled *bool `field:"optional" json:"dataTraceEnabled" yaml:"dataTraceEnabled"`
	// Specifies the logging level for this method, which effects the log entries pushed to Amazon CloudWatch Logs.
	LoggingLevel MethodLoggingLevel `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
	// Specifies whether Amazon CloudWatch metrics are enabled for this method.
	MetricsEnabled *bool `field:"optional" json:"metricsEnabled" yaml:"metricsEnabled"`
	// Specifies the throttling burst limit.
	//
	// The total rate of all requests in your AWS account is limited to 5,000 requests.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html
	//
	ThrottlingBurstLimit *float64 `field:"optional" json:"throttlingBurstLimit" yaml:"throttlingBurstLimit"`
	// Specifies the throttling rate limit.
	//
	// The total rate of all requests in your AWS account is limited to 10,000 requests per second (rps).
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html
	//
	ThrottlingRateLimit *float64 `field:"optional" json:"throttlingRateLimit" yaml:"throttlingRateLimit"`
}

// Example:
//   api := apigateway.NewRestApi(this, jsii.String("books"), &restApiProps{
//   	deployOptions: &stageOptions{
//   		loggingLevel: apigateway.methodLoggingLevel_INFO,
//   		dataTraceEnabled: jsii.Boolean(true),
//   	},
//   })
//
type MethodLoggingLevel string

const (
	MethodLoggingLevel_OFF MethodLoggingLevel = "OFF"
	MethodLoggingLevel_ERROR MethodLoggingLevel = "ERROR"
	MethodLoggingLevel_INFO MethodLoggingLevel = "INFO"
)

// Example:
//   var api restApi
//   var userLambda function
//
//
//   userModel := api.addModel(jsii.String("UserModel"), &modelOptions{
//   	schema: &jsonSchema{
//   		type: apigateway.jsonSchemaType_OBJECT,
//   		properties: map[string]*jsonSchema{
//   			"userId": &jsonSchema{
//   				"type": apigateway.*jsonSchemaType_STRING,
//   			},
//   			"name": &jsonSchema{
//   				"type": apigateway.*jsonSchemaType_STRING,
//   			},
//   		},
//   		required: []*string{
//   			jsii.String("userId"),
//   		},
//   	},
//   })
//   api.root.addResource(jsii.String("user")).addMethod(jsii.String("POST"),
//   apigateway.NewLambdaIntegration(userLambda), &methodOptions{
//   	requestModels: map[string]iModel{
//   		"application/json": userModel,
//   	},
//   })
//
type MethodOptions struct {
	// Indicates whether the method requires clients to submit a valid API key.
	ApiKeyRequired *bool `field:"optional" json:"apiKeyRequired" yaml:"apiKeyRequired"`
	// A list of authorization scopes configured on the method.
	//
	// The scopes are used with
	// a COGNITO_USER_POOLS authorizer to authorize the method invocation.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-authorizationscopes
	//
	AuthorizationScopes *[]*string `field:"optional" json:"authorizationScopes" yaml:"authorizationScopes"`
	// Method authorization. If the value is set of `Custom`, an `authorizer` must also be specified.
	//
	// If you're using one of the authorizers that are available via the {@link Authorizer} class, such as {@link Authorizer#token()},
	// it is recommended that this option not be specified. The authorizer will take care of setting the correct authorization type.
	// However, specifying an authorization type using this property that conflicts with what is expected by the {@link Authorizer}
	// will result in an error.
	AuthorizationType AuthorizationType `field:"optional" json:"authorizationType" yaml:"authorizationType"`
	// If `authorizationType` is `Custom`, this specifies the ID of the method authorizer resource.
	//
	// If specified, the value of `authorizationType` must be set to `Custom`.
	Authorizer IAuthorizer `field:"optional" json:"authorizer" yaml:"authorizer"`
	// The responses that can be sent to the client who calls the method.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-method-settings-method-response.html
	//
	MethodResponses *[]*MethodResponse `field:"optional" json:"methodResponses" yaml:"methodResponses"`
	// A friendly operation name for the method.
	//
	// For example, you can assign the
	// OperationName of ListPets for the GET /pets method.
	OperationName *string `field:"optional" json:"operationName" yaml:"operationName"`
	// The models which describe data structure of request payload.
	//
	// When
	// combined with `requestValidator` or `requestValidatorOptions`, the service
	// will validate the API request payload before it reaches the API's Integration (including proxies).
	// Specify `requestModels` as key-value pairs, with a content type
	// (e.g. `'application/json'`) as the key and an API Gateway Model as the value.
	//
	// Example:
	//   var api restApi
	//   var userLambda function
	//
	//
	//   userModel := api.addModel(jsii.String("UserModel"), &modelOptions{
	//   	schema: &jsonSchema{
	//   		type: apigateway.jsonSchemaType_OBJECT,
	//   		properties: map[string]*jsonSchema{
	//   			"userId": &jsonSchema{
	//   				"type": apigateway.*jsonSchemaType_STRING,
	//   			},
	//   			"name": &jsonSchema{
	//   				"type": apigateway.*jsonSchemaType_STRING,
	//   			},
	//   		},
	//   		required: []*string{
	//   			jsii.String("userId"),
	//   		},
	//   	},
	//   })
	//   api.root.addResource(jsii.String("user")).addMethod(jsii.String("POST"),
	//   apigateway.NewLambdaIntegration(userLambda), &methodOptions{
	//   	requestModels: map[string]iModel{
	//   		"application/json": userModel,
	//   	},
	//   })
	//
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-method-settings-method-request.html#setup-method-request-model
	//
	RequestModels *map[string]IModel `field:"optional" json:"requestModels" yaml:"requestModels"`
	// The request parameters that API Gateway accepts.
	//
	// Specify request parameters
	// as key-value pairs (string-to-Boolean mapping), with a source as the key and
	// a Boolean as the value. The Boolean specifies whether a parameter is required.
	// A source must match the format method.request.location.name, where the location
	// is querystring, path, or header, and name is a valid, unique parameter name.
	RequestParameters *map[string]*bool `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// The ID of the associated request validator.
	//
	// Only one of `requestValidator` or `requestValidatorOptions` must be specified.
	// Works together with `requestModels` or `requestParameters` to validate
	// the request before it reaches integration like Lambda Proxy Integration.
	RequestValidator IRequestValidator `field:"optional" json:"requestValidator" yaml:"requestValidator"`
	// Request validator options to create new validator Only one of `requestValidator` or `requestValidatorOptions` must be specified.
	//
	// Works together with `requestModels` or `requestParameters` to validate
	// the request before it reaches integration like Lambda Proxy Integration.
	RequestValidatorOptions *RequestValidatorOptions `field:"optional" json:"requestValidatorOptions" yaml:"requestValidatorOptions"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var authorizer authorizer
//   var integration integration
//   var model model
//   var requestValidator requestValidator
//   var resource resource
//
//   methodProps := &methodProps{
//   	httpMethod: jsii.String("httpMethod"),
//   	resource: resource,
//
//   	// the properties below are optional
//   	integration: integration,
//   	options: &methodOptions{
//   		apiKeyRequired: jsii.Boolean(false),
//   		authorizationScopes: []*string{
//   			jsii.String("authorizationScopes"),
//   		},
//   		authorizationType: awscdk.Aws_apigateway.authorizationType_NONE,
//   		authorizer: authorizer,
//   		methodResponses: []methodResponse{
//   			&methodResponse{
//   				statusCode: jsii.String("statusCode"),
//
//   				// the properties below are optional
//   				responseModels: map[string]iModel{
//   					"responseModelsKey": model,
//   				},
//   				responseParameters: map[string]*bool{
//   					"responseParametersKey": jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		operationName: jsii.String("operationName"),
//   		requestModels: map[string]*iModel{
//   			"requestModelsKey": model,
//   		},
//   		requestParameters: map[string]*bool{
//   			"requestParametersKey": jsii.Boolean(false),
//   		},
//   		requestValidator: requestValidator,
//   		requestValidatorOptions: &requestValidatorOptions{
//   			requestValidatorName: jsii.String("requestValidatorName"),
//   			validateRequestBody: jsii.Boolean(false),
//   			validateRequestParameters: jsii.Boolean(false),
//   		},
//   	},
//   }
//
type MethodProps struct {
	// The HTTP method ("GET", "POST", "PUT", ...) that clients use to call this method.
	HttpMethod *string `field:"required" json:"httpMethod" yaml:"httpMethod"`
	// The resource this method is associated with.
	//
	// For root resource methods,
	// specify the `RestApi` object.
	Resource IResource `field:"required" json:"resource" yaml:"resource"`
	// The backend system that the method calls when it receives a request.
	Integration Integration `field:"optional" json:"integration" yaml:"integration"`
	// Method options.
	Options *MethodOptions `field:"optional" json:"options" yaml:"options"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var model model
//
//   methodResponse := &methodResponse{
//   	statusCode: jsii.String("statusCode"),
//
//   	// the properties below are optional
//   	responseModels: map[string]iModel{
//   		"responseModelsKey": model,
//   	},
//   	responseParameters: map[string]*bool{
//   		"responseParametersKey": jsii.Boolean(false),
//   	},
//   }
//
type MethodResponse struct {
	// The method response's status code, which you map to an IntegrationResponse.
	//
	// Required.
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
	// The resources used for the response's content type.
	//
	// Specify response models as
	// key-value pairs (string-to-string maps), with a content type as the key and a Model
	// resource name as the value.
	ResponseModels *map[string]IModel `field:"optional" json:"responseModels" yaml:"responseModels"`
	// Response parameters that API Gateway sends to the client that called a method.
	//
	// Specify response parameters as key-value pairs (string-to-Boolean maps), with
	// a destination as the key and a Boolean as the value. Specify the destination
	// using the following pattern: method.response.header.name, where the name is a
	// valid, unique header name. The Boolean specifies whether a parameter is required.
	ResponseParameters *map[string]*bool `field:"optional" json:"responseParameters" yaml:"responseParameters"`
}

// This type of integration lets API Gateway return a response without sending the request further to the backend.
//
// This is useful for API testing because it
// can be used to test the integration set up without incurring charges for
// using the backend and to enable collaborative development of an API. In
// collaborative development, a team can isolate their development effort by
// setting up simulations of API components owned by other teams by using the
// MOCK integrations. It is also used to return CORS-related headers to ensure
// that the API method permits CORS access. In fact, the API Gateway console
// integrates the OPTIONS method to support CORS with a mock integration.
// Gateway responses are other examples of mock integrations.
//
// Example:
//   import path "github.com/aws-samples/dummy/path"
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   // Against the RestApi endpoint from the stack output, run
//   // `curl -s -o /dev/null -w "%{http_code}" <url>` should return 401
//   // `curl -s -o /dev/null -w "%{http_code}" -H 'Authorization: deny' <url>?allow=yes` should return 403
//   // `curl -s -o /dev/null -w "%{http_code}" -H 'Authorization: allow' <url>?allow=yes` should return 200
//
//   app := awscdk.NewApp()
//   stack := awscdk.NewStack(app, jsii.String("RequestAuthorizerInteg"))
//
//   authorizerFn := lambda.NewFunction(stack, jsii.String("MyAuthorizerFunction"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.assetCode.fromAsset(path.join(__dirname, jsii.String("integ.request-authorizer.handler"))),
//   })
//
//   restapi := awscdk.NewRestApi(stack, jsii.String("MyRestApi"))
//
//   authorizer := awscdk.NewRequestAuthorizer(stack, jsii.String("MyAuthorizer"), &requestAuthorizerProps{
//   	handler: authorizerFn,
//   	identitySources: []*string{
//   		awscdk.IdentitySource.header(jsii.String("Authorization")),
//   		awscdk.IdentitySource.queryString(jsii.String("allow")),
//   	},
//   })
//
//   restapi.root.addMethod(jsii.String("ANY"), awscdk.NewMockIntegration(&integrationOptions{
//   	integrationResponses: []integrationResponse{
//   		&integrationResponse{
//   			statusCode: jsii.String("200"),
//   		},
//   	},
//   	passthroughBehavior: awscdk.PassthroughBehavior_NEVER,
//   	requestTemplates: map[string]*string{
//   		"application/json": jsii.String("{ \"statusCode\": 200 }"),
//   	},
//   }), &methodOptions{
//   	methodResponses: []methodResponse{
//   		&methodResponse{
//   			statusCode: jsii.String("200"),
//   		},
//   	},
//   	authorizer: authorizer,
//   })
//
type MockIntegration interface {
	Integration
	// Can be overridden by subclasses to allow the integration to interact with the method being integrated, access the REST API object, method ARNs, etc.
	Bind(_method Method) *IntegrationConfig
}

// The jsii proxy struct for MockIntegration
type jsiiProxy_MockIntegration struct {
	jsiiProxy_Integration
}

func NewMockIntegration(options *IntegrationOptions) MockIntegration {
	_init_.Initialize()

	j := jsiiProxy_MockIntegration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.MockIntegration",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewMockIntegration_Override(m MockIntegration, options *IntegrationOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.MockIntegration",
		[]interface{}{options},
		m,
	)
}

func (m *jsiiProxy_MockIntegration) Bind(_method Method) *IntegrationConfig {
	var returns *IntegrationConfig

	_jsii_.Invoke(
		m,
		"bind",
		[]interface{}{_method},
		&returns,
	)

	return returns
}

// Example:
//   var api restApi
//
//
//   // We define the JSON Schema for the transformed valid response
//   responseModel := api.addModel(jsii.String("ResponseModel"), &modelOptions{
//   	contentType: jsii.String("application/json"),
//   	modelName: jsii.String("ResponseModel"),
//   	schema: &jsonSchema{
//   		schema: apigateway.jsonSchemaVersion_DRAFT4,
//   		title: jsii.String("pollResponse"),
//   		type: apigateway.jsonSchemaType_OBJECT,
//   		properties: map[string]*jsonSchema{
//   			"state": &jsonSchema{
//   				"type": apigateway.*jsonSchemaType_STRING,
//   			},
//   			"greeting": &jsonSchema{
//   				"type": apigateway.*jsonSchemaType_STRING,
//   			},
//   		},
//   	},
//   })
//
//   // We define the JSON Schema for the transformed error response
//   errorResponseModel := api.addModel(jsii.String("ErrorResponseModel"), &modelOptions{
//   	contentType: jsii.String("application/json"),
//   	modelName: jsii.String("ErrorResponseModel"),
//   	schema: &jsonSchema{
//   		schema: apigateway.*jsonSchemaVersion_DRAFT4,
//   		title: jsii.String("errorResponse"),
//   		type: apigateway.*jsonSchemaType_OBJECT,
//   		properties: map[string]*jsonSchema{
//   			"state": &jsonSchema{
//   				"type": apigateway.*jsonSchemaType_STRING,
//   			},
//   			"message": &jsonSchema{
//   				"type": apigateway.*jsonSchemaType_STRING,
//   			},
//   		},
//   	},
//   })
//
type Model interface {
	awscdk.Resource
	IModel
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// Returns the model name, such as 'myModel'.
	ModelId() *string
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Model
type jsiiProxy_Model struct {
	internal.Type__awscdkResource
	jsiiProxy_IModel
}

func (j *jsiiProxy_Model) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Model) ModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Model) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Model) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Model) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewModel(scope constructs.Construct, id *string, props *ModelProps) Model {
	_init_.Initialize()

	j := jsiiProxy_Model{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.Model",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewModel_Override(m Model, scope constructs.Construct, id *string, props *ModelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.Model",
		[]interface{}{scope, id, props},
		m,
	)
}

func Model_FromModelName(scope constructs.Construct, id *string, modelName *string) IModel {
	_init_.Initialize()

	var returns IModel

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.Model",
		"fromModelName",
		[]interface{}{scope, id, modelName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Model_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.Model",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func Model_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.Model",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Model_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.Model",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func Model_EMPTY_MODEL() IModel {
	_init_.Initialize()
	var returns IModel
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.Model",
		"EMPTY_MODEL",
		&returns,
	)
	return returns
}

func Model_ERROR_MODEL() IModel {
	_init_.Initialize()
	var returns IModel
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.Model",
		"ERROR_MODEL",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_Model) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		m,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (m *jsiiProxy_Model) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Model) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Model) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Model) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Example:
//   var api restApi
//
//
//   // We define the JSON Schema for the transformed valid response
//   responseModel := api.addModel(jsii.String("ResponseModel"), &modelOptions{
//   	contentType: jsii.String("application/json"),
//   	modelName: jsii.String("ResponseModel"),
//   	schema: &jsonSchema{
//   		schema: apigateway.jsonSchemaVersion_DRAFT4,
//   		title: jsii.String("pollResponse"),
//   		type: apigateway.jsonSchemaType_OBJECT,
//   		properties: map[string]*jsonSchema{
//   			"state": &jsonSchema{
//   				"type": apigateway.*jsonSchemaType_STRING,
//   			},
//   			"greeting": &jsonSchema{
//   				"type": apigateway.*jsonSchemaType_STRING,
//   			},
//   		},
//   	},
//   })
//
//   // We define the JSON Schema for the transformed error response
//   errorResponseModel := api.addModel(jsii.String("ErrorResponseModel"), &modelOptions{
//   	contentType: jsii.String("application/json"),
//   	modelName: jsii.String("ErrorResponseModel"),
//   	schema: &jsonSchema{
//   		schema: apigateway.*jsonSchemaVersion_DRAFT4,
//   		title: jsii.String("errorResponse"),
//   		type: apigateway.*jsonSchemaType_OBJECT,
//   		properties: map[string]*jsonSchema{
//   			"state": &jsonSchema{
//   				"type": apigateway.*jsonSchemaType_STRING,
//   			},
//   			"message": &jsonSchema{
//   				"type": apigateway.*jsonSchemaType_STRING,
//   			},
//   		},
//   	},
//   })
//
type ModelOptions struct {
	// The schema to use to transform data to one or more output formats.
	//
	// Specify null ({}) if you don't want to specify a schema.
	Schema *JsonSchema `field:"required" json:"schema" yaml:"schema"`
	// The content type for the model.
	//
	// You can also force a
	// content type in the request or response model mapping.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// A description that identifies this model.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name for the model.
	//
	// Important
	//   If you specify a name, you cannot perform updates that
	//   require replacement of this resource. You can perform
	//   updates that require no or some interruption. If you
	//   must replace the resource, specify a new name.
	ModelName *string `field:"optional" json:"modelName" yaml:"modelName"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var default_ interface{}
//   var enum_ interface{}
//   var jsonSchema_ jsonSchema
//   var restApi restApi
//
//   modelProps := &modelProps{
//   	restApi: restApi,
//   	schema: &jsonSchema{
//   		additionalItems: []*jsonSchema{
//   			jsonSchema_,
//   		},
//   		additionalProperties: jsii.Boolean(false),
//   		allOf: []*jsonSchema{
//   			jsonSchema_,
//   		},
//   		anyOf: []*jsonSchema{
//   			jsonSchema_,
//   		},
//   		contains: jsonSchema_,
//   		default: default_,
//   		definitions: map[string]*jsonSchema{
//   			"definitionsKey": jsonSchema_,
//   		},
//   		dependencies: map[string]interface{}{
//   			"dependenciesKey": jsonSchema_,
//   		},
//   		description: jsii.String("description"),
//   		enum: []interface{}{
//   			enum_,
//   		},
//   		exclusiveMaximum: jsii.Boolean(false),
//   		exclusiveMinimum: jsii.Boolean(false),
//   		format: jsii.String("format"),
//   		id: jsii.String("id"),
//   		items: jsonSchema_,
//   		maximum: jsii.Number(123),
//   		maxItems: jsii.Number(123),
//   		maxLength: jsii.Number(123),
//   		maxProperties: jsii.Number(123),
//   		minimum: jsii.Number(123),
//   		minItems: jsii.Number(123),
//   		minLength: jsii.Number(123),
//   		minProperties: jsii.Number(123),
//   		multipleOf: jsii.Number(123),
//   		not: jsonSchema_,
//   		oneOf: []*jsonSchema{
//   			jsonSchema_,
//   		},
//   		pattern: jsii.String("pattern"),
//   		patternProperties: map[string]*jsonSchema{
//   			"patternPropertiesKey": jsonSchema_,
//   		},
//   		properties: map[string]*jsonSchema{
//   			"propertiesKey": jsonSchema_,
//   		},
//   		propertyNames: jsonSchema_,
//   		ref: jsii.String("ref"),
//   		required: []*string{
//   			jsii.String("required"),
//   		},
//   		schema: awscdk.Aws_apigateway.jsonSchemaVersion_DRAFT4,
//   		title: jsii.String("title"),
//   		type: awscdk.*Aws_apigateway.jsonSchemaType_NULL,
//   		uniqueItems: jsii.Boolean(false),
//   	},
//
//   	// the properties below are optional
//   	contentType: jsii.String("contentType"),
//   	description: jsii.String("description"),
//   	modelName: jsii.String("modelName"),
//   }
//
type ModelProps struct {
	// The schema to use to transform data to one or more output formats.
	//
	// Specify null ({}) if you don't want to specify a schema.
	Schema *JsonSchema `field:"required" json:"schema" yaml:"schema"`
	// The content type for the model.
	//
	// You can also force a
	// content type in the request or response model mapping.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// A description that identifies this model.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name for the model.
	//
	// Important
	//   If you specify a name, you cannot perform updates that
	//   require replacement of this resource. You can perform
	//   updates that require no or some interruption. If you
	//   must replace the resource, specify a new name.
	ModelName *string `field:"optional" json:"modelName" yaml:"modelName"`
	// The rest API that this model is part of.
	//
	// The reason we need the RestApi object itself and not just the ID is because the model
	// is being tracked by the top-level RestApi object for the purpose of calculating it's
	// hash to determine the ID of the deployment. This allows us to automatically update
	// the deployment when the model of the REST API changes.
	RestApi IRestApi `field:"required" json:"restApi" yaml:"restApi"`
}

// Example:
//   import path "github.com/aws-samples/dummy/path"
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   // Against the RestApi endpoint from the stack output, run
//   // `curl -s -o /dev/null -w "%{http_code}" <url>` should return 401
//   // `curl -s -o /dev/null -w "%{http_code}" -H 'Authorization: deny' <url>?allow=yes` should return 403
//   // `curl -s -o /dev/null -w "%{http_code}" -H 'Authorization: allow' <url>?allow=yes` should return 200
//
//   app := awscdk.NewApp()
//   stack := awscdk.NewStack(app, jsii.String("RequestAuthorizerInteg"))
//
//   authorizerFn := lambda.NewFunction(stack, jsii.String("MyAuthorizerFunction"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.assetCode.fromAsset(path.join(__dirname, jsii.String("integ.request-authorizer.handler"))),
//   })
//
//   restapi := awscdk.NewRestApi(stack, jsii.String("MyRestApi"))
//
//   authorizer := awscdk.NewRequestAuthorizer(stack, jsii.String("MyAuthorizer"), &requestAuthorizerProps{
//   	handler: authorizerFn,
//   	identitySources: []*string{
//   		awscdk.IdentitySource.header(jsii.String("Authorization")),
//   		awscdk.IdentitySource.queryString(jsii.String("allow")),
//   	},
//   })
//
//   restapi.root.addMethod(jsii.String("ANY"), awscdk.NewMockIntegration(&integrationOptions{
//   	integrationResponses: []integrationResponse{
//   		&integrationResponse{
//   			statusCode: jsii.String("200"),
//   		},
//   	},
//   	passthroughBehavior: awscdk.PassthroughBehavior_NEVER,
//   	requestTemplates: map[string]*string{
//   		"application/json": jsii.String("{ \"statusCode\": 200 }"),
//   	},
//   }), &methodOptions{
//   	methodResponses: []methodResponse{
//   		&methodResponse{
//   			statusCode: jsii.String("200"),
//   		},
//   	},
//   	authorizer: authorizer,
//   })
//
type PassthroughBehavior string

const (
	// Passes the request body for unmapped content types through to the integration back end without transformation.
	PassthroughBehavior_WHEN_NO_MATCH PassthroughBehavior = "WHEN_NO_MATCH"
	// Rejects unmapped content types with an HTTP 415 'Unsupported Media Type' response.
	PassthroughBehavior_NEVER PassthroughBehavior = "NEVER"
	// Allows pass-through when the integration has NO content types mapped to templates.
	//
	// However if there is at least one content type defined,
	// unmapped content types will be rejected with the same 415 response.
	PassthroughBehavior_WHEN_NO_TEMPLATES PassthroughBehavior = "WHEN_NO_TEMPLATES"
)

// Time period for which quota settings apply.
//
// Example:
//   var api restApi
//
//
//   key := apigateway.NewRateLimitedApiKey(this, jsii.String("rate-limited-api-key"), &rateLimitedApiKeyProps{
//   	customerId: jsii.String("hello-customer"),
//   	resources: []iRestApi{
//   		api,
//   	},
//   	quota: &quotaSettings{
//   		limit: jsii.Number(10000),
//   		period: apigateway.period_MONTH,
//   	},
//   })
//
type Period string

const (
	Period_DAY Period = "DAY"
	Period_WEEK Period = "WEEK"
	Period_MONTH Period = "MONTH"
)

// Defines a {proxy+} greedy resource and an ANY method on a route.
//
// Example:
//   var resource resource
//   var handler function
//
//   proxy := resource.addProxy(&proxyResourceOptions{
//   	defaultIntegration: apigateway.NewLambdaIntegration(handler),
//
//   	// "false" will require explicitly adding methods on the `proxy` resource
//   	anyMethod: jsii.Boolean(true),
//   })
//
// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-set-up-simple-proxy.html
//
type ProxyResource interface {
	Resource
	// If `props.anyMethod` is `true`, this will be the reference to the 'ANY' method associated with this proxy resource.
	AnyMethod() Method
	// The rest API that this resource is part of.
	//
	// The reason we need the RestApi object itself and not just the ID is because the model
	// is being tracked by the top-level RestApi object for the purpose of calculating it's
	// hash to determine the ID of the deployment. This allows us to automatically update
	// the deployment when the model of the REST API changes.
	Api() IRestApi
	// Default options for CORS preflight OPTIONS method.
	DefaultCorsPreflightOptions() *CorsOptions
	// An integration to use as a default for all methods created within this API unless an integration is specified.
	DefaultIntegration() Integration
	// Method options to use as a default for all methods created within this API unless custom options are specified.
	DefaultMethodOptions() *MethodOptions
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	Node() constructs.Node
	// The parent of this resource or undefined for the root resource.
	ParentResource() IResource
	// The full path of this resource.
	Path() *string
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The ID of the resource.
	ResourceId() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Adds an OPTIONS method to this resource which responds to Cross-Origin Resource Sharing (CORS) preflight requests.
	//
	// Cross-Origin Resource Sharing (CORS) is a mechanism that uses additional
	// HTTP headers to tell browsers to give a web application running at one
	// origin, access to selected resources from a different origin. A web
	// application executes a cross-origin HTTP request when it requests a
	// resource that has a different origin (domain, protocol, or port) from its
	// own.
	AddCorsPreflight(options *CorsOptions) Method
	// Defines a new method for this resource.
	AddMethod(httpMethod *string, integration Integration, options *MethodOptions) Method
	// Adds a greedy proxy resource ("{proxy+}") and an ANY method to this route.
	AddProxy(options *ProxyResourceOptions) ProxyResource
	// Defines a new child resource where this resource is the parent.
	AddResource(pathPart *string, options *ResourceOptions) Resource
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Retrieves a child resource by path part.
	GetResource(pathPart *string) IResource
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Gets or create all resources leading up to the specified path.
	//
	// - Path may only start with "/" if this method is called on the root resource.
	// - All resources are created using default options.
	ResourceForPath(path *string) Resource
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for ProxyResource
type jsiiProxy_ProxyResource struct {
	jsiiProxy_Resource
}

func (j *jsiiProxy_ProxyResource) AnyMethod() Method {
	var returns Method
	_jsii_.Get(
		j,
		"anyMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProxyResource) Api() IRestApi {
	var returns IRestApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProxyResource) DefaultCorsPreflightOptions() *CorsOptions {
	var returns *CorsOptions
	_jsii_.Get(
		j,
		"defaultCorsPreflightOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProxyResource) DefaultIntegration() Integration {
	var returns Integration
	_jsii_.Get(
		j,
		"defaultIntegration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProxyResource) DefaultMethodOptions() *MethodOptions {
	var returns *MethodOptions
	_jsii_.Get(
		j,
		"defaultMethodOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProxyResource) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProxyResource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProxyResource) ParentResource() IResource {
	var returns IResource
	_jsii_.Get(
		j,
		"parentResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProxyResource) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProxyResource) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProxyResource) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProxyResource) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewProxyResource(scope constructs.Construct, id *string, props *ProxyResourceProps) ProxyResource {
	_init_.Initialize()

	j := jsiiProxy_ProxyResource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.ProxyResource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewProxyResource_Override(p ProxyResource, scope constructs.Construct, id *string, props *ProxyResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.ProxyResource",
		[]interface{}{scope, id, props},
		p,
	)
}

// Import an existing resource.
func ProxyResource_FromResourceAttributes(scope constructs.Construct, id *string, attrs *ResourceAttributes) IResource {
	_init_.Initialize()

	var returns IResource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.ProxyResource",
		"fromResourceAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func ProxyResource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.ProxyResource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func ProxyResource_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.ProxyResource",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ProxyResource_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.ProxyResource",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProxyResource) AddCorsPreflight(options *CorsOptions) Method {
	var returns Method

	_jsii_.Invoke(
		p,
		"addCorsPreflight",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProxyResource) AddMethod(httpMethod *string, integration Integration, options *MethodOptions) Method {
	var returns Method

	_jsii_.Invoke(
		p,
		"addMethod",
		[]interface{}{httpMethod, integration, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProxyResource) AddProxy(options *ProxyResourceOptions) ProxyResource {
	var returns ProxyResource

	_jsii_.Invoke(
		p,
		"addProxy",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProxyResource) AddResource(pathPart *string, options *ResourceOptions) Resource {
	var returns Resource

	_jsii_.Invoke(
		p,
		"addResource",
		[]interface{}{pathPart, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProxyResource) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (p *jsiiProxy_ProxyResource) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProxyResource) GetResource(pathPart *string) IResource {
	var returns IResource

	_jsii_.Invoke(
		p,
		"getResource",
		[]interface{}{pathPart},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProxyResource) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProxyResource) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProxyResource) ResourceForPath(path *string) Resource {
	var returns Resource

	_jsii_.Invoke(
		p,
		"resourceForPath",
		[]interface{}{path},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProxyResource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Example:
//   var resource resource
//   var handler function
//
//   proxy := resource.addProxy(&proxyResourceOptions{
//   	defaultIntegration: apigateway.NewLambdaIntegration(handler),
//
//   	// "false" will require explicitly adding methods on the `proxy` resource
//   	anyMethod: jsii.Boolean(true),
//   })
//
type ProxyResourceOptions struct {
	// Adds a CORS preflight OPTIONS method to this resource and all child resources.
	//
	// You can add CORS at the resource-level using `addCorsPreflight`.
	DefaultCorsPreflightOptions *CorsOptions `field:"optional" json:"defaultCorsPreflightOptions" yaml:"defaultCorsPreflightOptions"`
	// An integration to use as a default for all methods created within this API unless an integration is specified.
	DefaultIntegration Integration `field:"optional" json:"defaultIntegration" yaml:"defaultIntegration"`
	// Method options to use as a default for all methods created within this API unless custom options are specified.
	DefaultMethodOptions *MethodOptions `field:"optional" json:"defaultMethodOptions" yaml:"defaultMethodOptions"`
	// Adds an "ANY" method to this resource.
	//
	// If set to `false`, you will have to explicitly
	// add methods to this resource after it's created.
	AnyMethod *bool `field:"optional" json:"anyMethod" yaml:"anyMethod"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var authorizer authorizer
//   var integration integration
//   var model model
//   var requestValidator requestValidator
//   var resource resource
//
//   proxyResourceProps := &proxyResourceProps{
//   	parent: resource,
//
//   	// the properties below are optional
//   	anyMethod: jsii.Boolean(false),
//   	defaultCorsPreflightOptions: &corsOptions{
//   		allowOrigins: []*string{
//   			jsii.String("allowOrigins"),
//   		},
//
//   		// the properties below are optional
//   		allowCredentials: jsii.Boolean(false),
//   		allowHeaders: []*string{
//   			jsii.String("allowHeaders"),
//   		},
//   		allowMethods: []*string{
//   			jsii.String("allowMethods"),
//   		},
//   		disableCache: jsii.Boolean(false),
//   		exposeHeaders: []*string{
//   			jsii.String("exposeHeaders"),
//   		},
//   		maxAge: cdk.duration.minutes(jsii.Number(30)),
//   		statusCode: jsii.Number(123),
//   	},
//   	defaultIntegration: integration,
//   	defaultMethodOptions: &methodOptions{
//   		apiKeyRequired: jsii.Boolean(false),
//   		authorizationScopes: []*string{
//   			jsii.String("authorizationScopes"),
//   		},
//   		authorizationType: awscdk.Aws_apigateway.authorizationType_NONE,
//   		authorizer: authorizer,
//   		methodResponses: []methodResponse{
//   			&methodResponse{
//   				statusCode: jsii.String("statusCode"),
//
//   				// the properties below are optional
//   				responseModels: map[string]iModel{
//   					"responseModelsKey": model,
//   				},
//   				responseParameters: map[string]*bool{
//   					"responseParametersKey": jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		operationName: jsii.String("operationName"),
//   		requestModels: map[string]*iModel{
//   			"requestModelsKey": model,
//   		},
//   		requestParameters: map[string]*bool{
//   			"requestParametersKey": jsii.Boolean(false),
//   		},
//   		requestValidator: requestValidator,
//   		requestValidatorOptions: &requestValidatorOptions{
//   			requestValidatorName: jsii.String("requestValidatorName"),
//   			validateRequestBody: jsii.Boolean(false),
//   			validateRequestParameters: jsii.Boolean(false),
//   		},
//   	},
//   }
//
type ProxyResourceProps struct {
	// Adds a CORS preflight OPTIONS method to this resource and all child resources.
	//
	// You can add CORS at the resource-level using `addCorsPreflight`.
	DefaultCorsPreflightOptions *CorsOptions `field:"optional" json:"defaultCorsPreflightOptions" yaml:"defaultCorsPreflightOptions"`
	// An integration to use as a default for all methods created within this API unless an integration is specified.
	DefaultIntegration Integration `field:"optional" json:"defaultIntegration" yaml:"defaultIntegration"`
	// Method options to use as a default for all methods created within this API unless custom options are specified.
	DefaultMethodOptions *MethodOptions `field:"optional" json:"defaultMethodOptions" yaml:"defaultMethodOptions"`
	// Adds an "ANY" method to this resource.
	//
	// If set to `false`, you will have to explicitly
	// add methods to this resource after it's created.
	AnyMethod *bool `field:"optional" json:"anyMethod" yaml:"anyMethod"`
	// The parent resource of this resource.
	//
	// You can either pass another
	// `Resource` object or a `RestApi` object here.
	Parent IResource `field:"required" json:"parent" yaml:"parent"`
}

// Specifies the maximum number of requests that clients can make to API Gateway APIs.
//
// Example:
//   var api restApi
//
//
//   key := apigateway.NewRateLimitedApiKey(this, jsii.String("rate-limited-api-key"), &rateLimitedApiKeyProps{
//   	customerId: jsii.String("hello-customer"),
//   	resources: []iRestApi{
//   		api,
//   	},
//   	quota: &quotaSettings{
//   		limit: jsii.Number(10000),
//   		period: apigateway.period_MONTH,
//   	},
//   })
//
type QuotaSettings struct {
	// The maximum number of requests that users can make within the specified time period.
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// For the initial time period, the number of requests to subtract from the specified limit.
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
	// The time period for which the maximum limit of requests applies.
	Period Period `field:"optional" json:"period" yaml:"period"`
}

// An API Gateway ApiKey, for which a rate limiting configuration can be specified.
//
// Example:
//   var api restApi
//
//
//   key := apigateway.NewRateLimitedApiKey(this, jsii.String("rate-limited-api-key"), &rateLimitedApiKeyProps{
//   	customerId: jsii.String("hello-customer"),
//   	resources: []iRestApi{
//   		api,
//   	},
//   	quota: &quotaSettings{
//   		limit: jsii.Number(10000),
//   		period: apigateway.period_MONTH,
//   	},
//   })
//
type RateLimitedApiKey interface {
	awscdk.Resource
	IApiKey
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The API key ARN.
	KeyArn() *string
	// The API key ID.
	KeyId() *string
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Permits the IAM principal all read operations through this key.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Permits the IAM principal all read and write operations through this key.
	GrantReadWrite(grantee awsiam.IGrantable) awsiam.Grant
	// Permits the IAM principal all write operations through this key.
	GrantWrite(grantee awsiam.IGrantable) awsiam.Grant
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for RateLimitedApiKey
type jsiiProxy_RateLimitedApiKey struct {
	internal.Type__awscdkResource
	jsiiProxy_IApiKey
}

func (j *jsiiProxy_RateLimitedApiKey) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RateLimitedApiKey) KeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RateLimitedApiKey) KeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RateLimitedApiKey) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RateLimitedApiKey) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RateLimitedApiKey) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewRateLimitedApiKey(scope constructs.Construct, id *string, props *RateLimitedApiKeyProps) RateLimitedApiKey {
	_init_.Initialize()

	j := jsiiProxy_RateLimitedApiKey{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.RateLimitedApiKey",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewRateLimitedApiKey_Override(r RateLimitedApiKey, scope constructs.Construct, id *string, props *RateLimitedApiKeyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.RateLimitedApiKey",
		[]interface{}{scope, id, props},
		r,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func RateLimitedApiKey_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.RateLimitedApiKey",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func RateLimitedApiKey_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.RateLimitedApiKey",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func RateLimitedApiKey_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.RateLimitedApiKey",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RateLimitedApiKey) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_RateLimitedApiKey) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RateLimitedApiKey) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RateLimitedApiKey) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RateLimitedApiKey) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		r,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RateLimitedApiKey) GrantReadWrite(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		r,
		"grantReadWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RateLimitedApiKey) GrantWrite(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		r,
		"grantWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RateLimitedApiKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// RateLimitedApiKey properties.
//
// Example:
//   var api restApi
//
//
//   key := apigateway.NewRateLimitedApiKey(this, jsii.String("rate-limited-api-key"), &rateLimitedApiKeyProps{
//   	customerId: jsii.String("hello-customer"),
//   	resources: []iRestApi{
//   		api,
//   	},
//   	quota: &quotaSettings{
//   		limit: jsii.Number(10000),
//   		period: apigateway.period_MONTH,
//   	},
//   })
//
type RateLimitedApiKeyProps struct {
	// Adds a CORS preflight OPTIONS method to this resource and all child resources.
	//
	// You can add CORS at the resource-level using `addCorsPreflight`.
	DefaultCorsPreflightOptions *CorsOptions `field:"optional" json:"defaultCorsPreflightOptions" yaml:"defaultCorsPreflightOptions"`
	// An integration to use as a default for all methods created within this API unless an integration is specified.
	DefaultIntegration Integration `field:"optional" json:"defaultIntegration" yaml:"defaultIntegration"`
	// Method options to use as a default for all methods created within this API unless custom options are specified.
	DefaultMethodOptions *MethodOptions `field:"optional" json:"defaultMethodOptions" yaml:"defaultMethodOptions"`
	// A name for the API key.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the API key name.
	ApiKeyName *string `field:"optional" json:"apiKeyName" yaml:"apiKeyName"`
	// A description of the purpose of the API key.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The value of the API key.
	//
	// Must be at least 20 characters long.
	Value *string `field:"optional" json:"value" yaml:"value"`
	// An AWS Marketplace customer identifier to use when integrating with the AWS SaaS Marketplace.
	CustomerId *string `field:"optional" json:"customerId" yaml:"customerId"`
	// Indicates whether the API key can be used by clients.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Specifies whether the key identifier is distinct from the created API key value.
	GenerateDistinctId *bool `field:"optional" json:"generateDistinctId" yaml:"generateDistinctId"`
	// A list of resources this api key is associated with.
	Resources *[]IRestApi `field:"optional" json:"resources" yaml:"resources"`
	// API Stages to be associated with the RateLimitedApiKey.
	ApiStages *[]*UsagePlanPerApiStage `field:"optional" json:"apiStages" yaml:"apiStages"`
	// Number of requests clients can make in a given time period.
	Quota *QuotaSettings `field:"optional" json:"quota" yaml:"quota"`
	// Overall throttle settings for the API.
	Throttle *ThrottleSettings `field:"optional" json:"throttle" yaml:"throttle"`
}

// Request-based lambda authorizer that recognizes the caller's identity via request parameters, such as headers, paths, query strings, stage variables, or context variables.
//
// Based on the request, authorization is performed by a lambda function.
//
// Example:
//   var authFn function
//   var books resource
//
//
//   auth := apigateway.NewRequestAuthorizer(this, jsii.String("booksAuthorizer"), &requestAuthorizerProps{
//   	handler: authFn,
//   	identitySources: []*string{
//   		apigateway.identitySource.header(jsii.String("Authorization")),
//   	},
//   })
//
//   books.addMethod(jsii.String("GET"), apigateway.NewHttpIntegration(jsii.String("http://amazon.com")), &methodOptions{
//   	authorizer: auth,
//   })
//
type RequestAuthorizer interface {
	Authorizer
	IAuthorizer
	// The authorization type of this authorizer.
	AuthorizationType() AuthorizationType
	// The ARN of the authorizer to be used in permission policies, such as IAM and resource-based grants.
	AuthorizerArn() *string
	// The id of the authorizer.
	AuthorizerId() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The Lambda function handler that this authorizer uses.
	Handler() awslambda.IFunction
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	RestApiId() *string
	SetRestApiId(val *string)
	// The IAM role that the API Gateway service assumes while invoking the Lambda function.
	Role() awsiam.IRole
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a token that resolves to the Rest Api Id at the time of synthesis.
	//
	// Throws an error, during token resolution, if no RestApi is attached to this authorizer.
	LazyRestApiId() *string
	// Sets up the permissions necessary for the API Gateway service to invoke the Lambda function.
	SetupPermissions()
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for RequestAuthorizer
type jsiiProxy_RequestAuthorizer struct {
	jsiiProxy_Authorizer
	jsiiProxy_IAuthorizer
}

func (j *jsiiProxy_RequestAuthorizer) AuthorizationType() AuthorizationType {
	var returns AuthorizationType
	_jsii_.Get(
		j,
		"authorizationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequestAuthorizer) AuthorizerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequestAuthorizer) AuthorizerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequestAuthorizer) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequestAuthorizer) Handler() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequestAuthorizer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequestAuthorizer) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequestAuthorizer) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequestAuthorizer) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequestAuthorizer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewRequestAuthorizer(scope constructs.Construct, id *string, props *RequestAuthorizerProps) RequestAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_RequestAuthorizer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.RequestAuthorizer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewRequestAuthorizer_Override(r RequestAuthorizer, scope constructs.Construct, id *string, props *RequestAuthorizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.RequestAuthorizer",
		[]interface{}{scope, id, props},
		r,
	)
}

func (j *jsiiProxy_RequestAuthorizer) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

// Return whether the given object is an Authorizer.
func RequestAuthorizer_IsAuthorizer(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.RequestAuthorizer",
		"isAuthorizer",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func RequestAuthorizer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.RequestAuthorizer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func RequestAuthorizer_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.RequestAuthorizer",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func RequestAuthorizer_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.RequestAuthorizer",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RequestAuthorizer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_RequestAuthorizer) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RequestAuthorizer) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RequestAuthorizer) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RequestAuthorizer) LazyRestApiId() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"lazyRestApiId",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RequestAuthorizer) SetupPermissions() {
	_jsii_.InvokeVoid(
		r,
		"setupPermissions",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RequestAuthorizer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for RequestAuthorizer.
//
// Example:
//   var authFn function
//   var books resource
//
//
//   auth := apigateway.NewRequestAuthorizer(this, jsii.String("booksAuthorizer"), &requestAuthorizerProps{
//   	handler: authFn,
//   	identitySources: []*string{
//   		apigateway.identitySource.header(jsii.String("Authorization")),
//   	},
//   })
//
//   books.addMethod(jsii.String("GET"), apigateway.NewHttpIntegration(jsii.String("http://amazon.com")), &methodOptions{
//   	authorizer: auth,
//   })
//
type RequestAuthorizerProps struct {
	// The handler for the authorizer lambda function.
	//
	// The handler must follow a very specific protocol on the input it receives and the output it needs to produce.
	// API Gateway has documented the handler's input specification
	// {@link https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-lambda-authorizer-input.html | here} and output specification
	// {@link https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-lambda-authorizer-output.html | here}.
	Handler awslambda.IFunction `field:"required" json:"handler" yaml:"handler"`
	// An optional IAM role for APIGateway to assume before calling the Lambda-based authorizer.
	//
	// The IAM role must be
	// assumable by 'apigateway.amazonaws.com'.
	AssumeRole awsiam.IRole `field:"optional" json:"assumeRole" yaml:"assumeRole"`
	// An optional human friendly name for the authorizer.
	//
	// Note that, this is not the primary identifier of the authorizer.
	AuthorizerName *string `field:"optional" json:"authorizerName" yaml:"authorizerName"`
	// How long APIGateway should cache the results.
	//
	// Max 1 hour.
	// Disable caching by setting this to 0.
	ResultsCacheTtl awscdk.Duration `field:"optional" json:"resultsCacheTtl" yaml:"resultsCacheTtl"`
	// An array of request header mapping expressions for identities.
	//
	// Supported parameter types are
	// Header, Query String, Stage Variable, and Context. For instance, extracting an authorization
	// token from a header would use the identity source `IdentitySource.header('Authorizer')`.
	//
	// Note: API Gateway uses the specified identity sources as the request authorizer caching key. When caching is
	// enabled, API Gateway calls the authorizer's Lambda function only after successfully verifying that all the
	// specified identity sources are present at runtime. If a specified identify source is missing, null, or empty,
	// API Gateway returns a 401 Unauthorized response without calling the authorizer Lambda function.
	// See: https://docs.aws.amazon.com/apigateway/api-reference/link-relation/authorizer-create/#identitySource
	//
	IdentitySources *[]*string `field:"required" json:"identitySources" yaml:"identitySources"`
}

// Configure what must be included in the `requestContext`.
//
// More details can be found at mapping templates documentation.
//
// Example:
//   apigateway.NewStepFunctionsRestApi(this, jsii.String("StepFunctionsRestApi"), &stepFunctionsRestApiProps{
//   	stateMachine: machine,
//   	headers: jsii.Boolean(true),
//   	path: jsii.Boolean(false),
//   	querystring: jsii.Boolean(false),
//   	authorizer: jsii.Boolean(false),
//   	requestContext: &requestContext{
//   		caller: jsii.Boolean(true),
//   		user: jsii.Boolean(true),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html
//
type RequestContext struct {
	// Represents the information of $context.identity.accountId.
	//
	// Whether the AWS account of the API owner should be included in the request context.
	AccountId *bool `field:"optional" json:"accountId" yaml:"accountId"`
	// Represents the information of $context.apiId.
	//
	// Whether the identifier API Gateway assigns to your API should be included in the request context.
	ApiId *bool `field:"optional" json:"apiId" yaml:"apiId"`
	// Represents the information of $context.identity.apiKey.
	//
	// Whether the API key associated with the request should be included in request context.
	ApiKey *bool `field:"optional" json:"apiKey" yaml:"apiKey"`
	// Represents the information of $context.authorizer.principalId.
	//
	// Whether the principal user identifier associated with the token sent by the client and returned
	// from an API Gateway Lambda authorizer should be included in the request context.
	AuthorizerPrincipalId *bool `field:"optional" json:"authorizerPrincipalId" yaml:"authorizerPrincipalId"`
	// Represents the information of $context.identity.caller.
	//
	// Whether the principal identifier of the caller that signed the request should be included in the request context.
	// Supported for resources that use IAM authorization.
	Caller *bool `field:"optional" json:"caller" yaml:"caller"`
	// Represents the information of $context.identity.cognitoAuthenticationProvider.
	//
	// Whether the list of the Amazon Cognito authentication providers used by the caller making the request should be included in the request context.
	// Available only if the request was signed with Amazon Cognito credentials.
	CognitoAuthenticationProvider *bool `field:"optional" json:"cognitoAuthenticationProvider" yaml:"cognitoAuthenticationProvider"`
	// Represents the information of $context.identity.cognitoAuthenticationType.
	//
	// Whether the Amazon Cognito authentication type of the caller making the request should be included in the request context.
	// Available only if the request was signed with Amazon Cognito credentials.
	// Possible values include authenticated for authenticated identities and unauthenticated for unauthenticated identities.
	CognitoAuthenticationType *bool `field:"optional" json:"cognitoAuthenticationType" yaml:"cognitoAuthenticationType"`
	// Represents the information of $context.identity.cognitoIdentityId.
	//
	// Whether the Amazon Cognito identity ID of the caller making the request should be included in the request context.
	// Available only if the request was signed with Amazon Cognito credentials.
	CognitoIdentityId *bool `field:"optional" json:"cognitoIdentityId" yaml:"cognitoIdentityId"`
	// Represents the information of $context.identity.cognitoIdentityPoolId.
	//
	// Whether the Amazon Cognito identity pool ID of the caller making the request should be included in the request context.
	// Available only if the request was signed with Amazon Cognito credentials.
	CognitoIdentityPoolId *bool `field:"optional" json:"cognitoIdentityPoolId" yaml:"cognitoIdentityPoolId"`
	// Represents the information of $context.httpMethod.
	//
	// Whether the HTTP method used should be included in the request context.
	// Valid values include: DELETE, GET, HEAD, OPTIONS, PATCH, POST, and PUT.
	HttpMethod *bool `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// Represents the information of $context.requestId.
	//
	// Whether the ID for the request should be included in the request context.
	RequestId *bool `field:"optional" json:"requestId" yaml:"requestId"`
	// Represents the information of $context.resourceId.
	//
	// Whether the identifier that API Gateway assigns to your resource should be included in the request context.
	ResourceId *bool `field:"optional" json:"resourceId" yaml:"resourceId"`
	// Represents the information of $context.resourcePath.
	//
	// Whether the path to the resource should be included in the request context.
	ResourcePath *bool `field:"optional" json:"resourcePath" yaml:"resourcePath"`
	// Represents the information of $context.identity.sourceIp.
	//
	// Whether the source IP address of the immediate TCP connection making the request
	// to API Gateway endpoint should be included in the request context.
	SourceIp *bool `field:"optional" json:"sourceIp" yaml:"sourceIp"`
	// Represents the information of $context.stage.
	//
	// Whether the deployment stage of the API request should be included in the request context.
	Stage *bool `field:"optional" json:"stage" yaml:"stage"`
	// Represents the information of $context.identity.user.
	//
	// Whether the principal identifier of the user that will be authorized should be included in the request context.
	// Supported for resources that use IAM authorization.
	User *bool `field:"optional" json:"user" yaml:"user"`
	// Represents the information of $context.identity.userAgent.
	//
	// Whether the User-Agent header of the API caller should be included in the request context.
	UserAgent *bool `field:"optional" json:"userAgent" yaml:"userAgent"`
	// Represents the information of $context.identity.userArn.
	//
	// Whether the Amazon Resource Name (ARN) of the effective user identified after authentication should be included in the request context.
	UserArn *bool `field:"optional" json:"userArn" yaml:"userArn"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var restApi restApi
//
//   requestValidator := awscdk.Aws_apigateway.NewRequestValidator(this, jsii.String("MyRequestValidator"), &requestValidatorProps{
//   	restApi: restApi,
//
//   	// the properties below are optional
//   	requestValidatorName: jsii.String("requestValidatorName"),
//   	validateRequestBody: jsii.Boolean(false),
//   	validateRequestParameters: jsii.Boolean(false),
//   })
//
type RequestValidator interface {
	awscdk.Resource
	IRequestValidator
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// ID of the request validator, such as abc123.
	RequestValidatorId() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for RequestValidator
type jsiiProxy_RequestValidator struct {
	internal.Type__awscdkResource
	jsiiProxy_IRequestValidator
}

func (j *jsiiProxy_RequestValidator) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequestValidator) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequestValidator) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequestValidator) RequestValidatorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestValidatorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequestValidator) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewRequestValidator(scope constructs.Construct, id *string, props *RequestValidatorProps) RequestValidator {
	_init_.Initialize()

	j := jsiiProxy_RequestValidator{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.RequestValidator",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewRequestValidator_Override(r RequestValidator, scope constructs.Construct, id *string, props *RequestValidatorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.RequestValidator",
		[]interface{}{scope, id, props},
		r,
	)
}

func RequestValidator_FromRequestValidatorId(scope constructs.Construct, id *string, requestValidatorId *string) IRequestValidator {
	_init_.Initialize()

	var returns IRequestValidator

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.RequestValidator",
		"fromRequestValidatorId",
		[]interface{}{scope, id, requestValidatorId},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func RequestValidator_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.RequestValidator",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func RequestValidator_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.RequestValidator",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func RequestValidator_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.RequestValidator",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RequestValidator) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_RequestValidator) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RequestValidator) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RequestValidator) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RequestValidator) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Example:
//   var integration lambdaIntegration
//   var resource resource
//   var responseModel model
//   var errorResponseModel model
//
//
//   resource.addMethod(jsii.String("GET"), integration, &methodOptions{
//   	// We can mark the parameters as required
//   	requestParameters: map[string]*bool{
//   		"method.request.querystring.who": jsii.Boolean(true),
//   	},
//   	// we can set request validator options like below
//   	requestValidatorOptions: &requestValidatorOptions{
//   		requestValidatorName: jsii.String("test-validator"),
//   		validateRequestBody: jsii.Boolean(true),
//   		validateRequestParameters: jsii.Boolean(false),
//   	},
//   	methodResponses: []methodResponse{
//   		&methodResponse{
//   			// Successful response from the integration
//   			statusCode: jsii.String("200"),
//   			// Define what parameters are allowed or not
//   			responseParameters: map[string]*bool{
//   				"method.response.header.Content-Type": jsii.Boolean(true),
//   				"method.response.header.Access-Control-Allow-Origin": jsii.Boolean(true),
//   				"method.response.header.Access-Control-Allow-Credentials": jsii.Boolean(true),
//   			},
//   			// Validate the schema on the response
//   			responseModels: map[string]iModel{
//   				"application/json": responseModel,
//   			},
//   		},
//   		&methodResponse{
//   			// Same thing for the error responses
//   			statusCode: jsii.String("400"),
//   			responseParameters: map[string]*bool{
//   				"method.response.header.Content-Type": jsii.Boolean(true),
//   				"method.response.header.Access-Control-Allow-Origin": jsii.Boolean(true),
//   				"method.response.header.Access-Control-Allow-Credentials": jsii.Boolean(true),
//   			},
//   			responseModels: map[string]*iModel{
//   				"application/json": errorResponseModel,
//   			},
//   		},
//   	},
//   })
//
type RequestValidatorOptions struct {
	// The name of this request validator.
	RequestValidatorName *string `field:"optional" json:"requestValidatorName" yaml:"requestValidatorName"`
	// Indicates whether to validate the request body according to the configured schema for the targeted API and method.
	ValidateRequestBody *bool `field:"optional" json:"validateRequestBody" yaml:"validateRequestBody"`
	// Indicates whether to validate request parameters.
	ValidateRequestParameters *bool `field:"optional" json:"validateRequestParameters" yaml:"validateRequestParameters"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var restApi restApi
//
//   requestValidatorProps := &requestValidatorProps{
//   	restApi: restApi,
//
//   	// the properties below are optional
//   	requestValidatorName: jsii.String("requestValidatorName"),
//   	validateRequestBody: jsii.Boolean(false),
//   	validateRequestParameters: jsii.Boolean(false),
//   }
//
type RequestValidatorProps struct {
	// The name of this request validator.
	RequestValidatorName *string `field:"optional" json:"requestValidatorName" yaml:"requestValidatorName"`
	// Indicates whether to validate the request body according to the configured schema for the targeted API and method.
	ValidateRequestBody *bool `field:"optional" json:"validateRequestBody" yaml:"validateRequestBody"`
	// Indicates whether to validate request parameters.
	ValidateRequestParameters *bool `field:"optional" json:"validateRequestParameters" yaml:"validateRequestParameters"`
	// The rest API that this model is part of.
	//
	// The reason we need the RestApi object itself and not just the ID is because the model
	// is being tracked by the top-level RestApi object for the purpose of calculating it's
	// hash to determine the ID of the deployment. This allows us to automatically update
	// the deployment when the model of the REST API changes.
	RestApi IRestApi `field:"required" json:"restApi" yaml:"restApi"`
}

// Example:
//   var booksBackend lambdaIntegration
//
//   api := apigateway.NewRestApi(this, jsii.String("books"), &restApiProps{
//   	defaultIntegration: booksBackend,
//   })
//
//   books := api.root.addResource(jsii.String("books"))
//   books.addMethod(jsii.String("GET")) // integrated with `booksBackend`
//   books.addMethod(jsii.String("POST")) // integrated with `booksBackend`
//
//   book := books.addResource(jsii.String("{book_id}"))
//   book.addMethod(jsii.String("GET"))
//
type Resource interface {
	ResourceBase
	// The rest API that this resource is part of.
	//
	// The reason we need the RestApi object itself and not just the ID is because the model
	// is being tracked by the top-level RestApi object for the purpose of calculating it's
	// hash to determine the ID of the deployment. This allows us to automatically update
	// the deployment when the model of the REST API changes.
	Api() IRestApi
	// Default options for CORS preflight OPTIONS method.
	DefaultCorsPreflightOptions() *CorsOptions
	// An integration to use as a default for all methods created within this API unless an integration is specified.
	DefaultIntegration() Integration
	// Method options to use as a default for all methods created within this API unless custom options are specified.
	DefaultMethodOptions() *MethodOptions
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	Node() constructs.Node
	// The parent of this resource or undefined for the root resource.
	ParentResource() IResource
	// The full path of this resource.
	Path() *string
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The ID of the resource.
	ResourceId() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Adds an OPTIONS method to this resource which responds to Cross-Origin Resource Sharing (CORS) preflight requests.
	//
	// Cross-Origin Resource Sharing (CORS) is a mechanism that uses additional
	// HTTP headers to tell browsers to give a web application running at one
	// origin, access to selected resources from a different origin. A web
	// application executes a cross-origin HTTP request when it requests a
	// resource that has a different origin (domain, protocol, or port) from its
	// own.
	AddCorsPreflight(options *CorsOptions) Method
	// Defines a new method for this resource.
	AddMethod(httpMethod *string, integration Integration, options *MethodOptions) Method
	// Adds a greedy proxy resource ("{proxy+}") and an ANY method to this route.
	AddProxy(options *ProxyResourceOptions) ProxyResource
	// Defines a new child resource where this resource is the parent.
	AddResource(pathPart *string, options *ResourceOptions) Resource
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Retrieves a child resource by path part.
	GetResource(pathPart *string) IResource
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Gets or create all resources leading up to the specified path.
	//
	// - Path may only start with "/" if this method is called on the root resource.
	// - All resources are created using default options.
	ResourceForPath(path *string) Resource
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Resource
type jsiiProxy_Resource struct {
	jsiiProxy_ResourceBase
}

func (j *jsiiProxy_Resource) Api() IRestApi {
	var returns IRestApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) DefaultCorsPreflightOptions() *CorsOptions {
	var returns *CorsOptions
	_jsii_.Get(
		j,
		"defaultCorsPreflightOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) DefaultIntegration() Integration {
	var returns Integration
	_jsii_.Get(
		j,
		"defaultIntegration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) DefaultMethodOptions() *MethodOptions {
	var returns *MethodOptions
	_jsii_.Get(
		j,
		"defaultMethodOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) ParentResource() IResource {
	var returns IResource
	_jsii_.Get(
		j,
		"parentResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewResource(scope constructs.Construct, id *string, props *ResourceProps) Resource {
	_init_.Initialize()

	j := jsiiProxy_Resource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.Resource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewResource_Override(r Resource, scope constructs.Construct, id *string, props *ResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.Resource",
		[]interface{}{scope, id, props},
		r,
	)
}

// Import an existing resource.
func Resource_FromResourceAttributes(scope constructs.Construct, id *string, attrs *ResourceAttributes) IResource {
	_init_.Initialize()

	var returns IResource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.Resource",
		"fromResourceAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Resource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.Resource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func Resource_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.Resource",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Resource_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.Resource",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) AddCorsPreflight(options *CorsOptions) Method {
	var returns Method

	_jsii_.Invoke(
		r,
		"addCorsPreflight",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) AddMethod(httpMethod *string, integration Integration, options *MethodOptions) Method {
	var returns Method

	_jsii_.Invoke(
		r,
		"addMethod",
		[]interface{}{httpMethod, integration, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) AddProxy(options *ProxyResourceOptions) ProxyResource {
	var returns ProxyResource

	_jsii_.Invoke(
		r,
		"addProxy",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) AddResource(pathPart *string, options *ResourceOptions) Resource {
	var returns Resource

	_jsii_.Invoke(
		r,
		"addResource",
		[]interface{}{pathPart, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_Resource) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) GetResource(pathPart *string) IResource {
	var returns IResource

	_jsii_.Invoke(
		r,
		"getResource",
		[]interface{}{pathPart},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) ResourceForPath(path *string) Resource {
	var returns Resource

	_jsii_.Invoke(
		r,
		"resourceForPath",
		[]interface{}{path},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Attributes that can be specified when importing a Resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var restApi restApi
//
//   resourceAttributes := &resourceAttributes{
//   	path: jsii.String("path"),
//   	resourceId: jsii.String("resourceId"),
//   	restApi: restApi,
//   }
//
type ResourceAttributes struct {
	// The full path of this resource.
	Path *string `field:"required" json:"path" yaml:"path"`
	// The ID of the resource.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// The rest API that this resource is part of.
	RestApi IRestApi `field:"required" json:"restApi" yaml:"restApi"`
}

type ResourceBase interface {
	awscdk.Resource
	IResource
	// The rest API that this resource is part of.
	//
	// The reason we need the RestApi object itself and not just the ID is because the model
	// is being tracked by the top-level RestApi object for the purpose of calculating it's
	// hash to determine the ID of the deployment. This allows us to automatically update
	// the deployment when the model of the REST API changes.
	Api() IRestApi
	// Default options for CORS preflight OPTIONS method.
	DefaultCorsPreflightOptions() *CorsOptions
	// An integration to use as a default for all methods created within this API unless an integration is specified.
	DefaultIntegration() Integration
	// Method options to use as a default for all methods created within this API unless custom options are specified.
	DefaultMethodOptions() *MethodOptions
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	Node() constructs.Node
	// The parent of this resource or undefined for the root resource.
	ParentResource() IResource
	// The full path of this resource.
	Path() *string
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The ID of the resource.
	ResourceId() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Adds an OPTIONS method to this resource which responds to Cross-Origin Resource Sharing (CORS) preflight requests.
	//
	// Cross-Origin Resource Sharing (CORS) is a mechanism that uses additional
	// HTTP headers to tell browsers to give a web application running at one
	// origin, access to selected resources from a different origin. A web
	// application executes a cross-origin HTTP request when it requests a
	// resource that has a different origin (domain, protocol, or port) from its
	// own.
	AddCorsPreflight(options *CorsOptions) Method
	// Defines a new method for this resource.
	AddMethod(httpMethod *string, integration Integration, options *MethodOptions) Method
	// Adds a greedy proxy resource ("{proxy+}") and an ANY method to this route.
	AddProxy(options *ProxyResourceOptions) ProxyResource
	// Defines a new child resource where this resource is the parent.
	AddResource(pathPart *string, options *ResourceOptions) Resource
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Retrieves a child resource by path part.
	GetResource(pathPart *string) IResource
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Gets or create all resources leading up to the specified path.
	//
	// - Path may only start with "/" if this method is called on the root resource.
	// - All resources are created using default options.
	ResourceForPath(path *string) Resource
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for ResourceBase
type jsiiProxy_ResourceBase struct {
	internal.Type__awscdkResource
	jsiiProxy_IResource
}

func (j *jsiiProxy_ResourceBase) Api() IRestApi {
	var returns IRestApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceBase) DefaultCorsPreflightOptions() *CorsOptions {
	var returns *CorsOptions
	_jsii_.Get(
		j,
		"defaultCorsPreflightOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceBase) DefaultIntegration() Integration {
	var returns Integration
	_jsii_.Get(
		j,
		"defaultIntegration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceBase) DefaultMethodOptions() *MethodOptions {
	var returns *MethodOptions
	_jsii_.Get(
		j,
		"defaultMethodOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceBase) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceBase) ParentResource() IResource {
	var returns IResource
	_jsii_.Get(
		j,
		"parentResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceBase) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceBase) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceBase) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewResourceBase_Override(r ResourceBase, scope constructs.Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.ResourceBase",
		[]interface{}{scope, id},
		r,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func ResourceBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.ResourceBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func ResourceBase_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.ResourceBase",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ResourceBase_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.ResourceBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceBase) AddCorsPreflight(options *CorsOptions) Method {
	var returns Method

	_jsii_.Invoke(
		r,
		"addCorsPreflight",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceBase) AddMethod(httpMethod *string, integration Integration, options *MethodOptions) Method {
	var returns Method

	_jsii_.Invoke(
		r,
		"addMethod",
		[]interface{}{httpMethod, integration, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceBase) AddProxy(options *ProxyResourceOptions) ProxyResource {
	var returns ProxyResource

	_jsii_.Invoke(
		r,
		"addProxy",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceBase) AddResource(pathPart *string, options *ResourceOptions) Resource {
	var returns Resource

	_jsii_.Invoke(
		r,
		"addResource",
		[]interface{}{pathPart, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_ResourceBase) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceBase) GetResource(pathPart *string) IResource {
	var returns IResource

	_jsii_.Invoke(
		r,
		"getResource",
		[]interface{}{pathPart},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceBase) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceBase) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceBase) ResourceForPath(path *string) Resource {
	var returns Resource

	_jsii_.Invoke(
		r,
		"resourceForPath",
		[]interface{}{path},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Example:
//   var resource resource
//
//
//   subtree := resource.addResource(jsii.String("subtree"), &resourceOptions{
//   	defaultCorsPreflightOptions: &corsOptions{
//   		allowOrigins: []*string{
//   			jsii.String("https://amazon.com"),
//   		},
//   	},
//   })
//
type ResourceOptions struct {
	// Adds a CORS preflight OPTIONS method to this resource and all child resources.
	//
	// You can add CORS at the resource-level using `addCorsPreflight`.
	DefaultCorsPreflightOptions *CorsOptions `field:"optional" json:"defaultCorsPreflightOptions" yaml:"defaultCorsPreflightOptions"`
	// An integration to use as a default for all methods created within this API unless an integration is specified.
	DefaultIntegration Integration `field:"optional" json:"defaultIntegration" yaml:"defaultIntegration"`
	// Method options to use as a default for all methods created within this API unless custom options are specified.
	DefaultMethodOptions *MethodOptions `field:"optional" json:"defaultMethodOptions" yaml:"defaultMethodOptions"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var authorizer authorizer
//   var integration integration
//   var model model
//   var requestValidator requestValidator
//   var resource resource
//
//   resourceProps := &resourceProps{
//   	parent: resource,
//   	pathPart: jsii.String("pathPart"),
//
//   	// the properties below are optional
//   	defaultCorsPreflightOptions: &corsOptions{
//   		allowOrigins: []*string{
//   			jsii.String("allowOrigins"),
//   		},
//
//   		// the properties below are optional
//   		allowCredentials: jsii.Boolean(false),
//   		allowHeaders: []*string{
//   			jsii.String("allowHeaders"),
//   		},
//   		allowMethods: []*string{
//   			jsii.String("allowMethods"),
//   		},
//   		disableCache: jsii.Boolean(false),
//   		exposeHeaders: []*string{
//   			jsii.String("exposeHeaders"),
//   		},
//   		maxAge: cdk.duration.minutes(jsii.Number(30)),
//   		statusCode: jsii.Number(123),
//   	},
//   	defaultIntegration: integration,
//   	defaultMethodOptions: &methodOptions{
//   		apiKeyRequired: jsii.Boolean(false),
//   		authorizationScopes: []*string{
//   			jsii.String("authorizationScopes"),
//   		},
//   		authorizationType: awscdk.Aws_apigateway.authorizationType_NONE,
//   		authorizer: authorizer,
//   		methodResponses: []methodResponse{
//   			&methodResponse{
//   				statusCode: jsii.String("statusCode"),
//
//   				// the properties below are optional
//   				responseModels: map[string]iModel{
//   					"responseModelsKey": model,
//   				},
//   				responseParameters: map[string]*bool{
//   					"responseParametersKey": jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		operationName: jsii.String("operationName"),
//   		requestModels: map[string]*iModel{
//   			"requestModelsKey": model,
//   		},
//   		requestParameters: map[string]*bool{
//   			"requestParametersKey": jsii.Boolean(false),
//   		},
//   		requestValidator: requestValidator,
//   		requestValidatorOptions: &requestValidatorOptions{
//   			requestValidatorName: jsii.String("requestValidatorName"),
//   			validateRequestBody: jsii.Boolean(false),
//   			validateRequestParameters: jsii.Boolean(false),
//   		},
//   	},
//   }
//
type ResourceProps struct {
	// Adds a CORS preflight OPTIONS method to this resource and all child resources.
	//
	// You can add CORS at the resource-level using `addCorsPreflight`.
	DefaultCorsPreflightOptions *CorsOptions `field:"optional" json:"defaultCorsPreflightOptions" yaml:"defaultCorsPreflightOptions"`
	// An integration to use as a default for all methods created within this API unless an integration is specified.
	DefaultIntegration Integration `field:"optional" json:"defaultIntegration" yaml:"defaultIntegration"`
	// Method options to use as a default for all methods created within this API unless custom options are specified.
	DefaultMethodOptions *MethodOptions `field:"optional" json:"defaultMethodOptions" yaml:"defaultMethodOptions"`
	// The parent resource of this resource.
	//
	// You can either pass another
	// `Resource` object or a `RestApi` object here.
	Parent IResource `field:"required" json:"parent" yaml:"parent"`
	// A path name for the resource.
	PathPart *string `field:"required" json:"pathPart" yaml:"pathPart"`
}

// Supported types of gateway responses.
//
// Example:
//   api := apigateway.NewRestApi(this, jsii.String("books-api"))
//   api.addGatewayResponse(jsii.String("test-response"), &gatewayResponseOptions{
//   	type: apigateway.responseType_ACCESS_DENIED(),
//   	statusCode: jsii.String("500"),
//   	responseHeaders: map[string]*string{
//   		"Access-Control-Allow-Origin": jsii.String("test.com"),
//   		"test-key": jsii.String("test-value"),
//   	},
//   	templates: map[string]*string{
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

// Represents a REST API in Amazon API Gateway.
//
// Use `addResource` and `addMethod` to configure the API model.
//
// By default, the API will automatically be deployed and accessible from a
// public endpoint.
//
// Example:
//   stateMachine := stepfunctions.NewStateMachine(this, jsii.String("MyStateMachine"), &stateMachineProps{
//   	stateMachineType: stepfunctions.stateMachineType_EXPRESS,
//   	definition: stepfunctions.chain.start(stepfunctions.NewPass(this, jsii.String("Pass"))),
//   })
//
//   api := apigateway.NewRestApi(this, jsii.String("Api"), &restApiProps{
//   	restApiName: jsii.String("MyApi"),
//   })
//   api.root.addMethod(jsii.String("GET"), apigateway.stepFunctionsIntegration.startExecution(stateMachine))
//
type RestApi interface {
	RestApiBase
	CloudWatchAccount() CfnAccount
	SetCloudWatchAccount(val CfnAccount)
	// API Gateway stage that points to the latest deployment (if defined).
	//
	// If `deploy` is disabled, you will need to explicitly assign this value in order to
	// set up integrations.
	DeploymentStage() Stage
	SetDeploymentStage(val Stage)
	// The first domain name mapped to this API, if defined through the `domainName` configuration prop, or added via `addDomainName`.
	DomainName() DomainName
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// API Gateway deployment that represents the latest changes of the API.
	//
	// This resource will be automatically updated every time the REST API model changes.
	// This will be undefined if `deploy` is false.
	LatestDeployment() Deployment
	// The list of methods bound to this RestApi.
	Methods() *[]Method
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The ID of this API Gateway RestApi.
	RestApiId() *string
	// A human friendly name for this Rest API.
	//
	// Note that this is different from `restApiId`.
	RestApiName() *string
	// The resource ID of the root resource.
	RestApiRootResourceId() *string
	// Represents the root resource of this API endpoint ('/').
	//
	// Resources and Methods are added to this resource.
	Root() IResource
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The deployed root URL of this REST API.
	Url() *string
	// Add an ApiKey.
	AddApiKey(id *string, options *ApiKeyOptions) IApiKey
	// Defines an API Gateway domain name and maps it to this API.
	AddDomainName(id *string, options *DomainNameOptions) DomainName
	// Adds a new gateway response.
	AddGatewayResponse(id *string, options *GatewayResponseOptions) GatewayResponse
	// Adds a new model.
	AddModel(id *string, props *ModelOptions) Model
	// Adds a new request validator.
	AddRequestValidator(id *string, props *RequestValidatorOptions) RequestValidator
	// Adds a usage plan.
	AddUsagePlan(id *string, props *UsagePlanProps) UsagePlan
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Gets the "execute-api" ARN.
	ArnForExecuteApi(method *string, path *string, stage *string) *string
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns the given named metric for this API.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of requests served from the API cache in a given period.
	//
	// Default: sum over 5 minutes.
	MetricCacheHitCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of requests served from the backend in a given period, when API caching is enabled.
	//
	// Default: sum over 5 minutes.
	MetricCacheMissCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of client-side errors captured in a given period.
	//
	// Default: sum over 5 minutes.
	MetricClientError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the total number API requests in a given period.
	//
	// Default: sample count over 5 minutes.
	MetricCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the time between when API Gateway relays a request to the backend and when it receives a response from the backend.
	//
	// Default: average over 5 minutes.
	MetricIntegrationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time between when API Gateway receives a request from a client and when it returns a response to the client.
	//
	// The latency includes the integration latency and other API Gateway overhead.
	//
	// Default: average over 5 minutes.
	MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of server-side errors captured in a given period.
	//
	// Default: sum over 5 minutes.
	MetricServerError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns a string representation of this construct.
	ToString() *string
	// Returns the URL for an HTTP path.
	//
	// Fails if `deploymentStage` is not set either by `deploy` or explicitly.
	UrlForPath(path *string) *string
}

// The jsii proxy struct for RestApi
type jsiiProxy_RestApi struct {
	jsiiProxy_RestApiBase
}

func (j *jsiiProxy_RestApi) CloudWatchAccount() CfnAccount {
	var returns CfnAccount
	_jsii_.Get(
		j,
		"cloudWatchAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApi) DeploymentStage() Stage {
	var returns Stage
	_jsii_.Get(
		j,
		"deploymentStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApi) DomainName() DomainName {
	var returns DomainName
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApi) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApi) LatestDeployment() Deployment {
	var returns Deployment
	_jsii_.Get(
		j,
		"latestDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApi) Methods() *[]Method {
	var returns *[]Method
	_jsii_.Get(
		j,
		"methods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApi) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApi) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApi) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApi) RestApiName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApi) RestApiRootResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiRootResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApi) Root() IResource {
	var returns IResource
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApi) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApi) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}


func NewRestApi(scope constructs.Construct, id *string, props *RestApiProps) RestApi {
	_init_.Initialize()

	j := jsiiProxy_RestApi{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.RestApi",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewRestApi_Override(r RestApi, scope constructs.Construct, id *string, props *RestApiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.RestApi",
		[]interface{}{scope, id, props},
		r,
	)
}

func (j *jsiiProxy_RestApi) SetCloudWatchAccount(val CfnAccount) {
	_jsii_.Set(
		j,
		"cloudWatchAccount",
		val,
	)
}

func (j *jsiiProxy_RestApi) SetDeploymentStage(val Stage) {
	_jsii_.Set(
		j,
		"deploymentStage",
		val,
	)
}

// Import an existing RestApi that can be configured with additional Methods and Resources.
func RestApi_FromRestApiAttributes(scope constructs.Construct, id *string, attrs *RestApiAttributes) IRestApi {
	_init_.Initialize()

	var returns IRestApi

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.RestApi",
		"fromRestApiAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Import an existing RestApi.
func RestApi_FromRestApiId(scope constructs.Construct, id *string, restApiId *string) IRestApi {
	_init_.Initialize()

	var returns IRestApi

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.RestApi",
		"fromRestApiId",
		[]interface{}{scope, id, restApiId},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func RestApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.RestApi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func RestApi_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.RestApi",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func RestApi_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.RestApi",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) AddApiKey(id *string, options *ApiKeyOptions) IApiKey {
	var returns IApiKey

	_jsii_.Invoke(
		r,
		"addApiKey",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) AddDomainName(id *string, options *DomainNameOptions) DomainName {
	var returns DomainName

	_jsii_.Invoke(
		r,
		"addDomainName",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) AddGatewayResponse(id *string, options *GatewayResponseOptions) GatewayResponse {
	var returns GatewayResponse

	_jsii_.Invoke(
		r,
		"addGatewayResponse",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) AddModel(id *string, props *ModelOptions) Model {
	var returns Model

	_jsii_.Invoke(
		r,
		"addModel",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) AddRequestValidator(id *string, props *RequestValidatorOptions) RequestValidator {
	var returns RequestValidator

	_jsii_.Invoke(
		r,
		"addRequestValidator",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) AddUsagePlan(id *string, props *UsagePlanProps) UsagePlan {
	var returns UsagePlan

	_jsii_.Invoke(
		r,
		"addUsagePlan",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_RestApi) ArnForExecuteApi(method *string, path *string, stage *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"arnForExecuteApi",
		[]interface{}{method, path, stage},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) MetricCacheHitCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricCacheHitCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) MetricCacheMissCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricCacheMissCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) MetricClientError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricClientError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) MetricCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) MetricIntegrationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricIntegrationLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) MetricServerError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricServerError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) UrlForPath(path *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"urlForPath",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Attributes that can be specified when importing a RestApi.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/constructs-go/constructs"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   /**
//    * This file showcases how to split up a RestApi's Resources and Methods across nested stacks.
//    *
//    * The root stack 'RootStack' first defines a RestApi.
//    * Two nested stacks BooksStack and PetsStack, create corresponding Resources '/books' and '/pets'.
//    * They are then deployed to a 'prod' Stage via a third nested stack - DeployStack.
//    *
//    * To verify this worked, go to the APIGateway
//    */
//
//   type rootStack struct {
//   	stack
//   }
//
//   func newRootStack(scope construct) *rootStack {
//   	this := &rootStack{}
//   	newStack_Override(this, scope, jsii.String("integ-restapi-import-RootStack"))
//
//   	restApi := awscdk.NewRestApi(this, jsii.String("RestApi"), &restApiProps{
//   		deploy: jsii.Boolean(false),
//   	})
//   	restApi.root.addMethod(jsii.String("ANY"))
//
//   	petsStack := NewPetsStack(this, &resourceNestedStackProps{
//   		restApiId: restApi.restApiId,
//   		rootResourceId: restApi.restApiRootResourceId,
//   	})
//   	booksStack := NewBooksStack(this, &resourceNestedStackProps{
//   		restApiId: restApi.restApiId,
//   		rootResourceId: restApi.restApiRootResourceId,
//   	})
//   	NewDeployStack(this, &deployStackProps{
//   		restApiId: restApi.restApiId,
//   		methods: petsStack.methods.concat(booksStack.methods),
//   	})
//
//   	awscdk.NewCfnOutput(this, jsii.String("PetsURL"), &cfnOutputProps{
//   		value: fmt.Sprintf("https://%v.execute-api.%v.amazonaws.com/prod/pets", restApi.restApiId, this.region),
//   	})
//
//   	awscdk.NewCfnOutput(this, jsii.String("BooksURL"), &cfnOutputProps{
//   		value: fmt.Sprintf("https://%v.execute-api.%v.amazonaws.com/prod/books", restApi.restApiId, this.region),
//   	})
//   	return this
//   }
//
//   type resourceNestedStackProps struct {
//   	nestedStackProps
//   	restApiId *string
//   	rootResourceId *string
//   }
//
//   type petsStack struct {
//   	nestedStack
//   	methods []method
//   }
//
//   func newPetsStack(scope construct, props resourceNestedStackProps) *petsStack {
//   	this := &petsStack{}
//   	newNestedStack_Override(this, scope, jsii.String("integ-restapi-import-PetsStack"), props)
//
//   	api := awscdk.RestApi.fromRestApiAttributes(this, jsii.String("RestApi"), &restApiAttributes{
//   		restApiId: props.restApiId,
//   		rootResourceId: props.rootResourceId,
//   	})
//
//   	method := api.root.addResource(jsii.String("pets")).addMethod(jsii.String("GET"), awscdk.NewMockIntegration(&integrationOptions{
//   		integrationResponses: []integrationResponse{
//   			&integrationResponse{
//   				statusCode: jsii.String("200"),
//   			},
//   		},
//   		passthroughBehavior: awscdk.PassthroughBehavior_NEVER,
//   		requestTemplates: map[string]*string{
//   			"application/json": jsii.String("{ \"statusCode\": 200 }"),
//   		},
//   	}), &methodOptions{
//   		methodResponses: []methodResponse{
//   			&methodResponse{
//   				statusCode: jsii.String("200"),
//   			},
//   		},
//   	})
//
//   	this.methods.push(method)
//   	return this
//   }
//
//   type booksStack struct {
//   	nestedStack
//   	methods []*method
//   }
//
//   func newBooksStack(scope construct, props resourceNestedStackProps) *booksStack {
//   	this := &booksStack{}
//   	newNestedStack_Override(this, scope, jsii.String("integ-restapi-import-BooksStack"), props)
//
//   	api := awscdk.RestApi.fromRestApiAttributes(this, jsii.String("RestApi"), &restApiAttributes{
//   		restApiId: props.restApiId,
//   		rootResourceId: props.rootResourceId,
//   	})
//
//   	method := api.root.addResource(jsii.String("books")).addMethod(jsii.String("GET"), awscdk.NewMockIntegration(&integrationOptions{
//   		integrationResponses: []*integrationResponse{
//   			&integrationResponse{
//   				statusCode: jsii.String("200"),
//   			},
//   		},
//   		passthroughBehavior: awscdk.PassthroughBehavior_NEVER,
//   		requestTemplates: map[string]*string{
//   			"application/json": jsii.String("{ \"statusCode\": 200 }"),
//   		},
//   	}), &methodOptions{
//   		methodResponses: []*methodResponse{
//   			&methodResponse{
//   				statusCode: jsii.String("200"),
//   			},
//   		},
//   	})
//
//   	this.methods.push(method)
//   	return this
//   }
//
//   type deployStackProps struct {
//   	nestedStackProps
//   	restApiId *string
//   	methods []*method
//   }
//
//   type deployStack struct {
//   	nestedStack
//   }
//
//   func newDeployStack(scope construct, props deployStackProps) *deployStack {
//   	this := &deployStack{}
//   	newNestedStack_Override(this, scope, jsii.String("integ-restapi-import-DeployStack"), props)
//
//   	deployment := awscdk.NewDeployment(this, jsii.String("Deployment"), &deploymentProps{
//   		api: awscdk.RestApi.fromRestApiId(this, jsii.String("RestApi"), props.restApiId),
//   	})
//   	if *props.methods {
//   		for _, method := range *props.methods {
//   			deployment.node.addDependency(method)
//   		}
//   	}
//   	awscdk.NewStage(this, jsii.String("Stage"), &stageProps{
//   		deployment: deployment,
//   	})
//   	return this
//   }
//
//   NewRootStack(awscdk.NewApp())
//
type RestApiAttributes struct {
	// The ID of the API Gateway RestApi.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// The resource ID of the root resource.
	RootResourceId *string `field:"required" json:"rootResourceId" yaml:"rootResourceId"`
	// The name of the API Gateway RestApi.
	RestApiName *string `field:"optional" json:"restApiName" yaml:"restApiName"`
}

// Base implementation that are common to various implementations of IRestApi.
//
// Example:
//   import route53 "github.com/aws/aws-cdk-go/awscdk"
//   import targets "github.com/aws/aws-cdk-go/awscdk"
//
//   var api restApi
//   var hostedZoneForExampleCom interface{}
//
//
//   route53.NewARecord(this, jsii.String("CustomDomainAliasRecord"), &aRecordProps{
//   	zone: hostedZoneForExampleCom,
//   	target: route53.recordTarget.fromAlias(targets.NewApiGateway(api)),
//   })
//
type RestApiBase interface {
	awscdk.Resource
	IRestApi
	CloudWatchAccount() CfnAccount
	SetCloudWatchAccount(val CfnAccount)
	// API Gateway stage that points to the latest deployment (if defined).
	//
	// If `deploy` is disabled, you will need to explicitly assign this value in order to
	// set up integrations.
	DeploymentStage() Stage
	SetDeploymentStage(val Stage)
	// The first domain name mapped to this API, if defined through the `domainName` configuration prop, or added via `addDomainName`.
	DomainName() DomainName
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// API Gateway deployment that represents the latest changes of the API.
	//
	// This resource will be automatically updated every time the REST API model changes.
	// This will be undefined if `deploy` is false.
	LatestDeployment() Deployment
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The ID of this API Gateway RestApi.
	RestApiId() *string
	// A human friendly name for this Rest API.
	//
	// Note that this is different from `restApiId`.
	RestApiName() *string
	// The resource ID of the root resource.
	RestApiRootResourceId() *string
	// Represents the root resource of this API endpoint ('/').
	//
	// Resources and Methods are added to this resource.
	Root() IResource
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Add an ApiKey.
	AddApiKey(id *string, options *ApiKeyOptions) IApiKey
	// Defines an API Gateway domain name and maps it to this API.
	AddDomainName(id *string, options *DomainNameOptions) DomainName
	// Adds a new gateway response.
	AddGatewayResponse(id *string, options *GatewayResponseOptions) GatewayResponse
	// Adds a usage plan.
	AddUsagePlan(id *string, props *UsagePlanProps) UsagePlan
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Gets the "execute-api" ARN.
	ArnForExecuteApi(method *string, path *string, stage *string) *string
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns the given named metric for this API.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of requests served from the API cache in a given period.
	//
	// Default: sum over 5 minutes.
	MetricCacheHitCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of requests served from the backend in a given period, when API caching is enabled.
	//
	// Default: sum over 5 minutes.
	MetricCacheMissCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of client-side errors captured in a given period.
	//
	// Default: sum over 5 minutes.
	MetricClientError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the total number API requests in a given period.
	//
	// Default: sample count over 5 minutes.
	MetricCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the time between when API Gateway relays a request to the backend and when it receives a response from the backend.
	//
	// Default: average over 5 minutes.
	MetricIntegrationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time between when API Gateway receives a request from a client and when it returns a response to the client.
	//
	// The latency includes the integration latency and other API Gateway overhead.
	//
	// Default: average over 5 minutes.
	MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of server-side errors captured in a given period.
	//
	// Default: sum over 5 minutes.
	MetricServerError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns a string representation of this construct.
	ToString() *string
	// Returns the URL for an HTTP path.
	//
	// Fails if `deploymentStage` is not set either by `deploy` or explicitly.
	UrlForPath(path *string) *string
}

// The jsii proxy struct for RestApiBase
type jsiiProxy_RestApiBase struct {
	internal.Type__awscdkResource
	jsiiProxy_IRestApi
}

func (j *jsiiProxy_RestApiBase) CloudWatchAccount() CfnAccount {
	var returns CfnAccount
	_jsii_.Get(
		j,
		"cloudWatchAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApiBase) DeploymentStage() Stage {
	var returns Stage
	_jsii_.Get(
		j,
		"deploymentStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApiBase) DomainName() DomainName {
	var returns DomainName
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApiBase) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApiBase) LatestDeployment() Deployment {
	var returns Deployment
	_jsii_.Get(
		j,
		"latestDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApiBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApiBase) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApiBase) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApiBase) RestApiName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApiBase) RestApiRootResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiRootResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApiBase) Root() IResource {
	var returns IResource
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApiBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewRestApiBase_Override(r RestApiBase, scope constructs.Construct, id *string, props *RestApiBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.RestApiBase",
		[]interface{}{scope, id, props},
		r,
	)
}

func (j *jsiiProxy_RestApiBase) SetCloudWatchAccount(val CfnAccount) {
	_jsii_.Set(
		j,
		"cloudWatchAccount",
		val,
	)
}

func (j *jsiiProxy_RestApiBase) SetDeploymentStage(val Stage) {
	_jsii_.Set(
		j,
		"deploymentStage",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func RestApiBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.RestApiBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func RestApiBase_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.RestApiBase",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func RestApiBase_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.RestApiBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApiBase) AddApiKey(id *string, options *ApiKeyOptions) IApiKey {
	var returns IApiKey

	_jsii_.Invoke(
		r,
		"addApiKey",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApiBase) AddDomainName(id *string, options *DomainNameOptions) DomainName {
	var returns DomainName

	_jsii_.Invoke(
		r,
		"addDomainName",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApiBase) AddGatewayResponse(id *string, options *GatewayResponseOptions) GatewayResponse {
	var returns GatewayResponse

	_jsii_.Invoke(
		r,
		"addGatewayResponse",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApiBase) AddUsagePlan(id *string, props *UsagePlanProps) UsagePlan {
	var returns UsagePlan

	_jsii_.Invoke(
		r,
		"addUsagePlan",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApiBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_RestApiBase) ArnForExecuteApi(method *string, path *string, stage *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"arnForExecuteApi",
		[]interface{}{method, path, stage},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApiBase) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApiBase) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApiBase) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApiBase) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApiBase) MetricCacheHitCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricCacheHitCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApiBase) MetricCacheMissCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricCacheMissCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApiBase) MetricClientError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricClientError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApiBase) MetricCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApiBase) MetricIntegrationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricIntegrationLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApiBase) MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApiBase) MetricServerError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricServerError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApiBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApiBase) UrlForPath(path *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"urlForPath",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Represents the props that all Rest APIs share.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var accessLogDestination iAccessLogDestination
//   var accessLogFormat accessLogFormat
//   var bucket bucket
//   var certificate certificate
//   var policyDocument policyDocument
//
//   restApiBaseProps := &restApiBaseProps{
//   	cloudWatchRole: jsii.Boolean(false),
//   	deploy: jsii.Boolean(false),
//   	deployOptions: &stageOptions{
//   		accessLogDestination: accessLogDestination,
//   		accessLogFormat: accessLogFormat,
//   		cacheClusterEnabled: jsii.Boolean(false),
//   		cacheClusterSize: jsii.String("cacheClusterSize"),
//   		cacheDataEncrypted: jsii.Boolean(false),
//   		cacheTtl: cdk.duration.minutes(jsii.Number(30)),
//   		cachingEnabled: jsii.Boolean(false),
//   		clientCertificateId: jsii.String("clientCertificateId"),
//   		dataTraceEnabled: jsii.Boolean(false),
//   		description: jsii.String("description"),
//   		documentationVersion: jsii.String("documentationVersion"),
//   		loggingLevel: awscdk.Aws_apigateway.methodLoggingLevel_OFF,
//   		methodOptions: map[string]methodDeploymentOptions{
//   			"methodOptionsKey": &methodDeploymentOptions{
//   				"cacheDataEncrypted": jsii.Boolean(false),
//   				"cacheTtl": cdk.*duration.minutes(jsii.Number(30)),
//   				"cachingEnabled": jsii.Boolean(false),
//   				"dataTraceEnabled": jsii.Boolean(false),
//   				"loggingLevel": awscdk.*Aws_apigateway.*methodLoggingLevel_OFF,
//   				"metricsEnabled": jsii.Boolean(false),
//   				"throttlingBurstLimit": jsii.Number(123),
//   				"throttlingRateLimit": jsii.Number(123),
//   			},
//   		},
//   		metricsEnabled: jsii.Boolean(false),
//   		stageName: jsii.String("stageName"),
//   		throttlingBurstLimit: jsii.Number(123),
//   		throttlingRateLimit: jsii.Number(123),
//   		tracingEnabled: jsii.Boolean(false),
//   		variables: map[string]*string{
//   			"variablesKey": jsii.String("variables"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	disableExecuteApiEndpoint: jsii.Boolean(false),
//   	domainName: &domainNameOptions{
//   		certificate: certificate,
//   		domainName: jsii.String("domainName"),
//
//   		// the properties below are optional
//   		basePath: jsii.String("basePath"),
//   		endpointType: awscdk.*Aws_apigateway.endpointType_EDGE,
//   		mtls: &mTLSConfig{
//   			bucket: bucket,
//   			key: jsii.String("key"),
//
//   			// the properties below are optional
//   			version: jsii.String("version"),
//   		},
//   		securityPolicy: awscdk.*Aws_apigateway.securityPolicy_TLS_1_0,
//   	},
//   	endpointExportName: jsii.String("endpointExportName"),
//   	endpointTypes: []*endpointType{
//   		awscdk.*Aws_apigateway.*endpointType_EDGE,
//   	},
//   	failOnWarnings: jsii.Boolean(false),
//   	parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	policy: policyDocument,
//   	restApiName: jsii.String("restApiName"),
//   	retainDeployments: jsii.Boolean(false),
//   }
//
type RestApiBaseProps struct {
	// Automatically configure an AWS CloudWatch role for API Gateway.
	CloudWatchRole *bool `field:"optional" json:"cloudWatchRole" yaml:"cloudWatchRole"`
	// Indicates if a Deployment should be automatically created for this API, and recreated when the API model (resources, methods) changes.
	//
	// Since API Gateway deployments are immutable, When this option is enabled
	// (by default), an AWS::ApiGateway::Deployment resource will automatically
	// created with a logical ID that hashes the API model (methods, resources
	// and options). This means that when the model changes, the logical ID of
	// this CloudFormation resource will change, and a new deployment will be
	// created.
	//
	// If this is set, `latestDeployment` will refer to the `Deployment` object
	// and `deploymentStage` will refer to a `Stage` that points to this
	// deployment. To customize the stage options, use the `deployOptions`
	// property.
	//
	// A CloudFormation Output will also be defined with the root URL endpoint
	// of this REST API.
	Deploy *bool `field:"optional" json:"deploy" yaml:"deploy"`
	// Options for the API Gateway stage that will always point to the latest deployment when `deploy` is enabled.
	//
	// If `deploy` is disabled,
	// this value cannot be set.
	DeployOptions *StageOptions `field:"optional" json:"deployOptions" yaml:"deployOptions"`
	// A description of the RestApi construct.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether clients can invoke the API using the default execute-api endpoint.
	//
	// To require that clients use a custom domain name to invoke the
	// API, disable the default endpoint.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html
	//
	DisableExecuteApiEndpoint *bool `field:"optional" json:"disableExecuteApiEndpoint" yaml:"disableExecuteApiEndpoint"`
	// Configure a custom domain name and map it to this API.
	DomainName *DomainNameOptions `field:"optional" json:"domainName" yaml:"domainName"`
	// Export name for the CfnOutput containing the API endpoint.
	EndpointExportName *string `field:"optional" json:"endpointExportName" yaml:"endpointExportName"`
	// A list of the endpoint types of the API.
	//
	// Use this property when creating
	// an API.
	EndpointTypes *[]EndpointType `field:"optional" json:"endpointTypes" yaml:"endpointTypes"`
	// Indicates whether to roll back the resource if a warning occurs while API Gateway is creating the RestApi resource.
	FailOnWarnings *bool `field:"optional" json:"failOnWarnings" yaml:"failOnWarnings"`
	// Custom header parameters for the request.
	// See: https://docs.aws.amazon.com/cli/latest/reference/apigateway/import-rest-api.html
	//
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// A policy document that contains the permissions for this RestApi.
	Policy awsiam.PolicyDocument `field:"optional" json:"policy" yaml:"policy"`
	// A name for the API Gateway RestApi resource.
	RestApiName *string `field:"optional" json:"restApiName" yaml:"restApiName"`
	// Retains old deployment resources when the API changes.
	//
	// This allows
	// manually reverting stages to point to old deployments via the AWS
	// Console.
	RetainDeployments *bool `field:"optional" json:"retainDeployments" yaml:"retainDeployments"`
}

// Props to create a new instance of RestApi.
//
// Example:
//   stateMachine := stepfunctions.NewStateMachine(this, jsii.String("MyStateMachine"), &stateMachineProps{
//   	stateMachineType: stepfunctions.stateMachineType_EXPRESS,
//   	definition: stepfunctions.chain.start(stepfunctions.NewPass(this, jsii.String("Pass"))),
//   })
//
//   api := apigateway.NewRestApi(this, jsii.String("Api"), &restApiProps{
//   	restApiName: jsii.String("MyApi"),
//   })
//   api.root.addMethod(jsii.String("GET"), apigateway.stepFunctionsIntegration.startExecution(stateMachine))
//
type RestApiProps struct {
	// Adds a CORS preflight OPTIONS method to this resource and all child resources.
	//
	// You can add CORS at the resource-level using `addCorsPreflight`.
	DefaultCorsPreflightOptions *CorsOptions `field:"optional" json:"defaultCorsPreflightOptions" yaml:"defaultCorsPreflightOptions"`
	// An integration to use as a default for all methods created within this API unless an integration is specified.
	DefaultIntegration Integration `field:"optional" json:"defaultIntegration" yaml:"defaultIntegration"`
	// Method options to use as a default for all methods created within this API unless custom options are specified.
	DefaultMethodOptions *MethodOptions `field:"optional" json:"defaultMethodOptions" yaml:"defaultMethodOptions"`
	// Automatically configure an AWS CloudWatch role for API Gateway.
	CloudWatchRole *bool `field:"optional" json:"cloudWatchRole" yaml:"cloudWatchRole"`
	// Indicates if a Deployment should be automatically created for this API, and recreated when the API model (resources, methods) changes.
	//
	// Since API Gateway deployments are immutable, When this option is enabled
	// (by default), an AWS::ApiGateway::Deployment resource will automatically
	// created with a logical ID that hashes the API model (methods, resources
	// and options). This means that when the model changes, the logical ID of
	// this CloudFormation resource will change, and a new deployment will be
	// created.
	//
	// If this is set, `latestDeployment` will refer to the `Deployment` object
	// and `deploymentStage` will refer to a `Stage` that points to this
	// deployment. To customize the stage options, use the `deployOptions`
	// property.
	//
	// A CloudFormation Output will also be defined with the root URL endpoint
	// of this REST API.
	Deploy *bool `field:"optional" json:"deploy" yaml:"deploy"`
	// Options for the API Gateway stage that will always point to the latest deployment when `deploy` is enabled.
	//
	// If `deploy` is disabled,
	// this value cannot be set.
	DeployOptions *StageOptions `field:"optional" json:"deployOptions" yaml:"deployOptions"`
	// A description of the RestApi construct.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether clients can invoke the API using the default execute-api endpoint.
	//
	// To require that clients use a custom domain name to invoke the
	// API, disable the default endpoint.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html
	//
	DisableExecuteApiEndpoint *bool `field:"optional" json:"disableExecuteApiEndpoint" yaml:"disableExecuteApiEndpoint"`
	// Configure a custom domain name and map it to this API.
	DomainName *DomainNameOptions `field:"optional" json:"domainName" yaml:"domainName"`
	// Export name for the CfnOutput containing the API endpoint.
	EndpointExportName *string `field:"optional" json:"endpointExportName" yaml:"endpointExportName"`
	// A list of the endpoint types of the API.
	//
	// Use this property when creating
	// an API.
	EndpointTypes *[]EndpointType `field:"optional" json:"endpointTypes" yaml:"endpointTypes"`
	// Indicates whether to roll back the resource if a warning occurs while API Gateway is creating the RestApi resource.
	FailOnWarnings *bool `field:"optional" json:"failOnWarnings" yaml:"failOnWarnings"`
	// Custom header parameters for the request.
	// See: https://docs.aws.amazon.com/cli/latest/reference/apigateway/import-rest-api.html
	//
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// A policy document that contains the permissions for this RestApi.
	Policy awsiam.PolicyDocument `field:"optional" json:"policy" yaml:"policy"`
	// A name for the API Gateway RestApi resource.
	RestApiName *string `field:"optional" json:"restApiName" yaml:"restApiName"`
	// Retains old deployment resources when the API changes.
	//
	// This allows
	// manually reverting stages to point to old deployments via the AWS
	// Console.
	RetainDeployments *bool `field:"optional" json:"retainDeployments" yaml:"retainDeployments"`
	// The source of the API key for metering requests according to a usage plan.
	ApiKeySourceType ApiKeySourceType `field:"optional" json:"apiKeySourceType" yaml:"apiKeySourceType"`
	// The list of binary media mime-types that are supported by the RestApi resource, such as "image/png" or "application/octet-stream".
	BinaryMediaTypes *[]*string `field:"optional" json:"binaryMediaTypes" yaml:"binaryMediaTypes"`
	// The ID of the API Gateway RestApi resource that you want to clone.
	CloneFrom IRestApi `field:"optional" json:"cloneFrom" yaml:"cloneFrom"`
	// The EndpointConfiguration property type specifies the endpoint types of a REST API.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-restapi-endpointconfiguration.html
	//
	EndpointConfiguration *EndpointConfiguration `field:"optional" json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// A nullable integer that is used to enable compression (with non-negative between 0 and 10485760 (10M) bytes, inclusive) or disable compression (when undefined) on an API.
	//
	// When compression is enabled, compression or
	// decompression is not applied on the payload if the payload size is
	// smaller than this value. Setting it to zero allows compression for any
	// payload size.
	MinimumCompressionSize *float64 `field:"optional" json:"minimumCompressionSize" yaml:"minimumCompressionSize"`
}

// OpenAPI specification from an S3 archive.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   s3ApiDefinition := awscdk.Aws_apigateway.NewS3ApiDefinition(bucket, jsii.String("key"), jsii.String("objectVersion"))
//
type S3ApiDefinition interface {
	ApiDefinition
	// Called when the specification is initialized to allow this object to bind to the stack, add resources and have fun.
	Bind(_scope constructs.Construct) *ApiDefinitionConfig
	// Called after the CFN RestApi resource has been created to allow the Api Definition to bind to it.
	//
	// Specifically it's required to allow assets to add
	// metadata for tooling like SAM CLI to be able to find their origins.
	BindAfterCreate(_scope constructs.Construct, _restApi IRestApi)
}

// The jsii proxy struct for S3ApiDefinition
type jsiiProxy_S3ApiDefinition struct {
	jsiiProxy_ApiDefinition
}

func NewS3ApiDefinition(bucket awss3.IBucket, key *string, objectVersion *string) S3ApiDefinition {
	_init_.Initialize()

	j := jsiiProxy_S3ApiDefinition{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.S3ApiDefinition",
		[]interface{}{bucket, key, objectVersion},
		&j,
	)

	return &j
}

func NewS3ApiDefinition_Override(s S3ApiDefinition, bucket awss3.IBucket, key *string, objectVersion *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.S3ApiDefinition",
		[]interface{}{bucket, key, objectVersion},
		s,
	)
}

// Loads the API specification from a local disk asset.
func S3ApiDefinition_FromAsset(file *string, options *awss3assets.AssetOptions) AssetApiDefinition {
	_init_.Initialize()

	var returns AssetApiDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.S3ApiDefinition",
		"fromAsset",
		[]interface{}{file, options},
		&returns,
	)

	return returns
}

// Creates an API definition from a specification file in an S3 bucket.
func S3ApiDefinition_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3ApiDefinition {
	_init_.Initialize()

	var returns S3ApiDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.S3ApiDefinition",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Create an API definition from an inline object.
//
// The inline object must follow the
// schema of OpenAPI 2.0 or OpenAPI 3.0
//
// Example:
//     apigateway.ApiDefinition.fromInline({
//       openapi: '3.0.2',
//       paths: {
//         '/pets': {
//           get: {
//             'responses': {
//               200: {
//                 content: {
//                   'application/json': {
//                     schema: {
//                       $ref: '#/components/schemas/Empty',
//                     },
//                   },
//                 },
//               },
//             },
//             'x-amazon-apigateway-integration': {
//               responses: {
//                 default: {
//                   statusCode: '200',
//                 },
//               },
//               requestTemplates: {
//                 'application/json': '{"statusCode": 200}',
//               },
//               passthroughBehavior: 'when_no_match',
//               type: 'mock',
//             },
//           },
//         },
//       },
//       components: {
//         schemas: {
//           Empty: {
//             title: 'Empty Schema',
//             type: 'object',
//           },
//         },
//       },
//     });
//
func S3ApiDefinition_FromInline(definition interface{}) InlineApiDefinition {
	_init_.Initialize()

	var returns InlineApiDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.S3ApiDefinition",
		"fromInline",
		[]interface{}{definition},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ApiDefinition) Bind(_scope constructs.Construct) *ApiDefinitionConfig {
	var returns *ApiDefinitionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ApiDefinition) BindAfterCreate(_scope constructs.Construct, _restApi IRestApi) {
	_jsii_.InvokeVoid(
		s,
		"bindAfterCreate",
		[]interface{}{_scope, _restApi},
	)
}

// The minimum version of the SSL protocol that you want API Gateway to use for HTTPS connections.
//
// Example:
//   var acmCertificateForExampleCom interface{}
//
//
//   apigateway.NewDomainName(this, jsii.String("custom-domain"), &domainNameProps{
//   	domainName: jsii.String("example.com"),
//   	certificate: acmCertificateForExampleCom,
//   	endpointType: apigateway.endpointType_EDGE,
//   	 // default is REGIONAL
//   	securityPolicy: apigateway.securityPolicy_TLS_1_2,
//   })
//
type SecurityPolicy string

const (
	// Cipher suite TLS 1.0.
	SecurityPolicy_TLS_1_0 SecurityPolicy = "TLS_1_0"
	// Cipher suite TLS 1.2.
	SecurityPolicy_TLS_1_2 SecurityPolicy = "TLS_1_2"
)

// Represents a REST API in Amazon API Gateway, created with an OpenAPI specification.
//
// Some properties normally accessible on @see {@link RestApi} - such as the description -
// must be declared in the specification. All Resources and Methods need to be defined as
// part of the OpenAPI specification file, and cannot be added via the CDK.
//
// By default, the API will automatically be deployed and accessible from a
// public endpoint.
//
// Example:
//   var integration integration
//
//
//   api := apigateway.NewSpecRestApi(this, jsii.String("books-api"), &specRestApiProps{
//   	apiDefinition: apigateway.apiDefinition.fromAsset(jsii.String("path-to-file.json")),
//   })
//
//   booksResource := api.root.addResource(jsii.String("books"))
//   booksResource.addMethod(jsii.String("GET"), integration)
//
type SpecRestApi interface {
	RestApiBase
	CloudWatchAccount() CfnAccount
	SetCloudWatchAccount(val CfnAccount)
	// API Gateway stage that points to the latest deployment (if defined).
	//
	// If `deploy` is disabled, you will need to explicitly assign this value in order to
	// set up integrations.
	DeploymentStage() Stage
	SetDeploymentStage(val Stage)
	// The first domain name mapped to this API, if defined through the `domainName` configuration prop, or added via `addDomainName`.
	DomainName() DomainName
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// API Gateway deployment that represents the latest changes of the API.
	//
	// This resource will be automatically updated every time the REST API model changes.
	// This will be undefined if `deploy` is false.
	LatestDeployment() Deployment
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The ID of this API Gateway RestApi.
	RestApiId() *string
	// A human friendly name for this Rest API.
	//
	// Note that this is different from `restApiId`.
	RestApiName() *string
	// The resource ID of the root resource.
	RestApiRootResourceId() *string
	// Represents the root resource of this API endpoint ('/').
	//
	// Resources and Methods are added to this resource.
	Root() IResource
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Add an ApiKey.
	AddApiKey(id *string, options *ApiKeyOptions) IApiKey
	// Defines an API Gateway domain name and maps it to this API.
	AddDomainName(id *string, options *DomainNameOptions) DomainName
	// Adds a new gateway response.
	AddGatewayResponse(id *string, options *GatewayResponseOptions) GatewayResponse
	// Adds a usage plan.
	AddUsagePlan(id *string, props *UsagePlanProps) UsagePlan
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Gets the "execute-api" ARN.
	ArnForExecuteApi(method *string, path *string, stage *string) *string
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns the given named metric for this API.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of requests served from the API cache in a given period.
	//
	// Default: sum over 5 minutes.
	MetricCacheHitCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of requests served from the backend in a given period, when API caching is enabled.
	//
	// Default: sum over 5 minutes.
	MetricCacheMissCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of client-side errors captured in a given period.
	//
	// Default: sum over 5 minutes.
	MetricClientError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the total number API requests in a given period.
	//
	// Default: sample count over 5 minutes.
	MetricCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the time between when API Gateway relays a request to the backend and when it receives a response from the backend.
	//
	// Default: average over 5 minutes.
	MetricIntegrationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time between when API Gateway receives a request from a client and when it returns a response to the client.
	//
	// The latency includes the integration latency and other API Gateway overhead.
	//
	// Default: average over 5 minutes.
	MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of server-side errors captured in a given period.
	//
	// Default: sum over 5 minutes.
	MetricServerError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns a string representation of this construct.
	ToString() *string
	// Returns the URL for an HTTP path.
	//
	// Fails if `deploymentStage` is not set either by `deploy` or explicitly.
	UrlForPath(path *string) *string
}

// The jsii proxy struct for SpecRestApi
type jsiiProxy_SpecRestApi struct {
	jsiiProxy_RestApiBase
}

func (j *jsiiProxy_SpecRestApi) CloudWatchAccount() CfnAccount {
	var returns CfnAccount
	_jsii_.Get(
		j,
		"cloudWatchAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpecRestApi) DeploymentStage() Stage {
	var returns Stage
	_jsii_.Get(
		j,
		"deploymentStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpecRestApi) DomainName() DomainName {
	var returns DomainName
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpecRestApi) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpecRestApi) LatestDeployment() Deployment {
	var returns Deployment
	_jsii_.Get(
		j,
		"latestDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpecRestApi) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpecRestApi) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpecRestApi) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpecRestApi) RestApiName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpecRestApi) RestApiRootResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiRootResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpecRestApi) Root() IResource {
	var returns IResource
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpecRestApi) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewSpecRestApi(scope constructs.Construct, id *string, props *SpecRestApiProps) SpecRestApi {
	_init_.Initialize()

	j := jsiiProxy_SpecRestApi{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.SpecRestApi",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSpecRestApi_Override(s SpecRestApi, scope constructs.Construct, id *string, props *SpecRestApiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.SpecRestApi",
		[]interface{}{scope, id, props},
		s,
	)
}

func (j *jsiiProxy_SpecRestApi) SetCloudWatchAccount(val CfnAccount) {
	_jsii_.Set(
		j,
		"cloudWatchAccount",
		val,
	)
}

func (j *jsiiProxy_SpecRestApi) SetDeploymentStage(val Stage) {
	_jsii_.Set(
		j,
		"deploymentStage",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func SpecRestApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.SpecRestApi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func SpecRestApi_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.SpecRestApi",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func SpecRestApi_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.SpecRestApi",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) AddApiKey(id *string, options *ApiKeyOptions) IApiKey {
	var returns IApiKey

	_jsii_.Invoke(
		s,
		"addApiKey",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) AddDomainName(id *string, options *DomainNameOptions) DomainName {
	var returns DomainName

	_jsii_.Invoke(
		s,
		"addDomainName",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) AddGatewayResponse(id *string, options *GatewayResponseOptions) GatewayResponse {
	var returns GatewayResponse

	_jsii_.Invoke(
		s,
		"addGatewayResponse",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) AddUsagePlan(id *string, props *UsagePlanProps) UsagePlan {
	var returns UsagePlan

	_jsii_.Invoke(
		s,
		"addUsagePlan",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_SpecRestApi) ArnForExecuteApi(method *string, path *string, stage *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"arnForExecuteApi",
		[]interface{}{method, path, stage},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) MetricCacheHitCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricCacheHitCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) MetricCacheMissCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricCacheMissCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) MetricClientError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricClientError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) MetricCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) MetricIntegrationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricIntegrationLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) MetricServerError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricServerError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) UrlForPath(path *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"urlForPath",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Props to instantiate a new SpecRestApi.
//
// Example:
//   var integration integration
//
//
//   api := apigateway.NewSpecRestApi(this, jsii.String("books-api"), &specRestApiProps{
//   	apiDefinition: apigateway.apiDefinition.fromAsset(jsii.String("path-to-file.json")),
//   })
//
//   booksResource := api.root.addResource(jsii.String("books"))
//   booksResource.addMethod(jsii.String("GET"), integration)
//
type SpecRestApiProps struct {
	// Automatically configure an AWS CloudWatch role for API Gateway.
	CloudWatchRole *bool `field:"optional" json:"cloudWatchRole" yaml:"cloudWatchRole"`
	// Indicates if a Deployment should be automatically created for this API, and recreated when the API model (resources, methods) changes.
	//
	// Since API Gateway deployments are immutable, When this option is enabled
	// (by default), an AWS::ApiGateway::Deployment resource will automatically
	// created with a logical ID that hashes the API model (methods, resources
	// and options). This means that when the model changes, the logical ID of
	// this CloudFormation resource will change, and a new deployment will be
	// created.
	//
	// If this is set, `latestDeployment` will refer to the `Deployment` object
	// and `deploymentStage` will refer to a `Stage` that points to this
	// deployment. To customize the stage options, use the `deployOptions`
	// property.
	//
	// A CloudFormation Output will also be defined with the root URL endpoint
	// of this REST API.
	Deploy *bool `field:"optional" json:"deploy" yaml:"deploy"`
	// Options for the API Gateway stage that will always point to the latest deployment when `deploy` is enabled.
	//
	// If `deploy` is disabled,
	// this value cannot be set.
	DeployOptions *StageOptions `field:"optional" json:"deployOptions" yaml:"deployOptions"`
	// A description of the RestApi construct.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether clients can invoke the API using the default execute-api endpoint.
	//
	// To require that clients use a custom domain name to invoke the
	// API, disable the default endpoint.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html
	//
	DisableExecuteApiEndpoint *bool `field:"optional" json:"disableExecuteApiEndpoint" yaml:"disableExecuteApiEndpoint"`
	// Configure a custom domain name and map it to this API.
	DomainName *DomainNameOptions `field:"optional" json:"domainName" yaml:"domainName"`
	// Export name for the CfnOutput containing the API endpoint.
	EndpointExportName *string `field:"optional" json:"endpointExportName" yaml:"endpointExportName"`
	// A list of the endpoint types of the API.
	//
	// Use this property when creating
	// an API.
	EndpointTypes *[]EndpointType `field:"optional" json:"endpointTypes" yaml:"endpointTypes"`
	// Indicates whether to roll back the resource if a warning occurs while API Gateway is creating the RestApi resource.
	FailOnWarnings *bool `field:"optional" json:"failOnWarnings" yaml:"failOnWarnings"`
	// Custom header parameters for the request.
	// See: https://docs.aws.amazon.com/cli/latest/reference/apigateway/import-rest-api.html
	//
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// A policy document that contains the permissions for this RestApi.
	Policy awsiam.PolicyDocument `field:"optional" json:"policy" yaml:"policy"`
	// A name for the API Gateway RestApi resource.
	RestApiName *string `field:"optional" json:"restApiName" yaml:"restApiName"`
	// Retains old deployment resources when the API changes.
	//
	// This allows
	// manually reverting stages to point to old deployments via the AWS
	// Console.
	RetainDeployments *bool `field:"optional" json:"retainDeployments" yaml:"retainDeployments"`
	// An OpenAPI definition compatible with API Gateway.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-import-api.html
	//
	ApiDefinition ApiDefinition `field:"required" json:"apiDefinition" yaml:"apiDefinition"`
}

// Example:
//   // production stage
//   prdLogGroup := logs.NewLogGroup(this, jsii.String("PrdLogs"))
//   api := apigateway.NewRestApi(this, jsii.String("books"), &restApiProps{
//   	deployOptions: &stageOptions{
//   		accessLogDestination: apigateway.NewLogGroupLogDestination(prdLogGroup),
//   		accessLogFormat: apigateway.accessLogFormat.jsonWithStandardFields(),
//   	},
//   })
//   deployment := apigateway.NewDeployment(this, jsii.String("Deployment"), &deploymentProps{
//   	api: api,
//   })
//
//   // development stage
//   devLogGroup := logs.NewLogGroup(this, jsii.String("DevLogs"))
//   apigateway.NewStage(this, jsii.String("dev"), &stageProps{
//   	deployment: deployment,
//   	accessLogDestination: apigateway.NewLogGroupLogDestination(devLogGroup),
//   	accessLogFormat: apigateway.*accessLogFormat.jsonWithStandardFields(&jsonWithStandardFieldProps{
//   		caller: jsii.Boolean(false),
//   		httpMethod: jsii.Boolean(true),
//   		ip: jsii.Boolean(true),
//   		protocol: jsii.Boolean(true),
//   		requestTime: jsii.Boolean(true),
//   		resourcePath: jsii.Boolean(true),
//   		responseLength: jsii.Boolean(true),
//   		status: jsii.Boolean(true),
//   		user: jsii.Boolean(true),
//   	}),
//   })
//
type Stage interface {
	awscdk.Resource
	IStage
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// RestApi to which this stage is associated.
	RestApi() IRestApi
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Returns the resource ARN for this stage:.
	//
	// arn:aws:apigateway:{region}::/restapis/{restApiId}/stages/{stageName}
	//
	// Note that this is separate from the execute-api ARN for methods and resources
	// within this stage.
	StageArn() *string
	// Name of this stage.
	StageName() *string
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns the given named metric for this stage.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of requests served from the API cache in a given period.
	MetricCacheHitCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of requests served from the backend in a given period, when API caching is enabled.
	MetricCacheMissCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of client-side errors captured in a given period.
	MetricClientError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the total number API requests in a given period.
	MetricCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the time between when API Gateway relays a request to the backend and when it receives a response from the backend.
	MetricIntegrationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time between when API Gateway receives a request from a client and when it returns a response to the client.
	//
	// The latency includes the integration latency and other API Gateway overhead.
	MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of server-side errors captured in a given period.
	MetricServerError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns a string representation of this construct.
	ToString() *string
	// Returns the invoke URL for a certain path.
	UrlForPath(path *string) *string
}

// The jsii proxy struct for Stage
type jsiiProxy_Stage struct {
	internal.Type__awscdkResource
	jsiiProxy_IStage
}

func (j *jsiiProxy_Stage) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stage) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stage) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stage) RestApi() IRestApi {
	var returns IRestApi
	_jsii_.Get(
		j,
		"restApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stage) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stage) StageArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stage) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}


func NewStage(scope constructs.Construct, id *string, props *StageProps) Stage {
	_init_.Initialize()

	j := jsiiProxy_Stage{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.Stage",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewStage_Override(s Stage, scope constructs.Construct, id *string, props *StageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.Stage",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Stage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.Stage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func Stage_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.Stage",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Stage_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.Stage",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stage) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_Stage) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stage) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stage) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stage) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stage) MetricCacheHitCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricCacheHitCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stage) MetricCacheMissCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricCacheMissCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stage) MetricClientError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricClientError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stage) MetricCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stage) MetricIntegrationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricIntegrationLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stage) MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stage) MetricServerError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricServerError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stage) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stage) UrlForPath(path *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"urlForPath",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Example:
//   logGroup := logs.NewLogGroup(this, jsii.String("ApiGatewayAccessLogs"))
//   api := apigateway.NewRestApi(this, jsii.String("books"), &restApiProps{
//   	deployOptions: &stageOptions{
//   		accessLogDestination: apigateway.NewLogGroupLogDestination(logGroup),
//   		accessLogFormat: apigateway.accessLogFormat.clf(),
//   	},
//   })
//
type StageOptions struct {
	// Indicates whether the cached responses are encrypted.
	CacheDataEncrypted *bool `field:"optional" json:"cacheDataEncrypted" yaml:"cacheDataEncrypted"`
	// Specifies the time to live (TTL), in seconds, for cached responses.
	//
	// The
	// higher the TTL, the longer the response will be cached.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-caching.html
	//
	CacheTtl awscdk.Duration `field:"optional" json:"cacheTtl" yaml:"cacheTtl"`
	// Specifies whether responses should be cached and returned for requests.
	//
	// A
	// cache cluster must be enabled on the stage for responses to be cached.
	CachingEnabled *bool `field:"optional" json:"cachingEnabled" yaml:"cachingEnabled"`
	// Specifies whether data trace logging is enabled for this method.
	//
	// When enabled, API gateway will log the full API requests and responses.
	// This can be useful to troubleshoot APIs, but can result in logging sensitive data.
	// We recommend that you don't enable this feature for production APIs.
	DataTraceEnabled *bool `field:"optional" json:"dataTraceEnabled" yaml:"dataTraceEnabled"`
	// Specifies the logging level for this method, which effects the log entries pushed to Amazon CloudWatch Logs.
	LoggingLevel MethodLoggingLevel `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
	// Specifies whether Amazon CloudWatch metrics are enabled for this method.
	MetricsEnabled *bool `field:"optional" json:"metricsEnabled" yaml:"metricsEnabled"`
	// Specifies the throttling burst limit.
	//
	// The total rate of all requests in your AWS account is limited to 5,000 requests.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html
	//
	ThrottlingBurstLimit *float64 `field:"optional" json:"throttlingBurstLimit" yaml:"throttlingBurstLimit"`
	// Specifies the throttling rate limit.
	//
	// The total rate of all requests in your AWS account is limited to 10,000 requests per second (rps).
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html
	//
	ThrottlingRateLimit *float64 `field:"optional" json:"throttlingRateLimit" yaml:"throttlingRateLimit"`
	// The CloudWatch Logs log group.
	AccessLogDestination IAccessLogDestination `field:"optional" json:"accessLogDestination" yaml:"accessLogDestination"`
	// A single line format of access logs of data, as specified by selected $content variables.
	//
	// The format must include at least `AccessLogFormat.contextRequestId()`.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html#context-variable-reference
	//
	AccessLogFormat AccessLogFormat `field:"optional" json:"accessLogFormat" yaml:"accessLogFormat"`
	// Indicates whether cache clustering is enabled for the stage.
	CacheClusterEnabled *bool `field:"optional" json:"cacheClusterEnabled" yaml:"cacheClusterEnabled"`
	// The stage's cache cluster size.
	CacheClusterSize *string `field:"optional" json:"cacheClusterSize" yaml:"cacheClusterSize"`
	// The identifier of the client certificate that API Gateway uses to call your integration endpoints in the stage.
	ClientCertificateId *string `field:"optional" json:"clientCertificateId" yaml:"clientCertificateId"`
	// A description of the purpose of the stage.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The version identifier of the API documentation snapshot.
	DocumentationVersion *string `field:"optional" json:"documentationVersion" yaml:"documentationVersion"`
	// Method deployment options for specific resources/methods.
	//
	// These will
	// override common options defined in `StageOptions#methodOptions`.
	MethodOptions *map[string]*MethodDeploymentOptions `field:"optional" json:"methodOptions" yaml:"methodOptions"`
	// The name of the stage, which API Gateway uses as the first path segment in the invoked Uniform Resource Identifier (URI).
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
	// Specifies whether Amazon X-Ray tracing is enabled for this method.
	TracingEnabled *bool `field:"optional" json:"tracingEnabled" yaml:"tracingEnabled"`
	// A map that defines the stage variables.
	//
	// Variable names must consist of
	// alphanumeric characters, and the values must match the following regular
	// expression: [A-Za-z0-9-._~:/?#&amp;=,]+.
	Variables *map[string]*string `field:"optional" json:"variables" yaml:"variables"`
}

// Example:
//   // production stage
//   prdLogGroup := logs.NewLogGroup(this, jsii.String("PrdLogs"))
//   api := apigateway.NewRestApi(this, jsii.String("books"), &restApiProps{
//   	deployOptions: &stageOptions{
//   		accessLogDestination: apigateway.NewLogGroupLogDestination(prdLogGroup),
//   		accessLogFormat: apigateway.accessLogFormat.jsonWithStandardFields(),
//   	},
//   })
//   deployment := apigateway.NewDeployment(this, jsii.String("Deployment"), &deploymentProps{
//   	api: api,
//   })
//
//   // development stage
//   devLogGroup := logs.NewLogGroup(this, jsii.String("DevLogs"))
//   apigateway.NewStage(this, jsii.String("dev"), &stageProps{
//   	deployment: deployment,
//   	accessLogDestination: apigateway.NewLogGroupLogDestination(devLogGroup),
//   	accessLogFormat: apigateway.*accessLogFormat.jsonWithStandardFields(&jsonWithStandardFieldProps{
//   		caller: jsii.Boolean(false),
//   		httpMethod: jsii.Boolean(true),
//   		ip: jsii.Boolean(true),
//   		protocol: jsii.Boolean(true),
//   		requestTime: jsii.Boolean(true),
//   		resourcePath: jsii.Boolean(true),
//   		responseLength: jsii.Boolean(true),
//   		status: jsii.Boolean(true),
//   		user: jsii.Boolean(true),
//   	}),
//   })
//
type StageProps struct {
	// Indicates whether the cached responses are encrypted.
	CacheDataEncrypted *bool `field:"optional" json:"cacheDataEncrypted" yaml:"cacheDataEncrypted"`
	// Specifies the time to live (TTL), in seconds, for cached responses.
	//
	// The
	// higher the TTL, the longer the response will be cached.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-caching.html
	//
	CacheTtl awscdk.Duration `field:"optional" json:"cacheTtl" yaml:"cacheTtl"`
	// Specifies whether responses should be cached and returned for requests.
	//
	// A
	// cache cluster must be enabled on the stage for responses to be cached.
	CachingEnabled *bool `field:"optional" json:"cachingEnabled" yaml:"cachingEnabled"`
	// Specifies whether data trace logging is enabled for this method.
	//
	// When enabled, API gateway will log the full API requests and responses.
	// This can be useful to troubleshoot APIs, but can result in logging sensitive data.
	// We recommend that you don't enable this feature for production APIs.
	DataTraceEnabled *bool `field:"optional" json:"dataTraceEnabled" yaml:"dataTraceEnabled"`
	// Specifies the logging level for this method, which effects the log entries pushed to Amazon CloudWatch Logs.
	LoggingLevel MethodLoggingLevel `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
	// Specifies whether Amazon CloudWatch metrics are enabled for this method.
	MetricsEnabled *bool `field:"optional" json:"metricsEnabled" yaml:"metricsEnabled"`
	// Specifies the throttling burst limit.
	//
	// The total rate of all requests in your AWS account is limited to 5,000 requests.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html
	//
	ThrottlingBurstLimit *float64 `field:"optional" json:"throttlingBurstLimit" yaml:"throttlingBurstLimit"`
	// Specifies the throttling rate limit.
	//
	// The total rate of all requests in your AWS account is limited to 10,000 requests per second (rps).
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html
	//
	ThrottlingRateLimit *float64 `field:"optional" json:"throttlingRateLimit" yaml:"throttlingRateLimit"`
	// The CloudWatch Logs log group.
	AccessLogDestination IAccessLogDestination `field:"optional" json:"accessLogDestination" yaml:"accessLogDestination"`
	// A single line format of access logs of data, as specified by selected $content variables.
	//
	// The format must include at least `AccessLogFormat.contextRequestId()`.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html#context-variable-reference
	//
	AccessLogFormat AccessLogFormat `field:"optional" json:"accessLogFormat" yaml:"accessLogFormat"`
	// Indicates whether cache clustering is enabled for the stage.
	CacheClusterEnabled *bool `field:"optional" json:"cacheClusterEnabled" yaml:"cacheClusterEnabled"`
	// The stage's cache cluster size.
	CacheClusterSize *string `field:"optional" json:"cacheClusterSize" yaml:"cacheClusterSize"`
	// The identifier of the client certificate that API Gateway uses to call your integration endpoints in the stage.
	ClientCertificateId *string `field:"optional" json:"clientCertificateId" yaml:"clientCertificateId"`
	// A description of the purpose of the stage.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The version identifier of the API documentation snapshot.
	DocumentationVersion *string `field:"optional" json:"documentationVersion" yaml:"documentationVersion"`
	// Method deployment options for specific resources/methods.
	//
	// These will
	// override common options defined in `StageOptions#methodOptions`.
	MethodOptions *map[string]*MethodDeploymentOptions `field:"optional" json:"methodOptions" yaml:"methodOptions"`
	// The name of the stage, which API Gateway uses as the first path segment in the invoked Uniform Resource Identifier (URI).
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
	// Specifies whether Amazon X-Ray tracing is enabled for this method.
	TracingEnabled *bool `field:"optional" json:"tracingEnabled" yaml:"tracingEnabled"`
	// A map that defines the stage variables.
	//
	// Variable names must consist of
	// alphanumeric characters, and the values must match the following regular
	// expression: [A-Za-z0-9-._~:/?#&amp;=,]+.
	Variables *map[string]*string `field:"optional" json:"variables" yaml:"variables"`
	// The deployment that this stage points to [disable-awslint:ref-via-interface].
	Deployment Deployment `field:"required" json:"deployment" yaml:"deployment"`
}

// Options when configuring Step Functions synchronous integration with Rest API.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//   var vpcLink vpcLink
//
//   stepFunctionsExecutionIntegrationOptions := &stepFunctionsExecutionIntegrationOptions{
//   	authorizer: jsii.Boolean(false),
//   	cacheKeyParameters: []*string{
//   		jsii.String("cacheKeyParameters"),
//   	},
//   	cacheNamespace: jsii.String("cacheNamespace"),
//   	connectionType: awscdk.Aws_apigateway.connectionType_INTERNET,
//   	contentHandling: awscdk.*Aws_apigateway.contentHandling_CONVERT_TO_BINARY,
//   	credentialsPassthrough: jsii.Boolean(false),
//   	credentialsRole: role,
//   	headers: jsii.Boolean(false),
//   	integrationResponses: []integrationResponse{
//   		&integrationResponse{
//   			statusCode: jsii.String("statusCode"),
//
//   			// the properties below are optional
//   			contentHandling: awscdk.*Aws_apigateway.*contentHandling_CONVERT_TO_BINARY,
//   			responseParameters: map[string]*string{
//   				"responseParametersKey": jsii.String("responseParameters"),
//   			},
//   			responseTemplates: map[string]*string{
//   				"responseTemplatesKey": jsii.String("responseTemplates"),
//   			},
//   			selectionPattern: jsii.String("selectionPattern"),
//   		},
//   	},
//   	passthroughBehavior: awscdk.*Aws_apigateway.passthroughBehavior_WHEN_NO_MATCH,
//   	path: jsii.Boolean(false),
//   	querystring: jsii.Boolean(false),
//   	requestContext: &requestContext{
//   		accountId: jsii.Boolean(false),
//   		apiId: jsii.Boolean(false),
//   		apiKey: jsii.Boolean(false),
//   		authorizerPrincipalId: jsii.Boolean(false),
//   		caller: jsii.Boolean(false),
//   		cognitoAuthenticationProvider: jsii.Boolean(false),
//   		cognitoAuthenticationType: jsii.Boolean(false),
//   		cognitoIdentityId: jsii.Boolean(false),
//   		cognitoIdentityPoolId: jsii.Boolean(false),
//   		httpMethod: jsii.Boolean(false),
//   		requestId: jsii.Boolean(false),
//   		resourceId: jsii.Boolean(false),
//   		resourcePath: jsii.Boolean(false),
//   		sourceIp: jsii.Boolean(false),
//   		stage: jsii.Boolean(false),
//   		user: jsii.Boolean(false),
//   		userAgent: jsii.Boolean(false),
//   		userArn: jsii.Boolean(false),
//   	},
//   	requestParameters: map[string]*string{
//   		"requestParametersKey": jsii.String("requestParameters"),
//   	},
//   	requestTemplates: map[string]*string{
//   		"requestTemplatesKey": jsii.String("requestTemplates"),
//   	},
//   	timeout: cdk.duration.minutes(jsii.Number(30)),
//   	vpcLink: vpcLink,
//   }
//
type StepFunctionsExecutionIntegrationOptions struct {
	// A list of request parameters whose values are to be cached.
	//
	// It determines
	// request parameters that will make it into the cache key.
	CacheKeyParameters *[]*string `field:"optional" json:"cacheKeyParameters" yaml:"cacheKeyParameters"`
	// An API-specific tag group of related cached parameters.
	CacheNamespace *string `field:"optional" json:"cacheNamespace" yaml:"cacheNamespace"`
	// The type of network connection to the integration endpoint.
	ConnectionType ConnectionType `field:"optional" json:"connectionType" yaml:"connectionType"`
	// Specifies how to handle request payload content type conversions.
	ContentHandling ContentHandling `field:"optional" json:"contentHandling" yaml:"contentHandling"`
	// Requires that the caller's identity be passed through from the request.
	CredentialsPassthrough *bool `field:"optional" json:"credentialsPassthrough" yaml:"credentialsPassthrough"`
	// An IAM role that API Gateway assumes.
	//
	// Mutually exclusive with `credentialsPassThrough`.
	CredentialsRole awsiam.IRole `field:"optional" json:"credentialsRole" yaml:"credentialsRole"`
	// The response that API Gateway provides after a method's backend completes processing a request.
	//
	// API Gateway intercepts the response from the
	// backend so that you can control how API Gateway surfaces backend
	// responses. For example, you can map the backend status codes to codes
	// that you define.
	IntegrationResponses *[]*IntegrationResponse `field:"optional" json:"integrationResponses" yaml:"integrationResponses"`
	// Specifies the pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the requestTemplates property on the Integration resource.
	//
	// There are three valid values: WHEN_NO_MATCH, WHEN_NO_TEMPLATES, and
	// NEVER.
	PassthroughBehavior PassthroughBehavior `field:"optional" json:"passthroughBehavior" yaml:"passthroughBehavior"`
	// The request parameters that API Gateway sends with the backend request.
	//
	// Specify request parameters as key-value pairs (string-to-string
	// mappings), with a destination as the key and a source as the value.
	//
	// Specify the destination by using the following pattern
	// integration.request.location.name, where location is querystring, path,
	// or header, and name is a valid, unique parameter name.
	//
	// The source must be an existing method request parameter or a static
	// value. You must enclose static values in single quotation marks and
	// pre-encode these values based on their destination in the request.
	RequestParameters *map[string]*string `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// A map of Apache Velocity templates that are applied on the request payload.
	//
	// The template that API Gateway uses is based on the value of the
	// Content-Type header that's sent by the client. The content type value is
	// the key, and the template is the value (specified as a string), such as
	// the following snippet:
	//
	// ```
	//    { "application/json": "{ \"statusCode\": 200 }" }
	// ```.
	// See: http://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html
	//
	RequestTemplates *map[string]*string `field:"optional" json:"requestTemplates" yaml:"requestTemplates"`
	// The maximum amount of time an integration will run before it returns without a response.
	//
	// Must be between 50 milliseconds and 29 seconds.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The VpcLink used for the integration.
	//
	// Required if connectionType is VPC_LINK.
	VpcLink IVpcLink `field:"optional" json:"vpcLink" yaml:"vpcLink"`
	// If the whole authorizer object, including custom context values should be in the execution input.
	//
	// The execution input will include a new key `authorizer`:
	//
	// {
	//    "body": {},
	//    "authorizer": {
	//      "key": "value"
	//    }
	// }.
	Authorizer *bool `field:"optional" json:"authorizer" yaml:"authorizer"`
	// Check if header is to be included inside the execution input.
	//
	// The execution input will include a new key `headers`:
	//
	// {
	//    "body": {},
	//    "headers": {
	//       "header1": "value",
	//       "header2": "value"
	//    }
	// }.
	Headers *bool `field:"optional" json:"headers" yaml:"headers"`
	// Check if path is to be included inside the execution input.
	//
	// The execution input will include a new key `path`:
	//
	// {
	//    "body": {},
	//    "path": {
	//      "resourceName": "resourceValue"
	//    }
	// }.
	Path *bool `field:"optional" json:"path" yaml:"path"`
	// Check if querystring is to be included inside the execution input.
	//
	// The execution input will include a new key `queryString`:
	//
	// {
	//    "body": {},
	//    "querystring": {
	//      "key": "value"
	//    }
	// }.
	Querystring *bool `field:"optional" json:"querystring" yaml:"querystring"`
	// Which details of the incoming request must be passed onto the underlying state machine, such as, account id, user identity, request id, etc.
	//
	// The execution input will include a new key `requestContext`:
	//
	// {
	//    "body": {},
	//    "requestContext": {
	//        "key": "value"
	//    }
	// }.
	RequestContext *RequestContext `field:"optional" json:"requestContext" yaml:"requestContext"`
}

// Options to integrate with various StepFunction API.
//
// Example:
//   stateMachine := stepfunctions.NewStateMachine(this, jsii.String("MyStateMachine"), &stateMachineProps{
//   	stateMachineType: stepfunctions.stateMachineType_EXPRESS,
//   	definition: stepfunctions.chain.start(stepfunctions.NewPass(this, jsii.String("Pass"))),
//   })
//
//   api := apigateway.NewRestApi(this, jsii.String("Api"), &restApiProps{
//   	restApiName: jsii.String("MyApi"),
//   })
//   api.root.addMethod(jsii.String("GET"), apigateway.stepFunctionsIntegration.startExecution(stateMachine))
//
type StepFunctionsIntegration interface {
}

// The jsii proxy struct for StepFunctionsIntegration
type jsiiProxy_StepFunctionsIntegration struct {
	_ byte // padding
}

func NewStepFunctionsIntegration() StepFunctionsIntegration {
	_init_.Initialize()

	j := jsiiProxy_StepFunctionsIntegration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.StepFunctionsIntegration",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewStepFunctionsIntegration_Override(s StepFunctionsIntegration) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.StepFunctionsIntegration",
		nil, // no parameters
		s,
	)
}

// Integrates a Synchronous Express State Machine from AWS Step Functions to an API Gateway method.
//
// Example:
//   stateMachine := stepfunctions.NewStateMachine(this, jsii.String("MyStateMachine"), &stateMachineProps{
//   	stateMachineType: stepfunctions.stateMachineType_EXPRESS,
//   	definition: stepfunctions.chain.start(stepfunctions.NewPass(this, jsii.String("Pass"))),
//   })
//
//   api := apigateway.NewRestApi(this, jsii.String("Api"), &restApiProps{
//   	restApiName: jsii.String("MyApi"),
//   })
//   api.root.addMethod(jsii.String("GET"), apigateway.stepFunctionsIntegration.startExecution(stateMachine))
//
func StepFunctionsIntegration_StartExecution(stateMachine awsstepfunctions.IStateMachine, options *StepFunctionsExecutionIntegrationOptions) AwsIntegration {
	_init_.Initialize()

	var returns AwsIntegration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.StepFunctionsIntegration",
		"startExecution",
		[]interface{}{stateMachine, options},
		&returns,
	)

	return returns
}

// Defines an API Gateway REST API with a Synchrounous Express State Machine as a proxy integration.
//
// Example:
//   stateMachineDefinition := stepfunctions.NewPass(this, jsii.String("PassState"))
//
//   stateMachine := stepfunctions.NewStateMachine(this, jsii.String("StateMachine"), &stateMachineProps{
//   	definition: stateMachineDefinition,
//   	stateMachineType: stepfunctions.stateMachineType_EXPRESS,
//   })
//
//   apigateway.NewStepFunctionsRestApi(this, jsii.String("StepFunctionsRestApi"), &stepFunctionsRestApiProps{
//   	deploy: jsii.Boolean(true),
//   	stateMachine: stateMachine,
//   })
//
type StepFunctionsRestApi interface {
	RestApi
	CloudWatchAccount() CfnAccount
	SetCloudWatchAccount(val CfnAccount)
	// API Gateway stage that points to the latest deployment (if defined).
	//
	// If `deploy` is disabled, you will need to explicitly assign this value in order to
	// set up integrations.
	DeploymentStage() Stage
	SetDeploymentStage(val Stage)
	// The first domain name mapped to this API, if defined through the `domainName` configuration prop, or added via `addDomainName`.
	DomainName() DomainName
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// API Gateway deployment that represents the latest changes of the API.
	//
	// This resource will be automatically updated every time the REST API model changes.
	// This will be undefined if `deploy` is false.
	LatestDeployment() Deployment
	// The list of methods bound to this RestApi.
	Methods() *[]Method
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The ID of this API Gateway RestApi.
	RestApiId() *string
	// A human friendly name for this Rest API.
	//
	// Note that this is different from `restApiId`.
	RestApiName() *string
	// The resource ID of the root resource.
	RestApiRootResourceId() *string
	// Represents the root resource of this API endpoint ('/').
	//
	// Resources and Methods are added to this resource.
	Root() IResource
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The deployed root URL of this REST API.
	Url() *string
	// Add an ApiKey.
	AddApiKey(id *string, options *ApiKeyOptions) IApiKey
	// Defines an API Gateway domain name and maps it to this API.
	AddDomainName(id *string, options *DomainNameOptions) DomainName
	// Adds a new gateway response.
	AddGatewayResponse(id *string, options *GatewayResponseOptions) GatewayResponse
	// Adds a new model.
	AddModel(id *string, props *ModelOptions) Model
	// Adds a new request validator.
	AddRequestValidator(id *string, props *RequestValidatorOptions) RequestValidator
	// Adds a usage plan.
	AddUsagePlan(id *string, props *UsagePlanProps) UsagePlan
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Gets the "execute-api" ARN.
	ArnForExecuteApi(method *string, path *string, stage *string) *string
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns the given named metric for this API.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of requests served from the API cache in a given period.
	//
	// Default: sum over 5 minutes.
	MetricCacheHitCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of requests served from the backend in a given period, when API caching is enabled.
	//
	// Default: sum over 5 minutes.
	MetricCacheMissCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of client-side errors captured in a given period.
	//
	// Default: sum over 5 minutes.
	MetricClientError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the total number API requests in a given period.
	//
	// Default: sample count over 5 minutes.
	MetricCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the time between when API Gateway relays a request to the backend and when it receives a response from the backend.
	//
	// Default: average over 5 minutes.
	MetricIntegrationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time between when API Gateway receives a request from a client and when it returns a response to the client.
	//
	// The latency includes the integration latency and other API Gateway overhead.
	//
	// Default: average over 5 minutes.
	MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of server-side errors captured in a given period.
	//
	// Default: sum over 5 minutes.
	MetricServerError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns a string representation of this construct.
	ToString() *string
	// Returns the URL for an HTTP path.
	//
	// Fails if `deploymentStage` is not set either by `deploy` or explicitly.
	UrlForPath(path *string) *string
}

// The jsii proxy struct for StepFunctionsRestApi
type jsiiProxy_StepFunctionsRestApi struct {
	jsiiProxy_RestApi
}

func (j *jsiiProxy_StepFunctionsRestApi) CloudWatchAccount() CfnAccount {
	var returns CfnAccount
	_jsii_.Get(
		j,
		"cloudWatchAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsRestApi) DeploymentStage() Stage {
	var returns Stage
	_jsii_.Get(
		j,
		"deploymentStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsRestApi) DomainName() DomainName {
	var returns DomainName
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsRestApi) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsRestApi) LatestDeployment() Deployment {
	var returns Deployment
	_jsii_.Get(
		j,
		"latestDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsRestApi) Methods() *[]Method {
	var returns *[]Method
	_jsii_.Get(
		j,
		"methods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsRestApi) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsRestApi) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsRestApi) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsRestApi) RestApiName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsRestApi) RestApiRootResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiRootResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsRestApi) Root() IResource {
	var returns IResource
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsRestApi) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsRestApi) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}


func NewStepFunctionsRestApi(scope constructs.Construct, id *string, props *StepFunctionsRestApiProps) StepFunctionsRestApi {
	_init_.Initialize()

	j := jsiiProxy_StepFunctionsRestApi{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.StepFunctionsRestApi",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewStepFunctionsRestApi_Override(s StepFunctionsRestApi, scope constructs.Construct, id *string, props *StepFunctionsRestApiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.StepFunctionsRestApi",
		[]interface{}{scope, id, props},
		s,
	)
}

func (j *jsiiProxy_StepFunctionsRestApi) SetCloudWatchAccount(val CfnAccount) {
	_jsii_.Set(
		j,
		"cloudWatchAccount",
		val,
	)
}

func (j *jsiiProxy_StepFunctionsRestApi) SetDeploymentStage(val Stage) {
	_jsii_.Set(
		j,
		"deploymentStage",
		val,
	)
}

// Import an existing RestApi that can be configured with additional Methods and Resources.
func StepFunctionsRestApi_FromRestApiAttributes(scope constructs.Construct, id *string, attrs *RestApiAttributes) IRestApi {
	_init_.Initialize()

	var returns IRestApi

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.StepFunctionsRestApi",
		"fromRestApiAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Import an existing RestApi.
func StepFunctionsRestApi_FromRestApiId(scope constructs.Construct, id *string, restApiId *string) IRestApi {
	_init_.Initialize()

	var returns IRestApi

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.StepFunctionsRestApi",
		"fromRestApiId",
		[]interface{}{scope, id, restApiId},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func StepFunctionsRestApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.StepFunctionsRestApi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func StepFunctionsRestApi_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.StepFunctionsRestApi",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func StepFunctionsRestApi_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.StepFunctionsRestApi",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionsRestApi) AddApiKey(id *string, options *ApiKeyOptions) IApiKey {
	var returns IApiKey

	_jsii_.Invoke(
		s,
		"addApiKey",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionsRestApi) AddDomainName(id *string, options *DomainNameOptions) DomainName {
	var returns DomainName

	_jsii_.Invoke(
		s,
		"addDomainName",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionsRestApi) AddGatewayResponse(id *string, options *GatewayResponseOptions) GatewayResponse {
	var returns GatewayResponse

	_jsii_.Invoke(
		s,
		"addGatewayResponse",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionsRestApi) AddModel(id *string, props *ModelOptions) Model {
	var returns Model

	_jsii_.Invoke(
		s,
		"addModel",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionsRestApi) AddRequestValidator(id *string, props *RequestValidatorOptions) RequestValidator {
	var returns RequestValidator

	_jsii_.Invoke(
		s,
		"addRequestValidator",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionsRestApi) AddUsagePlan(id *string, props *UsagePlanProps) UsagePlan {
	var returns UsagePlan

	_jsii_.Invoke(
		s,
		"addUsagePlan",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionsRestApi) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_StepFunctionsRestApi) ArnForExecuteApi(method *string, path *string, stage *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"arnForExecuteApi",
		[]interface{}{method, path, stage},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionsRestApi) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionsRestApi) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionsRestApi) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionsRestApi) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionsRestApi) MetricCacheHitCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricCacheHitCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionsRestApi) MetricCacheMissCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricCacheMissCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionsRestApi) MetricClientError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricClientError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionsRestApi) MetricCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionsRestApi) MetricIntegrationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricIntegrationLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionsRestApi) MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionsRestApi) MetricServerError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricServerError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionsRestApi) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionsRestApi) UrlForPath(path *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"urlForPath",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Properties for StepFunctionsRestApi.
//
// Example:
//   stateMachineDefinition := stepfunctions.NewPass(this, jsii.String("PassState"))
//
//   stateMachine := stepfunctions.NewStateMachine(this, jsii.String("StateMachine"), &stateMachineProps{
//   	definition: stateMachineDefinition,
//   	stateMachineType: stepfunctions.stateMachineType_EXPRESS,
//   })
//
//   apigateway.NewStepFunctionsRestApi(this, jsii.String("StepFunctionsRestApi"), &stepFunctionsRestApiProps{
//   	deploy: jsii.Boolean(true),
//   	stateMachine: stateMachine,
//   })
//
type StepFunctionsRestApiProps struct {
	// Adds a CORS preflight OPTIONS method to this resource and all child resources.
	//
	// You can add CORS at the resource-level using `addCorsPreflight`.
	DefaultCorsPreflightOptions *CorsOptions `field:"optional" json:"defaultCorsPreflightOptions" yaml:"defaultCorsPreflightOptions"`
	// An integration to use as a default for all methods created within this API unless an integration is specified.
	DefaultIntegration Integration `field:"optional" json:"defaultIntegration" yaml:"defaultIntegration"`
	// Method options to use as a default for all methods created within this API unless custom options are specified.
	DefaultMethodOptions *MethodOptions `field:"optional" json:"defaultMethodOptions" yaml:"defaultMethodOptions"`
	// Automatically configure an AWS CloudWatch role for API Gateway.
	CloudWatchRole *bool `field:"optional" json:"cloudWatchRole" yaml:"cloudWatchRole"`
	// Indicates if a Deployment should be automatically created for this API, and recreated when the API model (resources, methods) changes.
	//
	// Since API Gateway deployments are immutable, When this option is enabled
	// (by default), an AWS::ApiGateway::Deployment resource will automatically
	// created with a logical ID that hashes the API model (methods, resources
	// and options). This means that when the model changes, the logical ID of
	// this CloudFormation resource will change, and a new deployment will be
	// created.
	//
	// If this is set, `latestDeployment` will refer to the `Deployment` object
	// and `deploymentStage` will refer to a `Stage` that points to this
	// deployment. To customize the stage options, use the `deployOptions`
	// property.
	//
	// A CloudFormation Output will also be defined with the root URL endpoint
	// of this REST API.
	Deploy *bool `field:"optional" json:"deploy" yaml:"deploy"`
	// Options for the API Gateway stage that will always point to the latest deployment when `deploy` is enabled.
	//
	// If `deploy` is disabled,
	// this value cannot be set.
	DeployOptions *StageOptions `field:"optional" json:"deployOptions" yaml:"deployOptions"`
	// A description of the RestApi construct.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether clients can invoke the API using the default execute-api endpoint.
	//
	// To require that clients use a custom domain name to invoke the
	// API, disable the default endpoint.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html
	//
	DisableExecuteApiEndpoint *bool `field:"optional" json:"disableExecuteApiEndpoint" yaml:"disableExecuteApiEndpoint"`
	// Configure a custom domain name and map it to this API.
	DomainName *DomainNameOptions `field:"optional" json:"domainName" yaml:"domainName"`
	// Export name for the CfnOutput containing the API endpoint.
	EndpointExportName *string `field:"optional" json:"endpointExportName" yaml:"endpointExportName"`
	// A list of the endpoint types of the API.
	//
	// Use this property when creating
	// an API.
	EndpointTypes *[]EndpointType `field:"optional" json:"endpointTypes" yaml:"endpointTypes"`
	// Indicates whether to roll back the resource if a warning occurs while API Gateway is creating the RestApi resource.
	FailOnWarnings *bool `field:"optional" json:"failOnWarnings" yaml:"failOnWarnings"`
	// Custom header parameters for the request.
	// See: https://docs.aws.amazon.com/cli/latest/reference/apigateway/import-rest-api.html
	//
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// A policy document that contains the permissions for this RestApi.
	Policy awsiam.PolicyDocument `field:"optional" json:"policy" yaml:"policy"`
	// A name for the API Gateway RestApi resource.
	RestApiName *string `field:"optional" json:"restApiName" yaml:"restApiName"`
	// Retains old deployment resources when the API changes.
	//
	// This allows
	// manually reverting stages to point to old deployments via the AWS
	// Console.
	RetainDeployments *bool `field:"optional" json:"retainDeployments" yaml:"retainDeployments"`
	// The source of the API key for metering requests according to a usage plan.
	ApiKeySourceType ApiKeySourceType `field:"optional" json:"apiKeySourceType" yaml:"apiKeySourceType"`
	// The list of binary media mime-types that are supported by the RestApi resource, such as "image/png" or "application/octet-stream".
	BinaryMediaTypes *[]*string `field:"optional" json:"binaryMediaTypes" yaml:"binaryMediaTypes"`
	// The ID of the API Gateway RestApi resource that you want to clone.
	CloneFrom IRestApi `field:"optional" json:"cloneFrom" yaml:"cloneFrom"`
	// The EndpointConfiguration property type specifies the endpoint types of a REST API.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-restapi-endpointconfiguration.html
	//
	EndpointConfiguration *EndpointConfiguration `field:"optional" json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// A nullable integer that is used to enable compression (with non-negative between 0 and 10485760 (10M) bytes, inclusive) or disable compression (when undefined) on an API.
	//
	// When compression is enabled, compression or
	// decompression is not applied on the payload if the payload size is
	// smaller than this value. Setting it to zero allows compression for any
	// payload size.
	MinimumCompressionSize *float64 `field:"optional" json:"minimumCompressionSize" yaml:"minimumCompressionSize"`
	// The default State Machine that handles all requests from this API.
	//
	// This stateMachine will be used as a the default integration for all methods in
	// this API, unless specified otherwise in `addMethod`.
	StateMachine awsstepfunctions.IStateMachine `field:"required" json:"stateMachine" yaml:"stateMachine"`
	// If the whole authorizer object, including custom context values should be in the execution input.
	//
	// The execution input will include a new key `authorizer`:
	//
	// {
	//    "body": {},
	//    "authorizer": {
	//      "key": "value"
	//    }
	// }.
	Authorizer *bool `field:"optional" json:"authorizer" yaml:"authorizer"`
	// Check if header is to be included inside the execution input.
	//
	// The execution input will include a new key `headers`:
	//
	// {
	//    "body": {},
	//    "headers": {
	//       "header1": "value",
	//       "header2": "value"
	//    }
	// }.
	Headers *bool `field:"optional" json:"headers" yaml:"headers"`
	// Check if path is to be included inside the execution input.
	//
	// The execution input will include a new key `path`:
	//
	// {
	//    "body": {},
	//    "path": {
	//      "resourceName": "resourceValue"
	//    }
	// }.
	Path *bool `field:"optional" json:"path" yaml:"path"`
	// Check if querystring is to be included inside the execution input.
	//
	// The execution input will include a new key `queryString`:
	//
	// {
	//    "body": {},
	//    "querystring": {
	//      "key": "value"
	//    }
	// }.
	Querystring *bool `field:"optional" json:"querystring" yaml:"querystring"`
	// Which details of the incoming request must be passed onto the underlying state machine, such as, account id, user identity, request id, etc.
	//
	// The execution input will include a new key `requestContext`:
	//
	// {
	//    "body": {},
	//    "requestContext": {
	//        "key": "value"
	//    }
	// }.
	RequestContext *RequestContext `field:"optional" json:"requestContext" yaml:"requestContext"`
	// An IAM role that API Gateway will assume to start the execution of the state machine.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

// Container for defining throttling parameters to API stages or methods.
//
// Example:
//   var integration lambdaIntegration
//
//
//   api := apigateway.NewRestApi(this, jsii.String("hello-api"))
//
//   v1 := api.root.addResource(jsii.String("v1"))
//   echo := v1.addResource(jsii.String("echo"))
//   echoMethod := echo.addMethod(jsii.String("GET"), integration, &methodOptions{
//   	apiKeyRequired: jsii.Boolean(true),
//   })
//
//   plan := api.addUsagePlan(jsii.String("UsagePlan"), &usagePlanProps{
//   	name: jsii.String("Easy"),
//   	throttle: &throttleSettings{
//   		rateLimit: jsii.Number(10),
//   		burstLimit: jsii.Number(2),
//   	},
//   })
//
//   key := api.addApiKey(jsii.String("ApiKey"))
//   plan.addApiKey(key)
//
type ThrottleSettings struct {
	// The maximum API request rate limit over a time ranging from one to a few seconds.
	BurstLimit *float64 `field:"optional" json:"burstLimit" yaml:"burstLimit"`
	// The API request steady-state rate limit (average requests per second over an extended period of time).
	RateLimit *float64 `field:"optional" json:"rateLimit" yaml:"rateLimit"`
}

// Represents per-method throttling for a resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var method method
//
//   throttlingPerMethod := &throttlingPerMethod{
//   	method: method,
//   	throttle: &throttleSettings{
//   		burstLimit: jsii.Number(123),
//   		rateLimit: jsii.Number(123),
//   	},
//   }
//
type ThrottlingPerMethod struct {
	// [disable-awslint:ref-via-interface] The method for which you specify the throttling settings.
	Method Method `field:"required" json:"method" yaml:"method"`
	// Specifies the overall request rate (average requests per second) and burst capacity.
	Throttle *ThrottleSettings `field:"required" json:"throttle" yaml:"throttle"`
}

// Token based lambda authorizer that recognizes the caller's identity as a bearer token, such as a JSON Web Token (JWT) or an OAuth token.
//
// Based on the token, authorization is performed by a lambda function.
//
// Example:
//   var authFn function
//   var books resource
//
//
//   auth := apigateway.NewTokenAuthorizer(this, jsii.String("booksAuthorizer"), &tokenAuthorizerProps{
//   	handler: authFn,
//   })
//
//   books.addMethod(jsii.String("GET"), apigateway.NewHttpIntegration(jsii.String("http://amazon.com")), &methodOptions{
//   	authorizer: auth,
//   })
//
type TokenAuthorizer interface {
	Authorizer
	IAuthorizer
	// The authorization type of this authorizer.
	AuthorizationType() AuthorizationType
	// The ARN of the authorizer to be used in permission policies, such as IAM and resource-based grants.
	AuthorizerArn() *string
	// The id of the authorizer.
	AuthorizerId() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The Lambda function handler that this authorizer uses.
	Handler() awslambda.IFunction
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	RestApiId() *string
	SetRestApiId(val *string)
	// The IAM role that the API Gateway service assumes while invoking the Lambda function.
	Role() awsiam.IRole
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a token that resolves to the Rest Api Id at the time of synthesis.
	//
	// Throws an error, during token resolution, if no RestApi is attached to this authorizer.
	LazyRestApiId() *string
	// Sets up the permissions necessary for the API Gateway service to invoke the Lambda function.
	SetupPermissions()
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for TokenAuthorizer
type jsiiProxy_TokenAuthorizer struct {
	jsiiProxy_Authorizer
	jsiiProxy_IAuthorizer
}

func (j *jsiiProxy_TokenAuthorizer) AuthorizationType() AuthorizationType {
	var returns AuthorizationType
	_jsii_.Get(
		j,
		"authorizationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthorizer) AuthorizerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthorizer) AuthorizerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthorizer) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthorizer) Handler() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthorizer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthorizer) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthorizer) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthorizer) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthorizer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewTokenAuthorizer(scope constructs.Construct, id *string, props *TokenAuthorizerProps) TokenAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_TokenAuthorizer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.TokenAuthorizer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewTokenAuthorizer_Override(t TokenAuthorizer, scope constructs.Construct, id *string, props *TokenAuthorizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.TokenAuthorizer",
		[]interface{}{scope, id, props},
		t,
	)
}

func (j *jsiiProxy_TokenAuthorizer) SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

// Return whether the given object is an Authorizer.
func TokenAuthorizer_IsAuthorizer(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.TokenAuthorizer",
		"isAuthorizer",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func TokenAuthorizer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.TokenAuthorizer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func TokenAuthorizer_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.TokenAuthorizer",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func TokenAuthorizer_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.TokenAuthorizer",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenAuthorizer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		t,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (t *jsiiProxy_TokenAuthorizer) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenAuthorizer) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenAuthorizer) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenAuthorizer) LazyRestApiId() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"lazyRestApiId",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenAuthorizer) SetupPermissions() {
	_jsii_.InvokeVoid(
		t,
		"setupPermissions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TokenAuthorizer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for TokenAuthorizer.
//
// Example:
//   var authFn function
//   var books resource
//
//
//   auth := apigateway.NewTokenAuthorizer(this, jsii.String("booksAuthorizer"), &tokenAuthorizerProps{
//   	handler: authFn,
//   })
//
//   books.addMethod(jsii.String("GET"), apigateway.NewHttpIntegration(jsii.String("http://amazon.com")), &methodOptions{
//   	authorizer: auth,
//   })
//
type TokenAuthorizerProps struct {
	// The handler for the authorizer lambda function.
	//
	// The handler must follow a very specific protocol on the input it receives and the output it needs to produce.
	// API Gateway has documented the handler's input specification
	// {@link https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-lambda-authorizer-input.html | here} and output specification
	// {@link https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-lambda-authorizer-output.html | here}.
	Handler awslambda.IFunction `field:"required" json:"handler" yaml:"handler"`
	// An optional IAM role for APIGateway to assume before calling the Lambda-based authorizer.
	//
	// The IAM role must be
	// assumable by 'apigateway.amazonaws.com'.
	AssumeRole awsiam.IRole `field:"optional" json:"assumeRole" yaml:"assumeRole"`
	// An optional human friendly name for the authorizer.
	//
	// Note that, this is not the primary identifier of the authorizer.
	AuthorizerName *string `field:"optional" json:"authorizerName" yaml:"authorizerName"`
	// How long APIGateway should cache the results.
	//
	// Max 1 hour.
	// Disable caching by setting this to 0.
	ResultsCacheTtl awscdk.Duration `field:"optional" json:"resultsCacheTtl" yaml:"resultsCacheTtl"`
	// The request header mapping expression for the bearer token.
	//
	// This is typically passed as part of the header, in which case
	// this should be `method.request.header.Authorizer` where Authorizer is the header containing the bearer token.
	// See: https://docs.aws.amazon.com/apigateway/api-reference/link-relation/authorizer-create/#identitySource
	//
	IdentitySource *string `field:"optional" json:"identitySource" yaml:"identitySource"`
	// An optional regex to be matched against the authorization token.
	//
	// When matched the authorizer lambda is invoked,
	// otherwise a 401 Unauthorized is returned to the client.
	ValidationRegex *string `field:"optional" json:"validationRegex" yaml:"validationRegex"`
}

// Example:
//   var integration lambdaIntegration
//
//
//   api := apigateway.NewRestApi(this, jsii.String("hello-api"))
//
//   v1 := api.root.addResource(jsii.String("v1"))
//   echo := v1.addResource(jsii.String("echo"))
//   echoMethod := echo.addMethod(jsii.String("GET"), integration, &methodOptions{
//   	apiKeyRequired: jsii.Boolean(true),
//   })
//
//   plan := api.addUsagePlan(jsii.String("UsagePlan"), &usagePlanProps{
//   	name: jsii.String("Easy"),
//   	throttle: &throttleSettings{
//   		rateLimit: jsii.Number(10),
//   		burstLimit: jsii.Number(2),
//   	},
//   })
//
//   key := api.addApiKey(jsii.String("ApiKey"))
//   plan.addApiKey(key)
//
type UsagePlan interface {
	awscdk.Resource
	IUsagePlan
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Id of the usage plan.
	UsagePlanId() *string
	// Adds an ApiKey.
	AddApiKey(apiKey IApiKey, options *AddApiKeyOptions)
	// Adds an apiStage.
	AddApiStage(apiStage *UsagePlanPerApiStage)
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for UsagePlan
type jsiiProxy_UsagePlan struct {
	internal.Type__awscdkResource
	jsiiProxy_IUsagePlan
}

func (j *jsiiProxy_UsagePlan) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UsagePlan) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UsagePlan) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UsagePlan) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UsagePlan) UsagePlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usagePlanId",
		&returns,
	)
	return returns
}


func NewUsagePlan(scope constructs.Construct, id *string, props *UsagePlanProps) UsagePlan {
	_init_.Initialize()

	j := jsiiProxy_UsagePlan{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.UsagePlan",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewUsagePlan_Override(u UsagePlan, scope constructs.Construct, id *string, props *UsagePlanProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.UsagePlan",
		[]interface{}{scope, id, props},
		u,
	)
}

// Import an externally defined usage plan using its ARN.
func UsagePlan_FromUsagePlanId(scope constructs.Construct, id *string, usagePlanId *string) IUsagePlan {
	_init_.Initialize()

	var returns IUsagePlan

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.UsagePlan",
		"fromUsagePlanId",
		[]interface{}{scope, id, usagePlanId},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func UsagePlan_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.UsagePlan",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func UsagePlan_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.UsagePlan",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func UsagePlan_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.UsagePlan",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UsagePlan) AddApiKey(apiKey IApiKey, options *AddApiKeyOptions) {
	_jsii_.InvokeVoid(
		u,
		"addApiKey",
		[]interface{}{apiKey, options},
	)
}

func (u *jsiiProxy_UsagePlan) AddApiStage(apiStage *UsagePlanPerApiStage) {
	_jsii_.InvokeVoid(
		u,
		"addApiStage",
		[]interface{}{apiStage},
	)
}

func (u *jsiiProxy_UsagePlan) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		u,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (u *jsiiProxy_UsagePlan) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UsagePlan) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UsagePlan) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UsagePlan) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Represents the API stages that a usage plan applies to.
//
// Example:
//   var plan usagePlan
//   var api restApi
//   var echoMethod method
//
//
//   plan.addApiStage(&usagePlanPerApiStage{
//   	stage: api.deploymentStage,
//   	throttle: []throttlingPerMethod{
//   		&throttlingPerMethod{
//   			method: echoMethod,
//   			throttle: &throttleSettings{
//   				rateLimit: jsii.Number(10),
//   				burstLimit: jsii.Number(2),
//   			},
//   		},
//   	},
//   })
//
type UsagePlanPerApiStage struct {
	Api IRestApi `field:"optional" json:"api" yaml:"api"`
	// [disable-awslint:ref-via-interface].
	Stage Stage `field:"optional" json:"stage" yaml:"stage"`
	Throttle *[]*ThrottlingPerMethod `field:"optional" json:"throttle" yaml:"throttle"`
}

// Example:
//   var integration lambdaIntegration
//
//
//   api := apigateway.NewRestApi(this, jsii.String("hello-api"))
//
//   v1 := api.root.addResource(jsii.String("v1"))
//   echo := v1.addResource(jsii.String("echo"))
//   echoMethod := echo.addMethod(jsii.String("GET"), integration, &methodOptions{
//   	apiKeyRequired: jsii.Boolean(true),
//   })
//
//   plan := api.addUsagePlan(jsii.String("UsagePlan"), &usagePlanProps{
//   	name: jsii.String("Easy"),
//   	throttle: &throttleSettings{
//   		rateLimit: jsii.Number(10),
//   		burstLimit: jsii.Number(2),
//   	},
//   })
//
//   key := api.addApiKey(jsii.String("ApiKey"))
//   plan.addApiKey(key)
//
type UsagePlanProps struct {
	// API Stages to be associated with the usage plan.
	ApiStages *[]*UsagePlanPerApiStage `field:"optional" json:"apiStages" yaml:"apiStages"`
	// Represents usage plan purpose.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Name for this usage plan.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Number of requests clients can make in a given time period.
	Quota *QuotaSettings `field:"optional" json:"quota" yaml:"quota"`
	// Overall throttle settings for the API.
	Throttle *ThrottleSettings `field:"optional" json:"throttle" yaml:"throttle"`
}

// Define a new VPC Link Specifies an API Gateway VPC link for a RestApi to access resources in an Amazon Virtual Private Cloud (VPC).
//
// Example:
//   import elbv2 "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("VPC"))
//   nlb := elbv2.NewNetworkLoadBalancer(this, jsii.String("NLB"), &networkLoadBalancerProps{
//   	vpc: vpc,
//   })
//   link := apigateway.NewVpcLink(this, jsii.String("link"), &vpcLinkProps{
//   	targets: []iNetworkLoadBalancer{
//   		nlb,
//   	},
//   })
//
//   integration := apigateway.NewIntegration(&integrationProps{
//   	type: apigateway.integrationType_HTTP_PROXY,
//   	options: &integrationOptions{
//   		connectionType: apigateway.connectionType_VPC_LINK,
//   		vpcLink: link,
//   	},
//   })
//
type VpcLink interface {
	awscdk.Resource
	IVpcLink
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Physical ID of the VpcLink resource.
	VpcLinkId() *string
	AddTargets(targets ...awselasticloadbalancingv2.INetworkLoadBalancer)
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for VpcLink
type jsiiProxy_VpcLink struct {
	internal.Type__awscdkResource
	jsiiProxy_IVpcLink
}

func (j *jsiiProxy_VpcLink) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcLink) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcLink) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcLink) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcLink) VpcLinkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcLinkId",
		&returns,
	)
	return returns
}


func NewVpcLink(scope constructs.Construct, id *string, props *VpcLinkProps) VpcLink {
	_init_.Initialize()

	j := jsiiProxy_VpcLink{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.VpcLink",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewVpcLink_Override(v VpcLink, scope constructs.Construct, id *string, props *VpcLinkProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.VpcLink",
		[]interface{}{scope, id, props},
		v,
	)
}

// Import a VPC Link by its Id.
func VpcLink_FromVpcLinkId(scope constructs.Construct, id *string, vpcLinkId *string) IVpcLink {
	_init_.Initialize()

	var returns IVpcLink

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.VpcLink",
		"fromVpcLinkId",
		[]interface{}{scope, id, vpcLinkId},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func VpcLink_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.VpcLink",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func VpcLink_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.VpcLink",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func VpcLink_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.VpcLink",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcLink) AddTargets(targets ...awselasticloadbalancingv2.INetworkLoadBalancer) {
	args := []interface{}{}
	for _, a := range targets {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		v,
		"addTargets",
		args,
	)
}

func (v *jsiiProxy_VpcLink) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		v,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (v *jsiiProxy_VpcLink) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcLink) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcLink) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcLink) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a VpcLink.
//
// Example:
//   import elbv2 "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("VPC"))
//   nlb := elbv2.NewNetworkLoadBalancer(this, jsii.String("NLB"), &networkLoadBalancerProps{
//   	vpc: vpc,
//   })
//   link := apigateway.NewVpcLink(this, jsii.String("link"), &vpcLinkProps{
//   	targets: []iNetworkLoadBalancer{
//   		nlb,
//   	},
//   })
//
//   integration := apigateway.NewIntegration(&integrationProps{
//   	type: apigateway.integrationType_HTTP_PROXY,
//   	options: &integrationOptions{
//   		connectionType: apigateway.connectionType_VPC_LINK,
//   		vpcLink: link,
//   	},
//   })
//
type VpcLinkProps struct {
	// The description of the VPC link.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The network load balancers of the VPC targeted by the VPC link.
	//
	// The network load balancers must be owned by the same AWS account of the API owner.
	Targets *[]awselasticloadbalancingv2.INetworkLoadBalancer `field:"optional" json:"targets" yaml:"targets"`
	// The name used to label and identify the VPC link.
	VpcLinkName *string `field:"optional" json:"vpcLinkName" yaml:"vpcLinkName"`
}

