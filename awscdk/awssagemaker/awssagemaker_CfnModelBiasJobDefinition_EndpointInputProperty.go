package awssagemaker


// Input object for the endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointInputProperty := &endpointInputProperty{
//   	endpointName: jsii.String("endpointName"),
//   	localPath: jsii.String("localPath"),
//
//   	// the properties below are optional
//   	endTimeOffset: jsii.String("endTimeOffset"),
//   	featuresAttribute: jsii.String("featuresAttribute"),
//   	inferenceAttribute: jsii.String("inferenceAttribute"),
//   	probabilityAttribute: jsii.String("probabilityAttribute"),
//   	probabilityThresholdAttribute: jsii.Number(123),
//   	s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   	s3InputMode: jsii.String("s3InputMode"),
//   	startTimeOffset: jsii.String("startTimeOffset"),
//   }
//
type CfnModelBiasJobDefinition_EndpointInputProperty struct {
	// An endpoint in customer's account which has enabled `DataCaptureConfig` enabled.
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// Path to the filesystem where the endpoint data is available to the container.
	LocalPath *string `field:"required" json:"localPath" yaml:"localPath"`
	// If specified, monitoring jobs substract this time from the end time.
	//
	// For information about using offsets for scheduling monitoring jobs, see [Schedule Model Quality Monitoring Jobs](https://docs.aws.amazon.com/sagemaker/latest/dg/model-monitor-model-quality-schedule.html) .
	EndTimeOffset *string `field:"optional" json:"endTimeOffset" yaml:"endTimeOffset"`
	// The attributes of the input data that are the input features.
	FeaturesAttribute *string `field:"optional" json:"featuresAttribute" yaml:"featuresAttribute"`
	// The attribute of the input data that represents the ground truth label.
	InferenceAttribute *string `field:"optional" json:"inferenceAttribute" yaml:"inferenceAttribute"`
	// In a classification problem, the attribute that represents the class probability.
	ProbabilityAttribute *string `field:"optional" json:"probabilityAttribute" yaml:"probabilityAttribute"`
	// The threshold for the class probability to be evaluated as a positive result.
	ProbabilityThresholdAttribute *float64 `field:"optional" json:"probabilityThresholdAttribute" yaml:"probabilityThresholdAttribute"`
	// Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key.
	//
	// Defaults to `FullyReplicated`.
	S3DataDistributionType *string `field:"optional" json:"s3DataDistributionType" yaml:"s3DataDistributionType"`
	// Whether the `Pipe` or `File` is used as the input mode for transferring data for the monitoring job.
	//
	// `Pipe` mode is recommended for large datasets. `File` mode is useful for small files that fit in memory. Defaults to `File` .
	S3InputMode *string `field:"optional" json:"s3InputMode" yaml:"s3InputMode"`
	// If specified, monitoring jobs substract this time from the start time.
	//
	// For information about using offsets for scheduling monitoring jobs, see [Schedule Model Quality Monitoring Jobs](https://docs.aws.amazon.com/sagemaker/latest/dg/model-monitor-model-quality-schedule.html) .
	StartTimeOffset *string `field:"optional" json:"startTimeOffset" yaml:"startTimeOffset"`
}

