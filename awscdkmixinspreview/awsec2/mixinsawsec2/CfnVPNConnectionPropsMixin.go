package mixinsawsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsec2/mixinsawsec2/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a VPN connection between a virtual private gateway and a VPN customer gateway or a transit gateway and a VPN customer gateway.
//
// To specify a VPN connection between a transit gateway and customer gateway, use the `TransitGatewayId` and `CustomerGatewayId` properties.
//
// To specify a VPN connection between a virtual private gateway and customer gateway, use the `VpnGatewayId` and `CustomerGatewayId` properties.
//
// For more information, see [AWS Site-to-Site VPN](https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html) in the *AWS Site-to-Site VPN User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVPNConnectionPropsMixin := awscdkmixinspreview.Mixins.NewCfnVPNConnectionPropsMixin(&CfnVPNConnectionMixinProps{
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
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnection.html
//
type CfnVPNConnectionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnVPNConnectionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnVPNConnectionPropsMixin
type jsiiProxy_CfnVPNConnectionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnVPNConnectionPropsMixin) Props() *CfnVPNConnectionMixinProps {
	var returns *CfnVPNConnectionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPNConnectionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::VPNConnection`.
func NewCfnVPNConnectionPropsMixin(props *CfnVPNConnectionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnVPNConnectionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnVPNConnectionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnVPNConnectionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::VPNConnection`.
func NewCfnVPNConnectionPropsMixin_Override(c CfnVPNConnectionPropsMixin, props *CfnVPNConnectionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnVPNConnectionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVPNConnectionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVPNConnectionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVPNConnectionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVPNConnectionPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

