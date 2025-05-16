package awsec2


// Properties for defining a `CfnGatewayRouteTableAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGatewayRouteTableAssociationProps := &CfnGatewayRouteTableAssociationProps{
//   	GatewayId: jsii.String("gatewayId"),
//   	RouteTableId: jsii.String("routeTableId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-gatewayroutetableassociation.html
//
type CfnGatewayRouteTableAssociationProps struct {
	// The ID of the gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-gatewayroutetableassociation.html#cfn-ec2-gatewayroutetableassociation-gatewayid
	//
	GatewayId *string `field:"required" json:"gatewayId" yaml:"gatewayId"`
	// The ID of the route table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-gatewayroutetableassociation.html#cfn-ec2-gatewayroutetableassociation-routetableid
	//
	RouteTableId *string `field:"required" json:"routeTableId" yaml:"routeTableId"`
}

