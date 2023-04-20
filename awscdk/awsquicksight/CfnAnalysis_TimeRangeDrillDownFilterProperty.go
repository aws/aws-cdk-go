package awsquicksight


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
	// `CfnAnalysis.TimeRangeDrillDownFilterProperty.Column`.
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// `CfnAnalysis.TimeRangeDrillDownFilterProperty.RangeMaximum`.
	RangeMaximum *string `field:"required" json:"rangeMaximum" yaml:"rangeMaximum"`
	// `CfnAnalysis.TimeRangeDrillDownFilterProperty.RangeMinimum`.
	RangeMinimum *string `field:"required" json:"rangeMinimum" yaml:"rangeMinimum"`
	// `CfnAnalysis.TimeRangeDrillDownFilterProperty.TimeGranularity`.
	TimeGranularity *string `field:"required" json:"timeGranularity" yaml:"timeGranularity"`
}

