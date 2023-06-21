package awsrds


// The retention period for Performance Insight data, in days.
//
// Per https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#cfn-rds-dbinstance-performanceinsightsretentionperiod
// This must be either
// - 7 days (the default, free tier)
// - month * 31, where month is a number of months from 1-23
// - 731 (2 years).
type PerformanceInsightRetention string

const (
	// Default retention period of 7 days.
	PerformanceInsightRetention_DEFAULT PerformanceInsightRetention = "DEFAULT"
	PerformanceInsightRetention_MONTHS_1 PerformanceInsightRetention = "MONTHS_1"
	PerformanceInsightRetention_MONTHS_2 PerformanceInsightRetention = "MONTHS_2"
	PerformanceInsightRetention_MONTHS_3 PerformanceInsightRetention = "MONTHS_3"
	PerformanceInsightRetention_MONTHS_4 PerformanceInsightRetention = "MONTHS_4"
	PerformanceInsightRetention_MONTHS_5 PerformanceInsightRetention = "MONTHS_5"
	PerformanceInsightRetention_MONTHS_6 PerformanceInsightRetention = "MONTHS_6"
	PerformanceInsightRetention_MONTHS_7 PerformanceInsightRetention = "MONTHS_7"
	PerformanceInsightRetention_MONTHS_8 PerformanceInsightRetention = "MONTHS_8"
	PerformanceInsightRetention_MONTHS_9 PerformanceInsightRetention = "MONTHS_9"
	PerformanceInsightRetention_MONTHS_10 PerformanceInsightRetention = "MONTHS_10"
	PerformanceInsightRetention_MONTHS_11 PerformanceInsightRetention = "MONTHS_11"
	PerformanceInsightRetention_MONTHS_12 PerformanceInsightRetention = "MONTHS_12"
	PerformanceInsightRetention_MONTHS_13 PerformanceInsightRetention = "MONTHS_13"
	PerformanceInsightRetention_MONTHS_14 PerformanceInsightRetention = "MONTHS_14"
	PerformanceInsightRetention_MONTHS_15 PerformanceInsightRetention = "MONTHS_15"
	PerformanceInsightRetention_MONTHS_16 PerformanceInsightRetention = "MONTHS_16"
	PerformanceInsightRetention_MONTHS_17 PerformanceInsightRetention = "MONTHS_17"
	PerformanceInsightRetention_MONTHS_18 PerformanceInsightRetention = "MONTHS_18"
	PerformanceInsightRetention_MONTHS_19 PerformanceInsightRetention = "MONTHS_19"
	PerformanceInsightRetention_MONTHS_20 PerformanceInsightRetention = "MONTHS_20"
	PerformanceInsightRetention_MONTHS_21 PerformanceInsightRetention = "MONTHS_21"
	PerformanceInsightRetention_MONTHS_22 PerformanceInsightRetention = "MONTHS_22"
	PerformanceInsightRetention_MONTHS_23 PerformanceInsightRetention = "MONTHS_23"
	// Long term retention period of 2 years.
	PerformanceInsightRetention_LONG_TERM PerformanceInsightRetention = "LONG_TERM"
)

