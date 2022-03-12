package awsservicediscovery

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicediscovery/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Instance that uses Route 53 Alias record type.
//
// Currently, the only resource types supported are Elastic Load
// Balancers.
//
// TODO: EXAMPLE
//
type AliasTargetInstance interface {
	InstanceBase
	DnsName() *string
	Env() *awscdk.ResourceEnvironment
	InstanceId() *string
	Node() constructs.Node
	PhysicalName() *string
	Service() IService
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
	UniqueInstanceId() *string
}

// The jsii proxy struct for AliasTargetInstance
type jsiiProxy_AliasTargetInstance struct {
	jsiiProxy_InstanceBase
}

func (j *jsiiProxy_AliasTargetInstance) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AliasTargetInstance) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AliasTargetInstance) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AliasTargetInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AliasTargetInstance) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AliasTargetInstance) Service() IService {
	var returns IService
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AliasTargetInstance) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewAliasTargetInstance(scope constructs.Construct, id *string, props *AliasTargetInstanceProps) AliasTargetInstance {
	_init_.Initialize()

	j := jsiiProxy_AliasTargetInstance{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.AliasTargetInstance",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewAliasTargetInstance_Override(a AliasTargetInstance, scope constructs.Construct, id *string, props *AliasTargetInstanceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.AliasTargetInstance",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func AliasTargetInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.AliasTargetInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func AliasTargetInstance_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.AliasTargetInstance",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (a *jsiiProxy_AliasTargetInstance) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (a *jsiiProxy_AliasTargetInstance) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (a *jsiiProxy_AliasTargetInstance) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (a *jsiiProxy_AliasTargetInstance) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_AliasTargetInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Generate a unique instance Id that is safe to pass to CloudMap.
func (a *jsiiProxy_AliasTargetInstance) UniqueInstanceId() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"uniqueInstanceId",
		nil, // no parameters
		&returns,
	)

	return returns
}

// TODO: EXAMPLE
//
type AliasTargetInstanceProps struct {
	// Custom attributes of the instance.
	CustomAttributes *map[string]*string `json:"customAttributes" yaml:"customAttributes"`
	// The id of the instance resource.
	InstanceId *string `json:"instanceId" yaml:"instanceId"`
	// DNS name of the target.
	DnsName *string `json:"dnsName" yaml:"dnsName"`
	// The Cloudmap service this resource is registered to.
	Service IService `json:"service" yaml:"service"`
}

// Used when the resource that's associated with the service instance is accessible using values other than an IP address or a domain name (CNAME), i.e. for non-ip-instances.
//
// TODO: EXAMPLE
//
type BaseInstanceProps struct {
	// Custom attributes of the instance.
	CustomAttributes *map[string]*string `json:"customAttributes" yaml:"customAttributes"`
	// The id of the instance resource.
	InstanceId *string `json:"instanceId" yaml:"instanceId"`
}

// TODO: EXAMPLE
//
type BaseNamespaceProps struct {
	// A name for the Namespace.
	Name *string `json:"name" yaml:"name"`
	// A description of the Namespace.
	Description *string `json:"description" yaml:"description"`
}

// Basic props needed to create a service in a given namespace.
//
// Used by HttpNamespace.createService
//
// TODO: EXAMPLE
//
type BaseServiceProps struct {
	// Structure containing failure threshold for a custom health checker.
	//
	// Only one of healthCheckConfig or healthCheckCustomConfig can be specified.
	// See: https://docs.aws.amazon.com/cloud-map/latest/api/API_HealthCheckCustomConfig.html
	CustomHealthCheck *HealthCheckCustomConfig `json:"customHealthCheck" yaml:"customHealthCheck"`
	// A description of the service.
	Description *string `json:"description" yaml:"description"`
	// Settings for an optional health check.
	//
	// If you specify health check settings, AWS Cloud Map associates the health
	// check with the records that you specify in DnsConfig. Only one of healthCheckConfig or healthCheckCustomConfig can
	// be specified. Not valid for PrivateDnsNamespaces. If you use healthCheck, you can only register IP instances to
	// this service.
	HealthCheck *HealthCheckConfig `json:"healthCheck" yaml:"healthCheck"`
	// A name for the Service.
	Name *string `json:"name" yaml:"name"`
}

// A CloudFormation `AWS::ServiceDiscovery::HttpNamespace`.
//
// The `HttpNamespace` resource is an AWS Cloud Map resource type that contains information about an HTTP namespace. Service instances that you register using an HTTP namespace can be discovered using a `DiscoverInstances` request but can't be discovered using DNS.
//
// For the current quota on the number of namespaces that you can create using the same AWS account, see [AWS Cloud Map quotas](https://docs.aws.amazon.com/cloud-map/latest/dg/cloud-map-limits.html) in the ** .
//
// TODO: EXAMPLE
//
type CfnHttpNamespace interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnHttpNamespace
type jsiiProxy_CfnHttpNamespace struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnHttpNamespace) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpNamespace) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpNamespace) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpNamespace) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpNamespace) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpNamespace) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpNamespace) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpNamespace) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpNamespace) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpNamespace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpNamespace) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpNamespace) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpNamespace) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpNamespace) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ServiceDiscovery::HttpNamespace`.
func NewCfnHttpNamespace(scope constructs.Construct, id *string, props *CfnHttpNamespaceProps) CfnHttpNamespace {
	_init_.Initialize()

	j := jsiiProxy_CfnHttpNamespace{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.CfnHttpNamespace",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ServiceDiscovery::HttpNamespace`.
func NewCfnHttpNamespace_Override(c CfnHttpNamespace, scope constructs.Construct, id *string, props *CfnHttpNamespaceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.CfnHttpNamespace",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnHttpNamespace) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnHttpNamespace) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnHttpNamespace_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.CfnHttpNamespace",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnHttpNamespace_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.CfnHttpNamespace",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnHttpNamespace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.CfnHttpNamespace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnHttpNamespace_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_servicediscovery.CfnHttpNamespace",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnHttpNamespace) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnHttpNamespace) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnHttpNamespace) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnHttpNamespace) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnHttpNamespace) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnHttpNamespace) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnHttpNamespace) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnHttpNamespace) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnHttpNamespace) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnHttpNamespace) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnHttpNamespace) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnHttpNamespace) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnHttpNamespace) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnHttpNamespace) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHttpNamespace) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnHttpNamespace`.
//
// TODO: EXAMPLE
//
type CfnHttpNamespaceProps struct {
	// The name that you want to assign to this namespace.
	Name *string `json:"name" yaml:"name"`
	// A description for the namespace.
	Description *string `json:"description" yaml:"description"`
	// The tags for the namespace.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::ServiceDiscovery::Instance`.
//
// A complex type that contains information about an instance that AWS Cloud Map creates when you submit a `RegisterInstance` request.
//
// TODO: EXAMPLE
//
type CfnInstance interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	InstanceAttributes() interface{}
	SetInstanceAttributes(val interface{})
	InstanceId() *string
	SetInstanceId(val *string)
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	ServiceId() *string
	SetServiceId(val *string)
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnInstance
type jsiiProxy_CfnInstance struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnInstance) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) InstanceAttributes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) ServiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ServiceDiscovery::Instance`.
func NewCfnInstance(scope constructs.Construct, id *string, props *CfnInstanceProps) CfnInstance {
	_init_.Initialize()

	j := jsiiProxy_CfnInstance{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.CfnInstance",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ServiceDiscovery::Instance`.
