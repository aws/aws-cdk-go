package awssagemaker


// Input object for the endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointInputProperty := &EndpointInputProperty{
//   	EndpointName: jsii.String("endpointName"),
//   	LocalPath: jsii.String("localPath"),
//
//   	// the properties below are optional
//   	EndTimeOffset: jsii.String("endTimeOffset"),
//   	InferenceAttribute: jsii.String("inferenceAttribute"),
//   	ProbabilityAttribute: jsii.String("probabilityAttribute"),
//   	ProbabilityThresholdAttribute: jsii.Number(123),
//   	S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   	S3InputMode: jsii.String("s3InputMode"),
//   	StartTimeOffset: jsii.String("startTimeOffset"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-endpointinput.html
//
type CfnModelQualityJobDefinition_EndpointInputProperty struct {
	// An endpoint in customer's account which has enabled `DataCaptureConfig` enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-endpointinput.html#cfn-sagemaker-modelqualityjobdefinition-endpointinput-endpointname
	//
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// Path to the filesystem where the endpoint data is available to the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-endpointinput.html#cfn-sagemaker-modelqualityjobdefinition-endpointinput-localpath
	//
	LocalPath *string `field:"required" json:"localPath" yaml:"localPath"`
	// If specified, monitoring jobs substract this time from the end time.
	//
	// For information about using offsets for scheduling monitoring jobs, see [Schedule Model Quality Monitoring Jobs](https://docs.aws.amazon.com/sagemaker/latest/dg/model-monitor-model-quality-schedule.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-endpointinput.html#cfn-sagemaker-modelqualityjobdefinition-endpointinput-endtimeoffset
	//
	EndTimeOffset *string `field:"optional" json:"endTimeOffset" yaml:"endTimeOffset"`
	// The attribute of the input data that represents the ground truth label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-endpointinput.html#cfn-sagemaker-modelqualityjobdefinition-endpointinput-inferenceattribute
	//
	InferenceAttribute *string `field:"optional" json:"inferenceAttribute" yaml:"inferenceAttribute"`
	// In a classification problem, the attribute that represents the class probability.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-endpointinput.html#cfn-sagemaker-modelqualityjobdefinition-endpointinput-probabilityattribute
	//
	ProbabilityAttribute *string `field:"optional" json:"probabilityAttribute" yaml:"probabilityAttribute"`
	// The threshold for the class probability to be evaluated as a positive result.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-endpointinput.html#cfn-sagemaker-modelqualityjobdefinition-endpointinput-probabilitythresholdattribute
	//
	ProbabilityThresholdAttribute *float64 `field:"optional" json:"probabilityThresholdAttribute" yaml:"probabilityThresholdAttribute"`
	// Whether input data distributed in Amazon S3 is fully replicated or sharded by an Amazon S3 key.
	//
	// Defaults to `FullyReplicated`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-endpointinput.html#cfn-sagemaker-modelqualityjobdefinition-endpointinput-s3datadistributiontype
	//
	S3DataDistributionType *string `field:"optional" json:"s3DataDistributionType" yaml:"s3DataDistributionType"`
	// Whether the `Pipe` or `File` is used as the input mode for transferring data for the monitoring job.
	//
	// `Pipe` mode is recommended for large datasets. `File` mode is useful for small files that fit in memory. Defaults to `File` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-endpointinput.html#cfn-sagemaker-modelqualityjobdefinition-endpointinput-s3inputmode
	//
	S3InputMode *string `field:"optional" json:"s3InputMode" yaml:"s3InputMode"`
	// If specified, monitoring jobs substract this time from the start time.
	//
	// For information about using offsets for scheduling monitoring jobs, see [Schedule Model Quality Monitoring Jobs](https://docs.aws.amazon.com/sagemaker/latest/dg/model-monitor-model-quality-schedule.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-endpointinput.html#cfn-sagemaker-modelqualityjobdefinition-endpointinput-starttimeoffset
	//
	StartTimeOffset *string `field:"optional" json:"startTimeOffset" yaml:"startTimeOffset"`
}

