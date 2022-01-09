package awsglobalaccelerator

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglobalaccelerator/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The Accelerator construct.
//
// TODO: EXAMPLE
//
type Accelerator interface {
	awscdk.Resource
	IAccelerator
	AcceleratorArn() *string
	DnsName() *string
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	AddListener(id *string, options *ListenerOptions) Listener
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for Accelerator
type jsiiProxy_Accelerator struct {
	internal.Type__awscdkResource
	jsiiProxy_IAccelerator
}

func (j *jsiiProxy_Accelerator) AcceleratorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceleratorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Accelerator) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Accelerator) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Accelerator) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Accelerator) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Accelerator) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewAccelerator(scope constructs.Construct, id *string, props *AcceleratorProps) Accelerator {
	_init_.Initialize()

	j := jsiiProxy_Accelerator{}

	_jsii_.Create(
		"aws-cdk-lib.aws_globalaccelerator.Accelerator",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewAccelerator_Override(a Accelerator, scope constructs.Construct, id *string, props *AcceleratorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_globalaccelerator.Accelerator",
		[]interface{}{scope, id, props},
		a,
	)
}

// import from attributes.
func Accelerator_FromAcceleratorAttributes(scope constructs.Construct, id *string, attrs *AcceleratorAttributes) IAccelerator {
	_init_.Initialize()

	var returns IAccelerator

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_globalaccelerator.Accelerator",
		"fromAcceleratorAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Accelerator_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_globalaccelerator.Accelerator",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Accelerator_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_globalaccelerator.Accelerator",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add a listener to the accelerator.
func (a *jsiiProxy_Accelerator) AddListener(id *string, options *ListenerOptions) Listener {
	var returns Listener

	_jsii_.Invoke(
		a,
		"addListener",
		[]interface{}{id, options},
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
func (a *jsiiProxy_Accelerator) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (a *jsiiProxy_Accelerator) GeneratePhysicalName() *string {
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
func (a *jsiiProxy_Accelerator) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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
func (a *jsiiProxy_Accelerator) GetResourceNameAttribute(nameAttr *string) *string {
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
func (a *jsiiProxy_Accelerator) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Attributes required to import an existing accelerator to the stack.
//
// TODO: EXAMPLE
//
type AcceleratorAttributes struct {
	// The ARN of the accelerator.
	AcceleratorArn *string `json:"acceleratorArn"`
	// The DNS name of the accelerator.
	DnsName *string `json:"dnsName"`
}

// Construct properties of the Accelerator.
//
// TODO: EXAMPLE
//
type AcceleratorProps struct {
	// The name of the accelerator.
	AcceleratorName *string `json:"acceleratorName"`
	// Indicates whether the accelerator is enabled.
	Enabled *bool `json:"enabled"`
}

// A CloudFormation `AWS::GlobalAccelerator::Accelerator`.
//
// The `AWS::GlobalAccelerator::Accelerator` resource is a Global Accelerator resource type that contains information about how you create an accelerator. An accelerator includes one or more listeners that process inbound connections and direct traffic to one or more endpoint groups, each of which includes endpoints, such as Application Load Balancers, Network Load Balancers, and Amazon EC2 instances.
//
// TODO: EXAMPLE
//
type CfnAccelerator interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrAcceleratorArn() *string
	AttrDnsName() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Enabled() interface{}
	SetEnabled(val interface{})
	IpAddresses() *[]*string
	SetIpAddresses(val *[]*string)
	IpAddressType() *string
	SetIpAddressType(val *string)
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

// The jsii proxy struct for CfnAccelerator
type jsiiProxy_CfnAccelerator struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAccelerator) AttrAcceleratorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAcceleratorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccelerator) AttrDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccelerator) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccelerator) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccelerator) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccelerator) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccelerator) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccelerator) IpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccelerator) IpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccelerator) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccelerator) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccelerator) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccelerator) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccelerator) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccelerator) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccelerator) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::GlobalAccelerator::Accelerator`.
func NewCfnAccelerator(scope constructs.Construct, id *string, props *CfnAcceleratorProps) CfnAccelerator {
	_init_.Initialize()

	j := jsiiProxy_CfnAccelerator{}

	_jsii_.Create(
		"aws-cdk-lib.aws_globalaccelerator.CfnAccelerator",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GlobalAccelerator::Accelerator`.
