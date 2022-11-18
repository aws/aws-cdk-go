package awsnetworkmanager

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnConnectAttachment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConnectAttachmentProps := &cfnConnectAttachmentProps{
//   	coreNetworkId: jsii.String("coreNetworkId"),
//   	edgeLocation: jsii.String("edgeLocation"),
//   	options: &connectAttachmentOptionsProperty{
//   		protocol: jsii.String("protocol"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	transportAttachmentId: jsii.String("transportAttachmentId"),
//   }
//
type CfnConnectAttachmentProps struct {
	// The core network ID.
	CoreNetworkId *string `field:"optional" json:"coreNetworkId" yaml:"coreNetworkId"`
	// The Region where the edge is located.
	EdgeLocation *string `field:"optional" json:"edgeLocation" yaml:"edgeLocation"`
	// Options for creating a Connect attachment.
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// The tags associated with the Connect attachment.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the attachment between the two connections.
	TransportAttachmentId *string `field:"optional" json:"transportAttachmentId" yaml:"transportAttachmentId"`
}

