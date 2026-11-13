package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var autoFilterSettings interface{}
//   var noneFilterSettings interface{}
//
//   av1FilterSettingsProperty := &Av1FilterSettingsProperty{
//   	AutoFilterSettings: autoFilterSettings,
//   	NoneFilterSettings: noneFilterSettings,
//   	PerceptualFilterSettings: &PerceptualFilterSettingsProperty{
//   		Strength: jsii.String("strength"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1filtersettings.html
//
type CfnChannelPropsMixin_Av1FilterSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1filtersettings.html#cfn-medialive-channel-av1filtersettings-autofiltersettings
	//
	AutoFilterSettings interface{} `field:"optional" json:"autoFilterSettings" yaml:"autoFilterSettings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1filtersettings.html#cfn-medialive-channel-av1filtersettings-nonefiltersettings
	//
	NoneFilterSettings interface{} `field:"optional" json:"noneFilterSettings" yaml:"noneFilterSettings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1filtersettings.html#cfn-medialive-channel-av1filtersettings-perceptualfiltersettings
	//
	PerceptualFilterSettings interface{} `field:"optional" json:"perceptualFilterSettings" yaml:"perceptualFilterSettings"`
}

