package awsmediapackagev2alpha


// The SPEKE Version 2.0 preset video associated with the encryption contract configuration of the origin endpoint.
//
// A collection of video encryption presets.
//
// Example:
//   var channel Channel
//   var spekeRole IRole
//
//
//   awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("Endpoint"), &OriginEndpointProps{
//   	Channel: Channel,
//   	Segment: awsmediapackagev2alpha.Segment_Cmaf(&CmafSegmentProps{
//   		Encryption: awsmediapackagev2alpha.CmafEncryption_Speke(&CmafSpekeEncryptionProps{
//   			Method: awsmediapackagev2alpha.CmafEncryptionMethod_CBCS,
//   			DrmSystems: []CmafDrmSystem{
//   				awsmediapackagev2alpha.CmafDrmSystem_FAIRPLAY,
//   				awsmediapackagev2alpha.CmafDrmSystem_WIDEVINE,
//   			},
//   			ResourceId: jsii.String("my-content-id"),
//   			Url: jsii.String("https://example.com/speke"),
//   			Role: spekeRole,
//   			KeyRotationInterval: awscdk.Duration_Seconds(jsii.Number(300)),
//   			AudioPreset: awsmediapackagev2alpha.PresetSpeke20Audio_PRESET_AUDIO_2,
//   			VideoPreset: awsmediapackagev2alpha.PresetSpeke20Video_PRESET_VIDEO_2,
//   		}),
//   	}),
//   	Manifests: []Manifest{
//   		awsmediapackagev2alpha.Manifest_Hls(&HlsManifestConfiguration{
//   			ManifestName: jsii.String("index"),
//   		}),
//   	},
//   })
//
// Experimental.
type PresetSpeke20Video string

const (
	// Use one content key to encrypt all of the video tracks in your stream.
	// Experimental.
	PresetSpeke20Video_PRESET_VIDEO_1 PresetSpeke20Video = "PRESET_VIDEO_1"
	// Use one content key to encrypt all of the SD video tracks and one content key for all HD and higher resolutions video tracks.
	// Experimental.
	PresetSpeke20Video_PRESET_VIDEO_2 PresetSpeke20Video = "PRESET_VIDEO_2"
	// Use one content key to encrypt all of the SD video tracks, one content key for HD video tracks and one content key for all UHD video tracks.
	// Experimental.
	PresetSpeke20Video_PRESET_VIDEO_3 PresetSpeke20Video = "PRESET_VIDEO_3"
	// Use one content key to encrypt all of the SD video tracks, one content key for HD video tracks, one content key for all UHD1 video tracks and one content key for all UHD2 video tracks.
	// Experimental.
	PresetSpeke20Video_PRESET_VIDEO_4 PresetSpeke20Video = "PRESET_VIDEO_4"
	// Use one content key to encrypt all of the SD video tracks, one content key for HD1 video tracks, one content key for HD2 video tracks, one content key for all UHD1 video tracks and one content key for all UHD2 video tracks.
	// Experimental.
	PresetSpeke20Video_PRESET_VIDEO_5 PresetSpeke20Video = "PRESET_VIDEO_5"
	// Use one content key to encrypt all of the SD video tracks, one content key for HD1 video tracks, one content key for HD2 video tracks and one content key for all UHD video tracks.
	// Experimental.
	PresetSpeke20Video_PRESET_VIDEO_6 PresetSpeke20Video = "PRESET_VIDEO_6"
	// Use one content key to encrypt all of the SD+HD1 video tracks, one content key for HD2 video tracks and one content key for all UHD video tracks.
	// Experimental.
	PresetSpeke20Video_PRESET_VIDEO_7 PresetSpeke20Video = "PRESET_VIDEO_7"
	// Use one content key to encrypt all of the SD+HD1 video tracks, one content key for HD2 video tracks, one content key for all UHD1 video tracks and one content key for all UHD2 video tracks.
	// Experimental.
	PresetSpeke20Video_PRESET_VIDEO_8 PresetSpeke20Video = "PRESET_VIDEO_8"
	// Use the same content key for all of the video and audio tracks in your stream.
	// Experimental.
	PresetSpeke20Video_SHARED PresetSpeke20Video = "SHARED"
	// Don't encrypt any of the video tracks in your stream.
	// Experimental.
	PresetSpeke20Video_UNENCRYPTED PresetSpeke20Video = "UNENCRYPTED"
)

