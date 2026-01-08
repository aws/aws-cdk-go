package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// The properties used to define an ECS cluster.
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
//   	ContainerOverrides: []ContainerOverride{
//   		&ContainerOverride{
//   			ContainerDefinition: *ContainerDefinition,
//   			Environment: []TaskEnvironmentVariable{
//   				&TaskEnvironmentVariable{
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
type ClusterProps struct {
	// The ec2 capacity to add to the cluster.
	// Default: - no EC2 capacity will be added, you can use `addCapacity` to add capacity later.
	//
	Capacity *AddCapacityOptions `field:"optional" json:"capacity" yaml:"capacity"`
	// The name for the cluster.
	// Default: CloudFormation-generated name.
	//
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// If true CloudWatch Container Insights will be enabled for the cluster.
	// Default: - Container Insights will be disabled for this cluster.
	//
	// Deprecated: See {@link containerInsightsV2 }.
	ContainerInsights *bool `field:"optional" json:"containerInsights" yaml:"containerInsights"`
	// The CloudWatch Container Insights configuration for the cluster.
	// Default: {@link ContainerInsights.DISABLED } This may be overridden by ECS account level settings.
	//
	ContainerInsightsV2 ContainerInsights `field:"optional" json:"containerInsightsV2" yaml:"containerInsightsV2"`
	// The service discovery namespace created in this cluster.
	// Default: - no service discovery namespace created, you can use `addDefaultCloudMapNamespace` to add a
	// default service discovery namespace later.
	//
	DefaultCloudMapNamespace *CloudMapNamespaceOptions `field:"optional" json:"defaultCloudMapNamespace" yaml:"defaultCloudMapNamespace"`
	// Whether to enable Fargate Capacity Providers.
	// Default: false.
	//
	EnableFargateCapacityProviders *bool `field:"optional" json:"enableFargateCapacityProviders" yaml:"enableFargateCapacityProviders"`
	// The execute command configuration for the cluster.
	// Default: - no configuration will be provided.
	//
	ExecuteCommandConfiguration *ExecuteCommandConfiguration `field:"optional" json:"executeCommandConfiguration" yaml:"executeCommandConfiguration"`
	// Encryption configuration for ECS Managed storage.
	// Default: - no encryption will be applied.
	//
	ManagedStorageConfiguration *ManagedStorageConfiguration `field:"optional" json:"managedStorageConfiguration" yaml:"managedStorageConfiguration"`
	// The VPC where your ECS instances will be running or your ENIs will be deployed.
	// Default: - creates a new VPC with two AZs.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