func NewCfnAccelerator_Override(c CfnAccelerator, scope constructs.Construct, id *string, props *CfnAcceleratorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_globalaccelerator.CfnAccelerator",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAccelerator) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CfnAccelerator) SetIpAddresses(val *[]*string) {
	_jsii_.Set(
		j,
		"ipAddresses",
		val,
	)
}

func (j *jsiiProxy_CfnAccelerator) SetIpAddressType(val *string) {
	_jsii_.Set(
		j,
		"ipAddressType",
		val,
	)
}

func (j *jsiiProxy_CfnAccelerator) SetName(val *string) {
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
func CfnAccelerator_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_globalaccelerator.CfnAccelerator",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnAccelerator_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_globalaccelerator.CfnAccelerator",
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
func CfnAccelerator_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_globalaccelerator.CfnAccelerator",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAccelerator_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_globalaccelerator.CfnAccelerator",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnAccelerator) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnAccelerator) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnAccelerator) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnAccelerator) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnAccelerator) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnAccelerator) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnAccelerator) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnAccelerator) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnAccelerator) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnAccelerator) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnAccelerator) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAccelerator) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnAccelerator) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnAccelerator) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAccelerator) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnAccelerator`.
//
// TODO: EXAMPLE
//
type CfnAcceleratorProps struct {
	// The name of the accelerator.
	//
	// The name must contain only alphanumeric characters or hyphens (-), and must not begin or end with a hyphen.
	Name *string `json:"name"`
	// Indicates whether the accelerator is enabled. The value is true or false. The default value is true.
	//
	// If the value is set to true, the accelerator cannot be deleted. If set to false, accelerator can be deleted.
	Enabled interface{} `json:"enabled"`
	// Optionally, if you've added your own IP address pool to Global Accelerator (BYOIP), you can choose IP addresses from your own pool to use for the accelerator's static IP addresses when you create an accelerator.
	//
	// You can specify one or two addresses, separated by a comma. Do not include the /32 suffix.
	//
	// Only one IP address from each of your IP address ranges can be used for each accelerator. If you specify only one IP address from your IP address range, Global Accelerator assigns a second static IP address for the accelerator from the AWS IP address pool.
	//
	// Note that you can't update IP addresses for an existing accelerator. To change them, you must create a new accelerator with the new addresses.
	//
	// For more information, see [Bring Your Own IP Addresses (BYOIP)](https://docs.aws.amazon.com/global-accelerator/latest/dg/using-byoip.html) in the *AWS Global Accelerator Developer Guide* .
	IpAddresses *[]*string `json:"ipAddresses"`
	// The value for the address type must be IPv4.
	IpAddressType *string `json:"ipAddressType"`
	// Create tags for an accelerator.
	//
	// For more information, see [Tagging](https://docs.aws.amazon.com/global-accelerator/latest/dg/tagging-in-global-accelerator.html) in the *AWS Global Accelerator Developer Guide* .
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::GlobalAccelerator::EndpointGroup`.
//
// The `AWS::GlobalAccelerator::EndpointGroup` resource is a Global Accelerator resource type that contains information about how you create an endpoint group for the specified listener. An endpoint group is a collection of endpoints in one AWS Region .
//
// TODO: EXAMPLE
//
type CfnEndpointGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrEndpointGroupArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	EndpointConfigurations() interface{}
	SetEndpointConfigurations(val interface{})
	EndpointGroupRegion() *string
	SetEndpointGroupRegion(val *string)
	HealthCheckIntervalSeconds() *float64
	SetHealthCheckIntervalSeconds(val *float64)
	HealthCheckPath() *string
	SetHealthCheckPath(val *string)
	HealthCheckPort() *float64
	SetHealthCheckPort(val *float64)
	HealthCheckProtocol() *string
	SetHealthCheckProtocol(val *string)
	ListenerArn() *string
	SetListenerArn(val *string)
	LogicalId() *string
	Node() constructs.Node
	PortOverrides() interface{}
	SetPortOverrides(val interface{})
	Ref() *string
	Stack() awscdk.Stack
	ThresholdCount() *float64
	SetThresholdCount(val *float64)
	TrafficDialPercentage() *float64
	SetTrafficDialPercentage(val *float64)
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

