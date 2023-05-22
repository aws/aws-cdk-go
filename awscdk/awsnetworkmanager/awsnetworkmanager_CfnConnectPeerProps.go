package awsnetworkmanager

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnConnectPeer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConnectPeerProps := &CfnConnectPeerProps{
//   	ConnectAttachmentId: jsii.String("connectAttachmentId"),
//   	InsideCidrBlocks: []*string{
//   		jsii.String("insideCidrBlocks"),
//   	},
//   	PeerAddress: jsii.String("peerAddress"),
//
//   	// the properties below are optional
//   	BgpOptions: &BgpOptionsProperty{
//   		PeerAsn: jsii.Number(123),
//   	},
//   	CoreNetworkAddress: jsii.String("coreNetworkAddress"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnConnectPeerProps struct {
	// The ID of the attachment to connect.
	ConnectAttachmentId *string `field:"required" json:"connectAttachmentId" yaml:"connectAttachmentId"`
	// The inside IP addresses used for a Connect peer configuration.
	InsideCidrBlocks *[]*string `field:"required" json:"insideCidrBlocks" yaml:"insideCidrBlocks"`
	// The IP address of the Connect peer.
	PeerAddress *string `field:"required" json:"peerAddress" yaml:"peerAddress"`
	// `AWS::NetworkManager::ConnectPeer.BgpOptions`.
	BgpOptions interface{} `field:"optional" json:"bgpOptions" yaml:"bgpOptions"`
	// The IP address of a core network.
	CoreNetworkAddress *string `field:"optional" json:"coreNetworkAddress" yaml:"coreNetworkAddress"`
	// The list of key-value tags associated with the Connect peer.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

