package previewawspinpointemailmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawspinpointemail/previewawspinpointemailmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Create a configuration set.
//
// *Configuration sets* are groups of rules that you can apply to the emails you send using Amazon Pinpoint. You apply a configuration set to an email by including a reference to the configuration set in the headers of the email. When you apply a configuration set to an email, all of the rules in that configuration set are applied to the email.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConfigurationSetPropsMixin := awscdkmixinspreview.Mixins.NewCfnConfigurationSetPropsMixin(&CfnConfigurationSetMixinProps{
//   	DeliveryOptions: &DeliveryOptionsProperty{
//   		SendingPoolName: jsii.String("sendingPoolName"),
//   	},
//   	Name: jsii.String("name"),
//   	ReputationOptions: &ReputationOptionsProperty{
//   		ReputationMetricsEnabled: jsii.Boolean(false),
//   	},
//   	SendingOptions: &SendingOptionsProperty{
//   		SendingEnabled: jsii.Boolean(false),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrackingOptions: &TrackingOptionsProperty{
//   		CustomRedirectDomain: jsii.String("customRedirectDomain"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-configurationset.html
//
type CfnConfigurationSetPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnConfigurationSetMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnConfigurationSetPropsMixin
type jsiiProxy_CfnConfigurationSetPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnConfigurationSetPropsMixin) Props() *CfnConfigurationSetMixinProps {
	var returns *CfnConfigurationSetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSetPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::PinpointEmail::ConfigurationSet`.
func NewCfnConfigurationSetPropsMixin(props *CfnConfigurationSetMixinProps, options *mixins.CfnPropertyMixinOptions) CfnConfigurationSetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnConfigurationSetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnConfigurationSetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pinpointemail.mixins.CfnConfigurationSetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::PinpointEmail::ConfigurationSet`.
func NewCfnConfigurationSetPropsMixin_Override(c CfnConfigurationSetPropsMixin, props *CfnConfigurationSetMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pinpointemail.mixins.CfnConfigurationSetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnConfigurationSetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConfigurationSetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_pinpointemail.mixins.CfnConfigurationSetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConfigurationSetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_pinpointemail.mixins.CfnConfigurationSetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConfigurationSetPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnConfigurationSetPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

