package awsiottwinmaker

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiottwinmaker/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::IoTTwinMaker::ComponentType`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataTypeProperty_ dataTypeProperty
//   var dataValueProperty_ dataValueProperty
//   var relationshipValue interface{}
//
//   cfnComponentType := awscdk.Aws_iottwinmaker.NewCfnComponentType(this, jsii.String("MyCfnComponentType"), &cfnComponentTypeProps{
//   	componentTypeId: jsii.String("componentTypeId"),
//   	workspaceId: jsii.String("workspaceId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	extendsFrom: []*string{
//   		jsii.String("extendsFrom"),
//   	},
//   	functions: map[string]interface{}{
//   		"functionsKey": &FunctionProperty{
//   			"implementedBy": &DataConnectorProperty{
//   				"isNative": jsii.Boolean(false),
//   				"lambda": &LambdaFunctionProperty{
//   					"arn": jsii.String("arn"),
//   				},
//   			},
//   			"requiredProperties": []*string{
//   				jsii.String("requiredProperties"),
//   			},
//   			"scope": jsii.String("scope"),
//   		},
//   	},
//   	isSingleton: jsii.Boolean(false),
//   	propertyDefinitions: map[string]interface{}{
//   		"propertyDefinitionsKey": &PropertyDefinitionProperty{
//   			"configurations": map[string]*string{
//   				"configurationsKey": jsii.String("configurations"),
//   			},
//   			"dataType": &dataTypeProperty{
//   				"type": jsii.String("type"),
//
//   				// the properties below are optional
//   				"allowedValues": []interface{}{
//   					&dataValueProperty{
//   						"booleanValue": jsii.Boolean(false),
//   						"doubleValue": jsii.Number(123),
//   						"expression": jsii.String("expression"),
//   						"integerValue": jsii.Number(123),
//   						"listValue": []interface{}{
//   							dataValueProperty_,
//   						},
//   						"longValue": jsii.Number(123),
//   						"mapValue": map[string]interface{}{
//   							"mapValueKey": dataValueProperty_,
//   						},
//   						"relationshipValue": relationshipValue,
//   						"stringValue": jsii.String("stringValue"),
//   					},
//   				},
//   				"nestedType": dataTypeProperty_,
//   				"relationship": &RelationshipProperty{
//   					"relationshipType": jsii.String("relationshipType"),
//   					"targetComponentTypeId": jsii.String("targetComponentTypeId"),
//   				},
//   				"unitOfMeasure": jsii.String("unitOfMeasure"),
//   			},
//   			"defaultValue": &dataValueProperty{
//   				"booleanValue": jsii.Boolean(false),
//   				"doubleValue": jsii.Number(123),
//   				"expression": jsii.String("expression"),
//   				"integerValue": jsii.Number(123),
//   				"listValue": []interface{}{
//   					dataValueProperty_,
//   				},
//   				"longValue": jsii.Number(123),
//   				"mapValue": map[string]interface{}{
//   					"mapValueKey": dataValueProperty_,
//   				},
//   				"relationshipValue": relationshipValue,
//   				"stringValue": jsii.String("stringValue"),
//   			},
//   			"isExternalId": jsii.Boolean(false),
//   			"isRequiredInEntity": jsii.Boolean(false),
//   			"isStoredExternally": jsii.Boolean(false),
//   			"isTimeSeries": jsii.Boolean(false),
//   		},
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   })
//
type CfnComponentType interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrCreationDateTime() *string
	AttrIsAbstract() awscdk.IResolvable
	AttrIsSchemaInitialized() awscdk.IResolvable
	AttrUpdateDateTime() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// `AWS::IoTTwinMaker::ComponentType.ComponentTypeId`.
	ComponentTypeId() *string
	SetComponentTypeId(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// `AWS::IoTTwinMaker::ComponentType.Description`.
	Description() *string
	SetDescription(val *string)
	// `AWS::IoTTwinMaker::ComponentType.ExtendsFrom`.
	ExtendsFrom() *[]*string
	SetExtendsFrom(val *[]*string)
	// `AWS::IoTTwinMaker::ComponentType.Functions`.
	Functions() interface{}
	SetFunctions(val interface{})
	// `AWS::IoTTwinMaker::ComponentType.IsSingleton`.
	IsSingleton() interface{}
	SetIsSingleton(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// `AWS::IoTTwinMaker::ComponentType.PropertyDefinitions`.
	PropertyDefinitions() interface{}
	SetPropertyDefinitions(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// `AWS::IoTTwinMaker::ComponentType.Tags`.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// `AWS::IoTTwinMaker::ComponentType.WorkspaceId`.
	WorkspaceId() *string
	SetWorkspaceId(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnComponentType
type jsiiProxy_CfnComponentType struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnComponentType) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) AttrCreationDateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationDateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) AttrIsAbstract() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrIsAbstract",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) AttrIsSchemaInitialized() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrIsSchemaInitialized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) AttrUpdateDateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUpdateDateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) ComponentTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"componentTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) ExtendsFrom() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"extendsFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) Functions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"functions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) IsSingleton() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSingleton",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) PropertyDefinitions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"propertyDefinitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) WorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoTTwinMaker::ComponentType`.
func NewCfnComponentType(scope awscdk.Construct, id *string, props *CfnComponentTypeProps) CfnComponentType {
	_init_.Initialize()

	j := jsiiProxy_CfnComponentType{}

	_jsii_.Create(
		"monocdk.aws_iottwinmaker.CfnComponentType",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTTwinMaker::ComponentType`.
func NewCfnComponentType_Override(c CfnComponentType, scope awscdk.Construct, id *string, props *CfnComponentTypeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iottwinmaker.CfnComponentType",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnComponentType) SetComponentTypeId(val *string) {
	_jsii_.Set(
		j,
		"componentTypeId",
		val,
	)
}

func (j *jsiiProxy_CfnComponentType) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnComponentType) SetExtendsFrom(val *[]*string) {
	_jsii_.Set(
		j,
		"extendsFrom",
		val,
	)
}

func (j *jsiiProxy_CfnComponentType) SetFunctions(val interface{}) {
	_jsii_.Set(
		j,
		"functions",
		val,
	)
}

func (j *jsiiProxy_CfnComponentType) SetIsSingleton(val interface{}) {
	_jsii_.Set(
		j,
		"isSingleton",
		val,
	)
}

func (j *jsiiProxy_CfnComponentType) SetPropertyDefinitions(val interface{}) {
	_jsii_.Set(
		j,
		"propertyDefinitions",
		val,
	)
}

func (j *jsiiProxy_CfnComponentType) SetWorkspaceId(val *string) {
	_jsii_.Set(
		j,
		"workspaceId",
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
func CfnComponentType_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iottwinmaker.CfnComponentType",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnComponentType_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iottwinmaker.CfnComponentType",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnComponentType_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iottwinmaker.CfnComponentType",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnComponentType_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_iottwinmaker.CfnComponentType",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnComponentType) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnComponentType) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnComponentType) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnComponentType) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnComponentType) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnComponentType) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnComponentType) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnComponentType) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponentType) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponentType) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnComponentType) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnComponentType) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnComponentType) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponentType) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnComponentType) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnComponentType) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponentType) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponentType) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnComponentType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponentType) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponentType) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataConnectorProperty := &dataConnectorProperty{
//   	isNative: jsii.Boolean(false),
//   	lambda: &lambdaFunctionProperty{
//   		arn: jsii.String("arn"),
//   	},
//   }
//
type CfnComponentType_DataConnectorProperty struct {
	// `CfnComponentType.DataConnectorProperty.IsNative`.
	IsNative interface{} `field:"optional" json:"isNative" yaml:"isNative"`
	// `CfnComponentType.DataConnectorProperty.Lambda`.
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataTypeProperty_ dataTypeProperty
//   var dataValueProperty_ dataValueProperty
//   var relationshipValue interface{}
//
//   dataTypeProperty := &dataTypeProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	allowedValues: []interface{}{
//   		&dataValueProperty{
//   			booleanValue: jsii.Boolean(false),
//   			doubleValue: jsii.Number(123),
//   			expression: jsii.String("expression"),
//   			integerValue: jsii.Number(123),
//   			listValue: []interface{}{
//   				dataValueProperty_,
//   			},
//   			longValue: jsii.Number(123),
//   			mapValue: map[string]interface{}{
//   				"mapValueKey": dataValueProperty_,
//   			},
//   			relationshipValue: relationshipValue,
//   			stringValue: jsii.String("stringValue"),
//   		},
//   	},
//   	nestedType: &dataTypeProperty{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		allowedValues: []interface{}{
//   			&dataValueProperty{
//   				booleanValue: jsii.Boolean(false),
//   				doubleValue: jsii.Number(123),
//   				expression: jsii.String("expression"),
//   				integerValue: jsii.Number(123),
//   				listValue: []interface{}{
//   					dataValueProperty_,
//   				},
//   				longValue: jsii.Number(123),
//   				mapValue: map[string]interface{}{
//   					"mapValueKey": dataValueProperty_,
//   				},
//   				relationshipValue: relationshipValue,
//   				stringValue: jsii.String("stringValue"),
//   			},
//   		},
//   		nestedType: dataTypeProperty_,
//   		relationship: &relationshipProperty{
//   			relationshipType: jsii.String("relationshipType"),
//   			targetComponentTypeId: jsii.String("targetComponentTypeId"),
//   		},
//   		unitOfMeasure: jsii.String("unitOfMeasure"),
//   	},
//   	relationship: &relationshipProperty{
//   		relationshipType: jsii.String("relationshipType"),
//   		targetComponentTypeId: jsii.String("targetComponentTypeId"),
//   	},
//   	unitOfMeasure: jsii.String("unitOfMeasure"),
//   }
//
type CfnComponentType_DataTypeProperty struct {
	// `CfnComponentType.DataTypeProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// `CfnComponentType.DataTypeProperty.AllowedValues`.
	AllowedValues interface{} `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// `CfnComponentType.DataTypeProperty.NestedType`.
	NestedType interface{} `field:"optional" json:"nestedType" yaml:"nestedType"`
	// `CfnComponentType.DataTypeProperty.Relationship`.
	Relationship interface{} `field:"optional" json:"relationship" yaml:"relationship"`
	// `CfnComponentType.DataTypeProperty.UnitOfMeasure`.
	UnitOfMeasure *string `field:"optional" json:"unitOfMeasure" yaml:"unitOfMeasure"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataValueProperty_ dataValueProperty
//   var relationshipValue interface{}
//
//   dataValueProperty := &dataValueProperty{
//   	booleanValue: jsii.Boolean(false),
//   	doubleValue: jsii.Number(123),
//   	expression: jsii.String("expression"),
//   	integerValue: jsii.Number(123),
//   	listValue: []interface{}{
//   		&dataValueProperty{
//   			booleanValue: jsii.Boolean(false),
//   			doubleValue: jsii.Number(123),
//   			expression: jsii.String("expression"),
//   			integerValue: jsii.Number(123),
//   			listValue: []interface{}{
//   				dataValueProperty_,
//   			},
//   			longValue: jsii.Number(123),
//   			mapValue: map[string]interface{}{
//   				"mapValueKey": dataValueProperty_,
//   			},
//   			relationshipValue: relationshipValue,
//   			stringValue: jsii.String("stringValue"),
//   		},
//   	},
//   	longValue: jsii.Number(123),
//   	mapValue: map[string]interface{}{
//   		"mapValueKey": &dataValueProperty{
//   			"booleanValue": jsii.Boolean(false),
//   			"doubleValue": jsii.Number(123),
//   			"expression": jsii.String("expression"),
//   			"integerValue": jsii.Number(123),
//   			"listValue": []interface{}{
//   				dataValueProperty_,
//   			},
//   			"longValue": jsii.Number(123),
//   			"mapValue": map[string]interface{}{
//   				"mapValueKey": dataValueProperty_,
//   			},
//   			"relationshipValue": relationshipValue,
//   			"stringValue": jsii.String("stringValue"),
//   		},
//   	},
//   	relationshipValue: relationshipValue,
//   	stringValue: jsii.String("stringValue"),
//   }
//
type CfnComponentType_DataValueProperty struct {
	// `CfnComponentType.DataValueProperty.BooleanValue`.
	BooleanValue interface{} `field:"optional" json:"booleanValue" yaml:"booleanValue"`
	// `CfnComponentType.DataValueProperty.DoubleValue`.
	DoubleValue *float64 `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// `CfnComponentType.DataValueProperty.Expression`.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// `CfnComponentType.DataValueProperty.IntegerValue`.
	IntegerValue *float64 `field:"optional" json:"integerValue" yaml:"integerValue"`
	// `CfnComponentType.DataValueProperty.ListValue`.
	ListValue interface{} `field:"optional" json:"listValue" yaml:"listValue"`
	// `CfnComponentType.DataValueProperty.LongValue`.
	LongValue *float64 `field:"optional" json:"longValue" yaml:"longValue"`
	// `CfnComponentType.DataValueProperty.MapValue`.
	MapValue interface{} `field:"optional" json:"mapValue" yaml:"mapValue"`
	// `CfnComponentType.DataValueProperty.RelationshipValue`.
	RelationshipValue interface{} `field:"optional" json:"relationshipValue" yaml:"relationshipValue"`
	// `CfnComponentType.DataValueProperty.StringValue`.
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   functionProperty := &functionProperty{
//   	implementedBy: &dataConnectorProperty{
//   		isNative: jsii.Boolean(false),
//   		lambda: &lambdaFunctionProperty{
//   			arn: jsii.String("arn"),
//   		},
//   	},
//   	requiredProperties: []*string{
//   		jsii.String("requiredProperties"),
//   	},
//   	scope: jsii.String("scope"),
//   }
//
type CfnComponentType_FunctionProperty struct {
	// `CfnComponentType.FunctionProperty.ImplementedBy`.
	ImplementedBy interface{} `field:"optional" json:"implementedBy" yaml:"implementedBy"`
	// `CfnComponentType.FunctionProperty.RequiredProperties`.
	RequiredProperties *[]*string `field:"optional" json:"requiredProperties" yaml:"requiredProperties"`
	// `CfnComponentType.FunctionProperty.Scope`.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaFunctionProperty := &lambdaFunctionProperty{
//   	arn: jsii.String("arn"),
//   }
//
type CfnComponentType_LambdaFunctionProperty struct {
	// `CfnComponentType.LambdaFunctionProperty.Arn`.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataTypeProperty_ dataTypeProperty
//   var dataValueProperty_ dataValueProperty
//   var relationshipValue interface{}
//
//   propertyDefinitionProperty := &propertyDefinitionProperty{
//   	configurations: map[string]*string{
//   		"configurationsKey": jsii.String("configurations"),
//   	},
//   	dataType: &dataTypeProperty{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		allowedValues: []interface{}{
//   			&dataValueProperty{
//   				booleanValue: jsii.Boolean(false),
//   				doubleValue: jsii.Number(123),
//   				expression: jsii.String("expression"),
//   				integerValue: jsii.Number(123),
//   				listValue: []interface{}{
//   					dataValueProperty_,
//   				},
//   				longValue: jsii.Number(123),
//   				mapValue: map[string]interface{}{
//   					"mapValueKey": dataValueProperty_,
//   				},
//   				relationshipValue: relationshipValue,
//   				stringValue: jsii.String("stringValue"),
//   			},
//   		},
//   		nestedType: dataTypeProperty_,
//   		relationship: &relationshipProperty{
//   			relationshipType: jsii.String("relationshipType"),
//   			targetComponentTypeId: jsii.String("targetComponentTypeId"),
//   		},
//   		unitOfMeasure: jsii.String("unitOfMeasure"),
//   	},
//   	defaultValue: &dataValueProperty{
//   		booleanValue: jsii.Boolean(false),
//   		doubleValue: jsii.Number(123),
//   		expression: jsii.String("expression"),
//   		integerValue: jsii.Number(123),
//   		listValue: []interface{}{
//   			dataValueProperty_,
//   		},
//   		longValue: jsii.Number(123),
//   		mapValue: map[string]interface{}{
//   			"mapValueKey": dataValueProperty_,
//   		},
//   		relationshipValue: relationshipValue,
//   		stringValue: jsii.String("stringValue"),
//   	},
//   	isExternalId: jsii.Boolean(false),
//   	isRequiredInEntity: jsii.Boolean(false),
//   	isStoredExternally: jsii.Boolean(false),
//   	isTimeSeries: jsii.Boolean(false),
//   }
//
type CfnComponentType_PropertyDefinitionProperty struct {
	// `CfnComponentType.PropertyDefinitionProperty.Configurations`.
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
	// `CfnComponentType.PropertyDefinitionProperty.DataType`.
	DataType interface{} `field:"optional" json:"dataType" yaml:"dataType"`
	// `CfnComponentType.PropertyDefinitionProperty.DefaultValue`.
	DefaultValue interface{} `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// `CfnComponentType.PropertyDefinitionProperty.IsExternalId`.
	IsExternalId interface{} `field:"optional" json:"isExternalId" yaml:"isExternalId"`
	// `CfnComponentType.PropertyDefinitionProperty.IsRequiredInEntity`.
	IsRequiredInEntity interface{} `field:"optional" json:"isRequiredInEntity" yaml:"isRequiredInEntity"`
	// `CfnComponentType.PropertyDefinitionProperty.IsStoredExternally`.
	IsStoredExternally interface{} `field:"optional" json:"isStoredExternally" yaml:"isStoredExternally"`
	// `CfnComponentType.PropertyDefinitionProperty.IsTimeSeries`.
	IsTimeSeries interface{} `field:"optional" json:"isTimeSeries" yaml:"isTimeSeries"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   relationshipProperty := &relationshipProperty{
//   	relationshipType: jsii.String("relationshipType"),
//   	targetComponentTypeId: jsii.String("targetComponentTypeId"),
//   }
//
type CfnComponentType_RelationshipProperty struct {
	// `CfnComponentType.RelationshipProperty.RelationshipType`.
	RelationshipType *string `field:"optional" json:"relationshipType" yaml:"relationshipType"`
	// `CfnComponentType.RelationshipProperty.TargetComponentTypeId`.
	TargetComponentTypeId *string `field:"optional" json:"targetComponentTypeId" yaml:"targetComponentTypeId"`
}

// Properties for defining a `CfnComponentType`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataTypeProperty_ dataTypeProperty
//   var dataValueProperty_ dataValueProperty
//   var relationshipValue interface{}
//
//   cfnComponentTypeProps := &cfnComponentTypeProps{
//   	componentTypeId: jsii.String("componentTypeId"),
//   	workspaceId: jsii.String("workspaceId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	extendsFrom: []*string{
//   		jsii.String("extendsFrom"),
//   	},
//   	functions: map[string]interface{}{
//   		"functionsKey": &FunctionProperty{
//   			"implementedBy": &DataConnectorProperty{
//   				"isNative": jsii.Boolean(false),
//   				"lambda": &LambdaFunctionProperty{
//   					"arn": jsii.String("arn"),
//   				},
//   			},
//   			"requiredProperties": []*string{
//   				jsii.String("requiredProperties"),
//   			},
//   			"scope": jsii.String("scope"),
//   		},
//   	},
//   	isSingleton: jsii.Boolean(false),
//   	propertyDefinitions: map[string]interface{}{
//   		"propertyDefinitionsKey": &PropertyDefinitionProperty{
//   			"configurations": map[string]*string{
//   				"configurationsKey": jsii.String("configurations"),
//   			},
//   			"dataType": &dataTypeProperty{
//   				"type": jsii.String("type"),
//
//   				// the properties below are optional
//   				"allowedValues": []interface{}{
//   					&dataValueProperty{
//   						"booleanValue": jsii.Boolean(false),
//   						"doubleValue": jsii.Number(123),
//   						"expression": jsii.String("expression"),
//   						"integerValue": jsii.Number(123),
//   						"listValue": []interface{}{
//   							dataValueProperty_,
//   						},
//   						"longValue": jsii.Number(123),
//   						"mapValue": map[string]interface{}{
//   							"mapValueKey": dataValueProperty_,
//   						},
//   						"relationshipValue": relationshipValue,
//   						"stringValue": jsii.String("stringValue"),
//   					},
//   				},
//   				"nestedType": dataTypeProperty_,
//   				"relationship": &RelationshipProperty{
//   					"relationshipType": jsii.String("relationshipType"),
//   					"targetComponentTypeId": jsii.String("targetComponentTypeId"),
//   				},
//   				"unitOfMeasure": jsii.String("unitOfMeasure"),
//   			},
//   			"defaultValue": &dataValueProperty{
//   				"booleanValue": jsii.Boolean(false),
//   				"doubleValue": jsii.Number(123),
//   				"expression": jsii.String("expression"),
//   				"integerValue": jsii.Number(123),
//   				"listValue": []interface{}{
//   					dataValueProperty_,
//   				},
//   				"longValue": jsii.Number(123),
//   				"mapValue": map[string]interface{}{
//   					"mapValueKey": dataValueProperty_,
//   				},
//   				"relationshipValue": relationshipValue,
//   				"stringValue": jsii.String("stringValue"),
//   			},
//   			"isExternalId": jsii.Boolean(false),
//   			"isRequiredInEntity": jsii.Boolean(false),
//   			"isStoredExternally": jsii.Boolean(false),
//   			"isTimeSeries": jsii.Boolean(false),
//   		},
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnComponentTypeProps struct {
	// `AWS::IoTTwinMaker::ComponentType.ComponentTypeId`.
	ComponentTypeId *string `field:"required" json:"componentTypeId" yaml:"componentTypeId"`
	// `AWS::IoTTwinMaker::ComponentType.WorkspaceId`.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// `AWS::IoTTwinMaker::ComponentType.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::IoTTwinMaker::ComponentType.ExtendsFrom`.
	ExtendsFrom *[]*string `field:"optional" json:"extendsFrom" yaml:"extendsFrom"`
	// `AWS::IoTTwinMaker::ComponentType.Functions`.
	Functions interface{} `field:"optional" json:"functions" yaml:"functions"`
	// `AWS::IoTTwinMaker::ComponentType.IsSingleton`.
	IsSingleton interface{} `field:"optional" json:"isSingleton" yaml:"isSingleton"`
	// `AWS::IoTTwinMaker::ComponentType.PropertyDefinitions`.
	PropertyDefinitions interface{} `field:"optional" json:"propertyDefinitions" yaml:"propertyDefinitions"`
	// `AWS::IoTTwinMaker::ComponentType.Tags`.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::IoTTwinMaker::Entity`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataValueProperty_ dataValueProperty
//   var definition interface{}
//   var error interface{}
//   var relationshipValue interface{}
//
//   cfnEntity := awscdk.Aws_iottwinmaker.NewCfnEntity(this, jsii.String("MyCfnEntity"), &cfnEntityProps{
//   	entityName: jsii.String("entityName"),
//   	workspaceId: jsii.String("workspaceId"),
//
//   	// the properties below are optional
//   	components: map[string]interface{}{
//   		"componentsKey": &ComponentProperty{
//   			"componentName": jsii.String("componentName"),
//   			"componentTypeId": jsii.String("componentTypeId"),
//   			"definedIn": jsii.String("definedIn"),
//   			"description": jsii.String("description"),
//   			"properties": map[string]interface{}{
//   				"propertiesKey": &PropertyProperty{
//   					"definition": definition,
//   					"value": &dataValueProperty{
//   						"booleanValue": jsii.Boolean(false),
//   						"doubleValue": jsii.Number(123),
//   						"expression": jsii.String("expression"),
//   						"integerValue": jsii.Number(123),
//   						"listValue": []interface{}{
//   							dataValueProperty_,
//   						},
//   						"longValue": jsii.Number(123),
//   						"mapValue": map[string]interface{}{
//   							"mapValueKey": dataValueProperty_,
//   						},
//   						"relationshipValue": relationshipValue,
//   						"stringValue": jsii.String("stringValue"),
//   					},
//   				},
//   			},
//   			"status": &StatusProperty{
//   				"error": error,
//   				"state": jsii.String("state"),
//   			},
//   		},
//   	},
//   	description: jsii.String("description"),
//   	entityId: jsii.String("entityId"),
//   	parentEntityId: jsii.String("parentEntityId"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   })
//
type CfnEntity interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrCreationDateTime() *string
	AttrHasChildEntities() awscdk.IResolvable
	AttrUpdateDateTime() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// `AWS::IoTTwinMaker::Entity.Components`.
	Components() interface{}
	SetComponents(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// `AWS::IoTTwinMaker::Entity.Description`.
	Description() *string
	SetDescription(val *string)
	// `AWS::IoTTwinMaker::Entity.EntityId`.
	EntityId() *string
	SetEntityId(val *string)
	// `AWS::IoTTwinMaker::Entity.EntityName`.
	EntityName() *string
	SetEntityName(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// `AWS::IoTTwinMaker::Entity.ParentEntityId`.
	ParentEntityId() *string
	SetParentEntityId(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// `AWS::IoTTwinMaker::Entity.Tags`.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// `AWS::IoTTwinMaker::Entity.WorkspaceId`.
	WorkspaceId() *string
	SetWorkspaceId(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnEntity
type jsiiProxy_CfnEntity struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnEntity) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEntity) AttrCreationDateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationDateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEntity) AttrHasChildEntities() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrHasChildEntities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEntity) AttrUpdateDateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUpdateDateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEntity) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEntity) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEntity) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEntity) Components() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEntity) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEntity) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEntity) EntityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEntity) EntityName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEntity) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEntity) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEntity) ParentEntityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentEntityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEntity) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEntity) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEntity) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEntity) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEntity) WorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoTTwinMaker::Entity`.
func NewCfnEntity(scope awscdk.Construct, id *string, props *CfnEntityProps) CfnEntity {
	_init_.Initialize()

	j := jsiiProxy_CfnEntity{}

	_jsii_.Create(
		"monocdk.aws_iottwinmaker.CfnEntity",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTTwinMaker::Entity`.
