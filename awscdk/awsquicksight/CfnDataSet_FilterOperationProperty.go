package awsquicksight


// A transform operation that filters rows based on a condition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterOperationProperty := &FilterOperationProperty{
//   	ConditionExpression: jsii.String("conditionExpression"),
//   	DateFilterCondition: &DataSetDateFilterConditionProperty{
//   		ColumnName: jsii.String("columnName"),
//   		ComparisonFilterCondition: &DataSetDateComparisonFilterConditionProperty{
//   			Operator: jsii.String("operator"),
//
//   			// the properties below are optional
//   			Value: &DataSetDateFilterValueProperty{
//   				StaticValue: jsii.String("staticValue"),
//   			},
//   		},
//   		RangeFilterCondition: &DataSetDateRangeFilterConditionProperty{
//   			IncludeMaximum: jsii.Boolean(false),
//   			IncludeMinimum: jsii.Boolean(false),
//   			RangeMaximum: &DataSetDateFilterValueProperty{
//   				StaticValue: jsii.String("staticValue"),
//   			},
//   			RangeMinimum: &DataSetDateFilterValueProperty{
//   				StaticValue: jsii.String("staticValue"),
//   			},
//   		},
//   	},
//   	NumericFilterCondition: &DataSetNumericFilterConditionProperty{
//   		ColumnName: jsii.String("columnName"),
//   		ComparisonFilterCondition: &DataSetNumericComparisonFilterConditionProperty{
//   			Operator: jsii.String("operator"),
//
//   			// the properties below are optional
//   			Value: &DataSetNumericFilterValueProperty{
//   				StaticValue: jsii.Number(123),
//   			},
//   		},
//   		RangeFilterCondition: &DataSetNumericRangeFilterConditionProperty{
//   			IncludeMaximum: jsii.Boolean(false),
//   			IncludeMinimum: jsii.Boolean(false),
//   			RangeMaximum: &DataSetNumericFilterValueProperty{
//   				StaticValue: jsii.Number(123),
//   			},
//   			RangeMinimum: &DataSetNumericFilterValueProperty{
//   				StaticValue: jsii.Number(123),
//   			},
//   		},
//   	},
//   	StringFilterCondition: &DataSetStringFilterConditionProperty{
//   		ColumnName: jsii.String("columnName"),
//   		ComparisonFilterCondition: &DataSetStringComparisonFilterConditionProperty{
//   			Operator: jsii.String("operator"),
//
//   			// the properties below are optional
//   			Value: &DataSetStringFilterValueProperty{
//   				StaticValue: jsii.String("staticValue"),
//   			},
//   		},
//   		ListFilterCondition: &DataSetStringListFilterConditionProperty{
//   			Operator: jsii.String("operator"),
//
//   			// the properties below are optional
//   			Values: &DataSetStringListFilterValueProperty{
//   				StaticValues: []*string{
//   					jsii.String("staticValues"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-filteroperation.html
//
type CfnDataSet_FilterOperationProperty struct {
	// An expression that must evaluate to a Boolean value.
	//
	// Rows for which the expression evaluates to true are kept in the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-filteroperation.html#cfn-quicksight-dataset-filteroperation-conditionexpression
	//
	ConditionExpression *string `field:"optional" json:"conditionExpression" yaml:"conditionExpression"`
	// A date-based filter condition within a filter operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-filteroperation.html#cfn-quicksight-dataset-filteroperation-datefiltercondition
	//
	DateFilterCondition interface{} `field:"optional" json:"dateFilterCondition" yaml:"dateFilterCondition"`
	// A numeric-based filter condition within a filter operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-filteroperation.html#cfn-quicksight-dataset-filteroperation-numericfiltercondition
	//
	NumericFilterCondition interface{} `field:"optional" json:"numericFilterCondition" yaml:"numericFilterCondition"`
	// A string-based filter condition within a filter operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-filteroperation.html#cfn-quicksight-dataset-filteroperation-stringfiltercondition
	//
	StringFilterCondition interface{} `field:"optional" json:"stringFilterCondition" yaml:"stringFilterCondition"`
}

