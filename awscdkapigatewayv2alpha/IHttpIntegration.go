// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents an Integration for an HTTP API.
// Experimental.
type IHttpIntegration interface {
	IIntegration
	// The HTTP API associated with this integration.
	// Experimental.
	HttpApi() IHttpApi
}

// The jsii proxy for IHttpIntegration
type jsiiProxy_IHttpIntegration struct {
	jsiiProxy_IIntegration
}

func (j *jsiiProxy_IHttpIntegration) HttpApi() IHttpApi {
	var returns IHttpApi
	_jsii_.Get(
		j,
		"httpApi",
		&returns,
	)
	return returns
}

