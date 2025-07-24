package awsmediatailor


// The output configuration for this channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   requestOutputItemProperty := &RequestOutputItemProperty{
//   	ManifestName: jsii.String("manifestName"),
//   	SourceGroup: jsii.String("sourceGroup"),
//
//   	// the properties below are optional
//   	DashPlaylistSettings: &DashPlaylistSettingsProperty{
//   		ManifestWindowSeconds: jsii.Number(123),
//   		MinBufferTimeSeconds: jsii.Number(123),
//   		MinUpdatePeriodSeconds: jsii.Number(123),
//   		SuggestedPresentationDelaySeconds: jsii.Number(123),
//   	},
//   	HlsPlaylistSettings: &HlsPlaylistSettingsProperty{
//   		AdMarkupType: []*string{
//   			jsii.String("adMarkupType"),
//   		},
//   		ManifestWindowSeconds: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-channel-requestoutputitem.html
//
type CfnChannel_RequestOutputItemProperty struct {
	// The name of the manifest for the channel.
	//
	// The name appears in the `PlaybackUrl` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-channel-requestoutputitem.html#cfn-mediatailor-channel-requestoutputitem-manifestname
	//
	ManifestName *string `field:"required" json:"manifestName" yaml:"manifestName"`
	// A string used to match which `HttpPackageConfiguration` is used for each `VodSource` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-channel-requestoutputitem.html#cfn-mediatailor-channel-requestoutputitem-sourcegroup
	//
	SourceGroup *string `field:"required" json:"sourceGroup" yaml:"sourceGroup"`
	// DASH manifest configuration parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-channel-requestoutputitem.html#cfn-mediatailor-channel-requestoutputitem-dashplaylistsettings
	//
	DashPlaylistSettings interface{} `field:"optional" json:"dashPlaylistSettings" yaml:"dashPlaylistSettings"`
	// HLS playlist configuration parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-channel-requestoutputitem.html#cfn-mediatailor-channel-requestoutputitem-hlsplaylistsettings
	//
	HlsPlaylistSettings interface{} `field:"optional" json:"hlsPlaylistSettings" yaml:"hlsPlaylistSettings"`
}

