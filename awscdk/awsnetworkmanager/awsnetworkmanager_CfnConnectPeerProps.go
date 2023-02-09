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
//   cfnConnectPeerProps := &cfnConnectPeerProps{
//   	connectAttachmentId: jsii.String("connectAttachmentId"),
//   	insideCidrBlocks: []*string{
//   		jsii.String("insideCidrBlocks"),
//   	},
//   	peerAddress: jsii.String("peerAddress"),
//
//   	// the properties below are optional
//   	bgpOptions: &bgpOptionsProperty{
//   		peerAsn: jsii.Number(123),
//   	},
//   	coreNetworkAddress: jsii.String("coreNetworkAddress"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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

