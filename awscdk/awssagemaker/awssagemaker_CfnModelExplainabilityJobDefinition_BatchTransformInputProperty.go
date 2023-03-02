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
//   	featuresAttribute: jsii.String("featuresAttribute"),
//   	inferenceAttribute: jsii.String("inferenceAttribute"),
//   	probabilityAttribute: jsii.String("probabilityAttribute"),
//   	s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   	s3InputMode: jsii.String("s3InputMode"),
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

