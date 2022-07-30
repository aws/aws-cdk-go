package awsiottwinmaker

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiottwinmaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::IoTTwinMaker::ComponentType`.
//
// Use the `AWS::IoTTwinMaker::ComponentType` resource to declare a component type.
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
	// The ARN of the component type.
	AttrArn() *string
	// The date and time when the component type was created.
	AttrCreationDateTime() *string
	// A boolean value that specifies whether the component type is abstract.
	AttrIsAbstract() awscdk.IResolvable
	// A boolean value that specifies whether the component type has a schema initializer and that the schema initializer has run.
	AttrIsSchemaInitialized() awscdk.IResolvable
	// The component type the update time.
	AttrUpdateDateTime() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The ID of the component type.
	ComponentTypeId() *string
	SetComponentTypeId(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of the component type.
	Description() *string
	SetDescription(val *string)
	// The name of the parent component type that this component type extends.
	ExtendsFrom() *[]*string
	SetExtendsFrom(val *[]*string)
	// An object that maps strings to the functions in the component type.
	//
	// Each string in the mapping must be unique to this object.
	//
	// For information on the FunctionResponse object see the [FunctionResponse](https://docs.aws.amazon.com//iot-twinmaker/latest/apireference/API_FunctionResponse.html) API reference.
	Functions() interface{}
	SetFunctions(val interface{})
	// A boolean value that specifies whether an entity can have more than one component of this type.
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
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// An object that maps strings to the property definitions in the component type.
	//
	// Each string in the mapping must be unique to this object.
	//
	// For information about the PropertyDefinitionResponse object, see the [PropertyDefinitionResponse](https://docs.aws.amazon.com//iot-twinmaker/latest/apireference/API_PropertyDefinitionResponse.html) API reference.
	PropertyDefinitions() interface{}
	SetPropertyDefinitions(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The ComponentType tags.
	Tags() awscdk.TagManager
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// The ID of the workspace.
	WorkspaceId() *string
	SetWorkspaceId(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
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

func (j *jsiiProxy_CfnComponentType) Node() constructs.Node {
	var returns constructs.Node
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

func (j *jsiiProxy_CfnComponentType) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
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
func NewCfnComponentType(scope constructs.Construct, id *string, props *CfnComponentTypeProps) CfnComponentType {
	_init_.Initialize()

	j := jsiiProxy_CfnComponentType{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iottwinmaker.CfnComponentType",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTTwinMaker::ComponentType`.
func NewCfnComponentType_Override(c CfnComponentType, scope constructs.Construct, id *string, props *CfnComponentTypeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iottwinmaker.CfnComponentType",
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
func CfnComponentType_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iottwinmaker.CfnComponentType",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnComponentType_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iottwinmaker.CfnComponentType",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnComponentType_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iottwinmaker.CfnComponentType",
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
		"aws-cdk-lib.aws_iottwinmaker.CfnComponentType",
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

func (c *jsiiProxy_CfnComponentType) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

func (c *jsiiProxy_CfnComponentType) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The data connector.
//
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
	// A boolean value that specifies whether the data connector is native to IoT TwinMaker.
	IsNative interface{} `field:"optional" json:"isNative" yaml:"isNative"`
	// The Lambda function associated with the data connector.
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
}

// An object that specifies the data type of a property.
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
	// The underlying type of the data type.
	//
	// Valid Values: `RELATIONSHIP | STRING | LONG | BOOLEAN | INTEGER | DOUBLE | LIST | MAP`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The allowed values for this data type.
	AllowedValues interface{} `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// The nested type in the data type.
	NestedType interface{} `field:"optional" json:"nestedType" yaml:"nestedType"`
	// A relationship that associates a component with another component.
	Relationship interface{} `field:"optional" json:"relationship" yaml:"relationship"`
	// The unit of measure used in this data type.
	UnitOfMeasure *string `field:"optional" json:"unitOfMeasure" yaml:"unitOfMeasure"`
}

// An object that specifies a value for a property.
//
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
	// A boolean value.
	BooleanValue interface{} `field:"optional" json:"booleanValue" yaml:"booleanValue"`
	// A double value.
	DoubleValue *float64 `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// An expression that produces the value.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// An integer value.
	IntegerValue *float64 `field:"optional" json:"integerValue" yaml:"integerValue"`
	// A list of multiple values.
	ListValue interface{} `field:"optional" json:"listValue" yaml:"listValue"`
	// A long value.
	LongValue *float64 `field:"optional" json:"longValue" yaml:"longValue"`
	// An object that maps strings to multiple `DataValue` objects.
	MapValue interface{} `field:"optional" json:"mapValue" yaml:"mapValue"`
	// A value that relates a component to another component.
	RelationshipValue interface{} `field:"optional" json:"relationshipValue" yaml:"relationshipValue"`
	// A string value.
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

