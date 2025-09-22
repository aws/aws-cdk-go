package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An authorizer that can attach to an Http Route.
type IHttpRouteAuthorizer interface {
	// Bind this authorizer to a specified Http route.
	Bind(options *HttpRouteAuthorizerBindOptions) *HttpRouteAuthorizerConfig
}

// The jsii proxy for IHttpRouteAuthorizer
type jsiiProxy_IHttpRouteAuthorizer struct {
	_ byte // padding
}

func (i *jsiiProxy_IHttpRouteAuthorizer) Bind(options *HttpRouteAuthorizerBindOptions) *HttpRouteAuthorizerConfig {
	if err := i.validateBindParameters(options); err != nil {
		panic(err)
	}
	var returns *HttpRouteAuthorizerConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

