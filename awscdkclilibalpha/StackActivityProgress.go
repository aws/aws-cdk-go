package awscdkclilibalpha


// Supported display modes for stack deployment activity.
// Deprecated.
type StackActivityProgress string

const (
	// Displays a progress bar with only the events for the resource currently being deployed.
	// Deprecated.
	StackActivityProgress_BAR StackActivityProgress = "BAR"
	// Displays complete history with all CloudFormation stack events.
	// Deprecated.
	StackActivityProgress_EVENTS StackActivityProgress = "EVENTS"
)

