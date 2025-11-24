package mixinsawsquicksight


// A filter condition that filters numeric values within a specified range.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSetNumericRangeFilterConditionProperty := &DataSetNumericRangeFilterConditionProperty{
//   	IncludeMaximum: jsii.Boolean(false),
//   	IncludeMinimum: jsii.Boolean(false),
//   	RangeMaximum: &DataSetNumericFilterValueProperty{
//   		StaticValue: jsii.Number(123),
//   	},
//   	RangeMinimum: &DataSetNumericFilterValueProperty{
//   		StaticValue: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetnumericrangefiltercondition.html
//
type CfnDataSetPropsMixin_DataSetNumericRangeFilterConditionProperty struct {
	// Whether to include the maximum value in the filter range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetnumericrangefiltercondition.html#cfn-quicksight-dataset-datasetnumericrangefiltercondition-includemaximum
	//
	IncludeMaximum interface{} `field:"optional" json:"includeMaximum" yaml:"includeMaximum"`
	// Whether to include the minimum value in the filter range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetnumericrangefiltercondition.html#cfn-quicksight-dataset-datasetnumericrangefiltercondition-includeminimum
	//
	IncludeMinimum interface{} `field:"optional" json:"includeMinimum" yaml:"includeMinimum"`
	// The maximum numeric value for the range filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetnumericrangefiltercondition.html#cfn-quicksight-dataset-datasetnumericrangefiltercondition-rangemaximum
	//
	RangeMaximum interface{} `field:"optional" json:"rangeMaximum" yaml:"rangeMaximum"`
	// The minimum numeric value for the range filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetnumericrangefiltercondition.html#cfn-quicksight-dataset-datasetnumericrangefiltercondition-rangeminimum
	//
	RangeMinimum interface{} `field:"optional" json:"rangeMinimum" yaml:"rangeMinimum"`
}

