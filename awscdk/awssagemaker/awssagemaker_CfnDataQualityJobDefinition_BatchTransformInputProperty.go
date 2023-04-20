package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var json interface{}
//
//   batchTransformInputProperty := &BatchTransformInputProperty{
//   	DataCapturedDestinationS3Uri: jsii.String("dataCapturedDestinationS3Uri"),
//   	DatasetFormat: &DatasetFormatProperty{
//   		Csv: &CsvProperty{
//   			Header: jsii.Boolean(false),
//   		},
//   		Json: json,
//   		Parquet: jsii.Boolean(false),
//   	},
//   	LocalPath: jsii.String("localPath"),
//
//   	// the properties below are optional
//   	S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   	S3InputMode: jsii.String("s3InputMode"),
//   }
//
type CfnDataQualityJobDefinition_BatchTransformInputProperty struct {
	// `CfnDataQualityJobDefinition.BatchTransformInputProperty.DataCapturedDestinationS3Uri`.
	DataCapturedDestinationS3Uri *string `field:"required" json:"dataCapturedDestinationS3Uri" yaml:"dataCapturedDestinationS3Uri"`
	// `CfnDataQualityJobDefinition.BatchTransformInputProperty.DatasetFormat`.
	DatasetFormat interface{} `field:"required" json:"datasetFormat" yaml:"datasetFormat"`
	// `CfnDataQualityJobDefinition.BatchTransformInputProperty.LocalPath`.
	LocalPath *string `field:"required" json:"localPath" yaml:"localPath"`
	// `CfnDataQualityJobDefinition.BatchTransformInputProperty.S3DataDistributionType`.
	S3DataDistributionType *string `field:"optional" json:"s3DataDistributionType" yaml:"s3DataDistributionType"`
	// `CfnDataQualityJobDefinition.BatchTransformInputProperty.S3InputMode`.
	S3InputMode *string `field:"optional" json:"s3InputMode" yaml:"s3InputMode"`
}

