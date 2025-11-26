package previewawsquicksightmixins


// The option that determines the hierarchy of the fields that are defined during data preparation.
//
// These fields are available to use in any analysis that uses the data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   predefinedHierarchyProperty := &PredefinedHierarchyProperty{
//   	Columns: []interface{}{
//   		&ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   	},
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
//   	HierarchyId: jsii.String("hierarchyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-predefinedhierarchy.html
//
type CfnDashboardPropsMixin_PredefinedHierarchyProperty struct {
	// The list of columns that define the predefined hierarchy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-predefinedhierarchy.html#cfn-quicksight-dashboard-predefinedhierarchy-columns
	//
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// The option that determines the drill down filters for the predefined hierarchy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-predefinedhierarchy.html#cfn-quicksight-dashboard-predefinedhierarchy-drilldownfilters
	//
	DrillDownFilters interface{} `field:"optional" json:"drillDownFilters" yaml:"drillDownFilters"`
	// The hierarchy ID of the predefined hierarchy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-predefinedhierarchy.html#cfn-quicksight-dashboard-predefinedhierarchy-hierarchyid
	//
	HierarchyId *string `field:"optional" json:"hierarchyId" yaml:"hierarchyId"`
}

