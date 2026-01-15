package previewawsmediapackagev2mixins


// Filter configuration includes settings for manifest filtering, start and end times, and time delay that apply to all of your egress requests for this manifest.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   filterConfigurationProperty := &FilterConfigurationProperty{
//   	ClipStartTime: jsii.String("clipStartTime"),
//   	DrmSettings: jsii.String("drmSettings"),
//   	End: jsii.String("end"),
//   	ManifestFilter: jsii.String("manifestFilter"),
//   	Start: jsii.String("start"),
//   	TimeDelaySeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-filterconfiguration.html
//
type CfnOriginEndpointPropsMixin_FilterConfigurationProperty struct {
	// Optionally specify the clip start time for all of your manifest egress requests.
	//
	// When you include clip start time, note that you cannot use clip start time query parameters for this manifest's endpoint URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-filterconfiguration.html#cfn-mediapackagev2-originendpoint-filterconfiguration-clipstarttime
	//
	ClipStartTime *string `field:"optional" json:"clipStartTime" yaml:"clipStartTime"`
	// <p>Optionally specify one or more DRM settings for all of your manifest egress requests.
	//
	// When you include a DRM setting, note that you cannot use an identical DRM setting query parameter for this manifest's endpoint URL.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-filterconfiguration.html#cfn-mediapackagev2-originendpoint-filterconfiguration-drmsettings
	//
	DrmSettings *string `field:"optional" json:"drmSettings" yaml:"drmSettings"`
	// Optionally specify the end time for all of your manifest egress requests.
	//
	// When you include end time, note that you cannot use end time query parameters for this manifest's endpoint URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-filterconfiguration.html#cfn-mediapackagev2-originendpoint-filterconfiguration-end
	//
	End *string `field:"optional" json:"end" yaml:"end"`
	// Optionally specify one or more manifest filters for all of your manifest egress requests.
	//
	// When you include a manifest filter, note that you cannot use an identical manifest filter query parameter for this manifest's endpoint URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-filterconfiguration.html#cfn-mediapackagev2-originendpoint-filterconfiguration-manifestfilter
	//
	ManifestFilter *string `field:"optional" json:"manifestFilter" yaml:"manifestFilter"`
	// Optionally specify the start time for all of your manifest egress requests.
	//
	// When you include start time, note that you cannot use start time query parameters for this manifest's endpoint URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-filterconfiguration.html#cfn-mediapackagev2-originendpoint-filterconfiguration-start
	//
	Start *string `field:"optional" json:"start" yaml:"start"`
	// Optionally specify the time delay for all of your manifest egress requests.
	//
	// Enter a value that is smaller than your endpoint's startover window. When you include time delay, note that you cannot use time delay query parameters for this manifest's endpoint URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-filterconfiguration.html#cfn-mediapackagev2-originendpoint-filterconfiguration-timedelayseconds
	//
	TimeDelaySeconds *float64 `field:"optional" json:"timeDelaySeconds" yaml:"timeDelaySeconds"`
}

