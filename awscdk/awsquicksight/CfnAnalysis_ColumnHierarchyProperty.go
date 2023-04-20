package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnHierarchyProperty := &ColumnHierarchyProperty{
//   	DateTimeHierarchy: &DateTimeHierarchyProperty{
//   		HierarchyId: jsii.String("hierarchyId"),
//
//   		// the properties below are optional
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
//   	},
//   	ExplicitHierarchy: &ExplicitHierarchyProperty{
//   		Columns: []interface{}{
//   			&ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   		},
//   		HierarchyId: jsii.String("hierarchyId"),
//
//   		// the properties below are optional
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
//   	},
//   	PredefinedHierarchy: &PredefinedHierarchyProperty{
//   		Columns: []interface{}{
//   			&ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   		},
//   		HierarchyId: jsii.String("hierarchyId"),
//
//   		// the properties below are optional
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
//   	},
//   }
//
type CfnAnalysis_ColumnHierarchyProperty struct {
	// `CfnAnalysis.ColumnHierarchyProperty.DateTimeHierarchy`.
	DateTimeHierarchy interface{} `field:"optional" json:"dateTimeHierarchy" yaml:"dateTimeHierarchy"`
	// `CfnAnalysis.ColumnHierarchyProperty.ExplicitHierarchy`.
	ExplicitHierarchy interface{} `field:"optional" json:"explicitHierarchy" yaml:"explicitHierarchy"`
	// `CfnAnalysis.ColumnHierarchyProperty.PredefinedHierarchy`.
	PredefinedHierarchy interface{} `field:"optional" json:"predefinedHierarchy" yaml:"predefinedHierarchy"`
}

