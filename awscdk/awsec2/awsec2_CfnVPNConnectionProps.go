package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnVPNConnection`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVPNConnectionProps := &cfnVPNConnectionProps{
//   	customerGatewayId: jsii.String("customerGatewayId"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	staticRoutesOnly: jsii.Boolean(false),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	transitGatewayId: jsii.String("transitGatewayId"),
//   	vpnGatewayId: jsii.String("vpnGatewayId"),
//   	vpnTunnelOptionsSpecifications: []interface{}{
//   		&vpnTunnelOptionsSpecificationProperty{
//   			preSharedKey: jsii.String("preSharedKey"),
//   			tunnelInsideCidr: jsii.String("tunnelInsideCidr"),
//   		},
//   	},
//   }
//
type CfnVPNConnectionProps struct {
	// The ID of the customer gateway at your end of the VPN connection.
	CustomerGatewayId *string `field:"required" json:"customerGatewayId" yaml:"customerGatewayId"`
	// The type of VPN connection.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Indicates whether the VPN connection uses static routes only.
	//
	// Static routes must be used for devices that don't support BGP.
	//
	// If you are creating a VPN connection for a device that does not support Border Gateway Protocol (BGP), you must specify `true` .
	StaticRoutesOnly interface{} `field:"optional" json:"staticRoutesOnly" yaml:"staticRoutesOnly"`
	// Any tags assigned to the VPN connection.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the transit gateway associated with the VPN connection.
	//
	// You must specify either `TransitGatewayId` or `VpnGatewayId` , but not both.
	TransitGatewayId *string `field:"optional" json:"transitGatewayId" yaml:"transitGatewayId"`
	// The ID of the virtual private gateway at the AWS side of the VPN connection.
	//
	// You must specify either `TransitGatewayId` or `VpnGatewayId` , but not both.
	VpnGatewayId *string `field:"optional" json:"vpnGatewayId" yaml:"vpnGatewayId"`
	// The tunnel options for the VPN connection.
	VpnTunnelOptionsSpecifications interface{} `field:"optional" json:"vpnTunnelOptionsSpecifications" yaml:"vpnTunnelOptionsSpecifications"`
}

