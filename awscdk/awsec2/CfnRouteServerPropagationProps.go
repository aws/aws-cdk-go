package awsec2


// Properties for defining a `CfnRouteServerPropagation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRouteServerPropagationProps := &CfnRouteServerPropagationProps{
//   	RouteServerId: jsii.String("routeServerId"),
//   	RouteTableId: jsii.String("routeTableId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserverpropagation.html
//
type CfnRouteServerPropagationProps struct {
	// The ID of the route server configured for route propagation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserverpropagation.html#cfn-ec2-routeserverpropagation-routeserverid
	//
	RouteServerId *string `field:"required" json:"routeServerId" yaml:"routeServerId"`
	// The ID of the route table configured for route server propagation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserverpropagation.html#cfn-ec2-routeserverpropagation-routetableid
	//
	RouteTableId *string `field:"required" json:"routeTableId" yaml:"routeTableId"`
}

