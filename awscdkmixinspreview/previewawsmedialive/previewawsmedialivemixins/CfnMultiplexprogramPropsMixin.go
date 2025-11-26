package previewawsmedialivemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsmedialive/previewawsmedialivemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource schema for AWS::MediaLive::Multiplexprogram.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMultiplexprogramPropsMixin := awscdkmixinspreview.Mixins.NewCfnMultiplexprogramPropsMixin(&CfnMultiplexprogramMixinProps{
//   	MultiplexId: jsii.String("multiplexId"),
//   	MultiplexProgramSettings: &MultiplexProgramSettingsProperty{
//   		PreferredChannelPipeline: jsii.String("preferredChannelPipeline"),
//   		ProgramNumber: jsii.Number(123),
//   		ServiceDescriptor: &MultiplexProgramServiceDescriptorProperty{
//   			ProviderName: jsii.String("providerName"),
//   			ServiceName: jsii.String("serviceName"),
//   		},
//   		VideoSettings: &MultiplexVideoSettingsProperty{
//   			ConstantBitrate: jsii.Number(123),
//   			StatmuxSettings: &MultiplexStatmuxVideoSettingsProperty{
//   				MaximumBitrate: jsii.Number(123),
//   				MinimumBitrate: jsii.Number(123),
//   				Priority: jsii.Number(123),
//   			},
//   		},
//   	},
//   	PacketIdentifiersMap: &MultiplexProgramPacketIdentifiersMapProperty{
//   		AudioPids: []interface{}{
//   			jsii.Number(123),
//   		},
//   		DvbSubPids: []interface{}{
//   			jsii.Number(123),
//   		},
//   		DvbTeletextPid: jsii.Number(123),
//   		EtvPlatformPid: jsii.Number(123),
//   		EtvSignalPid: jsii.Number(123),
//   		KlvDataPids: []interface{}{
//   			jsii.Number(123),
//   		},
//   		PcrPid: jsii.Number(123),
//   		PmtPid: jsii.Number(123),
//   		PrivateMetadataPid: jsii.Number(123),
//   		Scte27Pids: []interface{}{
//   			jsii.Number(123),
//   		},
//   		Scte35Pid: jsii.Number(123),
//   		TimedMetadataPid: jsii.Number(123),
//   		VideoPid: jsii.Number(123),
//   	},
//   	PipelineDetails: []interface{}{
//   		&MultiplexProgramPipelineDetailProperty{
//   			ActiveChannelPipeline: jsii.String("activeChannelPipeline"),
//   			PipelineId: jsii.String("pipelineId"),
//   		},
//   	},
//   	PreferredChannelPipeline: jsii.String("preferredChannelPipeline"),
//   	ProgramName: jsii.String("programName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-multiplexprogram.html
//
type CfnMultiplexprogramPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnMultiplexprogramMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMultiplexprogramPropsMixin
type jsiiProxy_CfnMultiplexprogramPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnMultiplexprogramPropsMixin) Props() *CfnMultiplexprogramMixinProps {
	var returns *CfnMultiplexprogramMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiplexprogramPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaLive::Multiplexprogram`.
func NewCfnMultiplexprogramPropsMixin(props *CfnMultiplexprogramMixinProps, options *mixins.CfnPropertyMixinOptions) CfnMultiplexprogramPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMultiplexprogramPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMultiplexprogramPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_medialive.mixins.CfnMultiplexprogramPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaLive::Multiplexprogram`.
func NewCfnMultiplexprogramPropsMixin_Override(c CfnMultiplexprogramPropsMixin, props *CfnMultiplexprogramMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_medialive.mixins.CfnMultiplexprogramPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnMultiplexprogramPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMultiplexprogramPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_medialive.mixins.CfnMultiplexprogramPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMultiplexprogramPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_medialive.mixins.CfnMultiplexprogramPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMultiplexprogramPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnMultiplexprogramPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

