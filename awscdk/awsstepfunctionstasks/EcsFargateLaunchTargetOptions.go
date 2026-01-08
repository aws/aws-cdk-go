package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// Properties to define an ECS service.
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
type EcsFargateLaunchTargetOptions struct {
	// Refers to a specific runtime environment for Fargate task infrastructure.
	//
	// Fargate platform version is a combination of the kernel and container runtime versions.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html
	//
	PlatformVersion awsecs.FargatePlatformVersion `field:"required" json:"platformVersion" yaml:"platformVersion"`
	// The capacity provider options to use for the task.
	//
	// This property allows you to set the capacity provider strategy for the task.
	//
	// If you want to set the capacity provider strategy for the task, specify
	// `CapacityProviderOptions.custom()`. This is required to use the FARGATE_SPOT
	// capacity provider.
	//
	// If you want to use the cluster's default capacity provider strategy, specify
	// `CapacityProviderOptions.default()`.
	// Default: - 'FARGATE' LaunchType running tasks on AWS Fargate On-Demand
	// infrastructure is used without the capacity provider strategy.
	//
	CapacityProviderOptions CapacityProviderOptions `field:"optional" json:"capacityProviderOptions" yaml:"capacityProviderOptions"`
}

