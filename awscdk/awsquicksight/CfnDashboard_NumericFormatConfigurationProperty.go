package awsquicksight


// The options that determine the numeric format configuration.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   numericFormatConfigurationProperty := &NumericFormatConfigurationProperty{
//   	CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   		DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   			DecimalPlaces: jsii.Number(123),
//   		},
//   		NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   			DisplayMode: jsii.String("displayMode"),
//   		},
//   		NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   			NullString: jsii.String("nullString"),
//   		},
//   		NumberScale: jsii.String("numberScale"),
//   		Prefix: jsii.String("prefix"),
//   		SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   			DecimalSeparator: jsii.String("decimalSeparator"),
//   			ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   				Symbol: jsii.String("symbol"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   		},
//   		Suffix: jsii.String("suffix"),
//   		Symbol: jsii.String("symbol"),
//   	},
//   	NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   		DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   			DecimalPlaces: jsii.Number(123),
//   		},
//   		NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   			DisplayMode: jsii.String("displayMode"),
//   		},
//   		NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   			NullString: jsii.String("nullString"),
//   		},
//   		NumberScale: jsii.String("numberScale"),
//   		Prefix: jsii.String("prefix"),
//   		SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   			DecimalSeparator: jsii.String("decimalSeparator"),
//   			ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   				Symbol: jsii.String("symbol"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   		},
//   		Suffix: jsii.String("suffix"),
//   	},
//   	PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   		DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   			DecimalPlaces: jsii.Number(123),
//   		},
//   		NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   			DisplayMode: jsii.String("displayMode"),
//   		},
//   		NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   			NullString: jsii.String("nullString"),
//   		},
//   		Prefix: jsii.String("prefix"),
//   		SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   			DecimalSeparator: jsii.String("decimalSeparator"),
//   			ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   				Symbol: jsii.String("symbol"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   		},
//   		Suffix: jsii.String("suffix"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericformatconfiguration.html
//
type CfnDashboard_NumericFormatConfigurationProperty struct {
	// The options that determine the currency display format configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericformatconfiguration.html#cfn-quicksight-dashboard-numericformatconfiguration-currencydisplayformatconfiguration
	//
	CurrencyDisplayFormatConfiguration interface{} `field:"optional" json:"currencyDisplayFormatConfiguration" yaml:"currencyDisplayFormatConfiguration"`
	// The options that determine the number display format configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericformatconfiguration.html#cfn-quicksight-dashboard-numericformatconfiguration-numberdisplayformatconfiguration
	//
	NumberDisplayFormatConfiguration interface{} `field:"optional" json:"numberDisplayFormatConfiguration" yaml:"numberDisplayFormatConfiguration"`
	// The options that determine the percentage display format configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericformatconfiguration.html#cfn-quicksight-dashboard-numericformatconfiguration-percentagedisplayformatconfiguration
	//
	PercentageDisplayFormatConfiguration interface{} `field:"optional" json:"percentageDisplayFormatConfiguration" yaml:"percentageDisplayFormatConfiguration"`
}

