package awsautoscaling


// The monitoring mode for instances launched in an autoscaling group.
type Monitoring string

const (
	// Generates metrics every 5 minutes.
	Monitoring_BASIC Monitoring = "BASIC"
	// Generates metrics every minute.
	Monitoring_DETAILED Monitoring = "DETAILED"
)

