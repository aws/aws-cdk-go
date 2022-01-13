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
// The AWS::AmplifyUIBuilder::Component resource specifies a component within an Amplify app. A component is a user interface (UI) element that you can customize. Use `ComponentChild` to configure an instance of a `Component` . A `ComponentChild` instance inherits the configuration of the main `Component` .
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

// The `ComponentBindingPropertiesValueProperties` property specifies the data binding configuration for a specific property using data stored in AWS .
//
// For AWS connected properties, you can bind a property to data stored in an Amazon S3 bucket, an Amplify DataStore model or an authenticated user attribute.
//
// TODO: EXAMPLE
//
type CfnComponent_ComponentBindingPropertiesValuePropertiesProperty struct {
	// An Amazon S3 bucket.
	Bucket *string `json:"bucket" yaml:"bucket"`
	// The default value to assign to the property.
	DefaultValue *string `json:"defaultValue" yaml:"defaultValue"`
	// The field to bind the data to.
	Field *string `json:"field" yaml:"field"`
	// The storage key for an Amazon S3 bucket.
	Key *string `json:"key" yaml:"key"`
	// An Amplify DataStore model.
	Model *string `json:"model" yaml:"model"`
	// A list of predicates for binding a component's properties to data.
	Predicates interface{} `json:"predicates" yaml:"predicates"`
	// An authenticated user attribute.
	UserAttribute *string `json:"userAttribute" yaml:"userAttribute"`
}

// The `ComponentBindingPropertiesValue` property specifies the data binding configuration for a component at runtime.
//
// You can use `ComponentBindingPropertiesValue` to add exposed properties to a component to allow different values to be entered when a component is reused in different places in an app.
//
// TODO: EXAMPLE
//
type CfnComponent_ComponentBindingPropertiesValueProperty struct {
	// Describes the properties to customize with data at runtime.
	BindingProperties interface{} `json:"bindingProperties" yaml:"bindingProperties"`
	// The default value of the property.
	DefaultValue *string `json:"defaultValue" yaml:"defaultValue"`
	// The property type.
	Type *string `json:"type" yaml:"type"`
}

// The `ComponentChild` property specifies a nested UI configuration within a parent `Component` .
//
// TODO: EXAMPLE
//
type CfnComponent_ComponentChildProperty struct {
	// The type of the child component.
	ComponentType *string `json:"componentType" yaml:"componentType"`
	// The name of the child component.
	Name *string `json:"name" yaml:"name"`
	// Describes the properties of the child component.
	Properties interface{} `json:"properties" yaml:"properties"`
	// The list of `ComponentChild` instances for this component.
	Children interface{} `json:"children" yaml:"children"`
}

// The `ComponentConditionProperty` property specifies a conditional expression for setting a component property.
//
// Use `ComponentConditionProperty` to set a property to different values conditionally, based on the value of another property.
//
// TODO: EXAMPLE
//
type CfnComponent_ComponentConditionPropertyProperty struct {
	// The value to assign to the property if the condition is not met.
	Else interface{} `json:"else" yaml:"else"`
	// The name of a field.
	//
	// Specify this when the property is a data model.
	Field *string `json:"field" yaml:"field"`
	// The value of the property to evaluate.
	Operand *string `json:"operand" yaml:"operand"`
	// The operator to use to perform the evaluation, such as `eq` to represent equals.
	Operator *string `json:"operator" yaml:"operator"`
	// The name of the conditional property.
	Property *string `json:"property" yaml:"property"`
	// The value to assign to the property if the condition is met.
	Then interface{} `json:"then" yaml:"then"`
}

