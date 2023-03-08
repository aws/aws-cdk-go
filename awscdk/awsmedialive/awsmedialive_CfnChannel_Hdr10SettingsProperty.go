package awsmedialive


// Hdr10 Settings.
//
// The parents of this entity are H265ColorSpaceSettings (for color space settings in the output) and VideoSelectorColorSpaceSettings (for color space settings in the input).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hdr10SettingsProperty := &hdr10SettingsProperty{
//   	maxCll: jsii.Number(123),
//   	maxFall: jsii.Number(123),
//   }
//
type CfnChannel_Hdr10SettingsProperty struct {
	// Maximum Content Light Level An integer metadata value defining the maximum light level, in nits, of any single pixel within an encoded HDR video stream or file.
	MaxCll *float64 `field:"optional" json:"maxCll" yaml:"maxCll"`
	// Maximum Frame Average Light Level An integer metadata value defining the maximum average light level, in nits, for any single frame within an encoded HDR video stream or file.
	MaxFall *float64 `field:"optional" json:"maxFall" yaml:"maxFall"`
}

