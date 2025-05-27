package awsecspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// The properties for the NetworkMultipleTargetGroupsFargateService service.
//
// Example:
//   // Two network load balancers, each with their own listener and target group.
//   var cluster cluster
//
//   loadBalancedFargateService := ecsPatterns.NewNetworkMultipleTargetGroupsFargateService(this, jsii.String("Service"), &NetworkMultipleTargetGroupsFargateServiceProps{
//   	Cluster: Cluster,
//   	MemoryLimitMiB: jsii.Number(512),
//   	TaskImageOptions: &NetworkLoadBalancedTaskImageProps{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   	LoadBalancers: []networkLoadBalancerProps{
//   		&networkLoadBalancerProps{
//   			Name: jsii.String("lb1"),
//   			Listeners: []networkListenerProps{
//   				&networkListenerProps{
//   					Name: jsii.String("listener1"),
//   				},
//   			},
//   		},
//   		&networkLoadBalancerProps{
//   			Name: jsii.String("lb2"),
//   			Listeners: []*networkListenerProps{
//   				&networkListenerProps{
//   					Name: jsii.String("listener2"),
//   				},
//   			},
//   		},
//   	},
//   	TargetGroups: []networkTargetProps{
//   		&networkTargetProps{
//   			ContainerPort: jsii.Number(80),
//   			Listener: jsii.String("listener1"),
//   		},
//   		&networkTargetProps{
//   			ContainerPort: jsii.Number(90),
//   			Listener: jsii.String("listener2"),
//   		},
//   	},
//   	MinHealthyPercent: jsii.Number(100),
//   	MaxHealthyPercent: jsii.Number(200),
//   })
//
type NetworkMultipleTargetGroupsFargateServiceProps struct {
	// The options for configuring an Amazon ECS service to use service discovery.
	// Default: - AWS Cloud Map service discovery is not enabled.
	//
	CloudMapOptions *awsecs.CloudMapOptions `field:"optional" json:"cloudMapOptions" yaml:"cloudMapOptions"`
	// The name of the cluster that hosts the service.
	//
	// If a cluster is specified, the vpc construct should be omitted. Alternatively, you can omit both cluster and vpc.
	// Default: - create a new cluster; if both cluster and vpc are omitted, a new VPC will be created for you.
	//
	Cluster awsecs.ICluster `field:"optional" json:"cluster" yaml:"cluster"`
	// The desired number of instantiations of the task definition to keep running on the service.
	//
	// The minimum value is 1.
	// Default: - The default is 1 for all new services and uses the existing service's desired count
	// when updating an existing service.
	//
	DesiredCount *float64 `field:"optional" json:"desiredCount" yaml:"desiredCount"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	// Default: false.
	//
	EnableECSManagedTags *bool `field:"optional" json:"enableECSManagedTags" yaml:"enableECSManagedTags"`
	// Whether ECS Exec should be enabled.
	// Default: - false.
	//
	EnableExecuteCommand *bool `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	// Default: - defaults to 60 seconds if at least one load balancer is in-use and it is not already set.
	//
	HealthCheckGracePeriod awscdk.Duration `field:"optional" json:"healthCheckGracePeriod" yaml:"healthCheckGracePeriod"`
	// The network load balancer that will serve traffic to the service.
	// Default: - a new load balancer with a listener will be created.
	//
	LoadBalancers *[]*NetworkLoadBalancerProps `field:"optional" json:"loadBalancers" yaml:"loadBalancers"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Tags can only be propagated to the tasks within the service during service creation.
	// Default: - none.
	//
	PropagateTags awsecs.PropagatedTagSource `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// Name of the service.
	// Default: - CloudFormation-generated name.
	//
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// Properties to specify NLB target groups.
	// Default: - default portMapping registered as target group and attached to the first defined listener.
	//
	TargetGroups *[]*NetworkTargetProps `field:"optional" json:"targetGroups" yaml:"targetGroups"`
	// The properties required to create a new task definition.
	//
	// Only one of TaskDefinition or TaskImageOptions must be specified.
	// Default: - none.
	//
	TaskImageOptions *NetworkLoadBalancedTaskImageProps `field:"optional" json:"taskImageOptions" yaml:"taskImageOptions"`
	// The VPC where the container instances will be launched or the elastic network interfaces (ENIs) will be deployed.
	//
	// If a vpc is specified, the cluster construct should be omitted. Alternatively, you can omit both vpc and cluster.
	// Default: - uses the VPC defined in the cluster or creates a new VPC.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// The number of cpu units used by the task.
	//
	// Valid values, which determines your range of valid values for the memory parameter:
	//
	// 256 (.25 vCPU) - Available memory values: 0.5GB, 1GB, 2GB
	//
	// 512 (.5 vCPU) - Available memory values: 1GB, 2GB, 3GB, 4GB
	//
	// 1024 (1 vCPU) - Available memory values: 2GB, 3GB, 4GB, 5GB, 6GB, 7GB, 8GB
	//
	// 2048 (2 vCPU) - Available memory values: Between 4GB and 16GB in 1GB increments
	//
	// 4096 (4 vCPU) - Available memory values: Between 8GB and 30GB in 1GB increments
	//
	// 8192 (8 vCPU) - Available memory values: Between 16GB and 60GB in 4GB increments
	//
	// 16384 (16 vCPU) - Available memory values: Between 32GB and 120GB in 8GB increments
	//
	// This default is set in the underlying FargateTaskDefinition construct.
	// Default: 256.
	//
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// The amount (in GiB) of ephemeral storage to be allocated to the task.
	//
	// The minimum supported value is `21` GiB and the maximum supported value is `200` GiB.
	//
	// Only supported in Fargate platform version 1.4.0 or later.
	// Default: Undefined, in which case, the task will receive 20GiB ephemeral storage.
	//
	EphemeralStorageGiB *float64 `field:"optional" json:"ephemeralStorageGiB" yaml:"ephemeralStorageGiB"`
	// The amount (in MiB) of memory used by the task.
	//
	// This field is required and you must use one of the following values, which determines your range of valid values
	// for the cpu parameter:
	//
	// 512 (0.5 GB), 1024 (1 GB), 2048 (2 GB) - Available cpu values: 256 (.25 vCPU)
	//
	// 1024 (1 GB), 2048 (2 GB), 3072 (3 GB), 4096 (4 GB) - Available cpu values: 512 (.5 vCPU)
	//
	// 2048 (2 GB), 3072 (3 GB), 4096 (4 GB), 5120 (5 GB), 6144 (6 GB), 7168 (7 GB), 8192 (8 GB) - Available cpu values: 1024 (1 vCPU)
	//
	// Between 4096 (4 GB) and 16384 (16 GB) in increments of 1024 (1 GB) - Available cpu values: 2048 (2 vCPU)
	//
	// Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB) - Available cpu values: 4096 (4 vCPU)
	//
	// Between 16384 (16 GB) and 61440 (60 GB) in increments of 4096 (4 GB) - Available cpu values: 8192 (8 vCPU)
	//
	// Between 32768 (32 GB) and 122880 (120 GB) in increments of 8192 (8 GB) - Available cpu values: 16384 (16 vCPU)
	//
	// This default is set in the underlying FargateTaskDefinition construct.
	// Default: 512.
	//
	MemoryLimitMiB *float64 `field:"optional" json:"memoryLimitMiB" yaml:"memoryLimitMiB"`
	// The platform version on which to run your service.
	//
	// If one is not specified, the LATEST platform version is used by default. For more information, see
	// [AWS Fargate Platform Versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html)
	// in the Amazon Elastic Container Service Developer Guide.
	// Default: Latest.
	//
	PlatformVersion awsecs.FargatePlatformVersion `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// The runtime platform of the task definition.
	// Default: - If the property is undefined, `operatingSystemFamily` is LINUX and `cpuArchitecture` is X86_64.
	//
	RuntimePlatform *awsecs.RuntimePlatform `field:"optional" json:"runtimePlatform" yaml:"runtimePlatform"`
	// The task definition to use for tasks in the service. TaskDefinition or TaskImageOptions must be specified, but not both.
	//
	// [disable-awslint:ref-via-interface].
	// Default: - none.
	//
	TaskDefinition awsecs.FargateTaskDefinition `field:"optional" json:"taskDefinition" yaml:"taskDefinition"`
	// Determines whether the service will be assigned a public IP address.
	// Default: false.
	//
	AssignPublicIp *bool `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// The maximum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that can run in a service during a deployment.
	// Default: - 200%.
	//
	MaxHealthyPercent *float64 `field:"optional" json:"maxHealthyPercent" yaml:"maxHealthyPercent"`
	// The minimum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that must continue to run and remain healthy during a deployment.
	// Default: - 50%.
	//
	MinHealthyPercent *float64 `field:"optional" json:"minHealthyPercent" yaml:"minHealthyPercent"`
}

