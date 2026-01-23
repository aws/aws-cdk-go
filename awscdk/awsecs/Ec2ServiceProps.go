package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// The properties for defining a service using the EC2 launch type.
//
// Example:
//   var vpc Vpc
//
//
//   // Create an ECS cluster
//   cluster := ecs.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
//   	Vpc: Vpc,
//   })
//
//   // Add capacity to it
//   cluster.AddCapacity(jsii.String("DefaultAutoScalingGroupCapacity"), &AddCapacityOptions{
//   	InstanceType: ec2.NewInstanceType(jsii.String("t2.xlarge")),
//   	DesiredCapacity: jsii.Number(3),
//   })
//
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//
//   taskDefinition.AddContainer(jsii.String("DefaultContainer"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	MemoryLimitMiB: jsii.Number(512),
//   })
//
//   // Instantiate an Amazon ECS Service
//   ecsService := ecs.NewEc2Service(this, jsii.String("Service"), &Ec2ServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	MinHealthyPercent: jsii.Number(100),
//   })
//
type Ec2ServiceProps struct {
	// The name of the cluster that hosts the service.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// bake time minutes for service.
	// Default: - none.
	//
	BakeTime awscdk.Duration `field:"optional" json:"bakeTime" yaml:"bakeTime"`
	// Configuration for canary deployment strategy.
	//
	// Only valid when deploymentStrategy is set to CANARY.
	// Default: - no canary configuration.
	//
	CanaryConfiguration *TrafficShiftConfig `field:"optional" json:"canaryConfiguration" yaml:"canaryConfiguration"`
	// A list of Capacity Provider strategies used to place a service.
	// Default: - undefined.
	//
	CapacityProviderStrategies *[]*CapacityProviderStrategy `field:"optional" json:"capacityProviderStrategies" yaml:"capacityProviderStrategies"`
	// Whether to enable the deployment circuit breaker.
	//
	// If this property is defined, circuit breaker will be implicitly
	// enabled.
	// Default: - disabled.
	//
	CircuitBreaker *DeploymentCircuitBreaker `field:"optional" json:"circuitBreaker" yaml:"circuitBreaker"`
	// The options for configuring an Amazon ECS service to use service discovery.
	// Default: - AWS Cloud Map service discovery is not enabled.
	//
	CloudMapOptions *CloudMapOptions `field:"optional" json:"cloudMapOptions" yaml:"cloudMapOptions"`
	// The alarm(s) to monitor during deployment, and behavior to apply if at least one enters a state of alarm during the deployment or bake time.
	// Default: - No alarms will be monitored during deployment.
	//
	DeploymentAlarms *DeploymentAlarmConfig `field:"optional" json:"deploymentAlarms" yaml:"deploymentAlarms"`
	// Specifies which deployment controller to use for the service.
	//
	// For more information, see
	// [Amazon ECS Deployment Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
	// Default: - Rolling update (ECS).
	//
	DeploymentController *DeploymentController `field:"optional" json:"deploymentController" yaml:"deploymentController"`
	// The deployment strategy to use for the service.
	// Default: ROLLING.
	//
	DeploymentStrategy DeploymentStrategy `field:"optional" json:"deploymentStrategy" yaml:"deploymentStrategy"`
	// The desired number of instantiations of the task definition to keep running on the service.
	// Default: - When creating the service, default is 1; when updating the service, default uses
	// the current task number.
	//
	DesiredCount *float64 `field:"optional" json:"desiredCount" yaml:"desiredCount"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	// Default: false.
	//
	EnableECSManagedTags *bool `field:"optional" json:"enableECSManagedTags" yaml:"enableECSManagedTags"`
	// Whether to enable the ability to execute into a container.
	// Default: - undefined.
	//
	EnableExecuteCommand *bool `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	// Default: - defaults to 60 seconds if at least one load balancer is in-use and it is not already set.
	//
	HealthCheckGracePeriod awscdk.Duration `field:"optional" json:"healthCheckGracePeriod" yaml:"healthCheckGracePeriod"`
	// The lifecycle hooks to execute during deployment stages.
	// Default: - none;.
	//
	LifecycleHooks *[]IDeploymentLifecycleHookTarget `field:"optional" json:"lifecycleHooks" yaml:"lifecycleHooks"`
	// Configuration for linear deployment strategy.
	//
	// Only valid when deploymentStrategy is set to LINEAR.
	// Default: - no linear configuration.
	//
	LinearConfiguration *TrafficShiftConfig `field:"optional" json:"linearConfiguration" yaml:"linearConfiguration"`
	// The maximum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that can run in a service during a deployment.
	// Default: - 100 if daemon, otherwise 200.
	//
	MaxHealthyPercent *float64 `field:"optional" json:"maxHealthyPercent" yaml:"maxHealthyPercent"`
	// The minimum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that must continue to run and remain healthy during a deployment.
	// Default: - 0 if daemon, otherwise 50.
	//
	MinHealthyPercent *float64 `field:"optional" json:"minHealthyPercent" yaml:"minHealthyPercent"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Valid values are: PropagatedTagSource.SERVICE, PropagatedTagSource.TASK_DEFINITION or PropagatedTagSource.NONE
	// Default: PropagatedTagSource.NONE
	//
	PropagateTags PropagatedTagSource `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// Configuration for Service Connect.
	// Default: No ports are advertised via Service Connect on this service, and the service
	// cannot make requests to other services via Service Connect.
	//
	ServiceConnectConfiguration *ServiceConnectProps `field:"optional" json:"serviceConnectConfiguration" yaml:"serviceConnectConfiguration"`
	// The name of the service.
	// Default: - CloudFormation-generated name.
	//
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// Revision number for the task definition or `latest` to use the latest active task revision.
	// Default: - Uses the revision of the passed task definition deployed by CloudFormation.
	//
	TaskDefinitionRevision TaskDefinitionRevision `field:"optional" json:"taskDefinitionRevision" yaml:"taskDefinitionRevision"`
	// Configuration details for a volume used by the service.
	//
	// This allows you to specify
	// details about the EBS volume that can be attched to ECS tasks.
	// Default: - undefined.
	//
	VolumeConfigurations *[]ServiceManagedVolume `field:"optional" json:"volumeConfigurations" yaml:"volumeConfigurations"`
	// The task definition to use for tasks in the service.
	//
	// [disable-awslint:ref-via-interface].
	TaskDefinition TaskDefinition `field:"required" json:"taskDefinition" yaml:"taskDefinition"`
	// Specifies whether the task's elastic network interface receives a public IP address.
	//
	// If true, each task will receive a public IP address.
	//
	// This property is only used for tasks that use the awsvpc network mode.
	// Default: false.
	//
	AssignPublicIp *bool `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// Whether to use Availability Zone rebalancing for the service.
	//
	// If enabled: `maxHealthyPercent` must be greater than 100; `daemon` must be false; if there
	// are any `placementStrategies`, the first must be "spread across Availability Zones"; there
	// must be no `placementConstraints` using `attribute:ecs.availability-zone`, and the
	// service must not be a target of a Classic Load Balancer.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-rebalancing.html
	//
	// Default: AvailabilityZoneRebalancing.ENABLED
	//
	AvailabilityZoneRebalancing AvailabilityZoneRebalancing `field:"optional" json:"availabilityZoneRebalancing" yaml:"availabilityZoneRebalancing"`
	// Specifies whether the service will use the daemon scheduling strategy.
	//
	// If true, the service scheduler deploys exactly one task on each container instance in your cluster.
	//
	// When you are using this strategy, do not specify a desired number of tasks or any task placement strategies.
	// Default: false.
	//
	Daemon *bool `field:"optional" json:"daemon" yaml:"daemon"`
	// The placement constraints to use for tasks in the service.
	//
	// For more information, see
	// [Amazon ECS Task Placement Constraints](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-constraints.html).
	// Default: - No constraints.
	//
	PlacementConstraints *[]PlacementConstraint `field:"optional" json:"placementConstraints" yaml:"placementConstraints"`
	// The placement strategies to use for tasks in the service.
	//
	// For more information, see
	// [Amazon ECS Task Placement Strategies](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-strategies.html).
	// Default: - No strategies.
	//
	PlacementStrategies *[]PlacementStrategy `field:"optional" json:"placementStrategies" yaml:"placementStrategies"`
	// The security groups to associate with the service.
	//
	// If you do not specify a security group, a new security group is created.
	//
	// This property is only used for tasks that use the awsvpc network mode.
	// Default: - A new security group is created.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The subnets to associate with the service.
	//
	// This property is only used for tasks that use the awsvpc network mode.
	// Default: - Public subnets if `assignPublicIp` is set, otherwise the first available one of Private, Isolated, Public, in that order.
	//
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

