package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Configuration for running an ECS task on Fargate.
//
// Example:
//   vpc := ec2.vpc.fromLookup(this, jsii.String("Vpc"), &vpcLookupOptions{
//   	isDefault: jsii.Boolean(true),
//   })
//
//   cluster := ecs.NewCluster(this, jsii.String("FargateCluster"), &clusterProps{
//   	vpc: vpc,
//   })
//
//   taskDefinition := ecs.NewTaskDefinition(this, jsii.String("TD"), &taskDefinitionProps{
//   	memoryMiB: jsii.String("512"),
//   	cpu: jsii.String("256"),
//   	compatibility: ecs.compatibility_FARGATE,
//   })
//
//   containerDefinition := taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("foo/bar")),
//   	memoryLimitMiB: jsii.Number(256),
//   })
//
//   runTask := tasks.NewEcsRunTask(this, jsii.String("RunFargate"), &ecsRunTaskProps{
//   	integrationPattern: sfn.integrationPattern_RUN_JOB,
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   	assignPublicIp: jsii.Boolean(true),
//   	containerOverrides: []containerOverride{
//   		&containerOverride{
//   			containerDefinition: containerDefinition,
//   			environment: []taskEnvironmentVariable{
//   				&taskEnvironmentVariable{
//   					name: jsii.String("SOME_KEY"),
//   					value: sfn.jsonPath.stringAt(jsii.String("$.SomeKey")),
//   				},
//   			},
//   		},
//   	},
//   	launchTarget: tasks.NewEcsFargateLaunchTarget(),
//   })
//
// See: https://docs.aws.amazon.com/AmazonECS/latest/userguide/launch_types.html#launch-type-fargate
//
// Experimental.
type EcsFargateLaunchTarget interface {
	IEcsLaunchTarget
	// Called when the Fargate launch type configured on RunTask.
	// Experimental.
	Bind(_task EcsRunTask, launchTargetOptions *LaunchTargetBindOptions) *EcsLaunchTargetConfig
}

// The jsii proxy struct for EcsFargateLaunchTarget
type jsiiProxy_EcsFargateLaunchTarget struct {
	jsiiProxy_IEcsLaunchTarget
}

// Experimental.
func NewEcsFargateLaunchTarget(options *EcsFargateLaunchTargetOptions) EcsFargateLaunchTarget {
	_init_.Initialize()

	j := jsiiProxy_EcsFargateLaunchTarget{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EcsFargateLaunchTarget",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewEcsFargateLaunchTarget_Override(e EcsFargateLaunchTarget, options *EcsFargateLaunchTargetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EcsFargateLaunchTarget",
		[]interface{}{options},
		e,
	)
}

func (e *jsiiProxy_EcsFargateLaunchTarget) Bind(_task EcsRunTask, launchTargetOptions *LaunchTargetBindOptions) *EcsLaunchTargetConfig {
	var returns *EcsLaunchTargetConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{_task, launchTargetOptions},
		&returns,
	)

	return returns
}

