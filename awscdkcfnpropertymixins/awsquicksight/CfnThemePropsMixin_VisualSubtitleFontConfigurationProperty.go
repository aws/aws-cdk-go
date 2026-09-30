package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   visualSubtitleFontConfigurationProperty := &VisualSubtitleFontConfigurationProperty{
//   	FontConfiguration: &FontConfigurationProperty{
//   		FontColor: jsii.String("fontColor"),
//   		FontDecoration: jsii.String("fontDecoration"),
//   		FontFamily: jsii.String("fontFamily"),
//   		FontSize: &FontSizeProperty{
//   			Absolute: jsii.String("absolute"),
//   		},
//   		FontStyle: jsii.String("fontStyle"),
//   		FontWeight: &FontWeightProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	TextAlignment: jsii.String("textAlignment"),
//   	TextTransform: jsii.String("textTransform"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-visualsubtitlefontconfiguration.html
//
type CfnThemePropsMixin_VisualSubtitleFontConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-visualsubtitlefontconfiguration.html#cfn-quicksight-theme-visualsubtitlefontconfiguration-fontconfiguration
	//
	FontConfiguration interface{} `field:"optional" json:"fontConfiguration" yaml:"fontConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-visualsubtitlefontconfiguration.html#cfn-quicksight-theme-visualsubtitlefontconfiguration-textalignment
	//
	TextAlignment *string `field:"optional" json:"textAlignment" yaml:"textAlignment"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-visualsubtitlefontconfiguration.html#cfn-quicksight-theme-visualsubtitlefontconfiguration-texttransform
	//
	TextTransform *string `field:"optional" json:"textTransform" yaml:"textTransform"`
}

