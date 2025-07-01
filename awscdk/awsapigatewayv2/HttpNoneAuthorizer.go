package awsapigatewayv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Explicitly configure no authorizers on specific HTTP API routes.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   issuer := "https://test.us.auth0.com"
//   authorizer := awscdk.NewHttpJwtAuthorizer(jsii.String("DefaultAuthorizer"), issuer, &HttpJwtAuthorizerProps{
//   	JwtAudience: []*string{
//   		jsii.String("3131231"),
//   	},
//   })
//
//   api := apigwv2.NewHttpApi(this, jsii.String("HttpApi"), &HttpApiProps{
//   	DefaultAuthorizer: authorizer,
//   	DefaultAuthorizationScopes: []*string{
//   		jsii.String("read:books"),
//   	},
//   })
//
//   api.AddRoutes(&AddRoutesOptions{
//   	Integration: awscdk.NewHttpUrlIntegration(jsii.String("BooksIntegration"), jsii.String("https://get-books-proxy.example.com")),
//   	Path: jsii.String("/books"),
//   	Methods: []httpMethod{
//   		apigwv2.*httpMethod_GET,
//   	},
//   })
//
//   api.AddRoutes(&AddRoutesOptions{
//   	Integration: awscdk.NewHttpUrlIntegration(jsii.String("BooksIdIntegration"), jsii.String("https://get-books-proxy.example.com")),
//   	Path: jsii.String("/books/{id}"),
//   	Methods: []*httpMethod{
//   		apigwv2.*httpMethod_GET,
//   	},
//   })
//
//   api.AddRoutes(&AddRoutesOptions{
//   	Integration: awscdk.NewHttpUrlIntegration(jsii.String("BooksIntegration"), jsii.String("https://get-books-proxy.example.com")),
//   	Path: jsii.String("/books"),
//   	Methods: []*httpMethod{
//   		apigwv2.*httpMethod_POST,
//   	},
//   	AuthorizationScopes: []*string{
//   		jsii.String("write:books"),
//   	},
//   })
//
//   api.AddRoutes(&AddRoutesOptions{
//   	Integration: awscdk.NewHttpUrlIntegration(jsii.String("LoginIntegration"), jsii.String("https://get-books-proxy.example.com")),
//   	Path: jsii.String("/login"),
//   	Methods: []*httpMethod{
//   		apigwv2.*httpMethod_POST,
//   	},
//   	Authorizer: apigwv2.NewHttpNoneAuthorizer(),
//   })
//
type HttpNoneAuthorizer interface {
	IHttpRouteAuthorizer
	// The authorizationType used for IAM Authorizer.
	AuthorizationType() *string
	// Bind this authorizer to a specified Http route.
	Bind(_options *HttpRouteAuthorizerBindOptions) *HttpRouteAuthorizerConfig
}

// The jsii proxy struct for HttpNoneAuthorizer
type jsiiProxy_HttpNoneAuthorizer struct {
	jsiiProxy_IHttpRouteAuthorizer
}

func (j *jsiiProxy_HttpNoneAuthorizer) AuthorizationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationType",
		&returns,
	)
	return returns
}


func NewHttpNoneAuthorizer() HttpNoneAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_HttpNoneAuthorizer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.HttpNoneAuthorizer",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewHttpNoneAuthorizer_Override(h HttpNoneAuthorizer) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.HttpNoneAuthorizer",
		nil, // no parameters
		h,
	)
}

func (h *jsiiProxy_HttpNoneAuthorizer) Bind(_options *HttpRouteAuthorizerBindOptions) *HttpRouteAuthorizerConfig {
	if err := h.validateBindParameters(_options); err != nil {
		panic(err)
	}
	var returns *HttpRouteAuthorizerConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{_options},
		&returns,
	)

	return returns
}

