package awscloudwatch


// Properties used to construct the Metric identifying part of an Alarm.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var value interface{}
//
//   metricAlarmConfig := &metricAlarmConfig{
//   	metricName: jsii.String("metricName"),
//   	namespace: jsii.String("namespace"),
//   	period: jsii.Number(123),
//
//   	// the properties below are optional
//   	dimensions: []dimension{
//   		&dimension{
//   			name: jsii.String("name"),
//   			value: value,
//   		},
//   	},
//   	extendedStatistic: jsii.String("extendedStatistic"),
//   	statistic: awscdk.Aws_cloudwatch.statistic_SAMPLE_COUNT,
//   	unit: awscdk.*Aws_cloudwatch.unit_SECONDS,
//   }
//
// Deprecated: Replaced by MetricConfig.
type MetricAlarmConfig struct {
	// Name of the metric.
	// Deprecated: Replaced by MetricConfig.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Namespace of the metric.
	// Deprecated: Replaced by MetricConfig.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// How many seconds to aggregate over.
	// Deprecated: Replaced by MetricConfig.
	Period *float64 `field:"required" json:"period" yaml:"period"`
	// The dimensions to apply to the alarm.
	// Deprecated: Replaced by MetricConfig.
	Dimensions *[]*Dimension `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Percentile aggregation function to use.
	// Deprecated: Replaced by MetricConfig.
	ExtendedStatistic *string `field:"optional" json:"extendedStatistic" yaml:"extendedStatistic"`
	// Simple aggregation function to use.
	// Deprecated: Replaced by MetricConfig.
	Statistic Statistic `field:"optional" json:"statistic" yaml:"statistic"`
	// The unit of the alarm.
	// Deprecated: Replaced by MetricConfig.
	Unit Unit `field:"optional" json:"unit" yaml:"unit"`
}

