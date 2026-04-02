package awsmediapackagev2alpha

import (
	"time"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Filter configuration includes settings for manifest filtering, start and end times, and time delay that apply to all of your egress requests for this manifest.
//
// Example:
//   var channel Channel
//
//
//   awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("Endpoint"), &OriginEndpointProps{
//   	Channel: Channel,
//   	Segment: awsmediapackagev2alpha.Segment_Cmaf(),
//   	Manifests: []Manifest{
//   		awsmediapackagev2alpha.Manifest_Hls(&HlsManifestConfiguration{
//   			ManifestName: jsii.String("index"),
//   			FilterConfiguration: &FilterConfiguration{
//   				ManifestFilter: []ManifestFilter{
//   					awsmediapackagev2alpha.ManifestFilter_NumericCombo(awsmediapackagev2alpha.NumericFilterKey_VIDEO_HEIGHT, []NumericExpression{
//   						awsmediapackagev2alpha.NumericExpression_Range(jsii.Number(240), jsii.Number(360)),
//   						awsmediapackagev2alpha.NumericExpression_*Range(jsii.Number(720), jsii.Number(1080)),
//   						awsmediapackagev2alpha.NumericExpression_Value(jsii.Number(1440)),
//   					}),
//   				},
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type FilterConfiguration struct {
	// Optionally specify the clip start time for all of your manifest egress requests.
	//
	// When you include clip start time, note that you cannot use clip start time query parameters for this manifest's endpoint URL.
	//
	// This will be converted to a UTC timestamp.
	// Default: - No clip start time.
	//
	// Experimental.
	ClipStartTime *time.Time `field:"optional" json:"clipStartTime" yaml:"clipStartTime"`
	// DRM settings for manifest egress requests.
	//
	// When you include a DRM setting, note that you cannot use an identical
	// DRM setting query parameter for this manifest's endpoint URL.
	// Default: - No DRM settings.
	//
	// Experimental.
	DrmSettings *[]DrmSettingsKey `field:"optional" json:"drmSettings" yaml:"drmSettings"`
	// Optionally specify the end time for all of your manifest egress requests.
	//
	// When you include end time, note that you cannot use end time query parameters for this manifest's endpoint URL.
	//
	// This will be converted to a UTC timestamp.
	// Default: - No end time.
	//
	// Experimental.
	End *time.Time `field:"optional" json:"end" yaml:"end"`
	// Optionally specify one or more manifest filters for all of your manifest egress requests.
	//
	// When you include a manifest filter, note that you cannot use an identical manifest filter query parameter for this manifest's endpoint URL.
	// Default: - No manifest filters.
	//
	// Experimental.
	ManifestFilter *[]ManifestFilter `field:"optional" json:"manifestFilter" yaml:"manifestFilter"`
	// Optionally specify the start time for all of your manifest egress requests.
	//
	// When you include start time, note that you cannot use start time query parameters for this manifest's endpoint URL.
	//
	// This will be converted to a UTC timestamp.
	// Default: - No start time.
	//
	// Experimental.
	Start *time.Time `field:"optional" json:"start" yaml:"start"`
	// Optionally specify the time delay for all of your manifest egress requests.
	//
	// Enter a value that is smaller than your endpoint's startover window.
	// When you include time delay, note that you cannot use time delay query parameters for this manifest's endpoint URL.
	// Default: - No time delay.
	//
	// Experimental.
	TimeDelay awscdk.Duration `field:"optional" json:"timeDelay" yaml:"timeDelay"`
}

