package awsec2


// Provides the options for specifying the CPU credit type for burstable EC2 instance types (T2, T3, T3a, etc).
// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-how-to.html
//
type CpuCredits string

const (
	// Standard bursting mode.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-standard-mode.html
	//
	CpuCredits_STANDARD CpuCredits = "STANDARD"
	// Unlimited bursting mode.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-unlimited-mode.html
	//
	CpuCredits_UNLIMITED CpuCredits = "UNLIMITED"
)

