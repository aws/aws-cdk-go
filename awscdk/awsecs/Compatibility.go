package awsecs


// The task launch type compatibility requirement.
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
type Compatibility string

const (
	// The task should specify the EC2 launch type.
	Compatibility_EC2 Compatibility = "EC2"
	// The task should specify the Fargate launch type.
	Compatibility_FARGATE Compatibility = "FARGATE"
	// The task can specify either the EC2 or Fargate launch types.
	Compatibility_EC2_AND_FARGATE Compatibility = "EC2_AND_FARGATE"
	// The task should specify the External launch type.
	Compatibility_EXTERNAL Compatibility = "EXTERNAL"
	// The task should specify the Managed Instances launch type.
	Compatibility_MANAGED_INSTANCES Compatibility = "MANAGED_INSTANCES"
	// The task can specify either the EC2 or Managed Instances launch types.
	Compatibility_EC2_AND_MANAGED_INSTANCES Compatibility = "EC2_AND_MANAGED_INSTANCES"
	// The task can specify either the Fargate or Managed Instances launch types.
	Compatibility_FARGATE_AND_MANAGED_INSTANCES Compatibility = "FARGATE_AND_MANAGED_INSTANCES"
	// The task can specify either the Fargate, EC2 or Managed Instances launch types.
	Compatibility_FARGATE_AND_EC2_AND_MANAGED_INSTANCES Compatibility = "FARGATE_AND_EC2_AND_MANAGED_INSTANCES"
)

