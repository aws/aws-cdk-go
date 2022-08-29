// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An authorizer that can attach to an Http Route.
// Experimental.
type IHttpRouteAuthorizer interface {
	// Bind this authorizer to a specified Http route.
	// Experimental.
	Bind(options *HttpRouteAuthorizerBindOptions) *HttpRouteAuthorizerConfig
}

// The jsii proxy for IHttpRouteAuthorizer
type jsiiProxy_IHttpRouteAuthorizer struct {
	_ byte // padding
}

func (i *jsiiProxy_IHttpRouteAuthorizer) Bind(options *HttpRouteAuthorizerBindOptions) *HttpRouteAuthorizerConfig {
	var returns *HttpRouteAuthorizerConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

