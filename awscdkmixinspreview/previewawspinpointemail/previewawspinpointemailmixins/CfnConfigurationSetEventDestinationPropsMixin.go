package previewawspinpointemailmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawspinpointemail/previewawspinpointemailmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Create an event destination.
//
// In Amazon Pinpoint, *events* include message sends, deliveries, opens, clicks, bounces, and complaints. *Event destinations* are places that you can send information about these events to. For example, you can send event data to Amazon SNS to receive notifications when you receive bounces or complaints, or you can use Amazon Kinesis Data Firehose to stream data to Amazon S3 for long-term storage.
//
// A single configuration set can include more than one event destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConfigurationSetEventDestinationPropsMixin := awscdkmixinspreview.Mixins.NewCfnConfigurationSetEventDestinationPropsMixin(&CfnConfigurationSetEventDestinationMixinProps{
//   	ConfigurationSetName: jsii.String("configurationSetName"),
//   	EventDestination: &EventDestinationProperty{
//   		CloudWatchDestination: &CloudWatchDestinationProperty{
//   			DimensionConfigurations: []interface{}{
//   				&DimensionConfigurationProperty{
//   					DefaultDimensionValue: jsii.String("defaultDimensionValue"),
//   					DimensionName: jsii.String("dimensionName"),
//   					DimensionValueSource: jsii.String("dimensionValueSource"),
//   				},
//   			},
//   		},
//   		Enabled: jsii.Boolean(false),
//   		KinesisFirehoseDestination: &KinesisFirehoseDestinationProperty{
//   			DeliveryStreamArn: jsii.String("deliveryStreamArn"),
//   			IamRoleArn: jsii.String("iamRoleArn"),
//   		},
//   		MatchingEventTypes: []*string{
//   			jsii.String("matchingEventTypes"),
//   		},
//   		PinpointDestination: &PinpointDestinationProperty{
//   			ApplicationArn: jsii.String("applicationArn"),
//   		},
//   		SnsDestination: &SnsDestinationProperty{
//   			TopicArn: jsii.String("topicArn"),
//   		},
//   	},
//   	EventDestinationName: jsii.String("eventDestinationName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-configurationseteventdestination.html
//
type CfnConfigurationSetEventDestinationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnConfigurationSetEventDestinationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnConfigurationSetEventDestinationPropsMixin
type jsiiProxy_CfnConfigurationSetEventDestinationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnConfigurationSetEventDestinationPropsMixin) Props() *CfnConfigurationSetEventDestinationMixinProps {
	var returns *CfnConfigurationSetEventDestinationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSetEventDestinationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::PinpointEmail::ConfigurationSetEventDestination`.
func NewCfnConfigurationSetEventDestinationPropsMixin(props *CfnConfigurationSetEventDestinationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnConfigurationSetEventDestinationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnConfigurationSetEventDestinationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnConfigurationSetEventDestinationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pinpointemail.mixins.CfnConfigurationSetEventDestinationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::PinpointEmail::ConfigurationSetEventDestination`.
func NewCfnConfigurationSetEventDestinationPropsMixin_Override(c CfnConfigurationSetEventDestinationPropsMixin, props *CfnConfigurationSetEventDestinationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pinpointemail.mixins.CfnConfigurationSetEventDestinationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnConfigurationSetEventDestinationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConfigurationSetEventDestinationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_pinpointemail.mixins.CfnConfigurationSetEventDestinationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConfigurationSetEventDestinationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_pinpointemail.mixins.CfnConfigurationSetEventDestinationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConfigurationSetEventDestinationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnConfigurationSetEventDestinationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

