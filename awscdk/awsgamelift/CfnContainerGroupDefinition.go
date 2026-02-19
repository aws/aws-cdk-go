package awsgamelift

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgamelift/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsgamelift"
	"github.com/aws/constructs-go/constructs/v10"
)

// The properties that describe a container group resource.
//
// You can update all properties of a container group definition properties. Updates to a container group definition are saved as new versions.
//
// *Used with:* [CreateContainerGroupDefinition](https://docs.aws.amazon.com/gamelift/latest/apireference/API_CreateContainerGroupDefinition.html)
//
// *Returned by:* [DescribeContainerGroupDefinition](https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeContainerGroupDefinition.html) , [ListContainerGroupDefinitions](https://docs.aws.amazon.com/gamelift/latest/apireference/API_ListContainerGroupDefinitions.html) , [UpdateContainerGroupDefinition](https://docs.aws.amazon.com/gamelift/latest/apireference/API_UpdateContainerGroupDefinition.html)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnContainerGroupDefinition := awscdk.Aws_gamelift.NewCfnContainerGroupDefinition(this, jsii.String("MyCfnContainerGroupDefinition"), &CfnContainerGroupDefinitionProps{
//   	Name: jsii.String("name"),
//   	OperatingSystem: jsii.String("operatingSystem"),
//   	TotalMemoryLimitMebibytes: jsii.Number(123),
//   	TotalVcpuLimit: jsii.Number(123),
//
//   	// the properties below are optional
//   	ContainerGroupType: jsii.String("containerGroupType"),
//   	GameServerContainerDefinition: &GameServerContainerDefinitionProperty{
//   		ContainerName: jsii.String("containerName"),
//   		ImageUri: jsii.String("imageUri"),
//   		ServerSdkVersion: jsii.String("serverSdkVersion"),
//
//   		// the properties below are optional
//   		DependsOn: []interface{}{
//   			&ContainerDependencyProperty{
//   				Condition: jsii.String("condition"),
//   				ContainerName: jsii.String("containerName"),
//   			},
//   		},
//   		EnvironmentOverride: []interface{}{
//   			&ContainerEnvironmentProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		MountPoints: []interface{}{
//   			&ContainerMountPointProperty{
//   				InstancePath: jsii.String("instancePath"),
//
//   				// the properties below are optional
//   				AccessLevel: jsii.String("accessLevel"),
//   				ContainerPath: jsii.String("containerPath"),
//   			},
//   		},
//   		PortConfiguration: &PortConfigurationProperty{
//   			ContainerPortRanges: []interface{}{
//   				&ContainerPortRangeProperty{
//   					FromPort: jsii.Number(123),
//   					Protocol: jsii.String("protocol"),
//   					ToPort: jsii.Number(123),
//   				},
//   			},
//   		},
//   		ResolvedImageDigest: jsii.String("resolvedImageDigest"),
//   	},
//   	SourceVersionNumber: jsii.Number(123),
//   	SupportContainerDefinitions: []interface{}{
//   		&SupportContainerDefinitionProperty{
//   			ContainerName: jsii.String("containerName"),
//   			ImageUri: jsii.String("imageUri"),
//
//   			// the properties below are optional
//   			DependsOn: []interface{}{
//   				&ContainerDependencyProperty{
//   					Condition: jsii.String("condition"),
//   					ContainerName: jsii.String("containerName"),
//   				},
//   			},
//   			EnvironmentOverride: []interface{}{
//   				&ContainerEnvironmentProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Essential: jsii.Boolean(false),
//   			HealthCheck: &ContainerHealthCheckProperty{
//   				Command: []*string{
//   					jsii.String("command"),
//   				},
//
//   				// the properties below are optional
//   				Interval: jsii.Number(123),
//   				Retries: jsii.Number(123),
//   				StartPeriod: jsii.Number(123),
//   				Timeout: jsii.Number(123),
//   			},
//   			MemoryHardLimitMebibytes: jsii.Number(123),
//   			MountPoints: []interface{}{
//   				&ContainerMountPointProperty{
//   					InstancePath: jsii.String("instancePath"),
//
//   					// the properties below are optional
//   					AccessLevel: jsii.String("accessLevel"),
//   					ContainerPath: jsii.String("containerPath"),
//   				},
//   			},
//   			PortConfiguration: &PortConfigurationProperty{
//   				ContainerPortRanges: []interface{}{
//   					&ContainerPortRangeProperty{
//   						FromPort: jsii.Number(123),
//   						Protocol: jsii.String("protocol"),
//   						ToPort: jsii.Number(123),
//   					},
//   				},
//   			},
//   			ResolvedImageDigest: jsii.String("resolvedImageDigest"),
//   			Vcpu: jsii.Number(123),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VersionDescription: jsii.String("versionDescription"),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html
//
type CfnContainerGroupDefinition interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsgamelift.IContainerGroupDefinitionRef
	awscdk.ITaggableV2
	// The Amazon Resource Name ( [ARN](https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html) ) that is assigned to an Amazon GameLift Servers `ContainerGroupDefinition` resource. It uniquely identifies the resource across all AWS Regions. Format is `arn:aws:gamelift:[region]::containergroupdefinition/[container group definition name]:[version]` .
	AttrContainerGroupDefinitionArn() *string
	// A time stamp indicating when this data object was created.
	//
	// Format is a number expressed in Unix time as milliseconds (for example `"1469498468.057"` ).
	AttrCreationTime() *string
	// Current status of the container group definition resource. Values include:.
	//
	// - `COPYING` -- Amazon GameLift Servers is in the process of making copies of all container images that are defined in the group. While in this state, the resource can't be used to create a container fleet.
	// - `READY` -- Amazon GameLift Servers has copied the registry images for all containers that are defined in the group. You can use a container group definition in this status to create a container fleet.
	// - `FAILED` -- Amazon GameLift Servers failed to create a valid container group definition resource. For more details on the cause of the failure, see `StatusReason` . A container group definition resource in failed status will be deleted within a few minutes.
	AttrStatus() *string
	// Additional information about a container group definition that's in `FAILED` status. Possible reasons include:.
	//
	// - An internal issue prevented Amazon GameLift Servers from creating the container group definition resource. Delete the failed resource and call [CreateContainerGroupDefinition](https://docs.aws.amazon.com/gamelift/latest/apireference/API_CreateContainerGroupDefinition.html) again.
	// - An access-denied message means that you don't have permissions to access the container image on ECR. See [IAM permission examples](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-iam-policy-examples.html.html) for help setting up required IAM permissions for Amazon GameLift Servers.
	// - The `ImageUri` value for at least one of the containers in the container group definition was invalid or not found in the current AWS account.
	// - At least one of the container images referenced in the container group definition exceeds the allowed size. For size limits, see [Amazon GameLift Servers endpoints and quotas](https://docs.aws.amazon.com/general/latest/gr/gamelift.html) .
	// - At least one of the container images referenced in the container group definition uses a different operating system than the one defined for the container group.
	AttrStatusReason() *string
	// Indicates the version of a particular container group definition.
	//
	// This number is incremented automatically when you update a container group definition. You can view, update, or delete individual versions or the entire container group definition.
	AttrVersionNumber() *float64
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// A reference to a ContainerGroupDefinition resource.
	ContainerGroupDefinitionRef() *interfacesawsgamelift.ContainerGroupDefinitionReference
	// The type of container group.
	ContainerGroupType() *string
	SetContainerGroupType(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	Env() *interfaces.ResourceEnvironment
	// The definition for the game server container in this group.
	GameServerContainerDefinition() interface{}
	SetGameServerContainerDefinition(val interface{})
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
	// A descriptive identifier for the container group definition.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// The platform that all containers in the container group definition run on.
	OperatingSystem() *string
	SetOperatingSystem(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// A specific ContainerGroupDefinition version to be updated.
	SourceVersionNumber() *float64
	SetSourceVersionNumber(val *float64)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The set of definitions for support containers in this group.
	SupportContainerDefinitions() interface{}
	SetSupportContainerDefinitions(val interface{})
	// An array of key-value pairs to apply to this resource.
	Tags() *[]*awscdk.CfnTag
	SetTags(val *[]*awscdk.CfnTag)
	// The amount of memory (in MiB) on a fleet instance to allocate for the container group.
	TotalMemoryLimitMebibytes() *float64
	SetTotalMemoryLimitMebibytes(val *float64)
	// The amount of vCPU units on a fleet instance to allocate for the container group (1 vCPU is equal to 1024 CPU units).
	TotalVcpuLimit() *float64
	SetTotalVcpuLimit(val *float64)
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
	// An optional description that was provided for a container group definition update.
	VersionDescription() *string
	SetVersionDescription(val *string)
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

// The jsii proxy struct for CfnContainerGroupDefinition
type jsiiProxy_CfnContainerGroupDefinition struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsgameliftIContainerGroupDefinitionRef
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnContainerGroupDefinition) AttrContainerGroupDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrContainerGroupDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerGroupDefinition) AttrCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerGroupDefinition) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerGroupDefinition) AttrStatusReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatusReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerGroupDefinition) AttrVersionNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrVersionNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerGroupDefinition) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerGroupDefinition) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerGroupDefinition) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerGroupDefinition) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerGroupDefinition) ContainerGroupDefinitionRef() *interfacesawsgamelift.ContainerGroupDefinitionReference {
	var returns *interfacesawsgamelift.ContainerGroupDefinitionReference
	_jsii_.Get(
		j,
		"containerGroupDefinitionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerGroupDefinition) ContainerGroupType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerGroupType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerGroupDefinition) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerGroupDefinition) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerGroupDefinition) GameServerContainerDefinition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gameServerContainerDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerGroupDefinition) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerGroupDefinition) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerGroupDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerGroupDefinition) OperatingSystem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatingSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerGroupDefinition) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerGroupDefinition) SourceVersionNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sourceVersionNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerGroupDefinition) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerGroupDefinition) SupportContainerDefinitions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportContainerDefinitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerGroupDefinition) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerGroupDefinition) TotalMemoryLimitMebibytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalMemoryLimitMebibytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerGroupDefinition) TotalVcpuLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalVcpuLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerGroupDefinition) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerGroupDefinition) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerGroupDefinition) VersionDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionDescription",
		&returns,
	)
	return returns
}


