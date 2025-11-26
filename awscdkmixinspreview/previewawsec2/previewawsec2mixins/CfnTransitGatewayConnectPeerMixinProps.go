package previewawsec2mixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnTransitGatewayConnectPeerPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTransitGatewayConnectPeerMixinProps := &CfnTransitGatewayConnectPeerMixinProps{
//   	ConnectPeerConfiguration: &TransitGatewayConnectPeerConfigurationProperty{
//   		BgpConfigurations: []interface{}{
//   			&TransitGatewayAttachmentBgpConfigurationProperty{
//   				BgpStatus: jsii.String("bgpStatus"),
//   				PeerAddress: jsii.String("peerAddress"),
//   				PeerAsn: jsii.Number(123),
//   				TransitGatewayAddress: jsii.String("transitGatewayAddress"),
//   				TransitGatewayAsn: jsii.Number(123),
//   			},
//   		},
//   		InsideCidrBlocks: []*string{
//   			jsii.String("insideCidrBlocks"),
//   		},
//   		PeerAddress: jsii.String("peerAddress"),
//   		Protocol: jsii.String("protocol"),
//   		TransitGatewayAddress: jsii.String("transitGatewayAddress"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TransitGatewayAttachmentId: jsii.String("transitGatewayAttachmentId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayconnectpeer.html
//
type CfnTransitGatewayConnectPeerMixinProps struct {
	// The Connect peer details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayconnectpeer.html#cfn-ec2-transitgatewayconnectpeer-connectpeerconfiguration
	//
	ConnectPeerConfiguration interface{} `field:"optional" json:"connectPeerConfiguration" yaml:"connectPeerConfiguration"`
	// The tags for the Connect peer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayconnectpeer.html#cfn-ec2-transitgatewayconnectpeer-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the Connect attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayconnectpeer.html#cfn-ec2-transitgatewayconnectpeer-transitgatewayattachmentid
	//
	TransitGatewayAttachmentId *string `field:"optional" json:"transitGatewayAttachmentId" yaml:"transitGatewayAttachmentId"`
}

