package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// The properties used to define an ECS cluster.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("Vpc"), &VpcProps{
//   	MaxAzs: jsii.Number(1),
//   })
//   cluster := ecs.NewCluster(this, jsii.String("EcsCluster"), &ClusterProps{
//   	Vpc: Vpc,
//   })
//   taskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"), &FargateTaskDefinitionProps{
//   	MemoryLimitMiB: jsii.Number(512),
//   	Cpu: jsii.Number(256),
//   })
//   taskDefinition.AddContainer(jsii.String("WebContainer"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   })
//   awscdk.Tags_Of(taskDefinition).Add(jsii.String("my-tag"), jsii.String("my-tag-value"))
//   scheduledFargateTask := ecsPatterns.NewScheduledFargateTask(this, jsii.String("ScheduledFargateTask"), &ScheduledFargateTaskProps{
//   	Cluster: Cluster,
//   	TaskDefinition: taskDefinition,
//   	Schedule: appscaling.Schedule_Expression(jsii.String("rate(1 minute)")),
//   	PropagateTags: ecs.PropagatedTagSource_TASK_DEFINITION,
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
	ContainerInsights *bool `field:"optional" json:"containerInsights" yaml:"containerInsights"`
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

