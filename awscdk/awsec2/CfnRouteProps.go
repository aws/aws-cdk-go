package awsec2


// Properties for defining a `CfnRoute`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRouteProps := &CfnRouteProps{
//   	RouteTableId: jsii.String("routeTableId"),
//
//   	// the properties below are optional
//   	CarrierGatewayId: jsii.String("carrierGatewayId"),
//   	DestinationCidrBlock: jsii.String("destinationCidrBlock"),
//   	DestinationIpv6CidrBlock: jsii.String("destinationIpv6CidrBlock"),
//   	EgressOnlyInternetGatewayId: jsii.String("egressOnlyInternetGatewayId"),
//   	GatewayId: jsii.String("gatewayId"),
//   	InstanceId: jsii.String("instanceId"),
//   	LocalGatewayId: jsii.String("localGatewayId"),
//   	NatGatewayId: jsii.String("natGatewayId"),
//   	NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   	TransitGatewayId: jsii.String("transitGatewayId"),
//   	VpcEndpointId: jsii.String("vpcEndpointId"),
//   	VpcPeeringConnectionId: jsii.String("vpcPeeringConnectionId"),
//   }
//
type CfnRouteProps struct {
	// The ID of the route table for the route.
	RouteTableId *string `field:"required" json:"routeTableId" yaml:"routeTableId"`
	// The ID of the carrier gateway.
	//
	// You can only use this option when the VPC contains a subnet which is associated with a Wavelength Zone.
	CarrierGatewayId *string `field:"optional" json:"carrierGatewayId" yaml:"carrierGatewayId"`
	// The IPv4 CIDR address block used for the destination match.
	//
	// Routing decisions are based on the most specific match. We modify the specified CIDR block to its canonical form; for example, if you specify `100.68.0.18/18` , we modify it to `100.68.0.0/18` .
	DestinationCidrBlock *string `field:"optional" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// The IPv6 CIDR block used for the destination match.
	//
	// Routing decisions are based on the most specific match.
	DestinationIpv6CidrBlock *string `field:"optional" json:"destinationIpv6CidrBlock" yaml:"destinationIpv6CidrBlock"`
	// [IPv6 traffic only] The ID of an egress-only internet gateway.
	EgressOnlyInternetGatewayId *string `field:"optional" json:"egressOnlyInternetGatewayId" yaml:"egressOnlyInternetGatewayId"`
	// The ID of an internet gateway or virtual private gateway attached to your VPC.
	GatewayId *string `field:"optional" json:"gatewayId" yaml:"gatewayId"`
	// The ID of a NAT instance in your VPC.
	//
	// The operation fails if you specify an instance ID unless exactly one network interface is attached.
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// The ID of the local gateway.
	LocalGatewayId *string `field:"optional" json:"localGatewayId" yaml:"localGatewayId"`
	// [IPv4 traffic only] The ID of a NAT gateway.
	NatGatewayId *string `field:"optional" json:"natGatewayId" yaml:"natGatewayId"`
	// The ID of a network interface.
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

