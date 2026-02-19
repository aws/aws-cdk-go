package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsecs"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an Express service that simplifies deploying containerized web applications on Amazon ECS with managed AWS infrastructure.
//
// This operation provisions and configures Application Load Balancers, target groups, security groups, and auto-scaling policies automatically.
//
// Specify a primary container configuration with your application image and basic settings. Amazon ECS creates the necessary AWS resources for traffic distribution, health monitoring, network access control, and capacity management.
//
// Provide an execution role for task operations and an infrastructure role for managing AWS resources on your behalf.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnExpressGatewayService := awscdk.Aws_ecs.NewCfnExpressGatewayService(this, jsii.String("MyCfnExpressGatewayService"), &CfnExpressGatewayServiceProps{
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	InfrastructureRoleArn: jsii.String("infrastructureRoleArn"),
//   	PrimaryContainer: &ExpressGatewayContainerProperty{
//   		Image: jsii.String("image"),
//
//   		// the properties below are optional
//   		AwsLogsConfiguration: &ExpressGatewayServiceAwsLogsConfigurationProperty{
//   			LogGroup: jsii.String("logGroup"),
//   			LogStreamPrefix: jsii.String("logStreamPrefix"),
//   		},
//   		Command: []*string{
//   			jsii.String("command"),
//   		},
//   		ContainerPort: jsii.Number(123),
//   		Environment: []interface{}{
//   			&KeyValuePairProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		RepositoryCredentials: &ExpressGatewayRepositoryCredentialsProperty{
//   			CredentialsParameter: jsii.String("credentialsParameter"),
//   		},
//   		Secrets: []interface{}{
//   			&SecretProperty{
//   				Name: jsii.String("name"),
//   				ValueFrom: jsii.String("valueFrom"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	Cluster: jsii.String("cluster"),
//   	Cpu: jsii.String("cpu"),
//   	HealthCheckPath: jsii.String("healthCheckPath"),
//   	Memory: jsii.String("memory"),
//   	NetworkConfiguration: &ExpressGatewayServiceNetworkConfigurationProperty{
//   		SecurityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   	ScalingTarget: &ExpressGatewayScalingTargetProperty{
//   		AutoScalingMetric: jsii.String("autoScalingMetric"),
//   		AutoScalingTargetValue: jsii.Number(123),
//   		MaxTaskCount: jsii.Number(123),
//   		MinTaskCount: jsii.Number(123),
//   	},
//   	ServiceName: jsii.String("serviceName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TaskRoleArn: jsii.String("taskRoleArn"),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-expressgatewayservice.html
//
type CfnExpressGatewayService interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsecs.IExpressGatewayServiceRef
	awscdk.ITaggableV2
	// The list of active service configurations for the Express service.
	AttrActiveConfigurations() awscdk.IResolvable
	// The Unix timestamp for when the Express service was created.
	AttrCreatedAt() *string
	AttrEcsManagedResourceArns() awscdk.IResolvable
	AttrEcsManagedResourceArnsAutoScaling() awscdk.IResolvable
	// The list of Auto Scaling policy ARNs associated with the express service.
	AttrEcsManagedResourceArnsAutoScalingApplicationAutoScalingPolicies() *[]*string
	// The Auto Scaling Scalable Target ARN associated with the express service.
	AttrEcsManagedResourceArnsAutoScalingScalableTarget() *string
	AttrEcsManagedResourceArnsIngressPath() awscdk.IResolvable
	// The Certificate ARN associated with the express service.
	AttrEcsManagedResourceArnsIngressPathCertificateArn() *string
	// The ARN of the Load Balancer listener associated with the express service.
	AttrEcsManagedResourceArnsIngressPathListenerArn() *string
	// The ARN of the Load Balancer listener rule associated with the express service.
	AttrEcsManagedResourceArnsIngressPathListenerRuleArn() *string
	// The ARN of the Load Balancer associated with the express service.
	AttrEcsManagedResourceArnsIngressPathLoadBalancerArn() *string
	// The list of Load Balancer Security Group ARNs associated with the express service.
	AttrEcsManagedResourceArnsIngressPathLoadBalancerSecurityGroups() *[]*string
	// The list of Target Group ARNs associated with the express service.
	AttrEcsManagedResourceArnsIngressPathTargetGroupArns() *[]*string
	// The list of Log Group ARNs associated with the express service.
	AttrEcsManagedResourceArnsLogGroups() *[]*string
	// The list of Metric Alarm ARNs associated with the express service.
	AttrEcsManagedResourceArnsMetricAlarms() *[]*string
	// The list of Security Group ARNs associated with the express service.
	AttrEcsManagedResourceArnsServiceSecurityGroups() *[]*string
	// The Endpoint of the express service.
	AttrEndpoint() *string
	// The ARN that identifies the Express service.
	AttrServiceArn() *string
	AttrStatus() awscdk.IResolvable
	// The Unix timestamp for when the Express service was last updated.
	AttrUpdatedAt() *string
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The short name or full ARN of the cluster that hosts the Express service.
	Cluster() *string
	SetCluster(val *string)
	// The CPU allocation for tasks in this service revision.
	Cpu() *string
	SetCpu(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	Env() *interfaces.ResourceEnvironment
	// The ARN of the task execution role for the service revision.
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	// A reference to a ExpressGatewayService resource.
	ExpressGatewayServiceRef() *interfacesawsecs.ExpressGatewayServiceReference
	// The health check path for this service revision.
	HealthCheckPath() *string
	SetHealthCheckPath(val *string)
	// The ARN of the infrastructure role that manages AWS resources for the Express service.
	InfrastructureRoleArn() *string
	SetInfrastructureRoleArn(val *string)
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
	// The memory allocation for tasks in this service revision.
	Memory() *string
	SetMemory(val *string)
	// The network configuration for tasks in this service revision.
	NetworkConfiguration() interface{}
	SetNetworkConfiguration(val interface{})
	// The tree node.
	Node() constructs.Node
	// The primary container configuration for this service revision.
	PrimaryContainer() interface{}
	SetPrimaryContainer(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The auto-scaling configuration for this service revision.
	ScalingTarget() interface{}
	SetScalingTarget(val interface{})
	// The name of the Express service.
	ServiceName() *string
	SetServiceName(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The metadata applied to the Express service.
	Tags() *[]*awscdk.CfnTag
	SetTags(val *[]*awscdk.CfnTag)
	// The ARN of the task role for the service revision.
	TaskRoleArn() *string
	SetTaskRoleArn(val *string)
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
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for CfnExpressGatewayService
type jsiiProxy_CfnExpressGatewayService struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsecsIExpressGatewayServiceRef
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnExpressGatewayService) AttrActiveConfigurations() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrActiveConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) AttrCreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) AttrEcsManagedResourceArns() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrEcsManagedResourceArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) AttrEcsManagedResourceArnsAutoScaling() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrEcsManagedResourceArnsAutoScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) AttrEcsManagedResourceArnsAutoScalingApplicationAutoScalingPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrEcsManagedResourceArnsAutoScalingApplicationAutoScalingPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) AttrEcsManagedResourceArnsAutoScalingScalableTarget() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEcsManagedResourceArnsAutoScalingScalableTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) AttrEcsManagedResourceArnsIngressPath() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrEcsManagedResourceArnsIngressPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) AttrEcsManagedResourceArnsIngressPathCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEcsManagedResourceArnsIngressPathCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) AttrEcsManagedResourceArnsIngressPathListenerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEcsManagedResourceArnsIngressPathListenerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) AttrEcsManagedResourceArnsIngressPathListenerRuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEcsManagedResourceArnsIngressPathListenerRuleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) AttrEcsManagedResourceArnsIngressPathLoadBalancerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEcsManagedResourceArnsIngressPathLoadBalancerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) AttrEcsManagedResourceArnsIngressPathLoadBalancerSecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrEcsManagedResourceArnsIngressPathLoadBalancerSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) AttrEcsManagedResourceArnsIngressPathTargetGroupArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrEcsManagedResourceArnsIngressPathTargetGroupArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) AttrEcsManagedResourceArnsLogGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrEcsManagedResourceArnsLogGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) AttrEcsManagedResourceArnsMetricAlarms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrEcsManagedResourceArnsMetricAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) AttrEcsManagedResourceArnsServiceSecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrEcsManagedResourceArnsServiceSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) AttrEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) AttrServiceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrServiceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) AttrStatus() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) AttrUpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUpdatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) Cluster() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) Cpu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) ExpressGatewayServiceRef() *interfacesawsecs.ExpressGatewayServiceReference {
	var returns *interfacesawsecs.ExpressGatewayServiceReference
	_jsii_.Get(
		j,
		"expressGatewayServiceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) HealthCheckPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) InfrastructureRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) Memory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) NetworkConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) PrimaryContainer() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primaryContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) ScalingTarget() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) TaskRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayService) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::ECS::ExpressGatewayService`.
func NewCfnExpressGatewayService(scope constructs.Construct, id *string, props *CfnExpressGatewayServiceProps) CfnExpressGatewayService {
	_init_.Initialize()

	if err := validateNewCfnExpressGatewayServiceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnExpressGatewayService{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.CfnExpressGatewayService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ECS::ExpressGatewayService`.
func NewCfnExpressGatewayService_Override(c CfnExpressGatewayService, scope constructs.Construct, id *string, props *CfnExpressGatewayServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.CfnExpressGatewayService",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnExpressGatewayService)SetCluster(val *string) {
	_jsii_.Set(
		j,
		"cluster",
		val,
	)
}

func (j *jsiiProxy_CfnExpressGatewayService)SetCpu(val *string) {
	_jsii_.Set(
		j,
		"cpu",
		val,
	)
}

func (j *jsiiProxy_CfnExpressGatewayService)SetExecutionRoleArn(val *string) {
	if err := j.validateSetExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnExpressGatewayService)SetHealthCheckPath(val *string) {
	_jsii_.Set(
		j,
		"healthCheckPath",
		val,
	)
}

func (j *jsiiProxy_CfnExpressGatewayService)SetInfrastructureRoleArn(val *string) {
	if err := j.validateSetInfrastructureRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"infrastructureRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnExpressGatewayService)SetMemory(val *string) {
	_jsii_.Set(
		j,
		"memory",
		val,
	)
}

func (j *jsiiProxy_CfnExpressGatewayService)SetNetworkConfiguration(val interface{}) {
	if err := j.validateSetNetworkConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnExpressGatewayService)SetPrimaryContainer(val interface{}) {
	if err := j.validateSetPrimaryContainerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryContainer",
		val,
	)
}

func (j *jsiiProxy_CfnExpressGatewayService)SetScalingTarget(val interface{}) {
	if err := j.validateSetScalingTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scalingTarget",
		val,
	)
}

func (j *jsiiProxy_CfnExpressGatewayService)SetServiceName(val *string) {
	_jsii_.Set(
		j,
		"serviceName",
		val,
	)
}

func (j *jsiiProxy_CfnExpressGatewayService)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CfnExpressGatewayService)SetTaskRoleArn(val *string) {
	_jsii_.Set(
		j,
		"taskRoleArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnExpressGatewayService_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnExpressGatewayService_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnExpressGatewayService",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnExpressGatewayService.
func CfnExpressGatewayService_IsCfnExpressGatewayService(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnExpressGatewayService_IsCfnExpressGatewayServiceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnExpressGatewayService",
		"isCfnExpressGatewayService",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnExpressGatewayService_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnExpressGatewayService_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnExpressGatewayService",
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
func CfnExpressGatewayService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnExpressGatewayService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnExpressGatewayService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnExpressGatewayService_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.CfnExpressGatewayService",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnExpressGatewayService) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnExpressGatewayService) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnExpressGatewayService) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnExpressGatewayService) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnExpressGatewayService) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnExpressGatewayService) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnExpressGatewayService) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnExpressGatewayService) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnExpressGatewayService) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnExpressGatewayService) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnExpressGatewayService) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnExpressGatewayService) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnExpressGatewayService) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnExpressGatewayService) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnExpressGatewayService) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnExpressGatewayService) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnExpressGatewayService) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnExpressGatewayService) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnExpressGatewayService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnExpressGatewayService) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

func (c *jsiiProxy_CfnExpressGatewayService) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"with",
		args,
		&returns,
	)

	return returns
}

