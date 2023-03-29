package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataDriven interface{}
//
//   histogramConfigurationProperty := &HistogramConfigurationProperty{
//   	BinOptions: &HistogramBinOptionsProperty{
//   		BinCount: &BinCountOptionsProperty{
//   			Value: jsii.Number(123),
//   		},
//   		BinWidth: &BinWidthOptionsProperty{
//   			BinCountLimit: jsii.Number(123),
//   			Value: jsii.Number(123),
//   		},
//   		SelectedBinType: jsii.String("selectedBinType"),
//   		StartValue: jsii.Number(123),
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
//   	FieldWells: &HistogramFieldWellsProperty{
//   		HistogramAggregatedFieldWells: &HistogramAggregatedFieldWellsProperty{
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
//   	XAxisDisplayOptions: &AxisDisplayOptionsProperty{
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
//   	XAxisLabelOptions: &ChartAxisLabelOptionsProperty{
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
//   	YAxisDisplayOptions: &AxisDisplayOptionsProperty{
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
//   }
//
type CfnTemplate_HistogramConfigurationProperty struct {
	// `CfnTemplate.HistogramConfigurationProperty.BinOptions`.
	BinOptions interface{} `field:"optional" json:"binOptions" yaml:"binOptions"`
	// `CfnTemplate.HistogramConfigurationProperty.DataLabels`.
	DataLabels interface{} `field:"optional" json:"dataLabels" yaml:"dataLabels"`
	// `CfnTemplate.HistogramConfigurationProperty.FieldWells`.
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// `CfnTemplate.HistogramConfigurationProperty.Tooltip`.
	Tooltip interface{} `field:"optional" json:"tooltip" yaml:"tooltip"`
	// `CfnTemplate.HistogramConfigurationProperty.VisualPalette`.
	VisualPalette interface{} `field:"optional" json:"visualPalette" yaml:"visualPalette"`
	// `CfnTemplate.HistogramConfigurationProperty.XAxisDisplayOptions`.
	XAxisDisplayOptions interface{} `field:"optional" json:"xAxisDisplayOptions" yaml:"xAxisDisplayOptions"`
	// `CfnTemplate.HistogramConfigurationProperty.XAxisLabelOptions`.
	XAxisLabelOptions interface{} `field:"optional" json:"xAxisLabelOptions" yaml:"xAxisLabelOptions"`
	// `CfnTemplate.HistogramConfigurationProperty.YAxisDisplayOptions`.
	YAxisDisplayOptions interface{} `field:"optional" json:"yAxisDisplayOptions" yaml:"yAxisDisplayOptions"`
}

