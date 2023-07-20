package awsquicksight


// <p>A parameter created in the dataset of string data type.</p>.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stringDatasetParameterProperty := &StringDatasetParameterProperty{
//   	Id: jsii.String("id"),
//   	Name: jsii.String("name"),
//   	ValueType: jsii.String("valueType"),
//
//   	// the properties below are optional
//   	DefaultValues: &StringDatasetParameterDefaultValuesProperty{
//   		StaticValues: []*string{
//   			jsii.String("staticValues"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-stringdatasetparameter.html
//
type CfnDataSet_StringDatasetParameterProperty struct {
	// <p>Identifier of the parameter created in the dataset.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-stringdatasetparameter.html#cfn-quicksight-dataset-stringdatasetparameter-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// <p>Name of the parameter created in the dataset.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-stringdatasetparameter.html#cfn-quicksight-dataset-stringdatasetparameter-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// <p>Every parameter value could be either a single value or multi value which helps to validate before evaluation.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-stringdatasetparameter.html#cfn-quicksight-dataset-stringdatasetparameter-valuetype
	//
	ValueType *string `field:"required" json:"valueType" yaml:"valueType"`
	// <p>List of default values defined for a given string dataset parameter type.
	//
	// Currently only static values are supported.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-stringdatasetparameter.html#cfn-quicksight-dataset-stringdatasetparameter-defaultvalues
	//
	DefaultValues interface{} `field:"optional" json:"defaultValues" yaml:"defaultValues"`
}

