package previewawsmedialivemixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   timecodeBurninSettingsProperty := &TimecodeBurninSettingsProperty{
//   	FontSize: jsii.String("fontSize"),
//   	Position: jsii.String("position"),
//   	Prefix: jsii.String("prefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-timecodeburninsettings.html
//
type CfnChannelPropsMixin_TimecodeBurninSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-timecodeburninsettings.html#cfn-medialive-channel-timecodeburninsettings-fontsize
	//
	FontSize *string `field:"optional" json:"fontSize" yaml:"fontSize"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-timecodeburninsettings.html#cfn-medialive-channel-timecodeburninsettings-position
	//
	Position *string `field:"optional" json:"position" yaml:"position"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-timecodeburninsettings.html#cfn-medialive-channel-timecodeburninsettings-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

