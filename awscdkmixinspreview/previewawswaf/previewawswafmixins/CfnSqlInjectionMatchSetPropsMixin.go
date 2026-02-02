package previewawswafmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawswaf/previewawswafmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// > AWS WAF Classic support will end on September 30, 2025.
//
// >
// > This is *AWS WAF Classic* documentation. For more information, see [AWS WAF Classic](https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html) in the developer guide.
// >
// > *For the latest version of AWS WAF* , use the AWS WAF V2 API and see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) . With the latest version, AWS WAF has a single set of endpoints for regional and global use.
//
// A complex type that contains `SqlInjectionMatchTuple` objects, which specify the parts of web requests that you want AWS WAF to inspect for snippets of malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header. If a `SqlInjectionMatchSet` contains more than one `SqlInjectionMatchTuple` object, a request needs to include snippets of SQL code in only one of the specified parts of the request to be considered a match.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSqlInjectionMatchSetPropsMixin := awscdkmixinspreview.Mixins.NewCfnSqlInjectionMatchSetPropsMixin(&CfnSqlInjectionMatchSetMixinProps{
//   	Name: jsii.String("name"),
//   	SqlInjectionMatchTuples: []interface{}{
//   		&SqlInjectionMatchTupleProperty{
//   			FieldToMatch: &FieldToMatchProperty{
//   				Data: jsii.String("data"),
//   				Type: jsii.String("type"),
//   			},
//   			TextTransformation: jsii.String("textTransformation"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sqlinjectionmatchset.html
//
type CfnSqlInjectionMatchSetPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSqlInjectionMatchSetMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSqlInjectionMatchSetPropsMixin
type jsiiProxy_CfnSqlInjectionMatchSetPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSqlInjectionMatchSetPropsMixin) Props() *CfnSqlInjectionMatchSetMixinProps {
	var returns *CfnSqlInjectionMatchSetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSqlInjectionMatchSetPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::WAF::SqlInjectionMatchSet`.
func NewCfnSqlInjectionMatchSetPropsMixin(props *CfnSqlInjectionMatchSetMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSqlInjectionMatchSetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSqlInjectionMatchSetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSqlInjectionMatchSetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_waf.mixins.CfnSqlInjectionMatchSetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::WAF::SqlInjectionMatchSet`.
func NewCfnSqlInjectionMatchSetPropsMixin_Override(c CfnSqlInjectionMatchSetPropsMixin, props *CfnSqlInjectionMatchSetMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_waf.mixins.CfnSqlInjectionMatchSetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSqlInjectionMatchSetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSqlInjectionMatchSetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_waf.mixins.CfnSqlInjectionMatchSetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSqlInjectionMatchSetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_waf.mixins.CfnSqlInjectionMatchSetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSqlInjectionMatchSetPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnSqlInjectionMatchSetPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

