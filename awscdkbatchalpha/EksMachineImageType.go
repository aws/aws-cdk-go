// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha


// Maps the image to instance types.
// Experimental.
type EksMachineImageType string

const (
	// Tells Batch that this machine image runs on non-GPU instances.
	// Experimental.
	EksMachineImageType_EKS_AL2 EksMachineImageType = "EKS_AL2"
	// Tells Batch that this machine image runs on GPU instances.
	// Experimental.
	EksMachineImageType_EKS_AL2_NVIDIA EksMachineImageType = "EKS_AL2_NVIDIA"
)

