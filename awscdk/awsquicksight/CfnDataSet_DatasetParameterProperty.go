package awsquicksight


// <p>A parameter created in the dataset that could be of any one data type such as string, integer, decimal or datetime.</p>.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetparameter.html
//
type CfnDataSet_DatasetParameterProperty struct {
	// <p>A parameter created in the dataset of date time data type.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetparameter.html#cfn-quicksight-dataset-datasetparameter-datetimedatasetparameter
	//
	DateTimeDatasetParameter interface{} `field:"optional" json:"dateTimeDatasetParameter" yaml:"dateTimeDatasetParameter"`
	// <p>A parameter created in the dataset of decimal data type.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetparameter.html#cfn-quicksight-dataset-datasetparameter-decimaldatasetparameter
	//
	DecimalDatasetParameter interface{} `field:"optional" json:"decimalDatasetParameter" yaml:"decimalDatasetParameter"`
	// <p>A parameter created in the dataset of integer data type.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetparameter.html#cfn-quicksight-dataset-datasetparameter-integerdatasetparameter
	//
	IntegerDatasetParameter interface{} `field:"optional" json:"integerDatasetParameter" yaml:"integerDatasetParameter"`
	// <p>A parameter created in the dataset of string data type.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetparameter.html#cfn-quicksight-dataset-datasetparameter-stringdatasetparameter
	//
	StringDatasetParameter interface{} `field:"optional" json:"stringDatasetParameter" yaml:"stringDatasetParameter"`
}

