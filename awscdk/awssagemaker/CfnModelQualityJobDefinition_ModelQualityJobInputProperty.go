package awssagemaker


// The input for the model quality monitoring job.
//
// Currently endpoints are supported for input for model quality monitoring jobs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelQualityJobInputProperty := &ModelQualityJobInputProperty{
//   	GroundTruthS3Input: &MonitoringGroundTruthS3InputProperty{
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//
//   	// the properties below are optional
//   	BatchTransformInput: &BatchTransformInputProperty{
//   		DataCapturedDestinationS3Uri: jsii.String("dataCapturedDestinationS3Uri"),
//   		DatasetFormat: &DatasetFormatProperty{
//   			Csv: &CsvProperty{
//   				Header: jsii.Boolean(false),
//   			},
//   			Json: &JsonProperty{
//   				Line: jsii.Boolean(false),
//   			},
//   			Parquet: jsii.Boolean(false),
//   		},
//   		LocalPath: jsii.String("localPath"),
//
//   		// the properties below are optional
//   		EndTimeOffset: jsii.String("endTimeOffset"),
//   		InferenceAttribute: jsii.String("inferenceAttribute"),
//   		ProbabilityAttribute: jsii.String("probabilityAttribute"),
//   		ProbabilityThresholdAttribute: jsii.Number(123),
//   		S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   		S3InputMode: jsii.String("s3InputMode"),
//   		StartTimeOffset: jsii.String("startTimeOffset"),
//   	},
//   	EndpointInput: &EndpointInputProperty{
//   		EndpointName: jsii.String("endpointName"),
//   		LocalPath: jsii.String("localPath"),
//
//   		// the properties below are optional
//   		EndTimeOffset: jsii.String("endTimeOffset"),
//   		InferenceAttribute: jsii.String("inferenceAttribute"),
//   		ProbabilityAttribute: jsii.String("probabilityAttribute"),
//   		ProbabilityThresholdAttribute: jsii.Number(123),
//   		S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   		S3InputMode: jsii.String("s3InputMode"),
//   		StartTimeOffset: jsii.String("startTimeOffset"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-modelqualityjobinput.html
//
type CfnModelQualityJobDefinition_ModelQualityJobInputProperty struct {
	// The ground truth label provided for the model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-modelqualityjobinput.html#cfn-sagemaker-modelqualityjobdefinition-modelqualityjobinput-groundtruths3input
	//
	GroundTruthS3Input interface{} `field:"required" json:"groundTruthS3Input" yaml:"groundTruthS3Input"`
	// Input object for the batch transform job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-modelqualityjobinput.html#cfn-sagemaker-modelqualityjobdefinition-modelqualityjobinput-batchtransforminput
	//
	BatchTransformInput interface{} `field:"optional" json:"batchTransformInput" yaml:"batchTransformInput"`
	// Input object for the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-modelqualityjobinput.html#cfn-sagemaker-modelqualityjobdefinition-modelqualityjobinput-endpointinput
	//
	EndpointInput interface{} `field:"optional" json:"endpointInput" yaml:"endpointInput"`
}

