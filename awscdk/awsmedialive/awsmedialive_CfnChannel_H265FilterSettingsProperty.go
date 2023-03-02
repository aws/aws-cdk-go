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
//   h265FilterSettingsProperty := &h265FilterSettingsProperty{
//   	temporalFilterSettings: &temporalFilterSettingsProperty{
//   		postFilterSharpening: jsii.String("postFilterSharpening"),
//   		strength: jsii.String("strength"),
//   	},
//   }
//
type CfnChannel_H265FilterSettingsProperty struct {
	// Settings for applying the temporal filter to the video.
	TemporalFilterSettings interface{} `field:"optional" json:"temporalFilterSettings" yaml:"temporalFilterSettings"`
}

