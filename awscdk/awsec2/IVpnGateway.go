package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

// The virtual private gateway interface.
type IVpnGateway interface {
	awscdk.IResource
	interfacesawsec2.IVPNGatewayRef
	// The virtual private gateway Id.
	GatewayId() *string
}

// The jsii proxy for IVpnGateway
type jsiiProxy_IVpnGateway struct {
	internal.Type__awscdkIResource
	internal.Type__interfacesawsec2IVPNGatewayRef
}

func (i *jsiiProxy_IVpnGateway) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
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

func (j *jsiiProxy_IVpnGateway) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpnGateway) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpnGateway) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpnGateway) VpnGatewayRef() *interfacesawsec2.VPNGatewayReference {
	var returns *interfacesawsec2.VPNGatewayReference
	_jsii_.Get(
		j,
		"vpnGatewayRef",
		&returns,
	)
	return returns
}