// Create a new `AWS::GameLift::ContainerGroupDefinition`.
func NewCfnContainerGroupDefinition(scope constructs.Construct, id *string, props *CfnContainerGroupDefinitionProps) CfnContainerGroupDefinition {
	_init_.Initialize()

	if err := validateNewCfnContainerGroupDefinitionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnContainerGroupDefinition{}

	_jsii_.Create(
		"aws-cdk-lib.aws_gamelift.CfnContainerGroupDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GameLift::ContainerGroupDefinition`.
func NewCfnContainerGroupDefinition_Override(c CfnContainerGroupDefinition, scope constructs.Construct, id *string, props *CfnContainerGroupDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_gamelift.CfnContainerGroupDefinition",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnContainerGroupDefinition)SetContainerGroupType(val *string) {
	_jsii_.Set(
		j,
		"containerGroupType",
		val,
	)
}

func (j *jsiiProxy_CfnContainerGroupDefinition)SetGameServerContainerDefinition(val interface{}) {
	if err := j.validateSetGameServerContainerDefinitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gameServerContainerDefinition",
		val,
	)
}

func (j *jsiiProxy_CfnContainerGroupDefinition)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnContainerGroupDefinition)SetOperatingSystem(val *string) {
	if err := j.validateSetOperatingSystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operatingSystem",
		val,
	)
}

