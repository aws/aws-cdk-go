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
type CfnChannel_H265ColorSpaceSettingsProperty struct {
	// Passthrough applies no color space conversion to the output.
	ColorSpacePassthroughSettings interface{} `field:"optional" json:"colorSpacePassthroughSettings" yaml:"colorSpacePassthroughSettings"`
	// Settings to configure the handling of HDR10 color space.
	Hdr10Settings interface{} `field:"optional" json:"hdr10Settings" yaml:"hdr10Settings"`
	// Settings to configure the handling of Rec601 color space.
	Rec601Settings interface{} `field:"optional" json:"rec601Settings" yaml:"rec601Settings"`
	// Settings to configure the handling of Rec709 color space.
	Rec709Settings interface{} `field:"optional" json:"rec709Settings" yaml:"rec709Settings"`
}

