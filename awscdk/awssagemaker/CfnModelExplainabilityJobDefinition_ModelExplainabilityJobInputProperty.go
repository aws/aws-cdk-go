package awssagemaker


// Inputs for the model explainability job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var json interface{}
//
//   modelExplainabilityJobInputProperty := &ModelExplainabilityJobInputProperty{
//   	BatchTransformInput: &BatchTransformInputProperty{
//   		DataCapturedDestinationS3Uri: jsii.String("dataCapturedDestinationS3Uri"),
//   		DatasetFormat: &DatasetFormatProperty{
//   			Csv: &CsvProperty{
//   				Header: jsii.Boolean(false),
//   			},
//   			Json: json,
//   			Parquet: jsii.Boolean(false),
//   		},
//   		LocalPath: jsii.String("localPath"),
//
//   		// the properties below are optional
//   		FeaturesAttribute: jsii.String("featuresAttribute"),
//   		InferenceAttribute: jsii.String("inferenceAttribute"),
//   		ProbabilityAttribute: jsii.String("probabilityAttribute"),
//   		S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   		S3InputMode: jsii.String("s3InputMode"),
//   	},
//   	EndpointInput: &EndpointInputProperty{
//   		EndpointName: jsii.String("endpointName"),
//   		LocalPath: jsii.String("localPath"),
//
//   		// the properties below are optional
//   		FeaturesAttribute: jsii.String("featuresAttribute"),
//   		InferenceAttribute: jsii.String("inferenceAttribute"),
//   		ProbabilityAttribute: jsii.String("probabilityAttribute"),
//   		S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   		S3InputMode: jsii.String("s3InputMode"),
//   	},
//   }
//
type CfnModelExplainabilityJobDefinition_ModelExplainabilityJobInputProperty struct {
	// `CfnModelExplainabilityJobDefinition.ModelExplainabilityJobInputProperty.BatchTransformInput`.
	BatchTransformInput interface{} `field:"optional" json:"batchTransformInput" yaml:"batchTransformInput"`
	// `CfnModelExplainabilityJobDefinition.ModelExplainabilityJobInputProperty.EndpointInput`.
	EndpointInput interface{} `field:"optional" json:"endpointInput" yaml:"endpointInput"`
}

