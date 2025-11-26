package previewawsmediapackagev2mixins


// Use `encryptionContractConfiguration` to configure one or more content encryption keys for your endpoints that use SPEKE Version 2.0. The encryption contract defines which content keys are used to encrypt the audio and video tracks in your stream. To configure the encryption contract, specify which audio and video encryption presets to use.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   encryptionContractConfigurationProperty := &EncryptionContractConfigurationProperty{
//   	PresetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   	PresetSpeke20Video: jsii.String("presetSpeke20Video"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-encryptioncontractconfiguration.html
//
type CfnOriginEndpointPropsMixin_EncryptionContractConfigurationProperty struct {
	// A collection of audio encryption presets.
	//
	// Value description:
	//
	// - `PRESET-AUDIO-1` - Use one content key to encrypt all of the audio tracks in your stream.
	// - `PRESET-AUDIO-2` - Use one content key to encrypt all of the stereo audio tracks and one content key to encrypt all of the multichannel audio tracks.
	// - `PRESET-AUDIO-3` - Use one content key to encrypt all of the stereo audio tracks, one content key to encrypt all of the multichannel audio tracks with 3 to 6 channels, and one content key to encrypt all of the multichannel audio tracks with more than 6 channels.
	// - `SHARED` - Use the same content key for all of the audio and video tracks in your stream.
	// - `UNENCRYPTED` - Don't encrypt any of the audio tracks in your stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-encryptioncontractconfiguration.html#cfn-mediapackagev2-originendpoint-encryptioncontractconfiguration-presetspeke20audio
	//
	PresetSpeke20Audio *string `field:"optional" json:"presetSpeke20Audio" yaml:"presetSpeke20Audio"`
	// The SPEKE Version 2.0 preset video associated with the encryption contract configuration of the origin endpoint.
	//
	// A collection of video encryption presets.
	//
	// Value description:
	//
	// - `PRESET-VIDEO-1` - Use one content key to encrypt all of the video tracks in your stream.
	// - `PRESET-VIDEO-2` - Use one content key to encrypt all of the SD video tracks and one content key for all HD and higher resolutions video tracks.
	// - `PRESET-VIDEO-3` - Use one content key to encrypt all of the SD video tracks, one content key for HD video tracks and one content key for all UHD video tracks.
	// - `PRESET-VIDEO-4` - Use one content key to encrypt all of the SD video tracks, one content key for HD video tracks, one content key for all UHD1 video tracks and one content key for all UHD2 video tracks.
	// - `PRESET-VIDEO-5` - Use one content key to encrypt all of the SD video tracks, one content key for HD1 video tracks, one content key for HD2 video tracks, one content key for all UHD1 video tracks and one content key for all UHD2 video tracks.
	// - `PRESET-VIDEO-6` - Use one content key to encrypt all of the SD video tracks, one content key for HD1 video tracks, one content key for HD2 video tracks and one content key for all UHD video tracks.
	// - `PRESET-VIDEO-7` - Use one content key to encrypt all of the SD+HD1 video tracks, one content key for HD2 video tracks and one content key for all UHD video tracks.
	// - `PRESET-VIDEO-8` - Use one content key to encrypt all of the SD+HD1 video tracks, one content key for HD2 video tracks, one content key for all UHD1 video tracks and one content key for all UHD2 video tracks.
	// - `SHARED` - Use the same content key for all of the video and audio tracks in your stream.
	// - `UNENCRYPTED` - Don't encrypt any of the video tracks in your stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-encryptioncontractconfiguration.html#cfn-mediapackagev2-originendpoint-encryptioncontractconfiguration-presetspeke20video
	//
	PresetSpeke20Video *string `field:"optional" json:"presetSpeke20Video" yaml:"presetSpeke20Video"`
}

