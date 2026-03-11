package awsiotanalytics

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsiotanalytics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The AWS::IoTAnalytics::Channel resource collects data from an MQTT topic and archives the raw, unprocessed messages before publishing the data to a pipeline.
//
// For more information, see [How to Use](https://docs.aws.amazon.com/iotanalytics/latest/userguide/welcome.html#aws-iot-analytics-how) in the *User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//   var serviceManagedS3 interface{}
//
//   cfnChannelPropsMixin := awscdkcfnpropertymixins.Aws_iotanalytics.NewCfnChannelPropsMixin(&CfnChannelMixinProps{
//   	ChannelName: jsii.String("channelName"),
//   	ChannelStorage: &ChannelStorageProperty{
//   		CustomerManagedS3: &CustomerManagedS3Property{
//   			Bucket: jsii.String("bucket"),
//   			KeyPrefix: jsii.String("keyPrefix"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   		ServiceManagedS3: serviceManagedS3,
//   	},
//   	RetentionPeriod: &RetentionPeriodProperty{
//   		NumberOfDays: jsii.Number(123),
//   		Unlimited: jsii.Boolean(false),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotanalytics-channel.html
//
type CfnChannelPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnChannelMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnChannelPropsMixin
type jsiiProxy_CfnChannelPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnChannelPropsMixin) Props() *CfnChannelMixinProps {
	var returns *CfnChannelMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannelPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoTAnalytics::Channel`.
func NewCfnChannelPropsMixin(props *CfnChannelMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnChannelPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnChannelPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnChannelPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iotanalytics.CfnChannelPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoTAnalytics::Channel`.
func NewCfnChannelPropsMixin_Override(c CfnChannelPropsMixin, props *CfnChannelMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iotanalytics.CfnChannelPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnChannelPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnChannelPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_iotanalytics.CfnChannelPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnChannelPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_iotanalytics.CfnChannelPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnChannelPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnChannelPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

