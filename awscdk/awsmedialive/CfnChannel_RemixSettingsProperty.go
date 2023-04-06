package awsmedialive


// The settings for remixing audio in the output.
//
// The parent of this entity is AudioDescription.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   remixSettingsProperty := &RemixSettingsProperty{
//   	ChannelMappings: []interface{}{
//   		&AudioChannelMappingProperty{
//   			InputChannelLevels: []interface{}{
//   				&InputChannelLevelProperty{
//   					Gain: jsii.Number(123),
//   					InputChannel: jsii.Number(123),
//   				},
//   			},
//   			OutputChannel: jsii.Number(123),
//   		},
//   	},
//   	ChannelsIn: jsii.Number(123),
//   	ChannelsOut: jsii.Number(123),
//   }
//
type CfnChannel_RemixSettingsProperty struct {
	// A mapping of input channels to output channels, with appropriate gain adjustments.
	ChannelMappings interface{} `field:"optional" json:"channelMappings" yaml:"channelMappings"`
	// The number of input channels to be used.
	ChannelsIn *float64 `field:"optional" json:"channelsIn" yaml:"channelsIn"`
	// The number of output channels to be produced.
	//
	// Valid values: 1, 2, 4, 6, 8.
	ChannelsOut *float64 `field:"optional" json:"channelsOut" yaml:"channelsOut"`
}

