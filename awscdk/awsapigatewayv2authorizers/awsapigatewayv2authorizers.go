package awsapigatewayv2authorizers

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2authorizers/internal"
	"github.com/aws/aws-cdk-go/awscdk/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
)

// Authorize Http Api routes on whether the requester is registered as part of an AWS Cognito user pool.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpJwtAuthorizer interface {
	awsapigatewayv2.IHttpRouteAuthorizer
	Bind(options *awsapigatewayv2.HttpRouteAuthorizerBindOptions) *awsapigatewayv2.HttpRouteAuthorizerConfig
}

// The jsii proxy struct for HttpJwtAuthorizer
type jsiiProxy_HttpJwtAuthorizer struct {
	internal.Type__awsapigatewayv2IHttpRouteAuthorizer
}

// Initialize a JWT authorizer to be bound with HTTP route.
// Experimental.
func NewHttpJwtAuthorizer(id *string, jwtIssuer *string, props *HttpJwtAuthorizerProps) HttpJwtAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_HttpJwtAuthorizer{}

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_authorizers.HttpJwtAuthorizer",
		[]interface{}{id, jwtIssuer, props},
		&j,
	)

	return &j
}

// Initialize a JWT authorizer to be bound with HTTP route.
// Experimental.
func NewHttpJwtAuthorizer_Override(h HttpJwtAuthorizer, id *string, jwtIssuer *string, props *HttpJwtAuthorizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_authorizers.HttpJwtAuthorizer",
		[]interface{}{id, jwtIssuer, props},
		h,
	)
}

// Bind this authorizer to a specified Http route.
// Experimental.
func (h *jsiiProxy_HttpJwtAuthorizer) Bind(options *awsapigatewayv2.HttpRouteAuthorizerBindOptions) *awsapigatewayv2.HttpRouteAuthorizerConfig {
	var returns *awsapigatewayv2.HttpRouteAuthorizerConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Properties to initialize HttpJwtAuthorizer.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpJwtAuthorizerProps struct {
	// A list of the intended recipients of the JWT.
	//
	// A valid JWT must provide an aud that matches at least one entry in this list.
	// Experimental.
	JwtAudience *[]*string `json:"jwtAudience"`
	// The name of the authorizer.
	// Experimental.
	AuthorizerName *string `json:"authorizerName"`
	// The identity source for which authorization is requested.
	// Experimental.
	IdentitySource *[]*string `json:"identitySource"`
}

// Authorize Http Api routes via a lambda function.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpLambdaAuthorizer interface {
	awsapigatewayv2.IHttpRouteAuthorizer
	Bind(options *awsapigatewayv2.HttpRouteAuthorizerBindOptions) *awsapigatewayv2.HttpRouteAuthorizerConfig
}

// The jsii proxy struct for HttpLambdaAuthorizer
type jsiiProxy_HttpLambdaAuthorizer struct {
	internal.Type__awsapigatewayv2IHttpRouteAuthorizer
}

// Initialize a lambda authorizer to be bound with HTTP route.
// Experimental.
func NewHttpLambdaAuthorizer(id *string, handler awslambda.IFunction, props *HttpLambdaAuthorizerProps) HttpLambdaAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_HttpLambdaAuthorizer{}

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_authorizers.HttpLambdaAuthorizer",
		[]interface{}{id, handler, props},
		&j,
	)

	return &j
}

// Initialize a lambda authorizer to be bound with HTTP route.
// Experimental.
func NewHttpLambdaAuthorizer_Override(h HttpLambdaAuthorizer, id *string, handler awslambda.IFunction, props *HttpLambdaAuthorizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_authorizers.HttpLambdaAuthorizer",
		[]interface{}{id, handler, props},
		h,
	)
}

// Bind this authorizer to a specified Http route.
// Experimental.
func (h *jsiiProxy_HttpLambdaAuthorizer) Bind(options *awsapigatewayv2.HttpRouteAuthorizerBindOptions) *awsapigatewayv2.HttpRouteAuthorizerConfig {
	var returns *awsapigatewayv2.HttpRouteAuthorizerConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Properties to initialize HttpTokenAuthorizer.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpLambdaAuthorizerProps struct {
	// Friendly authorizer name.
	// Experimental.
	AuthorizerName *string `json:"authorizerName"`
	// The identity source for which authorization is requested.
	// Experimental.
	IdentitySource *[]*string `json:"identitySource"`
	// The types of responses the lambda can return.
	//
	// If HttpLambdaResponseType.SIMPLE is included then
	// response format 2.0 will be used.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html#http-api-lambda-authorizer.payload-format-response
	//
	// Experimental.
	ResponseTypes *[]HttpLambdaResponseType `json:"responseTypes"`
	// How long APIGateway should cache the results.
	//
	// Max 1 hour.
	// Disable caching by setting this to `Duration.seconds(0)`.
	// Experimental.
	ResultsCacheTtl awscdk.Duration `json:"resultsCacheTtl"`
}

// Specifies the type responses the lambda returns.
// Experimental.
type HttpLambdaResponseType string

const (
	HttpLambdaResponseType_SIMPLE HttpLambdaResponseType = "SIMPLE"
	HttpLambdaResponseType_IAM HttpLambdaResponseType = "IAM"
)

// Authorize Http Api routes on whether the requester is registered as part of an AWS Cognito user pool.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpUserPoolAuthorizer interface {
	awsapigatewayv2.IHttpRouteAuthorizer
	Bind(options *awsapigatewayv2.HttpRouteAuthorizerBindOptions) *awsapigatewayv2.HttpRouteAuthorizerConfig
}

