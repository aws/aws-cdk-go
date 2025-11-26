package previewawsquicksightmixins


// The options that determine the presentation of a KPI visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kPIOptionsProperty := &KPIOptionsProperty{
//   	Comparison: &ComparisonConfigurationProperty{
//   		ComparisonFormat: &ComparisonFormatConfigurationProperty{
//   			NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   				DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   					DecimalPlaces: jsii.Number(123),
//   				},
//   				NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   					DisplayMode: jsii.String("displayMode"),
//   				},
//   				NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   					NullString: jsii.String("nullString"),
//   				},
//   				NumberScale: jsii.String("numberScale"),
//   				Prefix: jsii.String("prefix"),
//   				SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   					DecimalSeparator: jsii.String("decimalSeparator"),
//   					ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   						GroupingStyle: jsii.String("groupingStyle"),
//   						Symbol: jsii.String("symbol"),
//   						Visibility: jsii.String("visibility"),
//   					},
//   				},
//   				Suffix: jsii.String("suffix"),
//   			},
//   			PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   				DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   					DecimalPlaces: jsii.Number(123),
//   				},
//   				NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   					DisplayMode: jsii.String("displayMode"),
//   				},
//   				NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   					NullString: jsii.String("nullString"),
//   				},
//   				Prefix: jsii.String("prefix"),
//   				SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   					DecimalSeparator: jsii.String("decimalSeparator"),
//   					ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   						GroupingStyle: jsii.String("groupingStyle"),
//   						Symbol: jsii.String("symbol"),
//   						Visibility: jsii.String("visibility"),
//   					},
//   				},
//   				Suffix: jsii.String("suffix"),
//   			},
//   		},
//   		ComparisonMethod: jsii.String("comparisonMethod"),
//   	},
//   	PrimaryValueDisplayType: jsii.String("primaryValueDisplayType"),
//   	PrimaryValueFontConfiguration: &FontConfigurationProperty{
//   		FontColor: jsii.String("fontColor"),
//   		FontDecoration: jsii.String("fontDecoration"),
//   		FontFamily: jsii.String("fontFamily"),
//   		FontSize: &FontSizeProperty{
//   			Absolute: jsii.String("absolute"),
//   			Relative: jsii.String("relative"),
//   		},
//   		FontStyle: jsii.String("fontStyle"),
//   		FontWeight: &FontWeightProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	ProgressBar: &ProgressBarOptionsProperty{
//   		Visibility: jsii.String("visibility"),
//   	},
//   	SecondaryValue: &SecondaryValueOptionsProperty{
//   		Visibility: jsii.String("visibility"),
//   	},
//   	SecondaryValueFontConfiguration: &FontConfigurationProperty{
//   		FontColor: jsii.String("fontColor"),
//   		FontDecoration: jsii.String("fontDecoration"),
//   		FontFamily: jsii.String("fontFamily"),
//   		FontSize: &FontSizeProperty{
//   			Absolute: jsii.String("absolute"),
//   			Relative: jsii.String("relative"),
//   		},
//   		FontStyle: jsii.String("fontStyle"),
//   		FontWeight: &FontWeightProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Sparkline: &KPISparklineOptionsProperty{
//   		Color: jsii.String("color"),
//   		TooltipVisibility: jsii.String("tooltipVisibility"),
//   		Type: jsii.String("type"),
//   		Visibility: jsii.String("visibility"),
//   	},
//   	TrendArrows: &TrendArrowOptionsProperty{
//   		Visibility: jsii.String("visibility"),
//   	},
//   	VisualLayoutOptions: &KPIVisualLayoutOptionsProperty{
//   		StandardLayout: &KPIVisualStandardLayoutProperty{
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpioptions.html
//
type CfnTemplatePropsMixin_KPIOptionsProperty struct {
	// The comparison configuration of a KPI visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpioptions.html#cfn-quicksight-template-kpioptions-comparison
	//
	Comparison interface{} `field:"optional" json:"comparison" yaml:"comparison"`
	// The options that determine the primary value display type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpioptions.html#cfn-quicksight-template-kpioptions-primaryvaluedisplaytype
	//
	PrimaryValueDisplayType *string `field:"optional" json:"primaryValueDisplayType" yaml:"primaryValueDisplayType"`
	// The options that determine the primary value font configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpioptions.html#cfn-quicksight-template-kpioptions-primaryvaluefontconfiguration
	//
	PrimaryValueFontConfiguration interface{} `field:"optional" json:"primaryValueFontConfiguration" yaml:"primaryValueFontConfiguration"`
	// The options that determine the presentation of the progress bar of a KPI visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpioptions.html#cfn-quicksight-template-kpioptions-progressbar
	//
	ProgressBar interface{} `field:"optional" json:"progressBar" yaml:"progressBar"`
	// The options that determine the presentation of the secondary value of a KPI visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpioptions.html#cfn-quicksight-template-kpioptions-secondaryvalue
	//
	SecondaryValue interface{} `field:"optional" json:"secondaryValue" yaml:"secondaryValue"`
	// The options that determine the secondary value font configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpioptions.html#cfn-quicksight-template-kpioptions-secondaryvaluefontconfiguration
	//
	SecondaryValueFontConfiguration interface{} `field:"optional" json:"secondaryValueFontConfiguration" yaml:"secondaryValueFontConfiguration"`
	// The options that determine the visibility, color, type, and tooltip visibility of the sparkline of a KPI visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpioptions.html#cfn-quicksight-template-kpioptions-sparkline
	//
	Sparkline interface{} `field:"optional" json:"sparkline" yaml:"sparkline"`
	// The options that determine the presentation of trend arrows in a KPI visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpioptions.html#cfn-quicksight-template-kpioptions-trendarrows
	//
	TrendArrows interface{} `field:"optional" json:"trendArrows" yaml:"trendArrows"`
	// The options that determine the layout a KPI visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpioptions.html#cfn-quicksight-template-kpioptions-visuallayoutoptions
	//
	VisualLayoutOptions interface{} `field:"optional" json:"visualLayoutOptions" yaml:"visualLayoutOptions"`
}

