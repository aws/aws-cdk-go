package awsquicksight


// A `TimeRangeFilter` filters values that are between two specified values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timeRangeFilterProperty := &TimeRangeFilterProperty{
//   	Column: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	FilterId: jsii.String("filterId"),
//   	NullOption: jsii.String("nullOption"),
//
//   	// the properties below are optional
//   	ExcludePeriodConfiguration: &ExcludePeriodConfigurationProperty{
//   		Amount: jsii.Number(123),
//   		Granularity: jsii.String("granularity"),
//
//   		// the properties below are optional
//   		Status: jsii.String("status"),
//   	},
//   	IncludeMaximum: jsii.Boolean(false),
//   	IncludeMinimum: jsii.Boolean(false),
//   	RangeMaximumValue: &TimeRangeFilterValueProperty{
//   		Parameter: jsii.String("parameter"),
//   		RollingDate: &RollingDateConfigurationProperty{
//   			Expression: jsii.String("expression"),
//
//   			// the properties below are optional
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   		StaticValue: jsii.String("staticValue"),
//   	},
//   	RangeMinimumValue: &TimeRangeFilterValueProperty{
//   		Parameter: jsii.String("parameter"),
//   		RollingDate: &RollingDateConfigurationProperty{
//   			Expression: jsii.String("expression"),
//
//   			// the properties below are optional
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   		StaticValue: jsii.String("staticValue"),
//   	},
//   	TimeGranularity: jsii.String("timeGranularity"),
//   }
//
type CfnDashboard_TimeRangeFilterProperty struct {
	// The column that the filter is applied to.
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// An identifier that uniquely identifies a filter within a dashboard, analysis, or template.
	FilterId *string `field:"required" json:"filterId" yaml:"filterId"`
	// This option determines how null values should be treated when filtering data.
	//
	// - `ALL_VALUES` : Include null values in filtered results.
	// - `NULLS_ONLY` : Only include null values in filtered results.
	// - `NON_NULLS_ONLY` : Exclude null values from filtered results.
	NullOption *string `field:"required" json:"nullOption" yaml:"nullOption"`
	// The exclude period of the time range filter.
	ExcludePeriodConfiguration interface{} `field:"optional" json:"excludePeriodConfiguration" yaml:"excludePeriodConfiguration"`
	// Determines whether the maximum value in the filter value range should be included in the filtered results.
	IncludeMaximum interface{} `field:"optional" json:"includeMaximum" yaml:"includeMaximum"`
	// Determines whether the minimum value in the filter value range should be included in the filtered results.
	IncludeMinimum interface{} `field:"optional" json:"includeMinimum" yaml:"includeMinimum"`
	// The maximum value for the filter value range.
	RangeMaximumValue interface{} `field:"optional" json:"rangeMaximumValue" yaml:"rangeMaximumValue"`
	// The minimum value for the filter value range.
	RangeMinimumValue interface{} `field:"optional" json:"rangeMinimumValue" yaml:"rangeMinimumValue"`
	// The level of time precision that is used to aggregate `DateTime` values.
	TimeGranularity *string `field:"optional" json:"timeGranularity" yaml:"timeGranularity"`
}

