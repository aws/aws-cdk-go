package awsnetworkmanager

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Describes a proposed segment change.
//
// In some cases, the segment change must first be evaluated and accepted.
//
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
type CfnConnectAttachment_ProposedSegmentChangeProperty struct {
	// The rule number in the policy document that applies to this change.
	AttachmentPolicyRuleNumber *float64 `field:"optional" json:"attachmentPolicyRuleNumber" yaml:"attachmentPolicyRuleNumber"`
	// The name of the segment to change.
	SegmentName *string `field:"optional" json:"segmentName" yaml:"segmentName"`
	// The list of key-value tags that changed for the segment.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

