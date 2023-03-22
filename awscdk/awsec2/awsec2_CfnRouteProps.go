package awsec2


// Properties for defining a `CfnRoute`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRouteProps := &cfnRouteProps{
//   	routeTableId: jsii.String("routeTableId"),
//
//   	// the properties below are optional
//   	carrierGatewayId: jsii.String("carrierGatewayId"),
//   	destinationCidrBlock: jsii.String("destinationCidrBlock"),
//   	destinationIpv6CidrBlock: jsii.String("destinationIpv6CidrBlock"),
//   	egressOnlyInternetGatewayId: jsii.String("egressOnlyInternetGatewayId"),
//   	gatewayId: jsii.String("gatewayId"),
//   	instanceId: jsii.String("instanceId"),
//   	localGatewayId: jsii.String("localGatewayId"),
//   	natGatewayId: jsii.String("natGatewayId"),
//   	networkInterfaceId: jsii.String("networkInterfaceId"),
//   	transitGatewayId: jsii.String("transitGatewayId"),
//   	vpcEndpointId: jsii.String("vpcEndpointId"),
//   	vpcPeeringConnectionId: jsii.String("vpcPeeringConnectionId"),
//   }
//
type CfnRouteProps struct {
	// The ID of the route table.
	//
	// The routing table must be associated with the same VPC that the virtual private gateway is attached to.
	RouteTableId *string `field:"required" json:"routeTableId" yaml:"routeTableId"`
	// The ID of the carrier gateway.
	CarrierGatewayId *string `field:"optional" json:"carrierGatewayId" yaml:"carrierGatewayId"`
	// The IPv4 CIDR block used for the destination match.
	DestinationCidrBlock *string `field:"optional" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// The IPv6 CIDR block used for the destination match.
	DestinationIpv6CidrBlock *string `field:"optional" json:"destinationIpv6CidrBlock" yaml:"destinationIpv6CidrBlock"`
	// The ID of the egress-only internet gateway.
	EgressOnlyInternetGatewayId *string `field:"optional" json:"egressOnlyInternetGatewayId" yaml:"egressOnlyInternetGatewayId"`
	// The ID of an internet gateway or virtual private gateway attached to your VPC.
	GatewayId *string `field:"optional" json:"gatewayId" yaml:"gatewayId"`
	// The ID of a NAT instance in your VPC.
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// The ID of the local gateway.
	LocalGatewayId *string `field:"optional" json:"localGatewayId" yaml:"localGatewayId"`
	// The ID of a NAT gateway.
	NatGatewayId *string `field:"optional" json:"natGatewayId" yaml:"natGatewayId"`
	// The ID of the network interface.
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The ID of a transit gateway.
	TransitGatewayId *string `field:"optional" json:"transitGatewayId" yaml:"transitGatewayId"`
	// The ID of a VPC endpoint.
	//
	// Supported for Gateway Load Balancer endpoints only.
	VpcEndpointId *string `field:"optional" json:"vpcEndpointId" yaml:"vpcEndpointId"`
	// The ID of a VPC peering connection.
	VpcPeeringConnectionId *string `field:"optional" json:"vpcPeeringConnectionId" yaml:"vpcPeeringConnectionId"`
}

