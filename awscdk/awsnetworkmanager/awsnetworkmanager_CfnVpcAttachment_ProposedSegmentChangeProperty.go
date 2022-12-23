package awsnetworkmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   proposedSegmentChangeProperty := &proposedSegmentChangeProperty{
//   	attachmentPolicyRuleNumber: jsii.Number(123),
//   	segmentName: jsii.String("segmentName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnVpcAttachment_ProposedSegmentChangeProperty struct {
	// `CfnVpcAttachment.ProposedSegmentChangeProperty.AttachmentPolicyRuleNumber`.
	AttachmentPolicyRuleNumber *float64 `field:"optional" json:"attachmentPolicyRuleNumber" yaml:"attachmentPolicyRuleNumber"`
	// `CfnVpcAttachment.ProposedSegmentChangeProperty.SegmentName`.
	SegmentName *string `field:"optional" json:"segmentName" yaml:"segmentName"`
	// `CfnVpcAttachment.ProposedSegmentChangeProperty.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

