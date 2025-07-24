package awsmedialive


// Settings to configure video filters that apply to the H265 codec.
//
// The parent of this entity is H265Settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   h265FilterSettingsProperty := &H265FilterSettingsProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h265filtersettings.html
//
type CfnChannel_H265FilterSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h265filtersettings.html#cfn-medialive-channel-h265filtersettings-bandwidthreductionfiltersettings
	//
	BandwidthReductionFilterSettings interface{} `field:"optional" json:"bandwidthReductionFilterSettings" yaml:"bandwidthReductionFilterSettings"`
	// Settings for applying the temporal filter to the video.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h265filtersettings.html#cfn-medialive-channel-h265filtersettings-temporalfiltersettings
	//
	TemporalFilterSettings interface{} `field:"optional" json:"temporalFilterSettings" yaml:"temporalFilterSettings"`
}

