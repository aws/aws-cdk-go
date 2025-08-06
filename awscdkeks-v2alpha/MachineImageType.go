package awscdkeks-v2alpha


// The machine image type.
// Experimental.
type MachineImageType string

const (
	// Amazon EKS-optimized Linux AMI.
	// Experimental.
	MachineImageType_AMAZON_LINUX_2 MachineImageType = "AMAZON_LINUX_2"
	// Bottlerocket AMI.
	// Experimental.
	MachineImageType_BOTTLEROCKET MachineImageType = "BOTTLEROCKET"
)

