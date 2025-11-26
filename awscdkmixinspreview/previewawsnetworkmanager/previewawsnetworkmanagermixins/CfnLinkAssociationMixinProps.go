package previewawsnetworkmanagermixins


// Properties for CfnLinkAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLinkAssociationMixinProps := &CfnLinkAssociationMixinProps{
//   	DeviceId: jsii.String("deviceId"),
//   	GlobalNetworkId: jsii.String("globalNetworkId"),
//   	LinkId: jsii.String("linkId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-linkassociation.html
//
type CfnLinkAssociationMixinProps struct {
	// The device ID for the link association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-linkassociation.html#cfn-networkmanager-linkassociation-deviceid
	//
	DeviceId *string `field:"optional" json:"deviceId" yaml:"deviceId"`
	// The ID of the global network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-linkassociation.html#cfn-networkmanager-linkassociation-globalnetworkid
	//
	GlobalNetworkId *string `field:"optional" json:"globalNetworkId" yaml:"globalNetworkId"`
	// The ID of the link.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-linkassociation.html#cfn-networkmanager-linkassociation-linkid
	//
	LinkId *string `field:"optional" json:"linkId" yaml:"linkId"`
}

