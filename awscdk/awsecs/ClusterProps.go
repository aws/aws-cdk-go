package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// The properties used to define an ECS cluster.
//
// Example:
//   var vpc vpc
//
//
//   cluster := ecs.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
//   	Vpc: Vpc,
//   })
//
//   autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
//   	Vpc: Vpc,
//   	InstanceType: ec2.NewInstanceType(jsii.String("t2.micro")),
//   	MachineImage: ecs.EcsOptimizedImage_AmazonLinux2(),
//   	MinCapacity: jsii.Number(0),
//   	MaxCapacity: jsii.Number(100),
//   })
//
//   capacityProvider := ecs.NewAsgCapacityProvider(this, jsii.String("AsgCapacityProvider"), &AsgCapacityProviderProps{
//   	AutoScalingGroup: AutoScalingGroup,
//   	InstanceWarmupPeriod: jsii.Number(300),
//   })
//   cluster.AddAsgCapacityProvider(capacityProvider)
//
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//
//   taskDefinition.AddContainer(jsii.String("web"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	MemoryReservationMiB: jsii.Number(256),
//   })
//
//   ecs.NewEc2Service(this, jsii.String("EC2Service"), &Ec2ServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	CapacityProviderStrategies: []capacityProviderStrategy{
//   		&capacityProviderStrategy{
//   			CapacityProvider: capacityProvider.CapacityProviderName,
//   			Weight: jsii.Number(1),
//   		},
//   	},
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
	// The VPC where your ECS instances will be running or your ENIs will be deployed.
	// Default: - creates a new VPC with two AZs.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

