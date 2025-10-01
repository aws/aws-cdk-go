package awsecs


// The monitoring configuration for EC2 instances.
type InstanceMonitoring string

const (
	// Basic monitoring (5-minute intervals).
	InstanceMonitoring_BASIC InstanceMonitoring = "BASIC"
	// Detailed monitoring (1-minute intervals).
	InstanceMonitoring_DETAILED InstanceMonitoring = "DETAILED"
)

