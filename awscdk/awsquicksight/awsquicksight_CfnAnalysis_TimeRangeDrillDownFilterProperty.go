package awsquicksight


// The time range drill down filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timeRangeDrillDownFilterProperty := &TimeRangeDrillDownFilterProperty{
//   	Column: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	RangeMaximum: jsii.String("rangeMaximum"),
//   	RangeMinimum: jsii.String("rangeMinimum"),
//   	TimeGranularity: jsii.String("timeGranularity"),
//   }
//
type CfnAnalysis_TimeRangeDrillDownFilterProperty struct {
	// The column that the filter is applied to.
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// The maximum value for the filter value range.
	RangeMaximum *string `field:"required" json:"rangeMaximum" yaml:"rangeMaximum"`
	// The minimum value for the filter value range.
	RangeMinimum *string `field:"required" json:"rangeMinimum" yaml:"rangeMinimum"`
	// The level of time precision that is used to aggregate `DateTime` values.
	TimeGranularity *string `field:"required" json:"timeGranularity" yaml:"timeGranularity"`
}

