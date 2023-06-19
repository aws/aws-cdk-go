package awsec2


// Provides the options for specifying the instance initiated shutdown behavior.
// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingInstanceInitiatedShutdownBehavior
//
// Experimental.
type InstanceInitiatedShutdownBehavior string

const (
	// The instance will stop when it initiates a shutdown.
	// Experimental.
	InstanceInitiatedShutdownBehavior_STOP InstanceInitiatedShutdownBehavior = "STOP"
	// The instance will be terminated when it initiates a shutdown.
	// Experimental.
	InstanceInitiatedShutdownBehavior_TERMINATE InstanceInitiatedShutdownBehavior = "TERMINATE"
)

