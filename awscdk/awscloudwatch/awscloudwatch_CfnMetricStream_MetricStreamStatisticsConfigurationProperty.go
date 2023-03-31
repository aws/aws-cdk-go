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
//   metricStreamStatisticsConfigurationProperty := &metricStreamStatisticsConfigurationProperty{
//   	additionalStatistics: []*string{
//   		jsii.String("additionalStatistics"),
//   	},
//   	includeMetrics: []interface{}{
//   		&metricStreamStatisticsMetricProperty{
//   			metricName: jsii.String("metricName"),
//   			namespace: jsii.String("namespace"),
//   		},
//   	},
//   }
//
type CfnMetricStream_MetricStreamStatisticsConfigurationProperty struct {
	// The additional statistics to stream for the metrics listed in `IncludeMetrics` .
	AdditionalStatistics *[]*string `field:"required" json:"additionalStatistics" yaml:"additionalStatistics"`
	// An array that defines the metrics that are to have additional statistics streamed.
	IncludeMetrics interface{} `field:"required" json:"includeMetrics" yaml:"includeMetrics"`
}

