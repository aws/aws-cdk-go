package awssagemaker


// Inputs for the model bias job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var json interface{}
//
//   modelBiasJobInputProperty := &modelBiasJobInputProperty{
//   	groundTruthS3Input: &monitoringGroundTruthS3InputProperty{
//   		s3Uri: jsii.String("s3Uri"),
//   	},
//
//   	// the properties below are optional
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
//   		endTimeOffset: jsii.String("endTimeOffset"),
//   		featuresAttribute: jsii.String("featuresAttribute"),
//   		inferenceAttribute: jsii.String("inferenceAttribute"),
//   		probabilityAttribute: jsii.String("probabilityAttribute"),
//   		probabilityThresholdAttribute: jsii.Number(123),
//   		s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   		s3InputMode: jsii.String("s3InputMode"),
//   		startTimeOffset: jsii.String("startTimeOffset"),
//   	},
//   	endpointInput: &endpointInputProperty{
//   		endpointName: jsii.String("endpointName"),
//   		localPath: jsii.String("localPath"),
//
//   		// the properties below are optional
//   		endTimeOffset: jsii.String("endTimeOffset"),
//   		featuresAttribute: jsii.String("featuresAttribute"),
//   		inferenceAttribute: jsii.String("inferenceAttribute"),
//   		probabilityAttribute: jsii.String("probabilityAttribute"),
//   		probabilityThresholdAttribute: jsii.Number(123),
//   		s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   		s3InputMode: jsii.String("s3InputMode"),
//   		startTimeOffset: jsii.String("startTimeOffset"),
//   	},
//   }
//
type CfnModelBiasJobDefinition_ModelBiasJobInputProperty struct {
	// Location of ground truth labels to use in model bias job.
	GroundTruthS3Input interface{} `field:"required" json:"groundTruthS3Input" yaml:"groundTruthS3Input"`
	// `CfnModelBiasJobDefinition.ModelBiasJobInputProperty.BatchTransformInput`.
	BatchTransformInput interface{} `field:"optional" json:"batchTransformInput" yaml:"batchTransformInput"`
	// Input object for the endpoint.
	EndpointInput interface{} `field:"optional" json:"endpointInput" yaml:"endpointInput"`
}

