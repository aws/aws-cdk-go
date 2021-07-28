package awscustomerprofiles

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscustomerprofiles/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::CustomerProfiles::Domain`.
type CfnDomain interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrCreatedAt() *string
	AttrLastUpdatedAt() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DeadLetterQueueUrl() *string
	SetDeadLetterQueueUrl(val *string)
	DefaultEncryptionKey() *string
	SetDefaultEncryptionKey(val *string)
	DefaultExpirationDays() *float64
	SetDefaultExpirationDays(val *float64)
	DomainName() *string
	SetDomainName(val *string)
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

// The jsii proxy struct for CfnDomain
type jsiiProxy_CfnDomain struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDomain) AttrCreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrLastUpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastUpdatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) DeadLetterQueueUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deadLetterQueueUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) DefaultEncryptionKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) DefaultExpirationDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultExpirationDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CustomerProfiles::Domain`.
func NewCfnDomain(scope awscdk.Construct, id *string, props *CfnDomainProps) CfnDomain {
	_init_.Initialize()

	j := jsiiProxy_CfnDomain{}

	_jsii_.Create(
		"monocdk.aws_customerprofiles.CfnDomain",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CustomerProfiles::Domain`.
func NewCfnDomain_Override(c CfnDomain, scope awscdk.Construct, id *string, props *CfnDomainProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_customerprofiles.CfnDomain",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDomain) SetDeadLetterQueueUrl(val *string) {
	_jsii_.Set(
		j,
		"deadLetterQueueUrl",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetDefaultEncryptionKey(val *string) {
	_jsii_.Set(
		j,
		"defaultEncryptionKey",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetDefaultExpirationDays(val *float64) {
	_jsii_.Set(
		j,
		"defaultExpirationDays",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
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
func CfnDomain_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_customerprofiles.CfnDomain",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDomain_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_customerprofiles.CfnDomain",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnDomain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_customerprofiles.CfnDomain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDomain_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_customerprofiles.CfnDomain",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnDomain) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnDomain) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnDomain) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnDomain) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnDomain) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnDomain) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnDomain) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnDomain) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnDomain) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnDomain) Inspect(inspector awscdk.TreeInspector) {
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
func (c *jsiiProxy_CfnDomain) OnPrepare() {
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
func (c *jsiiProxy_CfnDomain) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnDomain) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnDomain) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnDomain) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDomain) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnDomain) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnDomain) Synthesize(session awscdk.ISynthesisSession) {
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
func (c *jsiiProxy_CfnDomain) ToString() *string {
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
func (c *jsiiProxy_CfnDomain) Validate() *[]*string {
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
func (c *jsiiProxy_CfnDomain) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::CustomerProfiles::Domain`.
type CfnDomainProps struct {
	// `AWS::CustomerProfiles::Domain.DomainName`.
	DomainName *string `json:"domainName"`
	// `AWS::CustomerProfiles::Domain.DeadLetterQueueUrl`.
	DeadLetterQueueUrl *string `json:"deadLetterQueueUrl"`
	// `AWS::CustomerProfiles::Domain.DefaultEncryptionKey`.
	DefaultEncryptionKey *string `json:"defaultEncryptionKey"`
	// `AWS::CustomerProfiles::Domain.DefaultExpirationDays`.
	DefaultExpirationDays *float64 `json:"defaultExpirationDays"`
	// `AWS::CustomerProfiles::Domain.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::CustomerProfiles::Integration`.
type CfnIntegration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrCreatedAt() *string
	AttrLastUpdatedAt() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DomainName() *string
	SetDomainName(val *string)
	FlowDefinition() interface{}
	SetFlowDefinition(val interface{})
	LogicalId() *string
	Node() awscdk.ConstructNode
	ObjectTypeName() *string
	SetObjectTypeName(val *string)
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	Uri() *string
	SetUri(val *string)
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

// The jsii proxy struct for CfnIntegration
type jsiiProxy_CfnIntegration struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnIntegration) AttrCreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) AttrLastUpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastUpdatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) FlowDefinition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"flowDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) ObjectTypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectTypeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}


