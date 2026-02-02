package previewawsgluemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsglue/previewawsgluemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Glue::Classifier` resource creates an AWS Glue classifier that categorizes data sources and specifies schemas.
//
// For more information, see [Adding Classifiers to a Crawler](https://docs.aws.amazon.com/glue/latest/dg/add-classifier.html) and [Classifier Structure](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-crawler-classifiers.html#aws-glue-api-crawler-classifiers-Classifier) in the *AWS Glue Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClassifierPropsMixin := awscdkmixinspreview.Mixins.NewCfnClassifierPropsMixin(&CfnClassifierMixinProps{
//   	CsvClassifier: &CsvClassifierProperty{
//   		AllowSingleColumn: jsii.Boolean(false),
//   		ContainsCustomDatatype: []*string{
//   			jsii.String("containsCustomDatatype"),
//   		},
//   		ContainsHeader: jsii.String("containsHeader"),
//   		CustomDatatypeConfigured: jsii.Boolean(false),
//   		Delimiter: jsii.String("delimiter"),
//   		DisableValueTrimming: jsii.Boolean(false),
//   		Header: []*string{
//   			jsii.String("header"),
//   		},
//   		Name: jsii.String("name"),
//   		QuoteSymbol: jsii.String("quoteSymbol"),
//   	},
//   	GrokClassifier: &GrokClassifierProperty{
//   		Classification: jsii.String("classification"),
//   		CustomPatterns: jsii.String("customPatterns"),
//   		GrokPattern: jsii.String("grokPattern"),
//   		Name: jsii.String("name"),
//   	},
//   	JsonClassifier: &JsonClassifierProperty{
//   		JsonPath: jsii.String("jsonPath"),
//   		Name: jsii.String("name"),
//   	},
//   	XmlClassifier: &XMLClassifierProperty{
//   		Classification: jsii.String("classification"),
//   		Name: jsii.String("name"),
//   		RowTag: jsii.String("rowTag"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-classifier.html
//
type CfnClassifierPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnClassifierMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnClassifierPropsMixin
type jsiiProxy_CfnClassifierPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnClassifierPropsMixin) Props() *CfnClassifierMixinProps {
	var returns *CfnClassifierMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClassifierPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Glue::Classifier`.
func NewCfnClassifierPropsMixin(props *CfnClassifierMixinProps, options *mixins.CfnPropertyMixinOptions) CfnClassifierPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnClassifierPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnClassifierPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnClassifierPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Glue::Classifier`.
func NewCfnClassifierPropsMixin_Override(c CfnClassifierPropsMixin, props *CfnClassifierMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnClassifierPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnClassifierPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnClassifierPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnClassifierPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnClassifierPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnClassifierPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnClassifierPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnClassifierPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

