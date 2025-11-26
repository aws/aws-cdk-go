package previewawsmedialivemixins


// Settings to configure video filters that apply to the H264 codec.
//
// The parent of this entity is H264Settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   h264FilterSettingsProperty := &H264FilterSettingsProperty{
//   	BandwidthReductionFilterSettings: &BandwidthReductionFilterSettingsProperty{
//   		PostFilterSharpening: jsii.String("postFilterSharpening"),
//   		Strength: jsii.String("strength"),
//   	},
//   	TemporalFilterSettings: &TemporalFilterSettingsProperty{
//   		PostFilterSharpening: jsii.String("postFilterSharpening"),
//   		Strength: jsii.String("strength"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264filtersettings.html
//
type CfnChannelPropsMixin_H264FilterSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264filtersettings.html#cfn-medialive-channel-h264filtersettings-bandwidthreductionfiltersettings
	//
	BandwidthReductionFilterSettings interface{} `field:"optional" json:"bandwidthReductionFilterSettings" yaml:"bandwidthReductionFilterSettings"`
	// Settings for applying the temporal filter to the video.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264filtersettings.html#cfn-medialive-channel-h264filtersettings-temporalfiltersettings
	//
	TemporalFilterSettings interface{} `field:"optional" json:"temporalFilterSettings" yaml:"temporalFilterSettings"`
}

