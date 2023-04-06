package awsecs


// Propagate tags from either service or task definition.
type PropagatedTagSource string

const (
	// Propagate tags from service.
	PropagatedTagSource_SERVICE PropagatedTagSource = "SERVICE"
	// Propagate tags from task definition.
	PropagatedTagSource_TASK_DEFINITION PropagatedTagSource = "TASK_DEFINITION"
	// Do not propagate.
	PropagatedTagSource_NONE PropagatedTagSource = "NONE"
)

