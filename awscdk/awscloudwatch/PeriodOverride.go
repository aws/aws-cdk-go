package awscloudwatch


// Specify the period for graphs when the CloudWatch dashboard loads.
// Experimental.
type PeriodOverride string

const (
	// Period of all graphs on the dashboard automatically adapt to the time range of the dashboard.
	// Experimental.
	PeriodOverride_AUTO PeriodOverride = "AUTO"
	// Period set for each graph will be used.
	// Experimental.
	PeriodOverride_INHERIT PeriodOverride = "INHERIT"
)

