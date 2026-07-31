package awsautoscaling


// The behavior of instances that are protected from scale in during an instance refresh.
type ScaleInProtectedInstances string

const (
	// Refresh instances that are protected from scale in.
	ScaleInProtectedInstances_REFRESH ScaleInProtectedInstances = "REFRESH"
	// Ignore instances that are protected from scale in.
	ScaleInProtectedInstances_IGNORE ScaleInProtectedInstances = "IGNORE"
	// Wait until instances are no longer protected from scale in, then refresh them.
	ScaleInProtectedInstances_WAIT ScaleInProtectedInstances = "WAIT"
)

