package awsec2


// Properties for defining a `CfnVPNGatewayRoutePropagation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVPNGatewayRoutePropagationProps := &CfnVPNGatewayRoutePropagationProps{
//   	RouteTableIds: []*string{
//   		jsii.String("routeTableIds"),
//   	},
//   	VpnGatewayId: jsii.String("vpnGatewayId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpngatewayroutepropagation.html
//
type CfnVPNGatewayRoutePropagationProps struct {
	// The ID of the route table.
	//
	// The routing table must be associated with the same VPC that the virtual private gateway is attached to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpngatewayroutepropagation.html#cfn-ec2-vpngatewayroutepropagation-routetableids
	//
	RouteTableIds *[]*string `field:"required" json:"routeTableIds" yaml:"routeTableIds"`
	// The ID of the virtual private gateway that is attached to a VPC.
	//
	// The virtual private gateway must be attached to the same VPC that the routing tables are associated with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpngatewayroutepropagation.html#cfn-ec2-vpngatewayroutepropagation-vpngatewayid
	//
	VpnGatewayId *string `field:"required" json:"vpnGatewayId" yaml:"vpnGatewayId"`
}

