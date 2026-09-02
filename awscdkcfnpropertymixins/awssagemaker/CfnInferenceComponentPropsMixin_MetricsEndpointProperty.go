package awssagemaker


// A metrics endpoint exposed by the container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   metricsEndpointProperty := &MetricsEndpointProperty{
//   	MetricPublishFrequencyInSeconds: jsii.Number(123),
//   	MetricsEndpointPath: jsii.String("metricsEndpointPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-metricsendpoint.html
//
type CfnInferenceComponentPropsMixin_MetricsEndpointProperty struct {
	// The interval, in seconds, at which container metrics scraped from the endpoint are published to Amazon CloudWatch.
	//
	// Valid values per the SageMaker API Reference are 10, 30, 60, 120, 180, 240 and 300; the service validates the value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-metricsendpoint.html#cfn-sagemaker-inferencecomponent-metricsendpoint-metricpublishfrequencyinseconds
	//
	MetricPublishFrequencyInSeconds *float64 `field:"optional" json:"metricPublishFrequencyInSeconds" yaml:"metricPublishFrequencyInSeconds"`
	// The path to the Prometheus formatted metrics endpoint exposed by the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-metricsendpoint.html#cfn-sagemaker-inferencecomponent-metricsendpoint-metricsendpointpath
	//
	MetricsEndpointPath *string `field:"optional" json:"metricsEndpointPath" yaml:"metricsEndpointPath"`
}

