package awsmedialivealpha


// Video selector settings for an input.
//
// Example:
//   var stack Stack
//   var input IInput
//   var bucket IBucket
//   var video EncodeConfiguration
//
//
//   medialive.NewChannel(stack, jsii.String("Channel"), &ChannelProps{
//   	Inputs: []InputAttachment{
//   		&InputAttachment{
//   			Input: *Input,
//   			AudioSelectors: []AudioSelector{
//   				medialive.AudioSelector_ByLanguage(jsii.String("eng"), jsii.String("eng"), medialive.AudioLanguageSelectionPolicy_STRICT()),
//   			},
//   			CaptionSelectors: []CaptionSelector{
//   				medialive.CaptionSelector_Embedded(jsii.String("embedded")),
//   			},
//   			VideoSelector: &VideoSelectorSettings{
//   				ColorSpace: medialive.VideoColorSpace_HDR10(),
//   				ColorSpaceUsage: medialive.VideoColorSpaceUsage_FORCE(),
//   				SelectBy: medialive.VideoSelection_ByProgramId(jsii.Number(1)),
//   			},
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
//   					},
//   					OutputName: jsii.String("hls_out"),
//   				},
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type VideoSelectorSettings struct {
	// The color space of the input video.
	// Default: - service default.
	//
	// Experimental.
	ColorSpace VideoColorSpace `field:"optional" json:"colorSpace" yaml:"colorSpace"`
	// How `colorSpace` is applied when it is not `FOLLOW`.
	// Default: - MediaLive service default.
	//
	// Experimental.
	ColorSpaceUsage VideoColorSpaceUsage `field:"optional" json:"colorSpaceUsage" yaml:"colorSpaceUsage"`
	// HDR10 color space metadata for the input.
	// Default: - none.
	//
	// Experimental.
	Hdr10 *Hdr10Settings `field:"optional" json:"hdr10" yaml:"hdr10"`
	// Selects the specific video to extract from the input (by PID or by program).
	// Default: - MediaLive selects the video automatically.
	//
	// Experimental.
	SelectBy VideoSelection `field:"optional" json:"selectBy" yaml:"selectBy"`
}

