package awsmedialive


// Information about one audio track to extract. You can select multiple tracks.
//
// The parent of this entity is AudioTrackSelection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioTrackProperty := &AudioTrackProperty{
//   	PremixSettings: &AudioPreMixerSettingsProperty{
//   		AudioNormalizationSettings: &AudioNormalizationSettingsProperty{
//   			Algorithm: jsii.String("algorithm"),
//   			AlgorithmControl: jsii.String("algorithmControl"),
//   			PeakCalculation: jsii.String("peakCalculation"),
//   			PeakLimiterThreshold: jsii.Number(123),
//   			TargetLkfs: jsii.Number(123),
//   		},
//   		Channels: jsii.Number(123),
//   		GainDb: jsii.Number(123),
//   		RemixSettings: &RemixSettingsProperty{
//   			ChannelMappings: []interface{}{
//   				&AudioChannelMappingProperty{
//   					InputChannelLevels: []interface{}{
//   						&InputChannelLevelProperty{
//   							Gain: jsii.Number(123),
//   							InputChannel: jsii.Number(123),
//   						},
//   					},
//   					OutputChannel: jsii.Number(123),
//   				},
//   			},
//   			ChannelsIn: jsii.Number(123),
//   			ChannelsOut: jsii.Number(123),
//   		},
//   	},
//   	Track: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiotrack.html
//
type CfnChannel_AudioTrackProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiotrack.html#cfn-medialive-channel-audiotrack-premixsettings
	//
	PremixSettings interface{} `field:"optional" json:"premixSettings" yaml:"premixSettings"`
	// 1-based integer value that maps to a specific audio track.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiotrack.html#cfn-medialive-channel-audiotrack-track
	//
	Track *float64 `field:"optional" json:"track" yaml:"track"`
}

