package awsquicksight


// Formatting configuration for number fields.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   numberFormatConfigurationProperty := &NumberFormatConfigurationProperty{
//   	FormatConfiguration: &NumericFormatConfigurationProperty{
//   		CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
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
//   			Symbol: jsii.String("symbol"),
//   		},
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-numberformatconfiguration.html
//
type CfnAnalysis_NumberFormatConfigurationProperty struct {
	// The options that determine the numeric format configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-numberformatconfiguration.html#cfn-quicksight-analysis-numberformatconfiguration-formatconfiguration
	//
	FormatConfiguration interface{} `field:"optional" json:"formatConfiguration" yaml:"formatConfiguration"`
}

