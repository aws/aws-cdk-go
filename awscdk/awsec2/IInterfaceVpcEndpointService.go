package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A service for an interface VPC endpoint.
type IInterfaceVpcEndpointService interface {
	// The name of the service.
	Name() *string
	// The port of the service.
	Port() *float64
	// Whether Private DNS is supported by default.
	PrivateDnsDefault() *bool
}

// The jsii proxy for IInterfaceVpcEndpointService
type jsiiProxy_IInterfaceVpcEndpointService struct {
	_ byte // padding
}

func (j *jsiiProxy_IInterfaceVpcEndpointService) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInterfaceVpcEndpointService) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInterfaceVpcEndpointService) PrivateDnsDefault() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"privateDnsDefault",
		&returns,
	)
	return returns
}

