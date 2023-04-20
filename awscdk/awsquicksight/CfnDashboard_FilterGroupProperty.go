package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//   						Column: &ColumnIdentifierProperty{
//   							ColumnName: jsii.String("columnName"),
//   							DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   						},
//   						SortDirection: jsii.String("sortDirection"),
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
type CfnDashboard_FilterGroupProperty struct {
	// `CfnDashboard.FilterGroupProperty.CrossDataset`.
	CrossDataset *string `field:"required" json:"crossDataset" yaml:"crossDataset"`
	// `CfnDashboard.FilterGroupProperty.FilterGroupId`.
	FilterGroupId *string `field:"required" json:"filterGroupId" yaml:"filterGroupId"`
	// `CfnDashboard.FilterGroupProperty.Filters`.
	Filters interface{} `field:"required" json:"filters" yaml:"filters"`
	// `CfnDashboard.FilterGroupProperty.ScopeConfiguration`.
	ScopeConfiguration interface{} `field:"required" json:"scopeConfiguration" yaml:"scopeConfiguration"`
	// `CfnDashboard.FilterGroupProperty.Status`.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

