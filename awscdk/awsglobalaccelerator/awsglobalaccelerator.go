package awsglobalaccelerator

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsglobalaccelerator/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// The Accelerator construct.
//
// TODO: EXAMPLE
//
// Experimental.
type Accelerator interface {
	awscdk.Resource
	IAccelerator
	AcceleratorArn() *string
	DnsName() *string
	Env() *awscdk.ResourceEnvironment
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	AddListener(id *string, options *ListenerOptions) Listener
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_Accelerator) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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


// Experimental.
func NewAccelerator(scope constructs.Construct, id *string, props *AcceleratorProps) Accelerator {
	_init_.Initialize()

	j := jsiiProxy_Accelerator{}

	_jsii_.Create(
		"monocdk.aws_globalaccelerator.Accelerator",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAccelerator_Override(a Accelerator, scope constructs.Construct, id *string, props *AcceleratorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_globalaccelerator.Accelerator",
		[]interface{}{scope, id, props},
		a,
	)
}

// import from attributes.
// Experimental.
func Accelerator_FromAcceleratorAttributes(scope constructs.Construct, id *string, attrs *AcceleratorAttributes) IAccelerator {
	_init_.Initialize()

	var returns IAccelerator

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.Accelerator",
		"fromAcceleratorAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func Accelerator_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.Accelerator",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Accelerator_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.Accelerator",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add a listener to the accelerator.
// Experimental.
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
// Experimental.
func (a *jsiiProxy_Accelerator) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
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
// Experimental.
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
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (a *jsiiProxy_Accelerator) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (a *jsiiProxy_Accelerator) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (a *jsiiProxy_Accelerator) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (a *jsiiProxy_Accelerator) Prepare() {
	_jsii_.InvokeVoid(
		a,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (a *jsiiProxy_Accelerator) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (a *jsiiProxy_Accelerator) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Attributes required to import an existing accelerator to the stack.
// Experimental.
type AcceleratorAttributes struct {
	// The ARN of the accelerator.
	// Experimental.
	AcceleratorArn *string `json:"acceleratorArn"`
	// The DNS name of the accelerator.
	// Experimental.
	DnsName *string `json:"dnsName"`
}

// Construct properties of the Accelerator.
// Experimental.
type AcceleratorProps struct {
	// The name of the accelerator.
	// Experimental.
	AcceleratorName *string `json:"acceleratorName"`
	// Indicates whether the accelerator is enabled.
	// Experimental.
	Enabled *bool `json:"enabled"`
}

// A CloudFormation `AWS::GlobalAccelerator::Accelerator`.
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
	Node() awscdk.ConstructNode
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_CfnAccelerator) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnAccelerator(scope awscdk.Construct, id *string, props *CfnAcceleratorProps) CfnAccelerator {
	_init_.Initialize()

	j := jsiiProxy_CfnAccelerator{}

	_jsii_.Create(
		"monocdk.aws_globalaccelerator.CfnAccelerator",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GlobalAccelerator::Accelerator`.
func NewCfnAccelerator_Override(c CfnAccelerator, scope awscdk.Construct, id *string, props *CfnAcceleratorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_globalaccelerator.CfnAccelerator",
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
// Experimental.
func CfnAccelerator_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.CfnAccelerator",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnAccelerator_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.CfnAccelerator",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnAccelerator_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.CfnAccelerator",
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
		"monocdk.aws_globalaccelerator.CfnAccelerator",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnAccelerator) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnAccelerator) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnAccelerator) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnAccelerator) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnAccelerator) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnAccelerator) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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
// Experimental.
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnAccelerator) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnAccelerator) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnAccelerator) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::GlobalAccelerator::Accelerator`.
type CfnAcceleratorProps struct {
	// `AWS::GlobalAccelerator::Accelerator.Name`.
	Name *string `json:"name"`
	// `AWS::GlobalAccelerator::Accelerator.Enabled`.
	Enabled interface{} `json:"enabled"`
	// `AWS::GlobalAccelerator::Accelerator.IpAddresses`.
	IpAddresses *[]*string `json:"ipAddresses"`
	// `AWS::GlobalAccelerator::Accelerator.IpAddressType`.
	IpAddressType *string `json:"ipAddressType"`
	// `AWS::GlobalAccelerator::Accelerator.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::GlobalAccelerator::EndpointGroup`.
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
	Node() awscdk.ConstructNode
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_CfnEndpointGroup) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnEndpointGroup(scope awscdk.Construct, id *string, props *CfnEndpointGroupProps) CfnEndpointGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnEndpointGroup{}

	_jsii_.Create(
		"monocdk.aws_globalaccelerator.CfnEndpointGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GlobalAccelerator::EndpointGroup`.
func NewCfnEndpointGroup_Override(c CfnEndpointGroup, scope awscdk.Construct, id *string, props *CfnEndpointGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_globalaccelerator.CfnEndpointGroup",
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
// Experimental.
func CfnEndpointGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.CfnEndpointGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnEndpointGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.CfnEndpointGroup",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnEndpointGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.CfnEndpointGroup",
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
		"monocdk.aws_globalaccelerator.CfnEndpointGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnEndpointGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnEndpointGroup) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnEndpointGroup) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnEndpointGroup) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnEndpointGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnEndpointGroup) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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
// Experimental.
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnEndpointGroup) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnEndpointGroup) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnEndpointGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnEndpointGroup_EndpointConfigurationProperty struct {
	// `CfnEndpointGroup.EndpointConfigurationProperty.EndpointId`.
	EndpointId *string `json:"endpointId"`
	// `CfnEndpointGroup.EndpointConfigurationProperty.ClientIPPreservationEnabled`.
	ClientIpPreservationEnabled interface{} `json:"clientIpPreservationEnabled"`
	// `CfnEndpointGroup.EndpointConfigurationProperty.Weight`.
	Weight *float64 `json:"weight"`
}

type CfnEndpointGroup_PortOverrideProperty struct {
	// `CfnEndpointGroup.PortOverrideProperty.EndpointPort`.
	EndpointPort *float64 `json:"endpointPort"`
	// `CfnEndpointGroup.PortOverrideProperty.ListenerPort`.
	ListenerPort *float64 `json:"listenerPort"`
}

// Properties for defining a `AWS::GlobalAccelerator::EndpointGroup`.
type CfnEndpointGroupProps struct {
	// `AWS::GlobalAccelerator::EndpointGroup.EndpointGroupRegion`.
	EndpointGroupRegion *string `json:"endpointGroupRegion"`
	// `AWS::GlobalAccelerator::EndpointGroup.ListenerArn`.
	ListenerArn *string `json:"listenerArn"`
	// `AWS::GlobalAccelerator::EndpointGroup.EndpointConfigurations`.
	EndpointConfigurations interface{} `json:"endpointConfigurations"`
	// `AWS::GlobalAccelerator::EndpointGroup.HealthCheckIntervalSeconds`.
	HealthCheckIntervalSeconds *float64 `json:"healthCheckIntervalSeconds"`
	// `AWS::GlobalAccelerator::EndpointGroup.HealthCheckPath`.
	HealthCheckPath *string `json:"healthCheckPath"`
	// `AWS::GlobalAccelerator::EndpointGroup.HealthCheckPort`.
	HealthCheckPort *float64 `json:"healthCheckPort"`
	// `AWS::GlobalAccelerator::EndpointGroup.HealthCheckProtocol`.
	HealthCheckProtocol *string `json:"healthCheckProtocol"`
	// `AWS::GlobalAccelerator::EndpointGroup.PortOverrides`.
	PortOverrides interface{} `json:"portOverrides"`
	// `AWS::GlobalAccelerator::EndpointGroup.ThresholdCount`.
	ThresholdCount *float64 `json:"thresholdCount"`
	// `AWS::GlobalAccelerator::EndpointGroup.TrafficDialPercentage`.
	TrafficDialPercentage *float64 `json:"trafficDialPercentage"`
}

// A CloudFormation `AWS::GlobalAccelerator::Listener`.
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
	Node() awscdk.ConstructNode
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_CfnListener) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnListener(scope awscdk.Construct, id *string, props *CfnListenerProps) CfnListener {
	_init_.Initialize()

	j := jsiiProxy_CfnListener{}

	_jsii_.Create(
		"monocdk.aws_globalaccelerator.CfnListener",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GlobalAccelerator::Listener`.
func NewCfnListener_Override(c CfnListener, scope awscdk.Construct, id *string, props *CfnListenerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_globalaccelerator.CfnListener",
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
// Experimental.
func CfnListener_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.CfnListener",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnListener_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.CfnListener",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnListener_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.CfnListener",
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
		"monocdk.aws_globalaccelerator.CfnListener",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnListener) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnListener) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnListener) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnListener) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnListener) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnListener) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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
// Experimental.
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnListener) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnListener) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnListener) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnListener_PortRangeProperty struct {
	// `CfnListener.PortRangeProperty.FromPort`.
	FromPort *float64 `json:"fromPort"`
	// `CfnListener.PortRangeProperty.ToPort`.
	ToPort *float64 `json:"toPort"`
}

// Properties for defining a `AWS::GlobalAccelerator::Listener`.
type CfnListenerProps struct {
	// `AWS::GlobalAccelerator::Listener.AcceleratorArn`.
	AcceleratorArn *string `json:"acceleratorArn"`
	// `AWS::GlobalAccelerator::Listener.PortRanges`.
	PortRanges interface{} `json:"portRanges"`
	// `AWS::GlobalAccelerator::Listener.Protocol`.
	Protocol *string `json:"protocol"`
	// `AWS::GlobalAccelerator::Listener.ClientAffinity`.
	ClientAffinity *string `json:"clientAffinity"`
}

// Client affinity gives you control over whether to always route each client to the same specific endpoint.
// See: https://docs.aws.amazon.com/global-accelerator/latest/dg/about-listeners.html#about-listeners-client-affinity
//
// Experimental.
type ClientAffinity string

const (
	ClientAffinity_NONE ClientAffinity = "NONE"
	ClientAffinity_SOURCE_IP ClientAffinity = "SOURCE_IP"
)

// The protocol for the connections from clients to the accelerator.
// Experimental.
type ConnectionProtocol string

const (
	ConnectionProtocol_TCP ConnectionProtocol = "TCP"
	ConnectionProtocol_UDP ConnectionProtocol = "UDP"
)

// EndpointGroup construct.
// Experimental.
type EndpointGroup interface {
	awscdk.Resource
	IEndpointGroup
	EndpointGroupArn() *string
	EndpointGroupName() *string
	Endpoints() *[]IEndpoint
	Env() *awscdk.ResourceEnvironment
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	AddEndpoint(endpoint IEndpoint)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	ConnectionsPeer(id *string, vpc awsec2.IVpc) awsec2.IPeer
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_EndpointGroup) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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


// Experimental.
func NewEndpointGroup(scope constructs.Construct, id *string, props *EndpointGroupProps) EndpointGroup {
	_init_.Initialize()

	j := jsiiProxy_EndpointGroup{}

	_jsii_.Create(
		"monocdk.aws_globalaccelerator.EndpointGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEndpointGroup_Override(e EndpointGroup, scope constructs.Construct, id *string, props *EndpointGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_globalaccelerator.EndpointGroup",
		[]interface{}{scope, id, props},
		e,
	)
}

// import from ARN.
// Experimental.
func EndpointGroup_FromEndpointGroupArn(scope constructs.Construct, id *string, endpointGroupArn *string) IEndpointGroup {
	_init_.Initialize()

	var returns IEndpointGroup

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.EndpointGroup",
		"fromEndpointGroupArn",
		[]interface{}{scope, id, endpointGroupArn},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func EndpointGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.EndpointGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func EndpointGroup_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.EndpointGroup",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add an endpoint.
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
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
// Experimental.
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
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EndpointGroup) OnPrepare() {
	_jsii_.InvokeVoid(
		e,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EndpointGroup) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (e *jsiiProxy_EndpointGroup) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EndpointGroup) Prepare() {
	_jsii_.InvokeVoid(
		e,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EndpointGroup) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (e *jsiiProxy_EndpointGroup) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Basic options for creating a new EndpointGroup.
//
// TODO: EXAMPLE
//
// Experimental.
type EndpointGroupOptions struct {
	// Name of the endpoint group.
	// Experimental.
	EndpointGroupName *string `json:"endpointGroupName"`
	// Initial list of endpoints for this group.
	// Experimental.
	Endpoints *[]IEndpoint `json:"endpoints"`
	// The time between health checks for each endpoint.
	//
	// Must be either 10 or 30 seconds.
	// Experimental.
	HealthCheckInterval awscdk.Duration `json:"healthCheckInterval"`
	// The ping path for health checks (if the protocol is HTTP(S)).
	// Experimental.
	HealthCheckPath *string `json:"healthCheckPath"`
	// The port used to perform health checks.
	// Experimental.
	HealthCheckPort *float64 `json:"healthCheckPort"`
	// The protocol used to perform health checks.
	// Experimental.
	HealthCheckProtocol HealthCheckProtocol `json:"healthCheckProtocol"`
	// The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy.
	// Experimental.
	HealthCheckThreshold *float64 `json:"healthCheckThreshold"`
	// Override the destination ports used to route traffic to an endpoint.
	//
	// Unless overridden, the port used to hit the endpoint will be the same as the port
	// that traffic arrives on at the listener.
	// Experimental.
	PortOverrides *[]*PortOverride `json:"portOverrides"`
	// The AWS Region where the endpoint group is located.
	// Experimental.
	Region *string `json:"region"`
	// The percentage of traffic to send to this AWS Region.
	//
	// The percentage is applied to the traffic that would otherwise have been
	// routed to the Region based on optimal routing. Additional traffic is
	// distributed to other endpoint groups for this listener.
	// Experimental.
	TrafficDialPercentage *float64 `json:"trafficDialPercentage"`
}

// Property of the EndpointGroup.
// Experimental.
type EndpointGroupProps struct {
	// Name of the endpoint group.
	// Experimental.
	EndpointGroupName *string `json:"endpointGroupName"`
	// Initial list of endpoints for this group.
	// Experimental.
	Endpoints *[]IEndpoint `json:"endpoints"`
	// The time between health checks for each endpoint.
	//
	// Must be either 10 or 30 seconds.
	// Experimental.
	HealthCheckInterval awscdk.Duration `json:"healthCheckInterval"`
	// The ping path for health checks (if the protocol is HTTP(S)).
	// Experimental.
	HealthCheckPath *string `json:"healthCheckPath"`
	// The port used to perform health checks.
	// Experimental.
	HealthCheckPort *float64 `json:"healthCheckPort"`
	// The protocol used to perform health checks.
	// Experimental.
	HealthCheckProtocol HealthCheckProtocol `json:"healthCheckProtocol"`
	// The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy.
	// Experimental.
	HealthCheckThreshold *float64 `json:"healthCheckThreshold"`
	// Override the destination ports used to route traffic to an endpoint.
	//
	// Unless overridden, the port used to hit the endpoint will be the same as the port
	// that traffic arrives on at the listener.
	// Experimental.
	PortOverrides *[]*PortOverride `json:"portOverrides"`
	// The AWS Region where the endpoint group is located.
	// Experimental.
	Region *string `json:"region"`
	// The percentage of traffic to send to this AWS Region.
	//
	// The percentage is applied to the traffic that would otherwise have been
	// routed to the Region based on optimal routing. Additional traffic is
	// distributed to other endpoint groups for this listener.
	// Experimental.
	TrafficDialPercentage *float64 `json:"trafficDialPercentage"`
	// The Amazon Resource Name (ARN) of the listener.
	// Experimental.
	Listener IListener `json:"listener"`
}

// The protocol for the connections from clients to the accelerator.
// Experimental.
type HealthCheckProtocol string

const (
	HealthCheckProtocol_TCP HealthCheckProtocol = "TCP"
	HealthCheckProtocol_HTTP HealthCheckProtocol = "HTTP"
	HealthCheckProtocol_HTTPS HealthCheckProtocol = "HTTPS"
)

// The interface of the Accelerator.
// Experimental.
type IAccelerator interface {
	awscdk.IResource
	// The ARN of the accelerator.
	// Experimental.
	AcceleratorArn() *string
	// The Domain Name System (DNS) name that Global Accelerator creates that points to your accelerator's static IP addresses.
	// Experimental.
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
// Experimental.
type IEndpoint interface {
	// Render the endpoint to an endpoint configuration.
	// Experimental.
	RenderEndpointConfiguration() interface{}
	// The region where the endpoint is located.
	//
	// If the region cannot be determined, `undefined` is returned
	// Experimental.
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
// Experimental.
type IEndpointGroup interface {
	awscdk.IResource
	// EndpointGroup ARN.
	// Experimental.
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
// Experimental.
type IListener interface {
	awscdk.IResource
	// The ARN of the listener.
	// Experimental.
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
// Experimental.
type Listener interface {
	awscdk.Resource
	IListener
	Env() *awscdk.ResourceEnvironment
	ListenerArn() *string
	ListenerName() *string
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	AddEndpointGroup(id *string, options *EndpointGroupOptions) EndpointGroup
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_Listener) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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


// Experimental.
func NewListener(scope constructs.Construct, id *string, props *ListenerProps) Listener {
	_init_.Initialize()

	j := jsiiProxy_Listener{}

	_jsii_.Create(
		"monocdk.aws_globalaccelerator.Listener",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewListener_Override(l Listener, scope constructs.Construct, id *string, props *ListenerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_globalaccelerator.Listener",
		[]interface{}{scope, id, props},
		l,
	)
}

// import from ARN.
// Experimental.
func Listener_FromListenerArn(scope constructs.Construct, id *string, listenerArn *string) IListener {
	_init_.Initialize()

	var returns IListener

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.Listener",
		"fromListenerArn",
		[]interface{}{scope, id, listenerArn},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func Listener_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.Listener",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Listener_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.Listener",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add a new endpoint group to this listener.
// Experimental.
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
// Experimental.
func (l *jsiiProxy_Listener) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		l,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
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
// Experimental.
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
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (l *jsiiProxy_Listener) OnPrepare() {
	_jsii_.InvokeVoid(
		l,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (l *jsiiProxy_Listener) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (l *jsiiProxy_Listener) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (l *jsiiProxy_Listener) Prepare() {
	_jsii_.InvokeVoid(
		l,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (l *jsiiProxy_Listener) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (l *jsiiProxy_Listener) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construct options for Listener.
//
// TODO: EXAMPLE
//
// Experimental.
type ListenerOptions struct {
	// The list of port ranges for the connections from clients to the accelerator.
	// Experimental.
	PortRanges *[]*PortRange `json:"portRanges"`
	// Client affinity to direct all requests from a user to the same endpoint.
	//
	// If you have stateful applications, client affinity lets you direct all
	// requests from a user to the same endpoint.
	//
	// By default, each connection from each client is routed to seperate
	// endpoints. Set client affinity to SOURCE_IP to route all connections from
	// a single client to the same endpoint.
	// Experimental.
	ClientAffinity ClientAffinity `json:"clientAffinity"`
	// Name of the listener.
	// Experimental.
	ListenerName *string `json:"listenerName"`
	// The protocol for the connections from clients to the accelerator.
	// Experimental.
	Protocol ConnectionProtocol `json:"protocol"`
}

// Construct properties for Listener.
// Experimental.
type ListenerProps struct {
	// The list of port ranges for the connections from clients to the accelerator.
	// Experimental.
	PortRanges *[]*PortRange `json:"portRanges"`
	// Client affinity to direct all requests from a user to the same endpoint.
	//
	// If you have stateful applications, client affinity lets you direct all
	// requests from a user to the same endpoint.
	//
	// By default, each connection from each client is routed to seperate
	// endpoints. Set client affinity to SOURCE_IP to route all connections from
	// a single client to the same endpoint.
	// Experimental.
	ClientAffinity ClientAffinity `json:"clientAffinity"`
	// Name of the listener.
	// Experimental.
	ListenerName *string `json:"listenerName"`
	// The protocol for the connections from clients to the accelerator.
	// Experimental.
	Protocol ConnectionProtocol `json:"protocol"`
	// The accelerator for this listener.
	// Experimental.
	Accelerator IAccelerator `json:"accelerator"`
}

// Override specific listener ports used to route traffic to endpoints that are part of an endpoint group.
// Experimental.
type PortOverride struct {
	// The endpoint port that you want a listener port to be mapped to.
	//
	// This is the port on the endpoint, such as the Application Load Balancer or Amazon EC2 instance.
	// Experimental.
	EndpointPort *float64 `json:"endpointPort"`
	// The listener port that you want to map to a specific endpoint port.
	//
	// This is the port that user traffic arrives to the Global Accelerator on.
	// Experimental.
	ListenerPort *float64 `json:"listenerPort"`
}

// The list of port ranges for the connections from clients to the accelerator.
// Experimental.
type PortRange struct {
	// The first port in the range of ports, inclusive.
	// Experimental.
	FromPort *float64 `json:"fromPort"`
	// The last port in the range of ports, inclusive.
	// Experimental.
	ToPort *float64 `json:"toPort"`
}

// Untyped endpoint implementation.
//
// Prefer using the classes in the `aws-globalaccelerator-endpoints` package instead,
// as they accept typed constructs. You can use this class if you want to use an
// endpoint type that does not have an appropriate class in that package yet.
// Experimental.
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


// Experimental.
func NewRawEndpoint(props *RawEndpointProps) RawEndpoint {
	_init_.Initialize()

	j := jsiiProxy_RawEndpoint{}

	_jsii_.Create(
		"monocdk.aws_globalaccelerator.RawEndpoint",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewRawEndpoint_Override(r RawEndpoint, props *RawEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_globalaccelerator.RawEndpoint",
		[]interface{}{props},
		r,
	)
}

// Render the endpoint to an endpoint configuration.
// Experimental.
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
// Experimental.
type RawEndpointProps struct {
	// Identifier of the endpoint.
	//
	// Load balancer ARN, instance ID or EIP allocation ID.
	// Experimental.
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
	// Experimental.
	PreserveClientIp *bool `json:"preserveClientIp"`
	// The region where this endpoint is located.
	// Experimental.
	Region *string `json:"region"`
	// Endpoint weight across all endpoints in the group.
	//
	// Must be a value between 0 and 255.
	// Experimental.
	Weight *float64 `json:"weight"`
}

