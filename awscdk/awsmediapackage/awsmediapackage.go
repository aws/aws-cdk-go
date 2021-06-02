package awsmediapackage

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsmediapackage/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::MediaPackage::Asset`.
type CfnAsset interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrCreatedAt() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	EgressEndpoints() interface{}
	SetEgressEndpoints(val interface{})
	Id() *string
	SetId(val *string)
	LogicalId() *string
	Node() awscdk.ConstructNode
	PackagingGroupId() *string
	SetPackagingGroupId(val *string)
	Ref() *string
	ResourceId() *string
	SetResourceId(val *string)
	SourceArn() *string
	SetSourceArn(val *string)
	SourceRoleArn() *string
	SetSourceRoleArn(val *string)
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

// The jsii proxy struct for CfnAsset
type jsiiProxy_CfnAsset struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAsset) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) AttrCreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) EgressEndpoints() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"egressEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) PackagingGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packagingGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) SourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) SourceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::MediaPackage::Asset`.
func NewCfnAsset(scope awscdk.Construct, id *string, props *CfnAssetProps) CfnAsset {
	_init_.Initialize()

	j := jsiiProxy_CfnAsset{}

	_jsii_.Create(
		"monocdk.aws_mediapackage.CfnAsset",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MediaPackage::Asset`.
func NewCfnAsset_Override(c CfnAsset, scope awscdk.Construct, id *string, props *CfnAssetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_mediapackage.CfnAsset",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAsset) SetEgressEndpoints(val interface{}) {
	_jsii_.Set(
		j,
		"egressEndpoints",
		val,
	)
}

func (j *jsiiProxy_CfnAsset) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CfnAsset) SetPackagingGroupId(val *string) {
	_jsii_.Set(
		j,
		"packagingGroupId",
		val,
	)
}

func (j *jsiiProxy_CfnAsset) SetResourceId(val *string) {
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_CfnAsset) SetSourceArn(val *string) {
	_jsii_.Set(
		j,
		"sourceArn",
		val,
	)
}

func (j *jsiiProxy_CfnAsset) SetSourceRoleArn(val *string) {
	_jsii_.Set(
		j,
		"sourceRoleArn",
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
func CfnAsset_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_mediapackage.CfnAsset",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnAsset_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_mediapackage.CfnAsset",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnAsset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_mediapackage.CfnAsset",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAsset_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_mediapackage.CfnAsset",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnAsset) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnAsset) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnAsset) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnAsset) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnAsset) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnAsset) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnAsset) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnAsset) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnAsset) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnAsset) Inspect(inspector awscdk.TreeInspector) {
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
func (c *jsiiProxy_CfnAsset) OnPrepare() {
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
func (c *jsiiProxy_CfnAsset) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnAsset) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnAsset) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnAsset) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnAsset) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnAsset) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnAsset) Synthesize(session awscdk.ISynthesisSession) {
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
func (c *jsiiProxy_CfnAsset) ToString() *string {
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
func (c *jsiiProxy_CfnAsset) Validate() *[]*string {
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
func (c *jsiiProxy_CfnAsset) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnAsset_EgressEndpointProperty struct {
	// `CfnAsset.EgressEndpointProperty.PackagingConfigurationId`.
	PackagingConfigurationId *string `json:"packagingConfigurationId"`
	// `CfnAsset.EgressEndpointProperty.Url`.
	Url *string `json:"url"`
}

// Properties for defining a `AWS::MediaPackage::Asset`.
type CfnAssetProps struct {
	// `AWS::MediaPackage::Asset.Id`.
	Id *string `json:"id"`
	// `AWS::MediaPackage::Asset.PackagingGroupId`.
	PackagingGroupId *string `json:"packagingGroupId"`
	// `AWS::MediaPackage::Asset.SourceArn`.
	SourceArn *string `json:"sourceArn"`
	// `AWS::MediaPackage::Asset.SourceRoleArn`.
	SourceRoleArn *string `json:"sourceRoleArn"`
	// `AWS::MediaPackage::Asset.EgressEndpoints`.
	EgressEndpoints interface{} `json:"egressEndpoints"`
	// `AWS::MediaPackage::Asset.ResourceId`.
	ResourceId *string `json:"resourceId"`
	// `AWS::MediaPackage::Asset.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::MediaPackage::Channel`.
type CfnChannel interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	EgressAccessLogs() interface{}
	SetEgressAccessLogs(val interface{})
	Id() *string
	SetId(val *string)
	IngressAccessLogs() interface{}
	SetIngressAccessLogs(val interface{})
	LogicalId() *string
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

// The jsii proxy struct for CfnChannel
type jsiiProxy_CfnChannel struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnChannel) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) EgressAccessLogs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"egressAccessLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) IngressAccessLogs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ingressAccessLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::MediaPackage::Channel`.
func NewCfnChannel(scope awscdk.Construct, id *string, props *CfnChannelProps) CfnChannel {
	_init_.Initialize()

	j := jsiiProxy_CfnChannel{}

	_jsii_.Create(
		"monocdk.aws_mediapackage.CfnChannel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MediaPackage::Channel`.
func NewCfnChannel_Override(c CfnChannel, scope awscdk.Construct, id *string, props *CfnChannelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_mediapackage.CfnChannel",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnChannel) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnChannel) SetEgressAccessLogs(val interface{}) {
	_jsii_.Set(
		j,
		"egressAccessLogs",
		val,
	)
}

func (j *jsiiProxy_CfnChannel) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CfnChannel) SetIngressAccessLogs(val interface{}) {
	_jsii_.Set(
		j,
		"ingressAccessLogs",
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
func CfnChannel_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_mediapackage.CfnChannel",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnChannel_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_mediapackage.CfnChannel",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_mediapackage.CfnChannel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnChannel_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_mediapackage.CfnChannel",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnChannel) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnChannel) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnChannel) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnChannel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnChannel) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnChannel) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnChannel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnChannel) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnChannel) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnChannel) Inspect(inspector awscdk.TreeInspector) {
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
func (c *jsiiProxy_CfnChannel) OnPrepare() {
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
func (c *jsiiProxy_CfnChannel) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnChannel) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnChannel) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnChannel) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnChannel) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnChannel) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnChannel) Synthesize(session awscdk.ISynthesisSession) {
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
func (c *jsiiProxy_CfnChannel) ToString() *string {
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
func (c *jsiiProxy_CfnChannel) Validate() *[]*string {
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
func (c *jsiiProxy_CfnChannel) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnChannel_LogConfigurationProperty struct {
	// `CfnChannel.LogConfigurationProperty.LogGroupName`.
	LogGroupName *string `json:"logGroupName"`
}

// Properties for defining a `AWS::MediaPackage::Channel`.
type CfnChannelProps struct {
	// `AWS::MediaPackage::Channel.Id`.
	Id *string `json:"id"`
	// `AWS::MediaPackage::Channel.Description`.
	Description *string `json:"description"`
	// `AWS::MediaPackage::Channel.EgressAccessLogs`.
	EgressAccessLogs interface{} `json:"egressAccessLogs"`
	// `AWS::MediaPackage::Channel.IngressAccessLogs`.
	IngressAccessLogs interface{} `json:"ingressAccessLogs"`
	// `AWS::MediaPackage::Channel.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::MediaPackage::OriginEndpoint`.
type CfnOriginEndpoint interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrUrl() *string
	Authorization() interface{}
	SetAuthorization(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ChannelId() *string
	SetChannelId(val *string)
	CmafPackage() interface{}
	SetCmafPackage(val interface{})
	CreationStack() *[]*string
	DashPackage() interface{}
	SetDashPackage(val interface{})
	Description() *string
	SetDescription(val *string)
	HlsPackage() interface{}
	SetHlsPackage(val interface{})
	Id() *string
	SetId(val *string)
	LogicalId() *string
	ManifestName() *string
	SetManifestName(val *string)
	MssPackage() interface{}
	SetMssPackage(val interface{})
	Node() awscdk.ConstructNode
	Origination() *string
	SetOrigination(val *string)
	Ref() *string
	Stack() awscdk.Stack
	StartoverWindowSeconds() *float64
	SetStartoverWindowSeconds(val *float64)
	Tags() awscdk.TagManager
	TimeDelaySeconds() *float64
	SetTimeDelaySeconds(val *float64)
	UpdatedProperites() *map[string]interface{}
	Whitelist() *[]*string
	SetWhitelist(val *[]*string)
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

// The jsii proxy struct for CfnOriginEndpoint
type jsiiProxy_CfnOriginEndpoint struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnOriginEndpoint) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) AttrUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) Authorization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) ChannelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) CmafPackage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cmafPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) DashPackage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dashPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) HlsPackage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hlsPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) ManifestName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) MssPackage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mssPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) Origination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"origination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) StartoverWindowSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startoverWindowSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) TimeDelaySeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeDelaySeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) Whitelist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"whitelist",
		&returns,
	)
	return returns
}


