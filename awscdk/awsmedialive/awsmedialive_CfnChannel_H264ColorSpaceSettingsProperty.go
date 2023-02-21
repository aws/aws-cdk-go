package awsmedialive


// Settings for configuring color space in an H264 video encode.
//
// The parent of this entity is H264Settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   h264ColorSpaceSettingsProperty := &h264ColorSpaceSettingsProperty{
//   	colorSpacePassthroughSettings: &colorSpacePassthroughSettingsProperty{
//   	},
//   	rec601Settings: &rec601SettingsProperty{
//   	},
//   	rec709Settings: &rec709SettingsProperty{
//   	},
//   }
//
type CfnChannel_H264ColorSpaceSettingsProperty struct {
	// Passthrough applies no color space conversion to the output.
	ColorSpacePassthroughSettings interface{} `field:"optional" json:"colorSpacePassthroughSettings" yaml:"colorSpacePassthroughSettings"`
	// Settings to configure the handling of Rec601 color space.
	Rec601Settings interface{} `field:"optional" json:"rec601Settings" yaml:"rec601Settings"`
	// Settings to configure the handling of Rec709 color space.
	Rec709Settings interface{} `field:"optional" json:"rec709Settings" yaml:"rec709Settings"`
}

