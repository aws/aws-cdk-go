package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   numberDisplayFormatConfigurationProperty := &NumberDisplayFormatConfigurationProperty{
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
//   }
//
type CfnTemplate_NumberDisplayFormatConfigurationProperty struct {
	// `CfnTemplate.NumberDisplayFormatConfigurationProperty.DecimalPlacesConfiguration`.
	DecimalPlacesConfiguration interface{} `field:"optional" json:"decimalPlacesConfiguration" yaml:"decimalPlacesConfiguration"`
	// `CfnTemplate.NumberDisplayFormatConfigurationProperty.NegativeValueConfiguration`.
	NegativeValueConfiguration interface{} `field:"optional" json:"negativeValueConfiguration" yaml:"negativeValueConfiguration"`
	// `CfnTemplate.NumberDisplayFormatConfigurationProperty.NullValueFormatConfiguration`.
	NullValueFormatConfiguration interface{} `field:"optional" json:"nullValueFormatConfiguration" yaml:"nullValueFormatConfiguration"`
	// `CfnTemplate.NumberDisplayFormatConfigurationProperty.NumberScale`.
	NumberScale *string `field:"optional" json:"numberScale" yaml:"numberScale"`
	// `CfnTemplate.NumberDisplayFormatConfigurationProperty.Prefix`.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// `CfnTemplate.NumberDisplayFormatConfigurationProperty.SeparatorConfiguration`.
	SeparatorConfiguration interface{} `field:"optional" json:"separatorConfiguration" yaml:"separatorConfiguration"`
	// `CfnTemplate.NumberDisplayFormatConfigurationProperty.Suffix`.
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

