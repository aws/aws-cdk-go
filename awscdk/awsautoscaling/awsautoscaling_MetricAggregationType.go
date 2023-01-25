package awsautoscaling


// How the scaling metric is going to be aggregated.
// Experimental.
type MetricAggregationType string

const (
	// Average.
	// Experimental.
	MetricAggregationType_AVERAGE MetricAggregationType = "AVERAGE"
	// Minimum.
	// Experimental.
	MetricAggregationType_MINIMUM MetricAggregationType = "MINIMUM"
	// Maximum.
	// Experimental.
	MetricAggregationType_MAXIMUM MetricAggregationType = "MAXIMUM"
)

