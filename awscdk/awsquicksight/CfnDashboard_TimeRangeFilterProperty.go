package awsquicksight


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
	// `CfnDashboard.TimeRangeFilterProperty.Column`.
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// `CfnDashboard.TimeRangeFilterProperty.FilterId`.
	FilterId *string `field:"required" json:"filterId" yaml:"filterId"`
	// `CfnDashboard.TimeRangeFilterProperty.NullOption`.
	NullOption *string `field:"required" json:"nullOption" yaml:"nullOption"`
	// `CfnDashboard.TimeRangeFilterProperty.ExcludePeriodConfiguration`.
	ExcludePeriodConfiguration interface{} `field:"optional" json:"excludePeriodConfiguration" yaml:"excludePeriodConfiguration"`
	// `CfnDashboard.TimeRangeFilterProperty.IncludeMaximum`.
	IncludeMaximum interface{} `field:"optional" json:"includeMaximum" yaml:"includeMaximum"`
	// `CfnDashboard.TimeRangeFilterProperty.IncludeMinimum`.
	IncludeMinimum interface{} `field:"optional" json:"includeMinimum" yaml:"includeMinimum"`
	// `CfnDashboard.TimeRangeFilterProperty.RangeMaximumValue`.
	RangeMaximumValue interface{} `field:"optional" json:"rangeMaximumValue" yaml:"rangeMaximumValue"`
	// `CfnDashboard.TimeRangeFilterProperty.RangeMinimumValue`.
	RangeMinimumValue interface{} `field:"optional" json:"rangeMinimumValue" yaml:"rangeMinimumValue"`
	// `CfnDashboard.TimeRangeFilterProperty.TimeGranularity`.
	TimeGranularity *string `field:"optional" json:"timeGranularity" yaml:"timeGranularity"`
}

