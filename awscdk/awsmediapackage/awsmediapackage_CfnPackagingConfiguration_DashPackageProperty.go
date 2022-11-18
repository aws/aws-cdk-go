package awsmediapackage


// Parameters for a packaging configuration that uses Dynamic Adaptive Streaming over HTTP (DASH) packaging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashPackageProperty := &dashPackageProperty{
//   	dashManifests: []interface{}{
//   		&dashManifestProperty{
//   			manifestLayout: jsii.String("manifestLayout"),
//   			manifestName: jsii.String("manifestName"),
//   			minBufferTimeSeconds: jsii.Number(123),
//   			profile: jsii.String("profile"),
//   			scteMarkersSource: jsii.String("scteMarkersSource"),
//   			streamSelection: &streamSelectionProperty{
//   				maxVideoBitsPerSecond: jsii.Number(123),
//   				minVideoBitsPerSecond: jsii.Number(123),
//   				streamOrder: jsii.String("streamOrder"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	encryption: &dashEncryptionProperty{
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
//   	periodTriggers: []*string{
//   		jsii.String("periodTriggers"),
//   	},
//   	segmentDurationSeconds: jsii.Number(123),
//   	segmentTemplateFormat: jsii.String("segmentTemplateFormat"),
//   }
//
type CfnPackagingConfiguration_DashPackageProperty struct {
	// A list of DASH manifest configurations that are available from this endpoint.
	DashManifests interface{} `field:"required" json:"dashManifests" yaml:"dashManifests"`
	// Parameters for encrypting content.
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// When includeEncoderConfigurationInSegments is set to true, MediaPackage places your encoder's Sequence Parameter Set (SPS), Picture Parameter Set (PPS), and Video Parameter Set (VPS) metadata in every video segment instead of in the init fragment.
	//
	// This lets you use different SPS/PPS/VPS settings for your assets during content playback.
	IncludeEncoderConfigurationInSegments interface{} `field:"optional" json:"includeEncoderConfigurationInSegments" yaml:"includeEncoderConfigurationInSegments"`
	// Controls whether MediaPackage produces single-period or multi-period DASH manifests.
	//
	// For more information about periods, see [Multi-period DASH in AWS Elemental MediaPackage](https://docs.aws.amazon.com/mediapackage/latest/ug/multi-period.html) .
	//
	// Valid values:
	//
	// - `ADS` - MediaPackage will produce multi-period DASH manifests. Periods are created based on the SCTE-35 ad markers present in the input manifest.
	// - *No value* - MediaPackage will produce single-period DASH manifests. This is the default setting.
	PeriodTriggers *[]*string `field:"optional" json:"periodTriggers" yaml:"periodTriggers"`
	// Duration (in seconds) of each fragment.
	//
	// Actual fragments are rounded to the nearest multiple of the source segment duration.
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
	// Determines the type of SegmentTemplate included in the Media Presentation Description (MPD).
	//
	// When set to `NUMBER_WITH_TIMELINE` , a full timeline is presented in each SegmentTemplate, with $Number$ media URLs. When set to `TIME_WITH_TIMELINE` , a full timeline is presented in each SegmentTemplate, with $Time$ media URLs. When set to `NUMBER_WITH_DURATION` , only a duration is included in each SegmentTemplate, with $Number$ media URLs.
	SegmentTemplateFormat *string `field:"optional" json:"segmentTemplateFormat" yaml:"segmentTemplateFormat"`
}

