package previewawsgroundstationmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsgroundstation/previewawsgroundstationmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a `Config` with the specified parameters.
//
// Config objects provide Ground Station with the details necessary in order to schedule and execute satellite contacts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConfigPropsMixin := awscdkmixinspreview.Mixins.NewCfnConfigPropsMixin(&CfnConfigMixinProps{
//   	ConfigData: &ConfigDataProperty{
//   		AntennaDownlinkConfig: &AntennaDownlinkConfigProperty{
//   			SpectrumConfig: &SpectrumConfigProperty{
//   				Bandwidth: &FrequencyBandwidthProperty{
//   					Units: jsii.String("units"),
//   					Value: jsii.Number(123),
//   				},
//   				CenterFrequency: &FrequencyProperty{
//   					Units: jsii.String("units"),
//   					Value: jsii.Number(123),
//   				},
//   				Polarization: jsii.String("polarization"),
//   			},
//   		},
//   		AntennaDownlinkDemodDecodeConfig: &AntennaDownlinkDemodDecodeConfigProperty{
//   			DecodeConfig: &DecodeConfigProperty{
//   				UnvalidatedJson: jsii.String("unvalidatedJson"),
//   			},
//   			DemodulationConfig: &DemodulationConfigProperty{
//   				UnvalidatedJson: jsii.String("unvalidatedJson"),
//   			},
//   			SpectrumConfig: &SpectrumConfigProperty{
//   				Bandwidth: &FrequencyBandwidthProperty{
//   					Units: jsii.String("units"),
//   					Value: jsii.Number(123),
//   				},
//   				CenterFrequency: &FrequencyProperty{
//   					Units: jsii.String("units"),
//   					Value: jsii.Number(123),
//   				},
//   				Polarization: jsii.String("polarization"),
//   			},
//   		},
//   		AntennaUplinkConfig: &AntennaUplinkConfigProperty{
//   			SpectrumConfig: &UplinkSpectrumConfigProperty{
//   				CenterFrequency: &FrequencyProperty{
//   					Units: jsii.String("units"),
//   					Value: jsii.Number(123),
//   				},
//   				Polarization: jsii.String("polarization"),
//   			},
//   			TargetEirp: &EirpProperty{
//   				Units: jsii.String("units"),
//   				Value: jsii.Number(123),
//   			},
//   			TransmitDisabled: jsii.Boolean(false),
//   		},
//   		DataflowEndpointConfig: &DataflowEndpointConfigProperty{
//   			DataflowEndpointName: jsii.String("dataflowEndpointName"),
//   			DataflowEndpointRegion: jsii.String("dataflowEndpointRegion"),
//   		},
//   		S3RecordingConfig: &S3RecordingConfigProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			Prefix: jsii.String("prefix"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   		TelemetrySinkConfig: &TelemetrySinkConfigProperty{
//   			TelemetrySinkData: &TelemetrySinkDataProperty{
//   				KinesisDataStreamData: &KinesisDataStreamDataProperty{
//   					KinesisDataStreamArn: jsii.String("kinesisDataStreamArn"),
//   					KinesisRoleArn: jsii.String("kinesisRoleArn"),
//   				},
//   			},
//   			TelemetrySinkType: jsii.String("telemetrySinkType"),
//   		},
//   		TrackingConfig: &TrackingConfigProperty{
//   			Autotrack: jsii.String("autotrack"),
//   		},
//   		UplinkEchoConfig: &UplinkEchoConfigProperty{
//   			AntennaUplinkConfigArn: jsii.String("antennaUplinkConfigArn"),
//   			Enabled: jsii.Boolean(false),
//   		},
//   	},
//   	Name: jsii.String("name"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-config.html
//
type CfnConfigPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnConfigMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnConfigPropsMixin
type jsiiProxy_CfnConfigPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnConfigPropsMixin) Props() *CfnConfigMixinProps {
	var returns *CfnConfigMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::GroundStation::Config`.
func NewCfnConfigPropsMixin(props *CfnConfigMixinProps, options *mixins.CfnPropertyMixinOptions) CfnConfigPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnConfigPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnConfigPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnConfigPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::GroundStation::Config`.
func NewCfnConfigPropsMixin_Override(c CfnConfigPropsMixin, props *CfnConfigMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnConfigPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnConfigPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConfigPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnConfigPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConfigPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnConfigPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConfigPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnConfigPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

