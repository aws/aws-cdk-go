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
//   	EndTimeOffset: jsii.String("endTimeOffset"),
//   	FeaturesAttribute: jsii.String("featuresAttribute"),
//   	InferenceAttribute: jsii.String("inferenceAttribute"),
//   	ProbabilityAttribute: jsii.String("probabilityAttribute"),
//   	ProbabilityThresholdAttribute: jsii.Number(123),
//   	S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   	S3InputMode: jsii.String("s3InputMode"),
//   	StartTimeOffset: jsii.String("startTimeOffset"),
//   }
//
type CfnModelBiasJobDefinition_BatchTransformInputProperty struct {
	// `CfnModelBiasJobDefinition.BatchTransformInputProperty.DataCapturedDestinationS3Uri`.
	DataCapturedDestinationS3Uri *string `field:"required" json:"dataCapturedDestinationS3Uri" yaml:"dataCapturedDestinationS3Uri"`
	// `CfnModelBiasJobDefinition.BatchTransformInputProperty.DatasetFormat`.
	DatasetFormat interface{} `field:"required" json:"datasetFormat" yaml:"datasetFormat"`
	// `CfnModelBiasJobDefinition.BatchTransformInputProperty.LocalPath`.
	LocalPath *string `field:"required" json:"localPath" yaml:"localPath"`
	// `CfnModelBiasJobDefinition.BatchTransformInputProperty.EndTimeOffset`.
	EndTimeOffset *string `field:"optional" json:"endTimeOffset" yaml:"endTimeOffset"`
	// `CfnModelBiasJobDefinition.BatchTransformInputProperty.FeaturesAttribute`.
	FeaturesAttribute *string `field:"optional" json:"featuresAttribute" yaml:"featuresAttribute"`
	// `CfnModelBiasJobDefinition.BatchTransformInputProperty.InferenceAttribute`.
	InferenceAttribute *string `field:"optional" json:"inferenceAttribute" yaml:"inferenceAttribute"`
	// `CfnModelBiasJobDefinition.BatchTransformInputProperty.ProbabilityAttribute`.
	ProbabilityAttribute *string `field:"optional" json:"probabilityAttribute" yaml:"probabilityAttribute"`
	// `CfnModelBiasJobDefinition.BatchTransformInputProperty.ProbabilityThresholdAttribute`.
	ProbabilityThresholdAttribute *float64 `field:"optional" json:"probabilityThresholdAttribute" yaml:"probabilityThresholdAttribute"`
	// `CfnModelBiasJobDefinition.BatchTransformInputProperty.S3DataDistributionType`.
	S3DataDistributionType *string `field:"optional" json:"s3DataDistributionType" yaml:"s3DataDistributionType"`
	// `CfnModelBiasJobDefinition.BatchTransformInputProperty.S3InputMode`.
	S3InputMode *string `field:"optional" json:"s3InputMode" yaml:"s3InputMode"`
	// `CfnModelBiasJobDefinition.BatchTransformInputProperty.StartTimeOffset`.
	StartTimeOffset *string `field:"optional" json:"startTimeOffset" yaml:"startTimeOffset"`
}

