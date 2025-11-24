package mixinsawsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mediaPackageV2DestinationSettingsProperty := &MediaPackageV2DestinationSettingsProperty{
//   	AudioGroupId: jsii.String("audioGroupId"),
//   	AudioRenditionSets: jsii.String("audioRenditionSets"),
//   	HlsAutoSelect: jsii.String("hlsAutoSelect"),
//   	HlsDefault: jsii.String("hlsDefault"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackagev2destinationsettings.html
//
type CfnChannelPropsMixin_MediaPackageV2DestinationSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackagev2destinationsettings.html#cfn-medialive-channel-mediapackagev2destinationsettings-audiogroupid
	//
	AudioGroupId *string `field:"optional" json:"audioGroupId" yaml:"audioGroupId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackagev2destinationsettings.html#cfn-medialive-channel-mediapackagev2destinationsettings-audiorenditionsets
	//
	AudioRenditionSets *string `field:"optional" json:"audioRenditionSets" yaml:"audioRenditionSets"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackagev2destinationsettings.html#cfn-medialive-channel-mediapackagev2destinationsettings-hlsautoselect
	//
	HlsAutoSelect *string `field:"optional" json:"hlsAutoSelect" yaml:"hlsAutoSelect"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackagev2destinationsettings.html#cfn-medialive-channel-mediapackagev2destinationsettings-hlsdefault
	//
	HlsDefault *string `field:"optional" json:"hlsDefault" yaml:"hlsDefault"`
}

