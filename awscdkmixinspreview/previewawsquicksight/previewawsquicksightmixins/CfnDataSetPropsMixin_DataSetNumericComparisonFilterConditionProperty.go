package previewawsquicksightmixins


// A filter condition that compares numeric values using operators like `EQUALS` , `GREATER_THAN` , or `LESS_THAN` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSetNumericComparisonFilterConditionProperty := &DataSetNumericComparisonFilterConditionProperty{
//   	Operator: jsii.String("operator"),
//   	Value: &DataSetNumericFilterValueProperty{
//   		StaticValue: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetnumericcomparisonfiltercondition.html
//
type CfnDataSetPropsMixin_DataSetNumericComparisonFilterConditionProperty struct {
	// The comparison operator to use, such as `EQUALS` , `GREATER_THAN` , `LESS_THAN` , or their variants.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetnumericcomparisonfiltercondition.html#cfn-quicksight-dataset-datasetnumericcomparisonfiltercondition-operator
	//
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// The numeric value to compare against.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetnumericcomparisonfiltercondition.html#cfn-quicksight-dataset-datasetnumericcomparisonfiltercondition-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

