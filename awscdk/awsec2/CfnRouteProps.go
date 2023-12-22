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
//   	CoreNetworkArn: jsii.String("coreNetworkArn"),
//   	DestinationCidrBlock: jsii.String("destinationCidrBlock"),
//   	DestinationIpv6CidrBlock: jsii.String("destinationIpv6CidrBlock"),
//   	DestinationPrefixListId: jsii.String("destinationPrefixListId"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html
//
type CfnRouteProps struct {
	// The ID of the route table for the route.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-routetableid
	//
	RouteTableId *string `field:"required" json:"routeTableId" yaml:"routeTableId"`
	// The ID of the carrier gateway.
	//
	// You can only use this option when the VPC contains a subnet which is associated with a Wavelength Zone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-carriergatewayid
	//
	CarrierGatewayId *string `field:"optional" json:"carrierGatewayId" yaml:"carrierGatewayId"`
	// The Amazon Resource Name (ARN) of the core network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-corenetworkarn
	//
	CoreNetworkArn *string `field:"optional" json:"coreNetworkArn" yaml:"coreNetworkArn"`
	// The IPv4 CIDR address block used for the destination match.
	//
	// Routing decisions are based on the most specific match. We modify the specified CIDR block to its canonical form; for example, if you specify `100.68.0.18/18` , we modify it to `100.68.0.0/18` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-destinationcidrblock
	//
	DestinationCidrBlock *string `field:"optional" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// The IPv6 CIDR block used for the destination match.
	//
	// Routing decisions are based on the most specific match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-destinationipv6cidrblock
	//
	DestinationIpv6CidrBlock *string `field:"optional" json:"destinationIpv6CidrBlock" yaml:"destinationIpv6CidrBlock"`
	// The ID of a prefix list used for the destination match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-destinationprefixlistid
	//
	DestinationPrefixListId *string `field:"optional" json:"destinationPrefixListId" yaml:"destinationPrefixListId"`
	// [IPv6 traffic only] The ID of an egress-only internet gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-egressonlyinternetgatewayid
	//
	EgressOnlyInternetGatewayId *string `field:"optional" json:"egressOnlyInternetGatewayId" yaml:"egressOnlyInternetGatewayId"`
	// The ID of an internet gateway or virtual private gateway attached to your VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-gatewayid
	//
	GatewayId *string `field:"optional" json:"gatewayId" yaml:"gatewayId"`
	// The ID of a NAT instance in your VPC.
	//
	// The operation fails if you specify an instance ID unless exactly one network interface is attached.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-instanceid
	//
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// The ID of the local gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-localgatewayid
	//
	LocalGatewayId *string `field:"optional" json:"localGatewayId" yaml:"localGatewayId"`
	// [IPv4 traffic only] The ID of a NAT gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-natgatewayid
	//
	NatGatewayId *string `field:"optional" json:"natGatewayId" yaml:"natGatewayId"`
	// The ID of a network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-networkinterfaceid
	//
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The ID of a transit gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-transitgatewayid
	//
	TransitGatewayId *string `field:"optional" json:"transitGatewayId" yaml:"transitGatewayId"`
	// The ID of a VPC endpoint.
	//
	// Supported for Gateway Load Balancer endpoints only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-vpcendpointid
	//
	VpcEndpointId *string `field:"optional" json:"vpcEndpointId" yaml:"vpcEndpointId"`
	// The ID of a VPC peering connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-vpcpeeringconnectionid
	//
	VpcPeeringConnectionId *string `field:"optional" json:"vpcPeeringConnectionId" yaml:"vpcPeeringConnectionId"`
}