func NewCfnInstance_Override(c CfnInstance, scope constructs.Construct, id *string, props *CfnInstanceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.CfnInstance",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnInstance) SetInstanceAttributes(val interface{}) {
	_jsii_.Set(
		j,
		"instanceAttributes",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetInstanceId(val *string) {
	_jsii_.Set(
		j,
		"instanceId",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetServiceId(val *string) {
	_jsii_.Set(
		j,
		"serviceId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnInstance_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.CfnInstance",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnInstance_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.CfnInstance",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.CfnInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInstance_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_servicediscovery.CfnInstance",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnInstance) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnInstance) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnInstance) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnInstance) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnInstance) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnInstance) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnInstance) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnInstance) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnInstance) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnInstance) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnInstance) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnInstance) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnInstance) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstance) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnInstance`.
//
// TODO: EXAMPLE
//
type CfnInstanceProps struct {
	// A string map that contains the following information for the service that you specify in `ServiceId` :.
	//
	// - The attributes that apply to the records that are defined in the service.
	// - For each attribute, the applicable value.
	//
	// Supported attribute keys include the following:
	//
	// - **AWS_ALIAS_DNS_NAME** - If you want AWS Cloud Map to create a Route 53 alias record that routes traffic to an Elastic Load Balancing load balancer, specify the DNS name that is associated with the load balancer. For information about how to get the DNS name, see [AliasTarget->DNSName](https://docs.aws.amazon.com/Route53/latest/APIReference/API_AliasTarget.html#Route53-Type-AliasTarget-DNSName) in the *Route 53 API Reference* .
	//
	// Note the following:
	//
	// - The configuration for the service that is specified by `ServiceId` must include settings for an `A` record, an `AAAA` record, or both.
	// - In the service that is specified by `ServiceId` , the value of `RoutingPolicy` must be `WEIGHTED` .
	// - If the service that is specified by `ServiceId` includes `HealthCheckConfig` settings, AWS Cloud Map will create the health check, but it won't associate the health check with the alias record.
	// - Auto naming currently doesn't support creating alias records that route traffic to AWS resources other than ELB load balancers.
	// - If you specify a value for `AWS_ALIAS_DNS_NAME` , don't specify values for any of the `AWS_INSTANCE` attributes.
	// - **AWS_EC2_INSTANCE_ID** - *HTTP namespaces only.* The Amazon EC2 instance ID for the instance. The `AWS_INSTANCE_IPV4` attribute contains the primary private IPv4 address. When creating resources with a type of [AWS::ServiceDiscovery::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-instance.html) , if the `AWS_EC2_INSTANCE_ID` attribute is specified, the only other attribute that can be specified is `AWS_INIT_HEALTH_STATUS` . After the resource has been created, the `AWS_INSTANCE_IPV4` attribute contains the primary private IPv4 address.
	// - **AWS_INIT_HEALTH_STATUS** - If the service configuration includes `HealthCheckCustomConfig` , when creating resources with a type of [AWS::ServiceDiscovery::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-instance.html) you can optionally use `AWS_INIT_HEALTH_STATUS` to specify the initial status of the custom health check, `HEALTHY` or `UNHEALTHY` . If you don't specify a value for `AWS_INIT_HEALTH_STATUS` , the initial status is `HEALTHY` . This attribute can only be used when creating resources and will not be seen on existing resources.
	// - **AWS_INSTANCE_CNAME** - If the service configuration includes a `CNAME` record, the domain name that you want Route 53 to return in response to DNS queries, for example, `example.com` .
	//
	// This value is required if the service specified by `ServiceId` includes settings for an `CNAME` record.
	// - **AWS_INSTANCE_IPV4** - If the service configuration includes an `A` record, the IPv4 address that you want Route 53 to return in response to DNS queries, for example, `192.0.2.44` .
	//
	// This value is required if the service specified by `ServiceId` includes settings for an `A` record. If the service includes settings for an `SRV` record, you must specify a value for `AWS_INSTANCE_IPV4` , `AWS_INSTANCE_IPV6` , or both.
	// - **AWS_INSTANCE_IPV6** - If the service configuration includes an `AAAA` record, the IPv6 address that you want Route 53 to return in response to DNS queries, for example, `2001:0db8:85a3:0000:0000:abcd:0001:2345` .
	//
	// This value is required if the service specified by `ServiceId` includes settings for an `AAAA` record. If the service includes settings for an `SRV` record, you must specify a value for `AWS_INSTANCE_IPV4` , `AWS_INSTANCE_IPV6` , or both.
	// - **AWS_INSTANCE_PORT** - If the service includes an `SRV` record, the value that you want Route 53 to return for the port.
	//
	// If the service includes `HealthCheckConfig` , the port on the endpoint that you want Route 53 to send requests to.
	//
	// This value is required if you specified settings for an `SRV` record or a Route 53 health check when you created the service.
	InstanceAttributes interface{} `json:"instanceAttributes" yaml:"instanceAttributes"`
	// The ID of the service that you want to use for settings for the instance.
	ServiceId *string `json:"serviceId" yaml:"serviceId"`
	// An identifier that you want to associate with the instance. Note the following:.
	//
	// - If the service that's specified by `ServiceId` includes settings for an `SRV` record, the value of `InstanceId` is automatically included as part of the value for the `SRV` record. For more information, see [DnsRecord > Type](https://docs.aws.amazon.com/cloud-map/latest/api/API_DnsRecord.html#cloudmap-Type-DnsRecord-Type) .
	// - You can use this value to update an existing instance.
	// - To register a new instance, you must specify a value that's unique among instances that you register by using the same service.
	// - If you specify an existing `InstanceId` and `ServiceId` , AWS Cloud Map updates the existing DNS records, if any. If there's also an existing health check, AWS Cloud Map deletes the old health check and creates a new one.
	//
	// > The health check isn't deleted immediately, so it will still appear for a while if you submit a `ListHealthChecks` request, for example.
	InstanceId *string `json:"instanceId" yaml:"instanceId"`
}

// A CloudFormation `AWS::ServiceDiscovery::PrivateDnsNamespace`.
//
// Creates a private namespace based on DNS, which is visible only inside a specified Amazon VPC. The namespace defines your service naming scheme. For example, if you name your namespace `example.com` and name your service `backend` , the resulting DNS name for the service is `backend.example.com` . Service instances that are registered using a private DNS namespace can be discovered using either a `DiscoverInstances` request or using DNS. For the current quota on the number of namespaces that you can create using the same AWS account , see [AWS Cloud Map quotas](https://docs.aws.amazon.com/cloud-map/latest/dg/cloud-map-limits.html) in the *AWS Cloud Map Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnPrivateDnsNamespace interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrHostedZoneId() *string
	AttrId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Properties() interface{}
	SetProperties(val interface{})
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	Vpc() *string
	SetVpc(val *string)
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnPrivateDnsNamespace
type jsiiProxy_CfnPrivateDnsNamespace struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPrivateDnsNamespace) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrivateDnsNamespace) AttrHostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrHostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrivateDnsNamespace) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrivateDnsNamespace) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrivateDnsNamespace) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrivateDnsNamespace) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrivateDnsNamespace) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrivateDnsNamespace) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrivateDnsNamespace) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrivateDnsNamespace) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrivateDnsNamespace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrivateDnsNamespace) Properties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrivateDnsNamespace) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrivateDnsNamespace) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrivateDnsNamespace) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrivateDnsNamespace) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrivateDnsNamespace) Vpc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


// Create a new `AWS::ServiceDiscovery::PrivateDnsNamespace`.
func NewCfnPrivateDnsNamespace(scope constructs.Construct, id *string, props *CfnPrivateDnsNamespaceProps) CfnPrivateDnsNamespace {
	_init_.Initialize()

	j := jsiiProxy_CfnPrivateDnsNamespace{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.CfnPrivateDnsNamespace",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ServiceDiscovery::PrivateDnsNamespace`.
func NewCfnPrivateDnsNamespace_Override(c CfnPrivateDnsNamespace, scope constructs.Construct, id *string, props *CfnPrivateDnsNamespaceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.CfnPrivateDnsNamespace",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPrivateDnsNamespace) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnPrivateDnsNamespace) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnPrivateDnsNamespace) SetProperties(val interface{}) {
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_CfnPrivateDnsNamespace) SetVpc(val *string) {
	_jsii_.Set(
		j,
		"vpc",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnPrivateDnsNamespace_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.CfnPrivateDnsNamespace",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnPrivateDnsNamespace_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.CfnPrivateDnsNamespace",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnPrivateDnsNamespace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.CfnPrivateDnsNamespace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPrivateDnsNamespace_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_servicediscovery.CfnPrivateDnsNamespace",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnPrivateDnsNamespace) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnPrivateDnsNamespace) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnPrivateDnsNamespace) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnPrivateDnsNamespace) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnPrivateDnsNamespace) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnPrivateDnsNamespace) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnPrivateDnsNamespace) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnPrivateDnsNamespace) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnPrivateDnsNamespace) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnPrivateDnsNamespace) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnPrivateDnsNamespace) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPrivateDnsNamespace) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnPrivateDnsNamespace) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnPrivateDnsNamespace) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPrivateDnsNamespace) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// DNS properties for the private DNS namespace.
//
// TODO: EXAMPLE
//
type CfnPrivateDnsNamespace_PrivateDnsPropertiesMutableProperty struct {
	// Fields for the Start of Authority (SOA) record for the hosted zone for the private DNS namespace.
	Soa interface{} `json:"soa" yaml:"soa"`
}

// Properties for the private DNS namespace.
//
// TODO: EXAMPLE
//
type CfnPrivateDnsNamespace_PropertiesProperty struct {
	// DNS properties for the private DNS namespace.
	DnsProperties interface{} `json:"dnsProperties" yaml:"dnsProperties"`
}

// Start of Authority (SOA) properties for a public or private DNS namespace.
//
// TODO: EXAMPLE
//
type CfnPrivateDnsNamespace_SOAProperty struct {
	// The time to live (TTL) for purposes of negative caching.
	Ttl *float64 `json:"ttl" yaml:"ttl"`
}

// Properties for defining a `CfnPrivateDnsNamespace`.
//
// TODO: EXAMPLE
//
type CfnPrivateDnsNamespaceProps struct {
	// The name that you want to assign to this namespace.
	//
	// When you create a private DNS namespace, AWS Cloud Map automatically creates an Amazon Route 53 private hosted zone that has the same name as the namespace.
	Name *string `json:"name" yaml:"name"`
	// The ID of the Amazon VPC that you want to associate the namespace with.
	Vpc *string `json:"vpc" yaml:"vpc"`
	// A description for the namespace.
	Description *string `json:"description" yaml:"description"`
	// Properties for the private DNS namespace.
	Properties interface{} `json:"properties" yaml:"properties"`
	// The tags for the namespace.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::ServiceDiscovery::PublicDnsNamespace`.
//
// Creates a public namespace based on DNS, which is visible on the internet. The namespace defines your service naming scheme. For example, if you name your namespace `example.com` and name your service `backend` , the resulting DNS name for the service is `backend.example.com` . You can discover instances that were registered with a public DNS namespace by using either a `DiscoverInstances` request or using DNS. For the current quota on the number of namespaces that you can create using the same AWS account , see [AWS Cloud Map quotas](https://docs.aws.amazon.com/cloud-map/latest/dg/cloud-map-limits.html) in the *AWS Cloud Map Developer Guide* .
//
// > The `CreatePublicDnsNamespace` API operation is not supported in the AWS GovCloud (US) Regions.
//
// TODO: EXAMPLE
//
type CfnPublicDnsNamespace interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrHostedZoneId() *string
	AttrId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Properties() interface{}
	SetProperties(val interface{})
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnPublicDnsNamespace
type jsiiProxy_CfnPublicDnsNamespace struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPublicDnsNamespace) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicDnsNamespace) AttrHostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrHostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicDnsNamespace) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicDnsNamespace) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicDnsNamespace) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicDnsNamespace) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicDnsNamespace) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicDnsNamespace) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicDnsNamespace) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicDnsNamespace) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicDnsNamespace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicDnsNamespace) Properties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicDnsNamespace) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicDnsNamespace) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicDnsNamespace) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicDnsNamespace) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ServiceDiscovery::PublicDnsNamespace`.
func NewCfnPublicDnsNamespace(scope constructs.Construct, id *string, props *CfnPublicDnsNamespaceProps) CfnPublicDnsNamespace {
	_init_.Initialize()

	j := jsiiProxy_CfnPublicDnsNamespace{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.CfnPublicDnsNamespace",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ServiceDiscovery::PublicDnsNamespace`.
