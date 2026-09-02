package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The class of the channel.
//
// Determines the pipeline redundancy.
//
// Example:
//   var stack Stack
//   var input IInput
//   var mpChannel IChannel
//
//
//   hdVideo := medialive.EncodeConfiguration_Video(&VideoEncodeProps{
//   	Name: jsii.String("video_1080p"),
//   	Codec: medialive.VideoCodecSettings_H265(&H265SettingsProps{
//   		RateControl: medialive.H265RateControl_Qvbr(&QvbrRateControlProps{
//   			MaxBitrate: awscdk.Bitrate_Mbps(jsii.Number(8)),
//   			QvbrQualityLevel: jsii.Number(7),
//   		}),
//   		Framerate: medialive.Framerate_FPS_30(),
//   	}),
//   	Width: jsii.Number(1920),
//   	Height: jsii.Number(1080),
//   })
//
//   sdVideo := medialive.EncodeConfiguration_Video(&VideoEncodeProps{
//   	Name: jsii.String("video_480p"),
//   	Codec: medialive.VideoCodecSettings_*H265(&H265SettingsProps{
//   		RateControl: medialive.H265RateControl_*Qvbr(&QvbrRateControlProps{
//   			MaxBitrate: awscdk.Bitrate_*Mbps(jsii.Number(2)),
//   			QvbrQualityLevel: jsii.Number(7),
//   		}),
//   		Framerate: medialive.Framerate_FPS_30(),
//   	}),
//   	Width: jsii.Number(854),
//   	Height: jsii.Number(480),
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
//   	ChannelClass: medialive.ChannelClass_STANDARD(),
//   	Inputs: []InputAttachment{
//   		&InputAttachment{
//   			Input: *Input,
//   		},
//   	},
//   	OutputGroups: []OutputGroupConfiguration{
//   		medialive.OutputGroupConfiguration_MediaPackageV2(&MediaPackageV2OutputGroupProps{
//   			Name: jsii.String("emp"),
//   			Channel: mpChannel,
//   			Outputs: []MediaPackageV2OutputDefinition{
//   				&MediaPackageV2OutputDefinition{
//   					Encode: hdVideo,
//   					OutputName: jsii.String("hd"),
//   				},
//   				&MediaPackageV2OutputDefinition{
//   					Encode: sdVideo,
//   					OutputName: jsii.String("sd"),
//   				},
//   				&MediaPackageV2OutputDefinition{
//   					Encode: audio,
//   					OutputName: jsii.String("audio"),
//   				},
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type ChannelClass interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for ChannelClass
type jsiiProxy_ChannelClass struct {
	_ byte // padding
}

func (j *jsiiProxy_ChannelClass) Value() *string {
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
func ChannelClass_Of(value *string) ChannelClass {
	_init_.Initialize()

	if err := validateChannelClass_OfParameters(value); err != nil {
		panic(err)
	}
	var returns ChannelClass

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.ChannelClass",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ChannelClass_SINGLE_PIPELINE() ChannelClass {
	_init_.Initialize()
	var returns ChannelClass
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.ChannelClass",
		"SINGLE_PIPELINE",
		&returns,
	)
	return returns
}

func ChannelClass_STANDARD() ChannelClass {
	_init_.Initialize()
	var returns ChannelClass
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.ChannelClass",
		"STANDARD",
		&returns,
	)
	return returns
}