func (j *jsiiProxy_CfnContainerGroupDefinition)SetSourceVersionNumber(val *float64) {
	_jsii_.Set(
		j,
		"sourceVersionNumber",
		val,
	)
}

func (j *jsiiProxy_CfnContainerGroupDefinition)SetSupportContainerDefinitions(val interface{}) {
	if err := j.validateSetSupportContainerDefinitionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportContainerDefinitions",
		val,
	)
}

func (j *jsiiProxy_CfnContainerGroupDefinition)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CfnContainerGroupDefinition)SetTotalMemoryLimitMebibytes(val *float64) {
	if err := j.validateSetTotalMemoryLimitMebibytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalMemoryLimitMebibytes",
		val,
	)
}

func (j *jsiiProxy_CfnContainerGroupDefinition)SetTotalVcpuLimit(val *float64) {
	if err := j.validateSetTotalVcpuLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalVcpuLimit",
		val,
	)
}

func (j *jsiiProxy_CfnContainerGroupDefinition)SetVersionDescription(val *string) {
	_jsii_.Set(
		j,
		"versionDescription",
		val,
	)
}

func CfnContainerGroupDefinition_ArnForContainerGroupDefinition(resource interfacesawsgamelift.IContainerGroupDefinitionRef) *string {
	_init_.Initialize()

	if err := validateCfnContainerGroupDefinition_ArnForContainerGroupDefinitionParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_gamelift.CfnContainerGroupDefinition",
		"arnForContainerGroupDefinition",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnContainerGroupDefinition.
func CfnContainerGroupDefinition_IsCfnContainerGroupDefinition(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnContainerGroupDefinition_IsCfnContainerGroupDefinitionParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_gamelift.CfnContainerGroupDefinition",
		"isCfnContainerGroupDefinition",
		[]interface{}{x},
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
func CfnContainerGroupDefinition_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnContainerGroupDefinition_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_gamelift.CfnContainerGroupDefinition",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnContainerGroupDefinition_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnContainerGroupDefinition_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_gamelift.CfnContainerGroupDefinition",
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
func CfnContainerGroupDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnContainerGroupDefinition_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_gamelift.CfnContainerGroupDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnContainerGroupDefinition_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_gamelift.CfnContainerGroupDefinition",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnContainerGroupDefinition) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnContainerGroupDefinition) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnContainerGroupDefinition) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnContainerGroupDefinition) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnContainerGroupDefinition) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnContainerGroupDefinition) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnContainerGroupDefinition) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnContainerGroupDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnContainerGroupDefinition) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnContainerGroupDefinition) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnContainerGroupDefinition) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnContainerGroupDefinition) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnContainerGroupDefinition) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnContainerGroupDefinition) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnContainerGroupDefinition) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnContainerGroupDefinition) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnContainerGroupDefinition) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnContainerGroupDefinition) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnContainerGroupDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnContainerGroupDefinition) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

func (c *jsiiProxy_CfnContainerGroupDefinition) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

