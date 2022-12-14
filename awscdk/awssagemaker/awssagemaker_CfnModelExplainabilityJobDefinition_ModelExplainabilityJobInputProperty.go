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
//   modelExplainabilityJobInputProperty := &modelExplainabilityJobInputProperty{
//   	batchTransformInput: &batchTransformInputProperty{
//   		dataCapturedDestinationS3Uri: jsii.String("dataCapturedDestinationS3Uri"),
//   		datasetFormat: &datasetFormatProperty{
//   			csv: &csvProperty{
//   				header: jsii.Boolean(false),
//   			},
//   			json: json,
//   			parquet: jsii.Boolean(false),
//   		},
//   		localPath: jsii.String("localPath"),
//
//   		// the properties below are optional
//   		featuresAttribute: jsii.String("featuresAttribute"),
//   		inferenceAttribute: jsii.String("inferenceAttribute"),
//   		probabilityAttribute: jsii.String("probabilityAttribute"),
//   		s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   		s3InputMode: jsii.String("s3InputMode"),
//   	},
//   	endpointInput: &endpointInputProperty{
//   		endpointName: jsii.String("endpointName"),
//   		localPath: jsii.String("localPath"),
//
//   		// the properties below are optional
//   		featuresAttribute: jsii.String("featuresAttribute"),
//   		inferenceAttribute: jsii.String("inferenceAttribute"),
//   		probabilityAttribute: jsii.String("probabilityAttribute"),
//   		s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   		s3InputMode: jsii.String("s3InputMode"),
//   	},
//   }
//
type CfnModelExplainabilityJobDefinition_ModelExplainabilityJobInputProperty struct {
	// `CfnModelExplainabilityJobDefinition.ModelExplainabilityJobInputProperty.BatchTransformInput`.
	BatchTransformInput interface{} `field:"optional" json:"batchTransformInput" yaml:"batchTransformInput"`
	// `CfnModelExplainabilityJobDefinition.ModelExplainabilityJobInputProperty.EndpointInput`.
	EndpointInput interface{} `field:"optional" json:"endpointInput" yaml:"endpointInput"`
}

