package previewawsec2mixins


// Properties for CfnClientVpnTargetNetworkAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClientVpnTargetNetworkAssociationMixinProps := &CfnClientVpnTargetNetworkAssociationMixinProps{
//   	ClientVpnEndpointId: jsii.String("clientVpnEndpointId"),
//   	SubnetId: jsii.String("subnetId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpntargetnetworkassociation.html
//
type CfnClientVpnTargetNetworkAssociationMixinProps struct {
	// The ID of the Client VPN endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpntargetnetworkassociation.html#cfn-ec2-clientvpntargetnetworkassociation-clientvpnendpointid
	//
	ClientVpnEndpointId *string `field:"optional" json:"clientVpnEndpointId" yaml:"clientVpnEndpointId"`
	// The ID of the subnet to associate with the Client VPN endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpntargetnetworkassociation.html#cfn-ec2-clientvpntargetnetworkassociation-subnetid
	//
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

