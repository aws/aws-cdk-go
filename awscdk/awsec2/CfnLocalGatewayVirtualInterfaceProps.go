package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnLocalGatewayVirtualInterface`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLocalGatewayVirtualInterfaceProps := &CfnLocalGatewayVirtualInterfaceProps{
//   	LocalAddress: jsii.String("localAddress"),
//   	LocalGatewayVirtualInterfaceGroupId: jsii.String("localGatewayVirtualInterfaceGroupId"),
//   	OutpostLagId: jsii.String("outpostLagId"),
//   	PeerAddress: jsii.String("peerAddress"),
//   	Vlan: jsii.Number(123),
//
//   	// the properties below are optional
//   	PeerBgpAsn: jsii.Number(123),
//   	PeerBgpAsnExtended: jsii.Number(123),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayvirtualinterface.html
//
type CfnLocalGatewayVirtualInterfaceProps struct {
	// The local address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayvirtualinterface.html#cfn-ec2-localgatewayvirtualinterface-localaddress
	//
	LocalAddress *string `field:"required" json:"localAddress" yaml:"localAddress"`
	// The ID of the local gateway virtual interface group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayvirtualinterface.html#cfn-ec2-localgatewayvirtualinterface-localgatewayvirtualinterfacegroupid
	//
	LocalGatewayVirtualInterfaceGroupId *string `field:"required" json:"localGatewayVirtualInterfaceGroupId" yaml:"localGatewayVirtualInterfaceGroupId"`
	// The Outpost LAG ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayvirtualinterface.html#cfn-ec2-localgatewayvirtualinterface-outpostlagid
	//
	OutpostLagId *string `field:"required" json:"outpostLagId" yaml:"outpostLagId"`
	// The peer address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayvirtualinterface.html#cfn-ec2-localgatewayvirtualinterface-peeraddress
	//
	PeerAddress *string `field:"required" json:"peerAddress" yaml:"peerAddress"`
	// The ID of the VLAN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayvirtualinterface.html#cfn-ec2-localgatewayvirtualinterface-vlan
	//
	Vlan *float64 `field:"required" json:"vlan" yaml:"vlan"`
	// The peer BGP ASN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayvirtualinterface.html#cfn-ec2-localgatewayvirtualinterface-peerbgpasn
	//
	PeerBgpAsn *float64 `field:"optional" json:"peerBgpAsn" yaml:"peerBgpAsn"`
	// The extended 32-bit ASN of the BGP peer for use with larger ASN values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayvirtualinterface.html#cfn-ec2-localgatewayvirtualinterface-peerbgpasnextended
	//
	PeerBgpAsnExtended *float64 `field:"optional" json:"peerBgpAsnExtended" yaml:"peerBgpAsnExtended"`
	// The tags assigned to the virtual interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayvirtualinterface.html#cfn-ec2-localgatewayvirtualinterface-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

