package previewawsquicksightmixins


// A filter condition for date columns, supporting both comparison and range-based filtering.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSetDateFilterConditionProperty := &DataSetDateFilterConditionProperty{
//   	ColumnName: jsii.String("columnName"),
//   	ComparisonFilterCondition: &DataSetDateComparisonFilterConditionProperty{
//   		Operator: jsii.String("operator"),
//   		Value: &DataSetDateFilterValueProperty{
//   			StaticValue: jsii.String("staticValue"),
//   		},
//   	},
//   	RangeFilterCondition: &DataSetDateRangeFilterConditionProperty{
//   		IncludeMaximum: jsii.Boolean(false),
//   		IncludeMinimum: jsii.Boolean(false),
//   		RangeMaximum: &DataSetDateFilterValueProperty{
//   			StaticValue: jsii.String("staticValue"),
//   		},
//   		RangeMinimum: &DataSetDateFilterValueProperty{
//   			StaticValue: jsii.String("staticValue"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetdatefiltercondition.html
//
type CfnDataSetPropsMixin_DataSetDateFilterConditionProperty struct {
	// The name of the date column to filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetdatefiltercondition.html#cfn-quicksight-dataset-datasetdatefiltercondition-columnname
	//
	ColumnName *string `field:"optional" json:"columnName" yaml:"columnName"`
	// A comparison-based filter condition for the date column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetdatefiltercondition.html#cfn-quicksight-dataset-datasetdatefiltercondition-comparisonfiltercondition
	//
	ComparisonFilterCondition interface{} `field:"optional" json:"comparisonFilterCondition" yaml:"comparisonFilterCondition"`
	// A range-based filter condition for the date column, filtering values between minimum and maximum dates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetdatefiltercondition.html#cfn-quicksight-dataset-datasetdatefiltercondition-rangefiltercondition
	//
	RangeFilterCondition interface{} `field:"optional" json:"rangeFilterCondition" yaml:"rangeFilterCondition"`
}

