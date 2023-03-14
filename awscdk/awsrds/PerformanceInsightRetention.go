package awsrds


// The retention period for Performance Insight.
type PerformanceInsightRetention string

const (
	// Default retention period of 7 days.
	PerformanceInsightRetention_DEFAULT PerformanceInsightRetention = "DEFAULT"
	// Long term retention period of 2 years.
	PerformanceInsightRetention_LONG_TERM PerformanceInsightRetention = "LONG_TERM"
)

