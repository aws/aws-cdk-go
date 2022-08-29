package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A service for a gateway VPC endpoint.
type IGatewayVpcEndpointService interface {
	// The name of the service.
	Name() *string
}

// The jsii proxy for IGatewayVpcEndpointService
type jsiiProxy_IGatewayVpcEndpointService struct {
	_ byte // padding
}

func (j *jsiiProxy_IGatewayVpcEndpointService) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

