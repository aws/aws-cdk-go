package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for AAC codec settings.
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
type AacSettingsProps struct {
	// The average bitrate.
	// Default: Bitrate.kbps(192)
	//
	// Experimental.
	Bitrate awscdk.Bitrate `field:"optional" json:"bitrate" yaml:"bitrate"`
	// The coding mode (mono, stereo, 5.1).
	// Default: AacCodingMode.CODING_MODE_2_0
	//
	// Experimental.
	CodingMode AacCodingMode `field:"optional" json:"codingMode" yaml:"codingMode"`
	// Set to broadcasterMixedAd when the input contains pre-mixed main audio + AD (narration) as a stereo pair.
	// Default: AacInputType.NORMAL
	//
	// Experimental.
	InputType AacInputType `field:"optional" json:"inputType" yaml:"inputType"`
	// The AAC profile.
	// Default: AacProfile.LC
	//
	// Experimental.
	Profile AacProfile `field:"optional" json:"profile" yaml:"profile"`
	// The rate control mode.
	// Default: AacRateControlMode.CBR
	//
	// Experimental.
	RateControlMode AacRateControlMode `field:"optional" json:"rateControlMode" yaml:"rateControlMode"`
	// Sets the LATM/LOAS AAC output for raw containers.
	// Default: AacRawFormat.NONE
	//
	// Experimental.
	RawFormat AacRawFormat `field:"optional" json:"rawFormat" yaml:"rawFormat"`
	// The sample rate.
	// Default: AudioSampleRate.HZ_48000
	//
	// Experimental.
	SampleRate AudioSampleRate `field:"optional" json:"sampleRate" yaml:"sampleRate"`
	// The AAC specification (MPEG-4 or MPEG-2) used to encode the audio.
	//
	// Set to `AacSpec.MPEG2` to emit MPEG-2 AAC instead of MPEG-4 AAC for raw or MPEG-2 Transport
	// Stream containers.
	// Default: AacSpec.MPEG4
	//
	// Experimental.
	Spec AacSpec `field:"optional" json:"spec" yaml:"spec"`
	// The VBR quality level.
	//
	// Used only if rateControlMode is VBR.
	// Default: - service default.
	//
	// Experimental.
	VbrQuality AacVbrQuality `field:"optional" json:"vbrQuality" yaml:"vbrQuality"`
}

