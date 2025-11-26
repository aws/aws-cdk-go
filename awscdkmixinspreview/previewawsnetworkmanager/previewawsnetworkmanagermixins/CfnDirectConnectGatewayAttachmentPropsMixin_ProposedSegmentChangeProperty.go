package previewawsnetworkmanagermixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Describes a proposed segment change.
//
// In some cases, the segment change must first be evaluated and accepted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   proposedSegmentChangeProperty := &ProposedSegmentChangeProperty{
//   	AttachmentPolicyRuleNumber: jsii.Number(123),
//   	SegmentName: jsii.String("segmentName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-directconnectgatewayattachment-proposedsegmentchange.html
//
type CfnDirectConnectGatewayAttachmentPropsMixin_ProposedSegmentChangeProperty struct {
	// The rule number in the policy document that applies to this change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-directconnectgatewayattachment-proposedsegmentchange.html#cfn-networkmanager-directconnectgatewayattachment-proposedsegmentchange-attachmentpolicyrulenumber
	//
	AttachmentPolicyRuleNumber *float64 `field:"optional" json:"attachmentPolicyRuleNumber" yaml:"attachmentPolicyRuleNumber"`
	// The name of the segment to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-directconnectgatewayattachment-proposedsegmentchange.html#cfn-networkmanager-directconnectgatewayattachment-proposedsegmentchange-segmentname
	//
	SegmentName *string `field:"optional" json:"segmentName" yaml:"segmentName"`
	// The list of key-value tags that changed for the segment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-directconnectgatewayattachment-proposedsegmentchange.html#cfn-networkmanager-directconnectgatewayattachment-proposedsegmentchange-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

