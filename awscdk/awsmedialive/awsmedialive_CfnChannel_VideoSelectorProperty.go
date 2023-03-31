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
//   videoSelectorProperty := &videoSelectorProperty{
//   	colorSpace: jsii.String("colorSpace"),
//   	colorSpaceSettings: &videoSelectorColorSpaceSettingsProperty{
//   		hdr10Settings: &hdr10SettingsProperty{
//   			maxCll: jsii.Number(123),
//   			maxFall: jsii.Number(123),
//   		},
//   	},
//   	colorSpaceUsage: jsii.String("colorSpaceUsage"),
//   	selectorSettings: &videoSelectorSettingsProperty{
//   		videoSelectorPid: &videoSelectorPidProperty{
//   			pid: jsii.Number(123),
//   		},
//   		videoSelectorProgramId: &videoSelectorProgramIdProperty{
//   			programId: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnChannel_VideoSelectorProperty struct {
	// Specifies the color space of an input.
	//
	// This setting works in tandem with colorSpaceConversion to determine if MediaLive will perform any conversion.
	ColorSpace *string `field:"optional" json:"colorSpace" yaml:"colorSpace"`
	// Settings to configure color space settings in the incoming video.
	ColorSpaceSettings interface{} `field:"optional" json:"colorSpaceSettings" yaml:"colorSpaceSettings"`
	// Applies only if colorSpace is a value other than Follow.
	//
	// This field controls how the value in the colorSpace field is used. Fallback means that when the input does include color space data, that data is used, but when the input has no color space data, the value in colorSpace is used. Choose fallback if your input is sometimes missing color space data, but when it does have color space data, that data is correct. Force means to always use the value in colorSpace. Choose force if your input usually has no color space data or might have unreliable color space data.
	ColorSpaceUsage *string `field:"optional" json:"colorSpaceUsage" yaml:"colorSpaceUsage"`
	// Information about the video to select from the content.
	SelectorSettings interface{} `field:"optional" json:"selectorSettings" yaml:"selectorSettings"`
}

