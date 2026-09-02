package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Video codec settings.
//
// Use the static factory methods to create.
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
type VideoCodecSettings interface {
}

// The jsii proxy struct for VideoCodecSettings
type jsiiProxy_VideoCodecSettings struct {
	_ byte // padding
}

// Experimental.
func NewVideoCodecSettings_Override(v VideoCodecSettings) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-medialive-alpha.VideoCodecSettings",
		nil, // no parameters
		v,
	)
}

// Create AV1 codec settings.
// Experimental.
func VideoCodecSettings_Av1(props *Av1SettingsProps) VideoCodecSettings {
	_init_.Initialize()

	if err := validateVideoCodecSettings_Av1Parameters(props); err != nil {
		panic(err)
	}
	var returns VideoCodecSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.VideoCodecSettings",
		"av1",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create frame capture codec settings.
// Experimental.
func VideoCodecSettings_FrameCapture(props *FrameCaptureSettingsProps) VideoCodecSettings {
	_init_.Initialize()

	if err := validateVideoCodecSettings_FrameCaptureParameters(props); err != nil {
		panic(err)
	}
	var returns VideoCodecSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.VideoCodecSettings",
		"frameCapture",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create H.264 (AVC) codec settings.
// Experimental.
func VideoCodecSettings_H264(props *H264SettingsProps) VideoCodecSettings {
	_init_.Initialize()

	if err := validateVideoCodecSettings_H264Parameters(props); err != nil {
		panic(err)
	}
	var returns VideoCodecSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.VideoCodecSettings",
		"h264",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create H.265 (HEVC) codec settings. Framerate is required for H.265.
// Experimental.
func VideoCodecSettings_H265(props *H265SettingsProps) VideoCodecSettings {
	_init_.Initialize()

	if err := validateVideoCodecSettings_H265Parameters(props); err != nil {
		panic(err)
	}
	var returns VideoCodecSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.VideoCodecSettings",
		"h265",
		[]interface{}{props},
		&returns,
	)

	return returns
}

