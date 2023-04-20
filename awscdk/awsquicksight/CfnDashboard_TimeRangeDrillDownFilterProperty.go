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
type CfnDashboard_TimeRangeDrillDownFilterProperty struct {
	// `CfnDashboard.TimeRangeDrillDownFilterProperty.Column`.
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// `CfnDashboard.TimeRangeDrillDownFilterProperty.RangeMaximum`.
	RangeMaximum *string `field:"required" json:"rangeMaximum" yaml:"rangeMaximum"`
	// `CfnDashboard.TimeRangeDrillDownFilterProperty.RangeMinimum`.
	RangeMinimum *string `field:"required" json:"rangeMinimum" yaml:"rangeMinimum"`
	// `CfnDashboard.TimeRangeDrillDownFilterProperty.TimeGranularity`.
	TimeGranularity *string `field:"required" json:"timeGranularity" yaml:"timeGranularity"`
}

