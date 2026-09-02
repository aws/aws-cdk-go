package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Source of output timing.
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
// Experimental.
type OutputTimingSource interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for OutputTimingSource
type jsiiProxy_OutputTimingSource struct {
	_ byte // padding
}

func (j *jsiiProxy_OutputTimingSource) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// A value not yet modelled by AWS CDK.
// Experimental.
func OutputTimingSource_Of(value *string) OutputTimingSource {
	_init_.Initialize()

	if err := validateOutputTimingSource_OfParameters(value); err != nil {
		panic(err)
	}
	var returns OutputTimingSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.OutputTimingSource",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func OutputTimingSource_INPUT_CLOCK() OutputTimingSource {
	_init_.Initialize()
	var returns OutputTimingSource
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.OutputTimingSource",
		"INPUT_CLOCK",
		&returns,
	)
	return returns
}

func OutputTimingSource_SYSTEM_CLOCK() OutputTimingSource {
	_init_.Initialize()
	var returns OutputTimingSource
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.OutputTimingSource",
		"SYSTEM_CLOCK",
		&returns,
	)
	return returns
}

