package awsquicksight


// Determines the typography options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   typographyProperty := &TypographyProperty{
//   	AxisLabelFontConfiguration: &FontConfigurationProperty{
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
//   	AxisTitleFontConfiguration: &FontConfigurationProperty{
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
//   	DataLabelFontConfiguration: &FontConfigurationProperty{
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
//   	FontFamilies: []interface{}{
//   		&FontProperty{
//   			FontFamily: jsii.String("fontFamily"),
//   		},
//   	},
//   	LegendTitleFontConfiguration: &FontConfigurationProperty{
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
//   	LegendValueFontConfiguration: &FontConfigurationProperty{
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
//   	VisualSubtitleFontConfiguration: &VisualSubtitleFontConfigurationProperty{
//   		FontConfiguration: &FontConfigurationProperty{
//   			FontColor: jsii.String("fontColor"),
//   			FontDecoration: jsii.String("fontDecoration"),
//   			FontFamily: jsii.String("fontFamily"),
//   			FontSize: &FontSizeProperty{
//   				Absolute: jsii.String("absolute"),
//   			},
//   			FontStyle: jsii.String("fontStyle"),
//   			FontWeight: &FontWeightProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		TextAlignment: jsii.String("textAlignment"),
//   		TextTransform: jsii.String("textTransform"),
//   	},
//   	VisualTitleFontConfiguration: &VisualTitleFontConfigurationProperty{
//   		FontConfiguration: &FontConfigurationProperty{
//   			FontColor: jsii.String("fontColor"),
//   			FontDecoration: jsii.String("fontDecoration"),
//   			FontFamily: jsii.String("fontFamily"),
//   			FontSize: &FontSizeProperty{
//   				Absolute: jsii.String("absolute"),
//   			},
//   			FontStyle: jsii.String("fontStyle"),
//   			FontWeight: &FontWeightProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		TextAlignment: jsii.String("textAlignment"),
//   		TextTransform: jsii.String("textTransform"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-typography.html
//
type CfnTheme_TypographyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-typography.html#cfn-quicksight-theme-typography-axislabelfontconfiguration
	//
	AxisLabelFontConfiguration interface{} `field:"optional" json:"axisLabelFontConfiguration" yaml:"axisLabelFontConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-typography.html#cfn-quicksight-theme-typography-axistitlefontconfiguration
	//
	AxisTitleFontConfiguration interface{} `field:"optional" json:"axisTitleFontConfiguration" yaml:"axisTitleFontConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-typography.html#cfn-quicksight-theme-typography-datalabelfontconfiguration
	//
	DataLabelFontConfiguration interface{} `field:"optional" json:"dataLabelFontConfiguration" yaml:"dataLabelFontConfiguration"`
	// Determines the list of font families.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-typography.html#cfn-quicksight-theme-typography-fontfamilies
	//
	FontFamilies interface{} `field:"optional" json:"fontFamilies" yaml:"fontFamilies"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-typography.html#cfn-quicksight-theme-typography-legendtitlefontconfiguration
	//
	LegendTitleFontConfiguration interface{} `field:"optional" json:"legendTitleFontConfiguration" yaml:"legendTitleFontConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-typography.html#cfn-quicksight-theme-typography-legendvaluefontconfiguration
	//
	LegendValueFontConfiguration interface{} `field:"optional" json:"legendValueFontConfiguration" yaml:"legendValueFontConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-typography.html#cfn-quicksight-theme-typography-visualsubtitlefontconfiguration
	//
	VisualSubtitleFontConfiguration interface{} `field:"optional" json:"visualSubtitleFontConfiguration" yaml:"visualSubtitleFontConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-typography.html#cfn-quicksight-theme-typography-visualtitlefontconfiguration
	//
	VisualTitleFontConfiguration interface{} `field:"optional" json:"visualTitleFontConfiguration" yaml:"visualTitleFontConfiguration"`
}

