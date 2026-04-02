package awsmediapackagev2alpha


// A collection of audio encryption presets.
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
type PresetSpeke20Audio string

const (
	// Use one content key to encrypt all of the audio tracks in your stream.
	// Experimental.
	PresetSpeke20Audio_PRESET_AUDIO_1 PresetSpeke20Audio = "PRESET_AUDIO_1"
	// Use one content key to encrypt all of the stereo audio tracks and one content key to encrypt all of the multichannel audio tracks.
	// Experimental.
	PresetSpeke20Audio_PRESET_AUDIO_2 PresetSpeke20Audio = "PRESET_AUDIO_2"
	// Use one content key to encrypt all of the stereo audio tracks, one content key to encrypt all of the multichannel audio tracks with 3 to 6 channels, and one content key to encrypt all of the multichannel audio tracks with more than 6 channels.
	// Experimental.
	PresetSpeke20Audio_PRESET_AUDIO_3 PresetSpeke20Audio = "PRESET_AUDIO_3"
	// Use the same content key for all of the audio and video tracks in your stream.
	// Experimental.
	PresetSpeke20Audio_SHARED PresetSpeke20Audio = "SHARED"
	// Don't encrypt any of the audio tracks in your stream.
	// Experimental.
	PresetSpeke20Audio_UNENCRYPTED PresetSpeke20Audio = "UNENCRYPTED"
)

