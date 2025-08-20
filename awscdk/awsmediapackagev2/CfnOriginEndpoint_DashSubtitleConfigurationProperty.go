package awsmediapackagev2


// The configuration for DASH subtitles.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashSubtitleConfigurationProperty := &DashSubtitleConfigurationProperty{
//   	TtmlConfiguration: &DashTtmlConfigurationProperty{
//   		TtmlProfile: jsii.String("ttmlProfile"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashsubtitleconfiguration.html
//
type CfnOriginEndpoint_DashSubtitleConfigurationProperty struct {
	// Settings for TTML subtitles.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashsubtitleconfiguration.html#cfn-mediapackagev2-originendpoint-dashsubtitleconfiguration-ttmlconfiguration
	//
	TtmlConfiguration interface{} `field:"optional" json:"ttmlConfiguration" yaml:"ttmlConfiguration"`
}

