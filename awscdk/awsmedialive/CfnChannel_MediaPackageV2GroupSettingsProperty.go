package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mediaPackageV2GroupSettingsProperty := &MediaPackageV2GroupSettingsProperty{
//   	CaptionLanguageMappings: []interface{}{
//   		&CaptionLanguageMappingProperty{
//   			CaptionChannel: jsii.Number(123),
//   			LanguageCode: jsii.String("languageCode"),
//   			LanguageDescription: jsii.String("languageDescription"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackagev2groupsettings.html
//
type CfnChannel_MediaPackageV2GroupSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackagev2groupsettings.html#cfn-medialive-channel-mediapackagev2groupsettings-captionlanguagemappings
	//
	CaptionLanguageMappings interface{} `field:"optional" json:"captionLanguageMappings" yaml:"captionLanguageMappings"`
}

