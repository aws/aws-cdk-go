package awscloudwatch


// This structure contains a metric namespace and optionally, a list of metric names, to either include in a metric ' stream or exclude from a metric stream.
//
// A metric stream's filters can include up to 1000 total names. This limit applies to the sum of namespace names and metric names in the filters. For example, this could include 10 metric namespace filters with 99 metrics each, or 20 namespace filters with 49 metrics specified in each filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricStreamFilterProperty := &MetricStreamFilterProperty{
//   	Namespace: jsii.String("namespace"),
//
//   	// the properties below are optional
//   	MetricNames: []*string{
//   		jsii.String("metricNames"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-metricstream-metricstreamfilter.html
//
type CfnMetricStream_MetricStreamFilterProperty struct {
	// The name of the metric namespace in the filter.
	//
	// The namespace can contain only ASCII printable characters (ASCII range 32 through 126). It must contain at least one non-whitespace character.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-metricstream-metricstreamfilter.html#cfn-cloudwatch-metricstream-metricstreamfilter-namespace
	//
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// The names of the metrics to either include or exclude from the metric stream.
	//
	// If you omit this parameter, all metrics in the namespace are included or excluded, depending on whether this filter is specified as an exclude filter or an include filter.
	//
	// Each metric name can contain only ASCII printable characters (ASCII range 32 through 126). Each metric name must contain at least one non-whitespace character.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-metricstream-metricstreamfilter.html#cfn-cloudwatch-metricstream-metricstreamfilter-metricnames
	//
	MetricNames *[]*string `field:"optional" json:"metricNames" yaml:"metricNames"`
}

