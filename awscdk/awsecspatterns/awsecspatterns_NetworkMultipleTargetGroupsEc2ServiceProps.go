package awsecspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsecs"
)

// The properties for the NetworkMultipleTargetGroupsEc2Service service.
//
// Example:
//   // Two network load balancers, each with their own listener and target group.
//   var cluster cluster
//
//   loadBalancedEc2Service := ecsPatterns.NewNetworkMultipleTargetGroupsEc2Service(this, jsii.String("Service"), &networkMultipleTargetGroupsEc2ServiceProps{
//   	cluster: cluster,
//   	memoryLimitMiB: jsii.Number(256),
//   	taskImageOptions: &networkLoadBalancedTaskImageProps{
//   		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   	loadBalancers: []networkLoadBalancerProps{
//   		&networkLoadBalancerProps{
//   			name: jsii.String("lb1"),
//   			listeners: []networkListenerProps{
//   				&networkListenerProps{
//   					name: jsii.String("listener1"),
//   				},
//   			},
//   		},
//   		&networkLoadBalancerProps{
//   			name: jsii.String("lb2"),
//   			listeners: []*networkListenerProps{
//   				&networkListenerProps{
//   					name: jsii.String("listener2"),
//   				},
//   			},
//   		},
//   	},
//   	targetGroups: []networkTargetProps{
//   		&networkTargetProps{
//   			containerPort: jsii.Number(80),
//   			listener: jsii.String("listener1"),
//   		},
//   		&networkTargetProps{
//   			containerPort: jsii.Number(90),
//   			listener: jsii.String("listener2"),
//   		},
//   	},
//   })
//
// Experimental.
type NetworkMultipleTargetGroupsEc2ServiceProps struct {
	// The options for configuring an Amazon ECS service to use service discovery.
	// Experimental.
	CloudMapOptions *awsecs.CloudMapOptions `field:"optional" json:"cloudMapOptions" yaml:"cloudMapOptions"`
	// The name of the cluster that hosts the service.
	//
	// If a cluster is specified, the vpc construct should be omitted. Alternatively, you can omit both cluster and vpc.
	// Experimental.
	Cluster awsecs.ICluster `field:"optional" json:"cluster" yaml:"cluster"`
	// The desired number of instantiations of the task definition to keep running on the service.
	//
	// The minimum value is 1.
	// Experimental.
	DesiredCount *float64 `field:"optional" json:"desiredCount" yaml:"desiredCount"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	// Experimental.
	EnableECSManagedTags *bool `field:"optional" json:"enableECSManagedTags" yaml:"enableECSManagedTags"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	// Experimental.
	HealthCheckGracePeriod awscdk.Duration `field:"optional" json:"healthCheckGracePeriod" yaml:"healthCheckGracePeriod"`
	// The network load balancer that will serve traffic to the service.
	// Experimental.
	LoadBalancers *[]*NetworkLoadBalancerProps `field:"optional" json:"loadBalancers" yaml:"loadBalancers"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Tags can only be propagated to the tasks within the service during service creation.
	// Experimental.
	PropagateTags awsecs.PropagatedTagSource `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// Name of the service.
	// Experimental.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// Properties to specify NLB target groups.
	// Experimental.
	TargetGroups *[]*NetworkTargetProps `field:"optional" json:"targetGroups" yaml:"targetGroups"`
	// The properties required to create a new task definition.
	//
	// Only one of TaskDefinition or TaskImageOptions must be specified.
	// Experimental.
	TaskImageOptions *NetworkLoadBalancedTaskImageProps `field:"optional" json:"taskImageOptions" yaml:"taskImageOptions"`
	// The VPC where the container instances will be launched or the elastic network interfaces (ENIs) will be deployed.
	//
	// If a vpc is specified, the cluster construct should be omitted. Alternatively, you can omit both vpc and cluster.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// The minimum number of CPU units to reserve for the container.
	//
	// Valid values, which determines your range of valid values for the memory parameter:.
	// Experimental.
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// The amount (in MiB) of memory to present to the container.
	//
	// If your container attempts to exceed the allocated memory, the container
	// is terminated.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required.
	// Experimental.
	MemoryLimitMiB *float64 `field:"optional" json:"memoryLimitMiB" yaml:"memoryLimitMiB"`
	// The soft limit (in MiB) of memory to reserve for the container.
	//
	// When system memory is under heavy contention, Docker attempts to keep the
	// container memory to this soft limit. However, your container can consume more
	// memory when it needs to, up to either the hard limit specified with the memory
	// parameter (if applicable), or all of the available memory on the container
	// instance, whichever comes first.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required.
	//
	// Note that this setting will be ignored if TaskImagesOptions is specified.
	// Experimental.
	MemoryReservationMiB *float64 `field:"optional" json:"memoryReservationMiB" yaml:"memoryReservationMiB"`
	// The placement constraints to use for tasks in the service.
	//
	// For more information, see
	// [Amazon ECS Task Placement Constraints](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-constraints.html).
	// Experimental.
	PlacementConstraints *[]awsecs.PlacementConstraint `field:"optional" json:"placementConstraints" yaml:"placementConstraints"`
	// The placement strategies to use for tasks in the service.
	//
	// For more information, see
	// [Amazon ECS Task Placement Strategies](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-strategies.html).
	// Experimental.
	PlacementStrategies *[]awsecs.PlacementStrategy `field:"optional" json:"placementStrategies" yaml:"placementStrategies"`
	// The task definition to use for tasks in the service. Only one of TaskDefinition or TaskImageOptions must be specified.
	//
	// [disable-awslint:ref-via-interface].
	// Experimental.
	TaskDefinition awsecs.Ec2TaskDefinition `field:"optional" json:"taskDefinition" yaml:"taskDefinition"`
}

