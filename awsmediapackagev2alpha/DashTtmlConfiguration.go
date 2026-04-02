package awsmediapackagev2alpha


// Configuration for TTML subtitles in DASH manifests.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediapackagev2_alpha "github.com/aws/aws-cdk-go/awsmediapackagev2alpha"
//
//   dashTtmlConfiguration := &DashTtmlConfiguration{
//   	TtmlProfile: mediapackagev2_alpha.TtmlProfile_IMSC_1,
//   }
//
// Experimental.
type DashTtmlConfiguration struct {
	// The profile that MediaPackage uses when signaling subtitles in the manifest.
	//
	// IMSC is the default profile. EBU-TT-D produces subtitles that are compliant with the EBU-TT-D TTML profile. MediaPackage passes through subtitle styles to the manifest.
	// For more information about EBU-TT-D subtitles, see EBU-TT-D Subtitling Distribution Format.
	// Experimental.
	TtmlProfile TtmlProfile `field:"required" json:"ttmlProfile" yaml:"ttmlProfile"`
}

