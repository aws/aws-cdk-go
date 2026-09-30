package awsnetworkmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDirectConnectGatewayAttachmentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnDirectConnectGatewayAttachmentMixinProps := &CfnDirectConnectGatewayAttachmentMixinProps{
//   	CoreNetworkId: jsii.String("coreNetworkId"),
//   	DirectConnectGatewayArn: jsii.String("directConnectGatewayArn"),
//   	EdgeLocations: []*string{
//   		jsii.String("edgeLocations"),
//   	},
//   	RoutingPolicyLabel: jsii.String("routingPolicyLabel"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-directconnectgatewayattachment.html
//
type CfnDirectConnectGatewayAttachmentMixinProps struct {
	// The ID of a core network for the Direct Connect Gateway attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-directconnectgatewayattachment.html#cfn-networkmanager-directconnectgatewayattachment-corenetworkid
	//
	CoreNetworkId *string `field:"optional" json:"coreNetworkId" yaml:"coreNetworkId"`
	// The Direct Connect gateway attachment ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-directconnectgatewayattachment.html#cfn-networkmanager-directconnectgatewayattachment-directconnectgatewayarn
	//
	DirectConnectGatewayArn *string `field:"optional" json:"directConnectGatewayArn" yaml:"directConnectGatewayArn"`
	// The Regions where the edges are located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-directconnectgatewayattachment.html#cfn-networkmanager-directconnectgatewayattachment-edgelocations
	//
	EdgeLocations *[]*string `field:"optional" json:"edgeLocations" yaml:"edgeLocations"`
	// Routing policy label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-directconnectgatewayattachment.html#cfn-networkmanager-directconnectgatewayattachment-routingpolicylabel
	//
	RoutingPolicyLabel *string `field:"optional" json:"routingPolicyLabel" yaml:"routingPolicyLabel"`
	// Tags for the attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-directconnectgatewayattachment.html#cfn-networkmanager-directconnectgatewayattachment-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

