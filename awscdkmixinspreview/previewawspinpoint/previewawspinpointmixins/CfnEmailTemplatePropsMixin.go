package previewawspinpointmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawspinpoint/previewawspinpointmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a message template that you can use in messages that are sent through the email channel.
//
// A *message template* is a set of content and settings that you can define, save, and reuse in messages for any of your Amazon Pinpoint applications.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var tags interface{}
//
//   cfnEmailTemplatePropsMixin := awscdkmixinspreview.Mixins.NewCfnEmailTemplatePropsMixin(&CfnEmailTemplateMixinProps{
//   	DefaultSubstitutions: jsii.String("defaultSubstitutions"),
//   	HtmlPart: jsii.String("htmlPart"),
//   	Subject: jsii.String("subject"),
//   	Tags: tags,
//   	TemplateDescription: jsii.String("templateDescription"),
//   	TemplateName: jsii.String("templateName"),
//   	TextPart: jsii.String("textPart"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-emailtemplate.html
//
type CfnEmailTemplatePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnEmailTemplateMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEmailTemplatePropsMixin
type jsiiProxy_CfnEmailTemplatePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnEmailTemplatePropsMixin) Props() *CfnEmailTemplateMixinProps {
	var returns *CfnEmailTemplateMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailTemplatePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Pinpoint::EmailTemplate`.
func NewCfnEmailTemplatePropsMixin(props *CfnEmailTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) CfnEmailTemplatePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEmailTemplatePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEmailTemplatePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pinpoint.mixins.CfnEmailTemplatePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Pinpoint::EmailTemplate`.
func NewCfnEmailTemplatePropsMixin_Override(c CfnEmailTemplatePropsMixin, props *CfnEmailTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pinpoint.mixins.CfnEmailTemplatePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnEmailTemplatePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEmailTemplatePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_pinpoint.mixins.CfnEmailTemplatePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEmailTemplatePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_pinpoint.mixins.CfnEmailTemplatePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEmailTemplatePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnEmailTemplatePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

