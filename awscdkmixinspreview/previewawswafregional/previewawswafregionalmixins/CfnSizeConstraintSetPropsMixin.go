package previewawswafregionalmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawswafregional/previewawswafregionalmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// > AWS WAF Classic support will end on September 30, 2025.
//
// >
// > This is *AWS WAF Classic* documentation. For more information, see [AWS WAF Classic](https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html) in the developer guide.
// >
// > *For the latest version of AWS WAF* , use the AWS WAF V2 API and see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) . With the latest version, AWS WAF has a single set of endpoints for regional and global use.
//
// A complex type that contains `SizeConstraint` objects, which specify the parts of web requests that you want AWS WAF to inspect the size of. If a `SizeConstraintSet` contains more than one `SizeConstraint` object, a request only needs to match one constraint to be considered a match.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSizeConstraintSetPropsMixin := awscdkmixinspreview.Mixins.NewCfnSizeConstraintSetPropsMixin(&CfnSizeConstraintSetMixinProps{
//   	Name: jsii.String("name"),
//   	SizeConstraints: []interface{}{
//   		&SizeConstraintProperty{
//   			ComparisonOperator: jsii.String("comparisonOperator"),
//   			FieldToMatch: &FieldToMatchProperty{
//   				Data: jsii.String("data"),
//   				Type: jsii.String("type"),
//   			},
//   			Size: jsii.Number(123),
//   			TextTransformation: jsii.String("textTransformation"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sizeconstraintset.html
//
type CfnSizeConstraintSetPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSizeConstraintSetMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSizeConstraintSetPropsMixin
type jsiiProxy_CfnSizeConstraintSetPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSizeConstraintSetPropsMixin) Props() *CfnSizeConstraintSetMixinProps {
	var returns *CfnSizeConstraintSetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSizeConstraintSetPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::WAFRegional::SizeConstraintSet`.
func NewCfnSizeConstraintSetPropsMixin(props *CfnSizeConstraintSetMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSizeConstraintSetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSizeConstraintSetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSizeConstraintSetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_wafregional.mixins.CfnSizeConstraintSetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::WAFRegional::SizeConstraintSet`.
func NewCfnSizeConstraintSetPropsMixin_Override(c CfnSizeConstraintSetPropsMixin, props *CfnSizeConstraintSetMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_wafregional.mixins.CfnSizeConstraintSetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSizeConstraintSetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSizeConstraintSetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_wafregional.mixins.CfnSizeConstraintSetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSizeConstraintSetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_wafregional.mixins.CfnSizeConstraintSetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSizeConstraintSetPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnSizeConstraintSetPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

