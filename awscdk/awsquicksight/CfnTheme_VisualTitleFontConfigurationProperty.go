package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   visualTitleFontConfigurationProperty := &VisualTitleFontConfigurationProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-visualtitlefontconfiguration.html
//
type CfnTheme_VisualTitleFontConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-visualtitlefontconfiguration.html#cfn-quicksight-theme-visualtitlefontconfiguration-fontconfiguration
	//
	FontConfiguration interface{} `field:"optional" json:"fontConfiguration" yaml:"fontConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-visualtitlefontconfiguration.html#cfn-quicksight-theme-visualtitlefontconfiguration-textalignment
	//
	TextAlignment *string `field:"optional" json:"textAlignment" yaml:"textAlignment"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-visualtitlefontconfiguration.html#cfn-quicksight-theme-visualtitlefontconfiguration-texttransform
	//
	TextTransform *string `field:"optional" json:"textTransform" yaml:"textTransform"`
}

