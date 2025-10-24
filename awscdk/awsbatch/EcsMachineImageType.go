package awsbatch


// Maps the image to instance types.
//
// Example:
//   var vpc IVpc
//
//
//   batch.NewManagedEc2EcsComputeEnvironment(this, jsii.String("myEc2ComputeEnv"), &ManagedEc2EcsComputeEnvironmentProps{
//   	Vpc: Vpc,
//   	Images: []EcsMachineImage{
//   		&EcsMachineImage{
//   			ImageType: batch.EcsMachineImageType_ECS_AL2023,
//   		},
//   	},
//   })
//
type EcsMachineImageType string

const (
	// Tells Batch that this machine image runs on non-GPU AL2 instances.
	EcsMachineImageType_ECS_AL2 EcsMachineImageType = "ECS_AL2"
	// Tells Batch that this machine image runs on non-GPU AL2023 instances.
	//
	// Amazon Linux 2023 does not support A1 instances.
	EcsMachineImageType_ECS_AL2023 EcsMachineImageType = "ECS_AL2023"
	// Tells Batch that this machine image runs on GPU instances.
	EcsMachineImageType_ECS_AL2_NVIDIA EcsMachineImageType = "ECS_AL2_NVIDIA"
)

