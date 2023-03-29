package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   drillDownFilterProperty := &DrillDownFilterProperty{
//   	CategoryFilter: &CategoryDrillDownFilterProperty{
//   		CategoryValues: []*string{
//   			jsii.String("categoryValues"),
//   		},
//   		Column: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   	},
//   	NumericEqualityFilter: &NumericEqualityDrillDownFilterProperty{
//   		Column: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   		Value: jsii.Number(123),
//   	},
//   	TimeRangeFilter: &TimeRangeDrillDownFilterProperty{
//   		Column: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   		RangeMaximum: jsii.String("rangeMaximum"),
//   		RangeMinimum: jsii.String("rangeMinimum"),
//   		TimeGranularity: jsii.String("timeGranularity"),
//   	},
//   }
//
type CfnAnalysis_DrillDownFilterProperty struct {
	// `CfnAnalysis.DrillDownFilterProperty.CategoryFilter`.
	CategoryFilter interface{} `field:"optional" json:"categoryFilter" yaml:"categoryFilter"`
	// `CfnAnalysis.DrillDownFilterProperty.NumericEqualityFilter`.
	NumericEqualityFilter interface{} `field:"optional" json:"numericEqualityFilter" yaml:"numericEqualityFilter"`
	// `CfnAnalysis.DrillDownFilterProperty.TimeRangeFilter`.
	TimeRangeFilter interface{} `field:"optional" json:"timeRangeFilter" yaml:"timeRangeFilter"`
}

