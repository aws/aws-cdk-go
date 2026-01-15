package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigatewayv2"
)

// Represents a reference to an HTTP API.
type IWebSocketApiRef interface {
	interfacesawsapigatewayv2.IApiRef
	// Indicates that this is a WebSocket API.
	//
	// Will always return true, but is necessary to prevent accidental structural
	// equality in TypeScript.
	IsWebsocketApi() *bool
}

// The jsii proxy for IWebSocketApiRef
type jsiiProxy_IWebSocketApiRef struct {
	internal.Type__interfacesawsapigatewayv2IApiRef
}

func (j *jsiiProxy_IWebSocketApiRef) IsWebsocketApi() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isWebsocketApi",
		&returns,
	)
	return returns
}

