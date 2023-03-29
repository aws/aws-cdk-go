package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geospatialMapVisualProperty := &GeospatialMapVisualProperty{
//   	VisualId: jsii.String("visualId"),
//
//   	// the properties below are optional
//   	Actions: []interface{}{
//   		&VisualCustomActionProperty{
//   			ActionOperations: []interface{}{
//   				&VisualCustomActionOperationProperty{
//   					FilterOperation: &CustomActionFilterOperationProperty{
//   						SelectedFieldsConfiguration: &FilterOperationSelectedFieldsConfigurationProperty{
//   							SelectedFieldOptions: jsii.String("selectedFieldOptions"),
//   							SelectedFields: []*string{
//   								jsii.String("selectedFields"),
//   							},
//   						},
//   						TargetVisualsConfiguration: &FilterOperationTargetVisualsConfigurationProperty{
//   							SameSheetTargetVisualConfiguration: &SameSheetTargetVisualConfigurationProperty{
//   								TargetVisualOptions: jsii.String("targetVisualOptions"),
//   								TargetVisuals: []*string{
//   									jsii.String("targetVisuals"),
//   								},
//   							},
//   						},
//   					},
//   					NavigationOperation: &CustomActionNavigationOperationProperty{
//   						LocalNavigationConfiguration: &LocalNavigationConfigurationProperty{
//   							TargetSheetId: jsii.String("targetSheetId"),
//   						},
//   					},
//   					SetParametersOperation: &CustomActionSetParametersOperationProperty{
//   						ParameterValueConfigurations: []interface{}{
//   							&SetParameterValueConfigurationProperty{
//   								DestinationParameterName: jsii.String("destinationParameterName"),
//   								Value: &DestinationParameterValueConfigurationProperty{
//   									CustomValuesConfiguration: &CustomValuesConfigurationProperty{
//   										CustomValues: &CustomParameterValuesProperty{
//   											DateTimeValues: []*string{
//   												jsii.String("dateTimeValues"),
//   											},
//   											DecimalValues: []interface{}{
//   												jsii.Number(123),
//   											},
//   											IntegerValues: []interface{}{
//   												jsii.Number(123),
//   											},
//   											StringValues: []*string{
//   												jsii.String("stringValues"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										IncludeNullValue: jsii.Boolean(false),
//   									},
//   									SelectAllValueOptions: jsii.String("selectAllValueOptions"),
//   									SourceField: jsii.String("sourceField"),
//   									SourceParameterName: jsii.String("sourceParameterName"),
//   								},
//   							},
//   						},
//   					},
//   					UrlOperation: &CustomActionURLOperationProperty{
//   						UrlTarget: jsii.String("urlTarget"),
//   						UrlTemplate: jsii.String("urlTemplate"),
//   					},
//   				},
//   			},
//   			CustomActionId: jsii.String("customActionId"),
//   			Name: jsii.String("name"),
//   			Trigger: jsii.String("trigger"),
//
//   			// the properties below are optional
//   			Status: jsii.String("status"),
//   		},
//   	},
//   	ChartConfiguration: &GeospatialMapConfigurationProperty{
//   		FieldWells: &GeospatialMapFieldWellsProperty{
//   			GeospatialMapAggregatedFieldWells: &GeospatialMapAggregatedFieldWellsProperty{
//   				Colors: []interface{}{
//   					&DimensionFieldProperty{
//   						CategoricalDimensionField: &CategoricalDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							FormatConfiguration: &StringFormatConfigurationProperty{
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   						DateDimensionField: &DateDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							DateGranularity: jsii.String("dateGranularity"),
//   							FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   								DateTimeFormat: jsii.String("dateTimeFormat"),
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   						NumericalDimensionField: &NumericalDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							FormatConfiguration: &NumberFormatConfigurationProperty{
//   								FormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   					},
//   				},
//   				Geospatial: []interface{}{
//   					&DimensionFieldProperty{
//   						CategoricalDimensionField: &CategoricalDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							FormatConfiguration: &StringFormatConfigurationProperty{
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   						DateDimensionField: &DateDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							DateGranularity: jsii.String("dateGranularity"),
//   							FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   								DateTimeFormat: jsii.String("dateTimeFormat"),
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   						NumericalDimensionField: &NumericalDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							FormatConfiguration: &NumberFormatConfigurationProperty{
//   								FormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   					},
//   				},
//   				Values: []interface{}{
//   					&MeasureFieldProperty{
//   						CalculatedMeasureField: &CalculatedMeasureFieldProperty{
//   							Expression: jsii.String("expression"),
//   							FieldId: jsii.String("fieldId"),
//   						},
//   						CategoricalMeasureField: &CategoricalMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: jsii.String("aggregationFunction"),
//   							FormatConfiguration: &StringFormatConfigurationProperty{
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   						DateMeasureField: &DateMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: jsii.String("aggregationFunction"),
//   							FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   								DateTimeFormat: jsii.String("dateTimeFormat"),
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   						NumericalMeasureField: &NumericalMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: &NumericalAggregationFunctionProperty{
//   								PercentileAggregation: &PercentileAggregationProperty{
//   									PercentileValue: jsii.Number(123),
//   								},
//   								SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   							},
//   							FormatConfiguration: &NumberFormatConfigurationProperty{
//   								FormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		Legend: &LegendOptionsProperty{
//   			Height: jsii.String("height"),
//   			Position: jsii.String("position"),
//   			Title: &LabelOptionsProperty{
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
//   			Visibility: jsii.String("visibility"),
//   			Width: jsii.String("width"),
//   		},
//   		MapStyleOptions: &GeospatialMapStyleOptionsProperty{
//   			BaseMapStyle: jsii.String("baseMapStyle"),
//   		},
//   		PointStyleOptions: &GeospatialPointStyleOptionsProperty{
//   			ClusterMarkerConfiguration: &ClusterMarkerConfigurationProperty{
//   				ClusterMarker: &ClusterMarkerProperty{
//   					SimpleClusterMarker: &SimpleClusterMarkerProperty{
//   						Color: jsii.String("color"),
//   					},
//   				},
//   			},
//   			SelectedPointStyle: jsii.String("selectedPointStyle"),
//   		},
//   		Tooltip: &TooltipOptionsProperty{
//   			FieldBasedTooltip: &FieldBasedTooltipProperty{
//   				AggregationVisibility: jsii.String("aggregationVisibility"),
//   				TooltipFields: []interface{}{
//   					&TooltipItemProperty{
//   						ColumnTooltipItem: &ColumnTooltipItemProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//
//   							// the properties below are optional
//   							Aggregation: &AggregationFunctionProperty{
//   								CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   								DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   								NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   									PercentileAggregation: &PercentileAggregationProperty{
//   										PercentileValue: jsii.Number(123),
//   									},
//   									SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   								},
//   							},
//   							Label: jsii.String("label"),
//   							Visibility: jsii.String("visibility"),
//   						},
//   						FieldTooltipItem: &FieldTooltipItemProperty{
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							Label: jsii.String("label"),
//   							Visibility: jsii.String("visibility"),
//   						},
//   					},
//   				},
//   				TooltipTitleType: jsii.String("tooltipTitleType"),
//   			},
//   			SelectedTooltipType: jsii.String("selectedTooltipType"),
//   			TooltipVisibility: jsii.String("tooltipVisibility"),
//   		},
//   		VisualPalette: &VisualPaletteProperty{
//   			ChartColor: jsii.String("chartColor"),
//   			ColorMap: []interface{}{
//   				&DataPathColorProperty{
//   					Color: jsii.String("color"),
//   					Element: &DataPathValueProperty{
//   						FieldId: jsii.String("fieldId"),
//   						FieldValue: jsii.String("fieldValue"),
//   					},
//
//   					// the properties below are optional
//   					TimeGranularity: jsii.String("timeGranularity"),
//   				},
//   			},
//   		},
//   		WindowOptions: &GeospatialWindowOptionsProperty{
//   			Bounds: &GeospatialCoordinateBoundsProperty{
//   				East: jsii.Number(123),
//   				North: jsii.Number(123),
//   				South: jsii.Number(123),
//   				West: jsii.Number(123),
//   			},
//   			MapZoomMode: jsii.String("mapZoomMode"),
//   		},
//   	},
//   	ColumnHierarchies: []interface{}{
//   		&ColumnHierarchyProperty{
//   			DateTimeHierarchy: &DateTimeHierarchyProperty{
//   				HierarchyId: jsii.String("hierarchyId"),
//
//   				// the properties below are optional
//   				DrillDownFilters: []interface{}{
//   					&DrillDownFilterProperty{
//   						CategoryFilter: &CategoryDrillDownFilterProperty{
//   							CategoryValues: []*string{
//   								jsii.String("categoryValues"),
//   							},
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   						},
//   						NumericEqualityFilter: &NumericEqualityDrillDownFilterProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							Value: jsii.Number(123),
//   						},
//   						TimeRangeFilter: &TimeRangeDrillDownFilterProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							RangeMaximum: jsii.String("rangeMaximum"),
//   							RangeMinimum: jsii.String("rangeMinimum"),
//   							TimeGranularity: jsii.String("timeGranularity"),
//   						},
//   					},
//   				},
//   			},
//   			ExplicitHierarchy: &ExplicitHierarchyProperty{
//   				Columns: []interface{}{
//   					&ColumnIdentifierProperty{
//   						ColumnName: jsii.String("columnName"),
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//   				},
//   				HierarchyId: jsii.String("hierarchyId"),
//
//   				// the properties below are optional
//   				DrillDownFilters: []interface{}{
//   					&DrillDownFilterProperty{
//   						CategoryFilter: &CategoryDrillDownFilterProperty{
//   							CategoryValues: []*string{
//   								jsii.String("categoryValues"),
//   							},
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   						},
//   						NumericEqualityFilter: &NumericEqualityDrillDownFilterProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							Value: jsii.Number(123),
//   						},
//   						TimeRangeFilter: &TimeRangeDrillDownFilterProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							RangeMaximum: jsii.String("rangeMaximum"),
//   							RangeMinimum: jsii.String("rangeMinimum"),
//   							TimeGranularity: jsii.String("timeGranularity"),
//   						},
//   					},
//   				},
//   			},
//   			PredefinedHierarchy: &PredefinedHierarchyProperty{
//   				Columns: []interface{}{
//   					&ColumnIdentifierProperty{
//   						ColumnName: jsii.String("columnName"),
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//   				},
//   				HierarchyId: jsii.String("hierarchyId"),
//
//   				// the properties below are optional
//   				DrillDownFilters: []interface{}{
//   					&DrillDownFilterProperty{
//   						CategoryFilter: &CategoryDrillDownFilterProperty{
//   							CategoryValues: []*string{
//   								jsii.String("categoryValues"),
//   							},
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   						},
//   						NumericEqualityFilter: &NumericEqualityDrillDownFilterProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							Value: jsii.Number(123),
//   						},
//   						TimeRangeFilter: &TimeRangeDrillDownFilterProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							RangeMaximum: jsii.String("rangeMaximum"),
//   							RangeMinimum: jsii.String("rangeMinimum"),
//   							TimeGranularity: jsii.String("timeGranularity"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Subtitle: &VisualSubtitleLabelOptionsProperty{
//   		FormatText: &LongFormatTextProperty{
//   			PlainText: jsii.String("plainText"),
//   			RichText: jsii.String("richText"),
//   		},
//   		Visibility: jsii.String("visibility"),
//   	},
//   	Title: &VisualTitleLabelOptionsProperty{
//   		FormatText: &ShortFormatTextProperty{
//   			PlainText: jsii.String("plainText"),
//   			RichText: jsii.String("richText"),
//   		},
//   		Visibility: jsii.String("visibility"),
//   	},
//   }
//
type CfnTemplate_GeospatialMapVisualProperty struct {
	// `CfnTemplate.GeospatialMapVisualProperty.VisualId`.
	VisualId *string `field:"required" json:"visualId" yaml:"visualId"`
	// `CfnTemplate.GeospatialMapVisualProperty.Actions`.
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// `CfnTemplate.GeospatialMapVisualProperty.ChartConfiguration`.
	ChartConfiguration interface{} `field:"optional" json:"chartConfiguration" yaml:"chartConfiguration"`
	// `CfnTemplate.GeospatialMapVisualProperty.ColumnHierarchies`.
	ColumnHierarchies interface{} `field:"optional" json:"columnHierarchies" yaml:"columnHierarchies"`
	// `CfnTemplate.GeospatialMapVisualProperty.Subtitle`.
	Subtitle interface{} `field:"optional" json:"subtitle" yaml:"subtitle"`
	// `CfnTemplate.GeospatialMapVisualProperty.Title`.
	Title interface{} `field:"optional" json:"title" yaml:"title"`
}

