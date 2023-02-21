package awsecs


// The process namespace to use for the containers in the task.
type PidMode string

const (
	// If host is specified, then all containers within the tasks that specified the host PID mode on the same container instance share the same process namespace with the host Amazon EC2 instance.
	PidMode_HOST PidMode = "HOST"
	// If task is specified, all containers within the specified task share the same process namespace.
	PidMode_TASK PidMode = "TASK"
)

