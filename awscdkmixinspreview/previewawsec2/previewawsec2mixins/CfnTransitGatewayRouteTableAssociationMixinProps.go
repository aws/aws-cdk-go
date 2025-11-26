package previewawsec2mixins


// Properties for CfnTransitGatewayRouteTableAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTransitGatewayRouteTableAssociationMixinProps := &CfnTransitGatewayRouteTableAssociationMixinProps{
//   	TransitGatewayAttachmentId: jsii.String("transitGatewayAttachmentId"),
//   	TransitGatewayRouteTableId: jsii.String("transitGatewayRouteTableId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetableassociation.html
//
type CfnTransitGatewayRouteTableAssociationMixinProps struct {
	// The ID of the attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetableassociation.html#cfn-ec2-transitgatewayroutetableassociation-transitgatewayattachmentid
	//
	TransitGatewayAttachmentId *string `field:"optional" json:"transitGatewayAttachmentId" yaml:"transitGatewayAttachmentId"`
	// The ID of the route table for the transit gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetableassociation.html#cfn-ec2-transitgatewayroutetableassociation-transitgatewayroutetableid
	//
	TransitGatewayRouteTableId *string `field:"optional" json:"transitGatewayRouteTableId" yaml:"transitGatewayRouteTableId"`
}

