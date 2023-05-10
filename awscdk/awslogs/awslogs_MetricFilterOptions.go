package awslogs


// Properties for a MetricFilter created from a LogGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var filterPattern iFilterPattern
//
//   metricFilterOptions := &MetricFilterOptions{
//   	FilterPattern: filterPattern,
//   	MetricName: jsii.String("metricName"),
//   	MetricNamespace: jsii.String("metricNamespace"),
//
//   	// the properties below are optional
//   	DefaultValue: jsii.Number(123),
//   	MetricValue: jsii.String("metricValue"),
//   }
//
// Experimental.
type MetricFilterOptions struct {
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
}