// Create a new `AWS::MediaPackage::OriginEndpoint`.
func NewCfnOriginEndpoint(scope awscdk.Construct, id *string, props *CfnOriginEndpointProps) CfnOriginEndpoint {
	_init_.Initialize()

	j := jsiiProxy_CfnOriginEndpoint{}

	_jsii_.Create(
		"monocdk.aws_mediapackage.CfnOriginEndpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MediaPackage::OriginEndpoint`.
func NewCfnOriginEndpoint_Override(c CfnOriginEndpoint, scope awscdk.Construct, id *string, props *CfnOriginEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_mediapackage.CfnOriginEndpoint",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint) SetAuthorization(val interface{}) {
	_jsii_.Set(
		j,
		"authorization",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint) SetChannelId(val *string) {
	_jsii_.Set(
		j,
		"channelId",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint) SetCmafPackage(val interface{}) {
	_jsii_.Set(
		j,
		"cmafPackage",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint) SetDashPackage(val interface{}) {
	_jsii_.Set(
		j,
		"dashPackage",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint) SetHlsPackage(val interface{}) {
	_jsii_.Set(
		j,
		"hlsPackage",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint) SetManifestName(val *string) {
	_jsii_.Set(
		j,
		"manifestName",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint) SetMssPackage(val interface{}) {
	_jsii_.Set(
		j,
		"mssPackage",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint) SetOrigination(val *string) {
	_jsii_.Set(
		j,
		"origination",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint) SetStartoverWindowSeconds(val *float64) {
	_jsii_.Set(
		j,
		"startoverWindowSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint) SetTimeDelaySeconds(val *float64) {
	_jsii_.Set(
		j,
		"timeDelaySeconds",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint) SetWhitelist(val *[]*string) {
	_jsii_.Set(
		j,
		"whitelist",
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
func CfnOriginEndpoint_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_mediapackage.CfnOriginEndpoint",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnOriginEndpoint_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_mediapackage.CfnOriginEndpoint",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnOriginEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_mediapackage.CfnOriginEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnOriginEndpoint_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_mediapackage.CfnOriginEndpoint",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnOriginEndpoint) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnOriginEndpoint) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnOriginEndpoint) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnOriginEndpoint) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnOriginEndpoint) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnOriginEndpoint) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnOriginEndpoint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnOriginEndpoint) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnOriginEndpoint) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnOriginEndpoint) Inspect(inspector awscdk.TreeInspector) {
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
func (c *jsiiProxy_CfnOriginEndpoint) OnPrepare() {
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
func (c *jsiiProxy_CfnOriginEndpoint) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnOriginEndpoint) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnOriginEndpoint) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnOriginEndpoint) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnOriginEndpoint) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnOriginEndpoint) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnOriginEndpoint) Synthesize(session awscdk.ISynthesisSession) {
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
func (c *jsiiProxy_CfnOriginEndpoint) ToString() *string {
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
func (c *jsiiProxy_CfnOriginEndpoint) Validate() *[]*string {
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
func (c *jsiiProxy_CfnOriginEndpoint) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnOriginEndpoint_AuthorizationProperty struct {
	// `CfnOriginEndpoint.AuthorizationProperty.CdnIdentifierSecret`.
	CdnIdentifierSecret *string `json:"cdnIdentifierSecret"`
	// `CfnOriginEndpoint.AuthorizationProperty.SecretsRoleArn`.
	SecretsRoleArn *string `json:"secretsRoleArn"`
}

type CfnOriginEndpoint_CmafEncryptionProperty struct {
	// `CfnOriginEndpoint.CmafEncryptionProperty.SpekeKeyProvider`.
	SpekeKeyProvider interface{} `json:"spekeKeyProvider"`
	// `CfnOriginEndpoint.CmafEncryptionProperty.ConstantInitializationVector`.
	ConstantInitializationVector *string `json:"constantInitializationVector"`
	// `CfnOriginEndpoint.CmafEncryptionProperty.KeyRotationIntervalSeconds`.
	KeyRotationIntervalSeconds *float64 `json:"keyRotationIntervalSeconds"`
}

type CfnOriginEndpoint_CmafPackageProperty struct {
	// `CfnOriginEndpoint.CmafPackageProperty.Encryption`.
	Encryption interface{} `json:"encryption"`
	// `CfnOriginEndpoint.CmafPackageProperty.HlsManifests`.
	HlsManifests interface{} `json:"hlsManifests"`
	// `CfnOriginEndpoint.CmafPackageProperty.SegmentDurationSeconds`.
	SegmentDurationSeconds *float64 `json:"segmentDurationSeconds"`
	// `CfnOriginEndpoint.CmafPackageProperty.SegmentPrefix`.
	SegmentPrefix *string `json:"segmentPrefix"`
	// `CfnOriginEndpoint.CmafPackageProperty.StreamSelection`.
	StreamSelection interface{} `json:"streamSelection"`
}

type CfnOriginEndpoint_DashEncryptionProperty struct {
	// `CfnOriginEndpoint.DashEncryptionProperty.SpekeKeyProvider`.
	SpekeKeyProvider interface{} `json:"spekeKeyProvider"`
	// `CfnOriginEndpoint.DashEncryptionProperty.KeyRotationIntervalSeconds`.
	KeyRotationIntervalSeconds *float64 `json:"keyRotationIntervalSeconds"`
}

type CfnOriginEndpoint_DashPackageProperty struct {
	// `CfnOriginEndpoint.DashPackageProperty.AdsOnDeliveryRestrictions`.
	AdsOnDeliveryRestrictions *string `json:"adsOnDeliveryRestrictions"`
	// `CfnOriginEndpoint.DashPackageProperty.AdTriggers`.
	AdTriggers *[]*string `json:"adTriggers"`
	// `CfnOriginEndpoint.DashPackageProperty.Encryption`.
	Encryption interface{} `json:"encryption"`
	// `CfnOriginEndpoint.DashPackageProperty.ManifestLayout`.
	ManifestLayout *string `json:"manifestLayout"`
	// `CfnOriginEndpoint.DashPackageProperty.ManifestWindowSeconds`.
	ManifestWindowSeconds *float64 `json:"manifestWindowSeconds"`
	// `CfnOriginEndpoint.DashPackageProperty.MinBufferTimeSeconds`.
	MinBufferTimeSeconds *float64 `json:"minBufferTimeSeconds"`
	// `CfnOriginEndpoint.DashPackageProperty.MinUpdatePeriodSeconds`.
	MinUpdatePeriodSeconds *float64 `json:"minUpdatePeriodSeconds"`
	// `CfnOriginEndpoint.DashPackageProperty.PeriodTriggers`.
	PeriodTriggers *[]*string `json:"periodTriggers"`
	// `CfnOriginEndpoint.DashPackageProperty.Profile`.
	Profile *string `json:"profile"`
	// `CfnOriginEndpoint.DashPackageProperty.SegmentDurationSeconds`.
	SegmentDurationSeconds *float64 `json:"segmentDurationSeconds"`
	// `CfnOriginEndpoint.DashPackageProperty.SegmentTemplateFormat`.
	SegmentTemplateFormat *string `json:"segmentTemplateFormat"`
	// `CfnOriginEndpoint.DashPackageProperty.StreamSelection`.
	StreamSelection interface{} `json:"streamSelection"`
	// `CfnOriginEndpoint.DashPackageProperty.SuggestedPresentationDelaySeconds`.
	SuggestedPresentationDelaySeconds *float64 `json:"suggestedPresentationDelaySeconds"`
	// `CfnOriginEndpoint.DashPackageProperty.UtcTiming`.
	UtcTiming *string `json:"utcTiming"`
	// `CfnOriginEndpoint.DashPackageProperty.UtcTimingUri`.
	UtcTimingUri *string `json:"utcTimingUri"`
}

type CfnOriginEndpoint_HlsEncryptionProperty struct {
	// `CfnOriginEndpoint.HlsEncryptionProperty.SpekeKeyProvider`.
	SpekeKeyProvider interface{} `json:"spekeKeyProvider"`
	// `CfnOriginEndpoint.HlsEncryptionProperty.ConstantInitializationVector`.
	ConstantInitializationVector *string `json:"constantInitializationVector"`
	// `CfnOriginEndpoint.HlsEncryptionProperty.EncryptionMethod`.
	EncryptionMethod *string `json:"encryptionMethod"`
	// `CfnOriginEndpoint.HlsEncryptionProperty.KeyRotationIntervalSeconds`.
	KeyRotationIntervalSeconds *float64 `json:"keyRotationIntervalSeconds"`
	// `CfnOriginEndpoint.HlsEncryptionProperty.RepeatExtXKey`.
	RepeatExtXKey interface{} `json:"repeatExtXKey"`
}

type CfnOriginEndpoint_HlsManifestProperty struct {
	// `CfnOriginEndpoint.HlsManifestProperty.Id`.
	Id *string `json:"id"`
	// `CfnOriginEndpoint.HlsManifestProperty.AdMarkers`.
	AdMarkers *string `json:"adMarkers"`
	// `CfnOriginEndpoint.HlsManifestProperty.AdsOnDeliveryRestrictions`.
	AdsOnDeliveryRestrictions *string `json:"adsOnDeliveryRestrictions"`
	// `CfnOriginEndpoint.HlsManifestProperty.AdTriggers`.
	AdTriggers *[]*string `json:"adTriggers"`
	// `CfnOriginEndpoint.HlsManifestProperty.IncludeIframeOnlyStream`.
	IncludeIframeOnlyStream interface{} `json:"includeIframeOnlyStream"`
	// `CfnOriginEndpoint.HlsManifestProperty.ManifestName`.
	ManifestName *string `json:"manifestName"`
	// `CfnOriginEndpoint.HlsManifestProperty.PlaylistType`.
	PlaylistType *string `json:"playlistType"`
	// `CfnOriginEndpoint.HlsManifestProperty.PlaylistWindowSeconds`.
	PlaylistWindowSeconds *float64 `json:"playlistWindowSeconds"`
	// `CfnOriginEndpoint.HlsManifestProperty.ProgramDateTimeIntervalSeconds`.
	ProgramDateTimeIntervalSeconds *float64 `json:"programDateTimeIntervalSeconds"`
	// `CfnOriginEndpoint.HlsManifestProperty.Url`.
	Url *string `json:"url"`
}

type CfnOriginEndpoint_HlsPackageProperty struct {
	// `CfnOriginEndpoint.HlsPackageProperty.AdMarkers`.
	AdMarkers *string `json:"adMarkers"`
	// `CfnOriginEndpoint.HlsPackageProperty.AdsOnDeliveryRestrictions`.
	AdsOnDeliveryRestrictions *string `json:"adsOnDeliveryRestrictions"`
	// `CfnOriginEndpoint.HlsPackageProperty.AdTriggers`.
	AdTriggers *[]*string `json:"adTriggers"`
	// `CfnOriginEndpoint.HlsPackageProperty.Encryption`.
	Encryption interface{} `json:"encryption"`
	// `CfnOriginEndpoint.HlsPackageProperty.IncludeIframeOnlyStream`.
	IncludeIframeOnlyStream interface{} `json:"includeIframeOnlyStream"`
	// `CfnOriginEndpoint.HlsPackageProperty.PlaylistType`.
	PlaylistType *string `json:"playlistType"`
	// `CfnOriginEndpoint.HlsPackageProperty.PlaylistWindowSeconds`.
	PlaylistWindowSeconds *float64 `json:"playlistWindowSeconds"`
	// `CfnOriginEndpoint.HlsPackageProperty.ProgramDateTimeIntervalSeconds`.
	ProgramDateTimeIntervalSeconds *float64 `json:"programDateTimeIntervalSeconds"`
	// `CfnOriginEndpoint.HlsPackageProperty.SegmentDurationSeconds`.
	SegmentDurationSeconds *float64 `json:"segmentDurationSeconds"`
	// `CfnOriginEndpoint.HlsPackageProperty.StreamSelection`.
	StreamSelection interface{} `json:"streamSelection"`
	// `CfnOriginEndpoint.HlsPackageProperty.UseAudioRenditionGroup`.
	UseAudioRenditionGroup interface{} `json:"useAudioRenditionGroup"`
}

type CfnOriginEndpoint_MssEncryptionProperty struct {
	// `CfnOriginEndpoint.MssEncryptionProperty.SpekeKeyProvider`.
	SpekeKeyProvider interface{} `json:"spekeKeyProvider"`
}

type CfnOriginEndpoint_MssPackageProperty struct {
	// `CfnOriginEndpoint.MssPackageProperty.Encryption`.
	Encryption interface{} `json:"encryption"`
	// `CfnOriginEndpoint.MssPackageProperty.ManifestWindowSeconds`.
	ManifestWindowSeconds *float64 `json:"manifestWindowSeconds"`
	// `CfnOriginEndpoint.MssPackageProperty.SegmentDurationSeconds`.
	SegmentDurationSeconds *float64 `json:"segmentDurationSeconds"`
	// `CfnOriginEndpoint.MssPackageProperty.StreamSelection`.
	StreamSelection interface{} `json:"streamSelection"`
}

type CfnOriginEndpoint_SpekeKeyProviderProperty struct {
	// `CfnOriginEndpoint.SpekeKeyProviderProperty.ResourceId`.
	ResourceId *string `json:"resourceId"`
	// `CfnOriginEndpoint.SpekeKeyProviderProperty.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `CfnOriginEndpoint.SpekeKeyProviderProperty.SystemIds`.
	SystemIds *[]*string `json:"systemIds"`
	// `CfnOriginEndpoint.SpekeKeyProviderProperty.Url`.
	Url *string `json:"url"`
	// `CfnOriginEndpoint.SpekeKeyProviderProperty.CertificateArn`.
	CertificateArn *string `json:"certificateArn"`
}

type CfnOriginEndpoint_StreamSelectionProperty struct {
	// `CfnOriginEndpoint.StreamSelectionProperty.MaxVideoBitsPerSecond`.
	MaxVideoBitsPerSecond *float64 `json:"maxVideoBitsPerSecond"`
	// `CfnOriginEndpoint.StreamSelectionProperty.MinVideoBitsPerSecond`.
	MinVideoBitsPerSecond *float64 `json:"minVideoBitsPerSecond"`
	// `CfnOriginEndpoint.StreamSelectionProperty.StreamOrder`.
	StreamOrder *string `json:"streamOrder"`
}

// Properties for defining a `AWS::MediaPackage::OriginEndpoint`.
type CfnOriginEndpointProps struct {
	// `AWS::MediaPackage::OriginEndpoint.ChannelId`.
	ChannelId *string `json:"channelId"`
	// `AWS::MediaPackage::OriginEndpoint.Id`.
	Id *string `json:"id"`
	// `AWS::MediaPackage::OriginEndpoint.Authorization`.
	Authorization interface{} `json:"authorization"`
	// `AWS::MediaPackage::OriginEndpoint.CmafPackage`.
	CmafPackage interface{} `json:"cmafPackage"`
	// `AWS::MediaPackage::OriginEndpoint.DashPackage`.
	DashPackage interface{} `json:"dashPackage"`
	// `AWS::MediaPackage::OriginEndpoint.Description`.
	Description *string `json:"description"`
	// `AWS::MediaPackage::OriginEndpoint.HlsPackage`.
	HlsPackage interface{} `json:"hlsPackage"`
	// `AWS::MediaPackage::OriginEndpoint.ManifestName`.
	ManifestName *string `json:"manifestName"`
	// `AWS::MediaPackage::OriginEndpoint.MssPackage`.
	MssPackage interface{} `json:"mssPackage"`
	// `AWS::MediaPackage::OriginEndpoint.Origination`.
	Origination *string `json:"origination"`
	// `AWS::MediaPackage::OriginEndpoint.StartoverWindowSeconds`.
	StartoverWindowSeconds *float64 `json:"startoverWindowSeconds"`
	// `AWS::MediaPackage::OriginEndpoint.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// `AWS::MediaPackage::OriginEndpoint.TimeDelaySeconds`.
	TimeDelaySeconds *float64 `json:"timeDelaySeconds"`
	// `AWS::MediaPackage::OriginEndpoint.Whitelist`.
	Whitelist *[]*string `json:"whitelist"`
}

// A CloudFormation `AWS::MediaPackage::PackagingConfiguration`.
type CfnPackagingConfiguration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CmafPackage() interface{}
	SetCmafPackage(val interface{})
	CreationStack() *[]*string
	DashPackage() interface{}
	SetDashPackage(val interface{})
	HlsPackage() interface{}
	SetHlsPackage(val interface{})
	Id() *string
	SetId(val *string)
	LogicalId() *string
	MssPackage() interface{}
	SetMssPackage(val interface{})
	Node() awscdk.ConstructNode
	PackagingGroupId() *string
	SetPackagingGroupId(val *string)
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

// The jsii proxy struct for CfnPackagingConfiguration
type jsiiProxy_CfnPackagingConfiguration struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPackagingConfiguration) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) CmafPackage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cmafPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) DashPackage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dashPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) HlsPackage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hlsPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) MssPackage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mssPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) PackagingGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packagingGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::MediaPackage::PackagingConfiguration`.
func NewCfnPackagingConfiguration(scope awscdk.Construct, id *string, props *CfnPackagingConfigurationProps) CfnPackagingConfiguration {
	_init_.Initialize()

	j := jsiiProxy_CfnPackagingConfiguration{}

	_jsii_.Create(
		"monocdk.aws_mediapackage.CfnPackagingConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MediaPackage::PackagingConfiguration`.
func NewCfnPackagingConfiguration_Override(c CfnPackagingConfiguration, scope awscdk.Construct, id *string, props *CfnPackagingConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_mediapackage.CfnPackagingConfiguration",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPackagingConfiguration) SetCmafPackage(val interface{}) {
	_jsii_.Set(
		j,
		"cmafPackage",
		val,
	)
}

func (j *jsiiProxy_CfnPackagingConfiguration) SetDashPackage(val interface{}) {
	_jsii_.Set(
		j,
		"dashPackage",
		val,
	)
}

func (j *jsiiProxy_CfnPackagingConfiguration) SetHlsPackage(val interface{}) {
	_jsii_.Set(
		j,
		"hlsPackage",
		val,
	)
}

func (j *jsiiProxy_CfnPackagingConfiguration) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CfnPackagingConfiguration) SetMssPackage(val interface{}) {
	_jsii_.Set(
		j,
		"mssPackage",
		val,
	)
}

func (j *jsiiProxy_CfnPackagingConfiguration) SetPackagingGroupId(val *string) {
	_jsii_.Set(
		j,
		"packagingGroupId",
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
func CfnPackagingConfiguration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_mediapackage.CfnPackagingConfiguration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnPackagingConfiguration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_mediapackage.CfnPackagingConfiguration",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnPackagingConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_mediapackage.CfnPackagingConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPackagingConfiguration_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_mediapackage.CfnPackagingConfiguration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnPackagingConfiguration) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnPackagingConfiguration) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnPackagingConfiguration) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnPackagingConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnPackagingConfiguration) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnPackagingConfiguration) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnPackagingConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnPackagingConfiguration) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnPackagingConfiguration) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnPackagingConfiguration) Inspect(inspector awscdk.TreeInspector) {
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
func (c *jsiiProxy_CfnPackagingConfiguration) OnPrepare() {
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
func (c *jsiiProxy_CfnPackagingConfiguration) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnPackagingConfiguration) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnPackagingConfiguration) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnPackagingConfiguration) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnPackagingConfiguration) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnPackagingConfiguration) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnPackagingConfiguration) Synthesize(session awscdk.ISynthesisSession) {
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
func (c *jsiiProxy_CfnPackagingConfiguration) ToString() *string {
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
func (c *jsiiProxy_CfnPackagingConfiguration) Validate() *[]*string {
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
func (c *jsiiProxy_CfnPackagingConfiguration) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnPackagingConfiguration_CmafEncryptionProperty struct {
	// `CfnPackagingConfiguration.CmafEncryptionProperty.SpekeKeyProvider`.
	SpekeKeyProvider interface{} `json:"spekeKeyProvider"`
}

type CfnPackagingConfiguration_CmafPackageProperty struct {
	// `CfnPackagingConfiguration.CmafPackageProperty.HlsManifests`.
	HlsManifests interface{} `json:"hlsManifests"`
	// `CfnPackagingConfiguration.CmafPackageProperty.Encryption`.
	Encryption interface{} `json:"encryption"`
	// `CfnPackagingConfiguration.CmafPackageProperty.IncludeEncoderConfigurationInSegments`.
	IncludeEncoderConfigurationInSegments interface{} `json:"includeEncoderConfigurationInSegments"`
	// `CfnPackagingConfiguration.CmafPackageProperty.SegmentDurationSeconds`.
	SegmentDurationSeconds *float64 `json:"segmentDurationSeconds"`
}

type CfnPackagingConfiguration_DashEncryptionProperty struct {
	// `CfnPackagingConfiguration.DashEncryptionProperty.SpekeKeyProvider`.
	SpekeKeyProvider interface{} `json:"spekeKeyProvider"`
}

type CfnPackagingConfiguration_DashManifestProperty struct {
	// `CfnPackagingConfiguration.DashManifestProperty.ManifestLayout`.
	ManifestLayout *string `json:"manifestLayout"`
	// `CfnPackagingConfiguration.DashManifestProperty.ManifestName`.
	ManifestName *string `json:"manifestName"`
	// `CfnPackagingConfiguration.DashManifestProperty.MinBufferTimeSeconds`.
	MinBufferTimeSeconds *float64 `json:"minBufferTimeSeconds"`
	// `CfnPackagingConfiguration.DashManifestProperty.Profile`.
	Profile *string `json:"profile"`
	// `CfnPackagingConfiguration.DashManifestProperty.StreamSelection`.
	StreamSelection interface{} `json:"streamSelection"`
}

type CfnPackagingConfiguration_DashPackageProperty struct {
	// `CfnPackagingConfiguration.DashPackageProperty.DashManifests`.
	DashManifests interface{} `json:"dashManifests"`
	// `CfnPackagingConfiguration.DashPackageProperty.Encryption`.
	Encryption interface{} `json:"encryption"`
	// `CfnPackagingConfiguration.DashPackageProperty.IncludeEncoderConfigurationInSegments`.
	IncludeEncoderConfigurationInSegments interface{} `json:"includeEncoderConfigurationInSegments"`
	// `CfnPackagingConfiguration.DashPackageProperty.PeriodTriggers`.
	PeriodTriggers *[]*string `json:"periodTriggers"`
	// `CfnPackagingConfiguration.DashPackageProperty.SegmentDurationSeconds`.
	SegmentDurationSeconds *float64 `json:"segmentDurationSeconds"`
	// `CfnPackagingConfiguration.DashPackageProperty.SegmentTemplateFormat`.
	SegmentTemplateFormat *string `json:"segmentTemplateFormat"`
}

type CfnPackagingConfiguration_HlsEncryptionProperty struct {
	// `CfnPackagingConfiguration.HlsEncryptionProperty.SpekeKeyProvider`.
	SpekeKeyProvider interface{} `json:"spekeKeyProvider"`
	// `CfnPackagingConfiguration.HlsEncryptionProperty.ConstantInitializationVector`.
	ConstantInitializationVector *string `json:"constantInitializationVector"`
	// `CfnPackagingConfiguration.HlsEncryptionProperty.EncryptionMethod`.
	EncryptionMethod *string `json:"encryptionMethod"`
}

type CfnPackagingConfiguration_HlsManifestProperty struct {
	// `CfnPackagingConfiguration.HlsManifestProperty.AdMarkers`.
	AdMarkers *string `json:"adMarkers"`
	// `CfnPackagingConfiguration.HlsManifestProperty.IncludeIframeOnlyStream`.
	IncludeIframeOnlyStream interface{} `json:"includeIframeOnlyStream"`
	// `CfnPackagingConfiguration.HlsManifestProperty.ManifestName`.
	ManifestName *string `json:"manifestName"`
	// `CfnPackagingConfiguration.HlsManifestProperty.ProgramDateTimeIntervalSeconds`.
	ProgramDateTimeIntervalSeconds *float64 `json:"programDateTimeIntervalSeconds"`
	// `CfnPackagingConfiguration.HlsManifestProperty.RepeatExtXKey`.
	RepeatExtXKey interface{} `json:"repeatExtXKey"`
	// `CfnPackagingConfiguration.HlsManifestProperty.StreamSelection`.
	StreamSelection interface{} `json:"streamSelection"`
}

type CfnPackagingConfiguration_HlsPackageProperty struct {
	// `CfnPackagingConfiguration.HlsPackageProperty.HlsManifests`.
	HlsManifests interface{} `json:"hlsManifests"`
	// `CfnPackagingConfiguration.HlsPackageProperty.Encryption`.
	Encryption interface{} `json:"encryption"`
	// `CfnPackagingConfiguration.HlsPackageProperty.SegmentDurationSeconds`.
	SegmentDurationSeconds *float64 `json:"segmentDurationSeconds"`
	// `CfnPackagingConfiguration.HlsPackageProperty.UseAudioRenditionGroup`.
	UseAudioRenditionGroup interface{} `json:"useAudioRenditionGroup"`
}

type CfnPackagingConfiguration_MssEncryptionProperty struct {
	// `CfnPackagingConfiguration.MssEncryptionProperty.SpekeKeyProvider`.
	SpekeKeyProvider interface{} `json:"spekeKeyProvider"`
}

type CfnPackagingConfiguration_MssManifestProperty struct {
	// `CfnPackagingConfiguration.MssManifestProperty.ManifestName`.
	ManifestName *string `json:"manifestName"`
	// `CfnPackagingConfiguration.MssManifestProperty.StreamSelection`.
	StreamSelection interface{} `json:"streamSelection"`
}

type CfnPackagingConfiguration_MssPackageProperty struct {
	// `CfnPackagingConfiguration.MssPackageProperty.MssManifests`.
	MssManifests interface{} `json:"mssManifests"`
	// `CfnPackagingConfiguration.MssPackageProperty.Encryption`.
	Encryption interface{} `json:"encryption"`
	// `CfnPackagingConfiguration.MssPackageProperty.SegmentDurationSeconds`.
	SegmentDurationSeconds *float64 `json:"segmentDurationSeconds"`
}

type CfnPackagingConfiguration_SpekeKeyProviderProperty struct {
	// `CfnPackagingConfiguration.SpekeKeyProviderProperty.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `CfnPackagingConfiguration.SpekeKeyProviderProperty.SystemIds`.
	SystemIds *[]*string `json:"systemIds"`
	// `CfnPackagingConfiguration.SpekeKeyProviderProperty.Url`.
	Url *string `json:"url"`
}

type CfnPackagingConfiguration_StreamSelectionProperty struct {
	// `CfnPackagingConfiguration.StreamSelectionProperty.MaxVideoBitsPerSecond`.
	MaxVideoBitsPerSecond *float64 `json:"maxVideoBitsPerSecond"`
	// `CfnPackagingConfiguration.StreamSelectionProperty.MinVideoBitsPerSecond`.
	MinVideoBitsPerSecond *float64 `json:"minVideoBitsPerSecond"`
	// `CfnPackagingConfiguration.StreamSelectionProperty.StreamOrder`.
	StreamOrder *string `json:"streamOrder"`
}

// Properties for defining a `AWS::MediaPackage::PackagingConfiguration`.
type CfnPackagingConfigurationProps struct {
	// `AWS::MediaPackage::PackagingConfiguration.Id`.
	Id *string `json:"id"`
	// `AWS::MediaPackage::PackagingConfiguration.PackagingGroupId`.
	PackagingGroupId *string `json:"packagingGroupId"`
	// `AWS::MediaPackage::PackagingConfiguration.CmafPackage`.
	CmafPackage interface{} `json:"cmafPackage"`
	// `AWS::MediaPackage::PackagingConfiguration.DashPackage`.
	DashPackage interface{} `json:"dashPackage"`
	// `AWS::MediaPackage::PackagingConfiguration.HlsPackage`.
	HlsPackage interface{} `json:"hlsPackage"`
	// `AWS::MediaPackage::PackagingConfiguration.MssPackage`.
	MssPackage interface{} `json:"mssPackage"`
	// `AWS::MediaPackage::PackagingConfiguration.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::MediaPackage::PackagingGroup`.
type CfnPackagingGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrDomainName() *string
	Authorization() interface{}
	SetAuthorization(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	EgressAccessLogs() interface{}
	SetEgressAccessLogs(val interface{})
	Id() *string
	SetId(val *string)
	LogicalId() *string
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

// The jsii proxy struct for CfnPackagingGroup
type jsiiProxy_CfnPackagingGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPackagingGroup) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingGroup) AttrDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingGroup) Authorization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingGroup) EgressAccessLogs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"egressAccessLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingGroup) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::MediaPackage::PackagingGroup`.
func NewCfnPackagingGroup(scope awscdk.Construct, id *string, props *CfnPackagingGroupProps) CfnPackagingGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnPackagingGroup{}

	_jsii_.Create(
		"monocdk.aws_mediapackage.CfnPackagingGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MediaPackage::PackagingGroup`.
func NewCfnPackagingGroup_Override(c CfnPackagingGroup, scope awscdk.Construct, id *string, props *CfnPackagingGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_mediapackage.CfnPackagingGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPackagingGroup) SetAuthorization(val interface{}) {
	_jsii_.Set(
		j,
		"authorization",
		val,
	)
}

func (j *jsiiProxy_CfnPackagingGroup) SetEgressAccessLogs(val interface{}) {
	_jsii_.Set(
		j,
		"egressAccessLogs",
		val,
	)
}

func (j *jsiiProxy_CfnPackagingGroup) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
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
func CfnPackagingGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_mediapackage.CfnPackagingGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnPackagingGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_mediapackage.CfnPackagingGroup",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnPackagingGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_mediapackage.CfnPackagingGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPackagingGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_mediapackage.CfnPackagingGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnPackagingGroup) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnPackagingGroup) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnPackagingGroup) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnPackagingGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnPackagingGroup) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnPackagingGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnPackagingGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnPackagingGroup) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnPackagingGroup) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnPackagingGroup) Inspect(inspector awscdk.TreeInspector) {
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
func (c *jsiiProxy_CfnPackagingGroup) OnPrepare() {
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
func (c *jsiiProxy_CfnPackagingGroup) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnPackagingGroup) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnPackagingGroup) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnPackagingGroup) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnPackagingGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnPackagingGroup) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnPackagingGroup) Synthesize(session awscdk.ISynthesisSession) {
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
func (c *jsiiProxy_CfnPackagingGroup) ToString() *string {
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
func (c *jsiiProxy_CfnPackagingGroup) Validate() *[]*string {
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
func (c *jsiiProxy_CfnPackagingGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnPackagingGroup_AuthorizationProperty struct {
	// `CfnPackagingGroup.AuthorizationProperty.CdnIdentifierSecret`.
	CdnIdentifierSecret *string `json:"cdnIdentifierSecret"`
	// `CfnPackagingGroup.AuthorizationProperty.SecretsRoleArn`.
	SecretsRoleArn *string `json:"secretsRoleArn"`
}

type CfnPackagingGroup_LogConfigurationProperty struct {
	// `CfnPackagingGroup.LogConfigurationProperty.LogGroupName`.
	LogGroupName *string `json:"logGroupName"`
}

// Properties for defining a `AWS::MediaPackage::PackagingGroup`.
type CfnPackagingGroupProps struct {
	// `AWS::MediaPackage::PackagingGroup.Id`.
	Id *string `json:"id"`
	// `AWS::MediaPackage::PackagingGroup.Authorization`.
	Authorization interface{} `json:"authorization"`
	// `AWS::MediaPackage::PackagingGroup.EgressAccessLogs`.
	EgressAccessLogs interface{} `json:"egressAccessLogs"`
	// `AWS::MediaPackage::PackagingGroup.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

