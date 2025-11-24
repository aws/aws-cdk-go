package mixinsawsmedialive


// Settings for the temporal filter to apply to the video.
//
// The parents of this entity are H264FilterSettings, H265FilterSettings, and Mpeg2FilterSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   temporalFilterSettingsProperty := &TemporalFilterSettingsProperty{
//   	PostFilterSharpening: jsii.String("postFilterSharpening"),
//   	Strength: jsii.String("strength"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-temporalfiltersettings.html
//
type CfnChannelPropsMixin_TemporalFilterSettingsProperty struct {
	// If you enable this filter, the results are the following: - If the source content is noisy (it contains excessive digital artifacts), the filter cleans up the source.
	//
	// - If the source content is already clean, the filter tends to decrease the bitrate, especially when the rate control mode is QVBR.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-temporalfiltersettings.html#cfn-medialive-channel-temporalfiltersettings-postfiltersharpening
	//
	PostFilterSharpening *string `field:"optional" json:"postFilterSharpening" yaml:"postFilterSharpening"`
	// Choose a filter strength.
	//
	// We recommend a strength of 1 or 2. A higher strength might take out good information, resulting in an image that is overly soft.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-temporalfiltersettings.html#cfn-medialive-channel-temporalfiltersettings-strength
	//
	Strength *string `field:"optional" json:"strength" yaml:"strength"`
}

