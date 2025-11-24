package mixinsawsquicksight


// A filter condition that compares date values using operators like `BEFORE` , `AFTER` , or their inclusive variants.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSetDateComparisonFilterConditionProperty := &DataSetDateComparisonFilterConditionProperty{
//   	Operator: jsii.String("operator"),
//   	Value: &DataSetDateFilterValueProperty{
//   		StaticValue: jsii.String("staticValue"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetdatecomparisonfiltercondition.html
//
type CfnDataSetPropsMixin_DataSetDateComparisonFilterConditionProperty struct {
	// The comparison operator to use, such as `BEFORE` , `BEFORE_OR_EQUALS_TO` , `AFTER` , or `AFTER_OR_EQUALS_TO` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetdatecomparisonfiltercondition.html#cfn-quicksight-dataset-datasetdatecomparisonfiltercondition-operator
	//
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// The date value to compare against.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetdatecomparisonfiltercondition.html#cfn-quicksight-dataset-datasetdatecomparisonfiltercondition-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

