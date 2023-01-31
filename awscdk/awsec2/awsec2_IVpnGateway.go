package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2/internal"
)

// The virtual private gateway interface.
// Experimental.
type IVpnGateway interface {
	awscdk.IResource
	// The virtual private gateway Id.
	// Experimental.
	GatewayId() *string
}

// The jsii proxy for IVpnGateway
type jsiiProxy_IVpnGateway struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IVpnGateway) GatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayId",
		&returns,
	)
	return returns
}

