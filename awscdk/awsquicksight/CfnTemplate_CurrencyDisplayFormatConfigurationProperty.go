package awsquicksight


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
type CfnTemplate_CurrencyDisplayFormatConfigurationProperty struct {
	// `CfnTemplate.CurrencyDisplayFormatConfigurationProperty.DecimalPlacesConfiguration`.
	DecimalPlacesConfiguration interface{} `field:"optional" json:"decimalPlacesConfiguration" yaml:"decimalPlacesConfiguration"`
	// `CfnTemplate.CurrencyDisplayFormatConfigurationProperty.NegativeValueConfiguration`.
	NegativeValueConfiguration interface{} `field:"optional" json:"negativeValueConfiguration" yaml:"negativeValueConfiguration"`
	// `CfnTemplate.CurrencyDisplayFormatConfigurationProperty.NullValueFormatConfiguration`.
	NullValueFormatConfiguration interface{} `field:"optional" json:"nullValueFormatConfiguration" yaml:"nullValueFormatConfiguration"`
	// `CfnTemplate.CurrencyDisplayFormatConfigurationProperty.NumberScale`.
	NumberScale *string `field:"optional" json:"numberScale" yaml:"numberScale"`
	// `CfnTemplate.CurrencyDisplayFormatConfigurationProperty.Prefix`.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// `CfnTemplate.CurrencyDisplayFormatConfigurationProperty.SeparatorConfiguration`.
	SeparatorConfiguration interface{} `field:"optional" json:"separatorConfiguration" yaml:"separatorConfiguration"`
	// `CfnTemplate.CurrencyDisplayFormatConfigurationProperty.Suffix`.
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
	// `CfnTemplate.CurrencyDisplayFormatConfigurationProperty.Symbol`.
	Symbol *string `field:"optional" json:"symbol" yaml:"symbol"`
}

