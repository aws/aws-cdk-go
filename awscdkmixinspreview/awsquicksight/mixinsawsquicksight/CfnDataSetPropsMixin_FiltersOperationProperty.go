package mixinsawsquicksight


// A transform operation that applies one or more filter conditions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   filtersOperationProperty := &FiltersOperationProperty{
//   	Alias: jsii.String("alias"),
//   	FilterOperations: []interface{}{
//   		&FilterOperationProperty{
//   			ConditionExpression: jsii.String("conditionExpression"),
//   			DateFilterCondition: &DataSetDateFilterConditionProperty{
//   				ColumnName: jsii.String("columnName"),
//   				ComparisonFilterCondition: &DataSetDateComparisonFilterConditionProperty{
//   					Operator: jsii.String("operator"),
//   					Value: &DataSetDateFilterValueProperty{
//   						StaticValue: jsii.String("staticValue"),
//   					},
//   				},
//   				RangeFilterCondition: &DataSetDateRangeFilterConditionProperty{
//   					IncludeMaximum: jsii.Boolean(false),
//   					IncludeMinimum: jsii.Boolean(false),
//   					RangeMaximum: &DataSetDateFilterValueProperty{
//   						StaticValue: jsii.String("staticValue"),
//   					},
//   					RangeMinimum: &DataSetDateFilterValueProperty{
//   						StaticValue: jsii.String("staticValue"),
//   					},
//   				},
//   			},
//   			NumericFilterCondition: &DataSetNumericFilterConditionProperty{
//   				ColumnName: jsii.String("columnName"),
//   				ComparisonFilterCondition: &DataSetNumericComparisonFilterConditionProperty{
//   					Operator: jsii.String("operator"),
//   					Value: &DataSetNumericFilterValueProperty{
//   						StaticValue: jsii.Number(123),
//   					},
//   				},
//   				RangeFilterCondition: &DataSetNumericRangeFilterConditionProperty{
//   					IncludeMaximum: jsii.Boolean(false),
//   					IncludeMinimum: jsii.Boolean(false),
//   					RangeMaximum: &DataSetNumericFilterValueProperty{
//   						StaticValue: jsii.Number(123),
//   					},
//   					RangeMinimum: &DataSetNumericFilterValueProperty{
//   						StaticValue: jsii.Number(123),
//   					},
//   				},
//   			},
//   			StringFilterCondition: &DataSetStringFilterConditionProperty{
//   				ColumnName: jsii.String("columnName"),
//   				ComparisonFilterCondition: &DataSetStringComparisonFilterConditionProperty{
//   					Operator: jsii.String("operator"),
//   					Value: &DataSetStringFilterValueProperty{
//   						StaticValue: jsii.String("staticValue"),
//   					},
//   				},
//   				ListFilterCondition: &DataSetStringListFilterConditionProperty{
//   					Operator: jsii.String("operator"),
//   					Values: &DataSetStringListFilterValueProperty{
//   						StaticValues: []*string{
//   							jsii.String("staticValues"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Source: &TransformOperationSourceProperty{
//   		ColumnIdMappings: []interface{}{
//   			&DataSetColumnIdMappingProperty{
//   				SourceColumnId: jsii.String("sourceColumnId"),
//   				TargetColumnId: jsii.String("targetColumnId"),
//   			},
//   		},
//   		TransformOperationId: jsii.String("transformOperationId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-filtersoperation.html
//
type CfnDataSetPropsMixin_FiltersOperationProperty struct {
	// Alias for this operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-filtersoperation.html#cfn-quicksight-dataset-filtersoperation-alias
	//
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The list of filter operations to apply.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-filtersoperation.html#cfn-quicksight-dataset-filtersoperation-filteroperations
	//
	FilterOperations interface{} `field:"optional" json:"filterOperations" yaml:"filterOperations"`
	// The source transform operation that provides input data for filtering.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-filtersoperation.html#cfn-quicksight-dataset-filtersoperation-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
}

