package awscdkclilibalpha


// Supported display modes for stack deployment activity.
// Experimental.
type StackActivityProgress string

const (
	// Displays a progress bar with only the events for the resource currently being deployed.
	// Experimental.
	StackActivityProgress_BAR StackActivityProgress = "BAR"
	// Displays complete history with all CloudFormation stack events.
	// Experimental.
	StackActivityProgress_EVENTS StackActivityProgress = "EVENTS"
)

