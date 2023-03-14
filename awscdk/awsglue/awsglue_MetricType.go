package awsglue


// The Glue CloudWatch metric type.
// See: https://docs.aws.amazon.com/glue/latest/dg/monitoring-awsglue-with-cloudwatch-metrics.html
//
// Experimental.
type MetricType string

const (
	// A value at a point in time.
	// Experimental.
	MetricType_GAUGE MetricType = "GAUGE"
	// An aggregate number.
	// Experimental.
	MetricType_COUNT MetricType = "COUNT"
)

