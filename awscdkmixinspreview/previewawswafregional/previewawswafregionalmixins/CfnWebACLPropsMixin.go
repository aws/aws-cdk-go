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
// Contains the `Rules` that identify the requests that you want to allow, block, or count. In a `WebACL` , you also specify a default action ( `ALLOW` or `BLOCK` ), and the action for each `Rule` that you add to a `WebACL` , for example, block requests from specified IP addresses or block requests from specified referrers. If you add more than one `Rule` to a `WebACL` , a request needs to match only one of the specifications to be allowed, blocked, or counted.
//
// To identify the requests that you want AWS WAF to filter, you associate the `WebACL` with an API Gateway API or an Application Load Balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWebACLPropsMixin := awscdkmixinspreview.Mixins.NewCfnWebACLPropsMixin(&CfnWebACLMixinProps{
//   	DefaultAction: &ActionProperty{
//   		Type: jsii.String("type"),
//   	},
//   	MetricName: jsii.String("metricName"),
//   	Name: jsii.String("name"),
//   	Rules: []interface{}{
//   		&RuleProperty{
//   			Action: &ActionProperty{
//   				Type: jsii.String("type"),
//   			},
//   			Priority: jsii.Number(123),
//   			RuleId: jsii.String("ruleId"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webacl.html
//
type CfnWebACLPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnWebACLMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnWebACLPropsMixin
type jsiiProxy_CfnWebACLPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnWebACLPropsMixin) Props() *CfnWebACLMixinProps {
	var returns *CfnWebACLMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::WAFRegional::WebACL`.
func NewCfnWebACLPropsMixin(props *CfnWebACLMixinProps, options *mixins.CfnPropertyMixinOptions) CfnWebACLPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnWebACLPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnWebACLPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_wafregional.mixins.CfnWebACLPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::WAFRegional::WebACL`.
func NewCfnWebACLPropsMixin_Override(c CfnWebACLPropsMixin, props *CfnWebACLMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_wafregional.mixins.CfnWebACLPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnWebACLPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWebACLPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_wafregional.mixins.CfnWebACLPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWebACLPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_wafregional.mixins.CfnWebACLPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWebACLPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnWebACLPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

