package awsquicksight


// The options that determine the currency display format configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   currencyDisplayFormatConfigurationProperty := &CurrencyDisplayFormatConfigurationProperty{
//   	DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   		DecimalPlaces: jsii.Number(123),
//   	},
//   	NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   		DisplayMode: jsii.String("displayMode"),
//   	},
//   	NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   		NullString: jsii.String("nullString"),
//   	},
//   	NumberScale: jsii.String("numberScale"),
//   	Prefix: jsii.String("prefix"),
//   	SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   		DecimalSeparator: jsii.String("decimalSeparator"),
//   		ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   			Symbol: jsii.String("symbol"),
//   			Visibility: jsii.String("visibility"),
//   		},
//   	},
//   	Suffix: jsii.String("suffix"),
//   	Symbol: jsii.String("symbol"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-currencydisplayformatconfiguration.html
//
type CfnAnalysis_CurrencyDisplayFormatConfigurationProperty struct {
	// The option that determines the decimal places configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-currencydisplayformatconfiguration.html#cfn-quicksight-analysis-currencydisplayformatconfiguration-decimalplacesconfiguration
	//
	DecimalPlacesConfiguration interface{} `field:"optional" json:"decimalPlacesConfiguration" yaml:"decimalPlacesConfiguration"`
	// The options that determine the negative value configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-currencydisplayformatconfiguration.html#cfn-quicksight-analysis-currencydisplayformatconfiguration-negativevalueconfiguration
	//
	NegativeValueConfiguration interface{} `field:"optional" json:"negativeValueConfiguration" yaml:"negativeValueConfiguration"`
	// The options that determine the null value format configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-currencydisplayformatconfiguration.html#cfn-quicksight-analysis-currencydisplayformatconfiguration-nullvalueformatconfiguration
	//
	NullValueFormatConfiguration interface{} `field:"optional" json:"nullValueFormatConfiguration" yaml:"nullValueFormatConfiguration"`
	// Determines the number scale value for the currency format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-currencydisplayformatconfiguration.html#cfn-quicksight-analysis-currencydisplayformatconfiguration-numberscale
	//
	NumberScale *string `field:"optional" json:"numberScale" yaml:"numberScale"`
	// Determines the prefix value of the currency format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-currencydisplayformatconfiguration.html#cfn-quicksight-analysis-currencydisplayformatconfiguration-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The options that determine the numeric separator configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-currencydisplayformatconfiguration.html#cfn-quicksight-analysis-currencydisplayformatconfiguration-separatorconfiguration
	//
	SeparatorConfiguration interface{} `field:"optional" json:"separatorConfiguration" yaml:"separatorConfiguration"`
	// Determines the suffix value of the currency format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-currencydisplayformatconfiguration.html#cfn-quicksight-analysis-currencydisplayformatconfiguration-suffix
	//
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
	// Determines the symbol for the currency format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-currencydisplayformatconfiguration.html#cfn-quicksight-analysis-currencydisplayformatconfiguration-symbol
	//
	Symbol *string `field:"optional" json:"symbol" yaml:"symbol"`
}

