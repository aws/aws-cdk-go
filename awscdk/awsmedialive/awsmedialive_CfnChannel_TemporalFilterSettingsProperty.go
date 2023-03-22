package awsmedialive


// Settings for the temporal filter to apply to the video.
//
// The parents of this entity are H264FilterSettings, H265FilterSettings, and Mpeg2FilterSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   temporalFilterSettingsProperty := &temporalFilterSettingsProperty{
//   	postFilterSharpening: jsii.String("postFilterSharpening"),
//   	strength: jsii.String("strength"),
//   }
//
type CfnChannel_TemporalFilterSettingsProperty struct {
	// If you enable this filter, the results are the following: - If the source content is noisy (it contains excessive digital artifacts), the filter cleans up the source.
	//
	// - If the source content is already clean, the filter tends to decrease the bitrate, especially when the rate control mode is QVBR.
	PostFilterSharpening *string `field:"optional" json:"postFilterSharpening" yaml:"postFilterSharpening"`
	// Choose a filter strength.
	//
	// We recommend a strength of 1 or 2. A higher strength might take out good information, resulting in an image that is overly soft.
	Strength *string `field:"optional" json:"strength" yaml:"strength"`
}

