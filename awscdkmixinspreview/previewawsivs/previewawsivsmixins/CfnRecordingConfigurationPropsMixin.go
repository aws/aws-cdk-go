package previewawsivsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsivs/previewawsivsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::IVS::RecordingConfiguration` resource specifies an  recording configuration.
//
// A recording configuration enables the recording of a channel’s live streams to a data store. Multiple channels can reference the same recording configuration. For more information, see [RecordingConfiguration](https://docs.aws.amazon.com/ivs/latest/LowLatencyAPIReference/API_RecordingConfiguration.html) in the *Amazon IVS Low-Latency Streaming API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRecordingConfigurationPropsMixin := awscdkmixinspreview.Mixins.NewCfnRecordingConfigurationPropsMixin(&CfnRecordingConfigurationMixinProps{
//   	DestinationConfiguration: &DestinationConfigurationProperty{
//   		S3: &S3DestinationConfigurationProperty{
//   			BucketName: jsii.String("bucketName"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	RecordingReconnectWindowSeconds: jsii.Number(123),
//   	RenditionConfiguration: &RenditionConfigurationProperty{
//   		Renditions: []*string{
//   			jsii.String("renditions"),
//   		},
//   		RenditionSelection: jsii.String("renditionSelection"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ThumbnailConfiguration: &ThumbnailConfigurationProperty{
//   		RecordingMode: jsii.String("recordingMode"),
//   		Resolution: jsii.String("resolution"),
//   		Storage: []*string{
//   			jsii.String("storage"),
//   		},
//   		TargetIntervalSeconds: jsii.Number(123),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-recordingconfiguration.html
//
type CfnRecordingConfigurationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnRecordingConfigurationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRecordingConfigurationPropsMixin
type jsiiProxy_CfnRecordingConfigurationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnRecordingConfigurationPropsMixin) Props() *CfnRecordingConfigurationMixinProps {
	var returns *CfnRecordingConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordingConfigurationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IVS::RecordingConfiguration`.
func NewCfnRecordingConfigurationPropsMixin(props *CfnRecordingConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnRecordingConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRecordingConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRecordingConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ivs.mixins.CfnRecordingConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IVS::RecordingConfiguration`.
func NewCfnRecordingConfigurationPropsMixin_Override(c CfnRecordingConfigurationPropsMixin, props *CfnRecordingConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ivs.mixins.CfnRecordingConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnRecordingConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRecordingConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ivs.mixins.CfnRecordingConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRecordingConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ivs.mixins.CfnRecordingConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRecordingConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnRecordingConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

