package awsec2


// Identifies an instance's CPU architecture.
type InstanceArchitecture string

const (
	// ARM64 architecture.
	InstanceArchitecture_ARM_64 InstanceArchitecture = "ARM_64"
	// x86-64 architecture.
	InstanceArchitecture_X86_64 InstanceArchitecture = "X86_64"
)

