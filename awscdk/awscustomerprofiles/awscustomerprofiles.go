package awscustomerprofiles

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscustomerprofiles/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::CustomerProfiles::Domain`.
//
// The AWS::CustomerProfiles::Domain resource specifies an Amazon Connect Customer Profiles Domain.
//
// TODO: EXAMPLE
//
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

func (j *jsiiProxy_CfnDomain) Node() constructs.Node {
	var returns constructs.Node
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
func NewCfnDomain(scope constructs.Construct, id *string, props *CfnDomainProps) CfnDomain {
	_init_.Initialize()

	j := jsiiProxy_CfnDomain{}

	_jsii_.Create(
		"aws-cdk-lib.aws_customerprofiles.CfnDomain",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CustomerProfiles::Domain`.
func NewCfnDomain_Override(c CfnDomain, scope constructs.Construct, id *string, props *CfnDomainProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_customerprofiles.CfnDomain",
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
func CfnDomain_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_customerprofiles.CfnDomain",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDomain_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_customerprofiles.CfnDomain",
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
func CfnDomain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_customerprofiles.CfnDomain",
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
		"aws-cdk-lib.aws_customerprofiles.CfnDomain",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnDomain) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
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
func (c *jsiiProxy_CfnDomain) AddPropertyOverride(propertyPath *string, value interface{}) {
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

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnDomain) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
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

func (c *jsiiProxy_CfnDomain) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnDomain`.
//
// TODO: EXAMPLE
//
type CfnDomainProps struct {
	// The unique name of the domain.
	DomainName *string `json:"domainName" yaml:"domainName"`
	// The URL of the SQS dead letter queue, which is used for reporting errors associated with ingesting data from third party applications.
	//
	// You must set up a policy on the DeadLetterQueue for the SendMessage operation to enable Amazon Connect Customer Profiles to send messages to the DeadLetterQueue.
	DeadLetterQueueUrl *string `json:"deadLetterQueueUrl" yaml:"deadLetterQueueUrl"`
	// The default encryption key, which is an AWS managed key, is used when no specific type of encryption key is specified.
	//
	// It is used to encrypt all data before it is placed in permanent or semi-permanent storage.
	DefaultEncryptionKey *string `json:"defaultEncryptionKey" yaml:"defaultEncryptionKey"`
	// The default number of days until the data within the domain expires.
	DefaultExpirationDays *float64 `json:"defaultExpirationDays" yaml:"defaultExpirationDays"`
	// The tags used to organize, track, or control access for this resource.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::CustomerProfiles::Integration`.
//
// The AWS::CustomerProfiles::Integration resource specifies an Amazon Connect Customer Profiles Integration.
//
// TODO: EXAMPLE
//
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
	Node() constructs.Node
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
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

