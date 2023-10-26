package awsnetworkmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
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
//   	PeerAddress: jsii.String("peerAddress"),
//
//   	// the properties below are optional
//   	BgpOptions: &BgpOptionsProperty{
//   		PeerAsn: jsii.Number(123),
//   	},
//   	CoreNetworkAddress: jsii.String("coreNetworkAddress"),
//   	InsideCidrBlocks: []*string{
//   		jsii.String("insideCidrBlocks"),
//   	},
//   	SubnetArn: jsii.String("subnetArn"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connectpeer.html
//
type CfnConnectPeerProps struct {
	// The ID of the attachment to connect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connectpeer.html#cfn-networkmanager-connectpeer-connectattachmentid
	//
	ConnectAttachmentId *string `field:"required" json:"connectAttachmentId" yaml:"connectAttachmentId"`
	// The IP address of the Connect peer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connectpeer.html#cfn-networkmanager-connectpeer-peeraddress
	//
	PeerAddress *string `field:"required" json:"peerAddress" yaml:"peerAddress"`
	// Describes the BGP options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connectpeer.html#cfn-networkmanager-connectpeer-bgpoptions
	//
	BgpOptions interface{} `field:"optional" json:"bgpOptions" yaml:"bgpOptions"`
	// The IP address of a core network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connectpeer.html#cfn-networkmanager-connectpeer-corenetworkaddress
	//
	CoreNetworkAddress *string `field:"optional" json:"coreNetworkAddress" yaml:"coreNetworkAddress"`
	// The inside IP addresses used for a Connect peer configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connectpeer.html#cfn-networkmanager-connectpeer-insidecidrblocks
	//
	InsideCidrBlocks *[]*string `field:"optional" json:"insideCidrBlocks" yaml:"insideCidrBlocks"`
	// The subnet ARN for the connect peer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connectpeer.html#cfn-networkmanager-connectpeer-subnetarn
	//
	SubnetArn *string `field:"optional" json:"subnetArn" yaml:"subnetArn"`
	// The list of key-value tags associated with the Connect peer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connectpeer.html#cfn-networkmanager-connectpeer-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

