package previewawsnetworkfirewallmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsnetworkfirewall/previewawsnetworkfirewallmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use the firewall to provide stateful, managed, network firewall and intrusion detection and prevention filtering for your VPCs in Amazon VPC .
//
// The firewall defines the configuration settings for an AWS Network Firewall firewall. The settings include the firewall policy, the subnets in your VPC to use for the firewall endpoints, and any tags that are attached to the firewall AWS resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFirewallPropsMixin := awscdkmixinspreview.Mixins.NewCfnFirewallPropsMixin(&CfnFirewallMixinProps{
//   	AvailabilityZoneChangeProtection: jsii.Boolean(false),
//   	AvailabilityZoneMappings: []interface{}{
//   		&AvailabilityZoneMappingProperty{
//   			AvailabilityZone: jsii.String("availabilityZone"),
//   		},
//   	},
//   	DeleteProtection: jsii.Boolean(false),
//   	Description: jsii.String("description"),
//   	EnabledAnalysisTypes: []*string{
//   		jsii.String("enabledAnalysisTypes"),
//   	},
//   	FirewallName: jsii.String("firewallName"),
//   	FirewallPolicyArn: jsii.String("firewallPolicyArn"),
//   	FirewallPolicyChangeProtection: jsii.Boolean(false),
//   	SubnetChangeProtection: jsii.Boolean(false),
//   	SubnetMappings: []interface{}{
//   		&SubnetMappingProperty{
//   			IpAddressType: jsii.String("ipAddressType"),
//   			SubnetId: jsii.String("subnetId"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TransitGatewayId: jsii.String("transitGatewayId"),
//   	VpcId: jsii.String("vpcId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-firewall.html
//
type CfnFirewallPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnFirewallMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFirewallPropsMixin
type jsiiProxy_CfnFirewallPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnFirewallPropsMixin) Props() *CfnFirewallMixinProps {
	var returns *CfnFirewallMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::NetworkFirewall::Firewall`.
func NewCfnFirewallPropsMixin(props *CfnFirewallMixinProps, options *mixins.CfnPropertyMixinOptions) CfnFirewallPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFirewallPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFirewallPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::NetworkFirewall::Firewall`.
func NewCfnFirewallPropsMixin_Override(c CfnFirewallPropsMixin, props *CfnFirewallMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnFirewallPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFirewallPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFirewallPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFirewallPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnFirewallPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

