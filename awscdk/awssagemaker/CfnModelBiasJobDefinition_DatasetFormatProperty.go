package awssagemaker


// The dataset format of the data to monitor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datasetFormatProperty := &DatasetFormatProperty{
//   	Csv: &CsvProperty{
//   		Header: jsii.Boolean(false),
//   	},
//   	Json: &JsonProperty{
//   		Line: jsii.Boolean(false),
//   	},
//   	Parquet: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-datasetformat.html
//
type CfnModelBiasJobDefinition_DatasetFormatProperty struct {
	// The CSV format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-datasetformat.html#cfn-sagemaker-modelbiasjobdefinition-datasetformat-csv
	//
	Csv interface{} `field:"optional" json:"csv" yaml:"csv"`
	// The Json format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-datasetformat.html#cfn-sagemaker-modelbiasjobdefinition-datasetformat-json
	//
	Json interface{} `field:"optional" json:"json" yaml:"json"`
	// A flag indicate if the dataset format is Parquet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-datasetformat.html#cfn-sagemaker-modelbiasjobdefinition-datasetformat-parquet
	//
	Parquet interface{} `field:"optional" json:"parquet" yaml:"parquet"`
}

