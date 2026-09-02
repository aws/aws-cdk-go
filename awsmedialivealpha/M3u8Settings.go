package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// M3U8 container settings for a standard HLS output.
//
// Use `M3u8Settings.of()` to control the transport stream produced by a standard HLS output.
// Omitting it uses MediaLive's service defaults.
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
type M3u8Settings interface {
}

// The jsii proxy struct for M3u8Settings
type jsiiProxy_M3u8Settings struct {
	_ byte // padding
}

// Create M3U8 container settings.
// Experimental.
func M3u8Settings_Of(props *M3u8SettingsProps) M3u8Settings {
	_init_.Initialize()

	if err := validateM3u8Settings_OfParameters(props); err != nil {
		panic(err)
	}
	var returns M3u8Settings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.M3u8Settings",
		"of",
		[]interface{}{props},
		&returns,
	)

	return returns
}

