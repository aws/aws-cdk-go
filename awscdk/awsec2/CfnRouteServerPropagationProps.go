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
	// Route Server ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserverpropagation.html#cfn-ec2-routeserverpropagation-routeserverid
	//
	RouteServerId *string `field:"required" json:"routeServerId" yaml:"routeServerId"`
	// Route Table ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserverpropagation.html#cfn-ec2-routeserverpropagation-routetableid
	//
	RouteTableId *string `field:"required" json:"routeTableId" yaml:"routeTableId"`
}

