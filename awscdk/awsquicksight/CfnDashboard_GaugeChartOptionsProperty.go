package awsquicksight


// The options that determine the presentation of the `GaugeChartVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gaugeChartOptionsProperty := &GaugeChartOptionsProperty{
//   	Arc: &ArcConfigurationProperty{
//   		ArcAngle: jsii.Number(123),
//   		ArcThickness: jsii.String("arcThickness"),
//   	},
//   	ArcAxis: &ArcAxisConfigurationProperty{
//   		Range: &ArcAxisDisplayRangeProperty{
//   			Max: jsii.Number(123),
//   			Min: jsii.Number(123),
//   		},
//   		ReserveRange: jsii.Number(123),
//   	},
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
//   }
//
type CfnDashboard_GaugeChartOptionsProperty struct {
	// The arc configuration of a `GaugeChartVisual` .
	Arc interface{} `field:"optional" json:"arc" yaml:"arc"`
	// The arc axis configuration of a `GaugeChartVisual` .
	ArcAxis interface{} `field:"optional" json:"arcAxis" yaml:"arcAxis"`
	// The comparison configuration of a `GaugeChartVisual` .
	Comparison interface{} `field:"optional" json:"comparison" yaml:"comparison"`
	// The options that determine the primary value display type.
	PrimaryValueDisplayType *string `field:"optional" json:"primaryValueDisplayType" yaml:"primaryValueDisplayType"`
	// The options that determine the primary value font configuration.
	PrimaryValueFontConfiguration interface{} `field:"optional" json:"primaryValueFontConfiguration" yaml:"primaryValueFontConfiguration"`
}

