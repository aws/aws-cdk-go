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
//   h265ColorSpaceSettingsProperty := &h265ColorSpaceSettingsProperty{
//   	colorSpacePassthroughSettings: &colorSpacePassthroughSettingsProperty{
//   	},
//   	hdr10Settings: &hdr10SettingsProperty{
//   		maxCll: jsii.Number(123),
//   		maxFall: jsii.Number(123),
//   	},
//   	rec601Settings: &rec601SettingsProperty{
//   	},
//   	rec709Settings: &rec709SettingsProperty{
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

