package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var colorSpacePassthroughSettings interface{}
//   var hlg2020Settings interface{}
//   var rec601Settings interface{}
//   var rec709Settings interface{}
//
//   av1ColorSpaceSettingsProperty := &Av1ColorSpaceSettingsProperty{
//   	ColorSpacePassthroughSettings: colorSpacePassthroughSettings,
//   	Hdr10Settings: &Hdr10SettingsProperty{
//   		MaxCll: jsii.Number(123),
//   		MaxFall: jsii.Number(123),
//   	},
//   	Hlg2020Settings: hlg2020Settings,
//   	Rec601Settings: rec601Settings,
//   	Rec709Settings: rec709Settings,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1colorspacesettings.html
//
type CfnChannelPropsMixin_Av1ColorSpaceSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1colorspacesettings.html#cfn-medialive-channel-av1colorspacesettings-colorspacepassthroughsettings
	//
	ColorSpacePassthroughSettings interface{} `field:"optional" json:"colorSpacePassthroughSettings" yaml:"colorSpacePassthroughSettings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1colorspacesettings.html#cfn-medialive-channel-av1colorspacesettings-hdr10settings
	//
	Hdr10Settings interface{} `field:"optional" json:"hdr10Settings" yaml:"hdr10Settings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1colorspacesettings.html#cfn-medialive-channel-av1colorspacesettings-hlg2020settings
	//
	Hlg2020Settings interface{} `field:"optional" json:"hlg2020Settings" yaml:"hlg2020Settings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1colorspacesettings.html#cfn-medialive-channel-av1colorspacesettings-rec601settings
	//
	Rec601Settings interface{} `field:"optional" json:"rec601Settings" yaml:"rec601Settings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1colorspacesettings.html#cfn-medialive-channel-av1colorspacesettings-rec709settings
	//
	Rec709Settings interface{} `field:"optional" json:"rec709Settings" yaml:"rec709Settings"`
}

