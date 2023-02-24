package awsec2


// Describes a route table route.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisRouteTableRouteProperty := &AnalysisRouteTableRouteProperty{
//   	DestinationCidr: jsii.String("destinationCidr"),
//   	DestinationPrefixListId: jsii.String("destinationPrefixListId"),
//   	EgressOnlyInternetGatewayId: jsii.String("egressOnlyInternetGatewayId"),
//   	GatewayId: jsii.String("gatewayId"),
//   	InstanceId: jsii.String("instanceId"),
//   	NatGatewayId: jsii.String("natGatewayId"),
//   	NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   	Origin: jsii.String("origin"),
//   	State: jsii.String("state"),
//   	TransitGatewayId: jsii.String("transitGatewayId"),
//   	VpcPeeringConnectionId: jsii.String("vpcPeeringConnectionId"),
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
	// The state. The following are the possible values:.
	//
	// - active
	// - blackhole.
	State *string `field:"optional" json:"state" yaml:"state"`
	// The ID of a transit gateway.
	TransitGatewayId *string `field:"optional" json:"transitGatewayId" yaml:"transitGatewayId"`
	// The ID of a VPC peering connection.
	VpcPeeringConnectionId *string `field:"optional" json:"vpcPeeringConnectionId" yaml:"vpcPeeringConnectionId"`
}

