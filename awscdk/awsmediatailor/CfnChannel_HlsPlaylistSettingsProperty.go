package awsmediatailor


// HLS playlist configuration parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsPlaylistSettingsProperty := &HlsPlaylistSettingsProperty{
//   	ManifestWindowSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-channel-hlsplaylistsettings.html
//
type CfnChannel_HlsPlaylistSettingsProperty struct {
	// The total duration (in seconds) of each manifest.
	//
	// Minimum value: `30` seconds. Maximum value: `3600` seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-channel-hlsplaylistsettings.html#cfn-mediatailor-channel-hlsplaylistsettings-manifestwindowseconds
	//
	// Default: - 0.
	//
	ManifestWindowSeconds *float64 `field:"optional" json:"manifestWindowSeconds" yaml:"manifestWindowSeconds"`
}

