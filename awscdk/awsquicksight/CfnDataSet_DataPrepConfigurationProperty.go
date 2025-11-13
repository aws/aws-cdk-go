package awsquicksight


// Configuration for data preparation operations, defining the complete pipeline from source tables through transformations to destination tables.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//   						"name": jsii.String("name"),
//   						"type": jsii.String("type"),
//
//   						// the properties below are optional
//   						"id": jsii.String("id"),
//   						"subType": jsii.String("subType"),
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
//   								"separator": jsii.String("separator"),
//
//   								// the properties below are optional
//   								"inputColumnName": jsii.String("inputColumnName"),
//   							},
//   							"percentileAggregation": &DataPrepPercentileAggregationFunctionProperty{
//   								"percentileValue": jsii.Number(123),
//
//   								// the properties below are optional
//   								"inputColumnName": jsii.String("inputColumnName"),
//   							},
//   							"simpleAggregation": &DataPrepSimpleAggregationFunctionProperty{
//   								"functionType": jsii.String("functionType"),
//
//   								// the properties below are optional
//   								"inputColumnName": jsii.String("inputColumnName"),
//   							},
//   						},
//   						"newColumnId": jsii.String("newColumnId"),
//   						"newColumnName": jsii.String("newColumnName"),
//   					},
//   				},
//   				"alias": jsii.String("alias"),
//   				"source": &TransformOperationSourceProperty{
//   					"transformOperationId": jsii.String("transformOperationId"),
//
//   					// the properties below are optional
//   					"columnIdMappings": []interface{}{
//   						&DataSetColumnIdMappingProperty{
//   							"sourceColumnId": jsii.String("sourceColumnId"),
//   							"targetColumnId": jsii.String("targetColumnId"),
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				"groupByColumnNames": []*string{
//   					jsii.String("groupByColumnNames"),
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
//
//   				// the properties below are optional
//   				"firstSource": &TransformOperationSourceProperty{
//   					"transformOperationId": jsii.String("transformOperationId"),
//
//   					// the properties below are optional
//   					"columnIdMappings": []interface{}{
//   						&DataSetColumnIdMappingProperty{
//   							"sourceColumnId": jsii.String("sourceColumnId"),
//   							"targetColumnId": jsii.String("targetColumnId"),
//   						},
//   					},
//   				},
//   				"secondSource": &TransformOperationSourceProperty{
//   					"transformOperationId": jsii.String("transformOperationId"),
//
//   					// the properties below are optional
//   					"columnIdMappings": []interface{}{
//   						&DataSetColumnIdMappingProperty{
//   							"sourceColumnId": jsii.String("sourceColumnId"),
//   							"targetColumnId": jsii.String("targetColumnId"),
//   						},
//   					},
//   				},
//   			},
//   			"castColumnTypesStep": &CastColumnTypesOperationProperty{
//   				"alias": jsii.String("alias"),
//   				"castColumnTypeOperations": []interface{}{
//   					&CastColumnTypeOperationProperty{
//   						"columnName": jsii.String("columnName"),
//   						"newColumnType": jsii.String("newColumnType"),
//
//   						// the properties below are optional
//   						"format": jsii.String("format"),
//   						"subType": jsii.String("subType"),
//   					},
//   				},
//   				"source": &TransformOperationSourceProperty{
//   					"transformOperationId": jsii.String("transformOperationId"),
//
//   					// the properties below are optional
//   					"columnIdMappings": []interface{}{
//   						&DataSetColumnIdMappingProperty{
//   							"sourceColumnId": jsii.String("sourceColumnId"),
//   							"targetColumnId": jsii.String("targetColumnId"),
//   						},
//   					},
//   				},
//   			},
//   			"createColumnsStep": &CreateColumnsOperationProperty{
//   				"columns": []interface{}{
//   					&CalculatedColumnProperty{
//   						"columnId": jsii.String("columnId"),
//   						"columnName": jsii.String("columnName"),
//   						"expression": jsii.String("expression"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				"alias": jsii.String("alias"),
//   				"source": &TransformOperationSourceProperty{
//   					"transformOperationId": jsii.String("transformOperationId"),
//
//   					// the properties below are optional
//   					"columnIdMappings": []interface{}{
//   						&DataSetColumnIdMappingProperty{
//   							"sourceColumnId": jsii.String("sourceColumnId"),
//   							"targetColumnId": jsii.String("targetColumnId"),
//   						},
//   					},
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
//
//   								// the properties below are optional
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
//
//   								// the properties below are optional
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
//
//   								// the properties below are optional
//   								"value": &DataSetStringFilterValueProperty{
//   									"staticValue": jsii.String("staticValue"),
//   								},
//   							},
//   							"listFilterCondition": &DataSetStringListFilterConditionProperty{
//   								"operator": jsii.String("operator"),
//
//   								// the properties below are optional
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
//   					"transformOperationId": jsii.String("transformOperationId"),
//
//   					// the properties below are optional
//   					"columnIdMappings": []interface{}{
//   						&DataSetColumnIdMappingProperty{
//   							"sourceColumnId": jsii.String("sourceColumnId"),
//   							"targetColumnId": jsii.String("targetColumnId"),
//   						},
//   					},
//   				},
//   			},
//   			"importTableStep": &ImportTableOperationProperty{
//   				"alias": jsii.String("alias"),
//   				"source": &ImportTableOperationSourceProperty{
//   					"sourceTableId": jsii.String("sourceTableId"),
//
//   					// the properties below are optional
//   					"columnIdMappings": []interface{}{
//   						&DataSetColumnIdMappingProperty{
//   							"sourceColumnId": jsii.String("sourceColumnId"),
//   							"targetColumnId": jsii.String("targetColumnId"),
//   						},
//   					},
//   				},
//   			},
//   			"joinStep": &JoinOperationProperty{
//   				"alias": jsii.String("alias"),
//   				"leftOperand": &TransformOperationSourceProperty{
//   					"transformOperationId": jsii.String("transformOperationId"),
//
//   					// the properties below are optional
//   					"columnIdMappings": []interface{}{
//   						&DataSetColumnIdMappingProperty{
//   							"sourceColumnId": jsii.String("sourceColumnId"),
//   							"targetColumnId": jsii.String("targetColumnId"),
//   						},
//   					},
//   				},
//   				"onClause": jsii.String("onClause"),
//   				"rightOperand": &TransformOperationSourceProperty{
//   					"transformOperationId": jsii.String("transformOperationId"),
//
//   					// the properties below are optional
//   					"columnIdMappings": []interface{}{
//   						&DataSetColumnIdMappingProperty{
//   							"sourceColumnId": jsii.String("sourceColumnId"),
//   							"targetColumnId": jsii.String("targetColumnId"),
//   						},
//   					},
//   				},
//   				"type": jsii.String("type"),
//
//   				// the properties below are optional
//   				"leftOperandProperties": &JoinOperandPropertiesProperty{
//   					"outputColumnNameOverrides": []interface{}{
//   						&OutputColumnNameOverrideProperty{
//   							"outputColumnName": jsii.String("outputColumnName"),
//
//   							// the properties below are optional
//   							"sourceColumnName": jsii.String("sourceColumnName"),
//   						},
//   					},
//   				},
//   				"rightOperandProperties": &JoinOperandPropertiesProperty{
//   					"outputColumnNameOverrides": []interface{}{
//   						&OutputColumnNameOverrideProperty{
//   							"outputColumnName": jsii.String("outputColumnName"),
//
//   							// the properties below are optional
//   							"sourceColumnName": jsii.String("sourceColumnName"),
//   						},
//   					},
//   				},
//   			},
//   			"pivotStep": &PivotOperationProperty{
//   				"alias": jsii.String("alias"),
//   				"pivotConfiguration": &PivotConfigurationProperty{
//   					"pivotedLabels": []interface{}{
//   						&PivotedLabelProperty{
//   							"labelName": jsii.String("labelName"),
//   							"newColumnId": jsii.String("newColumnId"),
//   							"newColumnName": jsii.String("newColumnName"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					"labelColumnName": jsii.String("labelColumnName"),
//   				},
//   				"source": &TransformOperationSourceProperty{
//   					"transformOperationId": jsii.String("transformOperationId"),
//
//   					// the properties below are optional
//   					"columnIdMappings": []interface{}{
//   						&DataSetColumnIdMappingProperty{
//   							"sourceColumnId": jsii.String("sourceColumnId"),
//   							"targetColumnId": jsii.String("targetColumnId"),
//   						},
//   					},
//   				},
//   				"valueColumnConfiguration": &ValueColumnConfigurationProperty{
//   					"aggregationFunction": &DataPrepAggregationFunctionProperty{
//   						"listAggregation": &DataPrepListAggregationFunctionProperty{
//   							"distinct": jsii.Boolean(false),
//   							"separator": jsii.String("separator"),
//
//   							// the properties below are optional
//   							"inputColumnName": jsii.String("inputColumnName"),
//   						},
//   						"percentileAggregation": &DataPrepPercentileAggregationFunctionProperty{
//   							"percentileValue": jsii.Number(123),
//
//   							// the properties below are optional
//   							"inputColumnName": jsii.String("inputColumnName"),
//   						},
//   						"simpleAggregation": &DataPrepSimpleAggregationFunctionProperty{
//   							"functionType": jsii.String("functionType"),
//
//   							// the properties below are optional
//   							"inputColumnName": jsii.String("inputColumnName"),
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				"groupByColumnNames": []*string{
//   					jsii.String("groupByColumnNames"),
//   				},
//   			},
//   			"projectStep": &ProjectOperationProperty{
//   				"alias": jsii.String("alias"),
//   				"projectedColumns": []*string{
//   					jsii.String("projectedColumns"),
//   				},
//   				"source": &TransformOperationSourceProperty{
//   					"transformOperationId": jsii.String("transformOperationId"),
//
//   					// the properties below are optional
//   					"columnIdMappings": []interface{}{
//   						&DataSetColumnIdMappingProperty{
//   							"sourceColumnId": jsii.String("sourceColumnId"),
//   							"targetColumnId": jsii.String("targetColumnId"),
//   						},
//   					},
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
//   					"transformOperationId": jsii.String("transformOperationId"),
//
//   					// the properties below are optional
//   					"columnIdMappings": []interface{}{
//   						&DataSetColumnIdMappingProperty{
//   							"sourceColumnId": jsii.String("sourceColumnId"),
//   							"targetColumnId": jsii.String("targetColumnId"),
//   						},
//   					},
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
//   					"transformOperationId": jsii.String("transformOperationId"),
//
//   					// the properties below are optional
//   					"columnIdMappings": []interface{}{
//   						&DataSetColumnIdMappingProperty{
//   							"sourceColumnId": jsii.String("sourceColumnId"),
//   							"targetColumnId": jsii.String("targetColumnId"),
//   						},
//   					},
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
type CfnDataSet_DataPrepConfigurationProperty struct {
	// A map of destination tables that receive the final prepared data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-dataprepconfiguration.html#cfn-quicksight-dataset-dataprepconfiguration-destinationtablemap
	//
	DestinationTableMap interface{} `field:"required" json:"destinationTableMap" yaml:"destinationTableMap"`
	// A map of source tables that provide information about underlying sources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-dataprepconfiguration.html#cfn-quicksight-dataset-dataprepconfiguration-sourcetablemap
	//
	SourceTableMap interface{} `field:"required" json:"sourceTableMap" yaml:"sourceTableMap"`
	// A map of transformation steps that process the data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-dataprepconfiguration.html#cfn-quicksight-dataset-dataprepconfiguration-transformstepmap
	//
	TransformStepMap interface{} `field:"required" json:"transformStepMap" yaml:"transformStepMap"`
}

