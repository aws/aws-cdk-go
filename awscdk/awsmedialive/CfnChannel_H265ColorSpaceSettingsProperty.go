package awsmedialive


// H265 Color Space Settings.
//
// The parent of this entity is H265Settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   h265ColorSpaceSettingsProperty := &H265ColorSpaceSettingsProperty{
//   	ColorSpacePassthroughSettings: &ColorSpacePassthroughSettingsProperty{
//   	},
//   	DolbyVision81Settings: &DolbyVision81SettingsProperty{
//   	},
//   	Hdr10Settings: &Hdr10SettingsProperty{
//   		MaxCll: jsii.Number(123),
//   		MaxFall: jsii.Number(123),
//   	},
//   	Rec601Settings: &Rec601SettingsProperty{
//   	},
//   	Rec709Settings: &Rec709SettingsProperty{
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h265colorspacesettings.html
//
type CfnChannel_H265ColorSpaceSettingsProperty struct {
	// Passthrough applies no color space conversion to the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h265colorspacesettings.html#cfn-medialive-channel-h265colorspacesettings-colorspacepassthroughsettings
	//
	ColorSpacePassthroughSettings interface{} `field:"optional" json:"colorSpacePassthroughSettings" yaml:"colorSpacePassthroughSettings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h265colorspacesettings.html#cfn-medialive-channel-h265colorspacesettings-dolbyvision81settings
	//
	DolbyVision81Settings interface{} `field:"optional" json:"dolbyVision81Settings" yaml:"dolbyVision81Settings"`
	// Settings to configure the handling of HDR10 color space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h265colorspacesettings.html#cfn-medialive-channel-h265colorspacesettings-hdr10settings
	//
	Hdr10Settings interface{} `field:"optional" json:"hdr10Settings" yaml:"hdr10Settings"`
	// Settings to configure the handling of Rec601 color space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h265colorspacesettings.html#cfn-medialive-channel-h265colorspacesettings-rec601settings
	//
	Rec601Settings interface{} `field:"optional" json:"rec601Settings" yaml:"rec601Settings"`
	// Settings to configure the handling of Rec709 color space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h265colorspacesettings.html#cfn-medialive-channel-h265colorspacesettings-rec709settings
	//
	Rec709Settings interface{} `field:"optional" json:"rec709Settings" yaml:"rec709Settings"`
}

