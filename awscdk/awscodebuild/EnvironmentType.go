package awscodebuild


// Build environment type.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   fleet := codebuild.NewFleet(this, jsii.String("MyFleet"), &FleetProps{
//   	BaseCapacity: jsii.Number(1),
//   	ComputeType: codebuild.FleetComputeType_CUSTOM_INSTANCE_TYPE,
//   	EnvironmentType: codebuild.EnvironmentType_LINUX_CONTAINER,
//   	ComputeConfiguration: &ComputeConfiguration{
//   		InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_T3, ec2.InstanceSize_MEDIUM),
//   		// By default, 64 GiB of disk space is included. Any value optionally
//   		// specified here is _incremental_ on top of the included disk space.
//   		Disk: awscdk.Size_Gibibytes(jsii.Number(10)),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-compute-types.html#environment.types
//
type EnvironmentType string

const (
	// ARM container.
	EnvironmentType_ARM_CONTAINER EnvironmentType = "ARM_CONTAINER"
	// Linux container.
	EnvironmentType_LINUX_CONTAINER EnvironmentType = "LINUX_CONTAINER"
	// Linux GPU container.
	EnvironmentType_LINUX_GPU_CONTAINER EnvironmentType = "LINUX_GPU_CONTAINER"
	// Windows Server 2019 container.
	EnvironmentType_WINDOWS_SERVER_2019_CONTAINER EnvironmentType = "WINDOWS_SERVER_2019_CONTAINER"
	// Windows Server 2022 container.
	EnvironmentType_WINDOWS_SERVER_2022_CONTAINER EnvironmentType = "WINDOWS_SERVER_2022_CONTAINER"
	// MacOS ARM container.
	EnvironmentType_MAC_ARM EnvironmentType = "MAC_ARM"
	// Linux EC2.
	EnvironmentType_LINUX_EC2 EnvironmentType = "LINUX_EC2"
	// ARM EC2.
	EnvironmentType_ARM_EC2 EnvironmentType = "ARM_EC2"
	// Windows EC2.
	EnvironmentType_WINDOWS_EC2 EnvironmentType = "WINDOWS_EC2"
)

