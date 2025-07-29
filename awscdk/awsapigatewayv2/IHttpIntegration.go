package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents an Integration for an HTTP API.
type IHttpIntegration interface {
	IIntegration
	// The HTTP API associated with this integration.
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

