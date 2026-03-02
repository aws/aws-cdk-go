package awsdirectconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPrivateVirtualInterface`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPrivateVirtualInterfaceProps := &CfnPrivateVirtualInterfaceProps{
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
//   	VirtualInterfaceName: jsii.String("virtualInterfaceName"),
//   	Vlan: jsii.Number(123),
//
//   	// the properties below are optional
//   	AllocatePrivateVirtualInterfaceRoleArn: jsii.String("allocatePrivateVirtualInterfaceRoleArn"),
//   	DirectConnectGatewayId: jsii.String("directConnectGatewayId"),
//   	EnableSiteLink: jsii.Boolean(false),
//   	Mtu: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VirtualGatewayId: jsii.String("virtualGatewayId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-privatevirtualinterface.html
//
type CfnPrivateVirtualInterfaceProps struct {
	// The BGP peers configured on this virtual interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-privatevirtualinterface.html#cfn-directconnect-privatevirtualinterface-bgppeers
	//
	BgpPeers interface{} `field:"required" json:"bgpPeers" yaml:"bgpPeers"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-privatevirtualinterface.html#cfn-directconnect-privatevirtualinterface-connectionid
	//
	ConnectionId interface{} `field:"required" json:"connectionId" yaml:"connectionId"`
	// The name of the virtual interface assigned by the customer network.
	//
	// The name has a maximum of 100 characters. The following are valid characters: a-z, 0-9 and a hyphen (-).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-privatevirtualinterface.html#cfn-directconnect-privatevirtualinterface-virtualinterfacename
	//
	VirtualInterfaceName *string `field:"required" json:"virtualInterfaceName" yaml:"virtualInterfaceName"`
	// The ID of the VLAN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-privatevirtualinterface.html#cfn-directconnect-privatevirtualinterface-vlan
	//
	Vlan *float64 `field:"required" json:"vlan" yaml:"vlan"`
	// The Amazon Resource Name (ARN) of the role to allocate the private virtual interface.
	//
	// Needs directconnect:AllocatePrivateVirtualInterface permissions and tag permissions if applicable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-privatevirtualinterface.html#cfn-directconnect-privatevirtualinterface-allocateprivatevirtualinterfacerolearn
	//
	AllocatePrivateVirtualInterfaceRoleArn *string `field:"optional" json:"allocatePrivateVirtualInterfaceRoleArn" yaml:"allocatePrivateVirtualInterfaceRoleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-privatevirtualinterface.html#cfn-directconnect-privatevirtualinterface-directconnectgatewayid
	//
	DirectConnectGatewayId interface{} `field:"optional" json:"directConnectGatewayId" yaml:"directConnectGatewayId"`
	// Indicates whether to enable or disable SiteLink.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-privatevirtualinterface.html#cfn-directconnect-privatevirtualinterface-enablesitelink
	//
	EnableSiteLink interface{} `field:"optional" json:"enableSiteLink" yaml:"enableSiteLink"`
	// The maximum transmission unit (MTU), in bytes.
	//
	// The supported values are 1500 and 9001. The default value is 1500.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-privatevirtualinterface.html#cfn-directconnect-privatevirtualinterface-mtu
	//
	Mtu *float64 `field:"optional" json:"mtu" yaml:"mtu"`
	// The tags associated with the private virtual interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-privatevirtualinterface.html#cfn-directconnect-privatevirtualinterface-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ID or ARN of the virtual private gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-privatevirtualinterface.html#cfn-directconnect-privatevirtualinterface-virtualgatewayid
	//
	VirtualGatewayId *string `field:"optional" json:"virtualGatewayId" yaml:"virtualGatewayId"`
}

