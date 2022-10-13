package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
)

// A VPC endpoint service.
type IVpcEndpointService interface {
	awscdk.IResource
	// The id of the VPC Endpoint Service that clients use to connect to, like vpce-svc-xxxxxxxxxxxxxxxx.
	VpcEndpointServiceId() *string
	// The service name of the VPC Endpoint Service that clients use to connect to, like com.amazonaws.vpce.<region>.vpce-svc-xxxxxxxxxxxxxxxx.
	VpcEndpointServiceName() *string
}

// The jsii proxy for IVpcEndpointService
type jsiiProxy_IVpcEndpointService struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IVpcEndpointService) VpcEndpointServiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointServiceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcEndpointService) VpcEndpointServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointServiceName",
		&returns,
	)
	return returns
}

