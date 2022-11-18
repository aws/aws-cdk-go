package awsmediapackage


// Parameters for a packaging configuration that uses Common Media Application Format (CMAF) packaging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cmafPackageProperty := &cmafPackageProperty{
//   	hlsManifests: []interface{}{
//   		&hlsManifestProperty{
//   			adMarkers: jsii.String("adMarkers"),
//   			includeIframeOnlyStream: jsii.Boolean(false),
//   			manifestName: jsii.String("manifestName"),
//   			programDateTimeIntervalSeconds: jsii.Number(123),
//   			repeatExtXKey: jsii.Boolean(false),
//   			streamSelection: &streamSelectionProperty{
//   				maxVideoBitsPerSecond: jsii.Number(123),
//   				minVideoBitsPerSecond: jsii.Number(123),
//   				streamOrder: jsii.String("streamOrder"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	encryption: &cmafEncryptionProperty{
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
//   	includeEncoderConfigurationInSegments: jsii.Boolean(false),
//   	segmentDurationSeconds: jsii.Number(123),
//   }
//
type CfnPackagingConfiguration_CmafPackageProperty struct {
	// A list of HLS manifest configurations that are available from this endpoint.
	HlsManifests interface{} `field:"required" json:"hlsManifests" yaml:"hlsManifests"`
	// Parameters for encrypting content.
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// When includeEncoderConfigurationInSegments is set to true, MediaPackage places your encoder's Sequence Parameter Set (SPS), Picture Parameter Set (PPS), and Video Parameter Set (VPS) metadata in every video segment instead of in the init fragment.
	//
	// This lets you use different SPS/PPS/VPS settings for your assets during content playback.
	IncludeEncoderConfigurationInSegments interface{} `field:"optional" json:"includeEncoderConfigurationInSegments" yaml:"includeEncoderConfigurationInSegments"`
	// Duration (in seconds) of each segment.
	//
	// Actual segments are rounded to the nearest multiple of the source fragment duration.
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
}