// Create a new `AWS::CustomerProfiles::Integration`.
func NewCfnIntegration(scope awscdk.Construct, id *string, props *CfnIntegrationProps) CfnIntegration {
	_init_.Initialize()

	j := jsiiProxy_CfnIntegration{}

	_jsii_.Create(
		"monocdk.aws_customerprofiles.CfnIntegration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CustomerProfiles::Integration`.
func NewCfnIntegration_Override(c CfnIntegration, scope awscdk.Construct, id *string, props *CfnIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_customerprofiles.CfnIntegration",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnIntegration) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration) SetFlowDefinition(val interface{}) {
	_jsii_.Set(
		j,
		"flowDefinition",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration) SetObjectTypeName(val *string) {
	_jsii_.Set(
		j,
		"objectTypeName",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration) SetUri(val *string) {
	_jsii_.Set(
		j,
		"uri",
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
func CfnIntegration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_customerprofiles.CfnIntegration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnIntegration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_customerprofiles.CfnIntegration",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnIntegration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_customerprofiles.CfnIntegration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIntegration_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_customerprofiles.CfnIntegration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnIntegration) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnIntegration) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnIntegration) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnIntegration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnIntegration) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnIntegration) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnIntegration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnIntegration) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnIntegration) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnIntegration) Inspect(inspector awscdk.TreeInspector) {
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
func (c *jsiiProxy_CfnIntegration) OnPrepare() {
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
func (c *jsiiProxy_CfnIntegration) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnIntegration) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnIntegration) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnIntegration) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnIntegration) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnIntegration) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnIntegration) Synthesize(session awscdk.ISynthesisSession) {
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
func (c *jsiiProxy_CfnIntegration) ToString() *string {
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
func (c *jsiiProxy_CfnIntegration) Validate() *[]*string {
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
func (c *jsiiProxy_CfnIntegration) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnIntegration_ConnectorOperatorProperty struct {
	// `CfnIntegration.ConnectorOperatorProperty.Marketo`.
	Marketo *string `json:"marketo"`
	// `CfnIntegration.ConnectorOperatorProperty.S3`.
	S3 *string `json:"s3"`
	// `CfnIntegration.ConnectorOperatorProperty.Salesforce`.
	Salesforce *string `json:"salesforce"`
	// `CfnIntegration.ConnectorOperatorProperty.ServiceNow`.
	ServiceNow *string `json:"serviceNow"`
	// `CfnIntegration.ConnectorOperatorProperty.Zendesk`.
	Zendesk *string `json:"zendesk"`
}

type CfnIntegration_FlowDefinitionProperty struct {
	// `CfnIntegration.FlowDefinitionProperty.FlowName`.
	FlowName *string `json:"flowName"`
	// `CfnIntegration.FlowDefinitionProperty.KmsArn`.
	KmsArn *string `json:"kmsArn"`
	// `CfnIntegration.FlowDefinitionProperty.SourceFlowConfig`.
	SourceFlowConfig interface{} `json:"sourceFlowConfig"`
	// `CfnIntegration.FlowDefinitionProperty.Tasks`.
	Tasks interface{} `json:"tasks"`
	// `CfnIntegration.FlowDefinitionProperty.TriggerConfig`.
	TriggerConfig interface{} `json:"triggerConfig"`
	// `CfnIntegration.FlowDefinitionProperty.Description`.
	Description *string `json:"description"`
}

type CfnIntegration_IncrementalPullConfigProperty struct {
	// `CfnIntegration.IncrementalPullConfigProperty.DatetimeTypeFieldName`.
	DatetimeTypeFieldName *string `json:"datetimeTypeFieldName"`
}

type CfnIntegration_MarketoSourcePropertiesProperty struct {
	// `CfnIntegration.MarketoSourcePropertiesProperty.Object`.
	Object *string `json:"object"`
}

type CfnIntegration_S3SourcePropertiesProperty struct {
	// `CfnIntegration.S3SourcePropertiesProperty.BucketName`.
	BucketName *string `json:"bucketName"`
	// `CfnIntegration.S3SourcePropertiesProperty.BucketPrefix`.
	BucketPrefix *string `json:"bucketPrefix"`
}

type CfnIntegration_SalesforceSourcePropertiesProperty struct {
	// `CfnIntegration.SalesforceSourcePropertiesProperty.Object`.
	Object *string `json:"object"`
	// `CfnIntegration.SalesforceSourcePropertiesProperty.EnableDynamicFieldUpdate`.
	EnableDynamicFieldUpdate interface{} `json:"enableDynamicFieldUpdate"`
	// `CfnIntegration.SalesforceSourcePropertiesProperty.IncludeDeletedRecords`.
	IncludeDeletedRecords interface{} `json:"includeDeletedRecords"`
}

type CfnIntegration_ScheduledTriggerPropertiesProperty struct {
	// `CfnIntegration.ScheduledTriggerPropertiesProperty.ScheduleExpression`.
	ScheduleExpression *string `json:"scheduleExpression"`
	// `CfnIntegration.ScheduledTriggerPropertiesProperty.DataPullMode`.
	DataPullMode *string `json:"dataPullMode"`
	// `CfnIntegration.ScheduledTriggerPropertiesProperty.FirstExecutionFrom`.
	FirstExecutionFrom *float64 `json:"firstExecutionFrom"`
	// `CfnIntegration.ScheduledTriggerPropertiesProperty.ScheduleEndTime`.
	ScheduleEndTime *float64 `json:"scheduleEndTime"`
	// `CfnIntegration.ScheduledTriggerPropertiesProperty.ScheduleOffset`.
	ScheduleOffset *float64 `json:"scheduleOffset"`
	// `CfnIntegration.ScheduledTriggerPropertiesProperty.ScheduleStartTime`.
	ScheduleStartTime *float64 `json:"scheduleStartTime"`
	// `CfnIntegration.ScheduledTriggerPropertiesProperty.Timezone`.
	Timezone *string `json:"timezone"`
}

type CfnIntegration_ServiceNowSourcePropertiesProperty struct {
	// `CfnIntegration.ServiceNowSourcePropertiesProperty.Object`.
	Object *string `json:"object"`
}

type CfnIntegration_SourceConnectorPropertiesProperty struct {
	// `CfnIntegration.SourceConnectorPropertiesProperty.Marketo`.
	Marketo interface{} `json:"marketo"`
	// `CfnIntegration.SourceConnectorPropertiesProperty.S3`.
	S3 interface{} `json:"s3"`
	// `CfnIntegration.SourceConnectorPropertiesProperty.Salesforce`.
	Salesforce interface{} `json:"salesforce"`
	// `CfnIntegration.SourceConnectorPropertiesProperty.ServiceNow`.
	ServiceNow interface{} `json:"serviceNow"`
	// `CfnIntegration.SourceConnectorPropertiesProperty.Zendesk`.
	Zendesk interface{} `json:"zendesk"`
}

type CfnIntegration_SourceFlowConfigProperty struct {
	// `CfnIntegration.SourceFlowConfigProperty.ConnectorType`.
	ConnectorType *string `json:"connectorType"`
	// `CfnIntegration.SourceFlowConfigProperty.SourceConnectorProperties`.
	SourceConnectorProperties interface{} `json:"sourceConnectorProperties"`
	// `CfnIntegration.SourceFlowConfigProperty.ConnectorProfileName`.
	ConnectorProfileName *string `json:"connectorProfileName"`
	// `CfnIntegration.SourceFlowConfigProperty.IncrementalPullConfig`.
	IncrementalPullConfig interface{} `json:"incrementalPullConfig"`
}

type CfnIntegration_TaskPropertiesMapProperty struct {
	// `CfnIntegration.TaskPropertiesMapProperty.OperatorPropertyKey`.
	OperatorPropertyKey *string `json:"operatorPropertyKey"`
	// `CfnIntegration.TaskPropertiesMapProperty.Property`.
	Property *string `json:"property"`
}

type CfnIntegration_TaskProperty struct {
	// `CfnIntegration.TaskProperty.SourceFields`.
	SourceFields *[]*string `json:"sourceFields"`
	// `CfnIntegration.TaskProperty.TaskType`.
	TaskType *string `json:"taskType"`
	// `CfnIntegration.TaskProperty.ConnectorOperator`.
	ConnectorOperator interface{} `json:"connectorOperator"`
	// `CfnIntegration.TaskProperty.DestinationField`.
	DestinationField *string `json:"destinationField"`
	// `CfnIntegration.TaskProperty.TaskProperties`.
	TaskProperties interface{} `json:"taskProperties"`
}

type CfnIntegration_TriggerConfigProperty struct {
	// `CfnIntegration.TriggerConfigProperty.TriggerType`.
	TriggerType *string `json:"triggerType"`
	// `CfnIntegration.TriggerConfigProperty.TriggerProperties`.
	TriggerProperties interface{} `json:"triggerProperties"`
}

type CfnIntegration_TriggerPropertiesProperty struct {
	// `CfnIntegration.TriggerPropertiesProperty.Scheduled`.
	Scheduled interface{} `json:"scheduled"`
}

type CfnIntegration_ZendeskSourcePropertiesProperty struct {
	// `CfnIntegration.ZendeskSourcePropertiesProperty.Object`.
	Object *string `json:"object"`
}

// Properties for defining a `AWS::CustomerProfiles::Integration`.
type CfnIntegrationProps struct {
	// `AWS::CustomerProfiles::Integration.DomainName`.
	DomainName *string `json:"domainName"`
	// `AWS::CustomerProfiles::Integration.ObjectTypeName`.
	ObjectTypeName *string `json:"objectTypeName"`
	// `AWS::CustomerProfiles::Integration.FlowDefinition`.
	FlowDefinition interface{} `json:"flowDefinition"`
	// `AWS::CustomerProfiles::Integration.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// `AWS::CustomerProfiles::Integration.Uri`.
	Uri *string `json:"uri"`
}

// A CloudFormation `AWS::CustomerProfiles::ObjectType`.
type CfnObjectType interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AllowProfileCreation() interface{}
	SetAllowProfileCreation(val interface{})
	AttrCreatedAt() *string
	AttrLastUpdatedAt() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	DomainName() *string
	SetDomainName(val *string)
	EncryptionKey() *string
	SetEncryptionKey(val *string)
	ExpirationDays() *float64
	SetExpirationDays(val *float64)
	Fields() interface{}
	SetFields(val interface{})
	Keys() interface{}
	SetKeys(val interface{})
	LogicalId() *string
	Node() awscdk.ConstructNode
	ObjectTypeName() *string
	SetObjectTypeName(val *string)
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	TemplateId() *string
	SetTemplateId(val *string)
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

// The jsii proxy struct for CfnObjectType
type jsiiProxy_CfnObjectType struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnObjectType) AllowProfileCreation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowProfileCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) AttrCreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) AttrLastUpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastUpdatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) EncryptionKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) ExpirationDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expirationDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) Fields() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) Keys() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) ObjectTypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectTypeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) TemplateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CustomerProfiles::ObjectType`.
func NewCfnObjectType(scope awscdk.Construct, id *string, props *CfnObjectTypeProps) CfnObjectType {
	_init_.Initialize()

	j := jsiiProxy_CfnObjectType{}

	_jsii_.Create(
		"monocdk.aws_customerprofiles.CfnObjectType",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CustomerProfiles::ObjectType`.
