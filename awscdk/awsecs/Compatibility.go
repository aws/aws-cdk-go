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
//   // Use custom() option - specify custom capacity provider strategy
//   runTaskWithCustom := tasks.NewEcsRunTask(this, jsii.String("RunTaskWithCustom"), &EcsRunTaskProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	LaunchTarget: tasks.NewEcsFargateLaunchTarget(&EcsFargateLaunchTargetOptions{
//   		PlatformVersion: ecs.FargatePlatformVersion_VERSION1_4,
//   		CapacityProviderOptions: tasks.CapacityProviderOptions_Custom([]CapacityProviderStrategy{
//   			&CapacityProviderStrategy{
//   				CapacityProvider: jsii.String("FARGATE_SPOT"),
//   				Weight: jsii.Number(2),
//   				Base: jsii.Number(1),
//   			},
//   			&CapacityProviderStrategy{
//   				CapacityProvider: jsii.String("FARGATE"),
//   				Weight: jsii.Number(1),
//   			},
//   		}),
//   	}),
//   })
//
//   // Use default() option - uses cluster's default capacity provider strategy
//   runTaskWithDefault := tasks.NewEcsRunTask(this, jsii.String("RunTaskWithDefault"), &EcsRunTaskProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	LaunchTarget: tasks.NewEcsFargateLaunchTarget(&EcsFargateLaunchTargetOptions{
//   		PlatformVersion: ecs.FargatePlatformVersion_VERSION1_4,
//   		CapacityProviderOptions: tasks.CapacityProviderOptions_Default(),
//   	}),
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

