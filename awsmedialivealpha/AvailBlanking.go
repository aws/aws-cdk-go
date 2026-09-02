package awsmedialivealpha


// Settings for blanking video, audio, and captions during ad avails.
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
//   		},
//   	},
//   	AvailBlanking: &AvailBlanking{
//   		State: medialive.AvailBlankingState_ENABLED(),
//   		Image: medialive.FileLocation_FromBucket(bucket, jsii.String("slates/avail.png")),
//   	},
//   	BlackoutSlate: &BlackoutSlate{
//   		State: medialive.BlackoutSlateState_ENABLED(),
//   		Image: medialive.FileLocation_*FromBucket(bucket, jsii.String("slates/blackout.png")),
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
type AvailBlanking struct {
	// A blanking image to display during ad avails.
	//
	// Provide a `FileLocation` referencing an S3
	// bucket (`FileLocation.fromBucket`, which auto-grants read access) or a URL (`FileLocation.url`).
	// Only .bmp and .png images are supported. If not set, solid black is used.
	// Default: - solid black.
	//
	// Experimental.
	Image FileLocation `field:"optional" json:"image" yaml:"image"`
	// Whether to blank the output during ad avails.
	// Default: - ENABLED if image is provided, DISABLED otherwise.
	//
	// Experimental.
	State AvailBlankingState `field:"optional" json:"state" yaml:"state"`
}

