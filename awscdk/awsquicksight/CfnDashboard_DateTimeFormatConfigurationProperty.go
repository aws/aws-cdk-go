package awsquicksight


// Formatting configuration for `DateTime` fields.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dateTimeFormatConfigurationProperty := &DateTimeFormatConfigurationProperty{
//   	DateTimeFormat: jsii.String("dateTimeFormat"),
//   	NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   		NullString: jsii.String("nullString"),
//   	},
//   	NumericFormatConfiguration: &NumericFormatConfigurationProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datetimeformatconfiguration.html
//
type CfnDashboard_DateTimeFormatConfigurationProperty struct {
	// Determines the `DateTime` format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datetimeformatconfiguration.html#cfn-quicksight-dashboard-datetimeformatconfiguration-datetimeformat
	//
	DateTimeFormat *string `field:"optional" json:"dateTimeFormat" yaml:"dateTimeFormat"`
	// The options that determine the null value format configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datetimeformatconfiguration.html#cfn-quicksight-dashboard-datetimeformatconfiguration-nullvalueformatconfiguration
	//
	NullValueFormatConfiguration interface{} `field:"optional" json:"nullValueFormatConfiguration" yaml:"nullValueFormatConfiguration"`
	// The formatting configuration for numeric `DateTime` fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datetimeformatconfiguration.html#cfn-quicksight-dashboard-datetimeformatconfiguration-numericformatconfiguration
	//
	NumericFormatConfiguration interface{} `field:"optional" json:"numericFormatConfiguration" yaml:"numericFormatConfiguration"`
}

