package awsnetworkmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnConnectionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnConnectionMixinProps := &CfnConnectionMixinProps{
//   	ConnectedDeviceId: jsii.String("connectedDeviceId"),
//   	ConnectedLinkId: jsii.String("connectedLinkId"),
//   	Description: jsii.String("description"),
//   	DeviceId: jsii.String("deviceId"),
//   	GlobalNetworkId: jsii.String("globalNetworkId"),
//   	LinkId: jsii.String("linkId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connection.html
//
type CfnConnectionMixinProps struct {
	// The ID of the second device in the connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connection.html#cfn-networkmanager-connection-connecteddeviceid
	//
	ConnectedDeviceId *string `field:"optional" json:"connectedDeviceId" yaml:"connectedDeviceId"`
	// The ID of the link for the second device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connection.html#cfn-networkmanager-connection-connectedlinkid
	//
	ConnectedLinkId *string `field:"optional" json:"connectedLinkId" yaml:"connectedLinkId"`
	// The description of the connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connection.html#cfn-networkmanager-connection-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the first device in the connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connection.html#cfn-networkmanager-connection-deviceid
	//
	DeviceId *string `field:"optional" json:"deviceId" yaml:"deviceId"`
	// The ID of the global network that the connection belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connection.html#cfn-networkmanager-connection-globalnetworkid
	//
	GlobalNetworkId *string `field:"optional" json:"globalNetworkId" yaml:"globalNetworkId"`
	// The ID of the link for the first device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connection.html#cfn-networkmanager-connection-linkid
	//
	LinkId *string `field:"optional" json:"linkId" yaml:"linkId"`
	// The tags for the connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-connection.html#cfn-networkmanager-connection-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

