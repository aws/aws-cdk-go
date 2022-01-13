package awswisdom

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Wisdom::Assistant`.
//
// Specifies an Amazon Connect Wisdom assistant.
//
// TODO: EXAMPLE
//
type CfnAssistant interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrAssistantArn() *string
	AttrAssistantId() *string
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
	ServerSideEncryptionConfiguration() interface{}
	SetServerSideEncryptionConfiguration(val interface{})
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

// The jsii proxy struct for CfnAssistant
type jsiiProxy_CfnAssistant struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAssistant) AttrAssistantArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAssistantArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistant) AttrAssistantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAssistantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistant) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistant) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistant) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistant) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistant) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistant) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistant) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistant) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistant) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistant) ServerSideEncryptionConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverSideEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistant) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistant) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistant) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistant) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Wisdom::Assistant`.
func NewCfnAssistant(scope constructs.Construct, id *string, props *CfnAssistantProps) CfnAssistant {
	_init_.Initialize()

	j := jsiiProxy_CfnAssistant{}

	_jsii_.Create(
		"aws-cdk-lib.aws_wisdom.CfnAssistant",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Wisdom::Assistant`.
func NewCfnAssistant_Override(c CfnAssistant, scope constructs.Construct, id *string, props *CfnAssistantProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_wisdom.CfnAssistant",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAssistant) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnAssistant) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnAssistant) SetServerSideEncryptionConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"serverSideEncryptionConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnAssistant) SetType(val *string) {
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
func CfnAssistant_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wisdom.CfnAssistant",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnAssistant_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wisdom.CfnAssistant",
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
func CfnAssistant_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wisdom.CfnAssistant",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAssistant_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_wisdom.CfnAssistant",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnAssistant) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnAssistant) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnAssistant) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnAssistant) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnAssistant) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnAssistant) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnAssistant) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnAssistant) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnAssistant) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnAssistant) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnAssistant) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAssistant) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnAssistant) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnAssistant) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAssistant) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The KMS key used for encryption.
//
// TODO: EXAMPLE
//
type CfnAssistant_ServerSideEncryptionConfigurationProperty struct {
	// The KMS key .
	//
	// For information about valid ID values, see [Key identifiers (KeyId)](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id) .
	KmsKeyId *string `json:"kmsKeyId" yaml:"kmsKeyId"`
}

// A CloudFormation `AWS::Wisdom::AssistantAssociation`.
//
// Specifies an association between an Amazon Connect Wisdom assistant and another resource. Currently, the only supported association is with a knowledge base. An assistant can have only a single association.
//
// TODO: EXAMPLE
//
type CfnAssistantAssociation interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AssistantId() *string
	SetAssistantId(val *string)
	Association() interface{}
	SetAssociation(val interface{})
	AssociationType() *string
	SetAssociationType(val *string)
	AttrAssistantArn() *string
	AttrAssistantAssociationArn() *string
	AttrAssistantAssociationId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
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

// The jsii proxy struct for CfnAssistantAssociation
type jsiiProxy_CfnAssistantAssociation struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAssistantAssociation) AssistantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assistantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistantAssociation) Association() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"association",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistantAssociation) AssociationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistantAssociation) AttrAssistantArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAssistantArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistantAssociation) AttrAssistantAssociationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAssistantAssociationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistantAssociation) AttrAssistantAssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAssistantAssociationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistantAssociation) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistantAssociation) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistantAssociation) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistantAssociation) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistantAssociation) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistantAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistantAssociation) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistantAssociation) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistantAssociation) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssistantAssociation) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Wisdom::AssistantAssociation`.
func NewCfnAssistantAssociation(scope constructs.Construct, id *string, props *CfnAssistantAssociationProps) CfnAssistantAssociation {
	_init_.Initialize()

	j := jsiiProxy_CfnAssistantAssociation{}

	_jsii_.Create(
		"aws-cdk-lib.aws_wisdom.CfnAssistantAssociation",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Wisdom::AssistantAssociation`.
func NewCfnAssistantAssociation_Override(c CfnAssistantAssociation, scope constructs.Construct, id *string, props *CfnAssistantAssociationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_wisdom.CfnAssistantAssociation",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAssistantAssociation) SetAssistantId(val *string) {
	_jsii_.Set(
		j,
		"assistantId",
		val,
	)
}

func (j *jsiiProxy_CfnAssistantAssociation) SetAssociation(val interface{}) {
	_jsii_.Set(
		j,
		"association",
		val,
	)
}

func (j *jsiiProxy_CfnAssistantAssociation) SetAssociationType(val *string) {
	_jsii_.Set(
		j,
		"associationType",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnAssistantAssociation_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wisdom.CfnAssistantAssociation",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnAssistantAssociation_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wisdom.CfnAssistantAssociation",
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
func CfnAssistantAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wisdom.CfnAssistantAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAssistantAssociation_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_wisdom.CfnAssistantAssociation",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnAssistantAssociation) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnAssistantAssociation) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnAssistantAssociation) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnAssistantAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnAssistantAssociation) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnAssistantAssociation) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnAssistantAssociation) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnAssistantAssociation) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnAssistantAssociation) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnAssistantAssociation) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnAssistantAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAssistantAssociation) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnAssistantAssociation) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnAssistantAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAssistantAssociation) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A union type that currently has a single argument, which is the knowledge base ID.
//
// TODO: EXAMPLE
//
type CfnAssistantAssociation_AssociationDataProperty struct {
	// The identifier of the knowledge base.
	KnowledgeBaseId *string `json:"knowledgeBaseId" yaml:"knowledgeBaseId"`
}

// Properties for defining a `CfnAssistantAssociation`.
//
// TODO: EXAMPLE
//
type CfnAssistantAssociationProps struct {
	// The identifier of the Wisdom assistant.
	AssistantId *string `json:"assistantId" yaml:"assistantId"`
	// The identifier of the associated resource.
	Association interface{} `json:"association" yaml:"association"`
	// The type of association.
	AssociationType *string `json:"associationType" yaml:"associationType"`
	// The tags used to organize, track, or control access for this resource.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// Properties for defining a `CfnAssistant`.
//
// TODO: EXAMPLE
//
type CfnAssistantProps struct {
	// The name of the assistant.
	Name *string `json:"name" yaml:"name"`
	// The type of assistant.
	Type *string `json:"type" yaml:"type"`
	// The description of the assistant.
	Description *string `json:"description" yaml:"description"`
	// The KMS key used for encryption.
	ServerSideEncryptionConfiguration interface{} `json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
	// The tags used to organize, track, or control access for this resource.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Wisdom::KnowledgeBase`.
//
// Specifies a knowledge base.
//
// TODO: EXAMPLE
//
type CfnKnowledgeBase interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrKnowledgeBaseArn() *string
	AttrKnowledgeBaseId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	KnowledgeBaseType() *string
	SetKnowledgeBaseType(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	RenderingConfiguration() interface{}
	SetRenderingConfiguration(val interface{})
	ServerSideEncryptionConfiguration() interface{}
	SetServerSideEncryptionConfiguration(val interface{})
	SourceConfiguration() interface{}
	SetSourceConfiguration(val interface{})
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

// The jsii proxy struct for CfnKnowledgeBase
type jsiiProxy_CfnKnowledgeBase struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnKnowledgeBase) AttrKnowledgeBaseArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrKnowledgeBaseArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKnowledgeBase) AttrKnowledgeBaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrKnowledgeBaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKnowledgeBase) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKnowledgeBase) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKnowledgeBase) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKnowledgeBase) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKnowledgeBase) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKnowledgeBase) KnowledgeBaseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"knowledgeBaseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKnowledgeBase) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKnowledgeBase) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKnowledgeBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKnowledgeBase) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKnowledgeBase) RenderingConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"renderingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKnowledgeBase) ServerSideEncryptionConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverSideEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKnowledgeBase) SourceConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKnowledgeBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKnowledgeBase) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKnowledgeBase) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Wisdom::KnowledgeBase`.
func NewCfnKnowledgeBase(scope constructs.Construct, id *string, props *CfnKnowledgeBaseProps) CfnKnowledgeBase {
	_init_.Initialize()

	j := jsiiProxy_CfnKnowledgeBase{}

	_jsii_.Create(
		"aws-cdk-lib.aws_wisdom.CfnKnowledgeBase",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Wisdom::KnowledgeBase`.
