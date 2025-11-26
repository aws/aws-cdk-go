package previewawsroute53resolvermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsroute53resolver/previewawsroute53resolvermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// High-level information for a firewall rule group.
//
// A firewall rule group is a collection of rules that DNS Firewall uses to filter DNS network traffic for a VPC. To retrieve the rules for the rule group, call [ListFirewallRules](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ListFirewallRules.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFirewallRuleGroupPropsMixin := awscdkmixinspreview.Mixins.NewCfnFirewallRuleGroupPropsMixin(&CfnFirewallRuleGroupMixinProps{
//   	FirewallRules: []interface{}{
//   		&FirewallRuleProperty{
//   			Action: jsii.String("action"),
//   			BlockOverrideDnsType: jsii.String("blockOverrideDnsType"),
//   			BlockOverrideDomain: jsii.String("blockOverrideDomain"),
//   			BlockOverrideTtl: jsii.Number(123),
//   			BlockResponse: jsii.String("blockResponse"),
//   			ConfidenceThreshold: jsii.String("confidenceThreshold"),
//   			DnsThreatProtection: jsii.String("dnsThreatProtection"),
//   			FirewallDomainListId: jsii.String("firewallDomainListId"),
//   			FirewallDomainRedirectionAction: jsii.String("firewallDomainRedirectionAction"),
//   			FirewallThreatProtectionId: jsii.String("firewallThreatProtectionId"),
//   			Priority: jsii.Number(123),
//   			Qtype: jsii.String("qtype"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-firewallrulegroup.html
//
type CfnFirewallRuleGroupPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnFirewallRuleGroupMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFirewallRuleGroupPropsMixin
type jsiiProxy_CfnFirewallRuleGroupPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnFirewallRuleGroupPropsMixin) Props() *CfnFirewallRuleGroupMixinProps {
	var returns *CfnFirewallRuleGroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroupPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Route53Resolver::FirewallRuleGroup`.
func NewCfnFirewallRuleGroupPropsMixin(props *CfnFirewallRuleGroupMixinProps, options *mixins.CfnPropertyMixinOptions) CfnFirewallRuleGroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFirewallRuleGroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFirewallRuleGroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53resolver.mixins.CfnFirewallRuleGroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Route53Resolver::FirewallRuleGroup`.
func NewCfnFirewallRuleGroupPropsMixin_Override(c CfnFirewallRuleGroupPropsMixin, props *CfnFirewallRuleGroupMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53resolver.mixins.CfnFirewallRuleGroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnFirewallRuleGroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFirewallRuleGroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_route53resolver.mixins.CfnFirewallRuleGroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFirewallRuleGroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_route53resolver.mixins.CfnFirewallRuleGroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFirewallRuleGroupPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnFirewallRuleGroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

