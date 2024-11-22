package awsec2alpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awsec2alpha/v2/internal"
)

// Placeholder to see what extra props we might need, will be added to original IVPC.
// Experimental.
type IVpcV2 interface {
	awsec2.IVpc
	// Add an Egress only Internet Gateway to current VPC.
	//
	// Can only be used for ipv6 enabled VPCs.
	// For more information, see the {@link https://docs.aws.amazon.com/vpc/latest/userguide/egress-only-internet-gateway-basics.html}.
	// Experimental.
	AddEgressOnlyInternetGateway(options *EgressOnlyInternetGatewayOptions)
	// Adds an Internet Gateway to current VPC.
	//
	// For more information, see the {@link https://docs.aws.amazon.com/vpc/latest/userguide/vpc-igw-internet-access.html}.
	// Default: - defines route for all ipv4('0.0.0.0') and ipv6 addresses('::/0')
	//
	// Experimental.
	AddInternetGateway(options *InternetGatewayOptions)
	// Adds a new NAT Gateway to VPC A NAT gateway is a Network Address Translation (NAT) service.
	//
	// NAT Gateway Connectivity can be of type `Public` or `Private`.
	// For more information, see the {@link https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html}.
	// Default: ConnectivityType.Public
	//
	// Experimental.
	AddNatGateway(options *NatGatewayOptions) NatGateway
	// Adds a new role to acceptor VPC account A cross account role is required for the VPC to peer with another account.
	//
	// For more information, see the {@link https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/peer-with-vpc-in-another-account.html}.
	// Experimental.
	CreateAcceptorVpcRole(requestorAccountId *string) awsiam.Role
	// Creates a new peering connection A peering connection is a private virtual network established between two VPCs.
	//
	// For more information, see the {@link https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html}.
	// Experimental.
	CreatePeeringConnection(id *string, options *VPCPeeringConnectionOptions) VPCPeeringConnection
	// Adds VPN Gateway to VPC and set route propogation.
	//
	// For more information, see the {@link https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpngateway.html}.
	// Default: - no route propogation.
	//
	// Experimental.
	EnableVpnGatewayV2(options *VPNGatewayV2Options) VPNGatewayV2
	// The primary IPv4 CIDR block associated with the VPC.
	//
	// Needed in order to validate the vpc range of subnet
	// current prop vpcCidrBlock refers to the token value
	// For more information, see the {@link https://docs.aws.amazon.com/vpc/latest/userguide/vpc-cidr-blocks.html#vpc-sizing-ipv4}.
	// Experimental.
	Ipv4CidrBlock() *string
	// IPv4 CIDR provisioned under pool Required to check for overlapping CIDRs after provisioning is complete under IPAM pool.
	// Experimental.
	Ipv4IpamProvisionedCidrs() *[]*string
	// The ID of the AWS account that owns the VPC.
	// Default: - the account id of the parent stack.
	//
	// Experimental.
	OwnerAccountId() *string
	// Optional to override inferred region.
	// Default: - current stack's environment region.
	//
	// Experimental.
	Region() *string
	// The secondary CIDR blocks associated with the VPC.
	//
	// For more information, see the {@link https://docs.aws.amazon.com/vpc/latest/userguide/vpc-cidr-blocks.html#vpc-resize}.
	// Experimental.
	SecondaryCidrBlock() *[]IVPCCidrBlock
}

// The jsii proxy for IVpcV2
type jsiiProxy_IVpcV2 struct {
	internal.Type__awsec2IVpc
}

func (i *jsiiProxy_IVpcV2) AddEgressOnlyInternetGateway(options *EgressOnlyInternetGatewayOptions) {
	if err := i.validateAddEgressOnlyInternetGatewayParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addEgressOnlyInternetGateway",
		[]interface{}{options},
	)
}

func (i *jsiiProxy_IVpcV2) AddInternetGateway(options *InternetGatewayOptions) {
	if err := i.validateAddInternetGatewayParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addInternetGateway",
		[]interface{}{options},
	)
}

func (i *jsiiProxy_IVpcV2) AddNatGateway(options *NatGatewayOptions) NatGateway {
	if err := i.validateAddNatGatewayParameters(options); err != nil {
		panic(err)
	}
	var returns NatGateway

	_jsii_.Invoke(
		i,
		"addNatGateway",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IVpcV2) CreateAcceptorVpcRole(requestorAccountId *string) awsiam.Role {
	if err := i.validateCreateAcceptorVpcRoleParameters(requestorAccountId); err != nil {
		panic(err)
	}
	var returns awsiam.Role

	_jsii_.Invoke(
		i,
		"createAcceptorVpcRole",
		[]interface{}{requestorAccountId},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IVpcV2) CreatePeeringConnection(id *string, options *VPCPeeringConnectionOptions) VPCPeeringConnection {
	if err := i.validateCreatePeeringConnectionParameters(id, options); err != nil {
		panic(err)
	}
	var returns VPCPeeringConnection

	_jsii_.Invoke(
		i,
		"createPeeringConnection",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IVpcV2) EnableVpnGatewayV2(options *VPNGatewayV2Options) VPNGatewayV2 {
	if err := i.validateEnableVpnGatewayV2Parameters(options); err != nil {
		panic(err)
	}
	var returns VPNGatewayV2

	_jsii_.Invoke(
		i,
		"enableVpnGatewayV2",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IVpcV2) Ipv4CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcV2) Ipv4IpamProvisionedCidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv4IpamProvisionedCidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcV2) OwnerAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcV2) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcV2) SecondaryCidrBlock() *[]IVPCCidrBlock {
	var returns *[]IVPCCidrBlock
	_jsii_.Get(
		j,
		"secondaryCidrBlock",
		&returns,
	)
	return returns
}

