package awsec2


// Describes a route table route.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisRouteTableRouteProperty := &analysisRouteTableRouteProperty{
//   	destinationCidr: jsii.String("destinationCidr"),
//   	destinationPrefixListId: jsii.String("destinationPrefixListId"),
//   	egressOnlyInternetGatewayId: jsii.String("egressOnlyInternetGatewayId"),
//   	gatewayId: jsii.String("gatewayId"),
//   	instanceId: jsii.String("instanceId"),
//   	natGatewayId: jsii.String("natGatewayId"),
//   	networkInterfaceId: jsii.String("networkInterfaceId"),
//   	origin: jsii.String("origin"),
//   	state: jsii.String("state"),
//   	transitGatewayId: jsii.String("transitGatewayId"),
//   	vpcPeeringConnectionId: jsii.String("vpcPeeringConnectionId"),
//   }
//
type CfnNetworkInsightsAnalysis_AnalysisRouteTableRouteProperty struct {
	// The destination IPv4 address, in CIDR notation.
	DestinationCidr *string `field:"optional" json:"destinationCidr" yaml:"destinationCidr"`
	// The prefix of the AWS service .
	DestinationPrefixListId *string `field:"optional" json:"destinationPrefixListId" yaml:"destinationPrefixListId"`
	// The ID of an egress-only internet gateway.
	EgressOnlyInternetGatewayId *string `field:"optional" json:"egressOnlyInternetGatewayId" yaml:"egressOnlyInternetGatewayId"`
	// The ID of the gateway, such as an internet gateway or virtual private gateway.
	GatewayId *string `field:"optional" json:"gatewayId" yaml:"gatewayId"`
	// The ID of the instance, such as a NAT instance.
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// The ID of a NAT gateway.
	NatGatewayId *string `field:"optional" json:"natGatewayId" yaml:"natGatewayId"`
	// The ID of a network interface.
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// Describes how the route was created. The following are the possible values:.
	//
	// - CreateRouteTable - The route was automatically created when the route table was created.
	// - CreateRoute - The route was manually added to the route table.
	// - EnableVgwRoutePropagation - The route was propagated by route propagation.
	Origin *string `field:"optional" json:"origin" yaml:"origin"`
	// `CfnNetworkInsightsAnalysis.AnalysisRouteTableRouteProperty.State`.
	State *string `field:"optional" json:"state" yaml:"state"`
	// The ID of a transit gateway.
	TransitGatewayId *string `field:"optional" json:"transitGatewayId" yaml:"transitGatewayId"`
	// The ID of a VPC peering connection.
	VpcPeeringConnectionId *string `field:"optional" json:"vpcPeeringConnectionId" yaml:"vpcPeeringConnectionId"`
}

