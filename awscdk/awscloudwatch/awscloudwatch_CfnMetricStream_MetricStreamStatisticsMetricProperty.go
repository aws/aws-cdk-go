package awscloudwatch


// A structure that specifies the metric name and namespace for one metric that is going to have additional statistics included in the stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricStreamStatisticsMetricProperty := &metricStreamStatisticsMetricProperty{
//   	metricName: jsii.String("metricName"),
//   	namespace: jsii.String("namespace"),
//   }
//
type CfnMetricStream_MetricStreamStatisticsMetricProperty struct {
	// The name of the metric.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The namespace of the metric.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

