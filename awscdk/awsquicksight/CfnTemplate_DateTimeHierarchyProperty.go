package awsquicksight


// The option that determines the hierarchy of any `DateTime` fields.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datetimehierarchy.html
//
type CfnTemplate_DateTimeHierarchyProperty struct {
	// The hierarchy ID of the `DateTime` hierarchy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datetimehierarchy.html#cfn-quicksight-template-datetimehierarchy-hierarchyid
	//
	HierarchyId *string `field:"required" json:"hierarchyId" yaml:"hierarchyId"`
	// The option that determines the drill down filters for the `DateTime` hierarchy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datetimehierarchy.html#cfn-quicksight-template-datetimehierarchy-drilldownfilters
	//
	DrillDownFilters interface{} `field:"optional" json:"drillDownFilters" yaml:"drillDownFilters"`
}

