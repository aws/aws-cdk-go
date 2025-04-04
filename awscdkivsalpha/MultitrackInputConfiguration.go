package awscdkivsalpha


// A complex type that specifies multitrack input configuration.
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
type MultitrackInputConfiguration struct {
	// Maximum resolution for multitrack input.
	// Experimental.
	MaximumResolution MaximumResolution `field:"required" json:"maximumResolution" yaml:"maximumResolution"`
	// Indicates whether multitrack input is allowed or required.
	// Experimental.
	Policy Policy `field:"required" json:"policy" yaml:"policy"`
}

