package awsmediapackage


// Parameters for Common Media Application Format (CMAF) packaging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cmafPackageProperty := &cmafPackageProperty{
//   	encryption: &cmafEncryptionProperty{
//   		spekeKeyProvider: &spekeKeyProviderProperty{
//   			resourceId: jsii.String("resourceId"),
//   			roleArn: jsii.String("roleArn"),
//   			systemIds: []*string{
//   				jsii.String("systemIds"),
//   			},
//   			url: jsii.String("url"),
//
//   			// the properties below are optional
//   			certificateArn: jsii.String("certificateArn"),
//   			encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   				presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   				presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		constantInitializationVector: jsii.String("constantInitializationVector"),
//   		encryptionMethod: jsii.String("encryptionMethod"),
//   		keyRotationIntervalSeconds: jsii.Number(123),
//   	},
//   	hlsManifests: []interface{}{
//   		&hlsManifestProperty{
//   			id: jsii.String("id"),
//
//   			// the properties below are optional
//   			adMarkers: jsii.String("adMarkers"),
//   			adsOnDeliveryRestrictions: jsii.String("adsOnDeliveryRestrictions"),
//   			adTriggers: []*string{
//   				jsii.String("adTriggers"),
//   			},
//   			includeIframeOnlyStream: jsii.Boolean(false),
//   			manifestName: jsii.String("manifestName"),
//   			playlistType: jsii.String("playlistType"),
//   			playlistWindowSeconds: jsii.Number(123),
//   			programDateTimeIntervalSeconds: jsii.Number(123),
//   			url: jsii.String("url"),
//   		},
//   	},
//   	segmentDurationSeconds: jsii.Number(123),
//   	segmentPrefix: jsii.String("segmentPrefix"),
//   	streamSelection: &streamSelectionProperty{
//   		maxVideoBitsPerSecond: jsii.Number(123),
//   		minVideoBitsPerSecond: jsii.Number(123),
//   		streamOrder: jsii.String("streamOrder"),
//   	},
//   }
//
type CfnOriginEndpoint_CmafPackageProperty struct {
	// Parameters for encrypting content.
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// A list of HLS manifest configurations that are available from this endpoint.
	HlsManifests interface{} `field:"optional" json:"hlsManifests" yaml:"hlsManifests"`
	// Duration (in seconds) of each segment.
	//
	// Actual segments are rounded to the nearest multiple of the source segment duration.
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
	// An optional custom string that is prepended to the name of each segment.
	//
	// If not specified, the segment prefix defaults to the ChannelId.
	SegmentPrefix *string `field:"optional" json:"segmentPrefix" yaml:"segmentPrefix"`
	// Limitations for outputs from the endpoint, based on the video bitrate.
	StreamSelection interface{} `field:"optional" json:"streamSelection" yaml:"streamSelection"`
}

