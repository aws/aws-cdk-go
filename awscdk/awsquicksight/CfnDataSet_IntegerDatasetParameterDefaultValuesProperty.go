package awsquicksight


// <p>List of default values defined for a given integer dataset parameter type.
//
// Currently only static values are supported.</p>
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integerDatasetParameterDefaultValuesProperty := &IntegerDatasetParameterDefaultValuesProperty{
//   	StaticValues: []interface{}{
//   		jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-integerdatasetparameterdefaultvalues.html
//
type CfnDataSet_IntegerDatasetParameterDefaultValuesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-integerdatasetparameterdefaultvalues.html#cfn-quicksight-dataset-integerdatasetparameterdefaultvalues-staticvalues
	//
	StaticValues interface{} `field:"optional" json:"staticValues" yaml:"staticValues"`
}

