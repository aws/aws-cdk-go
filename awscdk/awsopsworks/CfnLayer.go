package awsopsworks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopsworks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var customJson interface{}
//
//   cfnLayer := awscdk.Aws_opsworks.NewCfnLayer(this, jsii.String("MyCfnLayer"), &CfnLayerProps{
//   	AutoAssignElasticIps: jsii.Boolean(false),
//   	AutoAssignPublicIps: jsii.Boolean(false),
//   	EnableAutoHealing: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   	Shortname: jsii.String("shortname"),
//   	StackId: jsii.String("stackId"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Attributes: map[string]*string{
//   		"attributesKey": jsii.String("attributes"),
//   	},
//   	CustomInstanceProfileArn: jsii.String("customInstanceProfileArn"),
//   	CustomJson: customJson,
//   	CustomRecipes: &RecipesProperty{
//   		Configure: []*string{
//   			jsii.String("configure"),
//   		},
//   		Deploy: []*string{
//   			jsii.String("deploy"),
//   		},
//   		Setup: []*string{
//   			jsii.String("setup"),
//   		},
//   		Shutdown: []*string{
//   			jsii.String("shutdown"),
//   		},
//   		Undeploy: []*string{
//   			jsii.String("undeploy"),
//   		},
//   	},
//   	CustomSecurityGroupIds: []*string{
//   		jsii.String("customSecurityGroupIds"),
//   	},
//   	InstallUpdatesOnBoot: jsii.Boolean(false),
//   	LifecycleEventConfiguration: &LifecycleEventConfigurationProperty{
//   		ShutdownEventConfiguration: &ShutdownEventConfigurationProperty{
//   			DelayUntilElbConnectionsDrained: jsii.Boolean(false),
//   			ExecutionTimeout: jsii.Number(123),
//   		},
//   	},
//   	LoadBasedAutoScaling: &LoadBasedAutoScalingProperty{
//   		DownScaling: &AutoScalingThresholdsProperty{
//   			CpuThreshold: jsii.Number(123),
//   			IgnoreMetricsTime: jsii.Number(123),
//   			InstanceCount: jsii.Number(123),
//   			LoadThreshold: jsii.Number(123),
//   			MemoryThreshold: jsii.Number(123),
//   			ThresholdsWaitTime: jsii.Number(123),
//   		},
//   		Enable: jsii.Boolean(false),
//   		UpScaling: &AutoScalingThresholdsProperty{
//   			CpuThreshold: jsii.Number(123),
//   			IgnoreMetricsTime: jsii.Number(123),
//   			InstanceCount: jsii.Number(123),
//   			LoadThreshold: jsii.Number(123),
//   			MemoryThreshold: jsii.Number(123),
//   			ThresholdsWaitTime: jsii.Number(123),
//   		},
//   	},
//   	Packages: []*string{
//   		jsii.String("packages"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UseEbsOptimizedInstances: jsii.Boolean(false),
//   	VolumeConfigurations: []interface{}{
//   		&VolumeConfigurationProperty{
//   			Encrypted: jsii.Boolean(false),
//   			Iops: jsii.Number(123),
//   			MountPoint: jsii.String("mountPoint"),
//   			NumberOfDisks: jsii.Number(123),
//   			RaidLevel: jsii.Number(123),
//   			Size: jsii.Number(123),
//   			VolumeType: jsii.String("volumeType"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html
//
type CfnLayer interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
	// One or more user-defined key-value pairs to be added to the stack attributes.
	Attributes() interface{}
	SetAttributes(val interface{})
	AttrId() *string
	// Whether to automatically assign an [Elastic IP address](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html) to the layer's instances. For more information, see [How to Edit a Layer](https://docs.aws.amazon.com/opsworks/latest/userguide/workinglayers-basics-edit.html) .
	AutoAssignElasticIps() interface{}
	SetAutoAssignElasticIps(val interface{})
	// For stacks that are running in a VPC, whether to automatically assign a public IP address to the layer's instances.
	AutoAssignPublicIps() interface{}
	SetAutoAssignPublicIps(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The ARN of an IAM profile to be used for the layer's EC2 instances.
	CustomInstanceProfileArn() *string
	SetCustomInstanceProfileArn(val *string)
	// A JSON-formatted string containing custom stack configuration and deployment attributes to be installed on the layer's instances.
	CustomJson() interface{}
	SetCustomJson(val interface{})
	// A `LayerCustomRecipes` object that specifies the layer custom recipes.
	CustomRecipes() interface{}
	SetCustomRecipes(val interface{})
	// An array containing the layer custom security group IDs.
	CustomSecurityGroupIds() *[]*string
	SetCustomSecurityGroupIds(val *[]*string)
	// Whether to disable auto healing for the layer.
	EnableAutoHealing() interface{}
	SetEnableAutoHealing(val interface{})
	// Whether to install operating system and package updates when the instance boots.
	InstallUpdatesOnBoot() interface{}
	SetInstallUpdatesOnBoot(val interface{})
	// A `LifeCycleEventConfiguration` object that you can use to configure the Shutdown event to specify an execution timeout and enable or disable Elastic Load Balancer connection draining.
	LifecycleEventConfiguration() interface{}
	SetLifecycleEventConfiguration(val interface{})
	// The load-based scaling configuration for the AWS OpsWorks layer.
	LoadBasedAutoScaling() interface{}
	SetLoadBasedAutoScaling(val interface{})
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
	// The layer name, which is used by the console.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// An array of `Package` objects that describes the layer packages.
	Packages() *[]*string
	SetPackages(val *[]*string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// For custom layers only, use this parameter to specify the layer's short name, which is used internally by AWS OpsWorks Stacks and by Chef recipes.
	Shortname() *string
	SetShortname(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The layer stack ID.
	StackId() *string
	SetStackId(val *string)
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// Specifies one or more sets of tags (key–value pairs) to associate with this AWS OpsWorks layer.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
	// The layer type.
	Type() *string
	SetType(val *string)
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
	// Whether to use Amazon EBS-optimized instances.
	UseEbsOptimizedInstances() interface{}
	SetUseEbsOptimizedInstances(val interface{})
	// A `VolumeConfigurations` object that describes the layer's Amazon EBS volumes.
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

// The jsii proxy struct for CfnLayer
type jsiiProxy_CfnLayer struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnLayer) Attributes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) AutoAssignElasticIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignElasticIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) AutoAssignPublicIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignPublicIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) CustomInstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customInstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) CustomJson() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) CustomRecipes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) CustomSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) EnableAutoHealing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutoHealing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) InstallUpdatesOnBoot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) LifecycleEventConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lifecycleEventConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) LoadBasedAutoScaling() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loadBasedAutoScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) Packages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"packages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) Shortname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shortname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) UseEbsOptimizedInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useEbsOptimizedInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayer) VolumeConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeConfigurations",
		&returns,
	)
	return returns
}


