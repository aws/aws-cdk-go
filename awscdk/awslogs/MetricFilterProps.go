package awslogs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Properties for a MetricFilter.
//
// Example:
//   awscdk.NewMetricFilter(this, jsii.String("MetricFilter"), &MetricFilterProps{
//   	LogGroup: LogGroup,
//   	MetricNamespace: jsii.String("MyApp"),
//   	MetricName: jsii.String("Latency"),
//   	FilterPattern: awscdk.FilterPattern_Exists(jsii.String("$.latency")),
//   	MetricValue: jsii.String("$.latency"),
//   })
//
type MetricFilterProps struct {
	// Pattern to search for log events.
	FilterPattern IFilterPattern `field:"required" json:"filterPattern" yaml:"filterPattern"`
	// The name of the metric to emit.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The namespace of the metric to emit.
	MetricNamespace *string `field:"required" json:"metricNamespace" yaml:"metricNamespace"`
	// The value to emit if the pattern does not match a particular event.
	// Default: No metric emitted.
	//
	DefaultValue *float64 `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// The fields to use as dimensions for the metric.
	//
	// One metric filter can include as many as three dimensions.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-metricfilter-metrictransformation.html#cfn-logs-metricfilter-metrictransformation-dimensions
	//
	// Default: - No dimensions attached to metrics.
	//
	Dimensions *map[string]*string `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The name of the metric filter.
	// Default: - Cloudformation generated name.
	//
	FilterName *string `field:"optional" json:"filterName" yaml:"filterName"`
	// The value to emit for the metric.
	//
	// Can either be a literal number (typically "1"), or the name of a field in the structure
	// to take the value from the matched event. If you are using a field value, the field
	// value must have been matched using the pattern.
	//
	// If you want to specify a field from a matched JSON structure, use '$.fieldName',
	// and make sure the field is in the pattern (if only as '$.fieldName = *').
	//
	// If you want to specify a field from a matched space-delimited structure,
	// use '$fieldName'.
	// Default: "1".
	//
	MetricValue *string `field:"optional" json:"metricValue" yaml:"metricValue"`
	// The unit to assign to the metric.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-metricfilter-metrictransformation.html#cfn-logs-metricfilter-metrictransformation-unit
	//
	// Default: - No unit attached to metrics.
	//
	Unit awscloudwatch.Unit `field:"optional" json:"unit" yaml:"unit"`
	// The log group to create the filter on.
	LogGroup ILogGroup `field:"required" json:"logGroup" yaml:"logGroup"`
}

