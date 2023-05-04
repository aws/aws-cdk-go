package awsquicksight


// The configuration for a `PivotTableVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pivotTableConfigurationProperty := &PivotTableConfigurationProperty{
//   	FieldOptions: &PivotTableFieldOptionsProperty{
//   		DataPathOptions: []interface{}{
//   			&PivotTableDataPathOptionProperty{
//   				DataPathList: []interface{}{
//   					&DataPathValueProperty{
//   						FieldId: jsii.String("fieldId"),
//   						FieldValue: jsii.String("fieldValue"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				Width: jsii.String("width"),
//   			},
//   		},
//   		SelectedFieldOptions: []interface{}{
//   			&PivotTableFieldOptionProperty{
//   				FieldId: jsii.String("fieldId"),
//
//   				// the properties below are optional
//   				CustomLabel: jsii.String("customLabel"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   		},
//   	},
//   	FieldWells: &PivotTableFieldWellsProperty{
//   		PivotTableAggregatedFieldWells: &PivotTableAggregatedFieldWellsProperty{
//   			Columns: []interface{}{
//   				&DimensionFieldProperty{
//   					CategoricalDimensionField: &CategoricalDimensionFieldProperty{
//   						Column: &ColumnIdentifierProperty{
//   							ColumnName: jsii.String("columnName"),
//   							DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   						},
//   						FieldId: jsii.String("fieldId"),
//
//   						// the properties below are optional
//   						FormatConfiguration: &StringFormatConfigurationProperty{
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   								CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   									DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   										DecimalPlaces: jsii.Number(123),
//   									},
//   									NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   										DisplayMode: jsii.String("displayMode"),
//   									},
//   									NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   										NullString: jsii.String("nullString"),
//   									},
//   									NumberScale: jsii.String("numberScale"),
//   									Prefix: jsii.String("prefix"),
//   									SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   										DecimalSeparator: jsii.String("decimalSeparator"),
//   										ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   											Symbol: jsii.String("symbol"),
//   											Visibility: jsii.String("visibility"),
//   										},
//   									},
//   									Suffix: jsii.String("suffix"),
//   									Symbol: jsii.String("symbol"),
//   								},
//   								NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   									DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   										DecimalPlaces: jsii.Number(123),
//   									},
//   									NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   										DisplayMode: jsii.String("displayMode"),
//   									},
//   									NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   										NullString: jsii.String("nullString"),
//   									},
//   									NumberScale: jsii.String("numberScale"),
//   									Prefix: jsii.String("prefix"),
//   									SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   										DecimalSeparator: jsii.String("decimalSeparator"),
//   										ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   											Symbol: jsii.String("symbol"),
//   											Visibility: jsii.String("visibility"),
//   										},
//   									},
//   									Suffix: jsii.String("suffix"),
//   								},
//   								PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   									DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   										DecimalPlaces: jsii.Number(123),
//   									},
//   									NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   										DisplayMode: jsii.String("displayMode"),
//   									},
//   									NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   										NullString: jsii.String("nullString"),
//   									},
//   									Prefix: jsii.String("prefix"),
//   									SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   										DecimalSeparator: jsii.String("decimalSeparator"),
//   										ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   											Symbol: jsii.String("symbol"),
//   											Visibility: jsii.String("visibility"),
//   										},
//   									},
//   									Suffix: jsii.String("suffix"),
//   								},
//   							},
//   						},
//   						HierarchyId: jsii.String("hierarchyId"),
//   					},
//   					DateDimensionField: &DateDimensionFieldProperty{
//   						Column: &ColumnIdentifierProperty{
//   							ColumnName: jsii.String("columnName"),
//   							DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   						},
//   						FieldId: jsii.String("fieldId"),
//
//   						// the properties below are optional
//   						DateGranularity: jsii.String("dateGranularity"),
//   						FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   							DateTimeFormat: jsii.String("dateTimeFormat"),
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   								CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   									DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   										DecimalPlaces: jsii.Number(123),
//   									},
//   									NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   										DisplayMode: jsii.String("displayMode"),
//   									},
//   									NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   										NullString: jsii.String("nullString"),
//   									},
//   									NumberScale: jsii.String("numberScale"),
//   									Prefix: jsii.String("prefix"),
//   									SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   										DecimalSeparator: jsii.String("decimalSeparator"),
//   										ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   											Symbol: jsii.String("symbol"),
//   											Visibility: jsii.String("visibility"),
//   										},
//   									},
//   									Suffix: jsii.String("suffix"),
//   									Symbol: jsii.String("symbol"),
//   								},
//   								NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   									DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   										DecimalPlaces: jsii.Number(123),
//   									},
//   									NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   										DisplayMode: jsii.String("displayMode"),
//   									},
//   									NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   										NullString: jsii.String("nullString"),
//   									},
//   									NumberScale: jsii.String("numberScale"),
//   									Prefix: jsii.String("prefix"),
//   									SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   										DecimalSeparator: jsii.String("decimalSeparator"),
//   										ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   											Symbol: jsii.String("symbol"),
//   											Visibility: jsii.String("visibility"),
//   										},
//   									},
//   									Suffix: jsii.String("suffix"),
//   								},
//   								PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   									DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   										DecimalPlaces: jsii.Number(123),
//   									},
//   									NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   										DisplayMode: jsii.String("displayMode"),
//   									},
//   									NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   										NullString: jsii.String("nullString"),
//   									},
//   									Prefix: jsii.String("prefix"),
//   									SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   										DecimalSeparator: jsii.String("decimalSeparator"),
//   										ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   											Symbol: jsii.String("symbol"),
//   											Visibility: jsii.String("visibility"),
//   										},
//   									},
//   									Suffix: jsii.String("suffix"),
//   								},
//   							},
//   						},
//   						HierarchyId: jsii.String("hierarchyId"),
//   					},
//   					NumericalDimensionField: &NumericalDimensionFieldProperty{
//   						Column: &ColumnIdentifierProperty{
//   							ColumnName: jsii.String("columnName"),
//   							DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   						},
//   						FieldId: jsii.String("fieldId"),
//
//   						// the properties below are optional
//   						FormatConfiguration: &NumberFormatConfigurationProperty{
//   							FormatConfiguration: &NumericFormatConfigurationProperty{
//   								CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   									DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   										DecimalPlaces: jsii.Number(123),
//   									},
//   									NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   										DisplayMode: jsii.String("displayMode"),
//   									},
//   									NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   										NullString: jsii.String("nullString"),
//   									},
//   									NumberScale: jsii.String("numberScale"),
//   									Prefix: jsii.String("prefix"),
//   									SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   										DecimalSeparator: jsii.String("decimalSeparator"),
//   										ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   											Symbol: jsii.String("symbol"),
//   											Visibility: jsii.String("visibility"),
//   										},
//   									},
//   									Suffix: jsii.String("suffix"),
//   									Symbol: jsii.String("symbol"),
//   								},
//   								NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   									DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   										DecimalPlaces: jsii.Number(123),
//   									},
//   									NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   										DisplayMode: jsii.String("displayMode"),
//   									},
//   									NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   										NullString: jsii.String("nullString"),
//   									},
//   									NumberScale: jsii.String("numberScale"),
//   									Prefix: jsii.String("prefix"),
//   									SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   										DecimalSeparator: jsii.String("decimalSeparator"),
//   										ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   											Symbol: jsii.String("symbol"),
//   											Visibility: jsii.String("visibility"),
//   										},
//   									},
//   									Suffix: jsii.String("suffix"),
//   								},
//   								PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   									DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   										DecimalPlaces: jsii.Number(123),
//   									},
//   									NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   										DisplayMode: jsii.String("displayMode"),
//   									},
//   									NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   										NullString: jsii.String("nullString"),
//   									},
//   									Prefix: jsii.String("prefix"),
//   									SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   										DecimalSeparator: jsii.String("decimalSeparator"),
//   										ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   											Symbol: jsii.String("symbol"),
//   											Visibility: jsii.String("visibility"),
//   										},
//   									},
//   									Suffix: jsii.String("suffix"),
//   								},
//   							},
//   						},
//   						HierarchyId: jsii.String("hierarchyId"),
//   					},
//   				},
//   			},
//   			Rows: []interface{}{
//   				&DimensionFieldProperty{
//   					CategoricalDimensionField: &CategoricalDimensionFieldProperty{
//   						Column: &ColumnIdentifierProperty{
//   							ColumnName: jsii.String("columnName"),
//   							DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   						},
//   						FieldId: jsii.String("fieldId"),
//
//   						// the properties below are optional
//   						FormatConfiguration: &StringFormatConfigurationProperty{
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   								CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   									DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   										DecimalPlaces: jsii.Number(123),
//   									},
//   									NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   										DisplayMode: jsii.String("displayMode"),
//   									},
//   									NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   										NullString: jsii.String("nullString"),
//   									},
//   									NumberScale: jsii.String("numberScale"),
//   									Prefix: jsii.String("prefix"),
//   									SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   										DecimalSeparator: jsii.String("decimalSeparator"),
//   										ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   											Symbol: jsii.String("symbol"),
//   											Visibility: jsii.String("visibility"),
//   										},
//   									},
//   									Suffix: jsii.String("suffix"),
//   									Symbol: jsii.String("symbol"),
//   								},
//   								NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   									DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   										DecimalPlaces: jsii.Number(123),
//   									},
//   									NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   										DisplayMode: jsii.String("displayMode"),
//   									},
//   									NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   										NullString: jsii.String("nullString"),
//   									},
//   									NumberScale: jsii.String("numberScale"),
//   									Prefix: jsii.String("prefix"),
//   									SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   										DecimalSeparator: jsii.String("decimalSeparator"),
//   										ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   											Symbol: jsii.String("symbol"),
//   											Visibility: jsii.String("visibility"),
//   										},
//   									},
//   									Suffix: jsii.String("suffix"),
//   								},
//   								PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   									DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   										DecimalPlaces: jsii.Number(123),
//   									},
//   									NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   										DisplayMode: jsii.String("displayMode"),
//   									},
//   									NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   										NullString: jsii.String("nullString"),
//   									},
//   									Prefix: jsii.String("prefix"),
//   									SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   										DecimalSeparator: jsii.String("decimalSeparator"),
//   										ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   											Symbol: jsii.String("symbol"),
//   											Visibility: jsii.String("visibility"),
//   										},
//   									},
//   									Suffix: jsii.String("suffix"),
//   								},
//   							},
//   						},
//   						HierarchyId: jsii.String("hierarchyId"),
//   					},
//   					DateDimensionField: &DateDimensionFieldProperty{
//   						Column: &ColumnIdentifierProperty{
//   							ColumnName: jsii.String("columnName"),
//   							DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   						},
//   						FieldId: jsii.String("fieldId"),
//
//   						// the properties below are optional
//   						DateGranularity: jsii.String("dateGranularity"),
//   						FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   							DateTimeFormat: jsii.String("dateTimeFormat"),
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   								CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   									DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   										DecimalPlaces: jsii.Number(123),
//   									},
//   									NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   										DisplayMode: jsii.String("displayMode"),
//   									},
//   									NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   										NullString: jsii.String("nullString"),
//   									},
//   									NumberScale: jsii.String("numberScale"),
//   									Prefix: jsii.String("prefix"),
//   									SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   										DecimalSeparator: jsii.String("decimalSeparator"),
//   										ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   											Symbol: jsii.String("symbol"),
//   											Visibility: jsii.String("visibility"),
//   										},
//   									},
//   									Suffix: jsii.String("suffix"),
//   									Symbol: jsii.String("symbol"),
//   								},
//   								NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   									DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   										DecimalPlaces: jsii.Number(123),
//   									},
//   									NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   										DisplayMode: jsii.String("displayMode"),
//   									},
//   									NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   										NullString: jsii.String("nullString"),
//   									},
//   									NumberScale: jsii.String("numberScale"),
//   									Prefix: jsii.String("prefix"),
//   									SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   										DecimalSeparator: jsii.String("decimalSeparator"),
//   										ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   											Symbol: jsii.String("symbol"),
//   											Visibility: jsii.String("visibility"),
//   										},
//   									},
//   									Suffix: jsii.String("suffix"),
//   								},
//   								PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   									DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   										DecimalPlaces: jsii.Number(123),
//   									},
//   									NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   										DisplayMode: jsii.String("displayMode"),
//   									},
//   									NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   										NullString: jsii.String("nullString"),
//   									},
//   									Prefix: jsii.String("prefix"),
//   									SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   										DecimalSeparator: jsii.String("decimalSeparator"),
//   										ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   											Symbol: jsii.String("symbol"),
//   											Visibility: jsii.String("visibility"),
//   										},
//   									},
//   									Suffix: jsii.String("suffix"),
//   								},
//   							},
//   						},
//   						HierarchyId: jsii.String("hierarchyId"),
//   					},
//   					NumericalDimensionField: &NumericalDimensionFieldProperty{
//   						Column: &ColumnIdentifierProperty{
//   							ColumnName: jsii.String("columnName"),
//   							DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   						},
//   						FieldId: jsii.String("fieldId"),
//
//   						// the properties below are optional
//   						FormatConfiguration: &NumberFormatConfigurationProperty{
//   							FormatConfiguration: &NumericFormatConfigurationProperty{
//   								CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   									DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   										DecimalPlaces: jsii.Number(123),
//   									},
//   									NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   										DisplayMode: jsii.String("displayMode"),
//   									},
//   									NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   										NullString: jsii.String("nullString"),
//   									},
//   									NumberScale: jsii.String("numberScale"),
//   									Prefix: jsii.String("prefix"),
//   									SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   										DecimalSeparator: jsii.String("decimalSeparator"),
//   										ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   											Symbol: jsii.String("symbol"),
//   											Visibility: jsii.String("visibility"),
//   										},
//   									},
//   									Suffix: jsii.String("suffix"),
//   									Symbol: jsii.String("symbol"),
//   								},
//   								NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   									DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   										DecimalPlaces: jsii.Number(123),
//   									},
//   									NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   										DisplayMode: jsii.String("displayMode"),
//   									},
//   									NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   										NullString: jsii.String("nullString"),
//   									},
//   									NumberScale: jsii.String("numberScale"),
//   									Prefix: jsii.String("prefix"),
//   									SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   										DecimalSeparator: jsii.String("decimalSeparator"),
//   										ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   											Symbol: jsii.String("symbol"),
//   											Visibility: jsii.String("visibility"),
//   										},
//   									},
//   									Suffix: jsii.String("suffix"),
//   								},
//   								PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   									DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   										DecimalPlaces: jsii.Number(123),
//   									},
//   									NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   										DisplayMode: jsii.String("displayMode"),
//   									},
//   									NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   										NullString: jsii.String("nullString"),
//   									},
//   									Prefix: jsii.String("prefix"),
//   									SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   										DecimalSeparator: jsii.String("decimalSeparator"),
//   										ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   											Symbol: jsii.String("symbol"),
//   											Visibility: jsii.String("visibility"),
//   										},
//   									},
//   									Suffix: jsii.String("suffix"),
//   								},
//   							},
//   						},
//   						HierarchyId: jsii.String("hierarchyId"),
//   					},
//   				},
//   			},
//   			Values: []interface{}{
//   				&MeasureFieldProperty{
//   					CalculatedMeasureField: &CalculatedMeasureFieldProperty{
//   						Expression: jsii.String("expression"),
//   						FieldId: jsii.String("fieldId"),
//   					},
//   					CategoricalMeasureField: &CategoricalMeasureFieldProperty{
//   						Column: &ColumnIdentifierProperty{
//   							ColumnName: jsii.String("columnName"),
//   							DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   						},
//   						FieldId: jsii.String("fieldId"),
//
//   						// the properties below are optional
//   						AggregationFunction: jsii.String("aggregationFunction"),
//   						FormatConfiguration: &StringFormatConfigurationProperty{
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   								CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   									DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   										DecimalPlaces: jsii.Number(123),
//   									},
//   									NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   										DisplayMode: jsii.String("displayMode"),
//   									},
//   									NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   										NullString: jsii.String("nullString"),
//   									},
//   									NumberScale: jsii.String("numberScale"),
//   									Prefix: jsii.String("prefix"),
//   									SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   										DecimalSeparator: jsii.String("decimalSeparator"),
//   										ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   											Symbol: jsii.String("symbol"),
//   											Visibility: jsii.String("visibility"),
//   										},
//   									},
//   									Suffix: jsii.String("suffix"),
//   									Symbol: jsii.String("symbol"),
//   								},
//   								NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   									DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   										DecimalPlaces: jsii.Number(123),
//   									},
//   									NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   										DisplayMode: jsii.String("displayMode"),
//   									},
//   									NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   										NullString: jsii.String("nullString"),
//   									},
//   									NumberScale: jsii.String("numberScale"),
//   									Prefix: jsii.String("prefix"),
//   									SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   										DecimalSeparator: jsii.String("decimalSeparator"),
//   										ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   											Symbol: jsii.String("symbol"),
//   											Visibility: jsii.String("visibility"),
//   										},
//   									},
//   									Suffix: jsii.String("suffix"),
//   								},
//   								PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   									DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   										DecimalPlaces: jsii.Number(123),
//   									},
//   									NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   										DisplayMode: jsii.String("displayMode"),
//   									},
//   									NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   										NullString: jsii.String("nullString"),
//   									},
//   									Prefix: jsii.String("prefix"),
//   									SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   										DecimalSeparator: jsii.String("decimalSeparator"),
//   										ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   											Symbol: jsii.String("symbol"),
//   											Visibility: jsii.String("visibility"),
//   										},
//   									},
//   									Suffix: jsii.String("suffix"),
//   								},
//   							},
//   						},
//   					},
//   					DateMeasureField: &DateMeasureFieldProperty{
//   						Column: &ColumnIdentifierProperty{
//   							ColumnName: jsii.String("columnName"),
//   							DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   						},
//   						FieldId: jsii.String("fieldId"),
//
//   						// the properties below are optional
//   						AggregationFunction: jsii.String("aggregationFunction"),
//   						FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   							DateTimeFormat: jsii.String("dateTimeFormat"),
//   							NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   								NullString: jsii.String("nullString"),
//   							},
//   							NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   								CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   									DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   										DecimalPlaces: jsii.Number(123),
//   									},
//   									NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   										DisplayMode: jsii.String("displayMode"),
//   									},
//   									NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   										NullString: jsii.String("nullString"),
//   									},
//   									NumberScale: jsii.String("numberScale"),
//   									Prefix: jsii.String("prefix"),
//   									SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   										DecimalSeparator: jsii.String("decimalSeparator"),
//   										ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   											Symbol: jsii.String("symbol"),
//   											Visibility: jsii.String("visibility"),
//   										},
//   									},
//   									Suffix: jsii.String("suffix"),
//   									Symbol: jsii.String("symbol"),
//   								},
//   								NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   									DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   										DecimalPlaces: jsii.Number(123),
//   									},
//   									NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   										DisplayMode: jsii.String("displayMode"),
//   									},
//   									NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   										NullString: jsii.String("nullString"),
//   									},
//   									NumberScale: jsii.String("numberScale"),
//   									Prefix: jsii.String("prefix"),
//   									SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   										DecimalSeparator: jsii.String("decimalSeparator"),
//   										ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   											Symbol: jsii.String("symbol"),
//   											Visibility: jsii.String("visibility"),
//   										},
//   									},
//   									Suffix: jsii.String("suffix"),
//   								},
//   								PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   									DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   										DecimalPlaces: jsii.Number(123),
//   									},
//   									NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   										DisplayMode: jsii.String("displayMode"),
//   									},
//   									NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   										NullString: jsii.String("nullString"),
//   									},
//   									Prefix: jsii.String("prefix"),
//   									SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   										DecimalSeparator: jsii.String("decimalSeparator"),
//   										ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   											Symbol: jsii.String("symbol"),
//   											Visibility: jsii.String("visibility"),
//   										},
//   									},
//   									Suffix: jsii.String("suffix"),
//   								},
//   							},
//   						},
//   					},
//   					NumericalMeasureField: &NumericalMeasureFieldProperty{
//   						Column: &ColumnIdentifierProperty{
//   							ColumnName: jsii.String("columnName"),
//   							DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   						},
//   						FieldId: jsii.String("fieldId"),
//
//   						// the properties below are optional
//   						AggregationFunction: &NumericalAggregationFunctionProperty{
//   							PercentileAggregation: &PercentileAggregationProperty{
//   								PercentileValue: jsii.Number(123),
//   							},
//   							SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   						},
//   						FormatConfiguration: &NumberFormatConfigurationProperty{
//   							FormatConfiguration: &NumericFormatConfigurationProperty{
//   								CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   									DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   										DecimalPlaces: jsii.Number(123),
//   									},
//   									NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   										DisplayMode: jsii.String("displayMode"),
//   									},
//   									NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   										NullString: jsii.String("nullString"),
//   									},
//   									NumberScale: jsii.String("numberScale"),
//   									Prefix: jsii.String("prefix"),
//   									SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   										DecimalSeparator: jsii.String("decimalSeparator"),
//   										ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   											Symbol: jsii.String("symbol"),
//   											Visibility: jsii.String("visibility"),
//   										},
//   									},
//   									Suffix: jsii.String("suffix"),
//   									Symbol: jsii.String("symbol"),
//   								},
//   								NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   									DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   										DecimalPlaces: jsii.Number(123),
//   									},
//   									NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   										DisplayMode: jsii.String("displayMode"),
//   									},
//   									NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   										NullString: jsii.String("nullString"),
//   									},
//   									NumberScale: jsii.String("numberScale"),
//   									Prefix: jsii.String("prefix"),
//   									SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   										DecimalSeparator: jsii.String("decimalSeparator"),
//   										ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   											Symbol: jsii.String("symbol"),
//   											Visibility: jsii.String("visibility"),
//   										},
//   									},
//   									Suffix: jsii.String("suffix"),
//   								},
//   								PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   									DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   										DecimalPlaces: jsii.Number(123),
//   									},
//   									NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   										DisplayMode: jsii.String("displayMode"),
//   									},
//   									NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   										NullString: jsii.String("nullString"),
//   									},
//   									Prefix: jsii.String("prefix"),
//   									SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   										DecimalSeparator: jsii.String("decimalSeparator"),
//   										ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   											Symbol: jsii.String("symbol"),
//   											Visibility: jsii.String("visibility"),
//   										},
//   									},
//   									Suffix: jsii.String("suffix"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	PaginatedReportOptions: &PivotTablePaginatedReportOptionsProperty{
//   		OverflowColumnHeaderVisibility: jsii.String("overflowColumnHeaderVisibility"),
//   		VerticalOverflowVisibility: jsii.String("verticalOverflowVisibility"),
//   	},
//   	SortConfiguration: &PivotTableSortConfigurationProperty{
//   		FieldSortOptions: []interface{}{
//   			&PivotFieldSortOptionsProperty{
//   				FieldId: jsii.String("fieldId"),
//   				SortBy: &PivotTableSortByProperty{
//   					Column: &ColumnSortProperty{
//   						Direction: jsii.String("direction"),
//   						SortBy: &ColumnIdentifierProperty{
//   							ColumnName: jsii.String("columnName"),
//   							DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   						},
//
//   						// the properties below are optional
//   						AggregationFunction: &AggregationFunctionProperty{
//   							CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   							DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   							NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   								PercentileAggregation: &PercentileAggregationProperty{
//   									PercentileValue: jsii.Number(123),
//   								},
//   								SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   							},
//   						},
//   					},
//   					DataPath: &DataPathSortProperty{
//   						Direction: jsii.String("direction"),
//   						SortPaths: []interface{}{
//   							&DataPathValueProperty{
//   								FieldId: jsii.String("fieldId"),
//   								FieldValue: jsii.String("fieldValue"),
//   							},
//   						},
//   					},
//   					Field: &FieldSortProperty{
//   						Direction: jsii.String("direction"),
//   						FieldId: jsii.String("fieldId"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	TableOptions: &PivotTableOptionsProperty{
//   		CellStyle: &TableCellStyleProperty{
//   			BackgroundColor: jsii.String("backgroundColor"),
//   			Border: &GlobalTableBorderOptionsProperty{
//   				SideSpecificBorder: &TableSideBorderOptionsProperty{
//   					Bottom: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerHorizontal: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerVertical: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Left: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Right: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Top: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   				},
//   				UniformBorder: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   			},
//   			FontConfiguration: &FontConfigurationProperty{
//   				FontColor: jsii.String("fontColor"),
//   				FontDecoration: jsii.String("fontDecoration"),
//   				FontSize: &FontSizeProperty{
//   					Relative: jsii.String("relative"),
//   				},
//   				FontStyle: jsii.String("fontStyle"),
//   				FontWeight: &FontWeightProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Height: jsii.Number(123),
//   			HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   			TextWrap: jsii.String("textWrap"),
//   			VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   			Visibility: jsii.String("visibility"),
//   		},
//   		ColumnHeaderStyle: &TableCellStyleProperty{
//   			BackgroundColor: jsii.String("backgroundColor"),
//   			Border: &GlobalTableBorderOptionsProperty{
//   				SideSpecificBorder: &TableSideBorderOptionsProperty{
//   					Bottom: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerHorizontal: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerVertical: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Left: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Right: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Top: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   				},
//   				UniformBorder: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   			},
//   			FontConfiguration: &FontConfigurationProperty{
//   				FontColor: jsii.String("fontColor"),
//   				FontDecoration: jsii.String("fontDecoration"),
//   				FontSize: &FontSizeProperty{
//   					Relative: jsii.String("relative"),
//   				},
//   				FontStyle: jsii.String("fontStyle"),
//   				FontWeight: &FontWeightProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Height: jsii.Number(123),
//   			HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   			TextWrap: jsii.String("textWrap"),
//   			VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   			Visibility: jsii.String("visibility"),
//   		},
//   		ColumnNamesVisibility: jsii.String("columnNamesVisibility"),
//   		MetricPlacement: jsii.String("metricPlacement"),
//   		RowAlternateColorOptions: &RowAlternateColorOptionsProperty{
//   			RowAlternateColors: []*string{
//   				jsii.String("rowAlternateColors"),
//   			},
//   			Status: jsii.String("status"),
//   		},
//   		RowFieldNamesStyle: &TableCellStyleProperty{
//   			BackgroundColor: jsii.String("backgroundColor"),
//   			Border: &GlobalTableBorderOptionsProperty{
//   				SideSpecificBorder: &TableSideBorderOptionsProperty{
//   					Bottom: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerHorizontal: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerVertical: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Left: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Right: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Top: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   				},
//   				UniformBorder: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   			},
//   			FontConfiguration: &FontConfigurationProperty{
//   				FontColor: jsii.String("fontColor"),
//   				FontDecoration: jsii.String("fontDecoration"),
//   				FontSize: &FontSizeProperty{
//   					Relative: jsii.String("relative"),
//   				},
//   				FontStyle: jsii.String("fontStyle"),
//   				FontWeight: &FontWeightProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Height: jsii.Number(123),
//   			HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   			TextWrap: jsii.String("textWrap"),
//   			VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   			Visibility: jsii.String("visibility"),
//   		},
//   		RowHeaderStyle: &TableCellStyleProperty{
//   			BackgroundColor: jsii.String("backgroundColor"),
//   			Border: &GlobalTableBorderOptionsProperty{
//   				SideSpecificBorder: &TableSideBorderOptionsProperty{
//   					Bottom: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerHorizontal: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerVertical: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Left: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Right: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Top: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   				},
//   				UniformBorder: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   			},
//   			FontConfiguration: &FontConfigurationProperty{
//   				FontColor: jsii.String("fontColor"),
//   				FontDecoration: jsii.String("fontDecoration"),
//   				FontSize: &FontSizeProperty{
//   					Relative: jsii.String("relative"),
//   				},
//   				FontStyle: jsii.String("fontStyle"),
//   				FontWeight: &FontWeightProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Height: jsii.Number(123),
//   			HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   			TextWrap: jsii.String("textWrap"),
//   			VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   			Visibility: jsii.String("visibility"),
//   		},
//   		SingleMetricVisibility: jsii.String("singleMetricVisibility"),
//   		ToggleButtonsVisibility: jsii.String("toggleButtonsVisibility"),
//   	},
//   	TotalOptions: &PivotTableTotalOptionsProperty{
//   		ColumnSubtotalOptions: &SubtotalOptionsProperty{
//   			CustomLabel: jsii.String("customLabel"),
//   			FieldLevel: jsii.String("fieldLevel"),
//   			FieldLevelOptions: []interface{}{
//   				&PivotTableFieldSubtotalOptionsProperty{
//   					FieldId: jsii.String("fieldId"),
//   				},
//   			},
//   			MetricHeaderCellStyle: &TableCellStyleProperty{
//   				BackgroundColor: jsii.String("backgroundColor"),
//   				Border: &GlobalTableBorderOptionsProperty{
//   					SideSpecificBorder: &TableSideBorderOptionsProperty{
//   						Bottom: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						InnerHorizontal: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						InnerVertical: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Left: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Right: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Top: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   					},
//   					UniformBorder: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   				},
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontSize: &FontSizeProperty{
//   						Relative: jsii.String("relative"),
//   					},
//   					FontStyle: jsii.String("fontStyle"),
//   					FontWeight: &FontWeightProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Height: jsii.Number(123),
//   				HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   				TextWrap: jsii.String("textWrap"),
//   				VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   			TotalCellStyle: &TableCellStyleProperty{
//   				BackgroundColor: jsii.String("backgroundColor"),
//   				Border: &GlobalTableBorderOptionsProperty{
//   					SideSpecificBorder: &TableSideBorderOptionsProperty{
//   						Bottom: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						InnerHorizontal: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						InnerVertical: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Left: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Right: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Top: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   					},
//   					UniformBorder: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   				},
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontSize: &FontSizeProperty{
//   						Relative: jsii.String("relative"),
//   					},
//   					FontStyle: jsii.String("fontStyle"),
//   					FontWeight: &FontWeightProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Height: jsii.Number(123),
//   				HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   				TextWrap: jsii.String("textWrap"),
//   				VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   			TotalsVisibility: jsii.String("totalsVisibility"),
//   			ValueCellStyle: &TableCellStyleProperty{
//   				BackgroundColor: jsii.String("backgroundColor"),
//   				Border: &GlobalTableBorderOptionsProperty{
//   					SideSpecificBorder: &TableSideBorderOptionsProperty{
//   						Bottom: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						InnerHorizontal: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						InnerVertical: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Left: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Right: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Top: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   					},
//   					UniformBorder: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   				},
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontSize: &FontSizeProperty{
//   						Relative: jsii.String("relative"),
//   					},
//   					FontStyle: jsii.String("fontStyle"),
//   					FontWeight: &FontWeightProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Height: jsii.Number(123),
//   				HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   				TextWrap: jsii.String("textWrap"),
//   				VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   		},
//   		ColumnTotalOptions: &PivotTotalOptionsProperty{
//   			CustomLabel: jsii.String("customLabel"),
//   			MetricHeaderCellStyle: &TableCellStyleProperty{
//   				BackgroundColor: jsii.String("backgroundColor"),
//   				Border: &GlobalTableBorderOptionsProperty{
//   					SideSpecificBorder: &TableSideBorderOptionsProperty{
//   						Bottom: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						InnerHorizontal: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						InnerVertical: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Left: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Right: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Top: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   					},
//   					UniformBorder: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   				},
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontSize: &FontSizeProperty{
//   						Relative: jsii.String("relative"),
//   					},
//   					FontStyle: jsii.String("fontStyle"),
//   					FontWeight: &FontWeightProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Height: jsii.Number(123),
//   				HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   				TextWrap: jsii.String("textWrap"),
//   				VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   			Placement: jsii.String("placement"),
//   			ScrollStatus: jsii.String("scrollStatus"),
//   			TotalCellStyle: &TableCellStyleProperty{
//   				BackgroundColor: jsii.String("backgroundColor"),
//   				Border: &GlobalTableBorderOptionsProperty{
//   					SideSpecificBorder: &TableSideBorderOptionsProperty{
//   						Bottom: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						InnerHorizontal: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						InnerVertical: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Left: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Right: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Top: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   					},
//   					UniformBorder: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   				},
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontSize: &FontSizeProperty{
//   						Relative: jsii.String("relative"),
//   					},
//   					FontStyle: jsii.String("fontStyle"),
//   					FontWeight: &FontWeightProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Height: jsii.Number(123),
//   				HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   				TextWrap: jsii.String("textWrap"),
//   				VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   			TotalsVisibility: jsii.String("totalsVisibility"),
//   			ValueCellStyle: &TableCellStyleProperty{
//   				BackgroundColor: jsii.String("backgroundColor"),
//   				Border: &GlobalTableBorderOptionsProperty{
//   					SideSpecificBorder: &TableSideBorderOptionsProperty{
//   						Bottom: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						InnerHorizontal: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						InnerVertical: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Left: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Right: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Top: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   					},
//   					UniformBorder: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   				},
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontSize: &FontSizeProperty{
//   						Relative: jsii.String("relative"),
//   					},
//   					FontStyle: jsii.String("fontStyle"),
//   					FontWeight: &FontWeightProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Height: jsii.Number(123),
//   				HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   				TextWrap: jsii.String("textWrap"),
//   				VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   		},
//   		RowSubtotalOptions: &SubtotalOptionsProperty{
//   			CustomLabel: jsii.String("customLabel"),
//   			FieldLevel: jsii.String("fieldLevel"),
//   			FieldLevelOptions: []interface{}{
//   				&PivotTableFieldSubtotalOptionsProperty{
//   					FieldId: jsii.String("fieldId"),
//   				},
//   			},
//   			MetricHeaderCellStyle: &TableCellStyleProperty{
//   				BackgroundColor: jsii.String("backgroundColor"),
//   				Border: &GlobalTableBorderOptionsProperty{
//   					SideSpecificBorder: &TableSideBorderOptionsProperty{
//   						Bottom: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						InnerHorizontal: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						InnerVertical: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Left: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Right: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Top: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   					},
//   					UniformBorder: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   				},
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontSize: &FontSizeProperty{
//   						Relative: jsii.String("relative"),
//   					},
//   					FontStyle: jsii.String("fontStyle"),
//   					FontWeight: &FontWeightProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Height: jsii.Number(123),
//   				HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   				TextWrap: jsii.String("textWrap"),
//   				VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   			TotalCellStyle: &TableCellStyleProperty{
//   				BackgroundColor: jsii.String("backgroundColor"),
//   				Border: &GlobalTableBorderOptionsProperty{
//   					SideSpecificBorder: &TableSideBorderOptionsProperty{
//   						Bottom: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						InnerHorizontal: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						InnerVertical: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Left: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Right: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Top: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   					},
//   					UniformBorder: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   				},
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontSize: &FontSizeProperty{
//   						Relative: jsii.String("relative"),
//   					},
//   					FontStyle: jsii.String("fontStyle"),
//   					FontWeight: &FontWeightProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Height: jsii.Number(123),
//   				HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   				TextWrap: jsii.String("textWrap"),
//   				VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   			TotalsVisibility: jsii.String("totalsVisibility"),
//   			ValueCellStyle: &TableCellStyleProperty{
//   				BackgroundColor: jsii.String("backgroundColor"),
//   				Border: &GlobalTableBorderOptionsProperty{
//   					SideSpecificBorder: &TableSideBorderOptionsProperty{
//   						Bottom: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						InnerHorizontal: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						InnerVertical: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Left: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Right: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Top: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   					},
//   					UniformBorder: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   				},
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontSize: &FontSizeProperty{
//   						Relative: jsii.String("relative"),
//   					},
//   					FontStyle: jsii.String("fontStyle"),
//   					FontWeight: &FontWeightProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Height: jsii.Number(123),
//   				HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   				TextWrap: jsii.String("textWrap"),
//   				VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   		},
//   		RowTotalOptions: &PivotTotalOptionsProperty{
//   			CustomLabel: jsii.String("customLabel"),
//   			MetricHeaderCellStyle: &TableCellStyleProperty{
//   				BackgroundColor: jsii.String("backgroundColor"),
//   				Border: &GlobalTableBorderOptionsProperty{
//   					SideSpecificBorder: &TableSideBorderOptionsProperty{
//   						Bottom: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						InnerHorizontal: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						InnerVertical: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Left: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Right: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Top: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   					},
//   					UniformBorder: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   				},
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontSize: &FontSizeProperty{
//   						Relative: jsii.String("relative"),
//   					},
//   					FontStyle: jsii.String("fontStyle"),
//   					FontWeight: &FontWeightProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Height: jsii.Number(123),
//   				HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   				TextWrap: jsii.String("textWrap"),
//   				VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   			Placement: jsii.String("placement"),
//   			ScrollStatus: jsii.String("scrollStatus"),
//   			TotalCellStyle: &TableCellStyleProperty{
//   				BackgroundColor: jsii.String("backgroundColor"),
//   				Border: &GlobalTableBorderOptionsProperty{
//   					SideSpecificBorder: &TableSideBorderOptionsProperty{
//   						Bottom: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						InnerHorizontal: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						InnerVertical: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Left: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Right: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Top: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   					},
//   					UniformBorder: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   				},
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontSize: &FontSizeProperty{
//   						Relative: jsii.String("relative"),
//   					},
//   					FontStyle: jsii.String("fontStyle"),
//   					FontWeight: &FontWeightProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Height: jsii.Number(123),
//   				HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   				TextWrap: jsii.String("textWrap"),
//   				VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   			TotalsVisibility: jsii.String("totalsVisibility"),
//   			ValueCellStyle: &TableCellStyleProperty{
//   				BackgroundColor: jsii.String("backgroundColor"),
//   				Border: &GlobalTableBorderOptionsProperty{
//   					SideSpecificBorder: &TableSideBorderOptionsProperty{
//   						Bottom: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						InnerHorizontal: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						InnerVertical: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Left: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Right: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   						Top: &TableBorderOptionsProperty{
//   							Color: jsii.String("color"),
//   							Style: jsii.String("style"),
//   							Thickness: jsii.Number(123),
//   						},
//   					},
//   					UniformBorder: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   				},
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontSize: &FontSizeProperty{
//   						Relative: jsii.String("relative"),
//   					},
//   					FontStyle: jsii.String("fontStyle"),
//   					FontWeight: &FontWeightProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Height: jsii.Number(123),
//   				HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   				TextWrap: jsii.String("textWrap"),
//   				VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   		},
//   	},
//   }
//
type CfnDashboard_PivotTableConfigurationProperty struct {
	// The field options for a pivot table visual.
	FieldOptions interface{} `field:"optional" json:"fieldOptions" yaml:"fieldOptions"`
	// The field wells of the visual.
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The paginated report options for a pivot table visual.
	PaginatedReportOptions interface{} `field:"optional" json:"paginatedReportOptions" yaml:"paginatedReportOptions"`
	// The sort configuration for a `PivotTableVisual` .
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// The table options for a pivot table visual.
	TableOptions interface{} `field:"optional" json:"tableOptions" yaml:"tableOptions"`
	// The total options for a pivot table visual.
	TotalOptions interface{} `field:"optional" json:"totalOptions" yaml:"totalOptions"`
}

