package awsbatch


// Maps the image to instance types.
type EksMachineImageType string

const (
	// Tells Batch that this machine image runs on non-GPU instances.
	EksMachineImageType_EKS_AL2 EksMachineImageType = "EKS_AL2"
	// Tells Batch that this machine image runs on GPU instances.
	EksMachineImageType_EKS_AL2_NVIDIA EksMachineImageType = "EKS_AL2_NVIDIA"
)

