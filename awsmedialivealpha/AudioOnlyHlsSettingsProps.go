package awsmedialivealpha


// Properties for audio-only HLS settings.
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
type AudioOnlyHlsSettingsProps struct {
	// The group that this audio rendition belongs to.
	// Default: - service default.
	//
	// Experimental.
	AudioGroupId *string `field:"optional" json:"audioGroupId" yaml:"audioGroupId"`
	// A .jpg or .png cover-art image embedded in each audio-only segment. Provide a `FileLocation` referencing an S3 bucket (`FileLocation.fromBucket`, which auto-grants read access) or a URL (`FileLocation.url`).
	// Default: - no cover art.
	//
	// Experimental.
	AudioOnlyImage FileLocation `field:"optional" json:"audioOnlyImage" yaml:"audioOnlyImage"`
	// How the audio rendition is signaled in the HLS manifest.
	// Default: - service default.
	//
	// Experimental.
	AudioTrackType HlsAudioTrackType `field:"optional" json:"audioTrackType" yaml:"audioTrackType"`
	// The segment container type.
	// Default: - service default.
	//
	// Experimental.
	SegmentType HlsAudioOnlySegmentType `field:"optional" json:"segmentType" yaml:"segmentType"`
}

