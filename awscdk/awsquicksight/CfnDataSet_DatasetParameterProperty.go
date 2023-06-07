package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datasetParameterProperty := &DatasetParameterProperty{
//   	DateTimeDatasetParameter: &DateTimeDatasetParameterProperty{
//   		Id: jsii.String("id"),
//   		Name: jsii.String("name"),
//   		ValueType: jsii.String("valueType"),
//
//   		// the properties below are optional
//   		DefaultValues: &DateTimeDatasetParameterDefaultValuesProperty{
//   			StaticValues: []*string{
//   				jsii.String("staticValues"),
//   			},
//   		},
//   		TimeGranularity: jsii.String("timeGranularity"),
//   	},
//   	DecimalDatasetParameter: &DecimalDatasetParameterProperty{
//   		Id: jsii.String("id"),
//   		Name: jsii.String("name"),
//   		ValueType: jsii.String("valueType"),
//
//   		// the properties below are optional
//   		DefaultValues: &DecimalDatasetParameterDefaultValuesProperty{
//   			StaticValues: []interface{}{
//   				jsii.Number(123),
//   			},
//   		},
//   	},
//   	IntegerDatasetParameter: &IntegerDatasetParameterProperty{
//   		Id: jsii.String("id"),
//   		Name: jsii.String("name"),
//   		ValueType: jsii.String("valueType"),
//
//   		// the properties below are optional
//   		DefaultValues: &IntegerDatasetParameterDefaultValuesProperty{
//   			StaticValues: []interface{}{
//   				jsii.Number(123),
//   			},
//   		},
//   	},
//   	StringDatasetParameter: &StringDatasetParameterProperty{
//   		Id: jsii.String("id"),
//   		Name: jsii.String("name"),
//   		ValueType: jsii.String("valueType"),
//
//   		// the properties below are optional
//   		DefaultValues: &StringDatasetParameterDefaultValuesProperty{
//   			StaticValues: []*string{
//   				jsii.String("staticValues"),
//   			},
//   		},
//   	},
//   }
//
type CfnDataSet_DatasetParameterProperty struct {
	// `CfnDataSet.DatasetParameterProperty.DateTimeDatasetParameter`.
	DateTimeDatasetParameter interface{} `field:"optional" json:"dateTimeDatasetParameter" yaml:"dateTimeDatasetParameter"`
	// `CfnDataSet.DatasetParameterProperty.DecimalDatasetParameter`.
	DecimalDatasetParameter interface{} `field:"optional" json:"decimalDatasetParameter" yaml:"decimalDatasetParameter"`
	// `CfnDataSet.DatasetParameterProperty.IntegerDatasetParameter`.
	IntegerDatasetParameter interface{} `field:"optional" json:"integerDatasetParameter" yaml:"integerDatasetParameter"`
	// `CfnDataSet.DatasetParameterProperty.StringDatasetParameter`.
	StringDatasetParameter interface{} `field:"optional" json:"stringDatasetParameter" yaml:"stringDatasetParameter"`
}

