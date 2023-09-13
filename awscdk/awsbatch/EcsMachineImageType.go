package awsbatch


// Maps the image to instance types.
type EcsMachineImageType string

const (
	// Tells Batch that this machine image runs on non-GPU instances.
	EcsMachineImageType_ECS_AL2 EcsMachineImageType = "ECS_AL2"
	// Tells Batch that this machine image runs on GPU instances.
	EcsMachineImageType_ECS_AL2_NVIDIA EcsMachineImageType = "ECS_AL2_NVIDIA"
)

