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

// Creates a capacity provider.
//
// Capacity providers are associated with a cluster and are used in capacity provider strategies to facilitate cluster auto scaling. You can create capacity providers for Amazon ECS Managed Instances and EC2 instances. AWS Fargate has the predefined `FARGATE` and `FARGATE_SPOT` capacity providers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCapacityProvider := awscdk.Aws_ecs.NewCfnCapacityProvider(this, jsii.String("MyCfnCapacityProvider"), &CfnCapacityProviderProps{
//   	AutoScalingGroupProvider: &AutoScalingGroupProviderProperty{
//   		AutoScalingGroupArn: jsii.String("autoScalingGroupArn"),
//
//   		// the properties below are optional
//   		ManagedDraining: jsii.String("managedDraining"),
//   		ManagedScaling: &ManagedScalingProperty{
//   			InstanceWarmupPeriod: jsii.Number(123),
//   			MaximumScalingStepSize: jsii.Number(123),
//   			MinimumScalingStepSize: jsii.Number(123),
//   			Status: jsii.String("status"),
//   			TargetCapacity: jsii.Number(123),
//   		},
//   		ManagedTerminationProtection: jsii.String("managedTerminationProtection"),
//   	},
//   	ClusterName: jsii.String("clusterName"),
//   	ManagedInstancesProvider: &ManagedInstancesProviderProperty{
//   		InfrastructureRoleArn: jsii.String("infrastructureRoleArn"),
//   		InstanceLaunchTemplate: &InstanceLaunchTemplateProperty{
//   			Ec2InstanceProfileArn: jsii.String("ec2InstanceProfileArn"),
//   			NetworkConfiguration: &ManagedInstancesNetworkConfigurationProperty{
//   				SecurityGroups: []*string{
//   					jsii.String("securityGroups"),
//   				},
//   				Subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			CapacityOptionType: jsii.String("capacityOptionType"),
//   			FipsEnabled: jsii.Boolean(false),
//   			InstanceRequirements: &InstanceRequirementsRequestProperty{
//   				MemoryMiB: &MemoryMiBRequestProperty{
//   					Min: jsii.Number(123),
//
//   					// the properties below are optional
//   					Max: jsii.Number(123),
//   				},
//   				VCpuCount: &VCpuCountRangeRequestProperty{
//   					Min: jsii.Number(123),
//
//   					// the properties below are optional
//   					Max: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				AcceleratorCount: &AcceleratorCountRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				AcceleratorManufacturers: []*string{
//   					jsii.String("acceleratorManufacturers"),
//   				},
//   				AcceleratorNames: []*string{
//   					jsii.String("acceleratorNames"),
//   				},
//   				AcceleratorTotalMemoryMiB: &AcceleratorTotalMemoryMiBRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				AcceleratorTypes: []*string{
//   					jsii.String("acceleratorTypes"),
//   				},
//   				AllowedInstanceTypes: []*string{
//   					jsii.String("allowedInstanceTypes"),
//   				},
//   				BareMetal: jsii.String("bareMetal"),
//   				BaselineEbsBandwidthMbps: &BaselineEbsBandwidthMbpsRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				BurstablePerformance: jsii.String("burstablePerformance"),
//   				CpuManufacturers: []*string{
//   					jsii.String("cpuManufacturers"),
//   				},
//   				ExcludedInstanceTypes: []*string{
//   					jsii.String("excludedInstanceTypes"),
//   				},
//   				InstanceGenerations: []*string{
//   					jsii.String("instanceGenerations"),
//   				},
//   				LocalStorage: jsii.String("localStorage"),
//   				LocalStorageTypes: []*string{
//   					jsii.String("localStorageTypes"),
//   				},
//   				MaxSpotPriceAsPercentageOfOptimalOnDemandPrice: jsii.Number(123),
//   				MemoryGiBPerVCpu: &MemoryGiBPerVCpuRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				NetworkBandwidthGbps: &NetworkBandwidthGbpsRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				NetworkInterfaceCount: &NetworkInterfaceCountRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				OnDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   				RequireHibernateSupport: jsii.Boolean(false),
//   				SpotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   				TotalLocalStorageGb: &TotalLocalStorageGBRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   			},
//   			Monitoring: jsii.String("monitoring"),
//   			StorageConfiguration: &ManagedInstancesStorageConfigurationProperty{
//   				StorageSizeGiB: jsii.Number(123),
//   			},
//   		},
//
//   		// the properties below are optional
//   		InfrastructureOptimization: &InfrastructureOptimizationProperty{
//   			ScaleInAfter: jsii.Number(123),
//   		},
//   		PropagateTags: jsii.String("propagateTags"),
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-capacityprovider.html
//
type CfnCapacityProvider interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsecs.ICapacityProviderRef
	awscdk.ITaggable
	// The Auto Scaling group settings for the capacity provider.
	AutoScalingGroupProvider() interface{}
	SetAutoScalingGroupProvider(val interface{})
	// A reference to a CapacityProvider resource.
	CapacityProviderRef() *interfacesawsecs.CapacityProviderReference
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The cluster that this capacity provider is associated with.
	ClusterName() *string
	SetClusterName(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	Env() *interfaces.ResourceEnvironment
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
	// The configuration for the Amazon ECS Managed Instances provider.
	ManagedInstancesProvider() interface{}
	SetManagedInstancesProvider(val interface{})
	// The name of the capacity provider.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// The metadata that you apply to the capacity provider to help you categorize and organize it.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
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

// The jsii proxy struct for CfnCapacityProvider
type jsiiProxy_CfnCapacityProvider struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsecsICapacityProviderRef
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnCapacityProvider) AutoScalingGroupProvider() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoScalingGroupProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapacityProvider) CapacityProviderRef() *interfacesawsecs.CapacityProviderReference {
	var returns *interfacesawsecs.CapacityProviderReference
	_jsii_.Get(
		j,
		"capacityProviderRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapacityProvider) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapacityProvider) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapacityProvider) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapacityProvider) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapacityProvider) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapacityProvider) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapacityProvider) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapacityProvider) ManagedInstancesProvider() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedInstancesProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapacityProvider) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapacityProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapacityProvider) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapacityProvider) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapacityProvider) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapacityProvider) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapacityProvider) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapacityProvider) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::ECS::CapacityProvider`.
func NewCfnCapacityProvider(scope constructs.Construct, id *string, props *CfnCapacityProviderProps) CfnCapacityProvider {
	_init_.Initialize()

	if err := validateNewCfnCapacityProviderParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCapacityProvider{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.CfnCapacityProvider",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ECS::CapacityProvider`.
func NewCfnCapacityProvider_Override(c CfnCapacityProvider, scope constructs.Construct, id *string, props *CfnCapacityProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.CfnCapacityProvider",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCapacityProvider)SetAutoScalingGroupProvider(val interface{}) {
	if err := j.validateSetAutoScalingGroupProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoScalingGroupProvider",
		val,
	)
}

func (j *jsiiProxy_CfnCapacityProvider)SetClusterName(val *string) {
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_CfnCapacityProvider)SetManagedInstancesProvider(val interface{}) {
	if err := j.validateSetManagedInstancesProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedInstancesProvider",
		val,
	)
}

func (j *jsiiProxy_CfnCapacityProvider)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnCapacityProvider)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func CfnCapacityProvider_ArnForCapacityProvider(resource interfacesawsecs.ICapacityProviderRef) *string {
	_init_.Initialize()

	if err := validateCfnCapacityProvider_ArnForCapacityProviderParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnCapacityProvider",
		"arnForCapacityProvider",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// Creates a new ICapacityProviderRef from a capacityProviderName.
func CfnCapacityProvider_FromCapacityProviderName(scope constructs.Construct, id *string, capacityProviderName *string) interfacesawsecs.ICapacityProviderRef {
	_init_.Initialize()

	if err := validateCfnCapacityProvider_FromCapacityProviderNameParameters(scope, id, capacityProviderName); err != nil {
		panic(err)
	}
	var returns interfacesawsecs.ICapacityProviderRef

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnCapacityProvider",
		"fromCapacityProviderName",
		[]interface{}{scope, id, capacityProviderName},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnCapacityProvider.
func CfnCapacityProvider_IsCfnCapacityProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCapacityProvider_IsCfnCapacityProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnCapacityProvider",
		"isCfnCapacityProvider",
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
func CfnCapacityProvider_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCapacityProvider_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnCapacityProvider",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnCapacityProvider_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCapacityProvider_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnCapacityProvider",
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
func CfnCapacityProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCapacityProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnCapacityProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCapacityProvider_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.CfnCapacityProvider",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCapacityProvider) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCapacityProvider) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCapacityProvider) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCapacityProvider) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCapacityProvider) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCapacityProvider) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCapacityProvider) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCapacityProvider) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCapacityProvider) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnCapacityProvider) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnCapacityProvider) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCapacityProvider) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCapacityProvider) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCapacityProvider) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCapacityProvider) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCapacityProvider) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnCapacityProvider) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnCapacityProvider) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCapacityProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCapacityProvider) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

func (c *jsiiProxy_CfnCapacityProvider) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

