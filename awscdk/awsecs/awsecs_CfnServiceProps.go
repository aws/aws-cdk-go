package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnService`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServiceProps := &cfnServiceProps{
//   	capacityProviderStrategy: []interface{}{
//   		&capacityProviderStrategyItemProperty{
//   			base: jsii.Number(123),
//   			capacityProvider: jsii.String("capacityProvider"),
//   			weight: jsii.Number(123),
//   		},
//   	},
//   	cluster: jsii.String("cluster"),
//   	deploymentConfiguration: &deploymentConfigurationProperty{
//   		deploymentCircuitBreaker: &deploymentCircuitBreakerProperty{
//   			enable: jsii.Boolean(false),
//   			rollback: jsii.Boolean(false),
//   		},
//   		maximumPercent: jsii.Number(123),
//   		minimumHealthyPercent: jsii.Number(123),
//   	},
//   	deploymentController: &deploymentControllerProperty{
//   		type: jsii.String("type"),
//   	},
//   	desiredCount: jsii.Number(123),
//   	enableEcsManagedTags: jsii.Boolean(false),
//   	enableExecuteCommand: jsii.Boolean(false),
//   	healthCheckGracePeriodSeconds: jsii.Number(123),
//   	launchType: jsii.String("launchType"),
//   	loadBalancers: []interface{}{
//   		&loadBalancerProperty{
//   			containerPort: jsii.Number(123),
//
//   			// the properties below are optional
//   			containerName: jsii.String("containerName"),
//   			loadBalancerName: jsii.String("loadBalancerName"),
//   			targetGroupArn: jsii.String("targetGroupArn"),
//   		},
//   	},
//   	networkConfiguration: &networkConfigurationProperty{
//   		awsvpcConfiguration: &awsVpcConfigurationProperty{
//   			subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//
//   			// the properties below are optional
//   			assignPublicIp: jsii.String("assignPublicIp"),
//   			securityGroups: []*string{
//   				jsii.String("securityGroups"),
//   			},
//   		},
//   	},
//   	placementConstraints: []interface{}{
//   		&placementConstraintProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			expression: jsii.String("expression"),
//   		},
//   	},
//   	placementStrategies: []interface{}{
//   		&placementStrategyProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			field: jsii.String("field"),
//   		},
//   	},
//   	platformVersion: jsii.String("platformVersion"),
//   	propagateTags: jsii.String("propagateTags"),
//   	role: jsii.String("role"),
//   	schedulingStrategy: jsii.String("schedulingStrategy"),
//   	serviceConnectConfiguration: &serviceConnectConfigurationProperty{
//   		enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		logConfiguration: &logConfigurationProperty{
//   			logDriver: jsii.String("logDriver"),
//   			options: map[string]*string{
//   				"optionsKey": jsii.String("options"),
//   			},
//   			secretOptions: []interface{}{
//   				&secretProperty{
//   					name: jsii.String("name"),
//   					valueFrom: jsii.String("valueFrom"),
//   				},
//   			},
//   		},
//   		namespace: jsii.String("namespace"),
//   		services: []interface{}{
//   			&serviceConnectServiceProperty{
//   				portName: jsii.String("portName"),
//
//   				// the properties below are optional
//   				clientAliases: []interface{}{
//   					&serviceConnectClientAliasProperty{
//   						port: jsii.Number(123),
//
//   						// the properties below are optional
//   						dnsName: jsii.String("dnsName"),
//   					},
//   				},
//   				discoveryName: jsii.String("discoveryName"),
//   				ingressPortOverride: jsii.Number(123),
//   			},
//   		},
//   	},
//   	serviceName: jsii.String("serviceName"),
//   	serviceRegistries: []interface{}{
//   		&serviceRegistryProperty{
//   			containerName: jsii.String("containerName"),
//   			containerPort: jsii.Number(123),
//   			port: jsii.Number(123),
//   			registryArn: jsii.String("registryArn"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	taskDefinition: jsii.String("taskDefinition"),
//   }
//
type CfnServiceProps struct {
	// The capacity provider strategy to use for the service.
	//
	// A capacity provider strategy consists of one or more capacity providers along with the `base` and `weight` to assign to them. A capacity provider must be associated with the cluster to be used in a capacity provider strategy. The PutClusterCapacityProviders API is used to associate a capacity provider with a cluster. Only capacity providers with an `ACTIVE` or `UPDATING` status can be used.
	//
	// Review the [Capacity provider considerations](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-capacity-providers.html#capacity-providers-considerations) in the *Amazon Elastic Container Service Developer Guide.*
	//
	// If a `capacityProviderStrategy` is specified, the `launchType` parameter must be omitted. If no `capacityProviderStrategy` or `launchType` is specified, the `defaultCapacityProviderStrategy` for the cluster is used.
	//
	// If specifying a capacity provider that uses an Auto Scaling group, the capacity provider must already be created. New capacity providers can be created with the CreateCapacityProvider API operation.
	//
	// To use an AWS Fargate capacity provider, specify either the `FARGATE` or `FARGATE_SPOT` capacity providers. The AWS Fargate capacity providers are available to all accounts and only need to be associated with a cluster to be used.
	//
	// The PutClusterCapacityProviders API operation is used to update the list of available capacity providers for a cluster after the cluster is created.
	CapacityProviderStrategy interface{} `field:"optional" json:"capacityProviderStrategy" yaml:"capacityProviderStrategy"`
	// The short name or full Amazon Resource Name (ARN) of the cluster that you run your service on.
	//
	// If you do not specify a cluster, the default cluster is assumed.
	Cluster *string `field:"optional" json:"cluster" yaml:"cluster"`
	// Optional deployment parameters that control how many tasks run during the deployment and the ordering of stopping and starting tasks.
	DeploymentConfiguration interface{} `field:"optional" json:"deploymentConfiguration" yaml:"deploymentConfiguration"`
	// The deployment controller to use for the service.
	//
	// If no deployment controller is specified, the default value of `ECS` is used.
	DeploymentController interface{} `field:"optional" json:"deploymentController" yaml:"deploymentController"`
	// The number of instantiations of the specified task definition to place and keep running on your cluster.
	//
	// For new services, if a desired count is not specified, a default value of `1` is used. When using the `DAEMON` scheduling strategy, the desired count is not required.
	//
	// For existing services, if a desired count is not specified, it is omitted from the operation.
	DesiredCount *float64 `field:"optional" json:"desiredCount" yaml:"desiredCount"`
	// Specifies whether to turn on Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see [Tagging your Amazon ECS resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html) in the *Amazon Elastic Container Service Developer Guide* .
	EnableEcsManagedTags interface{} `field:"optional" json:"enableEcsManagedTags" yaml:"enableEcsManagedTags"`
	// Determines whether the execute command functionality is enabled for the service.
	//
	// If `true` , the execute command functionality is enabled for all containers in tasks as part of the service.
	EnableExecuteCommand interface{} `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	//
	// This is only used when your service is configured to use a load balancer. If your service has a load balancer defined and you don't specify a health check grace period value, the default value of `0` is used.
	//
	// If you do not use an Elastic Load Balancing, we recommend that you use the `startPeriod` in the task definition health check parameters. For more information, see [Health check](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_HealthCheck.html) .
	//
	// If your service's tasks take a while to start and respond to Elastic Load Balancing health checks, you can specify a health check grace period of up to 2,147,483,647 seconds (about 69 years). During that time, the Amazon ECS service scheduler ignores health check status. This grace period can prevent the service scheduler from marking tasks as unhealthy and stopping them before they have time to come up.
	HealthCheckGracePeriodSeconds *float64 `field:"optional" json:"healthCheckGracePeriodSeconds" yaml:"healthCheckGracePeriodSeconds"`
	// The launch type on which to run your service.
	//
	// For more information, see [Amazon ECS Launch Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html) in the *Amazon Elastic Container Service Developer Guide* .
	LaunchType *string `field:"optional" json:"launchType" yaml:"launchType"`
	// A list of load balancer objects to associate with the service.
	//
	// If you specify the `Role` property, `LoadBalancers` must be specified as well. For information about the number of load balancers that you can specify per service, see [Service Load Balancing](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-load-balancing.html) in the *Amazon Elastic Container Service Developer Guide* .
	LoadBalancers interface{} `field:"optional" json:"loadBalancers" yaml:"loadBalancers"`
	// The network configuration for the service.
	//
	// This parameter is required for task definitions that use the `awsvpc` network mode to receive their own elastic network interface, and it is not supported for other network modes. For more information, see [Task Networking](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html) in the *Amazon Elastic Container Service Developer Guide* .
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// An array of placement constraint objects to use for tasks in your service.
	//
	// You can specify a maximum of 10 constraints for each task. This limit includes constraints in the task definition and those specified at runtime.
	PlacementConstraints interface{} `field:"optional" json:"placementConstraints" yaml:"placementConstraints"`
	// The placement strategy objects to use for tasks in your service.
	//
	// You can specify a maximum of five strategy rules per service. For more information, see [Task Placement Strategies](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-strategies.html) in the *Amazon Elastic Container Service Developer Guide* .
	PlacementStrategies interface{} `field:"optional" json:"placementStrategies" yaml:"placementStrategies"`
	// The platform version that your tasks in the service are running on.
	//
	// A platform version is specified only for tasks using the Fargate launch type. If one isn't specified, the `LATEST` platform version is used. For more information, see [AWS Fargate platform versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html) in the *Amazon Elastic Container Service Developer Guide* .
	PlatformVersion *string `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// If no value is specified, the tags are not propagated. Tags can only be propagated to the tasks within the service during service creation. To add tags to a task after service creation, use the [TagResource](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TagResource.html) API action.
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// The name or full Amazon Resource Name (ARN) of the IAM role that allows Amazon ECS to make calls to your load balancer on your behalf.
	//
	// This parameter is only permitted if you are using a load balancer with your service and your task definition doesn't use the `awsvpc` network mode. If you specify the `role` parameter, you must also specify a load balancer object with the `loadBalancers` parameter.
	//
	// > If your account has already created the Amazon ECS service-linked role, that role is used for your service unless you specify a role here. The service-linked role is required if your task definition uses the `awsvpc` network mode or if the service is configured to use service discovery, an external deployment controller, multiple target groups, or Elastic Inference accelerators in which case you don't specify a role here. For more information, see [Using service-linked roles for Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using-service-linked-roles.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// If your specified role has a path other than `/` , then you must either specify the full role ARN (this is recommended) or prefix the role name with the path. For example, if a role with the name `bar` has a path of `/foo/` then you would specify `/foo/bar` as the role name. For more information, see [Friendly names and paths](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-friendly-names) in the *IAM User Guide* .
	Role *string `field:"optional" json:"role" yaml:"role"`
	// The scheduling strategy to use for the service. For more information, see [Services](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs_services.html) .
	//
	// There are two service scheduler strategies available:
	//
	// - `REPLICA` -The replica scheduling strategy places and maintains the desired number of tasks across your cluster. By default, the service scheduler spreads tasks across Availability Zones. You can use task placement strategies and constraints to customize task placement decisions. This scheduler strategy is required if the service uses the `CODE_DEPLOY` or `EXTERNAL` deployment controller types.
	// - `DAEMON` -The daemon scheduling strategy deploys exactly one task on each active container instance that meets all of the task placement constraints that you specify in your cluster. The service scheduler also evaluates the task placement constraints for running tasks and will stop tasks that don't meet the placement constraints. When you're using this strategy, you don't need to specify a desired number of tasks, a task placement strategy, or use Service Auto Scaling policies.
	//
	// > Tasks using the Fargate launch type or the `CODE_DEPLOY` or `EXTERNAL` deployment controller types don't support the `DAEMON` scheduling strategy.
	SchedulingStrategy *string `field:"optional" json:"schedulingStrategy" yaml:"schedulingStrategy"`
	// The configuration for this service to discover and connect to services, and be discovered by, and connected from, other services within a namespace.
	//
	// Tasks that run in a namespace can use short names to connect to services in the namespace. Tasks can connect to services across all of the clusters in the namespace. Tasks connect through a managed proxy container that collects logs and metrics for increased visibility. Only the tasks that Amazon ECS services create are supported with Service Connect. For more information, see [Service Connect](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-connect.html) in the *Amazon Elastic Container Service Developer Guide* .
	ServiceConnectConfiguration interface{} `field:"optional" json:"serviceConnectConfiguration" yaml:"serviceConnectConfiguration"`
	// The name of your service.
	//
	// Up to 255 letters (uppercase and lowercase), numbers, underscores, and hyphens are allowed. Service names must be unique within a cluster, but you can have similarly named services in multiple clusters within a Region or across multiple Regions.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// The details of the service discovery registry to associate with this service. For more information, see [Service discovery](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html) .
	//
	// > Each service may be associated with one service registry. Multiple service registries for each service isn't supported.
	ServiceRegistries interface{} `field:"optional" json:"serviceRegistries" yaml:"serviceRegistries"`
	// The metadata that you apply to the service to help you categorize and organize them.
	//
	// Each tag consists of a key and an optional value, both of which you define. When a service is deleted, the tags are deleted as well.
	//
	// The following basic restrictions apply to tags:
	//
	// - Maximum number of tags per resource - 50
	// - For each resource, each tag key must be unique, and each tag key can have only one value.
	// - Maximum key length - 128 Unicode characters in UTF-8
	// - Maximum value length - 256 Unicode characters in UTF-8
	// - If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.
	// - Tag keys and values are case-sensitive.
	// - Do not use `aws:` , `AWS:` , or any upper or lowercase combination of such as a prefix for either keys or values as it is reserved for AWS use. You cannot edit or delete tag keys or values with this prefix. Tags with this prefix do not count against your tags per resource limit.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The `family` and `revision` ( `family:revision` ) or full ARN of the task definition to run in your service.
	//
	// The `revision` is required in order for the resource to stabilize.
	//
	// A task definition must be specified if the service is using either the `ECS` or `CODE_DEPLOY` deployment controllers.
	//
	// For more information about deployment types, see [Amazon ECS deployment types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html) .
	TaskDefinition *string `field:"optional" json:"taskDefinition" yaml:"taskDefinition"`
}

