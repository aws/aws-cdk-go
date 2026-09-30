package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sparklinesOptionsProperty := &SparklinesOptionsProperty{
//   	FieldId: jsii.String("fieldId"),
//   	XAxisField: &DimensionFieldProperty{
//   		CategoricalDimensionField: &CategoricalDimensionFieldProperty{
//   			Column: &ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//
//   				// the properties below are optional
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				TopicIdentifier: jsii.String("topicIdentifier"),
//   			},
//   			FieldId: jsii.String("fieldId"),
//
//   			// the properties below are optional
//   			FormatConfiguration: &StringFormatConfigurationProperty{
//   				NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   					NullString: jsii.String("nullString"),
//   				},
//   				NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   					CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						NumberScale: jsii.String("numberScale"),
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								GroupingStyle: jsii.String("groupingStyle"),
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   						Symbol: jsii.String("symbol"),
//   					},
//   					NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						NumberScale: jsii.String("numberScale"),
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								GroupingStyle: jsii.String("groupingStyle"),
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   					},
//   					PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								GroupingStyle: jsii.String("groupingStyle"),
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   					},
//   				},
//   			},
//   			HierarchyId: jsii.String("hierarchyId"),
//   		},
//   		DateDimensionField: &DateDimensionFieldProperty{
//   			Column: &ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//
//   				// the properties below are optional
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				TopicIdentifier: jsii.String("topicIdentifier"),
//   			},
//   			FieldId: jsii.String("fieldId"),
//
//   			// the properties below are optional
//   			DateGranularity: jsii.String("dateGranularity"),
//   			FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   				DateTimeFormat: jsii.String("dateTimeFormat"),
//   				NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   					NullString: jsii.String("nullString"),
//   				},
//   				NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   					CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						NumberScale: jsii.String("numberScale"),
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								GroupingStyle: jsii.String("groupingStyle"),
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   						Symbol: jsii.String("symbol"),
//   					},
//   					NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						NumberScale: jsii.String("numberScale"),
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								GroupingStyle: jsii.String("groupingStyle"),
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   					},
//   					PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								GroupingStyle: jsii.String("groupingStyle"),
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   					},
//   				},
//   			},
//   			HierarchyId: jsii.String("hierarchyId"),
//   		},
//   		NumericalDimensionField: &NumericalDimensionFieldProperty{
//   			Column: &ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//
//   				// the properties below are optional
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				TopicIdentifier: jsii.String("topicIdentifier"),
//   			},
//   			FieldId: jsii.String("fieldId"),
//
//   			// the properties below are optional
//   			FormatConfiguration: &NumberFormatConfigurationProperty{
//   				FormatConfiguration: &NumericFormatConfigurationProperty{
//   					CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						NumberScale: jsii.String("numberScale"),
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								GroupingStyle: jsii.String("groupingStyle"),
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   						Symbol: jsii.String("symbol"),
//   					},
//   					NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						NumberScale: jsii.String("numberScale"),
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								GroupingStyle: jsii.String("groupingStyle"),
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   					},
//   					PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								GroupingStyle: jsii.String("groupingStyle"),
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   					},
//   				},
//   			},
//   			HierarchyId: jsii.String("hierarchyId"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	AllPointsMarker: &LineChartMarkerStyleSettingsProperty{
//   		MarkerColor: jsii.String("markerColor"),
//   		MarkerShape: jsii.String("markerShape"),
//   		MarkerSize: jsii.String("markerSize"),
//   		MarkerVisibility: jsii.String("markerVisibility"),
//   	},
//   	LineColor: jsii.String("lineColor"),
//   	LineInterpolation: jsii.String("lineInterpolation"),
//   	MaxValueMarker: &LineChartMarkerStyleSettingsProperty{
//   		MarkerColor: jsii.String("markerColor"),
//   		MarkerShape: jsii.String("markerShape"),
//   		MarkerSize: jsii.String("markerSize"),
//   		MarkerVisibility: jsii.String("markerVisibility"),
//   	},
//   	MinValueMarker: &LineChartMarkerStyleSettingsProperty{
//   		MarkerColor: jsii.String("markerColor"),
//   		MarkerShape: jsii.String("markerShape"),
//   		MarkerSize: jsii.String("markerSize"),
//   		MarkerVisibility: jsii.String("markerVisibility"),
//   	},
//   	VisualType: jsii.String("visualType"),
//   	YAxisBehavior: jsii.String("yAxisBehavior"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sparklinesoptions.html
//
type CfnTemplate_SparklinesOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sparklinesoptions.html#cfn-quicksight-template-sparklinesoptions-fieldid
	//
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sparklinesoptions.html#cfn-quicksight-template-sparklinesoptions-xaxisfield
	//
	XAxisField interface{} `field:"required" json:"xAxisField" yaml:"xAxisField"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sparklinesoptions.html#cfn-quicksight-template-sparklinesoptions-allpointsmarker
	//
	AllPointsMarker interface{} `field:"optional" json:"allPointsMarker" yaml:"allPointsMarker"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sparklinesoptions.html#cfn-quicksight-template-sparklinesoptions-linecolor
	//
	LineColor *string `field:"optional" json:"lineColor" yaml:"lineColor"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sparklinesoptions.html#cfn-quicksight-template-sparklinesoptions-lineinterpolation
	//
	LineInterpolation *string `field:"optional" json:"lineInterpolation" yaml:"lineInterpolation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sparklinesoptions.html#cfn-quicksight-template-sparklinesoptions-maxvaluemarker
	//
	MaxValueMarker interface{} `field:"optional" json:"maxValueMarker" yaml:"maxValueMarker"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sparklinesoptions.html#cfn-quicksight-template-sparklinesoptions-minvaluemarker
	//
	MinValueMarker interface{} `field:"optional" json:"minValueMarker" yaml:"minValueMarker"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sparklinesoptions.html#cfn-quicksight-template-sparklinesoptions-visualtype
	//
	VisualType *string `field:"optional" json:"visualType" yaml:"visualType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sparklinesoptions.html#cfn-quicksight-template-sparklinesoptions-yaxisbehavior
	//
	YAxisBehavior *string `field:"optional" json:"yAxisBehavior" yaml:"yAxisBehavior"`
}

