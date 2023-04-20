package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   explicitHierarchyProperty := &ExplicitHierarchyProperty{
//   	Columns: []interface{}{
//   		&ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   	},
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
type CfnDashboard_ExplicitHierarchyProperty struct {
	// `CfnDashboard.ExplicitHierarchyProperty.Columns`.
	Columns interface{} `field:"required" json:"columns" yaml:"columns"`
	// `CfnDashboard.ExplicitHierarchyProperty.HierarchyId`.
	HierarchyId *string `field:"required" json:"hierarchyId" yaml:"hierarchyId"`
	// `CfnDashboard.ExplicitHierarchyProperty.DrillDownFilters`.
	DrillDownFilters interface{} `field:"optional" json:"drillDownFilters" yaml:"drillDownFilters"`
}

