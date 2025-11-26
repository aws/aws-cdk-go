package previewawsquicksightmixins


// A list of default values for a given integer parameter.
//
// This structure only accepts static values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   integerDatasetParameterDefaultValuesProperty := &IntegerDatasetParameterDefaultValuesProperty{
//   	StaticValues: []interface{}{
//   		jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-integerdatasetparameterdefaultvalues.html
//
type CfnDataSetPropsMixin_IntegerDatasetParameterDefaultValuesProperty struct {
	// A list of static default values for a given integer parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-integerdatasetparameterdefaultvalues.html#cfn-quicksight-dataset-integerdatasetparameterdefaultvalues-staticvalues
	//
	StaticValues interface{} `field:"optional" json:"staticValues" yaml:"staticValues"`
}

