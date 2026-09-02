package awsmedialivealpha


// Blackout slate configuration for the channel.
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
type BlackoutSlate struct {
	// The blackout slate image.
	//
	// Provide a `FileLocation` referencing an S3 bucket
	// (`FileLocation.fromBucket`, which auto-grants read access) or a URL (`FileLocation.url`).
	// Only .bmp and .png supported.
	// Default: - solid black.
	//
	// Experimental.
	Image FileLocation `field:"optional" json:"image" yaml:"image"`
	// Whether to enable network end blackout (triggered by SCTE-35 Network End Segmentation Descriptor).
	// Default: - ENABLED if networkEndBlackoutImage is provided, DISABLED otherwise.
	//
	// Experimental.
	NetworkEndBlackout NetworkEndBlackout `field:"optional" json:"networkEndBlackout" yaml:"networkEndBlackout"`
	// The network end blackout image.
	//
	// Provide a `FileLocation` referencing an S3 bucket
	// (`FileLocation.fromBucket`, which auto-grants read access) or a URL (`FileLocation.url`).
	// Default: - solid black.
	//
	// Experimental.
	NetworkEndBlackoutImage FileLocation `field:"optional" json:"networkEndBlackoutImage" yaml:"networkEndBlackoutImage"`
	// The EIDR network ID (e.g. '10.XXXX/XXXX-XXXX-XXXX-XXXX-XXXX-C').
	// Default: - no network ID.
	//
	// Experimental.
	NetworkId *string `field:"optional" json:"networkId" yaml:"networkId"`
	// Whether to enable the blackout slate.
	// Default: - ENABLED if image is provided, DISABLED otherwise.
	//
	// Experimental.
	State BlackoutSlateState `field:"optional" json:"state" yaml:"state"`
}

