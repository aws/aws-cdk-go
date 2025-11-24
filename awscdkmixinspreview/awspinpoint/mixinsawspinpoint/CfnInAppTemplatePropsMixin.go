package mixinsawspinpoint

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awspinpoint/mixinsawspinpoint/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a message template that you can use to send in-app messages.
//
// A message template is a set of content and settings that you can define, save, and reuse in messages for any of your Amazon Pinpoint applications. The In-App channel is unavailable in AWS GovCloud (US).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var customConfig interface{}
//   var tags interface{}
//
//   cfnInAppTemplatePropsMixin := awscdkmixinspreview.Mixins.NewCfnInAppTemplatePropsMixin(&CfnInAppTemplateMixinProps{
//   	Content: []interface{}{
//   		&InAppMessageContentProperty{
//   			BackgroundColor: jsii.String("backgroundColor"),
//   			BodyConfig: &BodyConfigProperty{
//   				Alignment: jsii.String("alignment"),
//   				Body: jsii.String("body"),
//   				TextColor: jsii.String("textColor"),
//   			},
//   			HeaderConfig: &HeaderConfigProperty{
//   				Alignment: jsii.String("alignment"),
//   				Header: jsii.String("header"),
//   				TextColor: jsii.String("textColor"),
//   			},
//   			ImageUrl: jsii.String("imageUrl"),
//   			PrimaryBtn: &ButtonConfigProperty{
//   				Android: &OverrideButtonConfigurationProperty{
//   					ButtonAction: jsii.String("buttonAction"),
//   					Link: jsii.String("link"),
//   				},
//   				DefaultConfig: &DefaultButtonConfigurationProperty{
//   					BackgroundColor: jsii.String("backgroundColor"),
//   					BorderRadius: jsii.Number(123),
//   					ButtonAction: jsii.String("buttonAction"),
//   					Link: jsii.String("link"),
//   					Text: jsii.String("text"),
//   					TextColor: jsii.String("textColor"),
//   				},
//   				Ios: &OverrideButtonConfigurationProperty{
//   					ButtonAction: jsii.String("buttonAction"),
//   					Link: jsii.String("link"),
//   				},
//   				Web: &OverrideButtonConfigurationProperty{
//   					ButtonAction: jsii.String("buttonAction"),
//   					Link: jsii.String("link"),
//   				},
//   			},
//   			SecondaryBtn: &ButtonConfigProperty{
//   				Android: &OverrideButtonConfigurationProperty{
//   					ButtonAction: jsii.String("buttonAction"),
//   					Link: jsii.String("link"),
//   				},
//   				DefaultConfig: &DefaultButtonConfigurationProperty{
//   					BackgroundColor: jsii.String("backgroundColor"),
//   					BorderRadius: jsii.Number(123),
//   					ButtonAction: jsii.String("buttonAction"),
//   					Link: jsii.String("link"),
//   					Text: jsii.String("text"),
//   					TextColor: jsii.String("textColor"),
//   				},
//   				Ios: &OverrideButtonConfigurationProperty{
//   					ButtonAction: jsii.String("buttonAction"),
//   					Link: jsii.String("link"),
//   				},
//   				Web: &OverrideButtonConfigurationProperty{
//   					ButtonAction: jsii.String("buttonAction"),
//   					Link: jsii.String("link"),
//   				},
//   			},
//   		},
//   	},
//   	CustomConfig: customConfig,
//   	Layout: jsii.String("layout"),
//   	Tags: tags,
//   	TemplateDescription: jsii.String("templateDescription"),
//   	TemplateName: jsii.String("templateName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-inapptemplate.html
//
type CfnInAppTemplatePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnInAppTemplateMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnInAppTemplatePropsMixin
type jsiiProxy_CfnInAppTemplatePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnInAppTemplatePropsMixin) Props() *CfnInAppTemplateMixinProps {
	var returns *CfnInAppTemplateMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplatePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Pinpoint::InAppTemplate`.
func NewCfnInAppTemplatePropsMixin(props *CfnInAppTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) CfnInAppTemplatePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnInAppTemplatePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnInAppTemplatePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pinpoint.mixins.CfnInAppTemplatePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Pinpoint::InAppTemplate`.
func NewCfnInAppTemplatePropsMixin_Override(c CfnInAppTemplatePropsMixin, props *CfnInAppTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pinpoint.mixins.CfnInAppTemplatePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnInAppTemplatePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInAppTemplatePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_pinpoint.mixins.CfnInAppTemplatePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInAppTemplatePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_pinpoint.mixins.CfnInAppTemplatePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInAppTemplatePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnInAppTemplatePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

