package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A video frame rate expressed as a rational number (numerator/denominator).
//
// Use the predefined constants for standard rates, or {@link Framerate.of} for a custom rate.
//
// Example:
//   var stack Stack
//   var input IInput
//   var bucket IBucket
//
//
//   video := medialive.EncodeConfiguration_Video(&VideoEncodeProps{
//   	Name: jsii.String("video_720p"),
//   	Codec: medialive.VideoCodecSettings_H264(&H264SettingsProps{
//   		RateControl: medialive.H264RateControl_Cbr(&CbrRateControlProps{
//   			Bitrate: awscdk.Bitrate_Mbps(jsii.Number(3)),
//   		}),
//   		Framerate: medialive.Framerate_FPS_30(),
//   	}),
//   	Width: jsii.Number(1280),
//   	Height: jsii.Number(720),
//   })
//
//   audio := medialive.EncodeConfiguration_Audio(&AudioEncodeProps{
//   	Name: jsii.String("audio_aac"),
//   	Codec: medialive.AudioCodecSettings_Aac(&AacSettingsProps{
//   		Bitrate: awscdk.Bitrate_Kbps(jsii.Number(192)),
//   	}),
//   })
//
//   medialive.NewChannel(stack, jsii.String("Channel"), &ChannelProps{
//   	Inputs: []InputAttachment{
//   		&InputAttachment{
//   			Input: *Input,
//   		},
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
type Framerate interface {
	// Returns the string value.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Framerate
type jsiiProxy_Framerate struct {
	_ byte // padding
}

// Define a custom frame rate.
// Experimental.
func Framerate_Of(numerator *float64, denominator *float64) Framerate {
	_init_.Initialize()

	if err := validateFramerate_OfParameters(numerator, denominator); err != nil {
		panic(err)
	}
	var returns Framerate

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Framerate",
		"of",
		[]interface{}{numerator, denominator},
		&returns,
	)

	return returns
}

func Framerate_FPS_24() Framerate {
	_init_.Initialize()
	var returns Framerate
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Framerate",
		"FPS_24",
		&returns,
	)
	return returns
}

func Framerate_FPS_25() Framerate {
	_init_.Initialize()
	var returns Framerate
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Framerate",
		"FPS_25",
		&returns,
	)
	return returns
}

func Framerate_FPS_29_97() Framerate {
	_init_.Initialize()
	var returns Framerate
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Framerate",
		"FPS_29_97",
		&returns,
	)
	return returns
}

func Framerate_FPS_30() Framerate {
	_init_.Initialize()
	var returns Framerate
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Framerate",
		"FPS_30",
		&returns,
	)
	return returns
}

func Framerate_FPS_50() Framerate {
	_init_.Initialize()
	var returns Framerate
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Framerate",
		"FPS_50",
		&returns,
	)
	return returns
}

func Framerate_FPS_59_94() Framerate {
	_init_.Initialize()
	var returns Framerate
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Framerate",
		"FPS_59_94",
		&returns,
	)
	return returns
}

func Framerate_FPS_60() Framerate {
	_init_.Initialize()
	var returns Framerate
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Framerate",
		"FPS_60",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_Framerate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

