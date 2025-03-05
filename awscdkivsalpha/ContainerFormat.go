package awscdkivsalpha


// Container Format.
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
type ContainerFormat string

const (
	// Use MPEG-TS.
	// Experimental.
	ContainerFormat_TS ContainerFormat = "TS"
	// Use fMP4.
	// Experimental.
	ContainerFormat_FRAGMENTED_MP4 ContainerFormat = "FRAGMENTED_MP4"
)

