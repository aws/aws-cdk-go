package awsquicksight


// A transform operation that overrides the dataset parameter values that are defined in another dataset.
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
type CfnDataSet_OverrideDatasetParameterOperationProperty struct {
	// `CfnDataSet.OverrideDatasetParameterOperationProperty.ParameterName`.
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// The new default values for the parameter.
	NewDefaultValues interface{} `field:"optional" json:"newDefaultValues" yaml:"newDefaultValues"`
	// The new name for the parameter.
	NewParameterName *string `field:"optional" json:"newParameterName" yaml:"newParameterName"`
}

