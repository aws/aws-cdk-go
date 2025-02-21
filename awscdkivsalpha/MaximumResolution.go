package awscdkivsalpha


// Maximum resolution for multitrack input.
//
// Example:
//   ivs.NewChannel(this, jsii.String("ChannelWithMultitrackVideo"), &ChannelProps{
//   	Type: ivs.ChannelType_STANDARD,
//   	ContainerFormat: ivs.ContainerFormat_FRAGMENTED_MP4,
//   	MultitrackInputConfiguration: &MultitrackInputConfiguration{
//   		MaximumResolution: ivs.MaximumResolution_HD,
//   		Policy: ivs.Policy_ALLOW,
//   	},
//   })
//
// Experimental.
type MaximumResolution string

const (
	// Full HD (1080p).
	// Experimental.
	MaximumResolution_FULL_HD MaximumResolution = "FULL_HD"
	// HD (720p).
	// Experimental.
	MaximumResolution_HD MaximumResolution = "HD"
	// SD (480p).
	// Experimental.
	MaximumResolution_SD MaximumResolution = "SD"
)

