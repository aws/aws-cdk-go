package awsquicksight


// A filter condition that filters date values within a specified range.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSetDateRangeFilterConditionProperty := &DataSetDateRangeFilterConditionProperty{
//   	IncludeMaximum: jsii.Boolean(false),
//   	IncludeMinimum: jsii.Boolean(false),
//   	RangeMaximum: &DataSetDateFilterValueProperty{
//   		StaticValue: jsii.String("staticValue"),
//   	},
//   	RangeMinimum: &DataSetDateFilterValueProperty{
//   		StaticValue: jsii.String("staticValue"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetdaterangefiltercondition.html
//
type CfnDataSet_DataSetDateRangeFilterConditionProperty struct {
	// Whether to include the maximum value in the filter range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetdaterangefiltercondition.html#cfn-quicksight-dataset-datasetdaterangefiltercondition-includemaximum
	//
	IncludeMaximum interface{} `field:"optional" json:"includeMaximum" yaml:"includeMaximum"`
	// Whether to include the minimum value in the filter range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetdaterangefiltercondition.html#cfn-quicksight-dataset-datasetdaterangefiltercondition-includeminimum
	//
	IncludeMinimum interface{} `field:"optional" json:"includeMinimum" yaml:"includeMinimum"`
	// The maximum date value for the range filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetdaterangefiltercondition.html#cfn-quicksight-dataset-datasetdaterangefiltercondition-rangemaximum
	//
	RangeMaximum interface{} `field:"optional" json:"rangeMaximum" yaml:"rangeMaximum"`
	// The minimum date value for the range filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetdaterangefiltercondition.html#cfn-quicksight-dataset-datasetdaterangefiltercondition-rangeminimum
	//
	RangeMinimum interface{} `field:"optional" json:"rangeMinimum" yaml:"rangeMinimum"`
}

