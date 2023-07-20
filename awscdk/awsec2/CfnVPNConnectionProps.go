package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnVPNConnection`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVPNConnectionProps := &CfnVPNConnectionProps{
//   	CustomerGatewayId: jsii.String("customerGatewayId"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	StaticRoutesOnly: jsii.Boolean(false),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TransitGatewayId: jsii.String("transitGatewayId"),
//   	VpnGatewayId: jsii.String("vpnGatewayId"),
//   	VpnTunnelOptionsSpecifications: []interface{}{
//   		&VpnTunnelOptionsSpecificationProperty{
//   			PreSharedKey: jsii.String("preSharedKey"),
//   			TunnelInsideCidr: jsii.String("tunnelInsideCidr"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnection.html
//
type CfnVPNConnectionProps struct {
	// The ID of the customer gateway at your end of the VPN connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnection.html#cfn-ec2-vpnconnection-customergatewayid
	//
	CustomerGatewayId *string `field:"required" json:"customerGatewayId" yaml:"customerGatewayId"`
	// The type of VPN connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnection.html#cfn-ec2-vpnconnection-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Indicates whether the VPN connection uses static routes only.
	//
	// Static routes must be used for devices that don't support BGP.
	//
	// If you are creating a VPN connection for a device that does not support Border Gateway Protocol (BGP), you must specify `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnection.html#cfn-ec2-vpnconnection-staticroutesonly
	//
	StaticRoutesOnly interface{} `field:"optional" json:"staticRoutesOnly" yaml:"staticRoutesOnly"`
	// Any tags assigned to the VPN connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnection.html#cfn-ec2-vpnconnection-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the transit gateway associated with the VPN connection.
	//
	// You must specify either `TransitGatewayId` or `VpnGatewayId` , but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnection.html#cfn-ec2-vpnconnection-transitgatewayid
	//
	TransitGatewayId *string `field:"optional" json:"transitGatewayId" yaml:"transitGatewayId"`
	// The ID of the virtual private gateway at the AWS side of the VPN connection.
	//
	// You must specify either `TransitGatewayId` or `VpnGatewayId` , but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnection.html#cfn-ec2-vpnconnection-vpngatewayid
	//
	VpnGatewayId *string `field:"optional" json:"vpnGatewayId" yaml:"vpnGatewayId"`
	// The tunnel options for the VPN connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnection.html#cfn-ec2-vpnconnection-vpntunneloptionsspecifications
	//
	VpnTunnelOptionsSpecifications interface{} `field:"optional" json:"vpnTunnelOptionsSpecifications" yaml:"vpnTunnelOptionsSpecifications"`
}