// The jsii proxy struct for CfnEndpointGroup
type jsiiProxy_CfnEndpointGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnEndpointGroup) AttrEndpointGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpointGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) EndpointConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) EndpointGroupRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointGroupRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) HealthCheckIntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckIntervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) HealthCheckPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) HealthCheckPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) HealthCheckProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) ListenerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) PortOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) ThresholdCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) TrafficDialPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"trafficDialPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::GlobalAccelerator::EndpointGroup`.
func NewCfnEndpointGroup(scope constructs.Construct, id *string, props *CfnEndpointGroupProps) CfnEndpointGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnEndpointGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_globalaccelerator.CfnEndpointGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GlobalAccelerator::EndpointGroup`.
func NewCfnEndpointGroup_Override(c CfnEndpointGroup, scope constructs.Construct, id *string, props *CfnEndpointGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_globalaccelerator.CfnEndpointGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEndpointGroup) SetEndpointConfigurations(val interface{}) {
	_jsii_.Set(
		j,
		"endpointConfigurations",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointGroup) SetEndpointGroupRegion(val *string) {
	_jsii_.Set(
		j,
		"endpointGroupRegion",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointGroup) SetHealthCheckIntervalSeconds(val *float64) {
	_jsii_.Set(
		j,
		"healthCheckIntervalSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointGroup) SetHealthCheckPath(val *string) {
	_jsii_.Set(
		j,
		"healthCheckPath",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointGroup) SetHealthCheckPort(val *float64) {
	_jsii_.Set(
		j,
		"healthCheckPort",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointGroup) SetHealthCheckProtocol(val *string) {
	_jsii_.Set(
		j,
		"healthCheckProtocol",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointGroup) SetListenerArn(val *string) {
	_jsii_.Set(
		j,
		"listenerArn",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointGroup) SetPortOverrides(val interface{}) {
	_jsii_.Set(
		j,
		"portOverrides",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointGroup) SetThresholdCount(val *float64) {
	_jsii_.Set(
		j,
		"thresholdCount",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointGroup) SetTrafficDialPercentage(val *float64) {
	_jsii_.Set(
		j,
		"trafficDialPercentage",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnEndpointGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_globalaccelerator.CfnEndpointGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnEndpointGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_globalaccelerator.CfnEndpointGroup",
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
func CfnEndpointGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_globalaccelerator.CfnEndpointGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEndpointGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_globalaccelerator.CfnEndpointGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnEndpointGroup) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnEndpointGroup) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnEndpointGroup) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnEndpointGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnEndpointGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnEndpointGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnEndpointGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnEndpointGroup) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnEndpointGroup) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnEndpointGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnEndpointGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnEndpointGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnEndpointGroup) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnEndpointGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpointGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A complex type for endpoints.
//
// A resource must be valid and active when you add it as an endpoint.
//
// TODO: EXAMPLE
//
type CfnEndpointGroup_EndpointConfigurationProperty struct {
	// An ID for the endpoint.
	//
	// If the endpoint is a Network Load Balancer or Application Load Balancer, this is the Amazon Resource Name (ARN) of the resource. If the endpoint is an Elastic IP address, this is the Elastic IP address allocation ID. For Amazon EC2 instances, this is the EC2 instance ID. A resource must be valid and active when you add it as an endpoint.
	//
	// An Application Load Balancer can be either internal or internet-facing.
	EndpointId *string `json:"endpointId"`
	// Indicates whether client IP address preservation is enabled for an Application Load Balancer endpoint.
	//
	// The value is true or false. The default value is true for new accelerators.
	//
	// If the value is set to true, the client's IP address is preserved in the `X-Forwarded-For` request header as traffic travels to applications on the Application Load Balancer endpoint fronted by the accelerator.
	//
	// For more information, see [Preserve Client IP Addresses](https://docs.aws.amazon.com/global-accelerator/latest/dg/preserve-client-ip-address.html) in the *AWS Global Accelerator Developer Guide* .
	ClientIpPreservationEnabled interface{} `json:"clientIpPreservationEnabled"`
	// The weight associated with the endpoint.
	//
	// When you add weights to endpoints, you configure Global Accelerator to route traffic based on proportions that you specify. For example, you might specify endpoint weights of 4, 5, 5, and 6 (sum=20). The result is that 4/20 of your traffic, on average, is routed to the first endpoint, 5/20 is routed both to the second and third endpoints, and 6/20 is routed to the last endpoint. For more information, see [Endpoint Weights](https://docs.aws.amazon.com/global-accelerator/latest/dg/about-endpoints-endpoint-weights.html) in the *AWS Global Accelerator Developer Guide* .
	Weight *float64 `json:"weight"`
}

// Override specific listener ports used to route traffic to endpoints that are part of an endpoint group.
//
// For example, you can create a port override in which the listener receives user traffic on ports 80 and 443, but your accelerator routes that traffic to ports 1080 and 1443, respectively, on the endpoints.
//
// For more information, see [Port overrides](https://docs.aws.amazon.com/global-accelerator/latest/dg/about-endpoint-groups-port-override.html) in the *AWS Global Accelerator Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnEndpointGroup_PortOverrideProperty struct {
	// The endpoint port that you want a listener port to be mapped to.
	//
	// This is the port on the endpoint, such as the Application Load Balancer or Amazon EC2 instance.
	EndpointPort *float64 `json:"endpointPort"`
	// The listener port that you want to map to a specific endpoint port.
	//
	// This is the port that user traffic arrives to the Global Accelerator on.
	ListenerPort *float64 `json:"listenerPort"`
}

// Properties for defining a `CfnEndpointGroup`.
//
// TODO: EXAMPLE
//
type CfnEndpointGroupProps struct {
	// The AWS Regions where the endpoint group is located.
	EndpointGroupRegion *string `json:"endpointGroupRegion"`
	// The Amazon Resource Name (ARN) of the listener.
	ListenerArn *string `json:"listenerArn"`
	// The list of endpoint objects.
	EndpointConfigurations interface{} `json:"endpointConfigurations"`
	// The time—10 seconds or 30 seconds—between health checks for each endpoint.
	//
	// The default value is 30.
	HealthCheckIntervalSeconds *float64 `json:"healthCheckIntervalSeconds"`
	// If the protocol is HTTP/S, then this value provides the ping path that Global Accelerator uses for the destination on the endpoints for health checks.
	//
	// The default is slash (/).
	HealthCheckPath *string `json:"healthCheckPath"`
	// The port that Global Accelerator uses to perform health checks on endpoints that are part of this endpoint group.
	//
	// The default port is the port for the listener that this endpoint group is associated with. If the listener port is a list, Global Accelerator uses the first specified port in the list of ports.
	HealthCheckPort *float64 `json:"healthCheckPort"`
	// The protocol that Global Accelerator uses to perform health checks on endpoints that are part of this endpoint group.
	//
	// The default value is TCP.
	HealthCheckProtocol *string `json:"healthCheckProtocol"`
	// Allows you to override the destination ports used to route traffic to an endpoint.
	//
	// Using a port override lets you to map a list of external destination ports (that your users send traffic to) to a list of internal destination ports that you want an application endpoint to receive traffic on.
	PortOverrides interface{} `json:"portOverrides"`
	// The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy.
	//
	// The default value is 3.
	ThresholdCount *float64 `json:"thresholdCount"`
	// The percentage of traffic to send to an AWS Regions .
	//
	// Additional traffic is distributed to other endpoint groups for this listener.
	//
	// Use this action to increase (dial up) or decrease (dial down) traffic to a specific Region. The percentage is applied to the traffic that would otherwise have been routed to the Region based on optimal routing.
	//
	// The default value is 100.
	TrafficDialPercentage *float64 `json:"trafficDialPercentage"`
}

// A CloudFormation `AWS::GlobalAccelerator::Listener`.
//
// The `AWS::GlobalAccelerator::Listener` resource is a Global Accelerator resource type that contains information about how you create a listener to process inbound connections from clients to an accelerator. Connections arrive to assigned static IP addresses on a port, port range, or list of port ranges that you specify.
//
// TODO: EXAMPLE
//
type CfnListener interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AcceleratorArn() *string
	SetAcceleratorArn(val *string)
	AttrListenerArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ClientAffinity() *string
	SetClientAffinity(val *string)
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	PortRanges() interface{}
	SetPortRanges(val interface{})
	Protocol() *string
	SetProtocol(val *string)
	Ref() *string
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

// The jsii proxy struct for CfnListener
type jsiiProxy_CfnListener struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnListener) AcceleratorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceleratorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) AttrListenerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrListenerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) ClientAffinity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) PortRanges() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::GlobalAccelerator::Listener`.
func NewCfnListener(scope constructs.Construct, id *string, props *CfnListenerProps) CfnListener {
	_init_.Initialize()

	j := jsiiProxy_CfnListener{}

	_jsii_.Create(
		"aws-cdk-lib.aws_globalaccelerator.CfnListener",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GlobalAccelerator::Listener`.
func NewCfnListener_Override(c CfnListener, scope constructs.Construct, id *string, props *CfnListenerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_globalaccelerator.CfnListener",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnListener) SetAcceleratorArn(val *string) {
	_jsii_.Set(
		j,
		"acceleratorArn",
		val,
	)
}

func (j *jsiiProxy_CfnListener) SetClientAffinity(val *string) {
	_jsii_.Set(
		j,
		"clientAffinity",
		val,
	)
}

func (j *jsiiProxy_CfnListener) SetPortRanges(val interface{}) {
	_jsii_.Set(
		j,
		"portRanges",
		val,
	)
}

func (j *jsiiProxy_CfnListener) SetProtocol(val *string) {
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnListener_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_globalaccelerator.CfnListener",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnListener_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_globalaccelerator.CfnListener",
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
func CfnListener_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_globalaccelerator.CfnListener",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnListener_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_globalaccelerator.CfnListener",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnListener) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnListener) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnListener) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnListener) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnListener) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnListener) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnListener) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnListener) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnListener) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnListener) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnListener) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnListener) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnListener) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnListener) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnListener) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A complex type for a range of ports for a listener.
//
// TODO: EXAMPLE
//
type CfnListener_PortRangeProperty struct {
	// The first port in the range of ports, inclusive.
	FromPort *float64 `json:"fromPort"`
	// The last port in the range of ports, inclusive.
	ToPort *float64 `json:"toPort"`
}

// Properties for defining a `CfnListener`.
//
// TODO: EXAMPLE
//
type CfnListenerProps struct {
	// The Amazon Resource Name (ARN) of your accelerator.
	AcceleratorArn *string `json:"acceleratorArn"`
	// The list of port ranges for the connections from clients to the accelerator.
	PortRanges interface{} `json:"portRanges"`
	// The protocol for the connections from clients to the accelerator.
	Protocol *string `json:"protocol"`
	// Client affinity lets you direct all requests from a user to the same endpoint, if you have stateful applications, regardless of the port and protocol of the client request.
	//
	// Client affinity gives you control over whether to always route each client to the same specific endpoint.
	//
	// AWS Global Accelerator uses a consistent-flow hashing algorithm to choose the optimal endpoint for a connection. If client affinity is `NONE` , Global Accelerator uses the "five-tuple" (5-tuple) properties—source IP address, source port, destination IP address, destination port, and protocol—to select the hash value, and then chooses the best endpoint. However, with this setting, if someone uses different ports to connect to Global Accelerator, their connections might not be always routed to the same endpoint because the hash value changes.
	//
	// If you want a given client to always be routed to the same endpoint, set client affinity to `SOURCE_IP` instead. When you use the `SOURCE_IP` setting, Global Accelerator uses the "two-tuple" (2-tuple) properties— source (client) IP address and destination IP address—to select the hash value.
	//
	// The default value is `NONE` .
	ClientAffinity *string `json:"clientAffinity"`
}

// Client affinity gives you control over whether to always route each client to the same specific endpoint.
// See: https://docs.aws.amazon.com/global-accelerator/latest/dg/about-listeners.html#about-listeners-client-affinity
//
type ClientAffinity string

const (
	ClientAffinity_NONE ClientAffinity = "NONE"
	ClientAffinity_SOURCE_IP ClientAffinity = "SOURCE_IP"
)

// The protocol for the connections from clients to the accelerator.
type ConnectionProtocol string

const (
	ConnectionProtocol_TCP ConnectionProtocol = "TCP"
	ConnectionProtocol_UDP ConnectionProtocol = "UDP"
)

// EndpointGroup construct.
//
// TODO: EXAMPLE
//
type EndpointGroup interface {
	awscdk.Resource
	IEndpointGroup
	EndpointGroupArn() *string
	EndpointGroupName() *string
	Endpoints() *[]IEndpoint
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	AddEndpoint(endpoint IEndpoint)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	ConnectionsPeer(id *string, vpc awsec2.IVpc) awsec2.IPeer
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for EndpointGroup
type jsiiProxy_EndpointGroup struct {
	internal.Type__awscdkResource
	jsiiProxy_IEndpointGroup
}

func (j *jsiiProxy_EndpointGroup) EndpointGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointGroup) EndpointGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointGroup) Endpoints() *[]IEndpoint {
	var returns *[]IEndpoint
	_jsii_.Get(
		j,
		"endpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointGroup) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewEndpointGroup(scope constructs.Construct, id *string, props *EndpointGroupProps) EndpointGroup {
	_init_.Initialize()

	j := jsiiProxy_EndpointGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_globalaccelerator.EndpointGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewEndpointGroup_Override(e EndpointGroup, scope constructs.Construct, id *string, props *EndpointGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_globalaccelerator.EndpointGroup",
		[]interface{}{scope, id, props},
		e,
	)
}

// import from ARN.
func EndpointGroup_FromEndpointGroupArn(scope constructs.Construct, id *string, endpointGroupArn *string) IEndpointGroup {
	_init_.Initialize()

	var returns IEndpointGroup

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_globalaccelerator.EndpointGroup",
		"fromEndpointGroupArn",
		[]interface{}{scope, id, endpointGroupArn},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func EndpointGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_globalaccelerator.EndpointGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func EndpointGroup_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_globalaccelerator.EndpointGroup",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add an endpoint.
func (e *jsiiProxy_EndpointGroup) AddEndpoint(endpoint IEndpoint) {
	_jsii_.InvokeVoid(
		e,
		"addEndpoint",
		[]interface{}{endpoint},
	)
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
func (e *jsiiProxy_EndpointGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		e,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Return an object that represents the Accelerator's Security Group.
//
// Uses a Custom Resource to look up the Security Group that Accelerator
// creates at deploy time. Requires your VPC ID to perform the lookup.
//
// The Security Group will only be created if you enable **Client IP
// Preservation** on any of the endpoints.
//
// You cannot manipulate the rules inside this security group, but you can
// use this security group as a Peer in Connections rules on other
// constructs.
func (e *jsiiProxy_EndpointGroup) ConnectionsPeer(id *string, vpc awsec2.IVpc) awsec2.IPeer {
	var returns awsec2.IPeer

	_jsii_.Invoke(
		e,
		"connectionsPeer",
		[]interface{}{id, vpc},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EndpointGroup) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		e,
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
func (e *jsiiProxy_EndpointGroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		e,
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
func (e *jsiiProxy_EndpointGroup) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (e *jsiiProxy_EndpointGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Basic options for creating a new EndpointGroup.
//
// TODO: EXAMPLE
//
type EndpointGroupOptions struct {
	// Name of the endpoint group.
	EndpointGroupName *string `json:"endpointGroupName"`
	// Initial list of endpoints for this group.
	Endpoints *[]IEndpoint `json:"endpoints"`
	// The time between health checks for each endpoint.
	//
	// Must be either 10 or 30 seconds.
	HealthCheckInterval awscdk.Duration `json:"healthCheckInterval"`
	// The ping path for health checks (if the protocol is HTTP(S)).
	HealthCheckPath *string `json:"healthCheckPath"`
	// The port used to perform health checks.
	HealthCheckPort *float64 `json:"healthCheckPort"`
	// The protocol used to perform health checks.
	HealthCheckProtocol HealthCheckProtocol `json:"healthCheckProtocol"`
	// The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy.
	HealthCheckThreshold *float64 `json:"healthCheckThreshold"`
	// Override the destination ports used to route traffic to an endpoint.
	//
	// Unless overridden, the port used to hit the endpoint will be the same as the port
	// that traffic arrives on at the listener.
	PortOverrides *[]*PortOverride `json:"portOverrides"`
	// The AWS Region where the endpoint group is located.
	Region *string `json:"region"`
	// The percentage of traffic to send to this AWS Region.
	//
	// The percentage is applied to the traffic that would otherwise have been
	// routed to the Region based on optimal routing. Additional traffic is
	// distributed to other endpoint groups for this listener.
	TrafficDialPercentage *float64 `json:"trafficDialPercentage"`
}

// Property of the EndpointGroup.
//
// TODO: EXAMPLE
//
type EndpointGroupProps struct {
	// Name of the endpoint group.
	EndpointGroupName *string `json:"endpointGroupName"`
	// Initial list of endpoints for this group.
	Endpoints *[]IEndpoint `json:"endpoints"`
	// The time between health checks for each endpoint.
	//
	// Must be either 10 or 30 seconds.
	HealthCheckInterval awscdk.Duration `json:"healthCheckInterval"`
	// The ping path for health checks (if the protocol is HTTP(S)).
	HealthCheckPath *string `json:"healthCheckPath"`
	// The port used to perform health checks.
	HealthCheckPort *float64 `json:"healthCheckPort"`
	// The protocol used to perform health checks.
	HealthCheckProtocol HealthCheckProtocol `json:"healthCheckProtocol"`
	// The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy.
	HealthCheckThreshold *float64 `json:"healthCheckThreshold"`
	// Override the destination ports used to route traffic to an endpoint.
	//
	// Unless overridden, the port used to hit the endpoint will be the same as the port
	// that traffic arrives on at the listener.
	PortOverrides *[]*PortOverride `json:"portOverrides"`
	// The AWS Region where the endpoint group is located.
	Region *string `json:"region"`
	// The percentage of traffic to send to this AWS Region.
	//
	// The percentage is applied to the traffic that would otherwise have been
	// routed to the Region based on optimal routing. Additional traffic is
	// distributed to other endpoint groups for this listener.
	TrafficDialPercentage *float64 `json:"trafficDialPercentage"`
	// The Amazon Resource Name (ARN) of the listener.
	Listener IListener `json:"listener"`
}

// The protocol for the connections from clients to the accelerator.
type HealthCheckProtocol string

const (
	HealthCheckProtocol_TCP HealthCheckProtocol = "TCP"
	HealthCheckProtocol_HTTP HealthCheckProtocol = "HTTP"
	HealthCheckProtocol_HTTPS HealthCheckProtocol = "HTTPS"
)

// The interface of the Accelerator.
type IAccelerator interface {
	awscdk.IResource
	// The ARN of the accelerator.
	AcceleratorArn() *string
	// The Domain Name System (DNS) name that Global Accelerator creates that points to your accelerator's static IP addresses.
	DnsName() *string
}

// The jsii proxy for IAccelerator
type jsiiProxy_IAccelerator struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IAccelerator) AcceleratorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceleratorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccelerator) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

// An endpoint for the endpoint group.
//
// Implementations of `IEndpoint` can be found in the `aws-globalaccelerator-endpoints` package.
type IEndpoint interface {
	// Render the endpoint to an endpoint configuration.
	RenderEndpointConfiguration() interface{}
	// The region where the endpoint is located.
	//
	// If the region cannot be determined, `undefined` is returned
	Region() *string
}

// The jsii proxy for IEndpoint
type jsiiProxy_IEndpoint struct {
	_ byte // padding
}

func (i *jsiiProxy_IEndpoint) RenderEndpointConfiguration() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"renderEndpointConfiguration",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IEndpoint) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

// The interface of the EndpointGroup.
type IEndpointGroup interface {
	awscdk.IResource
	// EndpointGroup ARN.
	EndpointGroupArn() *string
}

// The jsii proxy for IEndpointGroup
type jsiiProxy_IEndpointGroup struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IEndpointGroup) EndpointGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointGroupArn",
		&returns,
	)
	return returns
}

// Interface of the Listener.
type IListener interface {
	awscdk.IResource
	// The ARN of the listener.
	ListenerArn() *string
}

// The jsii proxy for IListener
type jsiiProxy_IListener struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IListener) ListenerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerArn",
		&returns,
	)
	return returns
}

// The construct for the Listener.
//
// TODO: EXAMPLE
//
type Listener interface {
	awscdk.Resource
	IListener
	Env() *awscdk.ResourceEnvironment
	ListenerArn() *string
	ListenerName() *string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	AddEndpointGroup(id *string, options *EndpointGroupOptions) EndpointGroup
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for Listener
type jsiiProxy_Listener struct {
	internal.Type__awscdkResource
	jsiiProxy_IListener
}

func (j *jsiiProxy_Listener) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Listener) ListenerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Listener) ListenerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Listener) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Listener) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Listener) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewListener(scope constructs.Construct, id *string, props *ListenerProps) Listener {
	_init_.Initialize()

	j := jsiiProxy_Listener{}

	_jsii_.Create(
		"aws-cdk-lib.aws_globalaccelerator.Listener",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewListener_Override(l Listener, scope constructs.Construct, id *string, props *ListenerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_globalaccelerator.Listener",
		[]interface{}{scope, id, props},
		l,
	)
}

// import from ARN.
func Listener_FromListenerArn(scope constructs.Construct, id *string, listenerArn *string) IListener {
	_init_.Initialize()

	var returns IListener

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_globalaccelerator.Listener",
		"fromListenerArn",
		[]interface{}{scope, id, listenerArn},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Listener_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_globalaccelerator.Listener",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Listener_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_globalaccelerator.Listener",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add a new endpoint group to this listener.
func (l *jsiiProxy_Listener) AddEndpointGroup(id *string, options *EndpointGroupOptions) EndpointGroup {
	var returns EndpointGroup

	_jsii_.Invoke(
		l,
		"addEndpointGroup",
		[]interface{}{id, options},
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
func (l *jsiiProxy_Listener) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		l,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (l *jsiiProxy_Listener) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		l,
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
func (l *jsiiProxy_Listener) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		l,
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
func (l *jsiiProxy_Listener) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (l *jsiiProxy_Listener) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construct options for Listener.
//
// TODO: EXAMPLE
//
type ListenerOptions struct {
	// The list of port ranges for the connections from clients to the accelerator.
	PortRanges *[]*PortRange `json:"portRanges"`
	// Client affinity to direct all requests from a user to the same endpoint.
	//
	// If you have stateful applications, client affinity lets you direct all
	// requests from a user to the same endpoint.
	//
	// By default, each connection from each client is routed to seperate
	// endpoints. Set client affinity to SOURCE_IP to route all connections from
	// a single client to the same endpoint.
	ClientAffinity ClientAffinity `json:"clientAffinity"`
	// Name of the listener.
	ListenerName *string `json:"listenerName"`
	// The protocol for the connections from clients to the accelerator.
	Protocol ConnectionProtocol `json:"protocol"`
}

// Construct properties for Listener.
//
// TODO: EXAMPLE
//
type ListenerProps struct {
	// The list of port ranges for the connections from clients to the accelerator.
	PortRanges *[]*PortRange `json:"portRanges"`
	// Client affinity to direct all requests from a user to the same endpoint.
	//
	// If you have stateful applications, client affinity lets you direct all
	// requests from a user to the same endpoint.
	//
	// By default, each connection from each client is routed to seperate
	// endpoints. Set client affinity to SOURCE_IP to route all connections from
	// a single client to the same endpoint.
	ClientAffinity ClientAffinity `json:"clientAffinity"`
	// Name of the listener.
	ListenerName *string `json:"listenerName"`
	// The protocol for the connections from clients to the accelerator.
	Protocol ConnectionProtocol `json:"protocol"`
	// The accelerator for this listener.
	Accelerator IAccelerator `json:"accelerator"`
}

// Override specific listener ports used to route traffic to endpoints that are part of an endpoint group.
//
// TODO: EXAMPLE
//
type PortOverride struct {
	// The endpoint port that you want a listener port to be mapped to.
	//
	// This is the port on the endpoint, such as the Application Load Balancer or Amazon EC2 instance.
	EndpointPort *float64 `json:"endpointPort"`
	// The listener port that you want to map to a specific endpoint port.
	//
	// This is the port that user traffic arrives to the Global Accelerator on.
	ListenerPort *float64 `json:"listenerPort"`
}

// The list of port ranges for the connections from clients to the accelerator.
//
// TODO: EXAMPLE
//
type PortRange struct {
	// The first port in the range of ports, inclusive.
	FromPort *float64 `json:"fromPort"`
	// The last port in the range of ports, inclusive.
	ToPort *float64 `json:"toPort"`
}

// Untyped endpoint implementation.
//
// Prefer using the classes in the `aws-globalaccelerator-endpoints` package instead,
// as they accept typed constructs. You can use this class if you want to use an
// endpoint type that does not have an appropriate class in that package yet.
//
// TODO: EXAMPLE
//
type RawEndpoint interface {
	IEndpoint
	Region() *string
	RenderEndpointConfiguration() interface{}
}

// The jsii proxy struct for RawEndpoint
type jsiiProxy_RawEndpoint struct {
	jsiiProxy_IEndpoint
}

func (j *jsiiProxy_RawEndpoint) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}


func NewRawEndpoint(props *RawEndpointProps) RawEndpoint {
	_init_.Initialize()

	j := jsiiProxy_RawEndpoint{}

	_jsii_.Create(
		"aws-cdk-lib.aws_globalaccelerator.RawEndpoint",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewRawEndpoint_Override(r RawEndpoint, props *RawEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_globalaccelerator.RawEndpoint",
		[]interface{}{props},
		r,
	)
}

// Render the endpoint to an endpoint configuration.
func (r *jsiiProxy_RawEndpoint) RenderEndpointConfiguration() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"renderEndpointConfiguration",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for RawEndpoint.
//
// TODO: EXAMPLE
//
type RawEndpointProps struct {
	// Identifier of the endpoint.
	//
	// Load balancer ARN, instance ID or EIP allocation ID.
	EndpointId *string `json:"endpointId"`
	// Forward the client IP address.
	//
	// GlobalAccelerator will create Network Interfaces in your VPC in order
	// to preserve the client IP address.
	//
	// Only applies to Application Load Balancers and EC2 instances.
	//
	// Client IP address preservation is supported only in specific AWS Regions.
	// See the GlobalAccelerator Developer Guide for a list.
	PreserveClientIp *bool `json:"preserveClientIp"`
	// The region where this endpoint is located.
	Region *string `json:"region"`
	// Endpoint weight across all endpoints in the group.
	//
	// Must be a value between 0 and 255.
	Weight *float64 `json:"weight"`
}

