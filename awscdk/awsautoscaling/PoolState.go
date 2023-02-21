package awsautoscaling


// The instance state in the warm pool.
type PoolState string

const (
	// Hibernated.
	//
	// To use this state, prerequisites must be in place.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/hibernating-prerequisites.html
	//
	PoolState_HIBERNATED PoolState = "HIBERNATED"
	// Running.
	PoolState_RUNNING PoolState = "RUNNING"
	// Stopped.
	PoolState_STOPPED PoolState = "STOPPED"
)

