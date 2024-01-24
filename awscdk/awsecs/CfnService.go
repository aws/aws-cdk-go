package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::ECS::Service` resource creates an Amazon Elastic Container Service (Amazon ECS) service that runs and maintains the requested number of tasks and associated load balancers.
//
// > The stack update fails if you change any properties that require replacement and at least one Amazon ECS Service Connect `ServiceConnectService` is configured. This is because AWS CloudFormation creates the replacement service first, but each `ServiceConnectService` must have a name that is unique in the namespace. > Starting April 15, 2023, AWS ; will not onboard new customers to Amazon Elastic Inference (EI), and will help current customers migrate their workloads to options that offer better price and performance. After April 15, 2023, new customers will not be able to launch instances with Amazon EI accelerators in Amazon SageMaker, Amazon ECS , or Amazon EC2 . However, customers who have used Amazon EI at least once during the past 30-day period are considered current customers and will be able to continue using the service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnService := awscdk.Aws_ecs.NewCfnService(this, jsii.String("MyCfnService"), &CfnServiceProps{
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
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html
//
type CfnService interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
	// The name of the Amazon ECS service, such as `sample-webapp` .
	AttrName() *string
	// Not currently supported in AWS CloudFormation .
	AttrServiceArn() *string
	// The capacity provider strategy to use for the service.
	CapacityProviderStrategy() interface{}
	SetCapacityProviderStrategy(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The short name or full Amazon Resource Name (ARN) of the cluster that you run your service on.
	Cluster() *string
	SetCluster(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Optional deployment parameters that control how many tasks run during the deployment and the ordering of stopping and starting tasks.
	DeploymentConfiguration() interface{}
	SetDeploymentConfiguration(val interface{})
	// The deployment controller to use for the service.
	DeploymentController() interface{}
	SetDeploymentController(val interface{})
	// The number of instantiations of the specified task definition to place and keep running in your service.
	DesiredCount() *float64
	SetDesiredCount(val *float64)
	// Specifies whether to turn on Amazon ECS managed tags for the tasks within the service.
	EnableEcsManagedTags() interface{}
	SetEnableEcsManagedTags(val interface{})
	// Determines whether the execute command functionality is turned on for the service.
	EnableExecuteCommand() interface{}
	SetEnableExecuteCommand(val interface{})
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	HealthCheckGracePeriodSeconds() *float64
	SetHealthCheckGracePeriodSeconds(val *float64)
	// The launch type on which to run your service.
	LaunchType() *string
	SetLaunchType(val *string)
	// A list of load balancer objects to associate with the service.
	LoadBalancers() interface{}
	SetLoadBalancers(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The network configuration for the service.
	NetworkConfiguration() interface{}
	SetNetworkConfiguration(val interface{})
	// The tree node.
	Node() constructs.Node
	// An array of placement constraint objects to use for tasks in your service.
	PlacementConstraints() interface{}
	SetPlacementConstraints(val interface{})
	// The placement strategy objects to use for tasks in your service.
	PlacementStrategies() interface{}
	SetPlacementStrategies(val interface{})
	// The platform version that your tasks in the service are running on.
	PlatformVersion() *string
	SetPlatformVersion(val *string)
	// Specifies whether to propagate the tags from the task definition to the task.
	PropagateTags() *string
	SetPropagateTags(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The name or full Amazon Resource Name (ARN) of the IAM role that allows Amazon ECS to make calls to your load balancer on your behalf.
	Role() *string
	SetRole(val *string)
	// The scheduling strategy to use for the service.
	//
	// For more information, see [Services](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs_services.html) .
	SchedulingStrategy() *string
	SetSchedulingStrategy(val *string)
	// The configuration for this service to discover and connect to services, and be discovered by, and connected from, other services within a namespace.
	ServiceConnectConfiguration() interface{}
	SetServiceConnectConfiguration(val interface{})
	// The name of your service.
	ServiceName() *string
	SetServiceName(val *string)
	// The details of the service discovery registry to associate with this service.
	//
	// For more information, see [Service discovery](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html) .
	ServiceRegistries() interface{}
	SetServiceRegistries(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// The metadata that you apply to the service to help you categorize and organize them.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
	// The `family` and `revision` ( `family:revision` ) or full ARN of the task definition to run in your service.
	TaskDefinition() *string
	SetTaskDefinition(val *string)
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// The configuration for a volume specified in the task definition as a volume that is configured at launch time.
	VolumeConfigurations() interface{}
	SetVolumeConfigurations(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnService
type jsiiProxy_CfnService struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnService) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) AttrServiceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrServiceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) CapacityProviderStrategy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityProviderStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) Cluster() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) DeploymentConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) DeploymentController() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentController",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) DesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) EnableEcsManagedTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEcsManagedTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) EnableExecuteCommand() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExecuteCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) HealthCheckGracePeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckGracePeriodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) LaunchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) LoadBalancers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loadBalancers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) NetworkConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) PlacementConstraints() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) PlacementStrategies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementStrategies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) PlatformVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) PropagateTags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) SchedulingStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) ServiceConnectConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceConnectConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) ServiceRegistries() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceRegistries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) TaskDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) VolumeConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeConfigurations",
		&returns,
	)
	return returns
}


func NewCfnService(scope constructs.Construct, id *string, props *CfnServiceProps) CfnService {
	_init_.Initialize()

	if err := validateNewCfnServiceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnService{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.CfnService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnService_Override(c CfnService, scope constructs.Construct, id *string, props *CfnServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.CfnService",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnService)SetCapacityProviderStrategy(val interface{}) {
	if err := j.validateSetCapacityProviderStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityProviderStrategy",
		val,
	)
}

func (j *jsiiProxy_CfnService)SetCluster(val *string) {
	_jsii_.Set(
		j,
		"cluster",
		val,
	)
}

func (j *jsiiProxy_CfnService)SetDeploymentConfiguration(val interface{}) {
	if err := j.validateSetDeploymentConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnService)SetDeploymentController(val interface{}) {
	if err := j.validateSetDeploymentControllerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentController",
		val,
	)
}

func (j *jsiiProxy_CfnService)SetDesiredCount(val *float64) {
	_jsii_.Set(
		j,
		"desiredCount",
		val,
	)
}

func (j *jsiiProxy_CfnService)SetEnableEcsManagedTags(val interface{}) {
	if err := j.validateSetEnableEcsManagedTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableEcsManagedTags",
		val,
	)
}

func (j *jsiiProxy_CfnService)SetEnableExecuteCommand(val interface{}) {
	if err := j.validateSetEnableExecuteCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableExecuteCommand",
		val,
	)
}

func (j *jsiiProxy_CfnService)SetHealthCheckGracePeriodSeconds(val *float64) {
	_jsii_.Set(
		j,
		"healthCheckGracePeriodSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnService)SetLaunchType(val *string) {
	_jsii_.Set(
		j,
		"launchType",
		val,
	)
}

func (j *jsiiProxy_CfnService)SetLoadBalancers(val interface{}) {
	if err := j.validateSetLoadBalancersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancers",
		val,
	)
}

func (j *jsiiProxy_CfnService)SetNetworkConfiguration(val interface{}) {
	if err := j.validateSetNetworkConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnService)SetPlacementConstraints(val interface{}) {
	if err := j.validateSetPlacementConstraintsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementConstraints",
		val,
	)
}

func (j *jsiiProxy_CfnService)SetPlacementStrategies(val interface{}) {
	if err := j.validateSetPlacementStrategiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementStrategies",
		val,
	)
}

func (j *jsiiProxy_CfnService)SetPlatformVersion(val *string) {
	_jsii_.Set(
		j,
		"platformVersion",
		val,
	)
}

func (j *jsiiProxy_CfnService)SetPropagateTags(val *string) {
	_jsii_.Set(
		j,
		"propagateTags",
		val,
	)
}

func (j *jsiiProxy_CfnService)SetRole(val *string) {
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_CfnService)SetSchedulingStrategy(val *string) {
	_jsii_.Set(
		j,
		"schedulingStrategy",
		val,
	)
}

func (j *jsiiProxy_CfnService)SetServiceConnectConfiguration(val interface{}) {
	if err := j.validateSetServiceConnectConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceConnectConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnService)SetServiceName(val *string) {
	_jsii_.Set(
		j,
		"serviceName",
		val,
	)
}

func (j *jsiiProxy_CfnService)SetServiceRegistries(val interface{}) {
	if err := j.validateSetServiceRegistriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRegistries",
		val,
	)
}

func (j *jsiiProxy_CfnService)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnService)SetTaskDefinition(val *string) {
	_jsii_.Set(
		j,
		"taskDefinition",
		val,
	)
}

func (j *jsiiProxy_CfnService)SetVolumeConfigurations(val interface{}) {
	if err := j.validateSetVolumeConfigurationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeConfigurations",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnService_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnService_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnService",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnService_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnService_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnService",
		"isCfnResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnService_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.CfnService",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnService) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnService) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnService) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnService) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnService) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnService) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnService) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnService) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnService) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnService) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnService) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnService) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnService) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnService) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnService) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnService) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnService) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnService) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnService) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

