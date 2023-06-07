package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dateTimeDatasetParameterProperty := &DateTimeDatasetParameterProperty{
//   	Id: jsii.String("id"),
//   	Name: jsii.String("name"),
//   	ValueType: jsii.String("valueType"),
//
//   	// the properties below are optional
//   	DefaultValues: &DateTimeDatasetParameterDefaultValuesProperty{
//   		StaticValues: []*string{
//   			jsii.String("staticValues"),
//   		},
//   	},
//   	TimeGranularity: jsii.String("timeGranularity"),
//   }
//
type CfnDataSet_DateTimeDatasetParameterProperty struct {
	// `CfnDataSet.DateTimeDatasetParameterProperty.Id`.
	Id *string `field:"required" json:"id" yaml:"id"`
	// `CfnDataSet.DateTimeDatasetParameterProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnDataSet.DateTimeDatasetParameterProperty.ValueType`.
	ValueType *string `field:"required" json:"valueType" yaml:"valueType"`
	// `CfnDataSet.DateTimeDatasetParameterProperty.DefaultValues`.
	DefaultValues interface{} `field:"optional" json:"defaultValues" yaml:"defaultValues"`
	// `CfnDataSet.DateTimeDatasetParameterProperty.TimeGranularity`.
	TimeGranularity *string `field:"optional" json:"timeGranularity" yaml:"timeGranularity"`
}

