package awsec2


// Properties for defining a `CfnLocalGatewayRoute`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLocalGatewayRouteProps := &CfnLocalGatewayRouteProps{
//   	DestinationCidrBlock: jsii.String("destinationCidrBlock"),
//   	LocalGatewayRouteTableId: jsii.String("localGatewayRouteTableId"),
//
//   	// the properties below are optional
//   	LocalGatewayVirtualInterfaceGroupId: jsii.String("localGatewayVirtualInterfaceGroupId"),
//   	NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   }
//
type CfnLocalGatewayRouteProps struct {
	// The CIDR block used for destination matches.
	DestinationCidrBlock *string `field:"required" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// The ID of the local gateway route table.
	LocalGatewayRouteTableId *string `field:"required" json:"localGatewayRouteTableId" yaml:"localGatewayRouteTableId"`
	// The ID of the virtual interface group.
	LocalGatewayVirtualInterfaceGroupId *string `field:"optional" json:"localGatewayVirtualInterfaceGroupId" yaml:"localGatewayVirtualInterfaceGroupId"`
	// The ID of the network interface.
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
}

