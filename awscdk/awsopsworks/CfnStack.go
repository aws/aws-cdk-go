package awsopsworks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopsworks/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsopsworks"
	"github.com/aws/constructs-go/constructs/v10"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var customJson interface{}
//
//   cfnStack := awscdk.Aws_opsworks.NewCfnStack(this, jsii.String("MyCfnStack"), &CfnStackProps{
//   	DefaultInstanceProfileArn: jsii.String("defaultInstanceProfileArn"),
//   	Name: jsii.String("name"),
//   	ServiceRoleArn: jsii.String("serviceRoleArn"),
//
//   	// the properties below are optional
//   	AgentVersion: jsii.String("agentVersion"),
//   	Attributes: map[string]*string{
//   		"attributesKey": jsii.String("attributes"),
//   	},
//   	ChefConfiguration: &ChefConfigurationProperty{
//   		BerkshelfVersion: jsii.String("berkshelfVersion"),
//   		ManageBerkshelf: jsii.Boolean(false),
//   	},
//   	CloneAppIds: []*string{
//   		jsii.String("cloneAppIds"),
//   	},
//   	ClonePermissions: jsii.Boolean(false),
//   	ConfigurationManager: &StackConfigurationManagerProperty{
//   		Name: jsii.String("name"),
//   		Version: jsii.String("version"),
//   	},
//   	CustomCookbooksSource: &SourceProperty{
//   		Password: jsii.String("password"),
//   		Revision: jsii.String("revision"),
//   		SshKey: jsii.String("sshKey"),
//   		Type: jsii.String("type"),
//   		Url: jsii.String("url"),
//   		Username: jsii.String("username"),
//   	},
//   	CustomJson: customJson,
//   	DefaultAvailabilityZone: jsii.String("defaultAvailabilityZone"),
//   	DefaultOs: jsii.String("defaultOs"),
//   	DefaultRootDeviceType: jsii.String("defaultRootDeviceType"),
//   	DefaultSshKeyName: jsii.String("defaultSshKeyName"),
//   	DefaultSubnetId: jsii.String("defaultSubnetId"),
//   	EcsClusterArn: jsii.String("ecsClusterArn"),
//   	ElasticIps: []interface{}{
//   		&ElasticIpProperty{
//   			Ip: jsii.String("ip"),
//
//   			// the properties below are optional
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	HostnameTheme: jsii.String("hostnameTheme"),
//   	RdsDbInstances: []interface{}{
//   		&RdsDbInstanceProperty{
//   			DbPassword: jsii.String("dbPassword"),
//   			DbUser: jsii.String("dbUser"),
//   			RdsDbInstanceArn: jsii.String("rdsDbInstanceArn"),
//   		},
//   	},
//   	SourceStackId: jsii.String("sourceStackId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UseCustomCookbooks: jsii.Boolean(false),
//   	UseOpsworksSecurityGroups: jsii.Boolean(false),
//   	VpcId: jsii.String("vpcId"),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html
//
type CfnStack interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsopsworks.IStackRef
	awscdk.ITaggable
	// The default OpsWorks Stacks agent version.
	//
	// You have the following options:.
	AgentVersion() *string
	SetAgentVersion(val *string)
	// One or more user-defined key-value pairs to be added to the stack attributes.
	Attributes() interface{}
	SetAttributes(val interface{})
	AttrId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// A `ChefConfiguration` object that specifies whether to enable Berkshelf and the Berkshelf version on Chef 11.10 stacks. For more information, see [Create a New Stack](https://docs.aws.amazon.com/opsworks/latest/userguide/workingstacks-creating.html) .
	ChefConfiguration() interface{}
	SetChefConfiguration(val interface{})
	// If you're cloning an OpsWorks stack, a list of OpsWorks application stack IDs from the source stack to include in the cloned stack.
	CloneAppIds() *[]*string
	SetCloneAppIds(val *[]*string)
	// If you're cloning an OpsWorks stack, indicates whether to clone the source stack's permissions.
	ClonePermissions() interface{}
	SetClonePermissions(val interface{})
	// The configuration manager.
	ConfigurationManager() interface{}
	SetConfigurationManager(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Contains the information required to retrieve an app or cookbook from a repository.
	CustomCookbooksSource() interface{}
	SetCustomCookbooksSource(val interface{})
	// A string that contains user-defined, custom JSON.
	CustomJson() interface{}
	SetCustomJson(val interface{})
	// The stack's default Availability Zone, which must be in the specified region.
	DefaultAvailabilityZone() *string
	SetDefaultAvailabilityZone(val *string)
	// The Amazon Resource Name (ARN) of an IAM profile that is the default profile for all of the stack's EC2 instances.
	DefaultInstanceProfileArn() *string
	SetDefaultInstanceProfileArn(val *string)
	// The stack's default operating system, which is installed on every instance unless you specify a different operating system when you create the instance.
	DefaultOs() *string
	SetDefaultOs(val *string)
	// The default root device type.
	DefaultRootDeviceType() *string
	SetDefaultRootDeviceType(val *string)
	// A default Amazon EC2 key pair name.
	DefaultSshKeyName() *string
	SetDefaultSshKeyName(val *string)
	// The stack's default subnet ID.
	DefaultSubnetId() *string
	SetDefaultSubnetId(val *string)
	// The Amazon Resource Name (ARN) of the  ( Amazon ECS ) cluster to register with the OpsWorks stack.
	EcsClusterArn() *string
	SetEcsClusterArn(val *string)
	// A list of Elastic IP addresses to register with the OpsWorks stack.
	ElasticIps() interface{}
	SetElasticIps(val interface{})
	Env() *interfaces.ResourceEnvironment
	// The stack's host name theme, with spaces replaced by underscores.
	HostnameTheme() *string
	SetHostnameTheme(val *string)
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
	// The stack name.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// The Amazon Relational Database Service ( Amazon RDS ) database instance to register with the OpsWorks stack.
	RdsDbInstances() interface{}
	SetRdsDbInstances(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack's IAM role, which allows OpsWorks Stacks to work with AWS resources on your behalf.
	ServiceRoleArn() *string
	SetServiceRoleArn(val *string)
	// If you're cloning an OpsWorks stack, the stack ID of the source OpsWorks stack to clone.
	SourceStackId() *string
	SetSourceStackId(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// A reference to a Stack resource.
	StackRef() *interfacesawsopsworks.StackReference
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// A map that contains tag keys and tag values that are attached to a stack or layer.
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
	// Whether the stack uses custom cookbooks.
	UseCustomCookbooks() interface{}
	SetUseCustomCookbooks(val interface{})
	// Whether to associate the OpsWorks Stacks built-in security groups with the stack's layers.
	UseOpsworksSecurityGroups() interface{}
	SetUseOpsworksSecurityGroups(val interface{})
	// The ID of the VPC that the stack is to be launched into.
	VpcId() *string
	SetVpcId(val *string)
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

// The jsii proxy struct for CfnStack
type jsiiProxy_CfnStack struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsopsworksIStackRef
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnStack) AgentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) Attributes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) ChefConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"chefConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) CloneAppIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cloneAppIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) ClonePermissions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clonePermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) ConfigurationManager() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configurationManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) CustomCookbooksSource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customCookbooksSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) CustomJson() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) DefaultAvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) DefaultInstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultInstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) DefaultOs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultOs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) DefaultRootDeviceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRootDeviceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) DefaultSshKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSshKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) DefaultSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) EcsClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ecsClusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) ElasticIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) HostnameTheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameTheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) RdsDbInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rdsDbInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) ServiceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) SourceStackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceStackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) StackRef() *interfacesawsopsworks.StackReference {
	var returns *interfacesawsopsworks.StackReference
	_jsii_.Get(
		j,
		"stackRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) UseCustomCookbooks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCustomCookbooks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) UseOpsworksSecurityGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useOpsworksSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}


// Create a new `AWS::OpsWorks::Stack`.
func NewCfnStack(scope constructs.Construct, id *string, props *CfnStackProps) CfnStack {
	_init_.Initialize()

	if err := validateNewCfnStackParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnStack{}

	_jsii_.Create(
		"aws-cdk-lib.aws_opsworks.CfnStack",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::OpsWorks::Stack`.
func NewCfnStack_Override(c CfnStack, scope constructs.Construct, id *string, props *CfnStackProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_opsworks.CfnStack",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnStack)SetAgentVersion(val *string) {
	_jsii_.Set(
		j,
		"agentVersion",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetAttributes(val interface{}) {
	if err := j.validateSetAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributes",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetChefConfiguration(val interface{}) {
	if err := j.validateSetChefConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"chefConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetCloneAppIds(val *[]*string) {
	_jsii_.Set(
		j,
		"cloneAppIds",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetClonePermissions(val interface{}) {
	if err := j.validateSetClonePermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clonePermissions",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetConfigurationManager(val interface{}) {
	if err := j.validateSetConfigurationManagerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationManager",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetCustomCookbooksSource(val interface{}) {
	if err := j.validateSetCustomCookbooksSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customCookbooksSource",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetCustomJson(val interface{}) {
	_jsii_.Set(
		j,
		"customJson",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetDefaultAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"defaultAvailabilityZone",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetDefaultInstanceProfileArn(val *string) {
	if err := j.validateSetDefaultInstanceProfileArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultInstanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetDefaultOs(val *string) {
	_jsii_.Set(
		j,
		"defaultOs",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetDefaultRootDeviceType(val *string) {
	_jsii_.Set(
		j,
		"defaultRootDeviceType",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetDefaultSshKeyName(val *string) {
	_jsii_.Set(
		j,
		"defaultSshKeyName",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetDefaultSubnetId(val *string) {
	_jsii_.Set(
		j,
		"defaultSubnetId",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetEcsClusterArn(val *string) {
	_jsii_.Set(
		j,
		"ecsClusterArn",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetElasticIps(val interface{}) {
	if err := j.validateSetElasticIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticIps",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetHostnameTheme(val *string) {
	_jsii_.Set(
		j,
		"hostnameTheme",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetRdsDbInstances(val interface{}) {
	if err := j.validateSetRdsDbInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rdsDbInstances",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetServiceRoleArn(val *string) {
	if err := j.validateSetServiceRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetSourceStackId(val *string) {
	_jsii_.Set(
		j,
		"sourceStackId",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetUseCustomCookbooks(val interface{}) {
	if err := j.validateSetUseCustomCookbooksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useCustomCookbooks",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetUseOpsworksSecurityGroups(val interface{}) {
	if err := j.validateSetUseOpsworksSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useOpsworksSecurityGroups",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetVpcId(val *string) {
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

func CfnStack_ArnForStack(resource interfacesawsopsworks.IStackRef) *string {
	_init_.Initialize()

	if err := validateCfnStack_ArnForStackParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworks.CfnStack",
		"arnForStack",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// Creates a new IStackRef from a stackId.
func CfnStack_FromStackId(scope constructs.Construct, id *string, stackId *string) interfacesawsopsworks.IStackRef {
	_init_.Initialize()

	if err := validateCfnStack_FromStackIdParameters(scope, id, stackId); err != nil {
		panic(err)
	}
	var returns interfacesawsopsworks.IStackRef

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworks.CfnStack",
		"fromStackId",
		[]interface{}{scope, id, stackId},
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
func CfnStack_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnStack_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworks.CfnStack",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnStack_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnStack_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworks.CfnStack",
		"isCfnResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnStack.
func CfnStack_IsCfnStack(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnStack_IsCfnStackParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworks.CfnStack",
		"isCfnStack",
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
func CfnStack_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnStack_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworks.CfnStack",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStack_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opsworks.CfnStack",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStack) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnStack) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnStack) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnStack) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnStack) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnStack) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnStack) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnStack) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnStack) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnStack) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnStack) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnStack) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStack) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStack) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnStack) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnStack) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnStack) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnStack) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStack) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStack) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

func (c *jsiiProxy_CfnStack) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

