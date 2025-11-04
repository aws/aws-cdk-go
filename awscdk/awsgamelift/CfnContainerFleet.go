package awsgamelift

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgamelift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Describes an Amazon GameLift Servers managed container fleet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnContainerFleet := awscdk.Aws_gamelift.NewCfnContainerFleet(this, jsii.String("MyCfnContainerFleet"), &CfnContainerFleetProps{
//   	FleetRoleArn: jsii.String("fleetRoleArn"),
//
//   	// the properties below are optional
//   	BillingType: jsii.String("billingType"),
//   	DeploymentConfiguration: &DeploymentConfigurationProperty{
//   		ImpairmentStrategy: jsii.String("impairmentStrategy"),
//   		MinimumHealthyPercentage: jsii.Number(123),
//   		ProtectionStrategy: jsii.String("protectionStrategy"),
//   	},
//   	Description: jsii.String("description"),
//   	GameServerContainerGroupDefinitionName: jsii.String("gameServerContainerGroupDefinitionName"),
//   	GameServerContainerGroupsPerInstance: jsii.Number(123),
//   	GameSessionCreationLimitPolicy: &GameSessionCreationLimitPolicyProperty{
//   		NewGameSessionsPerCreator: jsii.Number(123),
//   		PolicyPeriodInMinutes: jsii.Number(123),
//   	},
//   	InstanceConnectionPortRange: &ConnectionPortRangeProperty{
//   		FromPort: jsii.Number(123),
//   		ToPort: jsii.Number(123),
//   	},
//   	InstanceInboundPermissions: []interface{}{
//   		&IpPermissionProperty{
//   			FromPort: jsii.Number(123),
//   			IpRange: jsii.String("ipRange"),
//   			Protocol: jsii.String("protocol"),
//   			ToPort: jsii.Number(123),
//   		},
//   	},
//   	InstanceType: jsii.String("instanceType"),
//   	Locations: []interface{}{
//   		&LocationConfigurationProperty{
//   			Location: jsii.String("location"),
//
//   			// the properties below are optional
//   			LocationCapacity: &LocationCapacityProperty{
//   				MaxSize: jsii.Number(123),
//   				MinSize: jsii.Number(123),
//
//   				// the properties below are optional
//   				DesiredEc2Instances: jsii.Number(123),
//   			},
//   			StoppedActions: []*string{
//   				jsii.String("stoppedActions"),
//   			},
//   		},
//   	},
//   	LogConfiguration: &LogConfigurationProperty{
//   		LogDestination: jsii.String("logDestination"),
//   		LogGroupArn: jsii.String("logGroupArn"),
//   		S3BucketName: jsii.String("s3BucketName"),
//   	},
//   	MetricGroups: []*string{
//   		jsii.String("metricGroups"),
//   	},
//   	NewGameSessionProtectionPolicy: jsii.String("newGameSessionProtectionPolicy"),
//   	PerInstanceContainerGroupDefinitionName: jsii.String("perInstanceContainerGroupDefinitionName"),
//   	ScalingPolicies: []interface{}{
//   		&ScalingPolicyProperty{
//   			MetricName: jsii.String("metricName"),
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			ComparisonOperator: jsii.String("comparisonOperator"),
//   			EvaluationPeriods: jsii.Number(123),
//   			PolicyType: jsii.String("policyType"),
//   			ScalingAdjustment: jsii.Number(123),
//   			ScalingAdjustmentType: jsii.String("scalingAdjustmentType"),
//   			TargetConfiguration: &TargetConfigurationProperty{
//   				TargetValue: jsii.Number(123),
//   			},
//   			Threshold: jsii.Number(123),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html
//
type CfnContainerFleet interface {
	awscdk.CfnResource
	IContainerFleetRef
	awscdk.IInspectable
	awscdk.ITaggableV2
	// A time stamp indicating when this data object was created.
	//
	// Format is a number expressed in Unix time as milliseconds (for example `"1469498468.057"` ).
	AttrCreationTime() *string
	// Provides information about the last deployment ID and its status.
	AttrDeploymentDetails() awscdk.IResolvable
	// The Amazon Resource Name ( [ARN](https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html) ) that is assigned to a Amazon GameLift Servers fleet resource and uniquely identifies it. ARNs are unique across all Regions. Format is `arn:aws:gamelift:<region>::fleet/fleet-a1234567-b8c9-0d1e-2fa3-b45c6d7e8912` . In a GameLift fleet ARN, the resource ID matches the `FleetId` value.
	AttrFleetArn() *string
	// A unique identifier for the container fleet to retrieve.
	AttrFleetId() *string
	// The Amazon Resource Name ( [ARN](https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html) ) that is assigned to the fleet's game server container group. The ARN value also identifies the specific container group definition version in use.
	AttrGameServerContainerGroupDefinitionArn() *string
	// The calculated maximum number of game server container group that can be deployed on each fleet instance.
	//
	// The calculation depends on the resource needs of the container group and the CPU and memory resources of the fleet's instance type.
	AttrMaximumGameServerContainerGroupsPerInstance() *float64
	// The Amazon Resource Name ( [ARN](https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html) ) that is assigned to the fleet's per-instance container group. The ARN value also identifies the specific container group definition version in use.
	AttrPerInstanceContainerGroupDefinitionArn() *string
	// The current status of the container fleet.
	//
	// - `PENDING` -- A new container fleet has been requested.
	// - `CREATING` -- A new container fleet resource is being created.
	// - `CREATED` -- A new container fleet resource has been created. No fleet instances have been deployed.
	// - `ACTIVATING` -- New container fleet instances are being deployed.
	// - `ACTIVE` -- The container fleet has been deployed and is ready to host game sessions.
	// - `UPDATING` -- Updates to the container fleet is being updated. A deployment is in progress.
	AttrStatus() *string
	// Indicates whether the fleet uses On-Demand or Spot instances for this fleet.
	BillingType() *string
	SetBillingType(val *string)
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// A reference to a ContainerFleet resource.
	ContainerFleetRef() *ContainerFleetReference
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Set of rules for processing a deployment for a container fleet update.
	DeploymentConfiguration() interface{}
	SetDeploymentConfiguration(val interface{})
	// A meaningful description of the container fleet.
	Description() *string
	SetDescription(val *string)
	Env() *awscdk.ResourceEnvironment
	// The unique identifier for an AWS Identity and Access Management (IAM) role with permissions to run your containers on resources that are managed by Amazon GameLift Servers.
	FleetRoleArn() *string
	SetFleetRoleArn(val *string)
	// The name of the fleet's game server container group definition, which describes how to deploy containers with your game server build and support software onto each fleet instance.
	GameServerContainerGroupDefinitionName() *string
	SetGameServerContainerGroupDefinitionName(val *string)
	// The number of times to replicate the game server container group on each fleet instance.
	GameServerContainerGroupsPerInstance() *float64
	SetGameServerContainerGroupsPerInstance(val *float64)
	// A policy that limits the number of game sessions that each individual player can create on instances in this fleet.
	GameSessionCreationLimitPolicy() interface{}
	SetGameSessionCreationLimitPolicy(val interface{})
	// The set of port numbers to open on each instance in a container fleet.
	InstanceConnectionPortRange() interface{}
	SetInstanceConnectionPortRange(val interface{})
	// The IP address ranges and port settings that allow inbound traffic to access game server processes and other processes on this fleet.
	InstanceInboundPermissions() interface{}
	SetInstanceInboundPermissions(val interface{})
	// The Amazon EC2 instance type to use for all instances in the fleet.
	InstanceType() *string
	SetInstanceType(val *string)
	Locations() interface{}
	SetLocations(val interface{})
	// The method that is used to collect container logs for the fleet.
	LogConfiguration() interface{}
	SetLogConfiguration(val interface{})
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
	// The name of an AWS CloudWatch metric group to add this fleet to.
	MetricGroups() *[]*string
	SetMetricGroups(val *[]*string)
	// Determines whether Amazon GameLift Servers can shut down game sessions on the fleet that are actively running and hosting players.
	NewGameSessionProtectionPolicy() *string
	SetNewGameSessionProtectionPolicy(val *string)
	// The tree node.
	Node() constructs.Node
	// The name of the fleet's per-instance container group definition.
	PerInstanceContainerGroupDefinitionName() *string
	SetPerInstanceContainerGroupDefinitionName(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// A list of rules that control how a fleet is scaled.
	ScalingPolicies() interface{}
	SetScalingPolicies(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
	Tags() *[]*awscdk.CfnTag
	SetTags(val *[]*awscdk.CfnTag)
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
}

// The jsii proxy struct for CfnContainerFleet
type jsiiProxy_CfnContainerFleet struct {
	internal.Type__awscdkCfnResource
	jsiiProxy_IContainerFleetRef
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnContainerFleet) AttrCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) AttrDeploymentDetails() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrDeploymentDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) AttrFleetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFleetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) AttrFleetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFleetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) AttrGameServerContainerGroupDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrGameServerContainerGroupDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) AttrMaximumGameServerContainerGroupsPerInstance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrMaximumGameServerContainerGroupsPerInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) AttrPerInstanceContainerGroupDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPerInstanceContainerGroupDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) BillingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) ContainerFleetRef() *ContainerFleetReference {
	var returns *ContainerFleetReference
	_jsii_.Get(
		j,
		"containerFleetRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) DeploymentConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) FleetRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) GameServerContainerGroupDefinitionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gameServerContainerGroupDefinitionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) GameServerContainerGroupsPerInstance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gameServerContainerGroupsPerInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) GameSessionCreationLimitPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gameSessionCreationLimitPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) InstanceConnectionPortRange() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceConnectionPortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) InstanceInboundPermissions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceInboundPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) Locations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"locations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) LogConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) MetricGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metricGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) NewGameSessionProtectionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newGameSessionProtectionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) PerInstanceContainerGroupDefinitionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"perInstanceContainerGroupDefinitionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) ScalingPolicies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleet) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnContainerFleet(scope constructs.Construct, id *string, props *CfnContainerFleetProps) CfnContainerFleet {
	_init_.Initialize()

	if err := validateNewCfnContainerFleetParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnContainerFleet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_gamelift.CfnContainerFleet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnContainerFleet_Override(c CfnContainerFleet, scope constructs.Construct, id *string, props *CfnContainerFleetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_gamelift.CfnContainerFleet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnContainerFleet)SetBillingType(val *string) {
	_jsii_.Set(
		j,
		"billingType",
		val,
	)
}

