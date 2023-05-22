package awsquicksight


// The option that determines the hierarchy of the fields that are built within a visual's field wells.
//
// These fields can't be duplicated to other visuals.
//
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
	// The list of columns that define the explicit hierarchy.
	Columns interface{} `field:"required" json:"columns" yaml:"columns"`
	// The hierarchy ID of the explicit hierarchy.
	HierarchyId *string `field:"required" json:"hierarchyId" yaml:"hierarchyId"`
	// The option that determines the drill down filters for the explicit hierarchy.
	DrillDownFilters interface{} `field:"optional" json:"drillDownFilters" yaml:"drillDownFilters"`
}

