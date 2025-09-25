package awsmedialive


// The setting to remix the audio.
//
// The parent of this entity is AudioChannelMappings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputChannelLevelProperty := &InputChannelLevelProperty{
//   	Gain: jsii.Number(123),
//   	InputChannel: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputchannellevel.html
//
type CfnChannel_InputChannelLevelProperty struct {
	// The remixing value.
	//
	// Units are in dB, and acceptable values are within the range from -60 (mute) to 6 dB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputchannellevel.html#cfn-medialive-channel-inputchannellevel-gain
	//
	Gain *float64 `field:"optional" json:"gain" yaml:"gain"`
	// The index of the input channel that is used as a source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputchannellevel.html#cfn-medialive-channel-inputchannellevel-inputchannel
	//
	InputChannel *float64 `field:"optional" json:"inputChannel" yaml:"inputChannel"`
}

