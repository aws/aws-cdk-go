package awscdkredshiftalpha


// The type of logs to export.
// Experimental.
type LogExport string

const (
	// Connection log - logs authentication attempts and connections/disconnections.
	// Experimental.
	LogExport_CONNECTION_LOG LogExport = "CONNECTION_LOG"
	// User activity log - logs each query before it's run.
	// Experimental.
	LogExport_USER_ACTIVITY_LOG LogExport = "USER_ACTIVITY_LOG"
	// User log - logs user account changes.
	// Experimental.
	LogExport_USER_LOG LogExport = "USER_LOG"
)

