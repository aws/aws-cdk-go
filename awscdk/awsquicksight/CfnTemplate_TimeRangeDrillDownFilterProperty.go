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
type CfnTemplate_TimeRangeDrillDownFilterProperty struct {
	// `CfnTemplate.TimeRangeDrillDownFilterProperty.Column`.
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// `CfnTemplate.TimeRangeDrillDownFilterProperty.RangeMaximum`.
	RangeMaximum *string `field:"required" json:"rangeMaximum" yaml:"rangeMaximum"`
	// `CfnTemplate.TimeRangeDrillDownFilterProperty.RangeMinimum`.
	RangeMinimum *string `field:"required" json:"rangeMinimum" yaml:"rangeMinimum"`
	// `CfnTemplate.TimeRangeDrillDownFilterProperty.TimeGranularity`.
	TimeGranularity *string `field:"required" json:"timeGranularity" yaml:"timeGranularity"`
}

