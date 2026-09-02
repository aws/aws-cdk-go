package awsmedialivealpha


// Properties for audio pre-mixer settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var audioNormalizationAlgorithm AudioNormalizationAlgorithm
//   var audioNormalizationAlgorithmControl AudioNormalizationAlgorithmControl
//   var audioNormalizationPeakCalculation AudioNormalizationPeakCalculation
//
//   audioPreMixerSettingsProps := &AudioPreMixerSettingsProps{
//   	AudioNormalizationSettings: &AudioNormalizationSettings{
//   		Algorithm: audioNormalizationAlgorithm,
//   		AlgorithmControl: audioNormalizationAlgorithmControl,
//   		PeakCalculation: audioNormalizationPeakCalculation,
//   		PeakLimiterThreshold: jsii.Number(123),
//   		TargetLkfs: jsii.Number(123),
//   	},
//   	Channels: jsii.Number(123),
//   	GainDb: jsii.Number(123),
//   	RemixSettings: &RemixSettings{
//   		ChannelMappings: []AudioChannelMapping{
//   			&AudioChannelMapping{
//   				InputChannelLevels: []InputChannelLevel{
//   					&InputChannelLevel{
//   						InputChannel: jsii.Number(123),
//
//   						// the properties below are optional
//   						Gain: jsii.Number(123),
//   					},
//   				},
//   				OutputChannel: jsii.Number(123),
//   			},
//   		},
//
//   		// the properties below are optional
//   		ChannelsIn: jsii.Number(123),
//   		ChannelsOut: jsii.Number(123),
//   	},
//   }
//
// Experimental.
type AudioPreMixerSettingsProps struct {
	// Audio normalization settings for loudness control.
	// Default: - no normalization.
	//
	// Experimental.
	AudioNormalizationSettings *AudioNormalizationSettings `field:"optional" json:"audioNormalizationSettings" yaml:"audioNormalizationSettings"`
	// The number of audio channels to remix to.
	//
	// Overridden by `remixSettings` if specified.
	// Default: - pass through original channel count.
	//
	// Experimental.
	Channels *float64 `field:"optional" json:"channels" yaml:"channels"`
	// The gain adjustment in decibels (dB).
	// Default: - no gain adjustment.
	//
	// Experimental.
	GainDb *float64 `field:"optional" json:"gainDb" yaml:"gainDb"`
	// Remix settings for fine-grained channel mapping and gain levels.
	// Default: - no remixing.
	//
	// Experimental.
	RemixSettings *RemixSettings `field:"optional" json:"remixSettings" yaml:"remixSettings"`
}