// The jsii proxy struct for HttpUserPoolAuthorizer
type jsiiProxy_HttpUserPoolAuthorizer struct {
	internal.Type__awsapigatewayv2IHttpRouteAuthorizer
}

// Initialize a Cognito user pool authorizer to be bound with HTTP route.
// Experimental.
func NewHttpUserPoolAuthorizer(id *string, pool awscognito.IUserPool, props *HttpUserPoolAuthorizerProps) HttpUserPoolAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_HttpUserPoolAuthorizer{}

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_authorizers.HttpUserPoolAuthorizer",
		[]interface{}{id, pool, props},
		&j,
	)

	return &j
}

// Initialize a Cognito user pool authorizer to be bound with HTTP route.
// Experimental.
func NewHttpUserPoolAuthorizer_Override(h HttpUserPoolAuthorizer, id *string, pool awscognito.IUserPool, props *HttpUserPoolAuthorizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_authorizers.HttpUserPoolAuthorizer",
		[]interface{}{id, pool, props},
		h,
	)
}

// Bind this authorizer to a specified Http route.
// Experimental.
func (h *jsiiProxy_HttpUserPoolAuthorizer) Bind(options *awsapigatewayv2.HttpRouteAuthorizerBindOptions) *awsapigatewayv2.HttpRouteAuthorizerConfig {
	var returns *awsapigatewayv2.HttpRouteAuthorizerConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Properties to initialize HttpUserPoolAuthorizer.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpUserPoolAuthorizerProps struct {
	// Friendly name of the authorizer.
	// Experimental.
	AuthorizerName *string `json:"authorizerName"`
	// The identity source for which authorization is requested.
	// Experimental.
	IdentitySource *[]*string `json:"identitySource"`
	// The user pool clients that should be used to authorize requests with the user pool.
	// Experimental.
	UserPoolClients *[]awscognito.IUserPoolClient `json:"userPoolClients"`
	// The AWS region in which the user pool is present.
	// Experimental.
	UserPoolRegion *string `json:"userPoolRegion"`
}

// Authorize WebSocket Api routes via a lambda function.
//
// TODO: EXAMPLE
//
// Experimental.
type WebSocketLambdaAuthorizer interface {
	awsapigatewayv2.IWebSocketRouteAuthorizer
	Bind(options *awsapigatewayv2.WebSocketRouteAuthorizerBindOptions) *awsapigatewayv2.WebSocketRouteAuthorizerConfig
}

// The jsii proxy struct for WebSocketLambdaAuthorizer
type jsiiProxy_WebSocketLambdaAuthorizer struct {
	internal.Type__awsapigatewayv2IWebSocketRouteAuthorizer
}

// Experimental.
func NewWebSocketLambdaAuthorizer(id *string, handler awslambda.IFunction, props *WebSocketLambdaAuthorizerProps) WebSocketLambdaAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_WebSocketLambdaAuthorizer{}

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_authorizers.WebSocketLambdaAuthorizer",
		[]interface{}{id, handler, props},
		&j,
	)

	return &j
}

// Experimental.
func NewWebSocketLambdaAuthorizer_Override(w WebSocketLambdaAuthorizer, id *string, handler awslambda.IFunction, props *WebSocketLambdaAuthorizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_authorizers.WebSocketLambdaAuthorizer",
		[]interface{}{id, handler, props},
		w,
	)
}

// Bind this authorizer to a specified WebSocket route.
// Experimental.
func (w *jsiiProxy_WebSocketLambdaAuthorizer) Bind(options *awsapigatewayv2.WebSocketRouteAuthorizerBindOptions) *awsapigatewayv2.WebSocketRouteAuthorizerConfig {
	var returns *awsapigatewayv2.WebSocketRouteAuthorizerConfig

	_jsii_.Invoke(
		w,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Properties to initialize WebSocketTokenAuthorizer.
//
// TODO: EXAMPLE
//
// Experimental.
type WebSocketLambdaAuthorizerProps struct {
	// The name of the authorizer.
	// Experimental.
	AuthorizerName *string `json:"authorizerName"`
	// The identity source for which authorization is requested.
	// Experimental.
	IdentitySource *[]*string `json:"identitySource"`
}

