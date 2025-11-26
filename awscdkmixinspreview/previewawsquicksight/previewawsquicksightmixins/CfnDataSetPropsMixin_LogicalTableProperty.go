package previewawsquicksightmixins


// <p>A <i>logical table</i> is a unit that joins and that data             transformations operate on.
//
// A logical table has a source, which can be either a physical
//             table or result of a join. When a logical table points to a physical table, the logical
//             table acts as a mutable copy of that physical table through transform operations.</p>
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   logicalTableProperty := &LogicalTableProperty{
//   	Alias: jsii.String("alias"),
//   	DataTransforms: []interface{}{
//   		&TransformOperationProperty{
//   			CastColumnTypeOperation: &CastColumnTypeOperationProperty{
//   				ColumnName: jsii.String("columnName"),
//   				Format: jsii.String("format"),
//   				NewColumnType: jsii.String("newColumnType"),
//   				SubType: jsii.String("subType"),
//   			},
//   			CreateColumnsOperation: &CreateColumnsOperationProperty{
//   				Alias: jsii.String("alias"),
//   				Columns: []interface{}{
//   					&CalculatedColumnProperty{
//   						ColumnId: jsii.String("columnId"),
//   						ColumnName: jsii.String("columnName"),
//   						Expression: jsii.String("expression"),
//   					},
//   				},
//   				Source: &TransformOperationSourceProperty{
//   					ColumnIdMappings: []interface{}{
//   						&DataSetColumnIdMappingProperty{
//   							SourceColumnId: jsii.String("sourceColumnId"),
//   							TargetColumnId: jsii.String("targetColumnId"),
//   						},
//   					},
//   					TransformOperationId: jsii.String("transformOperationId"),
//   				},
//   			},
//   			FilterOperation: &FilterOperationProperty{
//   				ConditionExpression: jsii.String("conditionExpression"),
//   				DateFilterCondition: &DataSetDateFilterConditionProperty{
//   					ColumnName: jsii.String("columnName"),
//   					ComparisonFilterCondition: &DataSetDateComparisonFilterConditionProperty{
//   						Operator: jsii.String("operator"),
//   						Value: &DataSetDateFilterValueProperty{
//   							StaticValue: jsii.String("staticValue"),
//   						},
//   					},
//   					RangeFilterCondition: &DataSetDateRangeFilterConditionProperty{
//   						IncludeMaximum: jsii.Boolean(false),
//   						IncludeMinimum: jsii.Boolean(false),
//   						RangeMaximum: &DataSetDateFilterValueProperty{
//   							StaticValue: jsii.String("staticValue"),
//   						},
//   						RangeMinimum: &DataSetDateFilterValueProperty{
//   							StaticValue: jsii.String("staticValue"),
//   						},
//   					},
//   				},
//   				NumericFilterCondition: &DataSetNumericFilterConditionProperty{
//   					ColumnName: jsii.String("columnName"),
//   					ComparisonFilterCondition: &DataSetNumericComparisonFilterConditionProperty{
//   						Operator: jsii.String("operator"),
//   						Value: &DataSetNumericFilterValueProperty{
//   							StaticValue: jsii.Number(123),
//   						},
//   					},
//   					RangeFilterCondition: &DataSetNumericRangeFilterConditionProperty{
//   						IncludeMaximum: jsii.Boolean(false),
//   						IncludeMinimum: jsii.Boolean(false),
//   						RangeMaximum: &DataSetNumericFilterValueProperty{
//   							StaticValue: jsii.Number(123),
//   						},
//   						RangeMinimum: &DataSetNumericFilterValueProperty{
//   							StaticValue: jsii.Number(123),
//   						},
//   					},
//   				},
//   				StringFilterCondition: &DataSetStringFilterConditionProperty{
//   					ColumnName: jsii.String("columnName"),
//   					ComparisonFilterCondition: &DataSetStringComparisonFilterConditionProperty{
//   						Operator: jsii.String("operator"),
//   						Value: &DataSetStringFilterValueProperty{
//   							StaticValue: jsii.String("staticValue"),
//   						},
//   					},
//   					ListFilterCondition: &DataSetStringListFilterConditionProperty{
//   						Operator: jsii.String("operator"),
//   						Values: &DataSetStringListFilterValueProperty{
//   							StaticValues: []*string{
//   								jsii.String("staticValues"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			OverrideDatasetParameterOperation: &OverrideDatasetParameterOperationProperty{
//   				NewDefaultValues: &NewDefaultValuesProperty{
//   					DateTimeStaticValues: []*string{
//   						jsii.String("dateTimeStaticValues"),
//   					},
//   					DecimalStaticValues: []interface{}{
//   						jsii.Number(123),
//   					},
//   					IntegerStaticValues: []interface{}{
//   						jsii.Number(123),
//   					},
//   					StringStaticValues: []*string{
//   						jsii.String("stringStaticValues"),
//   					},
//   				},
//   				NewParameterName: jsii.String("newParameterName"),
//   				ParameterName: jsii.String("parameterName"),
//   			},
//   			ProjectOperation: &ProjectOperationProperty{
//   				Alias: jsii.String("alias"),
//   				ProjectedColumns: []*string{
//   					jsii.String("projectedColumns"),
//   				},
//   				Source: &TransformOperationSourceProperty{
//   					ColumnIdMappings: []interface{}{
//   						&DataSetColumnIdMappingProperty{
//   							SourceColumnId: jsii.String("sourceColumnId"),
//   							TargetColumnId: jsii.String("targetColumnId"),
//   						},
//   					},
//   					TransformOperationId: jsii.String("transformOperationId"),
//   				},
//   			},
//   			RenameColumnOperation: &RenameColumnOperationProperty{
//   				ColumnName: jsii.String("columnName"),
//   				NewColumnName: jsii.String("newColumnName"),
//   			},
//   			TagColumnOperation: &TagColumnOperationProperty{
//   				ColumnName: jsii.String("columnName"),
//   				Tags: []ColumnTagProperty{
//   					&ColumnTagProperty{
//   						ColumnDescription: &ColumnDescriptionProperty{
//   							Text: jsii.String("text"),
//   						},
//   						ColumnGeographicRole: jsii.String("columnGeographicRole"),
//   					},
//   				},
//   			},
//   			UntagColumnOperation: &UntagColumnOperationProperty{
//   				ColumnName: jsii.String("columnName"),
//   				TagNames: []*string{
//   					jsii.String("tagNames"),
//   				},
//   			},
//   		},
//   	},
//   	Source: &LogicalTableSourceProperty{
//   		DataSetArn: jsii.String("dataSetArn"),
//   		JoinInstruction: &JoinInstructionProperty{
//   			LeftJoinKeyProperties: &JoinKeyPropertiesProperty{
//   				UniqueKey: jsii.Boolean(false),
//   			},
//   			LeftOperand: jsii.String("leftOperand"),
//   			OnClause: jsii.String("onClause"),
//   			RightJoinKeyProperties: &JoinKeyPropertiesProperty{
//   				UniqueKey: jsii.Boolean(false),
//   			},
//   			RightOperand: jsii.String("rightOperand"),
//   			Type: jsii.String("type"),
//   		},
//   		PhysicalTableId: jsii.String("physicalTableId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-logicaltable.html
//
type CfnDataSetPropsMixin_LogicalTableProperty struct {
	// <p>A display name for the logical table.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-logicaltable.html#cfn-quicksight-dataset-logicaltable-alias
	//
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// <p>Transform operations that act on this logical table.
	//
	// For this structure to be valid, only one of the attributes can be non-null. </p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-logicaltable.html#cfn-quicksight-dataset-logicaltable-datatransforms
	//
	DataTransforms interface{} `field:"optional" json:"dataTransforms" yaml:"dataTransforms"`
	// <p>Information about the source of a logical table.
	//
	// This is a variant type structure. For
	//             this structure to be valid, only one of the attributes can be non-null.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-logicaltable.html#cfn-quicksight-dataset-logicaltable-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
}

