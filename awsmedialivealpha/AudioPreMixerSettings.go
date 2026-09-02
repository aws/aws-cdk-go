package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Audio pre-mixer settings for normalizing audio before interleaving.
//
// Applied per-PID or per-track before tracks are combined.
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
//   audioPreMixerSettings := medialive_alpha.AudioPreMixerSettings_Of(&AudioPreMixerSettingsProps{
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
//   })
//
// Experimental.
type AudioPreMixerSettings interface {
}

// The jsii proxy struct for AudioPreMixerSettings
type jsiiProxy_AudioPreMixerSettings struct {
	_ byte // padding
}

// Create pre-mixer settings.
// Experimental.
func AudioPreMixerSettings_Of(props *AudioPreMixerSettingsProps) AudioPreMixerSettings {
	_init_.Initialize()

	if err := validateAudioPreMixerSettings_OfParameters(props); err != nil {
		panic(err)
	}
	var returns AudioPreMixerSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AudioPreMixerSettings",
		"of",
		[]interface{}{props},
		&returns,
	)

	return returns
}

