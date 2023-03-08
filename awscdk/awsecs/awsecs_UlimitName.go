package awsecs


// Type of resource to set a limit on.
// Experimental.
type UlimitName string

const (
	// Experimental.
	UlimitName_CORE UlimitName = "CORE"
	// Experimental.
	UlimitName_CPU UlimitName = "CPU"
	// Experimental.
	UlimitName_DATA UlimitName = "DATA"
	// Experimental.
	UlimitName_FSIZE UlimitName = "FSIZE"
	// Experimental.
	UlimitName_LOCKS UlimitName = "LOCKS"
	// Experimental.
	UlimitName_MEMLOCK UlimitName = "MEMLOCK"
	// Experimental.
	UlimitName_MSGQUEUE UlimitName = "MSGQUEUE"
	// Experimental.
	UlimitName_NICE UlimitName = "NICE"
	// Experimental.
	UlimitName_NOFILE UlimitName = "NOFILE"
	// Experimental.
	UlimitName_NPROC UlimitName = "NPROC"
	// Experimental.
	UlimitName_RSS UlimitName = "RSS"
	// Experimental.
	UlimitName_RTPRIO UlimitName = "RTPRIO"
	// Experimental.
	UlimitName_RTTIME UlimitName = "RTTIME"
	// Experimental.
	UlimitName_SIGPENDING UlimitName = "SIGPENDING"
	// Experimental.
	UlimitName_STACK UlimitName = "STACK"
)