func (j *jsiiProxy_CfnIntegration) Node() constructs.Node {
	var returns constructs.Node
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
func NewCfnIntegration(scope constructs.Construct, id *string, props *CfnIntegrationProps) CfnIntegration {
	_init_.Initialize()

	j := jsiiProxy_CfnIntegration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_customerprofiles.CfnIntegration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CustomerProfiles::Integration`.
func NewCfnIntegration_Override(c CfnIntegration, scope constructs.Construct, id *string, props *CfnIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_customerprofiles.CfnIntegration",
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
func CfnIntegration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_customerprofiles.CfnIntegration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnIntegration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_customerprofiles.CfnIntegration",
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
func CfnIntegration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_customerprofiles.CfnIntegration",
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
		"aws-cdk-lib.aws_customerprofiles.CfnIntegration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnIntegration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
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
func (c *jsiiProxy_CfnIntegration) AddPropertyOverride(propertyPath *string, value interface{}) {
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

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnIntegration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
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

func (c *jsiiProxy_CfnIntegration) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnIntegration_ConnectorOperatorProperty struct {
	// `CfnIntegration.ConnectorOperatorProperty.Marketo`.
	Marketo *string `json:"marketo" yaml:"marketo"`
	// `CfnIntegration.ConnectorOperatorProperty.S3`.
	S3 *string `json:"s3" yaml:"s3"`
	// `CfnIntegration.ConnectorOperatorProperty.Salesforce`.
	Salesforce *string `json:"salesforce" yaml:"salesforce"`
	// `CfnIntegration.ConnectorOperatorProperty.ServiceNow`.
	ServiceNow *string `json:"serviceNow" yaml:"serviceNow"`
	// `CfnIntegration.ConnectorOperatorProperty.Zendesk`.
	Zendesk *string `json:"zendesk" yaml:"zendesk"`
}

// TODO: EXAMPLE
//
type CfnIntegration_FlowDefinitionProperty struct {
	// `CfnIntegration.FlowDefinitionProperty.FlowName`.
	FlowName *string `json:"flowName" yaml:"flowName"`
	// `CfnIntegration.FlowDefinitionProperty.KmsArn`.
	KmsArn *string `json:"kmsArn" yaml:"kmsArn"`
	// `CfnIntegration.FlowDefinitionProperty.SourceFlowConfig`.
	SourceFlowConfig interface{} `json:"sourceFlowConfig" yaml:"sourceFlowConfig"`
	// `CfnIntegration.FlowDefinitionProperty.Tasks`.
	Tasks interface{} `json:"tasks" yaml:"tasks"`
	// `CfnIntegration.FlowDefinitionProperty.TriggerConfig`.
	TriggerConfig interface{} `json:"triggerConfig" yaml:"triggerConfig"`
	// `CfnIntegration.FlowDefinitionProperty.Description`.
	Description *string `json:"description" yaml:"description"`
}

// TODO: EXAMPLE
//
type CfnIntegration_IncrementalPullConfigProperty struct {
	// `CfnIntegration.IncrementalPullConfigProperty.DatetimeTypeFieldName`.
	DatetimeTypeFieldName *string `json:"datetimeTypeFieldName" yaml:"datetimeTypeFieldName"`
}

// TODO: EXAMPLE
//
type CfnIntegration_MarketoSourcePropertiesProperty struct {
	// `CfnIntegration.MarketoSourcePropertiesProperty.Object`.
	Object *string `json:"object" yaml:"object"`
}

// TODO: EXAMPLE
//
type CfnIntegration_S3SourcePropertiesProperty struct {
	// `CfnIntegration.S3SourcePropertiesProperty.BucketName`.
	BucketName *string `json:"bucketName" yaml:"bucketName"`
	// `CfnIntegration.S3SourcePropertiesProperty.BucketPrefix`.
	BucketPrefix *string `json:"bucketPrefix" yaml:"bucketPrefix"`
}

// TODO: EXAMPLE
//
type CfnIntegration_SalesforceSourcePropertiesProperty struct {
	// `CfnIntegration.SalesforceSourcePropertiesProperty.Object`.
	Object *string `json:"object" yaml:"object"`
	// `CfnIntegration.SalesforceSourcePropertiesProperty.EnableDynamicFieldUpdate`.
	EnableDynamicFieldUpdate interface{} `json:"enableDynamicFieldUpdate" yaml:"enableDynamicFieldUpdate"`
	// `CfnIntegration.SalesforceSourcePropertiesProperty.IncludeDeletedRecords`.
	IncludeDeletedRecords interface{} `json:"includeDeletedRecords" yaml:"includeDeletedRecords"`
}

// TODO: EXAMPLE
//
type CfnIntegration_ScheduledTriggerPropertiesProperty struct {
	// `CfnIntegration.ScheduledTriggerPropertiesProperty.ScheduleExpression`.
	ScheduleExpression *string `json:"scheduleExpression" yaml:"scheduleExpression"`
	// `CfnIntegration.ScheduledTriggerPropertiesProperty.DataPullMode`.
	DataPullMode *string `json:"dataPullMode" yaml:"dataPullMode"`
	// `CfnIntegration.ScheduledTriggerPropertiesProperty.FirstExecutionFrom`.
	FirstExecutionFrom *float64 `json:"firstExecutionFrom" yaml:"firstExecutionFrom"`
	// `CfnIntegration.ScheduledTriggerPropertiesProperty.ScheduleEndTime`.
	ScheduleEndTime *float64 `json:"scheduleEndTime" yaml:"scheduleEndTime"`
	// `CfnIntegration.ScheduledTriggerPropertiesProperty.ScheduleOffset`.
	ScheduleOffset *float64 `json:"scheduleOffset" yaml:"scheduleOffset"`
	// `CfnIntegration.ScheduledTriggerPropertiesProperty.ScheduleStartTime`.
	ScheduleStartTime *float64 `json:"scheduleStartTime" yaml:"scheduleStartTime"`
	// `CfnIntegration.ScheduledTriggerPropertiesProperty.Timezone`.
	Timezone *string `json:"timezone" yaml:"timezone"`
}

// TODO: EXAMPLE
//
type CfnIntegration_ServiceNowSourcePropertiesProperty struct {
	// `CfnIntegration.ServiceNowSourcePropertiesProperty.Object`.
	Object *string `json:"object" yaml:"object"`
}

// TODO: EXAMPLE
//
type CfnIntegration_SourceConnectorPropertiesProperty struct {
	// `CfnIntegration.SourceConnectorPropertiesProperty.Marketo`.
	Marketo interface{} `json:"marketo" yaml:"marketo"`
	// `CfnIntegration.SourceConnectorPropertiesProperty.S3`.
	S3 interface{} `json:"s3" yaml:"s3"`
	// `CfnIntegration.SourceConnectorPropertiesProperty.Salesforce`.
	Salesforce interface{} `json:"salesforce" yaml:"salesforce"`
	// `CfnIntegration.SourceConnectorPropertiesProperty.ServiceNow`.
	ServiceNow interface{} `json:"serviceNow" yaml:"serviceNow"`
	// `CfnIntegration.SourceConnectorPropertiesProperty.Zendesk`.
	Zendesk interface{} `json:"zendesk" yaml:"zendesk"`
}

// TODO: EXAMPLE
//
type CfnIntegration_SourceFlowConfigProperty struct {
	// `CfnIntegration.SourceFlowConfigProperty.ConnectorType`.
	ConnectorType *string `json:"connectorType" yaml:"connectorType"`
	// `CfnIntegration.SourceFlowConfigProperty.SourceConnectorProperties`.
	SourceConnectorProperties interface{} `json:"sourceConnectorProperties" yaml:"sourceConnectorProperties"`
	// `CfnIntegration.SourceFlowConfigProperty.ConnectorProfileName`.
	ConnectorProfileName *string `json:"connectorProfileName" yaml:"connectorProfileName"`
	// `CfnIntegration.SourceFlowConfigProperty.IncrementalPullConfig`.
	IncrementalPullConfig interface{} `json:"incrementalPullConfig" yaml:"incrementalPullConfig"`
}

// TODO: EXAMPLE
//
type CfnIntegration_TaskPropertiesMapProperty struct {
	// `CfnIntegration.TaskPropertiesMapProperty.OperatorPropertyKey`.
	OperatorPropertyKey *string `json:"operatorPropertyKey" yaml:"operatorPropertyKey"`
	// `CfnIntegration.TaskPropertiesMapProperty.Property`.
	Property *string `json:"property" yaml:"property"`
}

// TODO: EXAMPLE
//
type CfnIntegration_TaskProperty struct {
	// `CfnIntegration.TaskProperty.SourceFields`.
	SourceFields *[]*string `json:"sourceFields" yaml:"sourceFields"`
	// `CfnIntegration.TaskProperty.TaskType`.
	TaskType *string `json:"taskType" yaml:"taskType"`
	// `CfnIntegration.TaskProperty.ConnectorOperator`.
	ConnectorOperator interface{} `json:"connectorOperator" yaml:"connectorOperator"`
	// `CfnIntegration.TaskProperty.DestinationField`.
	DestinationField *string `json:"destinationField" yaml:"destinationField"`
	// `CfnIntegration.TaskProperty.TaskProperties`.
	TaskProperties interface{} `json:"taskProperties" yaml:"taskProperties"`
}

// TODO: EXAMPLE
//
type CfnIntegration_TriggerConfigProperty struct {
	// `CfnIntegration.TriggerConfigProperty.TriggerType`.
	TriggerType *string `json:"triggerType" yaml:"triggerType"`
	// `CfnIntegration.TriggerConfigProperty.TriggerProperties`.
	TriggerProperties interface{} `json:"triggerProperties" yaml:"triggerProperties"`
}

// TODO: EXAMPLE
//
type CfnIntegration_TriggerPropertiesProperty struct {
	// `CfnIntegration.TriggerPropertiesProperty.Scheduled`.
	Scheduled interface{} `json:"scheduled" yaml:"scheduled"`
}

// TODO: EXAMPLE
//
type CfnIntegration_ZendeskSourcePropertiesProperty struct {
	// `CfnIntegration.ZendeskSourcePropertiesProperty.Object`.
	Object *string `json:"object" yaml:"object"`
}

// Properties for defining a `CfnIntegration`.
//
// TODO: EXAMPLE
//
type CfnIntegrationProps struct {
	// The unique name of the domain.
	DomainName *string `json:"domainName" yaml:"domainName"`
	// The name of the profile object type mapping to use.
	ObjectTypeName *string `json:"objectTypeName" yaml:"objectTypeName"`
	// `AWS::CustomerProfiles::Integration.FlowDefinition`.
	FlowDefinition interface{} `json:"flowDefinition" yaml:"flowDefinition"`
	// The tags used to organize, track, or control access for this resource.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// The URI of the S3 bucket or any other type of data source.
	Uri *string `json:"uri" yaml:"uri"`
}

// A CloudFormation `AWS::CustomerProfiles::ObjectType`.
//
// The AWS::CustomerProfiles::ObjectType resource specifies an Amazon Connect Customer Profiles Object Type Mapping.
//
// TODO: EXAMPLE
//
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
	Node() constructs.Node
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
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

func (j *jsiiProxy_CfnObjectType) Node() constructs.Node {
	var returns constructs.Node
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
func NewCfnObjectType(scope constructs.Construct, id *string, props *CfnObjectTypeProps) CfnObjectType {
	_init_.Initialize()

	j := jsiiProxy_CfnObjectType{}

	_jsii_.Create(
		"aws-cdk-lib.aws_customerprofiles.CfnObjectType",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CustomerProfiles::ObjectType`.
func NewCfnObjectType_Override(c CfnObjectType, scope constructs.Construct, id *string, props *CfnObjectTypeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_customerprofiles.CfnObjectType",
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
func CfnObjectType_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_customerprofiles.CfnObjectType",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnObjectType_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_customerprofiles.CfnObjectType",
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
func CfnObjectType_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_customerprofiles.CfnObjectType",
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
		"aws-cdk-lib.aws_customerprofiles.CfnObjectType",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnObjectType) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
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
func (c *jsiiProxy_CfnObjectType) AddPropertyOverride(propertyPath *string, value interface{}) {
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

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnObjectType) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
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

func (c *jsiiProxy_CfnObjectType) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A map of the name and ObjectType field.
//
// TODO: EXAMPLE
//
type CfnObjectType_FieldMapProperty struct {
	// Name of the field.
	Name *string `json:"name" yaml:"name"`
	// Represents a field in a ProfileObjectType.
	ObjectTypeField interface{} `json:"objectTypeField" yaml:"objectTypeField"`
}

// A unique key map that can be used to map data to the profile.
//
// TODO: EXAMPLE
//
type CfnObjectType_KeyMapProperty struct {
	// Name of the key.
	Name *string `json:"name" yaml:"name"`
	// A list of ObjectTypeKey.
	ObjectTypeKeyList interface{} `json:"objectTypeKeyList" yaml:"objectTypeKeyList"`
}

// Represents a field in a ProfileObjectType.
//
// TODO: EXAMPLE
//
type CfnObjectType_ObjectTypeFieldProperty struct {
	// The content type of the field.
	//
	// Used for determining equality when searching.
	ContentType *string `json:"contentType" yaml:"contentType"`
	// A field of a ProfileObject.
	//
	// For example: _source.FirstName, where “_source” is a ProfileObjectType of a Zendesk user and “FirstName” is a field in that ObjectType.
	Source *string `json:"source" yaml:"source"`
	// The location of the data in the standard ProfileObject model.
	//
	// For example: _profile.Address.PostalCode.
	Target *string `json:"target" yaml:"target"`
}

// An object that defines the Key element of a ProfileObject.
//
// A Key is a special element that can be used to search for a customer profile.
//
// TODO: EXAMPLE
//
type CfnObjectType_ObjectTypeKeyProperty struct {
	// The reference for the key name of the fields map.
	FieldNames *[]*string `json:"fieldNames" yaml:"fieldNames"`
	// The types of keys that a ProfileObject can have.
	//
	// Each ProfileObject can have only 1 UNIQUE key but multiple PROFILE keys. PROFILE means that this key can be used to tie an object to a PROFILE. UNIQUE means that it can be used to uniquely identify an object. If a key a is marked as SECONDARY, it will be used to search for profiles after all other PROFILE keys have been searched. A LOOKUP_ONLY key is only used to match a profile but is not persisted to be used for searching of the profile. A NEW_ONLY key is only used if the profile does not already exist before the object is ingested, otherwise it is only used for matching objects to profiles.
	StandardIdentifiers *[]*string `json:"standardIdentifiers" yaml:"standardIdentifiers"`
}

// Properties for defining a `CfnObjectType`.
//
// TODO: EXAMPLE
//
type CfnObjectTypeProps struct {
	// The unique name of the domain.
	DomainName *string `json:"domainName" yaml:"domainName"`
	// Indicates whether a profile should be created when data is received if one doesn’t exist for an object of this type.
	//
	// The default is `FALSE` . If the AllowProfileCreation flag is set to `FALSE` , then the service tries to fetch a standard profile and associate this object with the profile. If it is set to `TRUE` , and if no match is found, then the service creates a new standard profile.
	AllowProfileCreation interface{} `json:"allowProfileCreation" yaml:"allowProfileCreation"`
	// The description of the profile object type mapping.
	Description *string `json:"description" yaml:"description"`
	// The customer-provided key to encrypt the profile object that will be created in this profile object type mapping.
	//
	// If not specified the system will use the encryption key of the domain.
	EncryptionKey *string `json:"encryptionKey" yaml:"encryptionKey"`
	// The number of days until the data of this type expires.
	ExpirationDays *float64 `json:"expirationDays" yaml:"expirationDays"`
	// A list of field definitions for the object type mapping.
	Fields interface{} `json:"fields" yaml:"fields"`
	// A list of keys that can be used to map data to the profile or search for the profile.
	Keys interface{} `json:"keys" yaml:"keys"`
	// The name of the profile object type.
	ObjectTypeName *string `json:"objectTypeName" yaml:"objectTypeName"`
	// The tags used to organize, track, or control access for this resource.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// A unique identifier for the template mapping.
	//
	// This can be used instead of specifying the Keys and Fields properties directly.
	TemplateId *string `json:"templateId" yaml:"templateId"`
}

