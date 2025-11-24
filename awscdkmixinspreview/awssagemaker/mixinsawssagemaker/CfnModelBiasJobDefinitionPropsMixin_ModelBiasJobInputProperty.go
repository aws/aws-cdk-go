package mixinsawssagemaker


// Inputs for the model bias job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   modelBiasJobInputProperty := &ModelBiasJobInputProperty{
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
//   		EndTimeOffset: jsii.String("endTimeOffset"),
//   		FeaturesAttribute: jsii.String("featuresAttribute"),
//   		InferenceAttribute: jsii.String("inferenceAttribute"),
//   		LocalPath: jsii.String("localPath"),
//   		ProbabilityAttribute: jsii.String("probabilityAttribute"),
//   		ProbabilityThresholdAttribute: jsii.Number(123),
//   		S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   		S3InputMode: jsii.String("s3InputMode"),
//   		StartTimeOffset: jsii.String("startTimeOffset"),
//   	},
//   	EndpointInput: &EndpointInputProperty{
//   		EndpointName: jsii.String("endpointName"),
//   		EndTimeOffset: jsii.String("endTimeOffset"),
//   		FeaturesAttribute: jsii.String("featuresAttribute"),
//   		InferenceAttribute: jsii.String("inferenceAttribute"),
//   		LocalPath: jsii.String("localPath"),
//   		ProbabilityAttribute: jsii.String("probabilityAttribute"),
//   		ProbabilityThresholdAttribute: jsii.Number(123),
//   		S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   		S3InputMode: jsii.String("s3InputMode"),
//   		StartTimeOffset: jsii.String("startTimeOffset"),
//   	},
//   	GroundTruthS3Input: &MonitoringGroundTruthS3InputProperty{
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-modelbiasjobinput.html
//
type CfnModelBiasJobDefinitionPropsMixin_ModelBiasJobInputProperty struct {
	// Input object for the batch transform job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-modelbiasjobinput.html#cfn-sagemaker-modelbiasjobdefinition-modelbiasjobinput-batchtransforminput
	//
	BatchTransformInput interface{} `field:"optional" json:"batchTransformInput" yaml:"batchTransformInput"`
	// Input object for the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-modelbiasjobinput.html#cfn-sagemaker-modelbiasjobdefinition-modelbiasjobinput-endpointinput
	//
	EndpointInput interface{} `field:"optional" json:"endpointInput" yaml:"endpointInput"`
	// Location of ground truth labels to use in model bias job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-modelbiasjobinput.html#cfn-sagemaker-modelbiasjobdefinition-modelbiasjobinput-groundtruths3input
	//
	GroundTruthS3Input interface{} `field:"optional" json:"groundTruthS3Input" yaml:"groundTruthS3Input"`
}

