package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// $context variables that can be used to customize access log pattern.
//
// Example:
//   apigateway.AccessLogFormat_Custom(jSON.stringify(map[string]interface{}{
//   	"requestId": apigateway.AccessLogField_contextRequestId(),
//   	"sourceIp": apigateway.AccessLogField_contextIdentitySourceIp(),
//   	"method": apigateway.AccessLogField_contextHttpMethod(),
//   	"userContext": map[string]*string{
//   		"sub": apigateway.AccessLogField_contextAuthorizerClaims(jsii.String("sub")),
//   		"email": apigateway.AccessLogField_contextAuthorizerClaims(jsii.String("email")),
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

// The API callers AWS account ID.
// Deprecated: Use `contextCallerAccountId` or `contextOwnerAccountId` instead.
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

// The error message returned from an authentication attempt.
func AccessLogField_ContextAuthenticateError() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextAuthenticateError",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The authentication latency in ms.
func AccessLogField_ContextAuthenticateLatency() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextAuthenticateLatency",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The status code returned from an authentication attempt.
func AccessLogField_ContextAuthenticateStatus() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextAuthenticateStatus",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The authorization error message.
func AccessLogField_ContextAuthorizeError() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextAuthorizeError",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The authorization latency in ms.
func AccessLogField_ContextAuthorizeLatency() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextAuthorizeLatency",
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

	if err := validateAccessLogField_ContextAuthorizerParameters(property); err != nil {
		panic(err)
	}
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

	if err := validateAccessLogField_ContextAuthorizerClaimsParameters(property); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextAuthorizerClaims",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// The error message returned from an authorizer.
func AccessLogField_ContextAuthorizerError() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextAuthorizerError",
		nil, // no parameters
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

// The status code returned from a Lambda authorizer.
func AccessLogField_ContextAuthorizerIntegrationStatus() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextAuthorizerIntegrationStatus",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The authorizer latency in ms.
func AccessLogField_ContextAuthorizerLatency() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextAuthorizerLatency",
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
func AccessLogField_ContextAuthorizerRequestId() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextAuthorizerRequestId",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The status code returned from an authorizer.
func AccessLogField_ContextAuthorizerStatus() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextAuthorizerStatus",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The status code returned from an authorization attempt.
func AccessLogField_ContextAuthorizeStatus() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextAuthorizeStatus",
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

// The API callers AWS account ID.
func AccessLogField_ContextCallerAccountId() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextCallerAccountId",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The path for an API mapping that an incoming request matched.
//
// Applicable when a client uses a custom domain name to access an API. For example if a client sends a request to
// https://api.example.com/v1/orders/1234, and the request matches the API mapping with the path v1/orders, the value is v1/orders.
// See: https://docs.aws.amazon.com/en_jp/apigateway/latest/developerguide/rest-api-mappings.html
//
func AccessLogField_ContextCustomDomainBasePathMatched() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextCustomDomainBasePathMatched",
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

// The distinguished name of the issuer of the certificate that a client presents.
//
// Present when a client accesses an API by using a custom domain name that has mutual TLS enabled.
// Present only in access logs if mutual TLS authentication fails.
func AccessLogField_ContextIdentityClientCertIssunerDN() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextIdentityClientCertIssunerDN",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The PEM-encoded client certificate that the client presented during mutual TLS authentication.
//
// Present when a client accesses an API by using a custom domain name that has mutual TLS enabled.
// Present only in access logs if mutual TLS authentication fails.
func AccessLogField_ContextIdentityClientCertPem() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextIdentityClientCertPem",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The serial number of the certificate.
//
// Present when a client accesses an API by using a custom domain name that has mutual TLS enabled.
// Present only in access logs if mutual TLS authentication fails.
func AccessLogField_ContextIdentityClientCertSerialNumber() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextIdentityClientCertSerialNumber",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The distinguished name of the subject of the certificate that a client presents.
//
// Present when a client accesses an API by using a custom domain name that has mutual TLS enabled.
// Present only in access logs if mutual TLS authentication fails.
func AccessLogField_ContextIdentityClientCertSubjectDN() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextIdentityClientCertSubjectDN",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The date after which the certificate is invalid.
//
// Present when a client accesses an API by using a custom domain name that has mutual TLS enabled.
// Present only in access logs if mutual TLS authentication fails.
func AccessLogField_ContextIdentityClientCertValidityNotAfter() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextIdentityClientCertValidityNotAfter",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The date before which the certificate is invalid.
//
// Present when a client accesses an API by using a custom domain name that has mutual TLS enabled.
// Present only in access logs if mutual TLS authentication fails.
func AccessLogField_ContextIdentityClientCertValidityNotBefore() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextIdentityClientCertValidityNotBefore",
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

// A string that contains an integration error message.
func AccessLogField_ContextIntegrationErrorMessage() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextIntegrationErrorMessage",
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

// The API owner's AWS account ID.
func AccessLogField_ContextOwnerAccountId() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextOwnerAccountId",
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

	if err := validateAccessLogField_ContextRequestOverrideHeaderParameters(headerName); err != nil {
		panic(err)
	}
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

	if err := validateAccessLogField_ContextRequestOverridePathParameters(pathName); err != nil {
		panic(err)
	}
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

	if err := validateAccessLogField_ContextRequestOverrideQuerystringParameters(querystringName); err != nil {
		panic(err)
	}
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

	if err := validateAccessLogField_ContextResponseOverrideHeaderParameters(headerName); err != nil {
		panic(err)
	}
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

// The error message returned from AWS WAF.
func AccessLogField_ContextWafError() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextWafError",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The AWS WAF latency in ms.
func AccessLogField_ContextWafLatency() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextWafLatency",
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

// The status code returned from AWS WAF.
func AccessLogField_ContextWafStatus() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AccessLogField",
		"contextWafStatus",
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

