package awsquicksight


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
type CfnDashboard_ComparisonConfigurationProperty struct {
	// `CfnDashboard.ComparisonConfigurationProperty.ComparisonFormat`.
	ComparisonFormat interface{} `field:"optional" json:"comparisonFormat" yaml:"comparisonFormat"`
	// `CfnDashboard.ComparisonConfigurationProperty.ComparisonMethod`.
	ComparisonMethod *string `field:"optional" json:"comparisonMethod" yaml:"comparisonMethod"`
}

