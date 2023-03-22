package awsnetworkmanager

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnTransitGatewayRouteTableAttachment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransitGatewayRouteTableAttachmentProps := &cfnTransitGatewayRouteTableAttachmentProps{
//   	peeringId: jsii.String("peeringId"),
//   	transitGatewayRouteTableArn: jsii.String("transitGatewayRouteTableArn"),
//
//   	// the properties below are optional
//   	proposedSegmentChange: &proposedSegmentChangeProperty{
//   		attachmentPolicyRuleNumber: jsii.Number(123),
//   		segmentName: jsii.String("segmentName"),
//   		tags: []cfnTag{
//   			&cfnTag{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	tags: []*cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnTransitGatewayRouteTableAttachmentProps struct {
	// The ID of the transit gateway peering.
	PeeringId *string `field:"required" json:"peeringId" yaml:"peeringId"`
	// The ARN of the transit gateway attachment route table.
	//
	// For example, `"TransitGatewayRouteTableArn": "arn:aws:ec2:us-west-2:123456789012:transit-gateway-route-table/tgw-rtb-9876543210123456"` .
	TransitGatewayRouteTableArn *string `field:"required" json:"transitGatewayRouteTableArn" yaml:"transitGatewayRouteTableArn"`
	// This property is read-only.
	//
	// Values can't be assigned to it.
	ProposedSegmentChange interface{} `field:"optional" json:"proposedSegmentChange" yaml:"proposedSegmentChange"`
	// The list of key-value pairs associated with the transit gateway route table attachment.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

