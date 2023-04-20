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
type CfnAnalysis_TimeRangeFilterProperty struct {
	// `CfnAnalysis.TimeRangeFilterProperty.Column`.
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// `CfnAnalysis.TimeRangeFilterProperty.FilterId`.
	FilterId *string `field:"required" json:"filterId" yaml:"filterId"`
	// `CfnAnalysis.TimeRangeFilterProperty.NullOption`.
	NullOption *string `field:"required" json:"nullOption" yaml:"nullOption"`
	// `CfnAnalysis.TimeRangeFilterProperty.ExcludePeriodConfiguration`.
	ExcludePeriodConfiguration interface{} `field:"optional" json:"excludePeriodConfiguration" yaml:"excludePeriodConfiguration"`
	// `CfnAnalysis.TimeRangeFilterProperty.IncludeMaximum`.
	IncludeMaximum interface{} `field:"optional" json:"includeMaximum" yaml:"includeMaximum"`
	// `CfnAnalysis.TimeRangeFilterProperty.IncludeMinimum`.
	IncludeMinimum interface{} `field:"optional" json:"includeMinimum" yaml:"includeMinimum"`
	// `CfnAnalysis.TimeRangeFilterProperty.RangeMaximumValue`.
	RangeMaximumValue interface{} `field:"optional" json:"rangeMaximumValue" yaml:"rangeMaximumValue"`
	// `CfnAnalysis.TimeRangeFilterProperty.RangeMinimumValue`.
	RangeMinimumValue interface{} `field:"optional" json:"rangeMinimumValue" yaml:"rangeMinimumValue"`
	// `CfnAnalysis.TimeRangeFilterProperty.TimeGranularity`.
	TimeGranularity *string `field:"optional" json:"timeGranularity" yaml:"timeGranularity"`
}

