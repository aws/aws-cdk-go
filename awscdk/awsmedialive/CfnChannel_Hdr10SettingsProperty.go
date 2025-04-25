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
//   hdr10SettingsProperty := &Hdr10SettingsProperty{
//   	MaxCll: jsii.Number(123),
//   	MaxFall: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hdr10settings.html
//
type CfnChannel_Hdr10SettingsProperty struct {
	// Maximum Content Light Level An integer metadata value defining the maximum light level, in nits, of any single pixel within an encoded HDR video stream or file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hdr10settings.html#cfn-medialive-channel-hdr10settings-maxcll
	//
	MaxCll *float64 `field:"optional" json:"maxCll" yaml:"maxCll"`
	// Maximum Frame Average Light Level An integer metadata value defining the maximum average light level, in nits, for any single frame within an encoded HDR video stream or file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hdr10settings.html#cfn-medialive-channel-hdr10settings-maxfall
	//
	MaxFall *float64 `field:"optional" json:"maxFall" yaml:"maxFall"`
}

