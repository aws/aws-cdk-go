package mixinsawsmediapackagev2


// The settings for TTML subtitles.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dashTtmlConfigurationProperty := &DashTtmlConfigurationProperty{
//   	TtmlProfile: jsii.String("ttmlProfile"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashttmlconfiguration.html
//
type CfnOriginEndpointPropsMixin_DashTtmlConfigurationProperty struct {
	// The profile that MediaPackage uses when signaling subtitles in the manifest.
	//
	// `IMSC` is the default profile. `EBU-TT-D` produces subtitles that are compliant with the EBU-TT-D TTML profile. MediaPackage passes through subtitle styles to the manifest. For more information about EBU-TT-D subtitles, see [EBU-TT-D Subtitling Distribution Format](https://docs.aws.amazon.com/https://tech.ebu.ch/publications/tech3380) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashttmlconfiguration.html#cfn-mediapackagev2-originendpoint-dashttmlconfiguration-ttmlprofile
	//
	TtmlProfile *string `field:"optional" json:"ttmlProfile" yaml:"ttmlProfile"`
}