func NewCfnEntity_Override(c CfnEntity, scope awscdk.Construct, id *string, props *CfnEntityProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iottwinmaker.CfnEntity",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEntity) SetComponents(val interface{}) {
	_jsii_.Set(
		j,
		"components",
		val,
	)
}

func (j *jsiiProxy_CfnEntity) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnEntity) SetEntityId(val *string) {
	_jsii_.Set(
		j,
		"entityId",
		val,
	)
}

func (j *jsiiProxy_CfnEntity) SetEntityName(val *string) {
	_jsii_.Set(
		j,
		"entityName",
		val,
	)
}

func (j *jsiiProxy_CfnEntity) SetParentEntityId(val *string) {
	_jsii_.Set(
		j,
		"parentEntityId",
		val,
	)
}

func (j *jsiiProxy_CfnEntity) SetWorkspaceId(val *string) {
	_jsii_.Set(
		j,
		"workspaceId",
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
func CfnEntity_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iottwinmaker.CfnEntity",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnEntity_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iottwinmaker.CfnEntity",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnEntity_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iottwinmaker.CfnEntity",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEntity_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_iottwinmaker.CfnEntity",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEntity) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnEntity) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEntity) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnEntity) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnEntity) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnEntity) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnEntity) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnEntity) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEntity) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEntity) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnEntity) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnEntity) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnEntity) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEntity) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnEntity) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnEntity) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEntity) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEntity) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnEntity) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEntity) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEntity) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataValueProperty_ dataValueProperty
//   var definition interface{}
//   var error interface{}
//   var relationshipValue interface{}
//
//   componentProperty := &componentProperty{
//   	componentName: jsii.String("componentName"),
//   	componentTypeId: jsii.String("componentTypeId"),
//   	definedIn: jsii.String("definedIn"),
//   	description: jsii.String("description"),
//   	properties: map[string]interface{}{
//   		"propertiesKey": &PropertyProperty{
//   			"definition": definition,
//   			"value": &dataValueProperty{
//   				"booleanValue": jsii.Boolean(false),
//   				"doubleValue": jsii.Number(123),
//   				"expression": jsii.String("expression"),
//   				"integerValue": jsii.Number(123),
//   				"listValue": []interface{}{
//   					dataValueProperty_,
//   				},
//   				"longValue": jsii.Number(123),
//   				"mapValue": map[string]interface{}{
//   					"mapValueKey": dataValueProperty_,
//   				},
//   				"relationshipValue": relationshipValue,
//   				"stringValue": jsii.String("stringValue"),
//   			},
//   		},
//   	},
//   	status: &statusProperty{
//   		error: error,
//   		state: jsii.String("state"),
//   	},
//   }
//
type CfnEntity_ComponentProperty struct {
	// `CfnEntity.ComponentProperty.ComponentName`.
	ComponentName *string `field:"optional" json:"componentName" yaml:"componentName"`
	// `CfnEntity.ComponentProperty.ComponentTypeId`.
	ComponentTypeId *string `field:"optional" json:"componentTypeId" yaml:"componentTypeId"`
	// `CfnEntity.ComponentProperty.DefinedIn`.
	DefinedIn *string `field:"optional" json:"definedIn" yaml:"definedIn"`
	// `CfnEntity.ComponentProperty.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `CfnEntity.ComponentProperty.Properties`.
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
	// `CfnEntity.ComponentProperty.Status`.
	Status interface{} `field:"optional" json:"status" yaml:"status"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataValueProperty_ dataValueProperty
//   var relationshipValue interface{}
//
//   dataValueProperty := &dataValueProperty{
//   	booleanValue: jsii.Boolean(false),
//   	doubleValue: jsii.Number(123),
//   	expression: jsii.String("expression"),
//   	integerValue: jsii.Number(123),
//   	listValue: []interface{}{
//   		&dataValueProperty{
//   			booleanValue: jsii.Boolean(false),
//   			doubleValue: jsii.Number(123),
//   			expression: jsii.String("expression"),
//   			integerValue: jsii.Number(123),
//   			listValue: []interface{}{
//   				dataValueProperty_,
//   			},
//   			longValue: jsii.Number(123),
//   			mapValue: map[string]interface{}{
//   				"mapValueKey": dataValueProperty_,
//   			},
//   			relationshipValue: relationshipValue,
//   			stringValue: jsii.String("stringValue"),
//   		},
//   	},
//   	longValue: jsii.Number(123),
//   	mapValue: map[string]interface{}{
//   		"mapValueKey": &dataValueProperty{
//   			"booleanValue": jsii.Boolean(false),
//   			"doubleValue": jsii.Number(123),
//   			"expression": jsii.String("expression"),
//   			"integerValue": jsii.Number(123),
//   			"listValue": []interface{}{
//   				dataValueProperty_,
//   			},
//   			"longValue": jsii.Number(123),
//   			"mapValue": map[string]interface{}{
//   				"mapValueKey": dataValueProperty_,
//   			},
//   			"relationshipValue": relationshipValue,
//   			"stringValue": jsii.String("stringValue"),
//   		},
//   	},
//   	relationshipValue: relationshipValue,
//   	stringValue: jsii.String("stringValue"),
//   }
//
type CfnEntity_DataValueProperty struct {
	// `CfnEntity.DataValueProperty.BooleanValue`.
	BooleanValue interface{} `field:"optional" json:"booleanValue" yaml:"booleanValue"`
	// `CfnEntity.DataValueProperty.DoubleValue`.
	DoubleValue *float64 `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// `CfnEntity.DataValueProperty.Expression`.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// `CfnEntity.DataValueProperty.IntegerValue`.
	IntegerValue *float64 `field:"optional" json:"integerValue" yaml:"integerValue"`
	// `CfnEntity.DataValueProperty.ListValue`.
	ListValue interface{} `field:"optional" json:"listValue" yaml:"listValue"`
	// `CfnEntity.DataValueProperty.LongValue`.
	LongValue *float64 `field:"optional" json:"longValue" yaml:"longValue"`
	// `CfnEntity.DataValueProperty.MapValue`.
	MapValue interface{} `field:"optional" json:"mapValue" yaml:"mapValue"`
	// `CfnEntity.DataValueProperty.RelationshipValue`.
	RelationshipValue interface{} `field:"optional" json:"relationshipValue" yaml:"relationshipValue"`
	// `CfnEntity.DataValueProperty.StringValue`.
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataValueProperty_ dataValueProperty
//   var definition interface{}
//   var relationshipValue interface{}
//
//   propertyProperty := &propertyProperty{
//   	definition: definition,
//   	value: &dataValueProperty{
//   		booleanValue: jsii.Boolean(false),
//   		doubleValue: jsii.Number(123),
//   		expression: jsii.String("expression"),
//   		integerValue: jsii.Number(123),
//   		listValue: []interface{}{
//   			dataValueProperty_,
//   		},
//   		longValue: jsii.Number(123),
//   		mapValue: map[string]interface{}{
//   			"mapValueKey": dataValueProperty_,
//   		},
//   		relationshipValue: relationshipValue,
//   		stringValue: jsii.String("stringValue"),
//   	},
//   }
//
type CfnEntity_PropertyProperty struct {
	// `CfnEntity.PropertyProperty.Definition`.
	Definition interface{} `field:"optional" json:"definition" yaml:"definition"`
	// `CfnEntity.PropertyProperty.Value`.
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var error interface{}
//
//   statusProperty := &statusProperty{
//   	error: error,
//   	state: jsii.String("state"),
//   }
//
type CfnEntity_StatusProperty struct {
	// `CfnEntity.StatusProperty.Error`.
	Error interface{} `field:"optional" json:"error" yaml:"error"`
	// `CfnEntity.StatusProperty.State`.
	State *string `field:"optional" json:"state" yaml:"state"`
}

// Properties for defining a `CfnEntity`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataValueProperty_ dataValueProperty
//   var definition interface{}
//   var error interface{}
//   var relationshipValue interface{}
//
//   cfnEntityProps := &cfnEntityProps{
//   	entityName: jsii.String("entityName"),
//   	workspaceId: jsii.String("workspaceId"),
//
//   	// the properties below are optional
//   	components: map[string]interface{}{
//   		"componentsKey": &ComponentProperty{
//   			"componentName": jsii.String("componentName"),
//   			"componentTypeId": jsii.String("componentTypeId"),
//   			"definedIn": jsii.String("definedIn"),
//   			"description": jsii.String("description"),
//   			"properties": map[string]interface{}{
//   				"propertiesKey": &PropertyProperty{
//   					"definition": definition,
//   					"value": &dataValueProperty{
//   						"booleanValue": jsii.Boolean(false),
//   						"doubleValue": jsii.Number(123),
//   						"expression": jsii.String("expression"),
//   						"integerValue": jsii.Number(123),
//   						"listValue": []interface{}{
//   							dataValueProperty_,
//   						},
//   						"longValue": jsii.Number(123),
//   						"mapValue": map[string]interface{}{
//   							"mapValueKey": dataValueProperty_,
//   						},
//   						"relationshipValue": relationshipValue,
//   						"stringValue": jsii.String("stringValue"),
//   					},
//   				},
//   			},
//   			"status": &StatusProperty{
//   				"error": error,
//   				"state": jsii.String("state"),
//   			},
//   		},
//   	},
//   	description: jsii.String("description"),
//   	entityId: jsii.String("entityId"),
//   	parentEntityId: jsii.String("parentEntityId"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnEntityProps struct {
	// `AWS::IoTTwinMaker::Entity.EntityName`.
	EntityName *string `field:"required" json:"entityName" yaml:"entityName"`
	// `AWS::IoTTwinMaker::Entity.WorkspaceId`.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// `AWS::IoTTwinMaker::Entity.Components`.
	Components interface{} `field:"optional" json:"components" yaml:"components"`
	// `AWS::IoTTwinMaker::Entity.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::IoTTwinMaker::Entity.EntityId`.
	EntityId *string `field:"optional" json:"entityId" yaml:"entityId"`
	// `AWS::IoTTwinMaker::Entity.ParentEntityId`.
	ParentEntityId *string `field:"optional" json:"parentEntityId" yaml:"parentEntityId"`
	// `AWS::IoTTwinMaker::Entity.Tags`.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::IoTTwinMaker::Scene`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnScene := awscdk.Aws_iottwinmaker.NewCfnScene(this, jsii.String("MyCfnScene"), &cfnSceneProps{
//   	contentLocation: jsii.String("contentLocation"),
//   	sceneId: jsii.String("sceneId"),
//   	workspaceId: jsii.String("workspaceId"),
//
//   	// the properties below are optional
//   	capabilities: []*string{
//   		jsii.String("capabilities"),
//   	},
//   	description: jsii.String("description"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   })
//
type CfnScene interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrCreationDateTime() *string
	AttrUpdateDateTime() *string
	// `AWS::IoTTwinMaker::Scene.Capabilities`.
	Capabilities() *[]*string
	SetCapabilities(val *[]*string)
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// `AWS::IoTTwinMaker::Scene.ContentLocation`.
	ContentLocation() *string
	SetContentLocation(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// `AWS::IoTTwinMaker::Scene.Description`.
	Description() *string
	SetDescription(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// `AWS::IoTTwinMaker::Scene.SceneId`.
	SceneId() *string
	SetSceneId(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// `AWS::IoTTwinMaker::Scene.Tags`.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// `AWS::IoTTwinMaker::Scene.WorkspaceId`.
	WorkspaceId() *string
	SetWorkspaceId(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnScene
type jsiiProxy_CfnScene struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnScene) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScene) AttrCreationDateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationDateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScene) AttrUpdateDateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUpdateDateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScene) Capabilities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"capabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScene) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScene) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScene) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScene) ContentLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScene) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScene) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScene) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScene) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScene) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScene) SceneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sceneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScene) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScene) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScene) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScene) WorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoTTwinMaker::Scene`.
func NewCfnScene(scope awscdk.Construct, id *string, props *CfnSceneProps) CfnScene {
	_init_.Initialize()

	j := jsiiProxy_CfnScene{}

	_jsii_.Create(
		"monocdk.aws_iottwinmaker.CfnScene",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTTwinMaker::Scene`.
func NewCfnScene_Override(c CfnScene, scope awscdk.Construct, id *string, props *CfnSceneProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iottwinmaker.CfnScene",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnScene) SetCapabilities(val *[]*string) {
	_jsii_.Set(
		j,
		"capabilities",
		val,
	)
}

func (j *jsiiProxy_CfnScene) SetContentLocation(val *string) {
	_jsii_.Set(
		j,
		"contentLocation",
		val,
	)
}

func (j *jsiiProxy_CfnScene) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnScene) SetSceneId(val *string) {
	_jsii_.Set(
		j,
		"sceneId",
		val,
	)
}

func (j *jsiiProxy_CfnScene) SetWorkspaceId(val *string) {
	_jsii_.Set(
		j,
		"workspaceId",
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
func CfnScene_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iottwinmaker.CfnScene",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnScene_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iottwinmaker.CfnScene",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnScene_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iottwinmaker.CfnScene",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnScene_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_iottwinmaker.CfnScene",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnScene) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnScene) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnScene) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnScene) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnScene) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnScene) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnScene) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnScene) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScene) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScene) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnScene) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnScene) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnScene) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScene) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnScene) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnScene) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScene) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScene) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnScene) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScene) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScene) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnScene`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSceneProps := &cfnSceneProps{
//   	contentLocation: jsii.String("contentLocation"),
//   	sceneId: jsii.String("sceneId"),
//   	workspaceId: jsii.String("workspaceId"),
//
//   	// the properties below are optional
//   	capabilities: []*string{
//   		jsii.String("capabilities"),
//   	},
//   	description: jsii.String("description"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnSceneProps struct {
	// `AWS::IoTTwinMaker::Scene.ContentLocation`.
	ContentLocation *string `field:"required" json:"contentLocation" yaml:"contentLocation"`
	// `AWS::IoTTwinMaker::Scene.SceneId`.
	SceneId *string `field:"required" json:"sceneId" yaml:"sceneId"`
	// `AWS::IoTTwinMaker::Scene.WorkspaceId`.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// `AWS::IoTTwinMaker::Scene.Capabilities`.
	Capabilities *[]*string `field:"optional" json:"capabilities" yaml:"capabilities"`
	// `AWS::IoTTwinMaker::Scene.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::IoTTwinMaker::Scene.Tags`.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::IoTTwinMaker::Workspace`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWorkspace := awscdk.Aws_iottwinmaker.NewCfnWorkspace(this, jsii.String("MyCfnWorkspace"), &cfnWorkspaceProps{
//   	role: jsii.String("role"),
//   	s3Location: jsii.String("s3Location"),
//   	workspaceId: jsii.String("workspaceId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   })
//
type CfnWorkspace interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrCreationDateTime() *string
	AttrUpdateDateTime() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// `AWS::IoTTwinMaker::Workspace.Description`.
	Description() *string
	SetDescription(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// `AWS::IoTTwinMaker::Workspace.Role`.
	Role() *string
	SetRole(val *string)
	// `AWS::IoTTwinMaker::Workspace.S3Location`.
	S3Location() *string
	SetS3Location(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// `AWS::IoTTwinMaker::Workspace.Tags`.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// `AWS::IoTTwinMaker::Workspace.WorkspaceId`.
	WorkspaceId() *string
	SetWorkspaceId(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnWorkspace
type jsiiProxy_CfnWorkspace struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnWorkspace) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) AttrCreationDateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationDateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) AttrUpdateDateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUpdateDateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) S3Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) WorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoTTwinMaker::Workspace`.
func NewCfnWorkspace(scope awscdk.Construct, id *string, props *CfnWorkspaceProps) CfnWorkspace {
	_init_.Initialize()

	j := jsiiProxy_CfnWorkspace{}

	_jsii_.Create(
		"monocdk.aws_iottwinmaker.CfnWorkspace",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTTwinMaker::Workspace`.
func NewCfnWorkspace_Override(c CfnWorkspace, scope awscdk.Construct, id *string, props *CfnWorkspaceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iottwinmaker.CfnWorkspace",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnWorkspace) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnWorkspace) SetRole(val *string) {
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_CfnWorkspace) SetS3Location(val *string) {
	_jsii_.Set(
		j,
		"s3Location",
		val,
	)
}

func (j *jsiiProxy_CfnWorkspace) SetWorkspaceId(val *string) {
	_jsii_.Set(
		j,
		"workspaceId",
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
func CfnWorkspace_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iottwinmaker.CfnWorkspace",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnWorkspace_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iottwinmaker.CfnWorkspace",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnWorkspace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iottwinmaker.CfnWorkspace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWorkspace_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_iottwinmaker.CfnWorkspace",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWorkspace) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnWorkspace) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnWorkspace) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnWorkspace) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnWorkspace) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnWorkspace) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnWorkspace) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnWorkspace) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkspace) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkspace) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnWorkspace) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnWorkspace) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnWorkspace) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkspace) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnWorkspace) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnWorkspace) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkspace) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkspace) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnWorkspace) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkspace) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkspace) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnWorkspace`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWorkspaceProps := &cfnWorkspaceProps{
//   	role: jsii.String("role"),
//   	s3Location: jsii.String("s3Location"),
//   	workspaceId: jsii.String("workspaceId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnWorkspaceProps struct {
	// `AWS::IoTTwinMaker::Workspace.Role`.
	Role *string `field:"required" json:"role" yaml:"role"`
	// `AWS::IoTTwinMaker::Workspace.S3Location`.
	S3Location *string `field:"required" json:"s3Location" yaml:"s3Location"`
	// `AWS::IoTTwinMaker::Workspace.WorkspaceId`.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// `AWS::IoTTwinMaker::Workspace.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::IoTTwinMaker::Workspace.Tags`.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

