package awsamplifyuibuilder

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsamplifyuibuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::AmplifyUIBuilder::Component`.
//
// TODO: EXAMPLE
//
type CfnComponent interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrAppId() *string
	AttrCreatedAt() *string
	AttrEnvironmentName() *string
	AttrId() *string
	AttrModifiedAt() *string
	BindingProperties() interface{}
	SetBindingProperties(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	Children() interface{}
	SetChildren(val interface{})
	CollectionProperties() interface{}
	SetCollectionProperties(val interface{})
	ComponentType() *string
	SetComponentType(val *string)
	CreationStack() *[]*string
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Overrides() interface{}
	SetOverrides(val interface{})
	Properties() interface{}
	SetProperties(val interface{})
	Ref() *string
	SourceId() *string
	SetSourceId(val *string)
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	Variants() interface{}
	SetVariants(val interface{})
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

// The jsii proxy struct for CfnComponent
type jsiiProxy_CfnComponent struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnComponent) AttrAppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAppId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) AttrCreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) AttrEnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEnvironmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) AttrModifiedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrModifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) BindingProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bindingProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Children() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"children",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) CollectionProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"collectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) ComponentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"componentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Overrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Properties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) SourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Variants() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"variants",
		&returns,
	)
	return returns
}


// Create a new `AWS::AmplifyUIBuilder::Component`.
func NewCfnComponent(scope constructs.Construct, id *string, props *CfnComponentProps) CfnComponent {
	_init_.Initialize()

	j := jsiiProxy_CfnComponent{}

	_jsii_.Create(
		"aws-cdk-lib.aws_amplifyuibuilder.CfnComponent",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AmplifyUIBuilder::Component`.
func NewCfnComponent_Override(c CfnComponent, scope constructs.Construct, id *string, props *CfnComponentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_amplifyuibuilder.CfnComponent",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnComponent) SetBindingProperties(val interface{}) {
	_jsii_.Set(
		j,
		"bindingProperties",
		val,
	)
}

func (j *jsiiProxy_CfnComponent) SetChildren(val interface{}) {
	_jsii_.Set(
		j,
		"children",
		val,
	)
}

func (j *jsiiProxy_CfnComponent) SetCollectionProperties(val interface{}) {
	_jsii_.Set(
		j,
		"collectionProperties",
		val,
	)
}

func (j *jsiiProxy_CfnComponent) SetComponentType(val *string) {
	_jsii_.Set(
		j,
		"componentType",
		val,
	)
}

func (j *jsiiProxy_CfnComponent) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnComponent) SetOverrides(val interface{}) {
	_jsii_.Set(
		j,
		"overrides",
		val,
	)
}

func (j *jsiiProxy_CfnComponent) SetProperties(val interface{}) {
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_CfnComponent) SetSourceId(val *string) {
	_jsii_.Set(
		j,
		"sourceId",
		val,
	)
}

