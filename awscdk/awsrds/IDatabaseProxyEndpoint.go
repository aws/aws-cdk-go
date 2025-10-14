package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
)

// A DB proxy endpoint.
type IDatabaseProxyEndpoint interface {
	awscdk.IResource
	// DB Proxy Endpoint ARN.
	DbProxyEndpointArn() *string
	// DB Proxy Endpoint Name.
	DbProxyEndpointName() *string
	// Endpoint.
	Endpoint() *string
}

// The jsii proxy for IDatabaseProxyEndpoint
type jsiiProxy_IDatabaseProxyEndpoint struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IDatabaseProxyEndpoint) DbProxyEndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbProxyEndpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseProxyEndpoint) DbProxyEndpointName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbProxyEndpointName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseProxyEndpoint) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

