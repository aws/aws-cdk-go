package previewawsquicksightmixins


// The unique values computation configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   uniqueValuesComputationProperty := &UniqueValuesComputationProperty{
//   	Category: &DimensionFieldProperty{
//   		CategoricalDimensionField: &CategoricalDimensionFieldProperty{
//   			Column: &ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   			FieldId: jsii.String("fieldId"),
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
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   			DateGranularity: jsii.String("dateGranularity"),
//   			FieldId: jsii.String("fieldId"),
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
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   			FieldId: jsii.String("fieldId"),
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
//   	ComputationId: jsii.String("computationId"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-uniquevaluescomputation.html
//
type CfnDashboardPropsMixin_UniqueValuesComputationProperty struct {
	// The category field that is used in a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-uniquevaluescomputation.html#cfn-quicksight-dashboard-uniquevaluescomputation-category
	//
	Category interface{} `field:"optional" json:"category" yaml:"category"`
	// The ID for a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-uniquevaluescomputation.html#cfn-quicksight-dashboard-uniquevaluescomputation-computationid
	//
	ComputationId *string `field:"optional" json:"computationId" yaml:"computationId"`
	// The name of a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-uniquevaluescomputation.html#cfn-quicksight-dashboard-uniquevaluescomputation-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

