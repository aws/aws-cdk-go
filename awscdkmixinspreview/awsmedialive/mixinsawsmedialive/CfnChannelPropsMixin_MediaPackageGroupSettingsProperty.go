package mixinsawsmedialive


// The settings for the MediaPackage group.
//
// The parent of this entity is OutputGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mediaPackageGroupSettingsProperty := &MediaPackageGroupSettingsProperty{
//   	Destination: &OutputLocationRefProperty{
//   		DestinationRefId: jsii.String("destinationRefId"),
//   	},
//   	MediapackageV2GroupSettings: &MediaPackageV2GroupSettingsProperty{
//   		CaptionLanguageMappings: []interface{}{
//   			&CaptionLanguageMappingProperty{
//   				CaptionChannel: jsii.Number(123),
//   				LanguageCode: jsii.String("languageCode"),
//   				LanguageDescription: jsii.String("languageDescription"),
//   			},
//   		},
//   		Id3Behavior: jsii.String("id3Behavior"),
//   		KlvBehavior: jsii.String("klvBehavior"),
//   		NielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   		Scte35Type: jsii.String("scte35Type"),
//   		SegmentLength: jsii.Number(123),
//   		SegmentLengthUnits: jsii.String("segmentLengthUnits"),
//   		TimedMetadataId3Frame: jsii.String("timedMetadataId3Frame"),
//   		TimedMetadataId3Period: jsii.Number(123),
//   		TimedMetadataPassthrough: jsii.String("timedMetadataPassthrough"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackagegroupsettings.html
//
type CfnChannelPropsMixin_MediaPackageGroupSettingsProperty struct {
	// The MediaPackage channel destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackagegroupsettings.html#cfn-medialive-channel-mediapackagegroupsettings-destination
	//
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackagegroupsettings.html#cfn-medialive-channel-mediapackagegroupsettings-mediapackagev2groupsettings
	//
	MediapackageV2GroupSettings interface{} `field:"optional" json:"mediapackageV2GroupSettings" yaml:"mediapackageV2GroupSettings"`
}

