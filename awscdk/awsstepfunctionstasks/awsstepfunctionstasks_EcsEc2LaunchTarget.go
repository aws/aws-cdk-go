package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Configuration for running an ECS task on EC2.
//
// Example:
//   vpc := ec2.vpc.fromLookup(this, jsii.String("Vpc"), &vpcLookupOptions{
//   	isDefault: jsii.Boolean(true),
//   })
//
//   cluster := ecs.NewCluster(this, jsii.String("Ec2Cluster"), &clusterProps{
//   	vpc: vpc,
//   })
//   cluster.addCapacity(jsii.String("DefaultAutoScalingGroup"), &addCapacityOptions{
//   	instanceType: ec2.NewInstanceType(jsii.String("t2.micro")),
//   	vpcSubnets: &subnetSelection{
//   		subnetType: ec2.subnetType_PUBLIC,
//   	},
//   })
//
//   taskDefinition := ecs.NewTaskDefinition(this, jsii.String("TD"), &taskDefinitionProps{
//   	compatibility: ecs.compatibility_EC2,
//   })
//
//   taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("foo/bar")),
//   	memoryLimitMiB: jsii.Number(256),
//   })
//
//   runTask := tasks.NewEcsRunTask(this, jsii.String("Run"), &ecsRunTaskProps{
//   	integrationPattern: sfn.integrationPattern_RUN_JOB,
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   	launchTarget: tasks.NewEcsEc2LaunchTarget(&ecsEc2LaunchTargetOptions{
//   		placementStrategies: []placementStrategy{
//   			ecs.*placementStrategy.spreadAcrossInstances(),
//   			ecs.*placementStrategy.packedByCpu(),
//   			ecs.*placementStrategy.randomly(),
//   		},
//   		placementConstraints: []placementConstraint{
//   			ecs.*placementConstraint.memberOf(jsii.String("blieptuut")),
//   		},
//   	}),
//   })
//
// See: https://docs.aws.amazon.com/AmazonECS/latest/userguide/launch_types.html#launch-type-ec2
//
type EcsEc2LaunchTarget interface {
	IEcsLaunchTarget
	// Called when the EC2 launch type is configured on RunTask.
	Bind(_task EcsRunTask, launchTargetOptions *LaunchTargetBindOptions) *EcsLaunchTargetConfig
}

// The jsii proxy struct for EcsEc2LaunchTarget
type jsiiProxy_EcsEc2LaunchTarget struct {
	jsiiProxy_IEcsLaunchTarget
}

func NewEcsEc2LaunchTarget(options *EcsEc2LaunchTargetOptions) EcsEc2LaunchTarget {
	_init_.Initialize()

	if err := validateNewEcsEc2LaunchTargetParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsEc2LaunchTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.EcsEc2LaunchTarget",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewEcsEc2LaunchTarget_Override(e EcsEc2LaunchTarget, options *EcsEc2LaunchTargetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.EcsEc2LaunchTarget",
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

