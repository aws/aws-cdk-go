package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sankeyDiagramChartConfigurationProperty := &SankeyDiagramChartConfigurationProperty{
//   	DataLabels: &DataLabelOptionsProperty{
//   		CategoryLabelVisibility: jsii.String("categoryLabelVisibility"),
//   		DataLabelTypes: []interface{}{
//   			&DataLabelTypeProperty{
//   				DataPathLabelType: &DataPathLabelTypeProperty{
//   					FieldId: jsii.String("fieldId"),
//   					FieldValue: jsii.String("fieldValue"),
//   					Visibility: jsii.String("visibility"),
//   				},
//   				FieldLabelType: &FieldLabelTypeProperty{
//   					FieldId: jsii.String("fieldId"),
//   					Visibility: jsii.String("visibility"),
//   				},
//   				MaximumLabelType: &MaximumLabelTypeProperty{
//   					Visibility: jsii.String("visibility"),
//   				},
//   				MinimumLabelType: &MinimumLabelTypeProperty{
//   					Visibility: jsii.String("visibility"),
//   				},
//   				RangeEndsLabelType: &RangeEndsLabelTypeProperty{
//   					Visibility: jsii.String("visibility"),
//   				},
//   			},
//   		},
//   		LabelColor: jsii.String("labelColor"),
//   		LabelContent: jsii.String("labelContent"),
//   		LabelFontConfiguration: &FontConfigurationProperty{
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
//   		MeasureLabelVisibility: jsii.String("measureLabelVisibility"),
//   		Overlap: jsii.String("overlap"),
//   		Position: jsii.String("position"),
//   		Visibility: jsii.String("visibility"),
//   	},
//   	FieldWells: &SankeyDiagramFieldWellsProperty{
//   		SankeyDiagramAggregatedFieldWells: &SankeyDiagramAggregatedFieldWellsProperty{
//   			Destination: []interface{}{
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
//   			Source: []interface{}{
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
//   			Weight: []interface{}{
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
//   	SortConfiguration: &SankeyDiagramSortConfigurationProperty{
//   		DestinationItemsLimit: &ItemsLimitConfigurationProperty{
//   			ItemsLimit: jsii.Number(123),
//   			OtherCategories: jsii.String("otherCategories"),
//   		},
//   		SourceItemsLimit: &ItemsLimitConfigurationProperty{
//   			ItemsLimit: jsii.Number(123),
//   			OtherCategories: jsii.String("otherCategories"),
//   		},
//   		WeightSort: []interface{}{
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
type CfnAnalysis_SankeyDiagramChartConfigurationProperty struct {
	// `CfnAnalysis.SankeyDiagramChartConfigurationProperty.DataLabels`.
	DataLabels interface{} `field:"optional" json:"dataLabels" yaml:"dataLabels"`
	// `CfnAnalysis.SankeyDiagramChartConfigurationProperty.FieldWells`.
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// `CfnAnalysis.SankeyDiagramChartConfigurationProperty.SortConfiguration`.
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
}

