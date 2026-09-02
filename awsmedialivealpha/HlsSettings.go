package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Per-output HLS settings.
//
// Select the variant that matches the output: a standard (video)
// output, an audio-only rendition, an fMP4 output, or a frame-capture output.
//
// Example:
//   var bucket IBucket
//   var video EncodeConfiguration
//   var audio EncodeConfiguration
//
//
//   medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
//   	Name: jsii.String("hls"),
//   	Destinations: []OutputDestination{
//   		medialive.OutputDestination_ToBucket(bucket, jsii.String("live/stream")),
//   	},
//   	Outputs: []HlsOutputDefinition{
//   		&HlsOutputDefinition{
//   			Encodes: []EncodeConfiguration{
//   				video,
//   			},
//   			OutputName: jsii.String("video"),
//   			HlsSettings: medialive.HlsSettings_Standard(&StandardHlsSettingsProps{
//   				M3u8Settings: medialive.M3u8Settings_Of(&M3u8SettingsProps{
//   					Scte35Behavior: medialive.M3u8Scte35Behavior_PASSTHROUGH(),
//   					ProgramNum: jsii.Number(1),
//   				}),
//   			}),
//   		},
//   		&HlsOutputDefinition{
//   			Encodes: []EncodeConfiguration{
//   				audio,
//   			},
//   			OutputName: jsii.String("audio"),
//   			HlsSettings: medialive.HlsSettings_AudioOnly(&AudioOnlyHlsSettingsProps{
//   				AudioGroupId: jsii.String("program"),
//   				AudioOnlyImage: medialive.FileLocation_FromBucket(bucket, jsii.String("art/cover.png")),
//   			}),
//   		},
//   	},
//   })
//
// Experimental.
type HlsSettings interface {
}

// The jsii proxy struct for HlsSettings
type jsiiProxy_HlsSettings struct {
	_ byte // padding
}

// Experimental.
func NewHlsSettings_Override(h HlsSettings) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-medialive-alpha.HlsSettings",
		nil, // no parameters
		h,
	)
}

// Settings for an audio-only HLS rendition.
// Experimental.
func HlsSettings_AudioOnly(props *AudioOnlyHlsSettingsProps) HlsSettings {
	_init_.Initialize()

	if err := validateHlsSettings_AudioOnlyParameters(props); err != nil {
		panic(err)
	}
	var returns HlsSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsSettings",
		"audioOnly",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Settings for an fMP4 HLS output.
// Experimental.
func HlsSettings_Fmp4(props *Fmp4HlsSettingsProps) HlsSettings {
	_init_.Initialize()

	if err := validateHlsSettings_Fmp4Parameters(props); err != nil {
		panic(err)
	}
	var returns HlsSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsSettings",
		"fmp4",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Settings for a frame-capture output in an HLS output group.
// Experimental.
func HlsSettings_FrameCapture() HlsSettings {
	_init_.Initialize()

	var returns HlsSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsSettings",
		"frameCapture",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Settings for a standard (video) HLS output.
// Experimental.
func HlsSettings_Standard(props *StandardHlsSettingsProps) HlsSettings {
	_init_.Initialize()

	if err := validateHlsSettings_StandardParameters(props); err != nil {
		panic(err)
	}
	var returns HlsSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsSettings",
		"standard",
		[]interface{}{props},
		&returns,
	)

	return returns
}

