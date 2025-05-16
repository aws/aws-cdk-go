package awscdkivsalpha


// Whether multitrack input is allowed or required.
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
type Policy string

const (
	// Multitrack input is allowed.
	// Experimental.
	Policy_ALLOW Policy = "ALLOW"
	// Multitrack input is required.
	// Experimental.
	Policy_REQUIRE Policy = "REQUIRE"
)

