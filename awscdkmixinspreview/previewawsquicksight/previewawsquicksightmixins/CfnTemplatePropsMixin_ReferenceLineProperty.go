package previewawsquicksightmixins


// The reference line visual display options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   referenceLineProperty := &ReferenceLineProperty{
//   	DataConfiguration: &ReferenceLineDataConfigurationProperty{
//   		AxisBinding: jsii.String("axisBinding"),
//   		DynamicConfiguration: &ReferenceLineDynamicDataConfigurationProperty{
//   			Calculation: &NumericalAggregationFunctionProperty{
//   				PercentileAggregation: &PercentileAggregationProperty{
//   					PercentileValue: jsii.Number(123),
//   				},
//   				SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   			},
//   			Column: &ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   			MeasureAggregationFunction: &AggregationFunctionProperty{
//   				AttributeAggregationFunction: &AttributeAggregationFunctionProperty{
//   					SimpleAttributeAggregation: jsii.String("simpleAttributeAggregation"),
//   					ValueForMultipleValues: jsii.String("valueForMultipleValues"),
//   				},
//   				CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   				DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   				NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   					PercentileAggregation: &PercentileAggregationProperty{
//   						PercentileValue: jsii.Number(123),
//   					},
//   					SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   				},
//   			},
//   		},
//   		SeriesType: jsii.String("seriesType"),
//   		StaticConfiguration: &ReferenceLineStaticDataConfigurationProperty{
//   			Value: jsii.Number(123),
//   		},
//   	},
//   	LabelConfiguration: &ReferenceLineLabelConfigurationProperty{
//   		CustomLabelConfiguration: &ReferenceLineCustomLabelConfigurationProperty{
//   			CustomLabel: jsii.String("customLabel"),
//   		},
//   		FontColor: jsii.String("fontColor"),
//   		FontConfiguration: &FontConfigurationProperty{
//   			FontColor: jsii.String("fontColor"),
//   			FontDecoration: jsii.String("fontDecoration"),
//   			FontFamily: jsii.String("fontFamily"),
//   			FontSize: &FontSizeProperty{
//   				Absolute: jsii.String("absolute"),
//   				Relative: jsii.String("relative"),
//   			},
//   			FontStyle: jsii.String("fontStyle"),
//   			FontWeight: &FontWeightProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		HorizontalPosition: jsii.String("horizontalPosition"),
//   		ValueLabelConfiguration: &ReferenceLineValueLabelConfigurationProperty{
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
//   			RelativePosition: jsii.String("relativePosition"),
//   		},
//   		VerticalPosition: jsii.String("verticalPosition"),
//   	},
//   	Status: jsii.String("status"),
//   	StyleConfiguration: &ReferenceLineStyleConfigurationProperty{
//   		Color: jsii.String("color"),
//   		Pattern: jsii.String("pattern"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referenceline.html
//
type CfnTemplatePropsMixin_ReferenceLineProperty struct {
	// The data configuration of the reference line.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referenceline.html#cfn-quicksight-template-referenceline-dataconfiguration
	//
	DataConfiguration interface{} `field:"optional" json:"dataConfiguration" yaml:"dataConfiguration"`
	// The label configuration of the reference line.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referenceline.html#cfn-quicksight-template-referenceline-labelconfiguration
	//
	LabelConfiguration interface{} `field:"optional" json:"labelConfiguration" yaml:"labelConfiguration"`
	// The status of the reference line. Choose one of the following options:.
	//
	// - `ENABLE`
	// - `DISABLE`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referenceline.html#cfn-quicksight-template-referenceline-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The style configuration of the reference line.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referenceline.html#cfn-quicksight-template-referenceline-styleconfiguration
	//
	StyleConfiguration interface{} `field:"optional" json:"styleConfiguration" yaml:"styleConfiguration"`
}

