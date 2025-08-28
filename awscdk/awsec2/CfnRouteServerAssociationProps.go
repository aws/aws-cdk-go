package awsec2


// Properties for defining a `CfnRouteServerAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRouteServerAssociationProps := &CfnRouteServerAssociationProps{
//   	RouteServerId: jsii.String("routeServerId"),
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserverassociation.html
//
type CfnRouteServerAssociationProps struct {
	// The ID of the associated route server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserverassociation.html#cfn-ec2-routeserverassociation-routeserverid
	//
	RouteServerId *string `field:"required" json:"routeServerId" yaml:"routeServerId"`
	// The ID of the associated VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserverassociation.html#cfn-ec2-routeserverassociation-vpcid
	//
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

