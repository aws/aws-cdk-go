package previewawsquicksightmixins


// A list of default values for a given decimal parameter.
//
// This structure only accepts static values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   decimalDatasetParameterDefaultValuesProperty := &DecimalDatasetParameterDefaultValuesProperty{
//   	StaticValues: []interface{}{
//   		jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-decimaldatasetparameterdefaultvalues.html
//
type CfnDataSetPropsMixin_DecimalDatasetParameterDefaultValuesProperty struct {
	// A list of static default values for a given decimal parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-decimaldatasetparameterdefaultvalues.html#cfn-quicksight-dataset-decimaldatasetparameterdefaultvalues-staticvalues
	//
	StaticValues interface{} `field:"optional" json:"staticValues" yaml:"staticValues"`
}

