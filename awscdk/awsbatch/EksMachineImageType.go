package awsbatch


// Maps the image to instance types.
type EksMachineImageType string

const (
	// Tells Batch that this machine image runs on non-GPU instances.
	EksMachineImageType_EKS_AL2 EksMachineImageType = "EKS_AL2"
	// Tells Batch that this machine image runs on GPU instances.
	EksMachineImageType_EKS_AL2_NVIDIA EksMachineImageType = "EKS_AL2_NVIDIA"
	// Tells Batch that this machine image runs on non-GPU AL2023 instances.
	EksMachineImageType_EKS_AL2023 EksMachineImageType = "EKS_AL2023"
	// Tells Batch that this machine image runs on GPU AL2023 instances.
	EksMachineImageType_EKS_AL2023_NVIDIA EksMachineImageType = "EKS_AL2023_NVIDIA"
)

