package awsmedialivealpha


// Properties for an audio encode configuration.
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
type AudioEncodeProps struct {
	// The codec for the audio encode.
	//
	// Choose the codec explicitly (e.g. `AudioCodecSettings.aac(...)`)
	// Experimental.
	Codec AudioCodecSettings `field:"required" json:"codec" yaml:"codec"`
	// A unique name for this audio encode.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The DASH roles to assign to this audio output.
	//
	// Applies only when the output is configured
	// for DVB DASH accessibility signaling.
	// Default: - no DASH roles.
	//
	// Experimental.
	AudioDashRoles *[]AudioDashRole `field:"optional" json:"audioDashRoles" yaml:"audioDashRoles"`
	// Audio normalization settings for loudness correction.
	// Default: - no normalization.
	//
	// Experimental.
	AudioNormalization *AudioNormalizationSettings `field:"optional" json:"audioNormalization" yaml:"audioNormalization"`
	// The name of the audio selector in the input to use as the source.
	//
	// Must match the `name` of an
	// `AudioSelector` on the input attachment. When omitted, MediaLive uses the input's default audio.
	// Default: - the input's default audio.
	//
	// Experimental.
	AudioSelectorName *string `field:"optional" json:"audioSelectorName" yaml:"audioSelectorName"`
	// The audio type when audioTypeControl is USE_CONFIGURED.
	//
	// The values are defined in ISO-IEC 13818-1.
	// Default: - follow input.
	//
	// Experimental.
	AudioType AudioType `field:"optional" json:"audioType" yaml:"audioType"`
	// How the audio type is signaled in the output.
	// Default: - USE_CONFIGURED when `audioType` is set, otherwise FOLLOW_INPUT.
	//
	// Experimental.
	AudioTypeControl AudioTypeControl `field:"optional" json:"audioTypeControl" yaml:"audioTypeControl"`
	// Audio watermarking settings (e.g. Nielsen watermarks).
	// Default: - no watermarking.
	//
	// Experimental.
	AudioWatermarkSettings *AudioWatermarkSettings `field:"optional" json:"audioWatermarkSettings" yaml:"audioWatermarkSettings"`
	// DVB DASH accessibility signaling for this audio output.
	// Default: - no DVB DASH accessibility signaling.
	//
	// Experimental.
	DvbDashAccessibility DvbDashAccessibility `field:"optional" json:"dvbDashAccessibility" yaml:"dvbDashAccessibility"`
	// The ISO 639-2 language code for the audio output track (e.g. 'eng', 'spa').
	// Default: - follow input.
	//
	// Experimental.
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// How the audio language code is signaled in the output.
	//
	// When `FOLLOW_INPUT`, a configured
	// `languageCode` is used only as a fallback when the input has none.
	// Default: - USE_CONFIGURED when `languageCode` is set, otherwise FOLLOW_INPUT.
	//
	// Experimental.
	LanguageCodeControl AudioLanguageCodeControl `field:"optional" json:"languageCodeControl" yaml:"languageCodeControl"`
	// Audio remix settings for channel remapping.
	// Default: - no remixing.
	//
	// Experimental.
	RemixSettings *RemixSettings `field:"optional" json:"remixSettings" yaml:"remixSettings"`
	// The display name for the audio track (e.g. 'English', 'Director Commentary'). Used for HLS and MS Smooth outputs.
	// Default: - no stream name.
	//
	// Experimental.
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
}

