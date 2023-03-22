package awscloudwatch


// This structure contains the name of one of the metric namespaces that is listed in a filter of a metric stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricStreamFilterProperty := &metricStreamFilterProperty{
//   	namespace: jsii.String("namespace"),
//   }
//
type CfnMetricStream_MetricStreamFilterProperty struct {
	// The name of the metric namespace in the filter.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

