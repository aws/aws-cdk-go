package awsmedialive


// The settings for the MediaPackage group.
//
// The parent of this entity is OutputGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackagegroupsettings.html
//
type CfnChannel_MediaPackageGroupSettingsProperty struct {
	// The MediaPackage channel destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackagegroupsettings.html#cfn-medialive-channel-mediapackagegroupsettings-destination
	//
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackagegroupsettings.html#cfn-medialive-channel-mediapackagegroupsettings-mediapackagev2groupsettings
	//
	MediapackageV2GroupSettings interface{} `field:"optional" json:"mediapackageV2GroupSettings" yaml:"mediapackageV2GroupSettings"`
}

