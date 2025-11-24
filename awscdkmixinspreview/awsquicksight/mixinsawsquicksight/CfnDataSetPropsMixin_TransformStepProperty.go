package mixinsawsquicksight


// A step in data preparation that performs a specific operation on the data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   transformStepProperty := &TransformStepProperty{
//   	AggregateStep: &AggregateOperationProperty{
//   		Aggregations: []interface{}{
//   			&AggregationProperty{
//   				AggregationFunction: &DataPrepAggregationFunctionProperty{
//   					ListAggregation: &DataPrepListAggregationFunctionProperty{
//   						Distinct: jsii.Boolean(false),
//   						InputColumnName: jsii.String("inputColumnName"),
//   						Separator: jsii.String("separator"),
//   					},
//   					PercentileAggregation: &DataPrepPercentileAggregationFunctionProperty{
//   						InputColumnName: jsii.String("inputColumnName"),
//   						PercentileValue: jsii.Number(123),
//   					},
//   					SimpleAggregation: &DataPrepSimpleAggregationFunctionProperty{
//   						FunctionType: jsii.String("functionType"),
//   						InputColumnName: jsii.String("inputColumnName"),
//   					},
//   				},
//   				NewColumnId: jsii.String("newColumnId"),
//   				NewColumnName: jsii.String("newColumnName"),
//   			},
//   		},
//   		Alias: jsii.String("alias"),
//   		GroupByColumnNames: []*string{
//   			jsii.String("groupByColumnNames"),
//   		},
//   		Source: &TransformOperationSourceProperty{
//   			ColumnIdMappings: []interface{}{
//   				&DataSetColumnIdMappingProperty{
//   					SourceColumnId: jsii.String("sourceColumnId"),
//   					TargetColumnId: jsii.String("targetColumnId"),
//   				},
//   			},
//   			TransformOperationId: jsii.String("transformOperationId"),
//   		},
//   	},
//   	AppendStep: &AppendOperationProperty{
//   		Alias: jsii.String("alias"),
//   		AppendedColumns: []interface{}{
//   			&AppendedColumnProperty{
//   				ColumnName: jsii.String("columnName"),
//   				NewColumnId: jsii.String("newColumnId"),
//   			},
//   		},
//   		FirstSource: &TransformOperationSourceProperty{
//   			ColumnIdMappings: []interface{}{
//   				&DataSetColumnIdMappingProperty{
//   					SourceColumnId: jsii.String("sourceColumnId"),
//   					TargetColumnId: jsii.String("targetColumnId"),
//   				},
//   			},
//   			TransformOperationId: jsii.String("transformOperationId"),
//   		},
//   		SecondSource: &TransformOperationSourceProperty{
//   			ColumnIdMappings: []interface{}{
//   				&DataSetColumnIdMappingProperty{
//   					SourceColumnId: jsii.String("sourceColumnId"),
//   					TargetColumnId: jsii.String("targetColumnId"),
//   				},
//   			},
//   			TransformOperationId: jsii.String("transformOperationId"),
//   		},
//   	},
//   	CastColumnTypesStep: &CastColumnTypesOperationProperty{
//   		Alias: jsii.String("alias"),
//   		CastColumnTypeOperations: []interface{}{
//   			&CastColumnTypeOperationProperty{
//   				ColumnName: jsii.String("columnName"),
//   				Format: jsii.String("format"),
//   				NewColumnType: jsii.String("newColumnType"),
//   				SubType: jsii.String("subType"),
//   			},
//   		},
//   		Source: &TransformOperationSourceProperty{
//   			ColumnIdMappings: []interface{}{
//   				&DataSetColumnIdMappingProperty{
//   					SourceColumnId: jsii.String("sourceColumnId"),
//   					TargetColumnId: jsii.String("targetColumnId"),
//   				},
//   			},
//   			TransformOperationId: jsii.String("transformOperationId"),
//   		},
//   	},
//   	CreateColumnsStep: &CreateColumnsOperationProperty{
//   		Alias: jsii.String("alias"),
//   		Columns: []interface{}{
//   			&CalculatedColumnProperty{
//   				ColumnId: jsii.String("columnId"),
//   				ColumnName: jsii.String("columnName"),
//   				Expression: jsii.String("expression"),
//   			},
//   		},
//   		Source: &TransformOperationSourceProperty{
//   			ColumnIdMappings: []interface{}{
//   				&DataSetColumnIdMappingProperty{
//   					SourceColumnId: jsii.String("sourceColumnId"),
//   					TargetColumnId: jsii.String("targetColumnId"),
//   				},
//   			},
//   			TransformOperationId: jsii.String("transformOperationId"),
//   		},
//   	},
//   	FiltersStep: &FiltersOperationProperty{
//   		Alias: jsii.String("alias"),
//   		FilterOperations: []interface{}{
//   			&FilterOperationProperty{
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
//   		},
//   		Source: &TransformOperationSourceProperty{
//   			ColumnIdMappings: []interface{}{
//   				&DataSetColumnIdMappingProperty{
//   					SourceColumnId: jsii.String("sourceColumnId"),
//   					TargetColumnId: jsii.String("targetColumnId"),
//   				},
//   			},
//   			TransformOperationId: jsii.String("transformOperationId"),
//   		},
//   	},
//   	ImportTableStep: &ImportTableOperationProperty{
//   		Alias: jsii.String("alias"),
//   		Source: &ImportTableOperationSourceProperty{
//   			ColumnIdMappings: []interface{}{
//   				&DataSetColumnIdMappingProperty{
//   					SourceColumnId: jsii.String("sourceColumnId"),
//   					TargetColumnId: jsii.String("targetColumnId"),
//   				},
//   			},
//   			SourceTableId: jsii.String("sourceTableId"),
//   		},
//   	},
//   	JoinStep: &JoinOperationProperty{
//   		Alias: jsii.String("alias"),
//   		LeftOperand: &TransformOperationSourceProperty{
//   			ColumnIdMappings: []interface{}{
//   				&DataSetColumnIdMappingProperty{
//   					SourceColumnId: jsii.String("sourceColumnId"),
//   					TargetColumnId: jsii.String("targetColumnId"),
//   				},
//   			},
//   			TransformOperationId: jsii.String("transformOperationId"),
//   		},
//   		LeftOperandProperties: &JoinOperandPropertiesProperty{
//   			OutputColumnNameOverrides: []interface{}{
//   				&OutputColumnNameOverrideProperty{
//   					OutputColumnName: jsii.String("outputColumnName"),
//   					SourceColumnName: jsii.String("sourceColumnName"),
//   				},
//   			},
//   		},
//   		OnClause: jsii.String("onClause"),
//   		RightOperand: &TransformOperationSourceProperty{
//   			ColumnIdMappings: []interface{}{
//   				&DataSetColumnIdMappingProperty{
//   					SourceColumnId: jsii.String("sourceColumnId"),
//   					TargetColumnId: jsii.String("targetColumnId"),
//   				},
//   			},
//   			TransformOperationId: jsii.String("transformOperationId"),
//   		},
//   		RightOperandProperties: &JoinOperandPropertiesProperty{
//   			OutputColumnNameOverrides: []interface{}{
//   				&OutputColumnNameOverrideProperty{
//   					OutputColumnName: jsii.String("outputColumnName"),
//   					SourceColumnName: jsii.String("sourceColumnName"),
//   				},
//   			},
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	PivotStep: &PivotOperationProperty{
//   		Alias: jsii.String("alias"),
//   		GroupByColumnNames: []*string{
//   			jsii.String("groupByColumnNames"),
//   		},
//   		PivotConfiguration: &PivotConfigurationProperty{
//   			LabelColumnName: jsii.String("labelColumnName"),
//   			PivotedLabels: []interface{}{
//   				&PivotedLabelProperty{
//   					LabelName: jsii.String("labelName"),
//   					NewColumnId: jsii.String("newColumnId"),
//   					NewColumnName: jsii.String("newColumnName"),
//   				},
//   			},
//   		},
//   		Source: &TransformOperationSourceProperty{
//   			ColumnIdMappings: []interface{}{
//   				&DataSetColumnIdMappingProperty{
//   					SourceColumnId: jsii.String("sourceColumnId"),
//   					TargetColumnId: jsii.String("targetColumnId"),
//   				},
//   			},
//   			TransformOperationId: jsii.String("transformOperationId"),
//   		},
//   		ValueColumnConfiguration: &ValueColumnConfigurationProperty{
//   			AggregationFunction: &DataPrepAggregationFunctionProperty{
//   				ListAggregation: &DataPrepListAggregationFunctionProperty{
//   					Distinct: jsii.Boolean(false),
//   					InputColumnName: jsii.String("inputColumnName"),
//   					Separator: jsii.String("separator"),
//   				},
//   				PercentileAggregation: &DataPrepPercentileAggregationFunctionProperty{
//   					InputColumnName: jsii.String("inputColumnName"),
//   					PercentileValue: jsii.Number(123),
//   				},
//   				SimpleAggregation: &DataPrepSimpleAggregationFunctionProperty{
//   					FunctionType: jsii.String("functionType"),
//   					InputColumnName: jsii.String("inputColumnName"),
//   				},
//   			},
//   		},
//   	},
//   	ProjectStep: &ProjectOperationProperty{
//   		Alias: jsii.String("alias"),
//   		ProjectedColumns: []*string{
//   			jsii.String("projectedColumns"),
//   		},
//   		Source: &TransformOperationSourceProperty{
//   			ColumnIdMappings: []interface{}{
//   				&DataSetColumnIdMappingProperty{
//   					SourceColumnId: jsii.String("sourceColumnId"),
//   					TargetColumnId: jsii.String("targetColumnId"),
//   				},
//   			},
//   			TransformOperationId: jsii.String("transformOperationId"),
//   		},
//   	},
//   	RenameColumnsStep: &RenameColumnsOperationProperty{
//   		Alias: jsii.String("alias"),
//   		RenameColumnOperations: []interface{}{
//   			&RenameColumnOperationProperty{
//   				ColumnName: jsii.String("columnName"),
//   				NewColumnName: jsii.String("newColumnName"),
//   			},
//   		},
//   		Source: &TransformOperationSourceProperty{
//   			ColumnIdMappings: []interface{}{
//   				&DataSetColumnIdMappingProperty{
//   					SourceColumnId: jsii.String("sourceColumnId"),
//   					TargetColumnId: jsii.String("targetColumnId"),
//   				},
//   			},
//   			TransformOperationId: jsii.String("transformOperationId"),
//   		},
//   	},
//   	UnpivotStep: &UnpivotOperationProperty{
//   		Alias: jsii.String("alias"),
//   		ColumnsToUnpivot: []interface{}{
//   			&ColumnToUnpivotProperty{
//   				ColumnName: jsii.String("columnName"),
//   				NewValue: jsii.String("newValue"),
//   			},
//   		},
//   		Source: &TransformOperationSourceProperty{
//   			ColumnIdMappings: []interface{}{
//   				&DataSetColumnIdMappingProperty{
//   					SourceColumnId: jsii.String("sourceColumnId"),
//   					TargetColumnId: jsii.String("targetColumnId"),
//   				},
//   			},
//   			TransformOperationId: jsii.String("transformOperationId"),
//   		},
//   		UnpivotedLabelColumnId: jsii.String("unpivotedLabelColumnId"),
//   		UnpivotedLabelColumnName: jsii.String("unpivotedLabelColumnName"),
//   		UnpivotedValueColumnId: jsii.String("unpivotedValueColumnId"),
//   		UnpivotedValueColumnName: jsii.String("unpivotedValueColumnName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-transformstep.html
//
type CfnDataSetPropsMixin_TransformStepProperty struct {
	// A transform step that groups data and applies aggregation functions to calculate summary values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-transformstep.html#cfn-quicksight-dataset-transformstep-aggregatestep
	//
	AggregateStep interface{} `field:"optional" json:"aggregateStep" yaml:"aggregateStep"`
	// A transform step that combines rows from multiple sources by stacking them vertically.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-transformstep.html#cfn-quicksight-dataset-transformstep-appendstep
	//
	AppendStep interface{} `field:"optional" json:"appendStep" yaml:"appendStep"`
	// A transform step that changes the data types of one or more columns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-transformstep.html#cfn-quicksight-dataset-transformstep-castcolumntypesstep
	//
	CastColumnTypesStep interface{} `field:"optional" json:"castColumnTypesStep" yaml:"castColumnTypesStep"`
	// <p>A transform operation that creates calculated columns.
	//
	// Columns created in one such
	//             operation form a lexical closure.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-transformstep.html#cfn-quicksight-dataset-transformstep-createcolumnsstep
	//
	CreateColumnsStep interface{} `field:"optional" json:"createColumnsStep" yaml:"createColumnsStep"`
	// A transform step that applies filter conditions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-transformstep.html#cfn-quicksight-dataset-transformstep-filtersstep
	//
	FiltersStep interface{} `field:"optional" json:"filtersStep" yaml:"filtersStep"`
	// A transform step that brings data from a source table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-transformstep.html#cfn-quicksight-dataset-transformstep-importtablestep
	//
	ImportTableStep interface{} `field:"optional" json:"importTableStep" yaml:"importTableStep"`
	// A transform step that combines data from two sources based on specified join conditions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-transformstep.html#cfn-quicksight-dataset-transformstep-joinstep
	//
	JoinStep interface{} `field:"optional" json:"joinStep" yaml:"joinStep"`
	// A transform step that converts row values into columns to reshape the data structure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-transformstep.html#cfn-quicksight-dataset-transformstep-pivotstep
	//
	PivotStep interface{} `field:"optional" json:"pivotStep" yaml:"pivotStep"`
	// <p>A transform operation that projects columns.
	//
	// Operations that come after a projection
	//             can only refer to projected columns.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-transformstep.html#cfn-quicksight-dataset-transformstep-projectstep
	//
	ProjectStep interface{} `field:"optional" json:"projectStep" yaml:"projectStep"`
	// A transform step that changes the names of one or more columns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-transformstep.html#cfn-quicksight-dataset-transformstep-renamecolumnsstep
	//
	RenameColumnsStep interface{} `field:"optional" json:"renameColumnsStep" yaml:"renameColumnsStep"`
	// A transform step that converts columns into rows to normalize the data structure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-transformstep.html#cfn-quicksight-dataset-transformstep-unpivotstep
	//
	UnpivotStep interface{} `field:"optional" json:"unpivotStep" yaml:"unpivotStep"`
}

