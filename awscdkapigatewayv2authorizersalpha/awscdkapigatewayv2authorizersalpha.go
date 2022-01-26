// Authorizers for AWS APIGateway V2
package awscdkapigatewayv2authorizersalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapigatewayv2authorizersalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2authorizersalpha/v2/internal"
)

// Authorize HTTP API Routes with IAM.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpIamAuthorizer interface {
	awscdkapigatewayv2alpha.IHttpRouteAuthorizer
	Bind(_options *awscdkapigatewayv2alpha.HttpRouteAuthorizerBindOptions) *awscdkapigatewayv2alpha.HttpRouteAuthorizerConfig
}

// The jsii proxy struct for HttpIamAuthorizer
type jsiiProxy_HttpIamAuthorizer struct {
	internal.Type__awscdkapigatewayv2alphaIHttpRouteAuthorizer
}

// Experimental.
func NewHttpIamAuthorizer() HttpIamAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_HttpIamAuthorizer{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.HttpIamAuthorizer",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewHttpIamAuthorizer_Override(h HttpIamAuthorizer) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.HttpIamAuthorizer",
		nil, // no parameters
		h,
	)
}

// (experimental) Bind this authorizer to a specified Http route.
// Experimental.
func (h *jsiiProxy_HttpIamAuthorizer) Bind(_options *awscdkapigatewayv2alpha.HttpRouteAuthorizerBindOptions) *awscdkapigatewayv2alpha.HttpRouteAuthorizerConfig {
	var returns *awscdkapigatewayv2alpha.HttpRouteAuthorizerConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{_options},
		&returns,
	)

	return returns
}

// Authorize Http Api routes on whether the requester is registered as part of an AWS Cognito user pool.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpJwtAuthorizer interface {
	awscdkapigatewayv2alpha.IHttpRouteAuthorizer
	Bind(options *awscdkapigatewayv2alpha.HttpRouteAuthorizerBindOptions) *awscdkapigatewayv2alpha.HttpRouteAuthorizerConfig
}

// The jsii proxy struct for HttpJwtAuthorizer
type jsiiProxy_HttpJwtAuthorizer struct {
	internal.Type__awscdkapigatewayv2alphaIHttpRouteAuthorizer
}

// Initialize a JWT authorizer to be bound with HTTP route.
// Experimental.
func NewHttpJwtAuthorizer(id *string, jwtIssuer *string, props *HttpJwtAuthorizerProps) HttpJwtAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_HttpJwtAuthorizer{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.HttpJwtAuthorizer",
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
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.HttpJwtAuthorizer",
		[]interface{}{id, jwtIssuer, props},
		h,
	)
}

// (experimental) Bind this authorizer to a specified Http route.
// Experimental.
func (h *jsiiProxy_HttpJwtAuthorizer) Bind(options *awscdkapigatewayv2alpha.HttpRouteAuthorizerBindOptions) *awscdkapigatewayv2alpha.HttpRouteAuthorizerConfig {
	var returns *awscdkapigatewayv2alpha.HttpRouteAuthorizerConfig

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
	JwtAudience *[]*string `json:"jwtAudience" yaml:"jwtAudience"`
	// The name of the authorizer.
	// Experimental.
	AuthorizerName *string `json:"authorizerName" yaml:"authorizerName"`
	// The identity source for which authorization is requested.
	// Experimental.
	IdentitySource *[]*string `json:"identitySource" yaml:"identitySource"`
}

// Authorize Http Api routes via a lambda function.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpLambdaAuthorizer interface {
	awscdkapigatewayv2alpha.IHttpRouteAuthorizer
	Bind(options *awscdkapigatewayv2alpha.HttpRouteAuthorizerBindOptions) *awscdkapigatewayv2alpha.HttpRouteAuthorizerConfig
}

// The jsii proxy struct for HttpLambdaAuthorizer
type jsiiProxy_HttpLambdaAuthorizer struct {
	internal.Type__awscdkapigatewayv2alphaIHttpRouteAuthorizer
}

// Initialize a lambda authorizer to be bound with HTTP route.
// Experimental.
func NewHttpLambdaAuthorizer(id *string, handler awslambda.IFunction, props *HttpLambdaAuthorizerProps) HttpLambdaAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_HttpLambdaAuthorizer{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.HttpLambdaAuthorizer",
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
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.HttpLambdaAuthorizer",
		[]interface{}{id, handler, props},
		h,
	)
}

