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
//   metricAlarmConfig := &MetricAlarmConfig{
//   	MetricName: jsii.String("metricName"),
//   	Namespace: jsii.String("namespace"),
//   	Period: jsii.Number(123),
//
//   	// the properties below are optional
//   	Dimensions: []dimension{
//   		&dimension{
//   			Name: jsii.String("name"),
//   			Value: value,
//   		},
//   	},
//   	ExtendedStatistic: jsii.String("extendedStatistic"),
//   	Statistic: awscdk.Aws_cloudwatch.Statistic_SAMPLE_COUNT,
//   	Unit: awscdk.*Aws_cloudwatch.Unit_SECONDS,
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

