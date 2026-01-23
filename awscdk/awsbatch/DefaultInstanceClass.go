package awsbatch


// Batch default instances types.
//
// Example:
//   vpc := ec2.NewVpc(this, jsii.String("Vpc"))
//
//   // Use ARM64 instances (e.g., m6g, c6g, r6g, c7g families)
//   // Use ARM64 instances (e.g., m6g, c6g, r6g, c7g families)
//   batch.NewManagedEc2EcsComputeEnvironment(this, jsii.String("Arm64Ec2ComputeEnv"), &ManagedEc2EcsComputeEnvironmentProps{
//   	Vpc: Vpc,
//   	DefaultInstanceClasses: []DefaultInstanceClass{
//   		batch.DefaultInstanceClass_ARM64,
//   	},
//   })
//
//   // Use x86_64 instances (e.g., m6i, c6i, r6i, c7i families)
//   // Use x86_64 instances (e.g., m6i, c6i, r6i, c7i families)
//   batch.NewManagedEc2EcsComputeEnvironment(this, jsii.String("X86_64Ec2ComputeEnv"), &ManagedEc2EcsComputeEnvironmentProps{
//   	Vpc: Vpc,
//   	DefaultInstanceClasses: []DefaultInstanceClass{
//   		batch.DefaultInstanceClass_X86_64,
//   	},
//   })
//
// See: https://docs.aws.amazon.com/batch/latest/userguide/instance-type-compute-table.html
//
type DefaultInstanceClass string

const (
	// x86 based instance types (from the m6i, c6i, r6i, and c7i instance families).
	DefaultInstanceClass_X86_64 DefaultInstanceClass = "X86_64"
	// ARM64 based instance types (from the m6g, c6g, r6g, and c7g instance families).
	DefaultInstanceClass_ARM64 DefaultInstanceClass = "ARM64"
)

