package awseksv2


// The machine image type.
type MachineImageType string

const (
	// Amazon EKS-optimized Linux AMI.
	MachineImageType_AMAZON_LINUX_2 MachineImageType = "AMAZON_LINUX_2"
	// Bottlerocket AMI.
	MachineImageType_BOTTLEROCKET MachineImageType = "BOTTLEROCKET"
)

