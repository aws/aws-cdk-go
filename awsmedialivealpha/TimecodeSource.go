package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The source of timecode for the channel outputs.
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
type TimecodeSource interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for TimecodeSource
type jsiiProxy_TimecodeSource struct {
	_ byte // padding
}

func (j *jsiiProxy_TimecodeSource) Value() *string {
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
func TimecodeSource_Of(value *string) TimecodeSource {
	_init_.Initialize()

	if err := validateTimecodeSource_OfParameters(value); err != nil {
		panic(err)
	}
	var returns TimecodeSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.TimecodeSource",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func TimecodeSource_EMBEDDED() TimecodeSource {
	_init_.Initialize()
	var returns TimecodeSource
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TimecodeSource",
		"EMBEDDED",
		&returns,
	)
	return returns
}

func TimecodeSource_SYSTEMCLOCK() TimecodeSource {
	_init_.Initialize()
	var returns TimecodeSource
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TimecodeSource",
		"SYSTEMCLOCK",
		&returns,
	)
	return returns
}

func TimecodeSource_ZEROBASED() TimecodeSource {
	_init_.Initialize()
	var returns TimecodeSource
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TimecodeSource",
		"ZEROBASED",
		&returns,
	)
	return returns
}

