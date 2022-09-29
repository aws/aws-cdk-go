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
//   	bgpOptions: &bgpOptionsProperty{
//   		peerAsn: jsii.Number(123),
//   	},
//   	connectAttachmentId: jsii.String("connectAttachmentId"),
//   	coreNetworkAddress: jsii.String("coreNetworkAddress"),
//   	insideCidrBlocks: []*string{
//   		jsii.String("insideCidrBlocks"),
//   	},
//   	peerAddress: jsii.String("peerAddress"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnConnectPeerProps struct {
	// The BGP peer options.
	BgpOptions interface{} `field:"optional" json:"bgpOptions" yaml:"bgpOptions"`
	// The ID of Connect peer.
	ConnectAttachmentId *string `field:"optional" json:"connectAttachmentId" yaml:"connectAttachmentId"`
	// The IP address of a core network.
	CoreNetworkAddress *string `field:"optional" json:"coreNetworkAddress" yaml:"coreNetworkAddress"`
	// The inside IP addresses used for a Connect peer configuration.
	InsideCidrBlocks *[]*string `field:"optional" json:"insideCidrBlocks" yaml:"insideCidrBlocks"`
	// The IP address of the Connect peer.
	PeerAddress *string `field:"optional" json:"peerAddress" yaml:"peerAddress"`
	// The tags associated with the Connect peer.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

