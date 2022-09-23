package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
)

// A VPC endpoint.
type IVpcEndpoint interface {
	awscdk.IResource
	// The VPC endpoint identifier.
	VpcEndpointId() *string
}

// The jsii proxy for IVpcEndpoint
type jsiiProxy_IVpcEndpoint struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IVpcEndpoint) VpcEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointId",
		&returns,
	)
	return returns
}

