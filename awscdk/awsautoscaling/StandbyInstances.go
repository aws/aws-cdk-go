package awsautoscaling


// The behavior of instances in standby during an instance refresh.
type StandbyInstances string

const (
	// Terminate instances in standby and launch new ones to replace them.
	StandbyInstances_TERMINATE StandbyInstances = "TERMINATE"
	// Ignore instances in standby.
	StandbyInstances_IGNORE StandbyInstances = "IGNORE"
	// Wait until instances are taken out of standby, then refresh them.
	StandbyInstances_WAIT StandbyInstances = "WAIT"
)

