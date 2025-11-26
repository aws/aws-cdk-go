package previewawsec2mixins


// Properties for CfnTransitGatewayMulticastDomainAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTransitGatewayMulticastDomainAssociationMixinProps := &CfnTransitGatewayMulticastDomainAssociationMixinProps{
//   	SubnetId: jsii.String("subnetId"),
//   	TransitGatewayAttachmentId: jsii.String("transitGatewayAttachmentId"),
//   	TransitGatewayMulticastDomainId: jsii.String("transitGatewayMulticastDomainId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymulticastdomainassociation.html
//
type CfnTransitGatewayMulticastDomainAssociationMixinProps struct {
	// The IDs of the subnets to associate with the transit gateway multicast domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymulticastdomainassociation.html#cfn-ec2-transitgatewaymulticastdomainassociation-subnetid
	//
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// The ID of the transit gateway attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymulticastdomainassociation.html#cfn-ec2-transitgatewaymulticastdomainassociation-transitgatewayattachmentid
	//
	TransitGatewayAttachmentId *string `field:"optional" json:"transitGatewayAttachmentId" yaml:"transitGatewayAttachmentId"`
	// The ID of the transit gateway multicast domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymulticastdomainassociation.html#cfn-ec2-transitgatewaymulticastdomainassociation-transitgatewaymulticastdomainid
	//
	TransitGatewayMulticastDomainId *string `field:"optional" json:"transitGatewayMulticastDomainId" yaml:"transitGatewayMulticastDomainId"`
}

