package awsquicksight


// The format of the comparison.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   comparisonFormatConfigurationProperty := &ComparisonFormatConfigurationProperty{
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
type CfnTemplate_ComparisonFormatConfigurationProperty struct {
	// The number display format.
	NumberDisplayFormatConfiguration interface{} `field:"optional" json:"numberDisplayFormatConfiguration" yaml:"numberDisplayFormatConfiguration"`
	// The percentage display format.
	PercentageDisplayFormatConfiguration interface{} `field:"optional" json:"percentageDisplayFormatConfiguration" yaml:"percentageDisplayFormatConfiguration"`
}

