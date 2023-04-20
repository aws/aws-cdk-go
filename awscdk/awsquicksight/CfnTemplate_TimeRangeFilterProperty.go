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
type CfnTemplate_TimeRangeFilterProperty struct {
	// `CfnTemplate.TimeRangeFilterProperty.Column`.
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// `CfnTemplate.TimeRangeFilterProperty.FilterId`.
	FilterId *string `field:"required" json:"filterId" yaml:"filterId"`
	// `CfnTemplate.TimeRangeFilterProperty.NullOption`.
	NullOption *string `field:"required" json:"nullOption" yaml:"nullOption"`
	// `CfnTemplate.TimeRangeFilterProperty.ExcludePeriodConfiguration`.
	ExcludePeriodConfiguration interface{} `field:"optional" json:"excludePeriodConfiguration" yaml:"excludePeriodConfiguration"`
	// `CfnTemplate.TimeRangeFilterProperty.IncludeMaximum`.
	IncludeMaximum interface{} `field:"optional" json:"includeMaximum" yaml:"includeMaximum"`
	// `CfnTemplate.TimeRangeFilterProperty.IncludeMinimum`.
	IncludeMinimum interface{} `field:"optional" json:"includeMinimum" yaml:"includeMinimum"`
	// `CfnTemplate.TimeRangeFilterProperty.RangeMaximumValue`.
	RangeMaximumValue interface{} `field:"optional" json:"rangeMaximumValue" yaml:"rangeMaximumValue"`
	// `CfnTemplate.TimeRangeFilterProperty.RangeMinimumValue`.
	RangeMinimumValue interface{} `field:"optional" json:"rangeMinimumValue" yaml:"rangeMinimumValue"`
	// `CfnTemplate.TimeRangeFilterProperty.TimeGranularity`.
	TimeGranularity *string `field:"optional" json:"timeGranularity" yaml:"timeGranularity"`
}

