package previewawsec2mixins


// Properties for CfnRouteServerAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRouteServerAssociationMixinProps := &CfnRouteServerAssociationMixinProps{
//   	RouteServerId: jsii.String("routeServerId"),
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserverassociation.html
//
type CfnRouteServerAssociationMixinProps struct {
	// The ID of the associated route server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserverassociation.html#cfn-ec2-routeserverassociation-routeserverid
	//
	RouteServerId *string `field:"optional" json:"routeServerId" yaml:"routeServerId"`
	// The ID of the associated VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserverassociation.html#cfn-ec2-routeserverassociation-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

