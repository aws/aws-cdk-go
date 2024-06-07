package awscodebuild


// Fleet build machine compute type. Subset of Fleet compatible {@link ComputeType} values.
//
// The allocated memory, vCPU count and disk space of the build machine for a
// given compute type are dependent on the environment type.
// Some compute types may also not be available for all environment types.
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
type FleetComputeType string

const (
	// Small compute type.
	//
	// May not be available for all environment types, see
	// {@link https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-compute-types.html#environment.types docs}
	// for more information.
	FleetComputeType_SMALL FleetComputeType = "SMALL"
	// Medium compute type.
	//
	// May not be available for all environment types, see
	// {@link https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-compute-types.html#environment.types docs}
	// for more information.
	FleetComputeType_MEDIUM FleetComputeType = "MEDIUM"
	// Large compute type.
	FleetComputeType_LARGE FleetComputeType = "LARGE"
	// Extra Large compute type.
	//
	// May not be available for all environment types, see
	// {@link https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-compute-types.html#environment.types docs}
	// for more information.
	FleetComputeType_X_LARGE FleetComputeType = "X_LARGE"
	// Extra, Extra Large compute type.
	//
	// May not be available for all environment types, see
	// {@link https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-compute-types.html#environment.types docs}
	// for more information.
	FleetComputeType_X2_LARGE FleetComputeType = "X2_LARGE"
)

