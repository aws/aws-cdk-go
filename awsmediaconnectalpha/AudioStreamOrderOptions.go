package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Audio Stream Order Options.
//
// Example:
//   var stack Stack
//
//
//   // Create media streams
//   videoStream := awsmediaconnectalpha.MediaStream_Video(&MediaStreamVideo{
//   	MediaStreamId: jsii.Number(1),
//   	MediaStreamName: jsii.String("video-stream"),
//   	VideoFormat: awsmediaconnectalpha.MediaVideoFormat_HD_1080P(),
//   	Fmtp: &FmtpVideo{
//   		Colorimetry: awsmediaconnectalpha.Colorimetry_BT709(),
//   		ExactFramerate: awsmediaconnectalpha.Framerate_FPS_29_97(),
//   		Par: awsmediaconnectalpha.PixelAspectRatio_SQUARE(),
//   		VideoRange: awsmediaconnectalpha.VideoRange_NARROW(),
//   		ScanMode: awsmediaconnectalpha.ScanMode_PROGRESSIVE(),
//   		Tcs: awsmediaconnectalpha.Tcs_SDR(),
//   	},
//   })
//
//   audioStream := awsmediaconnectalpha.MediaStream_Audio(&MediaStreamAudio{
//   	MediaStreamId: jsii.Number(2),
//   	MediaStreamName: jsii.String("audio-stream"),
//   	ChannelOrder: awsmediaconnectalpha.AudioStreamOrderOptions_STANDARD_STEREO(),
//   })
//
//   // Add to flow
//   flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyFlow"), &FlowProps{
//   	Source: awsmediaconnectalpha.SourceConfiguration_Router(),
//   	MediaStreams: []MediaStream{
//   		videoStream,
//   		audioStream,
//   	},
//   })
//
// Experimental.
type AudioStreamOrderOptions interface {
	// The audio stream order string value.
	// Experimental.
	Value() *string
	// Returns the string value.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AudioStreamOrderOptions
type jsiiProxy_AudioStreamOrderOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_AudioStreamOrderOptions) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Use a custom audio stream order value.
// Experimental.
func AudioStreamOrderOptions_Of(value *string) AudioStreamOrderOptions {
	_init_.Initialize()

	if err := validateAudioStreamOrderOptions_OfParameters(value); err != nil {
		panic(err)
	}
	var returns AudioStreamOrderOptions

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.AudioStreamOrderOptions",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func AudioStreamOrderOptions_DUAL_MONO() AudioStreamOrderOptions {
	_init_.Initialize()
	var returns AudioStreamOrderOptions
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.AudioStreamOrderOptions",
		"DUAL_MONO",
		&returns,
	)
	return returns
}

func AudioStreamOrderOptions_LTRT_MATRIX_STEREO() AudioStreamOrderOptions {
	_init_.Initialize()
	var returns AudioStreamOrderOptions
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.AudioStreamOrderOptions",
		"LTRT_MATRIX_STEREO",
		&returns,
	)
	return returns
}

func AudioStreamOrderOptions_MONO() AudioStreamOrderOptions {
	_init_.Initialize()
	var returns AudioStreamOrderOptions
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.AudioStreamOrderOptions",
		"MONO",
		&returns,
	)
	return returns
}

func AudioStreamOrderOptions_ONE_SDI_AUDIO_GROUP() AudioStreamOrderOptions {
	_init_.Initialize()
	var returns AudioStreamOrderOptions
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.AudioStreamOrderOptions",
		"ONE_SDI_AUDIO_GROUP",
		&returns,
	)
	return returns
}

func AudioStreamOrderOptions_STANDARD_STEREO() AudioStreamOrderOptions {
	_init_.Initialize()
	var returns AudioStreamOrderOptions
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.AudioStreamOrderOptions",
		"STANDARD_STEREO",
		&returns,
	)
	return returns
}

func AudioStreamOrderOptions_SURROUND_22_2() AudioStreamOrderOptions {
	_init_.Initialize()
	var returns AudioStreamOrderOptions
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.AudioStreamOrderOptions",
		"SURROUND_22_2",
		&returns,
	)
	return returns
}

func AudioStreamOrderOptions_SURROUND_5_1() AudioStreamOrderOptions {
	_init_.Initialize()
	var returns AudioStreamOrderOptions
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.AudioStreamOrderOptions",
		"SURROUND_5_1",
		&returns,
	)
	return returns
}

func AudioStreamOrderOptions_SURROUND_7_1() AudioStreamOrderOptions {
	_init_.Initialize()
	var returns AudioStreamOrderOptions
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.AudioStreamOrderOptions",
		"SURROUND_7_1",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AudioStreamOrderOptions) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

