package awsquicksight


// A grouping of individual filters. Filter groups are applied to the same group of visuals.
//
// For more information, see [Adding filter conditions (group filters) with AND and OR operators](https://docs.aws.amazon.com/quicksight/latest/user/add-a-compound-filter.html) in the *Amazon QuickSight User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var allSheets interface{}
//
//   filterGroupProperty := &FilterGroupProperty{
//   	CrossDataset: jsii.String("crossDataset"),
//   	FilterGroupId: jsii.String("filterGroupId"),
//   	Filters: []interface{}{
//   		&FilterProperty{
//   			CategoryFilter: &CategoryFilterProperty{
//   				Column: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				Configuration: &CategoryFilterConfigurationProperty{
//   					CustomFilterConfiguration: &CustomFilterConfigurationProperty{
//   						MatchOperator: jsii.String("matchOperator"),
//   						NullOption: jsii.String("nullOption"),
//
//   						// the properties below are optional
//   						CategoryValue: jsii.String("categoryValue"),
//   						ParameterName: jsii.String("parameterName"),
//   						SelectAllOptions: jsii.String("selectAllOptions"),
//   					},
//   					CustomFilterListConfiguration: &CustomFilterListConfigurationProperty{
//   						MatchOperator: jsii.String("matchOperator"),
//   						NullOption: jsii.String("nullOption"),
//
//   						// the properties below are optional
//   						CategoryValues: []*string{
//   							jsii.String("categoryValues"),
//   						},
//   						SelectAllOptions: jsii.String("selectAllOptions"),
//   					},
//   					FilterListConfiguration: &FilterListConfigurationProperty{
//   						MatchOperator: jsii.String("matchOperator"),
//
//   						// the properties below are optional
//   						CategoryValues: []*string{
//   							jsii.String("categoryValues"),
//   						},
//   						SelectAllOptions: jsii.String("selectAllOptions"),
//   					},
//   				},
//   				FilterId: jsii.String("filterId"),
//   			},
//   			NumericEqualityFilter: &NumericEqualityFilterProperty{
//   				Column: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				FilterId: jsii.String("filterId"),
//   				MatchOperator: jsii.String("matchOperator"),
//   				NullOption: jsii.String("nullOption"),
//
//   				// the properties below are optional
//   				AggregationFunction: &AggregationFunctionProperty{
//   					AttributeAggregationFunction: &AttributeAggregationFunctionProperty{
//   						SimpleAttributeAggregation: jsii.String("simpleAttributeAggregation"),
//   						ValueForMultipleValues: jsii.String("valueForMultipleValues"),
//   					},
//   					CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   					DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   					NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   						PercentileAggregation: &PercentileAggregationProperty{
//   							PercentileValue: jsii.Number(123),
//   						},
//   						SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   					},
//   				},
//   				ParameterName: jsii.String("parameterName"),
//   				SelectAllOptions: jsii.String("selectAllOptions"),
//   				Value: jsii.Number(123),
//   			},
//   			NumericRangeFilter: &NumericRangeFilterProperty{
//   				Column: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				FilterId: jsii.String("filterId"),
//   				NullOption: jsii.String("nullOption"),
//
//   				// the properties below are optional
//   				AggregationFunction: &AggregationFunctionProperty{
//   					AttributeAggregationFunction: &AttributeAggregationFunctionProperty{
//   						SimpleAttributeAggregation: jsii.String("simpleAttributeAggregation"),
//   						ValueForMultipleValues: jsii.String("valueForMultipleValues"),
//   					},
//   					CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   					DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   					NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   						PercentileAggregation: &PercentileAggregationProperty{
//   							PercentileValue: jsii.Number(123),
//   						},
//   						SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   					},
//   				},
//   				IncludeMaximum: jsii.Boolean(false),
//   				IncludeMinimum: jsii.Boolean(false),
//   				RangeMaximum: &NumericRangeFilterValueProperty{
//   					Parameter: jsii.String("parameter"),
//   					StaticValue: jsii.Number(123),
//   				},
//   				RangeMinimum: &NumericRangeFilterValueProperty{
//   					Parameter: jsii.String("parameter"),
//   					StaticValue: jsii.Number(123),
//   				},
//   				SelectAllOptions: jsii.String("selectAllOptions"),
//   			},
//   			RelativeDatesFilter: &RelativeDatesFilterProperty{
//   				AnchorDateConfiguration: &AnchorDateConfigurationProperty{
//   					AnchorOption: jsii.String("anchorOption"),
//   					ParameterName: jsii.String("parameterName"),
//   				},
//   				Column: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				FilterId: jsii.String("filterId"),
//   				NullOption: jsii.String("nullOption"),
//   				RelativeDateType: jsii.String("relativeDateType"),
//   				TimeGranularity: jsii.String("timeGranularity"),
//
//   				// the properties below are optional
//   				ExcludePeriodConfiguration: &ExcludePeriodConfigurationProperty{
//   					Amount: jsii.Number(123),
//   					Granularity: jsii.String("granularity"),
//
//   					// the properties below are optional
//   					Status: jsii.String("status"),
//   				},
//   				MinimumGranularity: jsii.String("minimumGranularity"),
//   				ParameterName: jsii.String("parameterName"),
//   				RelativeDateValue: jsii.Number(123),
//   			},
//   			TimeEqualityFilter: &TimeEqualityFilterProperty{
//   				Column: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				FilterId: jsii.String("filterId"),
//
//   				// the properties below are optional
//   				ParameterName: jsii.String("parameterName"),
//   				TimeGranularity: jsii.String("timeGranularity"),
//   				Value: jsii.String("value"),
//   			},
//   			TimeRangeFilter: &TimeRangeFilterProperty{
//   				Column: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				FilterId: jsii.String("filterId"),
//   				NullOption: jsii.String("nullOption"),
//
//   				// the properties below are optional
//   				ExcludePeriodConfiguration: &ExcludePeriodConfigurationProperty{
//   					Amount: jsii.Number(123),
//   					Granularity: jsii.String("granularity"),
//
//   					// the properties below are optional
//   					Status: jsii.String("status"),
//   				},
//   				IncludeMaximum: jsii.Boolean(false),
//   				IncludeMinimum: jsii.Boolean(false),
//   				RangeMaximumValue: &TimeRangeFilterValueProperty{
//   					Parameter: jsii.String("parameter"),
//   					RollingDate: &RollingDateConfigurationProperty{
//   						Expression: jsii.String("expression"),
//
//   						// the properties below are optional
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//   					StaticValue: jsii.String("staticValue"),
//   				},
//   				RangeMinimumValue: &TimeRangeFilterValueProperty{
//   					Parameter: jsii.String("parameter"),
//   					RollingDate: &RollingDateConfigurationProperty{
//   						Expression: jsii.String("expression"),
//
//   						// the properties below are optional
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//   					StaticValue: jsii.String("staticValue"),
//   				},
//   				TimeGranularity: jsii.String("timeGranularity"),
//   			},
//   			TopBottomFilter: &TopBottomFilterProperty{
//   				AggregationSortConfigurations: []interface{}{
//   					&AggregationSortConfigurationProperty{
//   						Column: &ColumnIdentifierProperty{
//   							ColumnName: jsii.String("columnName"),
//   							DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   						},
//   						SortDirection: jsii.String("sortDirection"),
//
//   						// the properties below are optional
//   						AggregationFunction: &AggregationFunctionProperty{
//   							AttributeAggregationFunction: &AttributeAggregationFunctionProperty{
//   								SimpleAttributeAggregation: jsii.String("simpleAttributeAggregation"),
//   								ValueForMultipleValues: jsii.String("valueForMultipleValues"),
//   							},
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
//   				},
//   				Column: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				FilterId: jsii.String("filterId"),
//
//   				// the properties below are optional
//   				Limit: jsii.Number(123),
//   				ParameterName: jsii.String("parameterName"),
//   				TimeGranularity: jsii.String("timeGranularity"),
//   			},
//   		},
//   	},
//   	ScopeConfiguration: &FilterScopeConfigurationProperty{
//   		AllSheets: allSheets,
//   		SelectedSheets: &SelectedSheetsFilterScopeConfigurationProperty{
//   			SheetVisualScopingConfigurations: []interface{}{
//   				&SheetVisualScopingConfigurationProperty{
//   					Scope: jsii.String("scope"),
//   					SheetId: jsii.String("sheetId"),
//
//   					// the properties below are optional
//   					VisualIds: []*string{
//   						jsii.String("visualIds"),
//   					},
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filtergroup.html
//
type CfnAnalysis_FilterGroupProperty struct {
	// The filter new feature which can apply filter group to all data sets. Choose one of the following options:.
	//
	// - `ALL_DATASETS`
	// - `SINGLE_DATASET`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filtergroup.html#cfn-quicksight-analysis-filtergroup-crossdataset
	//
	CrossDataset *string `field:"required" json:"crossDataset" yaml:"crossDataset"`
	// The value that uniquely identifies a `FilterGroup` within a dashboard, template, or analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filtergroup.html#cfn-quicksight-analysis-filtergroup-filtergroupid
	//
	FilterGroupId *string `field:"required" json:"filterGroupId" yaml:"filterGroupId"`
	// The list of filters that are present in a `FilterGroup` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filtergroup.html#cfn-quicksight-analysis-filtergroup-filters
	//
	Filters interface{} `field:"required" json:"filters" yaml:"filters"`
	// The configuration that specifies what scope to apply to a `FilterGroup` .
	//
	// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filtergroup.html#cfn-quicksight-analysis-filtergroup-scopeconfiguration
	//
	ScopeConfiguration interface{} `field:"required" json:"scopeConfiguration" yaml:"scopeConfiguration"`
	// The status of the `FilterGroup` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filtergroup.html#cfn-quicksight-analysis-filtergroup-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

