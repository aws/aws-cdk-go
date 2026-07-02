package awsmedialive


// Information about the audio track to extract.
//
// The parent of this entity is AudioSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   audioTrackSelectionProperty := &AudioTrackSelectionProperty{
//   	DolbyEDecode: &AudioDolbyEDecodeProperty{
//   		ProgramSelection: jsii.String("programSelection"),
//   	},
//   	Tracks: []interface{}{
//   		&AudioTrackProperty{
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
//   			Track: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiotrackselection.html
//
type CfnChannelPropsMixin_AudioTrackSelectionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiotrackselection.html#cfn-medialive-channel-audiotrackselection-dolbyedecode
	//
	DolbyEDecode interface{} `field:"optional" json:"dolbyEDecode" yaml:"dolbyEDecode"`
	// Selects one or more unique audio tracks from within a source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiotrackselection.html#cfn-medialive-channel-audiotrackselection-tracks
	//
	Tracks interface{} `field:"optional" json:"tracks" yaml:"tracks"`
}

