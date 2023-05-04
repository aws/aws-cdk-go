package awsquicksight


// The configuration of a KPI visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kPIConfigurationProperty := &KPIConfigurationProperty{
//   	FieldWells: &KPIFieldWellsProperty{
//   		TargetValues: []interface{}{
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
//   		TrendGroups: []interface{}{
//   			&DimensionFieldProperty{
//   				CategoricalDimensionField: &CategoricalDimensionFieldProperty{
//   					Column: &ColumnIdentifierProperty{
//   						ColumnName: jsii.String("columnName"),
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//   					FieldId: jsii.String("fieldId"),
//
//   					// the properties below are optional
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
//   										Symbol: jsii.String("symbol"),
//   										Visibility: jsii.String("visibility"),
//   									},
//   								},
//   								Suffix: jsii.String("suffix"),
//   							},
//   						},
//   					},
//   					HierarchyId: jsii.String("hierarchyId"),
//   				},
//   				DateDimensionField: &DateDimensionFieldProperty{
//   					Column: &ColumnIdentifierProperty{
//   						ColumnName: jsii.String("columnName"),
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//   					FieldId: jsii.String("fieldId"),
//
//   					// the properties below are optional
//   					DateGranularity: jsii.String("dateGranularity"),
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
//   										Symbol: jsii.String("symbol"),
//   										Visibility: jsii.String("visibility"),
//   									},
//   								},
//   								Suffix: jsii.String("suffix"),
//   							},
//   						},
//   					},
//   					HierarchyId: jsii.String("hierarchyId"),
//   				},
//   				NumericalDimensionField: &NumericalDimensionFieldProperty{
//   					Column: &ColumnIdentifierProperty{
//   						ColumnName: jsii.String("columnName"),
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//   					FieldId: jsii.String("fieldId"),
//
//   					// the properties below are optional
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
//   										Symbol: jsii.String("symbol"),
//   										Visibility: jsii.String("visibility"),
//   									},
//   								},
//   								Suffix: jsii.String("suffix"),
//   							},
//   						},
//   					},
//   					HierarchyId: jsii.String("hierarchyId"),
//   				},
//   			},
//   		},
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
//   	KpiOptions: &KPIOptionsProperty{
//   		Comparison: &ComparisonConfigurationProperty{
//   			ComparisonFormat: &ComparisonFormatConfigurationProperty{
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
//   							Symbol: jsii.String("symbol"),
//   							Visibility: jsii.String("visibility"),
//   						},
//   					},
//   					Suffix: jsii.String("suffix"),
//   				},
//   			},
//   			ComparisonMethod: jsii.String("comparisonMethod"),
//   		},
//   		PrimaryValueDisplayType: jsii.String("primaryValueDisplayType"),
//   		PrimaryValueFontConfiguration: &FontConfigurationProperty{
//   			FontColor: jsii.String("fontColor"),
//   			FontDecoration: jsii.String("fontDecoration"),
//   			FontSize: &FontSizeProperty{
//   				Relative: jsii.String("relative"),
//   			},
//   			FontStyle: jsii.String("fontStyle"),
//   			FontWeight: &FontWeightProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		ProgressBar: &ProgressBarOptionsProperty{
//   			Visibility: jsii.String("visibility"),
//   		},
//   		SecondaryValue: &SecondaryValueOptionsProperty{
//   			Visibility: jsii.String("visibility"),
//   		},
//   		SecondaryValueFontConfiguration: &FontConfigurationProperty{
//   			FontColor: jsii.String("fontColor"),
//   			FontDecoration: jsii.String("fontDecoration"),
//   			FontSize: &FontSizeProperty{
//   				Relative: jsii.String("relative"),
//   			},
//   			FontStyle: jsii.String("fontStyle"),
//   			FontWeight: &FontWeightProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		TrendArrows: &TrendArrowOptionsProperty{
//   			Visibility: jsii.String("visibility"),
//   		},
//   	},
//   	SortConfiguration: &KPISortConfigurationProperty{
//   		TrendGroupSort: []interface{}{
//   			&FieldSortOptionsProperty{
//   				ColumnSort: &ColumnSortProperty{
//   					Direction: jsii.String("direction"),
//   					SortBy: &ColumnIdentifierProperty{
//   						ColumnName: jsii.String("columnName"),
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//
//   					// the properties below are optional
//   					AggregationFunction: &AggregationFunctionProperty{
//   						CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   						DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   						NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   							PercentileAggregation: &PercentileAggregationProperty{
//   								PercentileValue: jsii.Number(123),
//   							},
//   							SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   						},
//   					},
//   				},
//   				FieldSort: &FieldSortProperty{
//   					Direction: jsii.String("direction"),
//   					FieldId: jsii.String("fieldId"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnTemplate_KPIConfigurationProperty struct {
	// The field well configuration of a KPI visual.
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The options that determine the presentation of a KPI visual.
	KpiOptions interface{} `field:"optional" json:"kpiOptions" yaml:"kpiOptions"`
	// The sort configuration of a KPI visual.
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
}

