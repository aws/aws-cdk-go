package awsec2


// Provides the options for the types of interruption for spot instances.
// Experimental.
type SpotInstanceInterruption string

const (
	// The instance will stop when interrupted.
	// Experimental.
	SpotInstanceInterruption_STOP SpotInstanceInterruption = "STOP"
	// The instance will be terminated when interrupted.
	// Experimental.
	SpotInstanceInterruption_TERMINATE SpotInstanceInterruption = "TERMINATE"
	// The instance will hibernate when interrupted.
	// Experimental.
	SpotInstanceInterruption_HIBERNATE SpotInstanceInterruption = "HIBERNATE"
)

