package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2/internal"
)

// Experimental.
type IVpc interface {
	awscdk.IResource
	// Adds a new client VPN endpoint to this VPC.
	// Experimental.
	AddClientVpnEndpoint(id *string, options *ClientVpnEndpointOptions) ClientVpnEndpoint
	// Adds a new Flow Log to this VPC.
	// Experimental.
	AddFlowLog(id *string, options *FlowLogOptions) FlowLog
	// Adds a new gateway endpoint to this VPC.
	// Experimental.
	AddGatewayEndpoint(id *string, options *GatewayVpcEndpointOptions) GatewayVpcEndpoint
	// Adds a new interface endpoint to this VPC.
	// Experimental.
	AddInterfaceEndpoint(id *string, options *InterfaceVpcEndpointOptions) InterfaceVpcEndpoint
	// Adds a new VPN connection to this VPC.
	// Experimental.
	AddVpnConnection(id *string, options *VpnConnectionOptions) VpnConnection
	// Adds a VPN Gateway to this VPC.
	// Experimental.
	EnableVpnGateway(options *EnableVpnGatewayOptions)
	// Return information on the subnets appropriate for the given selection strategy.
	//
	// Requires that at least one subnet is matched, throws a descriptive
	// error message otherwise.
	// Experimental.
	SelectSubnets(selection *SubnetSelection) *SelectedSubnets
	// AZs for this VPC.
	// Experimental.
	AvailabilityZones() *[]*string
	// Dependable that can be depended upon to force internet connectivity established on the VPC.
	// Experimental.
	InternetConnectivityEstablished() awscdk.IDependable
	// List of isolated subnets in this VPC.
	// Experimental.
	IsolatedSubnets() *[]ISubnet
	// List of private subnets in this VPC.
	// Experimental.
	PrivateSubnets() *[]ISubnet
	// List of public subnets in this VPC.
	// Experimental.
	PublicSubnets() *[]ISubnet
	// ARN for this VPC.
	// Experimental.
	VpcArn() *string
	// CIDR range for this VPC.
	// Experimental.
	VpcCidrBlock() *string
	// Identifier for this VPC.
	// Experimental.
	VpcId() *string
	// Identifier for the VPN gateway.
	// Experimental.
	VpnGatewayId() *string
}

// The jsii proxy for IVpc
type jsiiProxy_IVpc struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IVpc) AddClientVpnEndpoint(id *string, options *ClientVpnEndpointOptions) ClientVpnEndpoint {
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
	_jsii_.InvokeVoid(
		i,
		"enableVpnGateway",
		[]interface{}{options},
	)
}

func (i *jsiiProxy_IVpc) SelectSubnets(selection *SubnetSelection) *SelectedSubnets {
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

func (j *jsiiProxy_IVpc) InternetConnectivityEstablished() awscdk.IDependable {
	var returns awscdk.IDependable
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

