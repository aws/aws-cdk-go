package previewawsec2mixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnVPNConnectionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVPNConnectionMixinProps := &CfnVPNConnectionMixinProps{
//   	CustomerGatewayId: jsii.String("customerGatewayId"),
//   	EnableAcceleration: jsii.Boolean(false),
//   	LocalIpv4NetworkCidr: jsii.String("localIpv4NetworkCidr"),
//   	LocalIpv6NetworkCidr: jsii.String("localIpv6NetworkCidr"),
//   	OutsideIpAddressType: jsii.String("outsideIpAddressType"),
//   	PreSharedKeyStorage: jsii.String("preSharedKeyStorage"),
//   	RemoteIpv4NetworkCidr: jsii.String("remoteIpv4NetworkCidr"),
//   	RemoteIpv6NetworkCidr: jsii.String("remoteIpv6NetworkCidr"),
//   	StaticRoutesOnly: jsii.Boolean(false),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TransitGatewayId: jsii.String("transitGatewayId"),
//   	TransportTransitGatewayAttachmentId: jsii.String("transportTransitGatewayAttachmentId"),
//   	TunnelBandwidth: jsii.String("tunnelBandwidth"),
//   	TunnelInsideIpVersion: jsii.String("tunnelInsideIpVersion"),
//   	Type: jsii.String("type"),
//   	VpnConcentratorId: jsii.String("vpnConcentratorId"),
//   	VpnGatewayId: jsii.String("vpnGatewayId"),
//   	VpnTunnelOptionsSpecifications: []interface{}{
//   		&VpnTunnelOptionsSpecificationProperty{
//   			DpdTimeoutAction: jsii.String("dpdTimeoutAction"),
//   			DpdTimeoutSeconds: jsii.Number(123),
//   			EnableTunnelLifecycleControl: jsii.Boolean(false),
//   			IkeVersions: []interface{}{
//   				map[string]*string{
//   					"value": jsii.String("value"),
//   				},
//   			},
//   			LogOptions: &VpnTunnelLogOptionsSpecificationProperty{
//   				CloudwatchLogOptions: &CloudwatchLogOptionsSpecificationProperty{
//   					BgpLogEnabled: jsii.Boolean(false),
//   					BgpLogGroupArn: jsii.String("bgpLogGroupArn"),
//   					BgpLogOutputFormat: jsii.String("bgpLogOutputFormat"),
//   					LogEnabled: jsii.Boolean(false),
//   					LogGroupArn: jsii.String("logGroupArn"),
//   					LogOutputFormat: jsii.String("logOutputFormat"),
//   				},
//   			},
//   			Phase1DhGroupNumbers: []interface{}{
//   				&Phase1DHGroupNumbersRequestListValueProperty{
//   					Value: jsii.Number(123),
//   				},
//   			},
//   			Phase1EncryptionAlgorithms: []interface{}{
//   				&Phase1EncryptionAlgorithmsRequestListValueProperty{
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Phase1IntegrityAlgorithms: []interface{}{
//   				&Phase1IntegrityAlgorithmsRequestListValueProperty{
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Phase1LifetimeSeconds: jsii.Number(123),
//   			Phase2DhGroupNumbers: []interface{}{
//   				&Phase2DHGroupNumbersRequestListValueProperty{
//   					Value: jsii.Number(123),
//   				},
//   			},
//   			Phase2EncryptionAlgorithms: []interface{}{
//   				&Phase2EncryptionAlgorithmsRequestListValueProperty{
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Phase2IntegrityAlgorithms: []interface{}{
//   				&Phase2IntegrityAlgorithmsRequestListValueProperty{
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Phase2LifetimeSeconds: jsii.Number(123),
//   			PreSharedKey: jsii.String("preSharedKey"),
//   			RekeyFuzzPercentage: jsii.Number(123),
//   			RekeyMarginTimeSeconds: jsii.Number(123),
//   			ReplayWindowSize: jsii.Number(123),
//   			StartupAction: jsii.String("startupAction"),
//   			TunnelInsideCidr: jsii.String("tunnelInsideCidr"),
//   			TunnelInsideIpv6Cidr: jsii.String("tunnelInsideIpv6Cidr"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnection.html
//
type CfnVPNConnectionMixinProps struct {
	// The ID of the customer gateway at your end of the VPN connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnection.html#cfn-ec2-vpnconnection-customergatewayid
	//
	CustomerGatewayId *string `field:"optional" json:"customerGatewayId" yaml:"customerGatewayId"`
	// Indicate whether to enable acceleration for the VPN connection.
	//
	// Default: `false`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnection.html#cfn-ec2-vpnconnection-enableacceleration
	//
	EnableAcceleration interface{} `field:"optional" json:"enableAcceleration" yaml:"enableAcceleration"`
	// The IPv4 CIDR on the customer gateway (on-premises) side of the VPN connection.
	//
	// Default: `0.0.0.0/0`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnection.html#cfn-ec2-vpnconnection-localipv4networkcidr
	//
	LocalIpv4NetworkCidr *string `field:"optional" json:"localIpv4NetworkCidr" yaml:"localIpv4NetworkCidr"`
	// The IPv6 CIDR on the customer gateway (on-premises) side of the VPN connection.
	//
	// Default: `::/0`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnection.html#cfn-ec2-vpnconnection-localipv6networkcidr
	//
	LocalIpv6NetworkCidr *string `field:"optional" json:"localIpv6NetworkCidr" yaml:"localIpv6NetworkCidr"`
	// The type of IP address assigned to the outside interface of the customer gateway device.
	//
	// Valid values: `PrivateIpv4` | `PublicIpv4` | `Ipv6`
	//
	// Default: `PublicIpv4`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnection.html#cfn-ec2-vpnconnection-outsideipaddresstype
	//
	OutsideIpAddressType *string `field:"optional" json:"outsideIpAddressType" yaml:"outsideIpAddressType"`
	// Describes the storage location for an instance store-backed AMI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnection.html#cfn-ec2-vpnconnection-presharedkeystorage
	//
	PreSharedKeyStorage *string `field:"optional" json:"preSharedKeyStorage" yaml:"preSharedKeyStorage"`
	// The IPv4 CIDR on the AWS side of the VPN connection.
	//
	// Default: `0.0.0.0/0`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnection.html#cfn-ec2-vpnconnection-remoteipv4networkcidr
	//
	RemoteIpv4NetworkCidr *string `field:"optional" json:"remoteIpv4NetworkCidr" yaml:"remoteIpv4NetworkCidr"`
	// The IPv6 CIDR on the AWS side of the VPN connection.
	//
	// Default: `::/0`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnection.html#cfn-ec2-vpnconnection-remoteipv6networkcidr
	//
	RemoteIpv6NetworkCidr *string `field:"optional" json:"remoteIpv6NetworkCidr" yaml:"remoteIpv6NetworkCidr"`
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
	// The transit gateway attachment ID to use for the VPN tunnel.
	//
	// Required if `OutsideIpAddressType` is set to `PrivateIpv4` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnection.html#cfn-ec2-vpnconnection-transporttransitgatewayattachmentid
	//
	TransportTransitGatewayAttachmentId *string `field:"optional" json:"transportTransitGatewayAttachmentId" yaml:"transportTransitGatewayAttachmentId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnection.html#cfn-ec2-vpnconnection-tunnelbandwidth
	//
	// Default: - "standard".
	//
	TunnelBandwidth *string `field:"optional" json:"tunnelBandwidth" yaml:"tunnelBandwidth"`
	// Indicate whether the VPN tunnels process IPv4 or IPv6 traffic.
	//
	// Default: `ipv4`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnection.html#cfn-ec2-vpnconnection-tunnelinsideipversion
	//
	TunnelInsideIpVersion *string `field:"optional" json:"tunnelInsideIpVersion" yaml:"tunnelInsideIpVersion"`
	// The type of VPN connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnection.html#cfn-ec2-vpnconnection-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnection.html#cfn-ec2-vpnconnection-vpnconcentratorid
	//
	VpnConcentratorId *string `field:"optional" json:"vpnConcentratorId" yaml:"vpnConcentratorId"`
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

