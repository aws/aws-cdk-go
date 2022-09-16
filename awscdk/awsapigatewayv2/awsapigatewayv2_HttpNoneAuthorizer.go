package awsapigatewayv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
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
//   authorizer := awscdk.NewHttpJwtAuthorizer(jsii.String("DefaultAuthorizer"), issuer, &httpJwtAuthorizerProps{
//   	jwtAudience: []*string{
//   		jsii.String("3131231"),
//   	},
//   })
//
//   api := apigwv2.NewHttpApi(this, jsii.String("HttpApi"), &httpApiProps{
//   	defaultAuthorizer: authorizer,
//   	defaultAuthorizationScopes: []*string{
//   		jsii.String("read:books"),
//   	},
//   })
//
//   api.addRoutes(&addRoutesOptions{
//   	integration: awscdk.NewHttpUrlIntegration(jsii.String("BooksIntegration"), jsii.String("https://get-books-proxy.myproxy.internal")),
//   	path: jsii.String("/books"),
//   	methods: []httpMethod{
//   		apigwv2.*httpMethod_GET,
//   	},
//   })
//
//   api.addRoutes(&addRoutesOptions{
//   	integration: awscdk.NewHttpUrlIntegration(jsii.String("BooksIdIntegration"), jsii.String("https://get-books-proxy.myproxy.internal")),
//   	path: jsii.String("/books/{id}"),
//   	methods: []*httpMethod{
//   		apigwv2.*httpMethod_GET,
//   	},
//   })
//
//   api.addRoutes(&addRoutesOptions{
//   	integration: awscdk.NewHttpUrlIntegration(jsii.String("BooksIntegration"), jsii.String("https://get-books-proxy.myproxy.internal")),
//   	path: jsii.String("/books"),
//   	methods: []*httpMethod{
//   		apigwv2.*httpMethod_POST,
//   	},
//   	authorizationScopes: []*string{
//   		jsii.String("write:books"),
//   	},
//   })
//
//   api.addRoutes(&addRoutesOptions{
//   	integration: awscdk.NewHttpUrlIntegration(jsii.String("LoginIntegration"), jsii.String("https://get-books-proxy.myproxy.internal")),
//   	path: jsii.String("/login"),
//   	methods: []*httpMethod{
//   		apigwv2.*httpMethod_POST,
//   	},
//   	authorizer: apigwv2.NewHttpNoneAuthorizer(),
//   })
//
// Experimental.
type HttpNoneAuthorizer interface {
	IHttpRouteAuthorizer
	// Bind this authorizer to a specified Http route.
	// Experimental.
	Bind(_arg *HttpRouteAuthorizerBindOptions) *HttpRouteAuthorizerConfig
}

// The jsii proxy struct for HttpNoneAuthorizer
type jsiiProxy_HttpNoneAuthorizer struct {
	jsiiProxy_IHttpRouteAuthorizer
}

// Experimental.
func NewHttpNoneAuthorizer() HttpNoneAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_HttpNoneAuthorizer{}

	_jsii_.Create(
		"monocdk.aws_apigatewayv2.HttpNoneAuthorizer",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewHttpNoneAuthorizer_Override(h HttpNoneAuthorizer) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigatewayv2.HttpNoneAuthorizer",
		nil, // no parameters
		h,
	)
}

func (h *jsiiProxy_HttpNoneAuthorizer) Bind(_arg *HttpRouteAuthorizerBindOptions) *HttpRouteAuthorizerConfig {
	if err := h.validateBindParameters(_arg); err != nil {
		panic(err)
	}
	var returns *HttpRouteAuthorizerConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{_arg},
		&returns,
	)

	return returns
}

