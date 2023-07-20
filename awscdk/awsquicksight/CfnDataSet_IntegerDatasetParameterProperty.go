package awsquicksight


// <p>A parameter created in the dataset of integer data type.</p>.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integerDatasetParameterProperty := &IntegerDatasetParameterProperty{
//   	Id: jsii.String("id"),
//   	Name: jsii.String("name"),
//   	ValueType: jsii.String("valueType"),
//
//   	// the properties below are optional
//   	DefaultValues: &IntegerDatasetParameterDefaultValuesProperty{
//   		StaticValues: []interface{}{
//   			jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-integerdatasetparameter.html
//
type CfnDataSet_IntegerDatasetParameterProperty struct {
	// <p>Identifier of the parameter created in the dataset.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-integerdatasetparameter.html#cfn-quicksight-dataset-integerdatasetparameter-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// <p>Name of the parameter created in the dataset.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-integerdatasetparameter.html#cfn-quicksight-dataset-integerdatasetparameter-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// <p>Every parameter value could be either a single value or multi value which helps to validate before evaluation.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-integerdatasetparameter.html#cfn-quicksight-dataset-integerdatasetparameter-valuetype
	//
	ValueType *string `field:"required" json:"valueType" yaml:"valueType"`
	// <p>List of default values defined for a given integer dataset parameter type.
	//
	// Currently only static values are supported.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-integerdatasetparameter.html#cfn-quicksight-dataset-integerdatasetparameter-defaultvalues
	//
	DefaultValues interface{} `field:"optional" json:"defaultValues" yaml:"defaultValues"`
}

