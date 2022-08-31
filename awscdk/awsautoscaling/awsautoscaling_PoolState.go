package awsautoscaling


// The instance state in the warm pool.
// Experimental.
type PoolState string

const (
	// Hibernated.
	//
	// To use this state, prerequisites must be in place.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/hibernating-prerequisites.html
	//
	// Experimental.
	PoolState_HIBERNATED PoolState = "HIBERNATED"
	// Running.
	// Experimental.
	PoolState_RUNNING PoolState = "RUNNING"
	// Stopped.
	// Experimental.
	PoolState_STOPPED PoolState = "STOPPED"
)

