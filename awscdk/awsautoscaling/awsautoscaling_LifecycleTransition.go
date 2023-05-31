package awsautoscaling


// What instance transition to attach the hook to.
// Experimental.
type LifecycleTransition string

const (
	// Execute the hook when an instance is about to be added.
	// Experimental.
	LifecycleTransition_INSTANCE_LAUNCHING LifecycleTransition = "INSTANCE_LAUNCHING"
	// Execute the hook when an instance is about to be terminated.
	// Experimental.
	LifecycleTransition_INSTANCE_TERMINATING LifecycleTransition = "INSTANCE_TERMINATING"
)

