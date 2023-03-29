package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataDriven interface{}
//
//   barChartConfigurationProperty := &BarChartConfigurationProperty{
//   	BarsArrangement: jsii.String("barsArrangement"),
//   	CategoryAxis: &AxisDisplayOptionsProperty{
//   		AxisLineVisibility: jsii.String("axisLineVisibility"),
//   		AxisOffset: jsii.String("axisOffset"),
//   		DataOptions: &AxisDataOptionsProperty{
//   			DateAxisOptions: &DateAxisOptionsProperty{
//   				MissingDateVisibility: jsii.String("missingDateVisibility"),
//   			},
//   			NumericAxisOptions: &NumericAxisOptionsProperty{
//   				Range: &AxisDisplayRangeProperty{
//   					DataDriven: dataDriven,
//   					MinMax: &AxisDisplayMinMaxRangeProperty{
//   						Maximum: jsii.Number(123),
//   						Minimum: jsii.Number(123),
//   					},
//   				},
//   				Scale: &AxisScaleProperty{
//   					Linear: &AxisLinearScaleProperty{
//   						StepCount: jsii.Number(123),
//   						StepSize: jsii.Number(123),
//   					},
//   					Logarithmic: &AxisLogarithmicScaleProperty{
//   						Base: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   		GridLineVisibility: jsii.String("gridLineVisibility"),
//   		ScrollbarOptions: &ScrollBarOptionsProperty{
//   			Visibility: jsii.String("visibility"),
//   			VisibleRange: &VisibleRangeOptionsProperty{
//   				PercentRange: &PercentVisibleRangeProperty{
//   					From: jsii.Number(123),
//   					To: jsii.Number(123),
//   				},
//   			},
//   		},
//   		TickLabelOptions: &AxisTickLabelOptionsProperty{
//   			LabelOptions: &LabelOptionsProperty{
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
//   				Visibility: jsii.String("visibility"),
//   			},
//   			RotationAngle: jsii.Number(123),
//   		},
//   	},
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
//   	ColorLabelOptions: &ChartAxisLabelOptionsProperty{
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
//   	ContributionAnalysisDefaults: []interface{}{
//   		&ContributionAnalysisDefaultProperty{
//   			ContributorDimensions: []interface{}{
//   				&ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   			},
//   			MeasureFieldId: jsii.String("measureFieldId"),
//   		},
//   	},
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
//   	FieldWells: &BarChartFieldWellsProperty{
//   		BarChartAggregatedFieldWells: &BarChartAggregatedFieldWellsProperty{
//   			Category: []interface{}{
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
//   			Colors: []interface{}{
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
//   			SmallMultiples: []interface{}{
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
//   	Legend: &LegendOptionsProperty{
//   		Height: jsii.String("height"),
//   		Position: jsii.String("position"),
//   		Title: &LabelOptionsProperty{
//   			CustomLabel: jsii.String("customLabel"),
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
//   			Visibility: jsii.String("visibility"),
//   		},
//   		Visibility: jsii.String("visibility"),
//   		Width: jsii.String("width"),
//   	},
//   	Orientation: jsii.String("orientation"),
//   	ReferenceLines: []interface{}{
//   		&ReferenceLineProperty{
//   			DataConfiguration: &ReferenceLineDataConfigurationProperty{
//   				AxisBinding: jsii.String("axisBinding"),
//   				DynamicConfiguration: &ReferenceLineDynamicDataConfigurationProperty{
//   					Calculation: &NumericalAggregationFunctionProperty{
//   						PercentileAggregation: &PercentileAggregationProperty{
//   							PercentileValue: jsii.Number(123),
//   						},
//   						SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   					},
//   					Column: &ColumnIdentifierProperty{
//   						ColumnName: jsii.String("columnName"),
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//   					MeasureAggregationFunction: &AggregationFunctionProperty{
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
//   				StaticConfiguration: &ReferenceLineStaticDataConfigurationProperty{
//   					Value: jsii.Number(123),
//   				},
//   			},
//
//   			// the properties below are optional
//   			LabelConfiguration: &ReferenceLineLabelConfigurationProperty{
//   				CustomLabelConfiguration: &ReferenceLineCustomLabelConfigurationProperty{
//   					CustomLabel: jsii.String("customLabel"),
//   				},
//   				FontColor: jsii.String("fontColor"),
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
//   				HorizontalPosition: jsii.String("horizontalPosition"),
//   				ValueLabelConfiguration: &ReferenceLineValueLabelConfigurationProperty{
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
//   					RelativePosition: jsii.String("relativePosition"),
//   				},
//   				VerticalPosition: jsii.String("verticalPosition"),
//   			},
//   			Status: jsii.String("status"),
//   			StyleConfiguration: &ReferenceLineStyleConfigurationProperty{
//   				Color: jsii.String("color"),
//   				Pattern: jsii.String("pattern"),
//   			},
//   		},
//   	},
//   	SmallMultiplesOptions: &SmallMultiplesOptionsProperty{
//   		MaxVisibleColumns: jsii.Number(123),
//   		MaxVisibleRows: jsii.Number(123),
//   		PanelConfiguration: &PanelConfigurationProperty{
//   			BackgroundColor: jsii.String("backgroundColor"),
//   			BackgroundVisibility: jsii.String("backgroundVisibility"),
//   			BorderColor: jsii.String("borderColor"),
//   			BorderStyle: jsii.String("borderStyle"),
//   			BorderThickness: jsii.String("borderThickness"),
//   			BorderVisibility: jsii.String("borderVisibility"),
//   			GutterSpacing: jsii.String("gutterSpacing"),
//   			GutterVisibility: jsii.String("gutterVisibility"),
//   			Title: &PanelTitleOptionsProperty{
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
//   				HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   		},
//   	},
//   	SortConfiguration: &BarChartSortConfigurationProperty{
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
//   		ColorItemsLimit: &ItemsLimitConfigurationProperty{
//   			ItemsLimit: jsii.Number(123),
//   			OtherCategories: jsii.String("otherCategories"),
//   		},
//   		ColorSort: []interface{}{
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
//   		SmallMultiplesLimitConfiguration: &ItemsLimitConfigurationProperty{
//   			ItemsLimit: jsii.Number(123),
//   			OtherCategories: jsii.String("otherCategories"),
//   		},
//   		SmallMultiplesSort: []interface{}{
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
//   	Tooltip: &TooltipOptionsProperty{
//   		FieldBasedTooltip: &FieldBasedTooltipProperty{
//   			AggregationVisibility: jsii.String("aggregationVisibility"),
//   			TooltipFields: []interface{}{
//   				&TooltipItemProperty{
//   					ColumnTooltipItem: &ColumnTooltipItemProperty{
//   						Column: &ColumnIdentifierProperty{
//   							ColumnName: jsii.String("columnName"),
//   							DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   						},
//
//   						// the properties below are optional
//   						Aggregation: &AggregationFunctionProperty{
//   							CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   							DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   							NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   								PercentileAggregation: &PercentileAggregationProperty{
//   									PercentileValue: jsii.Number(123),
//   								},
//   								SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   							},
//   						},
//   						Label: jsii.String("label"),
//   						Visibility: jsii.String("visibility"),
//   					},
//   					FieldTooltipItem: &FieldTooltipItemProperty{
//   						FieldId: jsii.String("fieldId"),
//
//   						// the properties below are optional
//   						Label: jsii.String("label"),
//   						Visibility: jsii.String("visibility"),
//   					},
//   				},
//   			},
//   			TooltipTitleType: jsii.String("tooltipTitleType"),
//   		},
//   		SelectedTooltipType: jsii.String("selectedTooltipType"),
//   		TooltipVisibility: jsii.String("tooltipVisibility"),
//   	},
//   	ValueAxis: &AxisDisplayOptionsProperty{
//   		AxisLineVisibility: jsii.String("axisLineVisibility"),
//   		AxisOffset: jsii.String("axisOffset"),
//   		DataOptions: &AxisDataOptionsProperty{
//   			DateAxisOptions: &DateAxisOptionsProperty{
//   				MissingDateVisibility: jsii.String("missingDateVisibility"),
//   			},
//   			NumericAxisOptions: &NumericAxisOptionsProperty{
//   				Range: &AxisDisplayRangeProperty{
//   					DataDriven: dataDriven,
//   					MinMax: &AxisDisplayMinMaxRangeProperty{
//   						Maximum: jsii.Number(123),
//   						Minimum: jsii.Number(123),
//   					},
//   				},
//   				Scale: &AxisScaleProperty{
//   					Linear: &AxisLinearScaleProperty{
//   						StepCount: jsii.Number(123),
//   						StepSize: jsii.Number(123),
//   					},
//   					Logarithmic: &AxisLogarithmicScaleProperty{
//   						Base: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   		GridLineVisibility: jsii.String("gridLineVisibility"),
//   		ScrollbarOptions: &ScrollBarOptionsProperty{
//   			Visibility: jsii.String("visibility"),
//   			VisibleRange: &VisibleRangeOptionsProperty{
//   				PercentRange: &PercentVisibleRangeProperty{
//   					From: jsii.Number(123),
//   					To: jsii.Number(123),
//   				},
//   			},
//   		},
//   		TickLabelOptions: &AxisTickLabelOptionsProperty{
//   			LabelOptions: &LabelOptionsProperty{
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
//   				Visibility: jsii.String("visibility"),
//   			},
//   			RotationAngle: jsii.Number(123),
//   		},
//   	},
//   	ValueLabelOptions: &ChartAxisLabelOptionsProperty{
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
//   	VisualPalette: &VisualPaletteProperty{
//   		ChartColor: jsii.String("chartColor"),
//   		ColorMap: []interface{}{
//   			&DataPathColorProperty{
//   				Color: jsii.String("color"),
//   				Element: &DataPathValueProperty{
//   					FieldId: jsii.String("fieldId"),
//   					FieldValue: jsii.String("fieldValue"),
//   				},
//
//   				// the properties below are optional
//   				TimeGranularity: jsii.String("timeGranularity"),
//   			},
//   		},
//   	},
//   }
//
type CfnDashboard_BarChartConfigurationProperty struct {
	// `CfnDashboard.BarChartConfigurationProperty.BarsArrangement`.
	BarsArrangement *string `field:"optional" json:"barsArrangement" yaml:"barsArrangement"`
	// `CfnDashboard.BarChartConfigurationProperty.CategoryAxis`.
	CategoryAxis interface{} `field:"optional" json:"categoryAxis" yaml:"categoryAxis"`
	// `CfnDashboard.BarChartConfigurationProperty.CategoryLabelOptions`.
	CategoryLabelOptions interface{} `field:"optional" json:"categoryLabelOptions" yaml:"categoryLabelOptions"`
	// `CfnDashboard.BarChartConfigurationProperty.ColorLabelOptions`.
	ColorLabelOptions interface{} `field:"optional" json:"colorLabelOptions" yaml:"colorLabelOptions"`
	// `CfnDashboard.BarChartConfigurationProperty.ContributionAnalysisDefaults`.
	ContributionAnalysisDefaults interface{} `field:"optional" json:"contributionAnalysisDefaults" yaml:"contributionAnalysisDefaults"`
	// `CfnDashboard.BarChartConfigurationProperty.DataLabels`.
	DataLabels interface{} `field:"optional" json:"dataLabels" yaml:"dataLabels"`
	// `CfnDashboard.BarChartConfigurationProperty.FieldWells`.
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// `CfnDashboard.BarChartConfigurationProperty.Legend`.
	Legend interface{} `field:"optional" json:"legend" yaml:"legend"`
	// `CfnDashboard.BarChartConfigurationProperty.Orientation`.
	Orientation *string `field:"optional" json:"orientation" yaml:"orientation"`
	// `CfnDashboard.BarChartConfigurationProperty.ReferenceLines`.
	ReferenceLines interface{} `field:"optional" json:"referenceLines" yaml:"referenceLines"`
	// `CfnDashboard.BarChartConfigurationProperty.SmallMultiplesOptions`.
	SmallMultiplesOptions interface{} `field:"optional" json:"smallMultiplesOptions" yaml:"smallMultiplesOptions"`
	// `CfnDashboard.BarChartConfigurationProperty.SortConfiguration`.
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// `CfnDashboard.BarChartConfigurationProperty.Tooltip`.
	Tooltip interface{} `field:"optional" json:"tooltip" yaml:"tooltip"`
	// `CfnDashboard.BarChartConfigurationProperty.ValueAxis`.
	ValueAxis interface{} `field:"optional" json:"valueAxis" yaml:"valueAxis"`
	// `CfnDashboard.BarChartConfigurationProperty.ValueLabelOptions`.
	ValueLabelOptions interface{} `field:"optional" json:"valueLabelOptions" yaml:"valueLabelOptions"`
	// `CfnDashboard.BarChartConfigurationProperty.VisualPalette`.
	VisualPalette interface{} `field:"optional" json:"visualPalette" yaml:"visualPalette"`
}