func NewCfnKnowledgeBase_Override(c CfnKnowledgeBase, scope constructs.Construct, id *string, props *CfnKnowledgeBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_wisdom.CfnKnowledgeBase",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnKnowledgeBase) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnKnowledgeBase) SetKnowledgeBaseType(val *string) {
	_jsii_.Set(
		j,
		"knowledgeBaseType",
		val,
	)
}

func (j *jsiiProxy_CfnKnowledgeBase) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnKnowledgeBase) SetRenderingConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"renderingConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnKnowledgeBase) SetServerSideEncryptionConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"serverSideEncryptionConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnKnowledgeBase) SetSourceConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"sourceConfiguration",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnKnowledgeBase_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wisdom.CfnKnowledgeBase",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnKnowledgeBase_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wisdom.CfnKnowledgeBase",
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
func CfnKnowledgeBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wisdom.CfnKnowledgeBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnKnowledgeBase_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_wisdom.CfnKnowledgeBase",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnKnowledgeBase) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnKnowledgeBase) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnKnowledgeBase) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnKnowledgeBase) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnKnowledgeBase) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnKnowledgeBase) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnKnowledgeBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnKnowledgeBase) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnKnowledgeBase) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnKnowledgeBase) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnKnowledgeBase) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnKnowledgeBase) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnKnowledgeBase) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnKnowledgeBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnKnowledgeBase) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Configuration information for Amazon AppIntegrations to automatically ingest content.
//
// TODO: EXAMPLE
//
type CfnKnowledgeBase_AppIntegrationsConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the AppIntegrations DataIntegration to use for ingesting content.
	AppIntegrationArn *string `json:"appIntegrationArn" yaml:"appIntegrationArn"`
	// The fields from the source that are made available to your agents in Wisdom.
	//
	// - For [Salesforce](https://docs.aws.amazon.com/https://developer.salesforce.com/docs/atlas.en-us.knowledge_dev.meta/knowledge_dev/sforce_api_objects_knowledge__kav.htm) , you must include at least `Id` , `ArticleNumber` , `VersionNumber` , `Title` , `PublishStatus` , and `IsDeleted` .
	// - For [ServiceNow](https://docs.aws.amazon.com/https://developer.servicenow.com/dev.do#!/reference/api/rome/rest/knowledge-management-api) , you must include at least `number` , `short_description` , `sys_mod_count` , `workflow_state` , and `active` .
	//
	// Make sure to include additional fields. These fields are indexed and used to source recommendations.
	ObjectFields *[]*string `json:"objectFields" yaml:"objectFields"`
}

