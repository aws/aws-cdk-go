package previewawsnetworkmanagermixins


// Properties for CfnCustomerGatewayAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCustomerGatewayAssociationMixinProps := &CfnCustomerGatewayAssociationMixinProps{
//   	CustomerGatewayArn: jsii.String("customerGatewayArn"),
//   	DeviceId: jsii.String("deviceId"),
//   	GlobalNetworkId: jsii.String("globalNetworkId"),
//   	LinkId: jsii.String("linkId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-customergatewayassociation.html
//
type CfnCustomerGatewayAssociationMixinProps struct {
	// The Amazon Resource Name (ARN) of the customer gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-customergatewayassociation.html#cfn-networkmanager-customergatewayassociation-customergatewayarn
	//
	CustomerGatewayArn *string `field:"optional" json:"customerGatewayArn" yaml:"customerGatewayArn"`
	// The ID of the device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-customergatewayassociation.html#cfn-networkmanager-customergatewayassociation-deviceid
	//
	DeviceId *string `field:"optional" json:"deviceId" yaml:"deviceId"`
	// The ID of the global network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-customergatewayassociation.html#cfn-networkmanager-customergatewayassociation-globalnetworkid
	//
	GlobalNetworkId *string `field:"optional" json:"globalNetworkId" yaml:"globalNetworkId"`
	// The ID of the link.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-customergatewayassociation.html#cfn-networkmanager-customergatewayassociation-linkid
	//
	LinkId *string `field:"optional" json:"linkId" yaml:"linkId"`
}

