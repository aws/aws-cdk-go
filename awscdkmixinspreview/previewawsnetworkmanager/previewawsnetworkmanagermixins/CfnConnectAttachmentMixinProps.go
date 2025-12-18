package previewawsnetworkmanagermixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnConnectAttachmentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConnectAttachmentMixinProps := &CfnConnectAttachmentMixinProps{
//   	CoreNetworkId: jsii.String("coreNetworkId"),
//   	EdgeLocation: jsii.String("edgeLocation"),
//   	NetworkFunctionGroupName: jsii.String("networkFunctionGroupName"),
//   	Options: &ConnectAttachmentOptionsProperty{
//   		Protocol: jsii.String("protocol"),
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
//   	RoutingPolicyLabel: jsii.String("routingPolicyLabel"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TransportAttachmentId: jsii.String("transportAttachmentId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connectattachment.html
//
type CfnConnectAttachmentMixinProps struct {
	// The ID of the core network where the Connect attachment is located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connectattachment.html#cfn-networkmanager-connectattachment-corenetworkid
	//
	CoreNetworkId *string `field:"optional" json:"coreNetworkId" yaml:"coreNetworkId"`
	// The Region where the edge is located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connectattachment.html#cfn-networkmanager-connectattachment-edgelocation
	//
	EdgeLocation *string `field:"optional" json:"edgeLocation" yaml:"edgeLocation"`
	// The name of the network function group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connectattachment.html#cfn-networkmanager-connectattachment-networkfunctiongroupname
	//
	NetworkFunctionGroupName *string `field:"optional" json:"networkFunctionGroupName" yaml:"networkFunctionGroupName"`
	// Options for connecting an attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connectattachment.html#cfn-networkmanager-connectattachment-options
	//
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// Describes proposed changes to a network function group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connectattachment.html#cfn-networkmanager-connectattachment-proposednetworkfunctiongroupchange
	//
	ProposedNetworkFunctionGroupChange interface{} `field:"optional" json:"proposedNetworkFunctionGroupChange" yaml:"proposedNetworkFunctionGroupChange"`
	// Describes a proposed segment change.
	//
	// In some cases, the segment change must first be evaluated and accepted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connectattachment.html#cfn-networkmanager-connectattachment-proposedsegmentchange
	//
	ProposedSegmentChange interface{} `field:"optional" json:"proposedSegmentChange" yaml:"proposedSegmentChange"`
	// Routing policy label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connectattachment.html#cfn-networkmanager-connectattachment-routingpolicylabel
	//
	RoutingPolicyLabel *string `field:"optional" json:"routingPolicyLabel" yaml:"routingPolicyLabel"`
	// The tags associated with the Connect attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connectattachment.html#cfn-networkmanager-connectattachment-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the transport attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connectattachment.html#cfn-networkmanager-connectattachment-transportattachmentid
	//
	TransportAttachmentId *string `field:"optional" json:"transportAttachmentId" yaml:"transportAttachmentId"`
}

