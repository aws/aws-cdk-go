package previewawsnetworkmanagermixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDirectConnectGatewayAttachmentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDirectConnectGatewayAttachmentMixinProps := &CfnDirectConnectGatewayAttachmentMixinProps{
//   	CoreNetworkId: jsii.String("coreNetworkId"),
//   	DirectConnectGatewayArn: jsii.String("directConnectGatewayArn"),
//   	EdgeLocations: []*string{
//   		jsii.String("edgeLocations"),
//   	},
//   	ProposedNetworkFunctionGroupChange: &ProposedNetworkFunctionGroupChangeProperty{
//   		AttachmentPolicyRuleNumber: jsii.Number(123),
//   		NetworkFunctionGroupName: jsii.String("networkFunctionGroupName"),
//   		Tags: []CfnTag{
//   			&CfnTag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	ProposedSegmentChange: &ProposedSegmentChangeProperty{
//   		AttachmentPolicyRuleNumber: jsii.Number(123),
//   		SegmentName: jsii.String("segmentName"),
//   		Tags: []CfnTag{
//   			&CfnTag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
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
	// Describes proposed changes to a network function group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-directconnectgatewayattachment.html#cfn-networkmanager-directconnectgatewayattachment-proposednetworkfunctiongroupchange
	//
	ProposedNetworkFunctionGroupChange interface{} `field:"optional" json:"proposedNetworkFunctionGroupChange" yaml:"proposedNetworkFunctionGroupChange"`
	// Describes a proposed segment change.
	//
	// In some cases, the segment change must first be evaluated and accepted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-directconnectgatewayattachment.html#cfn-networkmanager-directconnectgatewayattachment-proposedsegmentchange
	//
	ProposedSegmentChange interface{} `field:"optional" json:"proposedSegmentChange" yaml:"proposedSegmentChange"`
	// Tags for the attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-directconnectgatewayattachment.html#cfn-networkmanager-directconnectgatewayattachment-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

