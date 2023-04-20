package awsecs


// The scope for the Docker volume that determines its lifecycle.
//
// Docker volumes that are scoped to a task are automatically provisioned when the task starts and destroyed when the task stops.
// Docker volumes that are scoped as shared persist after the task stops.
// Experimental.
type Scope string

const (
	// Docker volumes that are scoped to a task are automatically provisioned when the task starts and destroyed when the task stops.
	// Experimental.
	Scope_TASK Scope = "TASK"
	// Docker volumes that are scoped as shared persist after the task stops.
	// Experimental.
	Scope_SHARED Scope = "SHARED"
)

