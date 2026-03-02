package awsdirectconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTransitVirtualInterface`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransitVirtualInterfaceProps := &CfnTransitVirtualInterfaceProps{
//   	BgpPeers: []interface{}{
//   		&BgpPeerProperty{
//   			AddressFamily: jsii.String("addressFamily"),
//   			Asn: jsii.String("asn"),
//
//   			// the properties below are optional
//   			AmazonAddress: jsii.String("amazonAddress"),
//   			AuthKey: jsii.String("authKey"),
//   			BgpPeerId: jsii.String("bgpPeerId"),
//   			CustomerAddress: jsii.String("customerAddress"),
//   		},
//   	},
//   	ConnectionId: jsii.String("connectionId"),
//   	DirectConnectGatewayId: jsii.String("directConnectGatewayId"),
//   	VirtualInterfaceName: jsii.String("virtualInterfaceName"),
//   	Vlan: jsii.Number(123),
//
//   	// the properties below are optional
//   	AllocateTransitVirtualInterfaceRoleArn: jsii.String("allocateTransitVirtualInterfaceRoleArn"),
//   	EnableSiteLink: jsii.Boolean(false),
//   	Mtu: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-transitvirtualinterface.html
//
type CfnTransitVirtualInterfaceProps struct {
	// The BGP peers configured on this virtual interface..
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-transitvirtualinterface.html#cfn-directconnect-transitvirtualinterface-bgppeers
	//
	BgpPeers interface{} `field:"required" json:"bgpPeers" yaml:"bgpPeers"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-transitvirtualinterface.html#cfn-directconnect-transitvirtualinterface-connectionid
	//
	ConnectionId interface{} `field:"required" json:"connectionId" yaml:"connectionId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-transitvirtualinterface.html#cfn-directconnect-transitvirtualinterface-directconnectgatewayid
	//
	DirectConnectGatewayId interface{} `field:"required" json:"directConnectGatewayId" yaml:"directConnectGatewayId"`
	// The name of the virtual interface assigned by the customer network.
	//
	// The name has a maximum of 100 characters. The following are valid characters: a-z, 0-9 and a hyphen (-).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-transitvirtualinterface.html#cfn-directconnect-transitvirtualinterface-virtualinterfacename
	//
	VirtualInterfaceName *string `field:"required" json:"virtualInterfaceName" yaml:"virtualInterfaceName"`
	// The ID of the VLAN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-transitvirtualinterface.html#cfn-directconnect-transitvirtualinterface-vlan
	//
	Vlan *float64 `field:"required" json:"vlan" yaml:"vlan"`
	// The Amazon Resource Name (ARN) of the role to allocate the TransitVifAllocation.
	//
	// Needs directconnect:AllocateTransitVirtualInterface permissions and tag permissions if applicable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-transitvirtualinterface.html#cfn-directconnect-transitvirtualinterface-allocatetransitvirtualinterfacerolearn
	//
	AllocateTransitVirtualInterfaceRoleArn *string `field:"optional" json:"allocateTransitVirtualInterfaceRoleArn" yaml:"allocateTransitVirtualInterfaceRoleArn"`
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
}

