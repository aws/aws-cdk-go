package awsroute53globalresolver

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsroute53globalresolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource schema for AWS::Route53GlobalResolver::FirewallRule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnFirewallRulePropsMixin := awscdkcfnpropertymixins.Aws_route53globalresolver.NewCfnFirewallRulePropsMixin(&CfnFirewallRuleMixinProps{
//   	Action: jsii.String("action"),
//   	BlockOverrideDnsType: jsii.String("blockOverrideDnsType"),
//   	BlockOverrideDomain: jsii.String("blockOverrideDomain"),
//   	BlockOverrideTtl: jsii.Number(123),
//   	BlockResponse: jsii.String("blockResponse"),
//   	ClientToken: jsii.String("clientToken"),
//   	ConfidenceThreshold: jsii.String("confidenceThreshold"),
//   	Description: jsii.String("description"),
//   	DnsAdvancedProtection: jsii.String("dnsAdvancedProtection"),
//   	DnsViewId: jsii.String("dnsViewId"),
//   	FirewallDomainListId: jsii.String("firewallDomainListId"),
//   	Name: jsii.String("name"),
//   	Priority: jsii.Number(123),
//   	QType: jsii.String("qType"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-firewallrule.html
//
type CfnFirewallRulePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnFirewallRuleMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFirewallRulePropsMixin
type jsiiProxy_CfnFirewallRulePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnFirewallRulePropsMixin) Props() *CfnFirewallRuleMixinProps {
	var returns *CfnFirewallRuleMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRulePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Route53GlobalResolver::FirewallRule`.
func NewCfnFirewallRulePropsMixin(props *CfnFirewallRuleMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnFirewallRulePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFirewallRulePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFirewallRulePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_route53globalresolver.CfnFirewallRulePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Route53GlobalResolver::FirewallRule`.
func NewCfnFirewallRulePropsMixin_Override(c CfnFirewallRulePropsMixin, props *CfnFirewallRuleMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_route53globalresolver.CfnFirewallRulePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnFirewallRulePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFirewallRulePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_route53globalresolver.CfnFirewallRulePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFirewallRulePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_route53globalresolver.CfnFirewallRulePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFirewallRulePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnFirewallRulePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