func NewCfnObjectType_Override(c CfnObjectType, scope awscdk.Construct, id *string, props *CfnObjectTypeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_customerprofiles.CfnObjectType",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnObjectType) SetAllowProfileCreation(val interface{}) {
	_jsii_.Set(
		j,
		"allowProfileCreation",
		val,
	)
}

func (j *jsiiProxy_CfnObjectType) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnObjectType) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_CfnObjectType) SetEncryptionKey(val *string) {
	_jsii_.Set(
		j,
		"encryptionKey",
		val,
	)
}

func (j *jsiiProxy_CfnObjectType) SetExpirationDays(val *float64) {
	_jsii_.Set(
		j,
		"expirationDays",
		val,
	)
}

func (j *jsiiProxy_CfnObjectType) SetFields(val interface{}) {
	_jsii_.Set(
		j,
		"fields",
		val,
	)
}

func (j *jsiiProxy_CfnObjectType) SetKeys(val interface{}) {
	_jsii_.Set(
		j,
		"keys",
		val,
	)
}

func (j *jsiiProxy_CfnObjectType) SetObjectTypeName(val *string) {
	_jsii_.Set(
		j,
		"objectTypeName",
		val,
	)
}

func (j *jsiiProxy_CfnObjectType) SetTemplateId(val *string) {
	_jsii_.Set(
		j,
		"templateId",
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
func CfnObjectType_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_customerprofiles.CfnObjectType",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnObjectType_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_customerprofiles.CfnObjectType",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnObjectType_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_customerprofiles.CfnObjectType",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnObjectType_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_customerprofiles.CfnObjectType",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnObjectType) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnObjectType) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnObjectType) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnObjectType) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnObjectType) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnObjectType) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnObjectType) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnObjectType) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnObjectType) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnObjectType) Inspect(inspector awscdk.TreeInspector) {
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
func (c *jsiiProxy_CfnObjectType) OnPrepare() {
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
func (c *jsiiProxy_CfnObjectType) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnObjectType) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnObjectType) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnObjectType) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnObjectType) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnObjectType) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnObjectType) Synthesize(session awscdk.ISynthesisSession) {
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
func (c *jsiiProxy_CfnObjectType) ToString() *string {
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
func (c *jsiiProxy_CfnObjectType) Validate() *[]*string {
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
func (c *jsiiProxy_CfnObjectType) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnObjectType_FieldMapProperty struct {
	// `CfnObjectType.FieldMapProperty.Name`.
	Name *string `json:"name"`
	// `CfnObjectType.FieldMapProperty.ObjectTypeField`.
	ObjectTypeField interface{} `json:"objectTypeField"`
}

type CfnObjectType_KeyMapProperty struct {
	// `CfnObjectType.KeyMapProperty.Name`.
	Name *string `json:"name"`
	// `CfnObjectType.KeyMapProperty.ObjectTypeKeyList`.
	ObjectTypeKeyList interface{} `json:"objectTypeKeyList"`
}

type CfnObjectType_ObjectTypeFieldProperty struct {
	// `CfnObjectType.ObjectTypeFieldProperty.ContentType`.
	ContentType *string `json:"contentType"`
	// `CfnObjectType.ObjectTypeFieldProperty.Source`.
	Source *string `json:"source"`
	// `CfnObjectType.ObjectTypeFieldProperty.Target`.
	Target *string `json:"target"`
}

type CfnObjectType_ObjectTypeKeyProperty struct {
	// `CfnObjectType.ObjectTypeKeyProperty.FieldNames`.
	FieldNames *[]*string `json:"fieldNames"`
	// `CfnObjectType.ObjectTypeKeyProperty.StandardIdentifiers`.
	StandardIdentifiers *[]*string `json:"standardIdentifiers"`
}

// Properties for defining a `AWS::CustomerProfiles::ObjectType`.
type CfnObjectTypeProps struct {
	// `AWS::CustomerProfiles::ObjectType.DomainName`.
	DomainName *string `json:"domainName"`
	// `AWS::CustomerProfiles::ObjectType.AllowProfileCreation`.
	AllowProfileCreation interface{} `json:"allowProfileCreation"`
	// `AWS::CustomerProfiles::ObjectType.Description`.
	Description *string `json:"description"`
	// `AWS::CustomerProfiles::ObjectType.EncryptionKey`.
	EncryptionKey *string `json:"encryptionKey"`
	// `AWS::CustomerProfiles::ObjectType.ExpirationDays`.
	ExpirationDays *float64 `json:"expirationDays"`
	// `AWS::CustomerProfiles::ObjectType.Fields`.
	Fields interface{} `json:"fields"`
	// `AWS::CustomerProfiles::ObjectType.Keys`.
	Keys interface{} `json:"keys"`
	// `AWS::CustomerProfiles::ObjectType.ObjectTypeName`.
	ObjectTypeName *string `json:"objectTypeName"`
	// `AWS::CustomerProfiles::ObjectType.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// `AWS::CustomerProfiles::ObjectType.TemplateId`.
	TemplateId *string `json:"templateId"`
}

