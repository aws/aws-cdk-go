package mixinsawsec2


// Properties for CfnRouteServerPropagationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRouteServerPropagationMixinProps := &CfnRouteServerPropagationMixinProps{
//   	RouteServerId: jsii.String("routeServerId"),
//   	RouteTableId: jsii.String("routeTableId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserverpropagation.html
//
type CfnRouteServerPropagationMixinProps struct {
	// The ID of the route server configured for route propagation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserverpropagation.html#cfn-ec2-routeserverpropagation-routeserverid
	//
	RouteServerId *string `field:"optional" json:"routeServerId" yaml:"routeServerId"`
	// The ID of the route table configured for route server propagation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserverpropagation.html#cfn-ec2-routeserverpropagation-routetableid
	//
	RouteTableId *string `field:"optional" json:"routeTableId" yaml:"routeTableId"`
}

