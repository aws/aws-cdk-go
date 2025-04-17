package awsec2


// Provides the options for the types of interruption for spot instances.
type SpotInstanceInterruption string

const (
	// The instance will stop when interrupted.
	SpotInstanceInterruption_STOP SpotInstanceInterruption = "STOP"
	// The instance will be terminated when interrupted.
	SpotInstanceInterruption_TERMINATE SpotInstanceInterruption = "TERMINATE"
	// The instance will hibernate when interrupted.
	SpotInstanceInterruption_HIBERNATE SpotInstanceInterruption = "HIBERNATE"
)

