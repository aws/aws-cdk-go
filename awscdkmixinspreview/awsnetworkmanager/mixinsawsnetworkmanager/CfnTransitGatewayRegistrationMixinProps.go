package mixinsawsnetworkmanager


// Properties for CfnTransitGatewayRegistrationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTransitGatewayRegistrationMixinProps := &CfnTransitGatewayRegistrationMixinProps{
//   	GlobalNetworkId: jsii.String("globalNetworkId"),
//   	TransitGatewayArn: jsii.String("transitGatewayArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-transitgatewayregistration.html
//
type CfnTransitGatewayRegistrationMixinProps struct {
	// The ID of the global network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-transitgatewayregistration.html#cfn-networkmanager-transitgatewayregistration-globalnetworkid
	//
	GlobalNetworkId *string `field:"optional" json:"globalNetworkId" yaml:"globalNetworkId"`
	// The Amazon Resource Name (ARN) of the transit gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-transitgatewayregistration.html#cfn-networkmanager-transitgatewayregistration-transitgatewayarn
	//
	TransitGatewayArn *string `field:"optional" json:"transitGatewayArn" yaml:"transitGatewayArn"`
}

