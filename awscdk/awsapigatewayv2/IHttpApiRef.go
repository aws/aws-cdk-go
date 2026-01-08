package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigatewayv2"
)

// Represents a reference to an HTTP API.
type IHttpApiRef interface {
	interfacesawsapigatewayv2.IApiRef
	// Indicates that this is an HTTP API.
	//
	// Will always return true, but is necessary to prevent accidental structural
	// equality in TypeScript.
	IsHttpApi() *bool
}

// The jsii proxy for IHttpApiRef
type jsiiProxy_IHttpApiRef struct {
	internal.Type__interfacesawsapigatewayv2IApiRef
}

func (j *jsiiProxy_IHttpApiRef) IsHttpApi() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isHttpApi",
		&returns,
	)
	return returns
}

