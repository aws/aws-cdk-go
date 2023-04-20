package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   percentageDisplayFormatConfigurationProperty := &PercentageDisplayFormatConfigurationProperty{
//   	DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   		DecimalPlaces: jsii.Number(123),
//   	},
//   	NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   		DisplayMode: jsii.String("displayMode"),
//   	},
//   	NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   		NullString: jsii.String("nullString"),
//   	},
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
type CfnTemplate_PercentageDisplayFormatConfigurationProperty struct {
	// `CfnTemplate.PercentageDisplayFormatConfigurationProperty.DecimalPlacesConfiguration`.
	DecimalPlacesConfiguration interface{} `field:"optional" json:"decimalPlacesConfiguration" yaml:"decimalPlacesConfiguration"`
	// `CfnTemplate.PercentageDisplayFormatConfigurationProperty.NegativeValueConfiguration`.
	NegativeValueConfiguration interface{} `field:"optional" json:"negativeValueConfiguration" yaml:"negativeValueConfiguration"`
	// `CfnTemplate.PercentageDisplayFormatConfigurationProperty.NullValueFormatConfiguration`.
	NullValueFormatConfiguration interface{} `field:"optional" json:"nullValueFormatConfiguration" yaml:"nullValueFormatConfiguration"`
	// `CfnTemplate.PercentageDisplayFormatConfigurationProperty.Prefix`.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// `CfnTemplate.PercentageDisplayFormatConfigurationProperty.SeparatorConfiguration`.
	SeparatorConfiguration interface{} `field:"optional" json:"separatorConfiguration" yaml:"separatorConfiguration"`
	// `CfnTemplate.PercentageDisplayFormatConfigurationProperty.Suffix`.
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