func (j *jsiiProxy_CfnComponent) SetVariants(val interface{}) {
	_jsii_.Set(
		j,
		"variants",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnComponent_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_amplifyuibuilder.CfnComponent",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnComponent_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_amplifyuibuilder.CfnComponent",
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
func CfnComponent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_amplifyuibuilder.CfnComponent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnComponent_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_amplifyuibuilder.CfnComponent",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnComponent) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnComponent) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnComponent) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnComponent) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnComponent) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnComponent) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnComponent) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnComponent) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnComponent) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnComponent) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnComponent) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnComponent) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnComponent) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnComponent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponent) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnComponent_ComponentBindingPropertiesValuePropertiesProperty struct {
	// `CfnComponent.ComponentBindingPropertiesValuePropertiesProperty.Bucket`.
	Bucket *string `json:"bucket"`
	// `CfnComponent.ComponentBindingPropertiesValuePropertiesProperty.DefaultValue`.
	DefaultValue *string `json:"defaultValue"`
	// `CfnComponent.ComponentBindingPropertiesValuePropertiesProperty.Field`.
	Field *string `json:"field"`
	// `CfnComponent.ComponentBindingPropertiesValuePropertiesProperty.Key`.
	Key *string `json:"key"`
	// `CfnComponent.ComponentBindingPropertiesValuePropertiesProperty.Model`.
	Model *string `json:"model"`
	// `CfnComponent.ComponentBindingPropertiesValuePropertiesProperty.Predicates`.
	Predicates interface{} `json:"predicates"`
	// `CfnComponent.ComponentBindingPropertiesValuePropertiesProperty.UserAttribute`.
	UserAttribute *string `json:"userAttribute"`
}

// TODO: EXAMPLE
//
type CfnComponent_ComponentBindingPropertiesValueProperty struct {
	// `CfnComponent.ComponentBindingPropertiesValueProperty.BindingProperties`.
	BindingProperties interface{} `json:"bindingProperties"`
	// `CfnComponent.ComponentBindingPropertiesValueProperty.DefaultValue`.
	DefaultValue *string `json:"defaultValue"`
	// `CfnComponent.ComponentBindingPropertiesValueProperty.Type`.
	Type *string `json:"type"`
}

// TODO: EXAMPLE
//
type CfnComponent_ComponentChildProperty struct {
	// `CfnComponent.ComponentChildProperty.Children`.
	Children interface{} `json:"children"`
	// `CfnComponent.ComponentChildProperty.ComponentType`.
	ComponentType *string `json:"componentType"`
	// `CfnComponent.ComponentChildProperty.Name`.
	Name *string `json:"name"`
	// `CfnComponent.ComponentChildProperty.Properties`.
	Properties interface{} `json:"properties"`
}

// TODO: EXAMPLE
//
type CfnComponent_ComponentConditionPropertyProperty struct {
	// `CfnComponent.ComponentConditionPropertyProperty.Else`.
	Else interface{} `json:"else"`
	// `CfnComponent.ComponentConditionPropertyProperty.Field`.
	Field *string `json:"field"`
	// `CfnComponent.ComponentConditionPropertyProperty.Operand`.
	Operand *string `json:"operand"`
	// `CfnComponent.ComponentConditionPropertyProperty.Operator`.
	Operator *string `json:"operator"`
	// `CfnComponent.ComponentConditionPropertyProperty.Property`.
	Property *string `json:"property"`
	// `CfnComponent.ComponentConditionPropertyProperty.Then`.
	Then interface{} `json:"then"`
}

// TODO: EXAMPLE
//
type CfnComponent_ComponentDataConfigurationProperty struct {
	// `CfnComponent.ComponentDataConfigurationProperty.Identifiers`.
	Identifiers *[]*string `json:"identifiers"`
	// `CfnComponent.ComponentDataConfigurationProperty.Model`.
	Model *string `json:"model"`
	// `CfnComponent.ComponentDataConfigurationProperty.Predicate`.
	Predicate interface{} `json:"predicate"`
	// `CfnComponent.ComponentDataConfigurationProperty.Sort`.
	Sort interface{} `json:"sort"`
}

// TODO: EXAMPLE
//
type CfnComponent_ComponentOverridesProperty struct {
}

// TODO: EXAMPLE
//
type CfnComponent_ComponentOverridesValueProperty struct {
}

// TODO: EXAMPLE
//
type CfnComponent_ComponentPropertiesProperty struct {
}

// TODO: EXAMPLE
//
type CfnComponent_ComponentPropertyBindingPropertiesProperty struct {
	// `CfnComponent.ComponentPropertyBindingPropertiesProperty.Field`.
	Field *string `json:"field"`
	// `CfnComponent.ComponentPropertyBindingPropertiesProperty.Property`.
	Property *string `json:"property"`
}

// TODO: EXAMPLE
//
type CfnComponent_ComponentPropertyProperty struct {
	// `CfnComponent.ComponentPropertyProperty.BindingProperties`.
	BindingProperties interface{} `json:"bindingProperties"`
	// `CfnComponent.ComponentPropertyProperty.Bindings`.
	Bindings interface{} `json:"bindings"`
	// `CfnComponent.ComponentPropertyProperty.CollectionBindingProperties`.
	CollectionBindingProperties interface{} `json:"collectionBindingProperties"`
	// `CfnComponent.ComponentPropertyProperty.Concat`.
	Concat interface{} `json:"concat"`
	// `CfnComponent.ComponentPropertyProperty.Condition`.
	Condition interface{} `json:"condition"`
	// `CfnComponent.ComponentPropertyProperty.Configured`.
	Configured interface{} `json:"configured"`
	// `CfnComponent.ComponentPropertyProperty.DefaultValue`.
	DefaultValue *string `json:"defaultValue"`
	// `CfnComponent.ComponentPropertyProperty.Event`.
	Event *string `json:"event"`
	// `CfnComponent.ComponentPropertyProperty.ImportedValue`.
	ImportedValue *string `json:"importedValue"`
	// `CfnComponent.ComponentPropertyProperty.Model`.
	Model *string `json:"model"`
	// `CfnComponent.ComponentPropertyProperty.Type`.
	Type *string `json:"type"`
	// `CfnComponent.ComponentPropertyProperty.UserAttribute`.
	UserAttribute *string `json:"userAttribute"`
	// `CfnComponent.ComponentPropertyProperty.Value`.
	Value *string `json:"value"`
}

// TODO: EXAMPLE
//
type CfnComponent_ComponentVariantProperty struct {
	// `CfnComponent.ComponentVariantProperty.Overrides`.
	Overrides interface{} `json:"overrides"`
	// `CfnComponent.ComponentVariantProperty.VariantValues`.
	VariantValues interface{} `json:"variantValues"`
}

// TODO: EXAMPLE
//
type CfnComponent_ComponentVariantValuesProperty struct {
}

// TODO: EXAMPLE
//
type CfnComponent_FormBindingsProperty struct {
}

// TODO: EXAMPLE
//
type CfnComponent_PredicateProperty struct {
	// `CfnComponent.PredicateProperty.And`.
	And interface{} `json:"and"`
	// `CfnComponent.PredicateProperty.Field`.
	Field *string `json:"field"`
	// `CfnComponent.PredicateProperty.Operand`.
	Operand *string `json:"operand"`
	// `CfnComponent.PredicateProperty.Operator`.
	Operator *string `json:"operator"`
	// `CfnComponent.PredicateProperty.Or`.
	Or interface{} `json:"or"`
}

// TODO: EXAMPLE
//
type CfnComponent_SortPropertyProperty struct {
	// `CfnComponent.SortPropertyProperty.Direction`.
	Direction *string `json:"direction"`
	// `CfnComponent.SortPropertyProperty.Field`.
	Field *string `json:"field"`
}

// Properties for defining a `AWS::AmplifyUIBuilder::Component`.
//
// TODO: EXAMPLE
//
type CfnComponentProps struct {
	// `AWS::AmplifyUIBuilder::Component.BindingProperties`.
	BindingProperties interface{} `json:"bindingProperties"`
	// `AWS::AmplifyUIBuilder::Component.Children`.
	Children interface{} `json:"children"`
	// `AWS::AmplifyUIBuilder::Component.CollectionProperties`.
	CollectionProperties interface{} `json:"collectionProperties"`
	// `AWS::AmplifyUIBuilder::Component.ComponentType`.
	ComponentType *string `json:"componentType"`
	// `AWS::AmplifyUIBuilder::Component.Name`.
	Name *string `json:"name"`
	// `AWS::AmplifyUIBuilder::Component.Overrides`.
	Overrides interface{} `json:"overrides"`
	// `AWS::AmplifyUIBuilder::Component.Properties`.
	Properties interface{} `json:"properties"`
	// `AWS::AmplifyUIBuilder::Component.SourceId`.
	SourceId *string `json:"sourceId"`
	// `AWS::AmplifyUIBuilder::Component.Tags`.
	Tags *map[string]*string `json:"tags"`
	// `AWS::AmplifyUIBuilder::Component.Variants`.
	Variants interface{} `json:"variants"`
}

// A CloudFormation `AWS::AmplifyUIBuilder::Theme`.
//
// TODO: EXAMPLE
//
type CfnTheme interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrAppId() *string
	AttrCreatedAt() *string
	AttrEnvironmentName() *string
	AttrId() *string
	AttrModifiedAt() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Overrides() interface{}
	SetOverrides(val interface{})
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	Values() interface{}
	SetValues(val interface{})
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

// The jsii proxy struct for CfnTheme
type jsiiProxy_CfnTheme struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTheme) AttrAppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAppId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) AttrCreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) AttrEnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEnvironmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) AttrModifiedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrModifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) Overrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) Values() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}


