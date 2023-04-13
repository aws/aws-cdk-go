// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha


// The resources to be limited.
// Experimental.
type UlimitName string

const (
	// max core dump file size.
	// Experimental.
	UlimitName_CORE UlimitName = "CORE"
	// max cpu time (seconds) for a process.
	// Experimental.
	UlimitName_CPU UlimitName = "CPU"
	// max data segment size.
	// Experimental.
	UlimitName_DATA UlimitName = "DATA"
	// max file size.
	// Experimental.
	UlimitName_FSIZE UlimitName = "FSIZE"
	// max number of file locks.
	// Experimental.
	UlimitName_LOCKS UlimitName = "LOCKS"
	// max locked memory.
	// Experimental.
	UlimitName_MEMLOCK UlimitName = "MEMLOCK"
	// max POSIX message queue size.
	// Experimental.
	UlimitName_MSGQUEUE UlimitName = "MSGQUEUE"
	// max nice value for any process this user is running.
	// Experimental.
	UlimitName_NICE UlimitName = "NICE"
	// maximum number of open file descriptors.
	// Experimental.
	UlimitName_NOFILE UlimitName = "NOFILE"
	// maximum number of processes.
	// Experimental.
	UlimitName_NPROC UlimitName = "NPROC"
	// size of the process' resident set (in pages).
	// Experimental.
	UlimitName_RSS UlimitName = "RSS"
	// max realtime priority.
	// Experimental.
	UlimitName_RTPRIO UlimitName = "RTPRIO"
	// timeout for realtime tasks.
	// Experimental.
	UlimitName_RTTIME UlimitName = "RTTIME"
	// max number of pending signals.
	// Experimental.
	UlimitName_SIGPENDING UlimitName = "SIGPENDING"
	// max stack size (in bytes).
	// Experimental.
	UlimitName_STACK UlimitName = "STACK"
)

