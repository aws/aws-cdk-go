package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTransitGatewayConnectPeer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransitGatewayConnectPeerProps := &CfnTransitGatewayConnectPeerProps{
//   	ConnectPeerConfiguration: &TransitGatewayConnectPeerConfigurationProperty{
//   		InsideCidrBlocks: []*string{
//   			jsii.String("insideCidrBlocks"),
//   		},
//   		PeerAddress: jsii.String("peerAddress"),
//
//   		// the properties below are optional
//   		BgpConfigurations: []interface{}{
//   			&TransitGatewayAttachmentBgpConfigurationProperty{
//   				BgpStatus: jsii.String("bgpStatus"),
//   				PeerAddress: jsii.String("peerAddress"),
//   				PeerAsn: jsii.Number(123),
//   				TransitGatewayAddress: jsii.String("transitGatewayAddress"),
//   				TransitGatewayAsn: jsii.Number(123),
//   			},
//   		},
//   		Protocol: jsii.String("protocol"),
//   		TransitGatewayAddress: jsii.String("transitGatewayAddress"),
//   	},
//   	TransitGatewayAttachmentId: jsii.String("transitGatewayAttachmentId"),
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayconnectpeer.html
//
type CfnTransitGatewayConnectPeerProps struct {
	// The Connect peer details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayconnectpeer.html#cfn-ec2-transitgatewayconnectpeer-connectpeerconfiguration
	//
	ConnectPeerConfiguration interface{} `field:"required" json:"connectPeerConfiguration" yaml:"connectPeerConfiguration"`
	// The ID of the Connect attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayconnectpeer.html#cfn-ec2-transitgatewayconnectpeer-transitgatewayattachmentid
	//
	TransitGatewayAttachmentId *string `field:"required" json:"transitGatewayAttachmentId" yaml:"transitGatewayAttachmentId"`
	// The tags for the Connect peer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayconnectpeer.html#cfn-ec2-transitgatewayconnectpeer-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

