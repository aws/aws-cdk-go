package awsapplicationautoscaling

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::ApplicationAutoScaling::ScalableTarget`.
//
// The `AWS::ApplicationAutoScaling::ScalableTarget` resource specifies a resource that Application Auto Scaling can scale, such as an AWS::DynamoDB::Table or AWS::ECS::Service resource.
//
// > If the resource that you want Application Auto Scaling to scale is not yet created in your account, add a dependency on the resource when registering it as a scalable target using the [DependsOn](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) attribute.
//
// For more information, see [RegisterScalableTarget](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html) in the *Application Auto Scaling API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnScalableTarget := awscdk.Aws_applicationautoscaling.NewCfnScalableTarget(this, jsii.String("MyCfnScalableTarget"), &cfnScalableTargetProps{
//   	maxCapacity: jsii.Number(123),
//   	minCapacity: jsii.Number(123),
//   	resourceId: jsii.String("resourceId"),
//   	roleArn: jsii.String("roleArn"),
//   	scalableDimension: jsii.String("scalableDimension"),
//   	serviceNamespace: jsii.String("serviceNamespace"),
//
//   	// the properties below are optional
//   	scheduledActions: []interface{}{
//   		&scheduledActionProperty{
//   			schedule: jsii.String("schedule"),
//   			scheduledActionName: jsii.String("scheduledActionName"),
//
//   			// the properties below are optional
//   			endTime: NewDate(),
//   			scalableTargetAction: &scalableTargetActionProperty{
//   				maxCapacity: jsii.Number(123),
//   				minCapacity: jsii.Number(123),
//   			},
//   			startTime: NewDate(),
//   			timezone: jsii.String("timezone"),
//   		},
//   	},
//   	suspendedState: &suspendedStateProperty{
//   		dynamicScalingInSuspended: jsii.Boolean(false),
//   		dynamicScalingOutSuspended: jsii.Boolean(false),
//   		scheduledScalingSuspended: jsii.Boolean(false),
//   	},
//   })
//
type CfnScalableTarget interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	// The maximum value that you plan to scale out to.
	//
	// When a scaling policy is in effect, Application Auto Scaling can scale out (expand) as needed to the maximum capacity limit in response to changing demand.
	MaxCapacity() *float64
	SetMaxCapacity(val *float64)
	// The minimum value that you plan to scale in to.
	//
	// When a scaling policy is in effect, Application Auto Scaling can scale in (contract) as needed to the minimum capacity limit in response to changing demand.
	MinCapacity() *float64
	SetMinCapacity(val *float64)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The identifier of the resource associated with the scalable target.
	//
	// This string consists of the resource type and unique identifier.
	//
	// - ECS service - The resource type is `service` and the unique identifier is the cluster name and service name. Example: `service/default/sample-webapp` .
	// - Spot Fleet - The resource type is `spot-fleet-request` and the unique identifier is the Spot Fleet request ID. Example: `spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE` .
	// - EMR cluster - The resource type is `instancegroup` and the unique identifier is the cluster ID and instance group ID. Example: `instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0` .
	// - AppStream 2.0 fleet - The resource type is `fleet` and the unique identifier is the fleet name. Example: `fleet/sample-fleet` .
	// - DynamoDB table - The resource type is `table` and the unique identifier is the table name. Example: `table/my-table` .
	// - DynamoDB global secondary index - The resource type is `index` and the unique identifier is the index name. Example: `table/my-table/index/my-table-index` .
	// - Aurora DB cluster - The resource type is `cluster` and the unique identifier is the cluster name. Example: `cluster:my-db-cluster` .
	// - SageMaker endpoint variant - The resource type is `variant` and the unique identifier is the resource ID. Example: `endpoint/my-end-point/variant/KMeansClustering` .
	// - Custom resources are not supported with a resource type. This parameter must specify the `OutputValue` from the CloudFormation template stack used to access the resources. The unique identifier is defined by the service provider. More information is available in our [GitHub repository](https://docs.aws.amazon.com/https://github.com/aws/aws-auto-scaling-custom-resource) .
	// - Amazon Comprehend document classification endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: `arn:aws:comprehend:us-west-2:123456789012:document-classifier-endpoint/EXAMPLE` .
	// - Amazon Comprehend entity recognizer endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: `arn:aws:comprehend:us-west-2:123456789012:entity-recognizer-endpoint/EXAMPLE` .
	// - Lambda provisioned concurrency - The resource type is `function` and the unique identifier is the function name with a function version or alias name suffix that is not `$LATEST` . Example: `function:my-function:prod` or `function:my-function:1` .
	// - Amazon Keyspaces table - The resource type is `table` and the unique identifier is the table name. Example: `keyspace/mykeyspace/table/mytable` .
	// - Amazon MSK cluster - The resource type and unique identifier are specified using the cluster ARN. Example: `arn:aws:kafka:us-east-1:123456789012:cluster/demo-cluster-1/6357e0b2-0e6a-4b86-a0b4-70df934c2e31-5` .
	// - Amazon ElastiCache replication group - The resource type is `replication-group` and the unique identifier is the replication group name. Example: `replication-group/mycluster` .
	// - Neptune cluster - The resource type is `cluster` and the unique identifier is the cluster name. Example: `cluster:mycluster` .
	ResourceId() *string
	SetResourceId(val *string)
	// Specify the Amazon Resource Name (ARN) of an Identity and Access Management (IAM) role that allows Application Auto Scaling to modify the scalable target on your behalf.
	//
	// This can be either an IAM service role that Application Auto Scaling can assume to make calls to other AWS resources on your behalf, or a service-linked role for the specified service. For more information, see [How Application Auto Scaling works with IAM](https://docs.aws.amazon.com/autoscaling/application/userguide/security_iam_service-with-iam.html) in the *Application Auto Scaling User Guide* .
	//
	// To automatically create a service-linked role (recommended), specify the full ARN of the service-linked role in your stack template. To find the exact ARN of the service-linked role for your AWS or custom resource, see the [Service-linked roles](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-service-linked-roles.html) topic in the *Application Auto Scaling User Guide* . Look for the ARN in the table at the bottom of the page.
	RoleArn() *string
	SetRoleArn(val *string)
	// The scalable dimension associated with the scalable target.
	//
	// This string consists of the service namespace, resource type, and scaling property.
	//
	// - `ecs:service:DesiredCount` - The desired task count of an ECS service.
	// - `elasticmapreduce:instancegroup:InstanceCount` - The instance count of an EMR Instance Group.
	// - `ec2:spot-fleet-request:TargetCapacity` - The target capacity of a Spot Fleet.
	// - `appstream:fleet:DesiredCapacity` - The desired capacity of an AppStream 2.0 fleet.
	// - `dynamodb:table:ReadCapacityUnits` - The provisioned read capacity for a DynamoDB table.
	// - `dynamodb:table:WriteCapacityUnits` - The provisioned write capacity for a DynamoDB table.
	// - `dynamodb:index:ReadCapacityUnits` - The provisioned read capacity for a DynamoDB global secondary index.
	// - `dynamodb:index:WriteCapacityUnits` - The provisioned write capacity for a DynamoDB global secondary index.
	// - `rds:cluster:ReadReplicaCount` - The count of Aurora Replicas in an Aurora DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible edition.
	// - `sagemaker:variant:DesiredInstanceCount` - The number of EC2 instances for a SageMaker model endpoint variant.
	// - `custom-resource:ResourceType:Property` - The scalable dimension for a custom resource provided by your own application or service.
	// - `comprehend:document-classifier-endpoint:DesiredInferenceUnits` - The number of inference units for an Amazon Comprehend document classification endpoint.
	// - `comprehend:entity-recognizer-endpoint:DesiredInferenceUnits` - The number of inference units for an Amazon Comprehend entity recognizer endpoint.
	// - `lambda:function:ProvisionedConcurrency` - The provisioned concurrency for a Lambda function.
	// - `cassandra:table:ReadCapacityUnits` - The provisioned read capacity for an Amazon Keyspaces table.
	// - `cassandra:table:WriteCapacityUnits` - The provisioned write capacity for an Amazon Keyspaces table.
	// - `kafka:broker-storage:VolumeSize` - The provisioned volume size (in GiB) for brokers in an Amazon MSK cluster.
	// - `elasticache:replication-group:NodeGroups` - The number of node groups for an Amazon ElastiCache replication group.
	// - `elasticache:replication-group:Replicas` - The number of replicas per node group for an Amazon ElastiCache replication group.
	// - `neptune:cluster:ReadReplicaCount` - The count of read replicas in an Amazon Neptune DB cluster.
	ScalableDimension() *string
	SetScalableDimension(val *string)
	// The scheduled actions for the scalable target. Duplicates aren't allowed.
	//
	// For more information about using scheduled scaling, see [Scheduled scaling](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-scheduled-scaling.html) in the *Application Auto Scaling User Guide* .
	ScheduledActions() interface{}
	SetScheduledActions(val interface{})
	// The namespace of the AWS service that provides the resource, or a `custom-resource` .
	ServiceNamespace() *string
	SetServiceNamespace(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An embedded object that contains attributes and attribute values that are used to suspend and resume automatic scaling.
	//
	// Setting the value of an attribute to `true` suspends the specified scaling activities. Setting it to `false` (default) resumes the specified scaling activities.
	//
	// *Suspension Outcomes*
	//
	// - For `DynamicScalingInSuspended` , while a suspension is in effect, all scale-in activities that are triggered by a scaling policy are suspended.
	// - For `DynamicScalingOutSuspended` , while a suspension is in effect, all scale-out activities that are triggered by a scaling policy are suspended.
	// - For `ScheduledScalingSuspended` , while a suspension is in effect, all scaling activities that involve scheduled actions are suspended.
	//
	// For more information, see [Suspending and resuming scaling](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-suspend-resume-scaling.html) in the *Application Auto Scaling User Guide* .
	SuspendedState() interface{}
	SetSuspendedState(val interface{})
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
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
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
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnScalableTarget
type jsiiProxy_CfnScalableTarget struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnScalableTarget) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) MaxCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) MinCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) ScalableDimension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalableDimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) ScheduledActions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduledActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) ServiceNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) SuspendedState() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suspendedState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApplicationAutoScaling::ScalableTarget`.
func NewCfnScalableTarget(scope constructs.Construct, id *string, props *CfnScalableTargetProps) CfnScalableTarget {
	_init_.Initialize()

	if err := validateNewCfnScalableTargetParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnScalableTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_applicationautoscaling.CfnScalableTarget",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApplicationAutoScaling::ScalableTarget`.
func NewCfnScalableTarget_Override(c CfnScalableTarget, scope constructs.Construct, id *string, props *CfnScalableTargetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_applicationautoscaling.CfnScalableTarget",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnScalableTarget)SetMaxCapacity(val *float64) {
	if err := j.validateSetMaxCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCapacity",
		val,
	)
}

