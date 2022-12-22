package awsecs


// The IPC resource namespace to use for the containers in the task.
// Experimental.
type IpcMode string

const (
	// If none is specified, then IPC resources within the containers of a task are private and not shared with other containers in a task or on the container instance.
	// Experimental.
	IpcMode_NONE IpcMode = "NONE"
	// If host is specified, then all containers within the tasks that specified the host IPC mode on the same container instance share the same IPC resources with the host Amazon EC2 instance.
	// Experimental.
	IpcMode_HOST IpcMode = "HOST"
	// If task is specified, all containers within the specified task share the same IPC resources.
	// Experimental.
	IpcMode_TASK IpcMode = "TASK"
)

