package awsmediatailor


// HLS playlist configuration parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsPlaylistSettingsProperty := &HlsPlaylistSettingsProperty{
//   	AdMarkupType: []*string{
//   		jsii.String("adMarkupType"),
//   	},
//   	ManifestWindowSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-channel-hlsplaylistsettings.html
//
type CfnChannel_HlsPlaylistSettingsProperty struct {
	// Determines the type of SCTE 35 tags to use in ad markup.
	//
	// Specify `DATERANGE` to use `DATERANGE` tags (for live or VOD content). Specify `SCTE35_ENHANCED` to use `EXT-X-CUE-OUT` and `EXT-X-CUE-IN` tags (for VOD content only).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-channel-hlsplaylistsettings.html#cfn-mediatailor-channel-hlsplaylistsettings-admarkuptype
	//
	AdMarkupType *[]*string `field:"optional" json:"adMarkupType" yaml:"adMarkupType"`
	// The total duration (in seconds) of each manifest.
	//
	// Minimum value: `30` seconds. Maximum value: `3600` seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-channel-hlsplaylistsettings.html#cfn-mediatailor-channel-hlsplaylistsettings-manifestwindowseconds
	//
	// Default: - 0.
	//
	ManifestWindowSeconds *float64 `field:"optional" json:"manifestWindowSeconds" yaml:"manifestWindowSeconds"`
}

