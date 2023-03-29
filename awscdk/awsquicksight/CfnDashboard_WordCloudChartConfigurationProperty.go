package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   wordCloudChartConfigurationProperty := &WordCloudChartConfigurationProperty{
//   	CategoryLabelOptions: &ChartAxisLabelOptionsProperty{
//   		AxisLabelOptions: []interface{}{
//   			&AxisLabelOptionsProperty{
//   				ApplyTo: &AxisLabelReferenceOptionsProperty{
//   					Column: &ColumnIdentifierProperty{
//   						ColumnName: jsii.String("columnName"),
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//   					FieldId: jsii.String("fieldId"),
//   				},
//   				CustomLabel: jsii.String("customLabel"),
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
//   			},
//   		},
//   		SortIconVisibility: jsii.String("sortIconVisibility"),
//   		Visibility: jsii.String("visibility"),
//   	},
//   	FieldWells: &WordCloudFieldWellsProperty{
//   		WordCloudAggregatedFieldWells: &WordCloudAggregatedFieldWellsProperty{
//   			GroupBy: []interface{}{
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
//   			Size: []interface{}{
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
//   	SortConfiguration: &WordCloudSortConfigurationProperty{
//   		CategoryItemsLimit: &ItemsLimitConfigurationProperty{
//   			ItemsLimit: jsii.Number(123),
//   			OtherCategories: jsii.String("otherCategories"),
//   		},
//   		CategorySort: []interface{}{
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
//   	WordCloudOptions: &WordCloudOptionsProperty{
//   		CloudLayout: jsii.String("cloudLayout"),
//   		MaximumStringLength: jsii.Number(123),
//   		WordCasing: jsii.String("wordCasing"),
//   		WordOrientation: jsii.String("wordOrientation"),
//   		WordPadding: jsii.String("wordPadding"),
//   		WordScaling: jsii.String("wordScaling"),
//   	},
//   }
//
type CfnDashboard_WordCloudChartConfigurationProperty struct {
	// `CfnDashboard.WordCloudChartConfigurationProperty.CategoryLabelOptions`.
	CategoryLabelOptions interface{} `field:"optional" json:"categoryLabelOptions" yaml:"categoryLabelOptions"`
	// `CfnDashboard.WordCloudChartConfigurationProperty.FieldWells`.
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// `CfnDashboard.WordCloudChartConfigurationProperty.SortConfiguration`.
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// `CfnDashboard.WordCloudChartConfigurationProperty.WordCloudOptions`.
	WordCloudOptions interface{} `field:"optional" json:"wordCloudOptions" yaml:"wordCloudOptions"`
}

