package previewawswafregionalmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawswafregional/previewawswafregionalmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// > This is *AWS WAF Classic* documentation.
//
// For more information, see [AWS WAF Classic](https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html) in the developer guide.
// >
// > *For the latest version of AWS WAF* , use the AWS WAF V2 API and see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) . With the latest version, AWS WAF has a single set of endpoints for regional and global use.
//
// A `RateBasedRule` is identical to a regular `Rule` , with one addition: a `RateBasedRule` counts the number of requests that arrive from a specified IP address every five minutes. For example, based on recent requests that you've seen from an attacker, you might create a `RateBasedRule` that includes the following conditions:
//
// - The requests come from 192.0.2.44.
// - They contain the value `BadBot` in the `User-Agent` header.
//
// In the rule, you also define the rate limit as 15,000.
//
// Requests that meet both of these conditions and exceed 15,000 requests every five minutes trigger the rule's action (block or count), which is defined in the web ACL.
//
// Note you can only create rate-based rules using an CloudFormation template. To add the rate-based rules created through CloudFormation to a web ACL, use the AWS WAF console, API, or command line interface (CLI). For more information, see [UpdateWebACL](https://docs.aws.amazon.com/waf/latest/APIReference/API_regional_UpdateWebACL.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRateBasedRulePropsMixin := awscdkmixinspreview.Mixins.NewCfnRateBasedRulePropsMixin(&CfnRateBasedRuleMixinProps{
//   	MatchPredicates: []interface{}{
//   		&PredicateProperty{
//   			DataId: jsii.String("dataId"),
//   			Negated: jsii.Boolean(false),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	MetricName: jsii.String("metricName"),
//   	Name: jsii.String("name"),
//   	RateKey: jsii.String("rateKey"),
//   	RateLimit: jsii.Number(123),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ratebasedrule.html
//
type CfnRateBasedRulePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnRateBasedRuleMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRateBasedRulePropsMixin
type jsiiProxy_CfnRateBasedRulePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnRateBasedRulePropsMixin) Props() *CfnRateBasedRuleMixinProps {
	var returns *CfnRateBasedRuleMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRateBasedRulePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::WAFRegional::RateBasedRule`.
func NewCfnRateBasedRulePropsMixin(props *CfnRateBasedRuleMixinProps, options *mixins.CfnPropertyMixinOptions) CfnRateBasedRulePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRateBasedRulePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRateBasedRulePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_wafregional.mixins.CfnRateBasedRulePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::WAFRegional::RateBasedRule`.
func NewCfnRateBasedRulePropsMixin_Override(c CfnRateBasedRulePropsMixin, props *CfnRateBasedRuleMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_wafregional.mixins.CfnRateBasedRulePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnRateBasedRulePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRateBasedRulePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_wafregional.mixins.CfnRateBasedRulePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRateBasedRulePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_wafregional.mixins.CfnRateBasedRulePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRateBasedRulePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnRateBasedRulePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

