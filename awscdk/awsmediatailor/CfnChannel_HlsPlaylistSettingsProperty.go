package awsmediatailor


// <p>HLS playlist configuration parameters.</p>.
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
	// <p>The total duration (in seconds) of each manifest.
	//
	// Minimum value: <code>30</code> seconds. Maximum value: <code>3600</code> seconds.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-channel-hlsplaylistsettings.html#cfn-mediatailor-channel-hlsplaylistsettings-manifestwindowseconds
	//
	// Default: - 0.
	//
	ManifestWindowSeconds *float64 `field:"optional" json:"manifestWindowSeconds" yaml:"manifestWindowSeconds"`
}

