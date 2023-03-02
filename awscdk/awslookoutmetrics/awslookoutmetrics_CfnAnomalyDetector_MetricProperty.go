package awslookoutmetrics


// A calculation made by contrasting a measure and a dimension from your source data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricProperty := &metricProperty{
//   	aggregationFunction: jsii.String("aggregationFunction"),
//   	metricName: jsii.String("metricName"),
//
//   	// the properties below are optional
//   	namespace: jsii.String("namespace"),
//   }
//
type CfnAnomalyDetector_MetricProperty struct {
	// The function with which the metric is calculated.
	AggregationFunction *string `field:"required" json:"aggregationFunction" yaml:"aggregationFunction"`
	// The name of the metric.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The namespace for the metric.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