// The function body.
//
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
	// The data connector.
	ImplementedBy interface{} `field:"optional" json:"implementedBy" yaml:"implementedBy"`
	// The required properties of the function.
	RequiredProperties *[]*string `field:"optional" json:"requiredProperties" yaml:"requiredProperties"`
	// The scope of the function.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

// The Lambda function.
//
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
	// The Lambda function ARN.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
}

// PropertyDefinition is an object that maps strings to the property definitions in the component type.
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
	// A mapping that specifies configuration information about the property.
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
	// `CfnComponentType.PropertyDefinitionProperty.DataType`.
	DataType interface{} `field:"optional" json:"dataType" yaml:"dataType"`
	// A boolean value that specifies whether the property ID comes from an external data store.
	DefaultValue interface{} `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// A boolean value that specifies whether the property ID comes from an external data store.
	IsExternalId interface{} `field:"optional" json:"isExternalId" yaml:"isExternalId"`
	// A boolean value that specifies whether the property is required in an entity.
	IsRequiredInEntity interface{} `field:"optional" json:"isRequiredInEntity" yaml:"isRequiredInEntity"`
	// A boolean value that specifies whether the property is stored externally.
	IsStoredExternally interface{} `field:"optional" json:"isStoredExternally" yaml:"isStoredExternally"`
	// A boolean value that specifies whether the property consists of time series data.
	IsTimeSeries interface{} `field:"optional" json:"isTimeSeries" yaml:"isTimeSeries"`
}

// An object that specifies a relationship with another component type.
//
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
	// The type of the relationship.
	RelationshipType *string `field:"optional" json:"relationshipType" yaml:"relationshipType"`
	// The ID of the target component type associated with this relationship.
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
	// The ID of the component type.
	ComponentTypeId *string `field:"required" json:"componentTypeId" yaml:"componentTypeId"`
	// The ID of the workspace.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// The description of the component type.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the parent component type that this component type extends.
	ExtendsFrom *[]*string `field:"optional" json:"extendsFrom" yaml:"extendsFrom"`
	// An object that maps strings to the functions in the component type.
	//
	// Each string in the mapping must be unique to this object.
	//
	// For information on the FunctionResponse object see the [FunctionResponse](https://docs.aws.amazon.com//iot-twinmaker/latest/apireference/API_FunctionResponse.html) API reference.
	Functions interface{} `field:"optional" json:"functions" yaml:"functions"`
	// A boolean value that specifies whether an entity can have more than one component of this type.
	IsSingleton interface{} `field:"optional" json:"isSingleton" yaml:"isSingleton"`
	// An object that maps strings to the property definitions in the component type.
	//
	// Each string in the mapping must be unique to this object.
	//
	// For information about the PropertyDefinitionResponse object, see the [PropertyDefinitionResponse](https://docs.aws.amazon.com//iot-twinmaker/latest/apireference/API_PropertyDefinitionResponse.html) API reference.
	PropertyDefinitions interface{} `field:"optional" json:"propertyDefinitions" yaml:"propertyDefinitions"`
	// The ComponentType tags.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::IoTTwinMaker::Entity`.
