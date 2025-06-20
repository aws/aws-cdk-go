package awscodebuild


// The compute type of the fleet.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   fleet := codebuild.NewFleet(this, jsii.String("MyFleet"), &FleetProps{
//   	BaseCapacity: jsii.Number(1),
//   	ComputeType: codebuild.FleetComputeType_ATTRIBUTE_BASED,
//   	EnvironmentType: codebuild.EnvironmentType_LINUX_CONTAINER,
//   	ComputeConfiguration: &ComputeConfiguration{
//   		VCpu: jsii.Number(2),
//   		Memory: awscdk.Size_Gibibytes(jsii.Number(4)),
//   		Disk: awscdk.Size_*Gibibytes(jsii.Number(10)),
//   		MachineType: codebuild.MachineType_GENERAL,
//   	},
//   })
//
type MachineType string

const (
	// General purpose compute type.
	MachineType_GENERAL MachineType = "GENERAL"
	// Non-Volatile Memory Express (NVMe) storage optimized compute type.
	MachineType_NVME MachineType = "NVME"
)

