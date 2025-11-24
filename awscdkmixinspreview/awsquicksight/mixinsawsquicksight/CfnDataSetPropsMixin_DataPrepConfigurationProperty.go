package mixinsawsquicksight


// Configuration for data preparation operations, defining the complete pipeline from source tables through transformations to destination tables.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataPrepConfigurationProperty := &DataPrepConfigurationProperty{
//   	DestinationTableMap: map[string]interface{}{
//   		"destinationTableMapKey": &DestinationTableProperty{
//   			"alias": jsii.String("alias"),
//   			"source": &DestinationTableSourceProperty{
//   				"transformOperationId": jsii.String("transformOperationId"),
//   			},
//   		},
//   	},
//   	SourceTableMap: map[string]interface{}{
//   		"sourceTableMapKey": &SourceTableProperty{
//   			"dataSet": &ParentDataSetProperty{
//   				"dataSetArn": jsii.String("dataSetArn"),
//   				"inputColumns": []interface{}{
//   					&InputColumnProperty{
//   						"id": jsii.String("id"),
//   						"name": jsii.String("name"),
//   						"subType": jsii.String("subType"),
//   						"type": jsii.String("type"),
//   					},
//   				},
//   			},
//   			"physicalTableId": jsii.String("physicalTableId"),
//   		},
//   	},
//   	TransformStepMap: map[string]interface{}{
//   		"transformStepMapKey": &TransformStepProperty{
//   			"aggregateStep": &AggregateOperationProperty{
//   				"aggregations": []interface{}{
//   					&AggregationProperty{
//   						"aggregationFunction": &DataPrepAggregationFunctionProperty{
//   							"listAggregation": &DataPrepListAggregationFunctionProperty{
//   								"distinct": jsii.Boolean(false),
//   								"inputColumnName": jsii.String("inputColumnName"),
//   								"separator": jsii.String("separator"),
//   							},
//   							"percentileAggregation": &DataPrepPercentileAggregationFunctionProperty{
//   								"inputColumnName": jsii.String("inputColumnName"),
//   								"percentileValue": jsii.Number(123),
//   							},
//   							"simpleAggregation": &DataPrepSimpleAggregationFunctionProperty{
//   								"functionType": jsii.String("functionType"),
//   								"inputColumnName": jsii.String("inputColumnName"),
//   							},
//   						},
//   						"newColumnId": jsii.String("newColumnId"),
//   						"newColumnName": jsii.String("newColumnName"),
//   					},
//   				},
//   				"alias": jsii.String("alias"),
//   				"groupByColumnNames": []*string{
//   					jsii.String("groupByColumnNames"),
//   				},
//   				"source": &TransformOperationSourceProperty{
//   					"columnIdMappings": []interface{}{
//   						&DataSetColumnIdMappingProperty{
//   							"sourceColumnId": jsii.String("sourceColumnId"),
//   							"targetColumnId": jsii.String("targetColumnId"),
//   						},
//   					},
//   					"transformOperationId": jsii.String("transformOperationId"),
//   				},
//   			},
//   			"appendStep": &AppendOperationProperty{
//   				"alias": jsii.String("alias"),
//   				"appendedColumns": []interface{}{
//   					&AppendedColumnProperty{
//   						"columnName": jsii.String("columnName"),
//   						"newColumnId": jsii.String("newColumnId"),
//   					},
//   				},
//   				"firstSource": &TransformOperationSourceProperty{
//   					"columnIdMappings": []interface{}{
//   						&DataSetColumnIdMappingProperty{
//   							"sourceColumnId": jsii.String("sourceColumnId"),
//   							"targetColumnId": jsii.String("targetColumnId"),
//   						},
//   					},
//   					"transformOperationId": jsii.String("transformOperationId"),
//   				},
//   				"secondSource": &TransformOperationSourceProperty{
//   					"columnIdMappings": []interface{}{
//   						&DataSetColumnIdMappingProperty{
//   							"sourceColumnId": jsii.String("sourceColumnId"),
//   							"targetColumnId": jsii.String("targetColumnId"),
//   						},
//   					},
//   					"transformOperationId": jsii.String("transformOperationId"),
//   				},
//   			},
//   			"castColumnTypesStep": &CastColumnTypesOperationProperty{
//   				"alias": jsii.String("alias"),
//   				"castColumnTypeOperations": []interface{}{
//   					&CastColumnTypeOperationProperty{
//   						"columnName": jsii.String("columnName"),
//   						"format": jsii.String("format"),
//   						"newColumnType": jsii.String("newColumnType"),
//   						"subType": jsii.String("subType"),
//   					},
//   				},
//   				"source": &TransformOperationSourceProperty{
//   					"columnIdMappings": []interface{}{
//   						&DataSetColumnIdMappingProperty{
//   							"sourceColumnId": jsii.String("sourceColumnId"),
//   							"targetColumnId": jsii.String("targetColumnId"),
//   						},
//   					},
//   					"transformOperationId": jsii.String("transformOperationId"),
//   				},
//   			},
//   			"createColumnsStep": &CreateColumnsOperationProperty{
//   				"alias": jsii.String("alias"),
//   				"columns": []interface{}{
//   					&CalculatedColumnProperty{
//   						"columnId": jsii.String("columnId"),
//   						"columnName": jsii.String("columnName"),
//   						"expression": jsii.String("expression"),
//   					},
//   				},
//   				"source": &TransformOperationSourceProperty{
//   					"columnIdMappings": []interface{}{
//   						&DataSetColumnIdMappingProperty{
//   							"sourceColumnId": jsii.String("sourceColumnId"),
//   							"targetColumnId": jsii.String("targetColumnId"),
//   						},
//   					},
//   					"transformOperationId": jsii.String("transformOperationId"),
//   				},
//   			},
//   			"filtersStep": &FiltersOperationProperty{
//   				"alias": jsii.String("alias"),
//   				"filterOperations": []interface{}{
//   					&FilterOperationProperty{
//   						"conditionExpression": jsii.String("conditionExpression"),
//   						"dateFilterCondition": &DataSetDateFilterConditionProperty{
//   							"columnName": jsii.String("columnName"),
//   							"comparisonFilterCondition": &DataSetDateComparisonFilterConditionProperty{
//   								"operator": jsii.String("operator"),
//   								"value": &DataSetDateFilterValueProperty{
//   									"staticValue": jsii.String("staticValue"),
//   								},
//   							},
//   							"rangeFilterCondition": &DataSetDateRangeFilterConditionProperty{
//   								"includeMaximum": jsii.Boolean(false),
//   								"includeMinimum": jsii.Boolean(false),
//   								"rangeMaximum": &DataSetDateFilterValueProperty{
//   									"staticValue": jsii.String("staticValue"),
//   								},
//   								"rangeMinimum": &DataSetDateFilterValueProperty{
//   									"staticValue": jsii.String("staticValue"),
//   								},
//   							},
//   						},
//   						"numericFilterCondition": &DataSetNumericFilterConditionProperty{
//   							"columnName": jsii.String("columnName"),
//   							"comparisonFilterCondition": &DataSetNumericComparisonFilterConditionProperty{
//   								"operator": jsii.String("operator"),
//   								"value": &DataSetNumericFilterValueProperty{
//   									"staticValue": jsii.Number(123),
//   								},
//   							},
//   							"rangeFilterCondition": &DataSetNumericRangeFilterConditionProperty{
//   								"includeMaximum": jsii.Boolean(false),
//   								"includeMinimum": jsii.Boolean(false),
//   								"rangeMaximum": &DataSetNumericFilterValueProperty{
//   									"staticValue": jsii.Number(123),
//   								},
//   								"rangeMinimum": &DataSetNumericFilterValueProperty{
//   									"staticValue": jsii.Number(123),
//   								},
//   							},
//   						},
//   						"stringFilterCondition": &DataSetStringFilterConditionProperty{
//   							"columnName": jsii.String("columnName"),
//   							"comparisonFilterCondition": &DataSetStringComparisonFilterConditionProperty{
//   								"operator": jsii.String("operator"),
//   								"value": &DataSetStringFilterValueProperty{
//   									"staticValue": jsii.String("staticValue"),
//   								},
//   							},
//   							"listFilterCondition": &DataSetStringListFilterConditionProperty{
//   								"operator": jsii.String("operator"),
//   								"values": &DataSetStringListFilterValueProperty{
//   									"staticValues": []*string{
//   										jsii.String("staticValues"),
//   									},
//   								},
//   							},
//   						},
//   					},
//   				},
//   				"source": &TransformOperationSourceProperty{
//   					"columnIdMappings": []interface{}{
//   						&DataSetColumnIdMappingProperty{
//   							"sourceColumnId": jsii.String("sourceColumnId"),
//   							"targetColumnId": jsii.String("targetColumnId"),
//   						},
//   					},
//   					"transformOperationId": jsii.String("transformOperationId"),
//   				},
//   			},
//   			"importTableStep": &ImportTableOperationProperty{
//   				"alias": jsii.String("alias"),
//   				"source": &ImportTableOperationSourceProperty{
//   					"columnIdMappings": []interface{}{
//   						&DataSetColumnIdMappingProperty{
//   							"sourceColumnId": jsii.String("sourceColumnId"),
//   							"targetColumnId": jsii.String("targetColumnId"),
//   						},
//   					},
//   					"sourceTableId": jsii.String("sourceTableId"),
//   				},
//   			},
//   			"joinStep": &JoinOperationProperty{
//   				"alias": jsii.String("alias"),
//   				"leftOperand": &TransformOperationSourceProperty{
//   					"columnIdMappings": []interface{}{
//   						&DataSetColumnIdMappingProperty{
//   							"sourceColumnId": jsii.String("sourceColumnId"),
//   							"targetColumnId": jsii.String("targetColumnId"),
//   						},
//   					},
//   					"transformOperationId": jsii.String("transformOperationId"),
//   				},
//   				"leftOperandProperties": &JoinOperandPropertiesProperty{
//   					"outputColumnNameOverrides": []interface{}{
//   						&OutputColumnNameOverrideProperty{
//   							"outputColumnName": jsii.String("outputColumnName"),
//   							"sourceColumnName": jsii.String("sourceColumnName"),
//   						},
//   					},
//   				},
//   				"onClause": jsii.String("onClause"),
//   				"rightOperand": &TransformOperationSourceProperty{
//   					"columnIdMappings": []interface{}{
//   						&DataSetColumnIdMappingProperty{
//   							"sourceColumnId": jsii.String("sourceColumnId"),
//   							"targetColumnId": jsii.String("targetColumnId"),
//   						},
//   					},
//   					"transformOperationId": jsii.String("transformOperationId"),
//   				},
//   				"rightOperandProperties": &JoinOperandPropertiesProperty{
//   					"outputColumnNameOverrides": []interface{}{
//   						&OutputColumnNameOverrideProperty{
//   							"outputColumnName": jsii.String("outputColumnName"),
//   							"sourceColumnName": jsii.String("sourceColumnName"),
//   						},
//   					},
//   				},
//   				"type": jsii.String("type"),
//   			},
//   			"pivotStep": &PivotOperationProperty{
//   				"alias": jsii.String("alias"),
//   				"groupByColumnNames": []*string{
//   					jsii.String("groupByColumnNames"),
//   				},
//   				"pivotConfiguration": &PivotConfigurationProperty{
//   					"labelColumnName": jsii.String("labelColumnName"),
//   					"pivotedLabels": []interface{}{
//   						&PivotedLabelProperty{
//   							"labelName": jsii.String("labelName"),
//   							"newColumnId": jsii.String("newColumnId"),
//   							"newColumnName": jsii.String("newColumnName"),
//   						},
//   					},
//   				},
//   				"source": &TransformOperationSourceProperty{
//   					"columnIdMappings": []interface{}{
//   						&DataSetColumnIdMappingProperty{
//   							"sourceColumnId": jsii.String("sourceColumnId"),
//   							"targetColumnId": jsii.String("targetColumnId"),
//   						},
//   					},
//   					"transformOperationId": jsii.String("transformOperationId"),
//   				},
//   				"valueColumnConfiguration": &ValueColumnConfigurationProperty{
//   					"aggregationFunction": &DataPrepAggregationFunctionProperty{
//   						"listAggregation": &DataPrepListAggregationFunctionProperty{
//   							"distinct": jsii.Boolean(false),
//   							"inputColumnName": jsii.String("inputColumnName"),
//   							"separator": jsii.String("separator"),
//   						},
//   						"percentileAggregation": &DataPrepPercentileAggregationFunctionProperty{
//   							"inputColumnName": jsii.String("inputColumnName"),
//   							"percentileValue": jsii.Number(123),
//   						},
//   						"simpleAggregation": &DataPrepSimpleAggregationFunctionProperty{
//   							"functionType": jsii.String("functionType"),
//   							"inputColumnName": jsii.String("inputColumnName"),
//   						},
//   					},
//   				},
//   			},
//   			"projectStep": &ProjectOperationProperty{
//   				"alias": jsii.String("alias"),
//   				"projectedColumns": []*string{
//   					jsii.String("projectedColumns"),
//   				},
//   				"source": &TransformOperationSourceProperty{
//   					"columnIdMappings": []interface{}{
//   						&DataSetColumnIdMappingProperty{
//   							"sourceColumnId": jsii.String("sourceColumnId"),
//   							"targetColumnId": jsii.String("targetColumnId"),
//   						},
//   					},
//   					"transformOperationId": jsii.String("transformOperationId"),
//   				},
//   			},
//   			"renameColumnsStep": &RenameColumnsOperationProperty{
//   				"alias": jsii.String("alias"),
//   				"renameColumnOperations": []interface{}{
//   					&RenameColumnOperationProperty{
//   						"columnName": jsii.String("columnName"),
//   						"newColumnName": jsii.String("newColumnName"),
//   					},
//   				},
//   				"source": &TransformOperationSourceProperty{
//   					"columnIdMappings": []interface{}{
//   						&DataSetColumnIdMappingProperty{
//   							"sourceColumnId": jsii.String("sourceColumnId"),
//   							"targetColumnId": jsii.String("targetColumnId"),
//   						},
//   					},
//   					"transformOperationId": jsii.String("transformOperationId"),
//   				},
//   			},
//   			"unpivotStep": &UnpivotOperationProperty{
//   				"alias": jsii.String("alias"),
//   				"columnsToUnpivot": []interface{}{
//   					&ColumnToUnpivotProperty{
//   						"columnName": jsii.String("columnName"),
//   						"newValue": jsii.String("newValue"),
//   					},
//   				},
//   				"source": &TransformOperationSourceProperty{
//   					"columnIdMappings": []interface{}{
//   						&DataSetColumnIdMappingProperty{
//   							"sourceColumnId": jsii.String("sourceColumnId"),
//   							"targetColumnId": jsii.String("targetColumnId"),
//   						},
//   					},
//   					"transformOperationId": jsii.String("transformOperationId"),
//   				},
//   				"unpivotedLabelColumnId": jsii.String("unpivotedLabelColumnId"),
//   				"unpivotedLabelColumnName": jsii.String("unpivotedLabelColumnName"),
//   				"unpivotedValueColumnId": jsii.String("unpivotedValueColumnId"),
//   				"unpivotedValueColumnName": jsii.String("unpivotedValueColumnName"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-dataprepconfiguration.html
//
type CfnDataSetPropsMixin_DataPrepConfigurationProperty struct {
	// A map of destination tables that receive the final prepared data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-dataprepconfiguration.html#cfn-quicksight-dataset-dataprepconfiguration-destinationtablemap
	//
	DestinationTableMap interface{} `field:"optional" json:"destinationTableMap" yaml:"destinationTableMap"`
	// A map of source tables that provide information about underlying sources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-dataprepconfiguration.html#cfn-quicksight-dataset-dataprepconfiguration-sourcetablemap
	//
	SourceTableMap interface{} `field:"optional" json:"sourceTableMap" yaml:"sourceTableMap"`
	// A map of transformation steps that process the data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-dataprepconfiguration.html#cfn-quicksight-dataset-dataprepconfiguration-transformstepmap
	//
	TransformStepMap interface{} `field:"optional" json:"transformStepMap" yaml:"transformStepMap"`
}

