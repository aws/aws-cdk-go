package awsmedialivealpha


// Timecode configuration for the channel.
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
type TimecodeConfig struct {
	// The source of timecode.
	// Default: TimecodeSource.EMBEDDED
	//
	// Experimental.
	Source TimecodeSource `field:"optional" json:"source" yaml:"source"`
	// The threshold in frames beyond which output timecode is resynchronized to the input timecode.
	// Default: - no sync threshold.
	//
	// Experimental.
	SyncThreshold *float64 `field:"optional" json:"syncThreshold" yaml:"syncThreshold"`
}