func NewCfnPublicDnsNamespace_Override(c CfnPublicDnsNamespace, scope constructs.Construct, id *string, props *CfnPublicDnsNamespaceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.CfnPublicDnsNamespace",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPublicDnsNamespace) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnPublicDnsNamespace) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnPublicDnsNamespace) SetProperties(val interface{}) {
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnPublicDnsNamespace_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.CfnPublicDnsNamespace",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnPublicDnsNamespace_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.CfnPublicDnsNamespace",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnPublicDnsNamespace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.CfnPublicDnsNamespace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPublicDnsNamespace_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_servicediscovery.CfnPublicDnsNamespace",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnPublicDnsNamespace) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnPublicDnsNamespace) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnPublicDnsNamespace) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnPublicDnsNamespace) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnPublicDnsNamespace) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnPublicDnsNamespace) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnPublicDnsNamespace) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnPublicDnsNamespace) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnPublicDnsNamespace) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnPublicDnsNamespace) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnPublicDnsNamespace) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPublicDnsNamespace) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnPublicDnsNamespace) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnPublicDnsNamespace) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPublicDnsNamespace) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for the public DNS namespace.
//
// TODO: EXAMPLE
//
type CfnPublicDnsNamespace_PropertiesProperty struct {
	// DNS properties for the public DNS namespace.
	DnsProperties interface{} `json:"dnsProperties" yaml:"dnsProperties"`
}

// DNS properties for the public DNS namespace.
//
// TODO: EXAMPLE
//
type CfnPublicDnsNamespace_PublicDnsPropertiesMutableProperty struct {
	// Start of Authority (SOA) record for the hosted zone for the public DNS namespace.
	Soa interface{} `json:"soa" yaml:"soa"`
}

// Start of Authority (SOA) properties for a public or private DNS namespace.
//
// TODO: EXAMPLE
//
type CfnPublicDnsNamespace_SOAProperty struct {
	// The time to live (TTL) for purposes of negative caching.
	Ttl *float64 `json:"ttl" yaml:"ttl"`
}

// Properties for defining a `CfnPublicDnsNamespace`.
//
// TODO: EXAMPLE
//
type CfnPublicDnsNamespaceProps struct {
	// The name that you want to assign to this namespace.
	Name *string `json:"name" yaml:"name"`
	// A description for the namespace.
	Description *string `json:"description" yaml:"description"`
	// Properties for the public DNS namespace.
	Properties interface{} `json:"properties" yaml:"properties"`
	// The tags for the namespace.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::ServiceDiscovery::Service`.
//
// A complex type that contains information about a service, which defines the configuration of the following entities:
//
// - For public and private DNS namespaces, one of the following combinations of DNS records in Amazon Route 53:
//
// - A
// - AAAA
// - A and AAAA
// - SRV
// - CNAME
// - Optionally, a health check
//
// TODO: EXAMPLE
//
type CfnService interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrId() *string
	AttrName() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	DnsConfig() interface{}
	SetDnsConfig(val interface{})
	HealthCheckConfig() interface{}
	SetHealthCheckConfig(val interface{})
	HealthCheckCustomConfig() interface{}
	SetHealthCheckCustomConfig(val interface{})
	LogicalId() *string
	Name() *string
	SetName(val *string)
	NamespaceId() *string
	SetNamespaceId(val *string)
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	Type() *string
	SetType(val *string)
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnService
type jsiiProxy_CfnService struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnService) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_CfnService) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) DnsConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dnsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) HealthCheckConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthCheckConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) HealthCheckCustomConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthCheckCustomConfig",
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

func (j *jsiiProxy_CfnService) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) NamespaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceId",
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

func (j *jsiiProxy_CfnService) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
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

func (j *jsiiProxy_CfnService) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
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