// Information about how to render the content.
//
// TODO: EXAMPLE
//
type CfnKnowledgeBase_RenderingConfigurationProperty struct {
	// A URI template containing exactly one variable in `${variableName}` format.
	//
	// This can only be set for `EXTERNAL` knowledge bases. For Salesforce and ServiceNow, the variable must be one of the following:
	//
	// - Salesforce: `Id` , `ArticleNumber` , `VersionNumber` , `Title` , `PublishStatus` , or `IsDeleted`
	// - ServiceNow: `number` , `short_description` , `sys_mod_count` , `workflow_state` , or `active`
	//
	// The variable is replaced with the actual value for a piece of content when calling [GetContent](https://docs.aws.amazon.com/wisdom/latest/APIReference/API_GetContent.html) .
	TemplateUri *string `json:"templateUri" yaml:"templateUri"`
}

// The KMS key used for encryption.
//
// TODO: EXAMPLE
//
type CfnKnowledgeBase_ServerSideEncryptionConfigurationProperty struct {
	// The KMS key .
	//
	// For information about valid ID values, see [Key identifiers (KeyId)](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id) .
	KmsKeyId *string `json:"kmsKeyId" yaml:"kmsKeyId"`
}

// Configuration information about the external data source.
//
// TODO: EXAMPLE
//
type CfnKnowledgeBase_SourceConfigurationProperty struct {
	// Configuration information for Amazon AppIntegrations to automatically ingest content.
	AppIntegrations interface{} `json:"appIntegrations" yaml:"appIntegrations"`
}

// Properties for defining a `CfnKnowledgeBase`.
//
// TODO: EXAMPLE
//
type CfnKnowledgeBaseProps struct {
	// The type of knowledge base.
	//
	// Only CUSTOM knowledge bases allow you to upload your own content. EXTERNAL knowledge bases support integrations with third-party systems whose content is synchronized automatically.
	KnowledgeBaseType *string `json:"knowledgeBaseType" yaml:"knowledgeBaseType"`
	// The name of the knowledge base.
	Name *string `json:"name" yaml:"name"`
	// The description.
	Description *string `json:"description" yaml:"description"`
	// Information about how to render the content.
	RenderingConfiguration interface{} `json:"renderingConfiguration" yaml:"renderingConfiguration"`
	// The KMS key used for encryption.
	ServerSideEncryptionConfiguration interface{} `json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
	// The source of the knowledge base content.
	//
	// Only set this argument for EXTERNAL knowledge bases.
	SourceConfiguration interface{} `json:"sourceConfiguration" yaml:"sourceConfiguration"`
	// The tags used to organize, track, or control access for this resource.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

