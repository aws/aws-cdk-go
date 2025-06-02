package awsquicksight


// The value label configuration of the label in a reference line.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   referenceLineValueLabelConfigurationProperty := &ReferenceLineValueLabelConfigurationProperty{
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
//   	RelativePosition: jsii.String("relativePosition"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-referencelinevaluelabelconfiguration.html
//
type CfnDashboard_ReferenceLineValueLabelConfigurationProperty struct {
	// The format configuration of the value label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-referencelinevaluelabelconfiguration.html#cfn-quicksight-dashboard-referencelinevaluelabelconfiguration-formatconfiguration
	//
	FormatConfiguration interface{} `field:"optional" json:"formatConfiguration" yaml:"formatConfiguration"`
	// The relative position of the value label. Choose one of the following options:.
	//
	// - `BEFORE_CUSTOM_LABEL`
	// - `AFTER_CUSTOM_LABEL`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-referencelinevaluelabelconfiguration.html#cfn-quicksight-dashboard-referencelinevaluelabelconfiguration-relativeposition
	//
	RelativePosition *string `field:"optional" json:"relativePosition" yaml:"relativePosition"`
}

