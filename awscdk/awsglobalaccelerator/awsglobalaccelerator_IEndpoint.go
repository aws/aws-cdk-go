package awsglobalaccelerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An endpoint for the endpoint group.
//
// Implementations of `IEndpoint` can be found in the `aws-globalaccelerator-endpoints` package.
type IEndpoint interface {
	// Render the endpoint to an endpoint configuration.
	RenderEndpointConfiguration() interface{}
	// The region where the endpoint is located.
	//
	// If the region cannot be determined, `undefined` is returned.
	Region() *string
}

// The jsii proxy for IEndpoint
type jsiiProxy_IEndpoint struct {
	_ byte // padding
}

func (i *jsiiProxy_IEndpoint) RenderEndpointConfiguration() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"renderEndpointConfiguration",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IEndpoint) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

