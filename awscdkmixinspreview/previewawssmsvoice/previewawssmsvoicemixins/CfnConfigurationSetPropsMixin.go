package previewawssmsvoicemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssmsvoice/previewawssmsvoicemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new configuration set.
//
// After you create the configuration set, you can add one or more event destinations to it.
//
// A configuration set is a set of rules that you apply to the SMS and voice messages that you send.
//
// When you send a message, you can optionally specify a single configuration set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConfigurationSetPropsMixin := awscdkmixinspreview.Mixins.NewCfnConfigurationSetPropsMixin(&CfnConfigurationSetMixinProps{
//   	ConfigurationSetName: jsii.String("configurationSetName"),
//   	DefaultSenderId: jsii.String("defaultSenderId"),
//   	EventDestinations: []interface{}{
//   		&EventDestinationProperty{
//   			CloudWatchLogsDestination: &CloudWatchLogsDestinationProperty{
//   				IamRoleArn: jsii.String("iamRoleArn"),
//   				LogGroupArn: jsii.String("logGroupArn"),
//   			},
//   			Enabled: jsii.Boolean(false),
//   			EventDestinationName: jsii.String("eventDestinationName"),
//   			KinesisFirehoseDestination: &KinesisFirehoseDestinationProperty{
//   				DeliveryStreamArn: jsii.String("deliveryStreamArn"),
//   				IamRoleArn: jsii.String("iamRoleArn"),
//   			},
//   			MatchingEventTypes: []*string{
//   				jsii.String("matchingEventTypes"),
//   			},
//   			SnsDestination: &SnsDestinationProperty{
//   				TopicArn: jsii.String("topicArn"),
//   			},
//   		},
//   	},
//   	MessageFeedbackEnabled: jsii.Boolean(false),
//   	ProtectConfigurationId: jsii.String("protectConfigurationId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-configurationset.html
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


// Create a mixin to apply properties to `AWS::SMSVOICE::ConfigurationSet`.
func NewCfnConfigurationSetPropsMixin(props *CfnConfigurationSetMixinProps, options *mixins.CfnPropertyMixinOptions) CfnConfigurationSetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnConfigurationSetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnConfigurationSetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_smsvoice.mixins.CfnConfigurationSetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SMSVOICE::ConfigurationSet`.
func NewCfnConfigurationSetPropsMixin_Override(c CfnConfigurationSetPropsMixin, props *CfnConfigurationSetMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_smsvoice.mixins.CfnConfigurationSetPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_smsvoice.mixins.CfnConfigurationSetPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_smsvoice.mixins.CfnConfigurationSetPropsMixin",
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

