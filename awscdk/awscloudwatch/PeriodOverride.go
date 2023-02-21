package awscloudwatch


// Specify the period for graphs when the CloudWatch dashboard loads.
type PeriodOverride string

const (
	// Period of all graphs on the dashboard automatically adapt to the time range of the dashboard.
	PeriodOverride_AUTO PeriodOverride = "AUTO"
	// Period set for each graph will be used.
	PeriodOverride_INHERIT PeriodOverride = "INHERIT"
)

