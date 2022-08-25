package awslogs


// Properties for a MetricFilter.
//
// Example:
//   awscdk.NewMetricFilter(this, jsii.String("MetricFilter"), &metricFilterProps{
//   	logGroup: logGroup,
//   	metricNamespace: jsii.String("MyApp"),
//   	metricName: jsii.String("Latency"),
//   	filterPattern: awscdk.FilterPattern.exists(jsii.String("$.latency")),
//   	metricValue: jsii.String("$.latency"),
//   })
//
// Experimental.
type MetricFilterProps struct {
	// Pattern to search for log events.
	// Experimental.
	FilterPattern IFilterPattern `field:"required" json:"filterPattern" yaml:"filterPattern"`
	// The name of the metric to emit.
	// Experimental.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The namespace of the metric to emit.
	// Experimental.
	MetricNamespace *string `field:"required" json:"metricNamespace" yaml:"metricNamespace"`
	// The value to emit if the pattern does not match a particular event.
	// Experimental.
	DefaultValue *float64 `field:"optional" json:"defaultValue" yaml:"defaultValue"`
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
	// Experimental.
	MetricValue *string `field:"optional" json:"metricValue" yaml:"metricValue"`
	// The log group to create the filter on.
	// Experimental.
	LogGroup ILogGroup `field:"required" json:"logGroup" yaml:"logGroup"`
}

