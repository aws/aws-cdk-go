package previewawsquicksightmixins


// The options that determine the currency display format configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
//   			GroupingStyle: jsii.String("groupingStyle"),
//   			Symbol: jsii.String("symbol"),
//   			Visibility: jsii.String("visibility"),
//   		},
//   	},
//   	Suffix: jsii.String("suffix"),
//   	Symbol: jsii.String("symbol"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-currencydisplayformatconfiguration.html
//
type CfnDashboardPropsMixin_CurrencyDisplayFormatConfigurationProperty struct {
	// The option that determines the decimal places configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-currencydisplayformatconfiguration.html#cfn-quicksight-dashboard-currencydisplayformatconfiguration-decimalplacesconfiguration
	//
	DecimalPlacesConfiguration interface{} `field:"optional" json:"decimalPlacesConfiguration" yaml:"decimalPlacesConfiguration"`
	// The options that determine the negative value configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-currencydisplayformatconfiguration.html#cfn-quicksight-dashboard-currencydisplayformatconfiguration-negativevalueconfiguration
	//
	NegativeValueConfiguration interface{} `field:"optional" json:"negativeValueConfiguration" yaml:"negativeValueConfiguration"`
	// The options that determine the null value format configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-currencydisplayformatconfiguration.html#cfn-quicksight-dashboard-currencydisplayformatconfiguration-nullvalueformatconfiguration
	//
	NullValueFormatConfiguration interface{} `field:"optional" json:"nullValueFormatConfiguration" yaml:"nullValueFormatConfiguration"`
	// Determines the number scale value for the currency format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-currencydisplayformatconfiguration.html#cfn-quicksight-dashboard-currencydisplayformatconfiguration-numberscale
	//
	NumberScale *string `field:"optional" json:"numberScale" yaml:"numberScale"`
	// Determines the prefix value of the currency format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-currencydisplayformatconfiguration.html#cfn-quicksight-dashboard-currencydisplayformatconfiguration-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The options that determine the numeric separator configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-currencydisplayformatconfiguration.html#cfn-quicksight-dashboard-currencydisplayformatconfiguration-separatorconfiguration
	//
	SeparatorConfiguration interface{} `field:"optional" json:"separatorConfiguration" yaml:"separatorConfiguration"`
	// Determines the suffix value of the currency format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-currencydisplayformatconfiguration.html#cfn-quicksight-dashboard-currencydisplayformatconfiguration-suffix
	//
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
	// Determines the symbol for the currency format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-currencydisplayformatconfiguration.html#cfn-quicksight-dashboard-currencydisplayformatconfiguration-symbol
	//
	Symbol *string `field:"optional" json:"symbol" yaml:"symbol"`
}

