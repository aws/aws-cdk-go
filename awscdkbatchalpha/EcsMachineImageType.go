// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha


// Maps the image to instance types.
// Experimental.
type EcsMachineImageType string

const (
	// Tells Batch that this machine image runs on non-GPU instances.
	// Experimental.
	EcsMachineImageType_ECS_AL2 EcsMachineImageType = "ECS_AL2"
	// Tells Batch that this machine image runs on GPU instances.
	// Experimental.
	EcsMachineImageType_ECS_AL2_NVIDIA EcsMachineImageType = "ECS_AL2_NVIDIA"
)

