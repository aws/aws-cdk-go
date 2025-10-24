package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// The properties for defining a service using the Fargate launch type.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster Cluster
//   var taskDefinition TaskDefinition
//
//   serviceName := "MyFargateService"
//   service := ecs.NewFargateService(this, jsii.String("Service"), &FargateServiceProps{
//   	ServiceName: jsii.String(ServiceName),
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	MinHealthyPercent: jsii.Number(100),
//   })
//
//   cpuMetric := cw.NewMetric(&MetricProps{
//   	MetricName: jsii.String("CPUUtilization"),
//   	Namespace: jsii.String("AWS/ECS"),
//   	Period: awscdk.Duration_Minutes(jsii.Number(5)),
//   	Statistic: jsii.String("Average"),
//   	DimensionsMap: map[string]*string{
//   		"ClusterName": cluster.clusterName,
//   		// Using `service.serviceName` here will cause a circular dependency
//   		"ServiceName": serviceName,
//   	},
//   })
//   myAlarm := cw.NewAlarm(this, jsii.String("CPUAlarm"), &AlarmProps{
//   	AlarmName: jsii.String("cpuAlarmName"),
//   	Metric: cpuMetric,
//   	EvaluationPeriods: jsii.Number(2),
//   	Threshold: jsii.Number(80),
//   })
//
//   service.EnableDeploymentAlarms([]*string{
//   	myAlarm.AlarmName,
//   }, &DeploymentAlarmOptions{
//   	Behavior: ecs.AlarmBehavior_FAIL_ON_ALARM,
//   })
//
type FargateServiceProps struct {
	// The name of the cluster that hosts the service.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// bake time minutes for service.
	// Default: - none.
	//
	BakeTime awscdk.Duration `field:"optional" json:"bakeTime" yaml:"bakeTime"`
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
	// Default: false.
	//
	AssignPublicIp *bool `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// Whether to use Availability Zone rebalancing for the service.
	//
	// If enabled, `maxHealthyPercent` must be greater than 100, and the service must not be a target
	// of a Classic Load Balancer.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-rebalancing.html
	//
	// Default: AvailabilityZoneRebalancing.ENABLED
	//
	AvailabilityZoneRebalancing AvailabilityZoneRebalancing `field:"optional" json:"availabilityZoneRebalancing" yaml:"availabilityZoneRebalancing"`
	// The platform version on which to run your service.
	//
	// If one is not specified, the LATEST platform version is used by default. For more information, see
	// [AWS Fargate Platform Versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html)
	// in the Amazon Elastic Container Service Developer Guide.
	// Default: Latest.
	//
	PlatformVersion FargatePlatformVersion `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// The security groups to associate with the service.
	//
	// If you do not specify a security group, a new security group is created.
	// Default: - A new security group is created.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The subnets to associate with the service.
	// Default: - Public subnets if `assignPublicIp` is set, otherwise the first available one of Private, Isolated, Public, in that order.
	//
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

