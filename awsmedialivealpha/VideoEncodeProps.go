package awsmedialivealpha


// Properties for a video encode configuration.
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
type VideoEncodeProps struct {
	// The codec for the video encode.
	//
	// Choose the codec explicitly (e.g. `VideoCodecSettings.h264(...)`)
	// Experimental.
	Codec VideoCodecSettings `field:"required" json:"codec" yaml:"codec"`
	// The height of the output video in pixels.
	//
	// Must be an even number.
	// Experimental.
	Height *float64 `field:"required" json:"height" yaml:"height"`
	// A unique name for this video encode.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The width of the output video in pixels.
	//
	// Must be an even number.
	// Experimental.
	Width *float64 `field:"required" json:"width" yaml:"width"`
	// How to respond to AFD values in the input stream.
	// Default: RespondToAfd.NONE
	//
	// Experimental.
	RespondToAfd RespondToAfd `field:"optional" json:"respondToAfd" yaml:"respondToAfd"`
	// The video scaling behavior.
	// Default: ScalingBehavior.DEFAULT
	//
	// Experimental.
	ScalingBehavior ScalingBehavior `field:"optional" json:"scalingBehavior" yaml:"scalingBehavior"`
	// The anti-alias filter strength (0-100).
	//
	// 0 is softest, 100 is sharpest.
	// Default: 50.
	//
	// Experimental.
	Sharpness *float64 `field:"optional" json:"sharpness" yaml:"sharpness"`
}

