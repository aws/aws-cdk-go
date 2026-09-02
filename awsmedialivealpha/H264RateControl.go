package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// H.264 rate control. Use the static factory methods to create.
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
type H264RateControl interface {
}

// The jsii proxy struct for H264RateControl
type jsiiProxy_H264RateControl struct {
	_ byte // padding
}

// Constant bitrate.
// Experimental.
func H264RateControl_Cbr(props *CbrRateControlProps) H264RateControl {
	_init_.Initialize()

	if err := validateH264RateControl_CbrParameters(props); err != nil {
		panic(err)
	}
	var returns H264RateControl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H264RateControl",
		"cbr",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Quality-defined variable bitrate.
// Experimental.
func H264RateControl_Qvbr(props *QvbrRateControlProps) H264RateControl {
	_init_.Initialize()

	if err := validateH264RateControl_QvbrParameters(props); err != nil {
		panic(err)
	}
	var returns H264RateControl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H264RateControl",
		"qvbr",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Variable bitrate.
// Experimental.
func H264RateControl_Vbr(props *VbrRateControlProps) H264RateControl {
	_init_.Initialize()

	if err := validateH264RateControl_VbrParameters(props); err != nil {
		panic(err)
	}
	var returns H264RateControl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H264RateControl",
		"vbr",
		[]interface{}{props},
		&returns,
	)

	return returns
}