// (experimental) Bind this authorizer to a specified Http route.
// Experimental.
func (h *jsiiProxy_HttpLambdaAuthorizer) Bind(options *awscdkapigatewayv2alpha.HttpRouteAuthorizerBindOptions) *awscdkapigatewayv2alpha.HttpRouteAuthorizerConfig {
	var returns *awscdkapigatewayv2alpha.HttpRouteAuthorizerConfig

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
	AuthorizerName *string `json:"authorizerName" yaml:"authorizerName"`
	// The identity source for which authorization is requested.
	// Experimental.
	IdentitySource *[]*string `json:"identitySource" yaml:"identitySource"`
	// The types of responses the lambda can return.
	//
	// If HttpLambdaResponseType.SIMPLE is included then
	// response format 2.0 will be used.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html#http-api-lambda-authorizer.payload-format-response
	//
	// Experimental.
	ResponseTypes *[]HttpLambdaResponseType `json:"responseTypes" yaml:"responseTypes"`
	// How long APIGateway should cache the results.
	//
	// Max 1 hour.
	// Disable caching by setting this to `Duration.seconds(0)`.
	// Experimental.
	ResultsCacheTtl awscdk.Duration `json:"resultsCacheTtl" yaml:"resultsCacheTtl"`
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
	awscdkapigatewayv2alpha.IHttpRouteAuthorizer
	Bind(options *awscdkapigatewayv2alpha.HttpRouteAuthorizerBindOptions) *awscdkapigatewayv2alpha.HttpRouteAuthorizerConfig
}

// The jsii proxy struct for HttpUserPoolAuthorizer
type jsiiProxy_HttpUserPoolAuthorizer struct {
	internal.Type__awscdkapigatewayv2alphaIHttpRouteAuthorizer
}

// Initialize a Cognito user pool authorizer to be bound with HTTP route.
// Experimental.
func NewHttpUserPoolAuthorizer(id *string, pool awscognito.IUserPool, props *HttpUserPoolAuthorizerProps) HttpUserPoolAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_HttpUserPoolAuthorizer{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.HttpUserPoolAuthorizer",
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
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.HttpUserPoolAuthorizer",
		[]interface{}{id, pool, props},
		h,
	)
}

// (experimental) Bind this authorizer to a specified Http route.
// Experimental.
func (h *jsiiProxy_HttpUserPoolAuthorizer) Bind(options *awscdkapigatewayv2alpha.HttpRouteAuthorizerBindOptions) *awscdkapigatewayv2alpha.HttpRouteAuthorizerConfig {
	var returns *awscdkapigatewayv2alpha.HttpRouteAuthorizerConfig

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
	AuthorizerName *string `json:"authorizerName" yaml:"authorizerName"`
	// The identity source for which authorization is requested.
	// Experimental.
	IdentitySource *[]*string `json:"identitySource" yaml:"identitySource"`
	// The user pool clients that should be used to authorize requests with the user pool.
	// Experimental.
	UserPoolClients *[]awscognito.IUserPoolClient `json:"userPoolClients" yaml:"userPoolClients"`
	// The AWS region in which the user pool is present.
	// Experimental.
	UserPoolRegion *string `json:"userPoolRegion" yaml:"userPoolRegion"`
}

// Authorize WebSocket Api routes via a lambda function.
//
// TODO: EXAMPLE
//
// Experimental.
type WebSocketLambdaAuthorizer interface {
	awscdkapigatewayv2alpha.IWebSocketRouteAuthorizer
	Bind(options *awscdkapigatewayv2alpha.WebSocketRouteAuthorizerBindOptions) *awscdkapigatewayv2alpha.WebSocketRouteAuthorizerConfig
}

// The jsii proxy struct for WebSocketLambdaAuthorizer
type jsiiProxy_WebSocketLambdaAuthorizer struct {
	internal.Type__awscdkapigatewayv2alphaIWebSocketRouteAuthorizer
}

// Experimental.
func NewWebSocketLambdaAuthorizer(id *string, handler awslambda.IFunction, props *WebSocketLambdaAuthorizerProps) WebSocketLambdaAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_WebSocketLambdaAuthorizer{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.WebSocketLambdaAuthorizer",
		[]interface{}{id, handler, props},
		&j,
	)

	return &j
}

// Experimental.
func NewWebSocketLambdaAuthorizer_Override(w WebSocketLambdaAuthorizer, id *string, handler awslambda.IFunction, props *WebSocketLambdaAuthorizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.WebSocketLambdaAuthorizer",
		[]interface{}{id, handler, props},
		w,
	)
}

// (experimental) Bind this authorizer to a specified WebSocket route.
// Experimental.
func (w *jsiiProxy_WebSocketLambdaAuthorizer) Bind(options *awscdkapigatewayv2alpha.WebSocketRouteAuthorizerBindOptions) *awscdkapigatewayv2alpha.WebSocketRouteAuthorizerConfig {
	var returns *awscdkapigatewayv2alpha.WebSocketRouteAuthorizerConfig

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
	AuthorizerName *string `json:"authorizerName" yaml:"authorizerName"`
	// The identity source for which authorization is requested.
	//
	// Request parameter match `'route.request.querystring|header.[a-zA-z0-9._-]+'`.
	// Staged variable match `'stageVariables.[a-zA-Z0-9._-]+'`.
	// Context parameter match `'context.[a-zA-Z0-9._-]+'`.
	// Experimental.
	IdentitySource *[]*string `json:"identitySource" yaml:"identitySource"`
}

