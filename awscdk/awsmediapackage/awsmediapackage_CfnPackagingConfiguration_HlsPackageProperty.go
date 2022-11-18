package awsmediapackage


// Parameters for a packaging configuration that uses HTTP Live Streaming (HLS) packaging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsPackageProperty := &hlsPackageProperty{
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
//   	encryption: &hlsEncryptionProperty{
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
//
//   		// the properties below are optional
//   		constantInitializationVector: jsii.String("constantInitializationVector"),
//   		encryptionMethod: jsii.String("encryptionMethod"),
//   	},
//   	includeDvbSubtitles: jsii.Boolean(false),
//   	segmentDurationSeconds: jsii.Number(123),
//   	useAudioRenditionGroup: jsii.Boolean(false),
//   }
//
type CfnPackagingConfiguration_HlsPackageProperty struct {
	// A list of HLS manifest configurations that are available from this endpoint.
	HlsManifests interface{} `field:"required" json:"hlsManifests" yaml:"hlsManifests"`
	// Parameters for encrypting content.
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// `CfnPackagingConfiguration.HlsPackageProperty.IncludeDvbSubtitles`.
	IncludeDvbSubtitles interface{} `field:"optional" json:"includeDvbSubtitles" yaml:"includeDvbSubtitles"`
	// Duration (in seconds) of each fragment.
	//
	// Actual fragments are rounded to the nearest multiple of the source fragment duration.
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
	// When true, AWS Elemental MediaPackage bundles all audio tracks in a rendition group.
	//
	// All other tracks in the stream can be used with any audio rendition from the group.
	UseAudioRenditionGroup interface{} `field:"optional" json:"useAudioRenditionGroup" yaml:"useAudioRenditionGroup"`
}