func NewCfnLayer(scope constructs.Construct, id *string, props *CfnLayerProps) CfnLayer {
	_init_.Initialize()

	if err := validateNewCfnLayerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLayer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_opsworks.CfnLayer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnLayer_Override(c CfnLayer, scope constructs.Construct, id *string, props *CfnLayerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_opsworks.CfnLayer",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLayer)SetAttributes(val interface{}) {
	if err := j.validateSetAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributes",
		val,
	)
}

func (j *jsiiProxy_CfnLayer)SetAutoAssignElasticIps(val interface{}) {
	if err := j.validateSetAutoAssignElasticIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoAssignElasticIps",
		val,
	)
}

func (j *jsiiProxy_CfnLayer)SetAutoAssignPublicIps(val interface{}) {
	if err := j.validateSetAutoAssignPublicIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoAssignPublicIps",
		val,
	)
}

func (j *jsiiProxy_CfnLayer)SetCustomInstanceProfileArn(val *string) {
	_jsii_.Set(
		j,
		"customInstanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_CfnLayer)SetCustomJson(val interface{}) {
	_jsii_.Set(
		j,
		"customJson",
		val,
	)
}

func (j *jsiiProxy_CfnLayer)SetCustomRecipes(val interface{}) {
	if err := j.validateSetCustomRecipesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customRecipes",
		val,
	)
}

func (j *jsiiProxy_CfnLayer)SetCustomSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"customSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_CfnLayer)SetEnableAutoHealing(val interface{}) {
	if err := j.validateSetEnableAutoHealingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAutoHealing",
		val,
	)
}

func (j *jsiiProxy_CfnLayer)SetInstallUpdatesOnBoot(val interface{}) {
	if err := j.validateSetInstallUpdatesOnBootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"installUpdatesOnBoot",
		val,
	)
}

func (j *jsiiProxy_CfnLayer)SetLifecycleEventConfiguration(val interface{}) {
	if err := j.validateSetLifecycleEventConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleEventConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnLayer)SetLoadBasedAutoScaling(val interface{}) {
	if err := j.validateSetLoadBasedAutoScalingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBasedAutoScaling",
		val,
	)
}

func (j *jsiiProxy_CfnLayer)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnLayer)SetPackages(val *[]*string) {
	_jsii_.Set(
		j,
		"packages",
		val,
	)
}

func (j *jsiiProxy_CfnLayer)SetShortname(val *string) {
	if err := j.validateSetShortnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shortname",
		val,
	)
}

func (j *jsiiProxy_CfnLayer)SetStackId(val *string) {
	if err := j.validateSetStackIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stackId",
		val,
	)
}

func (j *jsiiProxy_CfnLayer)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnLayer)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_CfnLayer)SetUseEbsOptimizedInstances(val interface{}) {
	if err := j.validateSetUseEbsOptimizedInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useEbsOptimizedInstances",
		val,
	)
}

func (j *jsiiProxy_CfnLayer)SetVolumeConfigurations(val interface{}) {
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
func CfnLayer_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLayer_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworks.CfnLayer",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnLayer_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnLayer_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworks.CfnLayer",
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
func CfnLayer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLayer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworks.CfnLayer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLayer_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opsworks.CfnLayer",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLayer) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnLayer) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnLayer) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnLayer) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnLayer) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnLayer) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnLayer) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnLayer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnLayer) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnLayer) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnLayer) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnLayer) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLayer) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLayer) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLayer) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnLayer) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnLayer) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnLayer) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLayer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLayer) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

