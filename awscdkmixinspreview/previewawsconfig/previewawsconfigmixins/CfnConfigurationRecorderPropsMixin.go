package previewawsconfigmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsconfig/previewawsconfigmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Config::ConfigurationRecorder` resource type describes the AWS resource types that AWS Config records for configuration changes.
//
// The configuration recorder stores the configuration changes of the specified resources in your account as configuration items.
//
// > To enable AWS Config , you must create a configuration recorder and a delivery channel.
// >
// > AWS Config uses the delivery channel to deliver the configuration changes to your Amazon S3 bucket or Amazon  topic. For more information, see [AWS::Config::DeliveryChannel](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html) .
//
// AWS CloudFormation starts the recorder as soon as the delivery channel is available.
//
// To stop the recorder and delete it, delete the configuration recorder from your stack. To stop the recorder without deleting it, call the [StopConfigurationRecorder](https://docs.aws.amazon.com/config/latest/APIReference/API_StopConfigurationRecorder.html) action of the AWS Config API directly.
//
// For more information, see [Configuration Recorder](https://docs.aws.amazon.com/config/latest/developerguide/config-concepts.html#config-recorder) in the AWS Config Developer Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConfigurationRecorderPropsMixin := awscdkmixinspreview.Mixins.NewCfnConfigurationRecorderPropsMixin(&CfnConfigurationRecorderMixinProps{
//   	Name: jsii.String("name"),
//   	RecordingGroup: &RecordingGroupProperty{
//   		AllSupported: jsii.Boolean(false),
//   		ExclusionByResourceTypes: &ExclusionByResourceTypesProperty{
//   			ResourceTypes: []*string{
//   				jsii.String("resourceTypes"),
//   			},
//   		},
//   		IncludeGlobalResourceTypes: jsii.Boolean(false),
//   		RecordingStrategy: &RecordingStrategyProperty{
//   			UseOnly: jsii.String("useOnly"),
//   		},
//   		ResourceTypes: []*string{
//   			jsii.String("resourceTypes"),
//   		},
//   	},
//   	RecordingMode: &RecordingModeProperty{
//   		RecordingFrequency: jsii.String("recordingFrequency"),
//   		RecordingModeOverrides: []interface{}{
//   			&RecordingModeOverrideProperty{
//   				Description: jsii.String("description"),
//   				RecordingFrequency: jsii.String("recordingFrequency"),
//   				ResourceTypes: []*string{
//   					jsii.String("resourceTypes"),
//   				},
//   			},
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationrecorder.html
//
type CfnConfigurationRecorderPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnConfigurationRecorderMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnConfigurationRecorderPropsMixin
type jsiiProxy_CfnConfigurationRecorderPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnConfigurationRecorderPropsMixin) Props() *CfnConfigurationRecorderMixinProps {
	var returns *CfnConfigurationRecorderMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationRecorderPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Config::ConfigurationRecorder`.
func NewCfnConfigurationRecorderPropsMixin(props *CfnConfigurationRecorderMixinProps, options *mixins.CfnPropertyMixinOptions) CfnConfigurationRecorderPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnConfigurationRecorderPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnConfigurationRecorderPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_config.mixins.CfnConfigurationRecorderPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Config::ConfigurationRecorder`.
func NewCfnConfigurationRecorderPropsMixin_Override(c CfnConfigurationRecorderPropsMixin, props *CfnConfigurationRecorderMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_config.mixins.CfnConfigurationRecorderPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnConfigurationRecorderPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConfigurationRecorderPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_config.mixins.CfnConfigurationRecorderPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConfigurationRecorderPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_config.mixins.CfnConfigurationRecorderPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConfigurationRecorderPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnConfigurationRecorderPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

