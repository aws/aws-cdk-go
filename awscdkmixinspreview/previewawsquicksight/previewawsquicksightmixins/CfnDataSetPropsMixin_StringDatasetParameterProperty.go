package previewawsquicksightmixins


// A string parameter that is created in the dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   stringDatasetParameterProperty := &StringDatasetParameterProperty{
//   	DefaultValues: &StringDatasetParameterDefaultValuesProperty{
//   		StaticValues: []*string{
//   			jsii.String("staticValues"),
//   		},
//   	},
//   	Id: jsii.String("id"),
//   	Name: jsii.String("name"),
//   	ValueType: jsii.String("valueType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-stringdatasetparameter.html
//
type CfnDataSetPropsMixin_StringDatasetParameterProperty struct {
	// A list of default values for a given string dataset parameter type.
	//
	// This structure only accepts static values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-stringdatasetparameter.html#cfn-quicksight-dataset-stringdatasetparameter-defaultvalues
	//
	DefaultValues interface{} `field:"optional" json:"defaultValues" yaml:"defaultValues"`
	// An identifier for the string parameter that is created in the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-stringdatasetparameter.html#cfn-quicksight-dataset-stringdatasetparameter-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the string parameter that is created in the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-stringdatasetparameter.html#cfn-quicksight-dataset-stringdatasetparameter-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value type of the dataset parameter.
	//
	// Valid values are `single value` or `multi value` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-stringdatasetparameter.html#cfn-quicksight-dataset-stringdatasetparameter-valuetype
	//
	ValueType *string `field:"optional" json:"valueType" yaml:"valueType"`
}

