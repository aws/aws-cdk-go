package awsmedialive


// Information about the video to extract from the input. An input can contain only one video selector.
//
// The parent of this entity is InputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   videoSelectorProperty := &VideoSelectorProperty{
//   	ColorSpace: jsii.String("colorSpace"),
//   	ColorSpaceSettings: &VideoSelectorColorSpaceSettingsProperty{
//   		Hdr10Settings: &Hdr10SettingsProperty{
//   			MaxCll: jsii.Number(123),
//   			MaxFall: jsii.Number(123),
//   		},
//   	},
//   	ColorSpaceUsage: jsii.String("colorSpaceUsage"),
//   	SelectorSettings: &VideoSelectorSettingsProperty{
//   		VideoSelectorPid: &VideoSelectorPidProperty{
//   			Pid: jsii.Number(123),
//   		},
//   		VideoSelectorProgramId: &VideoSelectorProgramIdProperty{
//   			ProgramId: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videoselector.html
//
type CfnChannel_VideoSelectorProperty struct {
	// Specifies the color space of an input.
	//
	// This setting works in tandem with colorSpaceConversion to determine if MediaLive will perform any conversion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videoselector.html#cfn-medialive-channel-videoselector-colorspace
	//
	ColorSpace *string `field:"optional" json:"colorSpace" yaml:"colorSpace"`
	// Settings to configure color space settings in the incoming video.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videoselector.html#cfn-medialive-channel-videoselector-colorspacesettings
	//
	ColorSpaceSettings interface{} `field:"optional" json:"colorSpaceSettings" yaml:"colorSpaceSettings"`
	// Applies only if colorSpace is a value other than Follow.
	//
	// This field controls how the value in the colorSpace field is used. Fallback means that when the input does include color space data, that data is used, but when the input has no color space data, the value in colorSpace is used. Choose fallback if your input is sometimes missing color space data, but when it does have color space data, that data is correct. Force means to always use the value in colorSpace. Choose force if your input usually has no color space data or might have unreliable color space data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videoselector.html#cfn-medialive-channel-videoselector-colorspaceusage
	//
	ColorSpaceUsage *string `field:"optional" json:"colorSpaceUsage" yaml:"colorSpaceUsage"`
	// Information about the video to select from the content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videoselector.html#cfn-medialive-channel-videoselector-selectorsettings
	//
	SelectorSettings interface{} `field:"optional" json:"selectorSettings" yaml:"selectorSettings"`
}

