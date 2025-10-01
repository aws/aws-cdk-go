package awscodebuild


// Fleet build machine compute type. Subset of Fleet compatible ComputeType values.
//
// The allocated memory, vCPU count and disk space of the build machine for a
// given compute type are dependent on the environment type.
// Some compute types may also not be available for all environment types.
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
type FleetComputeType string

const (
	// Small compute type.
	//
	// May not be available for all environment types.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-compute-types.html#environment.types
	//
	FleetComputeType_SMALL FleetComputeType = "SMALL"
	// Medium compute type.
	//
	// May not be available for all environment types.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-compute-types.html#environment.types
	//
	FleetComputeType_MEDIUM FleetComputeType = "MEDIUM"
	// Large compute type.
	FleetComputeType_LARGE FleetComputeType = "LARGE"
	// Extra Large compute type.
	//
	// May not be available for all environment types.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-compute-types.html#environment.types
	//
	FleetComputeType_X_LARGE FleetComputeType = "X_LARGE"
	// Extra, Extra Large compute type.
	//
	// May not be available for all environment types.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-compute-types.html#environment.types
	//
	FleetComputeType_X2_LARGE FleetComputeType = "X2_LARGE"
	// Specify the amount of vCPUs, memory, disk space, and the type of machine.
	//
	// AWS CodeBuild will select the cheapest instance that satisfies your specified attributes from `computeConfiguration`.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-compute-types.html#environment-reserved-capacity.types
	//
	FleetComputeType_ATTRIBUTE_BASED FleetComputeType = "ATTRIBUTE_BASED"
	// Specify a specific EC2 instance type to use for compute.
	//
	// You must set `instanceType` on `computeConfiguration`, and optionally set a
	// `disk` size if the provided 64GB is insufficient.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-compute-types.html#environment-reserved-capacity.instance-types
	//
	FleetComputeType_CUSTOM_INSTANCE_TYPE FleetComputeType = "CUSTOM_INSTANCE_TYPE"
)

