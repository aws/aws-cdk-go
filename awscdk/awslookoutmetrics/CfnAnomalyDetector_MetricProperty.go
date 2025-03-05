package awslookoutmetrics


// A calculation made by contrasting a measure and a dimension from your source data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricProperty := &MetricProperty{
//   	AggregationFunction: jsii.String("aggregationFunction"),
//   	MetricName: jsii.String("metricName"),
//
//   	// the properties below are optional
//   	Namespace: jsii.String("namespace"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-metric.html
//
type CfnAnomalyDetector_MetricProperty struct {
	// The function with which the metric is calculated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-metric.html#cfn-lookoutmetrics-anomalydetector-metric-aggregationfunction
	//
	AggregationFunction *string `field:"required" json:"aggregationFunction" yaml:"aggregationFunction"`
	// The name of the metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-metric.html#cfn-lookoutmetrics-anomalydetector-metric-metricname
	//
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The namespace for the metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-metric.html#cfn-lookoutmetrics-anomalydetector-metric-namespace
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

