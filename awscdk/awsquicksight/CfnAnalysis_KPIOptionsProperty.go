package awsquicksight


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
//   			Relative: jsii.String("relative"),
//   		},
//   		FontStyle: jsii.String("fontStyle"),
//   		FontWeight: &FontWeightProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	TrendArrows: &TrendArrowOptionsProperty{
//   		Visibility: jsii.String("visibility"),
//   	},
//   }
//
type CfnAnalysis_KPIOptionsProperty struct {
	// `CfnAnalysis.KPIOptionsProperty.Comparison`.
	Comparison interface{} `field:"optional" json:"comparison" yaml:"comparison"`
	// `CfnAnalysis.KPIOptionsProperty.PrimaryValueDisplayType`.
	PrimaryValueDisplayType *string `field:"optional" json:"primaryValueDisplayType" yaml:"primaryValueDisplayType"`
	// `CfnAnalysis.KPIOptionsProperty.PrimaryValueFontConfiguration`.
	PrimaryValueFontConfiguration interface{} `field:"optional" json:"primaryValueFontConfiguration" yaml:"primaryValueFontConfiguration"`
	// `CfnAnalysis.KPIOptionsProperty.ProgressBar`.
	ProgressBar interface{} `field:"optional" json:"progressBar" yaml:"progressBar"`
	// `CfnAnalysis.KPIOptionsProperty.SecondaryValue`.
	SecondaryValue interface{} `field:"optional" json:"secondaryValue" yaml:"secondaryValue"`
	// `CfnAnalysis.KPIOptionsProperty.SecondaryValueFontConfiguration`.
	SecondaryValueFontConfiguration interface{} `field:"optional" json:"secondaryValueFontConfiguration" yaml:"secondaryValueFontConfiguration"`
	// `CfnAnalysis.KPIOptionsProperty.TrendArrows`.
	TrendArrows interface{} `field:"optional" json:"trendArrows" yaml:"trendArrows"`
}

