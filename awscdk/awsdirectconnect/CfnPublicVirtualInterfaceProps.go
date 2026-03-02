package awsdirectconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPublicVirtualInterface`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPublicVirtualInterfaceProps := &CfnPublicVirtualInterfaceProps{
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
//   	AllocatePublicVirtualInterfaceRoleArn: jsii.String("allocatePublicVirtualInterfaceRoleArn"),
//   	RouteFilterPrefixes: []*string{
//   		jsii.String("routeFilterPrefixes"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-publicvirtualinterface.html
//
type CfnPublicVirtualInterfaceProps struct {
	// The BGP peers configured on this virtual interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-publicvirtualinterface.html#cfn-directconnect-publicvirtualinterface-bgppeers
	//
	BgpPeers interface{} `field:"required" json:"bgpPeers" yaml:"bgpPeers"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-publicvirtualinterface.html#cfn-directconnect-publicvirtualinterface-connectionid
	//
	ConnectionId interface{} `field:"required" json:"connectionId" yaml:"connectionId"`
	// The name of the virtual interface assigned by the customer network.
	//
	// The name has a maximum of 100 characters. The following are valid characters: a-z, 0-9 and a hyphen (-).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-publicvirtualinterface.html#cfn-directconnect-publicvirtualinterface-virtualinterfacename
	//
	VirtualInterfaceName *string `field:"required" json:"virtualInterfaceName" yaml:"virtualInterfaceName"`
	// The ID of the VLAN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-publicvirtualinterface.html#cfn-directconnect-publicvirtualinterface-vlan
	//
	Vlan *float64 `field:"required" json:"vlan" yaml:"vlan"`
	// The Amazon Resource Name (ARN) of the role to allocate the public virtual interface.
	//
	// Needs directconnect:AllocatePublicVirtualInterface permissions and tag permissions if applicable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-publicvirtualinterface.html#cfn-directconnect-publicvirtualinterface-allocatepublicvirtualinterfacerolearn
	//
	AllocatePublicVirtualInterfaceRoleArn *string `field:"optional" json:"allocatePublicVirtualInterfaceRoleArn" yaml:"allocatePublicVirtualInterfaceRoleArn"`
	// The routes to be advertised to the AWS network in this region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-publicvirtualinterface.html#cfn-directconnect-publicvirtualinterface-routefilterprefixes
	//
	RouteFilterPrefixes *[]*string `field:"optional" json:"routeFilterPrefixes" yaml:"routeFilterPrefixes"`
	// The tags associated with the public virtual interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-publicvirtualinterface.html#cfn-directconnect-publicvirtualinterface-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

