package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnTransitGatewayPeeringAttachment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransitGatewayPeeringAttachmentProps := &cfnTransitGatewayPeeringAttachmentProps{
//   	peerAccountId: jsii.String("peerAccountId"),
//   	peerRegion: jsii.String("peerRegion"),
//   	peerTransitGatewayId: jsii.String("peerTransitGatewayId"),
//   	transitGatewayId: jsii.String("transitGatewayId"),
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnTransitGatewayPeeringAttachmentProps struct {
	// The ID of the AWS account that owns the transit gateway.
	PeerAccountId *string `field:"required" json:"peerAccountId" yaml:"peerAccountId"`
	// The Region of the transit gateway.
	PeerRegion *string `field:"required" json:"peerRegion" yaml:"peerRegion"`
	// The ID of the transit gateway.
	PeerTransitGatewayId *string `field:"required" json:"peerTransitGatewayId" yaml:"peerTransitGatewayId"`
	// The ID of the transit gateway peering attachment.
	TransitGatewayId *string `field:"required" json:"transitGatewayId" yaml:"transitGatewayId"`
	// The tags for the transit gateway peering attachment.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