// The `ComponentDataConfiguration` property specifies the configuration for binding a component's properties to data.
//
// TODO: EXAMPLE
//
type CfnComponent_ComponentDataConfigurationProperty struct {
	// The name of the data model to use to bind data to a component.
	Model *string `json:"model" yaml:"model"`
	// A list of IDs to use to bind data to a component.
	//
	// Use this property to bind specifically chosen data, rather than data retrieved from a query.
	Identifiers *[]*string `json:"identifiers" yaml:"identifiers"`
	// Represents the conditional logic to use when binding data to a component.
	//
	// Use this property to retrieve only a subset of the data in a collection.
	Predicate interface{} `json:"predicate" yaml:"predicate"`
	// Describes how to sort the component's properties.
	Sort interface{} `json:"sort" yaml:"sort"`
}

// The `ComponentOverrides` property specifies the component's properties that can be overriden in a customized instance of the component.
//
// TODO: EXAMPLE
//
type CfnComponent_ComponentOverridesProperty struct {
}

// The `ComponentOverridesValue` property specifies the value of the component's properties that can be overriden in a customized instance of the component.
//
// TODO: EXAMPLE
//
type CfnComponent_ComponentOverridesValueProperty struct {
}

// The `ComponentProperties` property specifies the component's properties.
//
// TODO: EXAMPLE
//
type CfnComponent_ComponentPropertiesProperty struct {
}

// The `ComponentPropertyBindingProperties` property specifies a component property to associate with a binding property.
//
// This enables exposed properties on the top level component to propagate data to the component's property values.
//
// TODO: EXAMPLE
//
type CfnComponent_ComponentPropertyBindingPropertiesProperty struct {
	// The component property to bind to the data field.
	Property *string `json:"property" yaml:"property"`
	// The data field to bind the property to.
	Field *string `json:"field" yaml:"field"`
}

// The `ComponentProperty` property specifies the configuration for all of a component's properties.
//
// Use `ComponentProperty` to specify the values to render or bind by default.
//
// TODO: EXAMPLE
//
type CfnComponent_ComponentPropertyProperty struct {
	// The information to bind the component property to data at runtime.
	BindingProperties interface{} `json:"bindingProperties" yaml:"bindingProperties"`
	// The information to bind the component property to form data.
	Bindings interface{} `json:"bindings" yaml:"bindings"`
	// The information to bind the component property to data at runtime.
	//
	// Use this for collection components.
	CollectionBindingProperties interface{} `json:"collectionBindingProperties" yaml:"collectionBindingProperties"`
	// A list of component properties to concatenate to create the value to assign to this component property.
	Concat interface{} `json:"concat" yaml:"concat"`
	// The conditional expression to use to assign a value to the component property.
	Condition interface{} `json:"condition" yaml:"condition"`
	// Specifies whether the user configured the property in Amplify Studio after importing it.
	Configured interface{} `json:"configured" yaml:"configured"`
	// The default value to assign to the component property.
	DefaultValue *string `json:"defaultValue" yaml:"defaultValue"`
	// An event that occurs in your app.
	//
	// Use this for workflow data binding.
	Event *string `json:"event" yaml:"event"`
	// The default value assigned to the property when the component is imported into an app.
	ImportedValue *string `json:"importedValue" yaml:"importedValue"`
	// The data model to use to assign a value to the component property.
	Model *string `json:"model" yaml:"model"`
	// The component type.
	Type *string `json:"type" yaml:"type"`
	// An authenticated user attribute to use to assign a value to the component property.
	UserAttribute *string `json:"userAttribute" yaml:"userAttribute"`
	// The value to assign to the component property.
	Value *string `json:"value" yaml:"value"`
}

// The `ComponentVariant` property specifies the style configuration of a unique variation of a main component.
//
// TODO: EXAMPLE
//
type CfnComponent_ComponentVariantProperty struct {
	// The properties of the component variant that can be overriden when customizing an instance of the component.
	Overrides interface{} `json:"overrides" yaml:"overrides"`
	// The combination of variants that comprise this variant.
	VariantValues interface{} `json:"variantValues" yaml:"variantValues"`
}

// The `ComponentVariantValues` property specifies the combination of variants that comprise a `ComponentVariant` .
//
// TODO: EXAMPLE
//
type CfnComponent_ComponentVariantValuesProperty struct {
}