//
// Use the `AWS::IoTTwinMaker::Entity` resource to declare an entity.
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
	// The entity ARN.
	AttrArn() *string
	// The date and time the entity was created.
	AttrCreationDateTime() *string
	// A boolean value that specifies whether the entity has child entities or not.
	AttrHasChildEntities() awscdk.IResolvable
	// The date and time when the component type was last updated.
	AttrUpdateDateTime() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// An object that maps strings to the components in the entity.
	//
	// Each string in the mapping must be unique to this object.
	//
	// For information on the component object see the [component](https://docs.aws.amazon.com//iot-twinmaker/latest/apireference/API_ComponentResponse.html) API reference.
	Components() interface{}
	SetComponents(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of the entity.
	Description() *string
	SetDescription(val *string)
	// The entity ID.
	EntityId() *string
	SetEntityId(val *string)
	// The entity name.
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
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// The ID of the parent entity.
	ParentEntityId() *string
	SetParentEntityId(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Metadata that you can use to manage the entity.
	Tags() awscdk.TagManager
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// The ID of the workspace.
	WorkspaceId() *string
	SetWorkspaceId(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
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

func (j *jsiiProxy_CfnEntity) Node() constructs.Node {
	var returns constructs.Node
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

func (j *jsiiProxy_CfnEntity) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
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
func NewCfnEntity(scope constructs.Construct, id *string, props *CfnEntityProps) CfnEntity {
	_init_.Initialize()

	j := jsiiProxy_CfnEntity{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iottwinmaker.CfnEntity",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTTwinMaker::Entity`.
func NewCfnEntity_Override(c CfnEntity, scope constructs.Construct, id *string, props *CfnEntityProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iottwinmaker.CfnEntity",
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
func CfnEntity_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iottwinmaker.CfnEntity",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnEntity_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iottwinmaker.CfnEntity",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnEntity_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iottwinmaker.CfnEntity",
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
		"aws-cdk-lib.aws_iottwinmaker.CfnEntity",
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

func (c *jsiiProxy_CfnEntity) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

func (c *jsiiProxy_CfnEntity) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The entity componenet.
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
	// The name of the component.
	ComponentName *string `field:"optional" json:"componentName" yaml:"componentName"`
	// The ID of the ComponentType.
	ComponentTypeId *string `field:"optional" json:"componentTypeId" yaml:"componentTypeId"`
	// The name of the property definition set in the request.
	DefinedIn *string `field:"optional" json:"definedIn" yaml:"definedIn"`
	// The description of the component.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An object that maps strings to the properties to set in the component type.
	//
	// Each string in the mapping must be unique to this object.
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
	// The status of the component.
	Status interface{} `field:"optional" json:"status" yaml:"status"`
}

// An object that specifies a value for a property.
//
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
	// A boolean value.
	BooleanValue interface{} `field:"optional" json:"booleanValue" yaml:"booleanValue"`
	// A double value.
	DoubleValue *float64 `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// An expression that produces the value.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// An integer value.
	IntegerValue *float64 `field:"optional" json:"integerValue" yaml:"integerValue"`
	// A list of multiple values.
	ListValue interface{} `field:"optional" json:"listValue" yaml:"listValue"`
	// A long value.
	LongValue *float64 `field:"optional" json:"longValue" yaml:"longValue"`
	// An object that maps strings to multiple DataValue objects.
	MapValue interface{} `field:"optional" json:"mapValue" yaml:"mapValue"`
	// A value that relates a component to another component.
	RelationshipValue interface{} `field:"optional" json:"relationshipValue" yaml:"relationshipValue"`
	// A string value.
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

// An object that sets information about a property.
//
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
	// An object that specifies information about a property.
	Definition interface{} `field:"optional" json:"definition" yaml:"definition"`
	// An object that contains information about a value for a time series property.
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

// The current status of the entity.
//
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
	// The error message.
	Error interface{} `field:"optional" json:"error" yaml:"error"`
	// The current state of the entity, component, component type, or workspace.
	//
	// Valid Values: `CREATING | UPDATING | DELETING | ACTIVE | ERROR`.
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
	// The entity name.
	EntityName *string `field:"required" json:"entityName" yaml:"entityName"`
	// The ID of the workspace.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// An object that maps strings to the components in the entity.
	//
	// Each string in the mapping must be unique to this object.
	//
	// For information on the component object see the [component](https://docs.aws.amazon.com//iot-twinmaker/latest/apireference/API_ComponentResponse.html) API reference.
	Components interface{} `field:"optional" json:"components" yaml:"components"`
	// The description of the entity.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The entity ID.
	EntityId *string `field:"optional" json:"entityId" yaml:"entityId"`
	// The ID of the parent entity.
	ParentEntityId *string `field:"optional" json:"parentEntityId" yaml:"parentEntityId"`
	// Metadata that you can use to manage the entity.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::IoTTwinMaker::Scene`.
//
// Use the `AWS::IoTTwinMaker::Scene` resource to declare a scene.
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
	// The scene ARN.
	AttrArn() *string
	// The date and time when the scene was created.
	AttrCreationDateTime() *string
	// The scene the update time.
	AttrUpdateDateTime() *string
	// A list of capabilities that the scene uses to render.
	Capabilities() *[]*string
	SetCapabilities(val *[]*string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The relative path that specifies the location of the content definition file.
	ContentLocation() *string
	SetContentLocation(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of this scene.
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
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The scene ID.
	SceneId() *string
	SetSceneId(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The ComponentType tags.
	Tags() awscdk.TagManager
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// The ID of the workspace.
	WorkspaceId() *string
	SetWorkspaceId(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
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

func (j *jsiiProxy_CfnScene) Node() constructs.Node {
	var returns constructs.Node
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

func (j *jsiiProxy_CfnScene) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
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
func NewCfnScene(scope constructs.Construct, id *string, props *CfnSceneProps) CfnScene {
	_init_.Initialize()

	j := jsiiProxy_CfnScene{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iottwinmaker.CfnScene",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTTwinMaker::Scene`.
func NewCfnScene_Override(c CfnScene, scope constructs.Construct, id *string, props *CfnSceneProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iottwinmaker.CfnScene",
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
func CfnScene_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iottwinmaker.CfnScene",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnScene_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iottwinmaker.CfnScene",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnScene_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iottwinmaker.CfnScene",
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
		"aws-cdk-lib.aws_iottwinmaker.CfnScene",
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

func (c *jsiiProxy_CfnScene) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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
	// The relative path that specifies the location of the content definition file.
	ContentLocation *string `field:"required" json:"contentLocation" yaml:"contentLocation"`
	// The scene ID.
	SceneId *string `field:"required" json:"sceneId" yaml:"sceneId"`
	// The ID of the workspace.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// A list of capabilities that the scene uses to render.
	Capabilities *[]*string `field:"optional" json:"capabilities" yaml:"capabilities"`
	// The description of this scene.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ComponentType tags.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::IoTTwinMaker::Workspace`.
//
// Use the `AWS::IoTTwinMaker::Workspace` resource to declare a workspace.
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
	// The workspace ARN.
	AttrArn() *string
	// The date and time the workspace was created.
	AttrCreationDateTime() *string
	// The date and time the workspace was updated.
	AttrUpdateDateTime() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of the workspace.
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
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The ARN of the execution role associated with the workspace.
	Role() *string
	SetRole(val *string)
	// The ARN of the S3 bucket where resources associated with the workspace are stored.
	S3Location() *string
	SetS3Location(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Metadata that you can use to manage the workspace.
	Tags() awscdk.TagManager
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// The ID of the workspace.
	WorkspaceId() *string
	SetWorkspaceId(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
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

func (j *jsiiProxy_CfnWorkspace) Node() constructs.Node {
	var returns constructs.Node
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

func (j *jsiiProxy_CfnWorkspace) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
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
func NewCfnWorkspace(scope constructs.Construct, id *string, props *CfnWorkspaceProps) CfnWorkspace {
	_init_.Initialize()

	j := jsiiProxy_CfnWorkspace{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iottwinmaker.CfnWorkspace",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTTwinMaker::Workspace`.
func NewCfnWorkspace_Override(c CfnWorkspace, scope constructs.Construct, id *string, props *CfnWorkspaceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iottwinmaker.CfnWorkspace",
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
func CfnWorkspace_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iottwinmaker.CfnWorkspace",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnWorkspace_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iottwinmaker.CfnWorkspace",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnWorkspace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iottwinmaker.CfnWorkspace",
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
		"aws-cdk-lib.aws_iottwinmaker.CfnWorkspace",
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

func (c *jsiiProxy_CfnWorkspace) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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
	// The ARN of the execution role associated with the workspace.
	Role *string `field:"required" json:"role" yaml:"role"`
	// The ARN of the S3 bucket where resources associated with the workspace are stored.
	S3Location *string `field:"required" json:"s3Location" yaml:"s3Location"`
	// The ID of the workspace.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// The description of the workspace.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Metadata that you can use to manage the workspace.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

