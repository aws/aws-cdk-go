package awsautoscaling


// The type of update to perform on instances in this AutoScalingGroup.
// Deprecated: Use UpdatePolicy instead.
type UpdateType string

const (
	// Don't do anything.
	// Deprecated: Use UpdatePolicy instead.
	UpdateType_NONE UpdateType = "NONE"
	// Replace the entire AutoScalingGroup.
	//
	// Builds a new AutoScalingGroup first, then delete the old one.
	// Deprecated: Use UpdatePolicy instead.
	UpdateType_REPLACING_UPDATE UpdateType = "REPLACING_UPDATE"
	// Replace the instances in the AutoScalingGroup.
	// Deprecated: Use UpdatePolicy instead.
	UpdateType_ROLLING_UPDATE UpdateType = "ROLLING_UPDATE"
)

