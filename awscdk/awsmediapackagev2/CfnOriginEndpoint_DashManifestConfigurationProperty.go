package awsmediapackagev2


// The DASH manifest configuration associated with the origin endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashManifestConfigurationProperty := &DashManifestConfigurationProperty{
//   	ManifestName: jsii.String("manifestName"),
//
//   	// the properties below are optional
//   	DrmSignaling: jsii.String("drmSignaling"),
//   	FilterConfiguration: &FilterConfigurationProperty{
//   		ClipStartTime: jsii.String("clipStartTime"),
//   		End: jsii.String("end"),
//   		ManifestFilter: jsii.String("manifestFilter"),
//   		Start: jsii.String("start"),
//   		TimeDelaySeconds: jsii.Number(123),
//   	},
//   	ManifestWindowSeconds: jsii.Number(123),
//   	MinBufferTimeSeconds: jsii.Number(123),
//   	MinUpdatePeriodSeconds: jsii.Number(123),
//   	PeriodTriggers: []*string{
//   		jsii.String("periodTriggers"),
//   	},
//   	ScteDash: &ScteDashProperty{
//   		AdMarkerDash: jsii.String("adMarkerDash"),
//   	},
//   	SegmentTemplateFormat: jsii.String("segmentTemplateFormat"),
//   	SuggestedPresentationDelaySeconds: jsii.Number(123),
//   	UtcTiming: &DashUtcTimingProperty{
//   		TimingMode: jsii.String("timingMode"),
//   		TimingSource: jsii.String("timingSource"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashmanifestconfiguration.html
//
type CfnOriginEndpoint_DashManifestConfigurationProperty struct {
	// A short string that's appended to the endpoint URL.
	//
	// The child manifest name creates a unique path to this endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-dashmanifestconfiguration-manifestname
	//
	ManifestName *string `field:"required" json:"manifestName" yaml:"manifestName"`
	// Determines how the DASH manifest signals the DRM content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-dashmanifestconfiguration-drmsignaling
	//
	DrmSignaling *string `field:"optional" json:"drmSignaling" yaml:"drmSignaling"`
	// Filter configuration includes settings for manifest filtering, start and end times, and time delay that apply to all of your egress requests for this manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-dashmanifestconfiguration-filterconfiguration
	//
	FilterConfiguration interface{} `field:"optional" json:"filterConfiguration" yaml:"filterConfiguration"`
	// The total duration (in seconds) of the manifest's content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-dashmanifestconfiguration-manifestwindowseconds
	//
	ManifestWindowSeconds *float64 `field:"optional" json:"manifestWindowSeconds" yaml:"manifestWindowSeconds"`
	// Minimum amount of content (in seconds) that a player must keep available in the buffer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-dashmanifestconfiguration-minbuffertimeseconds
	//
	MinBufferTimeSeconds *float64 `field:"optional" json:"minBufferTimeSeconds" yaml:"minBufferTimeSeconds"`
	// Minimum amount of time (in seconds) that the player should wait before requesting updates to the manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-dashmanifestconfiguration-minupdateperiodseconds
	//
	MinUpdatePeriodSeconds *float64 `field:"optional" json:"minUpdatePeriodSeconds" yaml:"minUpdatePeriodSeconds"`
	// A list of triggers that controls when AWS Elemental MediaPackage separates the MPEG-DASH manifest into multiple periods.
	//
	// Type `ADS` to indicate that AWS Elemental MediaPackage must create periods in the output manifest that correspond to SCTE-35 ad markers in the input source. Leave this value empty to indicate that the manifest is contained all in one period. For more information about periods in the DASH manifest, see [Multi-period DASH in AWS Elemental MediaPackage](https://docs.aws.amazon.com/mediapackage/latest/userguide/multi-period.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-dashmanifestconfiguration-periodtriggers
	//
	PeriodTriggers *[]*string `field:"optional" json:"periodTriggers" yaml:"periodTriggers"`
	// The SCTE configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-dashmanifestconfiguration-sctedash
	//
	ScteDash interface{} `field:"optional" json:"scteDash" yaml:"scteDash"`
	// Determines the type of variable used in the `media` URL of the `SegmentTemplate` tag in the manifest.
	//
	// Also specifies if segment timeline information is included in `SegmentTimeline` or `SegmentTemplate` .
	//
	// Value description:
	//
	// - `NUMBER_WITH_TIMELINE` - The `$Number$` variable is used in the `media` URL. The value of this variable is the sequential number of the segment. A full `SegmentTimeline` object is presented in each `SegmentTemplate` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-dashmanifestconfiguration-segmenttemplateformat
	//
	SegmentTemplateFormat *string `field:"optional" json:"segmentTemplateFormat" yaml:"segmentTemplateFormat"`
	// The amount of time (in seconds) that the player should be from the end of the manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-dashmanifestconfiguration-suggestedpresentationdelayseconds
	//
	SuggestedPresentationDelaySeconds *float64 `field:"optional" json:"suggestedPresentationDelaySeconds" yaml:"suggestedPresentationDelaySeconds"`
	// Determines the type of UTC timing included in the DASH Media Presentation Description (MPD).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-dashmanifestconfiguration-utctiming
	//
	UtcTiming interface{} `field:"optional" json:"utcTiming" yaml:"utcTiming"`
}

