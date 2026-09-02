package awssagemaker


// Specifies the metrics that the endpoint publishes to Amazon CloudWatch, the frequency of publication, and whether to enable enhanced or detailed observability metrics.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricsConfigProperty := &MetricsConfigProperty{
//   	EnableDetailedObservability: jsii.Boolean(false),
//   	EnableEnhancedMetrics: jsii.Boolean(false),
//   	MetricPublishFrequencyInSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-metricsconfig.html
//
type CfnEndpointConfig_MetricsConfigProperty struct {
	// Specifies whether to enable detailed observability for the endpoint.
	//
	// When set to true, the endpoint publishes container-level inference metrics, per-GPU metrics, per-instance host metrics, and inference component placement metrics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-metricsconfig.html#cfn-sagemaker-endpointconfig-metricsconfig-enabledetailedobservability
	//
	EnableDetailedObservability interface{} `field:"optional" json:"enableDetailedObservability" yaml:"enableDetailedObservability"`
	// Specifies whether to enable enhanced metrics for the endpoint.
	//
	// Enhanced metrics provide utilization and invocation data at instance and container granularity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-metricsconfig.html#cfn-sagemaker-endpointconfig-metricsconfig-enableenhancedmetrics
	//
	EnableEnhancedMetrics interface{} `field:"optional" json:"enableEnhancedMetrics" yaml:"enableEnhancedMetrics"`
	// The interval, in seconds, at which the endpoint publishes metrics to Amazon CloudWatch.
	//
	// Valid values are 10, 30, 60, 120, 180, 240, and 300. The default is 60.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-metricsconfig.html#cfn-sagemaker-endpointconfig-metricsconfig-metricpublishfrequencyinseconds
	//
	MetricPublishFrequencyInSeconds *float64 `field:"optional" json:"metricPublishFrequencyInSeconds" yaml:"metricPublishFrequencyInSeconds"`
}

