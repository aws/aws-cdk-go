package previewawsdirectconnectmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnTransitVirtualInterfacePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTransitVirtualInterfaceMixinProps := &CfnTransitVirtualInterfaceMixinProps{
//   	AllocateTransitVirtualInterfaceRoleArn: jsii.String("allocateTransitVirtualInterfaceRoleArn"),
//   	BgpPeers: []interface{}{
//   		&BgpPeerProperty{
//   			AddressFamily: jsii.String("addressFamily"),
//   			AmazonAddress: jsii.String("amazonAddress"),
//   			Asn: jsii.String("asn"),
//   			AuthKey: jsii.String("authKey"),
//   			BgpPeerId: jsii.String("bgpPeerId"),
//   			CustomerAddress: jsii.String("customerAddress"),
//   		},
//   	},
//   	ConnectionId: jsii.String("connectionId"),
//   	DirectConnectGatewayId: jsii.String("directConnectGatewayId"),
//   	EnableSiteLink: jsii.Boolean(false),
//   	Mtu: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VirtualInterfaceName: jsii.String("virtualInterfaceName"),
//   	Vlan: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-transitvirtualinterface.html
//
type CfnTransitVirtualInterfaceMixinProps struct {
	// The Amazon Resource Name (ARN) of the role to allocate the TransitVifAllocation.
	//
	// Needs directconnect:AllocateTransitVirtualInterface permissions and tag permissions if applicable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-transitvirtualinterface.html#cfn-directconnect-transitvirtualinterface-allocatetransitvirtualinterfacerolearn
	//
	AllocateTransitVirtualInterfaceRoleArn *string `field:"optional" json:"allocateTransitVirtualInterfaceRoleArn" yaml:"allocateTransitVirtualInterfaceRoleArn"`
	// The BGP peers configured on this virtual interface..
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-transitvirtualinterface.html#cfn-directconnect-transitvirtualinterface-bgppeers
	//
	BgpPeers interface{} `field:"optional" json:"bgpPeers" yaml:"bgpPeers"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-transitvirtualinterface.html#cfn-directconnect-transitvirtualinterface-connectionid
	//
	ConnectionId *string `field:"optional" json:"connectionId" yaml:"connectionId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-transitvirtualinterface.html#cfn-directconnect-transitvirtualinterface-directconnectgatewayid
	//
	DirectConnectGatewayId *string `field:"optional" json:"directConnectGatewayId" yaml:"directConnectGatewayId"`
	// Indicates whether to enable or disable SiteLink.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-transitvirtualinterface.html#cfn-directconnect-transitvirtualinterface-enablesitelink
	//
	EnableSiteLink interface{} `field:"optional" json:"enableSiteLink" yaml:"enableSiteLink"`
	// The maximum transmission unit (MTU), in bytes.
	//
	// The supported values are 1500 and 9001. The default value is 1500.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-transitvirtualinterface.html#cfn-directconnect-transitvirtualinterface-mtu
	//
	Mtu *float64 `field:"optional" json:"mtu" yaml:"mtu"`
	// The tags associated with the private virtual interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-transitvirtualinterface.html#cfn-directconnect-transitvirtualinterface-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the virtual interface assigned by the customer network.
	//
	// The name has a maximum of 100 characters. The following are valid characters: a-z, 0-9 and a hyphen (-).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-transitvirtualinterface.html#cfn-directconnect-transitvirtualinterface-virtualinterfacename
	//
	VirtualInterfaceName *string `field:"optional" json:"virtualInterfaceName" yaml:"virtualInterfaceName"`
	// The ID of the VLAN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-transitvirtualinterface.html#cfn-directconnect-transitvirtualinterface-vlan
	//
	Vlan *float64 `field:"optional" json:"vlan" yaml:"vlan"`
}

