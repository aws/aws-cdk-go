package awsquicksight


// The field well configuration of a histogram.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   histogramFieldWellsProperty := &HistogramFieldWellsProperty{
//   	HistogramAggregatedFieldWells: &HistogramAggregatedFieldWellsProperty{
//   		Values: []interface{}{
//   			&MeasureFieldProperty{
//   				CalculatedMeasureField: &CalculatedMeasureFieldProperty{
//   					Expression: jsii.String("expression"),
//   					FieldId: jsii.String("fieldId"),
//   				},
//   				CategoricalMeasureField: &CategoricalMeasureFieldProperty{
//   					Column: &ColumnIdentifierProperty{
//   						ColumnName: jsii.String("columnName"),
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//   					FieldId: jsii.String("fieldId"),
//
//   					// the properties below are optional
//   					AggregationFunction: jsii.String("aggregationFunction"),
//   					FormatConfiguration: &StringFormatConfigurationProperty{
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   							CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   								DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   									DecimalPlaces: jsii.Number(123),
//   								},
//   								NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   									DisplayMode: jsii.String("displayMode"),
//   								},
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumberScale: jsii.String("numberScale"),
//   								Prefix: jsii.String("prefix"),
//   								SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   									DecimalSeparator: jsii.String("decimalSeparator"),
//   									ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   										GroupingStyle: jsii.String("groupingStyle"),
//   										Symbol: jsii.String("symbol"),
//   										Visibility: jsii.String("visibility"),
//   									},
//   								},
//   								Suffix: jsii.String("suffix"),
//   								Symbol: jsii.String("symbol"),
//   							},
//   							NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   								DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   									DecimalPlaces: jsii.Number(123),
//   								},
//   								NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   									DisplayMode: jsii.String("displayMode"),
//   								},
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumberScale: jsii.String("numberScale"),
//   								Prefix: jsii.String("prefix"),
//   								SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   									DecimalSeparator: jsii.String("decimalSeparator"),
//   									ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   										GroupingStyle: jsii.String("groupingStyle"),
//   										Symbol: jsii.String("symbol"),
//   										Visibility: jsii.String("visibility"),
//   									},
//   								},
//   								Suffix: jsii.String("suffix"),
//   							},
//   							PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   								DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   									DecimalPlaces: jsii.Number(123),
//   								},
//   								NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   									DisplayMode: jsii.String("displayMode"),
//   								},
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								Prefix: jsii.String("prefix"),
//   								SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   									DecimalSeparator: jsii.String("decimalSeparator"),
//   									ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   										GroupingStyle: jsii.String("groupingStyle"),
//   										Symbol: jsii.String("symbol"),
//   										Visibility: jsii.String("visibility"),
//   									},
//   								},
//   								Suffix: jsii.String("suffix"),
//   							},
//   						},
//   					},
//   				},
//   				DateMeasureField: &DateMeasureFieldProperty{
//   					Column: &ColumnIdentifierProperty{
//   						ColumnName: jsii.String("columnName"),
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//   					FieldId: jsii.String("fieldId"),
//
//   					// the properties below are optional
//   					AggregationFunction: jsii.String("aggregationFunction"),
//   					FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   						DateTimeFormat: jsii.String("dateTimeFormat"),
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   							CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   								DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   									DecimalPlaces: jsii.Number(123),
//   								},
//   								NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   									DisplayMode: jsii.String("displayMode"),
//   								},
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumberScale: jsii.String("numberScale"),
//   								Prefix: jsii.String("prefix"),
//   								SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   									DecimalSeparator: jsii.String("decimalSeparator"),
//   									ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   										GroupingStyle: jsii.String("groupingStyle"),
//   										Symbol: jsii.String("symbol"),
//   										Visibility: jsii.String("visibility"),
//   									},
//   								},
//   								Suffix: jsii.String("suffix"),
//   								Symbol: jsii.String("symbol"),
//   							},
//   							NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   								DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   									DecimalPlaces: jsii.Number(123),
//   								},
//   								NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   									DisplayMode: jsii.String("displayMode"),
//   								},
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumberScale: jsii.String("numberScale"),
//   								Prefix: jsii.String("prefix"),
//   								SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   									DecimalSeparator: jsii.String("decimalSeparator"),
//   									ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   										GroupingStyle: jsii.String("groupingStyle"),
//   										Symbol: jsii.String("symbol"),
//   										Visibility: jsii.String("visibility"),
//   									},
//   								},
//   								Suffix: jsii.String("suffix"),
//   							},
//   							PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   								DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   									DecimalPlaces: jsii.Number(123),
//   								},
//   								NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   									DisplayMode: jsii.String("displayMode"),
//   								},
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								Prefix: jsii.String("prefix"),
//   								SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   									DecimalSeparator: jsii.String("decimalSeparator"),
//   									ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   										GroupingStyle: jsii.String("groupingStyle"),
//   										Symbol: jsii.String("symbol"),
//   										Visibility: jsii.String("visibility"),
//   									},
//   								},
//   								Suffix: jsii.String("suffix"),
//   							},
//   						},
//   					},
//   				},
//   				NumericalMeasureField: &NumericalMeasureFieldProperty{
//   					Column: &ColumnIdentifierProperty{
//   						ColumnName: jsii.String("columnName"),
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//   					FieldId: jsii.String("fieldId"),
//
//   					// the properties below are optional
//   					AggregationFunction: &NumericalAggregationFunctionProperty{
//   						PercentileAggregation: &PercentileAggregationProperty{
//   							PercentileValue: jsii.Number(123),
//   						},
//   						SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   					},
//   					FormatConfiguration: &NumberFormatConfigurationProperty{
//   						FormatConfiguration: &NumericFormatConfigurationProperty{
//   							CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   								DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   									DecimalPlaces: jsii.Number(123),
//   								},
//   								NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   									DisplayMode: jsii.String("displayMode"),
//   								},
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumberScale: jsii.String("numberScale"),
//   								Prefix: jsii.String("prefix"),
//   								SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   									DecimalSeparator: jsii.String("decimalSeparator"),
//   									ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   										GroupingStyle: jsii.String("groupingStyle"),
//   										Symbol: jsii.String("symbol"),
//   										Visibility: jsii.String("visibility"),
//   									},
//   								},
//   								Suffix: jsii.String("suffix"),
//   								Symbol: jsii.String("symbol"),
//   							},
//   							NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   								DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   									DecimalPlaces: jsii.Number(123),
//   								},
//   								NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   									DisplayMode: jsii.String("displayMode"),
//   								},
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumberScale: jsii.String("numberScale"),
//   								Prefix: jsii.String("prefix"),
//   								SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   									DecimalSeparator: jsii.String("decimalSeparator"),
//   									ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   										GroupingStyle: jsii.String("groupingStyle"),
//   										Symbol: jsii.String("symbol"),
//   										Visibility: jsii.String("visibility"),
//   									},
//   								},
//   								Suffix: jsii.String("suffix"),
//   							},
//   							PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   								DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   									DecimalPlaces: jsii.Number(123),
//   								},
//   								NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   									DisplayMode: jsii.String("displayMode"),
//   								},
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								Prefix: jsii.String("prefix"),
//   								SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   									DecimalSeparator: jsii.String("decimalSeparator"),
//   									ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   										GroupingStyle: jsii.String("groupingStyle"),
//   										Symbol: jsii.String("symbol"),
//   										Visibility: jsii.String("visibility"),
//   									},
//   								},
//   								Suffix: jsii.String("suffix"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-histogramfieldwells.html
//
type CfnTemplate_HistogramFieldWellsProperty struct {
	// The field well configuration of a histogram.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-histogramfieldwells.html#cfn-quicksight-template-histogramfieldwells-histogramaggregatedfieldwells
	//
	HistogramAggregatedFieldWells interface{} `field:"optional" json:"histogramAggregatedFieldWells" yaml:"histogramAggregatedFieldWells"`
}

