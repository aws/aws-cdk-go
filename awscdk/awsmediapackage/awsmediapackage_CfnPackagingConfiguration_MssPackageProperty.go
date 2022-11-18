package awsmediapackage


// Parameters for a packaging configuration that uses Microsoft Smooth Streaming (MSS) packaging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mssPackageProperty := &mssPackageProperty{
//   	mssManifests: []interface{}{
//   		&mssManifestProperty{
//   			manifestName: jsii.String("manifestName"),
//   			streamSelection: &streamSelectionProperty{
//   				maxVideoBitsPerSecond: jsii.Number(123),
//   				minVideoBitsPerSecond: jsii.Number(123),
//   				streamOrder: jsii.String("streamOrder"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	encryption: &mssEncryptionProperty{
//   		spekeKeyProvider: &spekeKeyProviderProperty{
//   			roleArn: jsii.String("roleArn"),
//   			systemIds: []*string{
//   				jsii.String("systemIds"),
//   			},
//   			url: jsii.String("url"),
//
//   			// the properties below are optional
//   			encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   				presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   				presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   			},
//   		},
//   	},
//   	segmentDurationSeconds: jsii.Number(123),
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