// Create a new `AWS::AmplifyUIBuilder::Theme`.
func NewCfnTheme(scope constructs.Construct, id *string, props *CfnThemeProps) CfnTheme {
	_init_.Initialize()

	j := jsiiProxy_CfnTheme{}

	_jsii_.Create(
		"aws-cdk-lib.aws_amplifyuibuilder.CfnTheme",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AmplifyUIBuilder::Theme`.
func NewCfnTheme_Override(c CfnTheme, scope constructs.Construct, id *string, props *CfnThemeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_amplifyuibuilder.CfnTheme",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTheme) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnTheme) SetOverrides(val interface{}) {
	_jsii_.Set(
		j,
		"overrides",
		val,
	)
}

func (j *jsiiProxy_CfnTheme) SetValues(val interface{}) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnTheme_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_amplifyuibuilder.CfnTheme",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnTheme_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_amplifyuibuilder.CfnTheme",
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
func CfnTheme_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_amplifyuibuilder.CfnTheme",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTheme_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_amplifyuibuilder.CfnTheme",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnTheme) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnTheme) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnTheme) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnTheme) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnTheme) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnTheme) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnTheme) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnTheme) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnTheme) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnTheme) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnTheme) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTheme) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnTheme) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnTheme) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTheme) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnTheme_ThemeValueProperty struct {
	// `CfnTheme.ThemeValueProperty.Children`.
	Children interface{} `json:"children"`
	// `CfnTheme.ThemeValueProperty.Value`.
	Value *string `json:"value"`
}

// TODO: EXAMPLE
//
type CfnTheme_ThemeValuesProperty struct {
	// `CfnTheme.ThemeValuesProperty.Key`.
	Key *string `json:"key"`
	// `CfnTheme.ThemeValuesProperty.Value`.
	Value interface{} `json:"value"`
}

// Properties for defining a `AWS::AmplifyUIBuilder::Theme`.
//
// TODO: EXAMPLE
//
type CfnThemeProps struct {
	// `AWS::AmplifyUIBuilder::Theme.Name`.
	Name *string `json:"name"`
	// `AWS::AmplifyUIBuilder::Theme.Overrides`.
	Overrides interface{} `json:"overrides"`
	// `AWS::AmplifyUIBuilder::Theme.Tags`.
	Tags *map[string]*string `json:"tags"`
	// `AWS::AmplifyUIBuilder::Theme.Values`.
	Values interface{} `json:"values"`
}

