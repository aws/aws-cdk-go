package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   smartSubtitleSourceSettingsProperty := &SmartSubtitleSourceSettingsProperty{
//   	CaptionSynchronizationMode: jsii.String("captionSynchronizationMode"),
//   	InferenceFeedOutput: jsii.String("inferenceFeedOutput"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-smartsubtitlesourcesettings.html
//
type CfnChannelPropsMixin_SmartSubtitleSourceSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-smartsubtitlesourcesettings.html#cfn-medialive-channel-smartsubtitlesourcesettings-captionsynchronizationmode
	//
	CaptionSynchronizationMode *string `field:"optional" json:"captionSynchronizationMode" yaml:"captionSynchronizationMode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-smartsubtitlesourcesettings.html#cfn-medialive-channel-smartsubtitlesourcesettings-inferencefeedoutput
	//
	InferenceFeedOutput *string `field:"optional" json:"inferenceFeedOutput" yaml:"inferenceFeedOutput"`
}

