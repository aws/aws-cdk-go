package awsapplicationautoscaling


// How the scaling metric is going to be aggregated.
type MetricAggregationType string

const (
	// Average.
	MetricAggregationType_AVERAGE MetricAggregationType = "AVERAGE"
	// Minimum.
	MetricAggregationType_MINIMUM MetricAggregationType = "MINIMUM"
	// Maximum.
	MetricAggregationType_MAXIMUM MetricAggregationType = "MAXIMUM"
)

