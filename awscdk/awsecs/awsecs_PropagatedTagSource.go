package awsecs


// Propagate tags from either service or task definition.
// Experimental.
type PropagatedTagSource string

const (
	// Propagate tags from service.
	// Experimental.
	PropagatedTagSource_SERVICE PropagatedTagSource = "SERVICE"
	// Propagate tags from task definition.
	// Experimental.
	PropagatedTagSource_TASK_DEFINITION PropagatedTagSource = "TASK_DEFINITION"
	// Do not propagate.
	// Experimental.
	PropagatedTagSource_NONE PropagatedTagSource = "NONE"
)

