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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gaugechartoptions.html
//
type CfnDashboard_GaugeChartOptionsProperty struct {
	// The arc configuration of a `GaugeChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gaugechartoptions.html#cfn-quicksight-dashboard-gaugechartoptions-arc
	//
	Arc interface{} `field:"optional" json:"arc" yaml:"arc"`
	// The arc axis configuration of a `GaugeChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gaugechartoptions.html#cfn-quicksight-dashboard-gaugechartoptions-arcaxis
	//
	ArcAxis interface{} `field:"optional" json:"arcAxis" yaml:"arcAxis"`
	// The comparison configuration of a `GaugeChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gaugechartoptions.html#cfn-quicksight-dashboard-gaugechartoptions-comparison
	//
	Comparison interface{} `field:"optional" json:"comparison" yaml:"comparison"`
	// The options that determine the primary value display type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gaugechartoptions.html#cfn-quicksight-dashboard-gaugechartoptions-primaryvaluedisplaytype
	//
	PrimaryValueDisplayType *string `field:"optional" json:"primaryValueDisplayType" yaml:"primaryValueDisplayType"`
	// The options that determine the primary value font configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gaugechartoptions.html#cfn-quicksight-dashboard-gaugechartoptions-primaryvaluefontconfiguration
	//
	PrimaryValueFontConfiguration interface{} `field:"optional" json:"primaryValueFontConfiguration" yaml:"primaryValueFontConfiguration"`
}

