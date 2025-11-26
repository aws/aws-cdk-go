package previewawsquicksightmixins


// A filter condition that compares string values using operators like `EQUALS` , `CONTAINS` , or `STARTS_WITH` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSetStringComparisonFilterConditionProperty := &DataSetStringComparisonFilterConditionProperty{
//   	Operator: jsii.String("operator"),
//   	Value: &DataSetStringFilterValueProperty{
//   		StaticValue: jsii.String("staticValue"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetstringcomparisonfiltercondition.html
//
type CfnDataSetPropsMixin_DataSetStringComparisonFilterConditionProperty struct {
	// The comparison operator to use, such as `EQUALS` , `CONTAINS` , `STARTS_WITH` , `ENDS_WITH` , or their negations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetstringcomparisonfiltercondition.html#cfn-quicksight-dataset-datasetstringcomparisonfiltercondition-operator
	//
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// The string value to compare against.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetstringcomparisonfiltercondition.html#cfn-quicksight-dataset-datasetstringcomparisonfiltercondition-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

