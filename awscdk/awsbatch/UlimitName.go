package awsbatch


// The resources to be limited.
type UlimitName string

const (
	// max core dump file size.
	UlimitName_CORE UlimitName = "CORE"
	// max cpu time (seconds) for a process.
	UlimitName_CPU UlimitName = "CPU"
	// max data segment size.
	UlimitName_DATA UlimitName = "DATA"
	// max file size.
	UlimitName_FSIZE UlimitName = "FSIZE"
	// max number of file locks.
	UlimitName_LOCKS UlimitName = "LOCKS"
	// max locked memory.
	UlimitName_MEMLOCK UlimitName = "MEMLOCK"
	// max POSIX message queue size.
	UlimitName_MSGQUEUE UlimitName = "MSGQUEUE"
	// max nice value for any process this user is running.
	UlimitName_NICE UlimitName = "NICE"
	// maximum number of open file descriptors.
	UlimitName_NOFILE UlimitName = "NOFILE"
	// maximum number of processes.
	UlimitName_NPROC UlimitName = "NPROC"
	// size of the process' resident set (in pages).
	UlimitName_RSS UlimitName = "RSS"
	// max realtime priority.
	UlimitName_RTPRIO UlimitName = "RTPRIO"
	// timeout for realtime tasks.
	UlimitName_RTTIME UlimitName = "RTTIME"
	// max number of pending signals.
	UlimitName_SIGPENDING UlimitName = "SIGPENDING"
	// max stack size (in bytes).
	UlimitName_STACK UlimitName = "STACK"
)

