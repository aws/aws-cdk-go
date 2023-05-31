package awsmediapackage


// Parameters for a packaging configuration that uses Microsoft Smooth Streaming (MSS) packaging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mssPackageProperty := &MssPackageProperty{
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
//
//   	// the properties below are optional
//   	Encryption: &MssEncryptionProperty{
//   		SpekeKeyProvider: &SpekeKeyProviderProperty{
//   			RoleArn: jsii.String("roleArn"),
//   			SystemIds: []*string{
//   				jsii.String("systemIds"),
//   			},
//   			Url: jsii.String("url"),
//
//   			// the properties below are optional
//   			EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   			},
//   		},
//   	},
//   	SegmentDurationSeconds: jsii.Number(123),
//   }
//
type CfnPackagingConfiguration_MssPackageProperty struct {
	// A list of Microsoft Smooth manifest configurations that are available from this endpoint.
	MssManifests interface{} `field:"required" json:"mssManifests" yaml:"mssManifests"`
	// Parameters for encrypting content.
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// Duration (in seconds) of each fragment.
	//
	// Actual fragments are rounded to the nearest multiple of the source fragment duration.
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
}

