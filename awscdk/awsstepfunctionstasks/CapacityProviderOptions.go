package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// Capacity provider options.
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
type CapacityProviderOptions interface {
}

// The jsii proxy struct for CapacityProviderOptions
type jsiiProxy_CapacityProviderOptions struct {
	_ byte // padding
}

// Use a custom capacity provider strategy.
//
// You can specify between 1 and 20 capacity providers.
func CapacityProviderOptions_Custom(capacityProviderStrategy *[]*awsecs.CapacityProviderStrategy) CapacityProviderOptions {
	_init_.Initialize()

	if err := validateCapacityProviderOptions_CustomParameters(capacityProviderStrategy); err != nil {
		panic(err)
	}
	var returns CapacityProviderOptions

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.CapacityProviderOptions",
		"custom",
		[]interface{}{capacityProviderStrategy},
		&returns,
	)

	return returns
}

// Use the cluster's default capacity provider strategy.
func CapacityProviderOptions_Default() CapacityProviderOptions {
	_init_.Initialize()

	var returns CapacityProviderOptions

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.CapacityProviderOptions",
		"default",
		nil, // no parameters
		&returns,
	)

	return returns
}