// Create a new `AWS::ServiceDiscovery::Service`.
func NewCfnService(scope constructs.Construct, id *string, props *CfnServiceProps) CfnService {
	_init_.Initialize()

	j := jsiiProxy_CfnService{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.CfnService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ServiceDiscovery::Service`.
func NewCfnService_Override(c CfnService, scope constructs.Construct, id *string, props *CfnServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.CfnService",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnService) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetDnsConfig(val interface{}) {
	_jsii_.Set(
		j,
		"dnsConfig",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetHealthCheckConfig(val interface{}) {
	_jsii_.Set(
		j,
		"healthCheckConfig",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetHealthCheckCustomConfig(val interface{}) {
	_jsii_.Set(
		j,
		"healthCheckCustomConfig",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetNamespaceId(val *string) {
	_jsii_.Set(
		j,
		"namespaceId",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
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

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.CfnService",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnService_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.CfnService",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.CfnService",
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
		"aws-cdk-lib.aws_servicediscovery.CfnService",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnService) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnService) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnService) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnService) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnService) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnService) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnService) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnService) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnService) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnService) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnService) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnService) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
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

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
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
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A complex type that contains information about the Amazon Route 53 DNS records that you want AWS Cloud Map to create when you register an instance.
//
// TODO: EXAMPLE
//
type CfnService_DnsConfigProperty struct {
	// An array that contains one `DnsRecord` object for each Route 53 DNS record that you want AWS Cloud Map to create when you register an instance.
	DnsRecords interface{} `json:"dnsRecords" yaml:"dnsRecords"`
	// The ID of the namespace to use for DNS configuration.
	//
	// > You must specify a value for `NamespaceId` either for `DnsConfig` or for the [service properties](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-service.html) . Don't specify a value in both places.
	NamespaceId *string `json:"namespaceId" yaml:"namespaceId"`
	// The routing policy that you want to apply to all Route 53 DNS records that AWS Cloud Map creates when you register an instance and specify this service.
	//
	// > If you want to use this service to register instances that create alias records, specify `WEIGHTED` for the routing policy.
	//
	// You can specify the following values:
	//
	// - **MULTIVALUE** - If you define a health check for the service and the health check is healthy, Route 53 returns the applicable value for up to eight instances.
	//
	// For example, suppose that the service includes configurations for one `A` record and a health check. You use the service to register 10 instances. Route 53 responds to DNS queries with IP addresses for up to eight healthy instances. If fewer than eight instances are healthy, Route 53 responds to every DNS query with the IP addresses for all of the healthy instances.
	//
	// If you don't define a health check for the service, Route 53 assumes that all instances are healthy and returns the values for up to eight instances.
	//
	// For more information about the multivalue routing policy, see [Multivalue Answer Routing](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy.html#routing-policy-multivalue) in the *Route 53 Developer Guide* .
	// - **WEIGHTED** - Route 53 returns the applicable value from one randomly selected instance from among the instances that you registered using the same service. Currently, all records have the same weight, so you can't route more or less traffic to any instances.
	//
	// For example, suppose that the service includes configurations for one `A` record and a health check. You use the service to register 10 instances. Route 53 responds to DNS queries with the IP address for one randomly selected instance from among the healthy instances. If no instances are healthy, Route 53 responds to DNS queries as if all of the instances were healthy.
	//
	// If you don't define a health check for the service, Route 53 assumes that all instances are healthy and returns the applicable value for one randomly selected instance.
	//
	// For more information about the weighted routing policy, see [Weighted Routing](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy.html#routing-policy-weighted) in the *Route 53 Developer Guide* .
	RoutingPolicy *string `json:"routingPolicy" yaml:"routingPolicy"`
}

// A complex type that contains information about the Route 53 DNS records that you want AWS Cloud Map to create when you register an instance.
//
// TODO: EXAMPLE
//
type CfnService_DnsRecordProperty struct {
	// The amount of time, in seconds, that you want DNS resolvers to cache the settings for this record.
	//
	// > Alias records don't include a TTL because Route 53 uses the TTL for the AWS resource that an alias record routes traffic to. If you include the `AWS_ALIAS_DNS_NAME` attribute when you submit a [RegisterInstance](https://docs.aws.amazon.com/cloud-map/latest/api/API_RegisterInstance.html) request, the `TTL` value is ignored. Always specify a TTL for the service; you can use a service to register instances that create either alias or non-alias records.
	Ttl *float64 `json:"ttl" yaml:"ttl"`
	// The type of the resource, which indicates the type of value that Route 53 returns in response to DNS queries.
	//
	// You can specify values for `Type` in the following combinations:
	//
	// - `A`
	// - `AAAA`
	// - `A` and `AAAA`
	// - `SRV`
	// - `CNAME`
	//
	// If you want AWS Cloud Map to create a Route 53 alias record when you register an instance, specify `A` or `AAAA` for `Type` .
	//
	// You specify other settings, such as the IP address for `A` and `AAAA` records, when you register an instance. For more information, see [RegisterInstance](https://docs.aws.amazon.com/cloud-map/latest/api/API_RegisterInstance.html) .
	//
	// The following values are supported:
	//
	// - **A** - Route 53 returns the IP address of the resource in IPv4 format, such as 192.0.2.44.
	// - **AAAA** - Route 53 returns the IP address of the resource in IPv6 format, such as 2001:0db8:85a3:0000:0000:abcd:0001:2345.
	// - **CNAME** - Route 53 returns the domain name of the resource, such as www.example.com. Note the following:
	//
	// - You specify the domain name that you want to route traffic to when you register an instance. For more information, see [Attributes](https://docs.aws.amazon.com/cloud-map/latest/api/API_RegisterInstance.html#cloudmap-RegisterInstance-request-Attributes) in the topic [RegisterInstance](https://docs.aws.amazon.com/cloud-map/latest/api/API_RegisterInstance.html) .
	// - You must specify `WEIGHTED` for the value of `RoutingPolicy` .
	// - You can't specify both `CNAME` for `Type` and settings for `HealthCheckConfig` . If you do, the request will fail with an `InvalidInput` error.
	// - **SRV** - Route 53 returns the value for an `SRV` record. The value for an `SRV` record uses the following values:
	//
	// `priority weight port service-hostname`
	//
	// Note the following about the values:
	//
	// - The values of `priority` and `weight` are both set to `1` and can't be changed.
	// - The value of `port` comes from the value that you specify for the `AWS_INSTANCE_PORT` attribute when you submit a [RegisterInstance](https://docs.aws.amazon.com/cloud-map/latest/api/API_RegisterInstance.html) request.
	// - The value of `service-hostname` is a concatenation of the following values:
	//
	// - The value that you specify for `InstanceId` when you register an instance.
	// - The name of the service.
	// - The name of the namespace.
	//
	// For example, if the value of `InstanceId` is `test` , the name of the service is `backend` , and the name of the namespace is `example.com` , the value of `service-hostname` is:
	//
	// `test.backend.example.com`
	//
	// If you specify settings for an `SRV` record and if you specify values for `AWS_INSTANCE_IPV4` , `AWS_INSTANCE_IPV6` , or both in the `RegisterInstance` request, AWS Cloud Map automatically creates `A` and/or `AAAA` records that have the same name as the value of `service-hostname` in the `SRV` record. You can ignore these records.
	Type *string `json:"type" yaml:"type"`
}

// *Public DNS and HTTP namespaces only.* A complex type that contains settings for an optional health check. If you specify settings for a health check, AWS Cloud Map associates the health check with the records that you specify in `DnsConfig` .
//
// > If you specify a health check configuration, you can specify either `HealthCheckCustomConfig` or `HealthCheckConfig` but not both.
//
// Health checks are basic Route 53 health checks that monitor an AWS endpoint. For information about pricing for health checks, see [Amazon Route 53 Pricing](https://docs.aws.amazon.com/route53/pricing/) .
//
// Note the following about configuring health checks.
//
// - **A and AAAA records** - If `DnsConfig` includes configurations for both `A` and `AAAA` records, AWS Cloud Map creates a health check that uses the IPv4 address to check the health of the resource. If the endpoint tthat's specified by the IPv4 address is unhealthy, Route 53 considers both the `A` and `AAAA` records to be unhealthy.
// - **CNAME records** - You can't specify settings for `HealthCheckConfig` when the `DNSConfig` includes `CNAME` for the value of `Type` . If you do, the `CreateService` request will fail with an `InvalidInput` error.
// - **Request interval** - A Route 53 health checker in each health-checking AWS Region sends a health check request to an endpoint every 30 seconds. On average, your endpoint receives a health check request about every two seconds. However, health checkers don't coordinate with one another. Therefore, you might sometimes see several requests in one second that's followed by a few seconds with no health checks at all.
// - **Health checking regions** - Health checkers perform checks from all Route 53 health-checking Regions. For a list of the current Regions, see [Regions](https://docs.aws.amazon.com/Route53/latest/APIReference/API_HealthCheckConfig.html#Route53-Type-HealthCheckConfig-Regions) .
// - **Alias records** - When you register an instance, if you include the `AWS_ALIAS_DNS_NAME` attribute, AWS Cloud Map creates a Route 53 alias record. Note the following:
//
// - Route 53 automatically sets `EvaluateTargetHealth` to true for alias records. When `EvaluateTargetHealth` is true, the alias record inherits the health of the referenced AWS resource. such as an ELB load balancer. For more information, see [EvaluateTargetHealth](https://docs.aws.amazon.com/Route53/latest/APIReference/API_AliasTarget.html#Route53-Type-AliasTarget-EvaluateTargetHealth) .
// - If you include `HealthCheckConfig` and then use the service to register an instance that creates an alias record, Route 53 doesn't create the health check.
// - **Charges for health checks** - Health checks are basic Route 53 health checks that monitor an AWS endpoint. For information about pricing for health checks, see [Amazon Route 53 Pricing](https://docs.aws.amazon.com/route53/pricing/) .
//
// TODO: EXAMPLE
//
type CfnService_HealthCheckConfigProperty struct {
	// The type of health check that you want to create, which indicates how Route 53 determines whether an endpoint is healthy.
	//
	// > You can't change the value of `Type` after you create a health check.
	//
	// You can create the following types of health checks:
	//
	// - *HTTP* : Route 53 tries to establish a TCP connection. If successful, Route 53 submits an HTTP request and waits for an HTTP status code of 200 or greater and less than 400.
	// - *HTTPS* : Route 53 tries to establish a TCP connection. If successful, Route 53 submits an HTTPS request and waits for an HTTP status code of 200 or greater and less than 400.
	//
	// > If you specify HTTPS for the value of `Type` , the endpoint must support TLS v1.0 or later.
	// - *TCP* : Route 53 tries to establish a TCP connection.
	//
	// If you specify `TCP` for `Type` , don't specify a value for `ResourcePath` .
	//
	// For more information, see [How Route 53 Determines Whether an Endpoint Is Healthy](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-determining-health-of-endpoints.html) in the *Route 53 Developer Guide* .
	Type *string `json:"type" yaml:"type"`
	// The number of consecutive health checks that an endpoint must pass or fail for Route 53 to change the current status of the endpoint from unhealthy to healthy or the other way around.
	//
	// For more information, see [How Route 53 Determines Whether an Endpoint Is Healthy](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-determining-health-of-endpoints.html) in the *Route 53 Developer Guide* .
	FailureThreshold *float64 `json:"failureThreshold" yaml:"failureThreshold"`
	// The path that you want Route 53 to request when performing health checks.
	//
	// The path can be any value that your endpoint returns an HTTP status code of a 2xx or 3xx format for when the endpoint is healthy. An example file is `/docs/route53-health-check.html` . Route 53 automatically adds the DNS name for the service. If you don't specify a value for `ResourcePath` , the default value is `/` .
	//
	// If you specify `TCP` for `Type` , you must *not* specify a value for `ResourcePath` .
	ResourcePath *string `json:"resourcePath" yaml:"resourcePath"`
}

// A complex type that contains information about an optional custom health check.
//
// A custom health check, which requires that you use a third-party health checker to evaluate the health of your resources, is useful in the following circumstances:
//
// - You can't use a health check that's defined by `HealthCheckConfig` because the resource isn't available over the internet. For example, you can use a custom health check when the instance is in an Amazon VPC. (To check the health of resources in a VPC, the health checker must also be in the VPC.)
// - You want to use a third-party health checker regardless of where your resources are located.
//
// > If you specify a health check configuration, you can specify either `HealthCheckCustomConfig` or `HealthCheckConfig` but not both.
//
// To change the status of a custom health check, submit an `UpdateInstanceCustomHealthStatus` request. AWS Cloud Map doesn't monitor the status of the resource, it just keeps a record of the status specified in the most recent `UpdateInstanceCustomHealthStatus` request.
//
// Here's how custom health checks work:
//
// - You create a service.
// - You register an instance.
// - You configure a third-party health checker to monitor the resource that's associated with the new instance.
//
// > AWS Cloud Map doesn't check the health of the resource directly.
// - The third-party health-checker determines that the resource is unhealthy and notifies your application.
// - Your application submits an `UpdateInstanceCustomHealthStatus` request.
// - AWS Cloud Map waits for 30 seconds.
// - If another `UpdateInstanceCustomHealthStatus` request doesn't arrive during that time to change the status back to healthy, AWS Cloud Map stops routing traffic to the resource.
//
// TODO: EXAMPLE
//
type CfnService_HealthCheckCustomConfigProperty struct {
	// > This parameter is no longer supported and is always set to 1.
	//
	// AWS Cloud Map waits for approximately 30 seconds after receiving an `UpdateInstanceCustomHealthStatus` request before changing the status of the service instance.
	//
	// The number of 30-second intervals that you want AWS Cloud Map to wait after receiving an `UpdateInstanceCustomHealthStatus` request before it changes the health status of a service instance.
	//
	// Sending a second or subsequent `UpdateInstanceCustomHealthStatus` request with the same value before 30 seconds has passed doesn't accelerate the change. AWS Cloud Map still waits `30` seconds after the first request to make the change.
	FailureThreshold *float64 `json:"failureThreshold" yaml:"failureThreshold"`
}

// Properties for defining a `CfnService`.
//
// TODO: EXAMPLE
//
type CfnServiceProps struct {
	// The description of the service.
	Description *string `json:"description" yaml:"description"`
	// A complex type that contains information about the Route 53 DNS records that you want AWS Cloud Map to create when you register an instance.
	DnsConfig interface{} `json:"dnsConfig" yaml:"dnsConfig"`
	// *Public DNS and HTTP namespaces only.* A complex type that contains settings for an optional health check. If you specify settings for a health check, AWS Cloud Map associates the health check with the records that you specify in `DnsConfig` .
	//
	// For information about the charges for health checks, see [Amazon Route 53 Pricing](https://docs.aws.amazon.com/route53/pricing/) .
	HealthCheckConfig interface{} `json:"healthCheckConfig" yaml:"healthCheckConfig"`
	// A complex type that contains information about an optional custom health check.
	//
	// > If you specify a health check configuration, you can specify either `HealthCheckCustomConfig` or `HealthCheckConfig` but not both.
	HealthCheckCustomConfig interface{} `json:"healthCheckCustomConfig" yaml:"healthCheckCustomConfig"`
	// The name of the service.
	Name *string `json:"name" yaml:"name"`
	// The ID of the namespace that was used to create the service.
	//
	// > You must specify a value for `NamespaceId` either for the service properties or for [DnsConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicediscovery-service-dnsconfig.html) . Don't specify a value in both places.
	NamespaceId *string `json:"namespaceId" yaml:"namespaceId"`
	// The tags for the service.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// If present, specifies that the service instances are only discoverable using the `DiscoverInstances` API operation.
	//
	// No DNS records is registered for the service instances. The only valid value is `HTTP` .
	Type *string `json:"type" yaml:"type"`
}

// Instance that is accessible using a domain name (CNAME).
//
// TODO: EXAMPLE
//
type CnameInstance interface {
	InstanceBase
	Cname() *string
	Env() *awscdk.ResourceEnvironment
	InstanceId() *string
	Node() constructs.Node
	PhysicalName() *string
	Service() IService
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
	UniqueInstanceId() *string
}

// The jsii proxy struct for CnameInstance
type jsiiProxy_CnameInstance struct {
	jsiiProxy_InstanceBase
}

func (j *jsiiProxy_CnameInstance) Cname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CnameInstance) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CnameInstance) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CnameInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CnameInstance) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CnameInstance) Service() IService {
	var returns IService
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CnameInstance) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewCnameInstance(scope constructs.Construct, id *string, props *CnameInstanceProps) CnameInstance {
	_init_.Initialize()

	j := jsiiProxy_CnameInstance{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.CnameInstance",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCnameInstance_Override(c CnameInstance, scope constructs.Construct, id *string, props *CnameInstanceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.CnameInstance",
		[]interface{}{scope, id, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CnameInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.CnameInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func CnameInstance_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.CnameInstance",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CnameInstance) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (c *jsiiProxy_CnameInstance) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (c *jsiiProxy_CnameInstance) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (c *jsiiProxy_CnameInstance) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_CnameInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Generate a unique instance Id that is safe to pass to CloudMap.
func (c *jsiiProxy_CnameInstance) UniqueInstanceId() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"uniqueInstanceId",
		nil, // no parameters
		&returns,
	)

	return returns
}

// TODO: EXAMPLE
//
type CnameInstanceBaseProps struct {
	// Custom attributes of the instance.
	CustomAttributes *map[string]*string `json:"customAttributes" yaml:"customAttributes"`
	// The id of the instance resource.
	InstanceId *string `json:"instanceId" yaml:"instanceId"`
	// If the service configuration includes a CNAME record, the domain name that you want Route 53 to return in response to DNS queries, for example, example.com. This value is required if the service specified by ServiceId includes settings for an CNAME record.
	InstanceCname *string `json:"instanceCname" yaml:"instanceCname"`
}

// TODO: EXAMPLE
//
type CnameInstanceProps struct {
	// Custom attributes of the instance.
	CustomAttributes *map[string]*string `json:"customAttributes" yaml:"customAttributes"`
	// The id of the instance resource.
	InstanceId *string `json:"instanceId" yaml:"instanceId"`
	// If the service configuration includes a CNAME record, the domain name that you want Route 53 to return in response to DNS queries, for example, example.com. This value is required if the service specified by ServiceId includes settings for an CNAME record.
	InstanceCname *string `json:"instanceCname" yaml:"instanceCname"`
	// The Cloudmap service this resource is registered to.
	Service IService `json:"service" yaml:"service"`
}

// TODO: EXAMPLE
//
type DnsRecordType string

const (
	DnsRecordType_A DnsRecordType = "A"
	DnsRecordType_AAAA DnsRecordType = "AAAA"
	DnsRecordType_A_AAAA DnsRecordType = "A_AAAA"
	DnsRecordType_SRV DnsRecordType = "SRV"
	DnsRecordType_CNAME DnsRecordType = "CNAME"
)

// Service props needed to create a service in a given namespace.
//
// Used by createService() for PrivateDnsNamespace and
// PublicDnsNamespace
//
// TODO: EXAMPLE
//
type DnsServiceProps struct {
	// Structure containing failure threshold for a custom health checker.
	//
	// Only one of healthCheckConfig or healthCheckCustomConfig can be specified.
	// See: https://docs.aws.amazon.com/cloud-map/latest/api/API_HealthCheckCustomConfig.html
	CustomHealthCheck *HealthCheckCustomConfig `json:"customHealthCheck" yaml:"customHealthCheck"`
	// A description of the service.
	Description *string `json:"description" yaml:"description"`
	// Settings for an optional health check.
	//
	// If you specify health check settings, AWS Cloud Map associates the health
	// check with the records that you specify in DnsConfig. Only one of healthCheckConfig or healthCheckCustomConfig can
	// be specified. Not valid for PrivateDnsNamespaces. If you use healthCheck, you can only register IP instances to
	// this service.
	HealthCheck *HealthCheckConfig `json:"healthCheck" yaml:"healthCheck"`
	// A name for the Service.
	Name *string `json:"name" yaml:"name"`
	// The DNS type of the record that you want AWS Cloud Map to create.
	//
	// Supported record types
	// include A, AAAA, A and AAAA (A_AAAA), CNAME, and SRV.
	DnsRecordType DnsRecordType `json:"dnsRecordType" yaml:"dnsRecordType"`
	// The amount of time, in seconds, that you want DNS resolvers to cache the settings for this record.
	DnsTtl awscdk.Duration `json:"dnsTtl" yaml:"dnsTtl"`
	// Whether or not this service will have an Elastic LoadBalancer registered to it as an AliasTargetInstance.
	//
	// Setting this to `true` correctly configures the `routingPolicy`
	// and performs some additional validation.
	LoadBalancer *bool `json:"loadBalancer" yaml:"loadBalancer"`
	// The routing policy that you want to apply to all DNS records that AWS Cloud Map creates when you register an instance and specify this service.
	RoutingPolicy RoutingPolicy `json:"routingPolicy" yaml:"routingPolicy"`
}

// Settings for an optional Amazon Route 53 health check.
//
// If you specify settings for a health check, AWS Cloud Map
// associates the health check with all the records that you specify in DnsConfig. Only valid with a PublicDnsNamespace.
//
// TODO: EXAMPLE
//
type HealthCheckConfig struct {
	// The number of consecutive health checks that an endpoint must pass or fail for Route 53 to change the current status of the endpoint from unhealthy to healthy or vice versa.
	FailureThreshold *float64 `json:"failureThreshold" yaml:"failureThreshold"`
	// The path that you want Route 53 to request when performing health checks.
	//
	// Do not use when health check type is TCP.
	ResourcePath *string `json:"resourcePath" yaml:"resourcePath"`
	// The type of health check that you want to create, which indicates how Route 53 determines whether an endpoint is healthy.
	//
	// Cannot be modified once created. Supported values are HTTP, HTTPS, and TCP.
	Type HealthCheckType `json:"type" yaml:"type"`
}

// Specifies information about an optional custom health check.
//
// TODO: EXAMPLE
//
type HealthCheckCustomConfig struct {
	// The number of 30-second intervals that you want Cloud Map to wait after receiving an UpdateInstanceCustomHealthStatus request before it changes the health status of a service instance.
	FailureThreshold *float64 `json:"failureThreshold" yaml:"failureThreshold"`
}

// TODO: EXAMPLE
//
type HealthCheckType string

const (
	HealthCheckType_HTTP HealthCheckType = "HTTP"
	HealthCheckType_HTTPS HealthCheckType = "HTTPS"
	HealthCheckType_TCP HealthCheckType = "TCP"
)

// Define an HTTP Namespace.
//
// TODO: EXAMPLE
//
type HttpNamespace interface {
	awscdk.Resource
	IHttpNamespace
	Env() *awscdk.ResourceEnvironment
	HttpNamespaceArn() *string
	HttpNamespaceId() *string
	HttpNamespaceName() *string
	NamespaceArn() *string
	NamespaceId() *string
	NamespaceName() *string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	Type() NamespaceType
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	CreateService(id *string, props *BaseServiceProps) Service
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for HttpNamespace
type jsiiProxy_HttpNamespace struct {
	internal.Type__awscdkResource
	jsiiProxy_IHttpNamespace
}

func (j *jsiiProxy_HttpNamespace) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpNamespace) HttpNamespaceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpNamespaceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpNamespace) HttpNamespaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpNamespaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpNamespace) HttpNamespaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpNamespaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpNamespace) NamespaceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpNamespace) NamespaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpNamespace) NamespaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpNamespace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpNamespace) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpNamespace) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpNamespace) Type() NamespaceType {
	var returns NamespaceType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewHttpNamespace(scope constructs.Construct, id *string, props *HttpNamespaceProps) HttpNamespace {
	_init_.Initialize()

	j := jsiiProxy_HttpNamespace{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.HttpNamespace",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewHttpNamespace_Override(h HttpNamespace, scope constructs.Construct, id *string, props *HttpNamespaceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.HttpNamespace",
		[]interface{}{scope, id, props},
		h,
	)
}

func HttpNamespace_FromHttpNamespaceAttributes(scope constructs.Construct, id *string, attrs *HttpNamespaceAttributes) IHttpNamespace {
	_init_.Initialize()

	var returns IHttpNamespace

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.HttpNamespace",
		"fromHttpNamespaceAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func HttpNamespace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.HttpNamespace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func HttpNamespace_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.HttpNamespace",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (h *jsiiProxy_HttpNamespace) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		h,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Creates a service within the namespace.
func (h *jsiiProxy_HttpNamespace) CreateService(id *string, props *BaseServiceProps) Service {
	var returns Service

	_jsii_.Invoke(
		h,
		"createService",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpNamespace) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (h *jsiiProxy_HttpNamespace) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (h *jsiiProxy_HttpNamespace) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (h *jsiiProxy_HttpNamespace) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// TODO: EXAMPLE
//
type HttpNamespaceAttributes struct {
	// Namespace ARN for the Namespace.
	NamespaceArn *string `json:"namespaceArn" yaml:"namespaceArn"`
	// Namespace Id for the Namespace.
	NamespaceId *string `json:"namespaceId" yaml:"namespaceId"`
	// A name for the Namespace.
	NamespaceName *string `json:"namespaceName" yaml:"namespaceName"`
}

// TODO: EXAMPLE
//
type HttpNamespaceProps struct {
	// A name for the Namespace.
	Name *string `json:"name" yaml:"name"`
	// A description of the Namespace.
	Description *string `json:"description" yaml:"description"`
}

type IHttpNamespace interface {
	INamespace
}

// The jsii proxy for IHttpNamespace
type jsiiProxy_IHttpNamespace struct {
	jsiiProxy_INamespace
}

type IInstance interface {
	awscdk.IResource
	// The id of the instance resource.
	InstanceId() *string
	// The Cloudmap service this resource is registered to.
	Service() IService
}

// The jsii proxy for IInstance
type jsiiProxy_IInstance struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IInstance) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstance) Service() IService {
	var returns IService
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

type INamespace interface {
	awscdk.IResource
	// Namespace ARN for the Namespace.
	NamespaceArn() *string
	// Namespace Id for the Namespace.
	NamespaceId() *string
	// A name for the Namespace.
	NamespaceName() *string
	// Type of Namespace.
	Type() NamespaceType
}

// The jsii proxy for INamespace
type jsiiProxy_INamespace struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_INamespace) NamespaceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INamespace) NamespaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INamespace) NamespaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INamespace) Type() NamespaceType {
	var returns NamespaceType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

type IPrivateDnsNamespace interface {
	INamespace
}

// The jsii proxy for IPrivateDnsNamespace
type jsiiProxy_IPrivateDnsNamespace struct {
	jsiiProxy_INamespace
}

type IPublicDnsNamespace interface {
	INamespace
}

// The jsii proxy for IPublicDnsNamespace
type jsiiProxy_IPublicDnsNamespace struct {
	jsiiProxy_INamespace
}

type IService interface {
	awscdk.IResource
	// The DnsRecordType used by the service.
	DnsRecordType() DnsRecordType
	// The namespace for the Cloudmap Service.
	Namespace() INamespace
	// The Routing Policy used by the service.
	RoutingPolicy() RoutingPolicy
	// The Arn of the namespace that you want to use for DNS configuration.
	ServiceArn() *string
	// The ID of the namespace that you want to use for DNS configuration.
	ServiceId() *string
	// A name for the Cloudmap Service.
	ServiceName() *string
}

// The jsii proxy for IService
type jsiiProxy_IService struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IService) DnsRecordType() DnsRecordType {
	var returns DnsRecordType
	_jsii_.Get(
		j,
		"dnsRecordType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IService) Namespace() INamespace {
	var returns INamespace
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IService) RoutingPolicy() RoutingPolicy {
	var returns RoutingPolicy
	_jsii_.Get(
		j,
		"routingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IService) ServiceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IService) ServiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IService) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

type InstanceBase interface {
	awscdk.Resource
	IInstance
	Env() *awscdk.ResourceEnvironment
	InstanceId() *string
	Node() constructs.Node
	PhysicalName() *string
	Service() IService
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
	UniqueInstanceId() *string
}

// The jsii proxy struct for InstanceBase
type jsiiProxy_InstanceBase struct {
	internal.Type__awscdkResource
	jsiiProxy_IInstance
}

func (j *jsiiProxy_InstanceBase) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceBase) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceBase) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceBase) Service() IService {
	var returns IService
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewInstanceBase_Override(i InstanceBase, scope constructs.Construct, id *string, props *awscdk.ResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.InstanceBase",
		[]interface{}{scope, id, props},
		i,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func InstanceBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.InstanceBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func InstanceBase_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.InstanceBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (i *jsiiProxy_InstanceBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_InstanceBase) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (i *jsiiProxy_InstanceBase) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (i *jsiiProxy_InstanceBase) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (i *jsiiProxy_InstanceBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Generate a unique instance Id that is safe to pass to CloudMap.
func (i *jsiiProxy_InstanceBase) UniqueInstanceId() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"uniqueInstanceId",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Instance that is accessible using an IP address.
//
// TODO: EXAMPLE
//
type IpInstance interface {
	InstanceBase
	Env() *awscdk.ResourceEnvironment
	InstanceId() *string
	Ipv4() *string
	Ipv6() *string
	Node() constructs.Node
	PhysicalName() *string
	Port() *float64
	Service() IService
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
	UniqueInstanceId() *string
}

// The jsii proxy struct for IpInstance
type jsiiProxy_IpInstance struct {
	jsiiProxy_InstanceBase
}

func (j *jsiiProxy_IpInstance) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpInstance) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpInstance) Ipv4() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpInstance) Ipv6() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpInstance) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpInstance) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpInstance) Service() IService {
	var returns IService
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpInstance) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewIpInstance(scope constructs.Construct, id *string, props *IpInstanceProps) IpInstance {
	_init_.Initialize()

	j := jsiiProxy_IpInstance{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.IpInstance",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewIpInstance_Override(i IpInstance, scope constructs.Construct, id *string, props *IpInstanceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.IpInstance",
		[]interface{}{scope, id, props},
		i,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func IpInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.IpInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func IpInstance_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.IpInstance",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (i *jsiiProxy_IpInstance) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IpInstance) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (i *jsiiProxy_IpInstance) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (i *jsiiProxy_IpInstance) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (i *jsiiProxy_IpInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Generate a unique instance Id that is safe to pass to CloudMap.
func (i *jsiiProxy_IpInstance) UniqueInstanceId() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"uniqueInstanceId",
		nil, // no parameters
		&returns,
	)

	return returns
}

// TODO: EXAMPLE
//
type IpInstanceBaseProps struct {
	// Custom attributes of the instance.
	CustomAttributes *map[string]*string `json:"customAttributes" yaml:"customAttributes"`
	// The id of the instance resource.
	InstanceId *string `json:"instanceId" yaml:"instanceId"`
	// If the service that you specify contains a template for an A record, the IPv4 address that you want AWS Cloud Map to use for the value of the A record.
	Ipv4 *string `json:"ipv4" yaml:"ipv4"`
	// If the service that you specify contains a template for an AAAA record, the IPv6 address that you want AWS Cloud Map to use for the value of the AAAA record.
	Ipv6 *string `json:"ipv6" yaml:"ipv6"`
	// The port on the endpoint that you want AWS Cloud Map to perform health checks on.
	//
	// This value is also used for
	// the port value in an SRV record if the service that you specify includes an SRV record. You can also specify a
	// default port that is applied to all instances in the Service configuration.
	Port *float64 `json:"port" yaml:"port"`
}

// TODO: EXAMPLE
//
type IpInstanceProps struct {
	// Custom attributes of the instance.
	CustomAttributes *map[string]*string `json:"customAttributes" yaml:"customAttributes"`
	// The id of the instance resource.
	InstanceId *string `json:"instanceId" yaml:"instanceId"`
	// If the service that you specify contains a template for an A record, the IPv4 address that you want AWS Cloud Map to use for the value of the A record.
	Ipv4 *string `json:"ipv4" yaml:"ipv4"`
	// If the service that you specify contains a template for an AAAA record, the IPv6 address that you want AWS Cloud Map to use for the value of the AAAA record.
	Ipv6 *string `json:"ipv6" yaml:"ipv6"`
	// The port on the endpoint that you want AWS Cloud Map to perform health checks on.
	//
	// This value is also used for
	// the port value in an SRV record if the service that you specify includes an SRV record. You can also specify a
	// default port that is applied to all instances in the Service configuration.
	Port *float64 `json:"port" yaml:"port"`
	// The Cloudmap service this resource is registered to.
	Service IService `json:"service" yaml:"service"`
}

type NamespaceType string

const (
	NamespaceType_HTTP NamespaceType = "HTTP"
	NamespaceType_DNS_PRIVATE NamespaceType = "DNS_PRIVATE"
	NamespaceType_DNS_PUBLIC NamespaceType = "DNS_PUBLIC"
)

// Instance accessible using values other than an IP address or a domain name (CNAME).
//
// Specify the other values in Custom attributes.
//
// TODO: EXAMPLE
//
type NonIpInstance interface {
	InstanceBase
	Env() *awscdk.ResourceEnvironment
	InstanceId() *string
	Node() constructs.Node
	PhysicalName() *string
	Service() IService
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
	UniqueInstanceId() *string
}

// The jsii proxy struct for NonIpInstance
type jsiiProxy_NonIpInstance struct {
	jsiiProxy_InstanceBase
}

func (j *jsiiProxy_NonIpInstance) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NonIpInstance) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NonIpInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NonIpInstance) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NonIpInstance) Service() IService {
	var returns IService
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NonIpInstance) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewNonIpInstance(scope constructs.Construct, id *string, props *NonIpInstanceProps) NonIpInstance {
	_init_.Initialize()

	j := jsiiProxy_NonIpInstance{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.NonIpInstance",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewNonIpInstance_Override(n NonIpInstance, scope constructs.Construct, id *string, props *NonIpInstanceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.NonIpInstance",
		[]interface{}{scope, id, props},
		n,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func NonIpInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.NonIpInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func NonIpInstance_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.NonIpInstance",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (n *jsiiProxy_NonIpInstance) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		n,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (n *jsiiProxy_NonIpInstance) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (n *jsiiProxy_NonIpInstance) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (n *jsiiProxy_NonIpInstance) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (n *jsiiProxy_NonIpInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Generate a unique instance Id that is safe to pass to CloudMap.
func (n *jsiiProxy_NonIpInstance) UniqueInstanceId() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"uniqueInstanceId",
		nil, // no parameters
		&returns,
	)

	return returns
}

// TODO: EXAMPLE
//
type NonIpInstanceBaseProps struct {
	// Custom attributes of the instance.
	CustomAttributes *map[string]*string `json:"customAttributes" yaml:"customAttributes"`
	// The id of the instance resource.
	InstanceId *string `json:"instanceId" yaml:"instanceId"`
}

// TODO: EXAMPLE
//
type NonIpInstanceProps struct {
	// Custom attributes of the instance.
	CustomAttributes *map[string]*string `json:"customAttributes" yaml:"customAttributes"`
	// The id of the instance resource.
	InstanceId *string `json:"instanceId" yaml:"instanceId"`
	// The Cloudmap service this resource is registered to.
	Service IService `json:"service" yaml:"service"`
}

// Define a Service Discovery HTTP Namespace.
//
// TODO: EXAMPLE
//
type PrivateDnsNamespace interface {
	awscdk.Resource
	IPrivateDnsNamespace
	Env() *awscdk.ResourceEnvironment
	NamespaceArn() *string
	NamespaceId() *string
	NamespaceName() *string
	Node() constructs.Node
	PhysicalName() *string
	PrivateDnsNamespaceArn() *string
	PrivateDnsNamespaceId() *string
	PrivateDnsNamespaceName() *string
	Stack() awscdk.Stack
	Type() NamespaceType
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	CreateService(id *string, props *DnsServiceProps) Service
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for PrivateDnsNamespace
type jsiiProxy_PrivateDnsNamespace struct {
	internal.Type__awscdkResource
	jsiiProxy_IPrivateDnsNamespace
}

func (j *jsiiProxy_PrivateDnsNamespace) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateDnsNamespace) NamespaceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateDnsNamespace) NamespaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateDnsNamespace) NamespaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateDnsNamespace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateDnsNamespace) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateDnsNamespace) PrivateDnsNamespaceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsNamespaceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateDnsNamespace) PrivateDnsNamespaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsNamespaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateDnsNamespace) PrivateDnsNamespaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsNamespaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateDnsNamespace) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateDnsNamespace) Type() NamespaceType {
	var returns NamespaceType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewPrivateDnsNamespace(scope constructs.Construct, id *string, props *PrivateDnsNamespaceProps) PrivateDnsNamespace {
	_init_.Initialize()

	j := jsiiProxy_PrivateDnsNamespace{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.PrivateDnsNamespace",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewPrivateDnsNamespace_Override(p PrivateDnsNamespace, scope constructs.Construct, id *string, props *PrivateDnsNamespaceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.PrivateDnsNamespace",
		[]interface{}{scope, id, props},
		p,
	)
}

func PrivateDnsNamespace_FromPrivateDnsNamespaceAttributes(scope constructs.Construct, id *string, attrs *PrivateDnsNamespaceAttributes) IPrivateDnsNamespace {
	_init_.Initialize()

	var returns IPrivateDnsNamespace

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.PrivateDnsNamespace",
		"fromPrivateDnsNamespaceAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func PrivateDnsNamespace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.PrivateDnsNamespace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func PrivateDnsNamespace_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.PrivateDnsNamespace",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (p *jsiiProxy_PrivateDnsNamespace) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Creates a service within the namespace.
func (p *jsiiProxy_PrivateDnsNamespace) CreateService(id *string, props *DnsServiceProps) Service {
	var returns Service

	_jsii_.Invoke(
		p,
		"createService",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateDnsNamespace) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (p *jsiiProxy_PrivateDnsNamespace) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (p *jsiiProxy_PrivateDnsNamespace) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (p *jsiiProxy_PrivateDnsNamespace) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// TODO: EXAMPLE
//
type PrivateDnsNamespaceAttributes struct {
	// Namespace ARN for the Namespace.
	NamespaceArn *string `json:"namespaceArn" yaml:"namespaceArn"`
	// Namespace Id for the Namespace.
	NamespaceId *string `json:"namespaceId" yaml:"namespaceId"`
	// A name for the Namespace.
	NamespaceName *string `json:"namespaceName" yaml:"namespaceName"`
}

// TODO: EXAMPLE
//
type PrivateDnsNamespaceProps struct {
	// A name for the Namespace.
	Name *string `json:"name" yaml:"name"`
	// A description of the Namespace.
	Description *string `json:"description" yaml:"description"`
	// The Amazon VPC that you want to associate the namespace with.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
}

// Define a Public DNS Namespace.
//
// TODO: EXAMPLE
//
type PublicDnsNamespace interface {
	awscdk.Resource
	IPublicDnsNamespace
	Env() *awscdk.ResourceEnvironment
	NamespaceArn() *string
	NamespaceId() *string
	NamespaceName() *string
	Node() constructs.Node
	PhysicalName() *string
	PublicDnsNamespaceArn() *string
	PublicDnsNamespaceId() *string
	PublicDnsNamespaceName() *string
	Stack() awscdk.Stack
	Type() NamespaceType
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	CreateService(id *string, props *DnsServiceProps) Service
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for PublicDnsNamespace
type jsiiProxy_PublicDnsNamespace struct {
	internal.Type__awscdkResource
	jsiiProxy_IPublicDnsNamespace
}

func (j *jsiiProxy_PublicDnsNamespace) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) NamespaceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) NamespaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) NamespaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) PublicDnsNamespaceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicDnsNamespaceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) PublicDnsNamespaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicDnsNamespaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) PublicDnsNamespaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicDnsNamespaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) Type() NamespaceType {
	var returns NamespaceType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewPublicDnsNamespace(scope constructs.Construct, id *string, props *PublicDnsNamespaceProps) PublicDnsNamespace {
	_init_.Initialize()

	j := jsiiProxy_PublicDnsNamespace{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.PublicDnsNamespace",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewPublicDnsNamespace_Override(p PublicDnsNamespace, scope constructs.Construct, id *string, props *PublicDnsNamespaceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.PublicDnsNamespace",
		[]interface{}{scope, id, props},
		p,
	)
}

func PublicDnsNamespace_FromPublicDnsNamespaceAttributes(scope constructs.Construct, id *string, attrs *PublicDnsNamespaceAttributes) IPublicDnsNamespace {
	_init_.Initialize()

	var returns IPublicDnsNamespace

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.PublicDnsNamespace",
		"fromPublicDnsNamespaceAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func PublicDnsNamespace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.PublicDnsNamespace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func PublicDnsNamespace_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.PublicDnsNamespace",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (p *jsiiProxy_PublicDnsNamespace) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Creates a service within the namespace.
func (p *jsiiProxy_PublicDnsNamespace) CreateService(id *string, props *DnsServiceProps) Service {
	var returns Service

	_jsii_.Invoke(
		p,
		"createService",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PublicDnsNamespace) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (p *jsiiProxy_PublicDnsNamespace) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (p *jsiiProxy_PublicDnsNamespace) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (p *jsiiProxy_PublicDnsNamespace) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// TODO: EXAMPLE
//
type PublicDnsNamespaceAttributes struct {
	// Namespace ARN for the Namespace.
	NamespaceArn *string `json:"namespaceArn" yaml:"namespaceArn"`
	// Namespace Id for the Namespace.
	NamespaceId *string `json:"namespaceId" yaml:"namespaceId"`
	// A name for the Namespace.
	NamespaceName *string `json:"namespaceName" yaml:"namespaceName"`
}

// TODO: EXAMPLE
//
type PublicDnsNamespaceProps struct {
	// A name for the Namespace.
	Name *string `json:"name" yaml:"name"`
	// A description of the Namespace.
	Description *string `json:"description" yaml:"description"`
}

type RoutingPolicy string

const (
	RoutingPolicy_WEIGHTED RoutingPolicy = "WEIGHTED"
	RoutingPolicy_MULTIVALUE RoutingPolicy = "MULTIVALUE"
)

// Define a CloudMap Service.
//
// TODO: EXAMPLE
//
type Service interface {
	awscdk.Resource
	IService
	DnsRecordType() DnsRecordType
	Env() *awscdk.ResourceEnvironment
	Namespace() INamespace
	Node() constructs.Node
	PhysicalName() *string
	RoutingPolicy() RoutingPolicy
	ServiceArn() *string
	ServiceId() *string
	ServiceName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	RegisterCnameInstance(id *string, props *CnameInstanceBaseProps) IInstance
	RegisterIpInstance(id *string, props *IpInstanceBaseProps) IInstance
	RegisterLoadBalancer(id *string, loadBalancer awselasticloadbalancingv2.ILoadBalancerV2, customAttributes *map[string]*string) IInstance
	RegisterNonIpInstance(id *string, props *NonIpInstanceBaseProps) IInstance
	ToString() *string
}

// The jsii proxy struct for Service
type jsiiProxy_Service struct {
	internal.Type__awscdkResource
	jsiiProxy_IService
}

func (j *jsiiProxy_Service) DnsRecordType() DnsRecordType {
	var returns DnsRecordType
	_jsii_.Get(
		j,
		"dnsRecordType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Namespace() INamespace {
	var returns INamespace
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) RoutingPolicy() RoutingPolicy {
	var returns RoutingPolicy
	_jsii_.Get(
		j,
		"routingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ServiceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ServiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewService(scope constructs.Construct, id *string, props *ServiceProps) Service {
	_init_.Initialize()

	j := jsiiProxy_Service{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.Service",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewService_Override(s Service, scope constructs.Construct, id *string, props *ServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.Service",
		[]interface{}{scope, id, props},
		s,
	)
}

func Service_FromServiceAttributes(scope constructs.Construct, id *string, attrs *ServiceAttributes) IService {
	_init_.Initialize()

	var returns IService

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.Service",
		"fromServiceAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Service_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.Service",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Service_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.Service",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (s *jsiiProxy_Service) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_Service) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (s *jsiiProxy_Service) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (s *jsiiProxy_Service) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Registers a resource that is accessible using a CNAME.
func (s *jsiiProxy_Service) RegisterCnameInstance(id *string, props *CnameInstanceBaseProps) IInstance {
	var returns IInstance

	_jsii_.Invoke(
		s,
		"registerCnameInstance",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

// Registers a resource that is accessible using an IP address.
func (s *jsiiProxy_Service) RegisterIpInstance(id *string, props *IpInstanceBaseProps) IInstance {
	var returns IInstance

	_jsii_.Invoke(
		s,
		"registerIpInstance",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

// Registers an ELB as a new instance with unique name instanceId in this service.
func (s *jsiiProxy_Service) RegisterLoadBalancer(id *string, loadBalancer awselasticloadbalancingv2.ILoadBalancerV2, customAttributes *map[string]*string) IInstance {
	var returns IInstance

	_jsii_.Invoke(
		s,
		"registerLoadBalancer",
		[]interface{}{id, loadBalancer, customAttributes},
		&returns,
	)

	return returns
}

// Registers a resource that is accessible using values other than an IP address or a domain name (CNAME).
func (s *jsiiProxy_Service) RegisterNonIpInstance(id *string, props *NonIpInstanceBaseProps) IInstance {
	var returns IInstance

	_jsii_.Invoke(
		s,
		"registerNonIpInstance",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_Service) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// TODO: EXAMPLE
//
type ServiceAttributes struct {
	DnsRecordType DnsRecordType `json:"dnsRecordType" yaml:"dnsRecordType"`
	Namespace INamespace `json:"namespace" yaml:"namespace"`
	RoutingPolicy RoutingPolicy `json:"routingPolicy" yaml:"routingPolicy"`
	ServiceArn *string `json:"serviceArn" yaml:"serviceArn"`
	ServiceId *string `json:"serviceId" yaml:"serviceId"`
	ServiceName *string `json:"serviceName" yaml:"serviceName"`
}

// TODO: EXAMPLE
//
type ServiceProps struct {
	// Structure containing failure threshold for a custom health checker.
	//
	// Only one of healthCheckConfig or healthCheckCustomConfig can be specified.
	// See: https://docs.aws.amazon.com/cloud-map/latest/api/API_HealthCheckCustomConfig.html
	CustomHealthCheck *HealthCheckCustomConfig `json:"customHealthCheck" yaml:"customHealthCheck"`
	// A description of the service.
	Description *string `json:"description" yaml:"description"`
	// Settings for an optional health check.
	//
	// If you specify health check settings, AWS Cloud Map associates the health
	// check with the records that you specify in DnsConfig. Only one of healthCheckConfig or healthCheckCustomConfig can
	// be specified. Not valid for PrivateDnsNamespaces. If you use healthCheck, you can only register IP instances to
	// this service.
	HealthCheck *HealthCheckConfig `json:"healthCheck" yaml:"healthCheck"`
	// A name for the Service.
	Name *string `json:"name" yaml:"name"`
	// The DNS type of the record that you want AWS Cloud Map to create.
	//
	// Supported record types
	// include A, AAAA, A and AAAA (A_AAAA), CNAME, and SRV.
	DnsRecordType DnsRecordType `json:"dnsRecordType" yaml:"dnsRecordType"`
	// The amount of time, in seconds, that you want DNS resolvers to cache the settings for this record.
	DnsTtl awscdk.Duration `json:"dnsTtl" yaml:"dnsTtl"`
	// Whether or not this service will have an Elastic LoadBalancer registered to it as an AliasTargetInstance.
	//
	// Setting this to `true` correctly configures the `routingPolicy`
	// and performs some additional validation.
	LoadBalancer *bool `json:"loadBalancer" yaml:"loadBalancer"`
	// The routing policy that you want to apply to all DNS records that AWS Cloud Map creates when you register an instance and specify this service.
	RoutingPolicy RoutingPolicy `json:"routingPolicy" yaml:"routingPolicy"`
	// The namespace that you want to use for DNS configuration.
	Namespace INamespace `json:"namespace" yaml:"namespace"`
}

