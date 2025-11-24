package mixinsawsquicksight


// The option that determines the hierarchy of the fields for a visual element.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   columnHierarchyProperty := &ColumnHierarchyProperty{
//   	DateTimeHierarchy: &DateTimeHierarchyProperty{
//   		DrillDownFilters: []interface{}{
//   			&DrillDownFilterProperty{
//   				CategoryFilter: &CategoryDrillDownFilterProperty{
//   					CategoryValues: []*string{
//   						jsii.String("categoryValues"),
//   					},
//   					Column: &ColumnIdentifierProperty{
//   						ColumnName: jsii.String("columnName"),
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//   				},
//   				NumericEqualityFilter: &NumericEqualityDrillDownFilterProperty{
//   					Column: &ColumnIdentifierProperty{
//   						ColumnName: jsii.String("columnName"),
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//   					Value: jsii.Number(123),
//   				},
//   				TimeRangeFilter: &TimeRangeDrillDownFilterProperty{
//   					Column: &ColumnIdentifierProperty{
//   						ColumnName: jsii.String("columnName"),
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//   					RangeMaximum: jsii.String("rangeMaximum"),
//   					RangeMinimum: jsii.String("rangeMinimum"),
//   					TimeGranularity: jsii.String("timeGranularity"),
//   				},
//   			},
//   		},
//   		HierarchyId: jsii.String("hierarchyId"),
//   	},
//   	ExplicitHierarchy: &ExplicitHierarchyProperty{
//   		Columns: []interface{}{
//   			&ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   		},
//   		DrillDownFilters: []interface{}{
//   			&DrillDownFilterProperty{
//   				CategoryFilter: &CategoryDrillDownFilterProperty{
//   					CategoryValues: []*string{
//   						jsii.String("categoryValues"),
//   					},
//   					Column: &ColumnIdentifierProperty{
//   						ColumnName: jsii.String("columnName"),
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//   				},
//   				NumericEqualityFilter: &NumericEqualityDrillDownFilterProperty{
//   					Column: &ColumnIdentifierProperty{
//   						ColumnName: jsii.String("columnName"),
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//   					Value: jsii.Number(123),
//   				},
//   				TimeRangeFilter: &TimeRangeDrillDownFilterProperty{
//   					Column: &ColumnIdentifierProperty{
//   						ColumnName: jsii.String("columnName"),
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//   					RangeMaximum: jsii.String("rangeMaximum"),
//   					RangeMinimum: jsii.String("rangeMinimum"),
//   					TimeGranularity: jsii.String("timeGranularity"),
//   				},
//   			},
//   		},
//   		HierarchyId: jsii.String("hierarchyId"),
//   	},
//   	PredefinedHierarchy: &PredefinedHierarchyProperty{
//   		Columns: []interface{}{
//   			&ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   		},
//   		DrillDownFilters: []interface{}{
//   			&DrillDownFilterProperty{
//   				CategoryFilter: &CategoryDrillDownFilterProperty{
//   					CategoryValues: []*string{
//   						jsii.String("categoryValues"),
//   					},
//   					Column: &ColumnIdentifierProperty{
//   						ColumnName: jsii.String("columnName"),
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//   				},
//   				NumericEqualityFilter: &NumericEqualityDrillDownFilterProperty{
//   					Column: &ColumnIdentifierProperty{
//   						ColumnName: jsii.String("columnName"),
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//   					Value: jsii.Number(123),
//   				},
//   				TimeRangeFilter: &TimeRangeDrillDownFilterProperty{
//   					Column: &ColumnIdentifierProperty{
//   						ColumnName: jsii.String("columnName"),
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//   					RangeMaximum: jsii.String("rangeMaximum"),
//   					RangeMinimum: jsii.String("rangeMinimum"),
//   					TimeGranularity: jsii.String("timeGranularity"),
//   				},
//   			},
//   		},
//   		HierarchyId: jsii.String("hierarchyId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-columnhierarchy.html
//
type CfnAnalysisPropsMixin_ColumnHierarchyProperty struct {
	// The option that determines the hierarchy of any `DateTime` fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-columnhierarchy.html#cfn-quicksight-analysis-columnhierarchy-datetimehierarchy
	//
	DateTimeHierarchy interface{} `field:"optional" json:"dateTimeHierarchy" yaml:"dateTimeHierarchy"`
	// The option that determines the hierarchy of the fields that are built within a visual's field wells.
	//
	// These fields can't be duplicated to other visuals.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-columnhierarchy.html#cfn-quicksight-analysis-columnhierarchy-explicithierarchy
	//
	ExplicitHierarchy interface{} `field:"optional" json:"explicitHierarchy" yaml:"explicitHierarchy"`
	// The option that determines the hierarchy of the fields that are defined during data preparation.
	//
	// These fields are available to use in any analysis that uses the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-columnhierarchy.html#cfn-quicksight-analysis-columnhierarchy-predefinedhierarchy
	//
	PredefinedHierarchy interface{} `field:"optional" json:"predefinedHierarchy" yaml:"predefinedHierarchy"`
}

