package awsmediapackagev2alpha


// Configuration for subtitles in DASH manifests.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediapackagev2_alpha "github.com/aws/aws-cdk-go/awsmediapackagev2alpha"
//
//   dashSubtitleConfiguration := &DashSubtitleConfiguration{
//   	TtmlConfiguration: &DashTtmlConfiguration{
//   		TtmlProfile: mediapackagev2_alpha.TtmlProfile_IMSC_1,
//   	},
//   }
//
// Experimental.
type DashSubtitleConfiguration struct {
	// Settings for TTML subtitles.
	// Default: - No TTML configuration.
	//
	// Experimental.
	TtmlConfiguration *DashTtmlConfiguration `field:"optional" json:"ttmlConfiguration" yaml:"ttmlConfiguration"`
}

