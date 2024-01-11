package awsmediatailor


// Dash manifest configuration parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashPlaylistSettingsProperty := &DashPlaylistSettingsProperty{
//   	ManifestWindowSeconds: jsii.Number(123),
//   	MinBufferTimeSeconds: jsii.Number(123),
//   	MinUpdatePeriodSeconds: jsii.Number(123),
//   	SuggestedPresentationDelaySeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-channel-dashplaylistsettings.html
//
type CfnChannel_DashPlaylistSettingsProperty struct {
	// The total duration (in seconds) of each manifest.
	//
	// Minimum value: `30` seconds. Maximum value: `3600` seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-channel-dashplaylistsettings.html#cfn-mediatailor-channel-dashplaylistsettings-manifestwindowseconds
	//
	ManifestWindowSeconds *float64 `field:"optional" json:"manifestWindowSeconds" yaml:"manifestWindowSeconds"`
	// Minimum amount of content (measured in seconds) that a player must keep available in the buffer.
	//
	// Minimum value: `2` seconds. Maximum value: `60` seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-channel-dashplaylistsettings.html#cfn-mediatailor-channel-dashplaylistsettings-minbuffertimeseconds
	//
	MinBufferTimeSeconds *float64 `field:"optional" json:"minBufferTimeSeconds" yaml:"minBufferTimeSeconds"`
	// Minimum amount of time (in seconds) that the player should wait before requesting updates to the manifest.
	//
	// Minimum value: `2` seconds. Maximum value: `60` seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-channel-dashplaylistsettings.html#cfn-mediatailor-channel-dashplaylistsettings-minupdateperiodseconds
	//
	MinUpdatePeriodSeconds *float64 `field:"optional" json:"minUpdatePeriodSeconds" yaml:"minUpdatePeriodSeconds"`
	// Amount of time (in seconds) that the player should be from the live point at the end of the manifest.
	//
	// Minimum value: `2` seconds. Maximum value: `60` seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-channel-dashplaylistsettings.html#cfn-mediatailor-channel-dashplaylistsettings-suggestedpresentationdelayseconds
	//
	SuggestedPresentationDelaySeconds *float64 `field:"optional" json:"suggestedPresentationDelaySeconds" yaml:"suggestedPresentationDelaySeconds"`
}

