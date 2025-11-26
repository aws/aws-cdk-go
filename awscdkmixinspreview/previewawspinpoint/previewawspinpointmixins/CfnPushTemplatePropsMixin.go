package previewawspinpointmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawspinpoint/previewawspinpointmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a message template that you can use in messages that are sent through a push notification channel.
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
//   cfnPushTemplatePropsMixin := awscdkmixinspreview.Mixins.NewCfnPushTemplatePropsMixin(&CfnPushTemplateMixinProps{
//   	Adm: &AndroidPushNotificationTemplateProperty{
//   		Action: jsii.String("action"),
//   		Body: jsii.String("body"),
//   		ImageIconUrl: jsii.String("imageIconUrl"),
//   		ImageUrl: jsii.String("imageUrl"),
//   		SmallImageIconUrl: jsii.String("smallImageIconUrl"),
//   		Sound: jsii.String("sound"),
//   		Title: jsii.String("title"),
//   		Url: jsii.String("url"),
//   	},
//   	Apns: &APNSPushNotificationTemplateProperty{
//   		Action: jsii.String("action"),
//   		Body: jsii.String("body"),
//   		MediaUrl: jsii.String("mediaUrl"),
//   		Sound: jsii.String("sound"),
//   		Title: jsii.String("title"),
//   		Url: jsii.String("url"),
//   	},
//   	Baidu: &AndroidPushNotificationTemplateProperty{
//   		Action: jsii.String("action"),
//   		Body: jsii.String("body"),
//   		ImageIconUrl: jsii.String("imageIconUrl"),
//   		ImageUrl: jsii.String("imageUrl"),
//   		SmallImageIconUrl: jsii.String("smallImageIconUrl"),
//   		Sound: jsii.String("sound"),
//   		Title: jsii.String("title"),
//   		Url: jsii.String("url"),
//   	},
//   	Default: &DefaultPushNotificationTemplateProperty{
//   		Action: jsii.String("action"),
//   		Body: jsii.String("body"),
//   		Sound: jsii.String("sound"),
//   		Title: jsii.String("title"),
//   		Url: jsii.String("url"),
//   	},
//   	DefaultSubstitutions: jsii.String("defaultSubstitutions"),
//   	Gcm: &AndroidPushNotificationTemplateProperty{
//   		Action: jsii.String("action"),
//   		Body: jsii.String("body"),
//   		ImageIconUrl: jsii.String("imageIconUrl"),
//   		ImageUrl: jsii.String("imageUrl"),
//   		SmallImageIconUrl: jsii.String("smallImageIconUrl"),
//   		Sound: jsii.String("sound"),
//   		Title: jsii.String("title"),
//   		Url: jsii.String("url"),
//   	},
//   	Tags: tags,
//   	TemplateDescription: jsii.String("templateDescription"),
//   	TemplateName: jsii.String("templateName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html
//
type CfnPushTemplatePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPushTemplateMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPushTemplatePropsMixin
type jsiiProxy_CfnPushTemplatePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPushTemplatePropsMixin) Props() *CfnPushTemplateMixinProps {
	var returns *CfnPushTemplateMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPushTemplatePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Pinpoint::PushTemplate`.
func NewCfnPushTemplatePropsMixin(props *CfnPushTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPushTemplatePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPushTemplatePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPushTemplatePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pinpoint.mixins.CfnPushTemplatePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Pinpoint::PushTemplate`.
func NewCfnPushTemplatePropsMixin_Override(c CfnPushTemplatePropsMixin, props *CfnPushTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pinpoint.mixins.CfnPushTemplatePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPushTemplatePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPushTemplatePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_pinpoint.mixins.CfnPushTemplatePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPushTemplatePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_pinpoint.mixins.CfnPushTemplatePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPushTemplatePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnPushTemplatePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

