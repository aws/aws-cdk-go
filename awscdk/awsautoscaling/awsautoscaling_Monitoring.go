package awsautoscaling


// The monitoring mode for instances launched in an autoscaling group.
// Experimental.
type Monitoring string

const (
	// Generates metrics every 5 minutes.
	// Experimental.
	Monitoring_BASIC Monitoring = "BASIC"
	// Generates metrics every minute.
	// Experimental.
	Monitoring_DETAILED Monitoring = "DETAILED"
)

