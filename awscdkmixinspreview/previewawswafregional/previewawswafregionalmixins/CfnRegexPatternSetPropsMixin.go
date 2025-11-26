package previewawswafregionalmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawswafregional/previewawswafregionalmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `RegexPatternSet` specifies the regular expression (regex) pattern that you want AWS WAF to search for, such as `B[a@]dB[o0]t` .
//
// You can then configure AWS WAF to reject those requests.
//
// Note that you can only create regex pattern sets using a CloudFormation template. To add the regex pattern sets created through CloudFormation to a RegexMatchSet, use the AWS WAF console, API, or command line interface (CLI). For more information, see [UpdateRegexMatchSet](https://docs.aws.amazon.com/waf/latest/APIReference/API_regional_UpdateRegexMatchSet.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRegexPatternSetPropsMixin := awscdkmixinspreview.Mixins.NewCfnRegexPatternSetPropsMixin(&CfnRegexPatternSetMixinProps{
//   	Name: jsii.String("name"),
//   	RegexPatternStrings: []*string{
//   		jsii.String("regexPatternStrings"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-regexpatternset.html
//
type CfnRegexPatternSetPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnRegexPatternSetMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRegexPatternSetPropsMixin
type jsiiProxy_CfnRegexPatternSetPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnRegexPatternSetPropsMixin) Props() *CfnRegexPatternSetMixinProps {
	var returns *CfnRegexPatternSetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSetPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::WAFRegional::RegexPatternSet`.
func NewCfnRegexPatternSetPropsMixin(props *CfnRegexPatternSetMixinProps, options *mixins.CfnPropertyMixinOptions) CfnRegexPatternSetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRegexPatternSetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRegexPatternSetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_wafregional.mixins.CfnRegexPatternSetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::WAFRegional::RegexPatternSet`.
func NewCfnRegexPatternSetPropsMixin_Override(c CfnRegexPatternSetPropsMixin, props *CfnRegexPatternSetMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_wafregional.mixins.CfnRegexPatternSetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnRegexPatternSetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRegexPatternSetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_wafregional.mixins.CfnRegexPatternSetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRegexPatternSetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_wafregional.mixins.CfnRegexPatternSetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRegexPatternSetPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnRegexPatternSetPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

