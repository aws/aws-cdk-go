package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dateTimeHierarchyProperty := &DateTimeHierarchyProperty{
//   	HierarchyId: jsii.String("hierarchyId"),
//
//   	// the properties below are optional
//   	DrillDownFilters: []interface{}{
//   		&DrillDownFilterProperty{
//   			CategoryFilter: &CategoryDrillDownFilterProperty{
//   				CategoryValues: []*string{
//   					jsii.String("categoryValues"),
//   				},
//   				Column: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   			},
//   			NumericEqualityFilter: &NumericEqualityDrillDownFilterProperty{
//   				Column: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				Value: jsii.Number(123),
//   			},
//   			TimeRangeFilter: &TimeRangeDrillDownFilterProperty{
//   				Column: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				RangeMaximum: jsii.String("rangeMaximum"),
//   				RangeMinimum: jsii.String("rangeMinimum"),
//   				TimeGranularity: jsii.String("timeGranularity"),
//   			},
//   		},
//   	},
//   }
//
type CfnAnalysis_DateTimeHierarchyProperty struct {
	// `CfnAnalysis.DateTimeHierarchyProperty.HierarchyId`.
	HierarchyId *string `field:"required" json:"hierarchyId" yaml:"hierarchyId"`
	// `CfnAnalysis.DateTimeHierarchyProperty.DrillDownFilters`.
	DrillDownFilters interface{} `field:"optional" json:"drillDownFilters" yaml:"drillDownFilters"`
}

