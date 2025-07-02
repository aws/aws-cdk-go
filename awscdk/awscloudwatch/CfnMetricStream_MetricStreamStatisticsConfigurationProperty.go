package awscloudwatch


// This structure specifies a list of additional statistics to stream, and the metrics to stream those additional statistics for.
//
// All metrics that match the combination of metric name and namespace will be streamed with the additional statistics, no matter their dimensions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricStreamStatisticsConfigurationProperty := &MetricStreamStatisticsConfigurationProperty{
//   	AdditionalStatistics: []*string{
//   		jsii.String("additionalStatistics"),
//   	},
//   	IncludeMetrics: []interface{}{
//   		&MetricStreamStatisticsMetricProperty{
//   			MetricName: jsii.String("metricName"),
//   			Namespace: jsii.String("namespace"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-metricstream-metricstreamstatisticsconfiguration.html
//
type CfnMetricStream_MetricStreamStatisticsConfigurationProperty struct {
	// The additional statistics to stream for the metrics listed in `IncludeMetrics` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-metricstream-metricstreamstatisticsconfiguration.html#cfn-cloudwatch-metricstream-metricstreamstatisticsconfiguration-additionalstatistics
	//
	AdditionalStatistics *[]*string `field:"required" json:"additionalStatistics" yaml:"additionalStatistics"`
	// An array that defines the metrics that are to have additional statistics streamed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-metricstream-metricstreamstatisticsconfiguration.html#cfn-cloudwatch-metricstream-metricstreamstatisticsconfiguration-includemetrics
	//
	IncludeMetrics interface{} `field:"required" json:"includeMetrics" yaml:"includeMetrics"`
}

