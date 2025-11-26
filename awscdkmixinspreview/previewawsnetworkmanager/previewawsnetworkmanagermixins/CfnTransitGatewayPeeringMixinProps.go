package previewawsnetworkmanagermixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnTransitGatewayPeeringPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTransitGatewayPeeringMixinProps := &CfnTransitGatewayPeeringMixinProps{
//   	CoreNetworkId: jsii.String("coreNetworkId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TransitGatewayArn: jsii.String("transitGatewayArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-transitgatewaypeering.html
//
type CfnTransitGatewayPeeringMixinProps struct {
	// The ID of the core network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-transitgatewaypeering.html#cfn-networkmanager-transitgatewaypeering-corenetworkid
	//
	CoreNetworkId *string `field:"optional" json:"coreNetworkId" yaml:"coreNetworkId"`
	// The list of key-value tags associated with the peering.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-transitgatewaypeering.html#cfn-networkmanager-transitgatewaypeering-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ARN of the transit gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-transitgatewaypeering.html#cfn-networkmanager-transitgatewaypeering-transitgatewayarn
	//
	TransitGatewayArn *string `field:"optional" json:"transitGatewayArn" yaml:"transitGatewayArn"`
}

