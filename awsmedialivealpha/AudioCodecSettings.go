package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Audio codec settings.
//
// Use the static factory methods to create.
//
// Example:
//   var stack Stack
//   var input IInput
//   var bucket IBucket
//
//
//   video := medialive.EncodeConfiguration_Video(&VideoEncodeProps{
//   	Name: jsii.String("video_720p"),
//   	Codec: medialive.VideoCodecSettings_H264(&H264SettingsProps{
//   		RateControl: medialive.H264RateControl_Cbr(&CbrRateControlProps{
//   			Bitrate: awscdk.Bitrate_Mbps(jsii.Number(3)),
//   		}),
//   		Framerate: medialive.Framerate_FPS_30(),
//   	}),
//   	Width: jsii.Number(1280),
//   	Height: jsii.Number(720),
//   })
//
//   audio := medialive.EncodeConfiguration_Audio(&AudioEncodeProps{
//   	Name: jsii.String("audio_aac"),
//   	Codec: medialive.AudioCodecSettings_Aac(&AacSettingsProps{
//   		Bitrate: awscdk.Bitrate_Kbps(jsii.Number(192)),
//   	}),
//   })
//
//   medialive.NewChannel(stack, jsii.String("Channel"), &ChannelProps{
//   	Inputs: []InputAttachment{
//   		&InputAttachment{
//   			Input: *Input,
//   		},
//   	},
//   	OutputGroups: []OutputGroupConfiguration{
//   		medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
//   			Name: jsii.String("hls"),
//   			Destinations: []OutputDestination{
//   				medialive.OutputDestination_ToBucket(bucket, jsii.String("live/stream")),
//   			},
//   			Outputs: []HlsOutputDefinition{
//   				&HlsOutputDefinition{
//   					Encodes: []EncodeConfiguration{
//   						video,
//   						audio,
//   					},
//   					OutputName: jsii.String("hls_out"),
//   				},
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type AudioCodecSettings interface {
}

// The jsii proxy struct for AudioCodecSettings
type jsiiProxy_AudioCodecSettings struct {
	_ byte // padding
}

// Experimental.
func NewAudioCodecSettings_Override(a AudioCodecSettings) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-medialive-alpha.AudioCodecSettings",
		nil, // no parameters
		a,
	)
}

// Create AAC codec settings.
// Experimental.
func AudioCodecSettings_Aac(props *AacSettingsProps) AudioCodecSettings {
	_init_.Initialize()

	if err := validateAudioCodecSettings_AacParameters(props); err != nil {
		panic(err)
	}
	var returns AudioCodecSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AudioCodecSettings",
		"aac",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create AC3 codec settings.
// Experimental.
func AudioCodecSettings_Ac3(props *Ac3SettingsProps) AudioCodecSettings {
	_init_.Initialize()

	if err := validateAudioCodecSettings_Ac3Parameters(props); err != nil {
		panic(err)
	}
	var returns AudioCodecSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AudioCodecSettings",
		"ac3",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create EAC3 (Dolby Digital Plus) codec settings.
// Experimental.
func AudioCodecSettings_Eac3(props *Eac3SettingsProps) AudioCodecSettings {
	_init_.Initialize()

	if err := validateAudioCodecSettings_Eac3Parameters(props); err != nil {
		panic(err)
	}
	var returns AudioCodecSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AudioCodecSettings",
		"eac3",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create EAC3 Atmos (Dolby Digital Plus with Atmos) codec settings.
// Experimental.
func AudioCodecSettings_Eac3Atmos(props *Eac3AtmosSettingsProps) AudioCodecSettings {
	_init_.Initialize()

	if err := validateAudioCodecSettings_Eac3AtmosParameters(props); err != nil {
		panic(err)
	}
	var returns AudioCodecSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AudioCodecSettings",
		"eac3Atmos",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create MP2 codec settings.
// Experimental.
func AudioCodecSettings_Mp2(props *Mp2SettingsProps) AudioCodecSettings {
	_init_.Initialize()

	if err := validateAudioCodecSettings_Mp2Parameters(props); err != nil {
		panic(err)
	}
	var returns AudioCodecSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AudioCodecSettings",
		"mp2",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create passthrough audio settings (no transcoding).
// Experimental.
func AudioCodecSettings_Passthrough() AudioCodecSettings {
	_init_.Initialize()

	var returns AudioCodecSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AudioCodecSettings",
		"passthrough",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Create WAV codec settings.
// Experimental.
func AudioCodecSettings_Wav(props *WavSettingsProps) AudioCodecSettings {
	_init_.Initialize()

	if err := validateAudioCodecSettings_WavParameters(props); err != nil {
		panic(err)
	}
	var returns AudioCodecSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AudioCodecSettings",
		"wav",
		[]interface{}{props},
		&returns,
	)

	return returns
}

