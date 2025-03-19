package awsecs


// Propagate tags from either service or task definition.
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
//   	PropagatedTagSource: ecs.PropagatedTagSource_TASK_DEFINITION,
//   })
//
type PropagatedTagSource string

const (
	// Propagate tags from service.
	PropagatedTagSource_SERVICE PropagatedTagSource = "SERVICE"
	// Propagate tags from task definition.
	PropagatedTagSource_TASK_DEFINITION PropagatedTagSource = "TASK_DEFINITION"
	// Do not propagate.
	PropagatedTagSource_NONE PropagatedTagSource = "NONE"
)

