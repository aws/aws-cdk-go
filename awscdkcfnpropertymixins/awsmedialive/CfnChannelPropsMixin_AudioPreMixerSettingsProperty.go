package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   audioPreMixerSettingsProperty := &AudioPreMixerSettingsProperty{
//   	AudioNormalizationSettings: &AudioNormalizationSettingsProperty{
//   		Algorithm: jsii.String("algorithm"),
//   		AlgorithmControl: jsii.String("algorithmControl"),
//   		PeakCalculation: jsii.String("peakCalculation"),
//   		PeakLimiterThreshold: jsii.Number(123),
//   		TargetLkfs: jsii.Number(123),
//   	},
//   	Channels: jsii.Number(123),
//   	GainDb: jsii.Number(123),
//   	RemixSettings: &RemixSettingsProperty{
//   		ChannelMappings: []interface{}{
//   			&AudioChannelMappingProperty{
//   				InputChannelLevels: []interface{}{
//   					&InputChannelLevelProperty{
//   						Gain: jsii.Number(123),
//   						InputChannel: jsii.Number(123),
//   					},
//   				},
//   				OutputChannel: jsii.Number(123),
//   			},
//   		},
//   		ChannelsIn: jsii.Number(123),
//   		ChannelsOut: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiopremixersettings.html
//
type CfnChannelPropsMixin_AudioPreMixerSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiopremixersettings.html#cfn-medialive-channel-audiopremixersettings-audionormalizationsettings
	//
	AudioNormalizationSettings interface{} `field:"optional" json:"audioNormalizationSettings" yaml:"audioNormalizationSettings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiopremixersettings.html#cfn-medialive-channel-audiopremixersettings-channels
	//
	Channels *float64 `field:"optional" json:"channels" yaml:"channels"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiopremixersettings.html#cfn-medialive-channel-audiopremixersettings-gaindb
	//
	GainDb *float64 `field:"optional" json:"gainDb" yaml:"gainDb"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiopremixersettings.html#cfn-medialive-channel-audiopremixersettings-remixsettings
	//
	RemixSettings interface{} `field:"optional" json:"remixSettings" yaml:"remixSettings"`
}

