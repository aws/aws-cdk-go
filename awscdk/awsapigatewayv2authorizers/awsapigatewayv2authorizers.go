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

// Authorize HTTP API Routes with IAM.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"type HttpIamAuthorizer awscdk.HttpIamAuthorizerimport awscdk "github.com/aws/aws-cdk-go/awscdk"type HttpUrlIntegration awscdk.HttpUrlIntegration
//
//   var principal anyPrincipal
//
//   authorizer := NewHttpIamAuthorizer()
//
//   httpApi := apigwv2.NewHttpApi(this, jsii.String("HttpApi"), &httpApiProps{
//   	defaultAuthorizer: authorizer,
//   })
//
//   routes := httpApi.addRoutes(&addRoutesOptions{
//   	integration: NewHttpUrlIntegration(jsii.String("BooksIntegration"), jsii.String("https://get-books-proxy.myproxy.internal")),
//   	path: jsii.String("/books/{book}"),
//   })
//
//   routes[0].grantInvoke(principal)
//
// Experimental.
type HttpIamAuthorizer interface {
	awsapigatewayv2.IHttpRouteAuthorizer
	// Bind this authorizer to a specified Http route.
	// Experimental.
	Bind(_options *awsapigatewayv2.HttpRouteAuthorizerBindOptions) *awsapigatewayv2.HttpRouteAuthorizerConfig
}

// The jsii proxy struct for HttpIamAuthorizer
type jsiiProxy_HttpIamAuthorizer struct {
	internal.Type__awsapigatewayv2IHttpRouteAuthorizer
}

// Experimental.
func NewHttpIamAuthorizer() HttpIamAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_HttpIamAuthorizer{}

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_authorizers.HttpIamAuthorizer",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewHttpIamAuthorizer_Override(h HttpIamAuthorizer) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_authorizers.HttpIamAuthorizer",
		nil, // no parameters
		h,
	)
}

