package awsmedialivealpha


// Audio remix settings for channel remapping.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   remixSettings := &RemixSettings{
//   	ChannelMappings: []AudioChannelMapping{
//   		&AudioChannelMapping{
//   			InputChannelLevels: []InputChannelLevel{
//   				&InputChannelLevel{
//   					InputChannel: jsii.Number(123),
//
//   					// the properties below are optional
//   					Gain: jsii.Number(123),
//   				},
//   			},
//   			OutputChannel: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	ChannelsIn: jsii.Number(123),
//   	ChannelsOut: jsii.Number(123),
//   }
//
// Experimental.
type RemixSettings struct {
	// The channel mappings from input to output.
	// Experimental.
	ChannelMappings *[]*AudioChannelMapping `field:"required" json:"channelMappings" yaml:"channelMappings"`
	// The number of input channels.
	// Default: - auto-detected.
	//
	// Experimental.
	ChannelsIn *float64 `field:"optional" json:"channelsIn" yaml:"channelsIn"`
	// The number of output channels.
	//
	// Valid values: 1, 2, 4, 6, 8.
	// Default: - auto-detected.
	//
	// Experimental.
	ChannelsOut *float64 `field:"optional" json:"channelsOut" yaml:"channelsOut"`
}

