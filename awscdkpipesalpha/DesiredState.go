package awscdkpipesalpha


// The state the pipe should be in.
// Experimental.
type DesiredState string

const (
	// The pipe should be running.
	// Experimental.
	DesiredState_RUNNING DesiredState = "RUNNING"
	// The pipe should be stopped.
	// Experimental.
	DesiredState_STOPPED DesiredState = "STOPPED"
)

