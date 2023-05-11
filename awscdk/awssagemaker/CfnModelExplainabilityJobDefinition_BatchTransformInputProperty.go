package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   batchTransformInputProperty := &BatchTransformInputProperty{
//   	DataCapturedDestinationS3Uri: jsii.String("dataCapturedDestinationS3Uri"),
//   	DatasetFormat: &DatasetFormatProperty{
//   		Csv: &CsvProperty{
//   			Header: jsii.Boolean(false),
//   		},
//   		Json: &JsonProperty{
//   			Line: jsii.Boolean(false),
//   		},
//   		Parquet: jsii.Boolean(false),
//   	},
//   	LocalPath: jsii.String("localPath"),
//
//   	// the properties below are optional
//   	FeaturesAttribute: jsii.String("featuresAttribute"),
//   	InferenceAttribute: jsii.String("inferenceAttribute"),
//   	ProbabilityAttribute: jsii.String("probabilityAttribute"),
//   	S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   	S3InputMode: jsii.String("s3InputMode"),
//   }
//
type CfnModelExplainabilityJobDefinition_BatchTransformInputProperty struct {
	// `CfnModelExplainabilityJobDefinition.BatchTransformInputProperty.DataCapturedDestinationS3Uri`.
	DataCapturedDestinationS3Uri *string `field:"required" json:"dataCapturedDestinationS3Uri" yaml:"dataCapturedDestinationS3Uri"`
	// `CfnModelExplainabilityJobDefinition.BatchTransformInputProperty.DatasetFormat`.
	DatasetFormat interface{} `field:"required" json:"datasetFormat" yaml:"datasetFormat"`
	// `CfnModelExplainabilityJobDefinition.BatchTransformInputProperty.LocalPath`.
	LocalPath *string `field:"required" json:"localPath" yaml:"localPath"`
	// `CfnModelExplainabilityJobDefinition.BatchTransformInputProperty.FeaturesAttribute`.
	FeaturesAttribute *string `field:"optional" json:"featuresAttribute" yaml:"featuresAttribute"`
	// `CfnModelExplainabilityJobDefinition.BatchTransformInputProperty.InferenceAttribute`.
	InferenceAttribute *string `field:"optional" json:"inferenceAttribute" yaml:"inferenceAttribute"`
	// `CfnModelExplainabilityJobDefinition.BatchTransformInputProperty.ProbabilityAttribute`.
	ProbabilityAttribute *string `field:"optional" json:"probabilityAttribute" yaml:"probabilityAttribute"`
	// `CfnModelExplainabilityJobDefinition.BatchTransformInputProperty.S3DataDistributionType`.
	S3DataDistributionType *string `field:"optional" json:"s3DataDistributionType" yaml:"s3DataDistributionType"`
	// `CfnModelExplainabilityJobDefinition.BatchTransformInputProperty.S3InputMode`.
	S3InputMode *string `field:"optional" json:"s3InputMode" yaml:"s3InputMode"`
}

