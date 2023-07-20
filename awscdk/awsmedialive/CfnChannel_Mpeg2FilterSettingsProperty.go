package awsmedialive


// Settings to configure video filters that apply to the MPEG-2 codec.
//
// The parent of this entity is Mpeg2FilterSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mpeg2FilterSettingsProperty := &Mpeg2FilterSettingsProperty{
//   	TemporalFilterSettings: &TemporalFilterSettingsProperty{
//   		PostFilterSharpening: jsii.String("postFilterSharpening"),
//   		Strength: jsii.String("strength"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mpeg2filtersettings.html
//
type CfnChannel_Mpeg2FilterSettingsProperty struct {
	// Settings for applying the temporal filter to the video.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mpeg2filtersettings.html#cfn-medialive-channel-mpeg2filtersettings-temporalfiltersettings
	//
	TemporalFilterSettings interface{} `field:"optional" json:"temporalFilterSettings" yaml:"temporalFilterSettings"`
}

