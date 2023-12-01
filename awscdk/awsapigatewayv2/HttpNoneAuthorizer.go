package awsapigatewayv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Explicitly configure no authorizers on specific HTTP API routes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpNoneAuthorizer := awscdk.Aws_apigatewayv2.NewHttpNoneAuthorizer()
//
type HttpNoneAuthorizer interface {
	IHttpRouteAuthorizer
	// Bind this authorizer to a specified Http route.
	Bind(_options *HttpRouteAuthorizerBindOptions) *HttpRouteAuthorizerConfig
}

// The jsii proxy struct for HttpNoneAuthorizer
type jsiiProxy_HttpNoneAuthorizer struct {
	jsiiProxy_IHttpRouteAuthorizer
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

