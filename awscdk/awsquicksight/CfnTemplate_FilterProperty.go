package awsquicksight


// With a `Filter` , you can remove portions of data from a particular visual or view.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterProperty := &FilterProperty{
//   	CategoryFilter: &CategoryFilterProperty{
//   		Column: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   		Configuration: &CategoryFilterConfigurationProperty{
//   			CustomFilterConfiguration: &CustomFilterConfigurationProperty{
//   				MatchOperator: jsii.String("matchOperator"),
//   				NullOption: jsii.String("nullOption"),
//
//   				// the properties below are optional
//   				CategoryValue: jsii.String("categoryValue"),
//   				ParameterName: jsii.String("parameterName"),
//   				SelectAllOptions: jsii.String("selectAllOptions"),
//   			},
//   			CustomFilterListConfiguration: &CustomFilterListConfigurationProperty{
//   				MatchOperator: jsii.String("matchOperator"),
//   				NullOption: jsii.String("nullOption"),
//
//   				// the properties below are optional
//   				CategoryValues: []*string{
//   					jsii.String("categoryValues"),
//   				},
//   				SelectAllOptions: jsii.String("selectAllOptions"),
//   			},
//   			FilterListConfiguration: &FilterListConfigurationProperty{
//   				MatchOperator: jsii.String("matchOperator"),
//
//   				// the properties below are optional
//   				CategoryValues: []*string{
//   					jsii.String("categoryValues"),
//   				},
//   				SelectAllOptions: jsii.String("selectAllOptions"),
//   			},
//   		},
//   		FilterId: jsii.String("filterId"),
//   	},
//   	NumericEqualityFilter: &NumericEqualityFilterProperty{
//   		Column: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   		FilterId: jsii.String("filterId"),
//   		MatchOperator: jsii.String("matchOperator"),
//   		NullOption: jsii.String("nullOption"),
//
//   		// the properties below are optional
//   		AggregationFunction: &AggregationFunctionProperty{
//   			CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   			DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   			NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   				PercentileAggregation: &PercentileAggregationProperty{
//   					PercentileValue: jsii.Number(123),
//   				},
//   				SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   			},
//   		},
//   		ParameterName: jsii.String("parameterName"),
//   		SelectAllOptions: jsii.String("selectAllOptions"),
//   		Value: jsii.Number(123),
//   	},
//   	NumericRangeFilter: &NumericRangeFilterProperty{
//   		Column: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   		FilterId: jsii.String("filterId"),
//   		NullOption: jsii.String("nullOption"),
//
//   		// the properties below are optional
//   		AggregationFunction: &AggregationFunctionProperty{
//   			CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   			DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   			NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   				PercentileAggregation: &PercentileAggregationProperty{
//   					PercentileValue: jsii.Number(123),
//   				},
//   				SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   			},
//   		},
//   		IncludeMaximum: jsii.Boolean(false),
//   		IncludeMinimum: jsii.Boolean(false),
//   		RangeMaximum: &NumericRangeFilterValueProperty{
//   			Parameter: jsii.String("parameter"),
//   			StaticValue: jsii.Number(123),
//   		},
//   		RangeMinimum: &NumericRangeFilterValueProperty{
//   			Parameter: jsii.String("parameter"),
//   			StaticValue: jsii.Number(123),
//   		},
//   		SelectAllOptions: jsii.String("selectAllOptions"),
//   	},
//   	RelativeDatesFilter: &RelativeDatesFilterProperty{
//   		AnchorDateConfiguration: &AnchorDateConfigurationProperty{
//   			AnchorOption: jsii.String("anchorOption"),
//   			ParameterName: jsii.String("parameterName"),
//   		},
//   		Column: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   		FilterId: jsii.String("filterId"),
//   		NullOption: jsii.String("nullOption"),
//   		RelativeDateType: jsii.String("relativeDateType"),
//   		TimeGranularity: jsii.String("timeGranularity"),
//
//   		// the properties below are optional
//   		ExcludePeriodConfiguration: &ExcludePeriodConfigurationProperty{
//   			Amount: jsii.Number(123),
//   			Granularity: jsii.String("granularity"),
//
//   			// the properties below are optional
//   			Status: jsii.String("status"),
//   		},
//   		MinimumGranularity: jsii.String("minimumGranularity"),
//   		ParameterName: jsii.String("parameterName"),
//   		RelativeDateValue: jsii.Number(123),
//   	},
//   	TimeEqualityFilter: &TimeEqualityFilterProperty{
//   		Column: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   		FilterId: jsii.String("filterId"),
//
//   		// the properties below are optional
//   		ParameterName: jsii.String("parameterName"),
//   		TimeGranularity: jsii.String("timeGranularity"),
//   		Value: jsii.String("value"),
//   	},
//   	TimeRangeFilter: &TimeRangeFilterProperty{
//   		Column: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   		FilterId: jsii.String("filterId"),
//   		NullOption: jsii.String("nullOption"),
//
//   		// the properties below are optional
//   		ExcludePeriodConfiguration: &ExcludePeriodConfigurationProperty{
//   			Amount: jsii.Number(123),
//   			Granularity: jsii.String("granularity"),
//
//   			// the properties below are optional
//   			Status: jsii.String("status"),
//   		},
//   		IncludeMaximum: jsii.Boolean(false),
//   		IncludeMinimum: jsii.Boolean(false),
//   		RangeMaximumValue: &TimeRangeFilterValueProperty{
//   			Parameter: jsii.String("parameter"),
//   			RollingDate: &RollingDateConfigurationProperty{
//   				Expression: jsii.String("expression"),
//
//   				// the properties below are optional
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   			StaticValue: jsii.String("staticValue"),
//   		},
//   		RangeMinimumValue: &TimeRangeFilterValueProperty{
//   			Parameter: jsii.String("parameter"),
//   			RollingDate: &RollingDateConfigurationProperty{
//   				Expression: jsii.String("expression"),
//
//   				// the properties below are optional
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   			StaticValue: jsii.String("staticValue"),
//   		},
//   		TimeGranularity: jsii.String("timeGranularity"),
//   	},
//   	TopBottomFilter: &TopBottomFilterProperty{
//   		AggregationSortConfigurations: []interface{}{
//   			&AggregationSortConfigurationProperty{
//   				Column: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				SortDirection: jsii.String("sortDirection"),
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
//   			},
//   		},
//   		Column: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   		FilterId: jsii.String("filterId"),
//
//   		// the properties below are optional
//   		Limit: jsii.Number(123),
//   		ParameterName: jsii.String("parameterName"),
//   		TimeGranularity: jsii.String("timeGranularity"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filter.html
//
type CfnTemplate_FilterProperty struct {
	// A `CategoryFilter` filters text values.
	//
	// For more information, see [Adding text filters](https://docs.aws.amazon.com/quicksight/latest/user/add-a-text-filter-data-prep.html) in the *Amazon QuickSight User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filter.html#cfn-quicksight-template-filter-categoryfilter
	//
	CategoryFilter interface{} `field:"optional" json:"categoryFilter" yaml:"categoryFilter"`
	// A `NumericEqualityFilter` filters numeric values that equal or do not equal a given numeric value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filter.html#cfn-quicksight-template-filter-numericequalityfilter
	//
	NumericEqualityFilter interface{} `field:"optional" json:"numericEqualityFilter" yaml:"numericEqualityFilter"`
	// A `NumericRangeFilter` filters numeric values that are either inside or outside a given numeric range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filter.html#cfn-quicksight-template-filter-numericrangefilter
	//
	NumericRangeFilter interface{} `field:"optional" json:"numericRangeFilter" yaml:"numericRangeFilter"`
	// A `RelativeDatesFilter` filters date values that are relative to a given date.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filter.html#cfn-quicksight-template-filter-relativedatesfilter
	//
	RelativeDatesFilter interface{} `field:"optional" json:"relativeDatesFilter" yaml:"relativeDatesFilter"`
	// A `TimeEqualityFilter` filters date-time values that equal or do not equal a given date/time value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filter.html#cfn-quicksight-template-filter-timeequalityfilter
	//
	TimeEqualityFilter interface{} `field:"optional" json:"timeEqualityFilter" yaml:"timeEqualityFilter"`
	// A `TimeRangeFilter` filters date-time values that are either inside or outside a given date/time range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filter.html#cfn-quicksight-template-filter-timerangefilter
	//
	TimeRangeFilter interface{} `field:"optional" json:"timeRangeFilter" yaml:"timeRangeFilter"`
	// A `TopBottomFilter` filters data to the top or bottom values for a given column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filter.html#cfn-quicksight-template-filter-topbottomfilter
	//
	TopBottomFilter interface{} `field:"optional" json:"topBottomFilter" yaml:"topBottomFilter"`
}

