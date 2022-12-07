package awssagemaker


// Input object for the batch transform job.
//
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
	// The Amazon S3 location being used to capture the data.
	DataCapturedDestinationS3Uri *string `field:"required" json:"dataCapturedDestinationS3Uri" yaml:"dataCapturedDestinationS3Uri"`
	// The dataset format for your batch transform job.
	DatasetFormat interface{} `field:"required" json:"datasetFormat" yaml:"datasetFormat"`
	// Path to the filesystem where the batch transform data is available to the container.
	LocalPath *string `field:"required" json:"localPath" yaml:"localPath"`
	// The attributes of the input data that are the input features.
	FeaturesAttribute *string `field:"optional" json:"featuresAttribute" yaml:"featuresAttribute"`
	// The attribute of the input data that represents the ground truth label.
	InferenceAttribute *string `field:"optional" json:"inferenceAttribute" yaml:"inferenceAttribute"`
	// In a classification problem, the attribute that represents the class probability.
	ProbabilityAttribute *string `field:"optional" json:"probabilityAttribute" yaml:"probabilityAttribute"`
	// Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key.
	//
	// Defaults to `FullyReplicated`.
	S3DataDistributionType *string `field:"optional" json:"s3DataDistributionType" yaml:"s3DataDistributionType"`
	// Whether the `Pipe` or `File` is used as the input mode for transferring data for the monitoring job.
	//
	// `Pipe` mode is recommended for large datasets. `File` mode is useful for small files that fit in memory. Defaults to `File` .
	S3InputMode *string `field:"optional" json:"s3InputMode" yaml:"s3InputMode"`
}