func (j *jsiiProxy_CfnContainerFleet)SetDeploymentConfiguration(val interface{}) {
	if err := j.validateSetDeploymentConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnContainerFleet)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnContainerFleet)SetFleetRoleArn(val *string) {
	if err := j.validateSetFleetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fleetRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnContainerFleet)SetGameServerContainerGroupDefinitionName(val *string) {
	_jsii_.Set(
		j,
		"gameServerContainerGroupDefinitionName",
		val,
	)
}

func (j *jsiiProxy_CfnContainerFleet)SetGameServerContainerGroupsPerInstance(val *float64) {
	_jsii_.Set(
		j,
		"gameServerContainerGroupsPerInstance",
		val,
	)
}

func (j *jsiiProxy_CfnContainerFleet)SetGameSessionCreationLimitPolicy(val interface{}) {
	if err := j.validateSetGameSessionCreationLimitPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gameSessionCreationLimitPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnContainerFleet)SetInstanceConnectionPortRange(val interface{}) {
	if err := j.validateSetInstanceConnectionPortRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceConnectionPortRange",
		val,
	)
}

func (j *jsiiProxy_CfnContainerFleet)SetInstanceInboundPermissions(val interface{}) {
	if err := j.validateSetInstanceInboundPermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceInboundPermissions",
		val,
	)
}

