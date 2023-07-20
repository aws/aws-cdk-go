package awsnetworkmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnConnectAttachment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConnectAttachmentProps := &CfnConnectAttachmentProps{
//   	CoreNetworkId: jsii.String("coreNetworkId"),
//   	EdgeLocation: jsii.String("edgeLocation"),
//   	Options: &ConnectAttachmentOptionsProperty{
//   		Protocol: jsii.String("protocol"),
//   	},
//   	TransportAttachmentId: jsii.String("transportAttachmentId"),
//
//   	// the properties below are optional
//   	ProposedSegmentChange: &ProposedSegmentChangeProperty{
//   		AttachmentPolicyRuleNumber: jsii.Number(123),
//   		SegmentName: jsii.String("segmentName"),
//   		Tags: []cfnTag{
//   			&cfnTag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	Tags: []*cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connectattachment.html
//
type CfnConnectAttachmentProps struct {
	// The ID of the core network where the Connect attachment is located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connectattachment.html#cfn-networkmanager-connectattachment-corenetworkid
	//
	CoreNetworkId *string `field:"required" json:"coreNetworkId" yaml:"coreNetworkId"`
	// The Region where the edge is located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connectattachment.html#cfn-networkmanager-connectattachment-edgelocation
	//
	EdgeLocation *string `field:"required" json:"edgeLocation" yaml:"edgeLocation"`
	// Options for connecting an attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connectattachment.html#cfn-networkmanager-connectattachment-options
	//
	Options interface{} `field:"required" json:"options" yaml:"options"`
	// The ID of the transport attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connectattachment.html#cfn-networkmanager-connectattachment-transportattachmentid
	//
	TransportAttachmentId *string `field:"required" json:"transportAttachmentId" yaml:"transportAttachmentId"`
	// The attachment to move from one segment to another.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connectattachment.html#cfn-networkmanager-connectattachment-proposedsegmentchange
	//
	ProposedSegmentChange interface{} `field:"optional" json:"proposedSegmentChange" yaml:"proposedSegmentChange"`
	// Tags for the attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connectattachment.html#cfn-networkmanager-connectattachment-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

