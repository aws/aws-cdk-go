package awsmedialive


// The settings for remixing audio.
//
// The parent of this entity is RemixSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioChannelMappingProperty := &AudioChannelMappingProperty{
//   	InputChannelLevels: []interface{}{
//   		&InputChannelLevelProperty{
//   			Gain: jsii.Number(123),
//   			InputChannel: jsii.Number(123),
//   		},
//   	},
//   	OutputChannel: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiochannelmapping.html
//
type CfnChannel_AudioChannelMappingProperty struct {
	// The indices and gain values for each input channel that should be remixed into this output channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiochannelmapping.html#cfn-medialive-channel-audiochannelmapping-inputchannellevels
	//
	InputChannelLevels interface{} `field:"optional" json:"inputChannelLevels" yaml:"inputChannelLevels"`
	// The index of the output channel that is being produced.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiochannelmapping.html#cfn-medialive-channel-audiochannelmapping-outputchannel
	//
	OutputChannel *float64 `field:"optional" json:"outputChannel" yaml:"outputChannel"`
}

