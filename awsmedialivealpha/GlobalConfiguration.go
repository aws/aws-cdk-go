package awsmedialivealpha


// Global configuration settings that apply to the entire channel.
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
type GlobalConfiguration struct {
	// The initial audio gain for the channel (-60 to 60 dB).
	// Default: - service default.
	//
	// Experimental.
	InitialAudioGain *float64 `field:"optional" json:"initialAudioGain" yaml:"initialAudioGain"`
	// Action to take when the current input completes.
	// Default: - service default.
	//
	// Experimental.
	InputEndAction InputEndAction `field:"optional" json:"inputEndAction" yaml:"inputEndAction"`
	// Behavior on input loss (substitute black / repeat frame, then a color or slate image).
	// Default: - service default.
	//
	// Experimental.
	InputLossBehavior *InputLossBehavior `field:"optional" json:"inputLossBehavior" yaml:"inputLossBehavior"`
	// How MediaLive pipelines are synchronised — `OutputLocking.pipeline()`, `OutputLocking.epoch()`, or `OutputLocking.disabled()`.
	// See: https://docs.aws.amazon.com/medialive/latest/ug/pipeline-lock.html
	//
	// Default: - service default.
	//
	// Experimental.
	OutputLocking OutputLocking `field:"optional" json:"outputLocking" yaml:"outputLocking"`
	// Source of output timing.
	// Default: - service default.
	//
	// Experimental.
	OutputTimingSource OutputTimingSource `field:"optional" json:"outputTimingSource" yaml:"outputTimingSource"`
	// Enable support for low framerate inputs (e.g. music channels with less than 1 fps).
	// Default: false.
	//
	// Experimental.
	SupportLowFramerateInputs *bool `field:"optional" json:"supportLowFramerateInputs" yaml:"supportLowFramerateInputs"`
}

