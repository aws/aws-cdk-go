package awsrds


// The retention period for Performance Insight.
// Experimental.
type PerformanceInsightRetention string

const (
	// Default retention period of 7 days.
	// Experimental.
	PerformanceInsightRetention_DEFAULT PerformanceInsightRetention = "DEFAULT"
	// Long term retention period of 2 years.
	// Experimental.
	PerformanceInsightRetention_LONG_TERM PerformanceInsightRetention = "LONG_TERM"
)

