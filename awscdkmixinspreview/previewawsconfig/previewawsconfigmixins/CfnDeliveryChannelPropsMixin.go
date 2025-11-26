package previewawsconfigmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsconfig/previewawsconfigmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a delivery channel object to deliver configuration information to an Amazon S3 bucket and Amazon  topic.
//
// Before you can create a delivery channel, you must create a configuration recorder. You can use this action to change the Amazon S3 bucket or an Amazon  topic of the existing delivery channel. To change the Amazon S3 bucket or an Amazon  topic, call this action and specify the changed values for the S3 bucket and the SNS topic. If you specify a different value for either the S3 bucket or the SNS topic, this action will keep the existing value for the parameter that is not changed.
//
// > In the China (Beijing) Region, when you call this action, the Amazon S3 bucket must also be in the China (Beijing) Region. In all the other regions, AWS Config supports cross-region and cross-account delivery channels.
//
// You can have only one delivery channel per region per AWS account, and the delivery channel is required to use AWS Config .
//
// > AWS Config does not support the delivery channel to an Amazon S3 bucket bucket where object lock is enabled. For more information, see [How S3 Object Lock works](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock-overview.html) .
//
// When you create the delivery channel, you can specify; how often AWS Config delivers configuration snapshots to your Amazon S3 bucket (for example, 24 hours), the S3 bucket to which AWS Config sends configuration snapshots and configuration history files, and the Amazon  topic to which AWS Config sends notifications about configuration changes, such as updated resources, AWS Config rule evaluations, and when AWS Config delivers the configuration snapshot to your S3 bucket. For more information, see [Deliver Configuration Items](https://docs.aws.amazon.com/config/latest/developerguide/how-does-config-work.html#delivery-channel) in the AWS Config Developer Guide.
//
// > To enable AWS Config , you must create a configuration recorder and a delivery channel. If you want to create the resources separately, you must create a configuration recorder before you can create a delivery channel. AWS Config uses the configuration recorder to capture configuration changes to your resources. For more information, see [AWS::Config::ConfigurationRecorder](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationrecorder.html) .
//
// For more information, see [Managing the Delivery Channel](https://docs.aws.amazon.com/config/latest/developerguide/manage-delivery-channel.html) in the AWS Config Developer Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDeliveryChannelPropsMixin := awscdkmixinspreview.Mixins.NewCfnDeliveryChannelPropsMixin(&CfnDeliveryChannelMixinProps{
//   	ConfigSnapshotDeliveryProperties: &ConfigSnapshotDeliveryPropertiesProperty{
//   		DeliveryFrequency: jsii.String("deliveryFrequency"),
//   	},
//   	Name: jsii.String("name"),
//   	S3BucketName: jsii.String("s3BucketName"),
//   	S3KeyPrefix: jsii.String("s3KeyPrefix"),
//   	S3KmsKeyArn: jsii.String("s3KmsKeyArn"),
//   	SnsTopicArn: jsii.String("snsTopicArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html
//
type CfnDeliveryChannelPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDeliveryChannelMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDeliveryChannelPropsMixin
type jsiiProxy_CfnDeliveryChannelPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDeliveryChannelPropsMixin) Props() *CfnDeliveryChannelMixinProps {
	var returns *CfnDeliveryChannelMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryChannelPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Config::DeliveryChannel`.
func NewCfnDeliveryChannelPropsMixin(props *CfnDeliveryChannelMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDeliveryChannelPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDeliveryChannelPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDeliveryChannelPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_config.mixins.CfnDeliveryChannelPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Config::DeliveryChannel`.
func NewCfnDeliveryChannelPropsMixin_Override(c CfnDeliveryChannelPropsMixin, props *CfnDeliveryChannelMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_config.mixins.CfnDeliveryChannelPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDeliveryChannelPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDeliveryChannelPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_config.mixins.CfnDeliveryChannelPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDeliveryChannelPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_config.mixins.CfnDeliveryChannelPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDeliveryChannelPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnDeliveryChannelPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

