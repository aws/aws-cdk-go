package awsec2


// Properties for defining a `CfnLocalGatewayRoute`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLocalGatewayRouteProps := &cfnLocalGatewayRouteProps{
//   	destinationCidrBlock: jsii.String("destinationCidrBlock"),
//   	localGatewayRouteTableId: jsii.String("localGatewayRouteTableId"),
//
//   	// the properties below are optional
//   	localGatewayVirtualInterfaceGroupId: jsii.String("localGatewayVirtualInterfaceGroupId"),
//   	networkInterfaceId: jsii.String("networkInterfaceId"),
//   }
//
type CfnLocalGatewayRouteProps struct {
	// The CIDR block used for destination matches.
	DestinationCidrBlock *string `field:"required" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// The ID of the local gateway route table.
	LocalGatewayRouteTableId *string `field:"required" json:"localGatewayRouteTableId" yaml:"localGatewayRouteTableId"`
	// The ID of the virtual interface group.
	LocalGatewayVirtualInterfaceGroupId *string `field:"optional" json:"localGatewayVirtualInterfaceGroupId" yaml:"localGatewayVirtualInterfaceGroupId"`
	// `AWS::EC2::LocalGatewayRoute.NetworkInterfaceId`.
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
}

