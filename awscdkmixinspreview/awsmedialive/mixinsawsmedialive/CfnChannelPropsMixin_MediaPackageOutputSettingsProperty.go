package mixinsawsmedialive


// The settings for a MediaPackage output.
//
// The parent of this entity is OutputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mediaPackageOutputSettingsProperty := &MediaPackageOutputSettingsProperty{
//   	MediaPackageV2DestinationSettings: &MediaPackageV2DestinationSettingsProperty{
//   		AudioGroupId: jsii.String("audioGroupId"),
//   		AudioRenditionSets: jsii.String("audioRenditionSets"),
//   		HlsAutoSelect: jsii.String("hlsAutoSelect"),
//   		HlsDefault: jsii.String("hlsDefault"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackageoutputsettings.html
//
type CfnChannelPropsMixin_MediaPackageOutputSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackageoutputsettings.html#cfn-medialive-channel-mediapackageoutputsettings-mediapackagev2destinationsettings
	//
	MediaPackageV2DestinationSettings interface{} `field:"optional" json:"mediaPackageV2DestinationSettings" yaml:"mediaPackageV2DestinationSettings"`
}

