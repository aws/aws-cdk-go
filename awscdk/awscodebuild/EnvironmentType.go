package awscodebuild


// Build environment type.
//
// Example:
//   fleet := codebuild.NewFleet(this, jsii.String("Fleet"), &FleetProps{
//   	ComputeType: codebuild.FleetComputeType_MEDIUM,
//   	EnvironmentType: codebuild.EnvironmentType_LINUX_CONTAINER,
//   	BaseCapacity: jsii.Number(1),
//   })
//
//   codebuild.NewProject(this, jsii.String("Project"), &ProjectProps{
//   	Environment: &BuildEnvironment{
//   		Fleet: *Fleet,
//   		BuildImage: codebuild.LinuxBuildImage_STANDARD_7_0(),
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

