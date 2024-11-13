package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnService`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServiceProps := &CfnServiceProps{
//   	CapacityProviderStrategy: []interface{}{
//   		&CapacityProviderStrategyItemProperty{
//   			Base: jsii.Number(123),
//   			CapacityProvider: jsii.String("capacityProvider"),
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   	Cluster: jsii.String("cluster"),
//   	DeploymentConfiguration: &DeploymentConfigurationProperty{
//   		Alarms: &DeploymentAlarmsProperty{
//   			AlarmNames: []*string{
//   				jsii.String("alarmNames"),
//   			},
//   			Enable: jsii.Boolean(false),
//   			Rollback: jsii.Boolean(false),
//   		},
//   		DeploymentCircuitBreaker: &DeploymentCircuitBreakerProperty{
//   			Enable: jsii.Boolean(false),
//   			Rollback: jsii.Boolean(false),
//   		},
//   		MaximumPercent: jsii.Number(123),
//   		MinimumHealthyPercent: jsii.Number(123),
//   	},
//   	DeploymentController: &DeploymentControllerProperty{
//   		Type: jsii.String("type"),
//   	},
//   	DesiredCount: jsii.Number(123),
//   	EnableEcsManagedTags: jsii.Boolean(false),
//   	EnableExecuteCommand: jsii.Boolean(false),
//   	HealthCheckGracePeriodSeconds: jsii.Number(123),
//   	LaunchType: jsii.String("launchType"),
//   	LoadBalancers: []interface{}{
//   		&LoadBalancerProperty{
//   			ContainerName: jsii.String("containerName"),
//   			ContainerPort: jsii.Number(123),
//   			LoadBalancerName: jsii.String("loadBalancerName"),
//   			TargetGroupArn: jsii.String("targetGroupArn"),
//   		},
//   	},
//   	NetworkConfiguration: &NetworkConfigurationProperty{
//   		AwsvpcConfiguration: &AwsVpcConfigurationProperty{
//   			AssignPublicIp: jsii.String("assignPublicIp"),
//   			SecurityGroups: []*string{
//   				jsii.String("securityGroups"),
//   			},
//   			Subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   	},
//   	PlacementConstraints: []interface{}{
//   		&PlacementConstraintProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Expression: jsii.String("expression"),
//   		},
//   	},
//   	PlacementStrategies: []interface{}{
//   		&PlacementStrategyProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Field: jsii.String("field"),
//   		},
//   	},
//   	PlatformVersion: jsii.String("platformVersion"),
//   	PropagateTags: jsii.String("propagateTags"),
//   	Role: jsii.String("role"),
//   	SchedulingStrategy: jsii.String("schedulingStrategy"),
//   	ServiceConnectConfiguration: &ServiceConnectConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		LogConfiguration: &LogConfigurationProperty{
//   			LogDriver: jsii.String("logDriver"),
//   			Options: map[string]*string{
//   				"optionsKey": jsii.String("options"),
//   			},
//   			SecretOptions: []interface{}{
//   				&SecretProperty{
//   					Name: jsii.String("name"),
//   					ValueFrom: jsii.String("valueFrom"),
//   				},
//   			},
//   		},
//   		Namespace: jsii.String("namespace"),
//   		Services: []interface{}{
//   			&ServiceConnectServiceProperty{
//   				PortName: jsii.String("portName"),
//
//   				// the properties below are optional
//   				ClientAliases: []interface{}{
//   					&ServiceConnectClientAliasProperty{
//   						Port: jsii.Number(123),
//
//   						// the properties below are optional
//   						DnsName: jsii.String("dnsName"),
//   					},
//   				},
//   				DiscoveryName: jsii.String("discoveryName"),
//   				IngressPortOverride: jsii.Number(123),
//   				Timeout: &TimeoutConfigurationProperty{
//   					IdleTimeoutSeconds: jsii.Number(123),
//   					PerRequestTimeoutSeconds: jsii.Number(123),
//   				},
//   				Tls: &ServiceConnectTlsConfigurationProperty{
//   					IssuerCertificateAuthority: &ServiceConnectTlsCertificateAuthorityProperty{
//   						AwsPcaAuthorityArn: jsii.String("awsPcaAuthorityArn"),
//   					},
//
//   					// the properties below are optional
//   					KmsKey: jsii.String("kmsKey"),
//   					RoleArn: jsii.String("roleArn"),
//   				},
//   			},
//   		},
//   	},
//   	ServiceName: jsii.String("serviceName"),
//   	ServiceRegistries: []interface{}{
//   		&ServiceRegistryProperty{
//   			ContainerName: jsii.String("containerName"),
//   			ContainerPort: jsii.Number(123),
//   			Port: jsii.Number(123),
//   			RegistryArn: jsii.String("registryArn"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TaskDefinition: jsii.String("taskDefinition"),
//   	VolumeConfigurations: []interface{}{
//   		&ServiceVolumeConfigurationProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			ManagedEbsVolume: &ServiceManagedEBSVolumeConfigurationProperty{
//   				RoleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				Encrypted: jsii.Boolean(false),
//   				FilesystemType: jsii.String("filesystemType"),
//   				Iops: jsii.Number(123),
//   				KmsKeyId: jsii.String("kmsKeyId"),
//   				SizeInGiB: jsii.Number(123),
//   				SnapshotId: jsii.String("snapshotId"),
//   				TagSpecifications: []interface{}{
//   					&EBSTagSpecificationProperty{
//   						ResourceType: jsii.String("resourceType"),
//
//   						// the properties below are optional
//   						PropagateTags: jsii.String("propagateTags"),
//   						Tags: []*cfnTag{
//   							&cfnTag{
//   								Key: jsii.String("key"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				Throughput: jsii.Number(123),
//   				VolumeType: jsii.String("volumeType"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html
//
type CfnServiceProps struct {
	// The capacity provider strategy to use for the service.
	//
	// If a `capacityProviderStrategy` is specified, the `launchType` parameter must be omitted. If no `capacityProviderStrategy` or `launchType` is specified, the `defaultCapacityProviderStrategy` for the cluster is used.
	//
	// A capacity provider strategy may contain a maximum of 6 capacity providers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-capacityproviderstrategy
	//
	CapacityProviderStrategy interface{} `field:"optional" json:"capacityProviderStrategy" yaml:"capacityProviderStrategy"`
	// The short name or full Amazon Resource Name (ARN) of the cluster that you run your service on.
	//
	// If you do not specify a cluster, the default cluster is assumed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-cluster
	//
	Cluster *string `field:"optional" json:"cluster" yaml:"cluster"`
	// Optional deployment parameters that control how many tasks run during the deployment and the ordering of stopping and starting tasks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-deploymentconfiguration
	//
	DeploymentConfiguration interface{} `field:"optional" json:"deploymentConfiguration" yaml:"deploymentConfiguration"`
	// The deployment controller to use for the service.
	//
	// If no deployment controller is specified, the default value of `ECS` is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-deploymentcontroller
	//
	DeploymentController interface{} `field:"optional" json:"deploymentController" yaml:"deploymentController"`
	// The number of instantiations of the specified task definition to place and keep running in your service.
	//
	// For new services, if a desired count is not specified, a default value of `1` is used. When using the `DAEMON` scheduling strategy, the desired count is not required.
	//
	// For existing services, if a desired count is not specified, it is omitted from the operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-desiredcount
	//
	DesiredCount *float64 `field:"optional" json:"desiredCount" yaml:"desiredCount"`
	// Specifies whether to turn on Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see [Tagging your Amazon ECS resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// When you use Amazon ECS managed tags, you need to set the `propagateTags` request parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-enableecsmanagedtags
	//
	EnableEcsManagedTags interface{} `field:"optional" json:"enableEcsManagedTags" yaml:"enableEcsManagedTags"`
	// Determines whether the execute command functionality is turned on for the service.
	//
	// If `true` , the execute command functionality is turned on for all containers in tasks as part of the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-enableexecutecommand
	//
	EnableExecuteCommand interface{} `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	//
	// This is only used when your service is configured to use a load balancer. If your service has a load balancer defined and you don't specify a health check grace period value, the default value of `0` is used.
	//
	// If you do not use an Elastic Load Balancing, we recommend that you use the `startPeriod` in the task definition health check parameters. For more information, see [Health check](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_HealthCheck.html) .
	//
	// If your service's tasks take a while to start and respond to Elastic Load Balancing health checks, you can specify a health check grace period of up to 2,147,483,647 seconds (about 69 years). During that time, the Amazon ECS service scheduler ignores health check status. This grace period can prevent the service scheduler from marking tasks as unhealthy and stopping them before they have time to come up.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-healthcheckgraceperiodseconds
	//
	HealthCheckGracePeriodSeconds *float64 `field:"optional" json:"healthCheckGracePeriodSeconds" yaml:"healthCheckGracePeriodSeconds"`
	// The launch type on which to run your service.
	//
	// For more information, see [Amazon ECS Launch Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html) in the *Amazon Elastic Container Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-launchtype
	//
	LaunchType *string `field:"optional" json:"launchType" yaml:"launchType"`
	// A list of load balancer objects to associate with the service.
	//
	// If you specify the `Role` property, `LoadBalancers` must be specified as well. For information about the number of load balancers that you can specify per service, see [Service Load Balancing](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-load-balancing.html) in the *Amazon Elastic Container Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-loadbalancers
	//
	LoadBalancers interface{} `field:"optional" json:"loadBalancers" yaml:"loadBalancers"`
	// The network configuration for the service.
	//
	// This parameter is required for task definitions that use the `awsvpc` network mode to receive their own elastic network interface, and it is not supported for other network modes. For more information, see [Task Networking](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html) in the *Amazon Elastic Container Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-networkconfiguration
	//
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// An array of placement constraint objects to use for tasks in your service.
	//
	// You can specify a maximum of 10 constraints for each task. This limit includes constraints in the task definition and those specified at runtime.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-placementconstraints
	//
	PlacementConstraints interface{} `field:"optional" json:"placementConstraints" yaml:"placementConstraints"`
	// The placement strategy objects to use for tasks in your service.
	//
	// You can specify a maximum of 5 strategy rules for each service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-placementstrategies
	//
	PlacementStrategies interface{} `field:"optional" json:"placementStrategies" yaml:"placementStrategies"`
	// The platform version that your tasks in the service are running on.
	//
	// A platform version is specified only for tasks using the Fargate launch type. If one isn't specified, the `LATEST` platform version is used. For more information, see [AWS Fargate platform versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html) in the *Amazon Elastic Container Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-platformversion
	//
	// Default: - "LATEST".
	//
	PlatformVersion *string `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// Specifies whether to propagate the tags from the task definition to the task.
	//
	// If no value is specified, the tags aren't propagated. Tags can only be propagated to the task during task creation. To add tags to a task after task creation, use the [TagResource](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TagResource.html) API action.
	//
	// You must set this to a value other than `NONE` when you use Cost Explorer. For more information, see [Amazon ECS usage reports](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/usage-reports.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// The default is `NONE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-propagatetags
	//
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// The name or full Amazon Resource Name (ARN) of the IAM role that allows Amazon ECS to make calls to your load balancer on your behalf.
	//
	// This parameter is only permitted if you are using a load balancer with your service and your task definition doesn't use the `awsvpc` network mode. If you specify the `role` parameter, you must also specify a load balancer object with the `loadBalancers` parameter.
	//
	// > If your account has already created the Amazon ECS service-linked role, that role is used for your service unless you specify a role here. The service-linked role is required if your task definition uses the `awsvpc` network mode or if the service is configured to use service discovery, an external deployment controller, multiple target groups, or Elastic Inference accelerators in which case you don't specify a role here. For more information, see [Using service-linked roles for Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using-service-linked-roles.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// If your specified role has a path other than `/` , then you must either specify the full role ARN (this is recommended) or prefix the role name with the path. For example, if a role with the name `bar` has a path of `/foo/` then you would specify `/foo/bar` as the role name. For more information, see [Friendly names and paths](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-friendly-names) in the *IAM User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-role
	//
	Role *string `field:"optional" json:"role" yaml:"role"`
	// The scheduling strategy to use for the service. For more information, see [Services](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs_services.html) .
	//
	// There are two service scheduler strategies available:
	//
	// - `REPLICA` -The replica scheduling strategy places and maintains the desired number of tasks across your cluster. By default, the service scheduler spreads tasks across Availability Zones. You can use task placement strategies and constraints to customize task placement decisions. This scheduler strategy is required if the service uses the `CODE_DEPLOY` or `EXTERNAL` deployment controller types.
	// - `DAEMON` -The daemon scheduling strategy deploys exactly one task on each active container instance that meets all of the task placement constraints that you specify in your cluster. The service scheduler also evaluates the task placement constraints for running tasks and will stop tasks that don't meet the placement constraints. When you're using this strategy, you don't need to specify a desired number of tasks, a task placement strategy, or use Service Auto Scaling policies.
	//
	// > Tasks using the Fargate launch type or the `CODE_DEPLOY` or `EXTERNAL` deployment controller types don't support the `DAEMON` scheduling strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-schedulingstrategy
	//
	SchedulingStrategy *string `field:"optional" json:"schedulingStrategy" yaml:"schedulingStrategy"`
	// The configuration for this service to discover and connect to services, and be discovered by, and connected from, other services within a namespace.
	//
	// Tasks that run in a namespace can use short names to connect to services in the namespace. Tasks can connect to services across all of the clusters in the namespace. Tasks connect through a managed proxy container that collects logs and metrics for increased visibility. Only the tasks that Amazon ECS services create are supported with Service Connect. For more information, see [Service Connect](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-connect.html) in the *Amazon Elastic Container Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-serviceconnectconfiguration
	//
	ServiceConnectConfiguration interface{} `field:"optional" json:"serviceConnectConfiguration" yaml:"serviceConnectConfiguration"`
	// The name of your service.
	//
	// Up to 255 letters (uppercase and lowercase), numbers, underscores, and hyphens are allowed. Service names must be unique within a cluster, but you can have similarly named services in multiple clusters within a Region or across multiple Regions.
	//
	// > The stack update fails if you change any properties that require replacement and the `ServiceName` is configured. This is because AWS CloudFormation creates the replacement service first, but each `ServiceName` must be unique in the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-servicename
	//
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// The details of the service discovery registry to associate with this service. For more information, see [Service discovery](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html) .
	//
	// > Each service may be associated with one service registry. Multiple service registries for each service isn't supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-serviceregistries
	//
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
	// - If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : /
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The `family` and `revision` ( `family:revision` ) or full ARN of the task definition to run in your service.
	//
	// If a `revision` isn't specified, the latest `ACTIVE` revision is used.
	//
	// A task definition must be specified if the service uses either the `ECS` or `CODE_DEPLOY` deployment controllers.
	//
	// For more information about deployment types, see [Amazon ECS deployment types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-taskdefinition
	//
	TaskDefinition *string `field:"optional" json:"taskDefinition" yaml:"taskDefinition"`
	// The configuration for a volume specified in the task definition as a volume that is configured at launch time.
	//
	// Currently, the only supported volume type is an Amazon EBS volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-volumeconfigurations
	//
	VolumeConfigurations interface{} `field:"optional" json:"volumeConfigurations" yaml:"volumeConfigurations"`
}

