package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var json interface{}
//
//   batchTransformInputProperty := &batchTransformInputProperty{
//   	dataCapturedDestinationS3Uri: jsii.String("dataCapturedDestinationS3Uri"),
//   	datasetFormat: &datasetFormatProperty{
//   		csv: &csvProperty{
//   			header: jsii.Boolean(false),
//   		},
//   		json: json,
//   		parquet: jsii.Boolean(false),
//   	},
//   	localPath: jsii.String("localPath"),
//
//   	// the properties below are optional
//   	endTimeOffset: jsii.String("endTimeOffset"),
//   	inferenceAttribute: jsii.String("inferenceAttribute"),
//   	probabilityAttribute: jsii.String("probabilityAttribute"),
//   	probabilityThresholdAttribute: jsii.Number(123),
//   	s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   	s3InputMode: jsii.String("s3InputMode"),
//   	startTimeOffset: jsii.String("startTimeOffset"),
//   }
//
type CfnModelQualityJobDefinition_BatchTransformInputProperty struct {
	// `CfnModelQualityJobDefinition.BatchTransformInputProperty.DataCapturedDestinationS3Uri`.
	DataCapturedDestinationS3Uri *string `field:"required" json:"dataCapturedDestinationS3Uri" yaml:"dataCapturedDestinationS3Uri"`
	// `CfnModelQualityJobDefinition.BatchTransformInputProperty.DatasetFormat`.
	DatasetFormat interface{} `field:"required" json:"datasetFormat" yaml:"datasetFormat"`
	// `CfnModelQualityJobDefinition.BatchTransformInputProperty.LocalPath`.
	LocalPath *string `field:"required" json:"localPath" yaml:"localPath"`
	// `CfnModelQualityJobDefinition.BatchTransformInputProperty.EndTimeOffset`.
	EndTimeOffset *string `field:"optional" json:"endTimeOffset" yaml:"endTimeOffset"`
	// `CfnModelQualityJobDefinition.BatchTransformInputProperty.InferenceAttribute`.
	InferenceAttribute *string `field:"optional" json:"inferenceAttribute" yaml:"inferenceAttribute"`
	// `CfnModelQualityJobDefinition.BatchTransformInputProperty.ProbabilityAttribute`.
	ProbabilityAttribute *string `field:"optional" json:"probabilityAttribute" yaml:"probabilityAttribute"`
	// `CfnModelQualityJobDefinition.BatchTransformInputProperty.ProbabilityThresholdAttribute`.
	ProbabilityThresholdAttribute *float64 `field:"optional" json:"probabilityThresholdAttribute" yaml:"probabilityThresholdAttribute"`
	// `CfnModelQualityJobDefinition.BatchTransformInputProperty.S3DataDistributionType`.
	S3DataDistributionType *string `field:"optional" json:"s3DataDistributionType" yaml:"s3DataDistributionType"`
	// `CfnModelQualityJobDefinition.BatchTransformInputProperty.S3InputMode`.
	S3InputMode *string `field:"optional" json:"s3InputMode" yaml:"s3InputMode"`
	// `CfnModelQualityJobDefinition.BatchTransformInputProperty.StartTimeOffset`.
	StartTimeOffset *string `field:"optional" json:"startTimeOffset" yaml:"startTimeOffset"`
}

