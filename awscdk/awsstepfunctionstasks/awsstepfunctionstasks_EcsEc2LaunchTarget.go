package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Configuration for running an ECS task on EC2.
//
// Example:
//   vpc := ec2.Vpc_FromLookup(this, jsii.String("Vpc"), &VpcLookupOptions{
//   	IsDefault: jsii.Boolean(true),
//   })
//
//   cluster := ecs.NewCluster(this, jsii.String("Ec2Cluster"), &ClusterProps{
//   	Vpc: Vpc,
//   })
//   cluster.AddCapacity(jsii.String("DefaultAutoScalingGroup"), &AddCapacityOptions{
//   	InstanceType: ec2.NewInstanceType(jsii.String("t2.micro")),
//   	VpcSubnets: &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PUBLIC,
//   	},
//   })
//
//   taskDefinition := ecs.NewTaskDefinition(this, jsii.String("TD"), &TaskDefinitionProps{
//   	Compatibility: ecs.Compatibility_EC2,
//   })
//
//   taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("foo/bar")),
//   	MemoryLimitMiB: jsii.Number(256),
//   })
//
//   runTask := tasks.NewEcsRunTask(this, jsii.String("Run"), &EcsRunTaskProps{
//   	IntegrationPattern: sfn.IntegrationPattern_RUN_JOB,
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	LaunchTarget: tasks.NewEcsEc2LaunchTarget(&EcsEc2LaunchTargetOptions{
//   		PlacementStrategies: []placementStrategy{
//   			ecs.*placementStrategy_SpreadAcrossInstances(),
//   			ecs.*placementStrategy_PackedByCpu(),
//   			ecs.*placementStrategy_Randomly(),
//   		},
//   		PlacementConstraints: []placementConstraint{
//   			ecs.*placementConstraint_MemberOf(jsii.String("blieptuut")),
//   		},
//   	}),
//   })
//
// See: https://docs.aws.amazon.com/AmazonECS/latest/userguide/launch_types.html#launch-type-ec2
//
// Experimental.
type EcsEc2LaunchTarget interface {
	IEcsLaunchTarget
	// Called when the EC2 launch type is configured on RunTask.
	// Experimental.
	Bind(_task EcsRunTask, launchTargetOptions *LaunchTargetBindOptions) *EcsLaunchTargetConfig
}

// The jsii proxy struct for EcsEc2LaunchTarget
type jsiiProxy_EcsEc2LaunchTarget struct {
	jsiiProxy_IEcsLaunchTarget
}

// Experimental.
func NewEcsEc2LaunchTarget(options *EcsEc2LaunchTargetOptions) EcsEc2LaunchTarget {
	_init_.Initialize()

	if err := validateNewEcsEc2LaunchTargetParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsEc2LaunchTarget{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EcsEc2LaunchTarget",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewEcsEc2LaunchTarget_Override(e EcsEc2LaunchTarget, options *EcsEc2LaunchTargetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EcsEc2LaunchTarget",
		[]interface{}{options},
		e,
	)
}

func (e *jsiiProxy_EcsEc2LaunchTarget) Bind(_task EcsRunTask, launchTargetOptions *LaunchTargetBindOptions) *EcsLaunchTargetConfig {
	if err := e.validateBindParameters(_task, launchTargetOptions); err != nil {
		panic(err)
	}
	var returns *EcsLaunchTargetConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{_task, launchTargetOptions},
		&returns,
	)

	return returns
}

