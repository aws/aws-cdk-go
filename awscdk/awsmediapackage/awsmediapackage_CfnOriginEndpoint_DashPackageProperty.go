package awsmediapackage


// Parameters for DASH packaging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashPackageProperty := &dashPackageProperty{
//   	adsOnDeliveryRestrictions: jsii.String("adsOnDeliveryRestrictions"),
//   	adTriggers: []*string{
//   		jsii.String("adTriggers"),
//   	},
//   	encryption: &dashEncryptionProperty{
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
//   		keyRotationIntervalSeconds: jsii.Number(123),
//   	},
//   	includeIframeOnlyStream: jsii.Boolean(false),
//   	manifestLayout: jsii.String("manifestLayout"),
//   	manifestWindowSeconds: jsii.Number(123),
//   	minBufferTimeSeconds: jsii.Number(123),
//   	minUpdatePeriodSeconds: jsii.Number(123),
//   	periodTriggers: []*string{
//   		jsii.String("periodTriggers"),
//   	},
//   	profile: jsii.String("profile"),
//   	segmentDurationSeconds: jsii.Number(123),
//   	segmentTemplateFormat: jsii.String("segmentTemplateFormat"),
//   	streamSelection: &streamSelectionProperty{
//   		maxVideoBitsPerSecond: jsii.Number(123),
//   		minVideoBitsPerSecond: jsii.Number(123),
//   		streamOrder: jsii.String("streamOrder"),
//   	},
//   	suggestedPresentationDelaySeconds: jsii.Number(123),
//   	utcTiming: jsii.String("utcTiming"),
//   	utcTimingUri: jsii.String("utcTimingUri"),
//   }
//
type CfnOriginEndpoint_DashPackageProperty struct {
	// The flags on SCTE-35 segmentation descriptors that have to be present for MediaPackage to insert ad markers in the output manifest.
	//
	// For information about SCTE-35 in MediaPackage, see [SCTE-35 Message Options in AWS Elemental MediaPackage](https://docs.aws.amazon.com/mediapackage/latest/ug/scte.html) .
	AdsOnDeliveryRestrictions *string `field:"optional" json:"adsOnDeliveryRestrictions" yaml:"adsOnDeliveryRestrictions"`
	// Specifies the SCTE-35 message types that MediaPackage treats as ad markers in the output manifest.
	//
	// Valid values:
	//
	// - `BREAK`
	// - `DISTRIBUTOR_ADVERTISEMENT`
	// - `DISTRIBUTOR_OVERLAY_PLACEMENT_OPPORTUNITY` .
	// - `DISTRIBUTOR_PLACEMENT_OPPORTUNITY` .
	// - `PROVIDER_ADVERTISEMENT` .
	// - `PROVIDER_OVERLAY_PLACEMENT_OPPORTUNITY` .
	// - `PROVIDER_PLACEMENT_OPPORTUNITY` .
	// - `SPLICE_INSERT` .
	AdTriggers *[]*string `field:"optional" json:"adTriggers" yaml:"adTriggers"`
	// Parameters for encrypting content.
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// `CfnOriginEndpoint.DashPackageProperty.IncludeIframeOnlyStream`.
	IncludeIframeOnlyStream interface{} `field:"optional" json:"includeIframeOnlyStream" yaml:"includeIframeOnlyStream"`
	// Determines the position of some tags in the manifest.
	//
	// Options:
	//
	// - `FULL` - elements like `SegmentTemplate` and `ContentProtection` are included in each `Representation` .
	// - `COMPACT` - duplicate elements are combined and presented at the `AdaptationSet` level.
	ManifestLayout *string `field:"optional" json:"manifestLayout" yaml:"manifestLayout"`
	// Time window (in seconds) contained in each manifest.
	ManifestWindowSeconds *float64 `field:"optional" json:"manifestWindowSeconds" yaml:"manifestWindowSeconds"`
	// Minimum amount of content (measured in seconds) that a player must keep available in the buffer.
	MinBufferTimeSeconds *float64 `field:"optional" json:"minBufferTimeSeconds" yaml:"minBufferTimeSeconds"`
	// Minimum amount of time (in seconds) that the player should wait before requesting updates to the manifest.
	MinUpdatePeriodSeconds *float64 `field:"optional" json:"minUpdatePeriodSeconds" yaml:"minUpdatePeriodSeconds"`
	// Controls whether MediaPackage produces single-period or multi-period DASH manifests.
	//
	// For more information about periods, see [Multi-period DASH in AWS Elemental MediaPackage](https://docs.aws.amazon.com/mediapackage/latest/ug/multi-period.html) .
	//
	// Valid values:
	//
	// - `ADS` - MediaPackage will produce multi-period DASH manifests. Periods are created based on the SCTE-35 ad markers present in the input manifest.
	// - *No value* - MediaPackage will produce single-period DASH manifests. This is the default setting.
	PeriodTriggers *[]*string `field:"optional" json:"periodTriggers" yaml:"periodTriggers"`
	// DASH profile for the output, such as HbbTV.
	//
	// Valid values:
	//
	// - `NONE` - the output doesn't use a DASH profile.
	// - `HBBTV_1_5` - the output is HbbTV-compliant.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Duration (in seconds) of each fragment.
	//
	// Actual fragments are rounded to the nearest multiple of the source fragment duration.
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
	// Determines the type of variable used in the `media` URL of the `SegmentTemplate` tag in the manifest.
	//
	// Also specifies if segment timeline information is included in `SegmentTimeline` or `SegmentTemplate` .
	//
	// - `NUMBER_WITH_TIMELINE` - The `$Number$` variable is used in the `media` URL. The value of this variable is the sequential number of the segment. A full `SegmentTimeline` object is presented in each `SegmentTemplate` .
	// - `NUMBER_WITH_DURATION` - The `$Number$` variable is used in the `media` URL and a `duration` attribute is added to the segment template. The `SegmentTimeline` object is removed from the representation.
	// - `TIME_WITH_TIMELINE` - The `$Time$` variable is used in the `media` URL. The value of this variable is the timestamp of when the segment starts. A full `SegmentTimeline` object is presented in each `SegmentTemplate` .
	SegmentTemplateFormat *string `field:"optional" json:"segmentTemplateFormat" yaml:"segmentTemplateFormat"`
	// Limitations for outputs from the endpoint, based on the video bitrate.
	StreamSelection interface{} `field:"optional" json:"streamSelection" yaml:"streamSelection"`
	// Amount of time (in seconds) that the player should be from the live point at the end of the manifest.
	SuggestedPresentationDelaySeconds *float64 `field:"optional" json:"suggestedPresentationDelaySeconds" yaml:"suggestedPresentationDelaySeconds"`
	// Determines the type of UTC timing included in the DASH Media Presentation Description (MPD).
	UtcTiming *string `field:"optional" json:"utcTiming" yaml:"utcTiming"`
	// Specifies the value attribute of the UTC timing field when utcTiming is set to HTTP-ISO or HTTP-HEAD.
	UtcTimingUri *string `field:"optional" json:"utcTimingUri" yaml:"utcTimingUri"`
}

