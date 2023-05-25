package awsquicksight


// The field well configuration of a waterfall visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   waterfallChartAggregatedFieldWellsProperty := &WaterfallChartAggregatedFieldWellsProperty{
//   	Breakdowns: []interface{}{
//   		&DimensionFieldProperty{
//   			CategoricalDimensionField: &CategoricalDimensionFieldProperty{
//   				Column: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				FieldId: jsii.String("fieldId"),
//
//   				// the properties below are optional
//   				FormatConfiguration: &StringFormatConfigurationProperty{
//   					NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   						NullString: jsii.String("nullString"),
//   					},
//   					NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   						CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   							DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   								DecimalPlaces: jsii.Number(123),
//   							},
//   							NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   								DisplayMode: jsii.String("displayMode"),
//   							},
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							NumberScale: jsii.String("numberScale"),
//   							Prefix: jsii.String("prefix"),
//   							SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   								DecimalSeparator: jsii.String("decimalSeparator"),
//   								ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   									Symbol: jsii.String("symbol"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   							},
//   							Suffix: jsii.String("suffix"),
//   							Symbol: jsii.String("symbol"),
//   						},
//   						NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   							DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   								DecimalPlaces: jsii.Number(123),
//   							},
//   							NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   								DisplayMode: jsii.String("displayMode"),
//   							},
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							NumberScale: jsii.String("numberScale"),
//   							Prefix: jsii.String("prefix"),
//   							SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   								DecimalSeparator: jsii.String("decimalSeparator"),
//   								ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   									Symbol: jsii.String("symbol"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   							},
//   							Suffix: jsii.String("suffix"),
//   						},
//   						PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   							DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   								DecimalPlaces: jsii.Number(123),
//   							},
//   							NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   								DisplayMode: jsii.String("displayMode"),
//   							},
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							Prefix: jsii.String("prefix"),
//   							SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   								DecimalSeparator: jsii.String("decimalSeparator"),
//   								ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   									Symbol: jsii.String("symbol"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   							},
//   							Suffix: jsii.String("suffix"),
//   						},
//   					},
//   				},
//   				HierarchyId: jsii.String("hierarchyId"),
//   			},
//   			DateDimensionField: &DateDimensionFieldProperty{
//   				Column: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				FieldId: jsii.String("fieldId"),
//
//   				// the properties below are optional
//   				DateGranularity: jsii.String("dateGranularity"),
//   				FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   					DateTimeFormat: jsii.String("dateTimeFormat"),
//   					NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   						NullString: jsii.String("nullString"),
//   					},
//   					NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   						CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   							DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   								DecimalPlaces: jsii.Number(123),
//   							},
//   							NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   								DisplayMode: jsii.String("displayMode"),
//   							},
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							NumberScale: jsii.String("numberScale"),
//   							Prefix: jsii.String("prefix"),
//   							SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   								DecimalSeparator: jsii.String("decimalSeparator"),
//   								ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   									Symbol: jsii.String("symbol"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   							},
//   							Suffix: jsii.String("suffix"),
//   							Symbol: jsii.String("symbol"),
//   						},
//   						NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   							DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   								DecimalPlaces: jsii.Number(123),
//   							},
//   							NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   								DisplayMode: jsii.String("displayMode"),
//   							},
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							NumberScale: jsii.String("numberScale"),
//   							Prefix: jsii.String("prefix"),
//   							SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   								DecimalSeparator: jsii.String("decimalSeparator"),
//   								ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   									Symbol: jsii.String("symbol"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   							},
//   							Suffix: jsii.String("suffix"),
//   						},
//   						PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   							DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   								DecimalPlaces: jsii.Number(123),
//   							},
//   							NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   								DisplayMode: jsii.String("displayMode"),
//   							},
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							Prefix: jsii.String("prefix"),
//   							SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   								DecimalSeparator: jsii.String("decimalSeparator"),
//   								ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   									Symbol: jsii.String("symbol"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   							},
//   							Suffix: jsii.String("suffix"),
//   						},
//   					},
//   				},
//   				HierarchyId: jsii.String("hierarchyId"),
//   			},
//   			NumericalDimensionField: &NumericalDimensionFieldProperty{
//   				Column: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				FieldId: jsii.String("fieldId"),
//
//   				// the properties below are optional
//   				FormatConfiguration: &NumberFormatConfigurationProperty{
//   					FormatConfiguration: &NumericFormatConfigurationProperty{
//   						CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   							DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   								DecimalPlaces: jsii.Number(123),
//   							},
//   							NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   								DisplayMode: jsii.String("displayMode"),
//   							},
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							NumberScale: jsii.String("numberScale"),
//   							Prefix: jsii.String("prefix"),
//   							SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   								DecimalSeparator: jsii.String("decimalSeparator"),
//   								ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   									Symbol: jsii.String("symbol"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   							},
//   							Suffix: jsii.String("suffix"),
//   							Symbol: jsii.String("symbol"),
//   						},
//   						NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   							DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   								DecimalPlaces: jsii.Number(123),
//   							},
//   							NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   								DisplayMode: jsii.String("displayMode"),
//   							},
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							NumberScale: jsii.String("numberScale"),
//   							Prefix: jsii.String("prefix"),
//   							SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   								DecimalSeparator: jsii.String("decimalSeparator"),
//   								ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   									Symbol: jsii.String("symbol"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   							},
//   							Suffix: jsii.String("suffix"),
//   						},
//   						PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   							DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   								DecimalPlaces: jsii.Number(123),
//   							},
//   							NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   								DisplayMode: jsii.String("displayMode"),
//   							},
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							Prefix: jsii.String("prefix"),
//   							SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   								DecimalSeparator: jsii.String("decimalSeparator"),
//   								ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   									Symbol: jsii.String("symbol"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   							},
//   							Suffix: jsii.String("suffix"),
//   						},
//   					},
//   				},
//   				HierarchyId: jsii.String("hierarchyId"),
//   			},
//   		},
//   	},
//   	Categories: []interface{}{
//   		&DimensionFieldProperty{
//   			CategoricalDimensionField: &CategoricalDimensionFieldProperty{
//   				Column: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				FieldId: jsii.String("fieldId"),
//
//   				// the properties below are optional
//   				FormatConfiguration: &StringFormatConfigurationProperty{
//   					NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   						NullString: jsii.String("nullString"),
//   					},
//   					NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   						CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   							DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   								DecimalPlaces: jsii.Number(123),
//   							},
//   							NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   								DisplayMode: jsii.String("displayMode"),
//   							},
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							NumberScale: jsii.String("numberScale"),
//   							Prefix: jsii.String("prefix"),
//   							SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   								DecimalSeparator: jsii.String("decimalSeparator"),
//   								ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   									Symbol: jsii.String("symbol"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   							},
//   							Suffix: jsii.String("suffix"),
//   							Symbol: jsii.String("symbol"),
//   						},
//   						NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   							DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   								DecimalPlaces: jsii.Number(123),
//   							},
//   							NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   								DisplayMode: jsii.String("displayMode"),
//   							},
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							NumberScale: jsii.String("numberScale"),
//   							Prefix: jsii.String("prefix"),
//   							SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   								DecimalSeparator: jsii.String("decimalSeparator"),
//   								ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   									Symbol: jsii.String("symbol"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   							},
//   							Suffix: jsii.String("suffix"),
//   						},
//   						PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   							DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   								DecimalPlaces: jsii.Number(123),
//   							},
//   							NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   								DisplayMode: jsii.String("displayMode"),
//   							},
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							Prefix: jsii.String("prefix"),
//   							SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   								DecimalSeparator: jsii.String("decimalSeparator"),
//   								ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   									Symbol: jsii.String("symbol"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   							},
//   							Suffix: jsii.String("suffix"),
//   						},
//   					},
//   				},
//   				HierarchyId: jsii.String("hierarchyId"),
//   			},
//   			DateDimensionField: &DateDimensionFieldProperty{
//   				Column: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				FieldId: jsii.String("fieldId"),
//
//   				// the properties below are optional
//   				DateGranularity: jsii.String("dateGranularity"),
//   				FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   					DateTimeFormat: jsii.String("dateTimeFormat"),
//   					NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   						NullString: jsii.String("nullString"),
//   					},
//   					NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   						CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   							DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   								DecimalPlaces: jsii.Number(123),
//   							},
//   							NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   								DisplayMode: jsii.String("displayMode"),
//   							},
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							NumberScale: jsii.String("numberScale"),
//   							Prefix: jsii.String("prefix"),
//   							SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   								DecimalSeparator: jsii.String("decimalSeparator"),
//   								ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   									Symbol: jsii.String("symbol"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   							},
//   							Suffix: jsii.String("suffix"),
//   							Symbol: jsii.String("symbol"),
//   						},
//   						NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   							DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   								DecimalPlaces: jsii.Number(123),
//   							},
//   							NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   								DisplayMode: jsii.String("displayMode"),
//   							},
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							NumberScale: jsii.String("numberScale"),
//   							Prefix: jsii.String("prefix"),
//   							SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   								DecimalSeparator: jsii.String("decimalSeparator"),
//   								ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   									Symbol: jsii.String("symbol"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   							},
//   							Suffix: jsii.String("suffix"),
//   						},
//   						PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   							DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   								DecimalPlaces: jsii.Number(123),
//   							},
//   							NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   								DisplayMode: jsii.String("displayMode"),
//   							},
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							Prefix: jsii.String("prefix"),
//   							SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   								DecimalSeparator: jsii.String("decimalSeparator"),
//   								ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   									Symbol: jsii.String("symbol"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   							},
//   							Suffix: jsii.String("suffix"),
//   						},
//   					},
//   				},
//   				HierarchyId: jsii.String("hierarchyId"),
//   			},
//   			NumericalDimensionField: &NumericalDimensionFieldProperty{
//   				Column: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				FieldId: jsii.String("fieldId"),
//
//   				// the properties below are optional
//   				FormatConfiguration: &NumberFormatConfigurationProperty{
//   					FormatConfiguration: &NumericFormatConfigurationProperty{
//   						CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   							DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   								DecimalPlaces: jsii.Number(123),
//   							},
//   							NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   								DisplayMode: jsii.String("displayMode"),
//   							},
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							NumberScale: jsii.String("numberScale"),
//   							Prefix: jsii.String("prefix"),
//   							SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   								DecimalSeparator: jsii.String("decimalSeparator"),
//   								ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   									Symbol: jsii.String("symbol"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   							},
//   							Suffix: jsii.String("suffix"),
//   							Symbol: jsii.String("symbol"),
//   						},
//   						NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   							DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   								DecimalPlaces: jsii.Number(123),
//   							},
//   							NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   								DisplayMode: jsii.String("displayMode"),
//   							},
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							NumberScale: jsii.String("numberScale"),
//   							Prefix: jsii.String("prefix"),
//   							SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   								DecimalSeparator: jsii.String("decimalSeparator"),
//   								ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   									Symbol: jsii.String("symbol"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   							},
//   							Suffix: jsii.String("suffix"),
//   						},
//   						PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   							DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   								DecimalPlaces: jsii.Number(123),
//   							},
//   							NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   								DisplayMode: jsii.String("displayMode"),
//   							},
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							Prefix: jsii.String("prefix"),
//   							SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   								DecimalSeparator: jsii.String("decimalSeparator"),
//   								ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   									Symbol: jsii.String("symbol"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   							},
//   							Suffix: jsii.String("suffix"),
//   						},
//   					},
//   				},
//   				HierarchyId: jsii.String("hierarchyId"),
//   			},
//   		},
//   	},
//   	Values: []interface{}{
//   		&MeasureFieldProperty{
//   			CalculatedMeasureField: &CalculatedMeasureFieldProperty{
//   				Expression: jsii.String("expression"),
//   				FieldId: jsii.String("fieldId"),
//   			},
//   			CategoricalMeasureField: &CategoricalMeasureFieldProperty{
//   				Column: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				FieldId: jsii.String("fieldId"),
//
//   				// the properties below are optional
//   				AggregationFunction: jsii.String("aggregationFunction"),
//   				FormatConfiguration: &StringFormatConfigurationProperty{
//   					NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   						NullString: jsii.String("nullString"),
//   					},
//   					NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   						CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   							DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   								DecimalPlaces: jsii.Number(123),
//   							},
//   							NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   								DisplayMode: jsii.String("displayMode"),
//   							},
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							NumberScale: jsii.String("numberScale"),
//   							Prefix: jsii.String("prefix"),
//   							SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   								DecimalSeparator: jsii.String("decimalSeparator"),
//   								ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   									Symbol: jsii.String("symbol"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   							},
//   							Suffix: jsii.String("suffix"),
//   							Symbol: jsii.String("symbol"),
//   						},
//   						NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   							DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   								DecimalPlaces: jsii.Number(123),
//   							},
//   							NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   								DisplayMode: jsii.String("displayMode"),
//   							},
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							NumberScale: jsii.String("numberScale"),
//   							Prefix: jsii.String("prefix"),
//   							SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   								DecimalSeparator: jsii.String("decimalSeparator"),
//   								ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   									Symbol: jsii.String("symbol"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   							},
//   							Suffix: jsii.String("suffix"),
//   						},
//   						PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   							DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   								DecimalPlaces: jsii.Number(123),
//   							},
//   							NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   								DisplayMode: jsii.String("displayMode"),
//   							},
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							Prefix: jsii.String("prefix"),
//   							SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   								DecimalSeparator: jsii.String("decimalSeparator"),
//   								ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   									Symbol: jsii.String("symbol"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   							},
//   							Suffix: jsii.String("suffix"),
//   						},
//   					},
//   				},
//   			},
//   			DateMeasureField: &DateMeasureFieldProperty{
//   				Column: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				FieldId: jsii.String("fieldId"),
//
//   				// the properties below are optional
//   				AggregationFunction: jsii.String("aggregationFunction"),
//   				FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   					DateTimeFormat: jsii.String("dateTimeFormat"),
//   					NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   						NullString: jsii.String("nullString"),
//   					},
//   					NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   						CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   							DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   								DecimalPlaces: jsii.Number(123),
//   							},
//   							NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   								DisplayMode: jsii.String("displayMode"),
//   							},
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							NumberScale: jsii.String("numberScale"),
//   							Prefix: jsii.String("prefix"),
//   							SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   								DecimalSeparator: jsii.String("decimalSeparator"),
//   								ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   									Symbol: jsii.String("symbol"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   							},
//   							Suffix: jsii.String("suffix"),
//   							Symbol: jsii.String("symbol"),
//   						},
//   						NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   							DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   								DecimalPlaces: jsii.Number(123),
//   							},
//   							NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   								DisplayMode: jsii.String("displayMode"),
//   							},
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							NumberScale: jsii.String("numberScale"),
//   							Prefix: jsii.String("prefix"),
//   							SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   								DecimalSeparator: jsii.String("decimalSeparator"),
//   								ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   									Symbol: jsii.String("symbol"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   							},
//   							Suffix: jsii.String("suffix"),
//   						},
//   						PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   							DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   								DecimalPlaces: jsii.Number(123),
//   							},
//   							NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   								DisplayMode: jsii.String("displayMode"),
//   							},
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							Prefix: jsii.String("prefix"),
//   							SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   								DecimalSeparator: jsii.String("decimalSeparator"),
//   								ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   									Symbol: jsii.String("symbol"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   							},
//   							Suffix: jsii.String("suffix"),
//   						},
//   					},
//   				},
//   			},
//   			NumericalMeasureField: &NumericalMeasureFieldProperty{
//   				Column: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				FieldId: jsii.String("fieldId"),
//
//   				// the properties below are optional
//   				AggregationFunction: &NumericalAggregationFunctionProperty{
//   					PercentileAggregation: &PercentileAggregationProperty{
//   						PercentileValue: jsii.Number(123),
//   					},
//   					SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   				},
//   				FormatConfiguration: &NumberFormatConfigurationProperty{
//   					FormatConfiguration: &NumericFormatConfigurationProperty{
//   						CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   							DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   								DecimalPlaces: jsii.Number(123),
//   							},
//   							NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   								DisplayMode: jsii.String("displayMode"),
//   							},
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							NumberScale: jsii.String("numberScale"),
//   							Prefix: jsii.String("prefix"),
//   							SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   								DecimalSeparator: jsii.String("decimalSeparator"),
//   								ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   									Symbol: jsii.String("symbol"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   							},
//   							Suffix: jsii.String("suffix"),
//   							Symbol: jsii.String("symbol"),
//   						},
//   						NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   							DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   								DecimalPlaces: jsii.Number(123),
//   							},
//   							NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   								DisplayMode: jsii.String("displayMode"),
//   							},
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							NumberScale: jsii.String("numberScale"),
//   							Prefix: jsii.String("prefix"),
//   							SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   								DecimalSeparator: jsii.String("decimalSeparator"),
//   								ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   									Symbol: jsii.String("symbol"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   							},
//   							Suffix: jsii.String("suffix"),
//   						},
//   						PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   							DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   								DecimalPlaces: jsii.Number(123),
//   							},
//   							NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   								DisplayMode: jsii.String("displayMode"),
//   							},
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							Prefix: jsii.String("prefix"),
//   							SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   								DecimalSeparator: jsii.String("decimalSeparator"),
//   								ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   									Symbol: jsii.String("symbol"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   							},
//   							Suffix: jsii.String("suffix"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnTemplate_WaterfallChartAggregatedFieldWellsProperty struct {
	// The breakdown field wells of a waterfall visual.
	Breakdowns interface{} `field:"optional" json:"breakdowns" yaml:"breakdowns"`
	// The category field wells of a waterfall visual.
	Categories interface{} `field:"optional" json:"categories" yaml:"categories"`
	// The value field wells of a waterfall visual.
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

