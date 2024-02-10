package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Complete base service properties that are required to be supplied by the implementation of the BaseService class.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//   var containerDefinition containerDefinition
//   var logDriver logDriver
//   var namespace iNamespace
//   var serviceManagedVolume serviceManagedVolume
//   var taskDefinitionRevision taskDefinitionRevision
//
//   baseServiceProps := &BaseServiceProps{
//   	Cluster: cluster,
//   	LaunchType: awscdk.Aws_ecs.LaunchType_EC2,
//
//   	// the properties below are optional
//   	CapacityProviderStrategies: []capacityProviderStrategy{
//   		&capacityProviderStrategy{
//   			CapacityProvider: jsii.String("capacityProvider"),
//
//   			// the properties below are optional
//   			Base: jsii.Number(123),
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   	CircuitBreaker: &DeploymentCircuitBreaker{
//   		Enable: jsii.Boolean(false),
//   		Rollback: jsii.Boolean(false),
//   	},
//   	CloudMapOptions: &CloudMapOptions{
//   		CloudMapNamespace: namespace,
//   		Container: containerDefinition,
//   		ContainerPort: jsii.Number(123),
//   		DnsRecordType: awscdk.Aws_servicediscovery.DnsRecordType_A,
//   		DnsTtl: cdk.Duration_Minutes(jsii.Number(30)),
//   		FailureThreshold: jsii.Number(123),
//   		Name: jsii.String("name"),
//   	},
//   	DeploymentAlarms: &DeploymentAlarmConfig{
//   		AlarmNames: []*string{
//   			jsii.String("alarmNames"),
//   		},
//
//   		// the properties below are optional
//   		Behavior: awscdk.*Aws_ecs.AlarmBehavior_ROLLBACK_ON_ALARM,
//   	},
//   	DeploymentController: &DeploymentController{
//   		Type: awscdk.*Aws_ecs.DeploymentControllerType_ECS,
//   	},
//   	DesiredCount: jsii.Number(123),
//   	EnableECSManagedTags: jsii.Boolean(false),
//   	EnableExecuteCommand: jsii.Boolean(false),
//   	HealthCheckGracePeriod: cdk.Duration_*Minutes(jsii.Number(30)),
//   	MaxHealthyPercent: jsii.Number(123),
//   	MinHealthyPercent: jsii.Number(123),
//   	PropagateTags: awscdk.*Aws_ecs.PropagatedTagSource_SERVICE,
//   	ServiceConnectConfiguration: &ServiceConnectProps{
//   		LogDriver: logDriver,
//   		Namespace: jsii.String("namespace"),
//   		Services: []serviceConnectService{
//   			&serviceConnectService{
//   				PortMappingName: jsii.String("portMappingName"),
//
//   				// the properties below are optional
//   				DiscoveryName: jsii.String("discoveryName"),
//   				DnsName: jsii.String("dnsName"),
//   				IdleTimeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   				IngressPortOverride: jsii.Number(123),
//   				PerRequestTimeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   				Port: jsii.Number(123),
//   			},
//   		},
//   	},
//   	ServiceName: jsii.String("serviceName"),
//   	TaskDefinitionRevision: taskDefinitionRevision,
//   	VolumeConfigurations: []*serviceManagedVolume{
//   		serviceManagedVolume,
//   	},
//   }
//
type BaseServiceProps struct {
	// The name of the cluster that hosts the service.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
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
	// The launch type on which to run your service.
	//
	// LaunchType will be omitted if capacity provider strategies are specified on the service.
	// See:  - https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-capacityproviderstrategy
	//
	// Valid values are: LaunchType.ECS or LaunchType.FARGATE or LaunchType.EXTERNAL
	//
	LaunchType LaunchType `field:"required" json:"launchType" yaml:"launchType"`
}