func (h *jsiiProxy_HttpIamAuthorizer) Bind(_options *awsapigatewayv2.HttpRouteAuthorizerBindOptions) *awsapigatewayv2.HttpRouteAuthorizerConfig {
	var returns *awsapigatewayv2.HttpRouteAuthorizerConfig

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
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"type HttpJwtAuthorizer awscdk.HttpJwtAuthorizerimport awscdk "github.com/aws/aws-cdk-go/awscdk"type HttpUrlIntegration awscdk.HttpUrlIntegration
//
//   issuer := "https://test.us.auth0.com"
//   authorizer := NewHttpJwtAuthorizer(jsii.String("BooksAuthorizer"), issuer, &httpJwtAuthorizerProps{
//   	jwtAudience: []*string{
//   		jsii.String("3131231"),
//   	},
//   })
//
//   api := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))
//
//   api.addRoutes(&addRoutesOptions{
//   	integration: NewHttpUrlIntegration(jsii.String("BooksIntegration"), jsii.String("https://get-books-proxy.myproxy.internal")),
//   	path: jsii.String("/books"),
//   	authorizer: authorizer,
//   })
//
// Experimental.
type HttpJwtAuthorizer interface {
	awsapigatewayv2.IHttpRouteAuthorizer
	// Bind this authorizer to a specified Http route.
	// Experimental.
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
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"type HttpJwtAuthorizer awscdk.HttpJwtAuthorizerimport awscdk "github.com/aws/aws-cdk-go/awscdk"type HttpUrlIntegration awscdk.HttpUrlIntegration
//
//   issuer := "https://test.us.auth0.com"
//   authorizer := NewHttpJwtAuthorizer(jsii.String("BooksAuthorizer"), issuer, &httpJwtAuthorizerProps{
//   	jwtAudience: []*string{
//   		jsii.String("3131231"),
//   	},
//   })
//
//   api := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))
//
//   api.addRoutes(&addRoutesOptions{
//   	integration: NewHttpUrlIntegration(jsii.String("BooksIntegration"), jsii.String("https://get-books-proxy.myproxy.internal")),
//   	path: jsii.String("/books"),
//   	authorizer: authorizer,
//   })
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
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"type HttpLambdaAuthorizer awscdk.HttpLambdaAuthorizer
//   type HttpLambdaResponseType awscdk.HttpLambdaResponseTypeimport awscdk "github.com/aws/aws-cdk-go/awscdk"type HttpUrlIntegration awscdk.HttpUrlIntegration
//
//   // This function handles your auth logic
//   var authHandler function
//
//   authorizer := NewHttpLambdaAuthorizer(jsii.String("BooksAuthorizer"), authHandler, &httpLambdaAuthorizerProps{
//   	responseTypes: []httpLambdaResponseType{
//   		httpLambdaResponseType_SIMPLE,
//   	},
//   })
//
//   api := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))
//
//   api.addRoutes(&addRoutesOptions{
//   	integration: NewHttpUrlIntegration(jsii.String("BooksIntegration"), jsii.String("https://get-books-proxy.myproxy.internal")),
//   	path: jsii.String("/books"),
//   	authorizer: authorizer,
//   })
//
// Experimental.
type HttpLambdaAuthorizer interface {
	awsapigatewayv2.IHttpRouteAuthorizer
	// Bind this authorizer to a specified Http route.
	// Experimental.
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
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"type HttpLambdaAuthorizer awscdk.HttpLambdaAuthorizer
//   type HttpLambdaResponseType awscdk.HttpLambdaResponseTypeimport awscdk "github.com/aws/aws-cdk-go/awscdk"type HttpUrlIntegration awscdk.HttpUrlIntegration
//
//   // This function handles your auth logic
//   var authHandler function
//
//   authorizer := NewHttpLambdaAuthorizer(jsii.String("BooksAuthorizer"), authHandler, &httpLambdaAuthorizerProps{
//   	responseTypes: []httpLambdaResponseType{
//   		httpLambdaResponseType_SIMPLE,
//   	},
//   })
//
//   api := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))
//
//   api.addRoutes(&addRoutesOptions{
//   	integration: NewHttpUrlIntegration(jsii.String("BooksIntegration"), jsii.String("https://get-books-proxy.myproxy.internal")),
//   	path: jsii.String("/books"),
//   	authorizer: authorizer,
//   })
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
	// Returns simple boolean response.
	// Experimental.
	HttpLambdaResponseType_SIMPLE HttpLambdaResponseType = "SIMPLE"
	// Returns an IAM Policy.
	// Experimental.
	HttpLambdaResponseType_IAM HttpLambdaResponseType = "IAM"
)

// Authorize Http Api routes on whether the requester is registered as part of an AWS Cognito user pool.
//
// Example:
//   import cognito "github.com/aws/aws-cdk-go/awscdk"import awscdk "github.com/aws/aws-cdk-go/awscdk"type HttpUserPoolAuthorizer awscdk.HttpUserPoolAuthorizerimport awscdk "github.com/aws/aws-cdk-go/awscdk"type HttpUrlIntegration awscdk.HttpUrlIntegration
//
//   userPool := cognito.NewUserPool(this, jsii.String("UserPool"))
//
//   authorizer := NewHttpUserPoolAuthorizer(jsii.String("BooksAuthorizer"), userPool)
//
//   api := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))
//
//   api.addRoutes(&addRoutesOptions{
//   	integration: NewHttpUrlIntegration(jsii.String("BooksIntegration"), jsii.String("https://get-books-proxy.myproxy.internal")),
//   	path: jsii.String("/books"),
//   	authorizer: authorizer,
//   })
//
// Experimental.
type HttpUserPoolAuthorizer interface {
	awsapigatewayv2.IHttpRouteAuthorizer
	// Bind this authorizer to a specified Http route.
	// Experimental.
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
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apigatewayv2_authorizers "github.com/aws/aws-cdk-go/awscdk/aws_apigatewayv2_authorizers"import awscdk "github.com/aws/aws-cdk-go/awscdk"import cognito "github.com/aws/aws-cdk-go/awscdk/aws_cognito"
//
//   var userPoolClient userPoolClient
//   httpUserPoolAuthorizerProps := &httpUserPoolAuthorizerProps{
//   	authorizerName: jsii.String("authorizerName"),
//   	identitySource: []*string{
//   		jsii.String("identitySource"),
//   	},
//   	userPoolClients: []iUserPoolClient{
//   		userPoolClient,
//   	},
//   	userPoolRegion: jsii.String("userPoolRegion"),
//   }
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
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"type WebSocketLambdaAuthorizer awscdk.WebSocketLambdaAuthorizerimport awscdk "github.com/aws/aws-cdk-go/awscdk"type WebSocketLambdaIntegration awscdk.WebSocketLambdaIntegration
//
//   // This function handles your auth logic
//   var authHandler function
//
//   // This function handles your WebSocket requests
//   var handler function
//
//   authorizer := NewWebSocketLambdaAuthorizer(jsii.String("Authorizer"), authHandler)
//
//   integration := NewWebSocketLambdaIntegration(jsii.String("Integration"), handler)
//
//   apigwv2.NewWebSocketApi(this, jsii.String("WebSocketApi"), &webSocketApiProps{
//   	connectRouteOptions: &webSocketRouteOptions{
//   		integration: integration,
//   		authorizer: authorizer,
//   	},
//   })
//
// Experimental.
type WebSocketLambdaAuthorizer interface {
	awsapigatewayv2.IWebSocketRouteAuthorizer
	// Bind this authorizer to a specified WebSocket route.
	// Experimental.
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
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import apigatewayv2_authorizers "github.com/aws/aws-cdk-go/awscdk/aws_apigatewayv2_authorizers"
//   webSocketLambdaAuthorizerProps := &webSocketLambdaAuthorizerProps{
//   	authorizerName: jsii.String("authorizerName"),
//   	identitySource: []*string{
//   		jsii.String("identitySource"),
//   	},
//   }
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

