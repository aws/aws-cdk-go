package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

type IVpc interface {
	awscdk.IResource
	// Adds a new client VPN endpoint to this VPC.
	AddClientVpnEndpoint(id *string, options *ClientVpnEndpointOptions) ClientVpnEndpoint
	// Adds a new Flow Log to this VPC.
	AddFlowLog(id *string, options *FlowLogOptions) FlowLog
	// Adds a new gateway endpoint to this VPC.
	AddGatewayEndpoint(id *string, options *GatewayVpcEndpointOptions) GatewayVpcEndpoint
	// Adds a new interface endpoint to this VPC.
	AddInterfaceEndpoint(id *string, options *InterfaceVpcEndpointOptions) InterfaceVpcEndpoint
	// Adds a new VPN connection to this VPC.
	AddVpnConnection(id *string, options *VpnConnectionOptions) VpnConnection
	// Adds a VPN Gateway to this VPC.
	EnableVpnGateway(options *EnableVpnGatewayOptions)
	// Return information on the subnets appropriate for the given selection strategy.
	//
	// Requires that at least one subnet is matched, throws a descriptive
	// error message otherwise.
	SelectSubnets(selection *SubnetSelection) *SelectedSubnets
	// AZs for this VPC.
	AvailabilityZones() *[]*string
	// Dependable that can be depended upon to force internet connectivity established on the VPC.
	InternetConnectivityEstablished() constructs.IDependable
	// List of isolated subnets in this VPC.
	IsolatedSubnets() *[]ISubnet
	// List of private subnets in this VPC.
	PrivateSubnets() *[]ISubnet
	// List of public subnets in this VPC.
	PublicSubnets() *[]ISubnet
	// ARN for this VPC.
	VpcArn() *string
	// CIDR range for this VPC.
	VpcCidrBlock() *string
	// Identifier for this VPC.
	VpcId() *string
	// Identifier for the VPN gateway.
	VpnGatewayId() *string
}

// The jsii proxy for IVpc
type jsiiProxy_IVpc struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IVpc) AddClientVpnEndpoint(id *string, options *ClientVpnEndpointOptions) ClientVpnEndpoint {
	if err := i.validateAddClientVpnEndpointParameters(id, options); err != nil {
		panic(err)
	}
	var returns ClientVpnEndpoint

	_jsii_.Invoke(
		i,
		"addClientVpnEndpoint",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IVpc) AddFlowLog(id *string, options *FlowLogOptions) FlowLog {
	if err := i.validateAddFlowLogParameters(id, options); err != nil {
		panic(err)
	}
	var returns FlowLog

	_jsii_.Invoke(
		i,
		"addFlowLog",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IVpc) AddGatewayEndpoint(id *string, options *GatewayVpcEndpointOptions) GatewayVpcEndpoint {
	if err := i.validateAddGatewayEndpointParameters(id, options); err != nil {
		panic(err)
	}
	var returns GatewayVpcEndpoint

	_jsii_.Invoke(
		i,
		"addGatewayEndpoint",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IVpc) AddInterfaceEndpoint(id *string, options *InterfaceVpcEndpointOptions) InterfaceVpcEndpoint {
	if err := i.validateAddInterfaceEndpointParameters(id, options); err != nil {
		panic(err)
	}
	var returns InterfaceVpcEndpoint

	_jsii_.Invoke(
		i,
		"addInterfaceEndpoint",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IVpc) AddVpnConnection(id *string, options *VpnConnectionOptions) VpnConnection {
	if err := i.validateAddVpnConnectionParameters(id, options); err != nil {
		panic(err)
	}
	var returns VpnConnection

	_jsii_.Invoke(
		i,
		"addVpnConnection",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IVpc) EnableVpnGateway(options *EnableVpnGatewayOptions) {
	if err := i.validateEnableVpnGatewayParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"enableVpnGateway",
		[]interface{}{options},
	)
}

func (i *jsiiProxy_IVpc) SelectSubnets(selection *SubnetSelection) *SelectedSubnets {
	if err := i.validateSelectSubnetsParameters(selection); err != nil {
		panic(err)
	}
	var returns *SelectedSubnets

	_jsii_.Invoke(
		i,
		"selectSubnets",
		[]interface{}{selection},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IVpc) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpc) InternetConnectivityEstablished() constructs.IDependable {
	var returns constructs.IDependable
	_jsii_.Get(
		j,
		"internetConnectivityEstablished",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpc) IsolatedSubnets() *[]ISubnet {
	var returns *[]ISubnet
	_jsii_.Get(
		j,
		"isolatedSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpc) PrivateSubnets() *[]ISubnet {
	var returns *[]ISubnet
	_jsii_.Get(
		j,
		"privateSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpc) PublicSubnets() *[]ISubnet {
	var returns *[]ISubnet
	_jsii_.Get(
		j,
		"publicSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpc) VpcArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpc) VpcCidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcCidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpc) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpc) VpnGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnGatewayId",
		&returns,
	)
	return returns
}

