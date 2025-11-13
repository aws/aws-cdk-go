package awsquicksight


// A step in data preparation that performs a specific operation on the data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transformStepProperty := &TransformStepProperty{
//   	AggregateStep: &AggregateOperationProperty{
//   		Aggregations: []interface{}{
//   			&AggregationProperty{
//   				AggregationFunction: &DataPrepAggregationFunctionProperty{
//   					ListAggregation: &DataPrepListAggregationFunctionProperty{
//   						Distinct: jsii.Boolean(false),
//   						Separator: jsii.String("separator"),
//
//   						// the properties below are optional
//   						InputColumnName: jsii.String("inputColumnName"),
//   					},
//   					PercentileAggregation: &DataPrepPercentileAggregationFunctionProperty{
//   						PercentileValue: jsii.Number(123),
//
//   						// the properties below are optional
//   						InputColumnName: jsii.String("inputColumnName"),
//   					},
//   					SimpleAggregation: &DataPrepSimpleAggregationFunctionProperty{
//   						FunctionType: jsii.String("functionType"),
//
//   						// the properties below are optional
//   						InputColumnName: jsii.String("inputColumnName"),
//   					},
//   				},
//   				NewColumnId: jsii.String("newColumnId"),
//   				NewColumnName: jsii.String("newColumnName"),
//   			},
//   		},
//   		Alias: jsii.String("alias"),
//   		Source: &TransformOperationSourceProperty{
//   			TransformOperationId: jsii.String("transformOperationId"),
//
//   			// the properties below are optional
//   			ColumnIdMappings: []interface{}{
//   				&DataSetColumnIdMappingProperty{
//   					SourceColumnId: jsii.String("sourceColumnId"),
//   					TargetColumnId: jsii.String("targetColumnId"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		GroupByColumnNames: []*string{
//   			jsii.String("groupByColumnNames"),
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
//
//   		// the properties below are optional
//   		FirstSource: &TransformOperationSourceProperty{
//   			TransformOperationId: jsii.String("transformOperationId"),
//
//   			// the properties below are optional
//   			ColumnIdMappings: []interface{}{
//   				&DataSetColumnIdMappingProperty{
//   					SourceColumnId: jsii.String("sourceColumnId"),
//   					TargetColumnId: jsii.String("targetColumnId"),
//   				},
//   			},
//   		},
//   		SecondSource: &TransformOperationSourceProperty{
//   			TransformOperationId: jsii.String("transformOperationId"),
//
//   			// the properties below are optional
//   			ColumnIdMappings: []interface{}{
//   				&DataSetColumnIdMappingProperty{
//   					SourceColumnId: jsii.String("sourceColumnId"),
//   					TargetColumnId: jsii.String("targetColumnId"),
//   				},
//   			},
//   		},
//   	},
//   	CastColumnTypesStep: &CastColumnTypesOperationProperty{
//   		Alias: jsii.String("alias"),
//   		CastColumnTypeOperations: []interface{}{
//   			&CastColumnTypeOperationProperty{
//   				ColumnName: jsii.String("columnName"),
//   				NewColumnType: jsii.String("newColumnType"),
//
//   				// the properties below are optional
//   				Format: jsii.String("format"),
//   				SubType: jsii.String("subType"),
//   			},
//   		},
//   		Source: &TransformOperationSourceProperty{
//   			TransformOperationId: jsii.String("transformOperationId"),
//
//   			// the properties below are optional
//   			ColumnIdMappings: []interface{}{
//   				&DataSetColumnIdMappingProperty{
//   					SourceColumnId: jsii.String("sourceColumnId"),
//   					TargetColumnId: jsii.String("targetColumnId"),
//   				},
//   			},
//   		},
//   	},
//   	CreateColumnsStep: &CreateColumnsOperationProperty{
//   		Columns: []interface{}{
//   			&CalculatedColumnProperty{
//   				ColumnId: jsii.String("columnId"),
//   				ColumnName: jsii.String("columnName"),
//   				Expression: jsii.String("expression"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		Alias: jsii.String("alias"),
//   		Source: &TransformOperationSourceProperty{
//   			TransformOperationId: jsii.String("transformOperationId"),
//
//   			// the properties below are optional
//   			ColumnIdMappings: []interface{}{
//   				&DataSetColumnIdMappingProperty{
//   					SourceColumnId: jsii.String("sourceColumnId"),
//   					TargetColumnId: jsii.String("targetColumnId"),
//   				},
//   			},
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
//
//   						// the properties below are optional
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
//
//   						// the properties below are optional
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
//
//   						// the properties below are optional
//   						Value: &DataSetStringFilterValueProperty{
//   							StaticValue: jsii.String("staticValue"),
//   						},
//   					},
//   					ListFilterCondition: &DataSetStringListFilterConditionProperty{
//   						Operator: jsii.String("operator"),
//
//   						// the properties below are optional
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
//   			TransformOperationId: jsii.String("transformOperationId"),
//
//   			// the properties below are optional
//   			ColumnIdMappings: []interface{}{
//   				&DataSetColumnIdMappingProperty{
//   					SourceColumnId: jsii.String("sourceColumnId"),
//   					TargetColumnId: jsii.String("targetColumnId"),
//   				},
//   			},
//   		},
//   	},
//   	ImportTableStep: &ImportTableOperationProperty{
//   		Alias: jsii.String("alias"),
//   		Source: &ImportTableOperationSourceProperty{
//   			SourceTableId: jsii.String("sourceTableId"),
//
//   			// the properties below are optional
//   			ColumnIdMappings: []interface{}{
//   				&DataSetColumnIdMappingProperty{
//   					SourceColumnId: jsii.String("sourceColumnId"),
//   					TargetColumnId: jsii.String("targetColumnId"),
//   				},
//   			},
//   		},
//   	},
//   	JoinStep: &JoinOperationProperty{
//   		Alias: jsii.String("alias"),
//   		LeftOperand: &TransformOperationSourceProperty{
//   			TransformOperationId: jsii.String("transformOperationId"),
//
//   			// the properties below are optional
//   			ColumnIdMappings: []interface{}{
//   				&DataSetColumnIdMappingProperty{
//   					SourceColumnId: jsii.String("sourceColumnId"),
//   					TargetColumnId: jsii.String("targetColumnId"),
//   				},
//   			},
//   		},
//   		OnClause: jsii.String("onClause"),
//   		RightOperand: &TransformOperationSourceProperty{
//   			TransformOperationId: jsii.String("transformOperationId"),
//
//   			// the properties below are optional
//   			ColumnIdMappings: []interface{}{
//   				&DataSetColumnIdMappingProperty{
//   					SourceColumnId: jsii.String("sourceColumnId"),
//   					TargetColumnId: jsii.String("targetColumnId"),
//   				},
//   			},
//   		},
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		LeftOperandProperties: &JoinOperandPropertiesProperty{
//   			OutputColumnNameOverrides: []interface{}{
//   				&OutputColumnNameOverrideProperty{
//   					OutputColumnName: jsii.String("outputColumnName"),
//
//   					// the properties below are optional
//   					SourceColumnName: jsii.String("sourceColumnName"),
//   				},
//   			},
//   		},
//   		RightOperandProperties: &JoinOperandPropertiesProperty{
//   			OutputColumnNameOverrides: []interface{}{
//   				&OutputColumnNameOverrideProperty{
//   					OutputColumnName: jsii.String("outputColumnName"),
//
//   					// the properties below are optional
//   					SourceColumnName: jsii.String("sourceColumnName"),
//   				},
//   			},
//   		},
//   	},
//   	PivotStep: &PivotOperationProperty{
//   		Alias: jsii.String("alias"),
//   		PivotConfiguration: &PivotConfigurationProperty{
//   			PivotedLabels: []interface{}{
//   				&PivotedLabelProperty{
//   					LabelName: jsii.String("labelName"),
//   					NewColumnId: jsii.String("newColumnId"),
//   					NewColumnName: jsii.String("newColumnName"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			LabelColumnName: jsii.String("labelColumnName"),
//   		},
//   		Source: &TransformOperationSourceProperty{
//   			TransformOperationId: jsii.String("transformOperationId"),
//
//   			// the properties below are optional
//   			ColumnIdMappings: []interface{}{
//   				&DataSetColumnIdMappingProperty{
//   					SourceColumnId: jsii.String("sourceColumnId"),
//   					TargetColumnId: jsii.String("targetColumnId"),
//   				},
//   			},
//   		},
//   		ValueColumnConfiguration: &ValueColumnConfigurationProperty{
//   			AggregationFunction: &DataPrepAggregationFunctionProperty{
//   				ListAggregation: &DataPrepListAggregationFunctionProperty{
//   					Distinct: jsii.Boolean(false),
//   					Separator: jsii.String("separator"),
//
//   					// the properties below are optional
//   					InputColumnName: jsii.String("inputColumnName"),
//   				},
//   				PercentileAggregation: &DataPrepPercentileAggregationFunctionProperty{
//   					PercentileValue: jsii.Number(123),
//
//   					// the properties below are optional
//   					InputColumnName: jsii.String("inputColumnName"),
//   				},
//   				SimpleAggregation: &DataPrepSimpleAggregationFunctionProperty{
//   					FunctionType: jsii.String("functionType"),
//
//   					// the properties below are optional
//   					InputColumnName: jsii.String("inputColumnName"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		GroupByColumnNames: []*string{
//   			jsii.String("groupByColumnNames"),
//   		},
//   	},
//   	ProjectStep: &ProjectOperationProperty{
//   		Alias: jsii.String("alias"),
//   		ProjectedColumns: []*string{
//   			jsii.String("projectedColumns"),
//   		},
//   		Source: &TransformOperationSourceProperty{
//   			TransformOperationId: jsii.String("transformOperationId"),
//
//   			// the properties below are optional
//   			ColumnIdMappings: []interface{}{
//   				&DataSetColumnIdMappingProperty{
//   					SourceColumnId: jsii.String("sourceColumnId"),
//   					TargetColumnId: jsii.String("targetColumnId"),
//   				},
//   			},
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
//   			TransformOperationId: jsii.String("transformOperationId"),
//
//   			// the properties below are optional
//   			ColumnIdMappings: []interface{}{
//   				&DataSetColumnIdMappingProperty{
//   					SourceColumnId: jsii.String("sourceColumnId"),
//   					TargetColumnId: jsii.String("targetColumnId"),
//   				},
//   			},
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
//   			TransformOperationId: jsii.String("transformOperationId"),
//
//   			// the properties below are optional
//   			ColumnIdMappings: []interface{}{
//   				&DataSetColumnIdMappingProperty{
//   					SourceColumnId: jsii.String("sourceColumnId"),
//   					TargetColumnId: jsii.String("targetColumnId"),
//   				},
//   			},
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
type CfnDataSet_TransformStepProperty struct {
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

