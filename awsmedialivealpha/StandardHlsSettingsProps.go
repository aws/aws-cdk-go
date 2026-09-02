package awsmedialivealpha


// Properties for standard (video) HLS settings.
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
type StandardHlsSettingsProps struct {
	// The audio GROUP-IDs used with this video output stream, comma-separated.
	// Default: - service default.
	//
	// Experimental.
	AudioRenditionSets *string `field:"optional" json:"audioRenditionSets" yaml:"audioRenditionSets"`
	// The M3U8 container settings.
	// Default: - service defaults.
	//
	// Experimental.
	M3u8Settings M3u8Settings `field:"optional" json:"m3u8Settings" yaml:"m3u8Settings"`
}