func (j *jsiiProxy_CfnScalableTarget)SetMinCapacity(val *float64) {
	if err := j.validateSetMinCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCapacity",
		val,
	)
}

func (j *jsiiProxy_CfnScalableTarget)SetResourceId(val *string) {
	if err := j.validateSetResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_CfnScalableTarget)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnScalableTarget)SetScalableDimension(val *string) {
	if err := j.validateSetScalableDimensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scalableDimension",
		val,
	)
}

func (j *jsiiProxy_CfnScalableTarget)SetScheduledActions(val interface{}) {
	if err := j.validateSetScheduledActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduledActions",
		val,
	)
}

func (j *jsiiProxy_CfnScalableTarget)SetServiceNamespace(val *string) {
	if err := j.validateSetServiceNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceNamespace",
		val,
	)
}

func (j *jsiiProxy_CfnScalableTarget)SetSuspendedState(val interface{}) {
	if err := j.validateSetSuspendedStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspendedState",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnScalableTarget_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnScalableTarget_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationautoscaling.CfnScalableTarget",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnScalableTarget_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnScalableTarget_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationautoscaling.CfnScalableTarget",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnScalableTarget_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnScalableTarget_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationautoscaling.CfnScalableTarget",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnScalableTarget_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_applicationautoscaling.CfnScalableTarget",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnScalableTarget) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnScalableTarget) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnScalableTarget) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnScalableTarget) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnScalableTarget) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnScalableTarget) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnScalableTarget) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnScalableTarget) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScalableTarget) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnScalableTarget) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnScalableTarget) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnScalableTarget) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnScalableTarget) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScalableTarget) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScalableTarget) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

