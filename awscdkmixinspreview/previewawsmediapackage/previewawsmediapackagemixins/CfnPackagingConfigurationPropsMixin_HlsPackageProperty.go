package previewawsmediapackagemixins


// Parameters for a packaging configuration that uses HTTP Live Streaming (HLS) packaging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   hlsPackageProperty := &HlsPackageProperty{
//   	Encryption: &HlsEncryptionProperty{
//   		ConstantInitializationVector: jsii.String("constantInitializationVector"),
//   		EncryptionMethod: jsii.String("encryptionMethod"),
//   		SpekeKeyProvider: &SpekeKeyProviderProperty{
//   			EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   				PresetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   				PresetSpeke20Video: jsii.String("presetSpeke20Video"),
//   			},
//   			RoleArn: jsii.String("roleArn"),
//   			SystemIds: []*string{
//   				jsii.String("systemIds"),
//   			},
//   			Url: jsii.String("url"),
//   		},
//   	},
//   	HlsManifests: []interface{}{
//   		&HlsManifestProperty{
//   			AdMarkers: jsii.String("adMarkers"),
//   			IncludeIframeOnlyStream: jsii.Boolean(false),
//   			ManifestName: jsii.String("manifestName"),
//   			ProgramDateTimeIntervalSeconds: jsii.Number(123),
//   			RepeatExtXKey: jsii.Boolean(false),
//   			StreamSelection: &StreamSelectionProperty{
//   				MaxVideoBitsPerSecond: jsii.Number(123),
//   				MinVideoBitsPerSecond: jsii.Number(123),
//   				StreamOrder: jsii.String("streamOrder"),
//   			},
//   		},
//   	},
//   	IncludeDvbSubtitles: jsii.Boolean(false),
//   	SegmentDurationSeconds: jsii.Number(123),
//   	UseAudioRenditionGroup: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-hlspackage.html
//
type CfnPackagingConfigurationPropsMixin_HlsPackageProperty struct {
	// Parameters for encrypting content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-hlspackage.html#cfn-mediapackage-packagingconfiguration-hlspackage-encryption
	//
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// A list of HLS manifest configurations that are available from this endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-hlspackage.html#cfn-mediapackage-packagingconfiguration-hlspackage-hlsmanifests
	//
	HlsManifests interface{} `field:"optional" json:"hlsManifests" yaml:"hlsManifests"`
	// When enabled, MediaPackage passes through digital video broadcasting (DVB) subtitles into the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-hlspackage.html#cfn-mediapackage-packagingconfiguration-hlspackage-includedvbsubtitles
	//
	IncludeDvbSubtitles interface{} `field:"optional" json:"includeDvbSubtitles" yaml:"includeDvbSubtitles"`
	// Duration (in seconds) of each fragment.
	//
	// Actual fragments are rounded to the nearest multiple of the source fragment duration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-hlspackage.html#cfn-mediapackage-packagingconfiguration-hlspackage-segmentdurationseconds
	//
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
	// When true, AWS Elemental MediaPackage bundles all audio tracks in a rendition group.
	//
	// All other tracks in the stream can be used with any audio rendition from the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-hlspackage.html#cfn-mediapackage-packagingconfiguration-hlspackage-useaudiorenditiongroup
	//
	UseAudioRenditionGroup interface{} `field:"optional" json:"useAudioRenditionGroup" yaml:"useAudioRenditionGroup"`
}

