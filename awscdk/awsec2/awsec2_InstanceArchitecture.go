package awsec2


// Identifies an instance's CPU architecture.
// Experimental.
type InstanceArchitecture string

const (
	// ARM64 architecture.
	// Experimental.
	InstanceArchitecture_ARM_64 InstanceArchitecture = "ARM_64"
	// x86-64 architecture.
	// Experimental.
	InstanceArchitecture_X86_64 InstanceArchitecture = "X86_64"
)

