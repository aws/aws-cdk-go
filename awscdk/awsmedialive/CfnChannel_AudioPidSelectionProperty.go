package awsmedialive


// Used to extract audio by The PID.
//
// The parent of this entity is AudioSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioPidSelectionProperty := &AudioPidSelectionProperty{
//   	Pid: jsii.Number(123),
//   	Pids: []interface{}{
//   		&AudioPidProperty{
//   			DolbyEDecode: &AudioDolbyEDecodeProperty{
//   				ProgramSelection: jsii.String("programSelection"),
//   			},
//   			Pid: jsii.Number(123),
//   			PremixSettings: &AudioPreMixerSettingsProperty{
//   				AudioNormalizationSettings: &AudioNormalizationSettingsProperty{
//   					Algorithm: jsii.String("algorithm"),
//   					AlgorithmControl: jsii.String("algorithmControl"),
//   					PeakCalculation: jsii.String("peakCalculation"),
//   					PeakLimiterThreshold: jsii.Number(123),
//   					TargetLkfs: jsii.Number(123),
//   				},
//   				Channels: jsii.Number(123),
//   				GainDb: jsii.Number(123),
//   				RemixSettings: &RemixSettingsProperty{
//   					ChannelMappings: []interface{}{
//   						&AudioChannelMappingProperty{
//   							InputChannelLevels: []interface{}{
//   								&InputChannelLevelProperty{
//   									Gain: jsii.Number(123),
//   									InputChannel: jsii.Number(123),
//   								},
//   							},
//   							OutputChannel: jsii.Number(123),
//   						},
//   					},
//   					ChannelsIn: jsii.Number(123),
//   					ChannelsOut: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiopidselection.html
//
type CfnChannel_AudioPidSelectionProperty struct {
	// Select the audio by this PID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiopidselection.html#cfn-medialive-channel-audiopidselection-pid
	//
	Pid *float64 `field:"optional" json:"pid" yaml:"pid"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiopidselection.html#cfn-medialive-channel-audiopidselection-pids
	//
	Pids interface{} `field:"optional" json:"pids" yaml:"pids"`
}

