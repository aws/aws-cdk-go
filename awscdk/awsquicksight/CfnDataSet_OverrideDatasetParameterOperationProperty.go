package awsquicksight


// <p>A transform operation that overrides the dataset parameter values defined in another dataset.</p>.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   overrideDatasetParameterOperationProperty := &OverrideDatasetParameterOperationProperty{
//   	ParameterName: jsii.String("parameterName"),
//
//   	// the properties below are optional
//   	NewDefaultValues: &NewDefaultValuesProperty{
//   		DateTimeStaticValues: []*string{
//   			jsii.String("dateTimeStaticValues"),
//   		},
//   		DecimalStaticValues: []interface{}{
//   			jsii.Number(123),
//   		},
//   		IntegerStaticValues: []interface{}{
//   			jsii.Number(123),
//   		},
//   		StringStaticValues: []*string{
//   			jsii.String("stringStaticValues"),
//   		},
//   	},
//   	NewParameterName: jsii.String("newParameterName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-overridedatasetparameteroperation.html
//
type CfnDataSet_OverrideDatasetParameterOperationProperty struct {
	// <p>Name of the parameter created in the dataset.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-overridedatasetparameteroperation.html#cfn-quicksight-dataset-overridedatasetparameteroperation-parametername
	//
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-overridedatasetparameteroperation.html#cfn-quicksight-dataset-overridedatasetparameteroperation-newdefaultvalues
	//
	NewDefaultValues interface{} `field:"optional" json:"newDefaultValues" yaml:"newDefaultValues"`
	// <p>Name of the parameter created in the dataset.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-overridedatasetparameteroperation.html#cfn-quicksight-dataset-overridedatasetparameteroperation-newparametername
	//
	NewParameterName *string `field:"optional" json:"newParameterName" yaml:"newParameterName"`
}

