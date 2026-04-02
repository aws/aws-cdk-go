package awsmediapackagev2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The DASH manifest configuration associated with the origin endpoint.
//
// Example:
//   var channel Channel
//
//
//   awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("Endpoint"), &OriginEndpointProps{
//   	Channel: Channel,
//   	Segment: awsmediapackagev2alpha.Segment_Cmaf(),
//   	Manifests: []Manifest{
//   		awsmediapackagev2alpha.Manifest_Dash(&DashManifestConfiguration{
//   			ManifestName: jsii.String("index"),
//   			ManifestWindow: awscdk.Duration_Seconds(jsii.Number(60)),
//   			MinBufferTime: awscdk.Duration_*Seconds(jsii.Number(30)),
//   			MinUpdatePeriod: awscdk.Duration_*Seconds(jsii.Number(10)),
//   			SegmentTemplateFormat: awsmediapackagev2alpha.SegmentTemplateFormat_NUMBER_WITH_TIMELINE,
//   			PeriodTriggers: []DashPeriodTriggers{
//   				awsmediapackagev2alpha.DashPeriodTriggers_AVAILS,
//   				awsmediapackagev2alpha.DashPeriodTriggers_DRM_KEY_ROTATION,
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type DashManifestConfiguration struct {
	// The name of the manifest associated with the DASH manifest configuration.
	// Experimental.
	ManifestName *string `field:"required" json:"manifestName" yaml:"manifestName"`
	// The base URLs to use for retrieving segments.
	// Default: - No base URLs specified.
	//
	// Experimental.
	BaseUrls *[]*DashBaseUrlProperty `field:"optional" json:"baseUrls" yaml:"baseUrls"`
	// The layout of the DASH manifest that MediaPackage produces.
	// Default: DashManifestCompactness.STANDARD
	//
	// Experimental.
	Compactness DashManifestCompactness `field:"optional" json:"compactness" yaml:"compactness"`
	// DRM signaling determines the way DASH manifest signals the DRM content.
	// Default: - No DRM signaling specified.
	//
	// Experimental.
	DrmSignalling DrmSignalling `field:"optional" json:"drmSignalling" yaml:"drmSignalling"`
	// For endpoints that use the DVB-DASH profile only.
	// Default: - No DVB settings.
	//
	// Experimental.
	DvbSettings *DashDvbSettings `field:"optional" json:"dvbSettings" yaml:"dvbSettings"`
	// Filter configuration includes settings for manifest filtering, start and end times, and time delay that apply to all of your egress requests for this manifest.
	//
	// https://docs.aws.amazon.com/mediapackage/latest/userguide/manifest-filter-query-parameters.html
	// Default: - No filter configuration.
	//
	// Experimental.
	FilterConfiguration *FilterConfiguration `field:"optional" json:"filterConfiguration" yaml:"filterConfiguration"`
	// The total duration (in seconds) of the manifest's content.
	// Default: 60.
	//
	// Experimental.
	ManifestWindow awscdk.Duration `field:"optional" json:"manifestWindow" yaml:"manifestWindow"`
	// The minimum amount of content that the player must keep available in the buffer.
	// Default: 5.
	//
	// Experimental.
	MinBufferTime awscdk.Duration `field:"optional" json:"minBufferTime" yaml:"minBufferTime"`
	// The minimum amount of time for the player to wait before requesting an updated manifest.
	// Default: 2.
	//
	// Experimental.
	MinUpdatePeriod awscdk.Duration `field:"optional" json:"minUpdatePeriod" yaml:"minUpdatePeriod"`
	// Specify what triggers cause AWS Elemental MediaPackage to create media presentation description (MPD) periods in the output manifest.
	// Default: [DashPeriodTriggers.AVAILS, DashPeriodTriggers.DRM_KEY_ROTATION, DashPeriodTriggers.SOURCE_CHANGES, DashPeriodTriggers.SOURCE_DISRUPTIONS]
	//
	// Experimental.
	PeriodTriggers *[]DashPeriodTriggers `field:"optional" json:"periodTriggers" yaml:"periodTriggers"`
	// The profile that the output is compliant with.
	// Default: - No profiles specified.
	//
	// Experimental.
	Profiles *[]*string `field:"optional" json:"profiles" yaml:"profiles"`
	// Details about the content that you want MediaPackage to pass through in the manifest to the playback device.
	// Default: - No program information.
	//
	// Experimental.
	ProgramInformation *DashProgramInformation `field:"optional" json:"programInformation" yaml:"programInformation"`
	// Choose how ad markers are included in the packaged content.
	//
	// If you include ad markers in the content stream in your upstream encoders,
	// then you need to inform MediaPackage what to do with the ad markers in the output.
	//
	// To choose this option STCE filtering needs to be enabled.
	// Default: AdMarkerDash.XML
	//
	// Experimental.
	ScteDashAdMarker AdMarkerDash `field:"optional" json:"scteDashAdMarker" yaml:"scteDashAdMarker"`
	// The type of variable that MediaPackage uses in the media attribute of the SegmentTemplate tag.
	// Default: SegmentTemplateFormat.NUMBER_WITH_TIMELINE
	//
	// Experimental.
	SegmentTemplateFormat SegmentTemplateFormat `field:"optional" json:"segmentTemplateFormat" yaml:"segmentTemplateFormat"`
	// The configuration for DASH subtitles.
	// Default: - No subtitle configuration.
	//
	// Experimental.
	SubtitleConfiguration *DashSubtitleConfiguration `field:"optional" json:"subtitleConfiguration" yaml:"subtitleConfiguration"`
	// The amount of time that the player should be from the end of the manifest.
	// Default: 10.
	//
	// Experimental.
	SuggestedPresentationDelay awscdk.Duration `field:"optional" json:"suggestedPresentationDelay" yaml:"suggestedPresentationDelay"`
	// The UTC timing mode.
	// Default: DashUtcTimingMode.UTC_DIRECT
	//
	// Experimental.
	UtcTimingMode DashUtcTimingMode `field:"optional" json:"utcTimingMode" yaml:"utcTimingMode"`
	// The method that the player uses to synchronize to coordinated universal time (UTC) wall clock time.
	// Default: undefined - No value is specified.
	//
	// Experimental.
	UtcTimingSource *string `field:"optional" json:"utcTimingSource" yaml:"utcTimingSource"`
}

