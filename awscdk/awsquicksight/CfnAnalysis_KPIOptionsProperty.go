package awsquicksight


// The options that determine the presentation of a KPI visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		Color: jsii.String("color"),
//   		TooltipVisibility: jsii.String("tooltipVisibility"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-kpioptions.html
//
type CfnAnalysis_KPIOptionsProperty struct {
	// The comparison configuration of a KPI visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-kpioptions.html#cfn-quicksight-analysis-kpioptions-comparison
	//
	Comparison interface{} `field:"optional" json:"comparison" yaml:"comparison"`
	// The options that determine the primary value display type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-kpioptions.html#cfn-quicksight-analysis-kpioptions-primaryvaluedisplaytype
	//
	PrimaryValueDisplayType *string `field:"optional" json:"primaryValueDisplayType" yaml:"primaryValueDisplayType"`
	// The options that determine the primary value font configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-kpioptions.html#cfn-quicksight-analysis-kpioptions-primaryvaluefontconfiguration
	//
	PrimaryValueFontConfiguration interface{} `field:"optional" json:"primaryValueFontConfiguration" yaml:"primaryValueFontConfiguration"`
	// The options that determine the presentation of the progress bar of a KPI visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-kpioptions.html#cfn-quicksight-analysis-kpioptions-progressbar
	//
	ProgressBar interface{} `field:"optional" json:"progressBar" yaml:"progressBar"`
	// The options that determine the presentation of the secondary value of a KPI visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-kpioptions.html#cfn-quicksight-analysis-kpioptions-secondaryvalue
	//
	SecondaryValue interface{} `field:"optional" json:"secondaryValue" yaml:"secondaryValue"`
	// The options that determine the secondary value font configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-kpioptions.html#cfn-quicksight-analysis-kpioptions-secondaryvaluefontconfiguration
	//
	SecondaryValueFontConfiguration interface{} `field:"optional" json:"secondaryValueFontConfiguration" yaml:"secondaryValueFontConfiguration"`
	// The options that determine the visibility, color, type, and tooltip visibility of the sparkline of a KPI visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-kpioptions.html#cfn-quicksight-analysis-kpioptions-sparkline
	//
	Sparkline interface{} `field:"optional" json:"sparkline" yaml:"sparkline"`
	// The options that determine the presentation of trend arrows in a KPI visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-kpioptions.html#cfn-quicksight-analysis-kpioptions-trendarrows
	//
	TrendArrows interface{} `field:"optional" json:"trendArrows" yaml:"trendArrows"`
	// The options that determine the layout a KPI visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-kpioptions.html#cfn-quicksight-analysis-kpioptions-visuallayoutoptions
	//
	VisualLayoutOptions interface{} `field:"optional" json:"visualLayoutOptions" yaml:"visualLayoutOptions"`
}

