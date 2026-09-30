package awsquicksight


// The general configuration of a column.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnConfigurationProperty := &ColumnConfigurationProperty{
//   	Column: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//
//   		// the properties below are optional
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		TopicIdentifier: jsii.String("topicIdentifier"),
//   	},
//
//   	// the properties below are optional
//   	ColorsConfiguration: &ColorsConfigurationProperty{
//   		CustomColors: []interface{}{
//   			&CustomColorProperty{
//   				Color: jsii.String("color"),
//
//   				// the properties below are optional
//   				FieldValue: jsii.String("fieldValue"),
//   				SpecialValue: jsii.String("specialValue"),
//   			},
//   		},
//   	},
//   	DecalSettingsConfiguration: &DecalSettingsConfigurationProperty{
//   		CustomDecalSettings: []interface{}{
//   			&DecalSettingsProperty{
//   				DecalColor: jsii.String("decalColor"),
//   				DecalPatternType: jsii.String("decalPatternType"),
//   				DecalStyleType: jsii.String("decalStyleType"),
//   				DecalVisibility: jsii.String("decalVisibility"),
//   				ElementValue: jsii.String("elementValue"),
//   			},
//   		},
//   	},
//   	FormatConfiguration: &FormatConfigurationProperty{
//   		DateTimeFormatConfiguration: &DateTimeFormatConfigurationProperty{
//   			DateTimeFormat: jsii.String("dateTimeFormat"),
//   			NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   				NullString: jsii.String("nullString"),
//   			},
//   			NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   				CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   					DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   						DecimalPlaces: jsii.Number(123),
//   					},
//   					NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   						DisplayMode: jsii.String("displayMode"),
//   					},
//   					NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   						NullString: jsii.String("nullString"),
//   					},
//   					NumberScale: jsii.String("numberScale"),
//   					Prefix: jsii.String("prefix"),
//   					SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   						DecimalSeparator: jsii.String("decimalSeparator"),
//   						ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   							GroupingStyle: jsii.String("groupingStyle"),
//   							Symbol: jsii.String("symbol"),
//   							Visibility: jsii.String("visibility"),
//   						},
//   					},
//   					Suffix: jsii.String("suffix"),
//   					Symbol: jsii.String("symbol"),
//   				},
//   				NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   					DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   						DecimalPlaces: jsii.Number(123),
//   					},
//   					NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   						DisplayMode: jsii.String("displayMode"),
//   					},
//   					NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   						NullString: jsii.String("nullString"),
//   					},
//   					NumberScale: jsii.String("numberScale"),
//   					Prefix: jsii.String("prefix"),
//   					SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   						DecimalSeparator: jsii.String("decimalSeparator"),
//   						ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   							GroupingStyle: jsii.String("groupingStyle"),
//   							Symbol: jsii.String("symbol"),
//   							Visibility: jsii.String("visibility"),
//   						},
//   					},
//   					Suffix: jsii.String("suffix"),
//   				},
//   				PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   					DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   						DecimalPlaces: jsii.Number(123),
//   					},
//   					NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   						DisplayMode: jsii.String("displayMode"),
//   					},
//   					NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   						NullString: jsii.String("nullString"),
//   					},
//   					Prefix: jsii.String("prefix"),
//   					SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   						DecimalSeparator: jsii.String("decimalSeparator"),
//   						ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   							GroupingStyle: jsii.String("groupingStyle"),
//   							Symbol: jsii.String("symbol"),
//   							Visibility: jsii.String("visibility"),
//   						},
//   					},
//   					Suffix: jsii.String("suffix"),
//   				},
//   			},
//   		},
//   		NumberFormatConfiguration: &NumberFormatConfigurationProperty{
//   			FormatConfiguration: &NumericFormatConfigurationProperty{
//   				CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   					DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   						DecimalPlaces: jsii.Number(123),
//   					},
//   					NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   						DisplayMode: jsii.String("displayMode"),
//   					},
//   					NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   						NullString: jsii.String("nullString"),
//   					},
//   					NumberScale: jsii.String("numberScale"),
//   					Prefix: jsii.String("prefix"),
//   					SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   						DecimalSeparator: jsii.String("decimalSeparator"),
//   						ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   							GroupingStyle: jsii.String("groupingStyle"),
//   							Symbol: jsii.String("symbol"),
//   							Visibility: jsii.String("visibility"),
//   						},
//   					},
//   					Suffix: jsii.String("suffix"),
//   					Symbol: jsii.String("symbol"),
//   				},
//   				NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   					DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   						DecimalPlaces: jsii.Number(123),
//   					},
//   					NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   						DisplayMode: jsii.String("displayMode"),
//   					},
//   					NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   						NullString: jsii.String("nullString"),
//   					},
//   					NumberScale: jsii.String("numberScale"),
//   					Prefix: jsii.String("prefix"),
//   					SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   						DecimalSeparator: jsii.String("decimalSeparator"),
//   						ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   							GroupingStyle: jsii.String("groupingStyle"),
//   							Symbol: jsii.String("symbol"),
//   							Visibility: jsii.String("visibility"),
//   						},
//   					},
//   					Suffix: jsii.String("suffix"),
//   				},
//   				PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   					DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   						DecimalPlaces: jsii.Number(123),
//   					},
//   					NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   						DisplayMode: jsii.String("displayMode"),
//   					},
//   					NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   						NullString: jsii.String("nullString"),
//   					},
//   					Prefix: jsii.String("prefix"),
//   					SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   						DecimalSeparator: jsii.String("decimalSeparator"),
//   						ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   							GroupingStyle: jsii.String("groupingStyle"),
//   							Symbol: jsii.String("symbol"),
//   							Visibility: jsii.String("visibility"),
//   						},
//   					},
//   					Suffix: jsii.String("suffix"),
//   				},
//   			},
//   		},
//   		StringFormatConfiguration: &StringFormatConfigurationProperty{
//   			NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   				NullString: jsii.String("nullString"),
//   			},
//   			NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   				CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   					DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   						DecimalPlaces: jsii.Number(123),
//   					},
//   					NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   						DisplayMode: jsii.String("displayMode"),
//   					},
//   					NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   						NullString: jsii.String("nullString"),
//   					},
//   					NumberScale: jsii.String("numberScale"),
//   					Prefix: jsii.String("prefix"),
//   					SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   						DecimalSeparator: jsii.String("decimalSeparator"),
//   						ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   							GroupingStyle: jsii.String("groupingStyle"),
//   							Symbol: jsii.String("symbol"),
//   							Visibility: jsii.String("visibility"),
//   						},
//   					},
//   					Suffix: jsii.String("suffix"),
//   					Symbol: jsii.String("symbol"),
//   				},
//   				NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   					DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   						DecimalPlaces: jsii.Number(123),
//   					},
//   					NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   						DisplayMode: jsii.String("displayMode"),
//   					},
//   					NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   						NullString: jsii.String("nullString"),
//   					},
//   					NumberScale: jsii.String("numberScale"),
//   					Prefix: jsii.String("prefix"),
//   					SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   						DecimalSeparator: jsii.String("decimalSeparator"),
//   						ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   							GroupingStyle: jsii.String("groupingStyle"),
//   							Symbol: jsii.String("symbol"),
//   							Visibility: jsii.String("visibility"),
//   						},
//   					},
//   					Suffix: jsii.String("suffix"),
//   				},
//   				PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   					DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   						DecimalPlaces: jsii.Number(123),
//   					},
//   					NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   						DisplayMode: jsii.String("displayMode"),
//   					},
//   					NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   						NullString: jsii.String("nullString"),
//   					},
//   					Prefix: jsii.String("prefix"),
//   					SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   						DecimalSeparator: jsii.String("decimalSeparator"),
//   						ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   							GroupingStyle: jsii.String("groupingStyle"),
//   							Symbol: jsii.String("symbol"),
//   							Visibility: jsii.String("visibility"),
//   						},
//   					},
//   					Suffix: jsii.String("suffix"),
//   				},
//   			},
//   		},
//   	},
//   	Role: jsii.String("role"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-columnconfiguration.html
//
type CfnTemplate_ColumnConfigurationProperty struct {
	// The column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-columnconfiguration.html#cfn-quicksight-template-columnconfiguration-column
	//
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// The color configurations of the column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-columnconfiguration.html#cfn-quicksight-template-columnconfiguration-colorsconfiguration
	//
	ColorsConfiguration interface{} `field:"optional" json:"colorsConfiguration" yaml:"colorsConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-columnconfiguration.html#cfn-quicksight-template-columnconfiguration-decalsettingsconfiguration
	//
	DecalSettingsConfiguration interface{} `field:"optional" json:"decalSettingsConfiguration" yaml:"decalSettingsConfiguration"`
	// The format configuration of a column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-columnconfiguration.html#cfn-quicksight-template-columnconfiguration-formatconfiguration
	//
	FormatConfiguration interface{} `field:"optional" json:"formatConfiguration" yaml:"formatConfiguration"`
	// The role of the column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-columnconfiguration.html#cfn-quicksight-template-columnconfiguration-role
	//
	Role *string `field:"optional" json:"role" yaml:"role"`
}

