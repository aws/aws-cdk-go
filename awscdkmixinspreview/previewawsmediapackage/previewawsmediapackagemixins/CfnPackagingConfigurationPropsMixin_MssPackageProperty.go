package previewawsmediapackagemixins


// Parameters for a packaging configuration that uses Microsoft Smooth Streaming (MSS) packaging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mssPackageProperty := &MssPackageProperty{
//   	Encryption: &MssEncryptionProperty{
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
//   	MssManifests: []interface{}{
//   		&MssManifestProperty{
//   			ManifestName: jsii.String("manifestName"),
//   			StreamSelection: &StreamSelectionProperty{
//   				MaxVideoBitsPerSecond: jsii.Number(123),
//   				MinVideoBitsPerSecond: jsii.Number(123),
//   				StreamOrder: jsii.String("streamOrder"),
//   			},
//   		},
//   	},
//   	SegmentDurationSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-msspackage.html
//
type CfnPackagingConfigurationPropsMixin_MssPackageProperty struct {
	// Parameters for encrypting content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-msspackage.html#cfn-mediapackage-packagingconfiguration-msspackage-encryption
	//
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// A list of Microsoft Smooth manifest configurations that are available from this endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-msspackage.html#cfn-mediapackage-packagingconfiguration-msspackage-mssmanifests
	//
	MssManifests interface{} `field:"optional" json:"mssManifests" yaml:"mssManifests"`
	// Duration (in seconds) of each fragment.
	//
	// Actual fragments are rounded to the nearest multiple of the source fragment duration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-msspackage.html#cfn-mediapackage-packagingconfiguration-msspackage-segmentdurationseconds
	//
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
}

