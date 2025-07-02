package awsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// A ManagedComputeEnvironment that uses ECS orchestration on EC2 instances.
type IManagedEc2EcsComputeEnvironment interface {
	IManagedComputeEnvironment
	// Add an instance class to this compute environment.
	AddInstanceClass(instanceClass awsec2.InstanceClass)
	// Add an instance type to this compute environment.
	AddInstanceType(instanceType awsec2.InstanceType)
	// The allocation strategy to use if not enough instances of the best fitting instance type can be allocated.
	// Default: - `BEST_FIT_PROGRESSIVE` if not using Spot instances,
	// `SPOT_CAPACITY_OPTIMIZED` if using Spot instances.
	//
	AllocationStrategy() AllocationStrategy
	// Configure which AMIs this Compute Environment can launch.
	//
	// Leave this `undefined` to allow Batch to choose the latest AMIs it supports for each instance that it launches.
	// Default: - ECS_AL2 compatible AMI ids for non-GPU instances, ECS_AL2_NVIDIA compatible AMI ids for GPU instances.
	//
	Images() *[]*EcsMachineImage
	// The instance classes that this Compute Environment can launch.
	//
	// Which one is chosen depends on the `AllocationStrategy` used.
	// Batch will automatically choose the size.
	InstanceClasses() *[]awsec2.InstanceClass
	// The execution Role that instances launched by this Compute Environment will use.
	// Default: - a role will be created.
	//
	InstanceRole() awsiam.IRole
	// The instance types that this Compute Environment can launch.
	//
	// Which one is chosen depends on the `AllocationStrategy` used.
	InstanceTypes() *[]awsec2.InstanceType
	// The Launch Template that this Compute Environment will use to provision EC2 Instances.
	//
	// *Note*: if `securityGroups` is specified on both your
	// launch template and this Compute Environment, **the
	// `securityGroup`s on the Compute Environment override the
	// ones on the launch template.
	// Default: no launch template.
	//
	LaunchTemplate() awsec2.ILaunchTemplate
	// The minimum vCPUs that an environment should maintain, even if the compute environment is DISABLED.
	// Default: 0.
	//
	MinvCpus() *float64
	// The EC2 placement group to associate with your compute resources.
	//
	// If you intend to submit multi-node parallel jobs to this Compute Environment,
	// you should consider creating a cluster placement group and associate it with your compute resources.
	// This keeps your multi-node parallel job on a logical grouping of instances
	// within a single Availability Zone with high network flow potential.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html
	//
	// Default: - no placement group.
	//
	PlacementGroup() awsec2.IPlacementGroup
	// The maximum percentage that a Spot Instance price can be when compared with the On-Demand price for that instance type before instances are launched.
	//
	// For example, if your maximum percentage is 20%, the Spot price must be
	// less than 20% of the current On-Demand price for that Instance.
	// You always pay the lowest market price and never more than your maximum percentage.
	// For most use cases, Batch recommends leaving this field empty.
	// Default: - 100%.
	//
	SpotBidPercentage() *float64
	// The service-linked role that Spot Fleet needs to launch instances on your behalf.
	// See: https://docs.aws.amazon.com/batch/latest/userguide/spot_fleet_IAM_role.html
	//
	// Default: - a new Role will be created.
	//
	SpotFleetRole() awsiam.IRole
	// Whether or not to use batch's optimal instance type.
	//
	// The optimal instance type is equivalent to adding the
	// C4, M4, and R4 instance classes. You can specify other instance classes
	// (of the same architecture) in addition to the optimal instance classes.
	// Default: true.
	//
	UseOptimalInstanceClasses() *bool
}

// The jsii proxy for IManagedEc2EcsComputeEnvironment
type jsiiProxy_IManagedEc2EcsComputeEnvironment struct {
	jsiiProxy_IManagedComputeEnvironment
}

func (i *jsiiProxy_IManagedEc2EcsComputeEnvironment) AddInstanceClass(instanceClass awsec2.InstanceClass) {
	if err := i.validateAddInstanceClassParameters(instanceClass); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addInstanceClass",
		[]interface{}{instanceClass},
	)
}

func (i *jsiiProxy_IManagedEc2EcsComputeEnvironment) AddInstanceType(instanceType awsec2.InstanceType) {
	if err := i.validateAddInstanceTypeParameters(instanceType); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addInstanceType",
		[]interface{}{instanceType},
	)
}

func (j *jsiiProxy_IManagedEc2EcsComputeEnvironment) AllocationStrategy() AllocationStrategy {
	var returns AllocationStrategy
	_jsii_.Get(
		j,
		"allocationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedEc2EcsComputeEnvironment) Images() *[]*EcsMachineImage {
	var returns *[]*EcsMachineImage
	_jsii_.Get(
		j,
		"images",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedEc2EcsComputeEnvironment) InstanceClasses() *[]awsec2.InstanceClass {
	var returns *[]awsec2.InstanceClass
	_jsii_.Get(
		j,
		"instanceClasses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedEc2EcsComputeEnvironment) InstanceRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"instanceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedEc2EcsComputeEnvironment) InstanceTypes() *[]awsec2.InstanceType {
	var returns *[]awsec2.InstanceType
	_jsii_.Get(
		j,
		"instanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedEc2EcsComputeEnvironment) LaunchTemplate() awsec2.ILaunchTemplate {
	var returns awsec2.ILaunchTemplate
	_jsii_.Get(
		j,
		"launchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedEc2EcsComputeEnvironment) MinvCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minvCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedEc2EcsComputeEnvironment) PlacementGroup() awsec2.IPlacementGroup {
	var returns awsec2.IPlacementGroup
	_jsii_.Get(
		j,
		"placementGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedEc2EcsComputeEnvironment) SpotBidPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotBidPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedEc2EcsComputeEnvironment) SpotFleetRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"spotFleetRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedEc2EcsComputeEnvironment) UseOptimalInstanceClasses() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"useOptimalInstanceClasses",
		&returns,
	)
	return returns
}

