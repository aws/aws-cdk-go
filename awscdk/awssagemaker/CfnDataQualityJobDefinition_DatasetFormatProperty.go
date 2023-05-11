package awssagemaker


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
type CfnDataQualityJobDefinition_DatasetFormatProperty struct {
	// `CfnDataQualityJobDefinition.DatasetFormatProperty.Csv`.
	Csv interface{} `field:"optional" json:"csv" yaml:"csv"`
	// `CfnDataQualityJobDefinition.DatasetFormatProperty.Json`.
	Json interface{} `field:"optional" json:"json" yaml:"json"`
	// `CfnDataQualityJobDefinition.DatasetFormatProperty.Parquet`.
	Parquet interface{} `field:"optional" json:"parquet" yaml:"parquet"`
}

