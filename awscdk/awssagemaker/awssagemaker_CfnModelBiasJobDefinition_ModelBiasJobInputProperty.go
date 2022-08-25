package awssagemaker


// Inputs for the model bias job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelBiasJobInputProperty := &modelBiasJobInputProperty{
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
//   	groundTruthS3Input: &monitoringGroundTruthS3InputProperty{
//   		s3Uri: jsii.String("s3Uri"),
//   	},
//   }
//
type CfnModelBiasJobDefinition_ModelBiasJobInputProperty struct {
	// Input object for the endpoint.
	EndpointInput interface{} `field:"required" json:"endpointInput" yaml:"endpointInput"`
	// Location of ground truth labels to use in model bias job.
	GroundTruthS3Input interface{} `field:"required" json:"groundTruthS3Input" yaml:"groundTruthS3Input"`
}

