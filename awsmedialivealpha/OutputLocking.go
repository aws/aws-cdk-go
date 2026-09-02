package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output locking synchronises the frames emitted by a channel's two pipelines. Use the static factory methods to select a strategy:.
//
// - `OutputLocking.pipeline()` — synchronise each pipeline's output to the other.
// - `OutputLocking.epoch()` — synchronise each pipeline's output to the Unix epoch.
// - `OutputLocking.disabled()` — do not synchronise pipelines.
//
// Example:
//   var stack Stack
//   var input IInput
//   var bucket IBucket
//   var video EncodeConfiguration
//   var audio EncodeConfiguration
//
//
//   medialive.NewChannel(stack, jsii.String("Channel"), &ChannelProps{
//   	Inputs: []InputAttachment{
//   		&InputAttachment{
//   			Input: *Input,
//   		},
//   	},
//   	TimecodeConfig: &TimecodeConfig{
//   		Source: medialive.TimecodeSource_EMBEDDED(),
//   	},
//   	GlobalConfiguration: &GlobalConfiguration{
//   		OutputLocking: medialive.OutputLocking_Epoch(),
//   		OutputTimingSource: medialive.OutputTimingSource_INPUT_CLOCK(),
//   	},
//   	OutputGroups: []OutputGroupConfiguration{
//   		medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
//   			Name: jsii.String("hls"),
//   			Destinations: []OutputDestination{
//   				medialive.OutputDestination_ToBucket(bucket, jsii.String("live/stream")),
//   			},
//   			Outputs: []HlsOutputDefinition{
//   				&HlsOutputDefinition{
//   					Encodes: []EncodeConfiguration{
//   						video,
//   						audio,
//   					},
//   					OutputName: jsii.String("hls_out"),
//   				},
//   			},
//   		}),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/medialive/latest/ug/pipeline-lock.html
//
// Experimental.
type OutputLocking interface {
}

// The jsii proxy struct for OutputLocking
type jsiiProxy_OutputLocking struct {
	_ byte // padding
}

// Experimental.
func NewOutputLocking_Override(o OutputLocking) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-medialive-alpha.OutputLocking",
		nil, // no parameters
		o,
	)
}

// Disable output locking (optionally with a custom epoch).
// Experimental.
func OutputLocking_Disabled(customEpoch *string) OutputLocking {
	_init_.Initialize()

	var returns OutputLocking

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.OutputLocking",
		"disabled",
		[]interface{}{customEpoch},
		&returns,
	)

	return returns
}

// Lock outputs to an epoch.
// Experimental.
func OutputLocking_Epoch(props *EpochOutputLockingProps) OutputLocking {
	_init_.Initialize()

	if err := validateOutputLocking_EpochParameters(props); err != nil {
		panic(err)
	}
	var returns OutputLocking

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.OutputLocking",
		"epoch",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Lock pipelines to each other.
// Experimental.
func OutputLocking_Pipeline(props *PipelineOutputLockingProps) OutputLocking {
	_init_.Initialize()

	if err := validateOutputLocking_PipelineParameters(props); err != nil {
		panic(err)
	}
	var returns OutputLocking

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.OutputLocking",
		"pipeline",
		[]interface{}{props},
		&returns,
	)

	return returns
}