func (j *jsiiProxy_CfnContainerFleet)SetInstanceType(val *string) {
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_CfnContainerFleet)SetLocations(val interface{}) {
	if err := j.validateSetLocationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locations",
		val,
	)
}

func (j *jsiiProxy_CfnContainerFleet)SetLogConfiguration(val interface{}) {
	if err := j.validateSetLogConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnContainerFleet)SetMetricGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"metricGroups",
		val,
	)
}

func (j *jsiiProxy_CfnContainerFleet)SetNewGameSessionProtectionPolicy(val *string) {
	_jsii_.Set(
		j,
		"newGameSessionProtectionPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnContainerFleet)SetPerInstanceContainerGroupDefinitionName(val *string) {
	_jsii_.Set(
		j,
		"perInstanceContainerGroupDefinitionName",
		val,
	)
}

func (j *jsiiProxy_CfnContainerFleet)SetScalingPolicies(val interface{}) {
	if err := j.validateSetScalingPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scalingPolicies",
		val,
	)
}

func (j *jsiiProxy_CfnContainerFleet)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Creates a new IContainerFleetRef from a fleetId.
func CfnContainerFleet_FromFleetId(scope constructs.Construct, id *string, fleetId *string) IContainerFleetRef {
	_init_.Initialize()

	if err := validateCfnContainerFleet_FromFleetIdParameters(scope, id, fleetId); err != nil {
		panic(err)
	}
	var returns IContainerFleetRef

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_gamelift.CfnContainerFleet",
		"fromFleetId",
		[]interface{}{scope, id, fleetId},
		&returns,
	)

	return returns
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnContainerFleet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnContainerFleet_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_gamelift.CfnContainerFleet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnContainerFleet_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnContainerFleet_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_gamelift.CfnContainerFleet",
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
func CfnContainerFleet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnContainerFleet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_gamelift.CfnContainerFleet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnContainerFleet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_gamelift.CfnContainerFleet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnContainerFleet) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnContainerFleet) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnContainerFleet) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnContainerFleet) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnContainerFleet) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnContainerFleet) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnContainerFleet) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnContainerFleet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnContainerFleet) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnContainerFleet) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnContainerFleet) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnContainerFleet) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnContainerFleet) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnContainerFleet) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnContainerFleet) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnContainerFleet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnContainerFleet) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnContainerFleet) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnContainerFleet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnContainerFleet) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

