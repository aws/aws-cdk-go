package awsglue

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a custom pattern that is used to detect sensitive data across the columns and rows of your structured data.
//
// Each custom pattern you create specifies a regular expression and an optional list of context words. If no context words are passed only a regular expression is checked.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//   var tags interface{}
//
//   cfnCustomEntityTypePropsMixin := awscdkcfnpropertymixins.Aws_glue.NewCfnCustomEntityTypePropsMixin(&CfnCustomEntityTypeMixinProps{
//   	ContextWords: []*string{
//   		jsii.String("contextWords"),
//   	},
//   	Name: jsii.String("name"),
//   	RegexString: jsii.String("regexString"),
//   	Tags: tags,
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-customentitytype.html
//
type CfnCustomEntityTypePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnCustomEntityTypeMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCustomEntityTypePropsMixin
type jsiiProxy_CfnCustomEntityTypePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnCustomEntityTypePropsMixin) Props() *CfnCustomEntityTypeMixinProps {
	var returns *CfnCustomEntityTypeMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomEntityTypePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Glue::CustomEntityType`.
func NewCfnCustomEntityTypePropsMixin(props *CfnCustomEntityTypeMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnCustomEntityTypePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCustomEntityTypePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCustomEntityTypePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCustomEntityTypePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Glue::CustomEntityType`.
func NewCfnCustomEntityTypePropsMixin_Override(c CfnCustomEntityTypePropsMixin, props *CfnCustomEntityTypeMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCustomEntityTypePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnCustomEntityTypePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCustomEntityTypePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCustomEntityTypePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCustomEntityTypePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCustomEntityTypePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCustomEntityTypePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnCustomEntityTypePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

