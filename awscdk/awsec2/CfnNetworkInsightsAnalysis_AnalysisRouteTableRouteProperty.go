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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisroutetableroute.html
//
type CfnNetworkInsightsAnalysis_AnalysisRouteTableRouteProperty struct {
	// The destination IPv4 address, in CIDR notation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisroutetableroute.html#cfn-ec2-networkinsightsanalysis-analysisroutetableroute-destinationcidr
	//
	DestinationCidr *string `field:"optional" json:"destinationCidr" yaml:"destinationCidr"`
	// The prefix of the AWS service .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisroutetableroute.html#cfn-ec2-networkinsightsanalysis-analysisroutetableroute-destinationprefixlistid
	//
	DestinationPrefixListId *string `field:"optional" json:"destinationPrefixListId" yaml:"destinationPrefixListId"`
	// The ID of an egress-only internet gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisroutetableroute.html#cfn-ec2-networkinsightsanalysis-analysisroutetableroute-egressonlyinternetgatewayid
	//
	EgressOnlyInternetGatewayId *string `field:"optional" json:"egressOnlyInternetGatewayId" yaml:"egressOnlyInternetGatewayId"`
	// The ID of the gateway, such as an internet gateway or virtual private gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisroutetableroute.html#cfn-ec2-networkinsightsanalysis-analysisroutetableroute-gatewayid
	//
	GatewayId *string `field:"optional" json:"gatewayId" yaml:"gatewayId"`
	// The ID of the instance, such as a NAT instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisroutetableroute.html#cfn-ec2-networkinsightsanalysis-analysisroutetableroute-instanceid
	//
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// The ID of a NAT gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisroutetableroute.html#cfn-ec2-networkinsightsanalysis-analysisroutetableroute-natgatewayid
	//
	NatGatewayId *string `field:"optional" json:"natGatewayId" yaml:"natGatewayId"`
	// The ID of a network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisroutetableroute.html#cfn-ec2-networkinsightsanalysis-analysisroutetableroute-networkinterfaceid
	//
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// Describes how the route was created. The following are the possible values:.
	//
	// - CreateRouteTable - The route was automatically created when the route table was created.
	// - CreateRoute - The route was manually added to the route table.
	// - EnableVgwRoutePropagation - The route was propagated by route propagation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisroutetableroute.html#cfn-ec2-networkinsightsanalysis-analysisroutetableroute-origin
	//
	Origin *string `field:"optional" json:"origin" yaml:"origin"`
	// The state. The following are the possible values:.
	//
	// - active
	// - blackhole.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisroutetableroute.html#cfn-ec2-networkinsightsanalysis-analysisroutetableroute-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
	// The ID of a transit gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisroutetableroute.html#cfn-ec2-networkinsightsanalysis-analysisroutetableroute-transitgatewayid
	//
	TransitGatewayId *string `field:"optional" json:"transitGatewayId" yaml:"transitGatewayId"`
	// The ID of a VPC peering connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisroutetableroute.html#cfn-ec2-networkinsightsanalysis-analysisroutetableroute-vpcpeeringconnectionid
	//
	VpcPeeringConnectionId *string `field:"optional" json:"vpcPeeringConnectionId" yaml:"vpcPeeringConnectionId"`
}

