package previewawsquicksightmixins


// A filter condition that includes or excludes string values from a specified list.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSetStringListFilterConditionProperty := &DataSetStringListFilterConditionProperty{
//   	Operator: jsii.String("operator"),
//   	Values: &DataSetStringListFilterValueProperty{
//   		StaticValues: []*string{
//   			jsii.String("staticValues"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetstringlistfiltercondition.html
//
type CfnDataSetPropsMixin_DataSetStringListFilterConditionProperty struct {
	// The list operator to use, either `INCLUDE` to match values in the list or `EXCLUDE` to filter out values in the list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetstringlistfiltercondition.html#cfn-quicksight-dataset-datasetstringlistfiltercondition-operator
	//
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// The list of string values to include or exclude in the filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetstringlistfiltercondition.html#cfn-quicksight-dataset-datasetstringlistfiltercondition-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

