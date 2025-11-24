package mixinsawsquicksight


// A filter condition for numeric columns, supporting both comparison and range-based filtering.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSetNumericFilterConditionProperty := &DataSetNumericFilterConditionProperty{
//   	ColumnName: jsii.String("columnName"),
//   	ComparisonFilterCondition: &DataSetNumericComparisonFilterConditionProperty{
//   		Operator: jsii.String("operator"),
//   		Value: &DataSetNumericFilterValueProperty{
//   			StaticValue: jsii.Number(123),
//   		},
//   	},
//   	RangeFilterCondition: &DataSetNumericRangeFilterConditionProperty{
//   		IncludeMaximum: jsii.Boolean(false),
//   		IncludeMinimum: jsii.Boolean(false),
//   		RangeMaximum: &DataSetNumericFilterValueProperty{
//   			StaticValue: jsii.Number(123),
//   		},
//   		RangeMinimum: &DataSetNumericFilterValueProperty{
//   			StaticValue: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetnumericfiltercondition.html
//
type CfnDataSetPropsMixin_DataSetNumericFilterConditionProperty struct {
	// The name of the numeric column to filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetnumericfiltercondition.html#cfn-quicksight-dataset-datasetnumericfiltercondition-columnname
	//
	ColumnName *string `field:"optional" json:"columnName" yaml:"columnName"`
	// A comparison-based filter condition for the numeric column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetnumericfiltercondition.html#cfn-quicksight-dataset-datasetnumericfiltercondition-comparisonfiltercondition
	//
	ComparisonFilterCondition interface{} `field:"optional" json:"comparisonFilterCondition" yaml:"comparisonFilterCondition"`
	// A range-based filter condition for the numeric column, filtering values between minimum and maximum numbers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetnumericfiltercondition.html#cfn-quicksight-dataset-datasetnumericfiltercondition-rangefiltercondition
	//
	RangeFilterCondition interface{} `field:"optional" json:"rangeFilterCondition" yaml:"rangeFilterCondition"`
}

