package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Configuration for running an ECS task on Fargate.
//
// Example:
//   vpc := ec2.Vpc_FromLookup(this, jsii.String("Vpc"), &VpcLookupOptions{
//   	IsDefault: jsii.Boolean(true),
//   })
//
//   cluster := ecs.NewCluster(this, jsii.String("FargateCluster"), &ClusterProps{
//   	Vpc: Vpc,
//   })
//
//   taskDefinition := ecs.NewTaskDefinition(this, jsii.String("TD"), &TaskDefinitionProps{
//   	MemoryMiB: jsii.String("512"),
//   	Cpu: jsii.String("256"),
//   	Compatibility: ecs.Compatibility_FARGATE,
//   })
//
//   containerDefinition := taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("foo/bar")),
//   	MemoryLimitMiB: jsii.Number(256),
//   })
//
//   runTask := tasks.NewEcsRunTask(this, jsii.String("RunFargate"), &EcsRunTaskProps{
//   	IntegrationPattern: sfn.IntegrationPattern_RUN_JOB,
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	AssignPublicIp: jsii.Boolean(true),
//   	ContainerOverrides: []containerOverride{
//   		&containerOverride{
//   			ContainerDefinition: *ContainerDefinition,
//   			Environment: []taskEnvironmentVariable{
//   				&taskEnvironmentVariable{
//   					Name: jsii.String("SOME_KEY"),
//   					Value: sfn.JsonPath_StringAt(jsii.String("$.SomeKey")),
//   				},
//   			},
//   		},
//   	},
//   	LaunchTarget: tasks.NewEcsFargateLaunchTarget(),
//   })
//
// See: https://docs.aws.amazon.com/AmazonECS/latest/userguide/launch_types.html#launch-type-fargate
//
type EcsFargateLaunchTarget interface {
	IEcsLaunchTarget
	// Called when the Fargate launch type configured on RunTask.
	Bind(_task EcsRunTask, launchTargetOptions *LaunchTargetBindOptions) *EcsLaunchTargetConfig
}

// The jsii proxy struct for EcsFargateLaunchTarget
type jsiiProxy_EcsFargateLaunchTarget struct {
	jsiiProxy_IEcsLaunchTarget
}

func NewEcsFargateLaunchTarget(options *EcsFargateLaunchTargetOptions) EcsFargateLaunchTarget {
	_init_.Initialize()

	if err := validateNewEcsFargateLaunchTargetParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsFargateLaunchTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.EcsFargateLaunchTarget",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewEcsFargateLaunchTarget_Override(e EcsFargateLaunchTarget, options *EcsFargateLaunchTargetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.EcsFargateLaunchTarget",
		[]interface{}{options},
		e,
	)
}

func (e *jsiiProxy_EcsFargateLaunchTarget) Bind(_task EcsRunTask, launchTargetOptions *LaunchTargetBindOptions) *EcsLaunchTargetConfig {
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

