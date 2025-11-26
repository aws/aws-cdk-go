package previewawsmediapackagemixins


// Parameters for a packaging configuration that uses Common Media Application Format (CMAF) packaging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cmafPackageProperty := &CmafPackageProperty{
//   	Encryption: &CmafEncryptionProperty{
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
//   	IncludeEncoderConfigurationInSegments: jsii.Boolean(false),
//   	SegmentDurationSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-cmafpackage.html
//
type CfnPackagingConfigurationPropsMixin_CmafPackageProperty struct {
	// Parameters for encrypting content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-cmafpackage.html#cfn-mediapackage-packagingconfiguration-cmafpackage-encryption
	//
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// A list of HLS manifest configurations that are available from this endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-cmafpackage.html#cfn-mediapackage-packagingconfiguration-cmafpackage-hlsmanifests
	//
	HlsManifests interface{} `field:"optional" json:"hlsManifests" yaml:"hlsManifests"`
	// When includeEncoderConfigurationInSegments is set to true, AWS Elemental MediaPackage places your encoder's Sequence Parameter Set (SPS), Picture Parameter Set (PPS), and Video Parameter Set (VPS) metadata in every video segment instead of in the init fragment.
	//
	// This lets you use different SPS/PPS/VPS settings for your assets during content playback.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-cmafpackage.html#cfn-mediapackage-packagingconfiguration-cmafpackage-includeencoderconfigurationinsegments
	//
	IncludeEncoderConfigurationInSegments interface{} `field:"optional" json:"includeEncoderConfigurationInSegments" yaml:"includeEncoderConfigurationInSegments"`
	// Duration (in seconds) of each segment.
	//
	// Actual segments are rounded to the nearest multiple of the source fragment duration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-cmafpackage.html#cfn-mediapackage-packagingconfiguration-cmafpackage-segmentdurationseconds
	//
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
}