// The `FormBindings` property specifies how to bind a component's properties to form data.
//
// TODO: EXAMPLE
//
type CfnComponent_FormBindingsProperty struct {
}

// The `Predicate` property specifies information for generating Amplify DataStore queries.
//
// Use `Predicate` to retrieve a subset of the data in a collection.
//
// TODO: EXAMPLE
//
type CfnComponent_PredicateProperty struct {
	// A list of predicates to combine logically.
	And interface{} `json:"and" yaml:"and"`
	// The field to query.
	Field *string `json:"field" yaml:"field"`
	// The value to use when performing the evaluation.
	Operand *string `json:"operand" yaml:"operand"`
	// The operator to use to perform the evaluation.
	Operator *string `json:"operator" yaml:"operator"`
	// A list of predicates to combine logically.
	Or interface{} `json:"or" yaml:"or"`
}

// The `SortProperty` property specifies how to sort the data that you bind to a component.
//
// TODO: EXAMPLE
//
type CfnComponent_SortPropertyProperty struct {
	// The direction of the sort, either ascending or descending.
	Direction *string `json:"direction" yaml:"direction"`
	// The field to perform the sort on.
	Field *string `json:"field" yaml:"field"`
}

// Properties for defining a `CfnComponent`.
//
// TODO: EXAMPLE
//
type CfnComponentProps struct {
	// The information to connect a component's properties to data at runtime.
	BindingProperties interface{} `json:"bindingProperties" yaml:"bindingProperties"`
	// A list of the component's `ComponentChild` instances.
	Children interface{} `json:"children" yaml:"children"`
	// The data binding configuration for the component's properties.
	//
	// Use this for a collection component.
	CollectionProperties interface{} `json:"collectionProperties" yaml:"collectionProperties"`
	// The type of the component.
	//
	// This can be an Amplify custom UI component or another custom component.
	ComponentType *string `json:"componentType" yaml:"componentType"`
	// The name of the component.
	Name *string `json:"name" yaml:"name"`
	// Describes the component's properties that can be overriden in a customized instance of the component.
	Overrides interface{} `json:"overrides" yaml:"overrides"`
	// Describes the component's properties.
	Properties interface{} `json:"properties" yaml:"properties"`
	// The unique ID of the component in its original source system, such as Figma.
	SourceId *string `json:"sourceId" yaml:"sourceId"`
	// One or more key-value pairs to use when tagging the component.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// A list of the component's variants.
	//
	// A variant is a unique style configuration of a main component.
	Variants interface{} `json:"variants" yaml:"variants"`
}

// A CloudFormation `AWS::AmplifyUIBuilder::Theme`.
//
// The AWS::AmplifyUIBuilder::Theme resource specifies a theme within an Amplify app. A theme is a collection of style settings that apply globally to the components associated with the app.
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

// The `ThemeValue` property specifies the configuration of a theme's properties.
//
// TODO: EXAMPLE
//
type CfnTheme_ThemeValueProperty struct {
	// A list of key-value pairs that define the theme's properties.
	Children interface{} `json:"children" yaml:"children"`
	// The value of a theme property.
	Value *string `json:"value" yaml:"value"`
}

// The `ThemeValues` property specifies key-value pair that defines a property of a theme.
//
// TODO: EXAMPLE
//
type CfnTheme_ThemeValuesProperty struct {
	// The name of the property.
	Key *string `json:"key" yaml:"key"`
	// The value of the property.
	Value interface{} `json:"value" yaml:"value"`
}

// Properties for defining a `CfnTheme`.
//
// TODO: EXAMPLE
//
type CfnThemeProps struct {
	// The name of the theme.
	Name *string `json:"name" yaml:"name"`
	// A list of key-value pairs that defines the properties of the theme.
	Values interface{} `json:"values" yaml:"values"`
	// Describes the properties that can be overriden to customize a theme.
	Overrides interface{} `json:"overrides" yaml:"overrides"`
	// One or more key-value pairs to use when tagging the theme.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
}

