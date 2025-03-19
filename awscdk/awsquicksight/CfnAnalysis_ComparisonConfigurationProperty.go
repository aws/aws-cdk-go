package awsquicksight


// The comparison display configuration of a KPI or gauge chart.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   comparisonConfigurationProperty := &ComparisonConfigurationProperty{
//   	ComparisonFormat: &ComparisonFormatConfigurationProperty{
//   		NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   			DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   				DecimalPlaces: jsii.Number(123),
//   			},
//   			NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   				DisplayMode: jsii.String("displayMode"),
//   			},
//   			NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   				NullString: jsii.String("nullString"),
//   			},
//   			NumberScale: jsii.String("numberScale"),
//   			Prefix: jsii.String("prefix"),
//   			SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   				DecimalSeparator: jsii.String("decimalSeparator"),
//   				ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   					GroupingStyle: jsii.String("groupingStyle"),
//   					Symbol: jsii.String("symbol"),
//   					Visibility: jsii.String("visibility"),
//   				},
//   			},
//   			Suffix: jsii.String("suffix"),
//   		},
//   		PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   			DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   				DecimalPlaces: jsii.Number(123),
//   			},
//   			NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   				DisplayMode: jsii.String("displayMode"),
//   			},
//   			NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   				NullString: jsii.String("nullString"),
//   			},
//   			Prefix: jsii.String("prefix"),
//   			SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   				DecimalSeparator: jsii.String("decimalSeparator"),
//   				ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   					GroupingStyle: jsii.String("groupingStyle"),
//   					Symbol: jsii.String("symbol"),
//   					Visibility: jsii.String("visibility"),
//   				},
//   			},
//   			Suffix: jsii.String("suffix"),
//   		},
//   	},
//   	ComparisonMethod: jsii.String("comparisonMethod"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-comparisonconfiguration.html
//
type CfnAnalysis_ComparisonConfigurationProperty struct {
	// The format of the comparison.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-comparisonconfiguration.html#cfn-quicksight-analysis-comparisonconfiguration-comparisonformat
	//
	ComparisonFormat interface{} `field:"optional" json:"comparisonFormat" yaml:"comparisonFormat"`
	// The method of the comparison. Choose from the following options:.
	//
	// - `DIFFERENCE`
	// - `PERCENT_DIFFERENCE`
	// - `PERCENT`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-comparisonconfiguration.html#cfn-quicksight-analysis-comparisonconfiguration-comparisonmethod
	//
	ComparisonMethod *string `field:"optional" json:"comparisonMethod" yaml:"comparisonMethod"`
}

