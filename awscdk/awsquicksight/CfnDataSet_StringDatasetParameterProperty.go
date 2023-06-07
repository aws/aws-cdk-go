package awsquicksight


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
type CfnDataSet_StringDatasetParameterProperty struct {
	// `CfnDataSet.StringDatasetParameterProperty.Id`.
	Id *string `field:"required" json:"id" yaml:"id"`
	// `CfnDataSet.StringDatasetParameterProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnDataSet.StringDatasetParameterProperty.ValueType`.
	ValueType *string `field:"required" json:"valueType" yaml:"valueType"`
	// `CfnDataSet.StringDatasetParameterProperty.DefaultValues`.
	DefaultValues interface{} `field:"optional" json:"defaultValues" yaml:"defaultValues"`
}

